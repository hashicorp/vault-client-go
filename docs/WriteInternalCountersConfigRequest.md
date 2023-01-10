# WriteInternalCountersConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultReportMonths** | Pointer to **int32** | Number of months to report if no start date specified. | [optional] [default to 12]
**Enabled** | Pointer to **string** | Enable or disable collection of client count: enable, disable, or default. | [optional] [default to "default"]
**RetentionMonths** | Pointer to **int32** | Number of months of client data to retain. Setting to 0 will clear all existing data. | [optional] [default to 24]

## Methods

### NewWriteInternalCountersConfigRequest

`func NewWriteInternalCountersConfigRequest() *WriteInternalCountersConfigRequest`

NewWriteInternalCountersConfigRequest instantiates a new WriteInternalCountersConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteInternalCountersConfigRequestWithDefaults

`func NewWriteInternalCountersConfigRequestWithDefaults() *WriteInternalCountersConfigRequest`

NewWriteInternalCountersConfigRequestWithDefaults instantiates a new WriteInternalCountersConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultReportMonths

`func (o *WriteInternalCountersConfigRequest) GetDefaultReportMonths() int32`

GetDefaultReportMonths returns the DefaultReportMonths field if non-nil, zero value otherwise.

### GetDefaultReportMonthsOk

`func (o *WriteInternalCountersConfigRequest) GetDefaultReportMonthsOk() (*int32, bool)`

GetDefaultReportMonthsOk returns a tuple with the DefaultReportMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReportMonths

`func (o *WriteInternalCountersConfigRequest) SetDefaultReportMonths(v int32)`

SetDefaultReportMonths sets DefaultReportMonths field to given value.

### HasDefaultReportMonths

`func (o *WriteInternalCountersConfigRequest) HasDefaultReportMonths() bool`

HasDefaultReportMonths returns a boolean if a field has been set.

### GetEnabled

`func (o *WriteInternalCountersConfigRequest) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WriteInternalCountersConfigRequest) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WriteInternalCountersConfigRequest) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WriteInternalCountersConfigRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRetentionMonths

`func (o *WriteInternalCountersConfigRequest) GetRetentionMonths() int32`

GetRetentionMonths returns the RetentionMonths field if non-nil, zero value otherwise.

### GetRetentionMonthsOk

`func (o *WriteInternalCountersConfigRequest) GetRetentionMonthsOk() (*int32, bool)`

GetRetentionMonthsOk returns a tuple with the RetentionMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMonths

`func (o *WriteInternalCountersConfigRequest) SetRetentionMonths(v int32)`

SetRetentionMonths sets RetentionMonths field to given value.

### HasRetentionMonths

`func (o *WriteInternalCountersConfigRequest) HasRetentionMonths() bool`

HasRetentionMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


