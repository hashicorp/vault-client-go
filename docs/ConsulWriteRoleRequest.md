# ConsulWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**ConsulNamespace** | Pointer to **string** | Indicates which namespace that the token will be created within. Defaults to &#x27;default&#x27;. Available in Consul 1.7 and above. | [optional] 
**ConsulPolicies** | Pointer to **[]string** | List of policies to attach to the token. Either \&quot;consul_policies\&quot; or \&quot;consul_roles\&quot; are required for Consul 1.5 and above, or just \&quot;consul_policies\&quot; if using Consul 1.4. | [optional] 
**ConsulRoles** | Pointer to **[]string** | List of Consul roles to attach to the token. Either \&quot;policies\&quot; or \&quot;consul_roles\&quot; are required for Consul 1.5 and above. | [optional] 
**Lease** | Pointer to **int32** | Use \&quot;ttl\&quot; instead. | [optional] 
**Local** | Pointer to **bool** | Indicates that the token should not be replicated globally and instead be local to the current datacenter. Available in Consul 1.4 and above. | [optional] 
**MaxTtl** | Pointer to **int32** | Max TTL for the Consul token created from the role. | [optional] 
**NodeIdentities** | Pointer to **[]string** | List of Node Identities to attach to the token. Available in Consul 1.8.1 or above. | [optional] 
**Partition** | Pointer to **string** | Indicates which admin partition that the token will be created within. Defaults to &#x27;default&#x27;. Available in Consul 1.11 and above. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;consul_policies\&quot; instead. | [optional] 
**Policy** | Pointer to **string** | Policy document, base64 encoded. Required for &#x27;client&#x27; tokens. Required for Consul pre-1.4. | [optional] 
**ServiceIdentities** | Pointer to **[]string** | List of Service Identities to attach to the token, separated by semicolons. Available in Consul 1.5 or above. | [optional] 
**TokenType** | Pointer to **string** | Which type of token to create: &#x27;client&#x27; or &#x27;management&#x27;. If a &#x27;management&#x27; token, the \&quot;policy\&quot;, \&quot;policies\&quot;, and \&quot;consul_roles\&quot; parameters are not required. Defaults to &#x27;client&#x27;. | [optional] [default to "client"]
**Ttl** | Pointer to **int32** | TTL for the Consul token created from the role. | [optional] 



## Methods


### NewConsulWriteRoleRequest

`func NewConsulWriteRoleRequest() *ConsulWriteRoleRequest`

NewConsulWriteRoleRequest instantiates a new ConsulWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulWriteRoleRequestWithDefaults

`func NewConsulWriteRoleRequestWithDefaults() *ConsulWriteRoleRequest`

NewConsulWriteRoleRequestWithDefaults instantiates a new ConsulWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetConsulNamespace

`func (o *ConsulWriteRoleRequest) GetConsulNamespace() string`

GetConsulNamespace returns the ConsulNamespace field if non-nil, zero value otherwise.

### GetConsulNamespaceOk

`func (o *ConsulWriteRoleRequest) GetConsulNamespaceOk() (*string, bool)`

GetConsulNamespaceOk returns a tuple with the ConsulNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulNamespace

`func (o *ConsulWriteRoleRequest) SetConsulNamespace(v string)`

SetConsulNamespace sets ConsulNamespace field to given value.


### HasConsulNamespace

`func (o *ConsulWriteRoleRequest) HasConsulNamespace() bool`

HasConsulNamespace returns a boolean if a field has been set.




### GetConsulPolicies

`func (o *ConsulWriteRoleRequest) GetConsulPolicies() []string`

GetConsulPolicies returns the ConsulPolicies field if non-nil, zero value otherwise.

### GetConsulPoliciesOk

`func (o *ConsulWriteRoleRequest) GetConsulPoliciesOk() (*[]string, bool)`

GetConsulPoliciesOk returns a tuple with the ConsulPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulPolicies

`func (o *ConsulWriteRoleRequest) SetConsulPolicies(v []string)`

