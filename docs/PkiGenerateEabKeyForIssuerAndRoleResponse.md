# PkiGenerateEabKeyForIssuerAndRoleResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | The ACME directory to which the key belongs | [optional] 
**CreatedOn** | Pointer to **time.Time** | An RFC3339 formatted date time when the EAB token was created | [optional] 
**Id** | Pointer to **string** | The EAB key identifier | [optional] 
**Key** | Pointer to **string** | The EAB hmac key | [optional] 
**KeyType** | Pointer to **string** | The EAB key type | [optional] 



## Methods


### NewPkiGenerateEabKeyForIssuerAndRoleResponse

`func NewPkiGenerateEabKeyForIssuerAndRoleResponse() *PkiGenerateEabKeyForIssuerAndRoleResponse`

NewPkiGenerateEabKeyForIssuerAndRoleResponse instantiates a new PkiGenerateEabKeyForIssuerAndRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateEabKeyForIssuerAndRoleResponseWithDefaults

`func NewPkiGenerateEabKeyForIssuerAndRoleResponseWithDefaults() *PkiGenerateEabKeyForIssuerAndRoleResponse`

NewPkiGenerateEabKeyForIssuerAndRoleResponseWithDefaults instantiates a new PkiGenerateEabKeyForIssuerAndRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.


### HasAcmeDirectory

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.




### GetCreatedOn

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### HasCreatedOn

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.




### GetId

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.




### GetKey

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateEabKeyForIssuerAndRoleResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


