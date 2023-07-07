# MfaUpdateOktaMethodRequest


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


### NewMfaUpdateOktaMethodRequest

`func NewMfaUpdateOktaMethodRequest() *MfaUpdateOktaMethodRequest`

NewMfaUpdateOktaMethodRequest instantiates a new MfaUpdateOktaMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaUpdateOktaMethodRequestWithDefaults

`func NewMfaUpdateOktaMethodRequestWithDefaults() *MfaUpdateOktaMethodRequest`

NewMfaUpdateOktaMethodRequestWithDefaults instantiates a new MfaUpdateOktaMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetApiToken

`func (o *MfaUpdateOktaMethodRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *MfaUpdateOktaMethodRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *MfaUpdateOktaMethodRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### HasApiToken

`func (o *MfaUpdateOktaMethodRequest) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.




### GetBaseUrl

`func (o *MfaUpdateOktaMethodRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *MfaUpdateOktaMethodRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *MfaUpdateOktaMethodRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### HasBaseUrl

`func (o *MfaUpdateOktaMethodRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaUpdateOktaMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaUpdateOktaMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaUpdateOktaMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaUpdateOktaMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetOrgName

`func (o *MfaUpdateOktaMethodRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MfaUpdateOktaMethodRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MfaUpdateOktaMethodRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### HasOrgName

`func (o *MfaUpdateOktaMethodRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.




### GetPrimaryEmail

`func (o *MfaUpdateOktaMethodRequest) GetPrimaryEmail() bool`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *MfaUpdateOktaMethodRequest) GetPrimaryEmailOk() (*bool, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *MfaUpdateOktaMethodRequest) SetPrimaryEmail(v bool)`

SetPrimaryEmail sets PrimaryEmail field to given value.


### HasPrimaryEmail

`func (o *MfaUpdateOktaMethodRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.




### GetProduction

`func (o *MfaUpdateOktaMethodRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *MfaUpdateOktaMethodRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *MfaUpdateOktaMethodRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.


### HasProduction

`func (o *MfaUpdateOktaMethodRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MfaUpdateOktaMethodRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MfaUpdateOktaMethodRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MfaUpdateOktaMethodRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MfaUpdateOktaMethodRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


