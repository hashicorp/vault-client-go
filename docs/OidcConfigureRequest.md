# OidcConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | Issuer URL to be used in the iss claim of the token. If not set, Vault&#x27;s app_addr will be used. | [optional] 



## Methods


### NewOidcConfigureRequest

`func NewOidcConfigureRequest() *OidcConfigureRequest`

NewOidcConfigureRequest instantiates a new OidcConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcConfigureRequestWithDefaults

`func NewOidcConfigureRequestWithDefaults() *OidcConfigureRequest`

NewOidcConfigureRequestWithDefaults instantiates a new OidcConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIssuer

`func (o *OidcConfigureRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OidcConfigureRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OidcConfigureRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### HasIssuer

`func (o *OidcConfigureRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


