# WriteRekeyVerifyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Specifies a single unseal share key from the new set of shares. | [optional] 
**Nonce** | Pointer to **string** | Specifies the nonce of the rekey verification operation. | [optional] 



## Methods


### NewWriteRekeyVerifyRequest

`func NewWriteRekeyVerifyRequest() *WriteRekeyVerifyRequest`

NewWriteRekeyVerifyRequest instantiates a new WriteRekeyVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRekeyVerifyRequestWithDefaults

`func NewWriteRekeyVerifyRequestWithDefaults() *WriteRekeyVerifyRequest`

NewWriteRekeyVerifyRequestWithDefaults instantiates a new WriteRekeyVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKey

`func (o *WriteRekeyVerifyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WriteRekeyVerifyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WriteRekeyVerifyRequest) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *WriteRekeyVerifyRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetNonce

`func (o *WriteRekeyVerifyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *WriteRekeyVerifyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *WriteRekeyVerifyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *WriteRekeyVerifyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


