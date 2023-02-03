# KerberosLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Authorization** | Pointer to **string** | SPNEGO Authorization header. Required. | [optional] 



## Methods


### NewKerberosLoginRequest

`func NewKerberosLoginRequest() *KerberosLoginRequest`

NewKerberosLoginRequest instantiates a new KerberosLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosLoginRequestWithDefaults

`func NewKerberosLoginRequestWithDefaults() *KerberosLoginRequest`

NewKerberosLoginRequestWithDefaults instantiates a new KerberosLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAuthorization

`func (o *KerberosLoginRequest) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *KerberosLoginRequest) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *KerberosLoginRequest) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.


### HasAuthorization

`func (o *KerberosLoginRequest) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


