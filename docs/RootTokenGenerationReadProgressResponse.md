# RootTokenGenerationReadProgressResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**EncodedRootToken** | Pointer to **string** |  | [optional] 
**EncodedToken** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Otp** | Pointer to **string** |  | [optional] 
**OtpLength** | Pointer to **int32** |  | [optional] 
**PgpFingerprint** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **int32** |  | [optional] 
**Started** | Pointer to **bool** |  | [optional] 



## Methods


### NewRootTokenGenerationReadProgressResponse

`func NewRootTokenGenerationReadProgressResponse() *RootTokenGenerationReadProgressResponse`

NewRootTokenGenerationReadProgressResponse instantiates a new RootTokenGenerationReadProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootTokenGenerationReadProgressResponseWithDefaults

`func NewRootTokenGenerationReadProgressResponseWithDefaults() *RootTokenGenerationReadProgressResponse`

NewRootTokenGenerationReadProgressResponseWithDefaults instantiates a new RootTokenGenerationReadProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetComplete

`func (o *RootTokenGenerationReadProgressResponse) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *RootTokenGenerationReadProgressResponse) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *RootTokenGenerationReadProgressResponse) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### HasComplete

`func (o *RootTokenGenerationReadProgressResponse) HasComplete() bool`

HasComplete returns a boolean if a field has been set.




### GetEncodedRootToken

`func (o *RootTokenGenerationReadProgressResponse) GetEncodedRootToken() string`

GetEncodedRootToken returns the EncodedRootToken field if non-nil, zero value otherwise.

### GetEncodedRootTokenOk

`func (o *RootTokenGenerationReadProgressResponse) GetEncodedRootTokenOk() (*string, bool)`

GetEncodedRootTokenOk returns a tuple with the EncodedRootToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedRootToken

`func (o *RootTokenGenerationReadProgressResponse) SetEncodedRootToken(v string)`

SetEncodedRootToken sets EncodedRootToken field to given value.


### HasEncodedRootToken

`func (o *RootTokenGenerationReadProgressResponse) HasEncodedRootToken() bool`

HasEncodedRootToken returns a boolean if a field has been set.




### GetEncodedToken

`func (o *RootTokenGenerationReadProgressResponse) GetEncodedToken() string`

GetEncodedToken returns the EncodedToken field if non-nil, zero value otherwise.

### GetEncodedTokenOk

`func (o *RootTokenGenerationReadProgressResponse) GetEncodedTokenOk() (*string, bool)`

GetEncodedTokenOk returns a tuple with the EncodedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedToken

`func (o *RootTokenGenerationReadProgressResponse) SetEncodedToken(v string)`

SetEncodedToken sets EncodedToken field to given value.


### HasEncodedToken

`func (o *RootTokenGenerationReadProgressResponse) HasEncodedToken() bool`

HasEncodedToken returns a boolean if a field has been set.




### GetNonce

`func (o *RootTokenGenerationReadProgressResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *RootTokenGenerationReadProgressResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *RootTokenGenerationReadProgressResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *RootTokenGenerationReadProgressResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetOtp

`func (o *RootTokenGenerationReadProgressResponse) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *RootTokenGenerationReadProgressResponse) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *RootTokenGenerationReadProgressResponse) SetOtp(v string)`

SetOtp sets Otp field to given value.


### HasOtp

`func (o *RootTokenGenerationReadProgressResponse) HasOtp() bool`

HasOtp returns a boolean if a field has been set.




### GetOtpLength

`func (o *RootTokenGenerationReadProgressResponse) GetOtpLength() int32`

GetOtpLength returns the OtpLength field if non-nil, zero value otherwise.

### GetOtpLengthOk

`func (o *RootTokenGenerationReadProgressResponse) GetOtpLengthOk() (*int32, bool)`

GetOtpLengthOk returns a tuple with the OtpLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpLength

`func (o *RootTokenGenerationReadProgressResponse) SetOtpLength(v int32)`

SetOtpLength sets OtpLength field to given value.


### HasOtpLength

`func (o *RootTokenGenerationReadProgressResponse) HasOtpLength() bool`

HasOtpLength returns a boolean if a field has been set.




### GetPgpFingerprint

`func (o *RootTokenGenerationReadProgressResponse) GetPgpFingerprint() string`

GetPgpFingerprint returns the PgpFingerprint field if non-nil, zero value otherwise.

### GetPgpFingerprintOk

`func (o *RootTokenGenerationReadProgressResponse) GetPgpFingerprintOk() (*string, bool)`

GetPgpFingerprintOk returns a tuple with the PgpFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpFingerprint

`func (o *RootTokenGenerationReadProgressResponse) SetPgpFingerprint(v string)`

SetPgpFingerprint sets PgpFingerprint field to given value.


### HasPgpFingerprint

`func (o *RootTokenGenerationReadProgressResponse) HasPgpFingerprint() bool`

HasPgpFingerprint returns a boolean if a field has been set.




### GetProgress

`func (o *RootTokenGenerationReadProgressResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RootTokenGenerationReadProgressResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RootTokenGenerationReadProgressResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *RootTokenGenerationReadProgressResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRequired

`func (o *RootTokenGenerationReadProgressResponse) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RootTokenGenerationReadProgressResponse) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RootTokenGenerationReadProgressResponse) SetRequired(v int32)`

SetRequired sets Required field to given value.


### HasRequired

`func (o *RootTokenGenerationReadProgressResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.




### GetStarted

`func (o *RootTokenGenerationReadProgressResponse) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RootTokenGenerationReadProgressResponse) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RootTokenGenerationReadProgressResponse) SetStarted(v bool)`

SetStarted sets Started field to given value.


### HasStarted

`func (o *RootTokenGenerationReadProgressResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


