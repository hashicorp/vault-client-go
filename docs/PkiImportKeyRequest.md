# PkiImportKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | Pointer to **string** | Optional name to be used for this key | [optional] 
**PemBundle** | Pointer to **string** | PEM-format, unencrypted secret key | [optional] 



## Methods


### NewPkiImportKeyRequest

`func NewPkiImportKeyRequest() *PkiImportKeyRequest`

NewPkiImportKeyRequest instantiates a new PkiImportKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiImportKeyRequestWithDefaults

`func NewPkiImportKeyRequestWithDefaults() *PkiImportKeyRequest`

NewPkiImportKeyRequestWithDefaults instantiates a new PkiImportKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKeyName

`func (o *PkiImportKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiImportKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiImportKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### HasKeyName

`func (o *PkiImportKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.




### GetPemBundle

`func (o *PkiImportKeyRequest) GetPemBundle() string`

GetPemBundle returns the PemBundle field if non-nil, zero value otherwise.

### GetPemBundleOk

`func (o *PkiImportKeyRequest) GetPemBundleOk() (*string, bool)`

GetPemBundleOk returns a tuple with the PemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemBundle

`func (o *PkiImportKeyRequest) SetPemBundle(v string)`

SetPemBundle sets PemBundle field to given value.


### HasPemBundle

`func (o *PkiImportKeyRequest) HasPemBundle() bool`

HasPemBundle returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


