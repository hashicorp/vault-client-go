# \Secrets

All URIs are relative to *http://localhost*

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
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Secrets.DeleteAdConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAdConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAdLibraryName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAdLibraryName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAdRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAdRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAlicloudConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAlicloudConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAlicloudRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAlicloudRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAwsRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAwsRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAzureConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteAzureRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteAzureRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteConsulRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteConsulRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteCubbyholePath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteCubbyholePath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpRolesetName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpRolesetName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpStaticAccountName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpStaticAccountName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpkmsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpkmsConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpkmsKeysDeregisterKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpkmsKeysDeregisterKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpkmsKeysKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpkmsKeysKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteGcpkmsKeysTrimKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteGcpkmsKeysTrimKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteKubernetesConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteKubernetesConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteKubernetesRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteKubernetesRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteKvPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteKvPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteMongodbatlasRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteMongodbatlasRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteNomadConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteNomadConfigAccess``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteNomadConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteNomadConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteNomadRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteNomadRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteOpenldapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteOpenldapConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteOpenldapRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteOpenldapRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteOpenldapStaticRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteOpenldapStaticRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeletePkiIssuerRefDerPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeletePkiIssuerRefDerPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeletePkiJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeletePkiJson``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeletePkiKeyKeyRef(context.Background(), keyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeletePkiKeyKeyRef``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeletePkiRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeletePkiRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeletePkiRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeletePkiRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteRabbitmqRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteRabbitmqRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSecretDataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSecretDataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSecretMetadataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSecretMetadataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSshConfigCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSshConfigCa``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSshConfigZeroaddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSshConfigZeroaddress``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSshKeysKeyName(context.Background(), keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSshKeysKeyName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteSshRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteSshRolesRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteTerraformConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteTerraformConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteTerraformRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteTerraformRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteTotpKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteTotpKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.DeleteTransitKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.DeleteTransitKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdLibrary(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdLibrary``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdLibraryName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdLibraryName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdLibraryNameStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdLibraryNameStatus``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAdRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAdRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAlicloudConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAlicloudConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAlicloudCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAlicloudCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAlicloudRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAlicloudRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAlicloudRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAlicloudRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsConfigRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsConfigRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsCreds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsCreds``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAwsStsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAwsStsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAzureConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAzureCredsRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAzureCredsRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAzureRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAzureRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetAzureRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetAzureRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetConsulConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetConsulConfigAccess``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetConsulCredsRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetConsulCredsRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetConsulRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetConsulRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetConsulRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetConsulRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetCubbyholePath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetCubbyholePath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpKeyRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpKeyRoleset``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpRolesetName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpRolesetName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpRolesetRolesetKey(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpRolesetRolesetKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpRolesetRolesetToken(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpRolesetRolesetToken``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpRolesets(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpRolesets``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpStaticAccountName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpStaticAccountName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpStaticAccountNameKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpStaticAccountNameKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpStaticAccountNameToken(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpStaticAccountNameToken``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpStaticAccounts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpStaticAccounts``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpTokenRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpTokenRoleset``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpkmsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpkmsConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpkmsKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpkmsKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpkmsKeysConfigKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpkmsKeysConfigKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpkmsKeysKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpkmsKeysKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetGcpkmsPubkeyKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetGcpkmsPubkeyKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetKubernetesConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetKubernetesConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetKubernetesRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetKubernetesRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetKubernetesRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetKubernetesRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetKvPath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetKvPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetMongodbatlasConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetMongodbatlasConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetMongodbatlasCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetMongodbatlasCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetMongodbatlasRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetMongodbatlasRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetMongodbatlasRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetMongodbatlasRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetNomadConfigAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetNomadConfigAccess``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetNomadConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetNomadConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetNomadCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetNomadCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetNomadRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetNomadRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetNomadRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetNomadRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapStaticCredName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapStaticCredName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapStaticRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapStaticRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetOpenldapStaticRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetOpenldapStaticRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCa``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCaChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCaChain``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCaPem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCaPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCertCaChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCertCaChain``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCertCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCertCrl``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCertSerial(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCertSerial``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCertSerialRaw(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCertSerialRaw``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCertSerialRawPem(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCertSerialRawPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCerts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCerts``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiConfigCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiConfigCrl``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiConfigIssuers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiConfigIssuers``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiConfigKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiConfigKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiConfigUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiConfigUrls``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCrl``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCrlPem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCrlPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiCrlRotate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiCrlRotate``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiDer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiDer``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiIssuerRefCrlPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiIssuerRefCrlPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiIssuerRefDerPem(context.Background(), issuerRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiIssuerRefDerPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiIssuers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiIssuers``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiJson``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiKeyKeyRef(context.Background(), keyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiKeyKeyRef``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetPkiTidyStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetPkiTidyStatus``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetRabbitmqConfigLease(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetRabbitmqConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetRabbitmqCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetRabbitmqCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetRabbitmqRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetRabbitmqRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetRabbitmqRolesName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetRabbitmqRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSecretConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSecretConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSecretDataPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSecretDataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSecretMetadataPath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSecretMetadataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSecretSubkeysPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSecretSubkeysPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSshConfigCa(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSshConfigCa``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSshConfigZeroaddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSshConfigZeroaddress``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSshPublicKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSshPublicKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSshRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSshRoles``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetSshRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetSshRolesRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTerraformConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTerraformConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTerraformCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTerraformCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTerraformRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTerraformRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTerraformRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTerraformRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTotpCodeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTotpCodeName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTotpKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTotpKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTotpKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTotpKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitBackupName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitBackupName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitCacheConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitCacheConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitExportTypeName(context.Background(), name, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitExportTypeName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitExportTypeNameVersion(context.Background(), name, type_, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitExportTypeNameVersion``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitKeys(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitKeysName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.GetTransitWrappingKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.GetTransitWrappingKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdConfig(context.Background()).AdConfigRequest(adConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdLibraryManageNameCheckIn(context.Background(), name).AdLibraryManageCheckInRequest(adLibraryManageCheckInRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdLibraryManageNameCheckIn``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdLibraryName(context.Background(), name).AdLibraryRequest(adLibraryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdLibraryName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdLibraryNameCheckIn(context.Background(), name).AdLibraryCheckInRequest(adLibraryCheckInRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdLibraryNameCheckIn``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdLibraryNameCheckOut(context.Background(), name).AdLibraryCheckOutRequest(adLibraryCheckOutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdLibraryNameCheckOut``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdRolesName(context.Background(), name).AdRolesRequest(adRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdRotateRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAdRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAdRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAlicloudConfig(context.Background()).AlicloudConfigRequest(alicloudConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAlicloudConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAlicloudRoleName(context.Background(), name).AlicloudRoleRequest(alicloudRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAlicloudRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsConfigLease(context.Background()).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsConfigRoot(context.Background()).AwsConfigRootRequest(awsConfigRootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsConfigRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsConfigRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsConfigRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsCreds(context.Background()).AwsCredsRequest(awsCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsCreds``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsRolesName(context.Background(), name).AwsRolesRequest(awsRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAwsStsName(context.Background(), name).AwsStsRequest(awsStsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAwsStsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAzureConfig(context.Background()).AzureConfigRequest(azureConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAzureConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAzureRolesName(context.Background(), name).AzureRolesRequest(azureRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAzureRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostAzureRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostAzureRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostConsulConfigAccess(context.Background()).ConsulConfigAccessRequest(consulConfigAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostConsulConfigAccess``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostConsulRolesName(context.Background(), name).ConsulRolesRequest(consulRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostConsulRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostCubbyholePath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostCubbyholePath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpConfig(context.Background()).GcpConfigRequest(gcpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpConfigRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpConfigRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpKeyRoleset(context.Background(), roleset).GcpKeyRequest(gcpKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpKeyRoleset``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpRolesetName(context.Background(), name).GcpRolesetRequest(gcpRolesetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpRolesetName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpRolesetNameRotate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpRolesetNameRotate``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpRolesetNameRotateKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpRolesetNameRotateKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpRolesetRolesetKey(context.Background(), roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpRolesetRolesetKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpRolesetRolesetToken(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpRolesetRolesetToken``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpStaticAccountName(context.Background(), name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpStaticAccountName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpStaticAccountNameKey(context.Background(), name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpStaticAccountNameKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpStaticAccountNameRotateKey(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpStaticAccountNameRotateKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpStaticAccountNameToken(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpStaticAccountNameToken``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpTokenRoleset(context.Background(), roleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpTokenRoleset``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsConfig(context.Background()).GcpkmsConfigRequest(gcpkmsConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsDecryptKey(context.Background(), key).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsDecryptKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsEncryptKey(context.Background(), key).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsEncryptKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysConfigKey(context.Background(), key).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysConfigKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysDeregisterKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysDeregisterKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysKey(context.Background(), key).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysRegisterKey(context.Background(), key).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysRegisterKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysRotateKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysRotateKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsKeysTrimKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsKeysTrimKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsReencryptKey(context.Background(), key).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsReencryptKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsSignKey(context.Background(), key).GcpkmsSignRequest(gcpkmsSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsSignKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostGcpkmsVerifyKey(context.Background(), key).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostGcpkmsVerifyKey``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostKubernetesConfig(context.Background()).KubernetesConfigRequest(kubernetesConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostKubernetesConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostKubernetesCredsName(context.Background(), name).KubernetesCredsRequest(kubernetesCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostKubernetesCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostKubernetesRolesName(context.Background(), name).KubernetesRolesRequest(kubernetesRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostKubernetesRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostKvPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostKvPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostMongodbatlasConfig(context.Background()).MongodbatlasConfigRequest(mongodbatlasConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostMongodbatlasConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostMongodbatlasCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostMongodbatlasCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostMongodbatlasRolesName(context.Background(), name).MongodbatlasRolesRequest(mongodbatlasRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostMongodbatlasRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostNomadConfigAccess(context.Background()).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostNomadConfigAccess``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostNomadConfigLease(context.Background()).NomadConfigLeaseRequest(nomadConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostNomadConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostNomadRoleName(context.Background(), name).NomadRoleRequest(nomadRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostNomadRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostOpenldapConfig(context.Background()).OpenldapConfigRequest(openldapConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostOpenldapConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostOpenldapRoleName(context.Background(), name).OpenldapRoleRequest(openldapRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostOpenldapRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostOpenldapRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostOpenldapRotateRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostOpenldapRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostOpenldapRotateRoot``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostOpenldapStaticRoleName(context.Background(), name).OpenldapStaticRoleRequest(openldapStaticRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostOpenldapStaticRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiBundle(context.Background()).PkiBundleRequest(pkiBundleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiBundle``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiCert(context.Background()).PkiCertRequest(pkiCertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiCert``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiConfigCa(context.Background()).PkiConfigCaRequest(pkiConfigCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiConfigCa``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiConfigCrl(context.Background()).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiConfigCrl``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiConfigIssuers(context.Background()).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiConfigIssuers``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiConfigKeys(context.Background()).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiConfigKeys``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiConfigUrls(context.Background()).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiConfigUrls``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIntermediateCrossSign(context.Background()).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIntermediateCrossSign``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIntermediateGenerateExported(context.Background(), exported).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIntermediateGenerateExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIntermediateSetSigned(context.Background()).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIntermediateSetSigned``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiInternalExported(context.Background()).PkiInternalExportedRequest(pkiInternalExportedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiInternalExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssueRole(context.Background(), role).PkiIssueRequest(pkiIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssueRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefIssueRole(context.Background(), issuerRef, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefIssueRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefSignIntermediate(context.Background(), issuerRef).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefSignIntermediate``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefSignRole(context.Background(), issuerRef, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefSignRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefSignSelfIssued(context.Background(), issuerRef).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefSignSelfIssued``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefSignVerbatim(context.Background(), issuerRef).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefSignVerbatim``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerIssuerRefSignVerbatimRole(context.Background(), issuerRef, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerIssuerRefSignVerbatimRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuerRefDerPem(context.Background(), issuerRef).PkiDerPemRequest(pkiDerPemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuerRefDerPem``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuersGenerateIntermediateExported(context.Background(), exported).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuersGenerateIntermediateExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiIssuersGenerateRootExported(context.Background(), exported).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiIssuersGenerateRootExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiJson(context.Background()).PkiJsonRequest(pkiJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiJson``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiKeyKeyRef(context.Background(), keyRef).PkiKeyRequest(pkiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiKeyKeyRef``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiKeysImport(context.Background()).PkiKeysImportRequest(pkiKeysImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiKeysImport``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiKms(context.Background()).PkiKmsRequest(pkiKmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiKms``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRevoke(context.Background()).PkiRevokeRequest(pkiRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRevoke``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRolesName(context.Background(), name).PkiRolesRequest(pkiRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRootGenerateExported(context.Background(), exported).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRootGenerateExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRootReplace(context.Background()).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRootReplace``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRootRotateExported(context.Background(), exported).PkiRootRotateRequest(pkiRootRotateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRootRotateExported``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRootSignIntermediate(context.Background()).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRootSignIntermediate``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiRootSignSelfIssued(context.Background()).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiRootSignSelfIssued``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiSignRole(context.Background(), role).PkiSignRequest(pkiSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiSignRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiSignVerbatim(context.Background()).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiSignVerbatim``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiSignVerbatimRole(context.Background(), role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiSignVerbatimRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostPkiTidy(context.Background()).PkiTidyRequest(pkiTidyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostPkiTidy``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostRabbitmqConfigConnection(context.Background()).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostRabbitmqConfigConnection``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostRabbitmqConfigLease(context.Background()).RabbitmqConfigLeaseRequest(rabbitmqConfigLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostRabbitmqConfigLease``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostRabbitmqRolesName(context.Background(), name).RabbitmqRolesRequest(rabbitmqRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostRabbitmqRolesName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretConfig(context.Background()).KvConfigRequest(kvConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretDataPath(context.Background(), path).KvDataRequest(kvDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretDataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretDeletePath(context.Background(), path).KvDeleteRequest(kvDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretDeletePath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretDestroyPath(context.Background(), path).KvDestroyRequest(kvDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretDestroyPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretMetadataPath(context.Background(), path).KvMetadataRequest(kvMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretMetadataPath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSecretUndeletePath(context.Background(), path).KvUndeleteRequest(kvUndeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSecretUndeletePath``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshConfigCa(context.Background()).SshConfigCaRequest(sshConfigCaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshConfigCa``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshConfigZeroaddress(context.Background()).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshConfigZeroaddress``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshCredsRole(context.Background(), role).SshCredsRequest(sshCredsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshCredsRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshKeysKeyName(context.Background(), keyName).SshKeysRequest(sshKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshKeysKeyName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshLookup(context.Background()).SshLookupRequest(sshLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshLookup``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshRolesRole(context.Background(), role).SshRolesRequest(sshRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshRolesRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshSignRole(context.Background(), role).SshSignRequest(sshSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshSignRole``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostSshVerify(context.Background()).SshVerifyRequest(sshVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostSshVerify``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTerraformConfig(context.Background()).TerraformConfigRequest(terraformConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTerraformConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTerraformCredsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTerraformCredsName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTerraformRoleName(context.Background(), name).TerraformRoleRequest(terraformRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTerraformRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTerraformRotateRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTerraformRotateRoleName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTotpCodeName(context.Background(), name).TotpCodeRequest(totpCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTotpCodeName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTotpKeysName(context.Background(), name).TotpKeysRequest(totpKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTotpKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitCacheConfig(context.Background()).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitCacheConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitDatakeyPlaintextName(context.Background(), name, plaintext).TransitDatakeyRequest(transitDatakeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitDatakeyPlaintextName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitDecryptName(context.Background(), name).TransitDecryptRequest(transitDecryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitDecryptName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitEncryptName(context.Background(), name).TransitEncryptRequest(transitEncryptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitEncryptName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitHash(context.Background()).TransitHashRequest(transitHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitHash``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitHashUrlalgorithm(context.Background(), urlalgorithm).TransitHashRequest(transitHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitHashUrlalgorithm``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitHmacName(context.Background(), name).TransitHmacRequest(transitHmacRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitHmacName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitHmacNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitHmacNameUrlalgorithm``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysName(context.Background(), name).TransitKeysRequest(transitKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysNameConfig(context.Background(), name).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysNameConfig``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysNameImport(context.Background(), name).TransitKeysImportRequest(transitKeysImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysNameImport``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysNameImportVersion(context.Background(), name).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysNameImportVersion``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysNameRotate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysNameRotate``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitKeysNameTrim(context.Background(), name).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitKeysNameTrim``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRandom(context.Background()).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRandom``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRandomSource(context.Background(), source).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRandomSource``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRandomSourceUrlbytes(context.Background(), source, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRandomSourceUrlbytes``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRandomUrlbytes(context.Background(), urlbytes).TransitRandomRequest(transitRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRandomUrlbytes``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRestore(context.Background()).TransitRestoreRequest(transitRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRestore``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRestoreName(context.Background(), name).TransitRestoreRequest(transitRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRestoreName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitRewrapName(context.Background(), name).TransitRewrapRequest(transitRewrapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitRewrapName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitSignName(context.Background(), name).TransitSignRequest(transitSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitSignName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitSignNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitSignNameUrlalgorithm``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitVerifyName(context.Background(), name).TransitVerifyRequest(transitVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitVerifyName``: %v\n", err)
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
    resp, r, err := apiClient.Secrets.PostTransitVerifyNameUrlalgorithm(context.Background(), name, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Secrets.PostTransitVerifyNameUrlalgorithm``: %v\n", err)
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

