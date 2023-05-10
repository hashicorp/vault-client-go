# InternalClientActivityConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultReportMonths** | Pointer to **int32** | Number of months to report if no start date specified. | [optional] [default to 12]
**Enabled** | Pointer to **string** | Enable or disable collection of client count: enable, disable, or default. | [optional] [default to "default"]
**RetentionMonths** | Pointer to **int32** | Number of months of client data to retain. Setting to 0 will clear all existing data. | [optional] [default to 24]



## Methods


### NewInternalClientActivityConfigureRequest

`func NewInternalClientActivityConfigureRequest() *InternalClientActivityConfigureRequest`

NewInternalClientActivityConfigureRequest instantiates a new InternalClientActivityConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalClientActivityConfigureRequestWithDefaults

`func NewInternalClientActivityConfigureRequestWithDefaults() *InternalClientActivityConfigureRequest`

NewInternalClientActivityConfigureRequestWithDefaults instantiates a new InternalClientActivityConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDefaultReportMonths

`func (o *InternalClientActivityConfigureRequest) GetDefaultReportMonths() int32`

GetDefaultReportMonths returns the DefaultReportMonths field if non-nil, zero value otherwise.

### GetDefaultReportMonthsOk

`func (o *InternalClientActivityConfigureRequest) GetDefaultReportMonthsOk() (*int32, bool)`

GetDefaultReportMonthsOk returns a tuple with the DefaultReportMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReportMonths

`func (o *InternalClientActivityConfigureRequest) SetDefaultReportMonths(v int32)`

SetDefaultReportMonths sets DefaultReportMonths field to given value.


### HasDefaultReportMonths

`func (o *InternalClientActivityConfigureRequest) HasDefaultReportMonths() bool`

HasDefaultReportMonths returns a boolean if a field has been set.




### GetEnabled

`func (o *InternalClientActivityConfigureRequest) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InternalClientActivityConfigureRequest) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InternalClientActivityConfigureRequest) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *InternalClientActivityConfigureRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.




### GetRetentionMonths

`func (o *InternalClientActivityConfigureRequest) GetRetentionMonths() int32`

GetRetentionMonths returns the RetentionMonths field if non-nil, zero value otherwise.

### GetRetentionMonthsOk

`func (o *InternalClientActivityConfigureRequest) GetRetentionMonthsOk() (*int32, bool)`

GetRetentionMonthsOk returns a tuple with the RetentionMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMonths

`func (o *InternalClientActivityConfigureRequest) SetRetentionMonths(v int32)`

SetRetentionMonths sets RetentionMonths field to given value.


### HasRetentionMonths

`func (o *InternalClientActivityConfigureRequest) HasRetentionMonths() bool`

HasRetentionMonths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


