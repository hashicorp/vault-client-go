# OidcRotateKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationTtl** | Pointer to **string** | Controls how long the public portion of a key will be available for verification after being rotated. Setting verification_ttl here will override the verification_ttl set on the key. | [optional] 



## Methods


### NewOidcRotateKeyRequest

`func NewOidcRotateKeyRequest() *OidcRotateKeyRequest`

NewOidcRotateKeyRequest instantiates a new OidcRotateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcRotateKeyRequestWithDefaults

`func NewOidcRotateKeyRequestWithDefaults() *OidcRotateKeyRequest`

NewOidcRotateKeyRequestWithDefaults instantiates a new OidcRotateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetVerificationTtl

`func (o *OidcRotateKeyRequest) GetVerificationTtl() string`

GetVerificationTtl returns the VerificationTtl field if non-nil, zero value otherwise.

### GetVerificationTtlOk

`func (o *OidcRotateKeyRequest) GetVerificationTtlOk() (*string, bool)`

GetVerificationTtlOk returns a tuple with the VerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTtl

`func (o *OidcRotateKeyRequest) SetVerificationTtl(v string)`

SetVerificationTtl sets VerificationTtl field to given value.


### HasVerificationTtl

`func (o *OidcRotateKeyRequest) HasVerificationTtl() bool`

HasVerificationTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


