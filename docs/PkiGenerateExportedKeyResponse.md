# PkiGenerateExportedKeyResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** | ID assigned to this key. | [optional] 
**KeyName** | Pointer to **string** | Name assigned to this key. | [optional] 
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot; and \&quot;ed25519\&quot; are the only valid values. | [optional] 
**PrivateKey** | Pointer to **string** | The private key string | [optional] 



## Methods


### NewPkiGenerateExportedKeyResponse

`func NewPkiGenerateExportedKeyResponse() *PkiGenerateExportedKeyResponse`

NewPkiGenerateExportedKeyResponse instantiates a new PkiGenerateExportedKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateExportedKeyResponseWithDefaults

`func NewPkiGenerateExportedKeyResponseWithDefaults() *PkiGenerateExportedKeyResponse`

NewPkiGenerateExportedKeyResponseWithDefaults instantiates a new PkiGenerateExportedKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKeyId

`func (o *PkiGenerateExportedKeyResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiGenerateExportedKeyResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiGenerateExportedKeyResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiGenerateExportedKeyResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetKeyName

`func (o *PkiGenerateExportedKeyResponse) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiGenerateExportedKeyResponse) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiGenerateExportedKeyResponse) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### HasKeyName

`func (o *PkiGenerateExportedKeyResponse) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateExportedKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateExportedKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateExportedKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateExportedKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetPrivateKey

`func (o *PkiGenerateExportedKeyResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiGenerateExportedKeyResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiGenerateExportedKeyResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *PkiGenerateExportedKeyResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


