# MfaCreateOktaMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** | Okta API key. | [optional] 
**BaseUrl** | Pointer to **string** | The base domain to use for the Okta API. When not specified in the configuration, \&quot;okta.com\&quot; is used. | [optional] 
**MethodName** | Pointer to **string** | The unique name identifier for this MFA method. | [optional] 
**OrgName** | Pointer to **string** | Name of the organization to be used in the Okta API. | [optional] 
**PrimaryEmail** | Pointer to **bool** | If true, the username will only match the primary email for the account. Defaults to false. | [optional] 
**Production** | Pointer to **bool** | (DEPRECATED) Use base_url instead. | [optional] 
**UsernameFormat** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to substitute should be placed in {{}}. For example, \&quot;{{entity.name}}@example.com\&quot;. If blank, the Entity&#x27;s name field will be used as-is. | [optional] 



## Methods


### NewMfaCreateOktaMethodRequest

`func NewMfaCreateOktaMethodRequest() *MfaCreateOktaMethodRequest`

NewMfaCreateOktaMethodRequest instantiates a new MfaCreateOktaMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaCreateOktaMethodRequestWithDefaults

`func NewMfaCreateOktaMethodRequestWithDefaults() *MfaCreateOktaMethodRequest`

NewMfaCreateOktaMethodRequestWithDefaults instantiates a new MfaCreateOktaMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetApiToken

`func (o *MfaCreateOktaMethodRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *MfaCreateOktaMethodRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *MfaCreateOktaMethodRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### HasApiToken

`func (o *MfaCreateOktaMethodRequest) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.




### GetBaseUrl

`func (o *MfaCreateOktaMethodRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *MfaCreateOktaMethodRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *MfaCreateOktaMethodRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### HasBaseUrl

`func (o *MfaCreateOktaMethodRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaCreateOktaMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaCreateOktaMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaCreateOktaMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaCreateOktaMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetOrgName

`func (o *MfaCreateOktaMethodRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MfaCreateOktaMethodRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MfaCreateOktaMethodRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### HasOrgName

`func (o *MfaCreateOktaMethodRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.




### GetPrimaryEmail

`func (o *MfaCreateOktaMethodRequest) GetPrimaryEmail() bool`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *MfaCreateOktaMethodRequest) GetPrimaryEmailOk() (*bool, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *MfaCreateOktaMethodRequest) SetPrimaryEmail(v bool)`

SetPrimaryEmail sets PrimaryEmail field to given value.


### HasPrimaryEmail

`func (o *MfaCreateOktaMethodRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.




### GetProduction

`func (o *MfaCreateOktaMethodRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *MfaCreateOktaMethodRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *MfaCreateOktaMethodRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.


### HasProduction

`func (o *MfaCreateOktaMethodRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MfaCreateOktaMethodRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MfaCreateOktaMethodRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MfaCreateOktaMethodRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MfaCreateOktaMethodRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


