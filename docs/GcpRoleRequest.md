# GcpRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddGroupAliases** | Pointer to **bool** | If true, will add group aliases to auth tokens generated under this role. This will add the full list of ancestors (projects, folders, organizations) for the given entity&#39;s project. Requires IAM permission &#x60;resourcemanager.projects.get&#x60; on this project. | [optional] [default to false]
**AllowGceInference** | Pointer to **bool** | &#39;iam&#39; roles only. If false, Vault will not not allow GCE instances to login in against this role | [optional] [default to true]
**BoundInstanceGroup** | Pointer to **string** | Deprecated: use \&quot;bound_instance_groups\&quot; instead. | [optional] 
**BoundInstanceGroups** | Pointer to **[]string** | Comma-separated list of permitted instance groups to which the GCE instance must belong. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundLabels** | Pointer to **[]string** | Comma-separated list of GCP labels formatted as\&quot;key:value\&quot; strings that must be present on the GCE instance in order to authenticate. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundProjects** | Pointer to **[]string** | GCP Projects that authenticating entities must belong to. | [optional] 
**BoundRegion** | Pointer to **string** | Deprecated: use \&quot;bound_regions\&quot; instead. | [optional] 
**BoundRegions** | Pointer to **[]string** | Comma-separated list of permitted regions to which the GCE instance must belong. If a group is provided, it is assumed to be a regional group. If \&quot;zone\&quot; is provided, this option is ignored. This can be a self-link or region name. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**BoundServiceAccounts** | Pointer to **[]string** | Can be set for both &#39;iam&#39; and &#39;gce&#39; roles (required for &#39;iam&#39;). A comma-seperated list of authorized service accounts. If the single value \&quot;*\&quot; is given, this is assumed to be all service accounts under the role&#39;s project. If this is set on a GCE role, the inferred service account from the instance metadata token will be used. | [optional] 
**BoundZone** | Pointer to **string** | Deprecated: use \&quot;bound_zones\&quot; instead. | [optional] 
**BoundZones** | Pointer to **[]string** | Comma-separated list of permitted zones to which the GCE instance must belong. If a group is provided, it is assumed to be a zonal group. This can be a self-link or zone name. This option only applies to \&quot;gce\&quot; roles. | [optional] 
**MaxJwtExp** | Pointer to **int32** | Currently enabled for &#39;iam&#39; only. Duration in seconds from time of validation that a JWT must expire within. | [optional] [default to 900]
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**ProjectId** | Pointer to **string** | Deprecated: use \&quot;bound_projects\&quot; instead | [optional] 
**ServiceAccounts** | Pointer to **[]string** | Deprecated: use \&quot;bound_service_accounts\&quot; instead. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 
**Type** | Pointer to **string** | Type of the role. Currently supported: iam, gce | [optional] 

## Methods

### NewGcpRoleRequest

`func NewGcpRoleRequest() *GcpRoleRequest`

NewGcpRoleRequest instantiates a new GcpRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpRoleRequestWithDefaults

`func NewGcpRoleRequestWithDefaults() *GcpRoleRequest`

NewGcpRoleRequestWithDefaults instantiates a new GcpRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddGroupAliases

`func (o *GcpRoleRequest) GetAddGroupAliases() bool`

GetAddGroupAliases returns the AddGroupAliases field if non-nil, zero value otherwise.

### GetAddGroupAliasesOk

`func (o *GcpRoleRequest) GetAddGroupAliasesOk() (*bool, bool)`

GetAddGroupAliasesOk returns a tuple with the AddGroupAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGroupAliases

`func (o *GcpRoleRequest) SetAddGroupAliases(v bool)`

SetAddGroupAliases sets AddGroupAliases field to given value.

### HasAddGroupAliases

`func (o *GcpRoleRequest) HasAddGroupAliases() bool`

HasAddGroupAliases returns a boolean if a field has been set.

### GetAllowGceInference

`func (o *GcpRoleRequest) GetAllowGceInference() bool`

GetAllowGceInference returns the AllowGceInference field if non-nil, zero value otherwise.

