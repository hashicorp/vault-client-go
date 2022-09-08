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

	"github.com/hashicorp/go-secure-stdlib/strutil"
)

const (
	HeaderIndex   = "X-Vault-Index"
	HeaderForward = "X-Vault-Forward"
)

// RecordReplicationState returns a response callback that will record the
// replication state returned by Vault in a response header.
//
// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func RecordReplicationState(state *string) ResponseCallback {
	return func(req *http.Request, resp *http.Response) {
		*state = resp.Header.Get(HeaderIndex)
	}
}

// RequireReplicationStates returns a request callback that will add request
// headers to specify the replication states we require of Vault. These states
// were obtained from the previously-seen response headers captured with
// RecordReplicationState(...).
//
// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func RequireReplicationStates(states ...string) RequestCallback {
	return func(req *http.Request) {
		for _, state := range states {
			req.Header.Add(HeaderIndex, state)
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

func ParseReplicationState(raw string, hmacKey []byte) (ReplicationState, error) {
	d, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return ReplicationState{}, fmt.Errorf("could not decode replication state: %w", err)
	}
	decoded := string(d)

	lastIndex := strings.LastIndexByte(decoded, ':')
	if lastIndex == -1 {
		return ReplicationState{}, fmt.Errorf("invalid replication state header full format")
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
// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
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
type replicationStateCache []string

// recordReplicationState merges the state from the given response into the
// existing cached replication states.
//
// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (states replicationStateCache) recordReplicationState(resp *http.Response) {
	new := resp.Header.Get(HeaderIndex)
	if new != "" {
		states = MergeReplicationStates(states, new)
	}
}

// requireReplicationStates adds headers to specify the replication states we
// require of Vault. These states were obtained from the previously-seen
// response headers captured with replicationStateCache.recordReplicationState.
//
// https://www.vaultproject.io/docs/enterprise/consistency#vault-1-7-mitigations
func (states replicationStateCache) requireReplicationStates(req *http.Request) {
	for _, state := range states {
		req.Header.Add(HeaderIndex, state)
	}
}