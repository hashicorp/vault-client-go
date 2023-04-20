# LdapLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password for this user. | [optional] 



## Methods


### NewLdapLoginRequest

`func NewLdapLoginRequest() *LdapLoginRequest`

NewLdapLoginRequest instantiates a new LdapLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapLoginRequestWithDefaults

`func NewLdapLoginRequestWithDefaults() *LdapLoginRequest`

NewLdapLoginRequestWithDefaults instantiates a new LdapLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPassword

`func (o *LdapLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### HasPassword

`func (o *LdapLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


