package vault

import (
	"net/http"

	"golang.org/x/time/rate"
)

type ClientOption func(*Configuration) error

func New(opts ...ClientOption) (*Client, error) {
	configuration := DefaultConfiguration()
	for _, opt := range opts {
		if err := opt(&configuration); err != nil {
			return nil, err
		}
	}

	return NewClient(configuration)
}

func FromEnv(configuration *Configuration) error {
	return configuration.LoadEnvironment()
}

func WithBaseAddress(address string) ClientOption {
	return func(configuration *Configuration) error {
		configuration.BaseAddress = address
		return nil
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(configuration *Configuration) error {
		configuration.BaseClient = client
		return nil
	}
}

func WithTLSConfiguration(tlsConfig TLSConfiguration) ClientOption {
	return func(configuration *Configuration) error {
		configuration.TLS = tlsConfig
		return nil
	}
}

func WithRetryConfiguration(retryConfig RetryConfiguration) ClientOption {
	return func(configuration *Configuration) error {
		configuration.Retry = retryConfig
		return nil
	}
}

func WithRateLimiter(limiter *rate.Limiter) ClientOption {
	return func(configuration *Configuration) error {
		configuration.RateLimiter = limiter
		return nil
	}
}

func WithEnforceReadYourWritesConsistency(enabled bool) ClientOption {
	return func(configuration *Configuration) error {
		configuration.EnforceReadYourWritesConsistency = enabled
		return nil
	}
}

func WithInitialToken(token string) ClientOption {
	return func(configuration *Configuration) error {
		configuration.InitialToken = token
		return nil
	}
}

func WithInitialNamespace(namespace string) ClientOption {
	return func(configuration *Configuration) error {
		configuration.InitialNamespace = namespace
		return nil
	}
}

func WithEnableSRVLookup(enabled bool) ClientOption {
	return func(configuration *Configuration) error {
		configuration.EnableSRVLookup = enabled
		return nil
	}
}

func WithDisableRedirects(disable bool) ClientOption {
	return func(configuration *Configuration) error {
		configuration.DisableRedirects = disable
		return nil
	}
}
