# Secrets

Method | HTTP request | Description
------------- | ------------- | -------------
[**AWSConfigReadLease**](Secrets.md#AWSConfigReadLease) | **Get** /{aws_mount_path}/config/lease | Configure the default lease information for generated credentials.
[**AWSConfigReadRootIAMCredentials**](Secrets.md#AWSConfigReadRootIAMCredentials) | **Get** /{aws_mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**AWSConfigRotateRootIAMCredentials**](Secrets.md#AWSConfigRotateRootIAMCredentials) | **Post** /{aws_mount_path}/config/rotate-root | 
[**AWSConfigWriteLease**](Secrets.md#AWSConfigWriteLease) | **Post** /{aws_mount_path}/config/lease | Configure the default lease information for generated credentials.
[**AWSConfigWriteRootIAMCredentials**](Secrets.md#AWSConfigWriteRootIAMCredentials) | **Post** /{aws_mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**AWSDeleteRole**](Secrets.md#AWSDeleteRole) | **Delete** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AWSListRoles**](Secrets.md#AWSListRoles) | **Get** /{aws_mount_path}/roles | List the existing roles in this backend
[**AWSReadCredentials**](Secrets.md#AWSReadCredentials) | **Get** /{aws_mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**AWSReadRole**](Secrets.md#AWSReadRole) | **Get** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AWSReadSecurityTokenService**](Secrets.md#AWSReadSecurityTokenService) | **Get** /{aws_mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**AWSWriteCredentials**](Secrets.md#AWSWriteCredentials) | **Post** /{aws_mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**AWSWriteRole**](Secrets.md#AWSWriteRole) | **Post** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**AWSWriteSecurityTokenService**](Secrets.md#AWSWriteSecurityTokenService) | **Post** /{aws_mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**ActiveDirectoryCheckInLibrary**](Secrets.md#ActiveDirectoryCheckInLibrary) | **Post** /{ad_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**ActiveDirectoryCheckInManageLibrary**](Secrets.md#ActiveDirectoryCheckInManageLibrary) | **Post** /{ad_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**ActiveDirectoryCheckOutLibrary**](Secrets.md#ActiveDirectoryCheckOutLibrary) | **Post** /{ad_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**ActiveDirectoryDeleteConfig**](Secrets.md#ActiveDirectoryDeleteConfig) | **Delete** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**ActiveDirectoryDeleteLibrary**](Secrets.md#ActiveDirectoryDeleteLibrary) | **Delete** /{ad_mount_path}/library/{name} | Delete a library set.
[**ActiveDirectoryDeleteRole**](Secrets.md#ActiveDirectoryDeleteRole) | **Delete** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**ActiveDirectoryListLibraries**](Secrets.md#ActiveDirectoryListLibraries) | **Get** /{ad_mount_path}/library | 
[**ActiveDirectoryListRoles**](Secrets.md#ActiveDirectoryListRoles) | **Get** /{ad_mount_path}/roles | List the name of each role currently stored.
[**ActiveDirectoryReadConfig**](Secrets.md#ActiveDirectoryReadConfig) | **Get** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**ActiveDirectoryReadCredentials**](Secrets.md#ActiveDirectoryReadCredentials) | **Get** /{ad_mount_path}/creds/{name} | 
[**ActiveDirectoryReadLibrary**](Secrets.md#ActiveDirectoryReadLibrary) | **Get** /{ad_mount_path}/library/{name} | Read a library set.
[**ActiveDirectoryReadLibraryStatus**](Secrets.md#ActiveDirectoryReadLibraryStatus) | **Get** /{ad_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**ActiveDirectoryReadRole**](Secrets.md#ActiveDirectoryReadRole) | **Get** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**ActiveDirectoryRotateRole**](Secrets.md#ActiveDirectoryRotateRole) | **Post** /{ad_mount_path}/rotate-role/{name} | 
[**ActiveDirectoryRotateRoot**](Secrets.md#ActiveDirectoryRotateRoot) | **Post** /{ad_mount_path}/rotate-root | 
[**ActiveDirectoryWriteConfig**](Secrets.md#ActiveDirectoryWriteConfig) | **Post** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**ActiveDirectoryWriteLibrary**](Secrets.md#ActiveDirectoryWriteLibrary) | **Post** /{ad_mount_path}/library/{name} | Update a library set.
[**ActiveDirectoryWriteRole**](Secrets.md#ActiveDirectoryWriteRole) | **Post** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**AliCloudDeleteConfig**](Secrets.md#AliCloudDeleteConfig) | **Delete** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**AliCloudDeleteRole**](Secrets.md#AliCloudDeleteRole) | **Delete** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AliCloudListRoles**](Secrets.md#AliCloudListRoles) | **Get** /{alicloud_mount_path}/role | List the existing roles in this backend.
[**AliCloudReadConfig**](Secrets.md#AliCloudReadConfig) | **Get** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**AliCloudReadCredentials**](Secrets.md#AliCloudReadCredentials) | **Get** /{alicloud_mount_path}/creds/{name} | Generate an API key or STS credential using the given role&#x27;s configuration.&#x27;
[**AliCloudReadRole**](Secrets.md#AliCloudReadRole) | **Get** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AliCloudWriteConfig**](Secrets.md#AliCloudWriteConfig) | **Post** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**AliCloudWriteRole**](Secrets.md#AliCloudWriteRole) | **Post** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**AzureDeleteConfig**](Secrets.md#AzureDeleteConfig) | **Delete** /{azure_mount_path}/config | 
[**AzureDeleteRole**](Secrets.md#AzureDeleteRole) | **Delete** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**AzureListRoles**](Secrets.md#AzureListRoles) | **Get** /{azure_mount_path}/roles | List existing roles.
[**AzureReadConfig**](Secrets.md#AzureReadConfig) | **Get** /{azure_mount_path}/config | 
[**AzureReadCredentials**](Secrets.md#AzureReadCredentials) | **Get** /{azure_mount_path}/creds/{role} | 
[**AzureReadRole**](Secrets.md#AzureReadRole) | **Get** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**AzureRotateRoot**](Secrets.md#AzureRotateRoot) | **Post** /{azure_mount_path}/rotate-root | 
[**AzureWriteConfig**](Secrets.md#AzureWriteConfig) | **Post** /{azure_mount_path}/config | 
[**AzureWriteRole**](Secrets.md#AzureWriteRole) | **Post** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**ConsulDeleteRole**](Secrets.md#ConsulDeleteRole) | **Delete** /{consul_mount_path}/roles/{name} | 
[**ConsulListRoles**](Secrets.md#ConsulListRoles) | **Get** /{consul_mount_path}/roles | 
[**ConsulReadAccessConfig**](Secrets.md#ConsulReadAccessConfig) | **Get** /{consul_mount_path}/config/access | 
[**ConsulReadCredentials**](Secrets.md#ConsulReadCredentials) | **Get** /{consul_mount_path}/creds/{role} | 
[**ConsulReadRole**](Secrets.md#ConsulReadRole) | **Get** /{consul_mount_path}/roles/{name} | 
[**ConsulWriteAccessConfig**](Secrets.md#ConsulWriteAccessConfig) | **Post** /{consul_mount_path}/config/access | 
[**ConsulWriteRole**](Secrets.md#ConsulWriteRole) | **Post** /{consul_mount_path}/roles/{name} | 
[**CubbyholeDelete**](Secrets.md#CubbyholeDelete) | **Delete** /{cubbyhole_mount_path}/{path} | Deletes the secret at the specified location.
[**CubbyholeRead**](Secrets.md#CubbyholeRead) | **Get** /{cubbyhole_mount_path}/{path} | Retrieve the secret at the specified location.
[**CubbyholeWrite**](Secrets.md#CubbyholeWrite) | **Post** /{cubbyhole_mount_path}/{path} | Store a secret at the specified location.
[**GoogleCloudDeleteRoleset**](Secrets.md#GoogleCloudDeleteRoleset) | **Delete** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudDeleteStaticAccount**](Secrets.md#GoogleCloudDeleteStaticAccount) | **Delete** /{gcp_mount_path}/static-account/{name} | 
[**GoogleCloudKMSDecrypt**](Secrets.md#GoogleCloudKMSDecrypt) | **Post** /{gcpkms_mount_path}/decrypt/{key} | Decrypt a ciphertext value using a named key
[**GoogleCloudKMSDeleteConfig**](Secrets.md#GoogleCloudKMSDeleteConfig) | **Delete** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**GoogleCloudKMSDeleteKey**](Secrets.md#GoogleCloudKMSDeleteKey) | **Delete** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudKMSDeregisterKey**](Secrets.md#GoogleCloudKMSDeregisterKey) | **Post** /{gcpkms_mount_path}/keys/deregister/{key} | Deregister an existing key in Vault
[**GoogleCloudKMSEncrypt**](Secrets.md#GoogleCloudKMSEncrypt) | **Post** /{gcpkms_mount_path}/encrypt/{key} | Encrypt a plaintext value using a named key
[**GoogleCloudKMSListKeys**](Secrets.md#GoogleCloudKMSListKeys) | **Get** /{gcpkms_mount_path}/keys | List named keys
[**GoogleCloudKMSReadConfig**](Secrets.md#GoogleCloudKMSReadConfig) | **Get** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**GoogleCloudKMSReadKey**](Secrets.md#GoogleCloudKMSReadKey) | **Get** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudKMSReadKeyConfig**](Secrets.md#GoogleCloudKMSReadKeyConfig) | **Get** /{gcpkms_mount_path}/keys/config/{key} | Configure the key in Vault
[**GoogleCloudKMSReadPubkey**](Secrets.md#GoogleCloudKMSReadPubkey) | **Get** /{gcpkms_mount_path}/pubkey/{key} | Retrieve the public key associated with the named key
[**GoogleCloudKMSReencrypt**](Secrets.md#GoogleCloudKMSReencrypt) | **Post** /{gcpkms_mount_path}/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**GoogleCloudKMSRegisterKey**](Secrets.md#GoogleCloudKMSRegisterKey) | **Post** /{gcpkms_mount_path}/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**GoogleCloudKMSRotateKey**](Secrets.md#GoogleCloudKMSRotateKey) | **Post** /{gcpkms_mount_path}/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**GoogleCloudKMSSign**](Secrets.md#GoogleCloudKMSSign) | **Post** /{gcpkms_mount_path}/sign/{key} | Signs a message or digest using a named key
[**GoogleCloudKMSTrimKey**](Secrets.md#GoogleCloudKMSTrimKey) | **Post** /{gcpkms_mount_path}/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**GoogleCloudKMSVerify**](Secrets.md#GoogleCloudKMSVerify) | **Post** /{gcpkms_mount_path}/verify/{key} | Verify a signature using a named key
[**GoogleCloudKMSWriteConfig**](Secrets.md#GoogleCloudKMSWriteConfig) | **Post** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**GoogleCloudKMSWriteKey**](Secrets.md#GoogleCloudKMSWriteKey) | **Post** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GoogleCloudKMSWriteKeyConfig**](Secrets.md#GoogleCloudKMSWriteKeyConfig) | **Post** /{gcpkms_mount_path}/keys/config/{key} | Configure the key in Vault
[**GoogleCloudListRolesets**](Secrets.md#GoogleCloudListRolesets) | **Get** /{gcp_mount_path}/rolesets | 
[**GoogleCloudListStaticAccounts**](Secrets.md#GoogleCloudListStaticAccounts) | **Get** /{gcp_mount_path}/static-accounts | 
[**GoogleCloudReadConfig**](Secrets.md#GoogleCloudReadConfig) | **Get** /{gcp_mount_path}/config | 
[**GoogleCloudReadKey**](Secrets.md#GoogleCloudReadKey) | **Get** /{gcp_mount_path}/key/{roleset} | 
[**GoogleCloudReadRoleset**](Secrets.md#GoogleCloudReadRoleset) | **Get** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudReadRolesetKey**](Secrets.md#GoogleCloudReadRolesetKey) | **Get** /{gcp_mount_path}/roleset/{roleset}/key | 
[**GoogleCloudReadRolesetToken**](Secrets.md#GoogleCloudReadRolesetToken) | **Get** /{gcp_mount_path}/roleset/{roleset}/token | 
[**GoogleCloudReadStaticAccount**](Secrets.md#GoogleCloudReadStaticAccount) | **Get** /{gcp_mount_path}/static-account/{name} | 
[**GoogleCloudReadStaticAccountKey**](Secrets.md#GoogleCloudReadStaticAccountKey) | **Get** /{gcp_mount_path}/static-account/{name}/key | 
[**GoogleCloudReadStaticAccountToken**](Secrets.md#GoogleCloudReadStaticAccountToken) | **Get** /{gcp_mount_path}/static-account/{name}/token | 
[**GoogleCloudReadToken**](Secrets.md#GoogleCloudReadToken) | **Get** /{gcp_mount_path}/token/{roleset} | 
[**GoogleCloudRotateRoleset**](Secrets.md#GoogleCloudRotateRoleset) | **Post** /{gcp_mount_path}/roleset/{name}/rotate | 
[**GoogleCloudRotateRolesetKey**](Secrets.md#GoogleCloudRotateRolesetKey) | **Post** /{gcp_mount_path}/roleset/{name}/rotate-key | 
[**GoogleCloudRotateRoot**](Secrets.md#GoogleCloudRotateRoot) | **Post** /{gcp_mount_path}/config/rotate-root | 
[**GoogleCloudRotateStaticAccountKey**](Secrets.md#GoogleCloudRotateStaticAccountKey) | **Post** /{gcp_mount_path}/static-account/{name}/rotate-key | 
[**GoogleCloudWriteConfig**](Secrets.md#GoogleCloudWriteConfig) | **Post** /{gcp_mount_path}/config | 
[**GoogleCloudWriteKey**](Secrets.md#GoogleCloudWriteKey) | **Post** /{gcp_mount_path}/key/{roleset} | 
[**GoogleCloudWriteRoleset**](Secrets.md#GoogleCloudWriteRoleset) | **Post** /{gcp_mount_path}/roleset/{name} | 
[**GoogleCloudWriteRolesetKey**](Secrets.md#GoogleCloudWriteRolesetKey) | **Post** /{gcp_mount_path}/roleset/{roleset}/key | 
[**GoogleCloudWriteRolesetToken**](Secrets.md#GoogleCloudWriteRolesetToken) | **Post** /{gcp_mount_path}/roleset/{roleset}/token | 
[**GoogleCloudWriteStaticAccount**](Secrets.md#GoogleCloudWriteStaticAccount) | **Post** /{gcp_mount_path}/static-account/{name} | 
[**GoogleCloudWriteStaticAccountKey**](Secrets.md#GoogleCloudWriteStaticAccountKey) | **Post** /{gcp_mount_path}/static-account/{name}/key | 
[**GoogleCloudWriteStaticAccountToken**](Secrets.md#GoogleCloudWriteStaticAccountToken) | **Post** /{gcp_mount_path}/static-account/{name}/token | 
[**GoogleCloudWriteToken**](Secrets.md#GoogleCloudWriteToken) | **Post** /{gcp_mount_path}/token/{roleset} | 
[**KVv1Delete**](Secrets.md#KVv1Delete) | **Delete** /{kv_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**KVv1Read**](Secrets.md#KVv1Read) | **Get** /{kv_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**KVv1Write**](Secrets.md#KVv1Write) | **Post** /{kv_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**KVv2Delete**](Secrets.md#KVv2Delete) | **Delete** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**KVv2DeleteMetadata**](Secrets.md#KVv2DeleteMetadata) | **Delete** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**KVv2DeleteVersions**](Secrets.md#KVv2DeleteVersions) | **Post** /{secret_mount_path}/delete/{path} | Marks one or more versions as deleted in the KV store.
[**KVv2DestroyVersions**](Secrets.md#KVv2DestroyVersions) | **Post** /{secret_mount_path}/destroy/{path} | Permanently removes one or more versions in the KV store
[**KVv2Read**](Secrets.md#KVv2Read) | **Get** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**KVv2ReadConfig**](Secrets.md#KVv2ReadConfig) | **Get** /{secret_mount_path}/config | Read the backend level settings.
[**KVv2ReadMetadata**](Secrets.md#KVv2ReadMetadata) | **Get** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**KVv2ReadSubkeys**](Secrets.md#KVv2ReadSubkeys) | **Get** /{secret_mount_path}/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
[**KVv2UndeleteVersions**](Secrets.md#KVv2UndeleteVersions) | **Post** /{secret_mount_path}/undelete/{path} | Undeletes one or more versions from the KV store.
[**KVv2Write**](Secrets.md#KVv2Write) | **Post** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**KVv2WriteConfig**](Secrets.md#KVv2WriteConfig) | **Post** /{secret_mount_path}/config | Configure backend level settings that are applied to every key in the key-value store.
[**KVv2WriteMetadata**](Secrets.md#KVv2WriteMetadata) | **Post** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**KubernetesDeleteConfig**](Secrets.md#KubernetesDeleteConfig) | **Delete** /{kubernetes_mount_path}/config | 
[**KubernetesDeleteRole**](Secrets.md#KubernetesDeleteRole) | **Delete** /{kubernetes_mount_path}/roles/{name} | 
[**KubernetesListRoles**](Secrets.md#KubernetesListRoles) | **Get** /{kubernetes_mount_path}/roles | 
[**KubernetesReadConfig**](Secrets.md#KubernetesReadConfig) | **Get** /{kubernetes_mount_path}/config | 
[**KubernetesReadRole**](Secrets.md#KubernetesReadRole) | **Get** /{kubernetes_mount_path}/roles/{name} | 
[**KubernetesWriteConfig**](Secrets.md#KubernetesWriteConfig) | **Post** /{kubernetes_mount_path}/config | 
[**KubernetesWriteCredentials**](Secrets.md#KubernetesWriteCredentials) | **Post** /{kubernetes_mount_path}/creds/{name} | 
[**KubernetesWriteRole**](Secrets.md#KubernetesWriteRole) | **Post** /{kubernetes_mount_path}/roles/{name} | 
[**LDAPCheckInLibrary**](Secrets.md#LDAPCheckInLibrary) | **Post** /{ldap_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**LDAPCheckInManageLibrary**](Secrets.md#LDAPCheckInManageLibrary) | **Post** /{ldap_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**LDAPCheckOutLibrary**](Secrets.md#LDAPCheckOutLibrary) | **Post** /{ldap_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**LDAPDeleteConfig**](Secrets.md#LDAPDeleteConfig) | **Delete** /{ldap_mount_path}/config | 
[**LDAPDeleteLibrary**](Secrets.md#LDAPDeleteLibrary) | **Delete** /{ldap_mount_path}/library/{name} | Delete a library set.
[**LDAPDeleteRole**](Secrets.md#LDAPDeleteRole) | **Delete** /{ldap_mount_path}/role/{name} | 
[**LDAPDeleteStaticRole**](Secrets.md#LDAPDeleteStaticRole) | **Delete** /{ldap_mount_path}/static-role/{name} | 
[**LDAPListLibraries**](Secrets.md#LDAPListLibraries) | **Get** /{ldap_mount_path}/library | 
[**LDAPListRoles**](Secrets.md#LDAPListRoles) | **Get** /{ldap_mount_path}/role | 
[**LDAPListStaticRoles**](Secrets.md#LDAPListStaticRoles) | **Get** /{ldap_mount_path}/static-role | 
[**LDAPReadConfig**](Secrets.md#LDAPReadConfig) | **Get** /{ldap_mount_path}/config | 
[**LDAPReadCredentials**](Secrets.md#LDAPReadCredentials) | **Get** /{ldap_mount_path}/creds/{name} | 
[**LDAPReadLibrary**](Secrets.md#LDAPReadLibrary) | **Get** /{ldap_mount_path}/library/{name} | Read a library set.
[**LDAPReadLibraryStatus**](Secrets.md#LDAPReadLibraryStatus) | **Get** /{ldap_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**LDAPReadRole**](Secrets.md#LDAPReadRole) | **Get** /{ldap_mount_path}/role/{name} | 
[**LDAPReadStaticCredentials**](Secrets.md#LDAPReadStaticCredentials) | **Get** /{ldap_mount_path}/static-cred/{name} | 
[**LDAPReadStaticRole**](Secrets.md#LDAPReadStaticRole) | **Get** /{ldap_mount_path}/static-role/{name} | 
[**LDAPRotateRole**](Secrets.md#LDAPRotateRole) | **Post** /{ldap_mount_path}/rotate-role/{name} | 
[**LDAPRotateRoot**](Secrets.md#LDAPRotateRoot) | **Post** /{ldap_mount_path}/rotate-root | 
[**LDAPWriteConfig**](Secrets.md#LDAPWriteConfig) | **Post** /{ldap_mount_path}/config | 
[**LDAPWriteLibrary**](Secrets.md#LDAPWriteLibrary) | **Post** /{ldap_mount_path}/library/{name} | Update a library set.
[**LDAPWriteRole**](Secrets.md#LDAPWriteRole) | **Post** /{ldap_mount_path}/role/{name} | 
[**LDAPWriteStaticRole**](Secrets.md#LDAPWriteStaticRole) | **Post** /{ldap_mount_path}/static-role/{name} | 
[**MongoDBAtlasDeleteRole**](Secrets.md#MongoDBAtlasDeleteRole) | **Delete** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**MongoDBAtlasListRoles**](Secrets.md#MongoDBAtlasListRoles) | **Get** /{mongodbatlas_mount_path}/roles | List the existing roles in this backend
[**MongoDBAtlasReadConfig**](Secrets.md#MongoDBAtlasReadConfig) | **Get** /{mongodbatlas_mount_path}/config | Configure the  credentials that are used to manage Database Users.
[**MongoDBAtlasReadCredentials**](Secrets.md#MongoDBAtlasReadCredentials) | **Get** /{mongodbatlas_mount_path}/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**MongoDBAtlasReadRole**](Secrets.md#MongoDBAtlasReadRole) | **Get** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**MongoDBAtlasWriteConfig**](Secrets.md#MongoDBAtlasWriteConfig) | **Post** /{mongodbatlas_mount_path}/config | Configure the  credentials that are used to manage Database Users.
[**MongoDBAtlasWriteCredentials**](Secrets.md#MongoDBAtlasWriteCredentials) | **Post** /{mongodbatlas_mount_path}/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**MongoDBAtlasWriteRole**](Secrets.md#MongoDBAtlasWriteRole) | **Post** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**NomadDeleteAccessConfig**](Secrets.md#NomadDeleteAccessConfig) | **Delete** /{nomad_mount_path}/config/access | 
[**NomadDeleteLeaseConfig**](Secrets.md#NomadDeleteLeaseConfig) | **Delete** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**NomadDeleteRole**](Secrets.md#NomadDeleteRole) | **Delete** /{nomad_mount_path}/role/{name} | 
[**NomadListRoles**](Secrets.md#NomadListRoles) | **Get** /{nomad_mount_path}/role | 
[**NomadReadAccessConfig**](Secrets.md#NomadReadAccessConfig) | **Get** /{nomad_mount_path}/config/access | 
[**NomadReadCredentials**](Secrets.md#NomadReadCredentials) | **Get** /{nomad_mount_path}/creds/{name} | 
[**NomadReadLeaseConfig**](Secrets.md#NomadReadLeaseConfig) | **Get** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**NomadReadRole**](Secrets.md#NomadReadRole) | **Get** /{nomad_mount_path}/role/{name} | 
[**NomadWriteAccessConfig**](Secrets.md#NomadWriteAccessConfig) | **Post** /{nomad_mount_path}/config/access | 
[**NomadWriteLeaseConfig**](Secrets.md#NomadWriteLeaseConfig) | **Post** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**NomadWriteRole**](Secrets.md#NomadWriteRole) | **Post** /{nomad_mount_path}/role/{name} | 
[**OpenLDAPCheckInLibrary**](Secrets.md#OpenLDAPCheckInLibrary) | **Post** /{openldap_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**OpenLDAPCheckInManageLibrary**](Secrets.md#OpenLDAPCheckInManageLibrary) | **Post** /{openldap_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**OpenLDAPCheckOutLibrary**](Secrets.md#OpenLDAPCheckOutLibrary) | **Post** /{openldap_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**OpenLDAPDeleteConfig**](Secrets.md#OpenLDAPDeleteConfig) | **Delete** /{openldap_mount_path}/config | 
[**OpenLDAPDeleteLibrary**](Secrets.md#OpenLDAPDeleteLibrary) | **Delete** /{openldap_mount_path}/library/{name} | Delete a library set.
[**OpenLDAPDeleteRole**](Secrets.md#OpenLDAPDeleteRole) | **Delete** /{openldap_mount_path}/role/{name} | 
[**OpenLDAPDeleteStaticRole**](Secrets.md#OpenLDAPDeleteStaticRole) | **Delete** /{openldap_mount_path}/static-role/{name} | 
[**OpenLDAPListLibraries**](Secrets.md#OpenLDAPListLibraries) | **Get** /{openldap_mount_path}/library | 
[**OpenLDAPListRoles**](Secrets.md#OpenLDAPListRoles) | **Get** /{openldap_mount_path}/role | 
[**OpenLDAPListStaticRoles**](Secrets.md#OpenLDAPListStaticRoles) | **Get** /{openldap_mount_path}/static-role | 
[**OpenLDAPReadConfig**](Secrets.md#OpenLDAPReadConfig) | **Get** /{openldap_mount_path}/config | 
[**OpenLDAPReadCredentials**](Secrets.md#OpenLDAPReadCredentials) | **Get** /{openldap_mount_path}/creds/{name} | 
[**OpenLDAPReadLibrary**](Secrets.md#OpenLDAPReadLibrary) | **Get** /{openldap_mount_path}/library/{name} | Read a library set.
[**OpenLDAPReadLibraryStatus**](Secrets.md#OpenLDAPReadLibraryStatus) | **Get** /{openldap_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**OpenLDAPReadRole**](Secrets.md#OpenLDAPReadRole) | **Get** /{openldap_mount_path}/role/{name} | 
[**OpenLDAPReadStaticCredentials**](Secrets.md#OpenLDAPReadStaticCredentials) | **Get** /{openldap_mount_path}/static-cred/{name} | 
[**OpenLDAPReadStaticRole**](Secrets.md#OpenLDAPReadStaticRole) | **Get** /{openldap_mount_path}/static-role/{name} | 
[**OpenLDAPRotateRole**](Secrets.md#OpenLDAPRotateRole) | **Post** /{openldap_mount_path}/rotate-role/{name} | 
[**OpenLDAPRotateRoot**](Secrets.md#OpenLDAPRotateRoot) | **Post** /{openldap_mount_path}/rotate-root | 
[**OpenLDAPWriteConfig**](Secrets.md#OpenLDAPWriteConfig) | **Post** /{openldap_mount_path}/config | 
[**OpenLDAPWriteLibrary**](Secrets.md#OpenLDAPWriteLibrary) | **Post** /{openldap_mount_path}/library/{name} | Update a library set.
[**OpenLDAPWriteRole**](Secrets.md#OpenLDAPWriteRole) | **Post** /{openldap_mount_path}/role/{name} | 
[**OpenLDAPWriteStaticRole**](Secrets.md#OpenLDAPWriteStaticRole) | **Post** /{openldap_mount_path}/static-role/{name} | 
[**PKIBundleWrite**](Secrets.md#PKIBundleWrite) | **Post** /{pki_mount_path}/bundle | 
[**PKIDeleteKey**](Secrets.md#PKIDeleteKey) | **Delete** /{pki_mount_path}/key/{key_ref} | 
[**PKIDeleteRole**](Secrets.md#PKIDeleteRole) | **Delete** /{pki_mount_path}/roles/{name} | 
[**PKIDeleteRoot**](Secrets.md#PKIDeleteRoot) | **Delete** /{pki_mount_path}/root | 
[**PKIGenerateRoot**](Secrets.md#PKIGenerateRoot) | **Post** /{pki_mount_path}/root/generate/{exported} | 
[**PKIImportKeys**](Secrets.md#PKIImportKeys) | **Post** /{pki_mount_path}/keys/import | 
[**PKIIssuerIssueRole**](Secrets.md#PKIIssuerIssueRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/issue/{role} | 
[**PKIIssuerResignCRLs**](Secrets.md#PKIIssuerResignCRLs) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/resign-crls | 
[**PKIIssuerRevoke**](Secrets.md#PKIIssuerRevoke) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/revoke | 
[**PKIIssuerSignIntermediate**](Secrets.md#PKIIssuerSignIntermediate) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate | 
[**PKIIssuerSignRevocationList**](Secrets.md#PKIIssuerSignRevocationList) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list | 
[**PKIIssuerSignRole**](Secrets.md#PKIIssuerSignRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign/{role} | 
[**PKIIssuerSignSelfIssued**](Secrets.md#PKIIssuerSignSelfIssued) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued | 
[**PKIIssuerSignVerbatim**](Secrets.md#PKIIssuerSignVerbatim) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim | 
[**PKIIssuerSignVerbatimRole**](Secrets.md#PKIIssuerSignVerbatimRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**PKIIssuersGenerateIntermediate**](Secrets.md#PKIIssuersGenerateIntermediate) | **Post** /{pki_mount_path}/issuers/generate/intermediate/{exported} | 
[**PKIIssuersGenerateRoot**](Secrets.md#PKIIssuersGenerateRoot) | **Post** /{pki_mount_path}/issuers/generate/root/{exported} | 
[**PKIIssuersList**](Secrets.md#PKIIssuersList) | **Get** /{pki_mount_path}/issuers | 
[**PKIListCerts**](Secrets.md#PKIListCerts) | **Get** /{pki_mount_path}/certs | 
[**PKIListCertsRevoked**](Secrets.md#PKIListCertsRevoked) | **Get** /{pki_mount_path}/certs/revoked | 
[**PKIListKeys**](Secrets.md#PKIListKeys) | **Get** /{pki_mount_path}/keys | 
[**PKIListRoles**](Secrets.md#PKIListRoles) | **Get** /{pki_mount_path}/roles | 
[**PKIReadAutoTidyConfig**](Secrets.md#PKIReadAutoTidyConfig) | **Get** /{pki_mount_path}/config/auto-tidy | 
[**PKIReadCA**](Secrets.md#PKIReadCA) | **Get** /{pki_mount_path}/ca | 
[**PKIReadCAChain**](Secrets.md#PKIReadCAChain) | **Get** /{pki_mount_path}/ca_chain | 
[**PKIReadCAPem**](Secrets.md#PKIReadCAPem) | **Get** /{pki_mount_path}/ca/pem | 
[**PKIReadCRL**](Secrets.md#PKIReadCRL) | **Get** /{pki_mount_path}/crl | 
[**PKIReadCRLConfig**](Secrets.md#PKIReadCRLConfig) | **Get** /{pki_mount_path}/config/crl | 
[**PKIReadCRLRotate**](Secrets.md#PKIReadCRLRotate) | **Get** /{pki_mount_path}/crl/rotate | 
[**PKIReadCRLRotateDelta**](Secrets.md#PKIReadCRLRotateDelta) | **Get** /{pki_mount_path}/crl/rotate-delta | 
[**PKIReadCert**](Secrets.md#PKIReadCert) | **Get** /{pki_mount_path}/cert/{serial} | 
[**PKIReadCertCAChain**](Secrets.md#PKIReadCertCAChain) | **Get** /{pki_mount_path}/cert/ca_chain | 
[**PKIReadCertRaw**](Secrets.md#PKIReadCertRaw) | **Get** /{pki_mount_path}/cert/{serial}/raw | 
[**PKIReadCertRawPem**](Secrets.md#PKIReadCertRawPem) | **Get** /{pki_mount_path}/cert/{serial}/raw/pem | 
[**PKIReadClusterConfig**](Secrets.md#PKIReadClusterConfig) | **Get** /{pki_mount_path}/config/cluster | 
[**PKIReadDeltaCRL**](Secrets.md#PKIReadDeltaCRL) | **Get** /{pki_mount_path}/delta-crl | 
[**PKIReadIssuersConfig**](Secrets.md#PKIReadIssuersConfig) | **Get** /{pki_mount_path}/config/issuers | 
[**PKIReadKey**](Secrets.md#PKIReadKey) | **Get** /{pki_mount_path}/key/{key_ref} | 
[**PKIReadKeysConfig**](Secrets.md#PKIReadKeysConfig) | **Get** /{pki_mount_path}/config/keys | 
[**PKIReadOCSPReq**](Secrets.md#PKIReadOCSPReq) | **Get** /{pki_mount_path}/ocsp/{req} | 
[**PKIReadRole**](Secrets.md#PKIReadRole) | **Get** /{pki_mount_path}/roles/{name} | 
[**PKIReadURLConfig**](Secrets.md#PKIReadURLConfig) | **Get** /{pki_mount_path}/config/urls | 
[**PKIReplaceRoot**](Secrets.md#PKIReplaceRoot) | **Post** /{pki_mount_path}/root/replace | 
[**PKIRevoke**](Secrets.md#PKIRevoke) | **Post** /{pki_mount_path}/revoke | 
[**PKIRevokeWithKey**](Secrets.md#PKIRevokeWithKey) | **Post** /{pki_mount_path}/revoke-with-key | 
[**PKIRootSignIntermediate**](Secrets.md#PKIRootSignIntermediate) | **Post** /{pki_mount_path}/root/sign-intermediate | 
[**PKIRootSignSelfIssued**](Secrets.md#PKIRootSignSelfIssued) | **Post** /{pki_mount_path}/root/sign-self-issued | 
[**PKIRotateRoot**](Secrets.md#PKIRotateRoot) | **Post** /{pki_mount_path}/root/rotate/{exported} | 
[**PKISignRole**](Secrets.md#PKISignRole) | **Post** /{pki_mount_path}/sign/{role} | 
[**PKISignVerbatim**](Secrets.md#PKISignVerbatim) | **Post** /{pki_mount_path}/sign-verbatim | 
[**PKISignVerbatimRole**](Secrets.md#PKISignVerbatimRole) | **Post** /{pki_mount_path}/sign-verbatim/{role} | 
[**PKITidy**](Secrets.md#PKITidy) | **Post** /{pki_mount_path}/tidy | 
[**PKITidyCancel**](Secrets.md#PKITidyCancel) | **Post** /{pki_mount_path}/tidy-cancel | 
[**PKITidyStatus**](Secrets.md#PKITidyStatus) | **Get** /{pki_mount_path}/tidy-status | 
[**PKIWriteAutoTidyConfig**](Secrets.md#PKIWriteAutoTidyConfig) | **Post** /{pki_mount_path}/config/auto-tidy | 
[**PKIWriteCAConfig**](Secrets.md#PKIWriteCAConfig) | **Post** /{pki_mount_path}/config/ca | 
[**PKIWriteCRLConfig**](Secrets.md#PKIWriteCRLConfig) | **Post** /{pki_mount_path}/config/crl | 
[**PKIWriteCerts**](Secrets.md#PKIWriteCerts) | **Post** /{pki_mount_path}/cert | 
[**PKIWriteClusterConfig**](Secrets.md#PKIWriteClusterConfig) | **Post** /{pki_mount_path}/config/cluster | 
[**PKIWriteIntermediateCrossSign**](Secrets.md#PKIWriteIntermediateCrossSign) | **Post** /{pki_mount_path}/intermediate/cross-sign | 
[**PKIWriteIntermediateGenerate**](Secrets.md#PKIWriteIntermediateGenerate) | **Post** /{pki_mount_path}/intermediate/generate/{exported} | 
[**PKIWriteIntermediateSetSigned**](Secrets.md#PKIWriteIntermediateSetSigned) | **Post** /{pki_mount_path}/intermediate/set-signed | 
[**PKIWriteInternalExported**](Secrets.md#PKIWriteInternalExported) | **Post** /{pki_mount_path}/internal|exported | 
[**PKIWriteIssueRole**](Secrets.md#PKIWriteIssueRole) | **Post** /{pki_mount_path}/issue/{role} | 
[**PKIWriteIssuersConfig**](Secrets.md#PKIWriteIssuersConfig) | **Post** /{pki_mount_path}/config/issuers | 
[**PKIWriteKMS**](Secrets.md#PKIWriteKMS) | **Post** /{pki_mount_path}/kms | 
[**PKIWriteKey**](Secrets.md#PKIWriteKey) | **Post** /{pki_mount_path}/key/{key_ref} | 
[**PKIWriteKeysConfig**](Secrets.md#PKIWriteKeysConfig) | **Post** /{pki_mount_path}/config/keys | 
[**PKIWriteOCSP**](Secrets.md#PKIWriteOCSP) | **Post** /{pki_mount_path}/ocsp | 
[**PKIWriteRole**](Secrets.md#PKIWriteRole) | **Post** /{pki_mount_path}/roles/{name} | 
[**PKIWriteURLConfig**](Secrets.md#PKIWriteURLConfig) | **Post** /{pki_mount_path}/config/urls | 
[**PkiDeleteIssuerRefDerPem**](Secrets.md#PkiDeleteIssuerRefDerPem) | **Delete** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**PkiDeleteJson**](Secrets.md#PkiDeleteJson) | **Delete** /{pki_mount_path}//json | 
[**PkiReadDelta**](Secrets.md#PkiReadDelta) | **Get** /{pki_mount_path}//delta | 
[**PkiReadDeltaPem**](Secrets.md#PkiReadDeltaPem) | **Get** /{pki_mount_path}//delta/pem | 
[**PkiReadDer**](Secrets.md#PkiReadDer) | **Get** /{pki_mount_path}//der | 
[**PkiReadIssuerRefCrlPemDerDeltaPem**](Secrets.md#PkiReadIssuerRefCrlPemDerDeltaPem) | **Get** /{pki_mount_path}/{issuer_ref}/crl/pem|/der|/delta/pem | 
[**PkiReadIssuerRefDerPem**](Secrets.md#PkiReadIssuerRefDerPem) | **Get** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**PkiReadJson**](Secrets.md#PkiReadJson) | **Get** /{pki_mount_path}//json | 
[**PkiReadPem**](Secrets.md#PkiReadPem) | **Get** /{pki_mount_path}//pem | 
[**PkiWriteIssuerRefDerPem**](Secrets.md#PkiWriteIssuerRefDerPem) | **Post** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**PkiWriteJson**](Secrets.md#PkiWriteJson) | **Post** /{pki_mount_path}//json | 
[**RabbitMQDeleteRole**](Secrets.md#RabbitMQDeleteRole) | **Delete** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**RabbitMQListRoles**](Secrets.md#RabbitMQListRoles) | **Get** /{rabbitmq_mount_path}/roles | Manage the roles that can be created with this backend.
[**RabbitMQReadCredentials**](Secrets.md#RabbitMQReadCredentials) | **Get** /{rabbitmq_mount_path}/creds/{name} | Request RabbitMQ credentials for a certain role.
[**RabbitMQReadLeaseConfig**](Secrets.md#RabbitMQReadLeaseConfig) | **Get** /{rabbitmq_mount_path}/config/lease | Configure the lease parameters for generated credentials
[**RabbitMQReadRole**](Secrets.md#RabbitMQReadRole) | **Get** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**RabbitMQWriteConnectionConfig**](Secrets.md#RabbitMQWriteConnectionConfig) | **Post** /{rabbitmq_mount_path}/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**RabbitMQWriteLeaseConfig**](Secrets.md#RabbitMQWriteLeaseConfig) | **Post** /{rabbitmq_mount_path}/config/lease | Configure the lease parameters for generated credentials
[**RabbitMQWriteRole**](Secrets.md#RabbitMQWriteRole) | **Post** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**SSHDeleteCAConfig**](Secrets.md#SSHDeleteCAConfig) | **Delete** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**SSHDeleteKeys**](Secrets.md#SSHDeleteKeys) | **Delete** /{ssh_mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**SSHDeleteRole**](Secrets.md#SSHDeleteRole) | **Delete** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SSHDeleteZeroAddressConfig**](Secrets.md#SSHDeleteZeroAddressConfig) | **Delete** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**SSHListRoles**](Secrets.md#SSHListRoles) | **Get** /{ssh_mount_path}/roles | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SSHLookup**](Secrets.md#SSHLookup) | **Post** /{ssh_mount_path}/lookup | List all the roles associated with the given IP address.
[**SSHReadCAConfig**](Secrets.md#SSHReadCAConfig) | **Get** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**SSHReadPublicKey**](Secrets.md#SSHReadPublicKey) | **Get** /{ssh_mount_path}/public_key | Retrieve the public key.
[**SSHReadRole**](Secrets.md#SSHReadRole) | **Get** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SSHReadZeroAddressConfig**](Secrets.md#SSHReadZeroAddressConfig) | **Get** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**SSHSign**](Secrets.md#SSHSign) | **Post** /{ssh_mount_path}/sign/{role} | Request signing an SSH key using a certain role with the provided details.
[**SSHVerify**](Secrets.md#SSHVerify) | **Post** /{ssh_mount_path}/verify | Validate the OTP provided by Vault SSH Agent.
[**SSHWriteCAConfig**](Secrets.md#SSHWriteCAConfig) | **Post** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**SSHWriteCredentials**](Secrets.md#SSHWriteCredentials) | **Post** /{ssh_mount_path}/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
[**SSHWriteIssue**](Secrets.md#SSHWriteIssue) | **Post** /{ssh_mount_path}/issue/{role} | 
[**SSHWriteKeys**](Secrets.md#SSHWriteKeys) | **Post** /{ssh_mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**SSHWriteRole**](Secrets.md#SSHWriteRole) | **Post** /{ssh_mount_path}/roles/{role} | Manage the &#x27;roles&#x27; that can be created with this backend.
[**SSHWriteZeroAddressConfig**](Secrets.md#SSHWriteZeroAddressConfig) | **Post** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**TOTPDeleteKey**](Secrets.md#TOTPDeleteKey) | **Delete** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**TOTPListKeys**](Secrets.md#TOTPListKeys) | **Get** /{totp_mount_path}/keys | Manage the keys that can be created with this backend.
[**TOTPReadCode**](Secrets.md#TOTPReadCode) | **Get** /{totp_mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**TOTPReadKey**](Secrets.md#TOTPReadKey) | **Get** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**TOTPWriteCode**](Secrets.md#TOTPWriteCode) | **Post** /{totp_mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**TOTPWriteKey**](Secrets.md#TOTPWriteKey) | **Post** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**TerraformDeleteConfig**](Secrets.md#TerraformDeleteConfig) | **Delete** /{terraform_mount_path}/config | 
[**TerraformDeleteRole**](Secrets.md#TerraformDeleteRole) | **Delete** /{terraform_mount_path}/role/{name} | 
[**TerraformListRoles**](Secrets.md#TerraformListRoles) | **Get** /{terraform_mount_path}/role | 
[**TerraformReadConfig**](Secrets.md#TerraformReadConfig) | **Get** /{terraform_mount_path}/config | 
[**TerraformReadCredentials**](Secrets.md#TerraformReadCredentials) | **Get** /{terraform_mount_path}/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**TerraformReadRole**](Secrets.md#TerraformReadRole) | **Get** /{terraform_mount_path}/role/{name} | 
[**TerraformRotateRole**](Secrets.md#TerraformRotateRole) | **Post** /{terraform_mount_path}/rotate-role/{name} | 
[**TerraformWriteConfig**](Secrets.md#TerraformWriteConfig) | **Post** /{terraform_mount_path}/config | 
[**TerraformWriteCredentials**](Secrets.md#TerraformWriteCredentials) | **Post** /{terraform_mount_path}/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**TerraformWriteRole**](Secrets.md#TerraformWriteRole) | **Post** /{terraform_mount_path}/role/{name} | 
[**TransitBackup**](Secrets.md#TransitBackup) | **Get** /{transit_mount_path}/backup/{name} | Backup the named key
[**TransitDecrypt**](Secrets.md#TransitDecrypt) | **Post** /{transit_mount_path}/decrypt/{name} | Decrypt a ciphertext value using a named key
[**TransitDeleteKey**](Secrets.md#TransitDeleteKey) | **Delete** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**TransitEncrypt**](Secrets.md#TransitEncrypt) | **Post** /{transit_mount_path}/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**TransitExport**](Secrets.md#TransitExport) | **Get** /{transit_mount_path}/export/{type}/{name} | Export named encryption or signing key
[**TransitExportVersion**](Secrets.md#TransitExportVersion) | **Get** /{transit_mount_path}/export/{type}/{name}/{version} | Export named encryption or signing key
[**TransitGenerateDataKey**](Secrets.md#TransitGenerateDataKey) | **Post** /{transit_mount_path}/datakey/{plaintext}/{name} | Generate a data key
[**TransitGenerateHMAC**](Secrets.md#TransitGenerateHMAC) | **Post** /{transit_mount_path}/hmac/{name} | Generate an HMAC for input data using the named key
[**TransitGenerateHMACWithAlgorithm**](Secrets.md#TransitGenerateHMACWithAlgorithm) | **Post** /{transit_mount_path}/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**TransitGenerateRandom**](Secrets.md#TransitGenerateRandom) | **Post** /{transit_mount_path}/random | Generate random bytes
[**TransitGenerateRandomSource**](Secrets.md#TransitGenerateRandomSource) | **Post** /{transit_mount_path}/random/{source} | Generate random bytes
[**TransitGenerateRandomSourceBytes**](Secrets.md#TransitGenerateRandomSourceBytes) | **Post** /{transit_mount_path}/random/{source}/{urlbytes} | Generate random bytes
[**TransitHash**](Secrets.md#TransitHash) | **Post** /{transit_mount_path}/hash | Generate a hash sum for input data
[**TransitHashWithAlgorithm**](Secrets.md#TransitHashWithAlgorithm) | **Post** /{transit_mount_path}/hash/{urlalgorithm} | Generate a hash sum for input data
[**TransitImportKey**](Secrets.md#TransitImportKey) | **Post** /{transit_mount_path}/keys/{name}/import | Imports an externally-generated key into a new transit key
[**TransitImportKeyVersion**](Secrets.md#TransitImportKeyVersion) | **Post** /{transit_mount_path}/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**TransitListKeys**](Secrets.md#TransitListKeys) | **Get** /{transit_mount_path}/keys | Managed named encryption keys
[**TransitReadCacheConfig**](Secrets.md#TransitReadCacheConfig) | **Get** /{transit_mount_path}/cache-config | Returns the size of the active cache
[**TransitReadConfigKeys**](Secrets.md#TransitReadConfigKeys) | **Get** /{transit_mount_path}/config/keys | Configuration common across all keys
[**TransitReadKey**](Secrets.md#TransitReadKey) | **Get** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**TransitReadWrappingKey**](Secrets.md#TransitReadWrappingKey) | **Get** /{transit_mount_path}/wrapping_key | Returns the public key to use for wrapping imported keys
[**TransitRestore**](Secrets.md#TransitRestore) | **Post** /{transit_mount_path}/restore | Restore the named key
[**TransitRestoreKey**](Secrets.md#TransitRestoreKey) | **Post** /{transit_mount_path}/restore/{name} | Restore the named key
[**TransitRewrap**](Secrets.md#TransitRewrap) | **Post** /{transit_mount_path}/rewrap/{name} | Rewrap ciphertext
[**TransitRotateKey**](Secrets.md#TransitRotateKey) | **Post** /{transit_mount_path}/keys/{name}/rotate | Rotate named encryption key
[**TransitSign**](Secrets.md#TransitSign) | **Post** /{transit_mount_path}/sign/{name} | Generate a signature for input data using the named key
[**TransitSignWithAlgorithm**](Secrets.md#TransitSignWithAlgorithm) | **Post** /{transit_mount_path}/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**TransitTrimKey**](Secrets.md#TransitTrimKey) | **Post** /{transit_mount_path}/keys/{name}/trim | Trim key versions of a named key
[**TransitVerify**](Secrets.md#TransitVerify) | **Post** /{transit_mount_path}/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**TransitVerifyWithAlgorithm**](Secrets.md#TransitVerifyWithAlgorithm) | **Post** /{transit_mount_path}/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key
[**TransitWriteCacheConfig**](Secrets.md#TransitWriteCacheConfig) | **Post** /{transit_mount_path}/cache-config | Configures a new cache of the specified size
[**TransitWriteConfigKeys**](Secrets.md#TransitWriteConfigKeys) | **Post** /{transit_mount_path}/config/keys | Configuration common across all keys
[**TransitWriteKey**](Secrets.md#TransitWriteKey) | **Post** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**TransitWriteKeyConfig**](Secrets.md#TransitWriteKeyConfig) | **Post** /{transit_mount_path}/keys/{name}/config | Configure a named encryption key
[**TransitWriteRandomUrlbytes**](Secrets.md#TransitWriteRandomUrlbytes) | **Post** /{transit_mount_path}/random/{urlbytes} | Generate random bytes





## AWSConfigReadLease

> AWSConfigReadLease(ctx, awsMountPath).Execute()

Configure the default lease information for generated credentials.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AWSConfigReadLease(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadRootIAMCredentials

> AWSConfigReadRootIAMCredentials(ctx, awsMountPath).Execute()

Configure the root credentials that are used to manage IAM.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AWSConfigReadRootIAMCredentials(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigRotateRootIAMCredentials

> AWSConfigRotateRootIAMCredentials(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AWSConfigRotateRootIAMCredentials(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteLease

> AWSConfigWriteLease(ctx, awsMountPath).AWSConfigWriteLeaseRequest(aWSConfigWriteLeaseRequest).Execute()

Configure the default lease information for generated credentials.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteLeaseRequestWithDefaults()

	resp, err := client.Secrets.AWSConfigWriteLease(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteLeaseRequest** | [**AWSConfigWriteLeaseRequest**](AWSConfigWriteLeaseRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteRootIAMCredentials

> AWSConfigWriteRootIAMCredentials(ctx, awsMountPath).AWSConfigWriteRootIAMCredentialsRequest(aWSConfigWriteRootIAMCredentialsRequest).Execute()

Configure the root credentials that are used to manage IAM.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteRootIAMCredentialsRequestWithDefaults()

	resp, err := client.Secrets.AWSConfigWriteRootIAMCredentials(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteRootIAMCredentialsRequest** | [**AWSConfigWriteRootIAMCredentialsRequest**](AWSConfigWriteRootIAMCredentialsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteRole

> AWSDeleteRole(ctx, awsMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy


	resp, err := client.Secrets.AWSDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListRoles

> AWSListRoles(ctx, awsMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AWSListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadCredentials

> AWSReadCredentials(ctx, awsMountPath).Execute()

Generate AWS credentials from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AWSReadCredentials(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadRole

> AWSReadRole(ctx, awsMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy


	resp, err := client.Secrets.AWSReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadSecurityTokenService

> AWSReadSecurityTokenService(ctx, awsMountPath, name).Execute()

Generate AWS credentials from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.AWSReadSecurityTokenService(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteCredentials

> AWSWriteCredentials(ctx, awsMountPath).AWSWriteCredentialsRequest(aWSWriteCredentialsRequest).Execute()

Generate AWS credentials from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSWriteCredentialsRequestWithDefaults()

	resp, err := client.Secrets.AWSWriteCredentials(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSWriteCredentialsRequest** | [**AWSWriteCredentialsRequest**](AWSWriteCredentialsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteRole

> AWSWriteRole(ctx, awsMountPath, name).AWSWriteRoleRequest(aWSWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy

	request := schema.NewAWSWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.AWSWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSWriteRoleRequest** | [**AWSWriteRoleRequest**](AWSWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteSecurityTokenService

> AWSWriteSecurityTokenService(ctx, awsMountPath, name).AWSWriteSecurityTokenServiceRequest(aWSWriteSecurityTokenServiceRequest).Execute()

Generate AWS credentials from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewAWSWriteSecurityTokenServiceRequestWithDefaults()

	resp, err := client.Secrets.AWSWriteSecurityTokenService(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSWriteSecurityTokenServiceRequest** | [**AWSWriteSecurityTokenServiceRequest**](AWSWriteSecurityTokenServiceRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryCheckInLibrary

> ActiveDirectoryCheckInLibrary(ctx, adMountPath, name).ActiveDirectoryCheckInLibraryRequest(activeDirectoryCheckInLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewActiveDirectoryCheckInLibraryRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryCheckInLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryCheckInLibraryRequest** | [**ActiveDirectoryCheckInLibraryRequest**](ActiveDirectoryCheckInLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryCheckInManageLibrary

> ActiveDirectoryCheckInManageLibrary(ctx, adMountPath, name).ActiveDirectoryCheckInManageLibraryRequest(activeDirectoryCheckInManageLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewActiveDirectoryCheckInManageLibraryRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryCheckInManageLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryCheckInManageLibraryRequest** | [**ActiveDirectoryCheckInManageLibraryRequest**](ActiveDirectoryCheckInManageLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryCheckOutLibrary

> ActiveDirectoryCheckOutLibrary(ctx, adMountPath, name).ActiveDirectoryCheckOutLibraryRequest(activeDirectoryCheckOutLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set

	request := schema.NewActiveDirectoryCheckOutLibraryRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryCheckOutLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryCheckOutLibraryRequest** | [**ActiveDirectoryCheckOutLibraryRequest**](ActiveDirectoryCheckOutLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryDeleteConfig

> ActiveDirectoryDeleteConfig(ctx, adMountPath).Execute()

Configure the AD server to connect to, along with password options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ActiveDirectoryDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryDeleteLibrary

> ActiveDirectoryDeleteLibrary(ctx, adMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.ActiveDirectoryDeleteLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryDeleteRole

> ActiveDirectoryDeleteRole(ctx, adMountPath, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.ActiveDirectoryDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryListLibraries

> ActiveDirectoryListLibraries(ctx, adMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ActiveDirectoryListLibraries(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryListRoles

> ActiveDirectoryListRoles(ctx, adMountPath).List(list).Execute()

List the name of each role currently stored.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ActiveDirectoryListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryReadConfig

> ActiveDirectoryReadConfig(ctx, adMountPath).Execute()

Configure the AD server to connect to, along with password options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ActiveDirectoryReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryReadCredentials

> ActiveDirectoryReadCredentials(ctx, adMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.ActiveDirectoryReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryReadLibrary

> ActiveDirectoryReadLibrary(ctx, adMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.ActiveDirectoryReadLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryReadLibraryStatus

> ActiveDirectoryReadLibraryStatus(ctx, adMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.ActiveDirectoryReadLibraryStatus(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryReadRole

> ActiveDirectoryReadRole(ctx, adMountPath, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.ActiveDirectoryReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryRotateRole

> ActiveDirectoryRotateRole(ctx, adMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role


	resp, err := client.Secrets.ActiveDirectoryRotateRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryRotateRoot

> ActiveDirectoryRotateRoot(ctx, adMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ActiveDirectoryRotateRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryWriteConfig

> ActiveDirectoryWriteConfig(ctx, adMountPath).ActiveDirectoryWriteConfigRequest(activeDirectoryWriteConfigRequest).Execute()

Configure the AD server to connect to, along with password options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewActiveDirectoryWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeDirectoryWriteConfigRequest** | [**ActiveDirectoryWriteConfigRequest**](ActiveDirectoryWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryWriteLibrary

> ActiveDirectoryWriteLibrary(ctx, adMountPath, name).ActiveDirectoryWriteLibraryRequest(activeDirectoryWriteLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewActiveDirectoryWriteLibraryRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryWriteLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryWriteLibraryRequest** | [**ActiveDirectoryWriteLibraryRequest**](ActiveDirectoryWriteLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ActiveDirectoryWriteRole

> ActiveDirectoryWriteRole(ctx, adMountPath, name).ActiveDirectoryWriteRoleRequest(activeDirectoryWriteRoleRequest).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewActiveDirectoryWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.ActiveDirectoryWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryWriteRoleRequest** | [**ActiveDirectoryWriteRoleRequest**](ActiveDirectoryWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudDeleteConfig

> AliCloudDeleteConfig(ctx, alicloudMountPath).Execute()

Configure the access key and secret to use for RAM and STS calls.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AliCloudDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudDeleteRole

> AliCloudDeleteRole(ctx, alicloudMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.


	resp, err := client.Secrets.AliCloudDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudListRoles

> AliCloudListRoles(ctx, alicloudMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AliCloudListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadConfig

> AliCloudReadConfig(ctx, alicloudMountPath).Execute()

Configure the access key and secret to use for RAM and STS calls.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AliCloudReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadCredentials

> AliCloudReadCredentials(ctx, alicloudMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.


	resp, err := client.Secrets.AliCloudReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadRole

> AliCloudReadRole(ctx, alicloudMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.


	resp, err := client.Secrets.AliCloudReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudWriteConfig

> AliCloudWriteConfig(ctx, alicloudMountPath).AliCloudWriteConfigRequest(aliCloudWriteConfigRequest).Execute()

Configure the access key and secret to use for RAM and STS calls.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAliCloudWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.AliCloudWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliCloudWriteConfigRequest** | [**AliCloudWriteConfigRequest**](AliCloudWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudWriteRole

> AliCloudWriteRole(ctx, alicloudMountPath, name).AliCloudWriteRoleRequest(aliCloudWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.

	request := schema.NewAliCloudWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.AliCloudWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliCloudWriteRoleRequest** | [**AliCloudWriteRoleRequest**](AliCloudWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteConfig

> AzureDeleteConfig(ctx, azureMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AzureDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteRole

> AzureDeleteRole(ctx, azureMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.AzureDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureListRoles

> AzureListRoles(ctx, azureMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AzureListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadConfig

> AzureReadConfig(ctx, azureMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AzureReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadCredentials

> AzureReadCredentials(ctx, azureMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the Vault role


	resp, err := client.Secrets.AzureReadCredentials(
		context.Background(),
		role,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadRole

> AzureReadRole(ctx, azureMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.AzureReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureRotateRoot

> AzureRotateRoot(ctx, azureMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.AzureRotateRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteConfig

> AzureWriteConfig(ctx, azureMountPath).AzureWriteConfigRequest(azureWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAzureWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.AzureWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureWriteConfigRequest** | [**AzureWriteConfigRequest**](AzureWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteRole

> AzureWriteRole(ctx, azureMountPath, name).AzureWriteRoleRequest(azureWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewAzureWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.AzureWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureWriteRoleRequest** | [**AzureWriteRoleRequest**](AzureWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulDeleteRole

> ConsulDeleteRole(ctx, consulMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.ConsulDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulListRoles

> ConsulListRoles(ctx, consulMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ConsulListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulReadAccessConfig

> ConsulReadAccessConfig(ctx, consulMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.ConsulReadAccessConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulReadCredentials

> ConsulReadCredentials(ctx, consulMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.


	resp, err := client.Secrets.ConsulReadCredentials(
		context.Background(),
		role,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulReadRole

> ConsulReadRole(ctx, consulMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.ConsulReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulWriteAccessConfig

> ConsulWriteAccessConfig(ctx, consulMountPath).ConsulWriteAccessConfigRequest(consulWriteAccessConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewConsulWriteAccessConfigRequestWithDefaults()

	resp, err := client.Secrets.ConsulWriteAccessConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulWriteAccessConfigRequest** | [**ConsulWriteAccessConfigRequest**](ConsulWriteAccessConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ConsulWriteRole

> ConsulWriteRole(ctx, consulMountPath, name).ConsulWriteRoleRequest(consulWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewConsulWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.ConsulWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consulWriteRoleRequest** | [**ConsulWriteRoleRequest**](ConsulWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CubbyholeDelete

> CubbyholeDelete(ctx, cubbyholeMountPath, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> CubbyholeRead(ctx, cubbyholeMountPath, path).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.


	list := "list_example" // string | Return a list if `true`
	resp, err := client.Secrets.CubbyholeRead(
		context.Background(),
		path,
		list,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

> CubbyholeWrite(ctx, cubbyholeMountPath, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## GoogleCloudDeleteRoleset

> GoogleCloudDeleteRoleset(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.


	resp, err := client.Secrets.GoogleCloudDeleteRoleset(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteStaticAccount

> GoogleCloudDeleteStaticAccount(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.


	resp, err := client.Secrets.GoogleCloudDeleteStaticAccount(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSDecrypt

> GoogleCloudKMSDecrypt(ctx, gcpkmsMountPath, key).GoogleCloudKMSDecryptRequest(googleCloudKMSDecryptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.

	request := schema.NewGoogleCloudKMSDecryptRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSDecrypt(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSDecryptRequest** | [**GoogleCloudKMSDecryptRequest**](GoogleCloudKMSDecryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSDeleteConfig

> GoogleCloudKMSDeleteConfig(ctx, gcpkmsMountPath).Execute()

Configure the GCP KMS secrets engine

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudKMSDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSDeleteKey

> GoogleCloudKMSDeleteKey(ctx, gcpkmsMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.


	resp, err := client.Secrets.GoogleCloudKMSDeleteKey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSDeregisterKey

> GoogleCloudKMSDeregisterKey(ctx, gcpkmsMountPath, key).Execute()

Deregister an existing key in Vault

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.


	resp, err := client.Secrets.GoogleCloudKMSDeregisterKey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSEncrypt

> GoogleCloudKMSEncrypt(ctx, gcpkmsMountPath, key).GoogleCloudKMSEncryptRequest(googleCloudKMSEncryptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.

	request := schema.NewGoogleCloudKMSEncryptRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSEncrypt(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSEncryptRequest** | [**GoogleCloudKMSEncryptRequest**](GoogleCloudKMSEncryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSListKeys

> GoogleCloudKMSListKeys(ctx, gcpkmsMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudKMSListKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSReadConfig

> GoogleCloudKMSReadConfig(ctx, gcpkmsMountPath).Execute()

Configure the GCP KMS secrets engine

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudKMSReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSReadKey

> GoogleCloudKMSReadKey(ctx, gcpkmsMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.


	resp, err := client.Secrets.GoogleCloudKMSReadKey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSReadKeyConfig

> GoogleCloudKMSReadKeyConfig(ctx, gcpkmsMountPath, key).Execute()

Configure the key in Vault

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.


	resp, err := client.Secrets.GoogleCloudKMSReadKeyConfig(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSReadPubkey

> GoogleCloudKMSReadPubkey(ctx, gcpkmsMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.


	resp, err := client.Secrets.GoogleCloudKMSReadPubkey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSReencrypt

> GoogleCloudKMSReencrypt(ctx, gcpkmsMountPath, key).GoogleCloudKMSReencryptRequest(googleCloudKMSReencryptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.

	request := schema.NewGoogleCloudKMSReencryptRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSReencrypt(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSReencryptRequest** | [**GoogleCloudKMSReencryptRequest**](GoogleCloudKMSReencryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSRegisterKey

> GoogleCloudKMSRegisterKey(ctx, gcpkmsMountPath, key).GoogleCloudKMSRegisterKeyRequest(googleCloudKMSRegisterKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.

	request := schema.NewGoogleCloudKMSRegisterKeyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSRegisterKey(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSRegisterKeyRequest** | [**GoogleCloudKMSRegisterKeyRequest**](GoogleCloudKMSRegisterKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSRotateKey

> GoogleCloudKMSRotateKey(ctx, gcpkmsMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.


	resp, err := client.Secrets.GoogleCloudKMSRotateKey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSSign

> GoogleCloudKMSSign(ctx, gcpkmsMountPath, key).GoogleCloudKMSSignRequest(googleCloudKMSSignRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.

	request := schema.NewGoogleCloudKMSSignRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSSign(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSSignRequest** | [**GoogleCloudKMSSignRequest**](GoogleCloudKMSSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSTrimKey

> GoogleCloudKMSTrimKey(ctx, gcpkmsMountPath, key).Execute()

Delete old crypto key versions from Google Cloud KMS

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.


	resp, err := client.Secrets.GoogleCloudKMSTrimKey(
		context.Background(),
		key,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSVerify

> GoogleCloudKMSVerify(ctx, gcpkmsMountPath, key).GoogleCloudKMSVerifyRequest(googleCloudKMSVerifyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.

	request := schema.NewGoogleCloudKMSVerifyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSVerify(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSVerifyRequest** | [**GoogleCloudKMSVerifyRequest**](GoogleCloudKMSVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSWriteConfig

> GoogleCloudKMSWriteConfig(ctx, gcpkmsMountPath).GoogleCloudKMSWriteConfigRequest(googleCloudKMSWriteConfigRequest).Execute()

Configure the GCP KMS secrets engine

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGoogleCloudKMSWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleCloudKMSWriteConfigRequest** | [**GoogleCloudKMSWriteConfigRequest**](GoogleCloudKMSWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSWriteKey

> GoogleCloudKMSWriteKey(ctx, gcpkmsMountPath, key).GoogleCloudKMSWriteKeyRequest(googleCloudKMSWriteKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.

	request := schema.NewGoogleCloudKMSWriteKeyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSWriteKey(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSWriteKeyRequest** | [**GoogleCloudKMSWriteKeyRequest**](GoogleCloudKMSWriteKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudKMSWriteKeyConfig

> GoogleCloudKMSWriteKeyConfig(ctx, gcpkmsMountPath, key).GoogleCloudKMSWriteKeyConfigRequest(googleCloudKMSWriteKeyConfigRequest).Execute()

Configure the key in Vault

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.

	request := schema.NewGoogleCloudKMSWriteKeyConfigRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudKMSWriteKeyConfig(
		context.Background(),
		key,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudKMSWriteKeyConfigRequest** | [**GoogleCloudKMSWriteKeyConfigRequest**](GoogleCloudKMSWriteKeyConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRolesets

> GoogleCloudListRolesets(ctx, gcpMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudListRolesets(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListStaticAccounts

> GoogleCloudListStaticAccounts(ctx, gcpMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudListStaticAccounts(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadConfig

> GoogleCloudReadConfig(ctx, gcpMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadKey

> GoogleCloudReadKey(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudReadKey(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRoleset

> GoogleCloudReadRoleset(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.


	resp, err := client.Secrets.GoogleCloudReadRoleset(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRolesetKey

> GoogleCloudReadRolesetKey(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudReadRolesetKey(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRolesetToken

> GoogleCloudReadRolesetToken(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudReadRolesetToken(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadStaticAccount

> GoogleCloudReadStaticAccount(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.


	resp, err := client.Secrets.GoogleCloudReadStaticAccount(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadStaticAccountKey

> GoogleCloudReadStaticAccountKey(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.


	resp, err := client.Secrets.GoogleCloudReadStaticAccountKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadStaticAccountToken

> GoogleCloudReadStaticAccountToken(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.


	resp, err := client.Secrets.GoogleCloudReadStaticAccountToken(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadToken

> GoogleCloudReadToken(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudReadToken(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRoleset

> GoogleCloudRotateRoleset(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.GoogleCloudRotateRoleset(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRolesetKey

> GoogleCloudRotateRolesetKey(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.GoogleCloudRotateRolesetKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateRoot

> GoogleCloudRotateRoot(ctx, gcpMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.GoogleCloudRotateRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudRotateStaticAccountKey

> GoogleCloudRotateStaticAccountKey(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the account.


	resp, err := client.Secrets.GoogleCloudRotateStaticAccountKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteConfig

> GoogleCloudWriteConfig(ctx, gcpMountPath).GoogleCloudWriteConfigRequest(googleCloudWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGoogleCloudWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleCloudWriteConfigRequest** | [**GoogleCloudWriteConfigRequest**](GoogleCloudWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteKey

> GoogleCloudWriteKey(ctx, gcpMountPath, roleset).GoogleCloudWriteKeyRequest(googleCloudWriteKeyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.

	request := schema.NewGoogleCloudWriteKeyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteKey(
		context.Background(),
		roleset,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteKeyRequest** | [**GoogleCloudWriteKeyRequest**](GoogleCloudWriteKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRoleset

> GoogleCloudWriteRoleset(ctx, gcpMountPath, name).GoogleCloudWriteRolesetRequest(googleCloudWriteRolesetRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.

	request := schema.NewGoogleCloudWriteRolesetRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteRoleset(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteRolesetRequest** | [**GoogleCloudWriteRolesetRequest**](GoogleCloudWriteRolesetRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRolesetKey

> GoogleCloudWriteRolesetKey(ctx, gcpMountPath, roleset).GoogleCloudWriteRolesetKeyRequest(googleCloudWriteRolesetKeyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.

	request := schema.NewGoogleCloudWriteRolesetKeyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteRolesetKey(
		context.Background(),
		roleset,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteRolesetKeyRequest** | [**GoogleCloudWriteRolesetKeyRequest**](GoogleCloudWriteRolesetKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRolesetToken

> GoogleCloudWriteRolesetToken(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudWriteRolesetToken(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteStaticAccount

> GoogleCloudWriteStaticAccount(ctx, gcpMountPath, name).GoogleCloudWriteStaticAccountRequest(googleCloudWriteStaticAccountRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.

	request := schema.NewGoogleCloudWriteStaticAccountRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteStaticAccount(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteStaticAccountRequest** | [**GoogleCloudWriteStaticAccountRequest**](GoogleCloudWriteStaticAccountRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteStaticAccountKey

> GoogleCloudWriteStaticAccountKey(ctx, gcpMountPath, name).GoogleCloudWriteStaticAccountKeyRequest(googleCloudWriteStaticAccountKeyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.

	request := schema.NewGoogleCloudWriteStaticAccountKeyRequestWithDefaults()

	resp, err := client.Secrets.GoogleCloudWriteStaticAccountKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteStaticAccountKeyRequest** | [**GoogleCloudWriteStaticAccountKeyRequest**](GoogleCloudWriteStaticAccountKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteStaticAccountToken

> GoogleCloudWriteStaticAccountToken(ctx, gcpMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.


	resp, err := client.Secrets.GoogleCloudWriteStaticAccountToken(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteToken

> GoogleCloudWriteToken(ctx, gcpMountPath, roleset).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.


	resp, err := client.Secrets.GoogleCloudWriteToken(
		context.Background(),
		roleset,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv1Delete

> KVv1Delete(ctx, kvMountPath, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv1Delete(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv1Read

> KVv1Read(ctx, kvMountPath, path).List(list).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	list := "list_example" // string | Return a list if `true`
	resp, err := client.Secrets.KVv1Read(
		context.Background(),
		path,
		list,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv1Write

> KVv1Write(ctx, kvMountPath, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv1Write(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2Delete

> KVv2Delete(ctx, path, secretMountPath).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv2Delete(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2DeleteMetadata

> KVv2DeleteMetadata(ctx, path, secretMountPath).Execute()

Configures settings for the KV store

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv2DeleteMetadata(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2DeleteVersions

> KVv2DeleteVersions(ctx, path, secretMountPath).KVv2DeleteVersionsRequest(kVv2DeleteVersionsRequest).Execute()

Marks one or more versions as deleted in the KV store.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.

	request := schema.NewKVv2DeleteVersionsRequestWithDefaults()

	resp, err := client.Secrets.KVv2DeleteVersions(
		context.Background(),
		path,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kVv2DeleteVersionsRequest** | [**KVv2DeleteVersionsRequest**](KVv2DeleteVersionsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2DestroyVersions

> KVv2DestroyVersions(ctx, path, secretMountPath).KVv2DestroyVersionsRequest(kVv2DestroyVersionsRequest).Execute()

Permanently removes one or more versions in the KV store

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.

	request := schema.NewKVv2DestroyVersionsRequestWithDefaults()

	resp, err := client.Secrets.KVv2DestroyVersions(
		context.Background(),
		path,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kVv2DestroyVersionsRequest** | [**KVv2DestroyVersionsRequest**](KVv2DestroyVersionsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2Read

> KVv2Read(ctx, path, secretMountPath).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv2Read(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2ReadConfig

> KVv2ReadConfig(ctx, secretMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.KVv2ReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2ReadMetadata

> KVv2ReadMetadata(ctx, path, secretMountPath).List(list).Execute()

Configures settings for the KV store

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	list := "list_example" // string | Return a list if `true`
	resp, err := client.Secrets.KVv2ReadMetadata(
		context.Background(),
		path,
		list,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2ReadSubkeys

> KVv2ReadSubkeys(ctx, path, secretMountPath).Execute()

Read the structure of a secret entry from the Key-Value store with the values removed.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.


	resp, err := client.Secrets.KVv2ReadSubkeys(
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
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2UndeleteVersions

> KVv2UndeleteVersions(ctx, path, secretMountPath).KVv2UndeleteVersionsRequest(kVv2UndeleteVersionsRequest).Execute()

Undeletes one or more versions from the KV store.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.

	request := schema.NewKVv2UndeleteVersionsRequestWithDefaults()

	resp, err := client.Secrets.KVv2UndeleteVersions(
		context.Background(),
		path,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kVv2UndeleteVersionsRequest** | [**KVv2UndeleteVersionsRequest**](KVv2UndeleteVersionsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2Write

> KVv2Write(ctx, path, secretMountPath).KVv2WriteRequest(kVv2WriteRequest).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.

	request := schema.NewKVv2WriteRequestWithDefaults()

	resp, err := client.Secrets.KVv2Write(
		context.Background(),
		path,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kVv2WriteRequest** | [**KVv2WriteRequest**](KVv2WriteRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2WriteConfig

> KVv2WriteConfig(ctx, secretMountPath).KVv2WriteConfigRequest(kVv2WriteConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKVv2WriteConfigRequestWithDefaults()

	resp, err := client.Secrets.KVv2WriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kVv2WriteConfigRequest** | [**KVv2WriteConfigRequest**](KVv2WriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KVv2WriteMetadata

> KVv2WriteMetadata(ctx, path, secretMountPath).KVv2WriteMetadataRequest(kVv2WriteMetadataRequest).Execute()

Configures settings for the KV store

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.

	request := schema.NewKVv2WriteMetadataRequestWithDefaults()

	resp, err := client.Secrets.KVv2WriteMetadata(
		context.Background(),
		path,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kVv2WriteMetadataRequest** | [**KVv2WriteMetadataRequest**](KVv2WriteMetadataRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteConfig

> KubernetesDeleteConfig(ctx, kubernetesMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.KubernetesDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteRole

> KubernetesDeleteRole(ctx, kubernetesMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.KubernetesDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesListRoles

> KubernetesListRoles(ctx, kubernetesMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.KubernetesListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadConfig

> KubernetesReadConfig(ctx, kubernetesMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.KubernetesReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadRole

> KubernetesReadRole(ctx, kubernetesMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.KubernetesReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteConfig

> KubernetesWriteConfig(ctx, kubernetesMountPath).KubernetesWriteConfigRequest(kubernetesWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKubernetesWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.KubernetesWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesWriteConfigRequest** | [**KubernetesWriteConfigRequest**](KubernetesWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteCredentials

> KubernetesWriteCredentials(ctx, kubernetesMountPath, name).KubernetesWriteCredentialsRequest(kubernetesWriteCredentialsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Vault role

	request := schema.NewKubernetesWriteCredentialsRequestWithDefaults()

	resp, err := client.Secrets.KubernetesWriteCredentials(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesWriteCredentialsRequest** | [**KubernetesWriteCredentialsRequest**](KubernetesWriteCredentialsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteRole

> KubernetesWriteRole(ctx, kubernetesMountPath, name).KubernetesWriteRoleRequest(kubernetesWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewKubernetesWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.KubernetesWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesWriteRoleRequest** | [**KubernetesWriteRoleRequest**](KubernetesWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPCheckInLibrary

> LDAPCheckInLibrary(ctx, ldapMountPath, name).LDAPCheckInLibraryRequest(lDAPCheckInLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewLDAPCheckInLibraryRequestWithDefaults()

	resp, err := client.Secrets.LDAPCheckInLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPCheckInLibraryRequest** | [**LDAPCheckInLibraryRequest**](LDAPCheckInLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPCheckInManageLibrary

> LDAPCheckInManageLibrary(ctx, ldapMountPath, name).LDAPCheckInManageLibraryRequest(lDAPCheckInManageLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewLDAPCheckInManageLibraryRequestWithDefaults()

	resp, err := client.Secrets.LDAPCheckInManageLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPCheckInManageLibraryRequest** | [**LDAPCheckInManageLibraryRequest**](LDAPCheckInManageLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPCheckOutLibrary

> LDAPCheckOutLibrary(ctx, ldapMountPath, name).LDAPCheckOutLibraryRequest(lDAPCheckOutLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set

	request := schema.NewLDAPCheckOutLibraryRequestWithDefaults()

	resp, err := client.Secrets.LDAPCheckOutLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPCheckOutLibraryRequest** | [**LDAPCheckOutLibraryRequest**](LDAPCheckOutLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPDeleteConfig

> LDAPDeleteConfig(ctx, ldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPDeleteLibrary

> LDAPDeleteLibrary(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.LDAPDeleteLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPDeleteRole

> LDAPDeleteRole(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)


	resp, err := client.Secrets.LDAPDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPDeleteStaticRole

> LDAPDeleteStaticRole(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.LDAPDeleteStaticRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPListLibraries

> LDAPListLibraries(ctx, ldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPListLibraries(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPListRoles

> LDAPListRoles(ctx, ldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPListStaticRoles

> LDAPListStaticRoles(ctx, ldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPListStaticRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadConfig

> LDAPReadConfig(ctx, ldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadCredentials

> LDAPReadCredentials(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the dynamic role.


	resp, err := client.Secrets.LDAPReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadLibrary

> LDAPReadLibrary(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.LDAPReadLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadLibraryStatus

> LDAPReadLibraryStatus(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.LDAPReadLibraryStatus(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadRole

> LDAPReadRole(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)


	resp, err := client.Secrets.LDAPReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadStaticCredentials

> LDAPReadStaticCredentials(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.


	resp, err := client.Secrets.LDAPReadStaticCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadStaticRole

> LDAPReadStaticRole(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.LDAPReadStaticRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPRotateRole

> LDAPRotateRole(ctx, ldapMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role


	resp, err := client.Secrets.LDAPRotateRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPRotateRoot

> LDAPRotateRoot(ctx, ldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.LDAPRotateRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteConfig

> LDAPWriteConfig(ctx, ldapMountPath).LDAPWriteConfigRequest(lDAPWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewLDAPWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.LDAPWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPWriteConfigRequest** | [**LDAPWriteConfigRequest**](LDAPWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteLibrary

> LDAPWriteLibrary(ctx, ldapMountPath, name).LDAPWriteLibraryRequest(lDAPWriteLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewLDAPWriteLibraryRequestWithDefaults()

	resp, err := client.Secrets.LDAPWriteLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPWriteLibraryRequest** | [**LDAPWriteLibraryRequest**](LDAPWriteLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteRole

> LDAPWriteRole(ctx, ldapMountPath, name).LDAPWriteRoleRequest(lDAPWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)

	request := schema.NewLDAPWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.LDAPWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPWriteRoleRequest** | [**LDAPWriteRoleRequest**](LDAPWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteStaticRole

> LDAPWriteStaticRole(ctx, ldapMountPath, name).LDAPWriteStaticRoleRequest(lDAPWriteStaticRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewLDAPWriteStaticRoleRequestWithDefaults()

	resp, err := client.Secrets.LDAPWriteStaticRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPWriteStaticRoleRequest** | [**LDAPWriteStaticRoleRequest**](LDAPWriteStaticRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasDeleteRole

> MongoDBAtlasDeleteRole(ctx, mongodbatlasMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles


	resp, err := client.Secrets.MongoDBAtlasDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasListRoles

> MongoDBAtlasListRoles(ctx, mongodbatlasMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.MongoDBAtlasListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasReadConfig

> MongoDBAtlasReadConfig(ctx, mongodbatlasMountPath).Execute()

Configure the  credentials that are used to manage Database Users.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.MongoDBAtlasReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasReadCredentials

> MongoDBAtlasReadCredentials(ctx, mongodbatlasMountPath, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.MongoDBAtlasReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasReadRole

> MongoDBAtlasReadRole(ctx, mongodbatlasMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles


	resp, err := client.Secrets.MongoDBAtlasReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasWriteConfig

> MongoDBAtlasWriteConfig(ctx, mongodbatlasMountPath).MongoDBAtlasWriteConfigRequest(mongoDBAtlasWriteConfigRequest).Execute()

Configure the  credentials that are used to manage Database Users.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewMongoDBAtlasWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.MongoDBAtlasWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mongoDBAtlasWriteConfigRequest** | [**MongoDBAtlasWriteConfigRequest**](MongoDBAtlasWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasWriteCredentials

> MongoDBAtlasWriteCredentials(ctx, mongodbatlasMountPath, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.MongoDBAtlasWriteCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MongoDBAtlasWriteRole

> MongoDBAtlasWriteRole(ctx, mongodbatlasMountPath, name).MongoDBAtlasWriteRoleRequest(mongoDBAtlasWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles

	request := schema.NewMongoDBAtlasWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.MongoDBAtlasWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mongoDBAtlasWriteRoleRequest** | [**MongoDBAtlasWriteRoleRequest**](MongoDBAtlasWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteAccessConfig

> NomadDeleteAccessConfig(ctx, nomadMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.NomadDeleteAccessConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteLeaseConfig

> NomadDeleteLeaseConfig(ctx, nomadMountPath).Execute()

Configure the lease parameters for generated tokens

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.NomadDeleteLeaseConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadDeleteRole

> NomadDeleteRole(ctx, name, nomadMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.NomadDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadListRoles

> NomadListRoles(ctx, nomadMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.NomadListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadAccessConfig

> NomadReadAccessConfig(ctx, nomadMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.NomadReadAccessConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadCredentials

> NomadReadCredentials(ctx, name, nomadMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.NomadReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadLeaseConfig

> NomadReadLeaseConfig(ctx, nomadMountPath).Execute()

Configure the lease parameters for generated tokens

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.NomadReadLeaseConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadReadRole

> NomadReadRole(ctx, name, nomadMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.NomadReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadWriteAccessConfig

> NomadWriteAccessConfig(ctx, nomadMountPath).NomadWriteAccessConfigRequest(nomadWriteAccessConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewNomadWriteAccessConfigRequestWithDefaults()

	resp, err := client.Secrets.NomadWriteAccessConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadWriteAccessConfigRequest** | [**NomadWriteAccessConfigRequest**](NomadWriteAccessConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadWriteLeaseConfig

> NomadWriteLeaseConfig(ctx, nomadMountPath).NomadWriteLeaseConfigRequest(nomadWriteLeaseConfigRequest).Execute()

Configure the lease parameters for generated tokens

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewNomadWriteLeaseConfigRequestWithDefaults()

	resp, err := client.Secrets.NomadWriteLeaseConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadWriteLeaseConfigRequest** | [**NomadWriteLeaseConfigRequest**](NomadWriteLeaseConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## NomadWriteRole

> NomadWriteRole(ctx, name, nomadMountPath).NomadWriteRoleRequest(nomadWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewNomadWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.NomadWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadWriteRoleRequest** | [**NomadWriteRoleRequest**](NomadWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPCheckInLibrary

> OpenLDAPCheckInLibrary(ctx, name, openldapMountPath).OpenLDAPCheckInLibraryRequest(openLDAPCheckInLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewOpenLDAPCheckInLibraryRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPCheckInLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPCheckInLibraryRequest** | [**OpenLDAPCheckInLibraryRequest**](OpenLDAPCheckInLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPCheckInManageLibrary

> OpenLDAPCheckInManageLibrary(ctx, name, openldapMountPath).OpenLDAPCheckInManageLibraryRequest(openLDAPCheckInManageLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewOpenLDAPCheckInManageLibraryRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPCheckInManageLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPCheckInManageLibraryRequest** | [**OpenLDAPCheckInManageLibraryRequest**](OpenLDAPCheckInManageLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPCheckOutLibrary

> OpenLDAPCheckOutLibrary(ctx, name, openldapMountPath).OpenLDAPCheckOutLibraryRequest(openLDAPCheckOutLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set

	request := schema.NewOpenLDAPCheckOutLibraryRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPCheckOutLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPCheckOutLibraryRequest** | [**OpenLDAPCheckOutLibraryRequest**](OpenLDAPCheckOutLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPDeleteConfig

> OpenLDAPDeleteConfig(ctx, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPDeleteLibrary

> OpenLDAPDeleteLibrary(ctx, name, openldapMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.OpenLDAPDeleteLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPDeleteRole

> OpenLDAPDeleteRole(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)


	resp, err := client.Secrets.OpenLDAPDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPDeleteStaticRole

> OpenLDAPDeleteStaticRole(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.OpenLDAPDeleteStaticRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPListLibraries

> OpenLDAPListLibraries(ctx, openldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPListLibraries(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPListRoles

> OpenLDAPListRoles(ctx, openldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPListStaticRoles

> OpenLDAPListStaticRoles(ctx, openldapMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPListStaticRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadConfig

> OpenLDAPReadConfig(ctx, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadCredentials

> OpenLDAPReadCredentials(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the dynamic role.


	resp, err := client.Secrets.OpenLDAPReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadLibrary

> OpenLDAPReadLibrary(ctx, name, openldapMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.OpenLDAPReadLibrary(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadLibraryStatus

> OpenLDAPReadLibraryStatus(ctx, name, openldapMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.


	resp, err := client.Secrets.OpenLDAPReadLibraryStatus(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadRole

> OpenLDAPReadRole(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)


	resp, err := client.Secrets.OpenLDAPReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadStaticCredentials

> OpenLDAPReadStaticCredentials(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.


	resp, err := client.Secrets.OpenLDAPReadStaticCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPReadStaticRole

> OpenLDAPReadStaticRole(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.OpenLDAPReadStaticRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPRotateRole

> OpenLDAPRotateRole(ctx, name, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role


	resp, err := client.Secrets.OpenLDAPRotateRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPRotateRoot

> OpenLDAPRotateRoot(ctx, openldapMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.OpenLDAPRotateRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPWriteConfig

> OpenLDAPWriteConfig(ctx, openldapMountPath).OpenLDAPWriteConfigRequest(openLDAPWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOpenLDAPWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openLDAPWriteConfigRequest** | [**OpenLDAPWriteConfigRequest**](OpenLDAPWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPWriteLibrary

> OpenLDAPWriteLibrary(ctx, name, openldapMountPath).OpenLDAPWriteLibraryRequest(openLDAPWriteLibraryRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.

	request := schema.NewOpenLDAPWriteLibraryRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPWriteLibrary(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPWriteLibraryRequest** | [**OpenLDAPWriteLibraryRequest**](OpenLDAPWriteLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPWriteRole

> OpenLDAPWriteRole(ctx, name, openldapMountPath).OpenLDAPWriteRoleRequest(openLDAPWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)

	request := schema.NewOpenLDAPWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPWriteRoleRequest** | [**OpenLDAPWriteRoleRequest**](OpenLDAPWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OpenLDAPWriteStaticRole

> OpenLDAPWriteStaticRole(ctx, name, openldapMountPath).OpenLDAPWriteStaticRoleRequest(openLDAPWriteStaticRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewOpenLDAPWriteStaticRoleRequestWithDefaults()

	resp, err := client.Secrets.OpenLDAPWriteStaticRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openLDAPWriteStaticRoleRequest** | [**OpenLDAPWriteStaticRoleRequest**](OpenLDAPWriteStaticRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIBundleWrite

> PKIBundleWrite(ctx, pkiMountPath).PKIBundleWriteRequest(pKIBundleWriteRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIBundleWriteRequestWithDefaults()

	resp, err := client.Secrets.PKIBundleWrite(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIBundleWriteRequest** | [**PKIBundleWriteRequest**](PKIBundleWriteRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIDeleteKey

> PKIDeleteKey(ctx, keyRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")


	resp, err := client.Secrets.PKIDeleteKey(
		context.Background(),
		keyRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIDeleteRole

> PKIDeleteRole(ctx, name, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.PKIDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIDeleteRoot

> PKIDeleteRoot(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIDeleteRoot(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIGenerateRoot

> PKIGenerateRoot(ctx, exported, pkiMountPath).PKIGenerateRootRequest(pKIGenerateRootRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!

	request := schema.NewPKIGenerateRootRequestWithDefaults()

	resp, err := client.Secrets.PKIGenerateRoot(
		context.Background(),
		exported,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIGenerateRootRequest** | [**PKIGenerateRootRequest**](PKIGenerateRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIImportKeys

> PKIImportKeys(ctx, pkiMountPath).PKIImportKeysRequest(pKIImportKeysRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIImportKeysRequestWithDefaults()

	resp, err := client.Secrets.PKIImportKeys(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIImportKeysRequest** | [**PKIImportKeysRequest**](PKIImportKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerIssueRole

> PKIIssuerIssueRole(ctx, issuerRef, pkiMountPath, role).PKIIssuerIssueRoleRequest(pKIIssuerIssueRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKIIssuerIssueRoleRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerIssueRole(
		context.Background(),
		issuerRef,
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pKIIssuerIssueRoleRequest** | [**PKIIssuerIssueRoleRequest**](PKIIssuerIssueRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerResignCRLs

> PKIIssuerResignCRLs(ctx, issuerRef, pkiMountPath).PKIIssuerResignCRLsRequest(pKIIssuerResignCRLsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPKIIssuerResignCRLsRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerResignCRLs(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuerResignCRLsRequest** | [**PKIIssuerResignCRLsRequest**](PKIIssuerResignCRLsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerRevoke

> PKIIssuerRevoke(ctx, issuerRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")


	resp, err := client.Secrets.PKIIssuerRevoke(
		context.Background(),
		issuerRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignIntermediate

> PKIIssuerSignIntermediate(ctx, issuerRef, pkiMountPath).PKIIssuerSignIntermediateRequest(pKIIssuerSignIntermediateRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPKIIssuerSignIntermediateRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignIntermediate(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuerSignIntermediateRequest** | [**PKIIssuerSignIntermediateRequest**](PKIIssuerSignIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignRevocationList

> PKIIssuerSignRevocationList(ctx, issuerRef, pkiMountPath).PKIIssuerSignRevocationListRequest(pKIIssuerSignRevocationListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPKIIssuerSignRevocationListRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignRevocationList(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuerSignRevocationListRequest** | [**PKIIssuerSignRevocationListRequest**](PKIIssuerSignRevocationListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignRole

> PKIIssuerSignRole(ctx, issuerRef, pkiMountPath, role).PKIIssuerSignRoleRequest(pKIIssuerSignRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKIIssuerSignRoleRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignRole(
		context.Background(),
		issuerRef,
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pKIIssuerSignRoleRequest** | [**PKIIssuerSignRoleRequest**](PKIIssuerSignRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignSelfIssued

> PKIIssuerSignSelfIssued(ctx, issuerRef, pkiMountPath).PKIIssuerSignSelfIssuedRequest(pKIIssuerSignSelfIssuedRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPKIIssuerSignSelfIssuedRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignSelfIssued(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuerSignSelfIssuedRequest** | [**PKIIssuerSignSelfIssuedRequest**](PKIIssuerSignSelfIssuedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignVerbatim

> PKIIssuerSignVerbatim(ctx, issuerRef, pkiMountPath).PKIIssuerSignVerbatimRequest(pKIIssuerSignVerbatimRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPKIIssuerSignVerbatimRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignVerbatim(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuerSignVerbatimRequest** | [**PKIIssuerSignVerbatimRequest**](PKIIssuerSignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuerSignVerbatimRole

> PKIIssuerSignVerbatimRole(ctx, issuerRef, pkiMountPath, role).PKIIssuerSignVerbatimRoleRequest(pKIIssuerSignVerbatimRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKIIssuerSignVerbatimRoleRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuerSignVerbatimRole(
		context.Background(),
		issuerRef,
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pKIIssuerSignVerbatimRoleRequest** | [**PKIIssuerSignVerbatimRoleRequest**](PKIIssuerSignVerbatimRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuersGenerateIntermediate

> PKIIssuersGenerateIntermediate(ctx, exported, pkiMountPath).PKIIssuersGenerateIntermediateRequest(pKIIssuersGenerateIntermediateRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!

	request := schema.NewPKIIssuersGenerateIntermediateRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuersGenerateIntermediate(
		context.Background(),
		exported,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuersGenerateIntermediateRequest** | [**PKIIssuersGenerateIntermediateRequest**](PKIIssuersGenerateIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuersGenerateRoot

> PKIIssuersGenerateRoot(ctx, exported, pkiMountPath).PKIIssuersGenerateRootRequest(pKIIssuersGenerateRootRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!

	request := schema.NewPKIIssuersGenerateRootRequestWithDefaults()

	resp, err := client.Secrets.PKIIssuersGenerateRoot(
		context.Background(),
		exported,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIIssuersGenerateRootRequest** | [**PKIIssuersGenerateRootRequest**](PKIIssuersGenerateRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIIssuersList

> PKIIssuersList(ctx, pkiMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIIssuersList(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIListCerts

> PKIListCerts(ctx, pkiMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIListCerts(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIListCertsRevoked

> PKIListCertsRevoked(ctx, pkiMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIListCertsRevoked(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIListKeys

> PKIListKeys(ctx, pkiMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIListKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIListRoles

> PKIListRoles(ctx, pkiMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadAutoTidyConfig

> PKIReadAutoTidyConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadAutoTidyConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCA

> PKIReadCA(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCA(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCAChain

> PKIReadCAChain(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCAChain(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCAPem

> PKIReadCAPem(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCAPem(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCRL

> PKIReadCRL(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCRL(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCRLConfig

> PKIReadCRLConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCRLConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCRLRotate

> PKIReadCRLRotate(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCRLRotate(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCRLRotateDelta

> PKIReadCRLRotateDelta(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCRLRotateDelta(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCert

> PKIReadCert(ctx, pkiMountPath, serial).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal


	resp, err := client.Secrets.PKIReadCert(
		context.Background(),
		serial,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCertCAChain

> PKIReadCertCAChain(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadCertCAChain(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCertRaw

> PKIReadCertRaw(ctx, pkiMountPath, serial).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal


	resp, err := client.Secrets.PKIReadCertRaw(
		context.Background(),
		serial,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadCertRawPem

> PKIReadCertRawPem(ctx, pkiMountPath, serial).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal


	resp, err := client.Secrets.PKIReadCertRawPem(
		context.Background(),
		serial,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadClusterConfig

> PKIReadClusterConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadClusterConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadDeltaCRL

> PKIReadDeltaCRL(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadDeltaCRL(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadIssuersConfig

> PKIReadIssuersConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadIssuersConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadKey

> PKIReadKey(ctx, keyRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")


	resp, err := client.Secrets.PKIReadKey(
		context.Background(),
		keyRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadKeysConfig

> PKIReadKeysConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadKeysConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadOCSPReq

> PKIReadOCSPReq(ctx, pkiMountPath, req).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	req := "req_example" // string | base-64 encoded ocsp request


	resp, err := client.Secrets.PKIReadOCSPReq(
		context.Background(),
		req,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadRole

> PKIReadRole(ctx, name, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.PKIReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReadURLConfig

> PKIReadURLConfig(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIReadURLConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIReplaceRoot

> PKIReplaceRoot(ctx, pkiMountPath).PKIReplaceRootRequest(pKIReplaceRootRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIReplaceRootRequestWithDefaults()

	resp, err := client.Secrets.PKIReplaceRoot(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIReplaceRootRequest** | [**PKIReplaceRootRequest**](PKIReplaceRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIRevoke

> PKIRevoke(ctx, pkiMountPath).PKIRevokeRequest(pKIRevokeRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIRevokeRequestWithDefaults()

	resp, err := client.Secrets.PKIRevoke(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIRevokeRequest** | [**PKIRevokeRequest**](PKIRevokeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIRevokeWithKey

> PKIRevokeWithKey(ctx, pkiMountPath).PKIRevokeWithKeyRequest(pKIRevokeWithKeyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIRevokeWithKeyRequestWithDefaults()

	resp, err := client.Secrets.PKIRevokeWithKey(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIRevokeWithKeyRequest** | [**PKIRevokeWithKeyRequest**](PKIRevokeWithKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIRootSignIntermediate

> PKIRootSignIntermediate(ctx, pkiMountPath).PKIRootSignIntermediateRequest(pKIRootSignIntermediateRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIRootSignIntermediateRequestWithDefaults()

	resp, err := client.Secrets.PKIRootSignIntermediate(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIRootSignIntermediateRequest** | [**PKIRootSignIntermediateRequest**](PKIRootSignIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIRootSignSelfIssued

> PKIRootSignSelfIssued(ctx, pkiMountPath).PKIRootSignSelfIssuedRequest(pKIRootSignSelfIssuedRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIRootSignSelfIssuedRequestWithDefaults()

	resp, err := client.Secrets.PKIRootSignSelfIssued(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIRootSignSelfIssuedRequest** | [**PKIRootSignSelfIssuedRequest**](PKIRootSignSelfIssuedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIRotateRoot

> PKIRotateRoot(ctx, exported, pkiMountPath).PKIRotateRootRequest(pKIRotateRootRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!

	request := schema.NewPKIRotateRootRequestWithDefaults()

	resp, err := client.Secrets.PKIRotateRoot(
		context.Background(),
		exported,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIRotateRootRequest** | [**PKIRotateRootRequest**](PKIRotateRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKISignRole

> PKISignRole(ctx, pkiMountPath, role).PKISignRoleRequest(pKISignRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKISignRoleRequestWithDefaults()

	resp, err := client.Secrets.PKISignRole(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKISignRoleRequest** | [**PKISignRoleRequest**](PKISignRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKISignVerbatim

> PKISignVerbatim(ctx, pkiMountPath).PKISignVerbatimRequest(pKISignVerbatimRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKISignVerbatimRequestWithDefaults()

	resp, err := client.Secrets.PKISignVerbatim(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKISignVerbatimRequest** | [**PKISignVerbatimRequest**](PKISignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKISignVerbatimRole

> PKISignVerbatimRole(ctx, pkiMountPath, role).PKISignVerbatimRoleRequest(pKISignVerbatimRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKISignVerbatimRoleRequestWithDefaults()

	resp, err := client.Secrets.PKISignVerbatimRole(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKISignVerbatimRoleRequest** | [**PKISignVerbatimRoleRequest**](PKISignVerbatimRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKITidy

> PKITidy(ctx, pkiMountPath).PKITidyRequest(pKITidyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKITidyRequestWithDefaults()

	resp, err := client.Secrets.PKITidy(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKITidyRequest** | [**PKITidyRequest**](PKITidyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKITidyCancel

> PKITidyCancel(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKITidyCancel(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKITidyStatus

> PKITidyStatus(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKITidyStatus(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteAutoTidyConfig

> PKIWriteAutoTidyConfig(ctx, pkiMountPath).PKIWriteAutoTidyConfigRequest(pKIWriteAutoTidyConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteAutoTidyConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteAutoTidyConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteAutoTidyConfigRequest** | [**PKIWriteAutoTidyConfigRequest**](PKIWriteAutoTidyConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteCAConfig

> PKIWriteCAConfig(ctx, pkiMountPath).PKIWriteCAConfigRequest(pKIWriteCAConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteCAConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteCAConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteCAConfigRequest** | [**PKIWriteCAConfigRequest**](PKIWriteCAConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteCRLConfig

> PKIWriteCRLConfig(ctx, pkiMountPath).PKIWriteCRLConfigRequest(pKIWriteCRLConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteCRLConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteCRLConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteCRLConfigRequest** | [**PKIWriteCRLConfigRequest**](PKIWriteCRLConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteCerts

> PKIWriteCerts(ctx, pkiMountPath).PKIWriteCertsRequest(pKIWriteCertsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteCertsRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteCerts(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteCertsRequest** | [**PKIWriteCertsRequest**](PKIWriteCertsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteClusterConfig

> PKIWriteClusterConfig(ctx, pkiMountPath).PKIWriteClusterConfigRequest(pKIWriteClusterConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteClusterConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteClusterConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteClusterConfigRequest** | [**PKIWriteClusterConfigRequest**](PKIWriteClusterConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteIntermediateCrossSign

> PKIWriteIntermediateCrossSign(ctx, pkiMountPath).PKIWriteIntermediateCrossSignRequest(pKIWriteIntermediateCrossSignRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteIntermediateCrossSignRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteIntermediateCrossSign(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteIntermediateCrossSignRequest** | [**PKIWriteIntermediateCrossSignRequest**](PKIWriteIntermediateCrossSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteIntermediateGenerate

> PKIWriteIntermediateGenerate(ctx, exported, pkiMountPath).PKIWriteIntermediateGenerateRequest(pKIWriteIntermediateGenerateRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!

	request := schema.NewPKIWriteIntermediateGenerateRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteIntermediateGenerate(
		context.Background(),
		exported,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIWriteIntermediateGenerateRequest** | [**PKIWriteIntermediateGenerateRequest**](PKIWriteIntermediateGenerateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteIntermediateSetSigned

> PKIWriteIntermediateSetSigned(ctx, pkiMountPath).PKIWriteIntermediateSetSignedRequest(pKIWriteIntermediateSetSignedRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteIntermediateSetSignedRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteIntermediateSetSigned(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteIntermediateSetSignedRequest** | [**PKIWriteIntermediateSetSignedRequest**](PKIWriteIntermediateSetSignedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteInternalExported

> PKIWriteInternalExported(ctx, pkiMountPath).PKIWriteInternalExportedRequest(pKIWriteInternalExportedRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteInternalExportedRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteInternalExported(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteInternalExportedRequest** | [**PKIWriteInternalExportedRequest**](PKIWriteInternalExportedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteIssueRole

> PKIWriteIssueRole(ctx, pkiMountPath, role).PKIWriteIssueRoleRequest(pKIWriteIssueRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request

	request := schema.NewPKIWriteIssueRoleRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteIssueRole(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIWriteIssueRoleRequest** | [**PKIWriteIssueRoleRequest**](PKIWriteIssueRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteIssuersConfig

> PKIWriteIssuersConfig(ctx, pkiMountPath).PKIWriteIssuersConfigRequest(pKIWriteIssuersConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteIssuersConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteIssuersConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteIssuersConfigRequest** | [**PKIWriteIssuersConfigRequest**](PKIWriteIssuersConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteKMS

> PKIWriteKMS(ctx, pkiMountPath).PKIWriteKMSRequest(pKIWriteKMSRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteKMSRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteKMS(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteKMSRequest** | [**PKIWriteKMSRequest**](PKIWriteKMSRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteKey

> PKIWriteKey(ctx, keyRef, pkiMountPath).PKIWriteKeyRequest(pKIWriteKeyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")

	request := schema.NewPKIWriteKeyRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteKey(
		context.Background(),
		keyRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIWriteKeyRequest** | [**PKIWriteKeyRequest**](PKIWriteKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteKeysConfig

> PKIWriteKeysConfig(ctx, pkiMountPath).PKIWriteKeysConfigRequest(pKIWriteKeysConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteKeysConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteKeysConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteKeysConfigRequest** | [**PKIWriteKeysConfigRequest**](PKIWriteKeysConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteOCSP

> PKIWriteOCSP(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PKIWriteOCSP(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteRole

> PKIWriteRole(ctx, name, pkiMountPath).PKIWriteRoleRequest(pKIWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewPKIWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pKIWriteRoleRequest** | [**PKIWriteRoleRequest**](PKIWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PKIWriteURLConfig

> PKIWriteURLConfig(ctx, pkiMountPath).PKIWriteURLConfigRequest(pKIWriteURLConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPKIWriteURLConfigRequestWithDefaults()

	resp, err := client.Secrets.PKIWriteURLConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIWriteURLConfigRequest** | [**PKIWriteURLConfigRequest**](PKIWriteURLConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteIssuerRefDerPem

> PkiDeleteIssuerRefDerPem(ctx, issuerRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")


	resp, err := client.Secrets.PkiDeleteIssuerRefDerPem(
		context.Background(),
		issuerRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiDeleteJson

> PkiDeleteJson(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiDeleteJson(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadDelta

> PkiReadDelta(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiReadDelta(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadDeltaPem

> PkiReadDeltaPem(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiReadDeltaPem(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadDer

> PkiReadDer(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiReadDer(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuerRefCrlPemDerDeltaPem

> PkiReadIssuerRefCrlPemDerDeltaPem(ctx, issuerRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")


	resp, err := client.Secrets.PkiReadIssuerRefCrlPemDerDeltaPem(
		context.Background(),
		issuerRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadIssuerRefDerPem

> PkiReadIssuerRefDerPem(ctx, issuerRef, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")


	resp, err := client.Secrets.PkiReadIssuerRefDerPem(
		context.Background(),
		issuerRef,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadJson

> PkiReadJson(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiReadJson(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiReadPem

> PkiReadPem(ctx, pkiMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.PkiReadPem(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiWriteIssuerRefDerPem

> PkiWriteIssuerRefDerPem(ctx, issuerRef, pkiMountPath).PkiWriteIssuerRefDerPemRequest(pkiWriteIssuerRefDerPemRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")

	request := schema.NewPkiWriteIssuerRefDerPemRequestWithDefaults()

	resp, err := client.Secrets.PkiWriteIssuerRefDerPem(
		context.Background(),
		issuerRef,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiWriteIssuerRefDerPemRequest** | [**PkiWriteIssuerRefDerPemRequest**](PkiWriteIssuerRefDerPemRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PkiWriteJson

> PkiWriteJson(ctx, pkiMountPath).PkiWriteJsonRequest(pkiWriteJsonRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewPkiWriteJsonRequestWithDefaults()

	resp, err := client.Secrets.PkiWriteJson(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiWriteJsonRequest** | [**PkiWriteJsonRequest**](PkiWriteJsonRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQDeleteRole

> RabbitMQDeleteRole(ctx, name, rabbitmqMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.RabbitMQDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQListRoles

> RabbitMQListRoles(ctx, rabbitmqMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.RabbitMQListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQReadCredentials

> RabbitMQReadCredentials(ctx, name, rabbitmqMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.RabbitMQReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQReadLeaseConfig

> RabbitMQReadLeaseConfig(ctx, rabbitmqMountPath).Execute()

Configure the lease parameters for generated credentials

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.RabbitMQReadLeaseConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQReadRole

> RabbitMQReadRole(ctx, name, rabbitmqMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Secrets.RabbitMQReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQWriteConnectionConfig

> RabbitMQWriteConnectionConfig(ctx, rabbitmqMountPath).RabbitMQWriteConnectionConfigRequest(rabbitMQWriteConnectionConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewRabbitMQWriteConnectionConfigRequestWithDefaults()

	resp, err := client.Secrets.RabbitMQWriteConnectionConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitMQWriteConnectionConfigRequest** | [**RabbitMQWriteConnectionConfigRequest**](RabbitMQWriteConnectionConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQWriteLeaseConfig

> RabbitMQWriteLeaseConfig(ctx, rabbitmqMountPath).RabbitMQWriteLeaseConfigRequest(rabbitMQWriteLeaseConfigRequest).Execute()

Configure the lease parameters for generated credentials

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewRabbitMQWriteLeaseConfigRequestWithDefaults()

	resp, err := client.Secrets.RabbitMQWriteLeaseConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitMQWriteLeaseConfigRequest** | [**RabbitMQWriteLeaseConfigRequest**](RabbitMQWriteLeaseConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RabbitMQWriteRole

> RabbitMQWriteRole(ctx, name, rabbitmqMountPath).RabbitMQWriteRoleRequest(rabbitMQWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewRabbitMQWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.RabbitMQWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitMQWriteRoleRequest** | [**RabbitMQWriteRoleRequest**](RabbitMQWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHDeleteCAConfig

> SSHDeleteCAConfig(ctx, sshMountPath).Execute()

Set the SSH private key used for signing certificates.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHDeleteCAConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHDeleteKeys

> SSHDeleteKeys(ctx, keyName, sshMountPath).Execute()

Register a shared private key with Vault.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyName := "keyName_example" // string | [Required] Name of the key


	resp, err := client.Secrets.SSHDeleteKeys(
		context.Background(),
		keyName,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**keyName** | **string** | [Required] Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHDeleteRole

> SSHDeleteRole(ctx, role, sshMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.


	resp, err := client.Secrets.SSHDeleteRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHDeleteZeroAddressConfig

> SSHDeleteZeroAddressConfig(ctx, sshMountPath).Execute()

Assign zero address as default CIDR block for select roles.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHDeleteZeroAddressConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHListRoles

> SSHListRoles(ctx, sshMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHLookup

> SSHLookup(ctx, sshMountPath).SSHLookupRequest(sSHLookupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewSSHLookupRequestWithDefaults()

	resp, err := client.Secrets.SSHLookup(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHLookupRequest** | [**SSHLookupRequest**](SSHLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHReadCAConfig

> SSHReadCAConfig(ctx, sshMountPath).Execute()

Set the SSH private key used for signing certificates.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHReadCAConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHReadPublicKey

> SSHReadPublicKey(ctx, sshMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHReadPublicKey(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHReadRole

> SSHReadRole(ctx, role, sshMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.


	resp, err := client.Secrets.SSHReadRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHReadZeroAddressConfig

> SSHReadZeroAddressConfig(ctx, sshMountPath).Execute()

Assign zero address as default CIDR block for select roles.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.SSHReadZeroAddressConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHSign

> SSHSign(ctx, role, sshMountPath).SSHSignRequest(sSHSignRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.

	request := schema.NewSSHSignRequestWithDefaults()

	resp, err := client.Secrets.SSHSign(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHSignRequest** | [**SSHSignRequest**](SSHSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHVerify

> SSHVerify(ctx, sshMountPath).SSHVerifyRequest(sSHVerifyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewSSHVerifyRequestWithDefaults()

	resp, err := client.Secrets.SSHVerify(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHVerifyRequest** | [**SSHVerifyRequest**](SSHVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteCAConfig

> SSHWriteCAConfig(ctx, sshMountPath).SSHWriteCAConfigRequest(sSHWriteCAConfigRequest).Execute()

Set the SSH private key used for signing certificates.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewSSHWriteCAConfigRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteCAConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHWriteCAConfigRequest** | [**SSHWriteCAConfigRequest**](SSHWriteCAConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteCredentials

> SSHWriteCredentials(ctx, role, sshMountPath).SSHWriteCredentialsRequest(sSHWriteCredentialsRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required] Name of the role

	request := schema.NewSSHWriteCredentialsRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteCredentials(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHWriteCredentialsRequest** | [**SSHWriteCredentialsRequest**](SSHWriteCredentialsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteIssue

> SSHWriteIssue(ctx, role, sshMountPath).SSHWriteIssueRequest(sSHWriteIssueRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.

	request := schema.NewSSHWriteIssueRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteIssue(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHWriteIssueRequest** | [**SSHWriteIssueRequest**](SSHWriteIssueRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteKeys

> SSHWriteKeys(ctx, keyName, sshMountPath).SSHWriteKeysRequest(sSHWriteKeysRequest).Execute()

Register a shared private key with Vault.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	keyName := "keyName_example" // string | [Required] Name of the key

	request := schema.NewSSHWriteKeysRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteKeys(
		context.Background(),
		keyName,
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
**keyName** | **string** | [Required] Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHWriteKeysRequest** | [**SSHWriteKeysRequest**](SSHWriteKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteRole

> SSHWriteRole(ctx, role, sshMountPath).SSHWriteRoleRequest(sSHWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.

	request := schema.NewSSHWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteRole(
		context.Background(),
		role,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHWriteRoleRequest** | [**SSHWriteRoleRequest**](SSHWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SSHWriteZeroAddressConfig

> SSHWriteZeroAddressConfig(ctx, sshMountPath).SSHWriteZeroAddressConfigRequest(sSHWriteZeroAddressConfigRequest).Execute()

Assign zero address as default CIDR block for select roles.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewSSHWriteZeroAddressConfigRequestWithDefaults()

	resp, err := client.Secrets.SSHWriteZeroAddressConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHWriteZeroAddressConfigRequest** | [**SSHWriteZeroAddressConfigRequest**](SSHWriteZeroAddressConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPDeleteKey

> TOTPDeleteKey(ctx, name, totpMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.


	resp, err := client.Secrets.TOTPDeleteKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPListKeys

> TOTPListKeys(ctx, totpMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TOTPListKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPReadCode

> TOTPReadCode(ctx, name, totpMountPath).Execute()

Request time-based one-time use password or validate a password for a certain key .

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.


	resp, err := client.Secrets.TOTPReadCode(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPReadKey

> TOTPReadKey(ctx, name, totpMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.


	resp, err := client.Secrets.TOTPReadKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPWriteCode

> TOTPWriteCode(ctx, name, totpMountPath).TOTPWriteCodeRequest(tOTPWriteCodeRequest).Execute()

Request time-based one-time use password or validate a password for a certain key .

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.

	request := schema.NewTOTPWriteCodeRequestWithDefaults()

	resp, err := client.Secrets.TOTPWriteCode(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tOTPWriteCodeRequest** | [**TOTPWriteCodeRequest**](TOTPWriteCodeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TOTPWriteKey

> TOTPWriteKey(ctx, name, totpMountPath).TOTPWriteKeyRequest(tOTPWriteKeyRequest).Execute()

Manage the keys that can be created with this backend.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.

	request := schema.NewTOTPWriteKeyRequestWithDefaults()

	resp, err := client.Secrets.TOTPWriteKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tOTPWriteKeyRequest** | [**TOTPWriteKeyRequest**](TOTPWriteKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformDeleteConfig

> TerraformDeleteConfig(ctx, terraformMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TerraformDeleteConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformDeleteRole

> TerraformDeleteRole(ctx, name, terraformMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.TerraformDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformListRoles

> TerraformListRoles(ctx, terraformMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TerraformListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformReadConfig

> TerraformReadConfig(ctx, terraformMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TerraformReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformReadCredentials

> TerraformReadCredentials(ctx, name, terraformMountPath).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.TerraformReadCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformReadRole

> TerraformReadRole(ctx, name, terraformMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.TerraformReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformRotateRole

> TerraformRotateRole(ctx, name, terraformMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the team or organization role


	resp, err := client.Secrets.TerraformRotateRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformWriteConfig

> TerraformWriteConfig(ctx, terraformMountPath).TerraformWriteConfigRequest(terraformWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTerraformWriteConfigRequestWithDefaults()

	resp, err := client.Secrets.TerraformWriteConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformWriteConfigRequest** | [**TerraformWriteConfigRequest**](TerraformWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformWriteCredentials

> TerraformWriteCredentials(ctx, name, terraformMountPath).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role


	resp, err := client.Secrets.TerraformWriteCredentials(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TerraformWriteRole

> TerraformWriteRole(ctx, name, terraformMountPath).TerraformWriteRoleRequest(terraformWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role

	request := schema.NewTerraformWriteRoleRequestWithDefaults()

	resp, err := client.Secrets.TerraformWriteRole(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformWriteRoleRequest** | [**TerraformWriteRoleRequest**](TerraformWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitBackup

> TransitBackup(ctx, name, transitMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key


	resp, err := client.Secrets.TransitBackup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitDecrypt

> TransitDecrypt(ctx, name, transitMountPath).TransitDecryptRequest(transitDecryptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitDecryptRequestWithDefaults()

	resp, err := client.Secrets.TransitDecrypt(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitDeleteKey

> TransitDeleteKey(ctx, name, transitMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key


	resp, err := client.Secrets.TransitDeleteKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitEncrypt

> TransitEncrypt(ctx, name, transitMountPath).TransitEncryptRequest(transitEncryptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitEncryptRequestWithDefaults()

	resp, err := client.Secrets.TransitEncrypt(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitExport

> TransitExport(ctx, name, transitMountPath, type_).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)


	resp, err := client.Secrets.TransitExport(
		context.Background(),
		name,
		type_,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitExportVersion

> TransitExportVersion(ctx, name, transitMountPath, type_, version).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
	version := "version_example" // string | Version of the key


	resp, err := client.Secrets.TransitExportVersion(
		context.Background(),
		name,
		type_,
		version,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateDataKey

> TransitGenerateDataKey(ctx, name, plaintext, transitMountPath).TransitGenerateDataKeyRequest(transitGenerateDataKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The backend key used for encrypting the data key
	plaintext := "plaintext_example" // string | \"plaintext\" will return the key in both plaintext and ciphertext; \"wrapped\" will return the ciphertext only.

	request := schema.NewTransitGenerateDataKeyRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateDataKey(
		context.Background(),
		name,
		plaintext,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateDataKeyRequest** | [**TransitGenerateDataKeyRequest**](TransitGenerateDataKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateHMAC

> TransitGenerateHMAC(ctx, name, transitMountPath).TransitGenerateHMACRequest(transitGenerateHMACRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function

	request := schema.NewTransitGenerateHMACRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateHMAC(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGenerateHMACRequest** | [**TransitGenerateHMACRequest**](TransitGenerateHMACRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateHMACWithAlgorithm

> TransitGenerateHMACWithAlgorithm(ctx, name, transitMountPath, urlalgorithm).TransitGenerateHMACWithAlgorithmRequest(transitGenerateHMACWithAlgorithmRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function
	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)

	request := schema.NewTransitGenerateHMACWithAlgorithmRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateHMACWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateHMACWithAlgorithmRequest** | [**TransitGenerateHMACWithAlgorithmRequest**](TransitGenerateHMACWithAlgorithmRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandom

> TransitGenerateRandom(ctx, transitMountPath).TransitGenerateRandomRequest(transitGenerateRandomRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTransitGenerateRandomRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateRandom(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitGenerateRandomRequest** | [**TransitGenerateRandomRequest**](TransitGenerateRandomRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandomSource

> TransitGenerateRandomSource(ctx, source, transitMountPath).TransitGenerateRandomSourceRequest(transitGenerateRandomSourceRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")

	request := schema.NewTransitGenerateRandomSourceRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateRandomSource(
		context.Background(),
		source,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGenerateRandomSourceRequest** | [**TransitGenerateRandomSourceRequest**](TransitGenerateRandomSourceRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitGenerateRandomSourceBytes

> TransitGenerateRandomSourceBytes(ctx, source, transitMountPath, urlbytes).TransitGenerateRandomSourceBytesRequest(transitGenerateRandomSourceBytesRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)

	request := schema.NewTransitGenerateRandomSourceBytesRequestWithDefaults()

	resp, err := client.Secrets.TransitGenerateRandomSourceBytes(
		context.Background(),
		source,
		urlbytes,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGenerateRandomSourceBytesRequest** | [**TransitGenerateRandomSourceBytesRequest**](TransitGenerateRandomSourceBytesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitHash

> TransitHash(ctx, transitMountPath).TransitHashRequest(transitHashRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTransitHashRequestWithDefaults()

	resp, err := client.Secrets.TransitHash(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitHashWithAlgorithm

> TransitHashWithAlgorithm(ctx, transitMountPath, urlalgorithm).TransitHashWithAlgorithmRequest(transitHashWithAlgorithmRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)

	request := schema.NewTransitHashWithAlgorithmRequestWithDefaults()

	resp, err := client.Secrets.TransitHashWithAlgorithm(
		context.Background(),
		urlalgorithm,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHashWithAlgorithmRequest** | [**TransitHashWithAlgorithmRequest**](TransitHashWithAlgorithmRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitImportKey

> TransitImportKey(ctx, name, transitMountPath).TransitImportKeyRequest(transitImportKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key

	request := schema.NewTransitImportKeyRequestWithDefaults()

	resp, err := client.Secrets.TransitImportKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitImportKeyRequest** | [**TransitImportKeyRequest**](TransitImportKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitImportKeyVersion

> TransitImportKeyVersion(ctx, name, transitMountPath).TransitImportKeyVersionRequest(transitImportKeyVersionRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key

	request := schema.NewTransitImportKeyVersionRequestWithDefaults()

	resp, err := client.Secrets.TransitImportKeyVersion(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitImportKeyVersionRequest** | [**TransitImportKeyVersionRequest**](TransitImportKeyVersionRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitListKeys

> TransitListKeys(ctx, transitMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TransitListKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadCacheConfig

> TransitReadCacheConfig(ctx, transitMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TransitReadCacheConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadConfigKeys

> TransitReadConfigKeys(ctx, transitMountPath).Execute()

Configuration common across all keys

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TransitReadConfigKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadKey

> TransitReadKey(ctx, name, transitMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key


	resp, err := client.Secrets.TransitReadKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitReadWrappingKey

> TransitReadWrappingKey(ctx, transitMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Secrets.TransitReadWrappingKey(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRestore

> TransitRestore(ctx, transitMountPath).TransitRestoreRequest(transitRestoreRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTransitRestoreRequestWithDefaults()

	resp, err := client.Secrets.TransitRestore(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRestoreKey

> TransitRestoreKey(ctx, name, transitMountPath).TransitRestoreKeyRequest(transitRestoreKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | If set, this will be the name of the restored key.

	request := schema.NewTransitRestoreKeyRequestWithDefaults()

	resp, err := client.Secrets.TransitRestoreKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRestoreKeyRequest** | [**TransitRestoreKeyRequest**](TransitRestoreKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRewrap

> TransitRewrap(ctx, name, transitMountPath).TransitRewrapRequest(transitRewrapRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitRewrapRequestWithDefaults()

	resp, err := client.Secrets.TransitRewrap(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRewrapRequest** | [**TransitRewrapRequest**](TransitRewrapRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitRotateKey

> TransitRotateKey(ctx, name, transitMountPath).Execute()

Rotate named encryption key

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key


	resp, err := client.Secrets.TransitRotateKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitSign

> TransitSign(ctx, name, transitMountPath).TransitSignRequest(transitSignRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use

	request := schema.NewTransitSignRequestWithDefaults()

	resp, err := client.Secrets.TransitSign(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitSignWithAlgorithm

> TransitSignWithAlgorithm(ctx, name, transitMountPath, urlalgorithm).TransitSignWithAlgorithmRequest(transitSignWithAlgorithmRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)

	request := schema.NewTransitSignWithAlgorithmRequestWithDefaults()

	resp, err := client.Secrets.TransitSignWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitSignWithAlgorithmRequest** | [**TransitSignWithAlgorithmRequest**](TransitSignWithAlgorithmRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitTrimKey

> TransitTrimKey(ctx, name, transitMountPath).TransitTrimKeyRequest(transitTrimKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitTrimKeyRequestWithDefaults()

	resp, err := client.Secrets.TransitTrimKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitTrimKeyRequest** | [**TransitTrimKeyRequest**](TransitTrimKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitVerify

> TransitVerify(ctx, name, transitMountPath).TransitVerifyRequest(transitVerifyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use

	request := schema.NewTransitVerifyRequestWithDefaults()

	resp, err := client.Secrets.TransitVerify(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitVerifyWithAlgorithm

> TransitVerifyWithAlgorithm(ctx, name, transitMountPath, urlalgorithm).TransitVerifyWithAlgorithmRequest(transitVerifyWithAlgorithmRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)

	request := schema.NewTransitVerifyWithAlgorithmRequestWithDefaults()

	resp, err := client.Secrets.TransitVerifyWithAlgorithm(
		context.Background(),
		name,
		urlalgorithm,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitVerifyWithAlgorithmRequest** | [**TransitVerifyWithAlgorithmRequest**](TransitVerifyWithAlgorithmRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitWriteCacheConfig

> TransitWriteCacheConfig(ctx, transitMountPath).TransitWriteCacheConfigRequest(transitWriteCacheConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTransitWriteCacheConfigRequestWithDefaults()

	resp, err := client.Secrets.TransitWriteCacheConfig(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitWriteCacheConfigRequest** | [**TransitWriteCacheConfigRequest**](TransitWriteCacheConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitWriteConfigKeys

> TransitWriteConfigKeys(ctx, transitMountPath).TransitWriteConfigKeysRequest(transitWriteConfigKeysRequest).Execute()

Configuration common across all keys

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTransitWriteConfigKeysRequestWithDefaults()

	resp, err := client.Secrets.TransitWriteConfigKeys(
		context.Background(),
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitWriteConfigKeysRequest** | [**TransitWriteConfigKeysRequest**](TransitWriteConfigKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitWriteKey

> TransitWriteKey(ctx, name, transitMountPath).TransitWriteKeyRequest(transitWriteKeyRequest).Execute()

Managed named encryption keys

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitWriteKeyRequestWithDefaults()

	resp, err := client.Secrets.TransitWriteKey(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitWriteKeyRequest** | [**TransitWriteKeyRequest**](TransitWriteKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitWriteKeyConfig

> TransitWriteKeyConfig(ctx, name, transitMountPath).TransitWriteKeyConfigRequest(transitWriteKeyConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key

	request := schema.NewTransitWriteKeyConfigRequestWithDefaults()

	resp, err := client.Secrets.TransitWriteKeyConfig(
		context.Background(),
		name,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitWriteKeyConfigRequest** | [**TransitWriteKeyConfigRequest**](TransitWriteKeyConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TransitWriteRandomUrlbytes

> TransitWriteRandomUrlbytes(ctx, transitMountPath, urlbytes).TransitWriteRandomUrlbytesRequest(transitWriteRandomUrlbytesRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)

	request := schema.NewTransitWriteRandomUrlbytesRequestWithDefaults()

	resp, err := client.Secrets.TransitWriteRandomUrlbytes(
		context.Background(),
		urlbytes,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitWriteRandomUrlbytesRequest** | [**TransitWriteRandomUrlbytesRequest**](TransitWriteRandomUrlbytesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



