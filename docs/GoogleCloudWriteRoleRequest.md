# GoogleCloudWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddGroupAliases** | Pointer to **bool** | If true, will add group aliases to auth tokens generated under this role. This will add the full list of ancestors (projects, folders, organizations) for the given entity&#x27;s project. Requires IAM permission &#x60;resourcemanager.projects.get&#x60; on this project. | [optional] [default to false]
**AllowGceInference** | Pointer to **bool** | &#x27;iam&#x27; roles only. If false, Vault will not not allow GCE instances to login in against this role | [optional] [default to true]
**BoundInstanceGroup** | Pointer to **string** | Deprecated: use \&quot;bound_instance_groups\&quot; instead. | [optional] 
**BoundInstanceGroups** | Pointer to **[]string** | Comma-separated list of permitted instance groups to which the GCE instance must belong. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundLabels** | Pointer to **[]string** | Comma-separated list of GCP labels formatted as\&quot;key:value\&quot; strings that must be present on the GCE instance in order to authenticate. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundProjects** | Pointer to **[]string** | GCP Projects that authenticating entities must belong to. | [optional] 
**BoundRegion** | Pointer to **string** | Deprecated: use \&quot;bound_regions\&quot; instead. | [optional] 
**BoundRegions** | Pointer to **[]string** | Comma-separated list of permitted regions to which the GCE instance must belong. If a group is provided, it is assumed to be a regional group. If \&quot;zone\&quot; is provided, this option is ignored. This can be a self-link or region name. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundServiceAccounts** | Pointer to **[]string** | Can be set for both &#x27;iam&#x27; and &#x27;gce&#x27; roles (required for &#x27;iam&#x27;). A comma-seperated list of authorized service accounts. If the single value \&quot;*\&quot; is given, this is assumed to be all service accounts under the role&#x27;s project. If this is set on a GCE role, the inferred service account from the instance metadata token will be used. | [optional] 
**BoundZone** | Pointer to **string** | Deprecated: use \&quot;bound_zones\&quot; instead. | [optional] 
**BoundZones** | Pointer to **[]string** | Comma-separated list of permitted zones to which the GCE instance must belong. If a group is provided, it is assumed to be a zonal group. This can be a self-link or zone name. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**MaxJwtExp** | Pointer to **string** | Currently enabled for &#x27;iam&#x27; only. Duration in seconds from time of validation that a JWT must expire within. | [optional] [default to "900"]
**MaxTtl** | Pointer to **string** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Period** | Pointer to **string** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**ProjectId** | Pointer to **string** | Deprecated: use \&quot;bound_projects\&quot; instead | [optional] 
**ServiceAccounts** | Pointer to **[]string** | Deprecated: use \&quot;bound_service_accounts\&quot; instead. | [optional] 
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
**Type** | Pointer to **string** | Type of the role. Currently supported: iam, gce | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


