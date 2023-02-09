// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build windows

package vault

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

// constructUserAgent returns a user agent string [executable/version (os version; go version)]
func constructUserAgent(clientVersion string) string {
	windowsVersion, err := windows.GetVersion()
	if err != nil {
		return fmt.Sprintf(
			"vault-client-go/%s (Go %s)",
			clientVersion,
			runtime.Version(),
		)
	}

	return fmt.Sprintf(
		"vault-client-go/%s (Windows %d.%d; Go %s)",
		clientVersion,
		windowsVersion&0xFF,
		windowsVersion&0xFF00>>8,
		runtime.Version(),
	)
}
