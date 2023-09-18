// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsLoginRequest struct for AwsLoginRequest
type AwsLoginRequest struct {
	// HTTP method to use for the AWS request when auth_type is iam. This must match what has been signed in the presigned request.
	IamHttpRequestMethod string `json:"iam_http_request_method,omitempty"`

	// Base64-encoded request body when auth_type is iam. This must match the request body included in the signature.
	IamRequestBody string `json:"iam_request_body,omitempty"`

	// Key/value pairs of headers for use in the sts:GetCallerIdentity HTTP requests headers when auth_type is iam. Can be either a Base64-encoded, JSON-serialized string, or a JSON object of key/value pairs. This must at a minimum include the headers over which AWS has included a signature.
	IamRequestHeaders string `json:"iam_request_headers,omitempty"`

	// Base64-encoded full URL against which to make the AWS request when using iam auth_type.
	IamRequestUrl string `json:"iam_request_url,omitempty"`

	// Base64 encoded EC2 instance identity document. This needs to be supplied along with the 'signature' parameter. If using 'curl' for fetching the identity document, consider using the option '-w 0' while piping the output to 'base64' binary.
	Identity string `json:"identity,omitempty"`

	// The nonce to be used for subsequent login requests when auth_type is ec2. If this parameter is not specified at all and if reauthentication is allowed, then the backend will generate a random nonce, attaches it to the instance's identity access list entry and returns the nonce back as part of auth metadata. This value should be used with further login requests, to establish client authenticity. Clients can choose to set a custom nonce if preferred, in which case, it is recommended that clients provide a strong nonce. If a nonce is provided but with an empty value, it indicates intent to disable reauthentication. Note that, when 'disallow_reauthentication' option is enabled on either the role or the role tag, the 'nonce' holds no significance.
	Nonce string `json:"nonce,omitempty"`

	// PKCS7 signature of the identity document when using an auth_type of ec2.
	Pkcs7 string `json:"pkcs7,omitempty"`

	// Name of the role against which the login is being attempted. If 'role' is not specified, then the login endpoint looks for a role bearing the name of the AMI ID of the EC2 instance that is trying to login. If a matching role is not found, login fails.
	Role string `json:"role,omitempty"`

	// Base64 encoded SHA256 RSA signature of the instance identity document. This needs to be supplied along with 'identity' parameter.
	Signature string `json:"signature,omitempty"`
}
