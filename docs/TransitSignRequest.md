# TransitSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Deprecated: use \&quot;hash_algorithm\&quot; instead. | [optional] [default to "sha2-256"]
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys. | [optional] 
**HashAlgorithm** | Pointer to **string** | Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \&quot;sha2-256\&quot;. Not valid for all key types, including ed25519. | [optional] [default to "sha2-256"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for signing. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**MarshalingAlgorithm** | Pointer to **string** | The method by which to marshal the signature. The default is &#39;asn1&#39; which is used by openssl and X.509. It can also be set to &#39;jws&#39; which is used for JWT signatures; setting it to this will also cause the encoding of the signature to be url-safe base64 instead of using standard base64 encoding. Currently only valid for ECDSA P-256 key types\&quot;. | [optional] [default to "asn1"]
**Prehashed** | Pointer to **bool** | Set to &#39;true&#39; when the input is already hashed. If the key type is &#39;rsa-2048&#39;, &#39;rsa-3072&#39; or &#39;rsa-4096&#39;, then the algorithm used to hash the input should be indicated by the &#39;algorithm&#39; parameter. | [optional] 
**SaltLength** | Pointer to **string** | The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are &#39;auto&#39; (the default used by Golang, causing the salt to be as large as possible when signing), &#39;hash&#39; (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to &#39;auto&#39;. | [optional] [default to "auto"]
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm to use for signing. Currently only applies to RSA key types. Options are &#39;pss&#39; or &#39;pkcs1v15&#39;. Defaults to &#39;pss&#39; | [optional] 
**Urlalgorithm** | Pointer to **string** | Hash algorithm to use (POST URL parameter) | [optional] 

## Methods

### NewTransitSignRequest

`func NewTransitSignRequest() *TransitSignRequest`

NewTransitSignRequest instantiates a new TransitSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitSignRequestWithDefaults

`func NewTransitSignRequestWithDefaults() *TransitSignRequest`

NewTransitSignRequestWithDefaults instantiates a new TransitSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *TransitSignRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransitSignRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransitSignRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TransitSignRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetContext

`func (o *TransitSignRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitSignRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitSignRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitSignRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *TransitSignRequest) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *TransitSignRequest) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *TransitSignRequest) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *TransitSignRequest) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetInput

`func (o *TransitSignRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransitSignRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransitSignRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TransitSignRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetKeyVersion

`func (o *TransitSignRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitSignRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitSignRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.

### HasKeyVersion

`func (o *TransitSignRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.

### GetMarshalingAlgorithm

`func (o *TransitSignRequest) GetMarshalingAlgorithm() string`

GetMarshalingAlgorithm returns the MarshalingAlgorithm field if non-nil, zero value otherwise.

### GetMarshalingAlgorithmOk

`func (o *TransitSignRequest) GetMarshalingAlgorithmOk() (*string, bool)`

GetMarshalingAlgorithmOk returns a tuple with the MarshalingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarshalingAlgorithm

`func (o *TransitSignRequest) SetMarshalingAlgorithm(v string)`

SetMarshalingAlgorithm sets MarshalingAlgorithm field to given value.

### HasMarshalingAlgorithm

`func (o *TransitSignRequest) HasMarshalingAlgorithm() bool`

HasMarshalingAlgorithm returns a boolean if a field has been set.

### GetPrehashed

`func (o *TransitSignRequest) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *TransitSignRequest) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *TransitSignRequest) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.

### HasPrehashed

`func (o *TransitSignRequest) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.

### GetSaltLength

`func (o *TransitSignRequest) GetSaltLength() string`

GetSaltLength returns the SaltLength field if non-nil, zero value otherwise.

### GetSaltLengthOk

`func (o *TransitSignRequest) GetSaltLengthOk() (*string, bool)`

GetSaltLengthOk returns a tuple with the SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLength

`func (o *TransitSignRequest) SetSaltLength(v string)`

SetSaltLength sets SaltLength field to given value.

### HasSaltLength

`func (o *TransitSignRequest) HasSaltLength() bool`

HasSaltLength returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *TransitSignRequest) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TransitSignRequest) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TransitSignRequest) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *TransitSignRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetUrlalgorithm

`func (o *TransitSignRequest) GetUrlalgorithm() string`

GetUrlalgorithm returns the Urlalgorithm field if non-nil, zero value otherwise.

### GetUrlalgorithmOk

`func (o *TransitSignRequest) GetUrlalgorithmOk() (*string, bool)`

GetUrlalgorithmOk returns a tuple with the Urlalgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlalgorithm

`func (o *TransitSignRequest) SetUrlalgorithm(v string)`

SetUrlalgorithm sets Urlalgorithm field to given value.

### HasUrlalgorithm

`func (o *TransitSignRequest) HasUrlalgorithm() bool`

HasUrlalgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


