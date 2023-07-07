# PkiGenerateEabKeyResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | The ACME directory to which the key belongs | [optional] 
**CreatedOn** | Pointer to **time.Time** | An RFC3339 formatted date time when the EAB token was created | [optional] 
**Id** | Pointer to **string** | The EAB key identifier | [optional] 
**Key** | Pointer to **string** | The EAB hmac key | [optional] 
**KeyType** | Pointer to **string** | The EAB key type | [optional] 



## Methods


### NewPkiGenerateEabKeyResponse

`func NewPkiGenerateEabKeyResponse() *PkiGenerateEabKeyResponse`

NewPkiGenerateEabKeyResponse instantiates a new PkiGenerateEabKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateEabKeyResponseWithDefaults

`func NewPkiGenerateEabKeyResponseWithDefaults() *PkiGenerateEabKeyResponse`

NewPkiGenerateEabKeyResponseWithDefaults instantiates a new PkiGenerateEabKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAcmeDirectory

`func (o *PkiGenerateEabKeyResponse) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *PkiGenerateEabKeyResponse) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *PkiGenerateEabKeyResponse) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.


### HasAcmeDirectory

`func (o *PkiGenerateEabKeyResponse) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.




### GetCreatedOn

`func (o *PkiGenerateEabKeyResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PkiGenerateEabKeyResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PkiGenerateEabKeyResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### HasCreatedOn

`func (o *PkiGenerateEabKeyResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.




### GetId

`func (o *PkiGenerateEabKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkiGenerateEabKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkiGenerateEabKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *PkiGenerateEabKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.




### GetKey

`func (o *PkiGenerateEabKeyResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PkiGenerateEabKeyResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PkiGenerateEabKeyResponse) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *PkiGenerateEabKeyResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateEabKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateEabKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateEabKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateEabKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


