# LdapMfaConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Enables MFA with given backend (available: duo) | [optional] 

## Methods

### NewLdapMfaConfigRequest

`func NewLdapMfaConfigRequest() *LdapMfaConfigRequest`

NewLdapMfaConfigRequest instantiates a new LdapMfaConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapMfaConfigRequestWithDefaults

`func NewLdapMfaConfigRequestWithDefaults() *LdapMfaConfigRequest`

NewLdapMfaConfigRequestWithDefaults instantiates a new LdapMfaConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LdapMfaConfigRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapMfaConfigRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapMfaConfigRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LdapMfaConfigRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


