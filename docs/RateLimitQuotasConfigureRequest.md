# RateLimitQuotasConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRateLimitAuditLogging** | Pointer to **bool** | If set, starts audit logging of requests that get rejected due to rate limit quota rule violations. | [optional] 
**EnableRateLimitResponseHeaders** | Pointer to **bool** | If set, additional rate limit quota HTTP headers will be added to responses. | [optional] 
**RateLimitExemptPaths** | Pointer to **[]string** | Specifies the list of exempt paths from all rate limit quotas. If empty no paths will be exempt. | [optional] 



## Methods


### NewRateLimitQuotasConfigureRequest

`func NewRateLimitQuotasConfigureRequest() *RateLimitQuotasConfigureRequest`

NewRateLimitQuotasConfigureRequest instantiates a new RateLimitQuotasConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitQuotasConfigureRequestWithDefaults

`func NewRateLimitQuotasConfigureRequestWithDefaults() *RateLimitQuotasConfigureRequest`

NewRateLimitQuotasConfigureRequestWithDefaults instantiates a new RateLimitQuotasConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEnableRateLimitAuditLogging

`func (o *RateLimitQuotasConfigureRequest) GetEnableRateLimitAuditLogging() bool`

GetEnableRateLimitAuditLogging returns the EnableRateLimitAuditLogging field if non-nil, zero value otherwise.

### GetEnableRateLimitAuditLoggingOk

`func (o *RateLimitQuotasConfigureRequest) GetEnableRateLimitAuditLoggingOk() (*bool, bool)`

GetEnableRateLimitAuditLoggingOk returns a tuple with the EnableRateLimitAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRateLimitAuditLogging

`func (o *RateLimitQuotasConfigureRequest) SetEnableRateLimitAuditLogging(v bool)`

SetEnableRateLimitAuditLogging sets EnableRateLimitAuditLogging field to given value.


### HasEnableRateLimitAuditLogging

`func (o *RateLimitQuotasConfigureRequest) HasEnableRateLimitAuditLogging() bool`

HasEnableRateLimitAuditLogging returns a boolean if a field has been set.




### GetEnableRateLimitResponseHeaders

`func (o *RateLimitQuotasConfigureRequest) GetEnableRateLimitResponseHeaders() bool`

GetEnableRateLimitResponseHeaders returns the EnableRateLimitResponseHeaders field if non-nil, zero value otherwise.

### GetEnableRateLimitResponseHeadersOk

`func (o *RateLimitQuotasConfigureRequest) GetEnableRateLimitResponseHeadersOk() (*bool, bool)`

GetEnableRateLimitResponseHeadersOk returns a tuple with the EnableRateLimitResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRateLimitResponseHeaders

`func (o *RateLimitQuotasConfigureRequest) SetEnableRateLimitResponseHeaders(v bool)`

SetEnableRateLimitResponseHeaders sets EnableRateLimitResponseHeaders field to given value.


### HasEnableRateLimitResponseHeaders

`func (o *RateLimitQuotasConfigureRequest) HasEnableRateLimitResponseHeaders() bool`

HasEnableRateLimitResponseHeaders returns a boolean if a field has been set.




### GetRateLimitExemptPaths

`func (o *RateLimitQuotasConfigureRequest) GetRateLimitExemptPaths() []string`

GetRateLimitExemptPaths returns the RateLimitExemptPaths field if non-nil, zero value otherwise.

### GetRateLimitExemptPathsOk

`func (o *RateLimitQuotasConfigureRequest) GetRateLimitExemptPathsOk() (*[]string, bool)`

GetRateLimitExemptPathsOk returns a tuple with the RateLimitExemptPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitExemptPaths

`func (o *RateLimitQuotasConfigureRequest) SetRateLimitExemptPaths(v []string)`

SetRateLimitExemptPaths sets RateLimitExemptPaths field to given value.


### HasRateLimitExemptPaths

`func (o *RateLimitQuotasConfigureRequest) HasRateLimitExemptPaths() bool`

HasRateLimitExemptPaths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


