# RateLimitQuotasWriteRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockInterval** | Pointer to **int32** | If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the &#x27;block_interval&#x27; has elapsed. | [optional] 
**Interval** | Pointer to **int32** | The duration to enforce rate limiting for (default &#x27;1s&#x27;). | [optional] 
**Path** | Pointer to **string** | Path of the mount or namespace to apply the quota. A blank path configures a global quota. For example namespace1/ adds a quota to a full namespace, namespace1/auth/userpass adds a quota to userpass in namespace1. | [optional] 
**Rate** | Pointer to **float32** | The maximum number of requests in a given interval to be allowed by the quota rule. The &#x27;rate&#x27; must be positive. | [optional] 
**Role** | Pointer to **string** | Login role to apply this quota to. Note that when set, path must be configured to a valid auth method with a concept of roles. | [optional] 
**Type** | Pointer to **string** | Type of the quota rule. | [optional] 



## Methods


### NewRateLimitQuotasWriteRequest

`func NewRateLimitQuotasWriteRequest() *RateLimitQuotasWriteRequest`

NewRateLimitQuotasWriteRequest instantiates a new RateLimitQuotasWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitQuotasWriteRequestWithDefaults

`func NewRateLimitQuotasWriteRequestWithDefaults() *RateLimitQuotasWriteRequest`

NewRateLimitQuotasWriteRequestWithDefaults instantiates a new RateLimitQuotasWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBlockInterval

`func (o *RateLimitQuotasWriteRequest) GetBlockInterval() int32`

GetBlockInterval returns the BlockInterval field if non-nil, zero value otherwise.

### GetBlockIntervalOk

`func (o *RateLimitQuotasWriteRequest) GetBlockIntervalOk() (*int32, bool)`

GetBlockIntervalOk returns a tuple with the BlockInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInterval

`func (o *RateLimitQuotasWriteRequest) SetBlockInterval(v int32)`

SetBlockInterval sets BlockInterval field to given value.


### HasBlockInterval

`func (o *RateLimitQuotasWriteRequest) HasBlockInterval() bool`

HasBlockInterval returns a boolean if a field has been set.




### GetInterval

`func (o *RateLimitQuotasWriteRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RateLimitQuotasWriteRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RateLimitQuotasWriteRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### HasInterval

`func (o *RateLimitQuotasWriteRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.




### GetPath

`func (o *RateLimitQuotasWriteRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RateLimitQuotasWriteRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RateLimitQuotasWriteRequest) SetPath(v string)`

SetPath sets Path field to given value.


### HasPath

`func (o *RateLimitQuotasWriteRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.




### GetRate

`func (o *RateLimitQuotasWriteRequest) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimitQuotasWriteRequest) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimitQuotasWriteRequest) SetRate(v float32)`

SetRate sets Rate field to given value.


### HasRate

`func (o *RateLimitQuotasWriteRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.




### GetRole

`func (o *RateLimitQuotasWriteRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RateLimitQuotasWriteRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RateLimitQuotasWriteRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *RateLimitQuotasWriteRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.




### GetType

`func (o *RateLimitQuotasWriteRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateLimitQuotasWriteRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateLimitQuotasWriteRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *RateLimitQuotasWriteRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


