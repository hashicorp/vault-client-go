# ConsulRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsulNamespace** | Pointer to **string** | Indicates which namespace that the token will be created within. Defaults to &#39;default&#39;. Available in Consul 1.7 and above. | [optional] 
**ConsulRoles** | Pointer to **[]string** | List of Consul roles to attach to the token. Either \&quot;policies\&quot; or \&quot;consul_roles\&quot; are required for Consul 1.5 and above. | [optional] 
**Lease** | Pointer to **int32** | Use \&quot;ttl\&quot; instead. | [optional] 
**Local** | Pointer to **bool** | Indicates that the token should not be replicated globally and instead be local to the current datacenter. Available in Consul 1.4 and above. | [optional] 
**MaxTtl** | Pointer to **int32** | Max TTL for the Consul token created from the role. | [optional] 
**Partition** | Pointer to **string** | Indicates which admin partition that the token will be created within. Defaults to &#39;default&#39;. Available in Consul 1.11 and above. | [optional] 
**Policies** | Pointer to **[]string** | List of policies to attach to the token. Either \&quot;policies\&quot; or \&quot;consul_roles\&quot; are required for Consul 1.5 and above, or just \&quot;policies\&quot; if using Consul 1.4. | [optional] 
**Policy** | Pointer to **string** | Policy document, base64 encoded. Required for &#39;client&#39; tokens. Required for Consul pre-1.4. | [optional] 
**TokenType** | Pointer to **string** | Which type of token to create: &#39;client&#39; or &#39;management&#39;. If a &#39;management&#39; token, the \&quot;policy\&quot;, \&quot;policies\&quot;, and \&quot;consul_roles\&quot; parameters are not required. Defaults to &#39;client&#39;. | [optional] [default to "client"]
**Ttl** | Pointer to **int32** | TTL for the Consul token created from the role. | [optional] 

## Methods

### NewConsulRolesRequest

`func NewConsulRolesRequest() *ConsulRolesRequest`

NewConsulRolesRequest instantiates a new ConsulRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulRolesRequestWithDefaults

`func NewConsulRolesRequestWithDefaults() *ConsulRolesRequest`

NewConsulRolesRequestWithDefaults instantiates a new ConsulRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsulNamespace

`func (o *ConsulRolesRequest) GetConsulNamespace() string`

GetConsulNamespace returns the ConsulNamespace field if non-nil, zero value otherwise.

### GetConsulNamespaceOk

`func (o *ConsulRolesRequest) GetConsulNamespaceOk() (*string, bool)`

GetConsulNamespaceOk returns a tuple with the ConsulNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulNamespace

`func (o *ConsulRolesRequest) SetConsulNamespace(v string)`

SetConsulNamespace sets ConsulNamespace field to given value.

### HasConsulNamespace

`func (o *ConsulRolesRequest) HasConsulNamespace() bool`

HasConsulNamespace returns a boolean if a field has been set.

### GetConsulRoles

`func (o *ConsulRolesRequest) GetConsulRoles() []string`

GetConsulRoles returns the ConsulRoles field if non-nil, zero value otherwise.

### GetConsulRolesOk

`func (o *ConsulRolesRequest) GetConsulRolesOk() (*[]string, bool)`

GetConsulRolesOk returns a tuple with the ConsulRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulRoles

`func (o *ConsulRolesRequest) SetConsulRoles(v []string)`

SetConsulRoles sets ConsulRoles field to given value.

### HasConsulRoles

`func (o *ConsulRolesRequest) HasConsulRoles() bool`

HasConsulRoles returns a boolean if a field has been set.

### GetLease

`func (o *ConsulRolesRequest) GetLease() int32`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *ConsulRolesRequest) GetLeaseOk() (*int32, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *ConsulRolesRequest) SetLease(v int32)`

SetLease sets Lease field to given value.

### HasLease

`func (o *ConsulRolesRequest) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetLocal

`func (o *ConsulRolesRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *ConsulRolesRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *ConsulRolesRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *ConsulRolesRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetMaxTtl

`func (o *ConsulRolesRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *ConsulRolesRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *ConsulRolesRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *ConsulRolesRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPartition

`func (o *ConsulRolesRequest) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ConsulRolesRequest) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ConsulRolesRequest) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ConsulRolesRequest) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPolicies

`func (o *ConsulRolesRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ConsulRolesRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ConsulRolesRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ConsulRolesRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetPolicy

`func (o *ConsulRolesRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ConsulRolesRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ConsulRolesRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ConsulRolesRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetTokenType

`func (o *ConsulRolesRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ConsulRolesRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ConsulRolesRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ConsulRolesRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *ConsulRolesRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ConsulRolesRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ConsulRolesRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ConsulRolesRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