### GetAllowGceInferenceOk

`func (o *GcpRoleRequest) GetAllowGceInferenceOk() (*bool, bool)`

GetAllowGceInferenceOk returns a tuple with the AllowGceInference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGceInference

`func (o *GcpRoleRequest) SetAllowGceInference(v bool)`

SetAllowGceInference sets AllowGceInference field to given value.

### HasAllowGceInference

`func (o *GcpRoleRequest) HasAllowGceInference() bool`

HasAllowGceInference returns a boolean if a field has been set.

### GetBoundInstanceGroup

`func (o *GcpRoleRequest) GetBoundInstanceGroup() string`

GetBoundInstanceGroup returns the BoundInstanceGroup field if non-nil, zero value otherwise.

### GetBoundInstanceGroupOk

`func (o *GcpRoleRequest) GetBoundInstanceGroupOk() (*string, bool)`

GetBoundInstanceGroupOk returns a tuple with the BoundInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceGroup

`func (o *GcpRoleRequest) SetBoundInstanceGroup(v string)`

SetBoundInstanceGroup sets BoundInstanceGroup field to given value.

### HasBoundInstanceGroup

`func (o *GcpRoleRequest) HasBoundInstanceGroup() bool`

HasBoundInstanceGroup returns a boolean if a field has been set.

### GetBoundInstanceGroups

`func (o *GcpRoleRequest) GetBoundInstanceGroups() []string`

GetBoundInstanceGroups returns the BoundInstanceGroups field if non-nil, zero value otherwise.

### GetBoundInstanceGroupsOk

`func (o *GcpRoleRequest) GetBoundInstanceGroupsOk() (*[]string, bool)`

GetBoundInstanceGroupsOk returns a tuple with the BoundInstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceGroups

`func (o *GcpRoleRequest) SetBoundInstanceGroups(v []string)`

SetBoundInstanceGroups sets BoundInstanceGroups field to given value.

### HasBoundInstanceGroups

`func (o *GcpRoleRequest) HasBoundInstanceGroups() bool`

HasBoundInstanceGroups returns a boolean if a field has been set.

### GetBoundLabels

`func (o *GcpRoleRequest) GetBoundLabels() []string`

GetBoundLabels returns the BoundLabels field if non-nil, zero value otherwise.

### GetBoundLabelsOk

`func (o *GcpRoleRequest) GetBoundLabelsOk() (*[]string, bool)`

GetBoundLabelsOk returns a tuple with the BoundLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLabels

`func (o *GcpRoleRequest) SetBoundLabels(v []string)`

SetBoundLabels sets BoundLabels field to given value.

### HasBoundLabels

`func (o *GcpRoleRequest) HasBoundLabels() bool`

HasBoundLabels returns a boolean if a field has been set.

### GetBoundProjects

`func (o *GcpRoleRequest) GetBoundProjects() []string`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *GcpRoleRequest) GetBoundProjectsOk() (*[]string, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *GcpRoleRequest) SetBoundProjects(v []string)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *GcpRoleRequest) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### GetBoundRegion

`func (o *GcpRoleRequest) GetBoundRegion() string`

GetBoundRegion returns the BoundRegion field if non-nil, zero value otherwise.

### GetBoundRegionOk

`func (o *GcpRoleRequest) GetBoundRegionOk() (*string, bool)`

GetBoundRegionOk returns a tuple with the BoundRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegion

`func (o *GcpRoleRequest) SetBoundRegion(v string)`

SetBoundRegion sets BoundRegion field to given value.

### HasBoundRegion

`func (o *GcpRoleRequest) HasBoundRegion() bool`

HasBoundRegion returns a boolean if a field has been set.

### GetBoundRegions

`func (o *GcpRoleRequest) GetBoundRegions() []string`

GetBoundRegions returns the BoundRegions field if non-nil, zero value otherwise.

### GetBoundRegionsOk

`func (o *GcpRoleRequest) GetBoundRegionsOk() (*[]string, bool)`

GetBoundRegionsOk returns a tuple with the BoundRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegions

