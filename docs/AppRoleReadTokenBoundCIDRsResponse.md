# AppRoleReadTokenBoundCIDRsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any. | [optional] 



## Methods


### NewAppRoleReadTokenBoundCIDRsResponse

`func NewAppRoleReadTokenBoundCIDRsResponse() *AppRoleReadTokenBoundCIDRsResponse`

NewAppRoleReadTokenBoundCIDRsResponse instantiates a new AppRoleReadTokenBoundCIDRsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleReadTokenBoundCIDRsResponseWithDefaults

`func NewAppRoleReadTokenBoundCIDRsResponseWithDefaults() *AppRoleReadTokenBoundCIDRsResponse`

NewAppRoleReadTokenBoundCIDRsResponseWithDefaults instantiates a new AppRoleReadTokenBoundCIDRsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetTokenBoundCidrs

`func (o *AppRoleReadTokenBoundCIDRsResponse) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AppRoleReadTokenBoundCIDRsResponse) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AppRoleReadTokenBoundCIDRsResponse) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *AppRoleReadTokenBoundCIDRsResponse) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


