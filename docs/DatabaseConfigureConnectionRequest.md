# DatabaseConfigureConnectionRequest


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


### NewDatabaseConfigureConnectionRequest

`func NewDatabaseConfigureConnectionRequest() *DatabaseConfigureConnectionRequest`

NewDatabaseConfigureConnectionRequest instantiates a new DatabaseConfigureConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfigureConnectionRequestWithDefaults

`func NewDatabaseConfigureConnectionRequestWithDefaults() *DatabaseConfigureConnectionRequest`

NewDatabaseConfigureConnectionRequestWithDefaults instantiates a new DatabaseConfigureConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedRoles

`func (o *DatabaseConfigureConnectionRequest) GetAllowedRoles() []string`

GetAllowedRoles returns the AllowedRoles field if non-nil, zero value otherwise.

### GetAllowedRolesOk

`func (o *DatabaseConfigureConnectionRequest) GetAllowedRolesOk() (*[]string, bool)`

GetAllowedRolesOk returns a tuple with the AllowedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRoles

`func (o *DatabaseConfigureConnectionRequest) SetAllowedRoles(v []string)`

SetAllowedRoles sets AllowedRoles field to given value.


### HasAllowedRoles

`func (o *DatabaseConfigureConnectionRequest) HasAllowedRoles() bool`

HasAllowedRoles returns a boolean if a field has been set.




### GetPasswordPolicy

`func (o *DatabaseConfigureConnectionRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *DatabaseConfigureConnectionRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *DatabaseConfigureConnectionRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### HasPasswordPolicy

`func (o *DatabaseConfigureConnectionRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.




### GetPluginName

`func (o *DatabaseConfigureConnectionRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *DatabaseConfigureConnectionRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *DatabaseConfigureConnectionRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### HasPluginName

`func (o *DatabaseConfigureConnectionRequest) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.




### GetPluginVersion

`func (o *DatabaseConfigureConnectionRequest) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *DatabaseConfigureConnectionRequest) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *DatabaseConfigureConnectionRequest) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *DatabaseConfigureConnectionRequest) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetRootRotationStatements

`func (o *DatabaseConfigureConnectionRequest) GetRootRotationStatements() []string`

GetRootRotationStatements returns the RootRotationStatements field if non-nil, zero value otherwise.

### GetRootRotationStatementsOk

`func (o *DatabaseConfigureConnectionRequest) GetRootRotationStatementsOk() (*[]string, bool)`

GetRootRotationStatementsOk returns a tuple with the RootRotationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRotationStatements

`func (o *DatabaseConfigureConnectionRequest) SetRootRotationStatements(v []string)`

SetRootRotationStatements sets RootRotationStatements field to given value.


### HasRootRotationStatements

`func (o *DatabaseConfigureConnectionRequest) HasRootRotationStatements() bool`

HasRootRotationStatements returns a boolean if a field has been set.




### GetVerifyConnection

`func (o *DatabaseConfigureConnectionRequest) GetVerifyConnection() bool`

GetVerifyConnection returns the VerifyConnection field if non-nil, zero value otherwise.

### GetVerifyConnectionOk

`func (o *DatabaseConfigureConnectionRequest) GetVerifyConnectionOk() (*bool, bool)`

GetVerifyConnectionOk returns a tuple with the VerifyConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyConnection

`func (o *DatabaseConfigureConnectionRequest) SetVerifyConnection(v bool)`

SetVerifyConnection sets VerifyConnection field to given value.


### HasVerifyConnection

`func (o *DatabaseConfigureConnectionRequest) HasVerifyConnection() bool`

HasVerifyConnection returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


