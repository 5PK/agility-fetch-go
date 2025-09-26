package agilitycms

import (
	"encoding/json"
	"testing"
	 

	"github.com/stretchr/testify/assert"
)

func TestPtrBool(t *testing.T) {
	value := true
	ptr := PtrBool(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrInt(t *testing.T) {
	value := 42
	ptr := PtrInt(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrInt32(t *testing.T) {
	value := int32(42)
	ptr := PtrInt32(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrInt64(t *testing.T) {
	value := int64(42)
	ptr := PtrInt64(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrFloat32(t *testing.T) {
	value := float32(3.14)
	ptr := PtrFloat32(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrFloat64(t *testing.T) {
	value := 3.14159
	ptr := PtrFloat64(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrString(t *testing.T) {
	value := "test"
	ptr := PtrString(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestPtrTime(t *testing.T) {
	value := time.Now()
	ptr := PtrTime(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestNullableBool(t *testing.T) {
	t.Run("NewNullableBool", func(t *testing.T) {
		value := true
		nb := NewNullableBool(&value)
		assert.NotNil(t, nb)
		assert.True(t, nb.IsSet())
		assert.Equal(t, &value, nb.Get())
	})

	t.Run("Set and Get", func(t *testing.T) {
		var nb NullableBool
		value := false
		nb.Set(&value)
		assert.True(t, nb.IsSet())
		assert.Equal(t, &value, nb.Get())
	})

	t.Run("Unset", func(t *testing.T) {
		value := true
		nb := NewNullableBool(&value)
		nb.Unset()
		assert.False(t, nb.IsSet())
		assert.Nil(t, nb.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := true
		nb := NewNullableBool(&value)
		data, err := nb.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "true", string(data))
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var nb NullableBool
		err := nb.UnmarshalJSON([]byte("false"))
		assert.NoError(t, err)
		assert.True(t, nb.IsSet())
		assert.False(t, *nb.Get())
	})
}

func TestNullableInt(t *testing.T) {
	t.Run("NewNullableInt", func(t *testing.T) {
		value := 42
		ni := NewNullableInt(&value)
		assert.NotNil(t, ni)
		assert.True(t, ni.IsSet())
		assert.Equal(t, &value, ni.Get())
	})

	t.Run("Set and Get", func(t *testing.T) {
		var ni NullableInt
		value := 123
		ni.Set(&value)
		assert.True(t, ni.IsSet())
		assert.Equal(t, &value, ni.Get())
	})

	t.Run("Unset", func(t *testing.T) {
		value := 42
		ni := NewNullableInt(&value)
		ni.Unset()
		assert.False(t, ni.IsSet())
		assert.Nil(t, ni.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := 42
		ni := NewNullableInt(&value)
		data, err := ni.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "42", string(data))
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var ni NullableInt
		err := ni.UnmarshalJSON([]byte("123"))
		assert.NoError(t, err)
		assert.True(t, ni.IsSet())
		assert.Equal(t, 123, *ni.Get())
	})
}

func TestNullableInt32(t *testing.T) {
	t.Run("NewNullableInt32", func(t *testing.T) {
		value := int32(42)
		ni := NewNullableInt32(&value)
		assert.NotNil(t, ni)
		assert.True(t, ni.IsSet())
		assert.Equal(t, &value, ni.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := int32(42)
		ni := NewNullableInt32(&value)
		data, err := ni.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "42", string(data))
	})
}

func TestNullableInt64(t *testing.T) {
	t.Run("NewNullableInt64", func(t *testing.T) {
		value := int64(42)
		ni := NewNullableInt64(&value)
		assert.NotNil(t, ni)
		assert.True(t, ni.IsSet())
		assert.Equal(t, &value, ni.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := int64(42)
		ni := NewNullableInt64(&value)
		data, err := ni.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "42", string(data))
	})
}

func TestNullableFloat32(t *testing.T) {
	t.Run("NewNullableFloat32", func(t *testing.T) {
		value := float32(3.14)
		nf := NewNullableFloat32(&value)
		assert.NotNil(t, nf)
		assert.True(t, nf.IsSet())
		assert.Equal(t, &value, nf.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := float32(3.14)
		nf := NewNullableFloat32(&value)
		data, err := nf.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "3.14", string(data))
	})
}

func TestNullableFloat64(t *testing.T) {
	t.Run("NewNullableFloat64", func(t *testing.T) {
		value := 3.14159
		nf := NewNullableFloat64(&value)
		assert.NotNil(t, nf)
		assert.True(t, nf.IsSet())
		assert.Equal(t, &value, nf.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := 3.14159
		nf := NewNullableFloat64(&value)
		data, err := nf.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "3.14159", string(data))
	})
}

func TestNullableString(t *testing.T) {
	t.Run("NewNullableString", func(t *testing.T) {
		value := "test"
		ns := NewNullableString(&value)
		assert.NotNil(t, ns)
		assert.True(t, ns.IsSet())
		assert.Equal(t, &value, ns.Get())
	})

	t.Run("Set and Get", func(t *testing.T) {
		var ns NullableString
		value := "hello"
		ns.Set(&value)
		assert.True(t, ns.IsSet())
		assert.Equal(t, &value, ns.Get())
	})

	t.Run("Unset", func(t *testing.T) {
		value := "test"
		ns := NewNullableString(&value)
		ns.Unset()
		assert.False(t, ns.IsSet())
		assert.Nil(t, ns.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		value := "test"
		ns := NewNullableString(&value)
		data, err := ns.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, `"test"`, string(data))
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var ns NullableString
		err := ns.UnmarshalJSON([]byte(`"hello"`))
		assert.NoError(t, err)
		assert.True(t, ns.IsSet())
		assert.Equal(t, "hello", *ns.Get())
	})
}

func TestNullableTime(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)

	t.Run("NewNullableTime", func(t *testing.T) {
		nt := NewNullableTime(&testTime)
		assert.NotNil(t, nt)
		assert.True(t, nt.IsSet())
		assert.Equal(t, &testTime, nt.Get())
	})

	t.Run("Set and Get", func(t *testing.T) {
		var nt NullableTime
		nt.Set(&testTime)
		assert.True(t, nt.IsSet())
		assert.Equal(t, &testTime, nt.Get())
	})

	t.Run("Unset", func(t *testing.T) {
		nt := NewNullableTime(&testTime)
		nt.Unset()
		assert.False(t, nt.IsSet())
		assert.Nil(t, nt.Get())
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		nt := NewNullableTime(&testTime)
		data, err := nt.MarshalJSON()
		assert.NoError(t, err)
		expected, _ := json.Marshal(testTime)
		assert.Equal(t, string(expected), string(data))
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var nt NullableTime
		timeJSON, _ := json.Marshal(testTime)
		err := nt.UnmarshalJSON(timeJSON)
		assert.NoError(t, err)
		assert.True(t, nt.IsSet())
		assert.Equal(t, testTime, *nt.Get())
	})
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"Nil interface", nil, true},
		{"Nil pointer", (*string)(nil), true},
		{"Nil slice", ([]string)(nil), true},
		{"Nil map", (map[string]string)(nil), true},
		{"Nil channel", (chan int)(nil), true},
		{"Non-nil string", "test", false},
		{"Non-nil int", 42, false},
		{"Non-nil pointer", PtrString("test"), false},
		{"Empty slice", []string{}, false},
		{"Empty map", map[string]string{}, false},
		{"Zero array", [3]int{}, true}, // Arrays are zero-valued when all elements are zero
		{"Non-zero array", [3]int{1, 0, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNil(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewStrictDecoder(t *testing.T) {
	data := []byte(`{"name": "test", "value": 42}`)
	decoder := newStrictDecoder(data)
	assert.NotNil(t, decoder)

	var result map[string]interface{}
	err := decoder.Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "test", result["name"])
	assert.Equal(t, float64(42), result["value"]) // JSON numbers decode as float64
}

func TestNewStrictDecoderWithUnknownFields(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
	}

	data := []byte(`{"name": "test", "unknown": "field"}`)
	decoder := newStrictDecoder(data)

	var result TestStruct
	err := decoder.Decode(&result)
	assert.Error(t, err) // Should error due to unknown field
}

// Benchmark tests for performance-critical functions
func BenchmarkIsNil(b *testing.B) {
	testCases := []interface{}{
		nil,
		"string",
		42,
		(*string)(nil),
		[]string{"a", "b", "c"},
		map[string]string{"key": "value"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			IsNil(tc)
		}
	}
}

func BenchmarkPtrString(b *testing.B) {
	value := "benchmark test string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PtrString(value)
	}
}

func BenchmarkNullableStringMarshal(b *testing.B) {
	value := "benchmark test"
	ns := NewNullableString(&value)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ns.MarshalJSON()
	}
}

func BenchmarkNullableStringUnmarshal(b *testing.B) {
	data := []byte(`"benchmark test"`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ns NullableString
		ns.UnmarshalJSON(data)
	}
}