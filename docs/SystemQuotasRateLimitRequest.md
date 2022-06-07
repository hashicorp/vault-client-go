# SystemQuotasRateLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockInterval** | Pointer to **int32** | If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the &#39;block_interval&#39; has elapsed. | [optional] 
**Interval** | Pointer to **int32** | The duration to enforce rate limiting for (default &#39;1s&#39;). | [optional] 
**Path** | Pointer to **string** | Path of the mount or namespace to apply the quota. A blank path configures a global quota. For example namespace1/ adds a quota to a full namespace, namespace1/auth/userpass adds a quota to userpass in namespace1. | [optional] 
**Rate** | Pointer to **float32** | The maximum number of requests in a given interval to be allowed by the quota rule. The &#39;rate&#39; must be positive. | [optional] 
**Type** | Pointer to **string** | Type of the quota rule. | [optional] 

## Methods

### NewSystemQuotasRateLimitRequest

`func NewSystemQuotasRateLimitRequest() *SystemQuotasRateLimitRequest`

NewSystemQuotasRateLimitRequest instantiates a new SystemQuotasRateLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemQuotasRateLimitRequestWithDefaults

`func NewSystemQuotasRateLimitRequestWithDefaults() *SystemQuotasRateLimitRequest`

NewSystemQuotasRateLimitRequestWithDefaults instantiates a new SystemQuotasRateLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockInterval

`func (o *SystemQuotasRateLimitRequest) GetBlockInterval() int32`

GetBlockInterval returns the BlockInterval field if non-nil, zero value otherwise.

### GetBlockIntervalOk

`func (o *SystemQuotasRateLimitRequest) GetBlockIntervalOk() (*int32, bool)`

GetBlockIntervalOk returns a tuple with the BlockInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInterval

`func (o *SystemQuotasRateLimitRequest) SetBlockInterval(v int32)`

SetBlockInterval sets BlockInterval field to given value.

### HasBlockInterval

`func (o *SystemQuotasRateLimitRequest) HasBlockInterval() bool`

HasBlockInterval returns a boolean if a field has been set.

### GetInterval

`func (o *SystemQuotasRateLimitRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SystemQuotasRateLimitRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SystemQuotasRateLimitRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SystemQuotasRateLimitRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPath

`func (o *SystemQuotasRateLimitRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemQuotasRateLimitRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemQuotasRateLimitRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SystemQuotasRateLimitRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRate

`func (o *SystemQuotasRateLimitRequest) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SystemQuotasRateLimitRequest) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SystemQuotasRateLimitRequest) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *SystemQuotasRateLimitRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetType

`func (o *SystemQuotasRateLimitRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemQuotasRateLimitRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemQuotasRateLimitRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemQuotasRateLimitRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


