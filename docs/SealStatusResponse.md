# SealStatusResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildDate** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**HcpLinkResourceID** | Pointer to **string** |  | [optional] 
**HcpLinkStatus** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Migration** | Pointer to **bool** |  | [optional] 
**N** | Pointer to **int32** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**RecoverySeal** | Pointer to **bool** |  | [optional] 
**Sealed** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 



## Methods


### NewSealStatusResponse

`func NewSealStatusResponse() *SealStatusResponse`

NewSealStatusResponse instantiates a new SealStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSealStatusResponseWithDefaults

`func NewSealStatusResponseWithDefaults() *SealStatusResponse`

NewSealStatusResponseWithDefaults instantiates a new SealStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBuildDate

`func (o *SealStatusResponse) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *SealStatusResponse) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *SealStatusResponse) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.


### HasBuildDate

`func (o *SealStatusResponse) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.




### GetClusterId

`func (o *SealStatusResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SealStatusResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SealStatusResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### HasClusterId

`func (o *SealStatusResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.




### GetClusterName

`func (o *SealStatusResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SealStatusResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SealStatusResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### HasClusterName

`func (o *SealStatusResponse) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.




### GetHcpLinkResourceID

`func (o *SealStatusResponse) GetHcpLinkResourceID() string`

GetHcpLinkResourceID returns the HcpLinkResourceID field if non-nil, zero value otherwise.

### GetHcpLinkResourceIDOk

`func (o *SealStatusResponse) GetHcpLinkResourceIDOk() (*string, bool)`

GetHcpLinkResourceIDOk returns a tuple with the HcpLinkResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpLinkResourceID

`func (o *SealStatusResponse) SetHcpLinkResourceID(v string)`

SetHcpLinkResourceID sets HcpLinkResourceID field to given value.


### HasHcpLinkResourceID

`func (o *SealStatusResponse) HasHcpLinkResourceID() bool`

HasHcpLinkResourceID returns a boolean if a field has been set.




### GetHcpLinkStatus

`func (o *SealStatusResponse) GetHcpLinkStatus() string`

GetHcpLinkStatus returns the HcpLinkStatus field if non-nil, zero value otherwise.

### GetHcpLinkStatusOk

`func (o *SealStatusResponse) GetHcpLinkStatusOk() (*string, bool)`

GetHcpLinkStatusOk returns a tuple with the HcpLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpLinkStatus

`func (o *SealStatusResponse) SetHcpLinkStatus(v string)`

SetHcpLinkStatus sets HcpLinkStatus field to given value.


### HasHcpLinkStatus

`func (o *SealStatusResponse) HasHcpLinkStatus() bool`

HasHcpLinkStatus returns a boolean if a field has been set.




### GetInitialized

`func (o *SealStatusResponse) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *SealStatusResponse) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *SealStatusResponse) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.


### HasInitialized

`func (o *SealStatusResponse) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.




### GetMigration

`func (o *SealStatusResponse) GetMigration() bool`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *SealStatusResponse) GetMigrationOk() (*bool, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *SealStatusResponse) SetMigration(v bool)`

SetMigration sets Migration field to given value.


### HasMigration

`func (o *SealStatusResponse) HasMigration() bool`

HasMigration returns a boolean if a field has been set.




### GetN

`func (o *SealStatusResponse) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *SealStatusResponse) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *SealStatusResponse) SetN(v int32)`

SetN sets N field to given value.


### HasN

`func (o *SealStatusResponse) HasN() bool`

HasN returns a boolean if a field has been set.




### GetNonce

`func (o *SealStatusResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SealStatusResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SealStatusResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *SealStatusResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetProgress

`func (o *SealStatusResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SealStatusResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SealStatusResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *SealStatusResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRecoverySeal

`func (o *SealStatusResponse) GetRecoverySeal() bool`

GetRecoverySeal returns the RecoverySeal field if non-nil, zero value otherwise.

### GetRecoverySealOk

`func (o *SealStatusResponse) GetRecoverySealOk() (*bool, bool)`

GetRecoverySealOk returns a tuple with the RecoverySeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySeal

`func (o *SealStatusResponse) SetRecoverySeal(v bool)`

SetRecoverySeal sets RecoverySeal field to given value.


### HasRecoverySeal

`func (o *SealStatusResponse) HasRecoverySeal() bool`

HasRecoverySeal returns a boolean if a field has been set.




### GetSealed

`func (o *SealStatusResponse) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *SealStatusResponse) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *SealStatusResponse) SetSealed(v bool)`

SetSealed sets Sealed field to given value.


### HasSealed

`func (o *SealStatusResponse) HasSealed() bool`

HasSealed returns a boolean if a field has been set.




### GetStorageType

`func (o *SealStatusResponse) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SealStatusResponse) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SealStatusResponse) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### HasStorageType

`func (o *SealStatusResponse) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.




### GetT

`func (o *SealStatusResponse) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SealStatusResponse) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SealStatusResponse) SetT(v int32)`

SetT sets T field to given value.


### HasT

`func (o *SealStatusResponse) HasT() bool`

HasT returns a boolean if a field has been set.




### GetType

`func (o *SealStatusResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SealStatusResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SealStatusResponse) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *SealStatusResponse) HasType() bool`

HasType returns a boolean if a field has been set.




### GetVersion

`func (o *SealStatusResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SealStatusResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SealStatusResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *SealStatusResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


