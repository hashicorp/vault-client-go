# AppIdMapUserIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | If not blank, restricts auth by this CIDR block | [optional] 
**Value** | Pointer to **string** | App IDs that this user associates with. | [optional] 

## Methods

### NewAppIdMapUserIdRequest

`func NewAppIdMapUserIdRequest() *AppIdMapUserIdRequest`

NewAppIdMapUserIdRequest instantiates a new AppIdMapUserIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIdMapUserIdRequestWithDefaults

`func NewAppIdMapUserIdRequestWithDefaults() *AppIdMapUserIdRequest`

NewAppIdMapUserIdRequestWithDefaults instantiates a new AppIdMapUserIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AppIdMapUserIdRequest) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AppIdMapUserIdRequest) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AppIdMapUserIdRequest) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AppIdMapUserIdRequest) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetValue

`func (o *AppIdMapUserIdRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppIdMapUserIdRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppIdMapUserIdRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppIdMapUserIdRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


