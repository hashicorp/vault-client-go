package vault

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-secure-stdlib/strutil"
)

// RequestOption is a functional parameter used to modify a request
type RequestOption func(*requestModifiers) error

type (
	RequestCallback  func(*http.Request)
	ResponseCallback func(*http.Request, *http.Response)
)

// requestModifiers contains headers, callbacks, etc. that will be added to
// each request
type requestModifiers struct {
	headers requestHeaders

	requestCallbacks  []RequestCallback
	responseCallbacks []ResponseCallback

	// mountPath, if specified, will overwrite the default mount path used in
	// client.Auth & client.Secrets requests
	mountPath string
}

// requestHeaders contains headers that will be added to requests
type requestHeaders struct {
	userAgent                 string                    // 'User-Agent'
	token                     string                    // 'X-Vault-Token'
	namespace                 string                    // 'X-Vault-Namespace'
	mfaCredentials            []string                  // 'X-Vault-MFA'
	responseWrappingTTL       time.Duration             // 'X-Vault-Wrap-TTL'
	replicationForwardingMode ReplicationForwardingMode // 'X-Vault-Forward' or 'X-Vault-Inconsistent'
	customHeaders             http.Header
}

// SetToken sets the token to be used with all subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on tokens.
func (c *Client) SetToken(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.token = token
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearToken clears the token for all subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on tokens.
func (c *Client) ClearToken() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.token = ""
	c.clientRequestModifiersLock.Unlock()
}

// WithToken sets the token for the next request; it takes precedence over the
// client-level token.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on tokens.
func WithToken(token string) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateToken(token); err != nil {
			return err
		}
		m.headers.token = token
		return nil
	}
}

// SetNamespace sets the namespace to be used with all subsequent requests.
// Use an empty string to clear the namespace.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) SetNamespace(namespace string) error {
	if err := validateNamespace(namespace); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.namespace = namespace
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearNamespace clears the namespace from all subsequent requests.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) ClearNamespace() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.namespace = ""
	c.clientRequestModifiersLock.Unlock()
}

// WithNamespace sets the namespace for the next request; it takes precedence
// over the client-level namespace. Use an empty string to clear the namespace
// from the next request.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) WithNamespace(namespace string) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateNamespace(namespace); err != nil {
			return err
		}
		m.headers.namespace = namespace
		return nil
	}
}

// SetMFACredentials sets multi-factor authentication credentials to be used
// with all subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) SetMFACredentials(credentials ...string) error {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.mfaCredentials = credentials
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearMFACredentials clears multi-factor authentication credentials from all
// subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) ClearMFACredentials() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.mfaCredentials = nil
	c.clientRequestModifiersLock.Unlock()
}

// WithMFACredentials sets the multi-factor authentication credentials for the
// next request, it takes precedence over the client-level MFA credentials.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) WithMFACredentials(credentials ...string) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.mfaCredentials = credentials
		return nil
	}
}

// SetResponseWrapping sets the response-wrapping TTL to the given duration for
// all subsequent requests, telling Vault to wrap responses and return
// response-wrapping tokens instead.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) SetResponseWrapping(ttl time.Duration) error {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.responseWrappingTTL = ttl
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearResponseWrapping clears the response-wrapping header from all
// subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) ClearResponseWrapping() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.responseWrappingTTL = 0
	c.clientRequestModifiersLock.Unlock()
}

// WithResponseWrapping sets the response-wrapping TTL to the given duration
// for the next request; it takes precedence over the client-level
// response-wrapping TTL. A non-zero duration will tell Vault to wrap the
// response and return a response-wrapping token instead. Set `ttl` to zero
// to clear the response-wrapping header from the next request.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) WithResponseWrapping(ttl time.Duration) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.responseWrappingTTL = ttl
		return nil
	}
}

// SetCustomHeaders sets custom headers to be used in all subsequent requests.
// The internal prefix 'X-Vault-' is not permitted for the header keys.
func (c *Client) SetCustomHeaders(headers http.Header) error {
	if err := validateCustomHeaders(headers); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.customHeaders = headers.Clone()
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearsCustomHeaders clears all custom headers from the subsequent requests.
func (c *Client) ClearCustomHeaders() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.customHeaders = nil
	c.clientRequestModifiersLock.Unlock()
}

