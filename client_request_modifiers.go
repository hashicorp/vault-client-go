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

// SetToken sets the token to be used with all subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) SetToken(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.token = token
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearToken clears the token for all subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) ClearToken() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.token = ""
	c.requestModifiersLock.Unlock()
}

// WithToken returns a shallow copy of the client with the token set to the
// given value.
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) WithToken(token string) *Client {
	clone := c.Clone()

	if err := validateToken(token); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.token = token
	}

	return clone
}

// SetNamespace sets the namespace to be used with all subsequent requests,
// set to "" to clear the namespace.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) SetNamespace(namespace string) error {
	if err := validateNamespace(namespace); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.namespace = namespace
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearNamespace clears the namespace from all subsequent requests.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) ClearNamespace() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.namespace = ""
	c.requestModifiersLock.Unlock()
}

// WithNamespace returns a shallow copy of the client with the namespace set to
// the given value, use "" to clear the namespace.
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) WithNamespace(namespace string) *Client {
	clone := c.Clone()

	if err := validateNamespace(namespace); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.namespace = namespace
	}

	return clone
}

// SetMFACredentials sets multi-factor authentication credentials to be used
// with all subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) SetMFACredentials(credentials ...string) error {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.mfaCredentials = credentials
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearMFACredentials clears multi-factor authentication credentials from all
// subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) ClearMFACredentials() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.mfaCredentials = nil
	c.requestModifiersLock.Unlock()
}

// WithMFACredentials returns a shallow copy of the client with multi-factor
// authentication credentials set to the given value(s); passing no parameters
// is equivalent to clearing the list.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) WithMFACredentials(credentials ...string) *Client {
	clone := c.Clone()
	clone.requestModifiers.headers.mfaCredentials = credentials

	return clone
}

// SetResponseWrapping sets the response-wrapping TTL to the given duration for
// all subsequent requests, telling Vault to wrap responses and return
// response-wrapping tokens instead.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) SetResponseWrapping(ttl time.Duration) error {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.responseWrappingTTL = ttl
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearResponseWrapping clears the response-wrapping header from all
// subsequent requests.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) ClearResponseWrapping() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.responseWrappingTTL = 0
	c.requestModifiersLock.Unlock()
}

// WithResponseWrapping returns a shallow copy of the client with the
// response-wrapping TTL set to the given duration. A non-zero duration will
// tell Vault to wrap the response and return a response-wrapping token
// instead. Set `ttl` to zero to clear the response-wrapping header.
//
// See https://www.vaultproject.io/docs/concepts/response-wrapping for more
// information on response wrapping.
func (c *Client) WithResponseWrapping(ttl time.Duration) *Client {
	clone := c.Clone()
	clone.requestModifiers.headers.responseWrappingTTL = ttl

	return clone
}

// SetCustomHeaders sets custom headers to be used in all subsequent requests.
// The internal prefix 'X-Vault-' is not permitted for the header keys.
func (c *Client) SetCustomHeaders(headers http.Header) error {
	if err := validateCustomHeaders(headers); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.customHeaders = headers.Clone()
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearsCustomHeaders clears all custom headers from the subsequent requests.
func (c *Client) ClearCustomHeaders() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.customHeaders = nil
	c.requestModifiersLock.Unlock()
}

// WithCustomHeaders returns a shallow copy of the client with custom headers
// set to the given value, use nil to clear out the headers.
func (c *Client) WithCustomHeaders(headers http.Header) *Client {
	clone := c.Clone()

	if err := validateCustomHeaders(headers); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.customHeaders = headers.Clone()
	}

	return clone
}

// SetRequestCallbacks sets callbacks which will be invoked before each request.
func (c *Client) SetRequestCallbacks(callbacks ...RequestCallback) error {
	c.requestModifiersLock.Lock()
	copy(c.requestModifiers.requestCallbacks, callbacks)
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearRequestCallbacks clears all request callbacks.
func (c *Client) ClearRequestCallbacks() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.requestCallbacks = nil
	c.requestModifiersLock.Unlock()
}

// WithRequestCallbacks returns a shallow copy of the client with request
// callbacks set to the given value, use nil to clear out the callbacks.
func (c *Client) WithRequestCallbacks(callbacks ...RequestCallback) *Client {
	clone := c.Clone()
	copy(clone.requestModifiers.requestCallbacks, callbacks)

	return clone
}

// SetResponseCallbacks sets callbacks which will be invoked after each
// successful response.
func (c *Client) SetResponseCallbacks(callbacks ...ResponseCallback) error {
	c.requestModifiersLock.Lock()
	copy(c.requestModifiers.responseCallbacks, callbacks)
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearResponseCallbacks clears all response callbacks.
func (c *Client) ClearResponseCallbacks() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.responseCallbacks = nil
	c.requestModifiersLock.Unlock()
}

// WithResponseCallbacks returns a shallow copy of the client with response
// callbacks set to the given value, use nil to clear out the callbacks.
func (c *Client) WithResponseCallbacks(callbacks ...ResponseCallback) *Client {
	clone := c.Clone()
	copy(clone.requestModifiers.responseCallbacks, callbacks)

	return clone
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

// SetReplicationForwardingMode will add a forwarding header to all subsequent
// requests:
//
//	ReplicationForwardNone         - no forwarding headers
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// Note: this feature must be enabled in Vault's configuration.
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) SetReplicationForwardingMode(mode ReplicationForwardingMode) {
	/* */ c.requestModifiersLock.Lock()
	defer c.requestModifiersLock.Unlock()

	c.requestModifiers.headers.replicationForwardingMode = mode
}

// ReplicationForwardingMode clears the X-Vault-Forward / X-Vault-Inconsistent
// headers from all subsequent requests.
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) ClearReplicationForwardingMode() {
	/* */ c.requestModifiersLock.Lock()
	defer c.requestModifiersLock.Unlock()

	c.requestModifiers.headers.replicationForwardingMode = ReplicationForwardNone
}

// WithReplicationForwardingMode returns a shallow copy of the client with
// a replication header set to the given value for subsequent requests:
//
//	ReplicationForwardNone         - no forwarding headers
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// See https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) WithReplicationForwardingMode(mode ReplicationForwardingMode) *Client {
	clone := c.Clone()
	clone.requestModifiers.headers.replicationForwardingMode = mode

	return clone
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
