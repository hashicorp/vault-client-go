# UnsealResponse


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


### NewUnsealResponse

`func NewUnsealResponse() *UnsealResponse`

NewUnsealResponse instantiates a new UnsealResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsealResponseWithDefaults

`func NewUnsealResponseWithDefaults() *UnsealResponse`

NewUnsealResponseWithDefaults instantiates a new UnsealResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBuildDate

`func (o *UnsealResponse) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *UnsealResponse) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *UnsealResponse) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.


### HasBuildDate

`func (o *UnsealResponse) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.




### GetClusterId

`func (o *UnsealResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *UnsealResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *UnsealResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### HasClusterId

`func (o *UnsealResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.




### GetClusterName

`func (o *UnsealResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *UnsealResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *UnsealResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### HasClusterName

`func (o *UnsealResponse) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.




### GetHcpLinkResourceID

`func (o *UnsealResponse) GetHcpLinkResourceID() string`

GetHcpLinkResourceID returns the HcpLinkResourceID field if non-nil, zero value otherwise.

### GetHcpLinkResourceIDOk

`func (o *UnsealResponse) GetHcpLinkResourceIDOk() (*string, bool)`

GetHcpLinkResourceIDOk returns a tuple with the HcpLinkResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpLinkResourceID

`func (o *UnsealResponse) SetHcpLinkResourceID(v string)`

SetHcpLinkResourceID sets HcpLinkResourceID field to given value.


### HasHcpLinkResourceID

`func (o *UnsealResponse) HasHcpLinkResourceID() bool`

HasHcpLinkResourceID returns a boolean if a field has been set.




### GetHcpLinkStatus

`func (o *UnsealResponse) GetHcpLinkStatus() string`

GetHcpLinkStatus returns the HcpLinkStatus field if non-nil, zero value otherwise.

### GetHcpLinkStatusOk

`func (o *UnsealResponse) GetHcpLinkStatusOk() (*string, bool)`

GetHcpLinkStatusOk returns a tuple with the HcpLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcpLinkStatus

`func (o *UnsealResponse) SetHcpLinkStatus(v string)`

SetHcpLinkStatus sets HcpLinkStatus field to given value.


### HasHcpLinkStatus

`func (o *UnsealResponse) HasHcpLinkStatus() bool`

HasHcpLinkStatus returns a boolean if a field has been set.




### GetInitialized

`func (o *UnsealResponse) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *UnsealResponse) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *UnsealResponse) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.


### HasInitialized

`func (o *UnsealResponse) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.




### GetMigration

`func (o *UnsealResponse) GetMigration() bool`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *UnsealResponse) GetMigrationOk() (*bool, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *UnsealResponse) SetMigration(v bool)`

SetMigration sets Migration field to given value.


### HasMigration

`func (o *UnsealResponse) HasMigration() bool`

HasMigration returns a boolean if a field has been set.




### GetN

`func (o *UnsealResponse) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *UnsealResponse) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *UnsealResponse) SetN(v int32)`

SetN sets N field to given value.


### HasN

`func (o *UnsealResponse) HasN() bool`

HasN returns a boolean if a field has been set.




### GetNonce

`func (o *UnsealResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *UnsealResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *UnsealResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *UnsealResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetProgress

`func (o *UnsealResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UnsealResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UnsealResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *UnsealResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRecoverySeal

`func (o *UnsealResponse) GetRecoverySeal() bool`

GetRecoverySeal returns the RecoverySeal field if non-nil, zero value otherwise.

### GetRecoverySealOk

`func (o *UnsealResponse) GetRecoverySealOk() (*bool, bool)`

GetRecoverySealOk returns a tuple with the RecoverySeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySeal

`func (o *UnsealResponse) SetRecoverySeal(v bool)`

SetRecoverySeal sets RecoverySeal field to given value.


### HasRecoverySeal

`func (o *UnsealResponse) HasRecoverySeal() bool`

HasRecoverySeal returns a boolean if a field has been set.




### GetSealed

`func (o *UnsealResponse) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *UnsealResponse) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *UnsealResponse) SetSealed(v bool)`

SetSealed sets Sealed field to given value.


### HasSealed

`func (o *UnsealResponse) HasSealed() bool`

HasSealed returns a boolean if a field has been set.




### GetStorageType

`func (o *UnsealResponse) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UnsealResponse) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UnsealResponse) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### HasStorageType

`func (o *UnsealResponse) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.




### GetT

`func (o *UnsealResponse) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UnsealResponse) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UnsealResponse) SetT(v int32)`

SetT sets T field to given value.


### HasT

`func (o *UnsealResponse) HasT() bool`

HasT returns a boolean if a field has been set.




### GetType

`func (o *UnsealResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnsealResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnsealResponse) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *UnsealResponse) HasType() bool`

HasType returns a boolean if a field has been set.




### GetVersion

`func (o *UnsealResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnsealResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnsealResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *UnsealResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


