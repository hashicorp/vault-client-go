// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !windows

package vault

import (
	"bytes"
	"fmt"
	"runtime"

	"golang.org/x/sys/unix"
)

// constructUserAgent returns a user agent string [executable/version (os version; go version)]
func constructUserAgent(clientVersion string) string {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return fmt.Sprintf(
			"vault-client-go/%s (Go %s)",
			clientVersion,
			runtime.Version(),
		)
	}

	return fmt.Sprintf(
		"vault-client-go/%s (%s %s; Go %s)",
		clientVersion,
		string(bytes.Trim(uname.Sysname[:], "\x00")),
		string(bytes.Trim(uname.Machine[:], "\x00")),
		runtime.Version(),
	)
}
