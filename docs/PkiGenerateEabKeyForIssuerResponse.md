# PkiGenerateEabKeyForIssuerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | The ACME directory to which the key belongs | [optional] 
**CreatedOn** | Pointer to **time.Time** | An RFC3339 formatted date time when the EAB token was created | [optional] 
**Id** | Pointer to **string** | The EAB key identifier | [optional] 
**Key** | Pointer to **string** | The EAB hmac key | [optional] 
**KeyType** | Pointer to **string** | The EAB key type | [optional] 



## Methods


### NewPkiGenerateEabKeyForIssuerResponse

`func NewPkiGenerateEabKeyForIssuerResponse() *PkiGenerateEabKeyForIssuerResponse`

NewPkiGenerateEabKeyForIssuerResponse instantiates a new PkiGenerateEabKeyForIssuerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateEabKeyForIssuerResponseWithDefaults

`func NewPkiGenerateEabKeyForIssuerResponseWithDefaults() *PkiGenerateEabKeyForIssuerResponse`

NewPkiGenerateEabKeyForIssuerResponseWithDefaults instantiates a new PkiGenerateEabKeyForIssuerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerResponse) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *PkiGenerateEabKeyForIssuerResponse) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerResponse) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.


### HasAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerResponse) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.




### GetCreatedOn

`func (o *PkiGenerateEabKeyForIssuerResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PkiGenerateEabKeyForIssuerResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PkiGenerateEabKeyForIssuerResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### HasCreatedOn

`func (o *PkiGenerateEabKeyForIssuerResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.




### GetId

`func (o *PkiGenerateEabKeyForIssuerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkiGenerateEabKeyForIssuerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkiGenerateEabKeyForIssuerResponse) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *PkiGenerateEabKeyForIssuerResponse) HasId() bool`

HasId returns a boolean if a field has been set.




### GetKey

`func (o *PkiGenerateEabKeyForIssuerResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PkiGenerateEabKeyForIssuerResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PkiGenerateEabKeyForIssuerResponse) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *PkiGenerateEabKeyForIssuerResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateEabKeyForIssuerResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateEabKeyForIssuerResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateEabKeyForIssuerResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateEabKeyForIssuerResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


