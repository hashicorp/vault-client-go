# PkiConfigureAcmeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedIssuers** | Pointer to **[]string** | which issuers are allowed for use with ACME; by default, this will only be the primary (default) issuer | [optional] [default to ["*"]]
**AllowedRoles** | Pointer to **[]string** | which roles are allowed for use with ACME; by default via &#x27;*&#x27;, these will be all roles including sign-verbatim; when concrete role names are specified, any default_directory_policy role must be included to allow usage of the default acme directories under /pki/acme/directory and /pki/issuer/:issuer_id/acme/directory. | [optional] [default to ["*"]]
**DefaultDirectoryPolicy** | Pointer to **string** | the policy to be used for non-role-qualified ACME requests; by default ACME issuance will be otherwise unrestricted, equivalent to the sign-verbatim endpoint; one may also specify a role to use as this policy, as \&quot;role:&lt;role_name&gt;\&quot;, the specified role must be allowed by allowed_roles | [optional] [default to "sign-verbatim"]
**DnsResolver** | Pointer to **string** | DNS resolver to use for domain resolution on this mount. Defaults to using the default system resolver. Must be in the format &lt;host&gt;:&lt;port&gt;, with both parts mandatory. | [optional] [default to ""]
**EabPolicy** | Pointer to **string** | Specify the policy to use for external account binding behaviour, &#x27;not-required&#x27;, &#x27;new-account-required&#x27; or &#x27;always-required&#x27; | [optional] [default to "always-required"]
**Enabled** | Pointer to **bool** | whether ACME is enabled, defaults to false meaning that clusters will by default not get ACME support | [optional] [default to false]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


