// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/go-rootcerts"

	"golang.org/x/time/rate"
)

// ClientConfiguration is used to configure the creation of the client
type ClientConfiguration struct {
	// Address specifies the Vault server's base address in the form of
	// scheme://host:port
	// Default: https://127.0.0.1:8200
	Address string `env:"VAULT_ADDR,VAULT_AGENT_ADDR"`

	// HTTPClient is the HTTP client to use for all API requests.
	// DefaultConfiguration() sets reasonable defaults for the HTTPClient and
	// its associated http.Transport. If you must modify Vault's defaults, it
	// is suggested that you start with that client and modify it as needed
	// rather than starting with an empty client or http.DefaultClient.
	HTTPClient *http.Client

	// RequestTimeout, given a non-negative value, will apply the timeout to
	// each request function unless an earlier deadline is passed to the
	// request function through context.Context. Note that this timeout is
	// not applicable to client.ReadRaw or client.ReadRawWithParameters.
	// Default: 60s
	RequestTimeout time.Duration `env:"VAULT_CLIENT_TIMEOUT"`

	// TLS is a collection of TLS settings used to configure the internal
	// http.Client.
	TLS TLSConfiguration

	// RetryConfiguration is a collection of settings used to configure the
	// internal go-retryablehttp client.
	RetryConfiguration RetryConfiguration

	// RateLimiter controls how frequently requests are allowed to happen.
	// If this pointer is nil, then there will be no limit set. Note that an
	// empty struct rate.Limiter is equivalent to blocking all requests.
	// Default: nil
	RateLimiter *rate.Limiter `env:"VAULT_RATE_LIMIT"`

	// EnforceReadYourWritesConsistency ensures isolated read-after-write
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
	EnforceReadYourWritesConsistency bool

	// DisableRedirects prevents the client from automatically following
	// redirects. Any redirect responses will result in `RedirectError` instead.
	//
	// Background: by default, the client follows a single redirect; disabling
	// redirects could cause issues with certain requests, e.g. raft-related
	// calls will fail to redirect to the primary node.
	DisableRedirects bool `env:"VAULT_DISABLE_REDIRECTS"`

	// initialToken is populated from environment variable VAULT_TOKEN and will
	// be used as the initial token in client requests; to programmatically
	// modify the token, use client.Set/ClearToken or WithToken(...) option
	initialToken string `env:"VAULT_TOKEN"`

	// initialNamespace is populated from environment variable VAULT_NAMESPACE
	// and will be used as the initial namespace in client requests; to
	// programmatically modify the namespace use client.Set/ClearNamespace
	// or WithNamespace(...) option
	initialNamespace string `env:"VAULT_NAMESPACE"`
}

// TLSConfiguration is a collection of TLS settings used to configure the internal
// http.Client.
type TLSConfiguration struct {
	// ServerCertificate is a PEM-encoded CA certificate, which  the client
	// will use to verify the Vault server TLS certificate. It can be sourced
	// from a file, from a directory or from raw bytes.
	ServerCertificate ServerCertificateEntry

	// ClientCertificate is a PEM-encoded client certificate (signed by a CA or
	// self-signed), which is used to authenticate with Vault via the cert auth
	// method (see https://developer.hashicorp.com/vault/docs/auth/cert)
	ClientCertificate ClientCertificateEntry

	// ClientCertificateKey is a private key, which is used together with
	// ClientCertificate to authenticate with Vault via the cert auth method
	// (see https://developer.hashicorp.com/vault/docs/auth/cert)
	// Default: ""
	ClientCertificateKey ClientCertificateKeyEntry

	// ServerName is used to verify the hostname on the returned certificates
	// unless InsecureSkipVerify is given.
	// Default: ""
	ServerName string `env:"VAULT_TLS_SERVER_NAME"`

	// InsecureSkipVerify controls whether the client verifies the server's
	// certificate chain and hostname.
	// Default: false
	InsecureSkipVerify bool `env:"VAULT_SKIP_VERIFY"`
}

type ServerCertificateEntry struct {
	// FromFile is the path to a PEM-encoded CA certificate file or bundle.
	// Default: "", takes precedence over 'FromBytes' and 'FromDirectory'.
	FromFile string `env:"VAULT_CACERT"`

	// FromBytes is PEM-encoded CA certificate data.
	// Default: nil, takes precedence over 'FromDirectory'.
	FromBytes []byte `env:"VAULT_CACERT_BYTES"`

	// FromDirectory is the path to a directory populated with PEM-encoded
	// certificates.
	// Default: ""
	FromDirectory string `env:"VAULT_CAPATH"`
}

