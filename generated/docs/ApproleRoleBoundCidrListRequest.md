# ApproleRoleBoundCidrListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundCidrList** | Pointer to **[]string** | Deprecated: Please use \&quot;secret_id_bound_cidrs\&quot; instead. Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation. | [optional] 

## Methods

### NewApproleRoleBoundCidrListRequest

`func NewApproleRoleBoundCidrListRequest() *ApproleRoleBoundCidrListRequest`

NewApproleRoleBoundCidrListRequest instantiates a new ApproleRoleBoundCidrListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleRoleBoundCidrListRequestWithDefaults

`func NewApproleRoleBoundCidrListRequestWithDefaults() *ApproleRoleBoundCidrListRequest`

NewApproleRoleBoundCidrListRequestWithDefaults instantiates a new ApproleRoleBoundCidrListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundCidrList

`func (o *ApproleRoleBoundCidrListRequest) GetBoundCidrList() []string`

GetBoundCidrList returns the BoundCidrList field if non-nil, zero value otherwise.

### GetBoundCidrListOk

`func (o *ApproleRoleBoundCidrListRequest) GetBoundCidrListOk() (*[]string, bool)`

GetBoundCidrListOk returns a tuple with the BoundCidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrList

`func (o *ApproleRoleBoundCidrListRequest) SetBoundCidrList(v []string)`

SetBoundCidrList sets BoundCidrList field to given value.

### HasBoundCidrList

`func (o *ApproleRoleBoundCidrListRequest) HasBoundCidrList() bool`

HasBoundCidrList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


