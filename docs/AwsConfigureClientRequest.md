# AwsConfigureClientRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AWS Access Key ID for the account used to make AWS API requests. | [optional] [default to ""]
**AllowedStsHeaderValues** | Pointer to **[]string** | List of additional headers that are allowed to be in AWS STS request headers | [optional] 
**Endpoint** | Pointer to **string** | URL to override the default generated endpoint for making AWS EC2 API calls. | [optional] [default to ""]
**IamEndpoint** | Pointer to **string** | URL to override the default generated endpoint for making AWS IAM API calls. | [optional] [default to ""]
**IamServerIdHeaderValue** | Pointer to **string** | Value to require in the X-Vault-AWS-IAM-Server-ID request header | [optional] [default to ""]
**MaxRetries** | Pointer to **int32** | Maximum number of retries for recoverable exceptions of AWS APIs | [optional] [default to -1]
**SecretKey** | Pointer to **string** | AWS Secret Access Key for the account used to make AWS API requests. | [optional] [default to ""]
**StsEndpoint** | Pointer to **string** | URL to override the default generated endpoint for making AWS STS API calls. | [optional] [default to ""]
**StsRegion** | Pointer to **string** | The region ID for the sts_endpoint, if set. | [optional] [default to ""]



## Methods


### NewAwsConfigureClientRequest

`func NewAwsConfigureClientRequest() *AwsConfigureClientRequest`

NewAwsConfigureClientRequest instantiates a new AwsConfigureClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigureClientRequestWithDefaults

`func NewAwsConfigureClientRequestWithDefaults() *AwsConfigureClientRequest`

NewAwsConfigureClientRequestWithDefaults instantiates a new AwsConfigureClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessKey

`func (o *AwsConfigureClientRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsConfigureClientRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsConfigureClientRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### HasAccessKey

`func (o *AwsConfigureClientRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.




### GetAllowedStsHeaderValues

`func (o *AwsConfigureClientRequest) GetAllowedStsHeaderValues() []string`

GetAllowedStsHeaderValues returns the AllowedStsHeaderValues field if non-nil, zero value otherwise.

### GetAllowedStsHeaderValuesOk

`func (o *AwsConfigureClientRequest) GetAllowedStsHeaderValuesOk() (*[]string, bool)`

GetAllowedStsHeaderValuesOk returns a tuple with the AllowedStsHeaderValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedStsHeaderValues

`func (o *AwsConfigureClientRequest) SetAllowedStsHeaderValues(v []string)`

SetAllowedStsHeaderValues sets AllowedStsHeaderValues field to given value.


### HasAllowedStsHeaderValues

`func (o *AwsConfigureClientRequest) HasAllowedStsHeaderValues() bool`

HasAllowedStsHeaderValues returns a boolean if a field has been set.




### GetEndpoint

`func (o *AwsConfigureClientRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AwsConfigureClientRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AwsConfigureClientRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### HasEndpoint

`func (o *AwsConfigureClientRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.




### GetIamEndpoint

`func (o *AwsConfigureClientRequest) GetIamEndpoint() string`

GetIamEndpoint returns the IamEndpoint field if non-nil, zero value otherwise.

### GetIamEndpointOk

`func (o *AwsConfigureClientRequest) GetIamEndpointOk() (*string, bool)`

GetIamEndpointOk returns a tuple with the IamEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamEndpoint

`func (o *AwsConfigureClientRequest) SetIamEndpoint(v string)`

SetIamEndpoint sets IamEndpoint field to given value.


### HasIamEndpoint

`func (o *AwsConfigureClientRequest) HasIamEndpoint() bool`

HasIamEndpoint returns a boolean if a field has been set.




### GetIamServerIdHeaderValue

`func (o *AwsConfigureClientRequest) GetIamServerIdHeaderValue() string`

GetIamServerIdHeaderValue returns the IamServerIdHeaderValue field if non-nil, zero value otherwise.

### GetIamServerIdHeaderValueOk

`func (o *AwsConfigureClientRequest) GetIamServerIdHeaderValueOk() (*string, bool)`

GetIamServerIdHeaderValueOk returns a tuple with the IamServerIdHeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamServerIdHeaderValue

`func (o *AwsConfigureClientRequest) SetIamServerIdHeaderValue(v string)`

SetIamServerIdHeaderValue sets IamServerIdHeaderValue field to given value.


### HasIamServerIdHeaderValue

`func (o *AwsConfigureClientRequest) HasIamServerIdHeaderValue() bool`

HasIamServerIdHeaderValue returns a boolean if a field has been set.




### GetMaxRetries

`func (o *AwsConfigureClientRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *AwsConfigureClientRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *AwsConfigureClientRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### HasMaxRetries

`func (o *AwsConfigureClientRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.




### GetSecretKey

`func (o *AwsConfigureClientRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsConfigureClientRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsConfigureClientRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### HasSecretKey

`func (o *AwsConfigureClientRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.




### GetStsEndpoint

`func (o *AwsConfigureClientRequest) GetStsEndpoint() string`

GetStsEndpoint returns the StsEndpoint field if non-nil, zero value otherwise.

### GetStsEndpointOk

`func (o *AwsConfigureClientRequest) GetStsEndpointOk() (*string, bool)`

GetStsEndpointOk returns a tuple with the StsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEndpoint

`func (o *AwsConfigureClientRequest) SetStsEndpoint(v string)`

SetStsEndpoint sets StsEndpoint field to given value.


### HasStsEndpoint

`func (o *AwsConfigureClientRequest) HasStsEndpoint() bool`

HasStsEndpoint returns a boolean if a field has been set.




### GetStsRegion

`func (o *AwsConfigureClientRequest) GetStsRegion() string`

GetStsRegion returns the StsRegion field if non-nil, zero value otherwise.

### GetStsRegionOk

`func (o *AwsConfigureClientRequest) GetStsRegionOk() (*string, bool)`

GetStsRegionOk returns a tuple with the StsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsRegion

`func (o *AwsConfigureClientRequest) SetStsRegion(v string)`

SetStsRegion sets StsRegion field to given value.


### HasStsRegion

`func (o *AwsConfigureClientRequest) HasStsRegion() bool`

HasStsRegion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