// WithCustomHeaders sets custom headers for the next request; these headers
// take precedence over the client-level custom headers. The internal prefix
// 'X-Vault-' is not permitted for the header keys.
func (c *Client) WithCustomHeaders(headers http.Header) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateCustomHeaders(headers); err != nil {
			return err
		}
		m.headers.customHeaders = headers
		return nil
	}
}

// SetRequestCallbacks sets callbacks which will be invoked before each request.
func (c *Client) SetRequestCallbacks(callbacks ...RequestCallback) error {
	c.clientRequestModifiersLock.Lock()
	copy(c.clientRequestModifiers.requestCallbacks, callbacks)
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearRequestCallbacks clears all request callbacks.
func (c *Client) ClearRequestCallbacks() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.requestCallbacks = nil
	c.clientRequestModifiersLock.Unlock()
}

// WithRequestCallbacks sets callbacks which will be invoked before the next
// request; these take precedence over the client-level request callbacks.
func (c *Client) WithRequestCallbacks(callbacks ...RequestCallback) RequestOption {
	return func(m *requestModifiers) error {
		m.requestCallbacks = callbacks
		return nil
	}
}

// SetResponseCallbacks sets callbacks which will be invoked after each
// successful response.
func (c *Client) SetResponseCallbacks(callbacks ...ResponseCallback) error {
	c.clientRequestModifiersLock.Lock()
	copy(c.clientRequestModifiers.responseCallbacks, callbacks)
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearResponseCallbacks clears all response callbacks.
func (c *Client) ClearResponseCallbacks() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.responseCallbacks = nil
	c.clientRequestModifiersLock.Unlock()
}

// WithResponseCallbacks sets callbacks which will be invoked after a
// successful response within the next request; these take precedence over the
// client-level response callbacks.
func (c *Client) WithResponseCallbacks(callbacks ...ResponseCallback) RequestOption {
	return func(m *requestModifiers) error {
		m.responseCallbacks = callbacks
		return nil
	}
}

// WithMountPath overwrites the default mount path in client.Auth and
// client.Secrets requests with the given mount path string.
func (c *Client) WithMountPath(path string) RequestOption {
	return func(m *requestModifiers) error {
		path = strings.TrimPrefix(path, "/")
		path = strings.TrimSuffix(path, "/")
		m.mountPath = path
		return nil
	}
}

func (m *requestModifiers) mountPathOr(defaultMountPath string) string {
	if m.mountPath == "" {
		return defaultMountPath
	}
	return m.mountPath
}

// mergeRequestModifiers merges the two objects, preferring the local modifiers
func mergeRequestModifiers(global, local requestModifiers) requestModifiers {
	merged := global

	if local.headers.userAgent != "" {
		merged.headers.userAgent = local.headers.userAgent
	}

	if local.headers.token != "" {
		merged.headers.token = local.headers.token
	}

	if local.headers.namespace != "" {
		merged.headers.namespace = local.headers.namespace
	}

	if len(local.headers.mfaCredentials) != 0 {
		merged.headers.mfaCredentials = local.headers.mfaCredentials
	}

	if local.headers.responseWrappingTTL != 0 {
		merged.headers.responseWrappingTTL = local.headers.responseWrappingTTL
	}

	if local.headers.replicationForwardingMode != ReplicationForwardNone {
		merged.headers.replicationForwardingMode = local.headers.replicationForwardingMode
	}

	if len(local.headers.customHeaders) != 0 {
		merged.headers.customHeaders = local.headers.customHeaders
	}

	if len(local.requestCallbacks) != 0 {
		merged.requestCallbacks = local.requestCallbacks
	}

	if len(local.responseCallbacks) != 0 {
		merged.responseCallbacks = local.responseCallbacks
	}

	return merged
}

