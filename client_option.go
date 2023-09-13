// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// ClientOption is a configuration option to initialize a client.
type ClientOption func(*ClientConfiguration) error

// WithAddress specifies the Vault server base address in the form of
// scheme://host:port
//
// Default: https://127.0.0.1:8200
func WithAddress(address string) ClientOption {
	return func(c *ClientConfiguration) error {
		c.Address = address
		return nil
	}
}

// WithHTTPClient sets the HTTP client to use for all API requests.
// The library sets reasonable defaults for the HTTPClient and its associated
// http.Transport. If you must modify Vault's defaults, it is suggested that
// you start with DefaultConfiguration().HTTPClient and modify it as needed
// rather than starting with an empty client or http.DefaultClient.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *ClientConfiguration) error {
		c.HTTPClient = client
		return nil
	}
}

// WithRequestTimeout, given a non-negative value, will apply the timeout to
// each request function unless an earlier deadline is passed to the request
// function through context.Context. Note that this timeout is not applicable
// to client.ReadRaw(...) or client.ReadRawWithParameters(...).
//
// Default: 60s
func WithRequestTimeout(timeout time.Duration) ClientOption {
	return func(c *ClientConfiguration) error {
		if timeout < 0 {
			return fmt.Errorf("request timeout must not be negative")
		}
		c.RequestTimeout = timeout
		return nil
	}
}

// WithTLS configures the TLS settings in the base http.Client.
func WithTLS(tls TLSConfiguration) ClientOption {
	return func(c *ClientConfiguration) error {
		c.TLS = tls
		return nil
	}
}

// WithRetryConfiguration configures the internal go-retryablehttp client.
// The library sets reasonable defaults for this setting.
func WithRetryConfiguration(retry RetryConfiguration) ClientOption {
	return func(c *ClientConfiguration) error {
		// if any of the required RetryConfiguration values are missing, use defaults
		defaultRetryConfiguration := DefaultConfiguration().RetryConfiguration

		if retry.CheckRetry == nil {
			retry.CheckRetry = defaultRetryConfiguration.CheckRetry
		}

		if retry.Backoff == nil {
			retry.Backoff = defaultRetryConfiguration.Backoff
		}

		if retry.ErrorHandler == nil {
			retry.ErrorHandler = defaultRetryConfiguration.ErrorHandler
		}

		c.RetryConfiguration = retry

		return nil
	}
}

// WithRateLimiter configures how frequently requests are allowed to happen.
// If this pointer is nil, then there will be no limit set. Note that an
// empty struct rate.Limiter is equivalent to blocking all requests.
//
// Default: nil
func WithRateLimiter(limiter *rate.Limiter) ClientOption {
	return func(c *ClientConfiguration) error {
		c.RateLimiter = limiter
		return nil
	}
}

// WithEnforceReadYourWritesConsistency ensures isolated read-after-write
// semantics by providing discovered cluster replication states in each
// request.
//
// Background: when running in a cluster, Vault has an eventual consistency
// model. Only one node (the leader) can write to Vault's storage. Users
// generally expect read-after-write consistency: in other words, after
// writing foo=1, a subsequent read of foo should return 1.
//
// Setting this to true will enable "Conditional Forwarding" as described in
// https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
//
// Note: careful consideration should be made prior to enabling this setting
// since there will be a performance penalty paid upon each request.
// This feature requires enterprise server-side.
func WithEnforceReadYourWritesConsistency() ClientOption {
	return func(c *ClientConfiguration) error {
		c.EnforceReadYourWritesConsistency = true
		return nil
	}
}

// WithDisableRedirects prevents the client from automatically following
// redirects. Any redirect responses will result in `RedirectError` instead.
//
// Background: by default, the client follows a single redirect; disabling
// redirects could cause issues with certain requests, e.g. raft-related
// calls will fail to redirect to the primary node.
func WithDisableRedirects() ClientOption {
	return func(c *ClientConfiguration) error {
		c.DisableRedirects = true
		return nil
	}
}

// WithConfiguration overwrites the default configuration object with the given
// one. It is recommended to start with DefaultConfiguration() and modify it as
// necessary. If only an individual configuration field needs to be modified,
// consider using other ClientOption functions.
func WithConfiguration(configuration ClientConfiguration) ClientOption {
	return func(c *ClientConfiguration) error {
		*c = configuration
		return nil
	}
}

// WithEnvironment populates the client's configuration object with values from
// environment values. The following environment variables are currently
// supported:
//
//	VAULT_ADDR, VAULT_AGENT_ADDR (vault's address, e.g. https://127.0.0.1:8200/)
//	VAULT_CLIENT_TIMEOUT         (request timeout)
//	VAULT_RATE_LIMIT             (rate[:burst] in operations per second)
//	VAULT_DISABLE_REDIRECTS      (prevents vault client from following redirects)
//	VAULT_TOKEN                  (the initial authentication token)
//	VAULT_NAMESPACE              (the initial namespace to use)
//	VAULT_SKIP_VERIFY            (do not veirfy vault's presented certificate)
//	VAULT_CACERT                 (PEM-encoded CA certificate file path)
//	VAULT_CACERT_BYTES           (PEM-encoded CA certificate bytes)
//	VAULT_CAPATH                 (PEM-encoded CA certificate directory path)
//	VAULT_CLIENT_CERT            (PEM-encoded client certificate file path)
//	VAULT_CLIENT_KEY             (PEM-encoded client certificate key file path)
//	VAULT_TLS_SERVER_NAME        (used to verify the hostname on returned certificates)
//	VAULT_RETRY_WAIT_MIN         (minimum time to wait before retrying)
//	VAULT_RETRY_WAIT_MAX         (maximum time to wait before retrying)
//	VAULT_MAX_RETRIES            (maximum number of retries for certain error codes)
func WithEnvironment() ClientOption {
	return func(c *ClientConfiguration) error {
		return c.populateFromEnvironment()
	}
}
