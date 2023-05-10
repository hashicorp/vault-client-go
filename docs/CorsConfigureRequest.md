# CorsConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedHeaders** | Pointer to **[]string** | A comma-separated string or array of strings indicating headers that are allowed on cross-origin requests. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | A comma-separated string or array of strings indicating origins that may make cross-origin requests. | [optional] 
**Enable** | Pointer to **bool** | Enables or disables CORS headers on requests. | [optional] 



## Methods


### NewCorsConfigureRequest

`func NewCorsConfigureRequest() *CorsConfigureRequest`

NewCorsConfigureRequest instantiates a new CorsConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsConfigureRequestWithDefaults

`func NewCorsConfigureRequestWithDefaults() *CorsConfigureRequest`

NewCorsConfigureRequestWithDefaults instantiates a new CorsConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedHeaders

`func (o *CorsConfigureRequest) GetAllowedHeaders() []string`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *CorsConfigureRequest) GetAllowedHeadersOk() (*[]string, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *CorsConfigureRequest) SetAllowedHeaders(v []string)`

SetAllowedHeaders sets AllowedHeaders field to given value.


### HasAllowedHeaders

`func (o *CorsConfigureRequest) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.




### GetAllowedOrigins

`func (o *CorsConfigureRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CorsConfigureRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CorsConfigureRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.


### HasAllowedOrigins

`func (o *CorsConfigureRequest) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.




### GetEnable

`func (o *CorsConfigureRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CorsConfigureRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CorsConfigureRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### HasEnable

`func (o *CorsConfigureRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


