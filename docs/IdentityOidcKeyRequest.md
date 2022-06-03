# IdentityOidcKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Signing algorithm to use. This will default to RS256. | [optional] [default to "RS256"]
**AllowedClientIds** | Pointer to **[]string** | Comma separated string or array of role client ids allowed to use this key for signing. If empty no roles are allowed. If \&quot;*\&quot; all roles are allowed. | [optional] 
**RotationPeriod** | Pointer to **int32** | How often to generate a new keypair. | [optional] 
**VerificationTtl** | Pointer to **int32** | Controls how long the public portion of a key will be available for verification after being rotated. | [optional] 

## Methods

### NewIdentityOidcKeyRequest

`func NewIdentityOidcKeyRequest() *IdentityOidcKeyRequest`

NewIdentityOidcKeyRequest instantiates a new IdentityOidcKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOidcKeyRequestWithDefaults

`func NewIdentityOidcKeyRequestWithDefaults() *IdentityOidcKeyRequest`

NewIdentityOidcKeyRequestWithDefaults instantiates a new IdentityOidcKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *IdentityOidcKeyRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *IdentityOidcKeyRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *IdentityOidcKeyRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *IdentityOidcKeyRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetAllowedClientIds

`func (o *IdentityOidcKeyRequest) GetAllowedClientIds() []string`

GetAllowedClientIds returns the AllowedClientIds field if non-nil, zero value otherwise.

### GetAllowedClientIdsOk

`func (o *IdentityOidcKeyRequest) GetAllowedClientIdsOk() (*[]string, bool)`

GetAllowedClientIdsOk returns a tuple with the AllowedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientIds

`func (o *IdentityOidcKeyRequest) SetAllowedClientIds(v []string)`

SetAllowedClientIds sets AllowedClientIds field to given value.

### HasAllowedClientIds

`func (o *IdentityOidcKeyRequest) HasAllowedClientIds() bool`

HasAllowedClientIds returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *IdentityOidcKeyRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *IdentityOidcKeyRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *IdentityOidcKeyRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *IdentityOidcKeyRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### GetVerificationTtl

`func (o *IdentityOidcKeyRequest) GetVerificationTtl() int32`

GetVerificationTtl returns the VerificationTtl field if non-nil, zero value otherwise.

### GetVerificationTtlOk

`func (o *IdentityOidcKeyRequest) GetVerificationTtlOk() (*int32, bool)`

GetVerificationTtlOk returns a tuple with the VerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTtl

`func (o *IdentityOidcKeyRequest) SetVerificationTtl(v int32)`

SetVerificationTtl sets VerificationTtl field to given value.

### HasVerificationTtl

`func (o *IdentityOidcKeyRequest) HasVerificationTtl() bool`

HasVerificationTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


