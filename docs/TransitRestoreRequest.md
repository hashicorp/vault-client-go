# TransitRestoreRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **string** | Backed up key data to be restored. This should be the output from the &#x27;backup/&#x27; endpoint. | [optional] 
**Force** | Pointer to **bool** | If set and a key by the given name exists, force the restore operation and override the key. | [optional] [default to false]
**Name** | Pointer to **string** | If set, this will be the name of the restored key. | [optional] 



## Methods


### NewTransitRestoreRequest

`func NewTransitRestoreRequest() *TransitRestoreRequest`

NewTransitRestoreRequest instantiates a new TransitRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitRestoreRequestWithDefaults

`func NewTransitRestoreRequestWithDefaults() *TransitRestoreRequest`

NewTransitRestoreRequestWithDefaults instantiates a new TransitRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBackup

`func (o *TransitRestoreRequest) GetBackup() string`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *TransitRestoreRequest) GetBackupOk() (*string, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *TransitRestoreRequest) SetBackup(v string)`

SetBackup sets Backup field to given value.


### HasBackup

`func (o *TransitRestoreRequest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.




### GetForce

`func (o *TransitRestoreRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *TransitRestoreRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *TransitRestoreRequest) SetForce(v bool)`

SetForce sets Force field to given value.


### HasForce

`func (o *TransitRestoreRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.




### GetName

`func (o *TransitRestoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitRestoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitRestoreRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *TransitRestoreRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


