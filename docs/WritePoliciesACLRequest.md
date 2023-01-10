# WritePoliciesACLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | The rules of the policy. | [optional] 

## Methods

### NewWritePoliciesACLRequest

`func NewWritePoliciesACLRequest() *WritePoliciesACLRequest`

NewWritePoliciesACLRequest instantiates a new WritePoliciesACLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritePoliciesACLRequestWithDefaults

`func NewWritePoliciesACLRequestWithDefaults() *WritePoliciesACLRequest`

NewWritePoliciesACLRequestWithDefaults instantiates a new WritePoliciesACLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *WritePoliciesACLRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *WritePoliciesACLRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *WritePoliciesACLRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *WritePoliciesACLRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


