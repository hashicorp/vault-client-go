# SshWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgorithmSigner** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512, default, or the empty string. | [optional] 
**AllowBareDomains** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use the base domains listed in \&quot;allowed_domains\&quot;, e.g. \&quot;example.com\&quot;. This is a separate option as in some cases this can be considered a security threat. | [optional] 
**AllowHostCertificates** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#x27;host&#x27;. | [optional] [default to false]
**AllowSubdomains** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use subdomains of those listed in \&quot;allowed_domains\&quot;. | [optional] 
**AllowUserCertificates** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#x27;user&#x27;. | [optional] [default to false]
**AllowUserKeyIds** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If true, users can override the key ID for a signed certificate with the \&quot;key_id\&quot; field. When false, the key ID will always be the token display name. The key ID is logged by the SSH server and can be useful for auditing. | [optional] 
**AllowedCriticalOptions** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] A comma-separated list of critical options that certificates can have when signed. To allow any critical options, set this to an empty string. | [optional] 
**AllowedDomains** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] If this option is not specified, client can request for a signed certificate for any valid host. If only certain domains are allowed, then this list enforces it. | [optional] 
**AllowedDomainsTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted. | [optional] [default to false]
**AllowedExtensions** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] A comma-separated list of extensions that certificates can have when signed. An empty list means that no extension overrides are allowed by an end-user; explicitly specify &#x27;*&#x27; to allow any extensions to be set. | [optional] 
**AllowedUserKeyLengths** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] If set, allows the enforcement of key types and minimum key sizes to be signed. | [optional] 
**AllowedUsers** | Pointer to **string** | [Optional for all types] [Works differently for CA type] If this option is not specified, or is &#x27;*&#x27;, client can request a credential for any valid user at the remote host, including the admin user. If only certain usernames are to be allowed, then this list enforces it. If this field is set, then credentials can only be created for default_user and usernames present in this list. Setting this option will enable all the users with access to this role to fetch credentials for all other usernames in this list. Use with caution. N.B.: with the CA type, an empty list means that no users are allowed; explicitly specify &#x27;*&#x27; to allow any user. | [optional] 
**AllowedUsersTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Allowed users can be specified using identity template policies. Non-templated users are also permitted. | [optional] [default to false]
**CidrList** | Pointer to **string** | [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks for which the role is applicable for. CIDR blocks can belong to more than one role. | [optional] 
**DefaultCriticalOptions** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] Critical options certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_critical_options\&quot;. Defaults to none. | [optional] 
**DefaultExtensions** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] Extensions certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_extensions\&quot;. Defaults to none. | [optional] 
**DefaultExtensionsTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Default extension values can be specified using identity template policies. Non-templated extension values are also permitted. | [optional] [default to false]
**DefaultUser** | Pointer to **string** | [Required for OTP type] [Optional for CA type] Default username for which a credential will be generated. When the endpoint &#x27;creds/&#x27; is used without a username, this value will be used as default username. | [optional] 
**DefaultUserTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Default user can be specified using identity template policies. Non-templated users are also permitted. | [optional] [default to false]
**ExcludeCidrList** | Pointer to **string** | [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks. IP addresses belonging to these blocks are not accepted by the role. This is particularly useful when big CIDR blocks are being used by the role and certain parts of it needs to be kept out. | [optional] 
**KeyIdFormat** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a custom format for the key id of a signed certificate. The following variables are available for use: &#x27;{{token_display_name}}&#x27; - The display name of the token used to make the request. &#x27;{{role_name}}&#x27; - The name of the role signing the request. &#x27;{{public_key_hash}}&#x27; - A SHA256 checksum of the public key that is being signed. | [optional] 
**KeyType** | Pointer to **string** | [Required for all types] Type of key used to login to hosts. It can be either &#x27;otp&#x27; or &#x27;ca&#x27;. &#x27;otp&#x27; type requires agent to be installed in remote hosts. | [optional] 
**MaxTtl** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The maximum allowed lease duration | [optional] 
**NotBeforeDuration** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The duration that the SSH certificate should be backdated by at issuance. | [optional] [default to "30"]
**Port** | Pointer to **int32** | [Optional for OTP type] [Not applicable for CA type] Port number for SSH connection. Default is &#x27;22&#x27;. Port number does not play any role in creation of OTP. For &#x27;otp&#x27; type, this is just a way to inform client about the port number to use. Port number will be returned to client by Vault server along with OTP. | [optional] 
**Ttl** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The lease duration if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the value of max_ttl. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


