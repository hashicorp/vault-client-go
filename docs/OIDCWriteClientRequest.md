# OIDCWriteClientRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**AccessTokenTtl** | Pointer to **int32** | The time-to-live for access tokens obtained by the client. | [optional] 
**Assignments** | Pointer to **[]string** | Comma separated string or array of assignment resources. | [optional] 
**ClientType** | Pointer to **string** | The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: &#x27;confidential&#x27;, &#x27;public&#x27;. Defaults to &#x27;confidential&#x27;. | [optional] [default to "confidential"]
**IdTokenTtl** | Pointer to **int32** | The time-to-live for ID tokens obtained by the client. | [optional] 
**Key** | Pointer to **string** | A reference to a named key resource. Cannot be modified after creation. Defaults to the &#x27;default&#x27; key. | [optional] [default to "default"]
**RedirectUris** | Pointer to **[]string** | Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request. | [optional] 



## Methods


### NewOIDCWriteClientRequest

`func NewOIDCWriteClientRequest() *OIDCWriteClientRequest`

NewOIDCWriteClientRequest instantiates a new OIDCWriteClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteClientRequestWithDefaults

`func NewOIDCWriteClientRequestWithDefaults() *OIDCWriteClientRequest`

NewOIDCWriteClientRequestWithDefaults instantiates a new OIDCWriteClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessTokenTtl

`func (o *OIDCWriteClientRequest) GetAccessTokenTtl() int32`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *OIDCWriteClientRequest) GetAccessTokenTtlOk() (*int32, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *OIDCWriteClientRequest) SetAccessTokenTtl(v int32)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.


### HasAccessTokenTtl

`func (o *OIDCWriteClientRequest) HasAccessTokenTtl() bool`

HasAccessTokenTtl returns a boolean if a field has been set.




### GetAssignments

`func (o *OIDCWriteClientRequest) GetAssignments() []string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *OIDCWriteClientRequest) GetAssignmentsOk() (*[]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *OIDCWriteClientRequest) SetAssignments(v []string)`

SetAssignments sets Assignments field to given value.


### HasAssignments

`func (o *OIDCWriteClientRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.




### GetClientType

`func (o *OIDCWriteClientRequest) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *OIDCWriteClientRequest) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *OIDCWriteClientRequest) SetClientType(v string)`

SetClientType sets ClientType field to given value.


### HasClientType

`func (o *OIDCWriteClientRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.




### GetIdTokenTtl

`func (o *OIDCWriteClientRequest) GetIdTokenTtl() int32`

GetIdTokenTtl returns the IdTokenTtl field if non-nil, zero value otherwise.

### GetIdTokenTtlOk

`func (o *OIDCWriteClientRequest) GetIdTokenTtlOk() (*int32, bool)`

GetIdTokenTtlOk returns a tuple with the IdTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenTtl

`func (o *OIDCWriteClientRequest) SetIdTokenTtl(v int32)`

SetIdTokenTtl sets IdTokenTtl field to given value.


### HasIdTokenTtl

`func (o *OIDCWriteClientRequest) HasIdTokenTtl() bool`

HasIdTokenTtl returns a boolean if a field has been set.




### GetKey

`func (o *OIDCWriteClientRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OIDCWriteClientRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OIDCWriteClientRequest) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *OIDCWriteClientRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetRedirectUris

`func (o *OIDCWriteClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OIDCWriteClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OIDCWriteClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### HasRedirectUris

`func (o *OIDCWriteClientRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


