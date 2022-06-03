# NomadRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to **bool** | Boolean value describing if the token should be global or not. Defaults to false. | [optional] 
**Policies** | Pointer to **[]string** | Comma-separated string or list of policies as previously created in Nomad. Required for &#39;client&#39; token. | [optional] 
**Type** | Pointer to **string** | Which type of token to create: &#39;client&#39; or &#39;management&#39;. If a &#39;management&#39; token, the \&quot;policies\&quot; parameter is not required. Defaults to &#39;client&#39;. | [optional] [default to "client"]

## Methods

### NewNomadRoleRequest

`func NewNomadRoleRequest() *NomadRoleRequest`

NewNomadRoleRequest instantiates a new NomadRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadRoleRequestWithDefaults

`func NewNomadRoleRequestWithDefaults() *NomadRoleRequest`

NewNomadRoleRequestWithDefaults instantiates a new NomadRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *NomadRoleRequest) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *NomadRoleRequest) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *NomadRoleRequest) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *NomadRoleRequest) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetPolicies

`func (o *NomadRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *NomadRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *NomadRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *NomadRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetType

`func (o *NomadRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NomadRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NomadRoleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NomadRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


