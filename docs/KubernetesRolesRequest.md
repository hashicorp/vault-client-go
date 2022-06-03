# KubernetesRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedKubernetesNamespaces** | **[]string** | A list of the Kubernetes namespaces in which credentials can be generated. If set to \&quot;*\&quot; all namespaces are allowed. | 
**ExtraAnnotations** | Pointer to **map[string]interface{}** | Additional annotations to apply to all generated Kubernetes objects. | [optional] 
**ExtraLabels** | Pointer to **map[string]interface{}** | Additional labels to apply to all generated Kubernetes objects. | [optional] 
**GeneratedRoleRules** | Pointer to **string** | The Role or ClusterRole rules to use when generating a role. Accepts either a JSON or YAML object. If set, the entire chain of Kubernetes objects will be generated. | [optional] 
**KubernetesRoleName** | Pointer to **string** | The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created. | [optional] 
**KubernetesRoleType** | Pointer to **string** | Specifies whether the Kubernetes role is a Role or ClusterRole. | [optional] [default to "Role"]
**NameTemplate** | Pointer to **string** | The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used. | [optional] 
**ServiceAccountName** | Pointer to **string** | The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes service account token will be created. | [optional] 
**TokenDefaultTtl** | Pointer to **int32** | The default ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default. | [optional] 

## Methods

### NewKubernetesRolesRequest

`func NewKubernetesRolesRequest(allowedKubernetesNamespaces []string, ) *KubernetesRolesRequest`

NewKubernetesRolesRequest instantiates a new KubernetesRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRolesRequestWithDefaults

`func NewKubernetesRolesRequestWithDefaults() *KubernetesRolesRequest`

NewKubernetesRolesRequestWithDefaults instantiates a new KubernetesRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedKubernetesNamespaces

`func (o *KubernetesRolesRequest) GetAllowedKubernetesNamespaces() []string`

GetAllowedKubernetesNamespaces returns the AllowedKubernetesNamespaces field if non-nil, zero value otherwise.

### GetAllowedKubernetesNamespacesOk

`func (o *KubernetesRolesRequest) GetAllowedKubernetesNamespacesOk() (*[]string, bool)`

GetAllowedKubernetesNamespacesOk returns a tuple with the AllowedKubernetesNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKubernetesNamespaces

`func (o *KubernetesRolesRequest) SetAllowedKubernetesNamespaces(v []string)`

SetAllowedKubernetesNamespaces sets AllowedKubernetesNamespaces field to given value.


### GetExtraAnnotations

`func (o *KubernetesRolesRequest) GetExtraAnnotations() map[string]interface{}`

GetExtraAnnotations returns the ExtraAnnotations field if non-nil, zero value otherwise.

### GetExtraAnnotationsOk

`func (o *KubernetesRolesRequest) GetExtraAnnotationsOk() (*map[string]interface{}, bool)`

GetExtraAnnotationsOk returns a tuple with the ExtraAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAnnotations

`func (o *KubernetesRolesRequest) SetExtraAnnotations(v map[string]interface{})`

SetExtraAnnotations sets ExtraAnnotations field to given value.

### HasExtraAnnotations

`func (o *KubernetesRolesRequest) HasExtraAnnotations() bool`

HasExtraAnnotations returns a boolean if a field has been set.

### GetExtraLabels

`func (o *KubernetesRolesRequest) GetExtraLabels() map[string]interface{}`

GetExtraLabels returns the ExtraLabels field if non-nil, zero value otherwise.

### GetExtraLabelsOk

`func (o *KubernetesRolesRequest) GetExtraLabelsOk() (*map[string]interface{}, bool)`

GetExtraLabelsOk returns a tuple with the ExtraLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraLabels

`func (o *KubernetesRolesRequest) SetExtraLabels(v map[string]interface{})`

SetExtraLabels sets ExtraLabels field to given value.

### HasExtraLabels

`func (o *KubernetesRolesRequest) HasExtraLabels() bool`

HasExtraLabels returns a boolean if a field has been set.

### GetGeneratedRoleRules

