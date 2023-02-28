# DatabaseWriteConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRoles** | Pointer to **[]string** | Comma separated string or array of the role names allowed to get creds from this database connection. If empty no roles are allowed. If \&quot;*\&quot; all roles are allowed. | [optional] 
**PasswordPolicy** | Pointer to **string** | Password policy to use when generating passwords. | [optional] 
**PluginName** | Pointer to **string** | The name of a builtin or previously registered plugin known to vault. This endpoint will create an instance of that plugin type. | [optional] 
**PluginVersion** | Pointer to **string** | The version of the plugin to use. | [optional] 
**RootRotationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to rotate the root user&#x27;s credentials. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**VerifyConnection** | Pointer to **bool** | If true, the connection details are verified by actually connecting to the database. Defaults to true. | [optional] [default to true]



## Methods


### NewDatabaseWriteConfigRequest

`func NewDatabaseWriteConfigRequest() *DatabaseWriteConfigRequest`

NewDatabaseWriteConfigRequest instantiates a new DatabaseWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWriteConfigRequestWithDefaults

`func NewDatabaseWriteConfigRequestWithDefaults() *DatabaseWriteConfigRequest`

NewDatabaseWriteConfigRequestWithDefaults instantiates a new DatabaseWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedRoles

`func (o *DatabaseWriteConfigRequest) GetAllowedRoles() []string`

GetAllowedRoles returns the AllowedRoles field if non-nil, zero value otherwise.

### GetAllowedRolesOk

`func (o *DatabaseWriteConfigRequest) GetAllowedRolesOk() (*[]string, bool)`

GetAllowedRolesOk returns a tuple with the AllowedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRoles

`func (o *DatabaseWriteConfigRequest) SetAllowedRoles(v []string)`

SetAllowedRoles sets AllowedRoles field to given value.


### HasAllowedRoles

`func (o *DatabaseWriteConfigRequest) HasAllowedRoles() bool`

HasAllowedRoles returns a boolean if a field has been set.




### GetPasswordPolicy

`func (o *DatabaseWriteConfigRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *DatabaseWriteConfigRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *DatabaseWriteConfigRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### HasPasswordPolicy

`func (o *DatabaseWriteConfigRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.




### GetPluginName

`func (o *DatabaseWriteConfigRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *DatabaseWriteConfigRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *DatabaseWriteConfigRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### HasPluginName

`func (o *DatabaseWriteConfigRequest) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.




### GetPluginVersion

`func (o *DatabaseWriteConfigRequest) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *DatabaseWriteConfigRequest) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *DatabaseWriteConfigRequest) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *DatabaseWriteConfigRequest) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetRootRotationStatements

`func (o *DatabaseWriteConfigRequest) GetRootRotationStatements() []string`

GetRootRotationStatements returns the RootRotationStatements field if non-nil, zero value otherwise.

### GetRootRotationStatementsOk

`func (o *DatabaseWriteConfigRequest) GetRootRotationStatementsOk() (*[]string, bool)`

GetRootRotationStatementsOk returns a tuple with the RootRotationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRotationStatements

`func (o *DatabaseWriteConfigRequest) SetRootRotationStatements(v []string)`

SetRootRotationStatements sets RootRotationStatements field to given value.


### HasRootRotationStatements

`func (o *DatabaseWriteConfigRequest) HasRootRotationStatements() bool`

HasRootRotationStatements returns a boolean if a field has been set.




### GetVerifyConnection

`func (o *DatabaseWriteConfigRequest) GetVerifyConnection() bool`

GetVerifyConnection returns the VerifyConnection field if non-nil, zero value otherwise.

### GetVerifyConnectionOk

`func (o *DatabaseWriteConfigRequest) GetVerifyConnectionOk() (*bool, bool)`

GetVerifyConnectionOk returns a tuple with the VerifyConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyConnection

`func (o *DatabaseWriteConfigRequest) SetVerifyConnection(v bool)`

SetVerifyConnection sets VerifyConnection field to given value.


### HasVerifyConnection

`func (o *DatabaseWriteConfigRequest) HasVerifyConnection() bool`

HasVerifyConnection returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


