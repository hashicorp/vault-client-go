# JWTWriteConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundIssuer** | Pointer to **string** | The value against which to match the &#39;iss&#39; claim in a JWT. Optional. | [optional] 
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
**OidcResponseMode** | Pointer to **string** | The response mode to be used in the OAuth2 request. Allowed values are &#39;query&#39; and &#39;form_post&#39;. | [optional] 
**OidcResponseTypes** | Pointer to **[]string** | The response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to &#39;code&#39;. | [optional] 
**ProviderConfig** | Pointer to **map[string]interface{}** | Provider-specific configuration. Optional. | [optional] 

## Methods

### NewJWTWriteConfigRequest

`func NewJWTWriteConfigRequest() *JWTWriteConfigRequest`

NewJWTWriteConfigRequest instantiates a new JWTWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTWriteConfigRequestWithDefaults

`func NewJWTWriteConfigRequestWithDefaults() *JWTWriteConfigRequest`

NewJWTWriteConfigRequestWithDefaults instantiates a new JWTWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundIssuer

`func (o *JWTWriteConfigRequest) GetBoundIssuer() string`

GetBoundIssuer returns the BoundIssuer field if non-nil, zero value otherwise.

### GetBoundIssuerOk

`func (o *JWTWriteConfigRequest) GetBoundIssuerOk() (*string, bool)`

GetBoundIssuerOk returns a tuple with the BoundIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIssuer

`func (o *JWTWriteConfigRequest) SetBoundIssuer(v string)`

SetBoundIssuer sets BoundIssuer field to given value.

### HasBoundIssuer

`func (o *JWTWriteConfigRequest) HasBoundIssuer() bool`

HasBoundIssuer returns a boolean if a field has been set.

### GetDefaultRole

`func (o *JWTWriteConfigRequest) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *JWTWriteConfigRequest) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *JWTWriteConfigRequest) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *JWTWriteConfigRequest) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetJwksCaPem

`func (o *JWTWriteConfigRequest) GetJwksCaPem() string`

GetJwksCaPem returns the JwksCaPem field if non-nil, zero value otherwise.

### GetJwksCaPemOk

`func (o *JWTWriteConfigRequest) GetJwksCaPemOk() (*string, bool)`

GetJwksCaPemOk returns a tuple with the JwksCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCaPem

`func (o *JWTWriteConfigRequest) SetJwksCaPem(v string)`

SetJwksCaPem sets JwksCaPem field to given value.

### HasJwksCaPem

`func (o *JWTWriteConfigRequest) HasJwksCaPem() bool`

HasJwksCaPem returns a boolean if a field has been set.

### GetJwksUrl

`func (o *JWTWriteConfigRequest) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *JWTWriteConfigRequest) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *JWTWriteConfigRequest) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *JWTWriteConfigRequest) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetJwtSupportedAlgs

`func (o *JWTWriteConfigRequest) GetJwtSupportedAlgs() []string`

GetJwtSupportedAlgs returns the JwtSupportedAlgs field if non-nil, zero value otherwise.

### GetJwtSupportedAlgsOk

`func (o *JWTWriteConfigRequest) GetJwtSupportedAlgsOk() (*[]string, bool)`

GetJwtSupportedAlgsOk returns a tuple with the JwtSupportedAlgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSupportedAlgs

`func (o *JWTWriteConfigRequest) SetJwtSupportedAlgs(v []string)`

SetJwtSupportedAlgs sets JwtSupportedAlgs field to given value.

### HasJwtSupportedAlgs

`func (o *JWTWriteConfigRequest) HasJwtSupportedAlgs() bool`

HasJwtSupportedAlgs returns a boolean if a field has been set.

### GetJwtValidationPubkeys

`func (o *JWTWriteConfigRequest) GetJwtValidationPubkeys() []string`

GetJwtValidationPubkeys returns the JwtValidationPubkeys field if non-nil, zero value otherwise.

### GetJwtValidationPubkeysOk

`func (o *JWTWriteConfigRequest) GetJwtValidationPubkeysOk() (*[]string, bool)`

GetJwtValidationPubkeysOk returns a tuple with the JwtValidationPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtValidationPubkeys

`func (o *JWTWriteConfigRequest) SetJwtValidationPubkeys(v []string)`

SetJwtValidationPubkeys sets JwtValidationPubkeys field to given value.

### HasJwtValidationPubkeys

`func (o *JWTWriteConfigRequest) HasJwtValidationPubkeys() bool`

HasJwtValidationPubkeys returns a boolean if a field has been set.

### GetNamespaceInState

`func (o *JWTWriteConfigRequest) GetNamespaceInState() bool`

