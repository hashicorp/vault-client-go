# DecodeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedToken** | Pointer to **string** | Specifies the encoded token (result from generate-root). | [optional] 
**Otp** | Pointer to **string** | Specifies the otp code for decode. | [optional] 



## Methods


### NewDecodeRequest

`func NewDecodeRequest() *DecodeRequest`

NewDecodeRequest instantiates a new DecodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRequestWithDefaults

`func NewDecodeRequestWithDefaults() *DecodeRequest`

NewDecodeRequestWithDefaults instantiates a new DecodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEncodedToken

`func (o *DecodeRequest) GetEncodedToken() string`

GetEncodedToken returns the EncodedToken field if non-nil, zero value otherwise.

### GetEncodedTokenOk

`func (o *DecodeRequest) GetEncodedTokenOk() (*string, bool)`

GetEncodedTokenOk returns a tuple with the EncodedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedToken

`func (o *DecodeRequest) SetEncodedToken(v string)`

SetEncodedToken sets EncodedToken field to given value.


### HasEncodedToken

`func (o *DecodeRequest) HasEncodedToken() bool`

HasEncodedToken returns a boolean if a field has been set.




### GetOtp

`func (o *DecodeRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DecodeRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DecodeRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.


### HasOtp

`func (o *DecodeRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


