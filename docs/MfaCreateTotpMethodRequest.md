# MfaCreateTotpMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The hashing algorithm used to generate the TOTP token. Options include SHA1, SHA256 and SHA512. | [optional] [default to "SHA1"]
**Digits** | Pointer to **int32** | The number of digits in the generated TOTP token. This value can either be 6 or 8. | [optional] [default to 6]
**Issuer** | Pointer to **string** | The name of the key&#x27;s issuing organization. | [optional] 
**KeySize** | Pointer to **int32** | Determines the size in bytes of the generated key. | [optional] [default to 20]
**MaxValidationAttempts** | Pointer to **int32** | Max number of allowed validation attempts. | [optional] 
**MethodName** | Pointer to **string** | The unique name identifier for this MFA method. | [optional] 
**Period** | Pointer to **string** | The length of time used to generate a counter for the TOTP token calculation. | [optional] [default to "30"]
**QrSize** | Pointer to **int32** | The pixel size of the generated square QR code. | [optional] [default to 200]
**Skew** | Pointer to **int32** | The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1. | [optional] [default to 1]



## Methods


### NewMfaCreateTotpMethodRequest

`func NewMfaCreateTotpMethodRequest() *MfaCreateTotpMethodRequest`

NewMfaCreateTotpMethodRequest instantiates a new MfaCreateTotpMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaCreateTotpMethodRequestWithDefaults

`func NewMfaCreateTotpMethodRequestWithDefaults() *MfaCreateTotpMethodRequest`

NewMfaCreateTotpMethodRequestWithDefaults instantiates a new MfaCreateTotpMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *MfaCreateTotpMethodRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *MfaCreateTotpMethodRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *MfaCreateTotpMethodRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *MfaCreateTotpMethodRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetDigits

`func (o *MfaCreateTotpMethodRequest) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *MfaCreateTotpMethodRequest) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *MfaCreateTotpMethodRequest) SetDigits(v int32)`

SetDigits sets Digits field to given value.


### HasDigits

`func (o *MfaCreateTotpMethodRequest) HasDigits() bool`

HasDigits returns a boolean if a field has been set.




### GetIssuer

`func (o *MfaCreateTotpMethodRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MfaCreateTotpMethodRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MfaCreateTotpMethodRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### HasIssuer

`func (o *MfaCreateTotpMethodRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.




### GetKeySize

`func (o *MfaCreateTotpMethodRequest) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *MfaCreateTotpMethodRequest) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *MfaCreateTotpMethodRequest) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.


### HasKeySize

`func (o *MfaCreateTotpMethodRequest) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.




### GetMaxValidationAttempts

`func (o *MfaCreateTotpMethodRequest) GetMaxValidationAttempts() int32`

GetMaxValidationAttempts returns the MaxValidationAttempts field if non-nil, zero value otherwise.

### GetMaxValidationAttemptsOk

`func (o *MfaCreateTotpMethodRequest) GetMaxValidationAttemptsOk() (*int32, bool)`

GetMaxValidationAttemptsOk returns a tuple with the MaxValidationAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValidationAttempts

`func (o *MfaCreateTotpMethodRequest) SetMaxValidationAttempts(v int32)`

SetMaxValidationAttempts sets MaxValidationAttempts field to given value.


### HasMaxValidationAttempts

`func (o *MfaCreateTotpMethodRequest) HasMaxValidationAttempts() bool`

HasMaxValidationAttempts returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaCreateTotpMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaCreateTotpMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaCreateTotpMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaCreateTotpMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetPeriod

`func (o *MfaCreateTotpMethodRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *MfaCreateTotpMethodRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *MfaCreateTotpMethodRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *MfaCreateTotpMethodRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetQrSize

`func (o *MfaCreateTotpMethodRequest) GetQrSize() int32`

GetQrSize returns the QrSize field if non-nil, zero value otherwise.

### GetQrSizeOk

`func (o *MfaCreateTotpMethodRequest) GetQrSizeOk() (*int32, bool)`

GetQrSizeOk returns a tuple with the QrSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrSize

`func (o *MfaCreateTotpMethodRequest) SetQrSize(v int32)`

SetQrSize sets QrSize field to given value.


### HasQrSize

`func (o *MfaCreateTotpMethodRequest) HasQrSize() bool`

HasQrSize returns a boolean if a field has been set.




### GetSkew

`func (o *MfaCreateTotpMethodRequest) GetSkew() int32`

GetSkew returns the Skew field if non-nil, zero value otherwise.

### GetSkewOk

`func (o *MfaCreateTotpMethodRequest) GetSkewOk() (*int32, bool)`

GetSkewOk returns a tuple with the Skew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkew

`func (o *MfaCreateTotpMethodRequest) SetSkew(v int32)`

SetSkew sets Skew field to given value.


### HasSkew

`func (o *MfaCreateTotpMethodRequest) HasSkew() bool`

HasSkew returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


