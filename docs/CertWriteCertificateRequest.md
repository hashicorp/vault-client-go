# CertWriteCertificateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCommonNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**AllowedDnsSans** | Pointer to **[]string** | A comma-separated list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedEmailSans** | Pointer to **[]string** | A comma-separated list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedMetadataExtensions** | Pointer to **[]string** | A comma-separated string or array of oid extensions. Upon successful authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates. | [optional] 
**AllowedNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in either the Common Name or SANs. Supports globbing. This parameter is deprecated, please use allowed_common_names, allowed_dns_sans, allowed_email_sans, allowed_uri_sans. | [optional] 
**AllowedOrganizationalUnits** | Pointer to **[]string** | A comma-separated list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**AllowedUriSans** | Pointer to **[]string** | A comma-separated list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**Certificate** | Pointer to **string** | The public certificate that should be trusted. Must be x509 PEM encoded. | [optional] 
**DisplayName** | Pointer to **string** | The display name to use for clients using this certificate. | [optional] 
**Lease** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 
**MaxTtl** | Pointer to **string** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**OcspCaCertificates** | Pointer to **string** | Any additional CA certificates needed to communicate with OCSP servers | [optional] 
**OcspEnabled** | Pointer to **bool** | Whether to attempt OCSP verification of certificates at login | [optional] 
**OcspFailOpen** | Pointer to **bool** | If set to true, if an OCSP revocation cannot be made successfully, login will proceed rather than failing. If false, failing to get an OCSP status fails the request. | [optional] [default to false]
**OcspQueryAllServers** | Pointer to **bool** | If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree. | [optional] [default to false]
**OcspServersOverride** | Pointer to **[]string** | A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected. | [optional] 
**Period** | Pointer to **string** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RequiredExtensions** | Pointer to **[]string** | A comma-separated string or array of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \&quot;value\&quot;. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **string** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


