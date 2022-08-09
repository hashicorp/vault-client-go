/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/hashicorp/go-retryablehttp"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// Client manages communication with the HashiCorp Vault API v1.12.0
// In most cases there should be only one, shared, Client.
type Client struct {
	configuration Configuration

	client            *http.Client
	clientWithRetries *retryablehttp.Client

	parsedBaseAddress *url.URL

	// API wrappers
	Auth     Auth
	Identity Identity
	Secrets  Secrets
	System   System
}

// NewClient returns a new Vault client with a copy of the given configuration
func NewClient(configuration Configuration) (*Client, error) {
	c := Client{
		configuration: configuration,

		// configured or default HTTP client
		client: configuration.HTTPClient,

		// retryablehttp wrapper around the HTTP client
		clientWithRetries: &retryablehttp.Client{
			HTTPClient:   configuration.HTTPClient,
			Logger:       configuration.RetryOptions.Logger,
			RetryWaitMin: configuration.RetryOptions.RetryWaitMin,
			RetryWaitMax: configuration.RetryOptions.RetryWaitMax,
			RetryMax:     configuration.RetryOptions.RetryMax,
			CheckRetry:   configuration.RetryOptions.CheckRetry,
			Backoff:      configuration.RetryOptions.Backoff,
			ErrorHandler: configuration.RetryOptions.ErrorHandler,
		},
	}

	a, err := url.Parse(configuration.BaseAddress)
	if err != nil {
		return nil, err
	}

	c.parsedBaseAddress = a

	// API wrappers
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

// NewStructuredRequest expects json.Marshaler encoded request body and returns a new request with vault-specific headers
func (c *Client) NewStructuredRequest(method, path string, body json.Marshaler) (*http.Request, error) {
	if body == nil {
		return c.NewRequest(method, path, nil)
	}

	buf := bytes.Buffer{}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return c.NewRequest(method, path, &buf)
}

// NewRequest returns a new request with vault-specific headers
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	// concatenate the base address with the given path
	url, err := c.parsedBaseAddress.Parse(path)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, url.String(), body)
}

// Do wraps http.client.Do or retryablehttp.client.Do to send an http request
func (c *Client) Do(ctx context.Context, req *http.Request, retry bool) (*http.Response, error) {
	var resp *http.Response

	req = req.WithContext(ctx)

	if retry {
		retryableReq, err := retryablehttp.FromRequest(req)
		if err != nil {
			return nil, fmt.Errorf("could not form a retryable request: %w", err)
		}

		r, err := c.clientWithRetries.Do(retryableReq)
		if err != nil {
			return nil, err
		}

		resp = r
	} else {
		r, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}

		resp = r
	}

	return resp, nil
}
