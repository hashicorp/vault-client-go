# LeasesRevokeLeaseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 
**Sync** | Pointer to **bool** | Whether or not to perform the revocation synchronously | [optional] [default to true]
**UrlLeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 



## Methods


### NewLeasesRevokeLeaseRequest

`func NewLeasesRevokeLeaseRequest() *LeasesRevokeLeaseRequest`

NewLeasesRevokeLeaseRequest instantiates a new LeasesRevokeLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesRevokeLeaseRequestWithDefaults

`func NewLeasesRevokeLeaseRequestWithDefaults() *LeasesRevokeLeaseRequest`

NewLeasesRevokeLeaseRequestWithDefaults instantiates a new LeasesRevokeLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetLeaseId

`func (o *LeasesRevokeLeaseRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *LeasesRevokeLeaseRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *LeasesRevokeLeaseRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *LeasesRevokeLeaseRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.




### GetSync

`func (o *LeasesRevokeLeaseRequest) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *LeasesRevokeLeaseRequest) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *LeasesRevokeLeaseRequest) SetSync(v bool)`

SetSync sets Sync field to given value.


### HasSync

`func (o *LeasesRevokeLeaseRequest) HasSync() bool`

HasSync returns a boolean if a field has been set.




### GetUrlLeaseId

`func (o *LeasesRevokeLeaseRequest) GetUrlLeaseId() string`

GetUrlLeaseId returns the UrlLeaseId field if non-nil, zero value otherwise.

### GetUrlLeaseIdOk

`func (o *LeasesRevokeLeaseRequest) GetUrlLeaseIdOk() (*string, bool)`

GetUrlLeaseIdOk returns a tuple with the UrlLeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLeaseId

`func (o *LeasesRevokeLeaseRequest) SetUrlLeaseId(v string)`

SetUrlLeaseId sets UrlLeaseId field to given value.


### HasUrlLeaseId

`func (o *LeasesRevokeLeaseRequest) HasUrlLeaseId() bool`

HasUrlLeaseId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


