# CertConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableBinding** | Pointer to **bool** | If set, during renewal, skips the matching of presented client identity with the client identity used during login. Defaults to false. | [optional] [default to false]
**EnableIdentityAliasMetadata** | Pointer to **bool** | If set, metadata of the certificate including the metadata corresponding to allowed_metadata_extensions will be stored in the alias. Defaults to false. | [optional] [default to false]
**OcspCacheSize** | Pointer to **int32** | The size of the in memory OCSP response cache, shared by all configured certs | [optional] [default to 100]

## Methods

### NewCertConfigRequest

`func NewCertConfigRequest() *CertConfigRequest`

NewCertConfigRequest instantiates a new CertConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertConfigRequestWithDefaults

`func NewCertConfigRequestWithDefaults() *CertConfigRequest`

NewCertConfigRequestWithDefaults instantiates a new CertConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableBinding

`func (o *CertConfigRequest) GetDisableBinding() bool`

GetDisableBinding returns the DisableBinding field if non-nil, zero value otherwise.

### GetDisableBindingOk

`func (o *CertConfigRequest) GetDisableBindingOk() (*bool, bool)`

GetDisableBindingOk returns a tuple with the DisableBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBinding

`func (o *CertConfigRequest) SetDisableBinding(v bool)`

SetDisableBinding sets DisableBinding field to given value.

### HasDisableBinding

`func (o *CertConfigRequest) HasDisableBinding() bool`

HasDisableBinding returns a boolean if a field has been set.

### GetEnableIdentityAliasMetadata

`func (o *CertConfigRequest) GetEnableIdentityAliasMetadata() bool`

GetEnableIdentityAliasMetadata returns the EnableIdentityAliasMetadata field if non-nil, zero value otherwise.

### GetEnableIdentityAliasMetadataOk

`func (o *CertConfigRequest) GetEnableIdentityAliasMetadataOk() (*bool, bool)`

GetEnableIdentityAliasMetadataOk returns a tuple with the EnableIdentityAliasMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIdentityAliasMetadata

`func (o *CertConfigRequest) SetEnableIdentityAliasMetadata(v bool)`

SetEnableIdentityAliasMetadata sets EnableIdentityAliasMetadata field to given value.

### HasEnableIdentityAliasMetadata

`func (o *CertConfigRequest) HasEnableIdentityAliasMetadata() bool`

HasEnableIdentityAliasMetadata returns a boolean if a field has been set.

### GetOcspCacheSize

`func (o *CertConfigRequest) GetOcspCacheSize() int32`

GetOcspCacheSize returns the OcspCacheSize field if non-nil, zero value otherwise.

### GetOcspCacheSizeOk

`func (o *CertConfigRequest) GetOcspCacheSizeOk() (*int32, bool)`

GetOcspCacheSizeOk returns a tuple with the OcspCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspCacheSize

`func (o *CertConfigRequest) SetOcspCacheSize(v int32)`

SetOcspCacheSize sets OcspCacheSize field to given value.

### HasOcspCacheSize

`func (o *CertConfigRequest) HasOcspCacheSize() bool`

HasOcspCacheSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