GetNamespaceInState returns the NamespaceInState field if non-nil, zero value otherwise.

### GetNamespaceInStateOk

`func (o *JWTWriteConfigRequest) GetNamespaceInStateOk() (*bool, bool)`

GetNamespaceInStateOk returns a tuple with the NamespaceInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceInState

`func (o *JWTWriteConfigRequest) SetNamespaceInState(v bool)`

SetNamespaceInState sets NamespaceInState field to given value.

### HasNamespaceInState

`func (o *JWTWriteConfigRequest) HasNamespaceInState() bool`

HasNamespaceInState returns a boolean if a field has been set.

### GetOidcClientId

`func (o *JWTWriteConfigRequest) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *JWTWriteConfigRequest) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *JWTWriteConfigRequest) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.

### HasOidcClientId

`func (o *JWTWriteConfigRequest) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.

### GetOidcClientSecret

`func (o *JWTWriteConfigRequest) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *JWTWriteConfigRequest) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *JWTWriteConfigRequest) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.

### HasOidcClientSecret

`func (o *JWTWriteConfigRequest) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.

### GetOidcDiscoveryCaPem

`func (o *JWTWriteConfigRequest) GetOidcDiscoveryCaPem() string`

GetOidcDiscoveryCaPem returns the OidcDiscoveryCaPem field if non-nil, zero value otherwise.

### GetOidcDiscoveryCaPemOk

`func (o *JWTWriteConfigRequest) GetOidcDiscoveryCaPemOk() (*string, bool)`

GetOidcDiscoveryCaPemOk returns a tuple with the OidcDiscoveryCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryCaPem

`func (o *JWTWriteConfigRequest) SetOidcDiscoveryCaPem(v string)`

SetOidcDiscoveryCaPem sets OidcDiscoveryCaPem field to given value.

### HasOidcDiscoveryCaPem

`func (o *JWTWriteConfigRequest) HasOidcDiscoveryCaPem() bool`

HasOidcDiscoveryCaPem returns a boolean if a field has been set.

### GetOidcDiscoveryUrl

`func (o *JWTWriteConfigRequest) GetOidcDiscoveryUrl() string`

GetOidcDiscoveryUrl returns the OidcDiscoveryUrl field if non-nil, zero value otherwise.

### GetOidcDiscoveryUrlOk

`func (o *JWTWriteConfigRequest) GetOidcDiscoveryUrlOk() (*string, bool)`

GetOidcDiscoveryUrlOk returns a tuple with the OidcDiscoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryUrl

`func (o *JWTWriteConfigRequest) SetOidcDiscoveryUrl(v string)`

SetOidcDiscoveryUrl sets OidcDiscoveryUrl field to given value.

### HasOidcDiscoveryUrl

`func (o *JWTWriteConfigRequest) HasOidcDiscoveryUrl() bool`

HasOidcDiscoveryUrl returns a boolean if a field has been set.

### GetOidcResponseMode

`func (o *JWTWriteConfigRequest) GetOidcResponseMode() string`

GetOidcResponseMode returns the OidcResponseMode field if non-nil, zero value otherwise.

### GetOidcResponseModeOk

`func (o *JWTWriteConfigRequest) GetOidcResponseModeOk() (*string, bool)`

GetOidcResponseModeOk returns a tuple with the OidcResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseMode

`func (o *JWTWriteConfigRequest) SetOidcResponseMode(v string)`

SetOidcResponseMode sets OidcResponseMode field to given value.

### HasOidcResponseMode

`func (o *JWTWriteConfigRequest) HasOidcResponseMode() bool`

HasOidcResponseMode returns a boolean if a field has been set.

### GetOidcResponseTypes

`func (o *JWTWriteConfigRequest) GetOidcResponseTypes() []string`

GetOidcResponseTypes returns the OidcResponseTypes field if non-nil, zero value otherwise.

### GetOidcResponseTypesOk

`func (o *JWTWriteConfigRequest) GetOidcResponseTypesOk() (*[]string, bool)`

GetOidcResponseTypesOk returns a tuple with the OidcResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseTypes

`func (o *JWTWriteConfigRequest) SetOidcResponseTypes(v []string)`

SetOidcResponseTypes sets OidcResponseTypes field to given value.

### HasOidcResponseTypes

`func (o *JWTWriteConfigRequest) HasOidcResponseTypes() bool`

HasOidcResponseTypes returns a boolean if a field has been set.

### GetProviderConfig

`func (o *JWTWriteConfigRequest) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *JWTWriteConfigRequest) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *JWTWriteConfigRequest) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.

### HasProviderConfig

`func (o *JWTWriteConfigRequest) HasProviderConfig() bool`

HasProviderConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


