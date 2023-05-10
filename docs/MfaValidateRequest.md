# MfaValidateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaPayload** | **map[string]interface{}** | A map from MFA method ID to a slice of passcodes or an empty slice if the method does not use passcodes | 
**MfaRequestId** | **string** | ID for this MFA request | 



## Methods


### NewMfaValidateRequest

`func NewMfaValidateRequest(mfaPayload map[string]interface{}, mfaRequestId string, ) *MfaValidateRequest`

NewMfaValidateRequest instantiates a new MfaValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaValidateRequestWithDefaults

`func NewMfaValidateRequestWithDefaults() *MfaValidateRequest`

NewMfaValidateRequestWithDefaults instantiates a new MfaValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMfaPayload

`func (o *MfaValidateRequest) GetMfaPayload() map[string]interface{}`

GetMfaPayload returns the MfaPayload field if non-nil, zero value otherwise.

### GetMfaPayloadOk

`func (o *MfaValidateRequest) GetMfaPayloadOk() (*map[string]interface{}, bool)`

GetMfaPayloadOk returns a tuple with the MfaPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaPayload

`func (o *MfaValidateRequest) SetMfaPayload(v map[string]interface{})`

SetMfaPayload sets MfaPayload field to given value.





### GetMfaRequestId

`func (o *MfaValidateRequest) GetMfaRequestId() string`

GetMfaRequestId returns the MfaRequestId field if non-nil, zero value otherwise.

### GetMfaRequestIdOk

`func (o *MfaValidateRequest) GetMfaRequestIdOk() (*string, bool)`

GetMfaRequestIdOk returns a tuple with the MfaRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequestId

`func (o *MfaValidateRequest) SetMfaRequestId(v string)`

SetMfaRequestId sets MfaRequestId field to given value.










[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


