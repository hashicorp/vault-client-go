// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import "fmt"

// parameterToString is a very simple function that returns the string
// representation of the given parameter value. We may need to expand this
// later to handle time.Time values, slices, etc.
func parameterToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
