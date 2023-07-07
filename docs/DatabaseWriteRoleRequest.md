# DatabaseWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **[]string** | Specifies the database statements executed to create and configure a user. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**CredentialConfig** | Pointer to **map[string]interface{}** | The configuration for the given credential_type. | [optional] 
**CredentialType** | Pointer to **string** | The type of credential to manage. Options include: &#x27;password&#x27;, &#x27;rsa_private_key&#x27;. Defaults to &#x27;password&#x27;. | [optional] [default to "password"]
**DbName** | Pointer to **string** | Name of the database this role acts on. | [optional] 
**DefaultTtl** | Pointer to **string** | Default ttl for role. | [optional] 
**MaxTtl** | Pointer to **string** | Maximum time a credential is valid for | [optional] 
**RenewStatements** | Pointer to **[]string** | Specifies the database statements to be executed to renew a user. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**RevocationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to revoke a user. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**RollbackStatements** | Pointer to **[]string** | Specifies the database statements to be executed rollback a create operation in the event of an error. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 



## Methods


### NewDatabaseWriteRoleRequest

`func NewDatabaseWriteRoleRequest() *DatabaseWriteRoleRequest`

NewDatabaseWriteRoleRequest instantiates a new DatabaseWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWriteRoleRequestWithDefaults

`func NewDatabaseWriteRoleRequestWithDefaults() *DatabaseWriteRoleRequest`

NewDatabaseWriteRoleRequestWithDefaults instantiates a new DatabaseWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCreationStatements

`func (o *DatabaseWriteRoleRequest) GetCreationStatements() []string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *DatabaseWriteRoleRequest) GetCreationStatementsOk() (*[]string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *DatabaseWriteRoleRequest) SetCreationStatements(v []string)`

SetCreationStatements sets CreationStatements field to given value.


### HasCreationStatements

`func (o *DatabaseWriteRoleRequest) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.




### GetCredentialConfig

`func (o *DatabaseWriteRoleRequest) GetCredentialConfig() map[string]interface{}`

GetCredentialConfig returns the CredentialConfig field if non-nil, zero value otherwise.

### GetCredentialConfigOk

`func (o *DatabaseWriteRoleRequest) GetCredentialConfigOk() (*map[string]interface{}, bool)`

GetCredentialConfigOk returns a tuple with the CredentialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialConfig

`func (o *DatabaseWriteRoleRequest) SetCredentialConfig(v map[string]interface{})`

SetCredentialConfig sets CredentialConfig field to given value.


### HasCredentialConfig

`func (o *DatabaseWriteRoleRequest) HasCredentialConfig() bool`

HasCredentialConfig returns a boolean if a field has been set.




### GetCredentialType

`func (o *DatabaseWriteRoleRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *DatabaseWriteRoleRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *DatabaseWriteRoleRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.


### HasCredentialType

`func (o *DatabaseWriteRoleRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.




### GetDbName

`func (o *DatabaseWriteRoleRequest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DatabaseWriteRoleRequest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DatabaseWriteRoleRequest) SetDbName(v string)`

SetDbName sets DbName field to given value.


### HasDbName

`func (o *DatabaseWriteRoleRequest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.




### GetDefaultTtl

`func (o *DatabaseWriteRoleRequest) GetDefaultTtl() string`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *DatabaseWriteRoleRequest) GetDefaultTtlOk() (*string, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *DatabaseWriteRoleRequest) SetDefaultTtl(v string)`

SetDefaultTtl sets DefaultTtl field to given value.


### HasDefaultTtl

`func (o *DatabaseWriteRoleRequest) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.




### GetMaxTtl

`func (o *DatabaseWriteRoleRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *DatabaseWriteRoleRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *DatabaseWriteRoleRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *DatabaseWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetRenewStatements

`func (o *DatabaseWriteRoleRequest) GetRenewStatements() []string`

GetRenewStatements returns the RenewStatements field if non-nil, zero value otherwise.

### GetRenewStatementsOk

`func (o *DatabaseWriteRoleRequest) GetRenewStatementsOk() (*[]string, bool)`

GetRenewStatementsOk returns a tuple with the RenewStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewStatements

`func (o *DatabaseWriteRoleRequest) SetRenewStatements(v []string)`

SetRenewStatements sets RenewStatements field to given value.


### HasRenewStatements

`func (o *DatabaseWriteRoleRequest) HasRenewStatements() bool`

HasRenewStatements returns a boolean if a field has been set.




### GetRevocationStatements

`func (o *DatabaseWriteRoleRequest) GetRevocationStatements() []string`

GetRevocationStatements returns the RevocationStatements field if non-nil, zero value otherwise.

### GetRevocationStatementsOk

`func (o *DatabaseWriteRoleRequest) GetRevocationStatementsOk() (*[]string, bool)`

GetRevocationStatementsOk returns a tuple with the RevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatements

`func (o *DatabaseWriteRoleRequest) SetRevocationStatements(v []string)`

SetRevocationStatements sets RevocationStatements field to given value.


### HasRevocationStatements

`func (o *DatabaseWriteRoleRequest) HasRevocationStatements() bool`

HasRevocationStatements returns a boolean if a field has been set.




### GetRollbackStatements

`func (o *DatabaseWriteRoleRequest) GetRollbackStatements() []string`

GetRollbackStatements returns the RollbackStatements field if non-nil, zero value otherwise.

### GetRollbackStatementsOk

`func (o *DatabaseWriteRoleRequest) GetRollbackStatementsOk() (*[]string, bool)`

GetRollbackStatementsOk returns a tuple with the RollbackStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackStatements

`func (o *DatabaseWriteRoleRequest) SetRollbackStatements(v []string)`

SetRollbackStatements sets RollbackStatements field to given value.


### HasRollbackStatements

`func (o *DatabaseWriteRoleRequest) HasRollbackStatements() bool`

HasRollbackStatements returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


