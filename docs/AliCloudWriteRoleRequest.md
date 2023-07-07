# AliCloudWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlinePolicies** | Pointer to **string** | JSON of policies to be dynamically applied to users of this role. | [optional] 
**MaxTtl** | Pointer to **string** | The maximum allowed lifetime of tokens issued using this role. | [optional] 
**RemotePolicies** | Pointer to **[]string** | The name and type of each remote policy to be applied. Example: \&quot;name:AliyunRDSReadOnlyAccess,type:System\&quot;. | [optional] 
**RoleArn** | Pointer to **string** | ARN of the role to be assumed. If provided, inline_policies and remote_policies should be blank. At creation time, this role must have configured trusted actors, and the access key and secret that will be used to assume the role (in /config) must qualify as a trusted actor. | [optional] 
**Ttl** | Pointer to **string** | Duration in seconds after which the issued token should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults. | [optional] 



## Methods


### NewAliCloudWriteRoleRequest

`func NewAliCloudWriteRoleRequest() *AliCloudWriteRoleRequest`

NewAliCloudWriteRoleRequest instantiates a new AliCloudWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliCloudWriteRoleRequestWithDefaults

`func NewAliCloudWriteRoleRequestWithDefaults() *AliCloudWriteRoleRequest`

NewAliCloudWriteRoleRequestWithDefaults instantiates a new AliCloudWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetInlinePolicies

`func (o *AliCloudWriteRoleRequest) GetInlinePolicies() string`

GetInlinePolicies returns the InlinePolicies field if non-nil, zero value otherwise.

### GetInlinePoliciesOk

`func (o *AliCloudWriteRoleRequest) GetInlinePoliciesOk() (*string, bool)`

GetInlinePoliciesOk returns a tuple with the InlinePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlinePolicies

`func (o *AliCloudWriteRoleRequest) SetInlinePolicies(v string)`

SetInlinePolicies sets InlinePolicies field to given value.


### HasInlinePolicies

`func (o *AliCloudWriteRoleRequest) HasInlinePolicies() bool`

HasInlinePolicies returns a boolean if a field has been set.




### GetMaxTtl

`func (o *AliCloudWriteRoleRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AliCloudWriteRoleRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AliCloudWriteRoleRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *AliCloudWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetRemotePolicies

`func (o *AliCloudWriteRoleRequest) GetRemotePolicies() []string`

GetRemotePolicies returns the RemotePolicies field if non-nil, zero value otherwise.

### GetRemotePoliciesOk

`func (o *AliCloudWriteRoleRequest) GetRemotePoliciesOk() (*[]string, bool)`

GetRemotePoliciesOk returns a tuple with the RemotePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePolicies

`func (o *AliCloudWriteRoleRequest) SetRemotePolicies(v []string)`

SetRemotePolicies sets RemotePolicies field to given value.


### HasRemotePolicies

`func (o *AliCloudWriteRoleRequest) HasRemotePolicies() bool`

HasRemotePolicies returns a boolean if a field has been set.




### GetRoleArn

`func (o *AliCloudWriteRoleRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AliCloudWriteRoleRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AliCloudWriteRoleRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### HasRoleArn

`func (o *AliCloudWriteRoleRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.




### GetTtl

`func (o *AliCloudWriteRoleRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AliCloudWriteRoleRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AliCloudWriteRoleRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AliCloudWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


