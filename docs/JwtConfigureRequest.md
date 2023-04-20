# JwtConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundIssuer** | Pointer to **string** | The value against which to match the &#x27;iss&#x27; claim in a JWT. Optional. | [optional] 
**DefaultRole** | Pointer to **string** | The default role to use if none is provided during login. If not set, a role is required during login. | [optional] 
**JwksCaPem** | Pointer to **string** | The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used. | [optional] 
**JwksUrl** | Pointer to **string** | JWKS URL to use to authenticate signatures. Cannot be used with \&quot;oidc_discovery_url\&quot; or \&quot;jwt_validation_pubkeys\&quot;. | [optional] 
**JwtSupportedAlgs** | Pointer to **[]string** | A list of supported signing algorithms. Defaults to RS256. | [optional] 
**JwtValidationPubkeys** | Pointer to **[]string** | A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with \&quot;jwks_url\&quot; or \&quot;oidc_discovery_url\&quot;. | [optional] 
**NamespaceInState** | Pointer to **bool** | Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs. | [optional] 
**OidcClientId** | Pointer to **string** | The OAuth Client ID configured with your OIDC provider. | [optional] 
**OidcClientSecret** | Pointer to **string** | The OAuth Client Secret configured with your OIDC provider. | [optional] 
**OidcDiscoveryCaPem** | Pointer to **string** | The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used. | [optional] 
**OidcDiscoveryUrl** | Pointer to **string** | OIDC Discovery URL, without any .well-known component (base path). Cannot be used with \&quot;jwks_url\&quot; or \&quot;jwt_validation_pubkeys\&quot;. | [optional] 
**OidcResponseMode** | Pointer to **string** | The response mode to be used in the OAuth2 request. Allowed values are &#x27;query&#x27; and &#x27;form_post&#x27;. | [optional] 
**OidcResponseTypes** | Pointer to **[]string** | The response types to request. Allowed values are &#x27;code&#x27; and &#x27;id_token&#x27;. Defaults to &#x27;code&#x27;. | [optional] 
**ProviderConfig** | Pointer to **map[string]interface{}** | Provider-specific configuration. Optional. | [optional] 



## Methods


### NewJwtConfigureRequest

`func NewJwtConfigureRequest() *JwtConfigureRequest`

NewJwtConfigureRequest instantiates a new JwtConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtConfigureRequestWithDefaults

`func NewJwtConfigureRequestWithDefaults() *JwtConfigureRequest`

NewJwtConfigureRequestWithDefaults instantiates a new JwtConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBoundIssuer

`func (o *JwtConfigureRequest) GetBoundIssuer() string`

GetBoundIssuer returns the BoundIssuer field if non-nil, zero value otherwise.

### GetBoundIssuerOk

`func (o *JwtConfigureRequest) GetBoundIssuerOk() (*string, bool)`

GetBoundIssuerOk returns a tuple with the BoundIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIssuer

`func (o *JwtConfigureRequest) SetBoundIssuer(v string)`

SetBoundIssuer sets BoundIssuer field to given value.


### HasBoundIssuer

`func (o *JwtConfigureRequest) HasBoundIssuer() bool`

HasBoundIssuer returns a boolean if a field has been set.




### GetDefaultRole

`func (o *JwtConfigureRequest) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *JwtConfigureRequest) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *JwtConfigureRequest) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.


### HasDefaultRole

`func (o *JwtConfigureRequest) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.




### GetJwksCaPem

`func (o *JwtConfigureRequest) GetJwksCaPem() string`

GetJwksCaPem returns the JwksCaPem field if non-nil, zero value otherwise.

### GetJwksCaPemOk

`func (o *JwtConfigureRequest) GetJwksCaPemOk() (*string, bool)`

GetJwksCaPemOk returns a tuple with the JwksCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCaPem

`func (o *JwtConfigureRequest) SetJwksCaPem(v string)`

SetJwksCaPem sets JwksCaPem field to given value.


### HasJwksCaPem

`func (o *JwtConfigureRequest) HasJwksCaPem() bool`

HasJwksCaPem returns a boolean if a field has been set.




### GetJwksUrl

`func (o *JwtConfigureRequest) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *JwtConfigureRequest) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *JwtConfigureRequest) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.


### HasJwksUrl

`func (o *JwtConfigureRequest) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.




### GetJwtSupportedAlgs

`func (o *JwtConfigureRequest) GetJwtSupportedAlgs() []string`

GetJwtSupportedAlgs returns the JwtSupportedAlgs field if non-nil, zero value otherwise.

### GetJwtSupportedAlgsOk

`func (o *JwtConfigureRequest) GetJwtSupportedAlgsOk() (*[]string, bool)`

GetJwtSupportedAlgsOk returns a tuple with the JwtSupportedAlgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSupportedAlgs

`func (o *JwtConfigureRequest) SetJwtSupportedAlgs(v []string)`

SetJwtSupportedAlgs sets JwtSupportedAlgs field to given value.


### HasJwtSupportedAlgs

`func (o *JwtConfigureRequest) HasJwtSupportedAlgs() bool`

HasJwtSupportedAlgs returns a boolean if a field has been set.




### GetJwtValidationPubkeys

`func (o *JwtConfigureRequest) GetJwtValidationPubkeys() []string`

