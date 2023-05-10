# OidcWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Optional client_id | [optional] 
**Key** | **string** | The OIDC key to use for generating tokens. The specified key must already exist. | 
**Template** | Pointer to **string** | The template string to use for generating tokens. This may be in string-ified JSON or base64 format. | [optional] 
**Ttl** | Pointer to **int32** | TTL of the tokens generated against the role. | [optional] 



## Methods


### NewOidcWriteRoleRequest

`func NewOidcWriteRoleRequest(key string, ) *OidcWriteRoleRequest`

NewOidcWriteRoleRequest instantiates a new OidcWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWriteRoleRequestWithDefaults

`func NewOidcWriteRoleRequestWithDefaults() *OidcWriteRoleRequest`

NewOidcWriteRoleRequestWithDefaults instantiates a new OidcWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OidcWriteRoleRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcWriteRoleRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcWriteRoleRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *OidcWriteRoleRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetKey

`func (o *OidcWriteRoleRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OidcWriteRoleRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OidcWriteRoleRequest) SetKey(v string)`

SetKey sets Key field to given value.





### GetTemplate

`func (o *OidcWriteRoleRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OidcWriteRoleRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OidcWriteRoleRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### HasTemplate

`func (o *OidcWriteRoleRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.




### GetTtl

`func (o *OidcWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *OidcWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *OidcWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *OidcWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


