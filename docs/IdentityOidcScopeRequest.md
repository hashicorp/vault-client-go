# IdentityOidcScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the scope | [optional] 
**Template** | Pointer to **string** | The template string to use for the scope. This may be in string-ified JSON or base64 format. | [optional] 

## Methods

### NewIdentityOidcScopeRequest

`func NewIdentityOidcScopeRequest() *IdentityOidcScopeRequest`

NewIdentityOidcScopeRequest instantiates a new IdentityOidcScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOidcScopeRequestWithDefaults

`func NewIdentityOidcScopeRequestWithDefaults() *IdentityOidcScopeRequest`

NewIdentityOidcScopeRequestWithDefaults instantiates a new IdentityOidcScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityOidcScopeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityOidcScopeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityOidcScopeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityOidcScopeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplate

`func (o *IdentityOidcScopeRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IdentityOidcScopeRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IdentityOidcScopeRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IdentityOidcScopeRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


