// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/vault-client-go/schema"
)

// System is a simple wrapper around the client for System requests
type System struct {
	client *Client
}

// CalculateAuditHash The hash of the given string via the given audit backend
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (a *System) CalculateAuditHash(ctx context.Context, path string, request schema.CalculateAuditHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit-hash/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteAuditDevice Disable the audit device at the given path.
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (a *System) DeleteAuditDevice(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteAuthMethod Disable the auth method at the given auth path
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (a *System) DeleteAuthMethod(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteConfigAuditingRequestHeader Disable auditing of the given request header.
func (a *System) DeleteConfigAuditingRequestHeader(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteConfigCORS Remove any CORS settings.
func (a *System) DeleteConfigCORS(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteConfigUIHeader Remove a UI header.
// header: The name of the header.
func (a *System) DeleteConfigUIHeader(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteGenerateRoot Cancels any in-progress root generation attempt.
func (a *System) DeleteGenerateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteGenerateRootAttempt Cancels any in-progress root generation attempt.
func (a *System) DeleteGenerateRootAttempt(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteLogger Revert a single logger to use log level provided in config.
// name: The name of the logger to be modified.
func (a *System) DeleteLogger(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteLoggers Revert the all loggers to use log level provided in config.
func (a *System) DeleteLoggers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteMount Disable the mount point specified at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (a *System) DeleteMount(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeletePluginsCatalogByTypeByName Remove the plugin with the given name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (a *System) DeletePluginsCatalogByTypeByName(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeletePoliciesACL Delete the ACL policy with the given name.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) DeletePoliciesACL(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeletePoliciesPassword Delete a password policy.
// name: The name of the password policy.
func (a *System) DeletePoliciesPassword(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeletePolicy Delete the policy with the given name.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) DeletePolicy(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policy/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteQuotasRateLimit
// name: Name of the quota rule.
func (a *System) DeleteQuotasRateLimit(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRaw Delete the key with given path.
func (a *System) DeleteRaw(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRawPath Delete the key with given path.
func (a *System) DeleteRawPath(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRekeyBackup Delete the backup copy of PGP-encrypted unseal keys.
func (a *System) DeleteRekeyBackup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/backup"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRekeyInit Cancels any in-progress rekey.
// This clears the rekey settings as well as any progress made. This must be called to change the parameters of the rekey. Note: verification is still a part of a rekey. If rekeying is canceled during the verification flow, the current unseal keys remain valid.
func (a *System) DeleteRekeyInit(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRekeyRecoveryKeyBackup Allows fetching or deleting the backup of the rotated unseal keys.
func (a *System) DeleteRekeyRecoveryKeyBackup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/recovery-key-backup"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DeleteRekeyVerify Cancel any in-progress rekey verification operation.
// This clears any progress made and resets the nonce. Unlike a `DELETE` against `sys/rekey/init`, this only resets the current verification operation, not the entire rekey atttempt.
func (a *System) DeleteRekeyVerify(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListConfigUIHeaders Return a list of configured UI headers.
func (a *System) ListConfigUIHeaders(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListLeasesLookupPrefix Returns a list of lease ids.
// prefix: The path to list leases under. Example: \&quot;aws/creds/deploy\&quot;
func (a *System) ListLeasesLookupPrefix(ctx context.Context, prefix string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/lookup/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListPluginsCatalogByType List the plugins in the catalog.
// type_: The type of the plugin, may be auth, secret, or database
func (a *System) ListPluginsCatalogByType(ctx context.Context, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}"
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListPoliciesACL List the configured access control policies.
func (a *System) ListPoliciesACL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListPoliciesPassword List the existing password policies.
func (a *System) ListPoliciesPassword(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListQuotasRateLimits
func (a *System) ListQuotasRateLimits(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListVersionHistory Returns map of historical version change entries
func (a *System) ListVersionHistory(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/version-history/"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MFAValidate Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
func (a *System) MFAValidate(ctx context.Context, request schema.MFAValidateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mfa/validate"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Monitor
// logFormat: Output format of logs. Supported values are \&quot;standard\&quot; and \&quot;json\&quot;. The default is \&quot;standard\&quot;.
// logLevel: Log level to view system logs at. Currently supported values are \&quot;trace\&quot;, \&quot;debug\&quot;, \&quot;info\&quot;, \&quot;warn\&quot;, \&quot;error\&quot;.
func (a *System) Monitor(ctx context.Context, logFormat string, logLevel string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/monitor"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("logFormat", url.QueryEscape(logFormat))
	requestQueryParameters.Set("logLevel", url.QueryEscape(logLevel))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofRead Returns an HTML page listing the available profiles.
// Returns an HTML page listing the available  profiles. This should be mainly accessed via browsers or applications that can  render pages.
func (a *System) PprofRead(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadAllocs Returns a sampling of all past memory allocations.
// Returns a sampling of all past memory allocations.
func (a *System) PprofReadAllocs(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/allocs"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadBlock Returns stack traces that led to blocking on synchronization primitives
// Returns stack traces that led to blocking on synchronization primitives
func (a *System) PprofReadBlock(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/block"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadCmdline Returns the running program's command line.
// Returns the running program's command line, with arguments separated by NUL bytes.
func (a *System) PprofReadCmdline(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/cmdline"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadGoroutine Returns stack traces of all current goroutines.
// Returns stack traces of all current goroutines.
func (a *System) PprofReadGoroutine(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/goroutine"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadHeap Returns a sampling of memory allocations of live object.
// Returns a sampling of memory allocations of live object.
func (a *System) PprofReadHeap(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/heap"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadMutex Returns stack traces of holders of contended mutexes
// Returns stack traces of holders of contended mutexes
func (a *System) PprofReadMutex(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/mutex"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadProfile Returns a pprof-formatted cpu profile payload.
// Returns a pprof-formatted cpu profile payload. Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
func (a *System) PprofReadProfile(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/profile"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadSymbol Returns the program counters listed in the request.
// Returns the program counters listed in the request.
func (a *System) PprofReadSymbol(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/symbol"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadThreadcreate Returns stack traces that led to the creation of new OS threads
// Returns stack traces that led to the creation of new OS threads
func (a *System) PprofReadThreadcreate(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/threadcreate"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofReadTrace Returns the execution trace in binary form.
// Returns  the execution trace in binary form. Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
func (a *System) PprofReadTrace(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/trace"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadAuditDevices List the enabled audit devices.
func (a *System) ReadAuditDevices(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadAuthMethod Read the configuration of the auth engine at the given path.
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (a *System) ReadAuthMethod(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadAuthMethodTune Reads the given auth path's configuration.
// This endpoint requires sudo capability on the final path, but the same functionality can be achieved without sudo via `sys/mounts/auth/[auth-path]/tune`.
// path: Tune the configuration parameters for an auth path.
func (a *System) ReadAuthMethodTune(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadAuthMethods List the currently enabled credential backends.
func (a *System) ReadAuthMethods(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadConfigAuditingRequestHeader List the information for the given request header.
func (a *System) ReadConfigAuditingRequestHeader(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadConfigAuditingRequestHeaders List the request headers that are configured to be audited.
func (a *System) ReadConfigAuditingRequestHeaders(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadConfigCORS Return the current CORS settings.
func (a *System) ReadConfigCORS(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadConfigStateSanitized Return a sanitized version of the Vault server configuration.
// The sanitized output strips configuration values in the storage, HA storage, and seals stanzas, which may contain sensitive values such as API tokens. It also removes any token or secret fields in other stanzas, such as the circonus_api_token from telemetry.
func (a *System) ReadConfigStateSanitized(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/state/sanitized"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadConfigUIHeader Return the given UI header's configuration
// header: The name of the header.
func (a *System) ReadConfigUIHeader(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadGenerateRoot Read the configuration and progress of the current root generation attempt.
func (a *System) ReadGenerateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadGenerateRootAttempt Read the configuration and progress of the current root generation attempt.
func (a *System) ReadGenerateRootAttempt(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadHAStatus Check the HA status of a Vault cluster
func (a *System) ReadHAStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/ha-status"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadHealth Returns the health status of Vault.
func (a *System) ReadHealth(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/health"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadHostInfo Information about the host instance that this Vault server is running on.
// Information about the host instance that this Vault server is running on.   The information that gets collected includes host hardware information, and CPU,   disk, and memory utilization
func (a *System) ReadHostInfo(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/host-info"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInFlightRequests reports in-flight requests
// This path responds to the following HTTP methods.   GET /    Returns a map of in-flight requests.
func (a *System) ReadInFlightRequests(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/in-flight-req"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInit Returns the initialization status of Vault.
func (a *System) ReadInit(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/init"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersActivity Report the client count metrics, for this namespace and all child namespaces.
func (a *System) ReadInternalCountersActivity(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersActivityExport Report the client count metrics, for this namespace and all child namespaces.
func (a *System) ReadInternalCountersActivityExport(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity/export"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersActivityMonthly Report the number of clients for this month, for this namespace and all child namespaces.
func (a *System) ReadInternalCountersActivityMonthly(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity/monthly"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersConfig Read the client count tracking configuration.
func (a *System) ReadInternalCountersConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/config"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersEntities Backwards compatibility is not guaranteed for this API
func (a *System) ReadInternalCountersEntities(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/entities"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersRequests Backwards compatibility is not guaranteed for this API
func (a *System) ReadInternalCountersRequests(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/requests"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalCountersTokens Backwards compatibility is not guaranteed for this API
func (a *System) ReadInternalCountersTokens(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/tokens"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalInspectRouter Expose the route entry and mount entry tables present in the router
// tag: Name of subtree being observed
func (a *System) ReadInternalInspectRouter(ctx context.Context, tag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/inspect/router/{tag}"
	requestPath = strings.Replace(requestPath, "{"+"tag"+"}", url.PathEscape(tag), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalSpecsOpenAPI Generate an OpenAPI 3 document of all mounted paths.
func (a *System) ReadInternalSpecsOpenAPI(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/specs/openapi"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalUIFeatureFlags Lists enabled feature flags.
func (a *System) ReadInternalUIFeatureFlags(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/feature-flags"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalUIMount Return information about the given mount.
// path: The path of the mount.
func (a *System) ReadInternalUIMount(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalUIMounts Lists all enabled and visible auth and secrets mounts.
func (a *System) ReadInternalUIMounts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/mounts"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalUINamespaces Backwards compatibility is not guaranteed for this API
func (a *System) ReadInternalUINamespaces(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/namespaces"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInternalUIResultantACL Backwards compatibility is not guaranteed for this API
func (a *System) ReadInternalUIResultantACL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/resultant-acl"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadKeyStatus Provides information about the backend encryption key.
func (a *System) ReadKeyStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/key-status"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadLeader Returns the high availability status and current leader instance of Vault.
func (a *System) ReadLeader(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leader"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadLeases List leases associated with this Vault cluster
func (a *System) ReadLeases(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadLeasesCount Count of leases associated with this Vault cluster
func (a *System) ReadLeasesCount(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/count"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadLogger Read the log level for a single logger.
// name: The name of the logger to be modified.
func (a *System) ReadLogger(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadLoggers Read the log level for all existing loggers.
func (a *System) ReadLoggers(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadMetrics Export the metrics aggregated for telemetry purpose.
// format: Format to export metrics into. Currently accepts only \&quot;prometheus\&quot;.
func (a *System) ReadMetrics(ctx context.Context, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/metrics"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadMount Read the configuration of the secret engine at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (a *System) ReadMount(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadMounts List the currently mounted backends.
func (a *System) ReadMounts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadMountsConfig Tune backend configuration parameters for this mount.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (a *System) ReadMountsConfig(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPluginsCatalog Lists all the plugins known to Vault
func (a *System) ReadPluginsCatalog(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPluginsCatalogByTypeByName Return the configuration data for the plugin with the given name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (a *System) ReadPluginsCatalogByTypeByName(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPolicies List the configured access control policies.
func (a *System) ReadPolicies(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policy"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPoliciesACL Retrieve information about the named ACL policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) ReadPoliciesACL(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPoliciesPassword Retrieve an existing password policy.
// name: The name of the password policy.
func (a *System) ReadPoliciesPassword(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPoliciesPasswordGenerate Generate a password from an existing password policy.
// name: The name of the password policy.
func (a *System) ReadPoliciesPasswordGenerate(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}/generate"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadPolicy Retrieve the policy body for the named policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) ReadPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policy/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadQuotasConfig
func (a *System) ReadQuotasConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/config"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadQuotasRateLimit
// name: Name of the quota rule.
func (a *System) ReadQuotasRateLimit(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRaw Read the value of the key at the given path.
func (a *System) ReadRaw(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRawPath Read the value of the key at the given path.
func (a *System) ReadRawPath(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRekeyBackup Return the backup copy of PGP-encrypted unseal keys.
func (a *System) ReadRekeyBackup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/backup"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRekeyInit Reads the configuration and progress of the current rekey attempt.
func (a *System) ReadRekeyInit(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRekeyRecoveryKeyBackup Allows fetching or deleting the backup of the rotated unseal keys.
func (a *System) ReadRekeyRecoveryKeyBackup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/recovery-key-backup"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRekeyVerify Read the configuration and progress of the current rekey verification attempt.
func (a *System) ReadRekeyVerify(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRemountStatus Check status of a mount migration
// migrationId: The ID of the migration operation
func (a *System) ReadRemountStatus(ctx context.Context, migrationId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/remount/status/{migration_id}"
	requestPath = strings.Replace(requestPath, "{"+"migration_id"+"}", url.PathEscape(migrationId), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadReplicationStatus
func (a *System) ReadReplicationStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/replication/status"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadRotateConfig
func (a *System) ReadRotateConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate/config"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadSealStatus Check the seal status of a Vault.
func (a *System) ReadSealStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/seal-status"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Remount Initiate a mount migration
func (a *System) Remount(ctx context.Context, request schema.RemountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/remount"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Renew Renews a lease, requesting to extend the lease.
func (a *System) Renew(ctx context.Context, request schema.RenewRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/renew"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RenewFor Renews a lease, requesting to extend the lease.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (a *System) RenewFor(ctx context.Context, urlLeaseId string, request schema.RenewForRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/renew/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Revoke Revokes a lease immediately.
func (a *System) Revoke(ctx context.Context, request schema.RevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/revoke"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RevokeForce Revokes all secrets or tokens generated under a given prefix immediately
// Unlike `/sys/leases/revoke-prefix`, this path ignores backend errors encountered during revocation. This is potentially very dangerous and should only be used in specific emergency situations where errors in the backend or the connected backend service prevent normal revocation.  By ignoring these errors, Vault abdicates responsibility for ensuring that the issued credentials or secrets are properly revoked and/or cleaned up. Access to this endpoint should be tightly controlled.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (a *System) RevokeForce(ctx context.Context, prefix string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/revoke-force/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RevokeLease Revokes a lease immediately.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (a *System) RevokeLease(ctx context.Context, urlLeaseId string, request schema.RevokeLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/revoke/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RevokePrefix Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (a *System) RevokePrefix(ctx context.Context, prefix string, request schema.RevokePrefixRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/revoke-prefix/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Rotate Rotates the backend encryption key used to persist data.
func (a *System) Rotate(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Seal Seal the Vault.
func (a *System) Seal(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/seal"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// StepDownLeader Cause the node to give up active status.
// This endpoint forces the node to give up active status. If the node does not have active status, this endpoint does nothing. Note that the node will sleep for ten seconds before attempting to grab the active lock again, but if no standby nodes grab the active lock in the interim, the same node may become the active node again.
func (a *System) StepDownLeader(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/step-down"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SysDeletePluginsCatalogName Remove the plugin with the given name.
// name: The name of the plugin
func (a *System) SysDeletePluginsCatalogName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SysListLeasesLookup Returns a list of lease ids.
func (a *System) SysListLeasesLookup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/lookup/"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SysReadPluginsCatalogName Return the configuration data for the plugin with the given name.
// name: The name of the plugin
func (a *System) SysReadPluginsCatalogName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SysWriteLockedusersMountAccessorUnlockAliasIdentifier Unlocks the user with given mount_accessor and alias_identifier
// aliasIdentifier: It is the name of the alias (user). For example, if the alias belongs to userpass backend, the name should be a valid username within userpass auth method. If the alias belongs to an approle auth method, the name should be a valid RoleID
// mountAccessor: MountAccessor is the identifier of the mount entry to which the user belongs
func (a *System) SysWriteLockedusersMountAccessorUnlockAliasIdentifier(ctx context.Context, aliasIdentifier string, mountAccessor string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/lockedusers/{mount_accessor}/unlock/{alias_identifier}"
	requestPath = strings.Replace(requestPath, "{"+"alias_identifier"+"}", url.PathEscape(aliasIdentifier), -1)
	requestPath = strings.Replace(requestPath, "{"+"mount_accessor"+"}", url.PathEscape(mountAccessor), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SysWritePluginsCatalogName Register a new plugin, or updates an existing one with the supplied name.
// name: The name of the plugin
func (a *System) SysWritePluginsCatalogName(ctx context.Context, name string, request schema.SysWritePluginsCatalogNameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SysWriteToolsRandomUrlbytes Generate random bytes
// urlbytes: The number of bytes to generate (POST URL parameter)
func (a *System) SysWriteToolsRandomUrlbytes(ctx context.Context, urlbytes string, request schema.SysWriteToolsRandomUrlbytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ToolsGenerateRandom Generate random bytes
func (a *System) ToolsGenerateRandom(ctx context.Context, request schema.ToolsGenerateRandomRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ToolsGenerateRandomSource Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
func (a *System) ToolsGenerateRandomSource(ctx context.Context, source string, request schema.ToolsGenerateRandomSourceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{source}"
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ToolsGenerateRandomSourceBytes Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// urlbytes: The number of bytes to generate (POST URL parameter)
func (a *System) ToolsGenerateRandomSourceBytes(ctx context.Context, source string, urlbytes string, request schema.ToolsGenerateRandomSourceBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{source}/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ToolsHash Generate a hash sum for input data
func (a *System) ToolsHash(ctx context.Context, request schema.ToolsHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/hash"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ToolsHashWith Generate a hash sum for input data
// urlalgorithm: Algorithm to use (POST URL parameter)
func (a *System) ToolsHashWith(ctx context.Context, urlalgorithm string, request schema.ToolsHashWithRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/hash/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Unseal Unseal the Vault.
func (a *System) Unseal(ctx context.Context, request schema.UnsealRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/unseal"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WrappingReadLookup Look up wrapping properties for the requester's token.
func (a *System) WrappingReadLookup(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/lookup"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// WrappingRewrap Rotates a response-wrapped token.
func (a *System) WrappingRewrap(ctx context.Context, request schema.WrappingRewrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/rewrap"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WrappingUnwrap Unwraps a response-wrapped token.
func (a *System) WrappingUnwrap(ctx context.Context, request schema.WrappingUnwrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/unwrap"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WrappingWrap Response-wraps an arbitrary JSON object.
func (a *System) WrappingWrap(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/wrap"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// WrappingWriteLookup Look up wrapping properties for the given token.
func (a *System) WrappingWriteLookup(ctx context.Context, request schema.WrappingWriteLookupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/lookup"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteAuditDevice Enable a new audit device at the supplied path.
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (a *System) WriteAuditDevice(ctx context.Context, path string, request schema.WriteAuditDeviceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteAuthMethod Enables a new auth method.
// After enabling, the auth method can be accessed and configured via the auth path specified as part of the URL. This auth path will be nested under the auth prefix.  For example, enable the \"foo\" auth method will make it accessible at /auth/foo.
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (a *System) WriteAuthMethod(ctx context.Context, path string, request schema.WriteAuthMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteAuthMethodTune Tune configuration parameters for a given auth path.
// This endpoint requires sudo capability on the final path, but the same functionality can be achieved without sudo via `sys/mounts/auth/[auth-path]/tune`.
// path: Tune the configuration parameters for an auth path.
func (a *System) WriteAuthMethodTune(ctx context.Context, path string, request schema.WriteAuthMethodTuneRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteCapabilities Fetches the capabilities of the given token on the given path.
func (a *System) WriteCapabilities(ctx context.Context, request schema.WriteCapabilitiesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteCapabilitiesAccessor Fetches the capabilities of the token associated with the given token, on the given path.
func (a *System) WriteCapabilitiesAccessor(ctx context.Context, request schema.WriteCapabilitiesAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities-accessor"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteCapabilitiesSelf Fetches the capabilities of the given token on the given path.
func (a *System) WriteCapabilitiesSelf(ctx context.Context, request schema.WriteCapabilitiesSelfRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities-self"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteConfigAuditingRequestHeader Enable auditing of a header.
func (a *System) WriteConfigAuditingRequestHeader(ctx context.Context, header string, request schema.WriteConfigAuditingRequestHeaderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteConfigCORS Configure the CORS settings.
func (a *System) WriteConfigCORS(ctx context.Context, request schema.WriteConfigCORSRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteConfigReloadSubsystem Reload the given subsystem
func (a *System) WriteConfigReloadSubsystem(ctx context.Context, subsystem string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/reload/{subsystem}"
	requestPath = strings.Replace(requestPath, "{"+"subsystem"+"}", url.PathEscape(subsystem), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteConfigUIHeader Configure the values to be returned for the UI header.
// header: The name of the header.
func (a *System) WriteConfigUIHeader(ctx context.Context, header string, request schema.WriteConfigUIHeaderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteGenerateRoot Initializes a new root generation attempt.
// Only a single root generation attempt can take place at a time. One (and only one) of otp or pgp_key are required.
func (a *System) WriteGenerateRoot(ctx context.Context, request schema.WriteGenerateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteGenerateRootAttempt Initializes a new root generation attempt.
// Only a single root generation attempt can take place at a time. One (and only one) of otp or pgp_key are required.
func (a *System) WriteGenerateRootAttempt(ctx context.Context, request schema.WriteGenerateRootAttemptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteGenerateRootUpdate Enter a single unseal key share to progress the root generation attempt.
// If the threshold number of unseal key shares is reached, Vault will complete the root generation and issue the new token. Otherwise, this API must be called multiple times until that threshold is met. The attempt nonce must be provided with each call.
func (a *System) WriteGenerateRootUpdate(ctx context.Context, request schema.WriteGenerateRootUpdateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/update"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteInit Initialize a new Vault.
// The Vault must not have been previously initialized. The recovery options, as well as the stored shares option, are only available when using Vault HSM.
func (a *System) WriteInit(ctx context.Context, request schema.WriteInitRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/init"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteInternalCountersConfig Enable or disable collection of client count, set retention period, or set default reporting period.
func (a *System) WriteInternalCountersConfig(ctx context.Context, request schema.WriteInternalCountersConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/config"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesLookup Retrieve lease metadata.
func (a *System) WriteLeasesLookup(ctx context.Context, request schema.WriteLeasesLookupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/lookup"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRenew Renews a lease, requesting to extend the lease.
func (a *System) WriteLeasesRenew(ctx context.Context, request schema.WriteLeasesRenewRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/renew"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRenew2 Renews a lease, requesting to extend the lease.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (a *System) WriteLeasesRenew2(ctx context.Context, urlLeaseId string, request schema.WriteLeasesRenew2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/renew/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRevoke Revokes a lease immediately.
func (a *System) WriteLeasesRevoke(ctx context.Context, request schema.WriteLeasesRevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRevoke2 Revokes a lease immediately.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (a *System) WriteLeasesRevoke2(ctx context.Context, urlLeaseId string, request schema.WriteLeasesRevoke2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRevokeForce Revokes all secrets or tokens generated under a given prefix immediately
// Unlike `/sys/leases/revoke-prefix`, this path ignores backend errors encountered during revocation. This is potentially very dangerous and should only be used in specific emergency situations where errors in the backend or the connected backend service prevent normal revocation.  By ignoring these errors, Vault abdicates responsibility for ensuring that the issued credentials or secrets are properly revoked and/or cleaned up. Access to this endpoint should be tightly controlled.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (a *System) WriteLeasesRevokeForce(ctx context.Context, prefix string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke-force/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesRevokePrefix Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (a *System) WriteLeasesRevokePrefix(ctx context.Context, prefix string, request schema.WriteLeasesRevokePrefixRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke-prefix/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLeasesTidy This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
func (a *System) WriteLeasesTidy(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/tidy"

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLogger Modify the log level of a single logger.
// name: The name of the logger to be modified.
func (a *System) WriteLogger(ctx context.Context, name string, request schema.WriteLoggerRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteLoggers Modify the log level for all existing loggers.
func (a *System) WriteLoggers(ctx context.Context, request schema.WriteLoggersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteMount Enable a new secrets engine at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (a *System) WriteMount(ctx context.Context, path string, request schema.WriteMountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteMountsConfig Tune backend configuration parameters for this mount.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (a *System) WriteMountsConfig(ctx context.Context, path string, request schema.WriteMountsConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WritePluginsCatalogByTypeByName Register a new plugin, or updates an existing one with the supplied name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (a *System) WritePluginsCatalogByTypeByName(ctx context.Context, name string, type_ string, request schema.WritePluginsCatalogByTypeByNameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WritePluginsReloadBackend Reload mounted plugin backends.
// Either the plugin name (`plugin`) or the desired plugin backend mounts (`mounts`) must be provided, but not both. In the case that the plugin name is provided, all mounted paths that use that plugin backend will be reloaded.  If (`scope`) is provided and is (`global`), the plugin(s) are reloaded globally.
func (a *System) WritePluginsReloadBackend(ctx context.Context, request schema.WritePluginsReloadBackendRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/reload/backend"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WritePoliciesACL Add a new or update an existing ACL policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) WritePoliciesACL(ctx context.Context, name string, request schema.WritePoliciesACLRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WritePoliciesPassword Add a new or update an existing password policy.
// name: The name of the password policy.
func (a *System) WritePoliciesPassword(ctx context.Context, name string, request schema.WritePoliciesPasswordRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WritePolicy Add a new or update an existing policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (a *System) WritePolicy(ctx context.Context, name string, request schema.WritePolicyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policy/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteQuotasConfig
func (a *System) WriteQuotasConfig(ctx context.Context, request schema.WriteQuotasConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/config"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteQuotasRateLimit
// name: Name of the quota rule.
func (a *System) WriteQuotasRateLimit(ctx context.Context, name string, request schema.WriteQuotasRateLimitRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRaw Update the value of the key at the given path.
func (a *System) WriteRaw(ctx context.Context, request schema.WriteRawRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRawPath Update the value of the key at the given path.
func (a *System) WriteRawPath(ctx context.Context, path string, request schema.WriteRawPathRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRekeyInit Initializes a new rekey attempt.
// Only a single rekey attempt can take place at a time, and changing the parameters of a rekey requires canceling and starting a new rekey, which will also provide a new nonce.
func (a *System) WriteRekeyInit(ctx context.Context, request schema.WriteRekeyInitRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRekeyUpdate Enter a single unseal key share to progress the rekey of the Vault.
func (a *System) WriteRekeyUpdate(ctx context.Context, request schema.WriteRekeyUpdateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/update"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRekeyVerify Enter a single new key share to progress the rekey verification operation.
func (a *System) WriteRekeyVerify(ctx context.Context, request schema.WriteRekeyVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// WriteRotateConfig
func (a *System) WriteRotateConfig(ctx context.Context, request schema.WriteRotateConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate/config"

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
