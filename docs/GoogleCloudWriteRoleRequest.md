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



## Methods


### NewGoogleCloudWriteRoleRequest

`func NewGoogleCloudWriteRoleRequest() *GoogleCloudWriteRoleRequest`

NewGoogleCloudWriteRoleRequest instantiates a new GoogleCloudWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudWriteRoleRequestWithDefaults

`func NewGoogleCloudWriteRoleRequestWithDefaults() *GoogleCloudWriteRoleRequest`

NewGoogleCloudWriteRoleRequestWithDefaults instantiates a new GoogleCloudWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAddGroupAliases

`func (o *GoogleCloudWriteRoleRequest) GetAddGroupAliases() bool`

GetAddGroupAliases returns the AddGroupAliases field if non-nil, zero value otherwise.

### GetAddGroupAliasesOk

`func (o *GoogleCloudWriteRoleRequest) GetAddGroupAliasesOk() (*bool, bool)`

GetAddGroupAliasesOk returns a tuple with the AddGroupAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGroupAliases

`func (o *GoogleCloudWriteRoleRequest) SetAddGroupAliases(v bool)`

SetAddGroupAliases sets AddGroupAliases field to given value.


### HasAddGroupAliases

`func (o *GoogleCloudWriteRoleRequest) HasAddGroupAliases() bool`

HasAddGroupAliases returns a boolean if a field has been set.




### GetAllowGceInference

`func (o *GoogleCloudWriteRoleRequest) GetAllowGceInference() bool`

GetAllowGceInference returns the AllowGceInference field if non-nil, zero value otherwise.

### GetAllowGceInferenceOk

`func (o *GoogleCloudWriteRoleRequest) GetAllowGceInferenceOk() (*bool, bool)`

GetAllowGceInferenceOk returns a tuple with the AllowGceInference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGceInference

`func (o *GoogleCloudWriteRoleRequest) SetAllowGceInference(v bool)`

SetAllowGceInference sets AllowGceInference field to given value.


### HasAllowGceInference

`func (o *GoogleCloudWriteRoleRequest) HasAllowGceInference() bool`

HasAllowGceInference returns a boolean if a field has been set.




### GetBoundInstanceGroup

`func (o *GoogleCloudWriteRoleRequest) GetBoundInstanceGroup() string`

GetBoundInstanceGroup returns the BoundInstanceGroup field if non-nil, zero value otherwise.

### GetBoundInstanceGroupOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundInstanceGroupOk() (*string, bool)`

GetBoundInstanceGroupOk returns a tuple with the BoundInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceGroup

`func (o *GoogleCloudWriteRoleRequest) SetBoundInstanceGroup(v string)`

SetBoundInstanceGroup sets BoundInstanceGroup field to given value.


### HasBoundInstanceGroup

`func (o *GoogleCloudWriteRoleRequest) HasBoundInstanceGroup() bool`

HasBoundInstanceGroup returns a boolean if a field has been set.




### GetBoundInstanceGroups

`func (o *GoogleCloudWriteRoleRequest) GetBoundInstanceGroups() []string`

GetBoundInstanceGroups returns the BoundInstanceGroups field if non-nil, zero value otherwise.

### GetBoundInstanceGroupsOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundInstanceGroupsOk() (*[]string, bool)`

GetBoundInstanceGroupsOk returns a tuple with the BoundInstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceGroups

`func (o *GoogleCloudWriteRoleRequest) SetBoundInstanceGroups(v []string)`

SetBoundInstanceGroups sets BoundInstanceGroups field to given value.


### HasBoundInstanceGroups

`func (o *GoogleCloudWriteRoleRequest) HasBoundInstanceGroups() bool`

HasBoundInstanceGroups returns a boolean if a field has been set.




### GetBoundLabels

`func (o *GoogleCloudWriteRoleRequest) GetBoundLabels() []string`

GetBoundLabels returns the BoundLabels field if non-nil, zero value otherwise.

### GetBoundLabelsOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundLabelsOk() (*[]string, bool)`

GetBoundLabelsOk returns a tuple with the BoundLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLabels

`func (o *GoogleCloudWriteRoleRequest) SetBoundLabels(v []string)`

SetBoundLabels sets BoundLabels field to given value.


### HasBoundLabels

`func (o *GoogleCloudWriteRoleRequest) HasBoundLabels() bool`

HasBoundLabels returns a boolean if a field has been set.




### GetBoundProjects

`func (o *GoogleCloudWriteRoleRequest) GetBoundProjects() []string`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundProjectsOk() (*[]string, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *GoogleCloudWriteRoleRequest) SetBoundProjects(v []string)`

SetBoundProjects sets BoundProjects field to given value.


### HasBoundProjects

`func (o *GoogleCloudWriteRoleRequest) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.




### GetBoundRegion

`func (o *GoogleCloudWriteRoleRequest) GetBoundRegion() string`

GetBoundRegion returns the BoundRegion field if non-nil, zero value otherwise.

### GetBoundRegionOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundRegionOk() (*string, bool)`

GetBoundRegionOk returns a tuple with the BoundRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegion

`func (o *GoogleCloudWriteRoleRequest) SetBoundRegion(v string)`

SetBoundRegion sets BoundRegion field to given value.


### HasBoundRegion

`func (o *GoogleCloudWriteRoleRequest) HasBoundRegion() bool`

HasBoundRegion returns a boolean if a field has been set.




### GetBoundRegions

`func (o *GoogleCloudWriteRoleRequest) GetBoundRegions() []string`

GetBoundRegions returns the BoundRegions field if non-nil, zero value otherwise.

### GetBoundRegionsOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundRegionsOk() (*[]string, bool)`

GetBoundRegionsOk returns a tuple with the BoundRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegions

`func (o *GoogleCloudWriteRoleRequest) SetBoundRegions(v []string)`

SetBoundRegions sets BoundRegions field to given value.


### HasBoundRegions

`func (o *GoogleCloudWriteRoleRequest) HasBoundRegions() bool`

HasBoundRegions returns a boolean if a field has been set.




### GetBoundServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) GetBoundServiceAccounts() []string`

GetBoundServiceAccounts returns the BoundServiceAccounts field if non-nil, zero value otherwise.

### GetBoundServiceAccountsOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundServiceAccountsOk() (*[]string, bool)`

GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) SetBoundServiceAccounts(v []string)`

SetBoundServiceAccounts sets BoundServiceAccounts field to given value.


### HasBoundServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) HasBoundServiceAccounts() bool`

HasBoundServiceAccounts returns a boolean if a field has been set.




### GetBoundZone

`func (o *GoogleCloudWriteRoleRequest) GetBoundZone() string`

GetBoundZone returns the BoundZone field if non-nil, zero value otherwise.

### GetBoundZoneOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundZoneOk() (*string, bool)`

GetBoundZoneOk returns a tuple with the BoundZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZone

`func (o *GoogleCloudWriteRoleRequest) SetBoundZone(v string)`

SetBoundZone sets BoundZone field to given value.


### HasBoundZone

`func (o *GoogleCloudWriteRoleRequest) HasBoundZone() bool`

HasBoundZone returns a boolean if a field has been set.




### GetBoundZones

`func (o *GoogleCloudWriteRoleRequest) GetBoundZones() []string`

GetBoundZones returns the BoundZones field if non-nil, zero value otherwise.

### GetBoundZonesOk

`func (o *GoogleCloudWriteRoleRequest) GetBoundZonesOk() (*[]string, bool)`

GetBoundZonesOk returns a tuple with the BoundZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZones

`func (o *GoogleCloudWriteRoleRequest) SetBoundZones(v []string)`

SetBoundZones sets BoundZones field to given value.


### HasBoundZones

`func (o *GoogleCloudWriteRoleRequest) HasBoundZones() bool`

HasBoundZones returns a boolean if a field has been set.




### GetMaxJwtExp

`func (o *GoogleCloudWriteRoleRequest) GetMaxJwtExp() string`

GetMaxJwtExp returns the MaxJwtExp field if non-nil, zero value otherwise.

### GetMaxJwtExpOk

`func (o *GoogleCloudWriteRoleRequest) GetMaxJwtExpOk() (*string, bool)`

GetMaxJwtExpOk returns a tuple with the MaxJwtExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJwtExp

`func (o *GoogleCloudWriteRoleRequest) SetMaxJwtExp(v string)`

SetMaxJwtExp sets MaxJwtExp field to given value.


### HasMaxJwtExp

`func (o *GoogleCloudWriteRoleRequest) HasMaxJwtExp() bool`

HasMaxJwtExp returns a boolean if a field has been set.




### GetMaxTtl

`func (o *GoogleCloudWriteRoleRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *GoogleCloudWriteRoleRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *GoogleCloudWriteRoleRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *GoogleCloudWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetPeriod

`func (o *GoogleCloudWriteRoleRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GoogleCloudWriteRoleRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GoogleCloudWriteRoleRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *GoogleCloudWriteRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *GoogleCloudWriteRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GoogleCloudWriteRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GoogleCloudWriteRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *GoogleCloudWriteRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetProjectId

`func (o *GoogleCloudWriteRoleRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleCloudWriteRoleRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleCloudWriteRoleRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### HasProjectId

`func (o *GoogleCloudWriteRoleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.




### GetServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *GoogleCloudWriteRoleRequest) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.


### HasServiceAccounts

`func (o *GoogleCloudWriteRoleRequest) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *GoogleCloudWriteRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *GoogleCloudWriteRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *GoogleCloudWriteRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *GoogleCloudWriteRoleRequest) GetTokenExplicitMaxTtl() string`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenExplicitMaxTtlOk() (*string, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *GoogleCloudWriteRoleRequest) SetTokenExplicitMaxTtl(v string)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *GoogleCloudWriteRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *GoogleCloudWriteRoleRequest) GetTokenMaxTtl() string`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenMaxTtlOk() (*string, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *GoogleCloudWriteRoleRequest) SetTokenMaxTtl(v string)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *GoogleCloudWriteRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *GoogleCloudWriteRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *GoogleCloudWriteRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *GoogleCloudWriteRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *GoogleCloudWriteRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *GoogleCloudWriteRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *GoogleCloudWriteRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *GoogleCloudWriteRoleRequest) GetTokenPeriod() string`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenPeriodOk() (*string, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *GoogleCloudWriteRoleRequest) SetTokenPeriod(v string)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *GoogleCloudWriteRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *GoogleCloudWriteRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *GoogleCloudWriteRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *GoogleCloudWriteRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *GoogleCloudWriteRoleRequest) GetTokenTtl() string`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenTtlOk() (*string, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *GoogleCloudWriteRoleRequest) SetTokenTtl(v string)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *GoogleCloudWriteRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *GoogleCloudWriteRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GoogleCloudWriteRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GoogleCloudWriteRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *GoogleCloudWriteRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *GoogleCloudWriteRoleRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GoogleCloudWriteRoleRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GoogleCloudWriteRoleRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *GoogleCloudWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetType

`func (o *GoogleCloudWriteRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleCloudWriteRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleCloudWriteRoleRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *GoogleCloudWriteRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


