package agilitycms

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	cfg := NewConfiguration()

	assert.NotNil(t, cfg)
	assert.NotNil(t, cfg.DefaultHeader)
	assert.Equal(t, "OpenAPI-Generator/1.0.0/go", cfg.UserAgent)
	assert.False(t, cfg.Debug)
	assert.Len(t, cfg.Servers, 1)
	assert.Equal(t, "", cfg.Servers[0].URL)
	assert.Equal(t, "No description provided", cfg.Servers[0].Description)
	assert.NotNil(t, cfg.OperationServers)
}

func TestConfiguration_AddDefaultHeader(t *testing.T) {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("X-API-Key", "test-key")

	assert.Equal(t, "test-key", cfg.DefaultHeader["X-API-Key"])
}

func TestServerConfigurations_URL(t *testing.T) {
	tests := []struct {
		name      string
		servers   ServerConfigurations
		index     int
		variables map[string]string
		expected  string
		wantErr   bool
	}{
		{
			name: "Valid index no variables",
			servers: ServerConfigurations{
				{URL: "https://api.example.com", Description: "Test server"},
			},
			index:    0,
			expected: "https://api.example.com",
			wantErr:  false,
		},
		{
			name: "Index out of range",
			servers: ServerConfigurations{
				{URL: "https://api.example.com", Description: "Test server"},
			},
			index:   1,
			wantErr: true,
		},
		{
			name: "Negative index",
			servers: ServerConfigurations{
				{URL: "https://api.example.com", Description: "Test server"},
			},
			index:   -1,
			wantErr: true,
		},
		{
			name: "With variables",
			servers: ServerConfigurations{
				{
					URL:         "https://{host}/api/{version}",
					Description: "Test server with variables",
					Variables: map[string]ServerVariable{
						"host": {
							Description:  "Host",
							DefaultValue: "api.example.com",
							EnumValues:   []string{"api.example.com", "staging.example.com"},
						},
						"version": {
							Description:  "API Version",
							DefaultValue: "v1",
						},
					},
				},
			},
			index: 0,
			variables: map[string]string{
				"host":    "staging.example.com",
				"version": "v2",
			},
			expected: "https://staging.example.com/api/v2",
			wantErr:  false,
		},
		{
			name: "With default variables",
			servers: ServerConfigurations{
				{
					URL:         "https://{host}/api/{version}",
					Description: "Test server with variables",
					Variables: map[string]ServerVariable{
						"host": {
							Description:  "Host",
							DefaultValue: "api.example.com",
						},
						"version": {
							Description:  "API Version",
							DefaultValue: "v1",
						},
					},
				},
			},
			index:    0,
			expected: "https://api.example.com/api/v1",
			wantErr:  false,
		},
		{
			name: "Invalid enum value",
			servers: ServerConfigurations{
				{
					URL:         "https://{host}/api",
					Description: "Test server with enum",
					Variables: map[string]ServerVariable{
						"host": {
							Description:  "Host",
							DefaultValue: "api.example.com",
							EnumValues:   []string{"api.example.com", "staging.example.com"},
						},
					},
				},
			},
			index: 0,
			variables: map[string]string{
				"host": "invalid.example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.servers.URL(tt.index, tt.variables)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestConfiguration_ServerURL(t *testing.T) {
	cfg := &Configuration{
		Servers: ServerConfigurations{
			{URL: "https://api.example.com", Description: "Test server"},
		},
	}

	result, err := cfg.ServerURL(0, nil)
	assert.NoError(t, err)
	assert.Equal(t, "https://api.example.com", result)
}

func TestGetServerIndex(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		expected int
		wantErr  bool
	}{
		{
			name:     "Empty context",
			ctx:      context.Background(),
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Context without server index",
			ctx:      context.Background(),
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Context with valid server index",
			ctx:      context.WithValue(context.Background(), ContextServerIndex, 2),
			expected: 2,
			wantErr:  false,
		},
		{
			name:    "Context with invalid server index type",
			ctx:     context.WithValue(context.Background(), ContextServerIndex, "invalid"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getServerIndex(tt.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetServerOperationIndex(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		endpoint string
		expected int
		wantErr  bool
	}{
		{
			name:     "Empty context",
			ctx:      context.Background(),
			endpoint: "test",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Context without operation indices",
			ctx:      context.Background(),
			endpoint: "test",
			expected: 0,
			wantErr:  false,
		},
		{
			name: "Context with valid operation index",
			ctx: context.WithValue(context.Background(), ContextOperationServerIndices, map[string]int{
				"test": 1,
			}),
			endpoint: "test",
			expected: 1,
			wantErr:  false,
		},
		{
			name: "Context with operation indices but endpoint not found",
			ctx: context.WithValue(context.Background(), ContextOperationServerIndices, map[string]int{
				"other": 1,
			}),
			endpoint: "test",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Context with invalid operation indices type",
			ctx:      context.WithValue(context.Background(), ContextOperationServerIndices, "invalid"),
			endpoint: "test",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getServerOperationIndex(tt.ctx, tt.endpoint)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetServerVariables(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		expected map[string]string
		wantErr  bool
	}{
		{
			name:     "Empty context",
			ctx:      context.Background(),
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Context without variables",
			ctx:      context.Background(),
			expected: nil,
			wantErr:  false,
		},
		{
			name: "Context with valid variables",
			ctx: context.WithValue(context.Background(), ContextServerVariables, map[string]string{
				"host": "api.example.com",
			}),
			expected: map[string]string{"host": "api.example.com"},
			wantErr:  false,
		},
		{
			name:    "Context with invalid variables type",
			ctx:     context.WithValue(context.Background(), ContextServerVariables, "invalid"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getServerVariables(tt.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetServerOperationVariables(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		endpoint string
		expected map[string]string
		wantErr  bool
	}{
		{
			name:     "Empty context",
			ctx:      context.Background(),
			endpoint: "test",
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Context without operation variables",
			ctx:      context.Background(),
			endpoint: "test",
			expected: nil,
			wantErr:  false,
		},
		{
			name: "Context with valid operation variables",
			ctx: context.WithValue(context.Background(), ContextOperationServerVariables, map[string]map[string]string{
				"test": {"host": "api.example.com"},
			}),
			endpoint: "test",
			expected: map[string]string{"host": "api.example.com"},
			wantErr:  false,
		},
		{
			name: "Context with operation variables but endpoint not found",
			ctx: context.WithValue(context.Background(), ContextOperationServerVariables, map[string]map[string]string{
				"other": {"host": "api.example.com"},
			}),
			endpoint: "test",
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Context with invalid operation variables type",
			ctx:      context.WithValue(context.Background(), ContextOperationServerVariables, "invalid"),
			endpoint: "test",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getServerOperationVariables(tt.ctx, tt.endpoint)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestConfiguration_ServerURLWithContext(t *testing.T) {
	cfg := &Configuration{
		Servers: ServerConfigurations{
			{
				URL:         "https://{host}/api",
				Description: "Default server",
				Variables: map[string]ServerVariable{
					"host": {DefaultValue: "api.example.com"},
				},
			},
		},
		OperationServers: map[string]ServerConfigurations{
			"test-endpoint": {
				{
					URL:         "https://{host}/test",
					Description: "Test endpoint server",
					Variables: map[string]ServerVariable{
						"host": {DefaultValue: "test.example.com"},
					},
				},
			},
		},
	}

	t.Run("Nil context uses default", func(t *testing.T) {
		result, err := cfg.ServerURLWithContext(nil, "unknown-endpoint")
		require.NoError(t, err)
		assert.Equal(t, "https://api.example.com/api", result)
	})

	t.Run("Context with operation server", func(t *testing.T) {
		ctx := context.Background()
		result, err := cfg.ServerURLWithContext(ctx, "test-endpoint")
		require.NoError(t, err)
		assert.Equal(t, "https://test.example.com/test", result)
	})

	t.Run("Context with server variables", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), ContextOperationServerVariables, map[string]map[string]string{
			"test-endpoint": {"host": "custom.example.com"},
		})
		result, err := cfg.ServerURLWithContext(ctx, "test-endpoint")
		require.NoError(t, err)
		assert.Equal(t, "https://custom.example.com/test", result)
	})
}

func TestContextKey_String(t *testing.T) {
	key := contextKey("test")
	assert.Equal(t, "auth test", key.String())
}

func TestBasicAuth(t *testing.T) {
	auth := BasicAuth{
		UserName: "user",
		Password: "pass",
	}

	assert.Equal(t, "user", auth.UserName)
	assert.Equal(t, "pass", auth.Password)
}

func TestAPIKey(t *testing.T) {
	key := APIKey{
		Key:    "secret-key",
		Prefix: "Bearer",
	}

	assert.Equal(t, "secret-key", key.Key)
	assert.Equal(t, "Bearer", key.Prefix)
}

func TestServerVariable(t *testing.T) {
	sv := ServerVariable{
		Description:  "Test variable",
		DefaultValue: "default",
		EnumValues:   []string{"default", "alternative"},
	}

	assert.Equal(t, "Test variable", sv.Description)
	assert.Equal(t, "default", sv.DefaultValue)
	assert.Equal(t, []string{"default", "alternative"}, sv.EnumValues)
}

func TestServerConfiguration(t *testing.T) {
	sc := ServerConfiguration{
		URL:         "https://api.example.com",
		Description: "Test server",
		Variables: map[string]ServerVariable{
			"version": {DefaultValue: "v1"},
		},
	}

	assert.Equal(t, "https://api.example.com", sc.URL)
	assert.Equal(t, "Test server", sc.Description)
	assert.NotNil(t, sc.Variables)
	assert.Equal(t, "v1", sc.Variables["version"].DefaultValue)
}

func TestConfiguration_Fields(t *testing.T) {
	cfg := &Configuration{
		Host:          "api.example.com",
		Scheme:        "https",
		DefaultHeader: map[string]string{"X-Test": "value"},
		UserAgent:     "Test-Agent/1.0",
		Debug:         true,
		HTTPClient:    &http.Client{},
	}

	assert.Equal(t, "api.example.com", cfg.Host)
	assert.Equal(t, "https", cfg.Scheme)
	assert.Equal(t, "value", cfg.DefaultHeader["X-Test"])
	assert.Equal(t, "Test-Agent/1.0", cfg.UserAgent)
	assert.True(t, cfg.Debug)
	assert.NotNil(t, cfg.HTTPClient)
}

func TestReportError(t *testing.T) {
	err := reportError("test error: %s", "details")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "test error: details")
}

func ExampleNewConfiguration() {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("X-API-Key", "your-api-key")
	cfg.Host = "api.agilitycms.com"
	cfg.Scheme = "https"

	fmt.Printf("Host: %s\n", cfg.Host)
	fmt.Printf("Scheme: %s\n", cfg.Scheme)
	fmt.Printf("User Agent: %s\n", cfg.UserAgent)
	// Output:
	// Host: api.agilitycms.com
	// Scheme: https
	// User Agent: OpenAPI-Generator/1.0.0/go
}