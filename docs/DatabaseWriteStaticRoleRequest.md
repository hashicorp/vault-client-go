# DatabaseWriteStaticRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialConfig** | Pointer to **map[string]interface{}** | The configuration for the given credential_type. | [optional] 
**CredentialType** | Pointer to **string** | The type of credential to manage. Options include: &#x27;password&#x27;, &#x27;rsa_private_key&#x27;. Defaults to &#x27;password&#x27;. | [optional] [default to "password"]
**DbName** | Pointer to **string** | Name of the database this role acts on. | [optional] 
**RotationPeriod** | Pointer to **string** | Period for automatic credential rotation of the given username. Not valid unless used with \&quot;username\&quot;. | [optional] 
**RotationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to rotate the accounts credentials. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**Username** | Pointer to **string** | Name of the static user account for Vault to manage. Requires \&quot;rotation_period\&quot; to be specified | [optional] 



## Methods


### NewDatabaseWriteStaticRoleRequest

`func NewDatabaseWriteStaticRoleRequest() *DatabaseWriteStaticRoleRequest`

NewDatabaseWriteStaticRoleRequest instantiates a new DatabaseWriteStaticRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWriteStaticRoleRequestWithDefaults

`func NewDatabaseWriteStaticRoleRequestWithDefaults() *DatabaseWriteStaticRoleRequest`

NewDatabaseWriteStaticRoleRequestWithDefaults instantiates a new DatabaseWriteStaticRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCredentialConfig

`func (o *DatabaseWriteStaticRoleRequest) GetCredentialConfig() map[string]interface{}`

GetCredentialConfig returns the CredentialConfig field if non-nil, zero value otherwise.

### GetCredentialConfigOk

`func (o *DatabaseWriteStaticRoleRequest) GetCredentialConfigOk() (*map[string]interface{}, bool)`

GetCredentialConfigOk returns a tuple with the CredentialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialConfig

`func (o *DatabaseWriteStaticRoleRequest) SetCredentialConfig(v map[string]interface{})`

SetCredentialConfig sets CredentialConfig field to given value.


### HasCredentialConfig

`func (o *DatabaseWriteStaticRoleRequest) HasCredentialConfig() bool`

HasCredentialConfig returns a boolean if a field has been set.




### GetCredentialType

`func (o *DatabaseWriteStaticRoleRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *DatabaseWriteStaticRoleRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *DatabaseWriteStaticRoleRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.


### HasCredentialType

`func (o *DatabaseWriteStaticRoleRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.




### GetDbName

`func (o *DatabaseWriteStaticRoleRequest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DatabaseWriteStaticRoleRequest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DatabaseWriteStaticRoleRequest) SetDbName(v string)`

SetDbName sets DbName field to given value.


### HasDbName

`func (o *DatabaseWriteStaticRoleRequest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.




### GetRotationPeriod

`func (o *DatabaseWriteStaticRoleRequest) GetRotationPeriod() string`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *DatabaseWriteStaticRoleRequest) GetRotationPeriodOk() (*string, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *DatabaseWriteStaticRoleRequest) SetRotationPeriod(v string)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *DatabaseWriteStaticRoleRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetRotationStatements

`func (o *DatabaseWriteStaticRoleRequest) GetRotationStatements() []string`

GetRotationStatements returns the RotationStatements field if non-nil, zero value otherwise.

### GetRotationStatementsOk

`func (o *DatabaseWriteStaticRoleRequest) GetRotationStatementsOk() (*[]string, bool)`

GetRotationStatementsOk returns a tuple with the RotationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationStatements

`func (o *DatabaseWriteStaticRoleRequest) SetRotationStatements(v []string)`

SetRotationStatements sets RotationStatements field to given value.


### HasRotationStatements

`func (o *DatabaseWriteStaticRoleRequest) HasRotationStatements() bool`

HasRotationStatements returns a boolean if a field has been set.




### GetUsername

`func (o *DatabaseWriteStaticRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseWriteStaticRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseWriteStaticRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *DatabaseWriteStaticRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


