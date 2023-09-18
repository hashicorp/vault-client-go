# KubernetesWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedKubernetesNamespaceSelector** | Pointer to **string** | A label selector for Kubernetes namespaces in which credentials can be generated. Accepts either a JSON or YAML object. If set with allowed_kubernetes_namespaces, the conditions are conjuncted. | [optional] 
**AllowedKubernetesNamespaces** | Pointer to **[]string** | A list of the Kubernetes namespaces in which credentials can be generated. If set to \&quot;*\&quot; all namespaces are allowed. | [optional] 
**ExtraAnnotations** | Pointer to **map[string]interface{}** | Additional annotations to apply to all generated Kubernetes objects. | [optional] 
**ExtraLabels** | Pointer to **map[string]interface{}** | Additional labels to apply to all generated Kubernetes objects. | [optional] 
**GeneratedRoleRules** | Pointer to **string** | The Role or ClusterRole rules to use when generating a role. Accepts either a JSON or YAML object. If set, the entire chain of Kubernetes objects will be generated. | [optional] 
**KubernetesRoleName** | Pointer to **string** | The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created. | [optional] 
**KubernetesRoleType** | Pointer to **string** | Specifies whether the Kubernetes role is a Role or ClusterRole. | [optional] [default to "Role"]
**NameTemplate** | Pointer to **string** | The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used. | [optional] 
**ServiceAccountName** | Pointer to **string** | The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes service account token will be created. | [optional] 
**TokenDefaultAudiences** | Pointer to **[]string** | The default audiences for generated Kubernetes service account tokens. If not set or set to \&quot;\&quot;, will use k8s cluster default. | [optional] 
**TokenDefaultTtl** | Pointer to **string** | The default ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


