# WriteAuditDeviceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Description** | Pointer to **string** | User-friendly description for this audit backend. | [optional] 
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | Configuration options for the audit backend. | [optional] 
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;mysql\&quot; | [optional] 



## Methods


### NewWriteAuditDeviceRequest

`func NewWriteAuditDeviceRequest() *WriteAuditDeviceRequest`

NewWriteAuditDeviceRequest instantiates a new WriteAuditDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteAuditDeviceRequestWithDefaults

`func NewWriteAuditDeviceRequestWithDefaults() *WriteAuditDeviceRequest`

NewWriteAuditDeviceRequestWithDefaults instantiates a new WriteAuditDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDescription

`func (o *WriteAuditDeviceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WriteAuditDeviceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WriteAuditDeviceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *WriteAuditDeviceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetLocal

`func (o *WriteAuditDeviceRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *WriteAuditDeviceRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *WriteAuditDeviceRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *WriteAuditDeviceRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetOptions

`func (o *WriteAuditDeviceRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WriteAuditDeviceRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WriteAuditDeviceRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *WriteAuditDeviceRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetType

`func (o *WriteAuditDeviceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteAuditDeviceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteAuditDeviceRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *WriteAuditDeviceRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


