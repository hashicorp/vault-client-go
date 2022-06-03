# IdentityMfaMethodTotpAdminGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Entity ID on which the generated secret needs to get stored. | 
**MethodId** | **string** | The unique identifier for this MFA method. | 

## Methods

### NewIdentityMfaMethodTotpAdminGenerateRequest

`func NewIdentityMfaMethodTotpAdminGenerateRequest(entityId string, methodId string, ) *IdentityMfaMethodTotpAdminGenerateRequest`

NewIdentityMfaMethodTotpAdminGenerateRequest instantiates a new IdentityMfaMethodTotpAdminGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults

`func NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults() *IdentityMfaMethodTotpAdminGenerateRequest`

NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults instantiates a new IdentityMfaMethodTotpAdminGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetMethodId

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *IdentityMfaMethodTotpAdminGenerateRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