type ClientCertificateEntry struct {
	// FromFile is the path to a PEM-encoded client certificate file.
	// Default: "", takes precedence over 'FromBytes'
	FromFile string `env:"VAULT_CLIENT_CERT"`

	// FromBytes is PEM-encoded certificate data.
	// Default: nil
	FromBytes []byte
}

type ClientCertificateKeyEntry struct {
	// FromFile is the path to a PEM-encoded private key file.
	// Default: "", takes precedence over 'FromBytes'
	FromFile string `env:"VAULT_CLIENT_KEY"`

	// FromBytes is PEM-encoded private key data.
	// Default: nil
	FromBytes []byte
}

// RetryConfiguration is a collection of settings used to configure the internal
// go-retryablehttp client.
type RetryConfiguration struct {
	// RetryWaitMin controls the minimum time to wait before retrying when
	// a 5xx or 412 error occurs.
	// Default: 1000 milliseconds
	RetryWaitMin time.Duration `env:"VAULT_RETRY_WAIT_MIN"`

	// RetryWaitMax controls the maximum time to wait before retrying when
	// a 5xx or 412 error occurs.
	// Default: 1500 milliseconds
	RetryWaitMax time.Duration `env:"VAULT_RETRY_WAIT_MAX"`

	// RetryMax controls the maximum number of times to retry when a 5xx or 412
	// error occurs. Set to -1 to disable retrying.
	// Default: 2 (for a total of three tries)
	RetryMax int `env:"VAULT_MAX_RETRIES"`

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
func DefaultConfiguration() ClientConfiguration {
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

	return ClientConfiguration{
		Address:        "https://127.0.0.1:8200",
		HTTPClient:     defaultClient,
		RequestTimeout: 60 * time.Second,
		RetryConfiguration: RetryConfiguration{
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

// populateFromEnvironment populates the configuration object with values from
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
func (c *ClientConfiguration) populateFromEnvironment() error {
	// this function will be recursively applied to each field within the configuration object
	assignFieldFromEnvironment := func(field reflect.Value, environmentTags []string) error {
		// for each 'env' tag ...
		for _, tag := range environmentTags {
			// try to fetch the environment variable
			env := os.Getenv(tag)
			if env == "" {
				continue
			}

			switch field.Type().String() {

			case "string":
				field.SetString(env)

			case "[]uint8":
				field.SetBytes([]byte(env))

			case "bool":
				v, err := strconv.ParseBool(env)
				if err != nil {
					return fmt.Errorf("could not convert %q environment variable value to bool", tag)
				}
				field.SetBool(v)

			case "int":
				v, err := strconv.Atoi(env)
				if err != nil {
					return fmt.Errorf("could not convert %q environment variable value to int", tag)
				}
				field.SetInt(int64(v))

			case "time.Duration":
				v, err := time.ParseDuration(env)
				if err != nil {
					return fmt.Errorf("could not convert %q environment variable value to time.Duration", tag)
				}
				field.SetInt(int64(v))

			case "*rate.Limiter":
				var (
					limiterRate  float64
					limiterBurst int
				)
				_, err := fmt.Sscanf(env, "%f:%d", &limiterRate, &limiterBurst)
				if err != nil {
					limiterRate, err = strconv.ParseFloat(env, 64)
					if err != nil {
						return fmt.Errorf("%q environment variable expects either 'rate:burst' or a float 'rate'", tag)
					}
					limiterBurst = int(limiterRate)
				}
				field.Set(reflect.ValueOf(rate.NewLimiter(rate.Limit(limiterRate), limiterBurst)))

			default:
				return fmt.Errorf("environment variable parsing not implemented for %q type", field.Type().String())
			}
		}

		return nil
	}

	// work on a copy of the configuration object to prevent modfications on errors
	copy := *c

	if err := walkConfigurationFields(&copy, assignFieldFromEnvironment); err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	// assign initial token & namespace manually since they are not exported
	if env := os.Getenv("VAULT_TOKEN"); env != "" {
		if err := validateToken(env); err != nil {
			return fmt.Errorf("configuration error: VAULT_TOKEN: %w", err)
		}
		copy.initialToken = env
	}

	if env := os.Getenv("VAULT_NAMESPACE"); env != "" {
		if err := validateNamespace(env); err != nil {
			return fmt.Errorf("configuration error: VAULT_NAMESPACE: %w", err)
		}
		copy.initialNamespace = env
	}

	*c = copy

	return nil
}

// DefaultRetryPolicy provides a default callback for RetryConfiguration.CheckRetry.
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

// empty returns true if t is equivalent to an empty TLSConfiguration{} object.
func (t *TLSConfiguration) empty() bool {
	if t.ServerCertificate.FromFile != "" {
		return false
	}

	if len(t.ServerCertificate.FromBytes) != 0 {
		return false
	}

	if t.ServerCertificate.FromDirectory != "" {
		return false
	}

	if t.ClientCertificate.FromFile != "" {
		return false
	}

	if len(t.ClientCertificate.FromBytes) != 0 {
		return false
	}

	if t.ClientCertificateKey.FromFile != "" {
		return false
	}

	if len(t.ClientCertificateKey.FromBytes) != 0 {
		return false
	}

	if t.ServerName != "" {
		return false
	}

	if t.InsecureSkipVerify {
		return false
	}

	return true
}

// applyTo applies the user-defined TLS configuration to the given client's
// *tls.Config pointer; it is used to configure the internal http.Client
func (from *TLSConfiguration) applyTo(to *tls.Config) error {
	if len(from.ServerCertificate.FromBytes) != 0 || from.ServerCertificate.FromFile != "" || from.ServerCertificate.FromDirectory != "" {
		rootCertificateConfig := rootcerts.Config{
			CAFile:        from.ServerCertificate.FromFile,
			CACertificate: from.ServerCertificate.FromBytes,
			CAPath:        from.ServerCertificate.FromDirectory,
		}
		if err := rootcerts.ConfigureTLS(
			to,
			&rootCertificateConfig,
		); err != nil {
			return fmt.Errorf("could not configure root certificate: %w", err)
		}
	}

	read := func(fromFile string, fromBytes []byte) ([]byte, error) {
		if fromFile != "" {
			b, err := os.ReadFile(fromFile)
			if err != nil {
				return nil, err
			}
			return b, nil
		}
		return fromBytes, nil
	}

	var (
		hasClientCertificate    = from.ClientCertificate.FromFile != "" || len(from.ClientCertificate.FromBytes) != 0
		hasClientCertificateKey = from.ClientCertificateKey.FromFile != "" || len(from.ClientCertificateKey.FromBytes) != 0
	)

	if hasClientCertificate != hasClientCertificateKey {
		return fmt.Errorf("client certificate and client certificate key must be provided together")
	}

	if hasClientCertificate && hasClientCertificateKey {
		clientCertificateBytes, err := read(from.ClientCertificate.FromFile, from.ClientCertificate.FromBytes)
		if err != nil {
			return fmt.Errorf("could not read certificate file: %w", err)
		}

		clientCertificateKeyBytes, err := read(from.ClientCertificateKey.FromFile, from.ClientCertificateKey.FromBytes)
		if err != nil {
			return fmt.Errorf("could not read certificate key file: %w", err)
		}

		clientCertificate, err := tls.X509KeyPair(clientCertificateBytes, clientCertificateKeyBytes)
		if err != nil {
			return fmt.Errorf("error parsing certificate pair: %w", err)
		}

		// Set this function to ignore the server's preferential list of CAs.
		// Otherwise, any CA used for the certificate auth backend must be in
		// the server's CA pool.
		to.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &clientCertificate, nil
		}
	}

	if from.InsecureSkipVerify {
		to.InsecureSkipVerify = from.InsecureSkipVerify
	}

	if from.ServerName != "" {
		to.ServerName = from.ServerName
	}

	return nil
}

// walkConfigurationFields is a helper function, which uses reflection to
// traverse the configuration fields, determine their `env` tags and apply the
// given function f to the modifiable fields.
func walkConfigurationFields(structPtr any, f func(field reflect.Value, environmentTags []string) error) error {
	// we expect a pointer to a struct to be able to modify the fields
	if reflect.ValueOf(structPtr).Kind() != reflect.Ptr {
		return fmt.Errorf("pointer input expected")
	}

	// struct value and type
	structV := reflect.ValueOf(structPtr).Elem()
	structT := reflect.TypeOf(structPtr).Elem()

	for i := 0; i < structT.NumField(); i++ {
		// field value and type
		fieldV := structV.Field(i)
		fieldT := structT.Field(i)

		switch {
		case !fieldV.CanSet():
			continue // unexported fields will be skipped

		case fieldV.Kind() == reflect.Struct:
			if err := walkConfigurationFields(fieldV.Addr().Interface(), f); err != nil {
				return err
			}

		default:
			tags := fieldT.Tag.Get("env")

			if tags == "" {
				continue // no 'env' tags found
			}

			if err := f(fieldV, strings.Split(tags, ",")); err != nil {
				return fmt.Errorf("could not configure %q: %w", fieldT.Name, err)
			}
		}
	}

	return nil
}
