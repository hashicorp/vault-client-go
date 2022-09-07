package vault

import (
	"net/http"
	"sync"
)

const (
	HeaderIndex   = "X-Vault-Index"
	HeaderForward = "X-Vault-Forward"
)

// RecordConsistencyState returns a response callback that will record the
// state returned by Vault in a response header.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func RecordConsistencyState(state *string) ResponseCallback {
	return func(req *http.Request, resp *http.Response) {
		*state = resp.Header.Get(HeaderIndex)
	}
}

// RequireConsistencyState returns a request callback that will add a request
// header to specify the state we require of Vault. This state was obtained
// from a response header seen previously and captured with
// RecordConsistencyState.
//
// https://www.vaultproject.io/docs/enterprise/consistency#conditional-forwarding-performance-standbys-only
func RequireConsistencyState(states ...string) RequestCallback {
	return func(req *http.Request) {
		for _, s := range states {
			req.Header.Add(HeaderIndex, s)
		}
	}
}

// MergeReplicationStates returns a merged array of replication states by iterating
// through all states in `old`. An iterated state is merged to the result before `new`
// based on the result of compareReplicationStates
func MergeReplicationStates(old []string, new string) []string {
	panic("not impl")
}

// replicationStateStore is used to track cluster replication states
// in order to ensure proper read-after-write semantics for a Client.
type replicationStateStore struct {
	m     sync.RWMutex
	store []string
}

// recordState updates the store's replication states with the merger of all
// states.
func (w *replicationStateStore) recordState(resp *http.Response) {
	w.m.Lock()
	defer w.m.Unlock()
	newState := resp.Header.Get(HeaderIndex)
	if newState != "" {
		w.store = MergeReplicationStates(w.store, newState)
	}
}

// requireState updates the Request with the store's current replication states.
func (w *replicationStateStore) requireState(req *http.Request) {
	w.m.RLock()
	defer w.m.RUnlock()
	for _, s := range w.store {
		req.Header.Add(HeaderIndex, s)
	}
}

// states currently stored.
func (w *replicationStateStore) states() []string {
	w.m.RLock()
	defer w.m.RUnlock()
	c := make([]string, len(w.store))
	copy(c, w.store)
	return c
}
