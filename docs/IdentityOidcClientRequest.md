# IdentityOidcClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtl** | Pointer to **int32** | The time-to-live for access tokens obtained by the client. | [optional] 
**Assignments** | Pointer to **[]string** | Comma separated string or array of assignment resources. | [optional] 
**ClientType** | Pointer to **string** | The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: &#39;confidential&#39;, &#39;public&#39;. Defaults to &#39;confidential&#39;. | [optional] [default to "confidential"]
**IdTokenTtl** | Pointer to **int32** | The time-to-live for ID tokens obtained by the client. | [optional] 
**Key** | Pointer to **string** | A reference to a named key resource. Cannot be modified after creation. Defaults to the &#39;default&#39; key. | [optional] [default to "default"]
**RedirectUris** | Pointer to **[]string** | Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request. | [optional] 

## Methods

### NewIdentityOidcClientRequest

`func NewIdentityOidcClientRequest() *IdentityOidcClientRequest`

NewIdentityOidcClientRequest instantiates a new IdentityOidcClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOidcClientRequestWithDefaults

`func NewIdentityOidcClientRequestWithDefaults() *IdentityOidcClientRequest`

NewIdentityOidcClientRequestWithDefaults instantiates a new IdentityOidcClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtl

`func (o *IdentityOidcClientRequest) GetAccessTokenTtl() int32`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *IdentityOidcClientRequest) GetAccessTokenTtlOk() (*int32, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *IdentityOidcClientRequest) SetAccessTokenTtl(v int32)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.

### HasAccessTokenTtl

`func (o *IdentityOidcClientRequest) HasAccessTokenTtl() bool`

HasAccessTokenTtl returns a boolean if a field has been set.

### GetAssignments

`func (o *IdentityOidcClientRequest) GetAssignments() []string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *IdentityOidcClientRequest) GetAssignmentsOk() (*[]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *IdentityOidcClientRequest) SetAssignments(v []string)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *IdentityOidcClientRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetClientType

`func (o *IdentityOidcClientRequest) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *IdentityOidcClientRequest) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *IdentityOidcClientRequest) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *IdentityOidcClientRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetIdTokenTtl

`func (o *IdentityOidcClientRequest) GetIdTokenTtl() int32`

GetIdTokenTtl returns the IdTokenTtl field if non-nil, zero value otherwise.

### GetIdTokenTtlOk

`func (o *IdentityOidcClientRequest) GetIdTokenTtlOk() (*int32, bool)`

GetIdTokenTtlOk returns a tuple with the IdTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenTtl

`func (o *IdentityOidcClientRequest) SetIdTokenTtl(v int32)`

SetIdTokenTtl sets IdTokenTtl field to given value.

### HasIdTokenTtl

`func (o *IdentityOidcClientRequest) HasIdTokenTtl() bool`

HasIdTokenTtl returns a boolean if a field has been set.

### GetKey

`func (o *IdentityOidcClientRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentityOidcClientRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentityOidcClientRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IdentityOidcClientRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IdentityOidcClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IdentityOidcClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IdentityOidcClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IdentityOidcClientRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


