# WriteLeasesRevoke2Request


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 
**Sync** | Pointer to **bool** | Whether or not to perform the revocation synchronously | [optional] [default to true]



## Methods


### NewWriteLeasesRevoke2Request

`func NewWriteLeasesRevoke2Request() *WriteLeasesRevoke2Request`

NewWriteLeasesRevoke2Request instantiates a new WriteLeasesRevoke2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteLeasesRevoke2RequestWithDefaults

`func NewWriteLeasesRevoke2RequestWithDefaults() *WriteLeasesRevoke2Request`

NewWriteLeasesRevoke2RequestWithDefaults instantiates a new WriteLeasesRevoke2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetLeaseId

`func (o *WriteLeasesRevoke2Request) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *WriteLeasesRevoke2Request) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *WriteLeasesRevoke2Request) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *WriteLeasesRevoke2Request) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.




### GetSync

`func (o *WriteLeasesRevoke2Request) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *WriteLeasesRevoke2Request) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *WriteLeasesRevoke2Request) SetSync(v bool)`

SetSync sets Sync field to given value.


### HasSync

`func (o *WriteLeasesRevoke2Request) HasSync() bool`

HasSync returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


