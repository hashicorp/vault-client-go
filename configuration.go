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
	"unsafe"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/go-rootcerts"

	"golang.org/x/time/rate"
)

// Configuration is used to configure the creation of the client
type Configuration struct {
	// BaseAddress specifies the Vault server base address in the form of
	// scheme://host:port
	// Default: https://127.0.0.1:8200
	BaseAddress string `env:"VAULT_ADDR,VAULT_AGENT_ADDR"`

	// BaseClient is the HTTP client to use for all API requests.
	// DefaultConfiguration() sets reasonable defaults for the BaseClient and
	// its associated http.Transport. If you must modify Vault's defaults, it
	// is suggested that you start with that client and modify it as needed
	// rather than starting with an empty client or http.DefaultClient.
	BaseClient *http.Client

	// TLS is a collection of TLS settings used to configure the internal
	// http.Client.
	TLS TLSConfiguration

	// Retry is a collection of settings used to configure the internal
	// go-retryablehttp client.
	Retry RetryConfiguration

	// RateLimiter controls how frequently requests are allowed to happen.
	// If this pointer is nil, then there will be no limit set. Note that an
	// empty struct rate.Limiter is equivalent blocking all requests.
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
	// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
	//
	// Note: careful consideration should be made prior to enabling this setting
	// since there will be a performance penalty paid upon each request.
	// This feature requires enterprise server-side.
	EnforceReadYourWritesConsistency bool

	// InitialToken will be used as the token in client requests unless
	// overwritten with client.SetToken or client.WithToken
	InitialToken string `env:"VAULT_TOKEN"`

	// InitialNamespace will be used as the namespace in client requests unless
	// overwritten with client.SetNamespace or client.WithNamespace
	InitialNamespace string `env:"VAULT_NAMESPACE"`

	// EnableSRVLookup enables the client to look up the Vault server host
	// through DNS SRV lookup. The lookup will happen on each request. The base
	// address' port must be empty for this setting to be respected.
	//
	// Note: this feature is not designed for high availability, just
	// discovery.
	//
	// See https://datatracker.ietf.org/doc/html/draft-andrews-http-srv-02
	// for more information
	EnableSRVLookup bool `env:"VAULT_SRV_LOOKUP"`
}

// TLSConfiguration is a collection of TLS settings used to configure the internal
// http.Client.
type TLSConfiguration struct {
	// ServerCACertificateFile is a path to a PEM-encoded CA certificate
	// file or bundle, which the client will use to verify the Vault server TLS
	// certificate.
	// Default: "", takes precedence over other ServerCACertificate* variables.
	ServerCACertificateFile string `env:"VAULT_CACERT"`

	// ServerCACertificate is a PEM-encoded CA certificate or bundle,
	// which the client will use to verify the Vault server TLS certificate.
	// Default: nil, takes precedence over ServerCACertificateDir.
	ServerCACertificate []byte `env:"VAULT_CACERT_BYTES"`

	// ServerCACertificateDir is a path to a directory populated with
	// PEM-encoded certificates, which the client will use to verify the Vault
	// server TLS certificate.
	// Default: ""
	ServerCACertificateDir string `env:"VAULT_CAPATH"`

	// ClientCertificateFile is the path to a client certificate (signed by a
	// CA or self-signed), which is used to authenticate with Vault via the
	// cert auth method (see https://www.vaultproject.io/docs/auth/cert)
	// Default: ""
	ClientCertificateFile string `env:"VAULT_CLIENT_CERT"`

	// ClientCertificateKeyFile is the path to a private key, which is used
	// together with ClientCertificateFile to authenticate with Vault via the
	// cert auth method (see https://www.vaultproject.io/docs/auth/cert)
	// Default: ""
	ClientCertificateKeyFile string `env:"VAULT_CLIENT_KEY"`

	// ServerName is used to verify the hostname on the returned certificates
	// unless InsecureSkipVerify is given.
	// Default: ""
	ServerName string `env:"VAULT_TLS_SERVER_NAME"`

	// InsecureSkipVerify controls whether the client verifies the server's
	// certificate chain and hostname.
	// Default: false
	InsecureSkipVerify bool `env:"VAULT_SKIP_VERIFY"`
}

