# TransitTrimKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAvailableVersion** | Pointer to **int32** | The minimum available version for the key ring. All versions before this version will be permanently deleted. This value can at most be equal to the lesser of &#39;min_decryption_version&#39; and &#39;min_encryption_version&#39;. This is not allowed to be set when either &#39;min_encryption_version&#39; or &#39;min_decryption_version&#39; is set to zero. | [optional] 

## Methods

### NewTransitTrimKeyRequest

`func NewTransitTrimKeyRequest() *TransitTrimKeyRequest`

NewTransitTrimKeyRequest instantiates a new TransitTrimKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitTrimKeyRequestWithDefaults

`func NewTransitTrimKeyRequestWithDefaults() *TransitTrimKeyRequest`

NewTransitTrimKeyRequestWithDefaults instantiates a new TransitTrimKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAvailableVersion

`func (o *TransitTrimKeyRequest) GetMinAvailableVersion() int32`

GetMinAvailableVersion returns the MinAvailableVersion field if non-nil, zero value otherwise.

### GetMinAvailableVersionOk

`func (o *TransitTrimKeyRequest) GetMinAvailableVersionOk() (*int32, bool)`

GetMinAvailableVersionOk returns a tuple with the MinAvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailableVersion

`func (o *TransitTrimKeyRequest) SetMinAvailableVersion(v int32)`

SetMinAvailableVersion sets MinAvailableVersion field to given value.

### HasMinAvailableVersion

`func (o *TransitTrimKeyRequest) HasMinAvailableVersion() bool`

HasMinAvailableVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


