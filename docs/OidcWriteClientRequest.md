# OidcWriteClientRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtl** | Pointer to **string** | The time-to-live for access tokens obtained by the client. | [optional] [default to "24h"]
**Assignments** | Pointer to **[]string** | Comma separated string or array of assignment resources. | [optional] 
**ClientType** | Pointer to **string** | The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: &#x27;confidential&#x27;, &#x27;public&#x27;. Defaults to &#x27;confidential&#x27;. | [optional] [default to "confidential"]
**IdTokenTtl** | Pointer to **string** | The time-to-live for ID tokens obtained by the client. | [optional] [default to "24h"]
**Key** | Pointer to **string** | A reference to a named key resource. Cannot be modified after creation. Defaults to the &#x27;default&#x27; key. | [optional] [default to "default"]
**RedirectUris** | Pointer to **[]string** | Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request. | [optional] 



## Methods


### NewOidcWriteClientRequest

`func NewOidcWriteClientRequest() *OidcWriteClientRequest`

NewOidcWriteClientRequest instantiates a new OidcWriteClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWriteClientRequestWithDefaults

`func NewOidcWriteClientRequestWithDefaults() *OidcWriteClientRequest`

NewOidcWriteClientRequestWithDefaults instantiates a new OidcWriteClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessTokenTtl

`func (o *OidcWriteClientRequest) GetAccessTokenTtl() string`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *OidcWriteClientRequest) GetAccessTokenTtlOk() (*string, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *OidcWriteClientRequest) SetAccessTokenTtl(v string)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.


### HasAccessTokenTtl

`func (o *OidcWriteClientRequest) HasAccessTokenTtl() bool`

HasAccessTokenTtl returns a boolean if a field has been set.




### GetAssignments

`func (o *OidcWriteClientRequest) GetAssignments() []string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *OidcWriteClientRequest) GetAssignmentsOk() (*[]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *OidcWriteClientRequest) SetAssignments(v []string)`

SetAssignments sets Assignments field to given value.


### HasAssignments

`func (o *OidcWriteClientRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.




### GetClientType

`func (o *OidcWriteClientRequest) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *OidcWriteClientRequest) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *OidcWriteClientRequest) SetClientType(v string)`

SetClientType sets ClientType field to given value.


### HasClientType

`func (o *OidcWriteClientRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.




### GetIdTokenTtl

`func (o *OidcWriteClientRequest) GetIdTokenTtl() string`

GetIdTokenTtl returns the IdTokenTtl field if non-nil, zero value otherwise.

### GetIdTokenTtlOk

`func (o *OidcWriteClientRequest) GetIdTokenTtlOk() (*string, bool)`

GetIdTokenTtlOk returns a tuple with the IdTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenTtl

`func (o *OidcWriteClientRequest) SetIdTokenTtl(v string)`

SetIdTokenTtl sets IdTokenTtl field to given value.


### HasIdTokenTtl

`func (o *OidcWriteClientRequest) HasIdTokenTtl() bool`

HasIdTokenTtl returns a boolean if a field has been set.




### GetKey

`func (o *OidcWriteClientRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OidcWriteClientRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OidcWriteClientRequest) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *OidcWriteClientRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetRedirectUris

`func (o *OidcWriteClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OidcWriteClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OidcWriteClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### HasRedirectUris

`func (o *OidcWriteClientRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


