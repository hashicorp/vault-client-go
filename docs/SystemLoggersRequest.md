# SystemLoggersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Log verbosity level. Supported values (in order of detail) are \&quot;trace\&quot;, \&quot;debug\&quot;, \&quot;info\&quot;, \&quot;warn\&quot;, and \&quot;error\&quot;. | [optional] 

## Methods

### NewSystemLoggersRequest

`func NewSystemLoggersRequest() *SystemLoggersRequest`

NewSystemLoggersRequest instantiates a new SystemLoggersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLoggersRequestWithDefaults

`func NewSystemLoggersRequestWithDefaults() *SystemLoggersRequest`

NewSystemLoggersRequestWithDefaults instantiates a new SystemLoggersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *SystemLoggersRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SystemLoggersRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SystemLoggersRequest) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SystemLoggersRequest) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