`func (o *KubernetesRolesRequest) GetGeneratedRoleRules() string`

GetGeneratedRoleRules returns the GeneratedRoleRules field if non-nil, zero value otherwise.

### GetGeneratedRoleRulesOk

`func (o *KubernetesRolesRequest) GetGeneratedRoleRulesOk() (*string, bool)`

GetGeneratedRoleRulesOk returns a tuple with the GeneratedRoleRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedRoleRules

`func (o *KubernetesRolesRequest) SetGeneratedRoleRules(v string)`

SetGeneratedRoleRules sets GeneratedRoleRules field to given value.

### HasGeneratedRoleRules

`func (o *KubernetesRolesRequest) HasGeneratedRoleRules() bool`

HasGeneratedRoleRules returns a boolean if a field has been set.

### GetKubernetesRoleName

`func (o *KubernetesRolesRequest) GetKubernetesRoleName() string`

GetKubernetesRoleName returns the KubernetesRoleName field if non-nil, zero value otherwise.

### GetKubernetesRoleNameOk

`func (o *KubernetesRolesRequest) GetKubernetesRoleNameOk() (*string, bool)`

GetKubernetesRoleNameOk returns a tuple with the KubernetesRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesRoleName

`func (o *KubernetesRolesRequest) SetKubernetesRoleName(v string)`

SetKubernetesRoleName sets KubernetesRoleName field to given value.

### HasKubernetesRoleName

`func (o *KubernetesRolesRequest) HasKubernetesRoleName() bool`

HasKubernetesRoleName returns a boolean if a field has been set.

### GetKubernetesRoleType

`func (o *KubernetesRolesRequest) GetKubernetesRoleType() string`

GetKubernetesRoleType returns the KubernetesRoleType field if non-nil, zero value otherwise.

### GetKubernetesRoleTypeOk

`func (o *KubernetesRolesRequest) GetKubernetesRoleTypeOk() (*string, bool)`

GetKubernetesRoleTypeOk returns a tuple with the KubernetesRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesRoleType

`func (o *KubernetesRolesRequest) SetKubernetesRoleType(v string)`

SetKubernetesRoleType sets KubernetesRoleType field to given value.

### HasKubernetesRoleType

`func (o *KubernetesRolesRequest) HasKubernetesRoleType() bool`

HasKubernetesRoleType returns a boolean if a field has been set.

### GetNameTemplate

`func (o *KubernetesRolesRequest) GetNameTemplate() string`

GetNameTemplate returns the NameTemplate field if non-nil, zero value otherwise.

### GetNameTemplateOk

`func (o *KubernetesRolesRequest) GetNameTemplateOk() (*string, bool)`

GetNameTemplateOk returns a tuple with the NameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTemplate

`func (o *KubernetesRolesRequest) SetNameTemplate(v string)`

SetNameTemplate sets NameTemplate field to given value.

### HasNameTemplate

`func (o *KubernetesRolesRequest) HasNameTemplate() bool`

HasNameTemplate returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *KubernetesRolesRequest) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *KubernetesRolesRequest) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *KubernetesRolesRequest) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *KubernetesRolesRequest) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetTokenDefaultTtl

`func (o *KubernetesRolesRequest) GetTokenDefaultTtl() int32`

GetTokenDefaultTtl returns the TokenDefaultTtl field if non-nil, zero value otherwise.

### GetTokenDefaultTtlOk

`func (o *KubernetesRolesRequest) GetTokenDefaultTtlOk() (*int32, bool)`

GetTokenDefaultTtlOk returns a tuple with the TokenDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDefaultTtl

`func (o *KubernetesRolesRequest) SetTokenDefaultTtl(v int32)`

SetTokenDefaultTtl sets TokenDefaultTtl field to given value.

### HasTokenDefaultTtl

`func (o *KubernetesRolesRequest) HasTokenDefaultTtl() bool`

HasTokenDefaultTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *KubernetesRolesRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *KubernetesRolesRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *KubernetesRolesRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *KubernetesRolesRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


