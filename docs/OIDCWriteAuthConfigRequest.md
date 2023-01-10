# OIDCWriteAuthConfigRequest

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

### NewOIDCWriteAuthConfigRequest

`func NewOIDCWriteAuthConfigRequest() *OIDCWriteAuthConfigRequest`

NewOIDCWriteAuthConfigRequest instantiates a new OIDCWriteAuthConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteAuthConfigRequestWithDefaults

`func NewOIDCWriteAuthConfigRequestWithDefaults() *OIDCWriteAuthConfigRequest`

NewOIDCWriteAuthConfigRequestWithDefaults instantiates a new OIDCWriteAuthConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundIssuer

`func (o *OIDCWriteAuthConfigRequest) GetBoundIssuer() string`

GetBoundIssuer returns the BoundIssuer field if non-nil, zero value otherwise.

### GetBoundIssuerOk

`func (o *OIDCWriteAuthConfigRequest) GetBoundIssuerOk() (*string, bool)`

GetBoundIssuerOk returns a tuple with the BoundIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIssuer

`func (o *OIDCWriteAuthConfigRequest) SetBoundIssuer(v string)`

SetBoundIssuer sets BoundIssuer field to given value.

### HasBoundIssuer

`func (o *OIDCWriteAuthConfigRequest) HasBoundIssuer() bool`

HasBoundIssuer returns a boolean if a field has been set.

### GetDefaultRole

`func (o *OIDCWriteAuthConfigRequest) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *OIDCWriteAuthConfigRequest) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *OIDCWriteAuthConfigRequest) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *OIDCWriteAuthConfigRequest) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetJwksCaPem

`func (o *OIDCWriteAuthConfigRequest) GetJwksCaPem() string`

GetJwksCaPem returns the JwksCaPem field if non-nil, zero value otherwise.

### GetJwksCaPemOk

`func (o *OIDCWriteAuthConfigRequest) GetJwksCaPemOk() (*string, bool)`

GetJwksCaPemOk returns a tuple with the JwksCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCaPem

`func (o *OIDCWriteAuthConfigRequest) SetJwksCaPem(v string)`

SetJwksCaPem sets JwksCaPem field to given value.

### HasJwksCaPem

`func (o *OIDCWriteAuthConfigRequest) HasJwksCaPem() bool`

HasJwksCaPem returns a boolean if a field has been set.

### GetJwksUrl

`func (o *OIDCWriteAuthConfigRequest) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *OIDCWriteAuthConfigRequest) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *OIDCWriteAuthConfigRequest) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *OIDCWriteAuthConfigRequest) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetJwtSupportedAlgs

`func (o *OIDCWriteAuthConfigRequest) GetJwtSupportedAlgs() []string`

GetJwtSupportedAlgs returns the JwtSupportedAlgs field if non-nil, zero value otherwise.

### GetJwtSupportedAlgsOk

`func (o *OIDCWriteAuthConfigRequest) GetJwtSupportedAlgsOk() (*[]string, bool)`

GetJwtSupportedAlgsOk returns a tuple with the JwtSupportedAlgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSupportedAlgs

`func (o *OIDCWriteAuthConfigRequest) SetJwtSupportedAlgs(v []string)`

SetJwtSupportedAlgs sets JwtSupportedAlgs field to given value.

### HasJwtSupportedAlgs

`func (o *OIDCWriteAuthConfigRequest) HasJwtSupportedAlgs() bool`

HasJwtSupportedAlgs returns a boolean if a field has been set.

### GetJwtValidationPubkeys

`func (o *OIDCWriteAuthConfigRequest) GetJwtValidationPubkeys() []string`

GetJwtValidationPubkeys returns the JwtValidationPubkeys field if non-nil, zero value otherwise.

### GetJwtValidationPubkeysOk

`func (o *OIDCWriteAuthConfigRequest) GetJwtValidationPubkeysOk() (*[]string, bool)`

GetJwtValidationPubkeysOk returns a tuple with the JwtValidationPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtValidationPubkeys

`func (o *OIDCWriteAuthConfigRequest) SetJwtValidationPubkeys(v []string)`

SetJwtValidationPubkeys sets JwtValidationPubkeys field to given value.

### HasJwtValidationPubkeys

`func (o *OIDCWriteAuthConfigRequest) HasJwtValidationPubkeys() bool`

HasJwtValidationPubkeys returns a boolean if a field has been set.

### GetNamespaceInState

`func (o *OIDCWriteAuthConfigRequest) GetNamespaceInState() bool`

GetNamespaceInState returns the NamespaceInState field if non-nil, zero value otherwise.

### GetNamespaceInStateOk

