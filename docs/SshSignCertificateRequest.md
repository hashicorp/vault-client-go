# SshSignCertificateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** | Type of certificate to be created; either \&quot;user\&quot; or \&quot;host\&quot;. | [optional] [default to "user"]
**CriticalOptions** | Pointer to **map[string]interface{}** | Critical options that the certificate should be signed for. | [optional] 
**Extensions** | Pointer to **map[string]interface{}** | Extensions that the certificate should be signed for. | [optional] 
**KeyId** | Pointer to **string** | Key id that the created certificate should have. If not specified, the display name of the token will be used. | [optional] 
**PublicKey** | Pointer to **string** | SSH public key that should be signed. | [optional] 
**Ttl** | Pointer to **string** | The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL. | [optional] 
**ValidPrincipals** | Pointer to **string** | Valid principals, either usernames or hostnames, that the certificate should be signed for. | [optional] 



## Methods


### NewSshSignCertificateRequest

`func NewSshSignCertificateRequest() *SshSignCertificateRequest`

NewSshSignCertificateRequest instantiates a new SshSignCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSignCertificateRequestWithDefaults

`func NewSshSignCertificateRequestWithDefaults() *SshSignCertificateRequest`

NewSshSignCertificateRequestWithDefaults instantiates a new SshSignCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertType

`func (o *SshSignCertificateRequest) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *SshSignCertificateRequest) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *SshSignCertificateRequest) SetCertType(v string)`

SetCertType sets CertType field to given value.


### HasCertType

`func (o *SshSignCertificateRequest) HasCertType() bool`

HasCertType returns a boolean if a field has been set.




### GetCriticalOptions

`func (o *SshSignCertificateRequest) GetCriticalOptions() map[string]interface{}`

GetCriticalOptions returns the CriticalOptions field if non-nil, zero value otherwise.

### GetCriticalOptionsOk

`func (o *SshSignCertificateRequest) GetCriticalOptionsOk() (*map[string]interface{}, bool)`

GetCriticalOptionsOk returns a tuple with the CriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalOptions

`func (o *SshSignCertificateRequest) SetCriticalOptions(v map[string]interface{})`

SetCriticalOptions sets CriticalOptions field to given value.


### HasCriticalOptions

`func (o *SshSignCertificateRequest) HasCriticalOptions() bool`

HasCriticalOptions returns a boolean if a field has been set.




### GetExtensions

`func (o *SshSignCertificateRequest) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SshSignCertificateRequest) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SshSignCertificateRequest) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.


### HasExtensions

`func (o *SshSignCertificateRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.




### GetKeyId

`func (o *SshSignCertificateRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *SshSignCertificateRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *SshSignCertificateRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *SshSignCertificateRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetPublicKey

`func (o *SshSignCertificateRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshSignCertificateRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshSignCertificateRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### HasPublicKey

`func (o *SshSignCertificateRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.




### GetTtl

`func (o *SshSignCertificateRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SshSignCertificateRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SshSignCertificateRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *SshSignCertificateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetValidPrincipals

`func (o *SshSignCertificateRequest) GetValidPrincipals() string`

GetValidPrincipals returns the ValidPrincipals field if non-nil, zero value otherwise.

### GetValidPrincipalsOk

`func (o *SshSignCertificateRequest) GetValidPrincipalsOk() (*string, bool)`

GetValidPrincipalsOk returns a tuple with the ValidPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPrincipals

`func (o *SshSignCertificateRequest) SetValidPrincipals(v string)`

SetValidPrincipals sets ValidPrincipals field to given value.


### HasValidPrincipals

`func (o *SshSignCertificateRequest) HasValidPrincipals() bool`

HasValidPrincipals returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


