# WritePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | The rules of the policy. | [optional] 
**Rules** | Pointer to **string** | The rules of the policy. | [optional] 

## Methods

### NewWritePolicyRequest

`func NewWritePolicyRequest() *WritePolicyRequest`

NewWritePolicyRequest instantiates a new WritePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritePolicyRequestWithDefaults

`func NewWritePolicyRequestWithDefaults() *WritePolicyRequest`

NewWritePolicyRequestWithDefaults instantiates a new WritePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *WritePolicyRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *WritePolicyRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *WritePolicyRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *WritePolicyRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRules

`func (o *WritePolicyRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *WritePolicyRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *WritePolicyRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *WritePolicyRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


