package vault

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseResponse_generic(t *testing.T) {
	cases := map[string]struct {
		body     string
		expected *Response[map[string]interface{}]
	}{
		"empty": {
			body:     "",
			expected: nil,
		},
		"empty-object": {
			body: `{}`,
			expected: &Response[map[string]interface{}]{
				Data: map[string]interface{}{},
			},
		},
		"with-data": {
			body: `{"data":{"key1":"value1","key2":"value2"}}`,
			expected: &Response[map[string]interface{}]{
				Data: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		"with-no-data": {
			body: `{"key1":"value1","key2":"value2"}`,
			expected: &Response[map[string]interface{}]{
				Data: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := parseResponse[map[string]interface{}](strings.NewReader(tc.body))
			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

type testStruct struct {
	TestString string `json:"test_string"`
	TestBool   bool   `json:"test_bool"`
}

func Test_parseResponse_structured(t *testing.T) {
	cases := map[string]struct {
		body     string
		expected *Response[testStruct]
	}{
		"empty": {
			body:     "",
			expected: nil,
		},
		"empty-object": {
			body: `{}`,
			expected: &Response[testStruct]{
				Data: testStruct{},
			},
		},
		"with-data": {
			body: `{"data":{"test_string":"test","test_bool":true}}`,
			expected: &Response[testStruct]{
				Data: testStruct{
					TestString: "test",
					TestBool:   true,
				},
			},
		},
		"with-no-data": {
			body: `{"test_string":"test","test_bool":false}`,
			expected: &Response[testStruct]{
				Data: testStruct{
					TestString: "test",
					TestBool:   false,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := parseResponse[testStruct](strings.NewReader(tc.body))
			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