SetConsulPolicies sets ConsulPolicies field to given value.


### HasConsulPolicies

`func (o *ConsulWriteRoleRequest) HasConsulPolicies() bool`

HasConsulPolicies returns a boolean if a field has been set.




### GetConsulRoles

`func (o *ConsulWriteRoleRequest) GetConsulRoles() []string`

GetConsulRoles returns the ConsulRoles field if non-nil, zero value otherwise.

### GetConsulRolesOk

`func (o *ConsulWriteRoleRequest) GetConsulRolesOk() (*[]string, bool)`

GetConsulRolesOk returns a tuple with the ConsulRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulRoles

`func (o *ConsulWriteRoleRequest) SetConsulRoles(v []string)`

SetConsulRoles sets ConsulRoles field to given value.


### HasConsulRoles

`func (o *ConsulWriteRoleRequest) HasConsulRoles() bool`

HasConsulRoles returns a boolean if a field has been set.




### GetLease

`func (o *ConsulWriteRoleRequest) GetLease() int32`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *ConsulWriteRoleRequest) GetLeaseOk() (*int32, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *ConsulWriteRoleRequest) SetLease(v int32)`

SetLease sets Lease field to given value.


### HasLease

`func (o *ConsulWriteRoleRequest) HasLease() bool`

HasLease returns a boolean if a field has been set.




### GetLocal

`func (o *ConsulWriteRoleRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *ConsulWriteRoleRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *ConsulWriteRoleRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *ConsulWriteRoleRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetMaxTtl

`func (o *ConsulWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *ConsulWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *ConsulWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *ConsulWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetNodeIdentities

`func (o *ConsulWriteRoleRequest) GetNodeIdentities() []string`

GetNodeIdentities returns the NodeIdentities field if non-nil, zero value otherwise.

### GetNodeIdentitiesOk

`func (o *ConsulWriteRoleRequest) GetNodeIdentitiesOk() (*[]string, bool)`

GetNodeIdentitiesOk returns a tuple with the NodeIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIdentities

`func (o *ConsulWriteRoleRequest) SetNodeIdentities(v []string)`

SetNodeIdentities sets NodeIdentities field to given value.


### HasNodeIdentities

`func (o *ConsulWriteRoleRequest) HasNodeIdentities() bool`

HasNodeIdentities returns a boolean if a field has been set.




### GetPartition

`func (o *ConsulWriteRoleRequest) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ConsulWriteRoleRequest) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ConsulWriteRoleRequest) SetPartition(v string)`

SetPartition sets Partition field to given value.


### HasPartition

`func (o *ConsulWriteRoleRequest) HasPartition() bool`

HasPartition returns a boolean if a field has been set.




### GetPolicies

`func (o *ConsulWriteRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ConsulWriteRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ConsulWriteRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *ConsulWriteRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetPolicy

`func (o *ConsulWriteRoleRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ConsulWriteRoleRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ConsulWriteRoleRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### HasPolicy

`func (o *ConsulWriteRoleRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.




### GetServiceIdentities

`func (o *ConsulWriteRoleRequest) GetServiceIdentities() []string`

GetServiceIdentities returns the ServiceIdentities field if non-nil, zero value otherwise.

### GetServiceIdentitiesOk

`func (o *ConsulWriteRoleRequest) GetServiceIdentitiesOk() (*[]string, bool)`

GetServiceIdentitiesOk returns a tuple with the ServiceIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIdentities

`func (o *ConsulWriteRoleRequest) SetServiceIdentities(v []string)`

SetServiceIdentities sets ServiceIdentities field to given value.


### HasServiceIdentities

`func (o *ConsulWriteRoleRequest) HasServiceIdentities() bool`

HasServiceIdentities returns a boolean if a field has been set.




### GetTokenType

`func (o *ConsulWriteRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ConsulWriteRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ConsulWriteRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *ConsulWriteRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *ConsulWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ConsulWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ConsulWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *ConsulWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


