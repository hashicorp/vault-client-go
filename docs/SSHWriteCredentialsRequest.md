# SSHWriteCredentialsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Ip** | Pointer to **string** | [Required] IP of the remote host | [optional] 
**Username** | Pointer to **string** | [Optional] Username in remote host | [optional] 



## Methods


### NewSSHWriteCredentialsRequest

`func NewSSHWriteCredentialsRequest() *SSHWriteCredentialsRequest`

NewSSHWriteCredentialsRequest instantiates a new SSHWriteCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHWriteCredentialsRequestWithDefaults

`func NewSSHWriteCredentialsRequestWithDefaults() *SSHWriteCredentialsRequest`

NewSSHWriteCredentialsRequestWithDefaults instantiates a new SSHWriteCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIp

`func (o *SSHWriteCredentialsRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SSHWriteCredentialsRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SSHWriteCredentialsRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### HasIp

`func (o *SSHWriteCredentialsRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.




### GetUsername

`func (o *SSHWriteCredentialsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SSHWriteCredentialsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SSHWriteCredentialsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *SSHWriteCredentialsRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


