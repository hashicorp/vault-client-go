# SshIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** | Type of certificate to be created; either \&quot;user\&quot; or \&quot;host\&quot;. | [optional] [default to "user"]
**CriticalOptions** | Pointer to **map[string]interface{}** | Critical options that the certificate should be signed for. | [optional] 
**Extensions** | Pointer to **map[string]interface{}** | Extensions that the certificate should be signed for. | [optional] 
**KeyBits** | Pointer to **int32** | Specifies the number of bits to use for the generated keys. | [optional] [default to 0]
**KeyId** | Pointer to **string** | Key id that the created certificate should have. If not specified, the display name of the token will be used. | [optional] 
**KeyType** | Pointer to **string** | Specifies the desired key type; must be &#x60;rsa&#x60;, &#x60;ed25519&#x60; or &#x60;ec&#x60; | [optional] [default to "rsa"]
**Ttl** | Pointer to **int32** | The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL. | [optional] 
**ValidPrincipals** | Pointer to **string** | Valid principals, either usernames or hostnames, that the certificate should be signed for. | [optional] 

## Methods

### NewSshIssueRequest

`func NewSshIssueRequest() *SshIssueRequest`

NewSshIssueRequest instantiates a new SshIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshIssueRequestWithDefaults

`func NewSshIssueRequestWithDefaults() *SshIssueRequest`

NewSshIssueRequestWithDefaults instantiates a new SshIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *SshIssueRequest) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *SshIssueRequest) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *SshIssueRequest) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *SshIssueRequest) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCriticalOptions

`func (o *SshIssueRequest) GetCriticalOptions() map[string]interface{}`

GetCriticalOptions returns the CriticalOptions field if non-nil, zero value otherwise.

### GetCriticalOptionsOk

`func (o *SshIssueRequest) GetCriticalOptionsOk() (*map[string]interface{}, bool)`

GetCriticalOptionsOk returns a tuple with the CriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalOptions

`func (o *SshIssueRequest) SetCriticalOptions(v map[string]interface{})`

SetCriticalOptions sets CriticalOptions field to given value.

### HasCriticalOptions

`func (o *SshIssueRequest) HasCriticalOptions() bool`

HasCriticalOptions returns a boolean if a field has been set.

### GetExtensions

`func (o *SshIssueRequest) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SshIssueRequest) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SshIssueRequest) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SshIssueRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetKeyBits

`func (o *SshIssueRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *SshIssueRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *SshIssueRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *SshIssueRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyId

`func (o *SshIssueRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *SshIssueRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *SshIssueRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *SshIssueRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeyType

`func (o *SshIssueRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshIssueRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshIssueRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *SshIssueRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetTtl

`func (o *SshIssueRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SshIssueRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SshIssueRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SshIssueRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValidPrincipals

`func (o *SshIssueRequest) GetValidPrincipals() string`

GetValidPrincipals returns the ValidPrincipals field if non-nil, zero value otherwise.

### GetValidPrincipalsOk

`func (o *SshIssueRequest) GetValidPrincipalsOk() (*string, bool)`

GetValidPrincipalsOk returns a tuple with the ValidPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPrincipals

`func (o *SshIssueRequest) SetValidPrincipals(v string)`

SetValidPrincipals sets ValidPrincipals field to given value.

### HasValidPrincipals

`func (o *SshIssueRequest) HasValidPrincipals() bool`

HasValidPrincipals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