`func (o *OIDCWriteAuthConfigRequest) GetNamespaceInStateOk() (*bool, bool)`

GetNamespaceInStateOk returns a tuple with the NamespaceInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceInState

`func (o *OIDCWriteAuthConfigRequest) SetNamespaceInState(v bool)`

SetNamespaceInState sets NamespaceInState field to given value.

### HasNamespaceInState

`func (o *OIDCWriteAuthConfigRequest) HasNamespaceInState() bool`

HasNamespaceInState returns a boolean if a field has been set.

### GetOidcClientId

`func (o *OIDCWriteAuthConfigRequest) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *OIDCWriteAuthConfigRequest) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.

### HasOidcClientId

`func (o *OIDCWriteAuthConfigRequest) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.

### GetOidcClientSecret

`func (o *OIDCWriteAuthConfigRequest) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *OIDCWriteAuthConfigRequest) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.

### HasOidcClientSecret

`func (o *OIDCWriteAuthConfigRequest) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.

### GetOidcDiscoveryCaPem

`func (o *OIDCWriteAuthConfigRequest) GetOidcDiscoveryCaPem() string`

GetOidcDiscoveryCaPem returns the OidcDiscoveryCaPem field if non-nil, zero value otherwise.

### GetOidcDiscoveryCaPemOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcDiscoveryCaPemOk() (*string, bool)`

GetOidcDiscoveryCaPemOk returns a tuple with the OidcDiscoveryCaPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryCaPem

`func (o *OIDCWriteAuthConfigRequest) SetOidcDiscoveryCaPem(v string)`

SetOidcDiscoveryCaPem sets OidcDiscoveryCaPem field to given value.

### HasOidcDiscoveryCaPem

`func (o *OIDCWriteAuthConfigRequest) HasOidcDiscoveryCaPem() bool`

HasOidcDiscoveryCaPem returns a boolean if a field has been set.

### GetOidcDiscoveryUrl

`func (o *OIDCWriteAuthConfigRequest) GetOidcDiscoveryUrl() string`

GetOidcDiscoveryUrl returns the OidcDiscoveryUrl field if non-nil, zero value otherwise.

### GetOidcDiscoveryUrlOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcDiscoveryUrlOk() (*string, bool)`

GetOidcDiscoveryUrlOk returns a tuple with the OidcDiscoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcDiscoveryUrl

`func (o *OIDCWriteAuthConfigRequest) SetOidcDiscoveryUrl(v string)`

SetOidcDiscoveryUrl sets OidcDiscoveryUrl field to given value.

### HasOidcDiscoveryUrl

`func (o *OIDCWriteAuthConfigRequest) HasOidcDiscoveryUrl() bool`

HasOidcDiscoveryUrl returns a boolean if a field has been set.

### GetOidcResponseMode

`func (o *OIDCWriteAuthConfigRequest) GetOidcResponseMode() string`

GetOidcResponseMode returns the OidcResponseMode field if non-nil, zero value otherwise.

### GetOidcResponseModeOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcResponseModeOk() (*string, bool)`

GetOidcResponseModeOk returns a tuple with the OidcResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseMode

`func (o *OIDCWriteAuthConfigRequest) SetOidcResponseMode(v string)`

SetOidcResponseMode sets OidcResponseMode field to given value.

### HasOidcResponseMode

`func (o *OIDCWriteAuthConfigRequest) HasOidcResponseMode() bool`

HasOidcResponseMode returns a boolean if a field has been set.

### GetOidcResponseTypes

`func (o *OIDCWriteAuthConfigRequest) GetOidcResponseTypes() []string`

GetOidcResponseTypes returns the OidcResponseTypes field if non-nil, zero value otherwise.

### GetOidcResponseTypesOk

`func (o *OIDCWriteAuthConfigRequest) GetOidcResponseTypesOk() (*[]string, bool)`

GetOidcResponseTypesOk returns a tuple with the OidcResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcResponseTypes

`func (o *OIDCWriteAuthConfigRequest) SetOidcResponseTypes(v []string)`

SetOidcResponseTypes sets OidcResponseTypes field to given value.

### HasOidcResponseTypes

`func (o *OIDCWriteAuthConfigRequest) HasOidcResponseTypes() bool`

HasOidcResponseTypes returns a boolean if a field has been set.

### GetProviderConfig

`func (o *OIDCWriteAuthConfigRequest) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *OIDCWriteAuthConfigRequest) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *OIDCWriteAuthConfigRequest) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.

### HasProviderConfig

`func (o *OIDCWriteAuthConfigRequest) HasProviderConfig() bool`

HasProviderConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


