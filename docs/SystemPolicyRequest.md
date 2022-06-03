# SystemPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | The rules of the policy. | [optional] 
**Rules** | Pointer to **string** | The rules of the policy. | [optional] 

## Methods

### NewSystemPolicyRequest

`func NewSystemPolicyRequest() *SystemPolicyRequest`

NewSystemPolicyRequest instantiates a new SystemPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPolicyRequestWithDefaults

`func NewSystemPolicyRequestWithDefaults() *SystemPolicyRequest`

NewSystemPolicyRequestWithDefaults instantiates a new SystemPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *SystemPolicyRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SystemPolicyRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SystemPolicyRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SystemPolicyRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRules

`func (o *SystemPolicyRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SystemPolicyRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SystemPolicyRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SystemPolicyRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


