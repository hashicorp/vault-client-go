# TerraformCloudConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address to access Terraform Cloud or Enterprise. Default is \&quot;https://app.terraform.io\&quot;. | [optional] [default to "https://app.terraform.io"]
**BasePath** | Pointer to **string** | The base path for the Terraform Cloud or Enterprise API. Default is \&quot;/api/v2/\&quot;. | [optional] [default to "/api/v2/"]
**Token** | **string** | The token to access Terraform Cloud | 



## Methods


### NewTerraformCloudConfigureRequest

`func NewTerraformCloudConfigureRequest(token string, ) *TerraformCloudConfigureRequest`

NewTerraformCloudConfigureRequest instantiates a new TerraformCloudConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformCloudConfigureRequestWithDefaults

`func NewTerraformCloudConfigureRequestWithDefaults() *TerraformCloudConfigureRequest`

NewTerraformCloudConfigureRequestWithDefaults instantiates a new TerraformCloudConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAddress

`func (o *TerraformCloudConfigureRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TerraformCloudConfigureRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TerraformCloudConfigureRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### HasAddress

`func (o *TerraformCloudConfigureRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.




### GetBasePath

`func (o *TerraformCloudConfigureRequest) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *TerraformCloudConfigureRequest) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *TerraformCloudConfigureRequest) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### HasBasePath

`func (o *TerraformCloudConfigureRequest) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.




### GetToken

`func (o *TerraformCloudConfigureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TerraformCloudConfigureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TerraformCloudConfigureRequest) SetToken(v string)`

SetToken sets Token field to given value.










[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