// requestOptionsToRequestModifiers constructs `requestModifiers` & propagates the errors
func requestOptionsToRequestModifiers(options []RequestOption) (_ requestModifiers, errs error) {
	var modifiers requestModifiers

	for _, option := range options {
		if err := option(&modifiers); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	if errs != nil {
		return requestModifiers{}, errs
	}

	return modifiers, nil
}

func validateToken(token string) error {
	if !printable(token) {
		return fmt.Errorf("vault token contains non-printable characters")
	}

	return nil
}

func validateNamespace(namespace string) error {
	if !printable(namespace) {
		return fmt.Errorf("vault namespace %q contains non-printable characters", namespace)
	}

	return nil
}

func validateCustomHeaders(headers http.Header) (errs error) {
	for key := range headers {
		if strings.HasPrefix(strings.ToLower(key), "x-vault-") {
			errs = multierror.Append(
				errs,
				fmt.Errorf("custom header key %q is not allowed: 'X-Vault-' prefix is for internal use only", key),
			)
		}
	}

	return errs
}

// printable returns true if the given string has no non-printable characters
func printable(str string) bool {
	idx := strings.IndexFunc(str, func(c rune) bool {
		return !unicode.IsPrint(c)
	})

	return idx == -1
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////// REPLICATION CONSISTENCY ///////////////////////////
///////////////////////////////////////////////////////////////////////////////

type ReplicationForwardingMode uint8

const (
	// Setting this mode will clear all forwarding headers
	ReplicationForwardNone ReplicationForwardingMode = iota

	// Setting this mode will add 'X-Vault-Forward' header to all subsequent
	// requests, telling any performance standbys handling the requests to
	// forward them to the active node.
	//
	// https://www.vaultproject.io/docs/enterprise/consistency#unconditional-forwarding-performance-standbys-only
	ReplicationForwardAlways

	// Setting this mode will add 'X-Vault-Inconsistent' header to  all
	// subsequent requests; any performance standbys handling the requests will
	// conditionally forward them to the active node if the state required
	// isn't present on the node receiving this request. This should be used
	// in conjunction with RequireReplicationState(...).
	//
	// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
	ReplicationForwardInconsistent
)

// SetReplicationForwardingMode sets a replication forwarding header for all
// subsequent requests:
//
//	ReplicationForwardNone         - no forwarding header
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// Note: this feature must be enabled in Vault's configuration.
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) SetReplicationForwardingMode(mode ReplicationForwardingMode) {
	/* */ c.clientRequestModifiersLock.Lock()
	defer c.clientRequestModifiersLock.Unlock()

	c.clientRequestModifiers.headers.replicationForwardingMode = mode
}

// ReplicationForwardingMode clears the X-Vault-Forward / X-Vault-Inconsistent
// headers from all subsequent requests.
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) ClearReplicationForwardingMode() {
	/* */ c.clientRequestModifiersLock.Lock()
	defer c.clientRequestModifiersLock.Unlock()

	c.clientRequestModifiers.headers.replicationForwardingMode = ReplicationForwardNone
}

// WithReplicationForwardingMode sets a replication forwarding header to the
// given value for the next request; it takes precedence over the client-level
// replication forwarding header.
//
//	ReplicationForwardNone         - no forwarding headers
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) WithReplicationForwardingMode(mode ReplicationForwardingMode) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.replicationForwardingMode = mode
		return nil
	}
}

// RecordReplicationState returns a response callback that will record the
// replication state returned by Vault in a response header.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func RecordReplicationState(state *string) ResponseCallback {
	return func(req *http.Request, resp *http.Response) {
		*state = resp.Header.Get("X-Vault-Index")
	}
}

// RequireReplicationStates returns a request callback that will add request
// headers to specify the replication states we require of Vault. These states
// were obtained from the previously-seen response headers captured with
// RecordReplicationState(...).
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func RequireReplicationStates(states ...string) RequestCallback {
	return func(req *http.Request) {
		for _, state := range states {
			req.Header.Add("X-Vault-Index", state)
		}
	}
}

// MergeReplicationStates returns a merged array of replication states by
// iterating through all states in the `old` slice. An iterated state is merged
// into the result before the `new` based on the result of
// compareReplicationStates
func MergeReplicationStates(old []string, new string) []string {
	if len(old) == 0 || len(old) > 2 {
		return []string{new}
	}

	var merged []string

	for _, o := range old {
		c, err := compareReplicationStates(o, new)
		if err != nil {
			return []string{new}
		}

		switch c {
		case 1:
			merged = append(merged, o)
		case -1:
			merged = append(merged, new)
		case 0:
			merged = append(merged, o, new)
		}
	}

	return strutil.RemoveDuplicates(merged, false)
}

// ReplicationState is analogous to the WALState in github.com/vault/sdk
type ReplicationState struct {
	Cluster         string
	LocalIndex      uint64
	ReplicatedIndex uint64
}

