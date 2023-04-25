# Secrets

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliCloudConfigure**](SecretsApi.md#AliCloudConfigure) | **Post** /{alicloud_mount_path}/config | 
[**AliCloudDeleteConfiguration**](SecretsApi.md#AliCloudDeleteConfiguration) | **Delete** /{alicloud_mount_path}/config | 
[**AliCloudDeleteRole**](SecretsApi.md#AliCloudDeleteRole) | **Delete** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AliCloudGenerateCredentials**](SecretsApi.md#AliCloudGenerateCredentials) | **Get** /{alicloud_mount_path}/creds/{name} | Generate an API key or STS credential using the given role&#x27;s configuration.&#x27;
[**AliCloudListRoles**](SecretsApi.md#AliCloudListRoles) | **Get** /{alicloud_mount_path}/role | List the existing roles in this backend.
[**AliCloudReadConfiguration**](SecretsApi.md#AliCloudReadConfiguration) | **Get** /{alicloud_mount_path}/config | 
[**AliCloudReadRole**](SecretsApi.md#AliCloudReadRole) | **Get** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AliCloudWriteRole**](SecretsApi.md#AliCloudWriteRole) | **Post** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AwsConfigureLease**](SecretsApi.md#AwsConfigureLease) | **Post** /{aws_mount_path}/config/lease | 
[**AwsConfigureRootIamCredentials**](SecretsApi.md#AwsConfigureRootIamCredentials) | **Post** /{aws_mount_path}/config/root | 
[**AwsDeleteRole**](SecretsApi.md#AwsDeleteRole) | **Delete** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AwsGenerateCredentials**](SecretsApi.md#AwsGenerateCredentials) | **Get** /{aws_mount_path}/creds/{name} | 
[**AwsGenerateCredentials2**](SecretsApi.md#AwsGenerateCredentials2) | **Post** /{aws_mount_path}/creds/{name} | 
[**AwsGenerateStsCredentials**](SecretsApi.md#AwsGenerateStsCredentials) | **Get** /{aws_mount_path}/sts/{name} | 
[**AwsGenerateStsCredentials2**](SecretsApi.md#AwsGenerateStsCredentials2) | **Post** /{aws_mount_path}/sts/{name} | 
[**AwsListRoles**](SecretsApi.md#AwsListRoles) | **Get** /{aws_mount_path}/roles | List the existing roles in this backend
[**AwsReadLeaseConfiguration**](SecretsApi.md#AwsReadLeaseConfiguration) | **Get** /{aws_mount_path}/config/lease | 
[**AwsReadRole**](SecretsApi.md#AwsReadRole) | **Get** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AwsReadRootIamCredentialsConfiguration**](SecretsApi.md#AwsReadRootIamCredentialsConfiguration) | **Get** /{aws_mount_path}/config/root | 
[**AwsRotateRootIamCredentials**](SecretsApi.md#AwsRotateRootIamCredentials) | **Post** /{aws_mount_path}/config/rotate-root | 
[**AwsWriteRole**](SecretsApi.md#AwsWriteRole) | **Post** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AzureConfigure**](SecretsApi.md#AzureConfigure) | **Post** /{azure_mount_path}/config | 
[**AzureDeleteConfiguration**](SecretsApi.md#AzureDeleteConfiguration) | **Delete** /{azure_mount_path}/config | 
[**AzureDeleteRole**](SecretsApi.md#AzureDeleteRole) | **Delete** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**AzureListRoles**](SecretsApi.md#AzureListRoles) | **Get** /{azure_mount_path}/roles | List existing roles.
[**AzureReadConfiguration**](SecretsApi.md#AzureReadConfiguration) | **Get** /{azure_mount_path}/config | 
[**AzureReadRole**](SecretsApi.md#AzureReadRole) | **Get** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**AzureRequestServicePrincipalCredentials**](SecretsApi.md#AzureRequestServicePrincipalCredentials) | **Get** /{azure_mount_path}/creds/{role} | 
[**AzureRotateRoot**](SecretsApi.md#AzureRotateRoot) | **Post** /{azure_mount_path}/rotate-root | 
[**AzureWriteRole**](SecretsApi.md#AzureWriteRole) | **Post** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**ConsulConfigureAccess**](SecretsApi.md#ConsulConfigureAccess) | **Post** /{consul_mount_path}/config/access | 
[**ConsulDeleteRole**](SecretsApi.md#ConsulDeleteRole) | **Delete** /{consul_mount_path}/roles/{name} | 
[**ConsulGenerateCredentials**](SecretsApi.md#ConsulGenerateCredentials) | **Get** /{consul_mount_path}/creds/{role} | 
[**ConsulListRoles**](SecretsApi.md#ConsulListRoles) | **Get** /{consul_mount_path}/roles | 
[**ConsulReadAccessConfiguration**](SecretsApi.md#ConsulReadAccessConfiguration) | **Get** /{consul_mount_path}/config/access | 
[**ConsulReadRole**](SecretsApi.md#ConsulReadRole) | **Get** /{consul_mount_path}/roles/{name} | 
[**ConsulWriteRole**](SecretsApi.md#ConsulWriteRole) | **Post** /{consul_mount_path}/roles/{name} | 
[**CubbyholeDelete**](SecretsApi.md#CubbyholeDelete) | **Delete** /cubbyhole/{path} | Deletes the secret at the specified location.
[**CubbyholeRead**](SecretsApi.md#CubbyholeRead) | **Get** /cubbyhole/{path} | Retrieve the secret at the specified location.
[**CubbyholeWrite**](SecretsApi.md#CubbyholeWrite) | **Post** /cubbyhole/{path} | Store a secret at the specified location.
[**DatabaseConfigureConnection**](SecretsApi.md#DatabaseConfigureConnection) | **Post** /{database_mount_path}/config/{name} | 
[**DatabaseDeleteConnectionConfiguration**](SecretsApi.md#DatabaseDeleteConnectionConfiguration) | **Delete** /{database_mount_path}/config/{name} | 
[**DatabaseDeleteRole**](SecretsApi.md#DatabaseDeleteRole) | **Delete** /{database_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**DatabaseDeleteStaticRole**](SecretsApi.md#DatabaseDeleteStaticRole) | **Delete** /{database_mount_path}/static-roles/{name} | Manage the static roles that can be created with this backend.
[**DatabaseGenerateCredentials**](SecretsApi.md#DatabaseGenerateCredentials) | **Get** /{database_mount_path}/creds/{name} | Request database credentials for a certain role.
[**DatabaseListConnections**](SecretsApi.md#DatabaseListConnections) | **Get** /{database_mount_path}/config | Configure connection details to a database plugin.
[**DatabaseListRoles**](SecretsApi.md#DatabaseListRoles) | **Get** /{database_mount_path}/roles | Manage the roles that can be created with this backend.
[**DatabaseListStaticRoles**](SecretsApi.md#DatabaseListStaticRoles) | **Get** /{database_mount_path}/static-roles | Manage the static roles that can be created with this backend.
[**DatabaseReadConnectionConfiguration**](SecretsApi.md#DatabaseReadConnectionConfiguration) | **Get** /{database_mount_path}/config/{name} | 
[**DatabaseReadRole**](SecretsApi.md#DatabaseReadRole) | **Get** /{database_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**DatabaseReadStaticRole**](SecretsApi.md#DatabaseReadStaticRole) | **Get** /{database_mount_path}/static-roles/{name} | Manage the static roles that can be created with this backend.
[**DatabaseReadStaticRoleCredentials**](SecretsApi.md#DatabaseReadStaticRoleCredentials) | **Get** /{database_mount_path}/static-creds/{name} | Request database credentials for a certain static role. These credentials are rotated periodically.
[**DatabaseResetConnection**](SecretsApi.md#DatabaseResetConnection) | **Post** /{database_mount_path}/reset/{name} | Resets a database plugin.
[**DatabaseRotateRootCredentials**](SecretsApi.md#DatabaseRotateRootCredentials) | **Post** /{database_mount_path}/rotate-root/{name} | 
[**DatabaseRotateStaticRoleCredentials**](SecretsApi.md#DatabaseRotateStaticRoleCredentials) | **Post** /{database_mount_path}/rotate-role/{name} | 
[**DatabaseWriteRole**](SecretsApi.md#DatabaseWriteRole) | **Post** /{database_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**DatabaseWriteStaticRole**](SecretsApi.md#DatabaseWriteStaticRole) | **Post** /{database_mount_path}/static-roles/{name} | Manage the static roles that can be created with this backend.
[**GoogleCloudConfigure**](SecretsApi.md#GoogleCloudConfigure) | **Post** /{gcp_mount_path}/config | 
[**GoogleCloudDeleteImpersonatedAccount**](SecretsApi.md#GoogleCloudDeleteImpersonatedAccount) | **Delete** /{gcp_mount_path}/impersonated-account/{name} | 
[**GoogleCloudDeleteRoleset**](SecretsApi.md#GoogleCloudDeleteRoleset) | **Delete** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudDeleteStaticAccount**](SecretsApi.md#GoogleCloudDeleteStaticAccount) | **Delete** /{gcp_mount_path}/static-account/{name} | 
[**GoogleCloudGenerateImpersonatedAccountAccessToken**](SecretsApi.md#GoogleCloudGenerateImpersonatedAccountAccessToken) | **Get** /{gcp_mount_path}/impersonated-account/{name}/token | 
[**GoogleCloudGenerateImpersonatedAccountAccessToken2**](SecretsApi.md#GoogleCloudGenerateImpersonatedAccountAccessToken2) | **Post** /{gcp_mount_path}/impersonated-account/{name}/token | 
[**GoogleCloudGenerateRolesetAccessToken**](SecretsApi.md#GoogleCloudGenerateRolesetAccessToken) | **Get** /{gcp_mount_path}/roleset/{roleset}/token | 
[**GoogleCloudGenerateRolesetAccessToken2**](SecretsApi.md#GoogleCloudGenerateRolesetAccessToken2) | **Get** /{gcp_mount_path}/token/{roleset} | 
[**GoogleCloudGenerateRolesetAccessTokenWithParameters**](SecretsApi.md#GoogleCloudGenerateRolesetAccessTokenWithParameters) | **Post** /{gcp_mount_path}/roleset/{roleset}/token | 
[**GoogleCloudGenerateRolesetAccessTokenWithParameters2**](SecretsApi.md#GoogleCloudGenerateRolesetAccessTokenWithParameters2) | **Post** /{gcp_mount_path}/token/{roleset} | 
[**GoogleCloudGenerateRolesetKey**](SecretsApi.md#GoogleCloudGenerateRolesetKey) | **Get** /{gcp_mount_path}/roleset/{roleset}/key | 
[**GoogleCloudGenerateRolesetKey2**](SecretsApi.md#GoogleCloudGenerateRolesetKey2) | **Get** /{gcp_mount_path}/key/{roleset} | 
[**GoogleCloudGenerateRolesetKeyWithParameters**](SecretsApi.md#GoogleCloudGenerateRolesetKeyWithParameters) | **Post** /{gcp_mount_path}/roleset/{roleset}/key | 
[**GoogleCloudGenerateRolesetKeyWithParameters2**](SecretsApi.md#GoogleCloudGenerateRolesetKeyWithParameters2) | **Post** /{gcp_mount_path}/key/{roleset} | 
[**GoogleCloudGenerateStaticAccountAccessToken**](SecretsApi.md#GoogleCloudGenerateStaticAccountAccessToken) | **Get** /{gcp_mount_path}/static-account/{name}/token | 
[**GoogleCloudGenerateStaticAccountAccessTokenWithParameters**](SecretsApi.md#GoogleCloudGenerateStaticAccountAccessTokenWithParameters) | **Post** /{gcp_mount_path}/static-account/{name}/token | 
[**GoogleCloudGenerateStaticAccountKey**](SecretsApi.md#GoogleCloudGenerateStaticAccountKey) | **Get** /{gcp_mount_path}/static-account/{name}/key | 
[**GoogleCloudGenerateStaticAccountKeyWithParameters**](SecretsApi.md#GoogleCloudGenerateStaticAccountKeyWithParameters) | **Post** /{gcp_mount_path}/static-account/{name}/key | 
[**GoogleCloudKmsConfigure**](SecretsApi.md#GoogleCloudKmsConfigure) | **Post** /{gcpkms_mount_path}/config | 
[**GoogleCloudKmsConfigureKey**](SecretsApi.md#GoogleCloudKmsConfigureKey) | **Post** /{gcpkms_mount_path}/keys/config/{key} | 
[**GoogleCloudKmsDecrypt**](SecretsApi.md#GoogleCloudKmsDecrypt) | **Post** /{gcpkms_mount_path}/decrypt/{key} | Decrypt a ciphertext value using a named key
[**GoogleCloudKmsDeleteConfiguration**](SecretsApi.md#GoogleCloudKmsDeleteConfiguration) | **Delete** /{gcpkms_mount_path}/config | 
[**GoogleCloudKmsDeleteKey**](SecretsApi.md#GoogleCloudKmsDeleteKey) | **Delete** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudKmsDeregisterKey**](SecretsApi.md#GoogleCloudKmsDeregisterKey) | **Post** /{gcpkms_mount_path}/keys/deregister/{key} | 
[**GoogleCloudKmsDeregisterKey2**](SecretsApi.md#GoogleCloudKmsDeregisterKey2) | **Delete** /{gcpkms_mount_path}/keys/deregister/{key} | 
[**GoogleCloudKmsEncrypt**](SecretsApi.md#GoogleCloudKmsEncrypt) | **Post** /{gcpkms_mount_path}/encrypt/{key} | Encrypt a plaintext value using a named key
[**GoogleCloudKmsListKeys**](SecretsApi.md#GoogleCloudKmsListKeys) | **Get** /{gcpkms_mount_path}/keys | List named keys
[**GoogleCloudKmsReadConfiguration**](SecretsApi.md#GoogleCloudKmsReadConfiguration) | **Get** /{gcpkms_mount_path}/config | 
[**GoogleCloudKmsReadKey**](SecretsApi.md#GoogleCloudKmsReadKey) | **Get** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudKmsReadKeyConfiguration**](SecretsApi.md#GoogleCloudKmsReadKeyConfiguration) | **Get** /{gcpkms_mount_path}/keys/config/{key} | 
[**GoogleCloudKmsReencrypt**](SecretsApi.md#GoogleCloudKmsReencrypt) | **Post** /{gcpkms_mount_path}/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**GoogleCloudKmsRegisterKey**](SecretsApi.md#GoogleCloudKmsRegisterKey) | **Post** /{gcpkms_mount_path}/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**GoogleCloudKmsRetrievePublicKey**](SecretsApi.md#GoogleCloudKmsRetrievePublicKey) | **Get** /{gcpkms_mount_path}/pubkey/{key} | Retrieve the public key associated with the named key
[**GoogleCloudKmsRotateKey**](SecretsApi.md#GoogleCloudKmsRotateKey) | **Post** /{gcpkms_mount_path}/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**GoogleCloudKmsSign**](SecretsApi.md#GoogleCloudKmsSign) | **Post** /{gcpkms_mount_path}/sign/{key} | Signs a message or digest using a named key
[**GoogleCloudKmsTrimKeyVersions**](SecretsApi.md#GoogleCloudKmsTrimKeyVersions) | **Post** /{gcpkms_mount_path}/keys/trim/{key} | 
[**GoogleCloudKmsTrimKeyVersions2**](SecretsApi.md#GoogleCloudKmsTrimKeyVersions2) | **Delete** /{gcpkms_mount_path}/keys/trim/{key} | 
[**GoogleCloudKmsVerify**](SecretsApi.md#GoogleCloudKmsVerify) | **Post** /{gcpkms_mount_path}/verify/{key} | Verify a signature using a named key
[**GoogleCloudKmsWriteKey**](SecretsApi.md#GoogleCloudKmsWriteKey) | **Post** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudListImpersonatedAccounts**](SecretsApi.md#GoogleCloudListImpersonatedAccounts) | **Get** /{gcp_mount_path}/impersonated-account | 
[**GoogleCloudListImpersonatedAccounts2**](SecretsApi.md#GoogleCloudListImpersonatedAccounts2) | **Get** /{gcp_mount_path}/impersonated-accounts | 
[**GoogleCloudListRolesets**](SecretsApi.md#GoogleCloudListRolesets) | **Get** /{gcp_mount_path}/roleset | 
[**GoogleCloudListRolesets2**](SecretsApi.md#GoogleCloudListRolesets2) | **Get** /{gcp_mount_path}/rolesets | 
[**GoogleCloudListStaticAccounts**](SecretsApi.md#GoogleCloudListStaticAccounts) | **Get** /{gcp_mount_path}/static-account | 
[**GoogleCloudListStaticAccounts2**](SecretsApi.md#GoogleCloudListStaticAccounts2) | **Get** /{gcp_mount_path}/static-accounts | 
[**GoogleCloudReadConfiguration**](SecretsApi.md#GoogleCloudReadConfiguration) | **Get** /{gcp_mount_path}/config | 
[**GoogleCloudReadImpersonatedAccount**](SecretsApi.md#GoogleCloudReadImpersonatedAccount) | **Get** /{gcp_mount_path}/impersonated-account/{name} | 
[**GoogleCloudReadRoleset**](SecretsApi.md#GoogleCloudReadRoleset) | **Get** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudReadStaticAccount**](SecretsApi.md#GoogleCloudReadStaticAccount) | **Get** /{gcp_mount_path}/static-account/{name} | 
[**GoogleCloudRotateRoleset**](SecretsApi.md#GoogleCloudRotateRoleset) | **Post** /{gcp_mount_path}/roleset/{name}/rotate | 
[**GoogleCloudRotateRolesetKey**](SecretsApi.md#GoogleCloudRotateRolesetKey) | **Post** /{gcp_mount_path}/roleset/{name}/rotate-key | 
[**GoogleCloudRotateRootCredentials**](SecretsApi.md#GoogleCloudRotateRootCredentials) | **Post** /{gcp_mount_path}/config/rotate-root | 
[**GoogleCloudRotateStaticAccountKey**](SecretsApi.md#GoogleCloudRotateStaticAccountKey) | **Post** /{gcp_mount_path}/static-account/{name}/rotate-key | 
[**GoogleCloudWriteImpersonatedAccount**](SecretsApi.md#GoogleCloudWriteImpersonatedAccount) | **Post** /{gcp_mount_path}/impersonated-account/{name} | 
[**GoogleCloudWriteRoleset**](SecretsApi.md#GoogleCloudWriteRoleset) | **Post** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudWriteStaticAccount**](SecretsApi.md#GoogleCloudWriteStaticAccount) | **Post** /{gcp_mount_path}/static-account/{name} | 
[**KubernetesCheckConfiguration**](SecretsApi.md#KubernetesCheckConfiguration) | **Get** /{kubernetes_mount_path}/check | 
[**KubernetesConfigure**](SecretsApi.md#KubernetesConfigure) | **Post** /{kubernetes_mount_path}/config | 
[**KubernetesDeleteConfiguration**](SecretsApi.md#KubernetesDeleteConfiguration) | **Delete** /{kubernetes_mount_path}/config | 
[**KubernetesDeleteRole**](SecretsApi.md#KubernetesDeleteRole) | **Delete** /{kubernetes_mount_path}/roles/{name} | 
[**KubernetesGenerateCredentials**](SecretsApi.md#KubernetesGenerateCredentials) | **Post** /{kubernetes_mount_path}/creds/{name} | 
[**KubernetesListRoles**](SecretsApi.md#KubernetesListRoles) | **Get** /{kubernetes_mount_path}/roles | 
[**KubernetesReadConfiguration**](SecretsApi.md#KubernetesReadConfiguration) | **Get** /{kubernetes_mount_path}/config | 
[**KubernetesReadRole**](SecretsApi.md#KubernetesReadRole) | **Get** /{kubernetes_mount_path}/roles/{name} | 
[**KubernetesWriteRole**](SecretsApi.md#KubernetesWriteRole) | **Post** /{kubernetes_mount_path}/roles/{name} | 
[**KvV1Delete**](SecretsApi.md#KvV1Delete) | **Delete** /{kv-v1_mount_path}/{path} | 
[**KvV1Read**](SecretsApi.md#KvV1Read) | **Get** /{kv-v1_mount_path}/{path} | 
[**KvV1Write**](SecretsApi.md#KvV1Write) | **Post** /{kv-v1_mount_path}/{path} | 
[**KvV2Configure**](SecretsApi.md#KvV2Configure) | **Post** /{kv-v2_mount_path}/config | Configure backend level settings that are applied to every key in the key-value store.
[**KvV2Delete**](SecretsApi.md#KvV2Delete) | **Delete** /{kv-v2_mount_path}/data/{path} | 
[**KvV2DeleteMetadata**](SecretsApi.md#KvV2DeleteMetadata) | **Delete** /{kv-v2_mount_path}/metadata/{path} | 
[**KvV2DeleteVersions**](SecretsApi.md#KvV2DeleteVersions) | **Post** /{kv-v2_mount_path}/delete/{path} | 
[**KvV2DestroyVersions**](SecretsApi.md#KvV2DestroyVersions) | **Post** /{kv-v2_mount_path}/destroy/{path} | 
[**KvV2Read**](SecretsApi.md#KvV2Read) | **Get** /{kv-v2_mount_path}/data/{path} | 
[**KvV2ReadConfiguration**](SecretsApi.md#KvV2ReadConfiguration) | **Get** /{kv-v2_mount_path}/config | Read the backend level settings.
[**KvV2ReadMetadata**](SecretsApi.md#KvV2ReadMetadata) | **Get** /{kv-v2_mount_path}/metadata/{path} | 
[**KvV2ReadSubkeys**](SecretsApi.md#KvV2ReadSubkeys) | **Get** /{kv-v2_mount_path}/subkeys/{path} | 
[**KvV2UndeleteVersions**](SecretsApi.md#KvV2UndeleteVersions) | **Post** /{kv-v2_mount_path}/undelete/{path} | 
[**KvV2Write**](SecretsApi.md#KvV2Write) | **Post** /{kv-v2_mount_path}/data/{path} | 
[**KvV2WriteMetadata**](SecretsApi.md#KvV2WriteMetadata) | **Post** /{kv-v2_mount_path}/metadata/{path} | 
[**LdapConfigure**](SecretsApi.md#LdapConfigure) | **Post** /{ldap_mount_path}/config | 
[**LdapDeleteConfiguration**](SecretsApi.md#LdapDeleteConfiguration) | **Delete** /{ldap_mount_path}/config | 
[**LdapDeleteDynamicRole**](SecretsApi.md#LdapDeleteDynamicRole) | **Delete** /{ldap_mount_path}/role/{name} | 
[**LdapDeleteStaticRole**](SecretsApi.md#LdapDeleteStaticRole) | **Delete** /{ldap_mount_path}/static-role/{name} | 
[**LdapLibraryCheckIn**](SecretsApi.md#LdapLibraryCheckIn) | **Post** /{ldap_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**LdapLibraryCheckOut**](SecretsApi.md#LdapLibraryCheckOut) | **Post** /{ldap_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**LdapLibraryCheckStatus**](SecretsApi.md#LdapLibraryCheckStatus) | **Get** /{ldap_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**LdapLibraryConfigure**](SecretsApi.md#LdapLibraryConfigure) | **Post** /{ldap_mount_path}/library/{name} | Update a library set.
[**LdapLibraryDelete**](SecretsApi.md#LdapLibraryDelete) | **Delete** /{ldap_mount_path}/library/{name} | Delete a library set.
[**LdapLibraryForceCheckIn**](SecretsApi.md#LdapLibraryForceCheckIn) | **Post** /{ldap_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**LdapLibraryList**](SecretsApi.md#LdapLibraryList) | **Get** /{ldap_mount_path}/library | 
[**LdapLibraryRead**](SecretsApi.md#LdapLibraryRead) | **Get** /{ldap_mount_path}/library/{name} | Read a library set.
[**LdapListDynamicRoles**](SecretsApi.md#LdapListDynamicRoles) | **Get** /{ldap_mount_path}/role | 
[**LdapListStaticRoles**](SecretsApi.md#LdapListStaticRoles) | **Get** /{ldap_mount_path}/static-role | 
[**LdapReadConfiguration**](SecretsApi.md#LdapReadConfiguration) | **Get** /{ldap_mount_path}/config | 
[**LdapReadDynamicRole**](SecretsApi.md#LdapReadDynamicRole) | **Get** /{ldap_mount_path}/role/{name} | 
[**LdapReadStaticRole**](SecretsApi.md#LdapReadStaticRole) | **Get** /{ldap_mount_path}/static-role/{name} | 
[**LdapRequestDynamicRoleCredentials**](SecretsApi.md#LdapRequestDynamicRoleCredentials) | **Get** /{ldap_mount_path}/creds/{name} | 
[**LdapRequestStaticRoleCredentials**](SecretsApi.md#LdapRequestStaticRoleCredentials) | **Get** /{ldap_mount_path}/static-cred/{name} | 
[**LdapRotateRootCredentials**](SecretsApi.md#LdapRotateRootCredentials) | **Post** /{ldap_mount_path}/rotate-root | 
[**LdapRotateStaticRole**](SecretsApi.md#LdapRotateStaticRole) | **Post** /{ldap_mount_path}/rotate-role/{name} | 
[**LdapWriteDynamicRole**](SecretsApi.md#LdapWriteDynamicRole) | **Post** /{ldap_mount_path}/role/{name} | 
[**LdapWriteStaticRole**](SecretsApi.md#LdapWriteStaticRole) | **Post** /{ldap_mount_path}/static-role/{name} | 
[**MongoDbAtlasConfigure**](SecretsApi.md#MongoDbAtlasConfigure) | **Post** /{mongodbatlas_mount_path}/config | 
[**MongoDbAtlasDeleteRole**](SecretsApi.md#MongoDbAtlasDeleteRole) | **Delete** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**MongoDbAtlasGenerateCredentials**](SecretsApi.md#MongoDbAtlasGenerateCredentials) | **Get** /{mongodbatlas_mount_path}/creds/{name} | 
[**MongoDbAtlasGenerateCredentials2**](SecretsApi.md#MongoDbAtlasGenerateCredentials2) | **Post** /{mongodbatlas_mount_path}/creds/{name} | 
[**MongoDbAtlasListRoles**](SecretsApi.md#MongoDbAtlasListRoles) | **Get** /{mongodbatlas_mount_path}/roles | List the existing roles in this backend
[**MongoDbAtlasReadConfiguration**](SecretsApi.md#MongoDbAtlasReadConfiguration) | **Get** /{mongodbatlas_mount_path}/config | 
[**MongoDbAtlasReadRole**](SecretsApi.md#MongoDbAtlasReadRole) | **Get** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**MongoDbAtlasWriteRole**](SecretsApi.md#MongoDbAtlasWriteRole) | **Post** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**NomadConfigureAccess**](SecretsApi.md#NomadConfigureAccess) | **Post** /{nomad_mount_path}/config/access | 
[**NomadConfigureLease**](SecretsApi.md#NomadConfigureLease) | **Post** /{nomad_mount_path}/config/lease | 
[**NomadDeleteAccessConfiguration**](SecretsApi.md#NomadDeleteAccessConfiguration) | **Delete** /{nomad_mount_path}/config/access | 
[**NomadDeleteLeaseConfiguration**](SecretsApi.md#NomadDeleteLeaseConfiguration) | **Delete** /{nomad_mount_path}/config/lease | 
[**NomadDeleteRole**](SecretsApi.md#NomadDeleteRole) | **Delete** /{nomad_mount_path}/role/{name} | 
[**NomadGenerateCredentials**](SecretsApi.md#NomadGenerateCredentials) | **Get** /{nomad_mount_path}/creds/{name} | 
[**NomadListRoles**](SecretsApi.md#NomadListRoles) | **Get** /{nomad_mount_path}/role | 
[**NomadReadAccessConfiguration**](SecretsApi.md#NomadReadAccessConfiguration) | **Get** /{nomad_mount_path}/config/access | 
[**NomadReadLeaseConfiguration**](SecretsApi.md#NomadReadLeaseConfiguration) | **Get** /{nomad_mount_path}/config/lease | 
[**NomadReadRole**](SecretsApi.md#NomadReadRole) | **Get** /{nomad_mount_path}/role/{name} | 
[**NomadWriteRole**](SecretsApi.md#NomadWriteRole) | **Post** /{nomad_mount_path}/role/{name} | 
[**PkiConfigureAutoTidy**](SecretsApi.md#PkiConfigureAutoTidy) | **Post** /{pki_mount_path}/config/auto-tidy | 
[**PkiConfigureCa**](SecretsApi.md#PkiConfigureCa) | **Post** /{pki_mount_path}/config/ca | 
[**PkiConfigureCluster**](SecretsApi.md#PkiConfigureCluster) | **Post** /{pki_mount_path}/config/cluster | 
[**PkiConfigureCrl**](SecretsApi.md#PkiConfigureCrl) | **Post** /{pki_mount_path}/config/crl | 
[**PkiConfigureIssuers**](SecretsApi.md#PkiConfigureIssuers) | **Post** /{pki_mount_path}/config/issuers | 
[**PkiConfigureKeys**](SecretsApi.md#PkiConfigureKeys) | **Post** /{pki_mount_path}/config/keys | 
[**PkiConfigureUrls**](SecretsApi.md#PkiConfigureUrls) | **Post** /{pki_mount_path}/config/urls | 
[**PkiCrossSignIntermediate**](SecretsApi.md#PkiCrossSignIntermediate) | **Post** /{pki_mount_path}/intermediate/cross-sign | 
[**PkiDeleteIssuer**](SecretsApi.md#PkiDeleteIssuer) | **Delete** /{pki_mount_path}/issuer/{issuer_ref} | 
[**PkiDeleteKey**](SecretsApi.md#PkiDeleteKey) | **Delete** /{pki_mount_path}/key/{key_ref} | 
[**PkiDeleteRole**](SecretsApi.md#PkiDeleteRole) | **Delete** /{pki_mount_path}/roles/{name} | 
[**PkiDeleteRoot**](SecretsApi.md#PkiDeleteRoot) | **Delete** /{pki_mount_path}/root | 
[**PkiGenerateExportedKey**](SecretsApi.md#PkiGenerateExportedKey) | **Post** /{pki_mount_path}/keys/generate/exported | 
[**PkiGenerateIntermediate**](SecretsApi.md#PkiGenerateIntermediate) | **Post** /{pki_mount_path}/intermediate/generate/{exported} | 
[**PkiGenerateInternalKey**](SecretsApi.md#PkiGenerateInternalKey) | **Post** /{pki_mount_path}/keys/generate/internal | 
[**PkiGenerateKmsKey**](SecretsApi.md#PkiGenerateKmsKey) | **Post** /{pki_mount_path}/keys/generate/kms | 
[**PkiGenerateRoot**](SecretsApi.md#PkiGenerateRoot) | **Post** /{pki_mount_path}/root/generate/{exported} | 
[**PkiImportKey**](SecretsApi.md#PkiImportKey) | **Post** /{pki_mount_path}/keys/import | 
[**PkiIssueWithRole**](SecretsApi.md#PkiIssueWithRole) | **Post** /{pki_mount_path}/issue/{role} | 
[**PkiIssuerIssueWithRole**](SecretsApi.md#PkiIssuerIssueWithRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/issue/{role} | 
[**PkiIssuerReadCrl**](SecretsApi.md#PkiIssuerReadCrl) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl | 
[**PkiIssuerReadCrlDelta**](SecretsApi.md#PkiIssuerReadCrlDelta) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl/delta | 
[**PkiIssuerReadCrlDeltaDer**](SecretsApi.md#PkiIssuerReadCrlDeltaDer) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl/delta/der | 
[**PkiIssuerReadCrlDeltaPem**](SecretsApi.md#PkiIssuerReadCrlDeltaPem) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl/delta/pem | 
[**PkiIssuerReadCrlDer**](SecretsApi.md#PkiIssuerReadCrlDer) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl/der | 
[**PkiIssuerReadCrlPem**](SecretsApi.md#PkiIssuerReadCrlPem) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/crl/pem | 
[**PkiIssuerResignCrls**](SecretsApi.md#PkiIssuerResignCrls) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/resign-crls | 
[**PkiIssuerSignIntermediate**](SecretsApi.md#PkiIssuerSignIntermediate) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate | 
[**PkiIssuerSignRevocationList**](SecretsApi.md#PkiIssuerSignRevocationList) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list | 
[**PkiIssuerSignSelfIssued**](SecretsApi.md#PkiIssuerSignSelfIssued) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued | 
[**PkiIssuerSignVerbatim**](SecretsApi.md#PkiIssuerSignVerbatim) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim | 
[**PkiIssuerSignVerbatimWithRole**](SecretsApi.md#PkiIssuerSignVerbatimWithRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**PkiIssuerSignWithRole**](SecretsApi.md#PkiIssuerSignWithRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign/{role} | 
[**PkiIssuersGenerateIntermediate**](SecretsApi.md#PkiIssuersGenerateIntermediate) | **Post** /{pki_mount_path}/issuers/generate/intermediate/{exported} | 
[**PkiIssuersGenerateRoot**](SecretsApi.md#PkiIssuersGenerateRoot) | **Post** /{pki_mount_path}/issuers/generate/root/{exported} | 
[**PkiIssuersImportBundle**](SecretsApi.md#PkiIssuersImportBundle) | **Post** /{pki_mount_path}/issuers/import/bundle | 
[**PkiIssuersImportCert**](SecretsApi.md#PkiIssuersImportCert) | **Post** /{pki_mount_path}/issuers/import/cert | 
[**PkiIssuersRotateRoot**](SecretsApi.md#PkiIssuersRotateRoot) | **Post** /{pki_mount_path}/root/rotate/{exported} | 
[**PkiListCerts**](SecretsApi.md#PkiListCerts) | **Get** /{pki_mount_path}/certs | 
[**PkiListIssuers**](SecretsApi.md#PkiListIssuers) | **Get** /{pki_mount_path}/issuers | 
[**PkiListKeys**](SecretsApi.md#PkiListKeys) | **Get** /{pki_mount_path}/keys | 
[**PkiListRevokedCerts**](SecretsApi.md#PkiListRevokedCerts) | **Get** /{pki_mount_path}/certs/revoked | 
[**PkiListRoles**](SecretsApi.md#PkiListRoles) | **Get** /{pki_mount_path}/roles | 
[**PkiQueryOcsp**](SecretsApi.md#PkiQueryOcsp) | **Post** /{pki_mount_path}/ocsp | 
[**PkiQueryOcspWithGetReq**](SecretsApi.md#PkiQueryOcspWithGetReq) | **Get** /{pki_mount_path}/ocsp/{req} | 
[**PkiReadAutoTidyConfiguration**](SecretsApi.md#PkiReadAutoTidyConfiguration) | **Get** /{pki_mount_path}/config/auto-tidy | 
[**PkiReadCaChainPem**](SecretsApi.md#PkiReadCaChainPem) | **Get** /{pki_mount_path}/ca_chain | 
[**PkiReadCaDer**](SecretsApi.md#PkiReadCaDer) | **Get** /{pki_mount_path}/ca | 
[**PkiReadCaPem**](SecretsApi.md#PkiReadCaPem) | **Get** /{pki_mount_path}/ca/pem | 
[**PkiReadCert**](SecretsApi.md#PkiReadCert) | **Get** /{pki_mount_path}/cert/{serial} | 
[**PkiReadCertCaChain**](SecretsApi.md#PkiReadCertCaChain) | **Get** /{pki_mount_path}/cert/ca_chain | 
[**PkiReadCertCrl**](SecretsApi.md#PkiReadCertCrl) | **Get** /{pki_mount_path}/cert/crl | 
[**PkiReadCertDeltaCrl**](SecretsApi.md#PkiReadCertDeltaCrl) | **Get** /{pki_mount_path}/cert/delta-crl | 
[**PkiReadCertRawDer**](SecretsApi.md#PkiReadCertRawDer) | **Get** /{pki_mount_path}/cert/{serial}/raw | 
[**PkiReadCertRawPem**](SecretsApi.md#PkiReadCertRawPem) | **Get** /{pki_mount_path}/cert/{serial}/raw/pem | 
[**PkiReadClusterConfiguration**](SecretsApi.md#PkiReadClusterConfiguration) | **Get** /{pki_mount_path}/config/cluster | 
[**PkiReadCrlConfiguration**](SecretsApi.md#PkiReadCrlConfiguration) | **Get** /{pki_mount_path}/config/crl | 
[**PkiReadCrlDelta**](SecretsApi.md#PkiReadCrlDelta) | **Get** /{pki_mount_path}/crl/delta | 
[**PkiReadCrlDeltaPem**](SecretsApi.md#PkiReadCrlDeltaPem) | **Get** /{pki_mount_path}/crl/delta/pem | 
[**PkiReadCrlDer**](SecretsApi.md#PkiReadCrlDer) | **Get** /{pki_mount_path}/crl | 
[**PkiReadCrlPem**](SecretsApi.md#PkiReadCrlPem) | **Get** /{pki_mount_path}/crl/pem | 
[**PkiReadIssuer**](SecretsApi.md#PkiReadIssuer) | **Get** /{pki_mount_path}/issuer/{issuer_ref} | 
[**PkiReadIssuerDer**](SecretsApi.md#PkiReadIssuerDer) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/der | 
[**PkiReadIssuerJson**](SecretsApi.md#PkiReadIssuerJson) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/json | 
[**PkiReadIssuerPem**](SecretsApi.md#PkiReadIssuerPem) | **Get** /{pki_mount_path}/issuer/{issuer_ref}/pem | 
[**PkiReadIssuersConfiguration**](SecretsApi.md#PkiReadIssuersConfiguration) | **Get** /{pki_mount_path}/config/issuers | 
[**PkiReadKey**](SecretsApi.md#PkiReadKey) | **Get** /{pki_mount_path}/key/{key_ref} | 
[**PkiReadKeysConfiguration**](SecretsApi.md#PkiReadKeysConfiguration) | **Get** /{pki_mount_path}/config/keys | 
[**PkiReadRole**](SecretsApi.md#PkiReadRole) | **Get** /{pki_mount_path}/roles/{name} | 
[**PkiReadUrlsConfiguration**](SecretsApi.md#PkiReadUrlsConfiguration) | **Get** /{pki_mount_path}/config/urls | 
[**PkiReplaceRoot**](SecretsApi.md#PkiReplaceRoot) | **Post** /{pki_mount_path}/root/replace | 
[**PkiRevoke**](SecretsApi.md#PkiRevoke) | **Post** /{pki_mount_path}/revoke | 
[**PkiRevokeIssuer**](SecretsApi.md#PkiRevokeIssuer) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/revoke | 
[**PkiRevokeWithKey**](SecretsApi.md#PkiRevokeWithKey) | **Post** /{pki_mount_path}/revoke-with-key | 
[**PkiRootSignIntermediate**](SecretsApi.md#PkiRootSignIntermediate) | **Post** /{pki_mount_path}/root/sign-intermediate | 
[**PkiRootSignSelfIssued**](SecretsApi.md#PkiRootSignSelfIssued) | **Post** /{pki_mount_path}/root/sign-self-issued | 
[**PkiRotateCrl**](SecretsApi.md#PkiRotateCrl) | **Get** /{pki_mount_path}/crl/rotate | 
[**PkiRotateDeltaCrl**](SecretsApi.md#PkiRotateDeltaCrl) | **Get** /{pki_mount_path}/crl/rotate-delta | 
[**PkiSetSignedIntermediate**](SecretsApi.md#PkiSetSignedIntermediate) | **Post** /{pki_mount_path}/intermediate/set-signed | 
[**PkiSignVerbatim**](SecretsApi.md#PkiSignVerbatim) | **Post** /{pki_mount_path}/sign-verbatim | 
[**PkiSignVerbatimWithRole**](SecretsApi.md#PkiSignVerbatimWithRole) | **Post** /{pki_mount_path}/sign-verbatim/{role} | 
[**PkiSignWithRole**](SecretsApi.md#PkiSignWithRole) | **Post** /{pki_mount_path}/sign/{role} | 
[**PkiTidy**](SecretsApi.md#PkiTidy) | **Post** /{pki_mount_path}/tidy | 
[**PkiTidyCancel**](SecretsApi.md#PkiTidyCancel) | **Post** /{pki_mount_path}/tidy-cancel | 
[**PkiTidyStatus**](SecretsApi.md#PkiTidyStatus) | **Get** /{pki_mount_path}/tidy-status | 
[**PkiWriteIssuer**](SecretsApi.md#PkiWriteIssuer) | **Post** /{pki_mount_path}/issuer/{issuer_ref} | 
[**PkiWriteKey**](SecretsApi.md#PkiWriteKey) | **Post** /{pki_mount_path}/key/{key_ref} | 
[**PkiWriteRole**](SecretsApi.md#PkiWriteRole) | **Post** /{pki_mount_path}/roles/{name} | 
[**RabbitMqConfigureConnection**](SecretsApi.md#RabbitMqConfigureConnection) | **Post** /{rabbitmq_mount_path}/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**RabbitMqConfigureLease**](SecretsApi.md#RabbitMqConfigureLease) | **Post** /{rabbitmq_mount_path}/config/lease | 
[**RabbitMqDeleteRole**](SecretsApi.md#RabbitMqDeleteRole) | **Delete** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**RabbitMqListRoles**](SecretsApi.md#RabbitMqListRoles) | **Get** /{rabbitmq_mount_path}/roles | Manage the roles that can be created with this backend.
[**RabbitMqReadLeaseConfiguration**](SecretsApi.md#RabbitMqReadLeaseConfiguration) | **Get** /{rabbitmq_mount_path}/config/lease | 
[**RabbitMqReadRole**](SecretsApi.md#RabbitMqReadRole) | **Get** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**RabbitMqRequestCredentials**](SecretsApi.md#RabbitMqRequestCredentials) | **Get** /{rabbitmq_mount_path}/creds/{name} | Request RabbitMQ credentials for a certain role.
[**RabbitMqWriteRole**](SecretsApi.md#RabbitMqWriteRole) | **Post** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**SshConfigureCa**](SecretsApi.md#SshConfigureCa) | **Post** /{ssh_mount_path}/config/ca | 
[**SshConfigureZeroAddress**](SecretsApi.md#SshConfigureZeroAddress) | **Post** /{ssh_mount_path}/config/zeroaddress | 
[**SshDeleteCaConfiguration**](SecretsApi.md#SshDeleteCaConfiguration) | **Delete** /{ssh_mount_path}/config/ca | 
[**SshDeleteRole**](SecretsApi.md#SshDeleteRole) | **Delete** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SshDeleteZeroAddressConfiguration**](SecretsApi.md#SshDeleteZeroAddressConfiguration) | **Delete** /{ssh_mount_path}/config/zeroaddress | 
[**SshGenerateCredentials**](SecretsApi.md#SshGenerateCredentials) | **Post** /{ssh_mount_path}/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
[**SshIssueCertificate**](SecretsApi.md#SshIssueCertificate) | **Post** /{ssh_mount_path}/issue/{role} | 
[**SshListRoles**](SecretsApi.md#SshListRoles) | **Get** /{ssh_mount_path}/roles | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SshListRolesByIp**](SecretsApi.md#SshListRolesByIp) | **Post** /{ssh_mount_path}/lookup | List all the roles associated with the given IP address.
[**SshReadCaConfiguration**](SecretsApi.md#SshReadCaConfiguration) | **Get** /{ssh_mount_path}/config/ca | 
[**SshReadPublicKey**](SecretsApi.md#SshReadPublicKey) | **Get** /{ssh_mount_path}/public_key | Retrieve the public key.
[**SshReadRole**](SecretsApi.md#SshReadRole) | **Get** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SshReadZeroAddressConfiguration**](SecretsApi.md#SshReadZeroAddressConfiguration) | **Get** /{ssh_mount_path}/config/zeroaddress | 
[**SshSignCertificate**](SecretsApi.md#SshSignCertificate) | **Post** /{ssh_mount_path}/sign/{role} | Request signing an SSH key using a certain role with the provided details.
[**SshTidyDynamicHostKeys**](SecretsApi.md#SshTidyDynamicHostKeys) | **Delete** /{ssh_mount_path}/tidy/dynamic-keys | This endpoint removes the stored host keys used for the removed Dynamic Key feature, if present.
[**SshVerifyOtp**](SecretsApi.md#SshVerifyOtp) | **Post** /{ssh_mount_path}/verify | Validate the OTP provided by Vault SSH Agent.
[**SshWriteRole**](SecretsApi.md#SshWriteRole) | **Post** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**TerraformCloudConfigure**](SecretsApi.md#TerraformCloudConfigure) | **Post** /{terraform_mount_path}/config | 
[**TerraformCloudDeleteConfiguration**](SecretsApi.md#TerraformCloudDeleteConfiguration) | **Delete** /{terraform_mount_path}/config | 
[**TerraformCloudDeleteRole**](SecretsApi.md#TerraformCloudDeleteRole) | **Delete** /{terraform_mount_path}/role/{name} | 
[**TerraformCloudGenerateCredentials**](SecretsApi.md#TerraformCloudGenerateCredentials) | **Get** /{terraform_mount_path}/creds/{name} | 
[**TerraformCloudGenerateCredentials2**](SecretsApi.md#TerraformCloudGenerateCredentials2) | **Post** /{terraform_mount_path}/creds/{name} | 
[**TerraformCloudListRoles**](SecretsApi.md#TerraformCloudListRoles) | **Get** /{terraform_mount_path}/role | 
[**TerraformCloudReadConfiguration**](SecretsApi.md#TerraformCloudReadConfiguration) | **Get** /{terraform_mount_path}/config | 
[**TerraformCloudReadRole**](SecretsApi.md#TerraformCloudReadRole) | **Get** /{terraform_mount_path}/role/{name} | 
[**TerraformCloudRotateRole**](SecretsApi.md#TerraformCloudRotateRole) | **Post** /{terraform_mount_path}/rotate-role/{name} | 
[**TerraformCloudWriteRole**](SecretsApi.md#TerraformCloudWriteRole) | **Post** /{terraform_mount_path}/role/{name} | 
[**TotpCreateKey**](SecretsApi.md#TotpCreateKey) | **Post** /{totp_mount_path}/keys/{name} | 
[**TotpDeleteKey**](SecretsApi.md#TotpDeleteKey) | **Delete** /{totp_mount_path}/keys/{name} | 
[**TotpGenerateCode**](SecretsApi.md#TotpGenerateCode) | **Get** /{totp_mount_path}/code/{name} | 
[**TotpListKeys**](SecretsApi.md#TotpListKeys) | **Get** /{totp_mount_path}/keys | Manage the keys that can be created with this backend.
[**TotpReadKey**](SecretsApi.md#TotpReadKey) | **Get** /{totp_mount_path}/keys/{name} | 
[**TotpValidateCode**](SecretsApi.md#TotpValidateCode) | **Post** /{totp_mount_path}/code/{name} | 
[**TransitBackUpKey**](SecretsApi.md#TransitBackUpKey) | **Get** /{transit_mount_path}/backup/{name} | Backup the named key
[**TransitConfigureCache**](SecretsApi.md#TransitConfigureCache) | **Post** /{transit_mount_path}/cache-config | Configures a new cache of the specified size
[**TransitConfigureKey**](SecretsApi.md#TransitConfigureKey) | **Post** /{transit_mount_path}/keys/{name}/config | Configure a named encryption key
[**TransitConfigureKeys**](SecretsApi.md#TransitConfigureKeys) | **Post** /{transit_mount_path}/config/keys | 
[**TransitCreateKey**](SecretsApi.md#TransitCreateKey) | **Post** /{transit_mount_path}/keys/{name} | 
[**TransitDecrypt**](SecretsApi.md#TransitDecrypt) | **Post** /{transit_mount_path}/decrypt/{name} | Decrypt a ciphertext value using a named key
[**TransitDeleteKey**](SecretsApi.md#TransitDeleteKey) | **Delete** /{transit_mount_path}/keys/{name} | 
[**TransitEncrypt**](SecretsApi.md#TransitEncrypt) | **Post** /{transit_mount_path}/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**TransitExportKey**](SecretsApi.md#TransitExportKey) | **Get** /{transit_mount_path}/export/{type}/{name} | Export named encryption or signing key
[**TransitExportKeyVersion**](SecretsApi.md#TransitExportKeyVersion) | **Get** /{transit_mount_path}/export/{type}/{name}/{version} | Export named encryption or signing key
[**TransitGenerateDataKey**](SecretsApi.md#TransitGenerateDataKey) | **Post** /{transit_mount_path}/datakey/{plaintext}/{name} | Generate a data key
[**TransitGenerateHmac**](SecretsApi.md#TransitGenerateHmac) | **Post** /{transit_mount_path}/hmac/{name} | Generate an HMAC for input data using the named key
[**TransitGenerateHmacWithAlgorithm**](SecretsApi.md#TransitGenerateHmacWithAlgorithm) | **Post** /{transit_mount_path}/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**TransitGenerateRandom**](SecretsApi.md#TransitGenerateRandom) | **Post** /{transit_mount_path}/random | Generate random bytes
[**TransitGenerateRandomWithBytes**](SecretsApi.md#TransitGenerateRandomWithBytes) | **Post** /{transit_mount_path}/random/{urlbytes} | Generate random bytes
[**TransitGenerateRandomWithSource**](SecretsApi.md#TransitGenerateRandomWithSource) | **Post** /{transit_mount_path}/random/{source} | Generate random bytes
[**TransitGenerateRandomWithSourceAndBytes**](SecretsApi.md#TransitGenerateRandomWithSourceAndBytes) | **Post** /{transit_mount_path}/random/{source}/{urlbytes} | Generate random bytes
[**TransitHash**](SecretsApi.md#TransitHash) | **Post** /{transit_mount_path}/hash | Generate a hash sum for input data
[**TransitHashWithAlgorithm**](SecretsApi.md#TransitHashWithAlgorithm) | **Post** /{transit_mount_path}/hash/{urlalgorithm} | Generate a hash sum for input data
[**TransitImportKey**](SecretsApi.md#TransitImportKey) | **Post** /{transit_mount_path}/keys/{name}/import | Imports an externally-generated key into a new transit key
[**TransitImportKeyVersion**](SecretsApi.md#TransitImportKeyVersion) | **Post** /{transit_mount_path}/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**TransitListKeys**](SecretsApi.md#TransitListKeys) | **Get** /{transit_mount_path}/keys | Managed named encryption keys
[**TransitReadCacheConfiguration**](SecretsApi.md#TransitReadCacheConfiguration) | **Get** /{transit_mount_path}/cache-config | Returns the size of the active cache
[**TransitReadKey**](SecretsApi.md#TransitReadKey) | **Get** /{transit_mount_path}/keys/{name} | 
[**TransitReadKeysConfiguration**](SecretsApi.md#TransitReadKeysConfiguration) | **Get** /{transit_mount_path}/config/keys | 
[**TransitReadWrappingKey**](SecretsApi.md#TransitReadWrappingKey) | **Get** /{transit_mount_path}/wrapping_key | Returns the public key to use for wrapping imported keys
[**TransitRestoreAndRenameKey**](SecretsApi.md#TransitRestoreAndRenameKey) | **Post** /{transit_mount_path}/restore/{name} | Restore the named key
[**TransitRestoreKey**](SecretsApi.md#TransitRestoreKey) | **Post** /{transit_mount_path}/restore | Restore the named key
[**TransitRewrap**](SecretsApi.md#TransitRewrap) | **Post** /{transit_mount_path}/rewrap/{name} | Rewrap ciphertext
[**TransitRotateKey**](SecretsApi.md#TransitRotateKey) | **Post** /{transit_mount_path}/keys/{name}/rotate | Rotate named encryption key
[**TransitSign**](SecretsApi.md#TransitSign) | **Post** /{transit_mount_path}/sign/{name} | Generate a signature for input data using the named key
[**TransitSignWithAlgorithm**](SecretsApi.md#TransitSignWithAlgorithm) | **Post** /{transit_mount_path}/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**TransitTrimKey**](SecretsApi.md#TransitTrimKey) | **Post** /{transit_mount_path}/keys/{name}/trim | Trim key versions of a named key
[**TransitVerify**](SecretsApi.md#TransitVerify) | **Post** /{transit_mount_path}/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**TransitVerifyWithAlgorithm**](SecretsApi.md#TransitVerifyWithAlgorithm) | **Post** /{transit_mount_path}/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key





## AliCloudConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	request := schema.NewAliCloudConfigureRequestWithDefaults()
	resp, err := client.Secrets.AliCloudConfigure(
		context.Background(),
		alicloudMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliCloudConfigureRequest** | [**AliCloudConfigureRequest**](AliCloudConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudDeleteConfiguration(
		context.Background(),
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudDeleteRole

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudDeleteRole(
		context.Background(),
		name,
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the role. | 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudGenerateCredentials

Generate an API key or STS credential using the given role's configuration.'

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudGenerateCredentials(
		context.Background(),
		name,
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the role. | 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudListRoles

List the existing roles in this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudListRoles(
		context.Background(),
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudReadConfiguration(
		context.Background(),
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadRole

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	resp, err := client.Secrets.AliCloudReadRole(
		context.Background(),
		name,
		alicloudMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the role. | 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudWriteRole

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path that the backend was mounted at (defaults to "alicloud")
	request := schema.NewAliCloudWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.AliCloudWriteRole(
		context.Background(),
		name,
		alicloudMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the role. | 
**alicloudMountPath** | **string** | Path that the backend was mounted at | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aliCloudWriteRoleRequest** | [**AliCloudWriteRoleRequest**](AliCloudWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureLease



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	request := schema.NewAwsConfigureLeaseRequestWithDefaults()
	resp, err := client.Secrets.AwsConfigureLease(
		context.Background(),
		awsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureLeaseRequest** | [**AwsConfigureLeaseRequest**](AwsConfigureLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureRootIamCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	request := schema.NewAwsConfigureRootIamCredentialsRequestWithDefaults()
	resp, err := client.Secrets.AwsConfigureRootIamCredentials(
		context.Background(),
		awsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureRootIamCredentialsRequest** | [**AwsConfigureRootIamCredentialsRequest**](AwsConfigureRootIamCredentialsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteRole

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsDeleteRole(
		context.Background(),
		name,
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the policy | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsGenerateCredentials(
		context.Background(),
		name,
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsGenerateCredentials2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	request := schema.NewAwsGenerateCredentials2RequestWithDefaults()
	resp, err := client.Secrets.AwsGenerateCredentials2(
		context.Background(),
		name,
		awsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsGenerateCredentials2Request** | [**AwsGenerateCredentials2Request**](AwsGenerateCredentials2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsGenerateStsCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsGenerateStsCredentials(
		context.Background(),
		name,
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsGenerateStsCredentials2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	request := schema.NewAwsGenerateStsCredentials2RequestWithDefaults()
	resp, err := client.Secrets.AwsGenerateStsCredentials2(
		context.Background(),
		name,
		awsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsGenerateStsCredentials2Request** | [**AwsGenerateStsCredentials2Request**](AwsGenerateStsCredentials2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListRoles

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsListRoles(
		context.Background(),
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadLeaseConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsReadLeaseConfiguration(
		context.Background(),
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadRole

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsReadRole(
		context.Background(),
		name,
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the policy | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadRootIamCredentialsConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsReadRootIamCredentialsConfiguration(
		context.Background(),
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsRotateRootIamCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	resp, err := client.Secrets.AwsRotateRootIamCredentials(
		context.Background(),
		awsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsWriteRole

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path that the backend was mounted at (defaults to "aws")
	request := schema.NewAwsWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.AwsWriteRole(
		context.Background(),
		name,
		awsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the policy | 
**awsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsWriteRoleRequest** | [**AwsWriteRoleRequest**](AwsWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	request := schema.NewAzureConfigureRequestWithDefaults()
	resp, err := client.Secrets.AzureConfigure(
		context.Background(),
		azureMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureConfigureRequest** | [**AzureConfigureRequest**](AzureConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureDeleteConfiguration(
		context.Background(),
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteRole

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureDeleteRole(
		context.Background(),
		name,
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureListRoles

List existing roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureListRoles(
		context.Background(),
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureReadConfiguration(
		context.Background(),
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadRole

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureReadRole(
		context.Background(),
		name,
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureRequestServicePrincipalCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the Vault role
	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureRequestServicePrincipalCredentials(
		context.Background(),
		role,
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | Name of the Vault role | 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureRotateRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	resp, err := client.Secrets.AzureRotateRoot(
		context.Background(),
		azureMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteRole

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path that the backend was mounted at (defaults to "azure")
	request := schema.NewAzureWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.AzureWriteRole(
		context.Background(),
		name,
		azureMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**azureMountPath** | **string** | Path that the backend was mounted at | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureWriteRoleRequest** | [**AzureWriteRoleRequest**](AzureWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulConfigureAccess



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	request := schema.NewConsulConfigureAccessRequestWithDefaults()
	resp, err := client.Secrets.ConsulConfigureAccess(
		context.Background(),
		consulMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consulConfigureAccessRequest** | [**ConsulConfigureAccessRequest**](ConsulConfigureAccessRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	resp, err := client.Secrets.ConsulDeleteRole(
		context.Background(),
		name,
		consulMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	resp, err := client.Secrets.ConsulGenerateCredentials(
		context.Background(),
		role,
		consulMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | Name of the role. | 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	resp, err := client.Secrets.ConsulListRoles(
		context.Background(),
		consulMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulReadAccessConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	resp, err := client.Secrets.ConsulReadAccessConfiguration(
		context.Background(),
		consulMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	resp, err := client.Secrets.ConsulReadRole(
		context.Background(),
		name,
		consulMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path that the backend was mounted at (defaults to "consul")
	request := schema.NewConsulWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.ConsulWriteRole(
		context.Background(),
		name,
		consulMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**consulMountPath** | **string** | Path that the backend was mounted at | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **consulWriteRoleRequest** | [**ConsulWriteRoleRequest**](ConsulWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CubbyholeDelete

Deletes the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	resp, err := client.Secrets.CubbyholeDelete(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CubbyholeRead

Retrieve the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	resp, err := client.Secrets.CubbyholeRead(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CubbyholeWrite

Store a secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	resp, err := client.Secrets.CubbyholeWrite(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseConfigureConnection



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of this database connection
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	request := schema.NewDatabaseConfigureConnectionRequestWithDefaults()
	resp, err := client.Secrets.DatabaseConfigureConnection(
		context.Background(),
		name,
		databaseMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of this database connection | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseConfigureConnectionRequest** | [**DatabaseConfigureConnectionRequest**](DatabaseConfigureConnectionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseDeleteConnectionConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of this database connection
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseDeleteConnectionConfiguration(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of this database connection | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseDeleteRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseDeleteRole(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseDeleteStaticRole

Manage the static roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseDeleteStaticRole(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseGenerateCredentials

Request database credentials for a certain role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseGenerateCredentials(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseListConnections

Configure connection details to a database plugin.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseListConnections(
		context.Background(),
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseListRoles

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseListRoles(
		context.Background(),
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseListStaticRoles

Manage the static roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseListStaticRoles(
		context.Background(),
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseReadConnectionConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of this database connection
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseReadConnectionConfiguration(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of this database connection | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseReadRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseReadRole(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseReadStaticRole

Manage the static roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseReadStaticRole(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseReadStaticRoleCredentials

Request database credentials for a certain static role. These credentials are rotated periodically.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseReadStaticRoleCredentials(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the static role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseResetConnection

Resets a database plugin.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of this database connection
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseResetConnection(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of this database connection | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseRotateRootCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of this database connection
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseRotateRootCredentials(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of this database connection | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseRotateStaticRoleCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	resp, err := client.Secrets.DatabaseRotateStaticRoleCredentials(
		context.Background(),
		name,
		databaseMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the static role | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseWriteRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	request := schema.NewDatabaseWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.DatabaseWriteRole(
		context.Background(),
		name,
		databaseMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseWriteRoleRequest** | [**DatabaseWriteRoleRequest**](DatabaseWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DatabaseWriteStaticRole

Manage the static roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	databaseMountPath := "databaseMountPath_example" // string | Path that the backend was mounted at (defaults to "database")
	request := schema.NewDatabaseWriteStaticRoleRequestWithDefaults()
	resp, err := client.Secrets.DatabaseWriteStaticRole(
		context.Background(),
		name,
		databaseMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**databaseMountPath** | **string** | Path that the backend was mounted at | [default to &quot;database&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseWriteStaticRoleRequest** | [**DatabaseWriteStaticRoleRequest**](DatabaseWriteStaticRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudConfigureRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudConfigure(
		context.Background(),
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudConfigureRequest** | [**GoogleCloudConfigureRequest**](GoogleCloudConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteImpersonatedAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this impersonated account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudDeleteImpersonatedAccount(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this impersonated account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteRoleset



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudDeleteRoleset(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the role. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteStaticAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudDeleteStaticAccount(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateImpersonatedAccountAccessToken



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the impersonated account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateImpersonatedAccountAccessToken(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the impersonated account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateImpersonatedAccountAccessToken2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the impersonated account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateImpersonatedAccountAccessToken2(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the impersonated account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetAccessToken



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetAccessToken(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetAccessToken2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetAccessToken2(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetAccessTokenWithParameters



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetAccessTokenWithParameters(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetAccessTokenWithParameters2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetAccessTokenWithParameters2(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetKey(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetKey2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateRolesetKey2(
		context.Background(),
		roleset,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetKeyWithParameters



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudGenerateRolesetKeyWithParametersRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudGenerateRolesetKeyWithParameters(
		context.Background(),
		roleset,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudGenerateRolesetKeyWithParametersRequest** | [**GoogleCloudGenerateRolesetKeyWithParametersRequest**](GoogleCloudGenerateRolesetKeyWithParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateRolesetKeyWithParameters2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudGenerateRolesetKeyWithParameters2RequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudGenerateRolesetKeyWithParameters2(
		context.Background(),
		roleset,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleset** | **string** | Required. Name of the role set. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudGenerateRolesetKeyWithParameters2Request** | [**GoogleCloudGenerateRolesetKeyWithParameters2Request**](GoogleCloudGenerateRolesetKeyWithParameters2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateStaticAccountAccessToken



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateStaticAccountAccessToken(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the static account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateStaticAccountAccessTokenWithParameters



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateStaticAccountAccessTokenWithParameters(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the static account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateStaticAccountKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudGenerateStaticAccountKey(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the static account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudGenerateStaticAccountKeyWithParameters



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudGenerateStaticAccountKeyWithParametersRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudGenerateStaticAccountKeyWithParameters(
		context.Background(),
		name,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the static account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudGenerateStaticAccountKeyWithParametersRequest** | [**GoogleCloudGenerateStaticAccountKeyWithParametersRequest**](GoogleCloudGenerateStaticAccountKeyWithParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsConfigureRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsConfigure(
		context.Background(),
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKmsConfigureRequest** | [**GoogleCloudKmsConfigureRequest**](GoogleCloudKmsConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsConfigureKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsConfigureKeyRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsConfigureKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsConfigureKeyRequest** | [**GoogleCloudKmsConfigureKeyRequest**](GoogleCloudKmsConfigureKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsDecrypt

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsDecryptRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsDecrypt(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsDecryptRequest** | [**GoogleCloudKmsDecryptRequest**](GoogleCloudKmsDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsDeleteConfiguration(
		context.Background(),
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsDeleteKey

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsDeleteKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsDeregisterKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsDeregisterKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsDeregisterKey2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsDeregisterKey2(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsEncrypt

Encrypt a plaintext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsEncryptRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsEncrypt(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsEncryptRequest** | [**GoogleCloudKmsEncryptRequest**](GoogleCloudKmsEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsListKeys

List named keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsListKeys(
		context.Background(),
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsReadConfiguration(
		context.Background(),
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsReadKey

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsReadKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsReadKeyConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsReadKeyConfiguration(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsReencrypt

Re-encrypt existing ciphertext data to a new version

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsReencryptRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsReencrypt(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsReencryptRequest** | [**GoogleCloudKmsReencryptRequest**](GoogleCloudKmsReencryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsRegisterKey

Register an existing crypto key in Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsRegisterKeyRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsRegisterKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsRegisterKeyRequest** | [**GoogleCloudKmsRegisterKeyRequest**](GoogleCloudKmsRegisterKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsRetrievePublicKey

Retrieve the public key associated with the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsRetrievePublicKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsRotateKey

Rotate a crypto key to a new primary version

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsRotateKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsSign

Signs a message or digest using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsSignRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsSign(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsSignRequest** | [**GoogleCloudKmsSignRequest**](GoogleCloudKmsSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsTrimKeyVersions



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsTrimKeyVersions(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsTrimKeyVersions2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	resp, err := client.Secrets.GoogleCloudKmsTrimKeyVersions2(
		context.Background(),
		key,
		gcpkmsMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsVerify

Verify a signature using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsVerifyRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsVerify(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsVerifyRequest** | [**GoogleCloudKmsVerifyRequest**](GoogleCloudKmsVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKmsWriteKey

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path that the backend was mounted at (defaults to "gcpkms")
	request := schema.NewGoogleCloudKmsWriteKeyRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudKmsWriteKey(
		context.Background(),
		key,
		gcpkmsMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Name of the key in Vault. | 
**gcpkmsMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudKmsWriteKeyRequest** | [**GoogleCloudKmsWriteKeyRequest**](GoogleCloudKmsWriteKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListImpersonatedAccounts



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListImpersonatedAccounts(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListImpersonatedAccounts2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListImpersonatedAccounts2(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRolesets



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListRolesets(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRolesets2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListRolesets2(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListStaticAccounts



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListStaticAccounts(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListStaticAccounts2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudListStaticAccounts2(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudReadConfiguration(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadImpersonatedAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this impersonated account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudReadImpersonatedAccount(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this impersonated account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRoleset



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudReadRoleset(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the role. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadStaticAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudReadStaticAccount(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRoleset



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudRotateRoleset(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRolesetKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudRotateRolesetKey(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRootCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudRotateRootCredentials(
		context.Background(),
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateStaticAccountKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the account.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	resp, err := client.Secrets.GoogleCloudRotateStaticAccountKey(
		context.Background(),
		name,
		gcpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the account. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteImpersonatedAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this impersonated account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudWriteImpersonatedAccount(
		context.Background(),
		name,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this impersonated account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudWriteImpersonatedAccountRequest** | [**GoogleCloudWriteImpersonatedAccountRequest**](GoogleCloudWriteImpersonatedAccountRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRoleset



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudWriteRolesetRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudWriteRoleset(
		context.Background(),
		name,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name of the role. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudWriteRolesetRequest** | [**GoogleCloudWriteRolesetRequest**](GoogleCloudWriteRolesetRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteStaticAccount



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path that the backend was mounted at (defaults to "gcp")
	request := schema.NewGoogleCloudWriteStaticAccountRequestWithDefaults()
	resp, err := client.Secrets.GoogleCloudWriteStaticAccount(
		context.Background(),
		name,
		gcpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 
**gcpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudWriteStaticAccountRequest** | [**GoogleCloudWriteStaticAccountRequest**](GoogleCloudWriteStaticAccountRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesCheckConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesCheckConfiguration(
		context.Background(),
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	request := schema.NewKubernetesConfigureRequestWithDefaults()
	resp, err := client.Secrets.KubernetesConfigure(
		context.Background(),
		kubernetesMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesConfigureRequest** | [**KubernetesConfigureRequest**](KubernetesConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesDeleteConfiguration(
		context.Background(),
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesDeleteRole(
		context.Background(),
		name,
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Vault role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	request := schema.NewKubernetesGenerateCredentialsRequestWithDefaults()
	resp, err := client.Secrets.KubernetesGenerateCredentials(
		context.Background(),
		name,
		kubernetesMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Vault role | 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kubernetesGenerateCredentialsRequest** | [**KubernetesGenerateCredentialsRequest**](KubernetesGenerateCredentialsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesListRoles(
		context.Background(),
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesReadConfiguration(
		context.Background(),
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	resp, err := client.Secrets.KubernetesReadRole(
		context.Background(),
		name,
		kubernetesMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path that the backend was mounted at (defaults to "kubernetes")
	request := schema.NewKubernetesWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.KubernetesWriteRole(
		context.Background(),
		name,
		kubernetesMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**kubernetesMountPath** | **string** | Path that the backend was mounted at | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kubernetesWriteRoleRequest** | [**KubernetesWriteRoleRequest**](KubernetesWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV1Delete



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV1MountPath := "kvV1MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v1")
	resp, err := client.Secrets.KvV1Delete(
		context.Background(),
		path,
		kvV1MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV1MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v1&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV1Read



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV1MountPath := "kvV1MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v1")
	resp, err := client.Secrets.KvV1Read(
		context.Background(),
		path,
		kvV1MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV1MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v1&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV1Write



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV1MountPath := "kvV1MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v1")
	resp, err := client.Secrets.KvV1Write(
		context.Background(),
		path,
		kvV1MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV1MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v1&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2Configure

Configure backend level settings that are applied to every key in the key-value store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2ConfigureRequestWithDefaults()
	resp, err := client.Secrets.KvV2Configure(
		context.Background(),
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvV2ConfigureRequest** | [**KvV2ConfigureRequest**](KvV2ConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2Delete



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2Delete(
		context.Background(),
		path,
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2DeleteMetadata



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2DeleteMetadata(
		context.Background(),
		path,
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2DeleteVersions



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2DeleteVersionsRequestWithDefaults()
	resp, err := client.Secrets.KvV2DeleteVersions(
		context.Background(),
		path,
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvV2DeleteVersionsRequest** | [**KvV2DeleteVersionsRequest**](KvV2DeleteVersionsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2DestroyVersions



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2DestroyVersionsRequestWithDefaults()
	resp, err := client.Secrets.KvV2DestroyVersions(
		context.Background(),
		path,
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvV2DestroyVersionsRequest** | [**KvV2DestroyVersionsRequest**](KvV2DestroyVersionsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2Read



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2Read(
		context.Background(),
		path,
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**KvV2ReadResponse**](KvV2ReadResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2ReadConfiguration

Read the backend level settings.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2ReadConfiguration(
		context.Background(),
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**KvV2ReadConfigurationResponse**](KvV2ReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2ReadMetadata



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2ReadMetadata(
		context.Background(),
		path,
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

[**KvV2ReadMetadataResponse**](KvV2ReadMetadataResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2ReadSubkeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	resp, err := client.Secrets.KvV2ReadSubkeys(
		context.Background(),
		path,
		kvV2MountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**KvV2ReadSubkeysResponse**](KvV2ReadSubkeysResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2UndeleteVersions



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2UndeleteVersionsRequestWithDefaults()
	resp, err := client.Secrets.KvV2UndeleteVersions(
		context.Background(),
		path,
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvV2UndeleteVersionsRequest** | [**KvV2UndeleteVersionsRequest**](KvV2UndeleteVersionsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2Write



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2WriteRequestWithDefaults()
	resp, err := client.Secrets.KvV2Write(
		context.Background(),
		path,
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvV2WriteRequest** | [**KvV2WriteRequest**](KvV2WriteRequest.md) |  | 

[**KvV2WriteResponse**](KvV2WriteResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## KvV2WriteMetadata



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	kvV2MountPath := "kvV2MountPath_example" // string | Path that the backend was mounted at (defaults to "kv-v2")
	request := schema.NewKvV2WriteMetadataRequestWithDefaults()
	resp, err := client.Secrets.KvV2WriteMetadata(
		context.Background(),
		path,
		kvV2MountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Location of the secret. | 
**kvV2MountPath** | **string** | Path that the backend was mounted at | [default to &quot;kv-v2&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvV2WriteMetadataRequest** | [**KvV2WriteMetadataRequest**](KvV2WriteMetadataRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapConfigureRequestWithDefaults()
	resp, err := client.Secrets.LdapConfigure(
		context.Background(),
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapConfigureRequest** | [**LdapConfigureRequest**](LdapConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapDeleteConfiguration(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapDeleteDynamicRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapDeleteDynamicRole(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role (lowercase) | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapDeleteStaticRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapDeleteStaticRole(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryCheckIn

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapLibraryCheckInRequestWithDefaults()
	resp, err := client.Secrets.LdapLibraryCheckIn(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryCheckInRequest** | [**LdapLibraryCheckInRequest**](LdapLibraryCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryCheckOut

Check a service account out from the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapLibraryCheckOutRequestWithDefaults()
	resp, err := client.Secrets.LdapLibraryCheckOut(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryCheckOutRequest** | [**LdapLibraryCheckOutRequest**](LdapLibraryCheckOutRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryCheckStatus

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapLibraryCheckStatus(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryConfigure

Update a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapLibraryConfigureRequestWithDefaults()
	resp, err := client.Secrets.LdapLibraryConfigure(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryConfigureRequest** | [**LdapLibraryConfigureRequest**](LdapLibraryConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryDelete

Delete a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapLibraryDelete(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryForceCheckIn

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapLibraryForceCheckInRequestWithDefaults()
	resp, err := client.Secrets.LdapLibraryForceCheckIn(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryForceCheckInRequest** | [**LdapLibraryForceCheckInRequest**](LdapLibraryForceCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapLibraryList(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLibraryRead

Read a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapLibraryRead(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the set. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapListDynamicRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapListDynamicRoles(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapListStaticRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapListStaticRoles(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapReadConfiguration(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapReadDynamicRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapReadDynamicRole(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role (lowercase) | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapReadStaticRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapReadStaticRole(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapRequestDynamicRoleCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the dynamic role.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapRequestDynamicRoleCredentials(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the dynamic role. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapRequestStaticRoleCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapRequestStaticRoleCredentials(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the static role. | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapRotateRootCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapRotateRootCredentials(
		context.Background(),
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapRotateStaticRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	resp, err := client.Secrets.LdapRotateStaticRole(
		context.Background(),
		name,
		ldapMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the static role | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapWriteDynamicRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapWriteDynamicRoleRequestWithDefaults()
	resp, err := client.Secrets.LdapWriteDynamicRole(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role (lowercase) | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapWriteDynamicRoleRequest** | [**LdapWriteDynamicRoleRequest**](LdapWriteDynamicRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapWriteStaticRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path that the backend was mounted at (defaults to "ldap")
	request := schema.NewLdapWriteStaticRoleRequestWithDefaults()
	resp, err := client.Secrets.LdapWriteStaticRole(
		context.Background(),
		name,
		ldapMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**ldapMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapWriteStaticRoleRequest** | [**LdapWriteStaticRoleRequest**](LdapWriteStaticRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	request := schema.NewMongoDbAtlasConfigureRequestWithDefaults()
	resp, err := client.Secrets.MongoDbAtlasConfigure(
		context.Background(),
		mongodbatlasMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mongoDbAtlasConfigureRequest** | [**MongoDbAtlasConfigureRequest**](MongoDbAtlasConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasDeleteRole

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasDeleteRole(
		context.Background(),
		name,
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Roles | 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasGenerateCredentials(
		context.Background(),
		name,
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasGenerateCredentials2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasGenerateCredentials2(
		context.Background(),
		name,
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasListRoles

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasListRoles(
		context.Background(),
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasReadConfiguration(
		context.Background(),
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasReadRole

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	resp, err := client.Secrets.MongoDbAtlasReadRole(
		context.Background(),
		name,
		mongodbatlasMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Roles | 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDbAtlasWriteRole

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path that the backend was mounted at (defaults to "mongodbatlas")
	request := schema.NewMongoDbAtlasWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.MongoDbAtlasWriteRole(
		context.Background(),
		name,
		mongodbatlasMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Roles | 
**mongodbatlasMountPath** | **string** | Path that the backend was mounted at | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mongoDbAtlasWriteRoleRequest** | [**MongoDbAtlasWriteRoleRequest**](MongoDbAtlasWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadConfigureAccess



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	request := schema.NewNomadConfigureAccessRequestWithDefaults()
	resp, err := client.Secrets.NomadConfigureAccess(
		context.Background(),
		nomadMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadConfigureAccessRequest** | [**NomadConfigureAccessRequest**](NomadConfigureAccessRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadConfigureLease



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	request := schema.NewNomadConfigureLeaseRequestWithDefaults()
	resp, err := client.Secrets.NomadConfigureLease(
		context.Background(),
		nomadMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadConfigureLeaseRequest** | [**NomadConfigureLeaseRequest**](NomadConfigureLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteAccessConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadDeleteAccessConfiguration(
		context.Background(),
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteLeaseConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadDeleteLeaseConfiguration(
		context.Background(),
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadDeleteRole(
		context.Background(),
		name,
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadGenerateCredentials(
		context.Background(),
		name,
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadListRoles(
		context.Background(),
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadAccessConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadReadAccessConfiguration(
		context.Background(),
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadLeaseConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadReadLeaseConfiguration(
		context.Background(),
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	resp, err := client.Secrets.NomadReadRole(
		context.Background(),
		name,
		nomadMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path that the backend was mounted at (defaults to "nomad")
	request := schema.NewNomadWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.NomadWriteRole(
		context.Background(),
		name,
		nomadMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**nomadMountPath** | **string** | Path that the backend was mounted at | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nomadWriteRoleRequest** | [**NomadWriteRoleRequest**](NomadWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureAutoTidy



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureAutoTidyRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureAutoTidy(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureAutoTidyRequest** | [**PkiConfigureAutoTidyRequest**](PkiConfigureAutoTidyRequest.md) |  | 

[**PkiConfigureAutoTidyResponse**](PkiConfigureAutoTidyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureCa



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureCaRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureCa(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureCaRequest** | [**PkiConfigureCaRequest**](PkiConfigureCaRequest.md) |  | 

[**PkiConfigureCaResponse**](PkiConfigureCaResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureCluster



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureClusterRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureCluster(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureClusterRequest** | [**PkiConfigureClusterRequest**](PkiConfigureClusterRequest.md) |  | 

[**PkiConfigureClusterResponse**](PkiConfigureClusterResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureCrlRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureCrl(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureCrlRequest** | [**PkiConfigureCrlRequest**](PkiConfigureCrlRequest.md) |  | 

[**PkiConfigureCrlResponse**](PkiConfigureCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureIssuers



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureIssuersRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureIssuers(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureIssuersRequest** | [**PkiConfigureIssuersRequest**](PkiConfigureIssuersRequest.md) |  | 

[**PkiConfigureIssuersResponse**](PkiConfigureIssuersResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureKeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureKeysRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureKeys(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureKeysRequest** | [**PkiConfigureKeysRequest**](PkiConfigureKeysRequest.md) |  | 

[**PkiConfigureKeysResponse**](PkiConfigureKeysResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiConfigureUrls



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiConfigureUrlsRequestWithDefaults()
	resp, err := client.Secrets.PkiConfigureUrls(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigureUrlsRequest** | [**PkiConfigureUrlsRequest**](PkiConfigureUrlsRequest.md) |  | 

[**PkiConfigureUrlsResponse**](PkiConfigureUrlsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiCrossSignIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiCrossSignIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiCrossSignIntermediate(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiCrossSignIntermediateRequest** | [**PkiCrossSignIntermediateRequest**](PkiCrossSignIntermediateRequest.md) |  | 

[**PkiCrossSignIntermediateResponse**](PkiCrossSignIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteIssuer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiDeleteIssuer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiDeleteKey(
		context.Background(),
		keyRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiDeleteRole(
		context.Background(),
		name,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiDeleteRoot(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiGenerateExportedKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiGenerateExportedKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiGenerateExportedKey(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiGenerateExportedKeyRequest** | [**PkiGenerateExportedKeyRequest**](PkiGenerateExportedKeyRequest.md) |  | 

[**PkiGenerateExportedKeyResponse**](PkiGenerateExportedKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiGenerateIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiGenerateIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiGenerateIntermediate(
		context.Background(),
		exported,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiGenerateIntermediateRequest** | [**PkiGenerateIntermediateRequest**](PkiGenerateIntermediateRequest.md) |  | 

[**PkiGenerateIntermediateResponse**](PkiGenerateIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiGenerateInternalKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiGenerateInternalKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiGenerateInternalKey(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiGenerateInternalKeyRequest** | [**PkiGenerateInternalKeyRequest**](PkiGenerateInternalKeyRequest.md) |  | 

[**PkiGenerateInternalKeyResponse**](PkiGenerateInternalKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiGenerateKmsKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiGenerateKmsKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiGenerateKmsKey(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiGenerateKmsKeyRequest** | [**PkiGenerateKmsKeyRequest**](PkiGenerateKmsKeyRequest.md) |  | 

[**PkiGenerateKmsKeyResponse**](PkiGenerateKmsKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiGenerateRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiGenerateRootRequestWithDefaults()
	resp, err := client.Secrets.PkiGenerateRoot(
		context.Background(),
		exported,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiGenerateRootRequest** | [**PkiGenerateRootRequest**](PkiGenerateRootRequest.md) |  | 

[**PkiGenerateRootResponse**](PkiGenerateRootResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiImportKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiImportKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiImportKey(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiImportKeyRequest** | [**PkiImportKeyRequest**](PkiImportKeyRequest.md) |  | 

[**PkiImportKeyResponse**](PkiImportKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssueWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssueWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiIssueWithRole(
		context.Background(),
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssueWithRoleRequest** | [**PkiIssueWithRoleRequest**](PkiIssueWithRoleRequest.md) |  | 

[**PkiIssueWithRoleResponse**](PkiIssueWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerIssueWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerIssueWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerIssueWithRole(
		context.Background(),
		issuerRef,
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerIssueWithRoleRequest** | [**PkiIssuerIssueWithRoleRequest**](PkiIssuerIssueWithRoleRequest.md) |  | 

[**PkiIssuerIssueWithRoleResponse**](PkiIssuerIssueWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrl(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlResponse**](PkiIssuerReadCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrlDelta



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrlDelta(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlDeltaResponse**](PkiIssuerReadCrlDeltaResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrlDeltaDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrlDeltaDer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlDeltaDerResponse**](PkiIssuerReadCrlDeltaDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrlDeltaPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrlDeltaPem(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlDeltaPemResponse**](PkiIssuerReadCrlDeltaPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrlDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrlDer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlDerResponse**](PkiIssuerReadCrlDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerReadCrlPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiIssuerReadCrlPem(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiIssuerReadCrlPemResponse**](PkiIssuerReadCrlPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerResignCrls



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerResignCrlsRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerResignCrls(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerResignCrlsRequest** | [**PkiIssuerResignCrlsRequest**](PkiIssuerResignCrlsRequest.md) |  | 

[**PkiIssuerResignCrlsResponse**](PkiIssuerResignCrlsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignIntermediate(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignIntermediateRequest** | [**PkiIssuerSignIntermediateRequest**](PkiIssuerSignIntermediateRequest.md) |  | 

[**PkiIssuerSignIntermediateResponse**](PkiIssuerSignIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignRevocationList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignRevocationListRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignRevocationList(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignRevocationListRequest** | [**PkiIssuerSignRevocationListRequest**](PkiIssuerSignRevocationListRequest.md) |  | 

[**PkiIssuerSignRevocationListResponse**](PkiIssuerSignRevocationListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignSelfIssued



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignSelfIssuedRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignSelfIssued(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignSelfIssuedRequest** | [**PkiIssuerSignSelfIssuedRequest**](PkiIssuerSignSelfIssuedRequest.md) |  | 

[**PkiIssuerSignSelfIssuedResponse**](PkiIssuerSignSelfIssuedResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignVerbatim



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignVerbatimRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignVerbatim(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

[**PkiIssuerSignVerbatimResponse**](PkiIssuerSignVerbatimResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignVerbatimWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignVerbatimWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignVerbatimWithRole(
		context.Background(),
		issuerRef,
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerSignVerbatimWithRoleRequest** | [**PkiIssuerSignVerbatimWithRoleRequest**](PkiIssuerSignVerbatimWithRoleRequest.md) |  | 

[**PkiIssuerSignVerbatimWithRoleResponse**](PkiIssuerSignVerbatimWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuerSignWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuerSignWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuerSignWithRole(
		context.Background(),
		issuerRef,
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerSignWithRoleRequest** | [**PkiIssuerSignWithRoleRequest**](PkiIssuerSignWithRoleRequest.md) |  | 

[**PkiIssuerSignWithRoleResponse**](PkiIssuerSignWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuersGenerateIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuersGenerateIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuersGenerateIntermediate(
		context.Background(),
		exported,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuersGenerateIntermediateRequest** | [**PkiIssuersGenerateIntermediateRequest**](PkiIssuersGenerateIntermediateRequest.md) |  | 

[**PkiIssuersGenerateIntermediateResponse**](PkiIssuersGenerateIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuersGenerateRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuersGenerateRootRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuersGenerateRoot(
		context.Background(),
		exported,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuersGenerateRootRequest** | [**PkiIssuersGenerateRootRequest**](PkiIssuersGenerateRootRequest.md) |  | 

[**PkiIssuersGenerateRootResponse**](PkiIssuersGenerateRootResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuersImportBundle



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuersImportBundleRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuersImportBundle(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersImportBundleRequest** | [**PkiIssuersImportBundleRequest**](PkiIssuersImportBundleRequest.md) |  | 

[**PkiIssuersImportBundleResponse**](PkiIssuersImportBundleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuersImportCert



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuersImportCertRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuersImportCert(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersImportCertRequest** | [**PkiIssuersImportCertRequest**](PkiIssuersImportCertRequest.md) |  | 

[**PkiIssuersImportCertResponse**](PkiIssuersImportCertResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiIssuersRotateRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiIssuersRotateRootRequestWithDefaults()
	resp, err := client.Secrets.PkiIssuersRotateRoot(
		context.Background(),
		exported,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuersRotateRootRequest** | [**PkiIssuersRotateRootRequest**](PkiIssuersRotateRootRequest.md) |  | 

[**PkiIssuersRotateRootResponse**](PkiIssuersRotateRootResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiListCerts



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiListCerts(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PkiListCertsResponse**](PkiListCertsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiListIssuers



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiListIssuers(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PkiListIssuersResponse**](PkiListIssuersResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiListKeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiListKeys(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PkiListKeysResponse**](PkiListKeysResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiListRevokedCerts



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiListRevokedCerts(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PkiListRevokedCertsResponse**](PkiListRevokedCertsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiListRoles(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PkiListRolesResponse**](PkiListRolesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiQueryOcsp



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiQueryOcsp(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiQueryOcspWithGetReq



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	req := "req_example" // string | base-64 encoded ocsp request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiQueryOcspWithGetReq(
		context.Background(),
		req,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**req** | **string** | base-64 encoded ocsp request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadAutoTidyConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadAutoTidyConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadAutoTidyConfigurationResponse**](PkiReadAutoTidyConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCaChainPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCaChainPem(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCaChainPemResponse**](PkiReadCaChainPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCaDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCaDer(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCaDerResponse**](PkiReadCaDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCaPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCaPem(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCaPemResponse**](PkiReadCaPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCert



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCert(
		context.Background(),
		serial,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadCertResponse**](PkiReadCertResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCertCaChain



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCertCaChain(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCertCaChainResponse**](PkiReadCertCaChainResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCertCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCertCrl(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCertCrlResponse**](PkiReadCertCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCertDeltaCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCertDeltaCrl(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCertDeltaCrlResponse**](PkiReadCertDeltaCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCertRawDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCertRawDer(
		context.Background(),
		serial,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadCertRawDerResponse**](PkiReadCertRawDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCertRawPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCertRawPem(
		context.Background(),
		serial,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadCertRawPemResponse**](PkiReadCertRawPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadClusterConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadClusterConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadClusterConfigurationResponse**](PkiReadClusterConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCrlConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCrlConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCrlConfigurationResponse**](PkiReadCrlConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCrlDelta



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCrlDelta(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCrlDeltaResponse**](PkiReadCrlDeltaResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCrlDeltaPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCrlDeltaPem(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCrlDeltaPemResponse**](PkiReadCrlDeltaPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCrlDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCrlDer(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCrlDerResponse**](PkiReadCrlDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadCrlPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadCrlPem(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadCrlPemResponse**](PkiReadCrlPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadIssuer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadIssuerResponse**](PkiReadIssuerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuerDer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadIssuerDer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadIssuerDerResponse**](PkiReadIssuerDerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuerJson



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadIssuerJson(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadIssuerJsonResponse**](PkiReadIssuerJsonResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuerPem



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadIssuerPem(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadIssuerPemResponse**](PkiReadIssuerPemResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuersConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadIssuersConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadIssuersConfigurationResponse**](PkiReadIssuersConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadKey(
		context.Background(),
		keyRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadKeyResponse**](PkiReadKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadKeysConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadKeysConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadKeysConfigurationResponse**](PkiReadKeysConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadRole(
		context.Background(),
		name,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiReadRoleResponse**](PkiReadRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadUrlsConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiReadUrlsConfiguration(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiReadUrlsConfigurationResponse**](PkiReadUrlsConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReplaceRoot



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiReplaceRootRequestWithDefaults()
	resp, err := client.Secrets.PkiReplaceRoot(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiReplaceRootRequest** | [**PkiReplaceRootRequest**](PkiReplaceRootRequest.md) |  | 

[**PkiReplaceRootResponse**](PkiReplaceRootResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRevoke



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiRevokeRequestWithDefaults()
	resp, err := client.Secrets.PkiRevoke(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRevokeRequest** | [**PkiRevokeRequest**](PkiRevokeRequest.md) |  | 

[**PkiRevokeResponse**](PkiRevokeResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRevokeIssuer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiRevokeIssuer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PkiRevokeIssuerResponse**](PkiRevokeIssuerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRevokeWithKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiRevokeWithKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiRevokeWithKey(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRevokeWithKeyRequest** | [**PkiRevokeWithKeyRequest**](PkiRevokeWithKeyRequest.md) |  | 

[**PkiRevokeWithKeyResponse**](PkiRevokeWithKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRootSignIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiRootSignIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiRootSignIntermediate(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootSignIntermediateRequest** | [**PkiRootSignIntermediateRequest**](PkiRootSignIntermediateRequest.md) |  | 

[**PkiRootSignIntermediateResponse**](PkiRootSignIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRootSignSelfIssued



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiRootSignSelfIssuedRequestWithDefaults()
	resp, err := client.Secrets.PkiRootSignSelfIssued(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootSignSelfIssuedRequest** | [**PkiRootSignSelfIssuedRequest**](PkiRootSignSelfIssuedRequest.md) |  | 

[**PkiRootSignSelfIssuedResponse**](PkiRootSignSelfIssuedResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRotateCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiRotateCrl(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiRotateCrlResponse**](PkiRotateCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiRotateDeltaCrl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiRotateDeltaCrl(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiRotateDeltaCrlResponse**](PkiRotateDeltaCrlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiSetSignedIntermediate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiSetSignedIntermediateRequestWithDefaults()
	resp, err := client.Secrets.PkiSetSignedIntermediate(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSetSignedIntermediateRequest** | [**PkiSetSignedIntermediateRequest**](PkiSetSignedIntermediateRequest.md) |  | 

[**PkiSetSignedIntermediateResponse**](PkiSetSignedIntermediateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiSignVerbatim



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiSignVerbatimRequestWithDefaults()
	resp, err := client.Secrets.PkiSignVerbatim(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

[**PkiSignVerbatimResponse**](PkiSignVerbatimResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiSignVerbatimWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiSignVerbatimWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiSignVerbatimWithRole(
		context.Background(),
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiSignVerbatimWithRoleRequest** | [**PkiSignVerbatimWithRoleRequest**](PkiSignVerbatimWithRoleRequest.md) |  | 

[**PkiSignVerbatimWithRoleResponse**](PkiSignVerbatimWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiSignWithRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiSignWithRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiSignWithRole(
		context.Background(),
		role,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | The desired role with configuration for this request | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiSignWithRoleRequest** | [**PkiSignWithRoleRequest**](PkiSignWithRoleRequest.md) |  | 

[**PkiSignWithRoleResponse**](PkiSignWithRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiTidy



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiTidyRequestWithDefaults()
	resp, err := client.Secrets.PkiTidy(
		context.Background(),
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiTidyRequest** | [**PkiTidyRequest**](PkiTidyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiTidyCancel



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiTidyCancel(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiTidyCancelResponse**](PkiTidyCancelResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiTidyStatus



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	resp, err := client.Secrets.PkiTidyStatus(
		context.Background(),
		pkiMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PkiTidyStatusResponse**](PkiTidyStatusResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiWriteIssuer



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiWriteIssuerRequestWithDefaults()
	resp, err := client.Secrets.PkiWriteIssuer(
		context.Background(),
		issuerRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiWriteIssuerRequest** | [**PkiWriteIssuerRequest**](PkiWriteIssuerRequest.md) |  | 

[**PkiWriteIssuerResponse**](PkiWriteIssuerResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiWriteKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiWriteKeyRequestWithDefaults()
	resp, err := client.Secrets.PkiWriteKey(
		context.Background(),
		keyRef,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiWriteKeyRequest** | [**PkiWriteKeyRequest**](PkiWriteKeyRequest.md) |  | 

[**PkiWriteKeyResponse**](PkiWriteKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path that the backend was mounted at (defaults to "pki")
	request := schema.NewPkiWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.PkiWriteRole(
		context.Background(),
		name,
		pkiMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**pkiMountPath** | **string** | Path that the backend was mounted at | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiWriteRoleRequest** | [**PkiWriteRoleRequest**](PkiWriteRoleRequest.md) |  | 

[**PkiWriteRoleResponse**](PkiWriteRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqConfigureConnection

Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	request := schema.NewRabbitMqConfigureConnectionRequestWithDefaults()
	resp, err := client.Secrets.RabbitMqConfigureConnection(
		context.Background(),
		rabbitmqMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitMqConfigureConnectionRequest** | [**RabbitMqConfigureConnectionRequest**](RabbitMqConfigureConnectionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqConfigureLease



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	request := schema.NewRabbitMqConfigureLeaseRequestWithDefaults()
	resp, err := client.Secrets.RabbitMqConfigureLease(
		context.Background(),
		rabbitmqMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitMqConfigureLeaseRequest** | [**RabbitMqConfigureLeaseRequest**](RabbitMqConfigureLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqDeleteRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	resp, err := client.Secrets.RabbitMqDeleteRole(
		context.Background(),
		name,
		rabbitmqMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqListRoles

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	resp, err := client.Secrets.RabbitMqListRoles(
		context.Background(),
		rabbitmqMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqReadLeaseConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	resp, err := client.Secrets.RabbitMqReadLeaseConfiguration(
		context.Background(),
		rabbitmqMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqReadRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	resp, err := client.Secrets.RabbitMqReadRole(
		context.Background(),
		name,
		rabbitmqMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqRequestCredentials

Request RabbitMQ credentials for a certain role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	resp, err := client.Secrets.RabbitMqRequestCredentials(
		context.Background(),
		name,
		rabbitmqMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMqWriteRole

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path that the backend was mounted at (defaults to "rabbitmq")
	request := schema.NewRabbitMqWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.RabbitMqWriteRole(
		context.Background(),
		name,
		rabbitmqMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role. | 
**rabbitmqMountPath** | **string** | Path that the backend was mounted at | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rabbitMqWriteRoleRequest** | [**RabbitMqWriteRoleRequest**](RabbitMqWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshConfigureCa



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshConfigureCaRequestWithDefaults()
	resp, err := client.Secrets.SshConfigureCa(
		context.Background(),
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshConfigureCaRequest** | [**SshConfigureCaRequest**](SshConfigureCaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshConfigureZeroAddress



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshConfigureZeroAddressRequestWithDefaults()
	resp, err := client.Secrets.SshConfigureZeroAddress(
		context.Background(),
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshConfigureZeroAddressRequest** | [**SshConfigureZeroAddressRequest**](SshConfigureZeroAddressRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshDeleteCaConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshDeleteCaConfiguration(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshDeleteRole

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshDeleteRole(
		context.Background(),
		role,
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | [Required for all types] Name of the role being created. | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshDeleteZeroAddressConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshDeleteZeroAddressConfiguration(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshGenerateCredentials

Creates a credential for establishing SSH connection with the remote host.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required] Name of the role
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshGenerateCredentialsRequestWithDefaults()
	resp, err := client.Secrets.SshGenerateCredentials(
		context.Background(),
		role,
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | [Required] Name of the role | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshGenerateCredentialsRequest** | [**SshGenerateCredentialsRequest**](SshGenerateCredentialsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshIssueCertificate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshIssueCertificateRequestWithDefaults()
	resp, err := client.Secrets.SshIssueCertificate(
		context.Background(),
		role,
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | The desired role with configuration for this request. | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshIssueCertificateRequest** | [**SshIssueCertificateRequest**](SshIssueCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshListRoles

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshListRoles(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshListRolesByIp

List all the roles associated with the given IP address.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshListRolesByIpRequestWithDefaults()
	resp, err := client.Secrets.SshListRolesByIp(
		context.Background(),
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshListRolesByIpRequest** | [**SshListRolesByIpRequest**](SshListRolesByIpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshReadCaConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshReadCaConfiguration(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshReadPublicKey

Retrieve the public key.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshReadPublicKey(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshReadRole

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshReadRole(
		context.Background(),
		role,
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | [Required for all types] Name of the role being created. | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshReadZeroAddressConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshReadZeroAddressConfiguration(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshSignCertificate

Request signing an SSH key using a certain role with the provided details.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshSignCertificateRequestWithDefaults()
	resp, err := client.Secrets.SshSignCertificate(
		context.Background(),
		role,
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | The desired role with configuration for this request. | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshSignCertificateRequest** | [**SshSignCertificateRequest**](SshSignCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshTidyDynamicHostKeys

This endpoint removes the stored host keys used for the removed Dynamic Key feature, if present.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	resp, err := client.Secrets.SshTidyDynamicHostKeys(
		context.Background(),
		sshMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshVerifyOtp

Validate the OTP provided by Vault SSH Agent.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshVerifyOtpRequestWithDefaults()
	resp, err := client.Secrets.SshVerifyOtp(
		context.Background(),
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshVerifyOtpRequest** | [**SshVerifyOtpRequest**](SshVerifyOtpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SshWriteRole

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path that the backend was mounted at (defaults to "ssh")
	request := schema.NewSshWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.SshWriteRole(
		context.Background(),
		role,
		sshMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**role** | **string** | [Required for all types] Name of the role being created. | 
**sshMountPath** | **string** | Path that the backend was mounted at | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshWriteRoleRequest** | [**SshWriteRoleRequest**](SshWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	request := schema.NewTerraformCloudConfigureRequestWithDefaults()
	resp, err := client.Secrets.TerraformCloudConfigure(
		context.Background(),
		terraformMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformCloudConfigureRequest** | [**TerraformCloudConfigureRequest**](TerraformCloudConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudDeleteConfiguration(
		context.Background(),
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudDeleteRole(
		context.Background(),
		name,
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudGenerateCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudGenerateCredentials(
		context.Background(),
		name,
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudGenerateCredentials2



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudGenerateCredentials2(
		context.Background(),
		name,
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudListRoles(
		context.Background(),
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudReadConfiguration(
		context.Background(),
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudReadRole(
		context.Background(),
		name,
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudRotateRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the team or organization role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	resp, err := client.Secrets.TerraformCloudRotateRole(
		context.Background(),
		name,
		terraformMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the team or organization role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformCloudWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path that the backend was mounted at (defaults to "terraform")
	request := schema.NewTerraformCloudWriteRoleRequestWithDefaults()
	resp, err := client.Secrets.TerraformCloudWriteRole(
		context.Background(),
		name,
		terraformMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 
**terraformMountPath** | **string** | Path that the backend was mounted at | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **terraformCloudWriteRoleRequest** | [**TerraformCloudWriteRoleRequest**](TerraformCloudWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpCreateKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	request := schema.NewTotpCreateKeyRequestWithDefaults()
	resp, err := client.Secrets.TotpCreateKey(
		context.Background(),
		name,
		totpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key. | 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **totpCreateKeyRequest** | [**TotpCreateKeyRequest**](TotpCreateKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpDeleteKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	resp, err := client.Secrets.TotpDeleteKey(
		context.Background(),
		name,
		totpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key. | 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpGenerateCode



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	resp, err := client.Secrets.TotpGenerateCode(
		context.Background(),
		name,
		totpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key. | 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpListKeys

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	resp, err := client.Secrets.TotpListKeys(
		context.Background(),
		totpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpReadKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	resp, err := client.Secrets.TotpReadKey(
		context.Background(),
		name,
		totpMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key. | 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TotpValidateCode



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path that the backend was mounted at (defaults to "totp")
	request := schema.NewTotpValidateCodeRequestWithDefaults()
	resp, err := client.Secrets.TotpValidateCode(
		context.Background(),
		name,
		totpMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key. | 
**totpMountPath** | **string** | Path that the backend was mounted at | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **totpValidateCodeRequest** | [**TotpValidateCodeRequest**](TotpValidateCodeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitBackUpKey

Backup the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitBackUpKey(
		context.Background(),
		name,
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitConfigureCache

Configures a new cache of the specified size

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitConfigureCacheRequestWithDefaults()
	resp, err := client.Secrets.TransitConfigureCache(
		context.Background(),
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitConfigureCacheRequest** | [**TransitConfigureCacheRequest**](TransitConfigureCacheRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitConfigureKey

Configure a named encryption key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitConfigureKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitConfigureKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitConfigureKeyRequest** | [**TransitConfigureKeyRequest**](TransitConfigureKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitConfigureKeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitConfigureKeysRequestWithDefaults()
	resp, err := client.Secrets.TransitConfigureKeys(
		context.Background(),
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitConfigureKeysRequest** | [**TransitConfigureKeysRequest**](TransitConfigureKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitCreateKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitCreateKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitCreateKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitCreateKeyRequest** | [**TransitCreateKeyRequest**](TransitCreateKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitDecrypt

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitDecryptRequestWithDefaults()
	resp, err := client.Secrets.TransitDecrypt(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitDeleteKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitDeleteKey(
		context.Background(),
		name,
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitEncrypt

Encrypt a plaintext value or a batch of plaintext blocks using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitEncryptRequestWithDefaults()
	resp, err := client.Secrets.TransitEncrypt(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitExportKey

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitExportKey(
		context.Background(),
		name,
		type_,
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitExportKeyVersion

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
	version := "version_example" // string | Version of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitExportKeyVersion(
		context.Background(),
		name,
		type_,
		version,
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 
**version** | **string** | Version of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateDataKey

Generate a data key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The backend key used for encrypting the data key
	plaintext := "plaintext_example" // string | \"plaintext\" will return the key in both plaintext and ciphertext; \"wrapped\" will return the ciphertext only.
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateDataKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateDataKey(
		context.Background(),
		name,
		plaintext,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The backend key used for encrypting the data key | 
**plaintext** | **string** | \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only. | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitGenerateDataKeyRequest** | [**TransitGenerateDataKeyRequest**](TransitGenerateDataKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateHmac

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateHmacRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateHmac(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use for the HMAC function | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateHmacRequest** | [**TransitGenerateHmacRequest**](TransitGenerateHmacRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateHmacWithAlgorithm

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function
	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateHmacWithAlgorithmRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateHmacWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use for the HMAC function | 
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitGenerateHmacWithAlgorithmRequest** | [**TransitGenerateHmacWithAlgorithmRequest**](TransitGenerateHmacWithAlgorithmRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandom

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateRandomRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateRandom(
		context.Background(),
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGenerateRandomRequest** | [**TransitGenerateRandomRequest**](TransitGenerateRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandomWithBytes

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateRandomWithBytesRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateRandomWithBytes(
		context.Background(),
		urlbytes,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateRandomWithBytesRequest** | [**TransitGenerateRandomWithBytesRequest**](TransitGenerateRandomWithBytesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandomWithSource

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateRandomWithSourceRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateRandomWithSource(
		context.Background(),
		source,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateRandomWithSourceRequest** | [**TransitGenerateRandomWithSourceRequest**](TransitGenerateRandomWithSourceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandomWithSourceAndBytes

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitGenerateRandomWithSourceAndBytesRequestWithDefaults()
	resp, err := client.Secrets.TransitGenerateRandomWithSourceAndBytes(
		context.Background(),
		source,
		urlbytes,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitGenerateRandomWithSourceAndBytesRequest** | [**TransitGenerateRandomWithSourceAndBytesRequest**](TransitGenerateRandomWithSourceAndBytesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitHash

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitHashRequestWithDefaults()
	resp, err := client.Secrets.TransitHash(
		context.Background(),
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitHashWithAlgorithm

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitHashWithAlgorithmRequestWithDefaults()
	resp, err := client.Secrets.TransitHashWithAlgorithm(
		context.Background(),
		urlalgorithm,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitHashWithAlgorithmRequest** | [**TransitHashWithAlgorithmRequest**](TransitHashWithAlgorithmRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitImportKey

Imports an externally-generated key into a new transit key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitImportKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitImportKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitImportKeyRequest** | [**TransitImportKeyRequest**](TransitImportKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitImportKeyVersion

Imports an externally-generated key into an existing imported key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitImportKeyVersionRequestWithDefaults()
	resp, err := client.Secrets.TransitImportKeyVersion(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitImportKeyVersionRequest** | [**TransitImportKeyVersionRequest**](TransitImportKeyVersionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitListKeys

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitListKeys(
		context.Background(),
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadCacheConfiguration

Returns the size of the active cache

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitReadCacheConfiguration(
		context.Background(),
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadKey



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitReadKey(
		context.Background(),
		name,
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadKeysConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitReadKeysConfiguration(
		context.Background(),
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadWrappingKey

Returns the public key to use for wrapping imported keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	resp, err := client.Secrets.TransitReadWrappingKey(
		context.Background(),
		transitMountPath,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRestoreAndRenameKey

Restore the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | If set, this will be the name of the restored key.
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitRestoreAndRenameKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitRestoreAndRenameKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | If set, this will be the name of the restored key. | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRestoreAndRenameKeyRequest** | [**TransitRestoreAndRenameKeyRequest**](TransitRestoreAndRenameKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRestoreKey

Restore the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitRestoreKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitRestoreKey(
		context.Background(),
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRestoreKeyRequest** | [**TransitRestoreKeyRequest**](TransitRestoreKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRewrap

Rewrap ciphertext

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitRewrapRequestWithDefaults()
	resp, err := client.Secrets.TransitRewrap(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRewrapRequest** | [**TransitRewrapRequest**](TransitRewrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRotateKey

Rotate named encryption key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitRotateKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitRotateKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRotateKeyRequest** | [**TransitRotateKeyRequest**](TransitRotateKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitSign

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitSignRequestWithDefaults()
	resp, err := client.Secrets.TransitSign(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitSignWithAlgorithm

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitSignWithAlgorithmRequestWithDefaults()
	resp, err := client.Secrets.TransitSignWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitSignWithAlgorithmRequest** | [**TransitSignWithAlgorithmRequest**](TransitSignWithAlgorithmRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitTrimKey

Trim key versions of a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitTrimKeyRequestWithDefaults()
	resp, err := client.Secrets.TransitTrimKey(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitTrimKeyRequest** | [**TransitTrimKeyRequest**](TransitTrimKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitVerify

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitVerifyRequestWithDefaults()
	resp, err := client.Secrets.TransitVerify(
		context.Background(),
		name,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitVerifyWithAlgorithm

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path that the backend was mounted at (defaults to "transit")
	request := schema.NewTransitVerifyWithAlgorithmRequestWithDefaults()
	resp, err := client.Secrets.TransitVerifyWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
		transitMountPath,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 
**transitMountPath** | **string** | Path that the backend was mounted at | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitVerifyWithAlgorithmRequest** | [**TransitVerifyWithAlgorithmRequest**](TransitVerifyWithAlgorithmRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



