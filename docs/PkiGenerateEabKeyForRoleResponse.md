# PkiGenerateEabKeyForRoleResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectory** | Pointer to **string** | The ACME directory to which the key belongs | [optional] 
**CreatedOn** | Pointer to **time.Time** | An RFC3339 formatted date time when the EAB token was created | [optional] 
**Id** | Pointer to **string** | The EAB key identifier | [optional] 
**Key** | Pointer to **string** | The EAB hmac key | [optional] 
**KeyType** | Pointer to **string** | The EAB key type | [optional] 



## Methods


### NewPkiGenerateEabKeyForRoleResponse

`func NewPkiGenerateEabKeyForRoleResponse() *PkiGenerateEabKeyForRoleResponse`

NewPkiGenerateEabKeyForRoleResponse instantiates a new PkiGenerateEabKeyForRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateEabKeyForRoleResponseWithDefaults

`func NewPkiGenerateEabKeyForRoleResponseWithDefaults() *PkiGenerateEabKeyForRoleResponse`

NewPkiGenerateEabKeyForRoleResponseWithDefaults instantiates a new PkiGenerateEabKeyForRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAcmeDirectory

`func (o *PkiGenerateEabKeyForRoleResponse) GetAcmeDirectory() string`

GetAcmeDirectory returns the AcmeDirectory field if non-nil, zero value otherwise.

### GetAcmeDirectoryOk

`func (o *PkiGenerateEabKeyForRoleResponse) GetAcmeDirectoryOk() (*string, bool)`

GetAcmeDirectoryOk returns a tuple with the AcmeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectory

`func (o *PkiGenerateEabKeyForRoleResponse) SetAcmeDirectory(v string)`

SetAcmeDirectory sets AcmeDirectory field to given value.


### HasAcmeDirectory

`func (o *PkiGenerateEabKeyForRoleResponse) HasAcmeDirectory() bool`

HasAcmeDirectory returns a boolean if a field has been set.




### GetCreatedOn

`func (o *PkiGenerateEabKeyForRoleResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PkiGenerateEabKeyForRoleResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PkiGenerateEabKeyForRoleResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### HasCreatedOn

`func (o *PkiGenerateEabKeyForRoleResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.




### GetId

`func (o *PkiGenerateEabKeyForRoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkiGenerateEabKeyForRoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkiGenerateEabKeyForRoleResponse) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *PkiGenerateEabKeyForRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.




### GetKey

`func (o *PkiGenerateEabKeyForRoleResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PkiGenerateEabKeyForRoleResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PkiGenerateEabKeyForRoleResponse) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *PkiGenerateEabKeyForRoleResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateEabKeyForRoleResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateEabKeyForRoleResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateEabKeyForRoleResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateEabKeyForRoleResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