`func (o *GcpRoleRequest) SetBoundRegions(v []string)`

SetBoundRegions sets BoundRegions field to given value.

### HasBoundRegions

`func (o *GcpRoleRequest) HasBoundRegions() bool`

HasBoundRegions returns a boolean if a field has been set.

### GetBoundServiceAccounts

`func (o *GcpRoleRequest) GetBoundServiceAccounts() []string`

GetBoundServiceAccounts returns the BoundServiceAccounts field if non-nil, zero value otherwise.

### GetBoundServiceAccountsOk

`func (o *GcpRoleRequest) GetBoundServiceAccountsOk() (*[]string, bool)`

GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccounts

`func (o *GcpRoleRequest) SetBoundServiceAccounts(v []string)`

SetBoundServiceAccounts sets BoundServiceAccounts field to given value.

### HasBoundServiceAccounts

`func (o *GcpRoleRequest) HasBoundServiceAccounts() bool`

HasBoundServiceAccounts returns a boolean if a field has been set.

### GetBoundZone

`func (o *GcpRoleRequest) GetBoundZone() string`

GetBoundZone returns the BoundZone field if non-nil, zero value otherwise.

### GetBoundZoneOk

`func (o *GcpRoleRequest) GetBoundZoneOk() (*string, bool)`

GetBoundZoneOk returns a tuple with the BoundZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZone

`func (o *GcpRoleRequest) SetBoundZone(v string)`

SetBoundZone sets BoundZone field to given value.

### HasBoundZone

`func (o *GcpRoleRequest) HasBoundZone() bool`

HasBoundZone returns a boolean if a field has been set.

### GetBoundZones

`func (o *GcpRoleRequest) GetBoundZones() []string`

GetBoundZones returns the BoundZones field if non-nil, zero value otherwise.

### GetBoundZonesOk

`func (o *GcpRoleRequest) GetBoundZonesOk() (*[]string, bool)`

GetBoundZonesOk returns a tuple with the BoundZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZones

`func (o *GcpRoleRequest) SetBoundZones(v []string)`

SetBoundZones sets BoundZones field to given value.

### HasBoundZones

`func (o *GcpRoleRequest) HasBoundZones() bool`

HasBoundZones returns a boolean if a field has been set.

### GetMaxJwtExp

`func (o *GcpRoleRequest) GetMaxJwtExp() int32`

GetMaxJwtExp returns the MaxJwtExp field if non-nil, zero value otherwise.

### GetMaxJwtExpOk

`func (o *GcpRoleRequest) GetMaxJwtExpOk() (*int32, bool)`

GetMaxJwtExpOk returns a tuple with the MaxJwtExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJwtExp

`func (o *GcpRoleRequest) SetMaxJwtExp(v int32)`

SetMaxJwtExp sets MaxJwtExp field to given value.

### HasMaxJwtExp

`func (o *GcpRoleRequest) HasMaxJwtExp() bool`

HasMaxJwtExp returns a boolean if a field has been set.

### GetMaxTtl

`func (o *GcpRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *GcpRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *GcpRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *GcpRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPeriod

`func (o *GcpRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GcpRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GcpRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GcpRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *GcpRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GcpRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GcpRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *GcpRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProjectId

`func (o *GcpRoleRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpRoleRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpRoleRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GcpRoleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServiceAccounts

`func (o *GcpRoleRequest) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *GcpRoleRequest) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *GcpRoleRequest) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *GcpRoleRequest) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *GcpRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *GcpRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *GcpRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *GcpRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *GcpRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *GcpRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *GcpRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *GcpRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *GcpRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *GcpRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *GcpRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *GcpRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *GcpRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *GcpRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *GcpRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *GcpRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *GcpRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *GcpRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *GcpRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *GcpRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *GcpRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *GcpRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *GcpRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *GcpRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *GcpRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *GcpRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *GcpRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *GcpRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *GcpRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *GcpRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *GcpRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *GcpRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *GcpRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GcpRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GcpRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *GcpRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *GcpRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GcpRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GcpRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GcpRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *GcpRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GcpRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GcpRoleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GcpRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


