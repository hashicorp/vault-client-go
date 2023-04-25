// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parameterToString(t *testing.T) {
	cases := []struct {
		value    interface{}
		expected string
	}{
		{value: false, expected: "false"},
		{value: 12345, expected: "12345"},
		{value: 12.45, expected: "12.45"},
		{value: "str", expected: "str"},
	}

	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			require.Equal(t, tc.expected, parameterToString(tc.value))
		})
	}
}
