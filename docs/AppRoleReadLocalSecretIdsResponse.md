# AppRoleReadLocalSecretIdsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSecretIds** | Pointer to **bool** | If true, the secret identifiers generated using this role will be cluster local. This can only be set during role creation and once set, it can&#x27;t be reset later | [optional] 



## Methods


### NewAppRoleReadLocalSecretIdsResponse

`func NewAppRoleReadLocalSecretIdsResponse() *AppRoleReadLocalSecretIdsResponse`

NewAppRoleReadLocalSecretIdsResponse instantiates a new AppRoleReadLocalSecretIdsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleReadLocalSecretIdsResponseWithDefaults

`func NewAppRoleReadLocalSecretIdsResponseWithDefaults() *AppRoleReadLocalSecretIdsResponse`

NewAppRoleReadLocalSecretIdsResponseWithDefaults instantiates a new AppRoleReadLocalSecretIdsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetLocalSecretIds

`func (o *AppRoleReadLocalSecretIdsResponse) GetLocalSecretIds() bool`

GetLocalSecretIds returns the LocalSecretIds field if non-nil, zero value otherwise.

### GetLocalSecretIdsOk

`func (o *AppRoleReadLocalSecretIdsResponse) GetLocalSecretIdsOk() (*bool, bool)`

GetLocalSecretIdsOk returns a tuple with the LocalSecretIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecretIds

`func (o *AppRoleReadLocalSecretIdsResponse) SetLocalSecretIds(v bool)`

SetLocalSecretIds sets LocalSecretIds field to given value.


### HasLocalSecretIds

`func (o *AppRoleReadLocalSecretIdsResponse) HasLocalSecretIds() bool`

HasLocalSecretIds returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


