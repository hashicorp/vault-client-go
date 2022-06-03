# IdentityMfaMethodOktaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** | Okta API key. | [optional] 
**BaseUrl** | Pointer to **string** | The base domain to use for the Okta API. When not specified in the configuration, \&quot;okta.com\&quot; is used. | [optional] 
**MethodId** | Pointer to **string** | The unique identifier for this MFA method. | [optional] 
**OrgName** | Pointer to **string** | Name of the organization to be used in the Okta API. | [optional] 
**PrimaryEmail** | Pointer to **bool** | If true, the username will only match the primary email for the account. Defaults to false. | [optional] 
**Production** | Pointer to **bool** | (DEPRECATED) Use base_url instead. | [optional] 
**UsernameTemplate** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to substitute should be placed in {{}}. For example, \&quot;{{entity.name}}@example.com\&quot;. If blank, the Entity&#39;s name field will be used as-is. | [optional] 

## Methods

### NewIdentityMfaMethodOktaRequest

`func NewIdentityMfaMethodOktaRequest() *IdentityMfaMethodOktaRequest`

NewIdentityMfaMethodOktaRequest instantiates a new IdentityMfaMethodOktaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMfaMethodOktaRequestWithDefaults

`func NewIdentityMfaMethodOktaRequestWithDefaults() *IdentityMfaMethodOktaRequest`

NewIdentityMfaMethodOktaRequestWithDefaults instantiates a new IdentityMfaMethodOktaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *IdentityMfaMethodOktaRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *IdentityMfaMethodOktaRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *IdentityMfaMethodOktaRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *IdentityMfaMethodOktaRequest) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetBaseUrl

`func (o *IdentityMfaMethodOktaRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *IdentityMfaMethodOktaRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *IdentityMfaMethodOktaRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *IdentityMfaMethodOktaRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetMethodId

`func (o *IdentityMfaMethodOktaRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *IdentityMfaMethodOktaRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *IdentityMfaMethodOktaRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.

### HasMethodId

`func (o *IdentityMfaMethodOktaRequest) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.

### GetOrgName

`func (o *IdentityMfaMethodOktaRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *IdentityMfaMethodOktaRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *IdentityMfaMethodOktaRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *IdentityMfaMethodOktaRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *IdentityMfaMethodOktaRequest) GetPrimaryEmail() bool`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *IdentityMfaMethodOktaRequest) GetPrimaryEmailOk() (*bool, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *IdentityMfaMethodOktaRequest) SetPrimaryEmail(v bool)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *IdentityMfaMethodOktaRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetProduction

`func (o *IdentityMfaMethodOktaRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *IdentityMfaMethodOktaRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *IdentityMfaMethodOktaRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *IdentityMfaMethodOktaRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetUsernameTemplate

`func (o *IdentityMfaMethodOktaRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *IdentityMfaMethodOktaRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *IdentityMfaMethodOktaRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.

### HasUsernameTemplate

`func (o *IdentityMfaMethodOktaRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


