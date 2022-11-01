# KerberosConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddGroupAliases** | Pointer to **bool** | If set to true, returns any groups found in LDAP as a group alias. | [optional] 
**Keytab** | Pointer to **string** | Base64 encoded keytab | [optional] 
**RemoveInstanceName** | Pointer to **bool** | Remove instance/FQDN from keytab principal names. | [optional] 
**ServiceAccount** | Pointer to **string** | Service Account | [optional] 

## Methods

### NewKerberosConfigRequest

`func NewKerberosConfigRequest() *KerberosConfigRequest`

NewKerberosConfigRequest instantiates a new KerberosConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosConfigRequestWithDefaults

`func NewKerberosConfigRequestWithDefaults() *KerberosConfigRequest`

NewKerberosConfigRequestWithDefaults instantiates a new KerberosConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddGroupAliases

`func (o *KerberosConfigRequest) GetAddGroupAliases() bool`

GetAddGroupAliases returns the AddGroupAliases field if non-nil, zero value otherwise.

### GetAddGroupAliasesOk

`func (o *KerberosConfigRequest) GetAddGroupAliasesOk() (*bool, bool)`

GetAddGroupAliasesOk returns a tuple with the AddGroupAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGroupAliases

`func (o *KerberosConfigRequest) SetAddGroupAliases(v bool)`

SetAddGroupAliases sets AddGroupAliases field to given value.

### HasAddGroupAliases

`func (o *KerberosConfigRequest) HasAddGroupAliases() bool`

HasAddGroupAliases returns a boolean if a field has been set.

### GetKeytab

`func (o *KerberosConfigRequest) GetKeytab() string`

GetKeytab returns the Keytab field if non-nil, zero value otherwise.

### GetKeytabOk

`func (o *KerberosConfigRequest) GetKeytabOk() (*string, bool)`

GetKeytabOk returns a tuple with the Keytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytab

`func (o *KerberosConfigRequest) SetKeytab(v string)`

SetKeytab sets Keytab field to given value.

### HasKeytab

`func (o *KerberosConfigRequest) HasKeytab() bool`

HasKeytab returns a boolean if a field has been set.

### GetRemoveInstanceName

`func (o *KerberosConfigRequest) GetRemoveInstanceName() bool`

GetRemoveInstanceName returns the RemoveInstanceName field if non-nil, zero value otherwise.

### GetRemoveInstanceNameOk

`func (o *KerberosConfigRequest) GetRemoveInstanceNameOk() (*bool, bool)`

GetRemoveInstanceNameOk returns a tuple with the RemoveInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveInstanceName

`func (o *KerberosConfigRequest) SetRemoveInstanceName(v bool)`

SetRemoveInstanceName sets RemoveInstanceName field to given value.

### HasRemoveInstanceName

`func (o *KerberosConfigRequest) HasRemoveInstanceName() bool`

HasRemoveInstanceName returns a boolean if a field has been set.

### GetServiceAccount

`func (o *KerberosConfigRequest) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *KerberosConfigRequest) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *KerberosConfigRequest) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *KerberosConfigRequest) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


