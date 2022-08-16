/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"

	"golang.org/x/time/rate"
)

// Configuration is used to configure the creation of the client
type Configuration struct {
	// BaseAddress specifies the Vault server base address in the form of
	// scheme://host:port
	// Default: http://127.0.0.1:8200
	BaseAddress string

	// BaseClient is the HTTP client to use for all API requests.
	// DefaultConfiguration() sets reasonable defaults for the BaseClient and
	// its associated http.Transport. If you must modify Vault's defaults, it
	// is suggested that you start with that client and modify it as needed
	// rather than starting with an empty client or http.DefaultClient.
	BaseClient *http.Client

	// TLS are a set of options used to configure TLS in the internal
	// base http.Client.
	TLSOptions TLSOptions

	// RetryOptions are a set of options used to configure the internal
	// go-retryablehttp client.
	RetryOptions RetryOptions

	// RateLimiter controls how frequently requests are allowed to happen.
	// If this pointer is nil, then there will be no limit set. Note that an
	// empty struct rate.Limiter is equivalent blocking all requests.
	// Default: nil
	RateLimiter *rate.Limiter
}

// TLSOptions are a set of options used to configure TLS in the internal
// base http.Client.
type TLSOptions struct {
	// VaultServerCAFile is a path to a PEM-encoded CA certificate file or
	// bundle, which the client will use to verify the Vault server TLS
	// certificate.
	// Default: "", takes precedence over 'CAPath' and 'CACertificate'.
	VaultServerCAFile string

	// VaultServerCACertificate is a PEM-encoded CA certificate or bundle,
	// which the client will use to verify the Vault server TLS certificate.
	// Default: nil, takes precedence over 'VaultServerCAPath'.
	VaultServerCACertificate []byte

	// VaultServerCAPath is a path to a directory populated with PEM-encoded
	// certificates, which the client will use to verify the Vault server TLS
	// certificate.
	// Default: ""
	VaultServerCAPath string

	// ClientCertFile is the path to the certificate for Vault communication
	// Default: ""
	ClientCertFile string

	// ClientCertKey is the path to the private key for Vault communication
	// Default: ""
	ClientCertKey string

	// ServerName is used to verify the hostname on the returned certificates
	// unless InsecureSkipVerify is given.
	// Default: ""
	ServerName string

	// InsecureSkipVerify controls whether the client verifies the server's
	// certificate chain and hostname.
	// Default: false
	InsecureSkipVerify bool
}

// RetryOptions are a set of options used to configure the internal
// go-retryablehttp client.
type RetryOptions struct {
	// RetryWaitMin controls the minimum time to wait before retrying when
	// a 5xx or 412 error occurs.
	// Default: 1000 milliseconds
	RetryWaitMin time.Duration

	// MaxRetryWait controls the maximum time to wait before retrying when
	// a 5xx or 412 error occurs.
	// Default: 1500 milliseconds
	RetryWaitMax time.Duration

	// RetryMax controls the maximum number of times to retry when a 5xx or 412
	// error occurs. Set to -1 to disable retrying.
	// Default: 2 (for a total of three tries)
	RetryMax int

	// CheckRetry specifies a policy for handling retries. It is called after
	// each request with the response and error values returned by the http.Client.
	// Default: retryablehttp.DefaultRetryPolicy + retry on 412 responses
	CheckRetry retryablehttp.CheckRetry

	// Backoff specifies a policy for how long to wait between retries.
	// Default: retryablehttp.LinearJitterBackoff
	Backoff retryablehttp.Backoff

	// ErrorHandler specifies the custom error handler to use if any.
	// Default: retryablehttp.PassthroughErrorHandler
	ErrorHandler retryablehttp.ErrorHandler

	// Logger is a custom retryablehttp.Logger or retryablehttp.LeveledLogger.
	// Default: nil
	Logger interface{}
}

// DefaultConfiguration returns the default configuration for the client. It is
// recommended to start with this configuration and modify it as needed.
func DefaultConfiguration() Configuration {
	// Use cleanhttp, which has the same default values as net/http client, but
	// does not share state with other clients (see: gh/hashicorp/go-cleanhttp)
	defaultClient := cleanhttp.DefaultPooledClient()

	defaultClientTransport := defaultClient.Transport.(*http.Transport)
	defaultClientTransport.TLSHandshakeTimeout = 10 * time.Second
	defaultClientTransport.TLSClientConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Ensure redirects are not automatically followed since the client has its
	// own redirect-handling logic.
	defaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		// ErrUseLastResponse will ensure that http.Client.Do will not send
		// the next redirect request. Instead, it will return the most recent
		// response with a nil error. A non-nil error from http.Client.Do would
		// cause redundant retries in retryablehttp.Client on every redirect.
		return http.ErrUseLastResponse
	}

	return Configuration{
		BaseAddress: "http://127.0.0.1:8200",
		BaseClient:  defaultClient,
		RetryOptions: RetryOptions{
			RetryWaitMin: time.Millisecond * 1000,
			RetryWaitMax: time.Millisecond * 1500,
			RetryMax:     2,
			CheckRetry:   DefaultRetryPolicy,
			Backoff:      retryablehttp.LinearJitterBackoff,
			ErrorHandler: retryablehttp.PassthroughErrorHandler,
			Logger:       nil,
		},
	}
}

// DefaultRetryPolicy provides a default callback for RetryOptions.CheckRetry.
// In addition to retryablehttp.DefaultRetryPolicy, it retries on 412 responses,
// which are returned by Vault when a X-Vault-Index header isn't satisfied.
func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	retry, err := retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	if err != nil || retry {
		return retry, err
	}

	if resp != nil && resp.StatusCode == http.StatusPreconditionFailed /* 412 */ {
		return true, nil
	}

	return false, nil
}

// SetDefaultsForUninitialized sets default values for uninitialized fields.
func (c *Configuration) SetDefaultsForUninitialized() {
	defaults := DefaultConfiguration()

	if c.BaseAddress == "" {
		c.BaseAddress = defaults.BaseAddress
	}

	if c.BaseClient == nil {
		c.BaseClient = defaults.BaseClient
	}

	if c.BaseClient.Transport == nil {
		c.BaseClient.Transport = defaults.BaseClient.Transport
	}

	if c.RetryOptions.RetryWaitMin == 0 {
		c.RetryOptions.RetryWaitMin = defaults.RetryOptions.RetryWaitMin
	}

	if c.RetryOptions.RetryWaitMax == 0 {
		c.RetryOptions.RetryWaitMax = defaults.RetryOptions.RetryWaitMax
	}

	if c.RetryOptions.RetryMax == 0 {
		c.RetryOptions.RetryMax = defaults.RetryOptions.RetryMax
	}

	if c.RetryOptions.CheckRetry == nil {
		c.RetryOptions.CheckRetry = defaults.RetryOptions.CheckRetry
	}

	if c.RetryOptions.Backoff == nil {
		c.RetryOptions.Backoff = defaults.RetryOptions.Backoff
	}

	if c.RetryOptions.ErrorHandler == nil {
		c.RetryOptions.ErrorHandler = defaults.RetryOptions.ErrorHandler
	}

	if c.RetryOptions.Logger == nil {
		c.RetryOptions.Logger = defaults.RetryOptions.Logger
	}
}
