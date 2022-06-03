# SystemRemountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | The previous mount point. | [optional] 
**To** | Pointer to **string** | The new mount point. | [optional] 

## Methods

### NewSystemRemountRequest

`func NewSystemRemountRequest() *SystemRemountRequest`

NewSystemRemountRequest instantiates a new SystemRemountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRemountRequestWithDefaults

`func NewSystemRemountRequestWithDefaults() *SystemRemountRequest`

NewSystemRemountRequestWithDefaults instantiates a new SystemRemountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SystemRemountRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SystemRemountRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SystemRemountRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SystemRemountRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SystemRemountRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SystemRemountRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SystemRemountRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SystemRemountRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


