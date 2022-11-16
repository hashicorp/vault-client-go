//go:build windows

package vault

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

// UserAgent returns a user agent string [executable/version (os version; go version)]
func UserAgent() string {
	windowsVersion, err := windows.GetVersion()
	if err != nil {
		return fmt.Sprintf(
			"vault-client-go/%s (Go %s)",
			version,
			runtime.Version(),
		)
	}

	return fmt.Sprintf(
		"vault-client-go/%s (Windows %d.%d; Go %s)",
		version,
		windowsVersion&0xFF,
		windowsVersion&0xFF00>>8,
		runtime.Version(),
	)
}
