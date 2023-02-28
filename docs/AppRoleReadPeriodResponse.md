# AppRoleReadPeriodResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 



## Methods


### NewAppRoleReadPeriodResponse

`func NewAppRoleReadPeriodResponse() *AppRoleReadPeriodResponse`

NewAppRoleReadPeriodResponse instantiates a new AppRoleReadPeriodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleReadPeriodResponseWithDefaults

`func NewAppRoleReadPeriodResponseWithDefaults() *AppRoleReadPeriodResponse`

NewAppRoleReadPeriodResponseWithDefaults instantiates a new AppRoleReadPeriodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPeriod

`func (o *AppRoleReadPeriodResponse) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AppRoleReadPeriodResponse) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AppRoleReadPeriodResponse) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *AppRoleReadPeriodResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *AppRoleReadPeriodResponse) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *AppRoleReadPeriodResponse) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *AppRoleReadPeriodResponse) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *AppRoleReadPeriodResponse) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


