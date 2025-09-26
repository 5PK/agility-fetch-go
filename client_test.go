package agilitycms

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIClient(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	assert.NotNil(t, client)
	assert.Equal(t, cfg, client.cfg)
	assert.NotNil(t, client.ContentModelsAPI)
	assert.NotNil(t, client.GalleryAPI)
	assert.NotNil(t, client.ItemAPI)
	assert.NotNil(t, client.ListAPI)
	assert.NotNil(t, client.PageAPI)
	assert.NotNil(t, client.SitemapAPI)
	assert.NotNil(t, client.SyncAPI)
	assert.NotNil(t, client.UrlRedirectionAPI)
}

func TestNewAPIClientWithCustomHTTPClient(t *testing.T) {
	cfg := NewConfiguration()
	customClient := &http.Client{}
	cfg.HTTPClient = customClient

	client := NewAPIClient(cfg)

	assert.Equal(t, customClient, client.cfg.HTTPClient)
}

func TestNewAPIClientWithNilHTTPClient(t *testing.T) {
	cfg := NewConfiguration()
	cfg.HTTPClient = nil

	client := NewAPIClient(cfg)

	assert.Equal(t, http.DefaultClient, client.cfg.HTTPClient)
}

func TestGetConfig(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	assert.Equal(t, cfg, client.GetConfig())
}

func TestSelectHeaderContentType(t *testing.T) {
	tests := []struct {
		name         string
		contentTypes []string
		expected     string
	}{
		{"Empty slice", []string{}, ""},
		{"JSON preferred", []string{"text/plain", "application/json", "application/xml"}, "application/json"},
		{"No JSON", []string{"text/plain", "application/xml"}, "text/plain"},
		{"JSON only", []string{"application/json"}, "application/json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := selectHeaderContentType(tt.contentTypes)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSelectHeaderAccept(t *testing.T) {
	tests := []struct {
		name     string
		accepts  []string
		expected string
	}{
		{"Empty slice", []string{}, ""},
		{"JSON preferred", []string{"text/plain", "application/json", "application/xml"}, "application/json"},
		{"No JSON", []string{"text/plain", "application/xml"}, "text/plain,application/xml"},
		{"JSON only", []string{"application/json"}, "application/json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := selectHeaderAccept(tt.accepts)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		haystack []string
		needle   string
		expected bool
	}{
		{"Found exact match", []string{"apple", "banana", "cherry"}, "banana", true},
		{"Found case insensitive", []string{"Apple", "BANANA", "cherry"}, "banana", true},
		{"Not found", []string{"apple", "banana", "cherry"}, "grape", false},
		{"Empty haystack", []string{}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := contains(tt.haystack, tt.needle)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTypeCheckParameter(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected string
		paramName string
		wantErr  bool
	}{
		{"Nil object", nil, "string", "test", false},
		{"Correct type", "hello", "string", "test", false},
		{"Wrong type", 123, "string", "test", true},
		{"Correct int type", 123, "int", "test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := typeCheckParameter(tt.obj, tt.expected, tt.paramName)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParameterValueToString(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		key      string
		expected string
	}{
		{"String value", "hello", "", "hello"},
		{"Int value", 123, "", "123"},
		{"Bool value", true, "", "true"},
		{"Pointer to string", stringPtr("world"), "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parameterValueToString(tt.obj, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDetectContentType(t *testing.T) {
	tests := []struct {
		name     string
		body     interface{}
		expected string
	}{
		{"String body", "hello", "text/plain; charset=utf-8"},
		{"Struct body", struct{ Name string }{Name: "test"}, "application/json; charset=utf-8"},
		{"Map body", map[string]string{"key": "value"}, "application/json; charset=utf-8"},
		{"Byte slice", []byte("hello"), "text/plain; charset=utf-8"},
		{"Slice", []string{"a", "b"}, "application/json; charset=utf-8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectContentType(tt.body)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParameterAddToHeaderOrQuery(t *testing.T) {
	t.Run("URL Values", func(t *testing.T) {
		params := url.Values{}
		parameterAddToHeaderOrQuery(params, "test", "value", "simple", "")

		assert.Equal(t, "value", params.Get("test"))
	})

	t.Run("Header Map", func(t *testing.T) {
		headers := make(map[string]string)
		parameterAddToHeaderOrQuery(headers, "test", "value", "simple", "")

		assert.Equal(t, "value", headers["test"])
	})

	t.Run("CSV Collection", func(t *testing.T) {
		params := url.Values{}
		params.Set("test", "first")
		parameterAddToHeaderOrQuery(params, "test", "second", "simple", "csv")

		assert.Equal(t, "first,second", params.Get("test"))
	})
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"123", 123, false},
		{"0", 0, false},
		{"-456", -456, false},
		{"abc", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := atoi(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGenericOpenAPIError(t *testing.T) {
	body := []byte("error body")
	errorMsg := "test error"
	model := map[string]interface{}{"error": "details"}

	err := GenericOpenAPIError{
		body:  body,
		error: errorMsg,
		model: model,
	}

	assert.Equal(t, errorMsg, err.Error())
	assert.Equal(t, body, err.Body())
	assert.Equal(t, model, err.Model())
}

func TestStrlen(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"🚀", 1}, // Unicode character
		{"hello 🚀", 7},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := strlen(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function for tests
func stringPtr(s string) *string {
	return &s
}