# SshSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** | Type of certificate to be created; either \&quot;user\&quot; or \&quot;host\&quot;. | [optional] [default to "user"]
**CriticalOptions** | Pointer to **map[string]interface{}** | Critical options that the certificate should be signed for. | [optional] 
**Extensions** | Pointer to **map[string]interface{}** | Extensions that the certificate should be signed for. | [optional] 
**KeyId** | Pointer to **string** | Key id that the created certificate should have. If not specified, the display name of the token will be used. | [optional] 
**PublicKey** | Pointer to **string** | SSH public key that should be signed. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL. | [optional] 
**ValidPrincipals** | Pointer to **string** | Valid principals, either usernames or hostnames, that the certificate should be signed for. | [optional] 

## Methods

### NewSshSignRequest

`func NewSshSignRequest() *SshSignRequest`

NewSshSignRequest instantiates a new SshSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSignRequestWithDefaults

`func NewSshSignRequestWithDefaults() *SshSignRequest`

NewSshSignRequestWithDefaults instantiates a new SshSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *SshSignRequest) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *SshSignRequest) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *SshSignRequest) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *SshSignRequest) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCriticalOptions

`func (o *SshSignRequest) GetCriticalOptions() map[string]interface{}`

GetCriticalOptions returns the CriticalOptions field if non-nil, zero value otherwise.

### GetCriticalOptionsOk

`func (o *SshSignRequest) GetCriticalOptionsOk() (*map[string]interface{}, bool)`

GetCriticalOptionsOk returns a tuple with the CriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalOptions

`func (o *SshSignRequest) SetCriticalOptions(v map[string]interface{})`

SetCriticalOptions sets CriticalOptions field to given value.

### HasCriticalOptions

`func (o *SshSignRequest) HasCriticalOptions() bool`

HasCriticalOptions returns a boolean if a field has been set.

### GetExtensions

`func (o *SshSignRequest) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SshSignRequest) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SshSignRequest) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SshSignRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetKeyId

`func (o *SshSignRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *SshSignRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *SshSignRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *SshSignRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetPublicKey

`func (o *SshSignRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshSignRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshSignRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SshSignRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTtl

`func (o *SshSignRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SshSignRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SshSignRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SshSignRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValidPrincipals

`func (o *SshSignRequest) GetValidPrincipals() string`

GetValidPrincipals returns the ValidPrincipals field if non-nil, zero value otherwise.

### GetValidPrincipalsOk

`func (o *SshSignRequest) GetValidPrincipalsOk() (*string, bool)`

GetValidPrincipalsOk returns a tuple with the ValidPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPrincipals

`func (o *SshSignRequest) SetValidPrincipals(v string)`

SetValidPrincipals sets ValidPrincipals field to given value.

### HasValidPrincipals

`func (o *SshSignRequest) HasValidPrincipals() bool`

HasValidPrincipals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


