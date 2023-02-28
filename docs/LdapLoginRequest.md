# LDAPLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password for this user. | [optional] 



## Methods


### NewLDAPLoginRequest

`func NewLDAPLoginRequest() *LDAPLoginRequest`

NewLDAPLoginRequest instantiates a new LDAPLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPLoginRequestWithDefaults

`func NewLDAPLoginRequestWithDefaults() *LDAPLoginRequest`

NewLDAPLoginRequestWithDefaults instantiates a new LDAPLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPassword

`func (o *LDAPLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LDAPLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LDAPLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### HasPassword

`func (o *LDAPLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