// RetryConfiguration is a collection of settings used to configure the internal
// go-retryablehttp client.
type RetryConfiguration struct {
	// RetryWaitMin controls the minimum time to wait before retrying when
	// a 5xx or 412 error occurs.
	// Default: 1000 milliseconds
	RetryWaitMin time.Duration `env:"VAULT_RETRY_WAIT_MIN"`

	// MaxRetryWait controls the maximum time to wait before retrying when
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
		BaseAddress: "https://127.0.0.1:8200",
		BaseClient:  defaultClient,
		Retry: RetryConfiguration{
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

// LoadEnvironment loads vault-specific environment variables and applies them
// to the given configuration object. The environment varialbes are specified
// in the 'env' tags defined next to each configuration field. In case of
// multiple environment variables defined for a field, the later ones take
// precedence.
func (c *Configuration) LoadEnvironment() error {
	// this function will be recursively applied to each field within the configuration object
	assignFieldFromEnv := func(field reflect.Value, environmentTags []string) error {
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
				field.SetPointer(unsafe.Pointer(rate.NewLimiter(rate.Limit(limiterRate), limiterBurst)))

			default:
				return fmt.Errorf("environment variable parsing not implemented for %q type", field.Type().String())
			}
		}

		return nil
	}

	// work on a copy of the configuration object to prevent modfications on errors
	copy := *c

	if err := walkConfigurationFields(&copy, assignFieldFromEnv); err != nil {
		return fmt.Errorf("configuration error: %w", err)
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

	if c.Retry.RetryWaitMin == 0 {
		c.Retry.RetryWaitMin = defaults.Retry.RetryWaitMin
	}

	if c.Retry.RetryWaitMax == 0 {
		c.Retry.RetryWaitMax = defaults.Retry.RetryWaitMax
	}

	if c.Retry.RetryMax == 0 {
		c.Retry.RetryMax = defaults.Retry.RetryMax
	}

	if c.Retry.CheckRetry == nil {
		c.Retry.CheckRetry = defaults.Retry.CheckRetry
	}

	if c.Retry.Backoff == nil {
		c.Retry.Backoff = defaults.Retry.Backoff
	}

	if c.Retry.ErrorHandler == nil {
		c.Retry.ErrorHandler = defaults.Retry.ErrorHandler
	}

	if c.Retry.Logger == nil {
		c.Retry.Logger = defaults.Retry.Logger
	}
}

// applyTo applies the user-defined TLS configuration to the given client's
// *tls.Config pointer; it is used to configure the internal http.Client
func (from *TLSConfiguration) applyTo(to *tls.Config) error {
	if len(from.ServerCACertificate) != 0 || from.ServerCACertificateFile != "" || from.ServerCACertificateDir != "" {
		rootCertificateConfig := rootcerts.Config{
			CAFile:        from.ServerCACertificateFile,
			CACertificate: from.ServerCACertificate,
			CAPath:        from.ServerCACertificateDir,
		}
		if err := rootcerts.ConfigureTLS(
			to,
			&rootCertificateConfig,
		); err != nil {
			return fmt.Errorf("could not configure root certificate: %w", err)
		}
	}

	switch {
	case from.ClientCertificateFile != "" && from.ClientCertificateKeyFile != "":
		clientCertificate, err := tls.LoadX509KeyPair(
			from.ClientCertificateFile,
			from.ClientCertificateKeyFile,
		)
		if err != nil {
			return fmt.Errorf(
				"could not load client certificate from certificate file %q and key file %q: %w",
				from.ClientCertificateFile,
				from.ClientCertificateKeyFile,
				err,
			)
		}

		// Set this function to ignore the server's preferential list of CAs.
		// Otherwise, any CA used for the certificate auth backend must be in
		// the server's CA pool.
		to.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &clientCertificate, nil
		}

	case from.ClientCertificateFile != "":
		return fmt.Errorf("client certificate file %q is specified but certificate key is missing", from.ClientCertificateFile)

	case from.ClientCertificateKeyFile != "":
		return fmt.Errorf("client certificate key %q is specified but certificate file is missing", from.ClientCertificateKeyFile)
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
			if err := f(fieldV, strings.Split(fieldT.Tag.Get("env"), ",")); err != nil {
				return fmt.Errorf("could not configure %q: %w", fieldT.Name, err)
			}
		}
	}

	return nil
}
