# PkiCrossSignIntermediateResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csr** | Pointer to **string** | Certificate signing request. | [optional] 
**KeyId** | Pointer to **string** | Id of the key. | [optional] 
**PrivateKey** | Pointer to **string** | Generated private key. | [optional] 
**PrivateKeyType** | Pointer to **string** | Specifies the format used for marshaling the private key. | [optional] 



## Methods


### NewPkiCrossSignIntermediateResponse

`func NewPkiCrossSignIntermediateResponse() *PkiCrossSignIntermediateResponse`

NewPkiCrossSignIntermediateResponse instantiates a new PkiCrossSignIntermediateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiCrossSignIntermediateResponseWithDefaults

`func NewPkiCrossSignIntermediateResponseWithDefaults() *PkiCrossSignIntermediateResponse`

NewPkiCrossSignIntermediateResponseWithDefaults instantiates a new PkiCrossSignIntermediateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCsr

`func (o *PkiCrossSignIntermediateResponse) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *PkiCrossSignIntermediateResponse) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *PkiCrossSignIntermediateResponse) SetCsr(v string)`

SetCsr sets Csr field to given value.


### HasCsr

`func (o *PkiCrossSignIntermediateResponse) HasCsr() bool`

HasCsr returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiCrossSignIntermediateResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiCrossSignIntermediateResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiCrossSignIntermediateResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiCrossSignIntermediateResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetPrivateKey

`func (o *PkiCrossSignIntermediateResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiCrossSignIntermediateResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiCrossSignIntermediateResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *PkiCrossSignIntermediateResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.




### GetPrivateKeyType

`func (o *PkiCrossSignIntermediateResponse) GetPrivateKeyType() string`

GetPrivateKeyType returns the PrivateKeyType field if non-nil, zero value otherwise.

### GetPrivateKeyTypeOk

`func (o *PkiCrossSignIntermediateResponse) GetPrivateKeyTypeOk() (*string, bool)`

GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyType

`func (o *PkiCrossSignIntermediateResponse) SetPrivateKeyType(v string)`

SetPrivateKeyType sets PrivateKeyType field to given value.


### HasPrivateKeyType

`func (o *PkiCrossSignIntermediateResponse) HasPrivateKeyType() bool`

HasPrivateKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


