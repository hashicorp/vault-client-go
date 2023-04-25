# TransitRotateKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedKeyId** | Pointer to **string** | The UUID of the managed key to use for the new version of this transit key | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use for the new version of this transit key | [optional] 



## Methods


### NewTransitRotateKeyRequest

`func NewTransitRotateKeyRequest() *TransitRotateKeyRequest`

NewTransitRotateKeyRequest instantiates a new TransitRotateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitRotateKeyRequestWithDefaults

`func NewTransitRotateKeyRequestWithDefaults() *TransitRotateKeyRequest`

NewTransitRotateKeyRequestWithDefaults instantiates a new TransitRotateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetManagedKeyId

`func (o *TransitRotateKeyRequest) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *TransitRotateKeyRequest) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *TransitRotateKeyRequest) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.


### HasManagedKeyId

`func (o *TransitRotateKeyRequest) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.




### GetManagedKeyName

`func (o *TransitRotateKeyRequest) GetManagedKeyName() string`

GetManagedKeyName returns the ManagedKeyName field if non-nil, zero value otherwise.

### GetManagedKeyNameOk

`func (o *TransitRotateKeyRequest) GetManagedKeyNameOk() (*string, bool)`

GetManagedKeyNameOk returns a tuple with the ManagedKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyName

`func (o *TransitRotateKeyRequest) SetManagedKeyName(v string)`

SetManagedKeyName sets ManagedKeyName field to given value.


### HasManagedKeyName

`func (o *TransitRotateKeyRequest) HasManagedKeyName() bool`

HasManagedKeyName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


