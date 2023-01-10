# TransitImportKeyVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | The base64-encoded ciphertext of the keys. The AES key should be encrypted using OAEP with the wrapping key and then concatenated with the import key, wrapped by the AES key. | [optional] 
**HashFunction** | Pointer to **string** | The hash function used as a random oracle in the OAEP wrapping of the user-generated, ephemeral AES key. Can be one of \&quot;SHA1\&quot;, \&quot;SHA224\&quot;, \&quot;SHA256\&quot; (default), \&quot;SHA384\&quot;, or \&quot;SHA512\&quot; | [optional] [default to "SHA256"]

## Methods

### NewTransitImportKeyVersionRequest

`func NewTransitImportKeyVersionRequest() *TransitImportKeyVersionRequest`

NewTransitImportKeyVersionRequest instantiates a new TransitImportKeyVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitImportKeyVersionRequestWithDefaults

`func NewTransitImportKeyVersionRequestWithDefaults() *TransitImportKeyVersionRequest`

NewTransitImportKeyVersionRequestWithDefaults instantiates a new TransitImportKeyVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *TransitImportKeyVersionRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TransitImportKeyVersionRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TransitImportKeyVersionRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *TransitImportKeyVersionRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetHashFunction

`func (o *TransitImportKeyVersionRequest) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *TransitImportKeyVersionRequest) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *TransitImportKeyVersionRequest) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *TransitImportKeyVersionRequest) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


