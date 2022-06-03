# PkiKmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyBits** | Pointer to **int32** | The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519. | [optional] [default to 0]
**KeyName** | Pointer to **string** | Optional name to be used for this key | [optional] 
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot; and \&quot;ed25519\&quot; are the only valid values. | [optional] [default to "rsa"]
**ManagedKeyId** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types. | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types. | [optional] 

## Methods

### NewPkiKmsRequest

`func NewPkiKmsRequest() *PkiKmsRequest`

NewPkiKmsRequest instantiates a new PkiKmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiKmsRequestWithDefaults

`func NewPkiKmsRequestWithDefaults() *PkiKmsRequest`

NewPkiKmsRequestWithDefaults instantiates a new PkiKmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyBits

`func (o *PkiKmsRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PkiKmsRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PkiKmsRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *PkiKmsRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyName

`func (o *PkiKmsRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiKmsRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiKmsRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *PkiKmsRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetKeyType

`func (o *PkiKmsRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiKmsRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiKmsRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PkiKmsRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetManagedKeyId

`func (o *PkiKmsRequest) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *PkiKmsRequest) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *PkiKmsRequest) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.

### HasManagedKeyId

`func (o *PkiKmsRequest) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.

### GetManagedKeyName

`func (o *PkiKmsRequest) GetManagedKeyName() string`

GetManagedKeyName returns the ManagedKeyName field if non-nil, zero value otherwise.

### GetManagedKeyNameOk

`func (o *PkiKmsRequest) GetManagedKeyNameOk() (*string, bool)`

GetManagedKeyNameOk returns a tuple with the ManagedKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyName

`func (o *PkiKmsRequest) SetManagedKeyName(v string)`

SetManagedKeyName sets ManagedKeyName field to given value.

### HasManagedKeyName

`func (o *PkiKmsRequest) HasManagedKeyName() bool`

HasManagedKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


