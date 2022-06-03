# SshKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | [Required] SSH private key with super user privileges in host | [optional] 

## Methods

### NewSshKeysRequest

`func NewSshKeysRequest() *SshKeysRequest`

NewSshKeysRequest instantiates a new SshKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeysRequestWithDefaults

`func NewSshKeysRequestWithDefaults() *SshKeysRequest`

NewSshKeysRequestWithDefaults instantiates a new SshKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SshKeysRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeysRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeysRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeysRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