GetJwtValidationPubkeys returns the JwtValidationPubkeys field if non-nil, zero value otherwise.

### GetJwtValidationPubkeysOk

`func (o *JwtConfigureRequest) GetJwtValidationPubkeysOk() (*[]string, bool)`

GetJwtValidationPubkeysOk returns a tuple with the JwtValidationPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtValidationPubkeys

`func (o *JwtConfigureRequest) SetJwtValidationPubkeys(v []string)`

SetJwtValidationPubkeys sets JwtValidationPubkeys field to given value.


### HasJwtValidationPubkeys

`func (o *JwtConfigureRequest) HasJwtValidationPubkeys() bool`

HasJwtValidationPubkeys returns a boolean if a field has been set.




### GetNamespaceInState

`func (o *JwtConfigureRequest) GetNamespaceInState() bool`

GetNamespaceInState returns the NamespaceInState field if non-nil, zero value otherwise.

### GetNamespaceInStateOk

`func (o *JwtConfigureRequest) GetNamespaceInStateOk() (*bool, bool)`

GetNamespaceInStateOk returns a tuple with the NamespaceInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceInState

`func (o *JwtConfigureRequest) SetNamespaceInState(v bool)`

SetNamespaceInState sets NamespaceInState field to given value.


### HasNamespaceInState

`func (o *JwtConfigureRequest) HasNamespaceInState() bool`

HasNamespaceInState returns a boolean if a field has been set.




### GetOidcClientId

`func (o *JwtConfigureRequest) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *JwtConfigureRequest) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *JwtConfigureRequest) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.


### HasOidcClientId

`func (o *JwtConfigureRequest) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.




### GetOidcClientSecret

`func (o *JwtConfigureRequest) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *JwtConfigureRequest) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *JwtConfigureRequest) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.


### HasOidcClientSecret

`func (o *JwtConfigureRequest) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.




### GetOidcDiscoveryCaPem

`func (o *JwtConfigureRequest) GetOidcDiscoveryCaPem() string`

GetOidcDiscoveryCaPem returns the OidcDiscoveryCaPem field if non-nil, zero value otherwise.

### GetOidcDiscoveryCaPemOk

`func (o *JwtConfigureRequest) GetOidcDiscoveryCaPemOk() (*string, bool)`

GetOidcDiscoveryCaPemOk returns a tuple with the OidcDiscoveryCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryCaPem

`func (o *JwtConfigureRequest) SetOidcDiscoveryCaPem(v string)`

SetOidcDiscoveryCaPem sets OidcDiscoveryCaPem field to given value.


### HasOidcDiscoveryCaPem

`func (o *JwtConfigureRequest) HasOidcDiscoveryCaPem() bool`

HasOidcDiscoveryCaPem returns a boolean if a field has been set.




### GetOidcDiscoveryUrl

`func (o *JwtConfigureRequest) GetOidcDiscoveryUrl() string`

GetOidcDiscoveryUrl returns the OidcDiscoveryUrl field if non-nil, zero value otherwise.

### GetOidcDiscoveryUrlOk

`func (o *JwtConfigureRequest) GetOidcDiscoveryUrlOk() (*string, bool)`

GetOidcDiscoveryUrlOk returns a tuple with the OidcDiscoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryUrl

`func (o *JwtConfigureRequest) SetOidcDiscoveryUrl(v string)`

SetOidcDiscoveryUrl sets OidcDiscoveryUrl field to given value.


### HasOidcDiscoveryUrl

`func (o *JwtConfigureRequest) HasOidcDiscoveryUrl() bool`

HasOidcDiscoveryUrl returns a boolean if a field has been set.




### GetOidcResponseMode

`func (o *JwtConfigureRequest) GetOidcResponseMode() string`

GetOidcResponseMode returns the OidcResponseMode field if non-nil, zero value otherwise.

### GetOidcResponseModeOk

`func (o *JwtConfigureRequest) GetOidcResponseModeOk() (*string, bool)`

GetOidcResponseModeOk returns a tuple with the OidcResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseMode

`func (o *JwtConfigureRequest) SetOidcResponseMode(v string)`

SetOidcResponseMode sets OidcResponseMode field to given value.


### HasOidcResponseMode

`func (o *JwtConfigureRequest) HasOidcResponseMode() bool`

HasOidcResponseMode returns a boolean if a field has been set.




### GetOidcResponseTypes

`func (o *JwtConfigureRequest) GetOidcResponseTypes() []string`

GetOidcResponseTypes returns the OidcResponseTypes field if non-nil, zero value otherwise.

### GetOidcResponseTypesOk

`func (o *JwtConfigureRequest) GetOidcResponseTypesOk() (*[]string, bool)`

GetOidcResponseTypesOk returns a tuple with the OidcResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseTypes

`func (o *JwtConfigureRequest) SetOidcResponseTypes(v []string)`

SetOidcResponseTypes sets OidcResponseTypes field to given value.


### HasOidcResponseTypes

`func (o *JwtConfigureRequest) HasOidcResponseTypes() bool`

HasOidcResponseTypes returns a boolean if a field has been set.




### GetProviderConfig

`func (o *JwtConfigureRequest) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *JwtConfigureRequest) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *JwtConfigureRequest) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.


### HasProviderConfig

`func (o *JwtConfigureRequest) HasProviderConfig() bool`

HasProviderConfig returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


