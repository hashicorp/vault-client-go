// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

const ClientVersion = "0.4.3"

// Client manages communication with Vault, initialize it with vault.New(...)
type Client struct {
	// the configuration object is immutable after the client has been initialized
	configuration ClientConfiguration

	parsedBaseAddress *url.URL

	client            *http.Client
	clientWithRetries *retryablehttp.Client

	// headers & callbacks that will be applied to each request
	clientRequestModifiers     requestModifiers
	clientRequestModifiersLock sync.RWMutex

	// replication state cache used to ensure read-after-write semantics
	replicationStates replicationStateCache

	// generated request methods
	Auth     Auth
	Identity Identity
	Secrets  Secrets
	System   System
}

// New returns a new client decorated with the given configuration options
func New(options ...ClientOption) (*Client, error) {
	configuration := DefaultConfiguration()

	for _, option := range options {
		if err := option(&configuration); err != nil {
			return nil, err
		}
	}

	return newClient(configuration)
}

// newClient returns a new Vault client with a copy of the given configuration
func newClient(configuration ClientConfiguration) (*Client, error) {
	c := Client{
		configuration: configuration,

		// configured or default HTTP client
		client: configuration.HTTPClient,

		// retryablehttp wrapper around the HTTP client
		clientWithRetries: &retryablehttp.Client{
			HTTPClient:   configuration.HTTPClient,
			Logger:       configuration.RetryConfiguration.Logger,
			RetryWaitMin: configuration.RetryConfiguration.RetryWaitMin,
			RetryWaitMax: configuration.RetryConfiguration.RetryWaitMax,
			RetryMax:     configuration.RetryConfiguration.RetryMax,
			CheckRetry:   configuration.RetryConfiguration.CheckRetry,
			Backoff:      configuration.RetryConfiguration.Backoff,
			ErrorHandler: configuration.RetryConfiguration.ErrorHandler,
		},

		clientRequestModifiers: requestModifiers{
			headers: requestHeaders{
				userAgent:                 constructUserAgent(ClientVersion),
				token:                     configuration.initialToken,
				namespace:                 configuration.initialNamespace,
				replicationForwardingMode: ReplicationForwardNone,
			},
		},
		clientRequestModifiersLock: sync.RWMutex{},
	}

	address, err := parseAddress(configuration.Address)
	if err != nil {
		return nil, err
	}
	c.parsedBaseAddress = address

	if strings.HasPrefix(configuration.Address, "unix://") {
		// Adjust the dial context for unix domain socket addresses in the
		// internal HTTP transport, if exists.
		if httpTransport, ok := c.client.Transport.(*http.Transport); ok {
			httpTransport.DialContext = func(context.Context, string, string) (net.Conn, error) {
				return net.Dial("unix", strings.TrimPrefix(configuration.Address, "unix://"))
			}
		} else {
			return nil, fmt.Errorf(
				"the configured base client's transport (%T) is not of type *http.Transport and cannot be used with the unix:// address",
				c.client.Transport,
			)
		}
	}

	if !configuration.TLS.empty() {
		// Apply TLS configuration to the internal HTTP transport, if exists.
		if httpTransport, ok := c.client.Transport.(*http.Transport); ok {
			if err := configuration.TLS.applyTo(httpTransport.TLSClientConfig); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf(
				"the configured base client's transport (%T) is not of type *http.Transport and cannot be used with TLS configuration",
				c.client.Transport,
			)
		}
	}

	c.Auth = Auth{
		client: &c,
	}

	c.Identity = Identity{
		client: &c,
	}

	c.Secrets = Secrets{
		client: &c,
	}

	c.System = System{
		client: &c,
	}

	return &c, nil
}

// parseAddress parses the given address string with special handling for unix
// domain sockets
func parseAddress(address string) (*url.URL, error) {
	parsed, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(address, "unix://") {
		// The address in the client is expected to be pointing to the protocol
		// used in the application layer and not to the transport layer. Hence,
		// setting the fields accordingly.
		parsed.Scheme = "http"
		parsed.Host = strings.TrimPrefix(address, "unix://") // socket
		parsed.Path = ""
	}

	return parsed, nil
}

// Clone creates a new client with the same configuration, request modifiers,
// and replication states as the original client. Note that the cloned client
// will point to the same base http.Client and retryablehttp.Client objects.
func (c *Client) Clone() *Client {
	clone := Client{
		configuration:     c.configuration,
		parsedBaseAddress: c.parsedBaseAddress,
		client:            c.client,
		clientWithRetries: c.clientWithRetries,
	}

	if c.configuration.EnforceReadYourWritesConsistency {
		clone.replicationStates = c.replicationStates.clone()
	}

	clone.clientRequestModifiers = c.cloneClientRequestModifiers()

	clone.Auth = Auth{
		client: &clone,
	}

	clone.Identity = Identity{
		client: &clone,
	}

	clone.Secrets = Secrets{
		client: &clone,
	}

	clone.System = System{
		client: &clone,
	}

	return &clone
}

// cloneClientRequestModifiers returns a copy of the modifiers behind a mutex
func (c *Client) cloneClientRequestModifiers() requestModifiers {
	/* */ c.clientRequestModifiersLock.RLock()
	defer c.clientRequestModifiersLock.RUnlock()

	var clone requestModifiers

	clone.requestCallbacks = slices.Clone(c.clientRequestModifiers.requestCallbacks)
	clone.responseCallbacks = slices.Clone(c.clientRequestModifiers.responseCallbacks)

	clone.headers = requestHeaders{
		userAgent:                 c.clientRequestModifiers.headers.userAgent,
		token:                     c.clientRequestModifiers.headers.token,
		namespace:                 c.clientRequestModifiers.headers.namespace,
		responseWrappingTTL:       c.clientRequestModifiers.headers.responseWrappingTTL,
		replicationForwardingMode: c.clientRequestModifiers.headers.replicationForwardingMode,
		customHeaders:             c.clientRequestModifiers.headers.customHeaders.Clone(),
	}

	clone.headers.mfaCredentials = slices.Clone(c.clientRequestModifiers.headers.mfaCredentials)

	return clone
}

// Configuration returns a copy of the configuration object used to initialize
// this client
func (c *Client) Configuration() ClientConfiguration {
	return c.configuration
}
