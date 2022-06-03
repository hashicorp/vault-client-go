# SystemRekeyVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Specifies a single unseal share key from the new set of shares. | [optional] 
**Nonce** | Pointer to **string** | Specifies the nonce of the rekey verification operation. | [optional] 

## Methods

### NewSystemRekeyVerifyRequest

`func NewSystemRekeyVerifyRequest() *SystemRekeyVerifyRequest`

NewSystemRekeyVerifyRequest instantiates a new SystemRekeyVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRekeyVerifyRequestWithDefaults

`func NewSystemRekeyVerifyRequestWithDefaults() *SystemRekeyVerifyRequest`

NewSystemRekeyVerifyRequestWithDefaults instantiates a new SystemRekeyVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SystemRekeyVerifyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemRekeyVerifyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemRekeyVerifyRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SystemRekeyVerifyRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNonce

`func (o *SystemRekeyVerifyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SystemRekeyVerifyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SystemRekeyVerifyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SystemRekeyVerifyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


