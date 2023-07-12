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





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


