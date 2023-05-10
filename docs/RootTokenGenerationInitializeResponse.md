# RootTokenGenerationInitializeResponse


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


### NewRootTokenGenerationInitializeResponse

`func NewRootTokenGenerationInitializeResponse() *RootTokenGenerationInitializeResponse`

NewRootTokenGenerationInitializeResponse instantiates a new RootTokenGenerationInitializeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootTokenGenerationInitializeResponseWithDefaults

`func NewRootTokenGenerationInitializeResponseWithDefaults() *RootTokenGenerationInitializeResponse`

NewRootTokenGenerationInitializeResponseWithDefaults instantiates a new RootTokenGenerationInitializeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetComplete

`func (o *RootTokenGenerationInitializeResponse) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *RootTokenGenerationInitializeResponse) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *RootTokenGenerationInitializeResponse) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### HasComplete

`func (o *RootTokenGenerationInitializeResponse) HasComplete() bool`

HasComplete returns a boolean if a field has been set.




### GetEncodedRootToken

`func (o *RootTokenGenerationInitializeResponse) GetEncodedRootToken() string`

GetEncodedRootToken returns the EncodedRootToken field if non-nil, zero value otherwise.

### GetEncodedRootTokenOk

`func (o *RootTokenGenerationInitializeResponse) GetEncodedRootTokenOk() (*string, bool)`

GetEncodedRootTokenOk returns a tuple with the EncodedRootToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedRootToken

`func (o *RootTokenGenerationInitializeResponse) SetEncodedRootToken(v string)`

SetEncodedRootToken sets EncodedRootToken field to given value.


### HasEncodedRootToken

`func (o *RootTokenGenerationInitializeResponse) HasEncodedRootToken() bool`

HasEncodedRootToken returns a boolean if a field has been set.




### GetEncodedToken

`func (o *RootTokenGenerationInitializeResponse) GetEncodedToken() string`

GetEncodedToken returns the EncodedToken field if non-nil, zero value otherwise.

### GetEncodedTokenOk

`func (o *RootTokenGenerationInitializeResponse) GetEncodedTokenOk() (*string, bool)`

GetEncodedTokenOk returns a tuple with the EncodedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedToken

`func (o *RootTokenGenerationInitializeResponse) SetEncodedToken(v string)`

SetEncodedToken sets EncodedToken field to given value.


### HasEncodedToken

`func (o *RootTokenGenerationInitializeResponse) HasEncodedToken() bool`

HasEncodedToken returns a boolean if a field has been set.




### GetNonce

`func (o *RootTokenGenerationInitializeResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *RootTokenGenerationInitializeResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *RootTokenGenerationInitializeResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *RootTokenGenerationInitializeResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetOtp

`func (o *RootTokenGenerationInitializeResponse) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *RootTokenGenerationInitializeResponse) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *RootTokenGenerationInitializeResponse) SetOtp(v string)`

SetOtp sets Otp field to given value.


### HasOtp

`func (o *RootTokenGenerationInitializeResponse) HasOtp() bool`

HasOtp returns a boolean if a field has been set.




### GetOtpLength

`func (o *RootTokenGenerationInitializeResponse) GetOtpLength() int32`

GetOtpLength returns the OtpLength field if non-nil, zero value otherwise.

### GetOtpLengthOk

`func (o *RootTokenGenerationInitializeResponse) GetOtpLengthOk() (*int32, bool)`

GetOtpLengthOk returns a tuple with the OtpLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpLength

`func (o *RootTokenGenerationInitializeResponse) SetOtpLength(v int32)`

SetOtpLength sets OtpLength field to given value.


### HasOtpLength

`func (o *RootTokenGenerationInitializeResponse) HasOtpLength() bool`

HasOtpLength returns a boolean if a field has been set.




### GetPgpFingerprint

`func (o *RootTokenGenerationInitializeResponse) GetPgpFingerprint() string`

GetPgpFingerprint returns the PgpFingerprint field if non-nil, zero value otherwise.

### GetPgpFingerprintOk

`func (o *RootTokenGenerationInitializeResponse) GetPgpFingerprintOk() (*string, bool)`

GetPgpFingerprintOk returns a tuple with the PgpFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpFingerprint

`func (o *RootTokenGenerationInitializeResponse) SetPgpFingerprint(v string)`

SetPgpFingerprint sets PgpFingerprint field to given value.


### HasPgpFingerprint

`func (o *RootTokenGenerationInitializeResponse) HasPgpFingerprint() bool`

HasPgpFingerprint returns a boolean if a field has been set.




### GetProgress

`func (o *RootTokenGenerationInitializeResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RootTokenGenerationInitializeResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RootTokenGenerationInitializeResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *RootTokenGenerationInitializeResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRequired

`func (o *RootTokenGenerationInitializeResponse) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RootTokenGenerationInitializeResponse) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RootTokenGenerationInitializeResponse) SetRequired(v int32)`

SetRequired sets Required field to given value.


### HasRequired

`func (o *RootTokenGenerationInitializeResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.




### GetStarted

`func (o *RootTokenGenerationInitializeResponse) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RootTokenGenerationInitializeResponse) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RootTokenGenerationInitializeResponse) SetStarted(v bool)`

SetStarted sets Started field to given value.


### HasStarted

`func (o *RootTokenGenerationInitializeResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


