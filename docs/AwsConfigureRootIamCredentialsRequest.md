# AwsConfigureRootIamCredentialsRequest


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


### NewAwsConfigureRootIamCredentialsRequest

`func NewAwsConfigureRootIamCredentialsRequest() *AwsConfigureRootIamCredentialsRequest`

NewAwsConfigureRootIamCredentialsRequest instantiates a new AwsConfigureRootIamCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigureRootIamCredentialsRequestWithDefaults

`func NewAwsConfigureRootIamCredentialsRequestWithDefaults() *AwsConfigureRootIamCredentialsRequest`

NewAwsConfigureRootIamCredentialsRequestWithDefaults instantiates a new AwsConfigureRootIamCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessKey

`func (o *AwsConfigureRootIamCredentialsRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsConfigureRootIamCredentialsRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### HasAccessKey

`func (o *AwsConfigureRootIamCredentialsRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.




### GetIamEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) GetIamEndpoint() string`

GetIamEndpoint returns the IamEndpoint field if non-nil, zero value otherwise.

### GetIamEndpointOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetIamEndpointOk() (*string, bool)`

GetIamEndpointOk returns a tuple with the IamEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) SetIamEndpoint(v string)`

SetIamEndpoint sets IamEndpoint field to given value.


### HasIamEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) HasIamEndpoint() bool`

HasIamEndpoint returns a boolean if a field has been set.




### GetMaxRetries

`func (o *AwsConfigureRootIamCredentialsRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *AwsConfigureRootIamCredentialsRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### HasMaxRetries

`func (o *AwsConfigureRootIamCredentialsRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.




### GetRegion

`func (o *AwsConfigureRootIamCredentialsRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsConfigureRootIamCredentialsRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### HasRegion

`func (o *AwsConfigureRootIamCredentialsRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.




### GetSecretKey

`func (o *AwsConfigureRootIamCredentialsRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsConfigureRootIamCredentialsRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### HasSecretKey

`func (o *AwsConfigureRootIamCredentialsRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.




### GetStsEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) GetStsEndpoint() string`

GetStsEndpoint returns the StsEndpoint field if non-nil, zero value otherwise.

### GetStsEndpointOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetStsEndpointOk() (*string, bool)`

GetStsEndpointOk returns a tuple with the StsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) SetStsEndpoint(v string)`

SetStsEndpoint sets StsEndpoint field to given value.


### HasStsEndpoint

`func (o *AwsConfigureRootIamCredentialsRequest) HasStsEndpoint() bool`

HasStsEndpoint returns a boolean if a field has been set.




### GetUsernameTemplate

`func (o *AwsConfigureRootIamCredentialsRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *AwsConfigureRootIamCredentialsRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *AwsConfigureRootIamCredentialsRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.


### HasUsernameTemplate

`func (o *AwsConfigureRootIamCredentialsRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


