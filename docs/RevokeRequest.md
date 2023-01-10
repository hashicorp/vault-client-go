# RevokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 
**Sync** | Pointer to **bool** | Whether or not to perform the revocation synchronously | [optional] [default to true]
**UrlLeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 

## Methods

### NewRevokeRequest

`func NewRevokeRequest() *RevokeRequest`

NewRevokeRequest instantiates a new RevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeRequestWithDefaults

`func NewRevokeRequestWithDefaults() *RevokeRequest`

NewRevokeRequestWithDefaults instantiates a new RevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseId

`func (o *RevokeRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *RevokeRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *RevokeRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *RevokeRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetSync

`func (o *RevokeRequest) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *RevokeRequest) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *RevokeRequest) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *RevokeRequest) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetUrlLeaseId

`func (o *RevokeRequest) GetUrlLeaseId() string`

GetUrlLeaseId returns the UrlLeaseId field if non-nil, zero value otherwise.

### GetUrlLeaseIdOk

`func (o *RevokeRequest) GetUrlLeaseIdOk() (*string, bool)`

GetUrlLeaseIdOk returns a tuple with the UrlLeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLeaseId

`func (o *RevokeRequest) SetUrlLeaseId(v string)`

SetUrlLeaseId sets UrlLeaseId field to given value.

### HasUrlLeaseId

`func (o *RevokeRequest) HasUrlLeaseId() bool`

HasUrlLeaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


