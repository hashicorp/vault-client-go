# AWSWriteAuthRoleTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInstanceMigration** | Pointer to **bool** | If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution. | [optional] [default to false]
**DisallowReauthentication** | Pointer to **bool** | If set, only allows a single token to be granted per instance ID. In order to perform a fresh login, the entry in access list for the instance ID needs to be cleared using the &#39;auth/aws-ec2/identity-accesslist/&lt;instance_id&gt;&#39; endpoint. | [optional] [default to false]
**InstanceId** | Pointer to **string** | Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID. | [optional] 
**MaxTtl** | Pointer to **int32** | If set, specifies the maximum allowed token lifetime. | [optional] [default to 0]
**Policies** | Pointer to **[]string** | Policies to be associated with the tag. If set, must be a subset of the role&#39;s policies. If set, but set to an empty value, only the &#39;default&#39; policy will be given to issued tokens. | [optional] 

## Methods

### NewAWSWriteAuthRoleTagRequest

`func NewAWSWriteAuthRoleTagRequest() *AWSWriteAuthRoleTagRequest`

NewAWSWriteAuthRoleTagRequest instantiates a new AWSWriteAuthRoleTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSWriteAuthRoleTagRequestWithDefaults

`func NewAWSWriteAuthRoleTagRequestWithDefaults() *AWSWriteAuthRoleTagRequest`

NewAWSWriteAuthRoleTagRequestWithDefaults instantiates a new AWSWriteAuthRoleTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInstanceMigration

`func (o *AWSWriteAuthRoleTagRequest) GetAllowInstanceMigration() bool`

GetAllowInstanceMigration returns the AllowInstanceMigration field if non-nil, zero value otherwise.

### GetAllowInstanceMigrationOk

`func (o *AWSWriteAuthRoleTagRequest) GetAllowInstanceMigrationOk() (*bool, bool)`

GetAllowInstanceMigrationOk returns a tuple with the AllowInstanceMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInstanceMigration

`func (o *AWSWriteAuthRoleTagRequest) SetAllowInstanceMigration(v bool)`

SetAllowInstanceMigration sets AllowInstanceMigration field to given value.

### HasAllowInstanceMigration

`func (o *AWSWriteAuthRoleTagRequest) HasAllowInstanceMigration() bool`

HasAllowInstanceMigration returns a boolean if a field has been set.

### GetDisallowReauthentication

`func (o *AWSWriteAuthRoleTagRequest) GetDisallowReauthentication() bool`

GetDisallowReauthentication returns the DisallowReauthentication field if non-nil, zero value otherwise.

### GetDisallowReauthenticationOk

`func (o *AWSWriteAuthRoleTagRequest) GetDisallowReauthenticationOk() (*bool, bool)`

GetDisallowReauthenticationOk returns a tuple with the DisallowReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowReauthentication

`func (o *AWSWriteAuthRoleTagRequest) SetDisallowReauthentication(v bool)`

SetDisallowReauthentication sets DisallowReauthentication field to given value.

### HasDisallowReauthentication

`func (o *AWSWriteAuthRoleTagRequest) HasDisallowReauthentication() bool`

HasDisallowReauthentication returns a boolean if a field has been set.

### GetInstanceId

`func (o *AWSWriteAuthRoleTagRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AWSWriteAuthRoleTagRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AWSWriteAuthRoleTagRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AWSWriteAuthRoleTagRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetMaxTtl

`func (o *AWSWriteAuthRoleTagRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AWSWriteAuthRoleTagRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AWSWriteAuthRoleTagRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *AWSWriteAuthRoleTagRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPolicies

`func (o *AWSWriteAuthRoleTagRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AWSWriteAuthRoleTagRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AWSWriteAuthRoleTagRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AWSWriteAuthRoleTagRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


