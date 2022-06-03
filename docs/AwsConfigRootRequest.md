# AwsConfigRootRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access key with permission to create new keys. | [optional] 
**IamEndpoint** | Pointer to **string** | Endpoint to custom IAM server URL | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of retries for recoverable exceptions of AWS APIs | [optional] [default to -1]
**Region** | Pointer to **string** | Region for API calls. | [optional] 
**SecretKey** | Pointer to **string** | Secret key with permission to create new keys. | [optional] 
**StsEndpoint** | Pointer to **string** | Endpoint to custom STS server URL | [optional] 
**UsernameTemplate** | Pointer to **string** | Template to generate custom IAM usernames | [optional] 

## Methods

### NewAwsConfigRootRequest

`func NewAwsConfigRootRequest() *AwsConfigRootRequest`

NewAwsConfigRootRequest instantiates a new AwsConfigRootRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigRootRequestWithDefaults

`func NewAwsConfigRootRequestWithDefaults() *AwsConfigRootRequest`

NewAwsConfigRootRequestWithDefaults instantiates a new AwsConfigRootRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AwsConfigRootRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsConfigRootRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsConfigRootRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AwsConfigRootRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetIamEndpoint

`func (o *AwsConfigRootRequest) GetIamEndpoint() string`

GetIamEndpoint returns the IamEndpoint field if non-nil, zero value otherwise.

### GetIamEndpointOk

`func (o *AwsConfigRootRequest) GetIamEndpointOk() (*string, bool)`

GetIamEndpointOk returns a tuple with the IamEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamEndpoint

`func (o *AwsConfigRootRequest) SetIamEndpoint(v string)`

SetIamEndpoint sets IamEndpoint field to given value.

### HasIamEndpoint

`func (o *AwsConfigRootRequest) HasIamEndpoint() bool`

HasIamEndpoint returns a boolean if a field has been set.

### GetMaxRetries

`func (o *AwsConfigRootRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *AwsConfigRootRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *AwsConfigRootRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *AwsConfigRootRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetRegion

`func (o *AwsConfigRootRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsConfigRootRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsConfigRootRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsConfigRootRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *AwsConfigRootRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsConfigRootRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsConfigRootRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AwsConfigRootRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStsEndpoint

`func (o *AwsConfigRootRequest) GetStsEndpoint() string`

GetStsEndpoint returns the StsEndpoint field if non-nil, zero value otherwise.

### GetStsEndpointOk

`func (o *AwsConfigRootRequest) GetStsEndpointOk() (*string, bool)`

GetStsEndpointOk returns a tuple with the StsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEndpoint

`func (o *AwsConfigRootRequest) SetStsEndpoint(v string)`

SetStsEndpoint sets StsEndpoint field to given value.

### HasStsEndpoint

`func (o *AwsConfigRootRequest) HasStsEndpoint() bool`

HasStsEndpoint returns a boolean if a field has been set.

### GetUsernameTemplate

`func (o *AwsConfigRootRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *AwsConfigRootRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *AwsConfigRootRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.

### HasUsernameTemplate

`func (o *AwsConfigRootRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


