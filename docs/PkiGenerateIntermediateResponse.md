# PkiGenerateIntermediateResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csr** | Pointer to **string** | Certificate signing request. | [optional] 
**KeyId** | Pointer to **string** | Id of the key. | [optional] 
**PrivateKey** | Pointer to **string** | Generated private key. | [optional] 
**PrivateKeyType** | Pointer to **string** | Specifies the format used for marshaling the private key. | [optional] 



## Methods


### NewPkiGenerateIntermediateResponse

`func NewPkiGenerateIntermediateResponse() *PkiGenerateIntermediateResponse`

NewPkiGenerateIntermediateResponse instantiates a new PkiGenerateIntermediateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateIntermediateResponseWithDefaults

`func NewPkiGenerateIntermediateResponseWithDefaults() *PkiGenerateIntermediateResponse`

NewPkiGenerateIntermediateResponseWithDefaults instantiates a new PkiGenerateIntermediateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCsr

`func (o *PkiGenerateIntermediateResponse) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *PkiGenerateIntermediateResponse) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *PkiGenerateIntermediateResponse) SetCsr(v string)`

SetCsr sets Csr field to given value.


### HasCsr

`func (o *PkiGenerateIntermediateResponse) HasCsr() bool`

HasCsr returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiGenerateIntermediateResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiGenerateIntermediateResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiGenerateIntermediateResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiGenerateIntermediateResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetPrivateKey

`func (o *PkiGenerateIntermediateResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiGenerateIntermediateResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiGenerateIntermediateResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *PkiGenerateIntermediateResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.




### GetPrivateKeyType

`func (o *PkiGenerateIntermediateResponse) GetPrivateKeyType() string`

GetPrivateKeyType returns the PrivateKeyType field if non-nil, zero value otherwise.

### GetPrivateKeyTypeOk

`func (o *PkiGenerateIntermediateResponse) GetPrivateKeyTypeOk() (*string, bool)`

GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyType

`func (o *PkiGenerateIntermediateResponse) SetPrivateKeyType(v string)`

SetPrivateKeyType sets PrivateKeyType field to given value.


### HasPrivateKeyType

`func (o *PkiGenerateIntermediateResponse) HasPrivateKeyType() bool`

HasPrivateKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


