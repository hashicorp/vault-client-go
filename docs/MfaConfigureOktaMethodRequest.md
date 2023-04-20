# MfaConfigureOktaMethodRequest


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


### NewMfaConfigureOktaMethodRequest

`func NewMfaConfigureOktaMethodRequest() *MfaConfigureOktaMethodRequest`

NewMfaConfigureOktaMethodRequest instantiates a new MfaConfigureOktaMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaConfigureOktaMethodRequestWithDefaults

`func NewMfaConfigureOktaMethodRequestWithDefaults() *MfaConfigureOktaMethodRequest`

NewMfaConfigureOktaMethodRequestWithDefaults instantiates a new MfaConfigureOktaMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetApiToken

`func (o *MfaConfigureOktaMethodRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *MfaConfigureOktaMethodRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *MfaConfigureOktaMethodRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### HasApiToken

`func (o *MfaConfigureOktaMethodRequest) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.




### GetBaseUrl

`func (o *MfaConfigureOktaMethodRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *MfaConfigureOktaMethodRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *MfaConfigureOktaMethodRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### HasBaseUrl

`func (o *MfaConfigureOktaMethodRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaConfigureOktaMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaConfigureOktaMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaConfigureOktaMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaConfigureOktaMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetOrgName

`func (o *MfaConfigureOktaMethodRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MfaConfigureOktaMethodRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MfaConfigureOktaMethodRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### HasOrgName

`func (o *MfaConfigureOktaMethodRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.




### GetPrimaryEmail

`func (o *MfaConfigureOktaMethodRequest) GetPrimaryEmail() bool`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *MfaConfigureOktaMethodRequest) GetPrimaryEmailOk() (*bool, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *MfaConfigureOktaMethodRequest) SetPrimaryEmail(v bool)`

SetPrimaryEmail sets PrimaryEmail field to given value.


### HasPrimaryEmail

`func (o *MfaConfigureOktaMethodRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.




### GetProduction

`func (o *MfaConfigureOktaMethodRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *MfaConfigureOktaMethodRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *MfaConfigureOktaMethodRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.


### HasProduction

`func (o *MfaConfigureOktaMethodRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MfaConfigureOktaMethodRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MfaConfigureOktaMethodRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MfaConfigureOktaMethodRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MfaConfigureOktaMethodRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


