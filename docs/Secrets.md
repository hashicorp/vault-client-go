# Secrets

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdConfig**](Secrets.md#DeleteAdConfig) | **Delete** /ad/config | Configure the AD server to connect to, along with password options.
[**DeleteAdLibraryName**](Secrets.md#DeleteAdLibraryName) | **Delete** /ad/library/{name} | Delete a library set.
[**DeleteAdRolesName**](Secrets.md#DeleteAdRolesName) | **Delete** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**DeleteAlicloudConfig**](Secrets.md#DeleteAlicloudConfig) | **Delete** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**DeleteAlicloudRoleName**](Secrets.md#DeleteAlicloudRoleName) | **Delete** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**DeleteAwsRolesName**](Secrets.md#DeleteAwsRolesName) | **Delete** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**DeleteAzureConfig**](Secrets.md#DeleteAzureConfig) | **Delete** /azure/config | 
[**DeleteAzureRolesName**](Secrets.md#DeleteAzureRolesName) | **Delete** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**DeleteConsulRolesName**](Secrets.md#DeleteConsulRolesName) | **Delete** /consul/roles/{name} | 
[**DeleteCubbyholePath**](Secrets.md#DeleteCubbyholePath) | **Delete** /cubbyhole/{path} | Deletes the secret at the specified location.
[**DeleteGcpRolesetName**](Secrets.md#DeleteGcpRolesetName) | **Delete** /gcp/roleset/{name} | 
[**DeleteGcpStaticAccountName**](Secrets.md#DeleteGcpStaticAccountName) | **Delete** /gcp/static-account/{name} | 
[**DeleteGcpkmsConfig**](Secrets.md#DeleteGcpkmsConfig) | **Delete** /gcpkms/config | Configure the GCP KMS secrets engine
[**DeleteGcpkmsKeysDeregisterKey**](Secrets.md#DeleteGcpkmsKeysDeregisterKey) | **Delete** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
[**DeleteGcpkmsKeysKey**](Secrets.md#DeleteGcpkmsKeysKey) | **Delete** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**DeleteGcpkmsKeysTrimKey**](Secrets.md#DeleteGcpkmsKeysTrimKey) | **Delete** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**DeleteKubernetesConfig**](Secrets.md#DeleteKubernetesConfig) | **Delete** /kubernetes/config | 
[**DeleteKubernetesRolesName**](Secrets.md#DeleteKubernetesRolesName) | **Delete** /kubernetes/roles/{name} | 
[**DeleteKvPath**](Secrets.md#DeleteKvPath) | **Delete** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**DeleteMongodbatlasRolesName**](Secrets.md#DeleteMongodbatlasRolesName) | **Delete** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**DeleteNomadConfigAccess**](Secrets.md#DeleteNomadConfigAccess) | **Delete** /nomad/config/access | 
[**DeleteNomadConfigLease**](Secrets.md#DeleteNomadConfigLease) | **Delete** /nomad/config/lease | Configure the lease parameters for generated tokens
[**DeleteNomadRoleName**](Secrets.md#DeleteNomadRoleName) | **Delete** /nomad/role/{name} | 
[**DeleteOpenldapConfig**](Secrets.md#DeleteOpenldapConfig) | **Delete** /openldap/config | 
[**DeleteOpenldapRoleName**](Secrets.md#DeleteOpenldapRoleName) | **Delete** /openldap/role/{name} | 
[**DeleteOpenldapStaticRoleName**](Secrets.md#DeleteOpenldapStaticRoleName) | **Delete** /openldap/static-role/{name} | 
[**DeletePkiIssuerRefDerPem**](Secrets.md#DeletePkiIssuerRefDerPem) | **Delete** /pki/{issuer_ref}/der|/pem | 
[**DeletePkiJson**](Secrets.md#DeletePkiJson) | **Delete** /pki//json | 
[**DeletePkiKeyKeyRef**](Secrets.md#DeletePkiKeyKeyRef) | **Delete** /pki/key/{key_ref} | 
[**DeletePkiRolesName**](Secrets.md#DeletePkiRolesName) | **Delete** /pki/roles/{name} | 
[**DeletePkiRoot**](Secrets.md#DeletePkiRoot) | **Delete** /pki/root | 
[**DeleteRabbitmqRolesName**](Secrets.md#DeleteRabbitmqRolesName) | **Delete** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**DeleteSecretDataPath**](Secrets.md#DeleteSecretDataPath) | **Delete** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**DeleteSecretMetadataPath**](Secrets.md#DeleteSecretMetadataPath) | **Delete** /secret/metadata/{path} | Configures settings for the KV store
[**DeleteSshConfigCa**](Secrets.md#DeleteSshConfigCa) | **Delete** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**DeleteSshConfigZeroaddress**](Secrets.md#DeleteSshConfigZeroaddress) | **Delete** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**DeleteSshKeysKeyName**](Secrets.md#DeleteSshKeysKeyName) | **Delete** /ssh/keys/{key_name} | Register a shared private key with Vault.
[**DeleteSshRolesRole**](Secrets.md#DeleteSshRolesRole) | **Delete** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**DeleteTerraformConfig**](Secrets.md#DeleteTerraformConfig) | **Delete** /terraform/config | 
[**DeleteTerraformRoleName**](Secrets.md#DeleteTerraformRoleName) | **Delete** /terraform/role/{name} | 
[**DeleteTotpKeysName**](Secrets.md#DeleteTotpKeysName) | **Delete** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**DeleteTransitKeysName**](Secrets.md#DeleteTransitKeysName) | **Delete** /transit/keys/{name} | Managed named encryption keys
[**GetAdConfig**](Secrets.md#GetAdConfig) | **Get** /ad/config | Configure the AD server to connect to, along with password options.
[**GetAdCredsName**](Secrets.md#GetAdCredsName) | **Get** /ad/creds/{name} | 
[**GetAdLibrary**](Secrets.md#GetAdLibrary) | **Get** /ad/library | 
[**GetAdLibraryName**](Secrets.md#GetAdLibraryName) | **Get** /ad/library/{name} | Read a library set.
[**GetAdLibraryNameStatus**](Secrets.md#GetAdLibraryNameStatus) | **Get** /ad/library/{name}/status | Check the status of the service accounts in a library set.
[**GetAdRoles**](Secrets.md#GetAdRoles) | **Get** /ad/roles | List the name of each role currently stored.
[**GetAdRolesName**](Secrets.md#GetAdRolesName) | **Get** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**GetAdRotateRoot**](Secrets.md#GetAdRotateRoot) | **Get** /ad/rotate-root | 
[**GetAlicloudConfig**](Secrets.md#GetAlicloudConfig) | **Get** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**GetAlicloudCredsName**](Secrets.md#GetAlicloudCredsName) | **Get** /alicloud/creds/{name} | Generate an API key or STS credential using the given role&#39;s configuration.&#39;
[**GetAlicloudRole**](Secrets.md#GetAlicloudRole) | **Get** /alicloud/role | List the existing roles in this backend.
[**GetAlicloudRoleName**](Secrets.md#GetAlicloudRoleName) | **Get** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**GetAwsConfigLease**](Secrets.md#GetAwsConfigLease) | **Get** /aws/config/lease | Configure the default lease information for generated credentials.
[**GetAwsConfigRoot**](Secrets.md#GetAwsConfigRoot) | **Get** /aws/config/root | Configure the root credentials that are used to manage IAM.
[**GetAwsCreds**](Secrets.md#GetAwsCreds) | **Get** /aws/creds | Generate AWS credentials from a specific Vault role.
[**GetAwsRoles**](Secrets.md#GetAwsRoles) | **Get** /aws/roles | List the existing roles in this backend
[**GetAwsRolesName**](Secrets.md#GetAwsRolesName) | **Get** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**GetAwsStsName**](Secrets.md#GetAwsStsName) | **Get** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
[**GetAzureConfig**](Secrets.md#GetAzureConfig) | **Get** /azure/config | 
[**GetAzureCredsRole**](Secrets.md#GetAzureCredsRole) | **Get** /azure/creds/{role} | 
[**GetAzureRoles**](Secrets.md#GetAzureRoles) | **Get** /azure/roles | List existing roles.
[**GetAzureRolesName**](Secrets.md#GetAzureRolesName) | **Get** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**GetConsulConfigAccess**](Secrets.md#GetConsulConfigAccess) | **Get** /consul/config/access | 
[**GetConsulCredsRole**](Secrets.md#GetConsulCredsRole) | **Get** /consul/creds/{role} | 
[**GetConsulRoles**](Secrets.md#GetConsulRoles) | **Get** /consul/roles | 
[**GetConsulRolesName**](Secrets.md#GetConsulRolesName) | **Get** /consul/roles/{name} | 
[**GetCubbyholePath**](Secrets.md#GetCubbyholePath) | **Get** /cubbyhole/{path} | Retrieve the secret at the specified location.
[**GetGcpConfig**](Secrets.md#GetGcpConfig) | **Get** /gcp/config | 
[**GetGcpKeyRoleset**](Secrets.md#GetGcpKeyRoleset) | **Get** /gcp/key/{roleset} | 
[**GetGcpRolesetName**](Secrets.md#GetGcpRolesetName) | **Get** /gcp/roleset/{name} | 
[**GetGcpRolesetRolesetKey**](Secrets.md#GetGcpRolesetRolesetKey) | **Get** /gcp/roleset/{roleset}/key | 
[**GetGcpRolesetRolesetToken**](Secrets.md#GetGcpRolesetRolesetToken) | **Get** /gcp/roleset/{roleset}/token | 
[**GetGcpRolesets**](Secrets.md#GetGcpRolesets) | **Get** /gcp/rolesets | 
[**GetGcpStaticAccountName**](Secrets.md#GetGcpStaticAccountName) | **Get** /gcp/static-account/{name} | 
[**GetGcpStaticAccountNameKey**](Secrets.md#GetGcpStaticAccountNameKey) | **Get** /gcp/static-account/{name}/key | 
[**GetGcpStaticAccountNameToken**](Secrets.md#GetGcpStaticAccountNameToken) | **Get** /gcp/static-account/{name}/token | 
[**GetGcpStaticAccounts**](Secrets.md#GetGcpStaticAccounts) | **Get** /gcp/static-accounts | 
[**GetGcpTokenRoleset**](Secrets.md#GetGcpTokenRoleset) | **Get** /gcp/token/{roleset} | 
[**GetGcpkmsConfig**](Secrets.md#GetGcpkmsConfig) | **Get** /gcpkms/config | Configure the GCP KMS secrets engine
[**GetGcpkmsKeys**](Secrets.md#GetGcpkmsKeys) | **Get** /gcpkms/keys | List named keys
[**GetGcpkmsKeysConfigKey**](Secrets.md#GetGcpkmsKeysConfigKey) | **Get** /gcpkms/keys/config/{key} | Configure the key in Vault
[**GetGcpkmsKeysKey**](Secrets.md#GetGcpkmsKeysKey) | **Get** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GetGcpkmsPubkeyKey**](Secrets.md#GetGcpkmsPubkeyKey) | **Get** /gcpkms/pubkey/{key} | Retrieve the public key associated with the named key
[**GetKubernetesConfig**](Secrets.md#GetKubernetesConfig) | **Get** /kubernetes/config | 
[**GetKubernetesRoles**](Secrets.md#GetKubernetesRoles) | **Get** /kubernetes/roles | 
[**GetKubernetesRolesName**](Secrets.md#GetKubernetesRolesName) | **Get** /kubernetes/roles/{name} | 
[**GetKvPath**](Secrets.md#GetKvPath) | **Get** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**GetMongodbatlasConfig**](Secrets.md#GetMongodbatlasConfig) | **Get** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
[**GetMongodbatlasCredsName**](Secrets.md#GetMongodbatlasCredsName) | **Get** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**GetMongodbatlasRoles**](Secrets.md#GetMongodbatlasRoles) | **Get** /mongodbatlas/roles | List the existing roles in this backend
[**GetMongodbatlasRolesName**](Secrets.md#GetMongodbatlasRolesName) | **Get** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**GetNomadConfigAccess**](Secrets.md#GetNomadConfigAccess) | **Get** /nomad/config/access | 
[**GetNomadConfigLease**](Secrets.md#GetNomadConfigLease) | **Get** /nomad/config/lease | Configure the lease parameters for generated tokens
[**GetNomadCredsName**](Secrets.md#GetNomadCredsName) | **Get** /nomad/creds/{name} | 
[**GetNomadRole**](Secrets.md#GetNomadRole) | **Get** /nomad/role | 
[**GetNomadRoleName**](Secrets.md#GetNomadRoleName) | **Get** /nomad/role/{name} | 
[**GetOpenldapConfig**](Secrets.md#GetOpenldapConfig) | **Get** /openldap/config | 
[**GetOpenldapCredsName**](Secrets.md#GetOpenldapCredsName) | **Get** /openldap/creds/{name} | 
[**GetOpenldapRole**](Secrets.md#GetOpenldapRole) | **Get** /openldap/role | 
[**GetOpenldapRoleName**](Secrets.md#GetOpenldapRoleName) | **Get** /openldap/role/{name} | 
[**GetOpenldapStaticCredName**](Secrets.md#GetOpenldapStaticCredName) | **Get** /openldap/static-cred/{name} | 
[**GetOpenldapStaticRole**](Secrets.md#GetOpenldapStaticRole) | **Get** /openldap/static-role | 
[**GetOpenldapStaticRoleName**](Secrets.md#GetOpenldapStaticRoleName) | **Get** /openldap/static-role/{name} | 
[**GetPkiCa**](Secrets.md#GetPkiCa) | **Get** /pki/ca | 
[**GetPkiCaChain**](Secrets.md#GetPkiCaChain) | **Get** /pki/ca_chain | 
[**GetPkiCaPem**](Secrets.md#GetPkiCaPem) | **Get** /pki/ca/pem | 
[**GetPkiCertCaChain**](Secrets.md#GetPkiCertCaChain) | **Get** /pki/cert/ca_chain | 
[**GetPkiCertCrl**](Secrets.md#GetPkiCertCrl) | **Get** /pki/cert/crl | 
[**GetPkiCertSerial**](Secrets.md#GetPkiCertSerial) | **Get** /pki/cert/{serial} | 
[**GetPkiCertSerialRaw**](Secrets.md#GetPkiCertSerialRaw) | **Get** /pki/cert/{serial}/raw | 
[**GetPkiCertSerialRawPem**](Secrets.md#GetPkiCertSerialRawPem) | **Get** /pki/cert/{serial}/raw/pem | 
[**GetPkiCerts**](Secrets.md#GetPkiCerts) | **Get** /pki/certs | 
[**GetPkiConfigCrl**](Secrets.md#GetPkiConfigCrl) | **Get** /pki/config/crl | 
[**GetPkiConfigIssuers**](Secrets.md#GetPkiConfigIssuers) | **Get** /pki/config/issuers | 
[**GetPkiConfigKeys**](Secrets.md#GetPkiConfigKeys) | **Get** /pki/config/keys | 
[**GetPkiConfigUrls**](Secrets.md#GetPkiConfigUrls) | **Get** /pki/config/urls | 
[**GetPkiCrl**](Secrets.md#GetPkiCrl) | **Get** /pki/crl | 
[**GetPkiCrlPem**](Secrets.md#GetPkiCrlPem) | **Get** /pki/crl/pem | 
[**GetPkiCrlRotate**](Secrets.md#GetPkiCrlRotate) | **Get** /pki/crl/rotate | 
[**GetPkiDer**](Secrets.md#GetPkiDer) | **Get** /pki//der | 
[**GetPkiIssuerRefCrlPem**](Secrets.md#GetPkiIssuerRefCrlPem) | **Get** /pki/{issuer_ref}/crl/pem | 
[**GetPkiIssuerRefDerPem**](Secrets.md#GetPkiIssuerRefDerPem) | **Get** /pki/{issuer_ref}/der|/pem | 
[**GetPkiIssuers**](Secrets.md#GetPkiIssuers) | **Get** /pki/issuers | 
[**GetPkiJson**](Secrets.md#GetPkiJson) | **Get** /pki//json | 
[**GetPkiKeyKeyRef**](Secrets.md#GetPkiKeyKeyRef) | **Get** /pki/key/{key_ref} | 
[**GetPkiKeys**](Secrets.md#GetPkiKeys) | **Get** /pki/keys | 
[**GetPkiRoles**](Secrets.md#GetPkiRoles) | **Get** /pki/roles | 
[**GetPkiRolesName**](Secrets.md#GetPkiRolesName) | **Get** /pki/roles/{name} | 
[**GetPkiTidyStatus**](Secrets.md#GetPkiTidyStatus) | **Get** /pki/tidy-status | 
[**GetRabbitmqConfigLease**](Secrets.md#GetRabbitmqConfigLease) | **Get** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
[**GetRabbitmqCredsName**](Secrets.md#GetRabbitmqCredsName) | **Get** /rabbitmq/creds/{name} | Request RabbitMQ credentials for a certain role.
[**GetRabbitmqRoles**](Secrets.md#GetRabbitmqRoles) | **Get** /rabbitmq/roles | Manage the roles that can be created with this backend.
[**GetRabbitmqRolesName**](Secrets.md#GetRabbitmqRolesName) | **Get** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**GetSecretConfig**](Secrets.md#GetSecretConfig) | **Get** /secret/config | Read the backend level settings.
[**GetSecretDataPath**](Secrets.md#GetSecretDataPath) | **Get** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**GetSecretMetadataPath**](Secrets.md#GetSecretMetadataPath) | **Get** /secret/metadata/{path} | Configures settings for the KV store
[**GetSecretSubkeysPath**](Secrets.md#GetSecretSubkeysPath) | **Get** /secret/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
[**GetSshConfigCa**](Secrets.md#GetSshConfigCa) | **Get** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**GetSshConfigZeroaddress**](Secrets.md#GetSshConfigZeroaddress) | **Get** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**GetSshPublicKey**](Secrets.md#GetSshPublicKey) | **Get** /ssh/public_key | Retrieve the public key.
[**GetSshRoles**](Secrets.md#GetSshRoles) | **Get** /ssh/roles | Manage the &#39;roles&#39; that can be created with this backend.
[**GetSshRolesRole**](Secrets.md#GetSshRolesRole) | **Get** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**GetTerraformConfig**](Secrets.md#GetTerraformConfig) | **Get** /terraform/config | 
[**GetTerraformCredsName**](Secrets.md#GetTerraformCredsName) | **Get** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**GetTerraformRole**](Secrets.md#GetTerraformRole) | **Get** /terraform/role | 
[**GetTerraformRoleName**](Secrets.md#GetTerraformRoleName) | **Get** /terraform/role/{name} | 
[**GetTotpCodeName**](Secrets.md#GetTotpCodeName) | **Get** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**GetTotpKeys**](Secrets.md#GetTotpKeys) | **Get** /totp/keys | Manage the keys that can be created with this backend.
[**GetTotpKeysName**](Secrets.md#GetTotpKeysName) | **Get** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**GetTransitBackupName**](Secrets.md#GetTransitBackupName) | **Get** /transit/backup/{name} | Backup the named key
[**GetTransitCacheConfig**](Secrets.md#GetTransitCacheConfig) | **Get** /transit/cache-config | Returns the size of the active cache
[**GetTransitExportTypeName**](Secrets.md#GetTransitExportTypeName) | **Get** /transit/export/{type}/{name} | Export named encryption or signing key
[**GetTransitExportTypeNameVersion**](Secrets.md#GetTransitExportTypeNameVersion) | **Get** /transit/export/{type}/{name}/{version} | Export named encryption or signing key
[**GetTransitKeys**](Secrets.md#GetTransitKeys) | **Get** /transit/keys | Managed named encryption keys
[**GetTransitKeysName**](Secrets.md#GetTransitKeysName) | **Get** /transit/keys/{name} | Managed named encryption keys
[**GetTransitWrappingKey**](Secrets.md#GetTransitWrappingKey) | **Get** /transit/wrapping_key | Returns the public key to use for wrapping imported keys
[**PostAdConfig**](Secrets.md#PostAdConfig) | **Post** /ad/config | Configure the AD server to connect to, along with password options.
[**PostAdLibraryManageNameCheckIn**](Secrets.md#PostAdLibraryManageNameCheckIn) | **Post** /ad/library/manage/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryName**](Secrets.md#PostAdLibraryName) | **Post** /ad/library/{name} | Update a library set.
[**PostAdLibraryNameCheckIn**](Secrets.md#PostAdLibraryNameCheckIn) | **Post** /ad/library/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryNameCheckOut**](Secrets.md#PostAdLibraryNameCheckOut) | **Post** /ad/library/{name}/check-out | Check a service account out from the library.
[**PostAdRolesName**](Secrets.md#PostAdRolesName) | **Post** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**PostAdRotateRoleName**](Secrets.md#PostAdRotateRoleName) | **Post** /ad/rotate-role/{name} | 
[**PostAdRotateRoot**](Secrets.md#PostAdRotateRoot) | **Post** /ad/rotate-root | 
[**PostAlicloudConfig**](Secrets.md#PostAlicloudConfig) | **Post** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
[**PostAlicloudRoleName**](Secrets.md#PostAlicloudRoleName) | **Post** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**PostAwsConfigLease**](Secrets.md#PostAwsConfigLease) | **Post** /aws/config/lease | Configure the default lease information for generated credentials.
[**PostAwsConfigRoot**](Secrets.md#PostAwsConfigRoot) | **Post** /aws/config/root | Configure the root credentials that are used to manage IAM.
[**PostAwsConfigRotateRoot**](Secrets.md#PostAwsConfigRotateRoot) | **Post** /aws/config/rotate-root | 
[**PostAwsCreds**](Secrets.md#PostAwsCreds) | **Post** /aws/creds | Generate AWS credentials from a specific Vault role.
[**PostAwsRolesName**](Secrets.md#PostAwsRolesName) | **Post** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**PostAwsStsName**](Secrets.md#PostAwsStsName) | **Post** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
[**PostAzureConfig**](Secrets.md#PostAzureConfig) | **Post** /azure/config | 
[**PostAzureRolesName**](Secrets.md#PostAzureRolesName) | **Post** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**PostAzureRotateRoot**](Secrets.md#PostAzureRotateRoot) | **Post** /azure/rotate-root | 
[**PostConsulConfigAccess**](Secrets.md#PostConsulConfigAccess) | **Post** /consul/config/access | 
[**PostConsulRolesName**](Secrets.md#PostConsulRolesName) | **Post** /consul/roles/{name} | 
[**PostCubbyholePath**](Secrets.md#PostCubbyholePath) | **Post** /cubbyhole/{path} | Store a secret at the specified location.
[**PostGcpConfig**](Secrets.md#PostGcpConfig) | **Post** /gcp/config | 
[**PostGcpConfigRotateRoot**](Secrets.md#PostGcpConfigRotateRoot) | **Post** /gcp/config/rotate-root | 
[**PostGcpKeyRoleset**](Secrets.md#PostGcpKeyRoleset) | **Post** /gcp/key/{roleset} | 
[**PostGcpRolesetName**](Secrets.md#PostGcpRolesetName) | **Post** /gcp/roleset/{name} | 
[**PostGcpRolesetNameRotate**](Secrets.md#PostGcpRolesetNameRotate) | **Post** /gcp/roleset/{name}/rotate | 
[**PostGcpRolesetNameRotateKey**](Secrets.md#PostGcpRolesetNameRotateKey) | **Post** /gcp/roleset/{name}/rotate-key | 
[**PostGcpRolesetRolesetKey**](Secrets.md#PostGcpRolesetRolesetKey) | **Post** /gcp/roleset/{roleset}/key | 
[**PostGcpRolesetRolesetToken**](Secrets.md#PostGcpRolesetRolesetToken) | **Post** /gcp/roleset/{roleset}/token | 
[**PostGcpStaticAccountName**](Secrets.md#PostGcpStaticAccountName) | **Post** /gcp/static-account/{name} | 
[**PostGcpStaticAccountNameKey**](Secrets.md#PostGcpStaticAccountNameKey) | **Post** /gcp/static-account/{name}/key | 
[**PostGcpStaticAccountNameRotateKey**](Secrets.md#PostGcpStaticAccountNameRotateKey) | **Post** /gcp/static-account/{name}/rotate-key | 
[**PostGcpStaticAccountNameToken**](Secrets.md#PostGcpStaticAccountNameToken) | **Post** /gcp/static-account/{name}/token | 
[**PostGcpTokenRoleset**](Secrets.md#PostGcpTokenRoleset) | **Post** /gcp/token/{roleset} | 
[**PostGcpkmsConfig**](Secrets.md#PostGcpkmsConfig) | **Post** /gcpkms/config | Configure the GCP KMS secrets engine
[**PostGcpkmsDecryptKey**](Secrets.md#PostGcpkmsDecryptKey) | **Post** /gcpkms/decrypt/{key} | Decrypt a ciphertext value using a named key
[**PostGcpkmsEncryptKey**](Secrets.md#PostGcpkmsEncryptKey) | **Post** /gcpkms/encrypt/{key} | Encrypt a plaintext value using a named key
[**PostGcpkmsKeysConfigKey**](Secrets.md#PostGcpkmsKeysConfigKey) | **Post** /gcpkms/keys/config/{key} | Configure the key in Vault
[**PostGcpkmsKeysDeregisterKey**](Secrets.md#PostGcpkmsKeysDeregisterKey) | **Post** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
[**PostGcpkmsKeysKey**](Secrets.md#PostGcpkmsKeysKey) | **Post** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**PostGcpkmsKeysRegisterKey**](Secrets.md#PostGcpkmsKeysRegisterKey) | **Post** /gcpkms/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**PostGcpkmsKeysRotateKey**](Secrets.md#PostGcpkmsKeysRotateKey) | **Post** /gcpkms/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**PostGcpkmsKeysTrimKey**](Secrets.md#PostGcpkmsKeysTrimKey) | **Post** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**PostGcpkmsReencryptKey**](Secrets.md#PostGcpkmsReencryptKey) | **Post** /gcpkms/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**PostGcpkmsSignKey**](Secrets.md#PostGcpkmsSignKey) | **Post** /gcpkms/sign/{key} | Signs a message or digest using a named key
[**PostGcpkmsVerifyKey**](Secrets.md#PostGcpkmsVerifyKey) | **Post** /gcpkms/verify/{key} | Verify a signature using a named key
[**PostKubernetesConfig**](Secrets.md#PostKubernetesConfig) | **Post** /kubernetes/config | 
[**PostKubernetesCredsName**](Secrets.md#PostKubernetesCredsName) | **Post** /kubernetes/creds/{name} | 
[**PostKubernetesRolesName**](Secrets.md#PostKubernetesRolesName) | **Post** /kubernetes/roles/{name} | 
[**PostKvPath**](Secrets.md#PostKvPath) | **Post** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**PostMongodbatlasConfig**](Secrets.md#PostMongodbatlasConfig) | **Post** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
[**PostMongodbatlasCredsName**](Secrets.md#PostMongodbatlasCredsName) | **Post** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**PostMongodbatlasRolesName**](Secrets.md#PostMongodbatlasRolesName) | **Post** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**PostNomadConfigAccess**](Secrets.md#PostNomadConfigAccess) | **Post** /nomad/config/access | 
[**PostNomadConfigLease**](Secrets.md#PostNomadConfigLease) | **Post** /nomad/config/lease | Configure the lease parameters for generated tokens
[**PostNomadRoleName**](Secrets.md#PostNomadRoleName) | **Post** /nomad/role/{name} | 
[**PostOpenldapConfig**](Secrets.md#PostOpenldapConfig) | **Post** /openldap/config | 
[**PostOpenldapRoleName**](Secrets.md#PostOpenldapRoleName) | **Post** /openldap/role/{name} | 
[**PostOpenldapRotateRoleName**](Secrets.md#PostOpenldapRotateRoleName) | **Post** /openldap/rotate-role/{name} | 
[**PostOpenldapRotateRoot**](Secrets.md#PostOpenldapRotateRoot) | **Post** /openldap/rotate-root | 
[**PostOpenldapStaticRoleName**](Secrets.md#PostOpenldapStaticRoleName) | **Post** /openldap/static-role/{name} | 
[**PostPkiBundle**](Secrets.md#PostPkiBundle) | **Post** /pki/bundle | 
[**PostPkiCert**](Secrets.md#PostPkiCert) | **Post** /pki/cert | 
[**PostPkiConfigCa**](Secrets.md#PostPkiConfigCa) | **Post** /pki/config/ca | 
[**PostPkiConfigCrl**](Secrets.md#PostPkiConfigCrl) | **Post** /pki/config/crl | 
[**PostPkiConfigIssuers**](Secrets.md#PostPkiConfigIssuers) | **Post** /pki/config/issuers | 
[**PostPkiConfigKeys**](Secrets.md#PostPkiConfigKeys) | **Post** /pki/config/keys | 
[**PostPkiConfigUrls**](Secrets.md#PostPkiConfigUrls) | **Post** /pki/config/urls | 
[**PostPkiIntermediateCrossSign**](Secrets.md#PostPkiIntermediateCrossSign) | **Post** /pki/intermediate/cross-sign | 
[**PostPkiIntermediateGenerateExported**](Secrets.md#PostPkiIntermediateGenerateExported) | **Post** /pki/intermediate/generate/{exported} | 
[**PostPkiIntermediateSetSigned**](Secrets.md#PostPkiIntermediateSetSigned) | **Post** /pki/intermediate/set-signed | 
[**PostPkiInternalExported**](Secrets.md#PostPkiInternalExported) | **Post** /pki/internal|exported | 
[**PostPkiIssueRole**](Secrets.md#PostPkiIssueRole) | **Post** /pki/issue/{role} | 
[**PostPkiIssuerIssuerRefIssueRole**](Secrets.md#PostPkiIssuerIssuerRefIssueRole) | **Post** /pki/issuer/{issuer_ref}/issue/{role} | 
[**PostPkiIssuerIssuerRefSignIntermediate**](Secrets.md#PostPkiIssuerIssuerRefSignIntermediate) | **Post** /pki/issuer/{issuer_ref}/sign-intermediate | 
[**PostPkiIssuerIssuerRefSignRole**](Secrets.md#PostPkiIssuerIssuerRefSignRole) | **Post** /pki/issuer/{issuer_ref}/sign/{role} | 
[**PostPkiIssuerIssuerRefSignSelfIssued**](Secrets.md#PostPkiIssuerIssuerRefSignSelfIssued) | **Post** /pki/issuer/{issuer_ref}/sign-self-issued | 
[**PostPkiIssuerIssuerRefSignVerbatim**](Secrets.md#PostPkiIssuerIssuerRefSignVerbatim) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim | 
[**PostPkiIssuerIssuerRefSignVerbatimRole**](Secrets.md#PostPkiIssuerIssuerRefSignVerbatimRole) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**PostPkiIssuerRefDerPem**](Secrets.md#PostPkiIssuerRefDerPem) | **Post** /pki/{issuer_ref}/der|/pem | 
[**PostPkiIssuersGenerateIntermediateExported**](Secrets.md#PostPkiIssuersGenerateIntermediateExported) | **Post** /pki/issuers/generate/intermediate/{exported} | 
[**PostPkiIssuersGenerateRootExported**](Secrets.md#PostPkiIssuersGenerateRootExported) | **Post** /pki/issuers/generate/root/{exported} | 
[**PostPkiJson**](Secrets.md#PostPkiJson) | **Post** /pki//json | 
[**PostPkiKeyKeyRef**](Secrets.md#PostPkiKeyKeyRef) | **Post** /pki/key/{key_ref} | 
[**PostPkiKeysImport**](Secrets.md#PostPkiKeysImport) | **Post** /pki/keys/import | 
[**PostPkiKms**](Secrets.md#PostPkiKms) | **Post** /pki/kms | 
[**PostPkiRevoke**](Secrets.md#PostPkiRevoke) | **Post** /pki/revoke | 
[**PostPkiRolesName**](Secrets.md#PostPkiRolesName) | **Post** /pki/roles/{name} | 
[**PostPkiRootGenerateExported**](Secrets.md#PostPkiRootGenerateExported) | **Post** /pki/root/generate/{exported} | 
[**PostPkiRootReplace**](Secrets.md#PostPkiRootReplace) | **Post** /pki/root/replace | 
[**PostPkiRootRotateExported**](Secrets.md#PostPkiRootRotateExported) | **Post** /pki/root/rotate/{exported} | 
[**PostPkiRootSignIntermediate**](Secrets.md#PostPkiRootSignIntermediate) | **Post** /pki/root/sign-intermediate | 
[**PostPkiRootSignSelfIssued**](Secrets.md#PostPkiRootSignSelfIssued) | **Post** /pki/root/sign-self-issued | 
[**PostPkiSignRole**](Secrets.md#PostPkiSignRole) | **Post** /pki/sign/{role} | 
[**PostPkiSignVerbatim**](Secrets.md#PostPkiSignVerbatim) | **Post** /pki/sign-verbatim | 
[**PostPkiSignVerbatimRole**](Secrets.md#PostPkiSignVerbatimRole) | **Post** /pki/sign-verbatim/{role} | 
[**PostPkiTidy**](Secrets.md#PostPkiTidy) | **Post** /pki/tidy | 
[**PostRabbitmqConfigConnection**](Secrets.md#PostRabbitmqConfigConnection) | **Post** /rabbitmq/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**PostRabbitmqConfigLease**](Secrets.md#PostRabbitmqConfigLease) | **Post** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
[**PostRabbitmqRolesName**](Secrets.md#PostRabbitmqRolesName) | **Post** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
[**PostSecretConfig**](Secrets.md#PostSecretConfig) | **Post** /secret/config | Configure backend level settings that are applied to every key in the key-value store.
[**PostSecretDataPath**](Secrets.md#PostSecretDataPath) | **Post** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**PostSecretDeletePath**](Secrets.md#PostSecretDeletePath) | **Post** /secret/delete/{path} | Marks one or more versions as deleted in the KV store.
[**PostSecretDestroyPath**](Secrets.md#PostSecretDestroyPath) | **Post** /secret/destroy/{path} | Permanently removes one or more versions in the KV store
[**PostSecretMetadataPath**](Secrets.md#PostSecretMetadataPath) | **Post** /secret/metadata/{path} | Configures settings for the KV store
[**PostSecretUndeletePath**](Secrets.md#PostSecretUndeletePath) | **Post** /secret/undelete/{path} | Undeletes one or more versions from the KV store.
[**PostSshConfigCa**](Secrets.md#PostSshConfigCa) | **Post** /ssh/config/ca | Set the SSH private key used for signing certificates.
[**PostSshConfigZeroaddress**](Secrets.md#PostSshConfigZeroaddress) | **Post** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**PostSshCredsRole**](Secrets.md#PostSshCredsRole) | **Post** /ssh/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
[**PostSshKeysKeyName**](Secrets.md#PostSshKeysKeyName) | **Post** /ssh/keys/{key_name} | Register a shared private key with Vault.
[**PostSshLookup**](Secrets.md#PostSshLookup) | **Post** /ssh/lookup | List all the roles associated with the given IP address.
[**PostSshRolesRole**](Secrets.md#PostSshRolesRole) | **Post** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**PostSshSignRole**](Secrets.md#PostSshSignRole) | **Post** /ssh/sign/{role} | Request signing an SSH key using a certain role with the provided details.
[**PostSshVerify**](Secrets.md#PostSshVerify) | **Post** /ssh/verify | Validate the OTP provided by Vault SSH Agent.
[**PostTerraformConfig**](Secrets.md#PostTerraformConfig) | **Post** /terraform/config | 
[**PostTerraformCredsName**](Secrets.md#PostTerraformCredsName) | **Post** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**PostTerraformRoleName**](Secrets.md#PostTerraformRoleName) | **Post** /terraform/role/{name} | 
[**PostTerraformRotateRoleName**](Secrets.md#PostTerraformRotateRoleName) | **Post** /terraform/rotate-role/{name} | 
[**PostTotpCodeName**](Secrets.md#PostTotpCodeName) | **Post** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**PostTotpKeysName**](Secrets.md#PostTotpKeysName) | **Post** /totp/keys/{name} | Manage the keys that can be created with this backend.
[**PostTransitCacheConfig**](Secrets.md#PostTransitCacheConfig) | **Post** /transit/cache-config | Configures a new cache of the specified size
[**PostTransitDatakeyPlaintextName**](Secrets.md#PostTransitDatakeyPlaintextName) | **Post** /transit/datakey/{plaintext}/{name} | Generate a data key
[**PostTransitDecryptName**](Secrets.md#PostTransitDecryptName) | **Post** /transit/decrypt/{name} | Decrypt a ciphertext value using a named key
[**PostTransitEncryptName**](Secrets.md#PostTransitEncryptName) | **Post** /transit/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**PostTransitHash**](Secrets.md#PostTransitHash) | **Post** /transit/hash | Generate a hash sum for input data
[**PostTransitHashUrlalgorithm**](Secrets.md#PostTransitHashUrlalgorithm) | **Post** /transit/hash/{urlalgorithm} | Generate a hash sum for input data
[**PostTransitHmacName**](Secrets.md#PostTransitHmacName) | **Post** /transit/hmac/{name} | Generate an HMAC for input data using the named key
[**PostTransitHmacNameUrlalgorithm**](Secrets.md#PostTransitHmacNameUrlalgorithm) | **Post** /transit/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**PostTransitKeysName**](Secrets.md#PostTransitKeysName) | **Post** /transit/keys/{name} | Managed named encryption keys
[**PostTransitKeysNameConfig**](Secrets.md#PostTransitKeysNameConfig) | **Post** /transit/keys/{name}/config | Configure a named encryption key
[**PostTransitKeysNameImport**](Secrets.md#PostTransitKeysNameImport) | **Post** /transit/keys/{name}/import | Imports an externally-generated key into a new transit key
[**PostTransitKeysNameImportVersion**](Secrets.md#PostTransitKeysNameImportVersion) | **Post** /transit/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**PostTransitKeysNameRotate**](Secrets.md#PostTransitKeysNameRotate) | **Post** /transit/keys/{name}/rotate | Rotate named encryption key
[**PostTransitKeysNameTrim**](Secrets.md#PostTransitKeysNameTrim) | **Post** /transit/keys/{name}/trim | Trim key versions of a named key
[**PostTransitRandom**](Secrets.md#PostTransitRandom) | **Post** /transit/random | Generate random bytes
[**PostTransitRandomSource**](Secrets.md#PostTransitRandomSource) | **Post** /transit/random/{source} | Generate random bytes
[**PostTransitRandomSourceUrlbytes**](Secrets.md#PostTransitRandomSourceUrlbytes) | **Post** /transit/random/{source}/{urlbytes} | Generate random bytes
[**PostTransitRandomUrlbytes**](Secrets.md#PostTransitRandomUrlbytes) | **Post** /transit/random/{urlbytes} | Generate random bytes
[**PostTransitRestore**](Secrets.md#PostTransitRestore) | **Post** /transit/restore | Restore the named key
[**PostTransitRestoreName**](Secrets.md#PostTransitRestoreName) | **Post** /transit/restore/{name} | Restore the named key
[**PostTransitRewrapName**](Secrets.md#PostTransitRewrapName) | **Post** /transit/rewrap/{name} | Rewrap ciphertext
[**PostTransitSignName**](Secrets.md#PostTransitSignName) | **Post** /transit/sign/{name} | Generate a signature for input data using the named key
[**PostTransitSignNameUrlalgorithm**](Secrets.md#PostTransitSignNameUrlalgorithm) | **Post** /transit/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**PostTransitVerifyName**](Secrets.md#PostTransitVerifyName) | **Post** /transit/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**PostTransitVerifyNameUrlalgorithm**](Secrets.md#PostTransitVerifyNameUrlalgorithm) | **Post** /transit/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key



## DeleteAdConfig

> DeleteAdConfig(ctx).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteAdConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAdLibraryName

> DeleteAdLibraryName(ctx, name).Execute()

Delete a library set.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	resp, err := client.Secrets.DeleteAdLibraryName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteAdRolesName

> DeleteAdRolesName(ctx, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeleteAdRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteAlicloudConfig

> DeleteAlicloudConfig(ctx).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteAlicloudConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAlicloudRoleName

> DeleteAlicloudRoleName(ctx, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the role.
	
	resp, err := client.Secrets.DeleteAlicloudRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteAwsRolesName

> DeleteAwsRolesName(ctx, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the policy
	
	resp, err := client.Secrets.DeleteAwsRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteAzureConfig

> DeleteAzureConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteAzureConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAzureRolesName

> DeleteAzureRolesName(ctx, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.DeleteAzureRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteConsulRolesName

> DeleteConsulRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.DeleteConsulRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteCubbyholePath

> DeleteCubbyholePath(ctx, path).Execute()

Deletes the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Specifies the path of the secret.
	
	resp, err := client.Secrets.DeleteCubbyholePath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteGcpRolesetName

> DeleteGcpRolesetName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the role.
	
	resp, err := client.Secrets.DeleteGcpRolesetName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteGcpStaticAccountName

> DeleteGcpStaticAccountName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	
	resp, err := client.Secrets.DeleteGcpStaticAccountName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteGcpkmsConfig

> DeleteGcpkmsConfig(ctx).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteGcpkmsConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysDeregisterKey

> DeleteGcpkmsKeysDeregisterKey(ctx, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	
	resp, err := client.Secrets.DeleteGcpkmsKeysDeregisterKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteGcpkmsKeysKey

> DeleteGcpkmsKeysKey(ctx, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	resp, err := client.Secrets.DeleteGcpkmsKeysKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteGcpkmsKeysTrimKey

> DeleteGcpkmsKeysTrimKey(ctx, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	resp, err := client.Secrets.DeleteGcpkmsKeysTrimKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteKubernetesConfig

> DeleteKubernetesConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteKubernetesConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteKubernetesRolesName

> DeleteKubernetesRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeleteKubernetesRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteKvPath

> DeleteKvPath(ctx, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.DeleteKvPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteMongodbatlasRolesName

> DeleteMongodbatlasRolesName(ctx, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Roles
	
	resp, err := client.Secrets.DeleteMongodbatlasRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteNomadConfigAccess

> DeleteNomadConfigAccess(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteNomadConfigAccess(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadConfigLease

> DeleteNomadConfigLease(ctx).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteNomadConfigLease(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadRoleName

> DeleteNomadRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeleteNomadRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteOpenldapConfig

> DeleteOpenldapConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteOpenldapConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOpenldapRoleName

> DeleteOpenldapRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role (lowercase)
	
	resp, err := client.Secrets.DeleteOpenldapRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteOpenldapStaticRoleName

> DeleteOpenldapStaticRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeleteOpenldapStaticRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeletePkiIssuerRefDerPem

> DeletePkiIssuerRefDerPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	resp, err := client.Secrets.DeletePkiIssuerRefDerPem(context.Background(), issuerRef)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeletePkiJson

> DeletePkiJson(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeletePkiJson(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiKeyKeyRef

> DeletePkiKeyKeyRef(ctx, keyRef).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	keyRef :=  // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	
	resp, err := client.Secrets.DeletePkiKeyKeyRef(context.Background(), keyRef)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeletePkiRolesName

> DeletePkiRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeletePkiRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeletePkiRoot

> DeletePkiRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeletePkiRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteRabbitmqRolesName

> DeleteRabbitmqRolesName(ctx, name).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.DeleteRabbitmqRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteSecretDataPath

> DeleteSecretDataPath(ctx, path).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.DeleteSecretDataPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteSecretMetadataPath

> DeleteSecretMetadataPath(ctx, path).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.DeleteSecretMetadataPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteSshConfigCa

> DeleteSshConfigCa(ctx).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteSshConfigCa(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshConfigZeroaddress

> DeleteSshConfigZeroaddress(ctx).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteSshConfigZeroaddress(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshKeysKeyName

> DeleteSshKeysKeyName(ctx, keyName).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	keyName :=  // string | [Required] Name of the key
	
	resp, err := client.Secrets.DeleteSshKeysKeyName(context.Background(), keyName)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteSshRolesRole

> DeleteSshRolesRole(ctx, role).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | [Required for all types] Name of the role being created.
	
	resp, err := client.Secrets.DeleteSshRolesRole(context.Background(), role)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteTerraformConfig

> DeleteTerraformConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.DeleteTerraformConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTerraformRoleName

> DeleteTerraformRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.DeleteTerraformRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteTotpKeysName

> DeleteTotpKeysName(ctx, name).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key.
	
	resp, err := client.Secrets.DeleteTotpKeysName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## DeleteTransitKeysName

> DeleteTransitKeysName(ctx, name).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Secrets.DeleteTransitKeysName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAdConfig

> GetAdConfig(ctx).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAdConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdCredsName

> GetAdCredsName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetAdCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAdLibrary

> GetAdLibrary(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetAdLibrary(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdLibraryName

> GetAdLibraryName(ctx, name).Execute()

Read a library set.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	resp, err := client.Secrets.GetAdLibraryName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAdLibraryNameStatus

> GetAdLibraryNameStatus(ctx, name).Execute()

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	resp, err := client.Secrets.GetAdLibraryNameStatus(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAdRoles

> GetAdRoles(ctx).List(list).Execute()

List the name of each role currently stored.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetAdRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdRolesName

> GetAdRolesName(ctx, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetAdRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAdRotateRoot

> GetAdRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAdRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudConfig

> GetAlicloudConfig(ctx).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAlicloudConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudCredsName

> GetAlicloudCredsName(ctx, name).Execute()

Generate an API key or STS credential using the given role's configuration.'

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the role.
	
	resp, err := client.Secrets.GetAlicloudCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAlicloudRole

> GetAlicloudRole(ctx).List(list).Execute()

List the existing roles in this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetAlicloudRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudRoleName

> GetAlicloudRoleName(ctx, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the role.
	
	resp, err := client.Secrets.GetAlicloudRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAwsConfigLease

> GetAwsConfigLease(ctx).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAwsConfigLease(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsConfigRoot

> GetAwsConfigRoot(ctx).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAwsConfigRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsCreds

> GetAwsCreds(ctx).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAwsCreds(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsRoles

> GetAwsRoles(ctx).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetAwsRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsRolesName

> GetAwsRolesName(ctx, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the policy
	
	resp, err := client.Secrets.GetAwsRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAwsStsName

> GetAwsStsName(ctx, name).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetAwsStsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAzureConfig

> GetAzureConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetAzureConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureCredsRole

> GetAzureCredsRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | Name of the Vault role
	
	resp, err := client.Secrets.GetAzureCredsRole(context.Background(), role)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetAzureRoles

> GetAzureRoles(ctx).List(list).Execute()

List existing roles.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetAzureRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureRolesName

> GetAzureRolesName(ctx, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.GetAzureRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetConsulConfigAccess

> GetConsulConfigAccess(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetConsulConfigAccess(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulCredsRole

> GetConsulCredsRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | Name of the role.
	
	resp, err := client.Secrets.GetConsulCredsRole(context.Background(), role)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetConsulRoles

> GetConsulRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetConsulRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulRolesName

> GetConsulRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.GetConsulRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetCubbyholePath

> GetCubbyholePath(ctx, path).List(list).Execute()

Retrieve the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Specifies the path of the secret.
	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetCubbyholePath(context.Background(), path, list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpConfig

> GetGcpConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetGcpConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpKeyRoleset

> GetGcpKeyRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.GetGcpKeyRoleset(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpRolesetName

> GetGcpRolesetName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the role.
	
	resp, err := client.Secrets.GetGcpRolesetName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpRolesetRolesetKey

> GetGcpRolesetRolesetKey(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.GetGcpRolesetRolesetKey(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpRolesetRolesetToken

> GetGcpRolesetRolesetToken(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.GetGcpRolesetRolesetToken(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpRolesets

> GetGcpRolesets(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetGcpRolesets(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpStaticAccountName

> GetGcpStaticAccountName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	
	resp, err := client.Secrets.GetGcpStaticAccountName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpStaticAccountNameKey

> GetGcpStaticAccountNameKey(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the static account.
	
	resp, err := client.Secrets.GetGcpStaticAccountNameKey(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpStaticAccountNameToken

> GetGcpStaticAccountNameToken(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the static account.
	
	resp, err := client.Secrets.GetGcpStaticAccountNameToken(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpStaticAccounts

> GetGcpStaticAccounts(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetGcpStaticAccounts(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpTokenRoleset

> GetGcpTokenRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.GetGcpTokenRoleset(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpkmsConfig

> GetGcpkmsConfig(ctx).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetGcpkmsConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsKeys

> GetGcpkmsKeys(ctx).List(list).Execute()

List named keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetGcpkmsKeys(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsKeysConfigKey

> GetGcpkmsKeysConfigKey(ctx, key).Execute()

Configure the key in Vault

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	resp, err := client.Secrets.GetGcpkmsKeysConfigKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpkmsKeysKey

> GetGcpkmsKeysKey(ctx, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	resp, err := client.Secrets.GetGcpkmsKeysKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetGcpkmsPubkeyKey

> GetGcpkmsPubkeyKey(ctx, key).Execute()

Retrieve the public key associated with the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
	
	resp, err := client.Secrets.GetGcpkmsPubkeyKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetKubernetesConfig

> GetKubernetesConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetKubernetesConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetKubernetesRoles

> GetKubernetesRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetKubernetesRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetKubernetesRolesName

> GetKubernetesRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetKubernetesRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetKvPath

> GetKvPath(ctx, path).List(list).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetKvPath(context.Background(), path, list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetMongodbatlasConfig

> GetMongodbatlasConfig(ctx).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetMongodbatlasConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasCredsName

> GetMongodbatlasCredsName(ctx, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetMongodbatlasCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetMongodbatlasRoles

> GetMongodbatlasRoles(ctx).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetMongodbatlasRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasRolesName

> GetMongodbatlasRolesName(ctx, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Roles
	
	resp, err := client.Secrets.GetMongodbatlasRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetNomadConfigAccess

> GetNomadConfigAccess(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetNomadConfigAccess(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadConfigLease

> GetNomadConfigLease(ctx).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetNomadConfigLease(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadCredsName

> GetNomadCredsName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetNomadCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetNomadRole

> GetNomadRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetNomadRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadRoleName

> GetNomadRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetNomadRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetOpenldapConfig

> GetOpenldapConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetOpenldapConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapCredsName

> GetOpenldapCredsName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the dynamic role.
	
	resp, err := client.Secrets.GetOpenldapCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetOpenldapRole

> GetOpenldapRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetOpenldapRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapRoleName

> GetOpenldapRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role (lowercase)
	
	resp, err := client.Secrets.GetOpenldapRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetOpenldapStaticCredName

> GetOpenldapStaticCredName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the static role.
	
	resp, err := client.Secrets.GetOpenldapStaticCredName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetOpenldapStaticRole

> GetOpenldapStaticRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetOpenldapStaticRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapStaticRoleName

> GetOpenldapStaticRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetOpenldapStaticRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiCa

> GetPkiCa(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCa(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCaChain

> GetPkiCaChain(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCaChain(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCaPem

> GetPkiCaPem(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCaPem(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertCaChain

> GetPkiCertCaChain(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCertCaChain(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertCrl

> GetPkiCertCrl(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCertCrl(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertSerial

> GetPkiCertSerial(ctx, serial).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	serial :=  // string | Certificate serial number, in colon- or hyphen-separated octal
	
	resp, err := client.Secrets.GetPkiCertSerial(context.Background(), serial)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiCertSerialRaw

> GetPkiCertSerialRaw(ctx, serial).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	serial :=  // string | Certificate serial number, in colon- or hyphen-separated octal
	
	resp, err := client.Secrets.GetPkiCertSerialRaw(context.Background(), serial)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiCertSerialRawPem

> GetPkiCertSerialRawPem(ctx, serial).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	serial :=  // string | Certificate serial number, in colon- or hyphen-separated octal
	
	resp, err := client.Secrets.GetPkiCertSerialRawPem(context.Background(), serial)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiCerts

> GetPkiCerts(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetPkiCerts(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigCrl

> GetPkiConfigCrl(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiConfigCrl(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigIssuers

> GetPkiConfigIssuers(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiConfigIssuers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigKeys

> GetPkiConfigKeys(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiConfigKeys(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigUrls

> GetPkiConfigUrls(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiConfigUrls(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrl

> GetPkiCrl(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCrl(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrlPem

> GetPkiCrlPem(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCrlPem(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrlRotate

> GetPkiCrlRotate(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiCrlRotate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiDer

> GetPkiDer(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiDer(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiIssuerRefCrlPem

> GetPkiIssuerRefCrlPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	resp, err := client.Secrets.GetPkiIssuerRefCrlPem(context.Background(), issuerRef)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiIssuerRefDerPem

> GetPkiIssuerRefDerPem(ctx, issuerRef).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	resp, err := client.Secrets.GetPkiIssuerRefDerPem(context.Background(), issuerRef)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiIssuers

> GetPkiIssuers(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetPkiIssuers(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiJson

> GetPkiJson(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiJson(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiKeyKeyRef

> GetPkiKeyKeyRef(ctx, keyRef).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	keyRef :=  // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	
	resp, err := client.Secrets.GetPkiKeyKeyRef(context.Background(), keyRef)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiKeys

> GetPkiKeys(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetPkiKeys(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiRoles

> GetPkiRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetPkiRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiRolesName

> GetPkiRolesName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetPkiRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetPkiTidyStatus

> GetPkiTidyStatus(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetPkiTidyStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqConfigLease

> GetRabbitmqConfigLease(ctx).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetRabbitmqConfigLease(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqCredsName

> GetRabbitmqCredsName(ctx, name).Execute()

Request RabbitMQ credentials for a certain role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.GetRabbitmqCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetRabbitmqRoles

> GetRabbitmqRoles(ctx).List(list).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetRabbitmqRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqRolesName

> GetRabbitmqRolesName(ctx, name).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.GetRabbitmqRolesName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetSecretConfig

> GetSecretConfig(ctx).Execute()

Read the backend level settings.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetSecretConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretDataPath

> GetSecretDataPath(ctx, path).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.GetSecretDataPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetSecretMetadataPath

> GetSecretMetadataPath(ctx, path).List(list).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetSecretMetadataPath(context.Background(), path, list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetSecretSubkeysPath

> GetSecretSubkeysPath(ctx, path).Execute()

Read the structure of a secret entry from the Key-Value store with the values removed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.GetSecretSubkeysPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetSshConfigCa

> GetSshConfigCa(ctx).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetSshConfigCa(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshConfigZeroaddress

> GetSshConfigZeroaddress(ctx).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetSshConfigZeroaddress(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshPublicKey

> GetSshPublicKey(ctx).Execute()

Retrieve the public key.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetSshPublicKey(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshRoles

> GetSshRoles(ctx).List(list).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetSshRoles(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshRolesRole

> GetSshRolesRole(ctx, role).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | [Required for all types] Name of the role being created.
	
	resp, err := client.Secrets.GetSshRolesRole(context.Background(), role)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTerraformConfig

> GetTerraformConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetTerraformConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformCredsName

> GetTerraformCredsName(ctx, name).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetTerraformCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTerraformRole

> GetTerraformRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetTerraformRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformRoleName

> GetTerraformRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.GetTerraformRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTotpCodeName

> GetTotpCodeName(ctx, name).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key.
	
	resp, err := client.Secrets.GetTotpCodeName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTotpKeys

> GetTotpKeys(ctx).List(list).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetTotpKeys(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTotpKeysName

> GetTotpKeysName(ctx, name).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key.
	
	resp, err := client.Secrets.GetTotpKeysName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTransitBackupName

> GetTransitBackupName(ctx, name).Execute()

Backup the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Secrets.GetTransitBackupName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTransitCacheConfig

> GetTransitCacheConfig(ctx).Execute()

Returns the size of the active cache

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetTransitCacheConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitExportTypeName

> GetTransitExportTypeName(ctx, name, type_).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	type_ :=  // string | Type of key to export (encryption-key, signing-key, hmac-key)
	
	resp, err := client.Secrets.GetTransitExportTypeName(context.Background(), name, type_)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTransitExportTypeNameVersion

> GetTransitExportTypeNameVersion(ctx, name, type_, version).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	type_ :=  // string | Type of key to export (encryption-key, signing-key, hmac-key)
	version :=  // string | Version of the key
	
	resp, err := client.Secrets.GetTransitExportTypeNameVersion(context.Background(), name, type_, version)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTransitKeys

> GetTransitKeys(ctx).List(list).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Secrets.GetTransitKeys(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitKeysName

> GetTransitKeysName(ctx, name).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Secrets.GetTransitKeysName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetTransitWrappingKey

> GetTransitWrappingKey(ctx).Execute()

Returns the public key to use for wrapping imported keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.GetTransitWrappingKey(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdConfig

> PostAdConfig(ctx).AdConfigRequest(adConfigRequest).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	adConfigRequest := NewAdConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdConfig(context.Background(), adConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adConfigRequest** | [**AdConfigRequest**](AdConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryManageNameCheckIn

> PostAdLibraryManageNameCheckIn(ctx, name).AdLibraryManageCheckInRequest(adLibraryManageCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	adLibraryManageCheckInRequest := NewAdLibraryManageCheckInRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdLibraryManageNameCheckIn(context.Background(), name, adLibraryManageCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **adLibraryManageCheckInRequest** | [**AdLibraryManageCheckInRequest**](AdLibraryManageCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryName

> PostAdLibraryName(ctx, name).AdLibraryRequest(adLibraryRequest).Execute()

Update a library set.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	adLibraryRequest := NewAdLibraryRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdLibraryName(context.Background(), name, adLibraryRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **adLibraryRequest** | [**AdLibraryRequest**](AdLibraryRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckIn

> PostAdLibraryNameCheckIn(ctx, name).AdLibraryCheckInRequest(adLibraryCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set.
	
	adLibraryCheckInRequest := NewAdLibraryCheckInRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdLibraryNameCheckIn(context.Background(), name, adLibraryCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **adLibraryCheckInRequest** | [**AdLibraryCheckInRequest**](AdLibraryCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckOut

> PostAdLibraryNameCheckOut(ctx, name).AdLibraryCheckOutRequest(adLibraryCheckOutRequest).Execute()

Check a service account out from the library.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the set
	
	adLibraryCheckOutRequest := NewAdLibraryCheckOutRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdLibraryNameCheckOut(context.Background(), name, adLibraryCheckOutRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **adLibraryCheckOutRequest** | [**AdLibraryCheckOutRequest**](AdLibraryCheckOutRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdRolesName

> PostAdRolesName(ctx, name).AdRolesRequest(adRolesRequest).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	adRolesRequest := NewAdRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostAdRolesName(context.Background(), name, adRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **adRolesRequest** | [**AdRolesRequest**](AdRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdRotateRoleName

> PostAdRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the static role
	
	resp, err := client.Secrets.PostAdRotateRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostAdRotateRoot

> PostAdRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.PostAdRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAlicloudConfig

> PostAlicloudConfig(ctx).AlicloudConfigRequest(alicloudConfigRequest).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	alicloudConfigRequest := NewAlicloudConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostAlicloudConfig(context.Background(), alicloudConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alicloudConfigRequest** | [**AlicloudConfigRequest**](AlicloudConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAlicloudRoleName

> PostAlicloudRoleName(ctx, name).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the role.
	
	alicloudRoleRequest := NewAlicloudRoleRequestWithDefaults()
	
	resp, err := client.Secrets.PostAlicloudRoleName(context.Background(), name, alicloudRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **alicloudRoleRequest** | [**AlicloudRoleRequest**](AlicloudRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigLease

> PostAwsConfigLease(ctx).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigLeaseRequest := NewAwsConfigLeaseRequestWithDefaults()
	
	resp, err := client.Secrets.PostAwsConfigLease(context.Background(), awsConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigLeaseRequest** | [**AwsConfigLeaseRequest**](AwsConfigLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigRoot

> PostAwsConfigRoot(ctx).AwsConfigRootRequest(awsConfigRootRequest).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigRootRequest := NewAwsConfigRootRequestWithDefaults()
	
	resp, err := client.Secrets.PostAwsConfigRoot(context.Background(), awsConfigRootRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigRootRequest** | [**AwsConfigRootRequest**](AwsConfigRootRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigRotateRoot

> PostAwsConfigRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.PostAwsConfigRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsCreds

> PostAwsCreds(ctx).AwsCredsRequest(awsCredsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsCredsRequest := NewAwsCredsRequestWithDefaults()
	
	resp, err := client.Secrets.PostAwsCreds(context.Background(), awsCredsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredsRequest** | [**AwsCredsRequest**](AwsCredsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsRolesName

> PostAwsRolesName(ctx, name).AwsRolesRequest(awsRolesRequest).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the policy
	
	awsRolesRequest := NewAwsRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostAwsRolesName(context.Background(), name, awsRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **awsRolesRequest** | [**AwsRolesRequest**](AwsRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsStsName

> PostAwsStsName(ctx, name).AwsStsRequest(awsStsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	awsStsRequest := NewAwsStsRequestWithDefaults()
	
	resp, err := client.Secrets.PostAwsStsName(context.Background(), name, awsStsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **awsStsRequest** | [**AwsStsRequest**](AwsStsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureConfig

> PostAzureConfig(ctx).AzureConfigRequest(azureConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	azureConfigRequest := NewAzureConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostAzureConfig(context.Background(), azureConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureConfigRequest** | [**AzureConfigRequest**](AzureConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureRolesName

> PostAzureRolesName(ctx, name).AzureRolesRequest(azureRolesRequest).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	azureRolesRequest := NewAzureRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostAzureRolesName(context.Background(), name, azureRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **azureRolesRequest** | [**AzureRolesRequest**](AzureRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureRotateRoot

> PostAzureRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.PostAzureRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostConsulConfigAccess

> PostConsulConfigAccess(ctx).ConsulConfigAccessRequest(consulConfigAccessRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	consulConfigAccessRequest := NewConsulConfigAccessRequestWithDefaults()
	
	resp, err := client.Secrets.PostConsulConfigAccess(context.Background(), consulConfigAccessRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulConfigAccessRequest** | [**ConsulConfigAccessRequest**](ConsulConfigAccessRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostConsulRolesName

> PostConsulRolesName(ctx, name).ConsulRolesRequest(consulRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	consulRolesRequest := NewConsulRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostConsulRolesName(context.Background(), name, consulRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **consulRolesRequest** | [**ConsulRolesRequest**](ConsulRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostCubbyholePath

> PostCubbyholePath(ctx, path).Execute()

Store a secret at the specified location.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Specifies the path of the secret.
	
	resp, err := client.Secrets.PostCubbyholePath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpConfig

> PostGcpConfig(ctx).GcpConfigRequest(gcpConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	gcpConfigRequest := NewGcpConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpConfig(context.Background(), gcpConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpConfigRequest** | [**GcpConfigRequest**](GcpConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpConfigRotateRoot

> PostGcpConfigRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.PostGcpConfigRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpKeyRoleset

> PostGcpKeyRoleset(ctx, roleset).GcpKeyRequest(gcpKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	gcpKeyRequest := NewGcpKeyRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpKeyRoleset(context.Background(), roleset, gcpKeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpKeyRequest** | [**GcpKeyRequest**](GcpKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetName

> PostGcpRolesetName(ctx, name).GcpRolesetRequest(gcpRolesetRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the role.
	
	gcpRolesetRequest := NewGcpRolesetRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpRolesetName(context.Background(), name, gcpRolesetRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpRolesetRequest** | [**GcpRolesetRequest**](GcpRolesetRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetNameRotate

> PostGcpRolesetNameRotate(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.PostGcpRolesetNameRotate(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpRolesetNameRotateKey

> PostGcpRolesetNameRotateKey(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	resp, err := client.Secrets.PostGcpRolesetNameRotateKey(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpRolesetRolesetKey

> PostGcpRolesetRolesetKey(ctx, roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	gcpRolesetKeyRequest := NewGcpRolesetKeyRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpRolesetRolesetKey(context.Background(), roleset, gcpRolesetKeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpRolesetKeyRequest** | [**GcpRolesetKeyRequest**](GcpRolesetKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetRolesetToken

> PostGcpRolesetRolesetToken(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.PostGcpRolesetRolesetToken(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpStaticAccountName

> PostGcpStaticAccountName(ctx, name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	
	gcpStaticAccountRequest := NewGcpStaticAccountRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpStaticAccountName(context.Background(), name, gcpStaticAccountRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpStaticAccountRequest** | [**GcpStaticAccountRequest**](GcpStaticAccountRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameKey

> PostGcpStaticAccountNameKey(ctx, name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the static account.
	
	gcpStaticAccountKeyRequest := NewGcpStaticAccountKeyRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpStaticAccountNameKey(context.Background(), name, gcpStaticAccountKeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpStaticAccountKeyRequest** | [**GcpStaticAccountKeyRequest**](GcpStaticAccountKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameRotateKey

> PostGcpStaticAccountNameRotateKey(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the account.
	
	resp, err := client.Secrets.PostGcpStaticAccountNameRotateKey(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpStaticAccountNameToken

> PostGcpStaticAccountNameToken(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Required. Name of the static account.
	
	resp, err := client.Secrets.PostGcpStaticAccountNameToken(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpTokenRoleset

> PostGcpTokenRoleset(ctx, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleset :=  // string | Required. Name of the role set.
	
	resp, err := client.Secrets.PostGcpTokenRoleset(context.Background(), roleset)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpkmsConfig

> PostGcpkmsConfig(ctx).GcpkmsConfigRequest(gcpkmsConfigRequest).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	gcpkmsConfigRequest := NewGcpkmsConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsConfig(context.Background(), gcpkmsConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpkmsConfigRequest** | [**GcpkmsConfigRequest**](GcpkmsConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsDecryptKey

> PostGcpkmsDecryptKey(ctx, key).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	
	gcpkmsDecryptRequest := NewGcpkmsDecryptRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsDecryptKey(context.Background(), key, gcpkmsDecryptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsDecryptRequest** | [**GcpkmsDecryptRequest**](GcpkmsDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsEncryptKey

> PostGcpkmsEncryptKey(ctx, key).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()

Encrypt a plaintext value using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	
	gcpkmsEncryptRequest := NewGcpkmsEncryptRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsEncryptKey(context.Background(), key, gcpkmsEncryptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsEncryptRequest** | [**GcpkmsEncryptRequest**](GcpkmsEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysConfigKey

> PostGcpkmsKeysConfigKey(ctx, key).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()

Configure the key in Vault

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	gcpkmsKeysConfigRequest := NewGcpkmsKeysConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsKeysConfigKey(context.Background(), key, gcpkmsKeysConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsKeysConfigRequest** | [**GcpkmsKeysConfigRequest**](GcpkmsKeysConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysDeregisterKey

> PostGcpkmsKeysDeregisterKey(ctx, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	
	resp, err := client.Secrets.PostGcpkmsKeysDeregisterKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpkmsKeysKey

> PostGcpkmsKeysKey(ctx, key).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	gcpkmsKeysRequest := NewGcpkmsKeysRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsKeysKey(context.Background(), key, gcpkmsKeysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsKeysRequest** | [**GcpkmsKeysRequest**](GcpkmsKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysRegisterKey

> PostGcpkmsKeysRegisterKey(ctx, key).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()

Register an existing crypto key in Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
	
	gcpkmsKeysRegisterRequest := NewGcpkmsKeysRegisterRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsKeysRegisterKey(context.Background(), key, gcpkmsKeysRegisterRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsKeysRegisterRequest** | [**GcpkmsKeysRegisterRequest**](GcpkmsKeysRegisterRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysRotateKey

> PostGcpkmsKeysRotateKey(ctx, key).Execute()

Rotate a crypto key to a new primary version

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
	
	resp, err := client.Secrets.PostGcpkmsKeysRotateKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpkmsKeysTrimKey

> PostGcpkmsKeysTrimKey(ctx, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault.
	
	resp, err := client.Secrets.PostGcpkmsKeysTrimKey(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostGcpkmsReencryptKey

> PostGcpkmsReencryptKey(ctx, key).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()

Re-encrypt existing ciphertext data to a new version

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
	
	gcpkmsReencryptRequest := NewGcpkmsReencryptRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsReencryptKey(context.Background(), key, gcpkmsReencryptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsReencryptRequest** | [**GcpkmsReencryptRequest**](GcpkmsReencryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsSignKey

> PostGcpkmsSignKey(ctx, key).GcpkmsSignRequest(gcpkmsSignRequest).Execute()

Signs a message or digest using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	
	gcpkmsSignRequest := NewGcpkmsSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsSignKey(context.Background(), key, gcpkmsSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsSignRequest** | [**GcpkmsSignRequest**](GcpkmsSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsVerifyKey

> PostGcpkmsVerifyKey(ctx, key).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()

Verify a signature using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	
	gcpkmsVerifyRequest := NewGcpkmsVerifyRequestWithDefaults()
	
	resp, err := client.Secrets.PostGcpkmsVerifyKey(context.Background(), key, gcpkmsVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **gcpkmsVerifyRequest** | [**GcpkmsVerifyRequest**](GcpkmsVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesConfig

> PostKubernetesConfig(ctx).KubernetesConfigRequest(kubernetesConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kubernetesConfigRequest := NewKubernetesConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostKubernetesConfig(context.Background(), kubernetesConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesConfigRequest** | [**KubernetesConfigRequest**](KubernetesConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesCredsName

> PostKubernetesCredsName(ctx, name).KubernetesCredsRequest(kubernetesCredsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Vault role
	
	kubernetesCredsRequest := NewKubernetesCredsRequestWithDefaults()
	
	resp, err := client.Secrets.PostKubernetesCredsName(context.Background(), name, kubernetesCredsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kubernetesCredsRequest** | [**KubernetesCredsRequest**](KubernetesCredsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesRolesName

> PostKubernetesRolesName(ctx, name).KubernetesRolesRequest(kubernetesRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	kubernetesRolesRequest := NewKubernetesRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostKubernetesRolesName(context.Background(), name, kubernetesRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kubernetesRolesRequest** | [**KubernetesRolesRequest**](KubernetesRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKvPath

> PostKvPath(ctx, path).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	resp, err := client.Secrets.PostKvPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostMongodbatlasConfig

> PostMongodbatlasConfig(ctx).MongodbatlasConfigRequest(mongodbatlasConfigRequest).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	mongodbatlasConfigRequest := NewMongodbatlasConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostMongodbatlasConfig(context.Background(), mongodbatlasConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mongodbatlasConfigRequest** | [**MongodbatlasConfigRequest**](MongodbatlasConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostMongodbatlasCredsName

> PostMongodbatlasCredsName(ctx, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.PostMongodbatlasCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostMongodbatlasRolesName

> PostMongodbatlasRolesName(ctx, name).MongodbatlasRolesRequest(mongodbatlasRolesRequest).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Roles
	
	mongodbatlasRolesRequest := NewMongodbatlasRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostMongodbatlasRolesName(context.Background(), name, mongodbatlasRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **mongodbatlasRolesRequest** | [**MongodbatlasRolesRequest**](MongodbatlasRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadConfigAccess

> PostNomadConfigAccess(ctx).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	nomadConfigAccessRequest := NewNomadConfigAccessRequestWithDefaults()
	
	resp, err := client.Secrets.PostNomadConfigAccess(context.Background(), nomadConfigAccessRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigAccessRequest** | [**NomadConfigAccessRequest**](NomadConfigAccessRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadConfigLease

> PostNomadConfigLease(ctx).NomadConfigLeaseRequest(nomadConfigLeaseRequest).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	nomadConfigLeaseRequest := NewNomadConfigLeaseRequestWithDefaults()
	
	resp, err := client.Secrets.PostNomadConfigLease(context.Background(), nomadConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigLeaseRequest** | [**NomadConfigLeaseRequest**](NomadConfigLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadRoleName

> PostNomadRoleName(ctx, name).NomadRoleRequest(nomadRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	nomadRoleRequest := NewNomadRoleRequestWithDefaults()
	
	resp, err := client.Secrets.PostNomadRoleName(context.Background(), name, nomadRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **nomadRoleRequest** | [**NomadRoleRequest**](NomadRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapConfig

> PostOpenldapConfig(ctx).OpenldapConfigRequest(openldapConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	openldapConfigRequest := NewOpenldapConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostOpenldapConfig(context.Background(), openldapConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openldapConfigRequest** | [**OpenldapConfigRequest**](OpenldapConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapRoleName

> PostOpenldapRoleName(ctx, name).OpenldapRoleRequest(openldapRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role (lowercase)
	
	openldapRoleRequest := NewOpenldapRoleRequestWithDefaults()
	
	resp, err := client.Secrets.PostOpenldapRoleName(context.Background(), name, openldapRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **openldapRoleRequest** | [**OpenldapRoleRequest**](OpenldapRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapRotateRoleName

> PostOpenldapRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the static role
	
	resp, err := client.Secrets.PostOpenldapRotateRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostOpenldapRotateRoot

> PostOpenldapRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Secrets.PostOpenldapRotateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapStaticRoleName

> PostOpenldapStaticRoleName(ctx, name).OpenldapStaticRoleRequest(openldapStaticRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	openldapStaticRoleRequest := NewOpenldapStaticRoleRequestWithDefaults()
	
	resp, err := client.Secrets.PostOpenldapStaticRoleName(context.Background(), name, openldapStaticRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **openldapStaticRoleRequest** | [**OpenldapStaticRoleRequest**](OpenldapStaticRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiBundle

> PostPkiBundle(ctx).PkiBundleRequest(pkiBundleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiBundleRequest := NewPkiBundleRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiBundle(context.Background(), pkiBundleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiBundleRequest** | [**PkiBundleRequest**](PkiBundleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiCert

> PostPkiCert(ctx).PkiCertRequest(pkiCertRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiCertRequest := NewPkiCertRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiCert(context.Background(), pkiCertRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiCertRequest** | [**PkiCertRequest**](PkiCertRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigCa

> PostPkiConfigCa(ctx).PkiConfigCaRequest(pkiConfigCaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiConfigCaRequest := NewPkiConfigCaRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiConfigCa(context.Background(), pkiConfigCaRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCaRequest** | [**PkiConfigCaRequest**](PkiConfigCaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigCrl

> PostPkiConfigCrl(ctx).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiConfigCrlRequest := NewPkiConfigCrlRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiConfigCrl(context.Background(), pkiConfigCrlRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCrlRequest** | [**PkiConfigCrlRequest**](PkiConfigCrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigIssuers

> PostPkiConfigIssuers(ctx).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiConfigIssuersRequest := NewPkiConfigIssuersRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiConfigIssuers(context.Background(), pkiConfigIssuersRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigIssuersRequest** | [**PkiConfigIssuersRequest**](PkiConfigIssuersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigKeys

> PostPkiConfigKeys(ctx).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiConfigKeysRequest := NewPkiConfigKeysRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiConfigKeys(context.Background(), pkiConfigKeysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigKeysRequest** | [**PkiConfigKeysRequest**](PkiConfigKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigUrls

> PostPkiConfigUrls(ctx).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiConfigUrlsRequest := NewPkiConfigUrlsRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiConfigUrls(context.Background(), pkiConfigUrlsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigUrlsRequest** | [**PkiConfigUrlsRequest**](PkiConfigUrlsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateCrossSign

> PostPkiIntermediateCrossSign(ctx).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiIntermediateCrossSignRequest := NewPkiIntermediateCrossSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIntermediateCrossSign(context.Background(), pkiIntermediateCrossSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateCrossSignRequest** | [**PkiIntermediateCrossSignRequest**](PkiIntermediateCrossSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateGenerateExported

> PostPkiIntermediateGenerateExported(ctx, exported).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	exported :=  // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	
	pkiIntermediateGenerateRequest := NewPkiIntermediateGenerateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIntermediateGenerateExported(context.Background(), exported, pkiIntermediateGenerateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIntermediateGenerateRequest** | [**PkiIntermediateGenerateRequest**](PkiIntermediateGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateSetSigned

> PostPkiIntermediateSetSigned(ctx).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiIntermediateSetSignedRequest := NewPkiIntermediateSetSignedRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIntermediateSetSigned(context.Background(), pkiIntermediateSetSignedRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateSetSignedRequest** | [**PkiIntermediateSetSignedRequest**](PkiIntermediateSetSignedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiInternalExported

> PostPkiInternalExported(ctx).PkiInternalExportedRequest(pkiInternalExportedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiInternalExportedRequest := NewPkiInternalExportedRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiInternalExported(context.Background(), pkiInternalExportedRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiInternalExportedRequest** | [**PkiInternalExportedRequest**](PkiInternalExportedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssueRole

> PostPkiIssueRole(ctx, role).PkiIssueRequest(pkiIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The desired role with configuration for this request
	
	pkiIssueRequest := NewPkiIssueRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssueRole(context.Background(), role, pkiIssueRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssueRequest** | [**PkiIssueRequest**](PkiIssueRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefIssueRole

> PostPkiIssuerIssuerRefIssueRole(ctx, issuerRef, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role :=  // string | The desired role with configuration for this request
	
	pkiIssuerIssueRequest := NewPkiIssuerIssueRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefIssueRole(context.Background(), issuerRef, role, pkiIssuerIssueRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **pkiIssuerIssueRequest** | [**PkiIssuerIssueRequest**](PkiIssuerIssueRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignIntermediate

> PostPkiIssuerIssuerRefSignIntermediate(ctx, issuerRef).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	pkiIssuerSignIntermediateRequest := NewPkiIssuerSignIntermediateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefSignIntermediate(context.Background(), issuerRef, pkiIssuerSignIntermediateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssuerSignIntermediateRequest** | [**PkiIssuerSignIntermediateRequest**](PkiIssuerSignIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignRole

> PostPkiIssuerIssuerRefSignRole(ctx, issuerRef, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role :=  // string | The desired role with configuration for this request
	
	pkiIssuerSignRequest := NewPkiIssuerSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefSignRole(context.Background(), issuerRef, role, pkiIssuerSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **pkiIssuerSignRequest** | [**PkiIssuerSignRequest**](PkiIssuerSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignSelfIssued

> PostPkiIssuerIssuerRefSignSelfIssued(ctx, issuerRef).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	pkiIssuerSignSelfIssuedRequest := NewPkiIssuerSignSelfIssuedRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefSignSelfIssued(context.Background(), issuerRef, pkiIssuerSignSelfIssuedRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssuerSignSelfIssuedRequest** | [**PkiIssuerSignSelfIssuedRequest**](PkiIssuerSignSelfIssuedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatim

> PostPkiIssuerIssuerRefSignVerbatim(ctx, issuerRef).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefSignVerbatim(context.Background(), issuerRef, pkiIssuerSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatimRole

> PostPkiIssuerIssuerRefSignVerbatimRole(ctx, issuerRef, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role :=  // string | The desired role with configuration for this request
	
	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerIssuerRefSignVerbatimRole(context.Background(), issuerRef, role, pkiIssuerSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerRefDerPem

> PostPkiIssuerRefDerPem(ctx, issuerRef).PkiDerPemRequest(pkiDerPemRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	issuerRef :=  // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	
	pkiDerPemRequest := NewPkiDerPemRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuerRefDerPem(context.Background(), issuerRef, pkiDerPemRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiDerPemRequest** | [**PkiDerPemRequest**](PkiDerPemRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateIntermediateExported

> PostPkiIssuersGenerateIntermediateExported(ctx, exported).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	exported :=  // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	
	pkiIssuersGenerateIntermediateRequest := NewPkiIssuersGenerateIntermediateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuersGenerateIntermediateExported(context.Background(), exported, pkiIssuersGenerateIntermediateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssuersGenerateIntermediateRequest** | [**PkiIssuersGenerateIntermediateRequest**](PkiIssuersGenerateIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateRootExported

> PostPkiIssuersGenerateRootExported(ctx, exported).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	exported :=  // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	
	pkiIssuersGenerateRootRequest := NewPkiIssuersGenerateRootRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiIssuersGenerateRootExported(context.Background(), exported, pkiIssuersGenerateRootRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiIssuersGenerateRootRequest** | [**PkiIssuersGenerateRootRequest**](PkiIssuersGenerateRootRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiJson

> PostPkiJson(ctx).PkiJsonRequest(pkiJsonRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiJsonRequest := NewPkiJsonRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiJson(context.Background(), pkiJsonRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiJsonRequest** | [**PkiJsonRequest**](PkiJsonRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKeyKeyRef

> PostPkiKeyKeyRef(ctx, keyRef).PkiKeyRequest(pkiKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	keyRef :=  // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	
	pkiKeyRequest := NewPkiKeyRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiKeyKeyRef(context.Background(), keyRef, pkiKeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiKeyRequest** | [**PkiKeyRequest**](PkiKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKeysImport

> PostPkiKeysImport(ctx).PkiKeysImportRequest(pkiKeysImportRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiKeysImportRequest := NewPkiKeysImportRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiKeysImport(context.Background(), pkiKeysImportRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKeysImportRequest** | [**PkiKeysImportRequest**](PkiKeysImportRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKms

> PostPkiKms(ctx).PkiKmsRequest(pkiKmsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiKmsRequest := NewPkiKmsRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiKms(context.Background(), pkiKmsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKmsRequest** | [**PkiKmsRequest**](PkiKmsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRevoke

> PostPkiRevoke(ctx).PkiRevokeRequest(pkiRevokeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiRevokeRequest := NewPkiRevokeRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRevoke(context.Background(), pkiRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRevokeRequest** | [**PkiRevokeRequest**](PkiRevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRolesName

> PostPkiRolesName(ctx, name).PkiRolesRequest(pkiRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	pkiRolesRequest := NewPkiRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRolesName(context.Background(), name, pkiRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiRolesRequest** | [**PkiRolesRequest**](PkiRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootGenerateExported

> PostPkiRootGenerateExported(ctx, exported).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	exported :=  // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	
	pkiRootGenerateRequest := NewPkiRootGenerateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRootGenerateExported(context.Background(), exported, pkiRootGenerateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiRootGenerateRequest** | [**PkiRootGenerateRequest**](PkiRootGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootReplace

> PostPkiRootReplace(ctx).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiRootReplaceRequest := NewPkiRootReplaceRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRootReplace(context.Background(), pkiRootReplaceRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootReplaceRequest** | [**PkiRootReplaceRequest**](PkiRootReplaceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootRotateExported

> PostPkiRootRotateExported(ctx, exported).PkiRootRotateRequest(pkiRootRotateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	exported :=  // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	
	pkiRootRotateRequest := NewPkiRootRotateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRootRotateExported(context.Background(), exported, pkiRootRotateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiRootRotateRequest** | [**PkiRootRotateRequest**](PkiRootRotateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootSignIntermediate

> PostPkiRootSignIntermediate(ctx).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiRootSignIntermediateRequest := NewPkiRootSignIntermediateRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRootSignIntermediate(context.Background(), pkiRootSignIntermediateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignIntermediateRequest** | [**PkiRootSignIntermediateRequest**](PkiRootSignIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootSignSelfIssued

> PostPkiRootSignSelfIssued(ctx).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiRootSignSelfIssuedRequest := NewPkiRootSignSelfIssuedRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiRootSignSelfIssued(context.Background(), pkiRootSignSelfIssuedRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignSelfIssuedRequest** | [**PkiRootSignSelfIssuedRequest**](PkiRootSignSelfIssuedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignRole

> PostPkiSignRole(ctx, role).PkiSignRequest(pkiSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The desired role with configuration for this request
	
	pkiSignRequest := NewPkiSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiSignRole(context.Background(), role, pkiSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiSignRequest** | [**PkiSignRequest**](PkiSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignVerbatim

> PostPkiSignVerbatim(ctx).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiSignVerbatim(context.Background(), pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignVerbatimRole

> PostPkiSignVerbatimRole(ctx, role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The desired role with configuration for this request
	
	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiSignVerbatimRole(context.Background(), role, pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiTidy

> PostPkiTidy(ctx).PkiTidyRequest(pkiTidyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	pkiTidyRequest := NewPkiTidyRequestWithDefaults()
	
	resp, err := client.Secrets.PostPkiTidy(context.Background(), pkiTidyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiTidyRequest** | [**PkiTidyRequest**](PkiTidyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqConfigConnection

> PostRabbitmqConfigConnection(ctx).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()

Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	rabbitmqConfigConnectionRequest := NewRabbitmqConfigConnectionRequestWithDefaults()
	
	resp, err := client.Secrets.PostRabbitmqConfigConnection(context.Background(), rabbitmqConfigConnectionRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigConnectionRequest** | [**RabbitmqConfigConnectionRequest**](RabbitmqConfigConnectionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqConfigLease

> PostRabbitmqConfigLease(ctx).RabbitmqConfigLeaseRequest(rabbitmqConfigLeaseRequest).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	rabbitmqConfigLeaseRequest := NewRabbitmqConfigLeaseRequestWithDefaults()
	
	resp, err := client.Secrets.PostRabbitmqConfigLease(context.Background(), rabbitmqConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigLeaseRequest** | [**RabbitmqConfigLeaseRequest**](RabbitmqConfigLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqRolesName

> PostRabbitmqRolesName(ctx, name).RabbitmqRolesRequest(rabbitmqRolesRequest).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role.
	
	rabbitmqRolesRequest := NewRabbitmqRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostRabbitmqRolesName(context.Background(), name, rabbitmqRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **rabbitmqRolesRequest** | [**RabbitmqRolesRequest**](RabbitmqRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretConfig

> PostSecretConfig(ctx).KvConfigRequest(kvConfigRequest).Execute()

Configure backend level settings that are applied to every key in the key-value store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kvConfigRequest := NewKvConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretConfig(context.Background(), kvConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvConfigRequest** | [**KvConfigRequest**](KvConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDataPath

> PostSecretDataPath(ctx, path).KvDataRequest(kvDataRequest).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	kvDataRequest := NewKvDataRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretDataPath(context.Background(), path, kvDataRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kvDataRequest** | [**KvDataRequest**](KvDataRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDeletePath

> PostSecretDeletePath(ctx, path).KvDeleteRequest(kvDeleteRequest).Execute()

Marks one or more versions as deleted in the KV store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	kvDeleteRequest := NewKvDeleteRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretDeletePath(context.Background(), path, kvDeleteRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kvDeleteRequest** | [**KvDeleteRequest**](KvDeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDestroyPath

> PostSecretDestroyPath(ctx, path).KvDestroyRequest(kvDestroyRequest).Execute()

Permanently removes one or more versions in the KV store

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	kvDestroyRequest := NewKvDestroyRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretDestroyPath(context.Background(), path, kvDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kvDestroyRequest** | [**KvDestroyRequest**](KvDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretMetadataPath

> PostSecretMetadataPath(ctx, path).KvMetadataRequest(kvMetadataRequest).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	kvMetadataRequest := NewKvMetadataRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretMetadataPath(context.Background(), path, kvMetadataRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kvMetadataRequest** | [**KvMetadataRequest**](KvMetadataRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretUndeletePath

> PostSecretUndeletePath(ctx, path).KvUndeleteRequest(kvUndeleteRequest).Execute()

Undeletes one or more versions from the KV store.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Location of the secret.
	
	kvUndeleteRequest := NewKvUndeleteRequestWithDefaults()
	
	resp, err := client.Secrets.PostSecretUndeletePath(context.Background(), path, kvUndeleteRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **kvUndeleteRequest** | [**KvUndeleteRequest**](KvUndeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshConfigCa

> PostSshConfigCa(ctx).SshConfigCaRequest(sshConfigCaRequest).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	sshConfigCaRequest := NewSshConfigCaRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshConfigCa(context.Background(), sshConfigCaRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigCaRequest** | [**SshConfigCaRequest**](SshConfigCaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshConfigZeroaddress

> PostSshConfigZeroaddress(ctx).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	sshConfigZeroaddressRequest := NewSshConfigZeroaddressRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshConfigZeroaddress(context.Background(), sshConfigZeroaddressRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigZeroaddressRequest** | [**SshConfigZeroaddressRequest**](SshConfigZeroaddressRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshCredsRole

> PostSshCredsRole(ctx, role).SshCredsRequest(sshCredsRequest).Execute()

Creates a credential for establishing SSH connection with the remote host.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | [Required] Name of the role
	
	sshCredsRequest := NewSshCredsRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshCredsRole(context.Background(), role, sshCredsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **sshCredsRequest** | [**SshCredsRequest**](SshCredsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshKeysKeyName

> PostSshKeysKeyName(ctx, keyName).SshKeysRequest(sshKeysRequest).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	keyName :=  // string | [Required] Name of the key
	
	sshKeysRequest := NewSshKeysRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshKeysKeyName(context.Background(), keyName, sshKeysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **sshKeysRequest** | [**SshKeysRequest**](SshKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshLookup

> PostSshLookup(ctx).SshLookupRequest(sshLookupRequest).Execute()

List all the roles associated with the given IP address.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	sshLookupRequest := NewSshLookupRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshLookup(context.Background(), sshLookupRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshLookupRequest** | [**SshLookupRequest**](SshLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshRolesRole

> PostSshRolesRole(ctx, role).SshRolesRequest(sshRolesRequest).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | [Required for all types] Name of the role being created.
	
	sshRolesRequest := NewSshRolesRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshRolesRole(context.Background(), role, sshRolesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **sshRolesRequest** | [**SshRolesRequest**](SshRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshSignRole

> PostSshSignRole(ctx, role).SshSignRequest(sshSignRequest).Execute()

Request signing an SSH key using a certain role with the provided details.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The desired role with configuration for this request.
	
	sshSignRequest := NewSshSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshSignRole(context.Background(), role, sshSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **sshSignRequest** | [**SshSignRequest**](SshSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshVerify

> PostSshVerify(ctx).SshVerifyRequest(sshVerifyRequest).Execute()

Validate the OTP provided by Vault SSH Agent.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	sshVerifyRequest := NewSshVerifyRequestWithDefaults()
	
	resp, err := client.Secrets.PostSshVerify(context.Background(), sshVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshVerifyRequest** | [**SshVerifyRequest**](SshVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformConfig

> PostTerraformConfig(ctx).TerraformConfigRequest(terraformConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	terraformConfigRequest := NewTerraformConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostTerraformConfig(context.Background(), terraformConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformConfigRequest** | [**TerraformConfigRequest**](TerraformConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformCredsName

> PostTerraformCredsName(ctx, name).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Secrets.PostTerraformCredsName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTerraformRoleName

> PostTerraformRoleName(ctx, name).TerraformRoleRequest(terraformRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	terraformRoleRequest := NewTerraformRoleRequestWithDefaults()
	
	resp, err := client.Secrets.PostTerraformRoleName(context.Background(), name, terraformRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **terraformRoleRequest** | [**TerraformRoleRequest**](TerraformRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformRotateRoleName

> PostTerraformRotateRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the team or organization role
	
	resp, err := client.Secrets.PostTerraformRotateRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTotpCodeName

> PostTotpCodeName(ctx, name).TotpCodeRequest(totpCodeRequest).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key.
	
	totpCodeRequest := NewTotpCodeRequestWithDefaults()
	
	resp, err := client.Secrets.PostTotpCodeName(context.Background(), name, totpCodeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **totpCodeRequest** | [**TotpCodeRequest**](TotpCodeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTotpKeysName

> PostTotpKeysName(ctx, name).TotpKeysRequest(totpKeysRequest).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key.
	
	totpKeysRequest := NewTotpKeysRequestWithDefaults()
	
	resp, err := client.Secrets.PostTotpKeysName(context.Background(), name, totpKeysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **totpKeysRequest** | [**TotpKeysRequest**](TotpKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitCacheConfig

> PostTransitCacheConfig(ctx).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()

Configures a new cache of the specified size

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	transitCacheConfigRequest := NewTransitCacheConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitCacheConfig(context.Background(), transitCacheConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitCacheConfigRequest** | [**TransitCacheConfigRequest**](TransitCacheConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitDatakeyPlaintextName

> PostTransitDatakeyPlaintextName(ctx, name, plaintext).TransitDatakeyRequest(transitDatakeyRequest).Execute()

Generate a data key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The backend key used for encrypting the data key
	plaintext :=  // string | \"plaintext\" will return the key in both plaintext and ciphertext; \"wrapped\" will return the ciphertext only.
	
	transitDatakeyRequest := NewTransitDatakeyRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitDatakeyPlaintextName(context.Background(), name, plaintext, transitDatakeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **transitDatakeyRequest** | [**TransitDatakeyRequest**](TransitDatakeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitDecryptName

> PostTransitDecryptName(ctx, name).TransitDecryptRequest(transitDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the policy
	
	transitDecryptRequest := NewTransitDecryptRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitDecryptName(context.Background(), name, transitDecryptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitEncryptName

> PostTransitEncryptName(ctx, name).TransitEncryptRequest(transitEncryptRequest).Execute()

Encrypt a plaintext value or a batch of plaintext blocks using a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the policy
	
	transitEncryptRequest := NewTransitEncryptRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitEncryptName(context.Background(), name, transitEncryptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHash

> PostTransitHash(ctx).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	transitHashRequest := NewTransitHashRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitHash(context.Background(), transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHashUrlalgorithm

> PostTransitHashUrlalgorithm(ctx, urlalgorithm).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlalgorithm :=  // string | Algorithm to use (POST URL parameter)
	
	transitHashRequest := NewTransitHashRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitHashUrlalgorithm(context.Background(), urlalgorithm, transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHmacName

> PostTransitHmacName(ctx, name).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use for the HMAC function
	
	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitHmacName(context.Background(), name, transitHmacRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHmacNameUrlalgorithm

> PostTransitHmacNameUrlalgorithm(ctx, name, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use for the HMAC function
	urlalgorithm :=  // string | Algorithm to use (POST URL parameter)
	
	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitHmacNameUrlalgorithm(context.Background(), name, urlalgorithm, transitHmacRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysName

> PostTransitKeysName(ctx, name).TransitKeysRequest(transitKeysRequest).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	transitKeysRequest := NewTransitKeysRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitKeysName(context.Background(), name, transitKeysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitKeysRequest** | [**TransitKeysRequest**](TransitKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameConfig

> PostTransitKeysNameConfig(ctx, name).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()

Configure a named encryption key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	transitKeysConfigRequest := NewTransitKeysConfigRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitKeysNameConfig(context.Background(), name, transitKeysConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitKeysConfigRequest** | [**TransitKeysConfigRequest**](TransitKeysConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameImport

> PostTransitKeysNameImport(ctx, name).TransitKeysImportRequest(transitKeysImportRequest).Execute()

Imports an externally-generated key into a new transit key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the key
	
	transitKeysImportRequest := NewTransitKeysImportRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitKeysNameImport(context.Background(), name, transitKeysImportRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitKeysImportRequest** | [**TransitKeysImportRequest**](TransitKeysImportRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameImportVersion

> PostTransitKeysNameImportVersion(ctx, name).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()

Imports an externally-generated key into an existing imported key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the key
	
	transitKeysImportVersionRequest := NewTransitKeysImportVersionRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitKeysNameImportVersion(context.Background(), name, transitKeysImportVersionRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitKeysImportVersionRequest** | [**TransitKeysImportVersionRequest**](TransitKeysImportVersionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameRotate

> PostTransitKeysNameRotate(ctx, name).Execute()

Rotate named encryption key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Secrets.PostTransitKeysNameRotate(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTransitKeysNameTrim

> PostTransitKeysNameTrim(ctx, name).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()

Trim key versions of a named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	transitKeysTrimRequest := NewTransitKeysTrimRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitKeysNameTrim(context.Background(), name, transitKeysTrimRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitKeysTrimRequest** | [**TransitKeysTrimRequest**](TransitKeysTrimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandom

> PostTransitRandom(ctx).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRandom(context.Background(), transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomSource

> PostTransitRandomSource(ctx, source).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	source :=  // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	
	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRandomSource(context.Background(), source, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomSourceUrlbytes

> PostTransitRandomSourceUrlbytes(ctx, source, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	source :=  // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes :=  // string | The number of bytes to generate (POST URL parameter)
	
	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRandomSourceUrlbytes(context.Background(), source, urlbytes, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomUrlbytes

> PostTransitRandomUrlbytes(ctx, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlbytes :=  // string | The number of bytes to generate (POST URL parameter)
	
	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRandomUrlbytes(context.Background(), urlbytes, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRestore

> PostTransitRestore(ctx).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRestore(context.Background(), transitRestoreRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRestoreName

> PostTransitRestoreName(ctx, name).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | If set, this will be the name of the restored key.
	
	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRestoreName(context.Background(), name, transitRestoreRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRewrapName

> PostTransitRewrapName(ctx, name).TransitRewrapRequest(transitRewrapRequest).Execute()

Rewrap ciphertext

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	transitRewrapRequest := NewTransitRewrapRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitRewrapName(context.Background(), name, transitRewrapRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTransitSignName

> PostTransitSignName(ctx, name).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use
	
	transitSignRequest := NewTransitSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitSignName(context.Background(), name, transitSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTransitSignNameUrlalgorithm

> PostTransitSignNameUrlalgorithm(ctx, name, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use
	urlalgorithm :=  // string | Hash algorithm to use (POST URL parameter)
	
	transitSignRequest := NewTransitSignRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitSignNameUrlalgorithm(context.Background(), name, urlalgorithm, transitSignRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitVerifyName

> PostTransitVerifyName(ctx, name).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use
	
	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitVerifyName(context.Background(), name, transitVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostTransitVerifyNameUrlalgorithm

> PostTransitVerifyNameUrlalgorithm(ctx, name, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The key to use
	urlalgorithm :=  // string | Hash algorithm to use (POST URL parameter)
	
	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	
	resp, err := client.Secrets.PostTransitVerifyNameUrlalgorithm(context.Background(), name, urlalgorithm, transitVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

