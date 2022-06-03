# SystemAuditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User-friendly description for this audit backend. | [optional] 
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | Configuration options for the audit backend. | [optional] 
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;mysql\&quot; | [optional] 

## Methods

### NewSystemAuditRequest

`func NewSystemAuditRequest() *SystemAuditRequest`

NewSystemAuditRequest instantiates a new SystemAuditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAuditRequestWithDefaults

`func NewSystemAuditRequestWithDefaults() *SystemAuditRequest`

NewSystemAuditRequestWithDefaults instantiates a new SystemAuditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SystemAuditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemAuditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemAuditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemAuditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocal

`func (o *SystemAuditRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *SystemAuditRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *SystemAuditRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *SystemAuditRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetOptions

`func (o *SystemAuditRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemAuditRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemAuditRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemAuditRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *SystemAuditRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemAuditRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemAuditRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemAuditRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


