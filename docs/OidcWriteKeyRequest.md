# OidcWriteKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Signing algorithm to use. This will default to RS256. | [optional] [default to "RS256"]
**AllowedClientIds** | Pointer to **[]string** | Comma separated string or array of role client ids allowed to use this key for signing. If empty no roles are allowed. If \&quot;*\&quot; all roles are allowed. | [optional] 
**RotationPeriod** | Pointer to **string** | How often to generate a new keypair. | [optional] [default to "24h"]
**VerificationTtl** | Pointer to **string** | Controls how long the public portion of a key will be available for verification after being rotated. | [optional] [default to "24h"]



## Methods


### NewOidcWriteKeyRequest

`func NewOidcWriteKeyRequest() *OidcWriteKeyRequest`

NewOidcWriteKeyRequest instantiates a new OidcWriteKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWriteKeyRequestWithDefaults

`func NewOidcWriteKeyRequestWithDefaults() *OidcWriteKeyRequest`

NewOidcWriteKeyRequestWithDefaults instantiates a new OidcWriteKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *OidcWriteKeyRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *OidcWriteKeyRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *OidcWriteKeyRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *OidcWriteKeyRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetAllowedClientIds

`func (o *OidcWriteKeyRequest) GetAllowedClientIds() []string`

GetAllowedClientIds returns the AllowedClientIds field if non-nil, zero value otherwise.

### GetAllowedClientIdsOk

`func (o *OidcWriteKeyRequest) GetAllowedClientIdsOk() (*[]string, bool)`

GetAllowedClientIdsOk returns a tuple with the AllowedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientIds

`func (o *OidcWriteKeyRequest) SetAllowedClientIds(v []string)`

SetAllowedClientIds sets AllowedClientIds field to given value.


### HasAllowedClientIds

`func (o *OidcWriteKeyRequest) HasAllowedClientIds() bool`

HasAllowedClientIds returns a boolean if a field has been set.




### GetRotationPeriod

`func (o *OidcWriteKeyRequest) GetRotationPeriod() string`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *OidcWriteKeyRequest) GetRotationPeriodOk() (*string, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *OidcWriteKeyRequest) SetRotationPeriod(v string)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *OidcWriteKeyRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetVerificationTtl

`func (o *OidcWriteKeyRequest) GetVerificationTtl() string`

GetVerificationTtl returns the VerificationTtl field if non-nil, zero value otherwise.

### GetVerificationTtlOk

`func (o *OidcWriteKeyRequest) GetVerificationTtlOk() (*string, bool)`

GetVerificationTtlOk returns a tuple with the VerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTtl

`func (o *OidcWriteKeyRequest) SetVerificationTtl(v string)`

SetVerificationTtl sets VerificationTtl field to given value.


### HasVerificationTtl

`func (o *OidcWriteKeyRequest) HasVerificationTtl() bool`

HasVerificationTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


