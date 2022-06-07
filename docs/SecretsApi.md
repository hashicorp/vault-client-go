# \SecretsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdConfig**](SecretsApi.md#DeleteAdConfig) | **Delete** /ad/config | Configure the AD server to connect to, along with password options.
[**DeleteAdLibraryName**](SecretsApi.md#DeleteAdLibraryName) | **Delete** /ad/library/{name} | Delete a library set.
[**DeleteAdRolesName**](SecretsApi.md#DeleteAdRolesName) | **Delete** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**DeleteAlicloudConfig**](SecretsApi.md#DeleteAlicloudConfig) | **Delete** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**DeleteAlicloudRoleName**](SecretsApi.md#DeleteAlicloudRoleName) | **Delete** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**DeleteAwsRolesName**](SecretsApi.md#DeleteAwsRolesName) | **Delete** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**DeleteAzureConfig**](SecretsApi.md#DeleteAzureConfig) | **Delete** /azure/config | 
[**DeleteAzureRolesName**](SecretsApi.md#DeleteAzureRolesName) | **Delete** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**DeleteConsulRolesName**](SecretsApi.md#DeleteConsulRolesName) | **Delete** /consul/roles/{name} | 
[**DeleteCubbyholePath**](SecretsApi.md#DeleteCubbyholePath) | **Delete** /cubbyhole/{path} | Deletes the secret at the specified location.
[**DeleteGcpRolesetName**](SecretsApi.md#DeleteGcpRolesetName) | **Delete** /gcp/roleset/{name} | 
[**DeleteGcpStaticAccountName**](SecretsApi.md#DeleteGcpStaticAccountName) | **Delete** /gcp/static-account/{name} | 
[**DeleteGcpkmsConfig**](SecretsApi.md#DeleteGcpkmsConfig) | **Delete** /gcpkms/config | Configure the GCP KMS secrets engine
[**DeleteGcpkmsKeysDeregisterKey**](SecretsApi.md#DeleteGcpkmsKeysDeregisterKey) | **Delete** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
[**DeleteGcpkmsKeysKey**](SecretsApi.md#DeleteGcpkmsKeysKey) | **Delete** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**DeleteGcpkmsKeysTrimKey**](SecretsApi.md#DeleteGcpkmsKeysTrimKey) | **Delete** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**DeleteKubernetesConfig**](SecretsApi.md#DeleteKubernetesConfig) | **Delete** /kubernetes/config | 
[**DeleteKubernetesRolesName**](SecretsApi.md#DeleteKubernetesRolesName) | **Delete** /kubernetes/roles/{name} | 
[**DeleteKvPath**](SecretsApi.md#DeleteKvPath) | **Delete** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**DeleteMongodbatlasRolesName**](SecretsApi.md#DeleteMongodbatlasRolesName) | **Delete** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**DeleteNomadConfigAccess**](SecretsApi.md#DeleteNomadConfigAccess) | **Delete** /nomad/config/access | 
[**DeleteNomadConfigLease**](SecretsApi.md#DeleteNomadConfigLease) | **Delete** /nomad/config/lease | Configure the lease parameters for generated tokens
[**DeleteNomadRoleName**](SecretsApi.md#DeleteNomadRoleName) | **Delete** /nomad/role/{name} | 
[**DeleteOpenldapConfig**](SecretsApi.md#DeleteOpenldapConfig) | **Delete** /openldap/config | 
[**DeleteOpenldapRoleName**](SecretsApi.md#DeleteOpenldapRoleName) | **Delete** /openldap/role/{name} | 
[**DeleteOpenldapStaticRoleName**](SecretsApi.md#DeleteOpenldapStaticRoleName) | **Delete** /openldap/static-role/{name} | 
[**DeletePkiIssuerRefDerPem**](SecretsApi.md#DeletePkiIssuerRefDerPem) | **Delete** /pki/{issuer_ref}/der|/pem | 
[**DeletePkiJson**](SecretsApi.md#DeletePkiJson) | **Delete** /pki//json | 
[**DeletePkiKeyKeyRef**](SecretsApi.md#DeletePkiKeyKeyRef) | **Delete** /pki/key/{key_ref} | 
[**DeletePkiRolesName**](SecretsApi.md#DeletePkiRolesName) | **Delete** /pki/roles/{name} | 
[**DeletePkiRoot**](SecretsApi.md#DeletePkiRoot) | **Delete** /pki/root | 
[**DeleteRabbitmqRolesName**](SecretsApi.md#DeleteRabbitmqRolesName) | **Delete** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**DeleteSecretDataPath**](SecretsApi.md#DeleteSecretDataPath) | **Delete** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**DeleteSecretMetadataPath**](SecretsApi.md#DeleteSecretMetadataPath) | **Delete** /secret/metadata/{path} | Configures settings for the KV store
[**DeleteSshConfigCa**](SecretsApi.md#DeleteSshConfigCa) | **Delete** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**DeleteSshConfigZeroaddress**](SecretsApi.md#DeleteSshConfigZeroaddress) | **Delete** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**DeleteSshKeysKeyName**](SecretsApi.md#DeleteSshKeysKeyName) | **Delete** /ssh/keys/{key_name} | Register a shared private key with Vault.
[**DeleteSshRolesRole**](SecretsApi.md#DeleteSshRolesRole) | **Delete** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**DeleteTerraformConfig**](SecretsApi.md#DeleteTerraformConfig) | **Delete** /terraform/config | 
[**DeleteTerraformRoleName**](SecretsApi.md#DeleteTerraformRoleName) | **Delete** /terraform/role/{name} | 
[**DeleteTotpKeysName**](SecretsApi.md#DeleteTotpKeysName) | **Delete** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**DeleteTransitKeysName**](SecretsApi.md#DeleteTransitKeysName) | **Delete** /transit/keys/{name} | Managed named encryption keys
[**GetAdConfig**](SecretsApi.md#GetAdConfig) | **Get** /ad/config | Configure the AD server to connect to, along with password options.
[**GetAdCredsName**](SecretsApi.md#GetAdCredsName) | **Get** /ad/creds/{name} | 
[**GetAdLibrary**](SecretsApi.md#GetAdLibrary) | **Get** /ad/library | 
[**GetAdLibraryName**](SecretsApi.md#GetAdLibraryName) | **Get** /ad/library/{name} | Read a library set.
[**GetAdLibraryNameStatus**](SecretsApi.md#GetAdLibraryNameStatus) | **Get** /ad/library/{name}/status | Check the status of the service accounts in a library set.
[**GetAdRoles**](SecretsApi.md#GetAdRoles) | **Get** /ad/roles | List the name of each role currently stored.
[**GetAdRolesName**](SecretsApi.md#GetAdRolesName) | **Get** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**GetAdRotateRoot**](SecretsApi.md#GetAdRotateRoot) | **Get** /ad/rotate-root | 
[**GetAlicloudConfig**](SecretsApi.md#GetAlicloudConfig) | **Get** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**GetAlicloudCredsName**](SecretsApi.md#GetAlicloudCredsName) | **Get** /alicloud/creds/{name} | Generate an API key or STS credential using the given role&#39;s configuration.&#39;
[**GetAlicloudRole**](SecretsApi.md#GetAlicloudRole) | **Get** /alicloud/role | List the existing roles in this backend.
[**GetAlicloudRoleName**](SecretsApi.md#GetAlicloudRoleName) | **Get** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**GetAwsConfigLease**](SecretsApi.md#GetAwsConfigLease) | **Get** /aws/config/lease | Configure the default lease information for generated credentials.
[**GetAwsConfigRoot**](SecretsApi.md#GetAwsConfigRoot) | **Get** /aws/config/root | Configure the root credentials that are used to manage IAM.
[**GetAwsCreds**](SecretsApi.md#GetAwsCreds) | **Get** /aws/creds | Generate AWS credentials from a specific Vault role.
[**GetAwsRoles**](SecretsApi.md#GetAwsRoles) | **Get** /aws/roles | List the existing roles in this backend
[**GetAwsRolesName**](SecretsApi.md#GetAwsRolesName) | **Get** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**GetAwsStsName**](SecretsApi.md#GetAwsStsName) | **Get** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
[**GetAzureConfig**](SecretsApi.md#GetAzureConfig) | **Get** /azure/config | 
[**GetAzureCredsRole**](SecretsApi.md#GetAzureCredsRole) | **Get** /azure/creds/{role} | 
[**GetAzureRoles**](SecretsApi.md#GetAzureRoles) | **Get** /azure/roles | List existing roles.
[**GetAzureRolesName**](SecretsApi.md#GetAzureRolesName) | **Get** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**GetConsulConfigAccess**](SecretsApi.md#GetConsulConfigAccess) | **Get** /consul/config/access | 
[**GetConsulCredsRole**](SecretsApi.md#GetConsulCredsRole) | **Get** /consul/creds/{role} | 
[**GetConsulRoles**](SecretsApi.md#GetConsulRoles) | **Get** /consul/roles | 
[**GetConsulRolesName**](SecretsApi.md#GetConsulRolesName) | **Get** /consul/roles/{name} | 
[**GetCubbyholePath**](SecretsApi.md#GetCubbyholePath) | **Get** /cubbyhole/{path} | Retrieve the secret at the specified location.
[**GetGcpConfig**](SecretsApi.md#GetGcpConfig) | **Get** /gcp/config | 
[**GetGcpKeyRoleset**](SecretsApi.md#GetGcpKeyRoleset) | **Get** /gcp/key/{roleset} | 
[**GetGcpRolesetName**](SecretsApi.md#GetGcpRolesetName) | **Get** /gcp/roleset/{name} | 
[**GetGcpRolesetRolesetKey**](SecretsApi.md#GetGcpRolesetRolesetKey) | **Get** /gcp/roleset/{roleset}/key | 
[**GetGcpRolesetRolesetToken**](SecretsApi.md#GetGcpRolesetRolesetToken) | **Get** /gcp/roleset/{roleset}/token | 
[**GetGcpRolesets**](SecretsApi.md#GetGcpRolesets) | **Get** /gcp/rolesets | 
[**GetGcpStaticAccountName**](SecretsApi.md#GetGcpStaticAccountName) | **Get** /gcp/static-account/{name} | 
[**GetGcpStaticAccountNameKey**](SecretsApi.md#GetGcpStaticAccountNameKey) | **Get** /gcp/static-account/{name}/key | 
[**GetGcpStaticAccountNameToken**](SecretsApi.md#GetGcpStaticAccountNameToken) | **Get** /gcp/static-account/{name}/token | 
[**GetGcpStaticAccounts**](SecretsApi.md#GetGcpStaticAccounts) | **Get** /gcp/static-accounts | 
[**GetGcpTokenRoleset**](SecretsApi.md#GetGcpTokenRoleset) | **Get** /gcp/token/{roleset} | 
[**GetGcpkmsConfig**](SecretsApi.md#GetGcpkmsConfig) | **Get** /gcpkms/config | Configure the GCP KMS secrets engine
[**GetGcpkmsKeys**](SecretsApi.md#GetGcpkmsKeys) | **Get** /gcpkms/keys | List named keys
[**GetGcpkmsKeysConfigKey**](SecretsApi.md#GetGcpkmsKeysConfigKey) | **Get** /gcpkms/keys/config/{key} | Configure the key in Vault
[**GetGcpkmsKeysKey**](SecretsApi.md#GetGcpkmsKeysKey) | **Get** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GetGcpkmsPubkeyKey**](SecretsApi.md#GetGcpkmsPubkeyKey) | **Get** /gcpkms/pubkey/{key} | Retrieve the public key associated with the named key
[**GetKubernetesConfig**](SecretsApi.md#GetKubernetesConfig) | **Get** /kubernetes/config | 
[**GetKubernetesRoles**](SecretsApi.md#GetKubernetesRoles) | **Get** /kubernetes/roles | 
[**GetKubernetesRolesName**](SecretsApi.md#GetKubernetesRolesName) | **Get** /kubernetes/roles/{name} | 
[**GetKvPath**](SecretsApi.md#GetKvPath) | **Get** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**GetMongodbatlasConfig**](SecretsApi.md#GetMongodbatlasConfig) | **Get** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
[**GetMongodbatlasCredsName**](SecretsApi.md#GetMongodbatlasCredsName) | **Get** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**GetMongodbatlasRoles**](SecretsApi.md#GetMongodbatlasRoles) | **Get** /mongodbatlas/roles | List the existing roles in this backend
[**GetMongodbatlasRolesName**](SecretsApi.md#GetMongodbatlasRolesName) | **Get** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**GetNomadConfigAccess**](SecretsApi.md#GetNomadConfigAccess) | **Get** /nomad/config/access | 
[**GetNomadConfigLease**](SecretsApi.md#GetNomadConfigLease) | **Get** /nomad/config/lease | Configure the lease parameters for generated tokens
[**GetNomadCredsName**](SecretsApi.md#GetNomadCredsName) | **Get** /nomad/creds/{name} | 
[**GetNomadRole**](SecretsApi.md#GetNomadRole) | **Get** /nomad/role | 
[**GetNomadRoleName**](SecretsApi.md#GetNomadRoleName) | **Get** /nomad/role/{name} | 
[**GetOpenldapConfig**](SecretsApi.md#GetOpenldapConfig) | **Get** /openldap/config | 
[**GetOpenldapCredsName**](SecretsApi.md#GetOpenldapCredsName) | **Get** /openldap/creds/{name} | 
[**GetOpenldapRole**](SecretsApi.md#GetOpenldapRole) | **Get** /openldap/role | 
[**GetOpenldapRoleName**](SecretsApi.md#GetOpenldapRoleName) | **Get** /openldap/role/{name} | 
[**GetOpenldapStaticCredName**](SecretsApi.md#GetOpenldapStaticCredName) | **Get** /openldap/static-cred/{name} | 
[**GetOpenldapStaticRole**](SecretsApi.md#GetOpenldapStaticRole) | **Get** /openldap/static-role | 
[**GetOpenldapStaticRoleName**](SecretsApi.md#GetOpenldapStaticRoleName) | **Get** /openldap/static-role/{name} | 
[**GetPkiCa**](SecretsApi.md#GetPkiCa) | **Get** /pki/ca | 
[**GetPkiCaChain**](SecretsApi.md#GetPkiCaChain) | **Get** /pki/ca_chain | 
[**GetPkiCaPem**](SecretsApi.md#GetPkiCaPem) | **Get** /pki/ca/pem | 
[**GetPkiCertCaChain**](SecretsApi.md#GetPkiCertCaChain) | **Get** /pki/cert/ca_chain | 
[**GetPkiCertCrl**](SecretsApi.md#GetPkiCertCrl) | **Get** /pki/cert/crl | 
[**GetPkiCertSerial**](SecretsApi.md#GetPkiCertSerial) | **Get** /pki/cert/{serial} | 
[**GetPkiCertSerialRaw**](SecretsApi.md#GetPkiCertSerialRaw) | **Get** /pki/cert/{serial}/raw | 
[**GetPkiCertSerialRawPem**](SecretsApi.md#GetPkiCertSerialRawPem) | **Get** /pki/cert/{serial}/raw/pem | 
[**GetPkiCerts**](SecretsApi.md#GetPkiCerts) | **Get** /pki/certs | 
[**GetPkiConfigCrl**](SecretsApi.md#GetPkiConfigCrl) | **Get** /pki/config/crl | 
[**GetPkiConfigIssuers**](SecretsApi.md#GetPkiConfigIssuers) | **Get** /pki/config/issuers | 
[**GetPkiConfigKeys**](SecretsApi.md#GetPkiConfigKeys) | **Get** /pki/config/keys | 
[**GetPkiConfigUrls**](SecretsApi.md#GetPkiConfigUrls) | **Get** /pki/config/urls | 
[**GetPkiCrl**](SecretsApi.md#GetPkiCrl) | **Get** /pki/crl | 
[**GetPkiCrlPem**](SecretsApi.md#GetPkiCrlPem) | **Get** /pki/crl/pem | 
[**GetPkiCrlRotate**](SecretsApi.md#GetPkiCrlRotate) | **Get** /pki/crl/rotate | 
[**GetPkiDer**](SecretsApi.md#GetPkiDer) | **Get** /pki//der | 
[**GetPkiIssuerRefCrlPem**](SecretsApi.md#GetPkiIssuerRefCrlPem) | **Get** /pki/{issuer_ref}/crl/pem | 
[**GetPkiIssuerRefDerPem**](SecretsApi.md#GetPkiIssuerRefDerPem) | **Get** /pki/{issuer_ref}/der|/pem | 
[**GetPkiIssuers**](SecretsApi.md#GetPkiIssuers) | **Get** /pki/issuers | 
[**GetPkiJson**](SecretsApi.md#GetPkiJson) | **Get** /pki//json | 
[**GetPkiKeyKeyRef**](SecretsApi.md#GetPkiKeyKeyRef) | **Get** /pki/key/{key_ref} | 
[**GetPkiKeys**](SecretsApi.md#GetPkiKeys) | **Get** /pki/keys | 
[**GetPkiRoles**](SecretsApi.md#GetPkiRoles) | **Get** /pki/roles | 
[**GetPkiRolesName**](SecretsApi.md#GetPkiRolesName) | **Get** /pki/roles/{name} | 
[**GetPkiTidyStatus**](SecretsApi.md#GetPkiTidyStatus) | **Get** /pki/tidy-status | 
[**GetRabbitmqConfigLease**](SecretsApi.md#GetRabbitmqConfigLease) | **Get** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
[**GetRabbitmqCredsName**](SecretsApi.md#GetRabbitmqCredsName) | **Get** /rabbitmq/creds/{name} | Request RabbitMQ credentials for a certain role.
[**GetRabbitmqRoles**](SecretsApi.md#GetRabbitmqRoles) | **Get** /rabbitmq/roles | Manage the roles that can be created with this backend.
[**GetRabbitmqRolesName**](SecretsApi.md#GetRabbitmqRolesName) | **Get** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**GetSecretConfig**](SecretsApi.md#GetSecretConfig) | **Get** /secret/config | Read the backend level settings.
[**GetSecretDataPath**](SecretsApi.md#GetSecretDataPath) | **Get** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**GetSecretMetadataPath**](SecretsApi.md#GetSecretMetadataPath) | **Get** /secret/metadata/{path} | Configures settings for the KV store
[**GetSecretSubkeysPath**](SecretsApi.md#GetSecretSubkeysPath) | **Get** /secret/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
[**GetSshConfigCa**](SecretsApi.md#GetSshConfigCa) | **Get** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**GetSshConfigZeroaddress**](SecretsApi.md#GetSshConfigZeroaddress) | **Get** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**GetSshPublicKey**](SecretsApi.md#GetSshPublicKey) | **Get** /ssh/public_key | Retrieve the public key.
[**GetSshRoles**](SecretsApi.md#GetSshRoles) | **Get** /ssh/roles | Manage the &#39;roles&#39; that can be created with this backend.
[**GetSshRolesRole**](SecretsApi.md#GetSshRolesRole) | **Get** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**GetTerraformConfig**](SecretsApi.md#GetTerraformConfig) | **Get** /terraform/config | 
[**GetTerraformCredsName**](SecretsApi.md#GetTerraformCredsName) | **Get** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**GetTerraformRole**](SecretsApi.md#GetTerraformRole) | **Get** /terraform/role | 
[**GetTerraformRoleName**](SecretsApi.md#GetTerraformRoleName) | **Get** /terraform/role/{name} | 
[**GetTotpCodeName**](SecretsApi.md#GetTotpCodeName) | **Get** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**GetTotpKeys**](SecretsApi.md#GetTotpKeys) | **Get** /totp/keys | Manage the keys that can be created with this backend.
[**GetTotpKeysName**](SecretsApi.md#GetTotpKeysName) | **Get** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**GetTransitBackupName**](SecretsApi.md#GetTransitBackupName) | **Get** /transit/backup/{name} | Backup the named key
[**GetTransitCacheConfig**](SecretsApi.md#GetTransitCacheConfig) | **Get** /transit/cache-config | Returns the size of the active cache
[**GetTransitExportTypeName**](SecretsApi.md#GetTransitExportTypeName) | **Get** /transit/export/{type}/{name} | Export named encryption or signing key
[**GetTransitExportTypeNameVersion**](SecretsApi.md#GetTransitExportTypeNameVersion) | **Get** /transit/export/{type}/{name}/{version} | Export named encryption or signing key
[**GetTransitKeys**](SecretsApi.md#GetTransitKeys) | **Get** /transit/keys | Managed named encryption keys
[**GetTransitKeysName**](SecretsApi.md#GetTransitKeysName) | **Get** /transit/keys/{name} | Managed named encryption keys
[**GetTransitWrappingKey**](SecretsApi.md#GetTransitWrappingKey) | **Get** /transit/wrapping_key | Returns the public key to use for wrapping imported keys
[**PostAdConfig**](SecretsApi.md#PostAdConfig) | **Post** /ad/config | Configure the AD server to connect to, along with password options.
[**PostAdLibraryManageNameCheckIn**](SecretsApi.md#PostAdLibraryManageNameCheckIn) | **Post** /ad/library/manage/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryName**](SecretsApi.md#PostAdLibraryName) | **Post** /ad/library/{name} | Update a library set.
[**PostAdLibraryNameCheckIn**](SecretsApi.md#PostAdLibraryNameCheckIn) | **Post** /ad/library/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryNameCheckOut**](SecretsApi.md#PostAdLibraryNameCheckOut) | **Post** /ad/library/{name}/check-out | Check a service account out from the library.
[**PostAdRolesName**](SecretsApi.md#PostAdRolesName) | **Post** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**PostAdRotateRoleName**](SecretsApi.md#PostAdRotateRoleName) | **Post** /ad/rotate-role/{name} | 
[**PostAdRotateRoot**](SecretsApi.md#PostAdRotateRoot) | **Post** /ad/rotate-root | 
[**PostAlicloudConfig**](SecretsApi.md#PostAlicloudConfig) | **Post** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**PostAlicloudRoleName**](SecretsApi.md#PostAlicloudRoleName) | **Post** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**PostAwsConfigLease**](SecretsApi.md#PostAwsConfigLease) | **Post** /aws/config/lease | Configure the default lease information for generated credentials.
[**PostAwsConfigRoot**](SecretsApi.md#PostAwsConfigRoot) | **Post** /aws/config/root | Configure the root credentials that are used to manage IAM.
[**PostAwsConfigRotateRoot**](SecretsApi.md#PostAwsConfigRotateRoot) | **Post** /aws/config/rotate-root | 
[**PostAwsCreds**](SecretsApi.md#PostAwsCreds) | **Post** /aws/creds | Generate AWS credentials from a specific Vault role.
[**PostAwsRolesName**](SecretsApi.md#PostAwsRolesName) | **Post** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**PostAwsStsName**](SecretsApi.md#PostAwsStsName) | **Post** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
[**PostAzureConfig**](SecretsApi.md#PostAzureConfig) | **Post** /azure/config | 
[**PostAzureRolesName**](SecretsApi.md#PostAzureRolesName) | **Post** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**PostAzureRotateRoot**](SecretsApi.md#PostAzureRotateRoot) | **Post** /azure/rotate-root | 
[**PostConsulConfigAccess**](SecretsApi.md#PostConsulConfigAccess) | **Post** /consul/config/access | 
[**PostConsulRolesName**](SecretsApi.md#PostConsulRolesName) | **Post** /consul/roles/{name} | 
[**PostCubbyholePath**](SecretsApi.md#PostCubbyholePath) | **Post** /cubbyhole/{path} | Store a secret at the specified location.
[**PostGcpConfig**](SecretsApi.md#PostGcpConfig) | **Post** /gcp/config | 
[**PostGcpConfigRotateRoot**](SecretsApi.md#PostGcpConfigRotateRoot) | **Post** /gcp/config/rotate-root | 
[**PostGcpKeyRoleset**](SecretsApi.md#PostGcpKeyRoleset) | **Post** /gcp/key/{roleset} | 
[**PostGcpRolesetName**](SecretsApi.md#PostGcpRolesetName) | **Post** /gcp/roleset/{name} | 
[**PostGcpRolesetNameRotate**](SecretsApi.md#PostGcpRolesetNameRotate) | **Post** /gcp/roleset/{name}/rotate | 
[**PostGcpRolesetNameRotateKey**](SecretsApi.md#PostGcpRolesetNameRotateKey) | **Post** /gcp/roleset/{name}/rotate-key | 
[**PostGcpRolesetRolesetKey**](SecretsApi.md#PostGcpRolesetRolesetKey) | **Post** /gcp/roleset/{roleset}/key | 
[**PostGcpRolesetRolesetToken**](SecretsApi.md#PostGcpRolesetRolesetToken) | **Post** /gcp/roleset/{roleset}/token | 
[**PostGcpStaticAccountName**](SecretsApi.md#PostGcpStaticAccountName) | **Post** /gcp/static-account/{name} | 
[**PostGcpStaticAccountNameKey**](SecretsApi.md#PostGcpStaticAccountNameKey) | **Post** /gcp/static-account/{name}/key | 
[**PostGcpStaticAccountNameRotateKey**](SecretsApi.md#PostGcpStaticAccountNameRotateKey) | **Post** /gcp/static-account/{name}/rotate-key | 
[**PostGcpStaticAccountNameToken**](SecretsApi.md#PostGcpStaticAccountNameToken) | **Post** /gcp/static-account/{name}/token | 
[**PostGcpTokenRoleset**](SecretsApi.md#PostGcpTokenRoleset) | **Post** /gcp/token/{roleset} | 
[**PostGcpkmsConfig**](SecretsApi.md#PostGcpkmsConfig) | **Post** /gcpkms/config | Configure the GCP KMS secrets engine
[**PostGcpkmsDecryptKey**](SecretsApi.md#PostGcpkmsDecryptKey) | **Post** /gcpkms/decrypt/{key} | Decrypt a ciphertext value using a named key
[**PostGcpkmsEncryptKey**](SecretsApi.md#PostGcpkmsEncryptKey) | **Post** /gcpkms/encrypt/{key} | Encrypt a plaintext value using a named key
[**PostGcpkmsKeysConfigKey**](SecretsApi.md#PostGcpkmsKeysConfigKey) | **Post** /gcpkms/keys/config/{key} | Configure the key in Vault
[**PostGcpkmsKeysDeregisterKey**](SecretsApi.md#PostGcpkmsKeysDeregisterKey) | **Post** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
[**PostGcpkmsKeysKey**](SecretsApi.md#PostGcpkmsKeysKey) | **Post** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**PostGcpkmsKeysRegisterKey**](SecretsApi.md#PostGcpkmsKeysRegisterKey) | **Post** /gcpkms/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**PostGcpkmsKeysRotateKey**](SecretsApi.md#PostGcpkmsKeysRotateKey) | **Post** /gcpkms/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**PostGcpkmsKeysTrimKey**](SecretsApi.md#PostGcpkmsKeysTrimKey) | **Post** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**PostGcpkmsReencryptKey**](SecretsApi.md#PostGcpkmsReencryptKey) | **Post** /gcpkms/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**PostGcpkmsSignKey**](SecretsApi.md#PostGcpkmsSignKey) | **Post** /gcpkms/sign/{key} | Signs a message or digest using a named key
[**PostGcpkmsVerifyKey**](SecretsApi.md#PostGcpkmsVerifyKey) | **Post** /gcpkms/verify/{key} | Verify a signature using a named key
[**PostKubernetesConfig**](SecretsApi.md#PostKubernetesConfig) | **Post** /kubernetes/config | 
[**PostKubernetesCredsName**](SecretsApi.md#PostKubernetesCredsName) | **Post** /kubernetes/creds/{name} | 
[**PostKubernetesRolesName**](SecretsApi.md#PostKubernetesRolesName) | **Post** /kubernetes/roles/{name} | 
[**PostKvPath**](SecretsApi.md#PostKvPath) | **Post** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**PostMongodbatlasConfig**](SecretsApi.md#PostMongodbatlasConfig) | **Post** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
[**PostMongodbatlasCredsName**](SecretsApi.md#PostMongodbatlasCredsName) | **Post** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**PostMongodbatlasRolesName**](SecretsApi.md#PostMongodbatlasRolesName) | **Post** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**PostNomadConfigAccess**](SecretsApi.md#PostNomadConfigAccess) | **Post** /nomad/config/access | 
[**PostNomadConfigLease**](SecretsApi.md#PostNomadConfigLease) | **Post** /nomad/config/lease | Configure the lease parameters for generated tokens
[**PostNomadRoleName**](SecretsApi.md#PostNomadRoleName) | **Post** /nomad/role/{name} | 
[**PostOpenldapConfig**](SecretsApi.md#PostOpenldapConfig) | **Post** /openldap/config | 
[**PostOpenldapRoleName**](SecretsApi.md#PostOpenldapRoleName) | **Post** /openldap/role/{name} | 
[**PostOpenldapRotateRoleName**](SecretsApi.md#PostOpenldapRotateRoleName) | **Post** /openldap/rotate-role/{name} | 
[**PostOpenldapRotateRoot**](SecretsApi.md#PostOpenldapRotateRoot) | **Post** /openldap/rotate-root | 
[**PostOpenldapStaticRoleName**](SecretsApi.md#PostOpenldapStaticRoleName) | **Post** /openldap/static-role/{name} | 
[**PostPkiBundle**](SecretsApi.md#PostPkiBundle) | **Post** /pki/bundle | 
[**PostPkiCert**](SecretsApi.md#PostPkiCert) | **Post** /pki/cert | 
[**PostPkiConfigCa**](SecretsApi.md#PostPkiConfigCa) | **Post** /pki/config/ca | 
[**PostPkiConfigCrl**](SecretsApi.md#PostPkiConfigCrl) | **Post** /pki/config/crl | 
[**PostPkiConfigIssuers**](SecretsApi.md#PostPkiConfigIssuers) | **Post** /pki/config/issuers | 
[**PostPkiConfigKeys**](SecretsApi.md#PostPkiConfigKeys) | **Post** /pki/config/keys | 
[**PostPkiConfigUrls**](SecretsApi.md#PostPkiConfigUrls) | **Post** /pki/config/urls | 
[**PostPkiIntermediateCrossSign**](SecretsApi.md#PostPkiIntermediateCrossSign) | **Post** /pki/intermediate/cross-sign | 
[**PostPkiIntermediateGenerateExported**](SecretsApi.md#PostPkiIntermediateGenerateExported) | **Post** /pki/intermediate/generate/{exported} | 
[**PostPkiIntermediateSetSigned**](SecretsApi.md#PostPkiIntermediateSetSigned) | **Post** /pki/intermediate/set-signed | 
[**PostPkiInternalExported**](SecretsApi.md#PostPkiInternalExported) | **Post** /pki/internal|exported | 
[**PostPkiIssueRole**](SecretsApi.md#PostPkiIssueRole) | **Post** /pki/issue/{role} | 
[**PostPkiIssuerIssuerRefIssueRole**](SecretsApi.md#PostPkiIssuerIssuerRefIssueRole) | **Post** /pki/issuer/{issuer_ref}/issue/{role} | 
[**PostPkiIssuerIssuerRefSignIntermediate**](SecretsApi.md#PostPkiIssuerIssuerRefSignIntermediate) | **Post** /pki/issuer/{issuer_ref}/sign-intermediate | 
[**PostPkiIssuerIssuerRefSignRole**](SecretsApi.md#PostPkiIssuerIssuerRefSignRole) | **Post** /pki/issuer/{issuer_ref}/sign/{role} | 
[**PostPkiIssuerIssuerRefSignSelfIssued**](SecretsApi.md#PostPkiIssuerIssuerRefSignSelfIssued) | **Post** /pki/issuer/{issuer_ref}/sign-self-issued | 
[**PostPkiIssuerIssuerRefSignVerbatim**](SecretsApi.md#PostPkiIssuerIssuerRefSignVerbatim) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim | 
[**PostPkiIssuerIssuerRefSignVerbatimRole**](SecretsApi.md#PostPkiIssuerIssuerRefSignVerbatimRole) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**PostPkiIssuerRefDerPem**](SecretsApi.md#PostPkiIssuerRefDerPem) | **Post** /pki/{issuer_ref}/der|/pem | 
[**PostPkiIssuersGenerateIntermediateExported**](SecretsApi.md#PostPkiIssuersGenerateIntermediateExported) | **Post** /pki/issuers/generate/intermediate/{exported} | 
[**PostPkiIssuersGenerateRootExported**](SecretsApi.md#PostPkiIssuersGenerateRootExported) | **Post** /pki/issuers/generate/root/{exported} | 
[**PostPkiJson**](SecretsApi.md#PostPkiJson) | **Post** /pki//json | 
[**PostPkiKeyKeyRef**](SecretsApi.md#PostPkiKeyKeyRef) | **Post** /pki/key/{key_ref} | 
[**PostPkiKeysImport**](SecretsApi.md#PostPkiKeysImport) | **Post** /pki/keys/import | 
[**PostPkiKms**](SecretsApi.md#PostPkiKms) | **Post** /pki/kms | 
[**PostPkiRevoke**](SecretsApi.md#PostPkiRevoke) | **Post** /pki/revoke | 
[**PostPkiRolesName**](SecretsApi.md#PostPkiRolesName) | **Post** /pki/roles/{name} | 
[**PostPkiRootGenerateExported**](SecretsApi.md#PostPkiRootGenerateExported) | **Post** /pki/root/generate/{exported} | 
[**PostPkiRootReplace**](SecretsApi.md#PostPkiRootReplace) | **Post** /pki/root/replace | 
[**PostPkiRootRotateExported**](SecretsApi.md#PostPkiRootRotateExported) | **Post** /pki/root/rotate/{exported} | 
[**PostPkiRootSignIntermediate**](SecretsApi.md#PostPkiRootSignIntermediate) | **Post** /pki/root/sign-intermediate | 
[**PostPkiRootSignSelfIssued**](SecretsApi.md#PostPkiRootSignSelfIssued) | **Post** /pki/root/sign-self-issued | 
[**PostPkiSignRole**](SecretsApi.md#PostPkiSignRole) | **Post** /pki/sign/{role} | 
[**PostPkiSignVerbatim**](SecretsApi.md#PostPkiSignVerbatim) | **Post** /pki/sign-verbatim | 
[**PostPkiSignVerbatimRole**](SecretsApi.md#PostPkiSignVerbatimRole) | **Post** /pki/sign-verbatim/{role} | 
[**PostPkiTidy**](SecretsApi.md#PostPkiTidy) | **Post** /pki/tidy | 
[**PostRabbitmqConfigConnection**](SecretsApi.md#PostRabbitmqConfigConnection) | **Post** /rabbitmq/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**PostRabbitmqConfigLease**](SecretsApi.md#PostRabbitmqConfigLease) | **Post** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
[**PostRabbitmqRolesName**](SecretsApi.md#PostRabbitmqRolesName) | **Post** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**PostSecretConfig**](SecretsApi.md#PostSecretConfig) | **Post** /secret/config | Configure backend level settings that are applied to every key in the key-value store.
[**PostSecretDataPath**](SecretsApi.md#PostSecretDataPath) | **Post** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**PostSecretDeletePath**](SecretsApi.md#PostSecretDeletePath) | **Post** /secret/delete/{path} | Marks one or more versions as deleted in the KV store.
[**PostSecretDestroyPath**](SecretsApi.md#PostSecretDestroyPath) | **Post** /secret/destroy/{path} | Permanently removes one or more versions in the KV store
[**PostSecretMetadataPath**](SecretsApi.md#PostSecretMetadataPath) | **Post** /secret/metadata/{path} | Configures settings for the KV store
[**PostSecretUndeletePath**](SecretsApi.md#PostSecretUndeletePath) | **Post** /secret/undelete/{path} | Undeletes one or more versions from the KV store.
[**PostSshConfigCa**](SecretsApi.md#PostSshConfigCa) | **Post** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**PostSshConfigZeroaddress**](SecretsApi.md#PostSshConfigZeroaddress) | **Post** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**PostSshCredsRole**](SecretsApi.md#PostSshCredsRole) | **Post** /ssh/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
[**PostSshKeysKeyName**](SecretsApi.md#PostSshKeysKeyName) | **Post** /ssh/keys/{key_name} | Register a shared private key with Vault.
[**PostSshLookup**](SecretsApi.md#PostSshLookup) | **Post** /ssh/lookup | List all the roles associated with the given IP address.
[**PostSshRolesRole**](SecretsApi.md#PostSshRolesRole) | **Post** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**PostSshSignRole**](SecretsApi.md#PostSshSignRole) | **Post** /ssh/sign/{role} | Request signing an SSH key using a certain role with the provided details.
[**PostSshVerify**](SecretsApi.md#PostSshVerify) | **Post** /ssh/verify | Validate the OTP provided by Vault SSH Agent.
[**PostTerraformConfig**](SecretsApi.md#PostTerraformConfig) | **Post** /terraform/config | 
[**PostTerraformCredsName**](SecretsApi.md#PostTerraformCredsName) | **Post** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**PostTerraformRoleName**](SecretsApi.md#PostTerraformRoleName) | **Post** /terraform/role/{name} | 
[**PostTerraformRotateRoleName**](SecretsApi.md#PostTerraformRotateRoleName) | **Post** /terraform/rotate-role/{name} | 
[**PostTotpCodeName**](SecretsApi.md#PostTotpCodeName) | **Post** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**PostTotpKeysName**](SecretsApi.md#PostTotpKeysName) | **Post** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**PostTransitCacheConfig**](SecretsApi.md#PostTransitCacheConfig) | **Post** /transit/cache-config | Configures a new cache of the specified size
[**PostTransitDatakeyPlaintextName**](SecretsApi.md#PostTransitDatakeyPlaintextName) | **Post** /transit/datakey/{plaintext}/{name} | Generate a data key
[**PostTransitDecryptName**](SecretsApi.md#PostTransitDecryptName) | **Post** /transit/decrypt/{name} | Decrypt a ciphertext value using a named key
[**PostTransitEncryptName**](SecretsApi.md#PostTransitEncryptName) | **Post** /transit/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**PostTransitHash**](SecretsApi.md#PostTransitHash) | **Post** /transit/hash | Generate a hash sum for input data
[**PostTransitHashUrlalgorithm**](SecretsApi.md#PostTransitHashUrlalgorithm) | **Post** /transit/hash/{urlalgorithm} | Generate a hash sum for input data
[**PostTransitHmacName**](SecretsApi.md#PostTransitHmacName) | **Post** /transit/hmac/{name} | Generate an HMAC for input data using the named key
[**PostTransitHmacNameUrlalgorithm**](SecretsApi.md#PostTransitHmacNameUrlalgorithm) | **Post** /transit/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**PostTransitKeysName**](SecretsApi.md#PostTransitKeysName) | **Post** /transit/keys/{name} | Managed named encryption keys
[**PostTransitKeysNameConfig**](SecretsApi.md#PostTransitKeysNameConfig) | **Post** /transit/keys/{name}/config | Configure a named encryption key
[**PostTransitKeysNameImport**](SecretsApi.md#PostTransitKeysNameImport) | **Post** /transit/keys/{name}/import | Imports an externally-generated key into a new transit key
[**PostTransitKeysNameImportVersion**](SecretsApi.md#PostTransitKeysNameImportVersion) | **Post** /transit/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**PostTransitKeysNameRotate**](SecretsApi.md#PostTransitKeysNameRotate) | **Post** /transit/keys/{name}/rotate | Rotate named encryption key
[**PostTransitKeysNameTrim**](SecretsApi.md#PostTransitKeysNameTrim) | **Post** /transit/keys/{name}/trim | Trim key versions of a named key
[**PostTransitRandom**](SecretsApi.md#PostTransitRandom) | **Post** /transit/random | Generate random bytes
[**PostTransitRandomSource**](SecretsApi.md#PostTransitRandomSource) | **Post** /transit/random/{source} | Generate random bytes
[**PostTransitRandomSourceUrlbytes**](SecretsApi.md#PostTransitRandomSourceUrlbytes) | **Post** /transit/random/{source}/{urlbytes} | Generate random bytes
[**PostTransitRandomUrlbytes**](SecretsApi.md#PostTransitRandomUrlbytes) | **Post** /transit/random/{urlbytes} | Generate random bytes
[**PostTransitRestore**](SecretsApi.md#PostTransitRestore) | **Post** /transit/restore | Restore the named key
[**PostTransitRestoreName**](SecretsApi.md#PostTransitRestoreName) | **Post** /transit/restore/{name} | Restore the named key
[**PostTransitRewrapName**](SecretsApi.md#PostTransitRewrapName) | **Post** /transit/rewrap/{name} | Rewrap ciphertext
[**PostTransitSignName**](SecretsApi.md#PostTransitSignName) | **Post** /transit/sign/{name} | Generate a signature for input data using the named key
[**PostTransitSignNameUrlalgorithm**](SecretsApi.md#PostTransitSignNameUrlalgorithm) | **Post** /transit/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**PostTransitVerifyName**](SecretsApi.md#PostTransitVerifyName) | **Post** /transit/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**PostTransitVerifyNameUrlalgorithm**](SecretsApi.md#PostTransitVerifyNameUrlalgorithm) | **Post** /transit/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key



## DeleteAdConfig

> DeleteAdConfig(ctx).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAdConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAdConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdLibraryName

> DeleteAdLibraryName(ctx, name).Execute()

Delete a library set.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAdLibraryName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAdLibraryName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdLibraryNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdRolesName

> DeleteAdRolesName(ctx, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAdRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAdRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlicloudConfig

> DeleteAlicloudConfig(ctx).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAlicloudConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAlicloudConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlicloudConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlicloudRoleName

> DeleteAlicloudRoleName(ctx, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAlicloudRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAlicloudRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlicloudRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAwsRolesName

> DeleteAwsRolesName(ctx, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAwsRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAwsRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureConfig

> DeleteAzureConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureRolesName

> DeleteAzureRolesName(ctx, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteAzureRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteAzureRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsulRolesName

> DeleteConsulRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteConsulRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteConsulRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsulRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCubbyholePath

> DeleteCubbyholePath(ctx, path).Execute()

Deletes the secret at the specified location.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Specifies the path of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteCubbyholePath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteCubbyholePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCubbyholePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpRolesetName

> DeleteGcpRolesetName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpRolesetName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpRolesetName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpRolesetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpStaticAccountName

> DeleteGcpStaticAccountName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpStaticAccountName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpStaticAccountName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpStaticAccountNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpkmsConfig

> DeleteGcpkmsConfig(ctx).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpkmsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpkmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpkmsConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysDeregisterKey

> DeleteGcpkmsKeysDeregisterKey(ctx, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpkmsKeysDeregisterKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpkmsKeysDeregisterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpkmsKeysDeregisterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysKey

> DeleteGcpkmsKeysKey(ctx, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpkmsKeysKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpkmsKeysKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpkmsKeysKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysTrimKey

> DeleteGcpkmsKeysTrimKey(ctx, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteGcpkmsKeysTrimKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteGcpkmsKeysTrimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpkmsKeysTrimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubernetesConfig

> DeleteKubernetesConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteKubernetesConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteKubernetesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubernetesRolesName

> DeleteKubernetesRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteKubernetesRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteKubernetesRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKvPath

> DeleteKvPath(ctx, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteKvPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteKvPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKvPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMongodbatlasRolesName

> DeleteMongodbatlasRolesName(ctx, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Roles

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteMongodbatlasRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteMongodbatlasRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMongodbatlasRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNomadConfigAccess

> DeleteNomadConfigAccess(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteNomadConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteNomadConfigAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNomadConfigAccessRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNomadConfigLease

> DeleteNomadConfigLease(ctx).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteNomadConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteNomadConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNomadConfigLeaseRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNomadRoleName

> DeleteNomadRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteNomadRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteNomadRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNomadRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOpenldapConfig

> DeleteOpenldapConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteOpenldapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteOpenldapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpenldapConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOpenldapRoleName

> DeleteOpenldapRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role (lowercase)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteOpenldapRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteOpenldapRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpenldapRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOpenldapStaticRoleName

> DeleteOpenldapStaticRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteOpenldapStaticRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteOpenldapStaticRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpenldapStaticRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePkiIssuerRefDerPem

> DeletePkiIssuerRefDerPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeletePkiIssuerRefDerPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeletePkiIssuerRefDerPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePkiIssuerRefDerPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePkiJson

> DeletePkiJson(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeletePkiJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeletePkiJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePkiJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePkiKeyKeyRef

> DeletePkiKeyKeyRef(ctx, keyRef).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (default to "default")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeletePkiKeyKeyRef(context.Background(), keyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeletePkiKeyKeyRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePkiKeyKeyRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePkiRolesName

> DeletePkiRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeletePkiRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeletePkiRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePkiRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePkiRoot

> DeletePkiRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeletePkiRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeletePkiRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePkiRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRabbitmqRolesName

> DeleteRabbitmqRolesName(ctx, name).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteRabbitmqRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteRabbitmqRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRabbitmqRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecretDataPath

> DeleteSecretDataPath(ctx, path).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSecretDataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSecretDataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretDataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecretMetadataPath

> DeleteSecretMetadataPath(ctx, path).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSecretMetadataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSecretMetadataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretMetadataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshConfigCa

> DeleteSshConfigCa(ctx).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSshConfigCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSshConfigCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshConfigCaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshConfigZeroaddress

> DeleteSshConfigZeroaddress(ctx).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSshConfigZeroaddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSshConfigZeroaddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshConfigZeroaddressRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshKeysKeyName

> DeleteSshKeysKeyName(ctx, keyName).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyName := "keyName_example" // string | [Required] Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSshKeysKeyName(context.Background(), keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSshKeysKeyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** | [Required] Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeysKeyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshRolesRole

> DeleteSshRolesRole(ctx, role).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | [Required for all types] Name of the role being created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteSshRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSshRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshRolesRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTerraformConfig

> DeleteTerraformConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteTerraformConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteTerraformConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTerraformConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTerraformRoleName

> DeleteTerraformRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteTerraformRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteTerraformRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTerraformRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTotpKeysName

> DeleteTotpKeysName(ctx, name).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteTotpKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteTotpKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTotpKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitKeysName

> DeleteTransitKeysName(ctx, name).Execute()

Managed named encryption keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.DeleteTransitKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteTransitKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdConfig

> GetAdConfig(ctx).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdCredsName

> GetAdCredsName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdLibrary

> GetAdLibrary(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdLibrary(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdLibraryName

> GetAdLibraryName(ctx, name).Execute()

Read a library set.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdLibraryName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdLibraryName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdLibraryNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdLibraryNameStatus

> GetAdLibraryNameStatus(ctx, name).Execute()

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdLibraryNameStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdLibraryNameStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdLibraryNameStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdRoles

> GetAdRoles(ctx).List(list).Execute()

List the name of each role currently stored.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdRolesName

> GetAdRolesName(ctx, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdRotateRoot

> GetAdRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAdRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAdRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlicloudConfig

> GetAlicloudConfig(ctx).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAlicloudConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAlicloudConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlicloudConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlicloudCredsName

> GetAlicloudCredsName(ctx, name).Execute()

Generate an API key or STS credential using the given role's configuration.'

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAlicloudCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAlicloudCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlicloudCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlicloudRole

> GetAlicloudRole(ctx).List(list).Execute()

List the existing roles in this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAlicloudRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAlicloudRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlicloudRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlicloudRoleName

> GetAlicloudRoleName(ctx, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAlicloudRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAlicloudRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlicloudRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsConfigLease

> GetAwsConfigLease(ctx).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsConfigLeaseRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsConfigRoot

> GetAwsConfigRoot(ctx).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsConfigRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsConfigRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsConfigRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsCreds

> GetAwsCreds(ctx).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsCreds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCredsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsRoles

> GetAwsRoles(ctx).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsRolesName

> GetAwsRolesName(ctx, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsStsName

> GetAwsStsName(ctx, name).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAwsStsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAwsStsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsStsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureConfig

> GetAzureConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureCredsRole

> GetAzureCredsRole(ctx, role).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | Name of the Vault role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAzureCredsRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAzureCredsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | Name of the Vault role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCredsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureRoles

> GetAzureRoles(ctx).List(list).Execute()

List existing roles.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAzureRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAzureRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureRolesName

> GetAzureRolesName(ctx, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetAzureRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetAzureRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsulConfigAccess

> GetConsulConfigAccess(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetConsulConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetConsulConfigAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsulConfigAccessRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsulCredsRole

> GetConsulCredsRole(ctx, role).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetConsulCredsRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetConsulCredsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsulCredsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsulRoles

> GetConsulRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetConsulRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetConsulRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsulRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsulRolesName

> GetConsulRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetConsulRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetConsulRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsulRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCubbyholePath

> GetCubbyholePath(ctx, path).List(list).Execute()

Retrieve the secret at the specified location.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Specifies the path of the secret.
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetCubbyholePath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetCubbyholePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCubbyholePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpConfig

> GetGcpConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpKeyRoleset

> GetGcpKeyRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpKeyRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpKeyRoleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpKeyRolesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpRolesetName

> GetGcpRolesetName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpRolesetName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpRolesetName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpRolesetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpRolesetRolesetKey

> GetGcpRolesetRolesetKey(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpRolesetRolesetKey(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpRolesetRolesetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpRolesetRolesetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpRolesetRolesetToken

> GetGcpRolesetRolesetToken(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpRolesetRolesetToken(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpRolesetRolesetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpRolesetRolesetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpRolesets

> GetGcpRolesets(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpRolesets(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpRolesets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpRolesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpStaticAccountName

> GetGcpStaticAccountName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpStaticAccountName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpStaticAccountName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpStaticAccountNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpStaticAccountNameKey

> GetGcpStaticAccountNameKey(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the static account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpStaticAccountNameKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpStaticAccountNameKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the static account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpStaticAccountNameKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpStaticAccountNameToken

> GetGcpStaticAccountNameToken(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the static account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpStaticAccountNameToken(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpStaticAccountNameToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the static account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpStaticAccountNameTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpStaticAccounts

> GetGcpStaticAccounts(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpStaticAccounts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpStaticAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpStaticAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpTokenRoleset

> GetGcpTokenRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpTokenRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpTokenRoleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpTokenRolesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpkmsConfig

> GetGcpkmsConfig(ctx).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpkmsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpkmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpkmsConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpkmsKeys

> GetGcpkmsKeys(ctx).List(list).Execute()

List named keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpkmsKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpkmsKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpkmsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpkmsKeysConfigKey

> GetGcpkmsKeysConfigKey(ctx, key).Execute()

Configure the key in Vault

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpkmsKeysConfigKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpkmsKeysConfigKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpkmsKeysConfigKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpkmsKeysKey

> GetGcpkmsKeysKey(ctx, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpkmsKeysKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpkmsKeysKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpkmsKeysKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpkmsPubkeyKey

> GetGcpkmsPubkeyKey(ctx, key).Execute()

Retrieve the public key associated with the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetGcpkmsPubkeyKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetGcpkmsPubkeyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpkmsPubkeyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesConfig

> GetKubernetesConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetKubernetesConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetKubernetesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesRoles

> GetKubernetesRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetKubernetesRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetKubernetesRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesRolesName

> GetKubernetesRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetKubernetesRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetKubernetesRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKvPath

> GetKvPath(ctx, path).List(list).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetKvPath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetKvPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKvPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMongodbatlasConfig

> GetMongodbatlasConfig(ctx).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetMongodbatlasConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetMongodbatlasConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMongodbatlasConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMongodbatlasCredsName

> GetMongodbatlasCredsName(ctx, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetMongodbatlasCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetMongodbatlasCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMongodbatlasCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMongodbatlasRoles

> GetMongodbatlasRoles(ctx).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetMongodbatlasRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetMongodbatlasRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMongodbatlasRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMongodbatlasRolesName

> GetMongodbatlasRolesName(ctx, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Roles

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetMongodbatlasRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetMongodbatlasRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMongodbatlasRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNomadConfigAccess

> GetNomadConfigAccess(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetNomadConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetNomadConfigAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNomadConfigAccessRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNomadConfigLease

> GetNomadConfigLease(ctx).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetNomadConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetNomadConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNomadConfigLeaseRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNomadCredsName

> GetNomadCredsName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetNomadCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetNomadCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNomadCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNomadRole

> GetNomadRole(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetNomadRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetNomadRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNomadRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNomadRoleName

> GetNomadRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetNomadRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetNomadRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNomadRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapConfig

> GetOpenldapConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapCredsName

> GetOpenldapCredsName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the dynamic role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the dynamic role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapRole

> GetOpenldapRole(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapRoleName

> GetOpenldapRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role (lowercase)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapStaticCredName

> GetOpenldapStaticCredName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the static role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapStaticCredName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapStaticCredName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the static role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapStaticCredNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapStaticRole

> GetOpenldapStaticRole(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapStaticRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapStaticRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapStaticRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenldapStaticRoleName

> GetOpenldapStaticRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetOpenldapStaticRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetOpenldapStaticRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenldapStaticRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCa

> GetPkiCa(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCaChain

> GetPkiCaChain(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCaChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCaChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCaChainRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCaPem

> GetPkiCaPem(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCaPem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCaPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCaPemRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCertCaChain

> GetPkiCertCaChain(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCertCaChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCertCaChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertCaChainRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCertCrl

> GetPkiCertCrl(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCertCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCertCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertCrlRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCertSerial

> GetPkiCertSerial(ctx, serial).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCertSerial(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCertSerial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertSerialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCertSerialRaw

> GetPkiCertSerialRaw(ctx, serial).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCertSerialRaw(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCertSerialRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertSerialRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCertSerialRawPem

> GetPkiCertSerialRawPem(ctx, serial).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCertSerialRawPem(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCertSerialRawPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertSerialRawPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCerts

> GetPkiCerts(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCerts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiConfigCrl

> GetPkiConfigCrl(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiConfigCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiConfigCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiConfigCrlRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiConfigIssuers

> GetPkiConfigIssuers(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiConfigIssuers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiConfigIssuers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiConfigIssuersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiConfigKeys

> GetPkiConfigKeys(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiConfigKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiConfigKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiConfigKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiConfigUrls

> GetPkiConfigUrls(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiConfigUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiConfigUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiConfigUrlsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCrl

> GetPkiCrl(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCrlRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCrlPem

> GetPkiCrlPem(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCrlPem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCrlPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCrlPemRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiCrlRotate

> GetPkiCrlRotate(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiCrlRotate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiCrlRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiCrlRotateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiDer

> GetPkiDer(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiDer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiDer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiDerRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiIssuerRefCrlPem

> GetPkiIssuerRefCrlPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiIssuerRefCrlPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiIssuerRefCrlPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiIssuerRefCrlPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiIssuerRefDerPem

> GetPkiIssuerRefDerPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiIssuerRefDerPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiIssuerRefDerPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiIssuerRefDerPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiIssuers

> GetPkiIssuers(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiIssuers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiIssuers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiIssuersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiJson

> GetPkiJson(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiKeyKeyRef

> GetPkiKeyKeyRef(ctx, keyRef).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (default to "default")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiKeyKeyRef(context.Background(), keyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiKeyKeyRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiKeyKeyRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiKeys

> GetPkiKeys(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiRoles

> GetPkiRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiRolesName

> GetPkiRolesName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPkiTidyStatus

> GetPkiTidyStatus(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetPkiTidyStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetPkiTidyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPkiTidyStatusRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRabbitmqConfigLease

> GetRabbitmqConfigLease(ctx).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetRabbitmqConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetRabbitmqConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRabbitmqConfigLeaseRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRabbitmqCredsName

> GetRabbitmqCredsName(ctx, name).Execute()

Request RabbitMQ credentials for a certain role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetRabbitmqCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetRabbitmqCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRabbitmqCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRabbitmqRoles

> GetRabbitmqRoles(ctx).List(list).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetRabbitmqRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetRabbitmqRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRabbitmqRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRabbitmqRolesName

> GetRabbitmqRolesName(ctx, name).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetRabbitmqRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetRabbitmqRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRabbitmqRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretConfig

> GetSecretConfig(ctx).Execute()

Read the backend level settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSecretConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecretConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretDataPath

> GetSecretDataPath(ctx, path).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSecretDataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecretDataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretDataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretMetadataPath

> GetSecretMetadataPath(ctx, path).List(list).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSecretMetadataPath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecretMetadataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretMetadataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretSubkeysPath

> GetSecretSubkeysPath(ctx, path).Execute()

Read the structure of a secret entry from the Key-Value store with the values removed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSecretSubkeysPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecretSubkeysPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretSubkeysPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshConfigCa

> GetSshConfigCa(ctx).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSshConfigCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSshConfigCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshConfigCaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshConfigZeroaddress

> GetSshConfigZeroaddress(ctx).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSshConfigZeroaddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSshConfigZeroaddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshConfigZeroaddressRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshPublicKey

> GetSshPublicKey(ctx).Execute()

Retrieve the public key.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSshPublicKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSshPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshPublicKeyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshRoles

> GetSshRoles(ctx).List(list).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSshRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSshRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshRolesRole

> GetSshRolesRole(ctx, role).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | [Required for all types] Name of the role being created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetSshRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSshRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshRolesRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformConfig

> GetTerraformConfig(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTerraformConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTerraformConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformCredsName

> GetTerraformCredsName(ctx, name).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTerraformCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTerraformCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformRole

> GetTerraformRole(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTerraformRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTerraformRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformRoleName

> GetTerraformRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTerraformRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTerraformRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotpCodeName

> GetTotpCodeName(ctx, name).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTotpCodeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTotpCodeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotpCodeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotpKeys

> GetTotpKeys(ctx).List(list).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTotpKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTotpKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotpKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotpKeysName

> GetTotpKeysName(ctx, name).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTotpKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTotpKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotpKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitBackupName

> GetTransitBackupName(ctx, name).Execute()

Backup the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitBackupName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitBackupName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitBackupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitCacheConfig

> GetTransitCacheConfig(ctx).Execute()

Returns the size of the active cache

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitCacheConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitCacheConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitCacheConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitExportTypeName

> GetTransitExportTypeName(ctx, name, type_).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitExportTypeName(context.Background(), name, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitExportTypeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitExportTypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitExportTypeNameVersion

> GetTransitExportTypeNameVersion(ctx, name, type_, version).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
    version := "version_example" // string | Version of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitExportTypeNameVersion(context.Background(), name, type_, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitExportTypeNameVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 
**version** | **string** | Version of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitExportTypeNameVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitKeys

> GetTransitKeys(ctx).List(list).Execute()

Managed named encryption keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitKeysName

> GetTransitKeysName(ctx, name).Execute()

Managed named encryption keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitWrappingKey

> GetTransitWrappingKey(ctx).Execute()

Returns the public key to use for wrapping imported keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.GetTransitWrappingKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetTransitWrappingKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitWrappingKeyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdConfig

> PostAdConfig(ctx).AdConfigRequest(adConfigRequest).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    adConfigRequest := *openapiclient.NewAdConfigRequest() // AdConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdConfig(context.Background()).AdConfigRequest(adConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adConfigRequest** | [**AdConfigRequest**](AdConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdLibraryManageNameCheckIn

> PostAdLibraryManageNameCheckIn(ctx, name).AdLibraryManageCheckInRequest(adLibraryManageCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.
    adLibraryManageCheckInRequest := *openapiclient.NewAdLibraryManageCheckInRequest() // AdLibraryManageCheckInRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdLibraryManageNameCheckIn(context.Background(), name).AdLibraryManageCheckInRequest(adLibraryManageCheckInRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdLibraryManageNameCheckIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdLibraryManageNameCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryManageCheckInRequest** | [**AdLibraryManageCheckInRequest**](AdLibraryManageCheckInRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdLibraryName

> PostAdLibraryName(ctx, name).AdLibraryRequest(adLibraryRequest).Execute()

Update a library set.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.
    adLibraryRequest := *openapiclient.NewAdLibraryRequest() // AdLibraryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdLibraryName(context.Background(), name).AdLibraryRequest(adLibraryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdLibraryName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdLibraryNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryRequest** | [**AdLibraryRequest**](AdLibraryRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckIn

> PostAdLibraryNameCheckIn(ctx, name).AdLibraryCheckInRequest(adLibraryCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set.
    adLibraryCheckInRequest := *openapiclient.NewAdLibraryCheckInRequest() // AdLibraryCheckInRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdLibraryNameCheckIn(context.Background(), name).AdLibraryCheckInRequest(adLibraryCheckInRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdLibraryNameCheckIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdLibraryNameCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryCheckInRequest** | [**AdLibraryCheckInRequest**](AdLibraryCheckInRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckOut

> PostAdLibraryNameCheckOut(ctx, name).AdLibraryCheckOutRequest(adLibraryCheckOutRequest).Execute()

Check a service account out from the library.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the set
    adLibraryCheckOutRequest := *openapiclient.NewAdLibraryCheckOutRequest() // AdLibraryCheckOutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdLibraryNameCheckOut(context.Background(), name).AdLibraryCheckOutRequest(adLibraryCheckOutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdLibraryNameCheckOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the set | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdLibraryNameCheckOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryCheckOutRequest** | [**AdLibraryCheckOutRequest**](AdLibraryCheckOutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdRolesName

> PostAdRolesName(ctx, name).AdRolesRequest(adRolesRequest).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    adRolesRequest := *openapiclient.NewAdRolesRequest() // AdRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdRolesName(context.Background(), name).AdRolesRequest(adRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adRolesRequest** | [**AdRolesRequest**](AdRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdRotateRoleName

> PostAdRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the static role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdRotateRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the static role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdRotateRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdRotateRoot

> PostAdRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAdRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAdRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAlicloudConfig

> PostAlicloudConfig(ctx).AlicloudConfigRequest(alicloudConfigRequest).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    alicloudConfigRequest := *openapiclient.NewAlicloudConfigRequest() // AlicloudConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAlicloudConfig(context.Background()).AlicloudConfigRequest(alicloudConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAlicloudConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAlicloudConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alicloudConfigRequest** | [**AlicloudConfigRequest**](AlicloudConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAlicloudRoleName

> PostAlicloudRoleName(ctx, name).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the role.
    alicloudRoleRequest := *openapiclient.NewAlicloudRoleRequest() // AlicloudRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAlicloudRoleName(context.Background(), name).AlicloudRoleRequest(alicloudRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAlicloudRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAlicloudRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alicloudRoleRequest** | [**AlicloudRoleRequest**](AlicloudRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsConfigLease

> PostAwsConfigLease(ctx).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigLeaseRequest := *openapiclient.NewAwsConfigLeaseRequest() // AwsConfigLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsConfigLease(context.Background()).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsConfigLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigLeaseRequest** | [**AwsConfigLeaseRequest**](AwsConfigLeaseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsConfigRoot

> PostAwsConfigRoot(ctx).AwsConfigRootRequest(awsConfigRootRequest).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigRootRequest := *openapiclient.NewAwsConfigRootRequest() // AwsConfigRootRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsConfigRoot(context.Background()).AwsConfigRootRequest(awsConfigRootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsConfigRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsConfigRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigRootRequest** | [**AwsConfigRootRequest**](AwsConfigRootRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsConfigRotateRoot

> PostAwsConfigRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsConfigRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsConfigRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsConfigRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsCreds

> PostAwsCreds(ctx).AwsCredsRequest(awsCredsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsCredsRequest := *openapiclient.NewAwsCredsRequest() // AwsCredsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsCreds(context.Background()).AwsCredsRequest(awsCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredsRequest** | [**AwsCredsRequest**](AwsCredsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsRolesName

> PostAwsRolesName(ctx, name).AwsRolesRequest(awsRolesRequest).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the policy
    awsRolesRequest := *openapiclient.NewAwsRolesRequest() // AwsRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsRolesName(context.Background(), name).AwsRolesRequest(awsRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRolesRequest** | [**AwsRolesRequest**](AwsRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAwsStsName

> PostAwsStsName(ctx, name).AwsStsRequest(awsStsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    awsStsRequest := *openapiclient.NewAwsStsRequest() // AwsStsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAwsStsName(context.Background(), name).AwsStsRequest(awsStsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAwsStsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAwsStsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsStsRequest** | [**AwsStsRequest**](AwsStsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAzureConfig

> PostAzureConfig(ctx).AzureConfigRequest(azureConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    azureConfigRequest := *openapiclient.NewAzureConfigRequest() // AzureConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAzureConfig(context.Background()).AzureConfigRequest(azureConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAzureConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureConfigRequest** | [**AzureConfigRequest**](AzureConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAzureRolesName

> PostAzureRolesName(ctx, name).AzureRolesRequest(azureRolesRequest).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.
    azureRolesRequest := *openapiclient.NewAzureRolesRequest() // AzureRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAzureRolesName(context.Background(), name).AzureRolesRequest(azureRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAzureRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAzureRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureRolesRequest** | [**AzureRolesRequest**](AzureRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAzureRotateRoot

> PostAzureRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostAzureRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostAzureRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAzureRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConsulConfigAccess

> PostConsulConfigAccess(ctx).ConsulConfigAccessRequest(consulConfigAccessRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    consulConfigAccessRequest := *openapiclient.NewConsulConfigAccessRequest() // ConsulConfigAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostConsulConfigAccess(context.Background()).ConsulConfigAccessRequest(consulConfigAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostConsulConfigAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConsulConfigAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulConfigAccessRequest** | [**ConsulConfigAccessRequest**](ConsulConfigAccessRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConsulRolesName

> PostConsulRolesName(ctx, name).ConsulRolesRequest(consulRolesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.
    consulRolesRequest := *openapiclient.NewConsulRolesRequest() // ConsulRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostConsulRolesName(context.Background(), name).ConsulRolesRequest(consulRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostConsulRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConsulRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consulRolesRequest** | [**ConsulRolesRequest**](ConsulRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCubbyholePath

> PostCubbyholePath(ctx, path).Execute()

Store a secret at the specified location.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Specifies the path of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostCubbyholePath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostCubbyholePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCubbyholePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpConfig

> PostGcpConfig(ctx).GcpConfigRequest(gcpConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gcpConfigRequest := *openapiclient.NewGcpConfigRequest() // GcpConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpConfig(context.Background()).GcpConfigRequest(gcpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpConfigRequest** | [**GcpConfigRequest**](GcpConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpConfigRotateRoot

> PostGcpConfigRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpConfigRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpConfigRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpConfigRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpKeyRoleset

> PostGcpKeyRoleset(ctx, roleset).GcpKeyRequest(gcpKeyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.
    gcpKeyRequest := *openapiclient.NewGcpKeyRequest() // GcpKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpKeyRoleset(context.Background(), roleset).GcpKeyRequest(gcpKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpKeyRoleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpKeyRolesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpKeyRequest** | [**GcpKeyRequest**](GcpKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpRolesetName

> PostGcpRolesetName(ctx, name).GcpRolesetRequest(gcpRolesetRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the role.
    gcpRolesetRequest := *openapiclient.NewGcpRolesetRequest() // GcpRolesetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpRolesetName(context.Background(), name).GcpRolesetRequest(gcpRolesetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpRolesetName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpRolesetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRolesetRequest** | [**GcpRolesetRequest**](GcpRolesetRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpRolesetNameRotate

> PostGcpRolesetNameRotate(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpRolesetNameRotate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpRolesetNameRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpRolesetNameRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpRolesetNameRotateKey

> PostGcpRolesetNameRotateKey(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpRolesetNameRotateKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpRolesetNameRotateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpRolesetNameRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpRolesetRolesetKey

> PostGcpRolesetRolesetKey(ctx, roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.
    gcpRolesetKeyRequest := *openapiclient.NewGcpRolesetKeyRequest() // GcpRolesetKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpRolesetRolesetKey(context.Background(), roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpRolesetRolesetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpRolesetRolesetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRolesetKeyRequest** | [**GcpRolesetKeyRequest**](GcpRolesetKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpRolesetRolesetToken

> PostGcpRolesetRolesetToken(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpRolesetRolesetToken(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpRolesetRolesetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpRolesetRolesetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpStaticAccountName

> PostGcpStaticAccountName(ctx, name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
    gcpStaticAccountRequest := *openapiclient.NewGcpStaticAccountRequest() // GcpStaticAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpStaticAccountName(context.Background(), name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpStaticAccountName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpStaticAccountNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpStaticAccountRequest** | [**GcpStaticAccountRequest**](GcpStaticAccountRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameKey

> PostGcpStaticAccountNameKey(ctx, name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the static account.
    gcpStaticAccountKeyRequest := *openapiclient.NewGcpStaticAccountKeyRequest() // GcpStaticAccountKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpStaticAccountNameKey(context.Background(), name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpStaticAccountNameKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the static account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpStaticAccountNameKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpStaticAccountKeyRequest** | [**GcpStaticAccountKeyRequest**](GcpStaticAccountKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameRotateKey

> PostGcpStaticAccountNameRotateKey(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpStaticAccountNameRotateKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpStaticAccountNameRotateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpStaticAccountNameRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameToken

> PostGcpStaticAccountNameToken(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Required. Name of the static account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpStaticAccountNameToken(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpStaticAccountNameToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Required. Name of the static account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpStaticAccountNameTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpTokenRoleset

> PostGcpTokenRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleset := "roleset_example" // string | Required. Name of the role set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpTokenRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpTokenRoleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpTokenRolesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsConfig

> PostGcpkmsConfig(ctx).GcpkmsConfigRequest(gcpkmsConfigRequest).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gcpkmsConfigRequest := *openapiclient.NewGcpkmsConfigRequest() // GcpkmsConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsConfig(context.Background()).GcpkmsConfigRequest(gcpkmsConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpkmsConfigRequest** | [**GcpkmsConfigRequest**](GcpkmsConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsDecryptKey

> PostGcpkmsDecryptKey(ctx, key).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
    gcpkmsDecryptRequest := *openapiclient.NewGcpkmsDecryptRequest() // GcpkmsDecryptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsDecryptKey(context.Background(), key).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsDecryptKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsDecryptKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsDecryptRequest** | [**GcpkmsDecryptRequest**](GcpkmsDecryptRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsEncryptKey

> PostGcpkmsEncryptKey(ctx, key).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()

Encrypt a plaintext value using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
    gcpkmsEncryptRequest := *openapiclient.NewGcpkmsEncryptRequest() // GcpkmsEncryptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsEncryptKey(context.Background(), key).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsEncryptKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsEncryptKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsEncryptRequest** | [**GcpkmsEncryptRequest**](GcpkmsEncryptRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysConfigKey

> PostGcpkmsKeysConfigKey(ctx, key).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()

Configure the key in Vault

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.
    gcpkmsKeysConfigRequest := *openapiclient.NewGcpkmsKeysConfigRequest() // GcpkmsKeysConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysConfigKey(context.Background(), key).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysConfigKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysConfigKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysConfigRequest** | [**GcpkmsKeysConfigRequest**](GcpkmsKeysConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysDeregisterKey

> PostGcpkmsKeysDeregisterKey(ctx, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysDeregisterKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysDeregisterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysDeregisterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysKey

> PostGcpkmsKeysKey(ctx, key).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.
    gcpkmsKeysRequest := *openapiclient.NewGcpkmsKeysRequest() // GcpkmsKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysKey(context.Background(), key).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysRequest** | [**GcpkmsKeysRequest**](GcpkmsKeysRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysRegisterKey

> PostGcpkmsKeysRegisterKey(ctx, key).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()

Register an existing crypto key in Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
    gcpkmsKeysRegisterRequest := *openapiclient.NewGcpkmsKeysRegisterRequest() // GcpkmsKeysRegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysRegisterKey(context.Background(), key).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysRegisterKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysRegisterKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysRegisterRequest** | [**GcpkmsKeysRegisterRequest**](GcpkmsKeysRegisterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysRotateKey

> PostGcpkmsKeysRotateKey(ctx, key).Execute()

Rotate a crypto key to a new primary version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysRotateKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysRotateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsKeysTrimKey

> PostGcpkmsKeysTrimKey(ctx, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsKeysTrimKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsKeysTrimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsKeysTrimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsReencryptKey

> PostGcpkmsReencryptKey(ctx, key).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()

Re-encrypt existing ciphertext data to a new version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
    gcpkmsReencryptRequest := *openapiclient.NewGcpkmsReencryptRequest() // GcpkmsReencryptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsReencryptKey(context.Background(), key).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsReencryptKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsReencryptKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsReencryptRequest** | [**GcpkmsReencryptRequest**](GcpkmsReencryptRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsSignKey

> PostGcpkmsSignKey(ctx, key).GcpkmsSignRequest(gcpkmsSignRequest).Execute()

Signs a message or digest using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
    gcpkmsSignRequest := *openapiclient.NewGcpkmsSignRequest() // GcpkmsSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsSignKey(context.Background(), key).GcpkmsSignRequest(gcpkmsSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsSignKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsSignKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsSignRequest** | [**GcpkmsSignRequest**](GcpkmsSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGcpkmsVerifyKey

> PostGcpkmsVerifyKey(ctx, key).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()

Verify a signature using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
    gcpkmsVerifyRequest := *openapiclient.NewGcpkmsVerifyRequest() // GcpkmsVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostGcpkmsVerifyKey(context.Background(), key).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostGcpkmsVerifyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGcpkmsVerifyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsVerifyRequest** | [**GcpkmsVerifyRequest**](GcpkmsVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKubernetesConfig

> PostKubernetesConfig(ctx).KubernetesConfigRequest(kubernetesConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kubernetesConfigRequest := *openapiclient.NewKubernetesConfigRequest() // KubernetesConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostKubernetesConfig(context.Background()).KubernetesConfigRequest(kubernetesConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostKubernetesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKubernetesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesConfigRequest** | [**KubernetesConfigRequest**](KubernetesConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKubernetesCredsName

> PostKubernetesCredsName(ctx, name).KubernetesCredsRequest(kubernetesCredsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Vault role
    kubernetesCredsRequest := *openapiclient.NewKubernetesCredsRequest("KubernetesNamespace_example") // KubernetesCredsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostKubernetesCredsName(context.Background(), name).KubernetesCredsRequest(kubernetesCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostKubernetesCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Vault role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKubernetesCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredsRequest** | [**KubernetesCredsRequest**](KubernetesCredsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKubernetesRolesName

> PostKubernetesRolesName(ctx, name).KubernetesRolesRequest(kubernetesRolesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    kubernetesRolesRequest := *openapiclient.NewKubernetesRolesRequest([]string{"AllowedKubernetesNamespaces_example"}) // KubernetesRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostKubernetesRolesName(context.Background(), name).KubernetesRolesRequest(kubernetesRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostKubernetesRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKubernetesRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesRolesRequest** | [**KubernetesRolesRequest**](KubernetesRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKvPath

> PostKvPath(ctx, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostKvPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostKvPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKvPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMongodbatlasConfig

> PostMongodbatlasConfig(ctx).MongodbatlasConfigRequest(mongodbatlasConfigRequest).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mongodbatlasConfigRequest := *openapiclient.NewMongodbatlasConfigRequest("PrivateKey_example", "PublicKey_example") // MongodbatlasConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostMongodbatlasConfig(context.Background()).MongodbatlasConfigRequest(mongodbatlasConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostMongodbatlasConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMongodbatlasConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mongodbatlasConfigRequest** | [**MongodbatlasConfigRequest**](MongodbatlasConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMongodbatlasCredsName

> PostMongodbatlasCredsName(ctx, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostMongodbatlasCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostMongodbatlasCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMongodbatlasCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMongodbatlasRolesName

> PostMongodbatlasRolesName(ctx, name).MongodbatlasRolesRequest(mongodbatlasRolesRequest).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Roles
    mongodbatlasRolesRequest := *openapiclient.NewMongodbatlasRolesRequest([]string{"Roles_example"}) // MongodbatlasRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostMongodbatlasRolesName(context.Background(), name).MongodbatlasRolesRequest(mongodbatlasRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostMongodbatlasRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMongodbatlasRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mongodbatlasRolesRequest** | [**MongodbatlasRolesRequest**](MongodbatlasRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNomadConfigAccess

> PostNomadConfigAccess(ctx).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nomadConfigAccessRequest := *openapiclient.NewNomadConfigAccessRequest() // NomadConfigAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostNomadConfigAccess(context.Background()).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostNomadConfigAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNomadConfigAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigAccessRequest** | [**NomadConfigAccessRequest**](NomadConfigAccessRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNomadConfigLease

> PostNomadConfigLease(ctx).NomadConfigLeaseRequest(nomadConfigLeaseRequest).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nomadConfigLeaseRequest := *openapiclient.NewNomadConfigLeaseRequest() // NomadConfigLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostNomadConfigLease(context.Background()).NomadConfigLeaseRequest(nomadConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostNomadConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNomadConfigLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigLeaseRequest** | [**NomadConfigLeaseRequest**](NomadConfigLeaseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNomadRoleName

> PostNomadRoleName(ctx, name).NomadRoleRequest(nomadRoleRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    nomadRoleRequest := *openapiclient.NewNomadRoleRequest() // NomadRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostNomadRoleName(context.Background(), name).NomadRoleRequest(nomadRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostNomadRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNomadRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadRoleRequest** | [**NomadRoleRequest**](NomadRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOpenldapConfig

> PostOpenldapConfig(ctx).OpenldapConfigRequest(openldapConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    openldapConfigRequest := *openapiclient.NewOpenldapConfigRequest() // OpenldapConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostOpenldapConfig(context.Background()).OpenldapConfigRequest(openldapConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostOpenldapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOpenldapConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openldapConfigRequest** | [**OpenldapConfigRequest**](OpenldapConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOpenldapRoleName

> PostOpenldapRoleName(ctx, name).OpenldapRoleRequest(openldapRoleRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role (lowercase)
    openldapRoleRequest := *openapiclient.NewOpenldapRoleRequest("CreationLdif_example", "DeletionLdif_example") // OpenldapRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostOpenldapRoleName(context.Background(), name).OpenldapRoleRequest(openldapRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostOpenldapRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOpenldapRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapRoleRequest** | [**OpenldapRoleRequest**](OpenldapRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOpenldapRotateRoleName

> PostOpenldapRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the static role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostOpenldapRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostOpenldapRotateRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the static role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOpenldapRotateRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOpenldapRotateRoot

> PostOpenldapRotateRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostOpenldapRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostOpenldapRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostOpenldapRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOpenldapStaticRoleName

> PostOpenldapStaticRoleName(ctx, name).OpenldapStaticRoleRequest(openldapStaticRoleRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    openldapStaticRoleRequest := *openapiclient.NewOpenldapStaticRoleRequest() // OpenldapStaticRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostOpenldapStaticRoleName(context.Background(), name).OpenldapStaticRoleRequest(openldapStaticRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostOpenldapStaticRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOpenldapStaticRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapStaticRoleRequest** | [**OpenldapStaticRoleRequest**](OpenldapStaticRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiBundle

> PostPkiBundle(ctx).PkiBundleRequest(pkiBundleRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiBundleRequest := *openapiclient.NewPkiBundleRequest() // PkiBundleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiBundle(context.Background()).PkiBundleRequest(pkiBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiBundleRequest** | [**PkiBundleRequest**](PkiBundleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiCert

> PostPkiCert(ctx).PkiCertRequest(pkiCertRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiCertRequest := *openapiclient.NewPkiCertRequest() // PkiCertRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiCert(context.Background()).PkiCertRequest(pkiCertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiCertRequest** | [**PkiCertRequest**](PkiCertRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiConfigCa

> PostPkiConfigCa(ctx).PkiConfigCaRequest(pkiConfigCaRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiConfigCaRequest := *openapiclient.NewPkiConfigCaRequest() // PkiConfigCaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiConfigCa(context.Background()).PkiConfigCaRequest(pkiConfigCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiConfigCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiConfigCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCaRequest** | [**PkiConfigCaRequest**](PkiConfigCaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiConfigCrl

> PostPkiConfigCrl(ctx).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiConfigCrlRequest := *openapiclient.NewPkiConfigCrlRequest() // PkiConfigCrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiConfigCrl(context.Background()).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiConfigCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiConfigCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCrlRequest** | [**PkiConfigCrlRequest**](PkiConfigCrlRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiConfigIssuers

> PostPkiConfigIssuers(ctx).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiConfigIssuersRequest := *openapiclient.NewPkiConfigIssuersRequest() // PkiConfigIssuersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiConfigIssuers(context.Background()).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiConfigIssuers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiConfigIssuersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigIssuersRequest** | [**PkiConfigIssuersRequest**](PkiConfigIssuersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiConfigKeys

> PostPkiConfigKeys(ctx).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiConfigKeysRequest := *openapiclient.NewPkiConfigKeysRequest() // PkiConfigKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiConfigKeys(context.Background()).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiConfigKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiConfigKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigKeysRequest** | [**PkiConfigKeysRequest**](PkiConfigKeysRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiConfigUrls

> PostPkiConfigUrls(ctx).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiConfigUrlsRequest := *openapiclient.NewPkiConfigUrlsRequest() // PkiConfigUrlsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiConfigUrls(context.Background()).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiConfigUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiConfigUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigUrlsRequest** | [**PkiConfigUrlsRequest**](PkiConfigUrlsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIntermediateCrossSign

> PostPkiIntermediateCrossSign(ctx).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiIntermediateCrossSignRequest := *openapiclient.NewPkiIntermediateCrossSignRequest() // PkiIntermediateCrossSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIntermediateCrossSign(context.Background()).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIntermediateCrossSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIntermediateCrossSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateCrossSignRequest** | [**PkiIntermediateCrossSignRequest**](PkiIntermediateCrossSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIntermediateGenerateExported

> PostPkiIntermediateGenerateExported(ctx, exported).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
    pkiIntermediateGenerateRequest := *openapiclient.NewPkiIntermediateGenerateRequest() // PkiIntermediateGenerateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIntermediateGenerateExported(context.Background(), exported).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIntermediateGenerateExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIntermediateGenerateExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIntermediateGenerateRequest** | [**PkiIntermediateGenerateRequest**](PkiIntermediateGenerateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIntermediateSetSigned

> PostPkiIntermediateSetSigned(ctx).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiIntermediateSetSignedRequest := *openapiclient.NewPkiIntermediateSetSignedRequest() // PkiIntermediateSetSignedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIntermediateSetSigned(context.Background()).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIntermediateSetSigned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIntermediateSetSignedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateSetSignedRequest** | [**PkiIntermediateSetSignedRequest**](PkiIntermediateSetSignedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiInternalExported

> PostPkiInternalExported(ctx).PkiInternalExportedRequest(pkiInternalExportedRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiInternalExportedRequest := *openapiclient.NewPkiInternalExportedRequest() // PkiInternalExportedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiInternalExported(context.Background()).PkiInternalExportedRequest(pkiInternalExportedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiInternalExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiInternalExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiInternalExportedRequest** | [**PkiInternalExportedRequest**](PkiInternalExportedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssueRole

> PostPkiIssueRole(ctx, role).PkiIssueRequest(pkiIssueRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The desired role with configuration for this request
    pkiIssueRequest := *openapiclient.NewPkiIssueRequest() // PkiIssueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssueRole(context.Background(), role).PkiIssueRequest(pkiIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssueRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssueRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssueRequest** | [**PkiIssueRequest**](PkiIssueRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefIssueRole

> PostPkiIssuerIssuerRefIssueRole(ctx, issuerRef, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    role := "role_example" // string | The desired role with configuration for this request
    pkiIssuerIssueRequest := *openapiclient.NewPkiIssuerIssueRequest() // PkiIssuerIssueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefIssueRole(context.Background(), issuerRef, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefIssueRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefIssueRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerIssueRequest** | [**PkiIssuerIssueRequest**](PkiIssuerIssueRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignIntermediate

> PostPkiIssuerIssuerRefSignIntermediate(ctx, issuerRef).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    pkiIssuerSignIntermediateRequest := *openapiclient.NewPkiIssuerSignIntermediateRequest() // PkiIssuerSignIntermediateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefSignIntermediate(context.Background(), issuerRef).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefSignIntermediate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefSignIntermediateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignIntermediateRequest** | [**PkiIssuerSignIntermediateRequest**](PkiIssuerSignIntermediateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignRole

> PostPkiIssuerIssuerRefSignRole(ctx, issuerRef, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    role := "role_example" // string | The desired role with configuration for this request
    pkiIssuerSignRequest := *openapiclient.NewPkiIssuerSignRequest() // PkiIssuerSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefSignRole(context.Background(), issuerRef, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefSignRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefSignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignRequest** | [**PkiIssuerSignRequest**](PkiIssuerSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignSelfIssued

> PostPkiIssuerIssuerRefSignSelfIssued(ctx, issuerRef).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    pkiIssuerSignSelfIssuedRequest := *openapiclient.NewPkiIssuerSignSelfIssuedRequest() // PkiIssuerSignSelfIssuedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefSignSelfIssued(context.Background(), issuerRef).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefSignSelfIssued``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefSignSelfIssuedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignSelfIssuedRequest** | [**PkiIssuerSignSelfIssuedRequest**](PkiIssuerSignSelfIssuedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatim

> PostPkiIssuerIssuerRefSignVerbatim(ctx, issuerRef).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    pkiIssuerSignVerbatimRequest := *openapiclient.NewPkiIssuerSignVerbatimRequest() // PkiIssuerSignVerbatimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefSignVerbatim(context.Background(), issuerRef).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefSignVerbatim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefSignVerbatimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatimRole

> PostPkiIssuerIssuerRefSignVerbatimRole(ctx, issuerRef, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    role := "role_example" // string | The desired role with configuration for this request
    pkiIssuerSignVerbatimRequest := *openapiclient.NewPkiIssuerSignVerbatimRequest() // PkiIssuerSignVerbatimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerIssuerRefSignVerbatimRole(context.Background(), issuerRef, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerIssuerRefSignVerbatimRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerIssuerRefSignVerbatimRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuerRefDerPem

> PostPkiIssuerRefDerPem(ctx, issuerRef).PkiDerPemRequest(pkiDerPemRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (default to "default")
    pkiDerPemRequest := *openapiclient.NewPkiDerPemRequest() // PkiDerPemRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuerRefDerPem(context.Background(), issuerRef).PkiDerPemRequest(pkiDerPemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuerRefDerPem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuerRef** | **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuerRefDerPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiDerPemRequest** | [**PkiDerPemRequest**](PkiDerPemRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateIntermediateExported

> PostPkiIssuersGenerateIntermediateExported(ctx, exported).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
    pkiIssuersGenerateIntermediateRequest := *openapiclient.NewPkiIssuersGenerateIntermediateRequest() // PkiIssuersGenerateIntermediateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuersGenerateIntermediateExported(context.Background(), exported).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuersGenerateIntermediateExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuersGenerateIntermediateExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersGenerateIntermediateRequest** | [**PkiIssuersGenerateIntermediateRequest**](PkiIssuersGenerateIntermediateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateRootExported

> PostPkiIssuersGenerateRootExported(ctx, exported).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
    pkiIssuersGenerateRootRequest := *openapiclient.NewPkiIssuersGenerateRootRequest() // PkiIssuersGenerateRootRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiIssuersGenerateRootExported(context.Background(), exported).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiIssuersGenerateRootExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiIssuersGenerateRootExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersGenerateRootRequest** | [**PkiIssuersGenerateRootRequest**](PkiIssuersGenerateRootRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiJson

> PostPkiJson(ctx).PkiJsonRequest(pkiJsonRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiJsonRequest := *openapiclient.NewPkiJsonRequest() // PkiJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiJson(context.Background()).PkiJsonRequest(pkiJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiJsonRequest** | [**PkiJsonRequest**](PkiJsonRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiKeyKeyRef

> PostPkiKeyKeyRef(ctx, keyRef).PkiKeyRequest(pkiKeyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (default to "default")
    pkiKeyRequest := *openapiclient.NewPkiKeyRequest() // PkiKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiKeyKeyRef(context.Background(), keyRef).PkiKeyRequest(pkiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiKeyKeyRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyRef** | **string** | Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key. | [default to &quot;default&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiKeyKeyRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiKeyRequest** | [**PkiKeyRequest**](PkiKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiKeysImport

> PostPkiKeysImport(ctx).PkiKeysImportRequest(pkiKeysImportRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiKeysImportRequest := *openapiclient.NewPkiKeysImportRequest() // PkiKeysImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiKeysImport(context.Background()).PkiKeysImportRequest(pkiKeysImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiKeysImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiKeysImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKeysImportRequest** | [**PkiKeysImportRequest**](PkiKeysImportRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiKms

> PostPkiKms(ctx).PkiKmsRequest(pkiKmsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiKmsRequest := *openapiclient.NewPkiKmsRequest() // PkiKmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiKms(context.Background()).PkiKmsRequest(pkiKmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiKms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiKmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKmsRequest** | [**PkiKmsRequest**](PkiKmsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRevoke

> PostPkiRevoke(ctx).PkiRevokeRequest(pkiRevokeRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiRevokeRequest := *openapiclient.NewPkiRevokeRequest() // PkiRevokeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRevoke(context.Background()).PkiRevokeRequest(pkiRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRevokeRequest** | [**PkiRevokeRequest**](PkiRevokeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRolesName

> PostPkiRolesName(ctx, name).PkiRolesRequest(pkiRolesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    pkiRolesRequest := *openapiclient.NewPkiRolesRequest() // PkiRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRolesName(context.Background(), name).PkiRolesRequest(pkiRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRolesRequest** | [**PkiRolesRequest**](PkiRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRootGenerateExported

> PostPkiRootGenerateExported(ctx, exported).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
    pkiRootGenerateRequest := *openapiclient.NewPkiRootGenerateRequest() // PkiRootGenerateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRootGenerateExported(context.Background(), exported).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRootGenerateExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRootGenerateExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootGenerateRequest** | [**PkiRootGenerateRequest**](PkiRootGenerateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRootReplace

> PostPkiRootReplace(ctx).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiRootReplaceRequest := *openapiclient.NewPkiRootReplaceRequest() // PkiRootReplaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRootReplace(context.Background()).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRootReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRootReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootReplaceRequest** | [**PkiRootReplaceRequest**](PkiRootReplaceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRootRotateExported

> PostPkiRootRotateExported(ctx, exported).PkiRootRotateRequest(pkiRootRotateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
    pkiRootRotateRequest := *openapiclient.NewPkiRootRotateRequest() // PkiRootRotateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRootRotateExported(context.Background(), exported).PkiRootRotateRequest(pkiRootRotateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRootRotateExported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exported** | **string** | Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key! | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRootRotateExportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootRotateRequest** | [**PkiRootRotateRequest**](PkiRootRotateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRootSignIntermediate

> PostPkiRootSignIntermediate(ctx).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiRootSignIntermediateRequest := *openapiclient.NewPkiRootSignIntermediateRequest() // PkiRootSignIntermediateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRootSignIntermediate(context.Background()).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRootSignIntermediate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRootSignIntermediateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignIntermediateRequest** | [**PkiRootSignIntermediateRequest**](PkiRootSignIntermediateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiRootSignSelfIssued

> PostPkiRootSignSelfIssued(ctx).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiRootSignSelfIssuedRequest := *openapiclient.NewPkiRootSignSelfIssuedRequest() // PkiRootSignSelfIssuedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiRootSignSelfIssued(context.Background()).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiRootSignSelfIssued``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiRootSignSelfIssuedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignSelfIssuedRequest** | [**PkiRootSignSelfIssuedRequest**](PkiRootSignSelfIssuedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiSignRole

> PostPkiSignRole(ctx, role).PkiSignRequest(pkiSignRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The desired role with configuration for this request
    pkiSignRequest := *openapiclient.NewPkiSignRequest() // PkiSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiSignRole(context.Background(), role).PkiSignRequest(pkiSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiSignRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiSignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignRequest** | [**PkiSignRequest**](PkiSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiSignVerbatim

> PostPkiSignVerbatim(ctx).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiSignVerbatimRequest := *openapiclient.NewPkiSignVerbatimRequest() // PkiSignVerbatimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiSignVerbatim(context.Background()).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiSignVerbatim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiSignVerbatimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiSignVerbatimRole

> PostPkiSignVerbatimRole(ctx, role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The desired role with configuration for this request
    pkiSignVerbatimRequest := *openapiclient.NewPkiSignVerbatimRequest() // PkiSignVerbatimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiSignVerbatimRole(context.Background(), role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiSignVerbatimRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiSignVerbatimRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPkiTidy

> PostPkiTidy(ctx).PkiTidyRequest(pkiTidyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkiTidyRequest := *openapiclient.NewPkiTidyRequest() // PkiTidyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostPkiTidy(context.Background()).PkiTidyRequest(pkiTidyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostPkiTidy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPkiTidyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiTidyRequest** | [**PkiTidyRequest**](PkiTidyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRabbitmqConfigConnection

> PostRabbitmqConfigConnection(ctx).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()

Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rabbitmqConfigConnectionRequest := *openapiclient.NewRabbitmqConfigConnectionRequest() // RabbitmqConfigConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostRabbitmqConfigConnection(context.Background()).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostRabbitmqConfigConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRabbitmqConfigConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigConnectionRequest** | [**RabbitmqConfigConnectionRequest**](RabbitmqConfigConnectionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRabbitmqConfigLease

> PostRabbitmqConfigLease(ctx).RabbitmqConfigLeaseRequest(rabbitmqConfigLeaseRequest).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rabbitmqConfigLeaseRequest := *openapiclient.NewRabbitmqConfigLeaseRequest() // RabbitmqConfigLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostRabbitmqConfigLease(context.Background()).RabbitmqConfigLeaseRequest(rabbitmqConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostRabbitmqConfigLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRabbitmqConfigLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigLeaseRequest** | [**RabbitmqConfigLeaseRequest**](RabbitmqConfigLeaseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRabbitmqRolesName

> PostRabbitmqRolesName(ctx, name).RabbitmqRolesRequest(rabbitmqRolesRequest).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role.
    rabbitmqRolesRequest := *openapiclient.NewRabbitmqRolesRequest() // RabbitmqRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostRabbitmqRolesName(context.Background(), name).RabbitmqRolesRequest(rabbitmqRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostRabbitmqRolesName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRabbitmqRolesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitmqRolesRequest** | [**RabbitmqRolesRequest**](RabbitmqRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretConfig

> PostSecretConfig(ctx).KvConfigRequest(kvConfigRequest).Execute()

Configure backend level settings that are applied to every key in the key-value store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kvConfigRequest := *openapiclient.NewKvConfigRequest() // KvConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretConfig(context.Background()).KvConfigRequest(kvConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvConfigRequest** | [**KvConfigRequest**](KvConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretDataPath

> PostSecretDataPath(ctx, path).KvDataRequest(kvDataRequest).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    kvDataRequest := *openapiclient.NewKvDataRequest() // KvDataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretDataPath(context.Background(), path).KvDataRequest(kvDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretDataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretDataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDataRequest** | [**KvDataRequest**](KvDataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretDeletePath

> PostSecretDeletePath(ctx, path).KvDeleteRequest(kvDeleteRequest).Execute()

Marks one or more versions as deleted in the KV store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    kvDeleteRequest := *openapiclient.NewKvDeleteRequest() // KvDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretDeletePath(context.Background(), path).KvDeleteRequest(kvDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretDeletePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretDeletePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDeleteRequest** | [**KvDeleteRequest**](KvDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretDestroyPath

> PostSecretDestroyPath(ctx, path).KvDestroyRequest(kvDestroyRequest).Execute()

Permanently removes one or more versions in the KV store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    kvDestroyRequest := *openapiclient.NewKvDestroyRequest() // KvDestroyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretDestroyPath(context.Background(), path).KvDestroyRequest(kvDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretDestroyPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretDestroyPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDestroyRequest** | [**KvDestroyRequest**](KvDestroyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretMetadataPath

> PostSecretMetadataPath(ctx, path).KvMetadataRequest(kvMetadataRequest).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    kvMetadataRequest := *openapiclient.NewKvMetadataRequest() // KvMetadataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretMetadataPath(context.Background(), path).KvMetadataRequest(kvMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretMetadataPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretMetadataPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvMetadataRequest** | [**KvMetadataRequest**](KvMetadataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretUndeletePath

> PostSecretUndeletePath(ctx, path).KvUndeleteRequest(kvUndeleteRequest).Execute()

Undeletes one or more versions from the KV store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Location of the secret.
    kvUndeleteRequest := *openapiclient.NewKvUndeleteRequest() // KvUndeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSecretUndeletePath(context.Background(), path).KvUndeleteRequest(kvUndeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSecretUndeletePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Location of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretUndeletePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvUndeleteRequest** | [**KvUndeleteRequest**](KvUndeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshConfigCa

> PostSshConfigCa(ctx).SshConfigCaRequest(sshConfigCaRequest).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sshConfigCaRequest := *openapiclient.NewSshConfigCaRequest() // SshConfigCaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshConfigCa(context.Background()).SshConfigCaRequest(sshConfigCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshConfigCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshConfigCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigCaRequest** | [**SshConfigCaRequest**](SshConfigCaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshConfigZeroaddress

> PostSshConfigZeroaddress(ctx).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sshConfigZeroaddressRequest := *openapiclient.NewSshConfigZeroaddressRequest() // SshConfigZeroaddressRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshConfigZeroaddress(context.Background()).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshConfigZeroaddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshConfigZeroaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigZeroaddressRequest** | [**SshConfigZeroaddressRequest**](SshConfigZeroaddressRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshCredsRole

> PostSshCredsRole(ctx, role).SshCredsRequest(sshCredsRequest).Execute()

Creates a credential for establishing SSH connection with the remote host.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | [Required] Name of the role
    sshCredsRequest := *openapiclient.NewSshCredsRequest() // SshCredsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshCredsRole(context.Background(), role).SshCredsRequest(sshCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshCredsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | [Required] Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSshCredsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshCredsRequest** | [**SshCredsRequest**](SshCredsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshKeysKeyName

> PostSshKeysKeyName(ctx, keyName).SshKeysRequest(sshKeysRequest).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyName := "keyName_example" // string | [Required] Name of the key
    sshKeysRequest := *openapiclient.NewSshKeysRequest() // SshKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshKeysKeyName(context.Background(), keyName).SshKeysRequest(sshKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshKeysKeyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** | [Required] Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSshKeysKeyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeysRequest** | [**SshKeysRequest**](SshKeysRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshLookup

> PostSshLookup(ctx).SshLookupRequest(sshLookupRequest).Execute()

List all the roles associated with the given IP address.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sshLookupRequest := *openapiclient.NewSshLookupRequest() // SshLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshLookup(context.Background()).SshLookupRequest(sshLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshLookupRequest** | [**SshLookupRequest**](SshLookupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshRolesRole

> PostSshRolesRole(ctx, role).SshRolesRequest(sshRolesRequest).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | [Required for all types] Name of the role being created.
    sshRolesRequest := *openapiclient.NewSshRolesRequest() // SshRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshRolesRole(context.Background(), role).SshRolesRequest(sshRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSshRolesRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshRolesRequest** | [**SshRolesRequest**](SshRolesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshSignRole

> PostSshSignRole(ctx, role).SshSignRequest(sshSignRequest).Execute()

Request signing an SSH key using a certain role with the provided details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The desired role with configuration for this request.
    sshSignRequest := *openapiclient.NewSshSignRequest() // SshSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshSignRole(context.Background(), role).SshSignRequest(sshSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshSignRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The desired role with configuration for this request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSshSignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshSignRequest** | [**SshSignRequest**](SshSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshVerify

> PostSshVerify(ctx).SshVerifyRequest(sshVerifyRequest).Execute()

Validate the OTP provided by Vault SSH Agent.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sshVerifyRequest := *openapiclient.NewSshVerifyRequest() // SshVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostSshVerify(context.Background()).SshVerifyRequest(sshVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostSshVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshVerifyRequest** | [**SshVerifyRequest**](SshVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTerraformConfig

> PostTerraformConfig(ctx).TerraformConfigRequest(terraformConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    terraformConfigRequest := *openapiclient.NewTerraformConfigRequest("Token_example") // TerraformConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTerraformConfig(context.Background()).TerraformConfigRequest(terraformConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTerraformConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTerraformConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformConfigRequest** | [**TerraformConfigRequest**](TerraformConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTerraformCredsName

> PostTerraformCredsName(ctx, name).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTerraformCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTerraformCredsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTerraformCredsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTerraformRoleName

> PostTerraformRoleName(ctx, name).TerraformRoleRequest(terraformRoleRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    terraformRoleRequest := *openapiclient.NewTerraformRoleRequest() // TerraformRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTerraformRoleName(context.Background(), name).TerraformRoleRequest(terraformRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTerraformRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTerraformRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformRoleRequest** | [**TerraformRoleRequest**](TerraformRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTerraformRotateRoleName

> PostTerraformRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the team or organization role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTerraformRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTerraformRotateRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the team or organization role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTerraformRotateRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTotpCodeName

> PostTotpCodeName(ctx, name).TotpCodeRequest(totpCodeRequest).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key.
    totpCodeRequest := *openapiclient.NewTotpCodeRequest() // TotpCodeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTotpCodeName(context.Background(), name).TotpCodeRequest(totpCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTotpCodeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTotpCodeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totpCodeRequest** | [**TotpCodeRequest**](TotpCodeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTotpKeysName

> PostTotpKeysName(ctx, name).TotpKeysRequest(totpKeysRequest).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key.
    totpKeysRequest := *openapiclient.NewTotpKeysRequest() // TotpKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTotpKeysName(context.Background(), name).TotpKeysRequest(totpKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTotpKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTotpKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totpKeysRequest** | [**TotpKeysRequest**](TotpKeysRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitCacheConfig

> PostTransitCacheConfig(ctx).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()

Configures a new cache of the specified size

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transitCacheConfigRequest := *openapiclient.NewTransitCacheConfigRequest() // TransitCacheConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitCacheConfig(context.Background()).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitCacheConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitCacheConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitCacheConfigRequest** | [**TransitCacheConfigRequest**](TransitCacheConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitDatakeyPlaintextName

> PostTransitDatakeyPlaintextName(ctx, name, plaintext).TransitDatakeyRequest(transitDatakeyRequest).Execute()

Generate a data key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The backend key used for encrypting the data key
    plaintext := "plaintext_example" // string | \"plaintext\" will return the key in both plaintext and ciphertext; \"wrapped\" will return the ciphertext only.
    transitDatakeyRequest := *openapiclient.NewTransitDatakeyRequest() // TransitDatakeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitDatakeyPlaintextName(context.Background(), name, plaintext).TransitDatakeyRequest(transitDatakeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitDatakeyPlaintextName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The backend key used for encrypting the data key | 
**plaintext** | **string** | \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitDatakeyPlaintextNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitDatakeyRequest** | [**TransitDatakeyRequest**](TransitDatakeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitDecryptName

> PostTransitDecryptName(ctx, name).TransitDecryptRequest(transitDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the policy
    transitDecryptRequest := *openapiclient.NewTransitDecryptRequest() // TransitDecryptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitDecryptName(context.Background(), name).TransitDecryptRequest(transitDecryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitDecryptName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitDecryptNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitEncryptName

> PostTransitEncryptName(ctx, name).TransitEncryptRequest(transitEncryptRequest).Execute()

Encrypt a plaintext value or a batch of plaintext blocks using a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the policy
    transitEncryptRequest := *openapiclient.NewTransitEncryptRequest() // TransitEncryptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitEncryptName(context.Background(), name).TransitEncryptRequest(transitEncryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitEncryptName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitEncryptNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitHash

> PostTransitHash(ctx).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transitHashRequest := *openapiclient.NewTransitHashRequest() // TransitHashRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitHash(context.Background()).TransitHashRequest(transitHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitHashUrlalgorithm

> PostTransitHashUrlalgorithm(ctx, urlalgorithm).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
    transitHashRequest := *openapiclient.NewTransitHashRequest() // TransitHashRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitHashUrlalgorithm(context.Background(), urlalgorithm).TransitHashRequest(transitHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitHashUrlalgorithm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitHashUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitHmacName

> PostTransitHmacName(ctx, name).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use for the HMAC function
    transitHmacRequest := *openapiclient.NewTransitHmacRequest() // TransitHmacRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitHmacName(context.Background(), name).TransitHmacRequest(transitHmacRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitHmacName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use for the HMAC function | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitHmacNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitHmacNameUrlalgorithm

> PostTransitHmacNameUrlalgorithm(ctx, name, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use for the HMAC function
    urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
    transitHmacRequest := *openapiclient.NewTransitHmacRequest() // TransitHmacRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitHmacNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitHmacNameUrlalgorithm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use for the HMAC function | 
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitHmacNameUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysName

> PostTransitKeysName(ctx, name).TransitKeysRequest(transitKeysRequest).Execute()

Managed named encryption keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    transitKeysRequest := *openapiclient.NewTransitKeysRequest() // TransitKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysName(context.Background(), name).TransitKeysRequest(transitKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysRequest** | [**TransitKeysRequest**](TransitKeysRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysNameConfig

> PostTransitKeysNameConfig(ctx, name).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()

Configure a named encryption key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    transitKeysConfigRequest := *openapiclient.NewTransitKeysConfigRequest() // TransitKeysConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysNameConfig(context.Background(), name).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysNameConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysConfigRequest** | [**TransitKeysConfigRequest**](TransitKeysConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysNameImport

> PostTransitKeysNameImport(ctx, name).TransitKeysImportRequest(transitKeysImportRequest).Execute()

Imports an externally-generated key into a new transit key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the key
    transitKeysImportRequest := *openapiclient.NewTransitKeysImportRequest() // TransitKeysImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysNameImport(context.Background(), name).TransitKeysImportRequest(transitKeysImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysNameImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysImportRequest** | [**TransitKeysImportRequest**](TransitKeysImportRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysNameImportVersion

> PostTransitKeysNameImportVersion(ctx, name).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()

Imports an externally-generated key into an existing imported key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the key
    transitKeysImportVersionRequest := *openapiclient.NewTransitKeysImportVersionRequest() // TransitKeysImportVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysNameImportVersion(context.Background(), name).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysNameImportVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameImportVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysImportVersionRequest** | [**TransitKeysImportVersionRequest**](TransitKeysImportVersionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysNameRotate

> PostTransitKeysNameRotate(ctx, name).Execute()

Rotate named encryption key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysNameRotate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysNameRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitKeysNameTrim

> PostTransitKeysNameTrim(ctx, name).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()

Trim key versions of a named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    transitKeysTrimRequest := *openapiclient.NewTransitKeysTrimRequest() // TransitKeysTrimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitKeysNameTrim(context.Background(), name).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitKeysNameTrim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitKeysNameTrimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysTrimRequest** | [**TransitKeysTrimRequest**](TransitKeysTrimRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRandom

> PostTransitRandom(ctx).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transitRandomRequest := *openapiclient.NewTransitRandomRequest() // TransitRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRandom(context.Background()).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRandom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRandomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRandomSource

> PostTransitRandomSource(ctx, source).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (default to "platform")
    transitRandomRequest := *openapiclient.NewTransitRandomRequest() // TransitRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRandomSource(context.Background(), source).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRandomSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRandomSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRandomSourceUrlbytes

> PostTransitRandomSourceUrlbytes(ctx, source, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (default to "platform")
    urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
    transitRandomRequest := *openapiclient.NewTransitRandomRequest() // TransitRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRandomSourceUrlbytes(context.Background(), source, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRandomSourceUrlbytes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRandomSourceUrlbytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRandomUrlbytes

> PostTransitRandomUrlbytes(ctx, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
    transitRandomRequest := *openapiclient.NewTransitRandomRequest() // TransitRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRandomUrlbytes(context.Background(), urlbytes).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRandomUrlbytes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRandomUrlbytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRestore

> PostTransitRestore(ctx).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transitRestoreRequest := *openapiclient.NewTransitRestoreRequest() // TransitRestoreRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRestore(context.Background()).TransitRestoreRequest(transitRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRestoreName

> PostTransitRestoreName(ctx, name).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | If set, this will be the name of the restored key.
    transitRestoreRequest := *openapiclient.NewTransitRestoreRequest() // TransitRestoreRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRestoreName(context.Background(), name).TransitRestoreRequest(transitRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRestoreName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | If set, this will be the name of the restored key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRestoreNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitRewrapName

> PostTransitRewrapName(ctx, name).TransitRewrapRequest(transitRewrapRequest).Execute()

Rewrap ciphertext

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    transitRewrapRequest := *openapiclient.NewTransitRewrapRequest() // TransitRewrapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitRewrapName(context.Background(), name).TransitRewrapRequest(transitRewrapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitRewrapName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitRewrapNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRewrapRequest** | [**TransitRewrapRequest**](TransitRewrapRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitSignName

> PostTransitSignName(ctx, name).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use
    transitSignRequest := *openapiclient.NewTransitSignRequest() // TransitSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitSignName(context.Background(), name).TransitSignRequest(transitSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitSignName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitSignNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitSignNameUrlalgorithm

> PostTransitSignNameUrlalgorithm(ctx, name, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use
    urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
    transitSignRequest := *openapiclient.NewTransitSignRequest() // TransitSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitSignNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitSignNameUrlalgorithm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitSignNameUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitVerifyName

> PostTransitVerifyName(ctx, name).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use
    transitVerifyRequest := *openapiclient.NewTransitVerifyRequest() // TransitVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitVerifyName(context.Background(), name).TransitVerifyRequest(transitVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitVerifyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitVerifyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransitVerifyNameUrlalgorithm

> PostTransitVerifyNameUrlalgorithm(ctx, name, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The key to use
    urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
    transitVerifyRequest := *openapiclient.NewTransitVerifyRequest() // TransitVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.PostTransitVerifyNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PostTransitVerifyNameUrlalgorithm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransitVerifyNameUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

