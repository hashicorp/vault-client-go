# RabbitMqWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **string** | Comma-separated list of tags for this role. | [optional] 
**VhostTopics** | Pointer to **string** | A nested map of virtual hosts and exchanges to topic permissions. | [optional] 
**Vhosts** | Pointer to **string** | A map of virtual hosts to permissions. | [optional] 



## Methods


### NewRabbitMqWriteRoleRequest

`func NewRabbitMqWriteRoleRequest() *RabbitMqWriteRoleRequest`

NewRabbitMqWriteRoleRequest instantiates a new RabbitMqWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRabbitMqWriteRoleRequestWithDefaults

`func NewRabbitMqWriteRoleRequestWithDefaults() *RabbitMqWriteRoleRequest`

NewRabbitMqWriteRoleRequestWithDefaults instantiates a new RabbitMqWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetTags

`func (o *RabbitMqWriteRoleRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RabbitMqWriteRoleRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RabbitMqWriteRoleRequest) SetTags(v string)`

SetTags sets Tags field to given value.


### HasTags

`func (o *RabbitMqWriteRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.




### GetVhostTopics

`func (o *RabbitMqWriteRoleRequest) GetVhostTopics() string`

GetVhostTopics returns the VhostTopics field if non-nil, zero value otherwise.

### GetVhostTopicsOk

`func (o *RabbitMqWriteRoleRequest) GetVhostTopicsOk() (*string, bool)`

GetVhostTopicsOk returns a tuple with the VhostTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhostTopics

`func (o *RabbitMqWriteRoleRequest) SetVhostTopics(v string)`

SetVhostTopics sets VhostTopics field to given value.


### HasVhostTopics

`func (o *RabbitMqWriteRoleRequest) HasVhostTopics() bool`

HasVhostTopics returns a boolean if a field has been set.




### GetVhosts

`func (o *RabbitMqWriteRoleRequest) GetVhosts() string`

GetVhosts returns the Vhosts field if non-nil, zero value otherwise.

### GetVhostsOk

`func (o *RabbitMqWriteRoleRequest) GetVhostsOk() (*string, bool)`

GetVhostsOk returns a tuple with the Vhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhosts

`func (o *RabbitMqWriteRoleRequest) SetVhosts(v string)`

SetVhosts sets Vhosts field to given value.


### HasVhosts

`func (o *RabbitMqWriteRoleRequest) HasVhosts() bool`

HasVhosts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