// ParseReplicationState will parse the raw base64-encoded replication state
// into its individual components. If an optional hmacKey is provided, it will
// used to verify the replication state contents. The format of the string
// (after decoding) is expected to be:
//
//	v1:cluster-id-string:local-index:replicated-index:hmac
func ParseReplicationState(raw string, hmacKey []byte) (ReplicationState, error) {
	d, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return ReplicationState{}, fmt.Errorf("could not decode replication state: %w", err)
	}
	decoded := string(d)

	lastIndex := strings.LastIndexByte(decoded, ':')
	if lastIndex == -1 {
		return ReplicationState{}, fmt.Errorf("invalid replication state header format")
	}

	state := decoded[:lastIndex]
	stateHMACRaw := decoded[lastIndex+1:]
	stateHMAC, err := hex.DecodeString(stateHMACRaw)
	if err != nil {
		return ReplicationState{}, fmt.Errorf("invalid replication state header HMAC: %s, %w", stateHMACRaw, err)
	}

	if len(hmacKey) != 0 {
		hm := hmac.New(sha256.New, hmacKey)
		hm.Write([]byte(state))
		if !hmac.Equal(hm.Sum(nil), stateHMAC) {
			return ReplicationState{}, fmt.Errorf("invalid replication state header HMAC (mismatch)")
		}
	}

	pieces := strings.Split(state, ":")
	if len(pieces) != 4 || pieces[0] != "v1" || pieces[1] == "" {
		return ReplicationState{}, fmt.Errorf("invalid replication state header format")
	}

	localIndex, err := strconv.ParseUint(pieces[2], 10, 64)
	if err != nil {
		return ReplicationState{}, fmt.Errorf("invalid local index in replication state header: %w", err)
	}

	replicatedIndex, err := strconv.ParseUint(pieces[3], 10, 64)
	if err != nil {
		return ReplicationState{}, fmt.Errorf("invalid replicated index in replication state header: %w", err)
	}

	return ReplicationState{
		Cluster:         pieces[1],
		LocalIndex:      localIndex,
		ReplicatedIndex: replicatedIndex,
	}, nil
}

// compareReplicationStates returns 1 if s1 is newer or identical, -1 if s1 is
// older, and 0 if neither s1 nor s2 is strictly greater. An error is returned
// if s1 or s2 are invalid or from different clusters.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func compareReplicationStates(s1, s2 string) (int, error) {
	r1, err := ParseReplicationState(s1, nil)
	if err != nil {
		return 0, err
	}

	r2, err := ParseReplicationState(s2, nil)
	if err != nil {
		return 0, err
	}

	if r1.Cluster != r2.Cluster {
		return 0, fmt.Errorf("can't compare replication states from different clusters")
	}

	switch {
	case r1.LocalIndex >= r2.LocalIndex && r1.ReplicatedIndex >= r2.ReplicatedIndex:
		return 1, nil
	// We've already handled the case where both are equal above, so really we're
	// asking here if one or both are lesser.
	case r1.LocalIndex <= r2.LocalIndex && r1.ReplicatedIndex <= r2.ReplicatedIndex:
		return -1, nil
	}

	return 0, nil
}

// replicationStateCache is used to track cluster replication states in order
// to ensure proper read-after-write semantics for the client. This will have
// at most two states due to how MergeReplicationStates works.
type replicationStateCache struct {
	states     []string
	statesLock sync.RWMutex
}

// recordReplicationState merges the state from the given response into the
// existing cached replication states.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func (c *replicationStateCache) recordReplicationState(resp *http.Response) {
	/* */ c.statesLock.Lock()
	defer c.statesLock.Unlock()

	if new := resp.Header.Get("X-Vault-Index"); new != "" {
		c.states = MergeReplicationStates(c.states, new)
	}
}

// requireReplicationStates adds headers to specify the replication states we
// require of Vault. These states were obtained from the previously-seen
// response headers captured with replicationStateCache.recordReplicationState.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func (c *replicationStateCache) requireReplicationStates(req *http.Request) {
	/* */ c.statesLock.RLock()
	defer c.statesLock.RUnlock()

	for _, state := range c.states {
		req.Header.Add("X-Vault-Index", state)
	}
}

// clone returns a deep copy of the replication state cache
func (c *replicationStateCache) clone() replicationStateCache {
	/* */ c.statesLock.RLock()
	defer c.statesLock.RUnlock()

	var cloned []string
	copy(cloned, c.states)

	return replicationStateCache{
		statesLock: sync.RWMutex{},
		states:     cloned,
	}
}
