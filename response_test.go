package vault

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResponseGeneric(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected *Response[map[string]interface{}]
	}{{
		name:     "empty",
		body:     "",
		expected: nil,
	}, {
		name: "with-data",
		body: `{"data":{"key1":"value1","key2":"value2"}}`,
		expected: &Response[map[string]interface{}]{
			Data: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}, {
		name: "with-no-data",
		body: `{"key1":"value1","key2":"value2"}`,
		expected: &Response[map[string]interface{}]{
			Data: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseResponse[map[string]interface{}](strings.NewReader(tt.body))
			require.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

type testStruct struct {
	TestString string `json:"test_string"`
	TestBool   bool   `json:"test_bool"`
}

func TestParseResponseStructured(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected *Response[testStruct]
	}{{
		name:     "empty",
		body:     "",
		expected: nil,
	}, {
		name: "with-data",
		body: `{"data":{"test_string":"test","test_bool":true}}`,
		expected: &Response[testStruct]{
			Data: testStruct{
				TestString: "test",
				TestBool:   true,
			},
		},
	}, {
		name: "with-no-data",
		body: `{"test_string":"test","test_bool":false}`,
		expected: &Response[testStruct]{
			Data: testStruct{
				TestString: "test",
				TestBool:   false,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseResponse[testStruct](strings.NewReader(tt.body))
			require.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
