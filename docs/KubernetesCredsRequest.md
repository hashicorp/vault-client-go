# KubernetesCredsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRoleBinding** | Pointer to **bool** | If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace. Requires the Vault role to have kubernetes_role_type set to ClusterRole. | [optional] 
**KubernetesNamespace** | **string** | The name of the Kubernetes namespace in which to generate the credentials | 
**Ttl** | Pointer to **int32** | The TTL of the generated credentials | [optional] 

## Methods

### NewKubernetesCredsRequest

`func NewKubernetesCredsRequest(kubernetesNamespace string, ) *KubernetesCredsRequest`

NewKubernetesCredsRequest instantiates a new KubernetesCredsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCredsRequestWithDefaults

`func NewKubernetesCredsRequestWithDefaults() *KubernetesCredsRequest`

NewKubernetesCredsRequestWithDefaults instantiates a new KubernetesCredsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRoleBinding

`func (o *KubernetesCredsRequest) GetClusterRoleBinding() bool`

GetClusterRoleBinding returns the ClusterRoleBinding field if non-nil, zero value otherwise.

### GetClusterRoleBindingOk

`func (o *KubernetesCredsRequest) GetClusterRoleBindingOk() (*bool, bool)`

GetClusterRoleBindingOk returns a tuple with the ClusterRoleBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleBinding

`func (o *KubernetesCredsRequest) SetClusterRoleBinding(v bool)`

SetClusterRoleBinding sets ClusterRoleBinding field to given value.

### HasClusterRoleBinding

`func (o *KubernetesCredsRequest) HasClusterRoleBinding() bool`

HasClusterRoleBinding returns a boolean if a field has been set.

### GetKubernetesNamespace

`func (o *KubernetesCredsRequest) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *KubernetesCredsRequest) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *KubernetesCredsRequest) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.


### GetTtl

`func (o *KubernetesCredsRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KubernetesCredsRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KubernetesCredsRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KubernetesCredsRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


