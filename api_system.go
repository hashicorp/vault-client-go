// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
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

// AuditingCalculateHash
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (s *System) AuditingCalculateHash(ctx context.Context, path string, request schema.AuditingCalculateHashRequest, options ...RequestOption) (*Response[schema.AuditingCalculateHashResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit-hash/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AuditingCalculateHashResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingDisableDevice Disable the audit device at the given path.
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (s *System) AuditingDisableDevice(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingDisableRequestHeader Disable auditing of the given request header.
func (s *System) AuditingDisableRequestHeader(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingEnableDevice Enable a new audit device at the supplied path.
// path: The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot;
func (s *System) AuditingEnableDevice(ctx context.Context, path string, request schema.AuditingEnableDeviceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingEnableRequestHeader Enable auditing of a header.
func (s *System) AuditingEnableRequestHeader(ctx context.Context, header string, request schema.AuditingEnableRequestHeaderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingListEnabledDevices List the enabled audit devices.
func (s *System) AuditingListEnabledDevices(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/audit"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingListRequestHeaders List the request headers that are configured to be audited.
func (s *System) AuditingListRequestHeaders(ctx context.Context, options ...RequestOption) (*Response[schema.AuditingListRequestHeadersResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AuditingListRequestHeadersResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuditingReadRequestHeaderInformation List the information for the given request header.
func (s *System) AuditingReadRequestHeaderInformation(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/auditing/request-headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthDisableMethod Disable the auth method at the given auth path
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (s *System) AuthDisableMethod(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthEnableMethod Enables a new auth method.
// After enabling, the auth method can be accessed and configured via the auth path specified as part of the URL. This auth path will be nested under the auth prefix.  For example, enable the \"foo\" auth method will make it accessible at /auth/foo.
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (s *System) AuthEnableMethod(ctx context.Context, path string, request schema.AuthEnableMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthListEnabledMethods
func (s *System) AuthListEnabledMethods(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthReadConfiguration Read the configuration of the auth engine at the given path.
// path: The path to mount to. Cannot be delimited. Example: \&quot;user\&quot;
func (s *System) AuthReadConfiguration(ctx context.Context, path string, options ...RequestOption) (*Response[schema.AuthReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AuthReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthReadTuningInformation Reads the given auth path's configuration.
// This endpoint requires sudo capability on the final path, but the same functionality can be achieved without sudo via `sys/mounts/auth/[auth-path]/tune`.
// path: Tune the configuration parameters for an auth path.
func (s *System) AuthReadTuningInformation(ctx context.Context, path string, options ...RequestOption) (*Response[schema.AuthReadTuningInformationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AuthReadTuningInformationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AuthTuneConfigurationParameters Tune configuration parameters for a given auth path.
// This endpoint requires sudo capability on the final path, but the same functionality can be achieved without sudo via `sys/mounts/auth/[auth-path]/tune`.
// path: Tune the configuration parameters for an auth path.
func (s *System) AuthTuneConfigurationParameters(ctx context.Context, path string, request schema.AuthTuneConfigurationParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/auth/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CollectHostInformation Information about the host instance that this Vault server is running on.
// Information about the host instance that this Vault server is running on.   The information that gets collected includes host hardware information, and CPU,   disk, and memory utilization
func (s *System) CollectHostInformation(ctx context.Context, options ...RequestOption) (*Response[schema.CollectHostInformationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/host-info"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.CollectHostInformationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CollectInFlightRequestInformation reports in-flight requests
// This path responds to the following HTTP methods.   GET /    Returns a map of in-flight requests.
func (s *System) CollectInFlightRequestInformation(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/in-flight-req"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CorsConfigure Configure the CORS settings.
func (s *System) CorsConfigure(ctx context.Context, request schema.CorsConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CorsDeleteConfiguration Remove any CORS settings.
func (s *System) CorsDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CorsReadConfiguration Return the current CORS settings.
func (s *System) CorsReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.CorsReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/cors"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.CorsReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Decode Decodes the encoded token with the otp.
func (s *System) Decode(ctx context.Context, request schema.DecodeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/decode-token"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EncryptionKeyConfigureRotation
func (s *System) EncryptionKeyConfigureRotation(ctx context.Context, request schema.EncryptionKeyConfigureRotationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EncryptionKeyReadRotationConfiguration
func (s *System) EncryptionKeyReadRotationConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.EncryptionKeyReadRotationConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.EncryptionKeyReadRotationConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EncryptionKeyRotate
func (s *System) EncryptionKeyRotate(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rotate"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EncryptionKeyStatus Provides information about the backend encryption key.
func (s *System) EncryptionKeyStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/key-status"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateHash
func (s *System) GenerateHash(ctx context.Context, request schema.GenerateHashRequest, options ...RequestOption) (*Response[schema.GenerateHashResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/hash"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateHashResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateHashWithAlgorithm
// urlalgorithm: Algorithm to use (POST URL parameter)
func (s *System) GenerateHashWithAlgorithm(ctx context.Context, urlalgorithm string, request schema.GenerateHashWithAlgorithmRequest, options ...RequestOption) (*Response[schema.GenerateHashWithAlgorithmResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/hash/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateHashWithAlgorithmResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateRandom
func (s *System) GenerateRandom(ctx context.Context, request schema.GenerateRandomRequest, options ...RequestOption) (*Response[schema.GenerateRandomResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateRandomResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateRandomWithBytes
// urlbytes: The number of bytes to generate (POST URL parameter)
func (s *System) GenerateRandomWithBytes(ctx context.Context, urlbytes string, request schema.GenerateRandomWithBytesRequest, options ...RequestOption) (*Response[schema.GenerateRandomWithBytesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateRandomWithBytesResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateRandomWithSource
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
func (s *System) GenerateRandomWithSource(ctx context.Context, source string, request schema.GenerateRandomWithSourceRequest, options ...RequestOption) (*Response[schema.GenerateRandomWithSourceResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{source}"
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateRandomWithSourceResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GenerateRandomWithSourceAndBytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// urlbytes: The number of bytes to generate (POST URL parameter)
func (s *System) GenerateRandomWithSourceAndBytes(ctx context.Context, source string, urlbytes string, request schema.GenerateRandomWithSourceAndBytesRequest, options ...RequestOption) (*Response[schema.GenerateRandomWithSourceAndBytesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/tools/random/{source}/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.GenerateRandomWithSourceAndBytesResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// HaStatus Check the HA status of a Vault cluster
func (s *System) HaStatus(ctx context.Context, options ...RequestOption) (*Response[schema.HaStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/ha-status"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.HaStatusResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Initialize Initialize a new Vault.
// The Vault must not have been previously initialized. The recovery options, as well as the stored shares option, are only available when using Vault HSM.
func (s *System) Initialize(ctx context.Context, request schema.InitializeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/init"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalClientActivityConfigure Enable or disable collection of client count, set retention period, or set default reporting period.
func (s *System) InternalClientActivityConfigure(ctx context.Context, request schema.InternalClientActivityConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalClientActivityExport Report the client count metrics, for this namespace and all child namespaces.
func (s *System) InternalClientActivityExport(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity/export"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalClientActivityReadConfiguration Read the client count tracking configuration.
func (s *System) InternalClientActivityReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalClientActivityReportCounts Report the client count metrics, for this namespace and all child namespaces.
func (s *System) InternalClientActivityReportCounts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalClientActivityReportCountsThisMonth Report the number of clients for this month, for this namespace and all child namespaces.
func (s *System) InternalClientActivityReportCountsThisMonth(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/activity/monthly"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalCountEntities Backwards compatibility is not guaranteed for this API
func (s *System) InternalCountEntities(ctx context.Context, options ...RequestOption) (*Response[schema.InternalCountEntitiesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/entities"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalCountEntitiesResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Deprecated
// InternalCountRequests Backwards compatibility is not guaranteed for this API
func (s *System) InternalCountRequests(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/requests"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalCountTokens Backwards compatibility is not guaranteed for this API
func (s *System) InternalCountTokens(ctx context.Context, options ...RequestOption) (*Response[schema.InternalCountTokensResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/counters/tokens"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalCountTokensResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalGenerateOpenApiDocument
// context: Context string appended to every operationId
// genericMountPaths: Use generic mount paths
func (s *System) InternalGenerateOpenApiDocument(ctx context.Context, context string, genericMountPaths bool, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/specs/openapi"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("context", parameterToString(context))
	requestQueryParameters.Add("generic_mount_paths", parameterToString(genericMountPaths))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalGenerateOpenApiDocumentWithParameters
func (s *System) InternalGenerateOpenApiDocumentWithParameters(ctx context.Context, request schema.InternalGenerateOpenApiDocumentWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/specs/openapi"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalInspectRouter Expose the route entry and mount entry tables present in the router
// tag: Name of subtree being observed
func (s *System) InternalInspectRouter(ctx context.Context, tag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/inspect/router/{tag}"
	requestPath = strings.Replace(requestPath, "{"+"tag"+"}", url.PathEscape(tag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalUiListEnabledFeatureFlags Lists enabled feature flags.
func (s *System) InternalUiListEnabledFeatureFlags(ctx context.Context, options ...RequestOption) (*Response[schema.InternalUiListEnabledFeatureFlagsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/feature-flags"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalUiListEnabledFeatureFlagsResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalUiListEnabledVisibleMounts Lists all enabled and visible auth and secrets mounts.
func (s *System) InternalUiListEnabledVisibleMounts(ctx context.Context, options ...RequestOption) (*Response[schema.InternalUiListEnabledVisibleMountsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/mounts"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalUiListEnabledVisibleMountsResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalUiListNamespaces Backwards compatibility is not guaranteed for this API
func (s *System) InternalUiListNamespaces(ctx context.Context, options ...RequestOption) (*Response[schema.InternalUiListNamespacesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/namespaces"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalUiListNamespacesResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalUiReadMountInformation Return information about the given mount.
// path: The path of the mount.
func (s *System) InternalUiReadMountInformation(ctx context.Context, path string, options ...RequestOption) (*Response[schema.InternalUiReadMountInformationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalUiReadMountInformationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// InternalUiReadResultantAcl Backwards compatibility is not guaranteed for this API
func (s *System) InternalUiReadResultantAcl(ctx context.Context, options ...RequestOption) (*Response[schema.InternalUiReadResultantAclResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/internal/ui/resultant-acl"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.InternalUiReadResultantAclResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeaderStatus Returns the high availability status and current leader instance of Vault.
func (s *System) LeaderStatus(ctx context.Context, options ...RequestOption) (*Response[schema.LeaderStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leader"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.LeaderStatusResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesCount
func (s *System) LeasesCount(ctx context.Context, options ...RequestOption) (*Response[schema.LeasesCountResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/count"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.LeasesCountResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesForceRevokeLeaseWithPrefix Revokes all secrets or tokens generated under a given prefix immediately
// Unlike `/sys/leases/revoke-prefix`, this path ignores backend errors encountered during revocation. This is potentially very dangerous and should only be used in specific emergency situations where errors in the backend or the connected backend service prevent normal revocation.  By ignoring these errors, Vault abdicates responsibility for ensuring that the issued credentials or secrets are properly revoked and/or cleaned up. Access to this endpoint should be tightly controlled.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (s *System) LeasesForceRevokeLeaseWithPrefix(ctx context.Context, prefix string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke-force/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesList
func (s *System) LeasesList(ctx context.Context, options ...RequestOption) (*Response[schema.LeasesListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.LeasesListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesLookUp
// prefix: The path to list leases under. Example: \&quot;aws/creds/deploy\&quot;
func (s *System) LeasesLookUp(ctx context.Context, prefix string, options ...RequestOption) (*Response[schema.LeasesLookUpResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/lookup/{prefix}/"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.LeasesLookUpResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesReadLease
func (s *System) LeasesReadLease(ctx context.Context, request schema.LeasesReadLeaseRequest, options ...RequestOption) (*Response[schema.LeasesReadLeaseResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/lookup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.LeasesReadLeaseResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesRenewLease Renews a lease, requesting to extend the lease.
func (s *System) LeasesRenewLease(ctx context.Context, request schema.LeasesRenewLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/renew"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesRenewLeaseWithId Renews a lease, requesting to extend the lease.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (s *System) LeasesRenewLeaseWithId(ctx context.Context, urlLeaseId string, request schema.LeasesRenewLeaseWithIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/renew/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesRevokeLease Revokes a lease immediately.
func (s *System) LeasesRevokeLease(ctx context.Context, request schema.LeasesRevokeLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesRevokeLeaseWithId Revokes a lease immediately.
// urlLeaseId: The lease identifier to renew. This is included with a lease.
func (s *System) LeasesRevokeLeaseWithId(ctx context.Context, urlLeaseId string, request schema.LeasesRevokeLeaseWithIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke/{url_lease_id}"
	requestPath = strings.Replace(requestPath, "{"+"url_lease_id"+"}", url.PathEscape(urlLeaseId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesRevokeLeaseWithPrefix Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.
// prefix: The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot;
func (s *System) LeasesRevokeLeaseWithPrefix(ctx context.Context, prefix string, request schema.LeasesRevokeLeaseWithPrefixRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/revoke-prefix/{prefix}"
	requestPath = strings.Replace(requestPath, "{"+"prefix"+"}", url.PathEscape(prefix), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LeasesTidy
func (s *System) LeasesTidy(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/leases/tidy"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ListExperimentalFeatures Returns the available and enabled experiments
func (s *System) ListExperimentalFeatures(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/experiments"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LockedUsersList Report the locked user count metrics, for this namespace and all child namespaces.
func (s *System) LockedUsersList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/locked-users"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LockedUsersUnlock Unlocks the user with given mount_accessor and alias_identifier
// aliasIdentifier: It is the name of the alias (user). For example, if the alias belongs to userpass backend, the name should be a valid username within userpass auth method. If the alias belongs to an approle auth method, the name should be a valid RoleID
// mountAccessor: MountAccessor is the identifier of the mount entry to which the user belongs
func (s *System) LockedUsersUnlock(ctx context.Context, aliasIdentifier string, mountAccessor string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/locked-users/{mount_accessor}/unlock/{alias_identifier}"
	requestPath = strings.Replace(requestPath, "{"+"alias_identifier"+"}", url.PathEscape(aliasIdentifier), -1)
	requestPath = strings.Replace(requestPath, "{"+"mount_accessor"+"}", url.PathEscape(mountAccessor), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersReadVerbosityLevel Read the log level for all existing loggers.
func (s *System) LoggersReadVerbosityLevel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersReadVerbosityLevelFor Read the log level for a single logger.
// name: The name of the logger to be modified.
func (s *System) LoggersReadVerbosityLevelFor(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersRevertVerbosityLevel Revert the all loggers to use log level provided in config.
func (s *System) LoggersRevertVerbosityLevel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersRevertVerbosityLevelFor Revert a single logger to use log level provided in config.
// name: The name of the logger to be modified.
func (s *System) LoggersRevertVerbosityLevelFor(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersUpdateVerbosityLevel Modify the log level for all existing loggers.
func (s *System) LoggersUpdateVerbosityLevel(ctx context.Context, request schema.LoggersUpdateVerbosityLevelRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LoggersUpdateVerbosityLevelFor Modify the log level of a single logger.
// name: The name of the logger to be modified.
func (s *System) LoggersUpdateVerbosityLevelFor(ctx context.Context, name string, request schema.LoggersUpdateVerbosityLevelForRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/loggers/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Metrics
// format: Format to export metrics into. Currently accepts only \&quot;prometheus\&quot;.
func (s *System) Metrics(ctx context.Context, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/metrics"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("format", parameterToString(format))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaValidate Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
func (s *System) MfaValidate(ctx context.Context, request schema.MfaValidateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mfa/validate"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
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
func (s *System) Monitor(ctx context.Context, logFormat string, logLevel string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/monitor"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("log_format", parameterToString(logFormat))
	requestQueryParameters.Add("log_level", parameterToString(logLevel))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsDisableSecretsEngine Disable the mount point specified at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (s *System) MountsDisableSecretsEngine(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsEnableSecretsEngine Enable a new secrets engine at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (s *System) MountsEnableSecretsEngine(ctx context.Context, path string, request schema.MountsEnableSecretsEngineRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsListSecretsEngines
func (s *System) MountsListSecretsEngines(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsReadConfiguration Read the configuration of the secret engine at the given path.
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (s *System) MountsReadConfiguration(ctx context.Context, path string, options ...RequestOption) (*Response[schema.MountsReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.MountsReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsReadTuningInformation
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (s *System) MountsReadTuningInformation(ctx context.Context, path string, options ...RequestOption) (*Response[schema.MountsReadTuningInformationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.MountsReadTuningInformationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MountsTuneConfigurationParameters
// path: The path to mount to. Example: \&quot;aws/east\&quot;
func (s *System) MountsTuneConfigurationParameters(ctx context.Context, path string, request schema.MountsTuneConfigurationParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/mounts/{path}/tune"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogListPlugins
func (s *System) PluginsCatalogListPlugins(ctx context.Context, options ...RequestOption) (*Response[schema.PluginsCatalogListPluginsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PluginsCatalogListPluginsResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogListPluginsWithType List the plugins in the catalog.
// type_: The type of the plugin, may be auth, secret, or database
func (s *System) PluginsCatalogListPluginsWithType(ctx context.Context, type_ string, options ...RequestOption) (*Response[schema.PluginsCatalogListPluginsWithTypeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/"
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.PluginsCatalogListPluginsWithTypeResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogReadPluginConfiguration Return the configuration data for the plugin with the given name.
// name: The name of the plugin
func (s *System) PluginsCatalogReadPluginConfiguration(ctx context.Context, name string, options ...RequestOption) (*Response[schema.PluginsCatalogReadPluginConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PluginsCatalogReadPluginConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogReadPluginConfigurationWithType Return the configuration data for the plugin with the given name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (s *System) PluginsCatalogReadPluginConfigurationWithType(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[schema.PluginsCatalogReadPluginConfigurationWithTypeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PluginsCatalogReadPluginConfigurationWithTypeResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogRegisterPlugin Register a new plugin, or updates an existing one with the supplied name.
// name: The name of the plugin
func (s *System) PluginsCatalogRegisterPlugin(ctx context.Context, name string, request schema.PluginsCatalogRegisterPluginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogRegisterPluginWithType Register a new plugin, or updates an existing one with the supplied name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (s *System) PluginsCatalogRegisterPluginWithType(ctx context.Context, name string, type_ string, request schema.PluginsCatalogRegisterPluginWithTypeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogRemovePlugin Remove the plugin with the given name.
// name: The name of the plugin
func (s *System) PluginsCatalogRemovePlugin(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsCatalogRemovePluginWithType Remove the plugin with the given name.
// name: The name of the plugin
// type_: The type of the plugin, may be auth, secret, or database
func (s *System) PluginsCatalogRemovePluginWithType(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsReloadBackends Reload mounted plugin backends.
// Either the plugin name (`plugin`) or the desired plugin backend mounts (`mounts`) must be provided, but not both. In the case that the plugin name is provided, all mounted paths that use that plugin backend will be reloaded.  If (`scope`) is provided and is (`global`), the plugin(s) are reloaded globally.
func (s *System) PluginsReloadBackends(ctx context.Context, request schema.PluginsReloadBackendsRequest, options ...RequestOption) (*Response[schema.PluginsReloadBackendsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/reload/backend"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.PluginsReloadBackendsResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsRuntimesCatalogListPluginsRuntimes
func (s *System) PluginsRuntimesCatalogListPluginsRuntimes(ctx context.Context, options ...RequestOption) (*Response[schema.PluginsRuntimesCatalogListPluginsRuntimesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/runtimes/catalog/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.PluginsRuntimesCatalogListPluginsRuntimesResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsRuntimesCatalogReadPluginRuntimeConfiguration Return the configuration data for the plugin runtime with the given name.
// name: The name of the plugin runtime
// type_: The type of the plugin runtime
func (s *System) PluginsRuntimesCatalogReadPluginRuntimeConfiguration(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[schema.PluginsRuntimesCatalogReadPluginRuntimeConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/runtimes/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PluginsRuntimesCatalogReadPluginRuntimeConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsRuntimesCatalogRegisterPluginRuntime Register a new plugin runtime, or updates an existing one with the supplied name.
// name: The name of the plugin runtime
// type_: The type of the plugin runtime
func (s *System) PluginsRuntimesCatalogRegisterPluginRuntime(ctx context.Context, name string, type_ string, request schema.PluginsRuntimesCatalogRegisterPluginRuntimeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/runtimes/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PluginsRuntimesCatalogRemovePluginRuntime Remove the plugin runtime with the given name.
// name: The name of the plugin runtime
// type_: The type of the plugin runtime
func (s *System) PluginsRuntimesCatalogRemovePluginRuntime(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/plugins/runtimes/catalog/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesDeleteAclPolicy Delete the ACL policy with the given name.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (s *System) PoliciesDeleteAclPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesDeletePasswordPolicy Delete a password policy.
// name: The name of the password policy.
func (s *System) PoliciesDeletePasswordPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesGeneratePasswordFromPasswordPolicy Generate a password from an existing password policy.
// name: The name of the password policy.
func (s *System) PoliciesGeneratePasswordFromPasswordPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[schema.PoliciesGeneratePasswordFromPasswordPolicyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}/generate"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PoliciesGeneratePasswordFromPasswordPolicyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesListAclPolicies
func (s *System) PoliciesListAclPolicies(ctx context.Context, options ...RequestOption) (*Response[schema.PoliciesListAclPoliciesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.PoliciesListAclPoliciesResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesListPasswordPolicies List the existing password policies.
func (s *System) PoliciesListPasswordPolicies(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesReadAclPolicy Retrieve information about the named ACL policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (s *System) PoliciesReadAclPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[schema.PoliciesReadAclPolicyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PoliciesReadAclPolicyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesReadPasswordPolicy Retrieve an existing password policy.
// name: The name of the password policy.
func (s *System) PoliciesReadPasswordPolicy(ctx context.Context, name string, options ...RequestOption) (*Response[schema.PoliciesReadPasswordPolicyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PoliciesReadPasswordPolicyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesWriteAclPolicy Add a new or update an existing ACL policy.
// name: The name of the policy. Example: \&quot;ops\&quot;
func (s *System) PoliciesWriteAclPolicy(ctx context.Context, name string, request schema.PoliciesWriteAclPolicyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/acl/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PoliciesWritePasswordPolicy Add a new or update an existing password policy.
// name: The name of the password policy.
func (s *System) PoliciesWritePasswordPolicy(ctx context.Context, name string, request schema.PoliciesWritePasswordPolicyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/policies/password/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofBlocking Returns stack traces that led to blocking on synchronization primitives
// Returns stack traces that led to blocking on synchronization primitives
func (s *System) PprofBlocking(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/block"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofCommandLine Returns the running program's command line.
// Returns the running program's command line, with arguments separated by NUL bytes.
func (s *System) PprofCommandLine(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/cmdline"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofCpuProfile Returns a pprof-formatted cpu profile payload.
// Returns a pprof-formatted cpu profile payload. Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
func (s *System) PprofCpuProfile(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/profile"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofExecutionTrace Returns the execution trace in binary form.
// Returns  the execution trace in binary form. Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
func (s *System) PprofExecutionTrace(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/trace"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofGoroutines Returns stack traces of all current goroutines.
// Returns stack traces of all current goroutines.
func (s *System) PprofGoroutines(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/goroutine"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofIndex Returns an HTML page listing the available profiles.
// Returns an HTML page listing the available  profiles. This should be mainly accessed via browsers or applications that can  render pages.
func (s *System) PprofIndex(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofMemoryAllocations Returns a sampling of all past memory allocations.
// Returns a sampling of all past memory allocations.
func (s *System) PprofMemoryAllocations(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/allocs"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofMemoryAllocationsLive Returns a sampling of memory allocations of live object.
// Returns a sampling of memory allocations of live object.
func (s *System) PprofMemoryAllocationsLive(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/heap"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofMutexes Returns stack traces of holders of contended mutexes
// Returns stack traces of holders of contended mutexes
func (s *System) PprofMutexes(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/mutex"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofSymbols Returns the program counters listed in the request.
// Returns the program counters listed in the request.
func (s *System) PprofSymbols(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/symbol"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PprofThreadCreations Returns stack traces that led to the creation of new OS threads
// Returns stack traces that led to the creation of new OS threads
func (s *System) PprofThreadCreations(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/pprof/threadcreate"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// QueryTokenAccessorCapabilities
func (s *System) QueryTokenAccessorCapabilities(ctx context.Context, request schema.QueryTokenAccessorCapabilitiesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities-accessor"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// QueryTokenCapabilities
func (s *System) QueryTokenCapabilities(ctx context.Context, request schema.QueryTokenCapabilitiesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// QueryTokenSelfCapabilities
func (s *System) QueryTokenSelfCapabilities(ctx context.Context, request schema.QueryTokenSelfCapabilitiesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/capabilities-self"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasConfigure
func (s *System) RateLimitQuotasConfigure(ctx context.Context, request schema.RateLimitQuotasConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasDelete
// name: Name of the quota rule.
func (s *System) RateLimitQuotasDelete(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasList
func (s *System) RateLimitQuotasList(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasRead
// name: Name of the quota rule.
func (s *System) RateLimitQuotasRead(ctx context.Context, name string, options ...RequestOption) (*Response[schema.RateLimitQuotasReadResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RateLimitQuotasReadResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasReadConfiguration
func (s *System) RateLimitQuotasReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.RateLimitQuotasReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RateLimitQuotasReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RateLimitQuotasWrite
// name: Name of the quota rule.
func (s *System) RateLimitQuotasWrite(ctx context.Context, name string, request schema.RateLimitQuotasWriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/quotas/rate-limit/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RawDelete Delete the key with given path.
func (s *System) RawDelete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RawList Return a list keys for a given path prefix.
func (s *System) RawList(ctx context.Context, path string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}/"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RawRead Read the value of the key at the given path.
func (s *System) RawRead(ctx context.Context, path string, options ...RequestOption) (*Response[schema.RawReadResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RawReadResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RawWrite Update the value of the key at the given path.
func (s *System) RawWrite(ctx context.Context, path string, request schema.RawWriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/raw/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadHealthStatus Returns the health status of Vault.
func (s *System) ReadHealthStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/health"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadInitializationStatus Returns the initialization status of Vault.
func (s *System) ReadInitializationStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/init"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadReplicationStatus
func (s *System) ReadReplicationStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/replication/status"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadSanitizedConfigurationState Return a sanitized version of the Vault server configuration.
// The sanitized output strips configuration values in the storage, HA storage, and seals stanzas, which may contain sensitive values such as API tokens. It also removes any token or secret fields in other stanzas, such as the circonus_api_token from telemetry.
func (s *System) ReadSanitizedConfigurationState(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/state/sanitized"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ReadWrappingProperties Look up wrapping properties for the given token.
func (s *System) ReadWrappingProperties(ctx context.Context, request schema.ReadWrappingPropertiesRequest, options ...RequestOption) (*Response[schema.ReadWrappingPropertiesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/lookup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.ReadWrappingPropertiesResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyAttemptCancel Cancels any in-progress rekey.
// This clears the rekey settings as well as any progress made. This must be called to change the parameters of the rekey. Note: verification is still a part of a rekey. If rekeying is canceled during the verification flow, the current unseal keys remain valid.
func (s *System) RekeyAttemptCancel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyAttemptInitialize Initializes a new rekey attempt.
// Only a single rekey attempt can take place at a time, and changing the parameters of a rekey requires canceling and starting a new rekey, which will also provide a new nonce.
func (s *System) RekeyAttemptInitialize(ctx context.Context, request schema.RekeyAttemptInitializeRequest, options ...RequestOption) (*Response[schema.RekeyAttemptInitializeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RekeyAttemptInitializeResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyAttemptReadProgress Reads the configuration and progress of the current rekey attempt.
func (s *System) RekeyAttemptReadProgress(ctx context.Context, options ...RequestOption) (*Response[schema.RekeyAttemptReadProgressResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/init"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RekeyAttemptReadProgressResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyAttemptUpdate Enter a single unseal key share to progress the rekey of the Vault.
func (s *System) RekeyAttemptUpdate(ctx context.Context, request schema.RekeyAttemptUpdateRequest, options ...RequestOption) (*Response[schema.RekeyAttemptUpdateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/update"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RekeyAttemptUpdateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyDeleteBackupKey Delete the backup copy of PGP-encrypted unseal keys.
func (s *System) RekeyDeleteBackupKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/backup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyDeleteBackupRecoveryKey
func (s *System) RekeyDeleteBackupRecoveryKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/recovery-key-backup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyReadBackupKey Return the backup copy of PGP-encrypted unseal keys.
func (s *System) RekeyReadBackupKey(ctx context.Context, options ...RequestOption) (*Response[schema.RekeyReadBackupKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/backup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RekeyReadBackupKeyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyReadBackupRecoveryKey
func (s *System) RekeyReadBackupRecoveryKey(ctx context.Context, options ...RequestOption) (*Response[schema.RekeyReadBackupRecoveryKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/recovery-key-backup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RekeyReadBackupRecoveryKeyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyVerificationCancel Cancel any in-progress rekey verification operation.
// This clears any progress made and resets the nonce. Unlike a `DELETE` against `sys/rekey/init`, this only resets the current verification operation, not the entire rekey atttempt.
func (s *System) RekeyVerificationCancel(ctx context.Context, options ...RequestOption) (*Response[schema.RekeyVerificationCancelResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RekeyVerificationCancelResponse](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyVerificationReadProgress Read the configuration and progress of the current rekey verification attempt.
func (s *System) RekeyVerificationReadProgress(ctx context.Context, options ...RequestOption) (*Response[schema.RekeyVerificationReadProgressResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RekeyVerificationReadProgressResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RekeyVerificationUpdate Enter a single new key share to progress the rekey verification operation.
func (s *System) RekeyVerificationUpdate(ctx context.Context, request schema.RekeyVerificationUpdateRequest, options ...RequestOption) (*Response[schema.RekeyVerificationUpdateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/rekey/verify"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RekeyVerificationUpdateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ReloadSubsystem Reload the given subsystem
func (s *System) ReloadSubsystem(ctx context.Context, subsystem string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/reload/{subsystem}"
	requestPath = strings.Replace(requestPath, "{"+"subsystem"+"}", url.PathEscape(subsystem), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Remount Initiate a mount migration
func (s *System) Remount(ctx context.Context, request schema.RemountRequest, options ...RequestOption) (*Response[schema.RemountResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/remount"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RemountResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RemountStatus Check status of a mount migration
// migrationId: The ID of the migration operation
func (s *System) RemountStatus(ctx context.Context, migrationId string, options ...RequestOption) (*Response[schema.RemountStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/remount/status/{migration_id}"
	requestPath = strings.Replace(requestPath, "{"+"migration_id"+"}", url.PathEscape(migrationId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RemountStatusResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Rewrap
func (s *System) Rewrap(ctx context.Context, request schema.RewrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/rewrap"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RootTokenGenerationCancel Cancels any in-progress root generation attempt.
func (s *System) RootTokenGenerationCancel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RootTokenGenerationInitialize Initializes a new root generation attempt.
// Only a single root generation attempt can take place at a time. One (and only one) of otp or pgp_key are required.
func (s *System) RootTokenGenerationInitialize(ctx context.Context, request schema.RootTokenGenerationInitializeRequest, options ...RequestOption) (*Response[schema.RootTokenGenerationInitializeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RootTokenGenerationInitializeResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RootTokenGenerationReadProgress Read the configuration and progress of the current root generation attempt.
func (s *System) RootTokenGenerationReadProgress(ctx context.Context, options ...RequestOption) (*Response[schema.RootTokenGenerationReadProgressResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/attempt"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.RootTokenGenerationReadProgressResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RootTokenGenerationUpdate Enter a single unseal key share to progress the root generation attempt.
// If the threshold number of unseal key shares is reached, Vault will complete the root generation and issue the new token. Otherwise, this API must be called multiple times until that threshold is met. The attempt nonce must be provided with each call.
func (s *System) RootTokenGenerationUpdate(ctx context.Context, request schema.RootTokenGenerationUpdateRequest, options ...RequestOption) (*Response[schema.RootTokenGenerationUpdateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/generate-root/update"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.RootTokenGenerationUpdateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Seal Seal the Vault.
func (s *System) Seal(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/seal"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SealStatus Check the seal status of a Vault.
func (s *System) SealStatus(ctx context.Context, options ...RequestOption) (*Response[schema.SealStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/seal-status"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.SealStatusResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// StepDownLeader Cause the node to give up active status.
// This endpoint forces the node to give up active status. If the node does not have active status, this endpoint does nothing. Note that the node will sleep for ten seconds before attempting to grab the active lock again, but if no standby nodes grab the active lock in the interim, the same node may become the active node again.
func (s *System) StepDownLeader(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/step-down"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UiHeadersConfigure Configure the values to be returned for the UI header.
// header: The name of the header.
func (s *System) UiHeadersConfigure(ctx context.Context, header string, request schema.UiHeadersConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UiHeadersDeleteConfiguration Remove a UI header.
// header: The name of the header.
func (s *System) UiHeadersDeleteConfiguration(ctx context.Context, header string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UiHeadersList Return a list of configured UI headers.
func (s *System) UiHeadersList(ctx context.Context, options ...RequestOption) (*Response[schema.UiHeadersListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.UiHeadersListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UiHeadersReadConfiguration Return the given UI header's configuration
// header: The name of the header.
func (s *System) UiHeadersReadConfiguration(ctx context.Context, header string, options ...RequestOption) (*Response[schema.UiHeadersReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/config/ui/headers/{header}"
	requestPath = strings.Replace(requestPath, "{"+"header"+"}", url.PathEscape(header), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.UiHeadersReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Unseal Unseal the Vault.
func (s *System) Unseal(ctx context.Context, request schema.UnsealRequest, options ...RequestOption) (*Response[schema.UnsealResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/unseal"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.UnsealResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// Unwrap
func (s *System) Unwrap(ctx context.Context, request schema.UnwrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/unwrap"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// VersionHistory Returns map of historical version change entries
func (s *System) VersionHistory(ctx context.Context, options ...RequestOption) (*Response[schema.VersionHistoryResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/version-history/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.VersionHistoryResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Wrap
func (s *System) Wrap(ctx context.Context, request map[string]interface{}, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/sys/wrapping/wrap"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
