# Secrets

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdConfig**](Secrets.md#DeleteAdConfig) | **Delete** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**DeleteAdLibraryName**](Secrets.md#DeleteAdLibraryName) | **Delete** /{ad_mount_path}/library/{name} | Delete a library set.
[**DeleteAdRolesName**](Secrets.md#DeleteAdRolesName) | **Delete** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**DeleteAlicloudConfig**](Secrets.md#DeleteAlicloudConfig) | **Delete** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**DeleteAlicloudRoleName**](Secrets.md#DeleteAlicloudRoleName) | **Delete** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**DeleteAwsRolesName**](Secrets.md#DeleteAwsRolesName) | **Delete** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**DeleteAzureConfig**](Secrets.md#DeleteAzureConfig) | **Delete** /{azure_mount_path}/config | 
[**DeleteAzureRolesName**](Secrets.md#DeleteAzureRolesName) | **Delete** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**DeleteConsulRolesName**](Secrets.md#DeleteConsulRolesName) | **Delete** /{consul_mount_path}/roles/{name} | 
[**DeleteCubbyholePath**](Secrets.md#DeleteCubbyholePath) | **Delete** /{cubbyhole_mount_path}/{path} | Deletes the secret at the specified location.
[**DeleteGcpRolesetName**](Secrets.md#DeleteGcpRolesetName) | **Delete** /{gcp_mount_path}/roleset/{name} | 
[**DeleteGcpStaticAccountName**](Secrets.md#DeleteGcpStaticAccountName) | **Delete** /{gcp_mount_path}/static-account/{name} | 
[**DeleteGcpkmsConfig**](Secrets.md#DeleteGcpkmsConfig) | **Delete** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**DeleteGcpkmsKeysDeregisterKey**](Secrets.md#DeleteGcpkmsKeysDeregisterKey) | **Delete** /{gcpkms_mount_path}/keys/deregister/{key} | Deregister an existing key in Vault
[**DeleteGcpkmsKeysKey**](Secrets.md#DeleteGcpkmsKeysKey) | **Delete** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**DeleteGcpkmsKeysTrimKey**](Secrets.md#DeleteGcpkmsKeysTrimKey) | **Delete** /{gcpkms_mount_path}/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**DeleteKubernetesConfig**](Secrets.md#DeleteKubernetesConfig) | **Delete** /{kubernetes_mount_path}/config | 
[**DeleteKubernetesRolesName**](Secrets.md#DeleteKubernetesRolesName) | **Delete** /{kubernetes_mount_path}/roles/{name} | 
[**DeleteLdapConfig**](Secrets.md#DeleteLdapConfig) | **Delete** /{ldap_mount_path}/config | 
[**DeleteLdapLibraryName**](Secrets.md#DeleteLdapLibraryName) | **Delete** /{ldap_mount_path}/library/{name} | Delete a library set.
[**DeleteLdapRoleName**](Secrets.md#DeleteLdapRoleName) | **Delete** /{ldap_mount_path}/role/{name} | 
[**DeleteLdapStaticRoleName**](Secrets.md#DeleteLdapStaticRoleName) | **Delete** /{ldap_mount_path}/static-role/{name} | 
[**DeleteMongodbatlasRolesName**](Secrets.md#DeleteMongodbatlasRolesName) | **Delete** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**DeleteNomadConfigAccess**](Secrets.md#DeleteNomadConfigAccess) | **Delete** /{nomad_mount_path}/config/access | 
[**DeleteNomadConfigLease**](Secrets.md#DeleteNomadConfigLease) | **Delete** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**DeleteNomadRoleName**](Secrets.md#DeleteNomadRoleName) | **Delete** /{nomad_mount_path}/role/{name} | 
[**DeleteOpenldapConfig**](Secrets.md#DeleteOpenldapConfig) | **Delete** /{openldap_mount_path}/config | 
[**DeleteOpenldapLibraryName**](Secrets.md#DeleteOpenldapLibraryName) | **Delete** /{openldap_mount_path}/library/{name} | Delete a library set.
[**DeleteOpenldapRoleName**](Secrets.md#DeleteOpenldapRoleName) | **Delete** /{openldap_mount_path}/role/{name} | 
[**DeleteOpenldapStaticRoleName**](Secrets.md#DeleteOpenldapStaticRoleName) | **Delete** /{openldap_mount_path}/static-role/{name} | 
[**DeletePkiIssuerRefDerPem**](Secrets.md#DeletePkiIssuerRefDerPem) | **Delete** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**DeletePkiJson**](Secrets.md#DeletePkiJson) | **Delete** /{pki_mount_path}//json | 
[**DeletePkiKeyKeyRef**](Secrets.md#DeletePkiKeyKeyRef) | **Delete** /{pki_mount_path}/key/{key_ref} | 
[**DeletePkiRolesName**](Secrets.md#DeletePkiRolesName) | **Delete** /{pki_mount_path}/roles/{name} | 
[**DeletePkiRoot**](Secrets.md#DeletePkiRoot) | **Delete** /{pki_mount_path}/root | 
[**DeleteRabbitmqRolesName**](Secrets.md#DeleteRabbitmqRolesName) | **Delete** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**DeleteSecretDataPath**](Secrets.md#DeleteSecretDataPath) | **Delete** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**DeleteSecretMetadataPath**](Secrets.md#DeleteSecretMetadataPath) | **Delete** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**DeleteSecretPath**](Secrets.md#DeleteSecretPath) | **Delete** /{secret_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**DeleteSshConfigCa**](Secrets.md#DeleteSshConfigCa) | **Delete** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**DeleteSshConfigZeroaddress**](Secrets.md#DeleteSshConfigZeroaddress) | **Delete** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**DeleteSshKeysKeyName**](Secrets.md#DeleteSshKeysKeyName) | **Delete** /{ssh_mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**DeleteSshRolesRole**](Secrets.md#DeleteSshRolesRole) | **Delete** /{ssh_mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**DeleteTerraformConfig**](Secrets.md#DeleteTerraformConfig) | **Delete** /{terraform_mount_path}/config | 
[**DeleteTerraformRoleName**](Secrets.md#DeleteTerraformRoleName) | **Delete** /{terraform_mount_path}/role/{name} | 
[**DeleteTotpKeysName**](Secrets.md#DeleteTotpKeysName) | **Delete** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**DeleteTransitKeysName**](Secrets.md#DeleteTransitKeysName) | **Delete** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**GetAdConfig**](Secrets.md#GetAdConfig) | **Get** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**GetAdCredsName**](Secrets.md#GetAdCredsName) | **Get** /{ad_mount_path}/creds/{name} | 
[**GetAdLibrary**](Secrets.md#GetAdLibrary) | **Get** /{ad_mount_path}/library | 
[**GetAdLibraryName**](Secrets.md#GetAdLibraryName) | **Get** /{ad_mount_path}/library/{name} | Read a library set.
[**GetAdLibraryNameStatus**](Secrets.md#GetAdLibraryNameStatus) | **Get** /{ad_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**GetAdRoles**](Secrets.md#GetAdRoles) | **Get** /{ad_mount_path}/roles | List the name of each role currently stored.
[**GetAdRolesName**](Secrets.md#GetAdRolesName) | **Get** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**GetAdRotateRoot**](Secrets.md#GetAdRotateRoot) | **Get** /{ad_mount_path}/rotate-root | 
[**GetAlicloudConfig**](Secrets.md#GetAlicloudConfig) | **Get** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**GetAlicloudCredsName**](Secrets.md#GetAlicloudCredsName) | **Get** /{alicloud_mount_path}/creds/{name} | Generate an API key or STS credential using the given role&#39;s configuration.&#39;
[**GetAlicloudRole**](Secrets.md#GetAlicloudRole) | **Get** /{alicloud_mount_path}/role | List the existing roles in this backend.
[**GetAlicloudRoleName**](Secrets.md#GetAlicloudRoleName) | **Get** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**GetAwsConfigLease**](Secrets.md#GetAwsConfigLease) | **Get** /{aws_mount_path}/config/lease | Configure the default lease information for generated credentials.
[**GetAwsConfigRoot**](Secrets.md#GetAwsConfigRoot) | **Get** /{aws_mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**GetAwsCreds**](Secrets.md#GetAwsCreds) | **Get** /{aws_mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**GetAwsRoles**](Secrets.md#GetAwsRoles) | **Get** /{aws_mount_path}/roles | List the existing roles in this backend
[**GetAwsRolesName**](Secrets.md#GetAwsRolesName) | **Get** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**GetAwsStsName**](Secrets.md#GetAwsStsName) | **Get** /{aws_mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**GetAzureConfig**](Secrets.md#GetAzureConfig) | **Get** /{azure_mount_path}/config | 
[**GetAzureCredsRole**](Secrets.md#GetAzureCredsRole) | **Get** /{azure_mount_path}/creds/{role} | 
[**GetAzureRoles**](Secrets.md#GetAzureRoles) | **Get** /{azure_mount_path}/roles | List existing roles.
[**GetAzureRolesName**](Secrets.md#GetAzureRolesName) | **Get** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**GetConsulConfigAccess**](Secrets.md#GetConsulConfigAccess) | **Get** /{consul_mount_path}/config/access | 
[**GetConsulCredsRole**](Secrets.md#GetConsulCredsRole) | **Get** /{consul_mount_path}/creds/{role} | 
[**GetConsulRoles**](Secrets.md#GetConsulRoles) | **Get** /{consul_mount_path}/roles | 
[**GetConsulRolesName**](Secrets.md#GetConsulRolesName) | **Get** /{consul_mount_path}/roles/{name} | 
[**GetCubbyholePath**](Secrets.md#GetCubbyholePath) | **Get** /{cubbyhole_mount_path}/{path} | Retrieve the secret at the specified location.
[**GetGcpConfig**](Secrets.md#GetGcpConfig) | **Get** /{gcp_mount_path}/config | 
[**GetGcpKeyRoleset**](Secrets.md#GetGcpKeyRoleset) | **Get** /{gcp_mount_path}/key/{roleset} | 
[**GetGcpRolesetName**](Secrets.md#GetGcpRolesetName) | **Get** /{gcp_mount_path}/roleset/{name} | 
[**GetGcpRolesetRolesetKey**](Secrets.md#GetGcpRolesetRolesetKey) | **Get** /{gcp_mount_path}/roleset/{roleset}/key | 
[**GetGcpRolesetRolesetToken**](Secrets.md#GetGcpRolesetRolesetToken) | **Get** /{gcp_mount_path}/roleset/{roleset}/token | 
[**GetGcpRolesets**](Secrets.md#GetGcpRolesets) | **Get** /{gcp_mount_path}/rolesets | 
[**GetGcpStaticAccountName**](Secrets.md#GetGcpStaticAccountName) | **Get** /{gcp_mount_path}/static-account/{name} | 
[**GetGcpStaticAccountNameKey**](Secrets.md#GetGcpStaticAccountNameKey) | **Get** /{gcp_mount_path}/static-account/{name}/key | 
[**GetGcpStaticAccountNameToken**](Secrets.md#GetGcpStaticAccountNameToken) | **Get** /{gcp_mount_path}/static-account/{name}/token | 
[**GetGcpStaticAccounts**](Secrets.md#GetGcpStaticAccounts) | **Get** /{gcp_mount_path}/static-accounts | 
[**GetGcpTokenRoleset**](Secrets.md#GetGcpTokenRoleset) | **Get** /{gcp_mount_path}/token/{roleset} | 
[**GetGcpkmsConfig**](Secrets.md#GetGcpkmsConfig) | **Get** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**GetGcpkmsKeys**](Secrets.md#GetGcpkmsKeys) | **Get** /{gcpkms_mount_path}/keys | List named keys
[**GetGcpkmsKeysConfigKey**](Secrets.md#GetGcpkmsKeysConfigKey) | **Get** /{gcpkms_mount_path}/keys/config/{key} | Configure the key in Vault
[**GetGcpkmsKeysKey**](Secrets.md#GetGcpkmsKeysKey) | **Get** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**GetGcpkmsPubkeyKey**](Secrets.md#GetGcpkmsPubkeyKey) | **Get** /{gcpkms_mount_path}/pubkey/{key} | Retrieve the public key associated with the named key
[**GetKubernetesConfig**](Secrets.md#GetKubernetesConfig) | **Get** /{kubernetes_mount_path}/config | 
[**GetKubernetesRoles**](Secrets.md#GetKubernetesRoles) | **Get** /{kubernetes_mount_path}/roles | 
[**GetKubernetesRolesName**](Secrets.md#GetKubernetesRolesName) | **Get** /{kubernetes_mount_path}/roles/{name} | 
[**GetLdapConfig**](Secrets.md#GetLdapConfig) | **Get** /{ldap_mount_path}/config | 
[**GetLdapCredsName**](Secrets.md#GetLdapCredsName) | **Get** /{ldap_mount_path}/creds/{name} | 
[**GetLdapLibrary**](Secrets.md#GetLdapLibrary) | **Get** /{ldap_mount_path}/library | 
[**GetLdapLibraryName**](Secrets.md#GetLdapLibraryName) | **Get** /{ldap_mount_path}/library/{name} | Read a library set.
[**GetLdapLibraryNameStatus**](Secrets.md#GetLdapLibraryNameStatus) | **Get** /{ldap_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**GetLdapRole**](Secrets.md#GetLdapRole) | **Get** /{ldap_mount_path}/role | 
[**GetLdapRoleName**](Secrets.md#GetLdapRoleName) | **Get** /{ldap_mount_path}/role/{name} | 
[**GetLdapStaticCredName**](Secrets.md#GetLdapStaticCredName) | **Get** /{ldap_mount_path}/static-cred/{name} | 
[**GetLdapStaticRole**](Secrets.md#GetLdapStaticRole) | **Get** /{ldap_mount_path}/static-role | 
[**GetLdapStaticRoleName**](Secrets.md#GetLdapStaticRoleName) | **Get** /{ldap_mount_path}/static-role/{name} | 
[**GetMongodbatlasConfig**](Secrets.md#GetMongodbatlasConfig) | **Get** /{mongodbatlas_mount_path}/config | Configure the  credentials that are used to manage Database Users.
[**GetMongodbatlasCredsName**](Secrets.md#GetMongodbatlasCredsName) | **Get** /{mongodbatlas_mount_path}/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**GetMongodbatlasRoles**](Secrets.md#GetMongodbatlasRoles) | **Get** /{mongodbatlas_mount_path}/roles | List the existing roles in this backend
[**GetMongodbatlasRolesName**](Secrets.md#GetMongodbatlasRolesName) | **Get** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**GetNomadConfigAccess**](Secrets.md#GetNomadConfigAccess) | **Get** /{nomad_mount_path}/config/access | 
[**GetNomadConfigLease**](Secrets.md#GetNomadConfigLease) | **Get** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**GetNomadCredsName**](Secrets.md#GetNomadCredsName) | **Get** /{nomad_mount_path}/creds/{name} | 
[**GetNomadRole**](Secrets.md#GetNomadRole) | **Get** /{nomad_mount_path}/role | 
[**GetNomadRoleName**](Secrets.md#GetNomadRoleName) | **Get** /{nomad_mount_path}/role/{name} | 
[**GetOpenldapConfig**](Secrets.md#GetOpenldapConfig) | **Get** /{openldap_mount_path}/config | 
[**GetOpenldapCredsName**](Secrets.md#GetOpenldapCredsName) | **Get** /{openldap_mount_path}/creds/{name} | 
[**GetOpenldapLibrary**](Secrets.md#GetOpenldapLibrary) | **Get** /{openldap_mount_path}/library | 
[**GetOpenldapLibraryName**](Secrets.md#GetOpenldapLibraryName) | **Get** /{openldap_mount_path}/library/{name} | Read a library set.
[**GetOpenldapLibraryNameStatus**](Secrets.md#GetOpenldapLibraryNameStatus) | **Get** /{openldap_mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**GetOpenldapRole**](Secrets.md#GetOpenldapRole) | **Get** /{openldap_mount_path}/role | 
[**GetOpenldapRoleName**](Secrets.md#GetOpenldapRoleName) | **Get** /{openldap_mount_path}/role/{name} | 
[**GetOpenldapStaticCredName**](Secrets.md#GetOpenldapStaticCredName) | **Get** /{openldap_mount_path}/static-cred/{name} | 
[**GetOpenldapStaticRole**](Secrets.md#GetOpenldapStaticRole) | **Get** /{openldap_mount_path}/static-role | 
[**GetOpenldapStaticRoleName**](Secrets.md#GetOpenldapStaticRoleName) | **Get** /{openldap_mount_path}/static-role/{name} | 
[**GetPkiCa**](Secrets.md#GetPkiCa) | **Get** /{pki_mount_path}/ca | 
[**GetPkiCaChain**](Secrets.md#GetPkiCaChain) | **Get** /{pki_mount_path}/ca_chain | 
[**GetPkiCaPem**](Secrets.md#GetPkiCaPem) | **Get** /{pki_mount_path}/ca/pem | 
[**GetPkiCertCaChain**](Secrets.md#GetPkiCertCaChain) | **Get** /{pki_mount_path}/cert/ca_chain | 
[**GetPkiCertSerial**](Secrets.md#GetPkiCertSerial) | **Get** /{pki_mount_path}/cert/{serial} | 
[**GetPkiCertSerialRaw**](Secrets.md#GetPkiCertSerialRaw) | **Get** /{pki_mount_path}/cert/{serial}/raw | 
[**GetPkiCertSerialRawPem**](Secrets.md#GetPkiCertSerialRawPem) | **Get** /{pki_mount_path}/cert/{serial}/raw/pem | 
[**GetPkiCerts**](Secrets.md#GetPkiCerts) | **Get** /{pki_mount_path}/certs | 
[**GetPkiCertsRevoked**](Secrets.md#GetPkiCertsRevoked) | **Get** /{pki_mount_path}/certs/revoked | 
[**GetPkiConfigAutoTidy**](Secrets.md#GetPkiConfigAutoTidy) | **Get** /{pki_mount_path}/config/auto-tidy | 
[**GetPkiConfigCrl**](Secrets.md#GetPkiConfigCrl) | **Get** /{pki_mount_path}/config/crl | 
[**GetPkiConfigIssuers**](Secrets.md#GetPkiConfigIssuers) | **Get** /{pki_mount_path}/config/issuers | 
[**GetPkiConfigKeys**](Secrets.md#GetPkiConfigKeys) | **Get** /{pki_mount_path}/config/keys | 
[**GetPkiConfigUrls**](Secrets.md#GetPkiConfigUrls) | **Get** /{pki_mount_path}/config/urls | 
[**GetPkiCrl**](Secrets.md#GetPkiCrl) | **Get** /{pki_mount_path}/crl | 
[**GetPkiCrlRotate**](Secrets.md#GetPkiCrlRotate) | **Get** /{pki_mount_path}/crl/rotate | 
[**GetPkiCrlRotateDelta**](Secrets.md#GetPkiCrlRotateDelta) | **Get** /{pki_mount_path}/crl/rotate-delta | 
[**GetPkiDelta**](Secrets.md#GetPkiDelta) | **Get** /{pki_mount_path}//delta | 
[**GetPkiDeltaCrl**](Secrets.md#GetPkiDeltaCrl) | **Get** /{pki_mount_path}/delta-crl | 
[**GetPkiDeltaPem**](Secrets.md#GetPkiDeltaPem) | **Get** /{pki_mount_path}//delta/pem | 
[**GetPkiDer**](Secrets.md#GetPkiDer) | **Get** /{pki_mount_path}//der | 
[**GetPkiIssuerRefCrlPemDerDeltaPem**](Secrets.md#GetPkiIssuerRefCrlPemDerDeltaPem) | **Get** /{pki_mount_path}/{issuer_ref}/crl/pem|/der|/delta/pem | 
[**GetPkiIssuerRefDerPem**](Secrets.md#GetPkiIssuerRefDerPem) | **Get** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**GetPkiIssuers**](Secrets.md#GetPkiIssuers) | **Get** /{pki_mount_path}/issuers | 
[**GetPkiJson**](Secrets.md#GetPkiJson) | **Get** /{pki_mount_path}//json | 
[**GetPkiKeyKeyRef**](Secrets.md#GetPkiKeyKeyRef) | **Get** /{pki_mount_path}/key/{key_ref} | 
[**GetPkiKeys**](Secrets.md#GetPkiKeys) | **Get** /{pki_mount_path}/keys | 
[**GetPkiOcspReq**](Secrets.md#GetPkiOcspReq) | **Get** /{pki_mount_path}/ocsp/{req} | 
[**GetPkiPem**](Secrets.md#GetPkiPem) | **Get** /{pki_mount_path}//pem | 
[**GetPkiRoles**](Secrets.md#GetPkiRoles) | **Get** /{pki_mount_path}/roles | 
[**GetPkiRolesName**](Secrets.md#GetPkiRolesName) | **Get** /{pki_mount_path}/roles/{name} | 
[**GetPkiTidyStatus**](Secrets.md#GetPkiTidyStatus) | **Get** /{pki_mount_path}/tidy-status | 
[**GetRabbitmqConfigLease**](Secrets.md#GetRabbitmqConfigLease) | **Get** /{rabbitmq_mount_path}/config/lease | Configure the lease parameters for generated credentials
[**GetRabbitmqCredsName**](Secrets.md#GetRabbitmqCredsName) | **Get** /{rabbitmq_mount_path}/creds/{name} | Request RabbitMQ credentials for a certain role.
[**GetRabbitmqRoles**](Secrets.md#GetRabbitmqRoles) | **Get** /{rabbitmq_mount_path}/roles | Manage the roles that can be created with this backend.
[**GetRabbitmqRolesName**](Secrets.md#GetRabbitmqRolesName) | **Get** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**GetSecretConfig**](Secrets.md#GetSecretConfig) | **Get** /{secret_mount_path}/config | Read the backend level settings.
[**GetSecretDataPath**](Secrets.md#GetSecretDataPath) | **Get** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**GetSecretMetadataPath**](Secrets.md#GetSecretMetadataPath) | **Get** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**GetSecretPath**](Secrets.md#GetSecretPath) | **Get** /{secret_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**GetSecretSubkeysPath**](Secrets.md#GetSecretSubkeysPath) | **Get** /{secret_mount_path}/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
[**GetSshConfigCa**](Secrets.md#GetSshConfigCa) | **Get** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**GetSshConfigZeroaddress**](Secrets.md#GetSshConfigZeroaddress) | **Get** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**GetSshPublicKey**](Secrets.md#GetSshPublicKey) | **Get** /{ssh_mount_path}/public_key | Retrieve the public key.
[**GetSshRoles**](Secrets.md#GetSshRoles) | **Get** /{ssh_mount_path}/roles | Manage the &#39;roles&#39; that can be created with this backend.
[**GetSshRolesRole**](Secrets.md#GetSshRolesRole) | **Get** /{ssh_mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**GetTerraformConfig**](Secrets.md#GetTerraformConfig) | **Get** /{terraform_mount_path}/config | 
[**GetTerraformCredsName**](Secrets.md#GetTerraformCredsName) | **Get** /{terraform_mount_path}/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**GetTerraformRole**](Secrets.md#GetTerraformRole) | **Get** /{terraform_mount_path}/role | 
[**GetTerraformRoleName**](Secrets.md#GetTerraformRoleName) | **Get** /{terraform_mount_path}/role/{name} | 
[**GetTotpCodeName**](Secrets.md#GetTotpCodeName) | **Get** /{totp_mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**GetTotpKeys**](Secrets.md#GetTotpKeys) | **Get** /{totp_mount_path}/keys | Manage the keys that can be created with this backend.
[**GetTotpKeysName**](Secrets.md#GetTotpKeysName) | **Get** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**GetTransitBackupName**](Secrets.md#GetTransitBackupName) | **Get** /{transit_mount_path}/backup/{name} | Backup the named key
[**GetTransitCacheConfig**](Secrets.md#GetTransitCacheConfig) | **Get** /{transit_mount_path}/cache-config | Returns the size of the active cache
[**GetTransitExportTypeName**](Secrets.md#GetTransitExportTypeName) | **Get** /{transit_mount_path}/export/{type}/{name} | Export named encryption or signing key
[**GetTransitExportTypeNameVersion**](Secrets.md#GetTransitExportTypeNameVersion) | **Get** /{transit_mount_path}/export/{type}/{name}/{version} | Export named encryption or signing key
[**GetTransitKeys**](Secrets.md#GetTransitKeys) | **Get** /{transit_mount_path}/keys | Managed named encryption keys
[**GetTransitKeysName**](Secrets.md#GetTransitKeysName) | **Get** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**GetTransitWrappingKey**](Secrets.md#GetTransitWrappingKey) | **Get** /{transit_mount_path}/wrapping_key | Returns the public key to use for wrapping imported keys
[**PostAdConfig**](Secrets.md#PostAdConfig) | **Post** /{ad_mount_path}/config | Configure the AD server to connect to, along with password options.
[**PostAdLibraryManageNameCheckIn**](Secrets.md#PostAdLibraryManageNameCheckIn) | **Post** /{ad_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryName**](Secrets.md#PostAdLibraryName) | **Post** /{ad_mount_path}/library/{name} | Update a library set.
[**PostAdLibraryNameCheckIn**](Secrets.md#PostAdLibraryNameCheckIn) | **Post** /{ad_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**PostAdLibraryNameCheckOut**](Secrets.md#PostAdLibraryNameCheckOut) | **Post** /{ad_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**PostAdRolesName**](Secrets.md#PostAdRolesName) | **Post** /{ad_mount_path}/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
[**PostAdRotateRoleName**](Secrets.md#PostAdRotateRoleName) | **Post** /{ad_mount_path}/rotate-role/{name} | 
[**PostAdRotateRoot**](Secrets.md#PostAdRotateRoot) | **Post** /{ad_mount_path}/rotate-root | 
[**PostAlicloudConfig**](Secrets.md#PostAlicloudConfig) | **Post** /{alicloud_mount_path}/config | Configure the access key and secret to use for RAM and STS calls.
[**PostAlicloudRoleName**](Secrets.md#PostAlicloudRoleName) | **Post** /{alicloud_mount_path}/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
[**PostAwsConfigLease**](Secrets.md#PostAwsConfigLease) | **Post** /{aws_mount_path}/config/lease | Configure the default lease information for generated credentials.
[**PostAwsConfigRoot**](Secrets.md#PostAwsConfigRoot) | **Post** /{aws_mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**PostAwsConfigRotateRoot**](Secrets.md#PostAwsConfigRotateRoot) | **Post** /{aws_mount_path}/config/rotate-root | 
[**PostAwsCreds**](Secrets.md#PostAwsCreds) | **Post** /{aws_mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**PostAwsRolesName**](Secrets.md#PostAwsRolesName) | **Post** /{aws_mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**PostAwsStsName**](Secrets.md#PostAwsStsName) | **Post** /{aws_mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**PostAzureConfig**](Secrets.md#PostAzureConfig) | **Post** /{azure_mount_path}/config | 
[**PostAzureRolesName**](Secrets.md#PostAzureRolesName) | **Post** /{azure_mount_path}/roles/{name} | Manage the Vault roles used to generate Azure credentials.
[**PostAzureRotateRoot**](Secrets.md#PostAzureRotateRoot) | **Post** /{azure_mount_path}/rotate-root | 
[**PostConsulConfigAccess**](Secrets.md#PostConsulConfigAccess) | **Post** /{consul_mount_path}/config/access | 
[**PostConsulRolesName**](Secrets.md#PostConsulRolesName) | **Post** /{consul_mount_path}/roles/{name} | 
[**PostCubbyholePath**](Secrets.md#PostCubbyholePath) | **Post** /{cubbyhole_mount_path}/{path} | Store a secret at the specified location.
[**PostGcpConfig**](Secrets.md#PostGcpConfig) | **Post** /{gcp_mount_path}/config | 
[**PostGcpConfigRotateRoot**](Secrets.md#PostGcpConfigRotateRoot) | **Post** /{gcp_mount_path}/config/rotate-root | 
[**PostGcpKeyRoleset**](Secrets.md#PostGcpKeyRoleset) | **Post** /{gcp_mount_path}/key/{roleset} | 
[**PostGcpRolesetName**](Secrets.md#PostGcpRolesetName) | **Post** /{gcp_mount_path}/roleset/{name} | 
[**PostGcpRolesetNameRotate**](Secrets.md#PostGcpRolesetNameRotate) | **Post** /{gcp_mount_path}/roleset/{name}/rotate | 
[**PostGcpRolesetNameRotateKey**](Secrets.md#PostGcpRolesetNameRotateKey) | **Post** /{gcp_mount_path}/roleset/{name}/rotate-key | 
[**PostGcpRolesetRolesetKey**](Secrets.md#PostGcpRolesetRolesetKey) | **Post** /{gcp_mount_path}/roleset/{roleset}/key | 
[**PostGcpRolesetRolesetToken**](Secrets.md#PostGcpRolesetRolesetToken) | **Post** /{gcp_mount_path}/roleset/{roleset}/token | 
[**PostGcpStaticAccountName**](Secrets.md#PostGcpStaticAccountName) | **Post** /{gcp_mount_path}/static-account/{name} | 
[**PostGcpStaticAccountNameKey**](Secrets.md#PostGcpStaticAccountNameKey) | **Post** /{gcp_mount_path}/static-account/{name}/key | 
[**PostGcpStaticAccountNameRotateKey**](Secrets.md#PostGcpStaticAccountNameRotateKey) | **Post** /{gcp_mount_path}/static-account/{name}/rotate-key | 
[**PostGcpStaticAccountNameToken**](Secrets.md#PostGcpStaticAccountNameToken) | **Post** /{gcp_mount_path}/static-account/{name}/token | 
[**PostGcpTokenRoleset**](Secrets.md#PostGcpTokenRoleset) | **Post** /{gcp_mount_path}/token/{roleset} | 
[**PostGcpkmsConfig**](Secrets.md#PostGcpkmsConfig) | **Post** /{gcpkms_mount_path}/config | Configure the GCP KMS secrets engine
[**PostGcpkmsDecryptKey**](Secrets.md#PostGcpkmsDecryptKey) | **Post** /{gcpkms_mount_path}/decrypt/{key} | Decrypt a ciphertext value using a named key
[**PostGcpkmsEncryptKey**](Secrets.md#PostGcpkmsEncryptKey) | **Post** /{gcpkms_mount_path}/encrypt/{key} | Encrypt a plaintext value using a named key
[**PostGcpkmsKeysConfigKey**](Secrets.md#PostGcpkmsKeysConfigKey) | **Post** /{gcpkms_mount_path}/keys/config/{key} | Configure the key in Vault
[**PostGcpkmsKeysDeregisterKey**](Secrets.md#PostGcpkmsKeysDeregisterKey) | **Post** /{gcpkms_mount_path}/keys/deregister/{key} | Deregister an existing key in Vault
[**PostGcpkmsKeysKey**](Secrets.md#PostGcpkmsKeysKey) | **Post** /{gcpkms_mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**PostGcpkmsKeysRegisterKey**](Secrets.md#PostGcpkmsKeysRegisterKey) | **Post** /{gcpkms_mount_path}/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**PostGcpkmsKeysRotateKey**](Secrets.md#PostGcpkmsKeysRotateKey) | **Post** /{gcpkms_mount_path}/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**PostGcpkmsKeysTrimKey**](Secrets.md#PostGcpkmsKeysTrimKey) | **Post** /{gcpkms_mount_path}/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**PostGcpkmsReencryptKey**](Secrets.md#PostGcpkmsReencryptKey) | **Post** /{gcpkms_mount_path}/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**PostGcpkmsSignKey**](Secrets.md#PostGcpkmsSignKey) | **Post** /{gcpkms_mount_path}/sign/{key} | Signs a message or digest using a named key
[**PostGcpkmsVerifyKey**](Secrets.md#PostGcpkmsVerifyKey) | **Post** /{gcpkms_mount_path}/verify/{key} | Verify a signature using a named key
[**PostKubernetesConfig**](Secrets.md#PostKubernetesConfig) | **Post** /{kubernetes_mount_path}/config | 
[**PostKubernetesCredsName**](Secrets.md#PostKubernetesCredsName) | **Post** /{kubernetes_mount_path}/creds/{name} | 
[**PostKubernetesRolesName**](Secrets.md#PostKubernetesRolesName) | **Post** /{kubernetes_mount_path}/roles/{name} | 
[**PostLdapConfig**](Secrets.md#PostLdapConfig) | **Post** /{ldap_mount_path}/config | 
[**PostLdapLibraryManageNameCheckIn**](Secrets.md#PostLdapLibraryManageNameCheckIn) | **Post** /{ldap_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**PostLdapLibraryName**](Secrets.md#PostLdapLibraryName) | **Post** /{ldap_mount_path}/library/{name} | Update a library set.
[**PostLdapLibraryNameCheckIn**](Secrets.md#PostLdapLibraryNameCheckIn) | **Post** /{ldap_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**PostLdapLibraryNameCheckOut**](Secrets.md#PostLdapLibraryNameCheckOut) | **Post** /{ldap_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**PostLdapRoleName**](Secrets.md#PostLdapRoleName) | **Post** /{ldap_mount_path}/role/{name} | 
[**PostLdapRotateRoleName**](Secrets.md#PostLdapRotateRoleName) | **Post** /{ldap_mount_path}/rotate-role/{name} | 
[**PostLdapRotateRoot**](Secrets.md#PostLdapRotateRoot) | **Post** /{ldap_mount_path}/rotate-root | 
[**PostLdapStaticRoleName**](Secrets.md#PostLdapStaticRoleName) | **Post** /{ldap_mount_path}/static-role/{name} | 
[**PostMongodbatlasConfig**](Secrets.md#PostMongodbatlasConfig) | **Post** /{mongodbatlas_mount_path}/config | Configure the  credentials that are used to manage Database Users.
[**PostMongodbatlasCredsName**](Secrets.md#PostMongodbatlasCredsName) | **Post** /{mongodbatlas_mount_path}/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
[**PostMongodbatlasRolesName**](Secrets.md#PostMongodbatlasRolesName) | **Post** /{mongodbatlas_mount_path}/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
[**PostNomadConfigAccess**](Secrets.md#PostNomadConfigAccess) | **Post** /{nomad_mount_path}/config/access | 
[**PostNomadConfigLease**](Secrets.md#PostNomadConfigLease) | **Post** /{nomad_mount_path}/config/lease | Configure the lease parameters for generated tokens
[**PostNomadRoleName**](Secrets.md#PostNomadRoleName) | **Post** /{nomad_mount_path}/role/{name} | 
[**PostOpenldapConfig**](Secrets.md#PostOpenldapConfig) | **Post** /{openldap_mount_path}/config | 
[**PostOpenldapLibraryManageNameCheckIn**](Secrets.md#PostOpenldapLibraryManageNameCheckIn) | **Post** /{openldap_mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**PostOpenldapLibraryName**](Secrets.md#PostOpenldapLibraryName) | **Post** /{openldap_mount_path}/library/{name} | Update a library set.
[**PostOpenldapLibraryNameCheckIn**](Secrets.md#PostOpenldapLibraryNameCheckIn) | **Post** /{openldap_mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**PostOpenldapLibraryNameCheckOut**](Secrets.md#PostOpenldapLibraryNameCheckOut) | **Post** /{openldap_mount_path}/library/{name}/check-out | Check a service account out from the library.
[**PostOpenldapRoleName**](Secrets.md#PostOpenldapRoleName) | **Post** /{openldap_mount_path}/role/{name} | 
[**PostOpenldapRotateRoleName**](Secrets.md#PostOpenldapRotateRoleName) | **Post** /{openldap_mount_path}/rotate-role/{name} | 
[**PostOpenldapRotateRoot**](Secrets.md#PostOpenldapRotateRoot) | **Post** /{openldap_mount_path}/rotate-root | 
[**PostOpenldapStaticRoleName**](Secrets.md#PostOpenldapStaticRoleName) | **Post** /{openldap_mount_path}/static-role/{name} | 
[**PostPkiBundle**](Secrets.md#PostPkiBundle) | **Post** /{pki_mount_path}/bundle | 
[**PostPkiCert**](Secrets.md#PostPkiCert) | **Post** /{pki_mount_path}/cert | 
[**PostPkiConfigAutoTidy**](Secrets.md#PostPkiConfigAutoTidy) | **Post** /{pki_mount_path}/config/auto-tidy | 
[**PostPkiConfigCa**](Secrets.md#PostPkiConfigCa) | **Post** /{pki_mount_path}/config/ca | 
[**PostPkiConfigCrl**](Secrets.md#PostPkiConfigCrl) | **Post** /{pki_mount_path}/config/crl | 
[**PostPkiConfigIssuers**](Secrets.md#PostPkiConfigIssuers) | **Post** /{pki_mount_path}/config/issuers | 
[**PostPkiConfigKeys**](Secrets.md#PostPkiConfigKeys) | **Post** /{pki_mount_path}/config/keys | 
[**PostPkiConfigUrls**](Secrets.md#PostPkiConfigUrls) | **Post** /{pki_mount_path}/config/urls | 
[**PostPkiIntermediateCrossSign**](Secrets.md#PostPkiIntermediateCrossSign) | **Post** /{pki_mount_path}/intermediate/cross-sign | 
[**PostPkiIntermediateGenerateExported**](Secrets.md#PostPkiIntermediateGenerateExported) | **Post** /{pki_mount_path}/intermediate/generate/{exported} | 
[**PostPkiIntermediateSetSigned**](Secrets.md#PostPkiIntermediateSetSigned) | **Post** /{pki_mount_path}/intermediate/set-signed | 
[**PostPkiInternalExported**](Secrets.md#PostPkiInternalExported) | **Post** /{pki_mount_path}/internal|exported | 
[**PostPkiIssueRole**](Secrets.md#PostPkiIssueRole) | **Post** /{pki_mount_path}/issue/{role} | 
[**PostPkiIssuerIssuerRefIssueRole**](Secrets.md#PostPkiIssuerIssuerRefIssueRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/issue/{role} | 
[**PostPkiIssuerIssuerRefRevoke**](Secrets.md#PostPkiIssuerIssuerRefRevoke) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/revoke | 
[**PostPkiIssuerIssuerRefSignIntermediate**](Secrets.md#PostPkiIssuerIssuerRefSignIntermediate) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate | 
[**PostPkiIssuerIssuerRefSignRole**](Secrets.md#PostPkiIssuerIssuerRefSignRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign/{role} | 
[**PostPkiIssuerIssuerRefSignSelfIssued**](Secrets.md#PostPkiIssuerIssuerRefSignSelfIssued) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued | 
[**PostPkiIssuerIssuerRefSignVerbatim**](Secrets.md#PostPkiIssuerIssuerRefSignVerbatim) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim | 
[**PostPkiIssuerIssuerRefSignVerbatimRole**](Secrets.md#PostPkiIssuerIssuerRefSignVerbatimRole) | **Post** /{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**PostPkiIssuerRefDerPem**](Secrets.md#PostPkiIssuerRefDerPem) | **Post** /{pki_mount_path}/{issuer_ref}/der|/pem | 
[**PostPkiIssuersGenerateIntermediateExported**](Secrets.md#PostPkiIssuersGenerateIntermediateExported) | **Post** /{pki_mount_path}/issuers/generate/intermediate/{exported} | 
[**PostPkiIssuersGenerateRootExported**](Secrets.md#PostPkiIssuersGenerateRootExported) | **Post** /{pki_mount_path}/issuers/generate/root/{exported} | 
[**PostPkiJson**](Secrets.md#PostPkiJson) | **Post** /{pki_mount_path}//json | 
[**PostPkiKeyKeyRef**](Secrets.md#PostPkiKeyKeyRef) | **Post** /{pki_mount_path}/key/{key_ref} | 
[**PostPkiKeysImport**](Secrets.md#PostPkiKeysImport) | **Post** /{pki_mount_path}/keys/import | 
[**PostPkiKms**](Secrets.md#PostPkiKms) | **Post** /{pki_mount_path}/kms | 
[**PostPkiOcsp**](Secrets.md#PostPkiOcsp) | **Post** /{pki_mount_path}/ocsp | 
[**PostPkiRevoke**](Secrets.md#PostPkiRevoke) | **Post** /{pki_mount_path}/revoke | 
[**PostPkiRevokeWithKey**](Secrets.md#PostPkiRevokeWithKey) | **Post** /{pki_mount_path}/revoke-with-key | 
[**PostPkiRolesName**](Secrets.md#PostPkiRolesName) | **Post** /{pki_mount_path}/roles/{name} | 
[**PostPkiRootGenerateExported**](Secrets.md#PostPkiRootGenerateExported) | **Post** /{pki_mount_path}/root/generate/{exported} | 
[**PostPkiRootReplace**](Secrets.md#PostPkiRootReplace) | **Post** /{pki_mount_path}/root/replace | 
[**PostPkiRootRotateExported**](Secrets.md#PostPkiRootRotateExported) | **Post** /{pki_mount_path}/root/rotate/{exported} | 
[**PostPkiRootSignIntermediate**](Secrets.md#PostPkiRootSignIntermediate) | **Post** /{pki_mount_path}/root/sign-intermediate | 
[**PostPkiRootSignSelfIssued**](Secrets.md#PostPkiRootSignSelfIssued) | **Post** /{pki_mount_path}/root/sign-self-issued | 
[**PostPkiSignRole**](Secrets.md#PostPkiSignRole) | **Post** /{pki_mount_path}/sign/{role} | 
[**PostPkiSignVerbatim**](Secrets.md#PostPkiSignVerbatim) | **Post** /{pki_mount_path}/sign-verbatim | 
[**PostPkiSignVerbatimRole**](Secrets.md#PostPkiSignVerbatimRole) | **Post** /{pki_mount_path}/sign-verbatim/{role} | 
[**PostPkiTidy**](Secrets.md#PostPkiTidy) | **Post** /{pki_mount_path}/tidy | 
[**PostPkiTidyCancel**](Secrets.md#PostPkiTidyCancel) | **Post** /{pki_mount_path}/tidy-cancel | 
[**PostRabbitmqConfigConnection**](Secrets.md#PostRabbitmqConfigConnection) | **Post** /{rabbitmq_mount_path}/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**PostRabbitmqConfigLease**](Secrets.md#PostRabbitmqConfigLease) | **Post** /{rabbitmq_mount_path}/config/lease | Configure the lease parameters for generated credentials
[**PostRabbitmqRolesName**](Secrets.md#PostRabbitmqRolesName) | **Post** /{rabbitmq_mount_path}/roles/{name} | Manage the roles that can be created with this backend.
[**PostSecretConfig**](Secrets.md#PostSecretConfig) | **Post** /{secret_mount_path}/config | Configure backend level settings that are applied to every key in the key-value store.
[**PostSecretDataPath**](Secrets.md#PostSecretDataPath) | **Post** /{secret_mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**PostSecretDeletePath**](Secrets.md#PostSecretDeletePath) | **Post** /{secret_mount_path}/delete/{path} | Marks one or more versions as deleted in the KV store.
[**PostSecretDestroyPath**](Secrets.md#PostSecretDestroyPath) | **Post** /{secret_mount_path}/destroy/{path} | Permanently removes one or more versions in the KV store
[**PostSecretMetadataPath**](Secrets.md#PostSecretMetadataPath) | **Post** /{secret_mount_path}/metadata/{path} | Configures settings for the KV store
[**PostSecretPath**](Secrets.md#PostSecretPath) | **Post** /{secret_mount_path}/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
[**PostSecretUndeletePath**](Secrets.md#PostSecretUndeletePath) | **Post** /{secret_mount_path}/undelete/{path} | Undeletes one or more versions from the KV store.
[**PostSshConfigCa**](Secrets.md#PostSshConfigCa) | **Post** /{ssh_mount_path}/config/ca | Set the SSH private key used for signing certificates.
[**PostSshConfigZeroaddress**](Secrets.md#PostSshConfigZeroaddress) | **Post** /{ssh_mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**PostSshCredsRole**](Secrets.md#PostSshCredsRole) | **Post** /{ssh_mount_path}/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
[**PostSshIssueRole**](Secrets.md#PostSshIssueRole) | **Post** /{ssh_mount_path}/issue/{role} | 
[**PostSshKeysKeyName**](Secrets.md#PostSshKeysKeyName) | **Post** /{ssh_mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**PostSshLookup**](Secrets.md#PostSshLookup) | **Post** /{ssh_mount_path}/lookup | List all the roles associated with the given IP address.
[**PostSshRolesRole**](Secrets.md#PostSshRolesRole) | **Post** /{ssh_mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**PostSshSignRole**](Secrets.md#PostSshSignRole) | **Post** /{ssh_mount_path}/sign/{role} | Request signing an SSH key using a certain role with the provided details.
[**PostSshVerify**](Secrets.md#PostSshVerify) | **Post** /{ssh_mount_path}/verify | Validate the OTP provided by Vault SSH Agent.
[**PostTerraformConfig**](Secrets.md#PostTerraformConfig) | **Post** /{terraform_mount_path}/config | 
[**PostTerraformCredsName**](Secrets.md#PostTerraformCredsName) | **Post** /{terraform_mount_path}/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
[**PostTerraformRoleName**](Secrets.md#PostTerraformRoleName) | **Post** /{terraform_mount_path}/role/{name} | 
[**PostTerraformRotateRoleName**](Secrets.md#PostTerraformRotateRoleName) | **Post** /{terraform_mount_path}/rotate-role/{name} | 
[**PostTotpCodeName**](Secrets.md#PostTotpCodeName) | **Post** /{totp_mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**PostTotpKeysName**](Secrets.md#PostTotpKeysName) | **Post** /{totp_mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**PostTransitCacheConfig**](Secrets.md#PostTransitCacheConfig) | **Post** /{transit_mount_path}/cache-config | Configures a new cache of the specified size
[**PostTransitDatakeyPlaintextName**](Secrets.md#PostTransitDatakeyPlaintextName) | **Post** /{transit_mount_path}/datakey/{plaintext}/{name} | Generate a data key
[**PostTransitDecryptName**](Secrets.md#PostTransitDecryptName) | **Post** /{transit_mount_path}/decrypt/{name} | Decrypt a ciphertext value using a named key
[**PostTransitEncryptName**](Secrets.md#PostTransitEncryptName) | **Post** /{transit_mount_path}/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**PostTransitHash**](Secrets.md#PostTransitHash) | **Post** /{transit_mount_path}/hash | Generate a hash sum for input data
[**PostTransitHashUrlalgorithm**](Secrets.md#PostTransitHashUrlalgorithm) | **Post** /{transit_mount_path}/hash/{urlalgorithm} | Generate a hash sum for input data
[**PostTransitHmacName**](Secrets.md#PostTransitHmacName) | **Post** /{transit_mount_path}/hmac/{name} | Generate an HMAC for input data using the named key
[**PostTransitHmacNameUrlalgorithm**](Secrets.md#PostTransitHmacNameUrlalgorithm) | **Post** /{transit_mount_path}/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**PostTransitKeysName**](Secrets.md#PostTransitKeysName) | **Post** /{transit_mount_path}/keys/{name} | Managed named encryption keys
[**PostTransitKeysNameConfig**](Secrets.md#PostTransitKeysNameConfig) | **Post** /{transit_mount_path}/keys/{name}/config | Configure a named encryption key
[**PostTransitKeysNameImport**](Secrets.md#PostTransitKeysNameImport) | **Post** /{transit_mount_path}/keys/{name}/import | Imports an externally-generated key into a new transit key
[**PostTransitKeysNameImportVersion**](Secrets.md#PostTransitKeysNameImportVersion) | **Post** /{transit_mount_path}/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**PostTransitKeysNameRotate**](Secrets.md#PostTransitKeysNameRotate) | **Post** /{transit_mount_path}/keys/{name}/rotate | Rotate named encryption key
[**PostTransitKeysNameTrim**](Secrets.md#PostTransitKeysNameTrim) | **Post** /{transit_mount_path}/keys/{name}/trim | Trim key versions of a named key
[**PostTransitRandom**](Secrets.md#PostTransitRandom) | **Post** /{transit_mount_path}/random | Generate random bytes
[**PostTransitRandomSource**](Secrets.md#PostTransitRandomSource) | **Post** /{transit_mount_path}/random/{source} | Generate random bytes
[**PostTransitRandomSourceUrlbytes**](Secrets.md#PostTransitRandomSourceUrlbytes) | **Post** /{transit_mount_path}/random/{source}/{urlbytes} | Generate random bytes
[**PostTransitRandomUrlbytes**](Secrets.md#PostTransitRandomUrlbytes) | **Post** /{transit_mount_path}/random/{urlbytes} | Generate random bytes
[**PostTransitRestore**](Secrets.md#PostTransitRestore) | **Post** /{transit_mount_path}/restore | Restore the named key
[**PostTransitRestoreName**](Secrets.md#PostTransitRestoreName) | **Post** /{transit_mount_path}/restore/{name} | Restore the named key
[**PostTransitRewrapName**](Secrets.md#PostTransitRewrapName) | **Post** /{transit_mount_path}/rewrap/{name} | Rewrap ciphertext
[**PostTransitSignName**](Secrets.md#PostTransitSignName) | **Post** /{transit_mount_path}/sign/{name} | Generate a signature for input data using the named key
[**PostTransitSignNameUrlalgorithm**](Secrets.md#PostTransitSignNameUrlalgorithm) | **Post** /{transit_mount_path}/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**PostTransitVerifyName**](Secrets.md#PostTransitVerifyName) | **Post** /{transit_mount_path}/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**PostTransitVerifyNameUrlalgorithm**](Secrets.md#PostTransitVerifyNameUrlalgorithm) | **Post** /{transit_mount_path}/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key



## DeleteAdConfig

> DeleteAdConfig(ctx, adMountPath).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.DeleteAdConfig(context.Background(), adMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAdLibraryName

> DeleteAdLibraryName(ctx, adMountPath, name).Execute()

Delete a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.DeleteAdLibraryName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAdRolesName

> DeleteAdRolesName(ctx, adMountPath, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.DeleteAdRolesName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAlicloudConfig

> DeleteAlicloudConfig(ctx, alicloudMountPath).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Secrets.DeleteAlicloudConfig(context.Background(), alicloudMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAlicloudRoleName

> DeleteAlicloudRoleName(ctx, alicloudMountPath, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Secrets.DeleteAlicloudRoleName(context.Background(), alicloudMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]
**name** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsRolesName

> DeleteAwsRolesName(ctx, awsMountPath, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.DeleteAwsRolesName(context.Background(), awsMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAzureConfig

> DeleteAzureConfig(ctx, azureMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.DeleteAzureConfig(context.Background(), azureMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAzureRolesName

> DeleteAzureRolesName(ctx, azureMountPath, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.DeleteAzureRolesName(context.Background(), azureMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteConsulRolesName

> DeleteConsulRolesName(ctx, consulMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	resp, err := client.WithToken("my-token").Secrets.DeleteConsulRolesName(context.Background(), consulMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteCubbyholePath

> DeleteCubbyholePath(ctx, cubbyholeMountPath, path).Execute()

Deletes the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	cubbyholeMountPath := "cubbyholeMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	resp, err := client.WithToken("my-token").Secrets.DeleteCubbyholePath(context.Background(), cubbyholeMountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cubbyholeMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpRolesetName

> DeleteGcpRolesetName(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpRolesetName(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpStaticAccountName

> DeleteGcpStaticAccountName(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpStaticAccountName(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsConfig

> DeleteGcpkmsConfig(ctx, gcpkmsMountPath).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsConfig(context.Background(), gcpkmsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysDeregisterKey

> DeleteGcpkmsKeysDeregisterKey(ctx, gcpkmsMountPath, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysDeregisterKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysKey

> DeleteGcpkmsKeysKey(ctx, gcpkmsMountPath, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysTrimKey

> DeleteGcpkmsKeysTrimKey(ctx, gcpkmsMountPath, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysTrimKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteKubernetesConfig

> DeleteKubernetesConfig(ctx, kubernetesMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Secrets.DeleteKubernetesConfig(context.Background(), kubernetesMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteKubernetesRolesName

> DeleteKubernetesRolesName(ctx, kubernetesMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Secrets.DeleteKubernetesRolesName(context.Background(), kubernetesMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapConfig

> DeleteLdapConfig(ctx, ldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapConfig(context.Background(), ldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapLibraryName

> DeleteLdapLibraryName(ctx, ldapMountPath, name).Execute()

Delete a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapLibraryName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapRoleName

> DeleteLdapRoleName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapRoleName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapStaticRoleName

> DeleteLdapStaticRoleName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapStaticRoleName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteMongodbatlasRolesName

> DeleteMongodbatlasRolesName(ctx, mongodbatlasMountPath, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	resp, err := client.WithToken("my-token").Secrets.DeleteMongodbatlasRolesName(context.Background(), mongodbatlasMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]
**name** | **string** | Name of the Roles | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadConfigAccess

> DeleteNomadConfigAccess(ctx, nomadMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.DeleteNomadConfigAccess(context.Background(), nomadMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadConfigLease

> DeleteNomadConfigLease(ctx, nomadMountPath).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.DeleteNomadConfigLease(context.Background(), nomadMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadRoleName

> DeleteNomadRoleName(ctx, name, nomadMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.DeleteNomadRoleName(context.Background(), name, nomadMountPath)
	if err != nil {
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
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOpenldapConfig

> DeleteOpenldapConfig(ctx, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteOpenldapConfig(context.Background(), openldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOpenldapLibraryName

> DeleteOpenldapLibraryName(ctx, name, openldapMountPath).Execute()

Delete a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteOpenldapLibraryName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOpenldapRoleName

> DeleteOpenldapRoleName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteOpenldapRoleName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOpenldapStaticRoleName

> DeleteOpenldapStaticRoleName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteOpenldapStaticRoleName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiIssuerRefDerPem

> DeletePkiIssuerRefDerPem(ctx, issuerRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiIssuerRefDerPem(context.Background(), issuerRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiJson

> DeletePkiJson(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiJson(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiKeyKeyRef

> DeletePkiKeyKeyRef(ctx, keyRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiKeyKeyRef(context.Background(), keyRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiRolesName

> DeletePkiRolesName(ctx, name, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiRolesName(context.Background(), name, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiRoot

> DeletePkiRoot(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiRoot(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteRabbitmqRolesName

> DeleteRabbitmqRolesName(ctx, name, rabbitmqMountPath).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	resp, err := client.WithToken("my-token").Secrets.DeleteRabbitmqRolesName(context.Background(), name, rabbitmqMountPath)
	if err != nil {
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
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSecretDataPath

> DeleteSecretDataPath(ctx, path, secretMountPath).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.DeleteSecretDataPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSecretMetadataPath

> DeleteSecretMetadataPath(ctx, path, secretMountPath).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.DeleteSecretMetadataPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSecretPath

> DeleteSecretPath(ctx, path, secretMountPath).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.DeleteSecretPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshConfigCa

> DeleteSshConfigCa(ctx, sshMountPath).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshConfigCa(context.Background(), sshMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshConfigZeroaddress

> DeleteSshConfigZeroaddress(ctx, sshMountPath).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshConfigZeroaddress(context.Background(), sshMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshKeysKeyName

> DeleteSshKeysKeyName(ctx, keyName, sshMountPath).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	keyName := "keyName_example" // string | [Required] Name of the key
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshKeysKeyName(context.Background(), keyName, sshMountPath)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshRolesRole

> DeleteSshRolesRole(ctx, role, sshMountPath).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshRolesRole(context.Background(), role, sshMountPath)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTerraformConfig

> DeleteTerraformConfig(ctx, terraformMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.DeleteTerraformConfig(context.Background(), terraformMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTerraformRoleName

> DeleteTerraformRoleName(ctx, name, terraformMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.DeleteTerraformRoleName(context.Background(), name, terraformMountPath)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTotpKeysName

> DeleteTotpKeysName(ctx, name, totpMountPath).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.DeleteTotpKeysName(context.Background(), name, totpMountPath)
	if err != nil {
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
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTransitKeysName

> DeleteTransitKeysName(ctx, name, transitMountPath).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.DeleteTransitKeysName(context.Background(), name, transitMountPath)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdConfig

> GetAdConfig(ctx, adMountPath).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdConfig(context.Background(), adMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdCredsName

> GetAdCredsName(ctx, adMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdCredsName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdLibrary

> GetAdLibrary(ctx, adMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetAdLibrary(context.Background(), adMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdLibraryName

> GetAdLibraryName(ctx, adMountPath, name).Execute()

Read a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdLibraryName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdLibraryNameStatus

> GetAdLibraryNameStatus(ctx, adMountPath, name).Execute()

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdLibraryNameStatus(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdRoles

> GetAdRoles(ctx, adMountPath).List(list).Execute()

List the name of each role currently stored.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetAdRoles(context.Background(), adMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdRolesName

> GetAdRolesName(ctx, adMountPath, name).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdRolesName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAdRotateRoot

> GetAdRotateRoot(ctx, adMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.GetAdRotateRoot(context.Background(), adMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudConfig

> GetAlicloudConfig(ctx, alicloudMountPath).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Secrets.GetAlicloudConfig(context.Background(), alicloudMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudCredsName

> GetAlicloudCredsName(ctx, alicloudMountPath, name).Execute()

Generate an API key or STS credential using the given role's configuration.'

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Secrets.GetAlicloudCredsName(context.Background(), alicloudMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]
**name** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudRole

> GetAlicloudRole(ctx, alicloudMountPath).List(list).Execute()

List the existing roles in this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetAlicloudRole(context.Background(), alicloudMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAlicloudRoleName

> GetAlicloudRoleName(ctx, alicloudMountPath, name).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Secrets.GetAlicloudRoleName(context.Background(), alicloudMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]
**name** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsConfigLease

> GetAwsConfigLease(ctx, awsMountPath).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.GetAwsConfigLease(context.Background(), awsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsConfigRoot

> GetAwsConfigRoot(ctx, awsMountPath).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.GetAwsConfigRoot(context.Background(), awsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsCreds

> GetAwsCreds(ctx, awsMountPath).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.GetAwsCreds(context.Background(), awsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsRoles

> GetAwsRoles(ctx, awsMountPath).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetAwsRoles(context.Background(), awsMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsRolesName

> GetAwsRolesName(ctx, awsMountPath, name).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.GetAwsRolesName(context.Background(), awsMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAwsStsName

> GetAwsStsName(ctx, awsMountPath, name).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.GetAwsStsName(context.Background(), awsMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureConfig

> GetAzureConfig(ctx, azureMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.GetAzureConfig(context.Background(), azureMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureCredsRole

> GetAzureCredsRole(ctx, azureMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the Vault role
	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.GetAzureCredsRole(context.Background(), azureMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]
**role** | **string** | Name of the Vault role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureRoles

> GetAzureRoles(ctx, azureMountPath).List(list).Execute()

List existing roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetAzureRoles(context.Background(), azureMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAzureRolesName

> GetAzureRolesName(ctx, azureMountPath, name).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.GetAzureRolesName(context.Background(), azureMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulConfigAccess

> GetConsulConfigAccess(ctx, consulMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	resp, err := client.WithToken("my-token").Secrets.GetConsulConfigAccess(context.Background(), consulMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulCredsRole

> GetConsulCredsRole(ctx, consulMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	resp, err := client.WithToken("my-token").Secrets.GetConsulCredsRole(context.Background(), consulMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulRoles

> GetConsulRoles(ctx, consulMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetConsulRoles(context.Background(), consulMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetConsulRolesName

> GetConsulRolesName(ctx, consulMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	resp, err := client.WithToken("my-token").Secrets.GetConsulRolesName(context.Background(), consulMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetCubbyholePath

> GetCubbyholePath(ctx, cubbyholeMountPath, path).List(list).Execute()

Retrieve the secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	cubbyholeMountPath := "cubbyholeMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetCubbyholePath(context.Background(), cubbyholeMountPath, path, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cubbyholeMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | [default to &quot;false&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpConfig

> GetGcpConfig(ctx, gcpMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpConfig(context.Background(), gcpMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpKeyRoleset

> GetGcpKeyRoleset(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpKeyRoleset(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpRolesetName

> GetGcpRolesetName(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpRolesetName(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpRolesetRolesetKey

> GetGcpRolesetRolesetKey(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpRolesetRolesetKey(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpRolesetRolesetToken

> GetGcpRolesetRolesetToken(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpRolesetRolesetToken(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpRolesets

> GetGcpRolesets(ctx, gcpMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetGcpRolesets(context.Background(), gcpMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpStaticAccountName

> GetGcpStaticAccountName(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpStaticAccountName(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpStaticAccountNameKey

> GetGcpStaticAccountNameKey(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpStaticAccountNameKey(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpStaticAccountNameToken

> GetGcpStaticAccountNameToken(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpStaticAccountNameToken(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpStaticAccounts

> GetGcpStaticAccounts(ctx, gcpMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetGcpStaticAccounts(context.Background(), gcpMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpTokenRoleset

> GetGcpTokenRoleset(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.GetGcpTokenRoleset(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsConfig

> GetGcpkmsConfig(ctx, gcpkmsMountPath).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.GetGcpkmsConfig(context.Background(), gcpkmsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsKeys

> GetGcpkmsKeys(ctx, gcpkmsMountPath).List(list).Execute()

List named keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetGcpkmsKeys(context.Background(), gcpkmsMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsKeysConfigKey

> GetGcpkmsKeysConfigKey(ctx, gcpkmsMountPath, key).Execute()

Configure the key in Vault

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.GetGcpkmsKeysConfigKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsKeysKey

> GetGcpkmsKeysKey(ctx, gcpkmsMountPath, key).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.GetGcpkmsKeysKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetGcpkmsPubkeyKey

> GetGcpkmsPubkeyKey(ctx, gcpkmsMountPath, key).Execute()

Retrieve the public key associated with the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.GetGcpkmsPubkeyKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetKubernetesConfig

> GetKubernetesConfig(ctx, kubernetesMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Secrets.GetKubernetesConfig(context.Background(), kubernetesMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetKubernetesRoles

> GetKubernetesRoles(ctx, kubernetesMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetKubernetesRoles(context.Background(), kubernetesMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetKubernetesRolesName

> GetKubernetesRolesName(ctx, kubernetesMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Secrets.GetKubernetesRolesName(context.Background(), kubernetesMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapConfig

> GetLdapConfig(ctx, ldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapConfig(context.Background(), ldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapCredsName

> GetLdapCredsName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the dynamic role.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapCredsName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the dynamic role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapLibrary

> GetLdapLibrary(ctx, ldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetLdapLibrary(context.Background(), ldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapLibraryName

> GetLdapLibraryName(ctx, ldapMountPath, name).Execute()

Read a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapLibraryName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapLibraryNameStatus

> GetLdapLibraryNameStatus(ctx, ldapMountPath, name).Execute()

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapLibraryNameStatus(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapRole

> GetLdapRole(ctx, ldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetLdapRole(context.Background(), ldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapRoleName

> GetLdapRoleName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapRoleName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapStaticCredName

> GetLdapStaticCredName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapStaticCredName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the static role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapStaticRole

> GetLdapStaticRole(ctx, ldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetLdapStaticRole(context.Background(), ldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetLdapStaticRoleName

> GetLdapStaticRoleName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.GetLdapStaticRoleName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasConfig

> GetMongodbatlasConfig(ctx, mongodbatlasMountPath).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	resp, err := client.WithToken("my-token").Secrets.GetMongodbatlasConfig(context.Background(), mongodbatlasMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasCredsName

> GetMongodbatlasCredsName(ctx, mongodbatlasMountPath, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	resp, err := client.WithToken("my-token").Secrets.GetMongodbatlasCredsName(context.Background(), mongodbatlasMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasRoles

> GetMongodbatlasRoles(ctx, mongodbatlasMountPath).List(list).Execute()

List the existing roles in this backend

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetMongodbatlasRoles(context.Background(), mongodbatlasMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetMongodbatlasRolesName

> GetMongodbatlasRolesName(ctx, mongodbatlasMountPath, name).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	resp, err := client.WithToken("my-token").Secrets.GetMongodbatlasRolesName(context.Background(), mongodbatlasMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]
**name** | **string** | Name of the Roles | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadConfigAccess

> GetNomadConfigAccess(ctx, nomadMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.GetNomadConfigAccess(context.Background(), nomadMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadConfigLease

> GetNomadConfigLease(ctx, nomadMountPath).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.GetNomadConfigLease(context.Background(), nomadMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadCredsName

> GetNomadCredsName(ctx, name, nomadMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.GetNomadCredsName(context.Background(), name, nomadMountPath)
	if err != nil {
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
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadRole

> GetNomadRole(ctx, nomadMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetNomadRole(context.Background(), nomadMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetNomadRoleName

> GetNomadRoleName(ctx, name, nomadMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.GetNomadRoleName(context.Background(), name, nomadMountPath)
	if err != nil {
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
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapConfig

> GetOpenldapConfig(ctx, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapConfig(context.Background(), openldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapCredsName

> GetOpenldapCredsName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the dynamic role.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapCredsName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapLibrary

> GetOpenldapLibrary(ctx, openldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetOpenldapLibrary(context.Background(), openldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapLibraryName

> GetOpenldapLibraryName(ctx, name, openldapMountPath).Execute()

Read a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapLibraryName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapLibraryNameStatus

> GetOpenldapLibraryNameStatus(ctx, name, openldapMountPath).Execute()

Check the status of the service accounts in a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapLibraryNameStatus(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapRole

> GetOpenldapRole(ctx, openldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetOpenldapRole(context.Background(), openldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapRoleName

> GetOpenldapRoleName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapRoleName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapStaticCredName

> GetOpenldapStaticCredName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapStaticCredName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapStaticRole

> GetOpenldapStaticRole(ctx, openldapMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetOpenldapStaticRole(context.Background(), openldapMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetOpenldapStaticRoleName

> GetOpenldapStaticRoleName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.GetOpenldapStaticRoleName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCa

> GetPkiCa(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCa(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCaChain

> GetPkiCaChain(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCaChain(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCaPem

> GetPkiCaPem(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCaPem(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertCaChain

> GetPkiCertCaChain(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCertCaChain(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertSerial

> GetPkiCertSerial(ctx, pkiMountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCertSerial(context.Background(), pkiMountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertSerialRaw

> GetPkiCertSerialRaw(ctx, pkiMountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCertSerialRaw(context.Background(), pkiMountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertSerialRawPem

> GetPkiCertSerialRawPem(ctx, pkiMountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	serial := "serial_example" // string | Certificate serial number, in colon- or hyphen-separated octal
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCertSerialRawPem(context.Background(), pkiMountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCerts

> GetPkiCerts(ctx, pkiMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetPkiCerts(context.Background(), pkiMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCertsRevoked

> GetPkiCertsRevoked(ctx, pkiMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetPkiCertsRevoked(context.Background(), pkiMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigAutoTidy

> GetPkiConfigAutoTidy(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiConfigAutoTidy(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigCrl

> GetPkiConfigCrl(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiConfigCrl(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigIssuers

> GetPkiConfigIssuers(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiConfigIssuers(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigKeys

> GetPkiConfigKeys(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiConfigKeys(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiConfigUrls

> GetPkiConfigUrls(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiConfigUrls(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrl

> GetPkiCrl(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCrl(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrlRotate

> GetPkiCrlRotate(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCrlRotate(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiCrlRotateDelta

> GetPkiCrlRotateDelta(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiCrlRotateDelta(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiDelta

> GetPkiDelta(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiDelta(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiDeltaCrl

> GetPkiDeltaCrl(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiDeltaCrl(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiDeltaPem

> GetPkiDeltaPem(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiDeltaPem(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiDer

> GetPkiDer(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiDer(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiIssuerRefCrlPemDerDeltaPem

> GetPkiIssuerRefCrlPemDerDeltaPem(ctx, issuerRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiIssuerRefCrlPemDerDeltaPem(context.Background(), issuerRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiIssuerRefDerPem

> GetPkiIssuerRefDerPem(ctx, issuerRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiIssuerRefDerPem(context.Background(), issuerRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiIssuers

> GetPkiIssuers(ctx, pkiMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetPkiIssuers(context.Background(), pkiMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiJson

> GetPkiJson(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiJson(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiKeyKeyRef

> GetPkiKeyKeyRef(ctx, keyRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiKeyKeyRef(context.Background(), keyRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiKeys

> GetPkiKeys(ctx, pkiMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetPkiKeys(context.Background(), pkiMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiOcspReq

> GetPkiOcspReq(ctx, pkiMountPath, req).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	req := "req_example" // string | base-64 encoded ocsp request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiOcspReq(context.Background(), pkiMountPath, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**req** | **string** | base-64 encoded ocsp request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiPem

> GetPkiPem(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiPem(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiRoles

> GetPkiRoles(ctx, pkiMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetPkiRoles(context.Background(), pkiMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiRolesName

> GetPkiRolesName(ctx, name, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiRolesName(context.Background(), name, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPkiTidyStatus

> GetPkiTidyStatus(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.GetPkiTidyStatus(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqConfigLease

> GetRabbitmqConfigLease(ctx, rabbitmqMountPath).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	resp, err := client.WithToken("my-token").Secrets.GetRabbitmqConfigLease(context.Background(), rabbitmqMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqCredsName

> GetRabbitmqCredsName(ctx, name, rabbitmqMountPath).Execute()

Request RabbitMQ credentials for a certain role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	resp, err := client.WithToken("my-token").Secrets.GetRabbitmqCredsName(context.Background(), name, rabbitmqMountPath)
	if err != nil {
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
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqRoles

> GetRabbitmqRoles(ctx, rabbitmqMountPath).List(list).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetRabbitmqRoles(context.Background(), rabbitmqMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRabbitmqRolesName

> GetRabbitmqRolesName(ctx, name, rabbitmqMountPath).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	resp, err := client.WithToken("my-token").Secrets.GetRabbitmqRolesName(context.Background(), name, rabbitmqMountPath)
	if err != nil {
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
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretConfig

> GetSecretConfig(ctx, secretMountPath).Execute()

Read the backend level settings.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.GetSecretConfig(context.Background(), secretMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretDataPath

> GetSecretDataPath(ctx, path, secretMountPath).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.GetSecretDataPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretMetadataPath

> GetSecretMetadataPath(ctx, path, secretMountPath).List(list).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetSecretMetadataPath(context.Background(), path, secretMountPath, list)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | [default to &quot;false&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretPath

> GetSecretPath(ctx, path, secretMountPath).List(list).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetSecretPath(context.Background(), path, secretMountPath, list)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | [default to &quot;false&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSecretSubkeysPath

> GetSecretSubkeysPath(ctx, path, secretMountPath).Execute()

Read the structure of a secret entry from the Key-Value store with the values removed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.GetSecretSubkeysPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshConfigCa

> GetSshConfigCa(ctx, sshMountPath).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.GetSshConfigCa(context.Background(), sshMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshConfigZeroaddress

> GetSshConfigZeroaddress(ctx, sshMountPath).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.GetSshConfigZeroaddress(context.Background(), sshMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshPublicKey

> GetSshPublicKey(ctx, sshMountPath).Execute()

Retrieve the public key.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.GetSshPublicKey(context.Background(), sshMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshRoles

> GetSshRoles(ctx, sshMountPath).List(list).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetSshRoles(context.Background(), sshMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetSshRolesRole

> GetSshRolesRole(ctx, role, sshMountPath).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.GetSshRolesRole(context.Background(), role, sshMountPath)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformConfig

> GetTerraformConfig(ctx, terraformMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.GetTerraformConfig(context.Background(), terraformMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformCredsName

> GetTerraformCredsName(ctx, name, terraformMountPath).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.GetTerraformCredsName(context.Background(), name, terraformMountPath)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformRole

> GetTerraformRole(ctx, terraformMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetTerraformRole(context.Background(), terraformMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTerraformRoleName

> GetTerraformRoleName(ctx, name, terraformMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.GetTerraformRoleName(context.Background(), name, terraformMountPath)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTotpCodeName

> GetTotpCodeName(ctx, name, totpMountPath).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.GetTotpCodeName(context.Background(), name, totpMountPath)
	if err != nil {
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
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTotpKeys

> GetTotpKeys(ctx, totpMountPath).List(list).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetTotpKeys(context.Background(), totpMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTotpKeysName

> GetTotpKeysName(ctx, name, totpMountPath).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.GetTotpKeysName(context.Background(), name, totpMountPath)
	if err != nil {
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
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitBackupName

> GetTransitBackupName(ctx, name, transitMountPath).Execute()

Backup the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitBackupName(context.Background(), name, transitMountPath)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitCacheConfig

> GetTransitCacheConfig(ctx, transitMountPath).Execute()

Returns the size of the active cache

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitCacheConfig(context.Background(), transitMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitExportTypeName

> GetTransitExportTypeName(ctx, name, transitMountPath, type_).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitExportTypeName(context.Background(), name, transitMountPath, type_)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitExportTypeNameVersion

> GetTransitExportTypeNameVersion(ctx, name, transitMountPath, type_, version).Execute()

Export named encryption or signing key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	type_ := "type__example" // string | Type of key to export (encryption-key, signing-key, hmac-key)
	version := "version_example" // string | Version of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitExportTypeNameVersion(context.Background(), name, transitMountPath, type_, version)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 
**version** | **string** | Version of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitKeys

> GetTransitKeys(ctx, transitMountPath).List(list).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.GetTransitKeys(context.Background(), transitMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitKeysName

> GetTransitKeysName(ctx, name, transitMountPath).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitKeysName(context.Background(), name, transitMountPath)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetTransitWrappingKey

> GetTransitWrappingKey(ctx, transitMountPath).Execute()

Returns the public key to use for wrapping imported keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.GetTransitWrappingKey(context.Background(), transitMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdConfig

> PostAdConfig(ctx, adMountPath).AdConfigRequest(adConfigRequest).Execute()

Configure the AD server to connect to, along with password options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adConfigRequest := NewAdConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdConfig(context.Background(), adMountPath, adConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adConfigRequest** | [**AdConfigRequest**](AdConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryManageNameCheckIn

> PostAdLibraryManageNameCheckIn(ctx, adMountPath, name).AdLibraryManageCheckInRequest(adLibraryManageCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adLibraryManageCheckInRequest := NewAdLibraryManageCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdLibraryManageNameCheckIn(context.Background(), adMountPath, name, adLibraryManageCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryManageCheckInRequest** | [**AdLibraryManageCheckInRequest**](AdLibraryManageCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryName

> PostAdLibraryName(ctx, adMountPath, name).AdLibraryRequest(adLibraryRequest).Execute()

Update a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adLibraryRequest := NewAdLibraryRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdLibraryName(context.Background(), adMountPath, name, adLibraryRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryRequest** | [**AdLibraryRequest**](AdLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckIn

> PostAdLibraryNameCheckIn(ctx, adMountPath, name).AdLibraryCheckInRequest(adLibraryCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adLibraryCheckInRequest := NewAdLibraryCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdLibraryNameCheckIn(context.Background(), adMountPath, name, adLibraryCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryCheckInRequest** | [**AdLibraryCheckInRequest**](AdLibraryCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdLibraryNameCheckOut

> PostAdLibraryNameCheckOut(ctx, adMountPath, name).AdLibraryCheckOutRequest(adLibraryCheckOutRequest).Execute()

Check a service account out from the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adLibraryCheckOutRequest := NewAdLibraryCheckOutRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdLibraryNameCheckOut(context.Background(), adMountPath, name, adLibraryCheckOutRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the set | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adLibraryCheckOutRequest** | [**AdLibraryCheckOutRequest**](AdLibraryCheckOutRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdRolesName

> PostAdRolesName(ctx, adMountPath, name).AdRolesRequest(adRolesRequest).Execute()

Manage roles to build links between Vault and Active Directory service accounts.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	adRolesRequest := NewAdRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAdRolesName(context.Background(), adMountPath, name, adRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adRolesRequest** | [**AdRolesRequest**](AdRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdRotateRoleName

> PostAdRotateRoleName(ctx, adMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role
	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.PostAdRotateRoleName(context.Background(), adMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]
**name** | **string** | Name of the static role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAdRotateRoot

> PostAdRotateRoot(ctx, adMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	adMountPath := "adMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ad")

	resp, err := client.WithToken("my-token").Secrets.PostAdRotateRoot(context.Background(), adMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**adMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAlicloudConfig

> PostAlicloudConfig(ctx, alicloudMountPath).AlicloudConfigRequest(alicloudConfigRequest).Execute()

Configure the access key and secret to use for RAM and STS calls.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	alicloudConfigRequest := NewAlicloudConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAlicloudConfig(context.Background(), alicloudMountPath, alicloudConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alicloudConfigRequest** | [**AlicloudConfigRequest**](AlicloudConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAlicloudRoleName

> PostAlicloudRoleName(ctx, alicloudMountPath, name).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Read, write and reference policies and roles that API keys or STS credentials can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the role.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	alicloudRoleRequest := NewAlicloudRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAlicloudRoleName(context.Background(), alicloudMountPath, name, alicloudRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**alicloudMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;alicloud&quot;]
**name** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alicloudRoleRequest** | [**AlicloudRoleRequest**](AlicloudRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigLease

> PostAwsConfigLease(ctx, awsMountPath).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()

Configure the default lease information for generated credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigLeaseRequest := NewAwsConfigLeaseRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAwsConfigLease(context.Background(), awsMountPath, awsConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigLeaseRequest** | [**AwsConfigLeaseRequest**](AwsConfigLeaseRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigRoot

> PostAwsConfigRoot(ctx, awsMountPath).AwsConfigRootRequest(awsConfigRootRequest).Execute()

Configure the root credentials that are used to manage IAM.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigRootRequest := NewAwsConfigRootRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAwsConfigRoot(context.Background(), awsMountPath, awsConfigRootRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigRootRequest** | [**AwsConfigRootRequest**](AwsConfigRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsConfigRotateRoot

> PostAwsConfigRotateRoot(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.PostAwsConfigRotateRoot(context.Background(), awsMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsCreds

> PostAwsCreds(ctx, awsMountPath).AwsCredsRequest(awsCredsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsCredsRequest := NewAwsCredsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAwsCreds(context.Background(), awsMountPath, awsCredsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredsRequest** | [**AwsCredsRequest**](AwsCredsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsRolesName

> PostAwsRolesName(ctx, awsMountPath, name).AwsRolesRequest(awsRolesRequest).Execute()

Read, write and reference IAM policies that access keys can be made for.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the policy
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRolesRequest := NewAwsRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAwsRolesName(context.Background(), awsMountPath, name, awsRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRolesRequest** | [**AwsRolesRequest**](AwsRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAwsStsName

> PostAwsStsName(ctx, awsMountPath, name).AwsStsRequest(awsStsRequest).Execute()

Generate AWS credentials from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsStsRequest := NewAwsStsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAwsStsName(context.Background(), awsMountPath, name, awsStsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsStsRequest** | [**AwsStsRequest**](AwsStsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureConfig

> PostAzureConfig(ctx, azureMountPath).AzureConfigRequest(azureConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	azureConfigRequest := NewAzureConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAzureConfig(context.Background(), azureMountPath, azureConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureConfigRequest** | [**AzureConfigRequest**](AzureConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureRolesName

> PostAzureRolesName(ctx, azureMountPath, name).AzureRolesRequest(azureRolesRequest).Execute()

Manage the Vault roles used to generate Azure credentials.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	azureRolesRequest := NewAzureRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostAzureRolesName(context.Background(), azureMountPath, name, azureRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureRolesRequest** | [**AzureRolesRequest**](AzureRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAzureRotateRoot

> PostAzureRotateRoot(ctx, azureMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	azureMountPath := "azureMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.PostAzureRotateRoot(context.Background(), azureMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**azureMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostConsulConfigAccess

> PostConsulConfigAccess(ctx, consulMountPath).ConsulConfigAccessRequest(consulConfigAccessRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	consulConfigAccessRequest := NewConsulConfigAccessRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostConsulConfigAccess(context.Background(), consulMountPath, consulConfigAccessRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consulConfigAccessRequest** | [**ConsulConfigAccessRequest**](ConsulConfigAccessRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostConsulRolesName

> PostConsulRolesName(ctx, consulMountPath, name).ConsulRolesRequest(consulRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	consulMountPath := "consulMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "consul")

	consulRolesRequest := NewConsulRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostConsulRolesName(context.Background(), consulMountPath, name, consulRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**consulMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;consul&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consulRolesRequest** | [**ConsulRolesRequest**](ConsulRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostCubbyholePath

> PostCubbyholePath(ctx, cubbyholeMountPath, path).Execute()

Store a secret at the specified location.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Specifies the path of the secret.
	cubbyholeMountPath := "cubbyholeMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	resp, err := client.WithToken("my-token").Secrets.PostCubbyholePath(context.Background(), cubbyholeMountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cubbyholeMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpConfig

> PostGcpConfig(ctx, gcpMountPath).GcpConfigRequest(gcpConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpConfigRequest := NewGcpConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpConfig(context.Background(), gcpMountPath, gcpConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpConfigRequest** | [**GcpConfigRequest**](GcpConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpConfigRotateRoot

> PostGcpConfigRotateRoot(ctx, gcpMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpConfigRotateRoot(context.Background(), gcpMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpKeyRoleset

> PostGcpKeyRoleset(ctx, gcpMountPath, roleset).GcpKeyRequest(gcpKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpKeyRequest := NewGcpKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpKeyRoleset(context.Background(), gcpMountPath, roleset, gcpKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpKeyRequest** | [**GcpKeyRequest**](GcpKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetName

> PostGcpRolesetName(ctx, gcpMountPath, name).GcpRolesetRequest(gcpRolesetRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRolesetRequest := NewGcpRolesetRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpRolesetName(context.Background(), gcpMountPath, name, gcpRolesetRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRolesetRequest** | [**GcpRolesetRequest**](GcpRolesetRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetNameRotate

> PostGcpRolesetNameRotate(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpRolesetNameRotate(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetNameRotateKey

> PostGcpRolesetNameRotateKey(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpRolesetNameRotateKey(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetRolesetKey

> PostGcpRolesetRolesetKey(ctx, gcpMountPath, roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRolesetKeyRequest := NewGcpRolesetKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpRolesetRolesetKey(context.Background(), gcpMountPath, roleset, gcpRolesetKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRolesetKeyRequest** | [**GcpRolesetKeyRequest**](GcpRolesetKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpRolesetRolesetToken

> PostGcpRolesetRolesetToken(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpRolesetRolesetToken(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountName

> PostGcpStaticAccountName(ctx, gcpMountPath, name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name to refer to this static account in Vault. Cannot be updated.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpStaticAccountRequest := NewGcpStaticAccountRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpStaticAccountName(context.Background(), gcpMountPath, name, gcpStaticAccountRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpStaticAccountRequest** | [**GcpStaticAccountRequest**](GcpStaticAccountRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameKey

> PostGcpStaticAccountNameKey(ctx, gcpMountPath, name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpStaticAccountKeyRequest := NewGcpStaticAccountKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpStaticAccountNameKey(context.Background(), gcpMountPath, name, gcpStaticAccountKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpStaticAccountKeyRequest** | [**GcpStaticAccountKeyRequest**](GcpStaticAccountKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameRotateKey

> PostGcpStaticAccountNameRotateKey(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the account.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpStaticAccountNameRotateKey(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpStaticAccountNameToken

> PostGcpStaticAccountNameToken(ctx, gcpMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Required. Name of the static account.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpStaticAccountNameToken(context.Background(), gcpMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpTokenRoleset

> PostGcpTokenRoleset(ctx, gcpMountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleset := "roleset_example" // string | Required. Name of the role set.
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.PostGcpTokenRoleset(context.Background(), gcpMountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsConfig

> PostGcpkmsConfig(ctx, gcpkmsMountPath).GcpkmsConfigRequest(gcpkmsConfigRequest).Execute()

Configure the GCP KMS secrets engine

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsConfigRequest := NewGcpkmsConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsConfig(context.Background(), gcpkmsMountPath, gcpkmsConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpkmsConfigRequest** | [**GcpkmsConfigRequest**](GcpkmsConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsDecryptKey

> PostGcpkmsDecryptKey(ctx, gcpkmsMountPath, key).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsDecryptRequest := NewGcpkmsDecryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsDecryptKey(context.Background(), gcpkmsMountPath, key, gcpkmsDecryptRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsDecryptRequest** | [**GcpkmsDecryptRequest**](GcpkmsDecryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsEncryptKey

> PostGcpkmsEncryptKey(ctx, gcpkmsMountPath, key).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()

Encrypt a plaintext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsEncryptRequest := NewGcpkmsEncryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsEncryptKey(context.Background(), gcpkmsMountPath, key, gcpkmsEncryptRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsEncryptRequest** | [**GcpkmsEncryptRequest**](GcpkmsEncryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysConfigKey

> PostGcpkmsKeysConfigKey(ctx, gcpkmsMountPath, key).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()

Configure the key in Vault

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysConfigRequest := NewGcpkmsKeysConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysConfigKey(context.Background(), gcpkmsMountPath, key, gcpkmsKeysConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysConfigRequest** | [**GcpkmsKeysConfigRequest**](GcpkmsKeysConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysDeregisterKey

> PostGcpkmsKeysDeregisterKey(ctx, gcpkmsMountPath, key).Execute()

Deregister an existing key in Vault

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysDeregisterKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysKey

> PostGcpkmsKeysKey(ctx, gcpkmsMountPath, key).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()

Interact with crypto keys in Vault and Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysRequest := NewGcpkmsKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysKey(context.Background(), gcpkmsMountPath, key, gcpkmsKeysRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysRequest** | [**GcpkmsKeysRequest**](GcpkmsKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysRegisterKey

> PostGcpkmsKeysRegisterKey(ctx, gcpkmsMountPath, key).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()

Register an existing crypto key in Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysRegisterRequest := NewGcpkmsKeysRegisterRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysRegisterKey(context.Background(), gcpkmsMountPath, key, gcpkmsKeysRegisterRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsKeysRegisterRequest** | [**GcpkmsKeysRegisterRequest**](GcpkmsKeysRegisterRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysRotateKey

> PostGcpkmsKeysRotateKey(ctx, gcpkmsMountPath, key).Execute()

Rotate a crypto key to a new primary version

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysRotateKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsKeysTrimKey

> PostGcpkmsKeysTrimKey(ctx, gcpkmsMountPath, key).Execute()

Delete old crypto key versions from Google Cloud KMS

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsKeysTrimKey(context.Background(), gcpkmsMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsReencryptKey

> PostGcpkmsReencryptKey(ctx, gcpkmsMountPath, key).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()

Re-encrypt existing ciphertext data to a new version

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsReencryptRequest := NewGcpkmsReencryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsReencryptKey(context.Background(), gcpkmsMountPath, key, gcpkmsReencryptRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsReencryptRequest** | [**GcpkmsReencryptRequest**](GcpkmsReencryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsSignKey

> PostGcpkmsSignKey(ctx, gcpkmsMountPath, key).GcpkmsSignRequest(gcpkmsSignRequest).Execute()

Signs a message or digest using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsSignRequest := NewGcpkmsSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsSignKey(context.Background(), gcpkmsMountPath, key, gcpkmsSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsSignRequest** | [**GcpkmsSignRequest**](GcpkmsSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostGcpkmsVerifyKey

> PostGcpkmsVerifyKey(ctx, gcpkmsMountPath, key).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()

Verify a signature using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
	gcpkmsMountPath := "gcpkmsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsVerifyRequest := NewGcpkmsVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostGcpkmsVerifyKey(context.Background(), gcpkmsMountPath, key, gcpkmsVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**gcpkmsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]
**key** | **string** | Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpkmsVerifyRequest** | [**GcpkmsVerifyRequest**](GcpkmsVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesConfig

> PostKubernetesConfig(ctx, kubernetesMountPath).KubernetesConfigRequest(kubernetesConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	kubernetesConfigRequest := NewKubernetesConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostKubernetesConfig(context.Background(), kubernetesMountPath, kubernetesConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesConfigRequest** | [**KubernetesConfigRequest**](KubernetesConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesCredsName

> PostKubernetesCredsName(ctx, kubernetesMountPath, name).KubernetesCredsRequest(kubernetesCredsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Vault role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	kubernetesCredsRequest := NewKubernetesCredsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostKubernetesCredsName(context.Background(), kubernetesMountPath, name, kubernetesCredsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]
**name** | **string** | Name of the Vault role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredsRequest** | [**KubernetesCredsRequest**](KubernetesCredsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostKubernetesRolesName

> PostKubernetesRolesName(ctx, kubernetesMountPath, name).KubernetesRolesRequest(kubernetesRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	kubernetesRolesRequest := NewKubernetesRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostKubernetesRolesName(context.Background(), kubernetesMountPath, name, kubernetesRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kubernetesMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kubernetes&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesRolesRequest** | [**KubernetesRolesRequest**](KubernetesRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapConfig

> PostLdapConfig(ctx, ldapMountPath).LdapConfigRequest(ldapConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapConfigRequest := NewLdapConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapConfig(context.Background(), ldapMountPath, ldapConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigRequest** | [**LdapConfigRequest**](LdapConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapLibraryManageNameCheckIn

> PostLdapLibraryManageNameCheckIn(ctx, ldapMountPath, name).LdapLibraryManageCheckInRequest(ldapLibraryManageCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryManageCheckInRequest := NewLdapLibraryManageCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapLibraryManageNameCheckIn(context.Background(), ldapMountPath, name, ldapLibraryManageCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLibraryManageCheckInRequest** | [**LdapLibraryManageCheckInRequest**](LdapLibraryManageCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapLibraryName

> PostLdapLibraryName(ctx, ldapMountPath, name).LdapLibraryRequest(ldapLibraryRequest).Execute()

Update a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryRequest := NewLdapLibraryRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapLibraryName(context.Background(), ldapMountPath, name, ldapLibraryRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLibraryRequest** | [**LdapLibraryRequest**](LdapLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapLibraryNameCheckIn

> PostLdapLibraryNameCheckIn(ctx, ldapMountPath, name).LdapLibraryCheckInRequest(ldapLibraryCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryCheckInRequest := NewLdapLibraryCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapLibraryNameCheckIn(context.Background(), ldapMountPath, name, ldapLibraryCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLibraryCheckInRequest** | [**LdapLibraryCheckInRequest**](LdapLibraryCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapLibraryNameCheckOut

> PostLdapLibraryNameCheckOut(ctx, ldapMountPath, name).LdapLibraryCheckOutRequest(ldapLibraryCheckOutRequest).Execute()

Check a service account out from the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryCheckOutRequest := NewLdapLibraryCheckOutRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapLibraryNameCheckOut(context.Background(), ldapMountPath, name, ldapLibraryCheckOutRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLibraryCheckOutRequest** | [**LdapLibraryCheckOutRequest**](LdapLibraryCheckOutRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapRoleName

> PostLdapRoleName(ctx, ldapMountPath, name).LdapRoleRequest(ldapRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapRoleRequest := NewLdapRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapRoleName(context.Background(), ldapMountPath, name, ldapRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapRoleRequest** | [**LdapRoleRequest**](LdapRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapRotateRoleName

> PostLdapRotateRoleName(ctx, ldapMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.PostLdapRotateRoleName(context.Background(), ldapMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the static role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapRotateRoot

> PostLdapRotateRoot(ctx, ldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.PostLdapRotateRoot(context.Background(), ldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostLdapStaticRoleName

> PostLdapStaticRoleName(ctx, ldapMountPath, name).LdapStaticRoleRequest(ldapStaticRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapStaticRoleRequest := NewLdapStaticRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostLdapStaticRoleName(context.Background(), ldapMountPath, name, ldapStaticRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapStaticRoleRequest** | [**LdapStaticRoleRequest**](LdapStaticRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostMongodbatlasConfig

> PostMongodbatlasConfig(ctx, mongodbatlasMountPath).MongodbatlasConfigRequest(mongodbatlasConfigRequest).Execute()

Configure the  credentials that are used to manage Database Users.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	mongodbatlasConfigRequest := NewMongodbatlasConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostMongodbatlasConfig(context.Background(), mongodbatlasMountPath, mongodbatlasConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mongodbatlasConfigRequest** | [**MongodbatlasConfigRequest**](MongodbatlasConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostMongodbatlasCredsName

> PostMongodbatlasCredsName(ctx, mongodbatlasMountPath, name).Execute()

Generate MongoDB Atlas Programmatic API from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	resp, err := client.WithToken("my-token").Secrets.PostMongodbatlasCredsName(context.Background(), mongodbatlasMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostMongodbatlasRolesName

> PostMongodbatlasRolesName(ctx, mongodbatlasMountPath, name).MongodbatlasRolesRequest(mongodbatlasRolesRequest).Execute()

Manage the roles used to generate MongoDB Atlas Programmatic API Keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Roles
	mongodbatlasMountPath := "mongodbatlasMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "mongodbatlas")

	mongodbatlasRolesRequest := NewMongodbatlasRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostMongodbatlasRolesName(context.Background(), mongodbatlasMountPath, name, mongodbatlasRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mongodbatlasMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;mongodbatlas&quot;]
**name** | **string** | Name of the Roles | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mongodbatlasRolesRequest** | [**MongodbatlasRolesRequest**](MongodbatlasRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadConfigAccess

> PostNomadConfigAccess(ctx, nomadMountPath).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	nomadConfigAccessRequest := NewNomadConfigAccessRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostNomadConfigAccess(context.Background(), nomadMountPath, nomadConfigAccessRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigAccessRequest** | [**NomadConfigAccessRequest**](NomadConfigAccessRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadConfigLease

> PostNomadConfigLease(ctx, nomadMountPath).NomadConfigLeaseRequest(nomadConfigLeaseRequest).Execute()

Configure the lease parameters for generated tokens

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	nomadConfigLeaseRequest := NewNomadConfigLeaseRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostNomadConfigLease(context.Background(), nomadMountPath, nomadConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nomadConfigLeaseRequest** | [**NomadConfigLeaseRequest**](NomadConfigLeaseRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostNomadRoleName

> PostNomadRoleName(ctx, name, nomadMountPath).NomadRoleRequest(nomadRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	nomadMountPath := "nomadMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	nomadRoleRequest := NewNomadRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostNomadRoleName(context.Background(), name, nomadMountPath, nomadRoleRequest)
	if err != nil {
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
**nomadMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadRoleRequest** | [**NomadRoleRequest**](NomadRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapConfig

> PostOpenldapConfig(ctx, openldapMountPath).OpenldapConfigRequest(openldapConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapConfigRequest := NewOpenldapConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapConfig(context.Background(), openldapMountPath, openldapConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openldapConfigRequest** | [**OpenldapConfigRequest**](OpenldapConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapLibraryManageNameCheckIn

> PostOpenldapLibraryManageNameCheckIn(ctx, name, openldapMountPath).OpenldapLibraryManageCheckInRequest(openldapLibraryManageCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapLibraryManageCheckInRequest := NewOpenldapLibraryManageCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapLibraryManageNameCheckIn(context.Background(), name, openldapMountPath, openldapLibraryManageCheckInRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapLibraryManageCheckInRequest** | [**OpenldapLibraryManageCheckInRequest**](OpenldapLibraryManageCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapLibraryName

> PostOpenldapLibraryName(ctx, name, openldapMountPath).OpenldapLibraryRequest(openldapLibraryRequest).Execute()

Update a library set.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapLibraryRequest := NewOpenldapLibraryRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapLibraryName(context.Background(), name, openldapMountPath, openldapLibraryRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapLibraryRequest** | [**OpenldapLibraryRequest**](OpenldapLibraryRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapLibraryNameCheckIn

> PostOpenldapLibraryNameCheckIn(ctx, name, openldapMountPath).OpenldapLibraryCheckInRequest(openldapLibraryCheckInRequest).Execute()

Check service accounts in to the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set.
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapLibraryCheckInRequest := NewOpenldapLibraryCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapLibraryNameCheckIn(context.Background(), name, openldapMountPath, openldapLibraryCheckInRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapLibraryCheckInRequest** | [**OpenldapLibraryCheckInRequest**](OpenldapLibraryCheckInRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapLibraryNameCheckOut

> PostOpenldapLibraryNameCheckOut(ctx, name, openldapMountPath).OpenldapLibraryCheckOutRequest(openldapLibraryCheckOutRequest).Execute()

Check a service account out from the library.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the set
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapLibraryCheckOutRequest := NewOpenldapLibraryCheckOutRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapLibraryNameCheckOut(context.Background(), name, openldapMountPath, openldapLibraryCheckOutRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapLibraryCheckOutRequest** | [**OpenldapLibraryCheckOutRequest**](OpenldapLibraryCheckOutRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapRoleName

> PostOpenldapRoleName(ctx, name, openldapMountPath).OpenldapRoleRequest(openldapRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role (lowercase)
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapRoleRequest := NewOpenldapRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapRoleName(context.Background(), name, openldapMountPath, openldapRoleRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapRoleRequest** | [**OpenldapRoleRequest**](OpenldapRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapRotateRoleName

> PostOpenldapRotateRoleName(ctx, name, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the static role
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.PostOpenldapRotateRoleName(context.Background(), name, openldapMountPath)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapRotateRoot

> PostOpenldapRotateRoot(ctx, openldapMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	resp, err := client.WithToken("my-token").Secrets.PostOpenldapRotateRoot(context.Background(), openldapMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostOpenldapStaticRoleName

> PostOpenldapStaticRoleName(ctx, name, openldapMountPath).OpenldapStaticRoleRequest(openldapStaticRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	openldapMountPath := "openldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "openldap")

	openldapStaticRoleRequest := NewOpenldapStaticRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostOpenldapStaticRoleName(context.Background(), name, openldapMountPath, openldapStaticRoleRequest)
	if err != nil {
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
**openldapMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;openldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openldapStaticRoleRequest** | [**OpenldapStaticRoleRequest**](OpenldapStaticRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiBundle

> PostPkiBundle(ctx, pkiMountPath).PkiBundleRequest(pkiBundleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiBundleRequest := NewPkiBundleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiBundle(context.Background(), pkiMountPath, pkiBundleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiBundleRequest** | [**PkiBundleRequest**](PkiBundleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiCert

> PostPkiCert(ctx, pkiMountPath).PkiCertRequest(pkiCertRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiCertRequest := NewPkiCertRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiCert(context.Background(), pkiMountPath, pkiCertRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiCertRequest** | [**PkiCertRequest**](PkiCertRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigAutoTidy

> PostPkiConfigAutoTidy(ctx, pkiMountPath).PkiConfigAutoTidyRequest(pkiConfigAutoTidyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigAutoTidyRequest := NewPkiConfigAutoTidyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigAutoTidy(context.Background(), pkiMountPath, pkiConfigAutoTidyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigAutoTidyRequest** | [**PkiConfigAutoTidyRequest**](PkiConfigAutoTidyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigCa

> PostPkiConfigCa(ctx, pkiMountPath).PkiConfigCaRequest(pkiConfigCaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigCaRequest := NewPkiConfigCaRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigCa(context.Background(), pkiMountPath, pkiConfigCaRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCaRequest** | [**PkiConfigCaRequest**](PkiConfigCaRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigCrl

> PostPkiConfigCrl(ctx, pkiMountPath).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigCrlRequest := NewPkiConfigCrlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigCrl(context.Background(), pkiMountPath, pkiConfigCrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigCrlRequest** | [**PkiConfigCrlRequest**](PkiConfigCrlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigIssuers

> PostPkiConfigIssuers(ctx, pkiMountPath).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigIssuersRequest := NewPkiConfigIssuersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigIssuers(context.Background(), pkiMountPath, pkiConfigIssuersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigIssuersRequest** | [**PkiConfigIssuersRequest**](PkiConfigIssuersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigKeys

> PostPkiConfigKeys(ctx, pkiMountPath).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigKeysRequest := NewPkiConfigKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigKeys(context.Background(), pkiMountPath, pkiConfigKeysRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigKeysRequest** | [**PkiConfigKeysRequest**](PkiConfigKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiConfigUrls

> PostPkiConfigUrls(ctx, pkiMountPath).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigUrlsRequest := NewPkiConfigUrlsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiConfigUrls(context.Background(), pkiMountPath, pkiConfigUrlsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConfigUrlsRequest** | [**PkiConfigUrlsRequest**](PkiConfigUrlsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateCrossSign

> PostPkiIntermediateCrossSign(ctx, pkiMountPath).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateCrossSignRequest := NewPkiIntermediateCrossSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIntermediateCrossSign(context.Background(), pkiMountPath, pkiIntermediateCrossSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateCrossSignRequest** | [**PkiIntermediateCrossSignRequest**](PkiIntermediateCrossSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateGenerateExported

> PostPkiIntermediateGenerateExported(ctx, exported, pkiMountPath).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateGenerateRequest := NewPkiIntermediateGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIntermediateGenerateExported(context.Background(), exported, pkiMountPath, pkiIntermediateGenerateRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIntermediateGenerateRequest** | [**PkiIntermediateGenerateRequest**](PkiIntermediateGenerateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIntermediateSetSigned

> PostPkiIntermediateSetSigned(ctx, pkiMountPath).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateSetSignedRequest := NewPkiIntermediateSetSignedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIntermediateSetSigned(context.Background(), pkiMountPath, pkiIntermediateSetSignedRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiIntermediateSetSignedRequest** | [**PkiIntermediateSetSignedRequest**](PkiIntermediateSetSignedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiInternalExported

> PostPkiInternalExported(ctx, pkiMountPath).PkiInternalExportedRequest(pkiInternalExportedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiInternalExportedRequest := NewPkiInternalExportedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiInternalExported(context.Background(), pkiMountPath, pkiInternalExportedRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiInternalExportedRequest** | [**PkiInternalExportedRequest**](PkiInternalExportedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssueRole

> PostPkiIssueRole(ctx, pkiMountPath, role).PkiIssueRequest(pkiIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssueRequest := NewPkiIssueRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssueRole(context.Background(), pkiMountPath, role, pkiIssueRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssueRequest** | [**PkiIssueRequest**](PkiIssueRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefIssueRole

> PostPkiIssuerIssuerRefIssueRole(ctx, issuerRef, pkiMountPath, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerIssueRequest := NewPkiIssuerIssueRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefIssueRole(context.Background(), issuerRef, pkiMountPath, role, pkiIssuerIssueRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerIssueRequest** | [**PkiIssuerIssueRequest**](PkiIssuerIssueRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefRevoke

> PostPkiIssuerIssuerRefRevoke(ctx, issuerRef, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefRevoke(context.Background(), issuerRef, pkiMountPath)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignIntermediate

> PostPkiIssuerIssuerRefSignIntermediate(ctx, issuerRef, pkiMountPath).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignIntermediateRequest := NewPkiIssuerSignIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefSignIntermediate(context.Background(), issuerRef, pkiMountPath, pkiIssuerSignIntermediateRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignIntermediateRequest** | [**PkiIssuerSignIntermediateRequest**](PkiIssuerSignIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignRole

> PostPkiIssuerIssuerRefSignRole(ctx, issuerRef, pkiMountPath, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignRequest := NewPkiIssuerSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefSignRole(context.Background(), issuerRef, pkiMountPath, role, pkiIssuerSignRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignRequest** | [**PkiIssuerSignRequest**](PkiIssuerSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignSelfIssued

> PostPkiIssuerIssuerRefSignSelfIssued(ctx, issuerRef, pkiMountPath).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignSelfIssuedRequest := NewPkiIssuerSignSelfIssuedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefSignSelfIssued(context.Background(), issuerRef, pkiMountPath, pkiIssuerSignSelfIssuedRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignSelfIssuedRequest** | [**PkiIssuerSignSelfIssuedRequest**](PkiIssuerSignSelfIssuedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatim

> PostPkiIssuerIssuerRefSignVerbatim(ctx, issuerRef, pkiMountPath).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefSignVerbatim(context.Background(), issuerRef, pkiMountPath, pkiIssuerSignVerbatimRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerIssuerRefSignVerbatimRole

> PostPkiIssuerIssuerRefSignVerbatimRole(ctx, issuerRef, pkiMountPath, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerIssuerRefSignVerbatimRole(context.Background(), issuerRef, pkiMountPath, role, pkiIssuerSignVerbatimRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuerRefDerPem

> PostPkiIssuerRefDerPem(ctx, issuerRef, pkiMountPath).PkiDerPemRequest(pkiDerPemRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	issuerRef := "issuerRef_example" // string | Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiDerPemRequest := NewPkiDerPemRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuerRefDerPem(context.Background(), issuerRef, pkiMountPath, pkiDerPemRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiDerPemRequest** | [**PkiDerPemRequest**](PkiDerPemRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateIntermediateExported

> PostPkiIssuersGenerateIntermediateExported(ctx, exported, pkiMountPath).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuersGenerateIntermediateRequest := NewPkiIssuersGenerateIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuersGenerateIntermediateExported(context.Background(), exported, pkiMountPath, pkiIssuersGenerateIntermediateRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersGenerateIntermediateRequest** | [**PkiIssuersGenerateIntermediateRequest**](PkiIssuersGenerateIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiIssuersGenerateRootExported

> PostPkiIssuersGenerateRootExported(ctx, exported, pkiMountPath).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuersGenerateRootRequest := NewPkiIssuersGenerateRootRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiIssuersGenerateRootExported(context.Background(), exported, pkiMountPath, pkiIssuersGenerateRootRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIssuersGenerateRootRequest** | [**PkiIssuersGenerateRootRequest**](PkiIssuersGenerateRootRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiJson

> PostPkiJson(ctx, pkiMountPath).PkiJsonRequest(pkiJsonRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiJsonRequest := NewPkiJsonRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiJson(context.Background(), pkiMountPath, pkiJsonRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiJsonRequest** | [**PkiJsonRequest**](PkiJsonRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKeyKeyRef

> PostPkiKeyKeyRef(ctx, keyRef, pkiMountPath).PkiKeyRequest(pkiKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	keyRef := "keyRef_example" // string | Reference to key; either \"default\" for the configured default key, an identifier of a key, or the name assigned to the key. (defaults to "default")
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKeyRequest := NewPkiKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiKeyKeyRef(context.Background(), keyRef, pkiMountPath, pkiKeyRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiKeyRequest** | [**PkiKeyRequest**](PkiKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKeysImport

> PostPkiKeysImport(ctx, pkiMountPath).PkiKeysImportRequest(pkiKeysImportRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKeysImportRequest := NewPkiKeysImportRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiKeysImport(context.Background(), pkiMountPath, pkiKeysImportRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKeysImportRequest** | [**PkiKeysImportRequest**](PkiKeysImportRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiKms

> PostPkiKms(ctx, pkiMountPath).PkiKmsRequest(pkiKmsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKmsRequest := NewPkiKmsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiKms(context.Background(), pkiMountPath, pkiKmsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiKmsRequest** | [**PkiKmsRequest**](PkiKmsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiOcsp

> PostPkiOcsp(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.PostPkiOcsp(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRevoke

> PostPkiRevoke(ctx, pkiMountPath).PkiRevokeRequest(pkiRevokeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRevokeRequest := NewPkiRevokeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRevoke(context.Background(), pkiMountPath, pkiRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRevokeRequest** | [**PkiRevokeRequest**](PkiRevokeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRevokeWithKey

> PostPkiRevokeWithKey(ctx, pkiMountPath).PkiRevokeWithKeyRequest(pkiRevokeWithKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRevokeWithKeyRequest := NewPkiRevokeWithKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRevokeWithKey(context.Background(), pkiMountPath, pkiRevokeWithKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRevokeWithKeyRequest** | [**PkiRevokeWithKeyRequest**](PkiRevokeWithKeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRolesName

> PostPkiRolesName(ctx, name, pkiMountPath).PkiRolesRequest(pkiRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRolesRequest := NewPkiRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRolesName(context.Background(), name, pkiMountPath, pkiRolesRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRolesRequest** | [**PkiRolesRequest**](PkiRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootGenerateExported

> PostPkiRootGenerateExported(ctx, exported, pkiMountPath).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootGenerateRequest := NewPkiRootGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRootGenerateExported(context.Background(), exported, pkiMountPath, pkiRootGenerateRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootGenerateRequest** | [**PkiRootGenerateRequest**](PkiRootGenerateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootReplace

> PostPkiRootReplace(ctx, pkiMountPath).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootReplaceRequest := NewPkiRootReplaceRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRootReplace(context.Background(), pkiMountPath, pkiRootReplaceRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootReplaceRequest** | [**PkiRootReplaceRequest**](PkiRootReplaceRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootRotateExported

> PostPkiRootRotateExported(ctx, exported, pkiMountPath).PkiRootRotateRequest(pkiRootRotateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	exported := "exported_example" // string | Must be \"internal\", \"exported\" or \"kms\". If set to \"exported\", the generated private key will be returned. This is your *only* chance to retrieve the private key!
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootRotateRequest := NewPkiRootRotateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRootRotateExported(context.Background(), exported, pkiMountPath, pkiRootRotateRequest)
	if err != nil {
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
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootRotateRequest** | [**PkiRootRotateRequest**](PkiRootRotateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootSignIntermediate

> PostPkiRootSignIntermediate(ctx, pkiMountPath).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootSignIntermediateRequest := NewPkiRootSignIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRootSignIntermediate(context.Background(), pkiMountPath, pkiRootSignIntermediateRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignIntermediateRequest** | [**PkiRootSignIntermediateRequest**](PkiRootSignIntermediateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiRootSignSelfIssued

> PostPkiRootSignSelfIssued(ctx, pkiMountPath).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootSignSelfIssuedRequest := NewPkiRootSignSelfIssuedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiRootSignSelfIssued(context.Background(), pkiMountPath, pkiRootSignSelfIssuedRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiRootSignSelfIssuedRequest** | [**PkiRootSignSelfIssuedRequest**](PkiRootSignSelfIssuedRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignRole

> PostPkiSignRole(ctx, pkiMountPath, role).PkiSignRequest(pkiSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignRequest := NewPkiSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiSignRole(context.Background(), pkiMountPath, role, pkiSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignRequest** | [**PkiSignRequest**](PkiSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignVerbatim

> PostPkiSignVerbatim(ctx, pkiMountPath).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiSignVerbatim(context.Background(), pkiMountPath, pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiSignVerbatimRole

> PostPkiSignVerbatimRole(ctx, pkiMountPath, role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request
	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiSignVerbatimRole(context.Background(), pkiMountPath, role, pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiTidy

> PostPkiTidy(ctx, pkiMountPath).PkiTidyRequest(pkiTidyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiTidyRequest := NewPkiTidyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostPkiTidy(context.Background(), pkiMountPath, pkiTidyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiTidyRequest** | [**PkiTidyRequest**](PkiTidyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostPkiTidyCancel

> PostPkiTidyCancel(ctx, pkiMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	pkiMountPath := "pkiMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.PostPkiTidyCancel(context.Background(), pkiMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**pkiMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqConfigConnection

> PostRabbitmqConfigConnection(ctx, rabbitmqMountPath).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()

Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	rabbitmqConfigConnectionRequest := NewRabbitmqConfigConnectionRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostRabbitmqConfigConnection(context.Background(), rabbitmqMountPath, rabbitmqConfigConnectionRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigConnectionRequest** | [**RabbitmqConfigConnectionRequest**](RabbitmqConfigConnectionRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqConfigLease

> PostRabbitmqConfigLease(ctx, rabbitmqMountPath).RabbitmqConfigLeaseRequest(rabbitmqConfigLeaseRequest).Execute()

Configure the lease parameters for generated credentials

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	rabbitmqConfigLeaseRequest := NewRabbitmqConfigLeaseRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostRabbitmqConfigLease(context.Background(), rabbitmqMountPath, rabbitmqConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rabbitmqConfigLeaseRequest** | [**RabbitmqConfigLeaseRequest**](RabbitmqConfigLeaseRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostRabbitmqRolesName

> PostRabbitmqRolesName(ctx, name, rabbitmqMountPath).RabbitmqRolesRequest(rabbitmqRolesRequest).Execute()

Manage the roles that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	rabbitmqMountPath := "rabbitmqMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	rabbitmqRolesRequest := NewRabbitmqRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostRabbitmqRolesName(context.Background(), name, rabbitmqMountPath, rabbitmqRolesRequest)
	if err != nil {
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
**rabbitmqMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitmqRolesRequest** | [**RabbitmqRolesRequest**](RabbitmqRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretConfig

> PostSecretConfig(ctx, secretMountPath).KvConfigRequest(kvConfigRequest).Execute()

Configure backend level settings that are applied to every key in the key-value store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvConfigRequest := NewKvConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretConfig(context.Background(), secretMountPath, kvConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvConfigRequest** | [**KvConfigRequest**](KvConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDataPath

> PostSecretDataPath(ctx, path, secretMountPath).KvDataRequest(kvDataRequest).Execute()

Write, Patch, Read, and Delete data in the Key-Value Store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDataRequest := NewKvDataRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretDataPath(context.Background(), path, secretMountPath, kvDataRequest)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDataRequest** | [**KvDataRequest**](KvDataRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDeletePath

> PostSecretDeletePath(ctx, path, secretMountPath).KvDeleteRequest(kvDeleteRequest).Execute()

Marks one or more versions as deleted in the KV store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDeleteRequest := NewKvDeleteRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretDeletePath(context.Background(), path, secretMountPath, kvDeleteRequest)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDeleteRequest** | [**KvDeleteRequest**](KvDeleteRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretDestroyPath

> PostSecretDestroyPath(ctx, path, secretMountPath).KvDestroyRequest(kvDestroyRequest).Execute()

Permanently removes one or more versions in the KV store

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDestroyRequest := NewKvDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretDestroyPath(context.Background(), path, secretMountPath, kvDestroyRequest)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvDestroyRequest** | [**KvDestroyRequest**](KvDestroyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretMetadataPath

> PostSecretMetadataPath(ctx, path, secretMountPath).KvMetadataRequest(kvMetadataRequest).Execute()

Configures settings for the KV store

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvMetadataRequest := NewKvMetadataRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretMetadataPath(context.Background(), path, secretMountPath, kvMetadataRequest)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvMetadataRequest** | [**KvMetadataRequest**](KvMetadataRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretPath

> PostSecretPath(ctx, path, secretMountPath).Execute()

Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.PostSecretPath(context.Background(), path, secretMountPath)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSecretUndeletePath

> PostSecretUndeletePath(ctx, path, secretMountPath).KvUndeleteRequest(kvUndeleteRequest).Execute()

Undeletes one or more versions from the KV store.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Location of the secret.
	secretMountPath := "secretMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvUndeleteRequest := NewKvUndeleteRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSecretUndeletePath(context.Background(), path, secretMountPath, kvUndeleteRequest)
	if err != nil {
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
**secretMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvUndeleteRequest** | [**KvUndeleteRequest**](KvUndeleteRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshConfigCa

> PostSshConfigCa(ctx, sshMountPath).SshConfigCaRequest(sshConfigCaRequest).Execute()

Set the SSH private key used for signing certificates.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshConfigCaRequest := NewSshConfigCaRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshConfigCa(context.Background(), sshMountPath, sshConfigCaRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigCaRequest** | [**SshConfigCaRequest**](SshConfigCaRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshConfigZeroaddress

> PostSshConfigZeroaddress(ctx, sshMountPath).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()

Assign zero address as default CIDR block for select roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshConfigZeroaddressRequest := NewSshConfigZeroaddressRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshConfigZeroaddress(context.Background(), sshMountPath, sshConfigZeroaddressRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshConfigZeroaddressRequest** | [**SshConfigZeroaddressRequest**](SshConfigZeroaddressRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshCredsRole

> PostSshCredsRole(ctx, role, sshMountPath).SshCredsRequest(sshCredsRequest).Execute()

Creates a credential for establishing SSH connection with the remote host.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required] Name of the role
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshCredsRequest := NewSshCredsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshCredsRole(context.Background(), role, sshMountPath, sshCredsRequest)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshCredsRequest** | [**SshCredsRequest**](SshCredsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshIssueRole

> PostSshIssueRole(ctx, role, sshMountPath).SshIssueRequest(sshIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshIssueRequest := NewSshIssueRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshIssueRole(context.Background(), role, sshMountPath, sshIssueRequest)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshIssueRequest** | [**SshIssueRequest**](SshIssueRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshKeysKeyName

> PostSshKeysKeyName(ctx, keyName, sshMountPath).SshKeysRequest(sshKeysRequest).Execute()

Register a shared private key with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	keyName := "keyName_example" // string | [Required] Name of the key
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshKeysRequest := NewSshKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshKeysKeyName(context.Background(), keyName, sshMountPath, sshKeysRequest)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeysRequest** | [**SshKeysRequest**](SshKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshLookup

> PostSshLookup(ctx, sshMountPath).SshLookupRequest(sshLookupRequest).Execute()

List all the roles associated with the given IP address.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshLookupRequest := NewSshLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshLookup(context.Background(), sshMountPath, sshLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshLookupRequest** | [**SshLookupRequest**](SshLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshRolesRole

> PostSshRolesRole(ctx, role, sshMountPath).SshRolesRequest(sshRolesRequest).Execute()

Manage the 'roles' that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | [Required for all types] Name of the role being created.
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshRolesRequest := NewSshRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshRolesRole(context.Background(), role, sshMountPath, sshRolesRequest)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshRolesRequest** | [**SshRolesRequest**](SshRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshSignRole

> PostSshSignRole(ctx, role, sshMountPath).SshSignRequest(sshSignRequest).Execute()

Request signing an SSH key using a certain role with the provided details.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The desired role with configuration for this request.
	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshSignRequest := NewSshSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshSignRole(context.Background(), role, sshMountPath, sshSignRequest)
	if err != nil {
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
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshSignRequest** | [**SshSignRequest**](SshSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostSshVerify

> PostSshVerify(ctx, sshMountPath).SshVerifyRequest(sshVerifyRequest).Execute()

Validate the OTP provided by Vault SSH Agent.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	sshMountPath := "sshMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshVerifyRequest := NewSshVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostSshVerify(context.Background(), sshMountPath, sshVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**sshMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshVerifyRequest** | [**SshVerifyRequest**](SshVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformConfig

> PostTerraformConfig(ctx, terraformMountPath).TerraformConfigRequest(terraformConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	terraformConfigRequest := NewTerraformConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTerraformConfig(context.Background(), terraformMountPath, terraformConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformConfigRequest** | [**TerraformConfigRequest**](TerraformConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformCredsName

> PostTerraformCredsName(ctx, name, terraformMountPath).Execute()

Generate a Terraform Cloud or Enterprise API token from a specific Vault role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.PostTerraformCredsName(context.Background(), name, terraformMountPath)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformRoleName

> PostTerraformRoleName(ctx, name, terraformMountPath).TerraformRoleRequest(terraformRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	terraformRoleRequest := NewTerraformRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTerraformRoleName(context.Background(), name, terraformMountPath, terraformRoleRequest)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformRoleRequest** | [**TerraformRoleRequest**](TerraformRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTerraformRotateRoleName

> PostTerraformRotateRoleName(ctx, name, terraformMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the team or organization role
	terraformMountPath := "terraformMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "terraform")

	resp, err := client.WithToken("my-token").Secrets.PostTerraformRotateRoleName(context.Background(), name, terraformMountPath)
	if err != nil {
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
**terraformMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;terraform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTotpCodeName

> PostTotpCodeName(ctx, name, totpMountPath).TotpCodeRequest(totpCodeRequest).Execute()

Request time-based one-time use password or validate a password for a certain key .

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	totpCodeRequest := NewTotpCodeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTotpCodeName(context.Background(), name, totpMountPath, totpCodeRequest)
	if err != nil {
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
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totpCodeRequest** | [**TotpCodeRequest**](TotpCodeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTotpKeysName

> PostTotpKeysName(ctx, name, totpMountPath).TotpKeysRequest(totpKeysRequest).Execute()

Manage the keys that can be created with this backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key.
	totpMountPath := "totpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	totpKeysRequest := NewTotpKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTotpKeysName(context.Background(), name, totpMountPath, totpKeysRequest)
	if err != nil {
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
**totpMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totpKeysRequest** | [**TotpKeysRequest**](TotpKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitCacheConfig

> PostTransitCacheConfig(ctx, transitMountPath).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()

Configures a new cache of the specified size

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitCacheConfigRequest := NewTransitCacheConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitCacheConfig(context.Background(), transitMountPath, transitCacheConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitCacheConfigRequest** | [**TransitCacheConfigRequest**](TransitCacheConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitDatakeyPlaintextName

> PostTransitDatakeyPlaintextName(ctx, name, plaintext, transitMountPath).TransitDatakeyRequest(transitDatakeyRequest).Execute()

Generate a data key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The backend key used for encrypting the data key
	plaintext := "plaintext_example" // string | \"plaintext\" will return the key in both plaintext and ciphertext; \"wrapped\" will return the ciphertext only.
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitDatakeyRequest := NewTransitDatakeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitDatakeyPlaintextName(context.Background(), name, plaintext, transitMountPath, transitDatakeyRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitDatakeyRequest** | [**TransitDatakeyRequest**](TransitDatakeyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitDecryptName

> PostTransitDecryptName(ctx, name, transitMountPath).TransitDecryptRequest(transitDecryptRequest).Execute()

Decrypt a ciphertext value using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitDecryptRequest := NewTransitDecryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitDecryptName(context.Background(), name, transitMountPath, transitDecryptRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitEncryptName

> PostTransitEncryptName(ctx, name, transitMountPath).TransitEncryptRequest(transitEncryptRequest).Execute()

Encrypt a plaintext value or a batch of plaintext blocks using a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitEncryptRequest := NewTransitEncryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitEncryptName(context.Background(), name, transitMountPath, transitEncryptRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHash

> PostTransitHash(ctx, transitMountPath).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHashRequest := NewTransitHashRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitHash(context.Background(), transitMountPath, transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHashUrlalgorithm

> PostTransitHashUrlalgorithm(ctx, transitMountPath, urlalgorithm).TransitHashRequest(transitHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHashRequest := NewTransitHashRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitHashUrlalgorithm(context.Background(), transitMountPath, urlalgorithm, transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHmacName

> PostTransitHmacName(ctx, name, transitMountPath).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitHmacName(context.Background(), name, transitMountPath, transitHmacRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitHmacNameUrlalgorithm

> PostTransitHmacNameUrlalgorithm(ctx, name, transitMountPath, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()

Generate an HMAC for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use for the HMAC function
	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitHmacNameUrlalgorithm(context.Background(), name, transitMountPath, urlalgorithm, transitHmacRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysName

> PostTransitKeysName(ctx, name, transitMountPath).TransitKeysRequest(transitKeysRequest).Execute()

Managed named encryption keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysRequest := NewTransitKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysName(context.Background(), name, transitMountPath, transitKeysRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysRequest** | [**TransitKeysRequest**](TransitKeysRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameConfig

> PostTransitKeysNameConfig(ctx, name, transitMountPath).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()

Configure a named encryption key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysConfigRequest := NewTransitKeysConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysNameConfig(context.Background(), name, transitMountPath, transitKeysConfigRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysConfigRequest** | [**TransitKeysConfigRequest**](TransitKeysConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameImport

> PostTransitKeysNameImport(ctx, name, transitMountPath).TransitKeysImportRequest(transitKeysImportRequest).Execute()

Imports an externally-generated key into a new transit key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysImportRequest := NewTransitKeysImportRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysNameImport(context.Background(), name, transitMountPath, transitKeysImportRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysImportRequest** | [**TransitKeysImportRequest**](TransitKeysImportRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameImportVersion

> PostTransitKeysNameImportVersion(ctx, name, transitMountPath).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()

Imports an externally-generated key into an existing imported key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysImportVersionRequest := NewTransitKeysImportVersionRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysNameImportVersion(context.Background(), name, transitMountPath, transitKeysImportVersionRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysImportVersionRequest** | [**TransitKeysImportVersionRequest**](TransitKeysImportVersionRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameRotate

> PostTransitKeysNameRotate(ctx, name, transitMountPath).Execute()

Rotate named encryption key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysNameRotate(context.Background(), name, transitMountPath)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitKeysNameTrim

> PostTransitKeysNameTrim(ctx, name, transitMountPath).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()

Trim key versions of a named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysTrimRequest := NewTransitKeysTrimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitKeysNameTrim(context.Background(), name, transitMountPath, transitKeysTrimRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitKeysTrimRequest** | [**TransitKeysTrimRequest**](TransitKeysTrimRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandom

> PostTransitRandom(ctx, transitMountPath).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRandom(context.Background(), transitMountPath, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomSource

> PostTransitRandomSource(ctx, source, transitMountPath).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRandomSource(context.Background(), source, transitMountPath, transitRandomRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomSourceUrlbytes

> PostTransitRandomSourceUrlbytes(ctx, source, transitMountPath, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRandomSourceUrlbytes(context.Background(), source, transitMountPath, urlbytes, transitRandomRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRandomUrlbytes

> PostTransitRandomUrlbytes(ctx, transitMountPath, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRandomUrlbytes(context.Background(), transitMountPath, urlbytes, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRestore

> PostTransitRestore(ctx, transitMountPath).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRestore(context.Background(), transitMountPath, transitRestoreRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRestoreName

> PostTransitRestoreName(ctx, name, transitMountPath).TransitRestoreRequest(transitRestoreRequest).Execute()

Restore the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | If set, this will be the name of the restored key.
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRestoreName(context.Background(), name, transitMountPath, transitRestoreRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitRewrapName

> PostTransitRewrapName(ctx, name, transitMountPath).TransitRewrapRequest(transitRewrapRequest).Execute()

Rewrap ciphertext

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRewrapRequest := NewTransitRewrapRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitRewrapName(context.Background(), name, transitMountPath, transitRewrapRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRewrapRequest** | [**TransitRewrapRequest**](TransitRewrapRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitSignName

> PostTransitSignName(ctx, name, transitMountPath).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitSignRequest := NewTransitSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitSignName(context.Background(), name, transitMountPath, transitSignRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitSignNameUrlalgorithm

> PostTransitSignNameUrlalgorithm(ctx, name, transitMountPath, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()

Generate a signature for input data using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitSignRequest := NewTransitSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitSignNameUrlalgorithm(context.Background(), name, transitMountPath, urlalgorithm, transitSignRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitVerifyName

> PostTransitVerifyName(ctx, name, transitMountPath).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitVerifyName(context.Background(), name, transitMountPath, transitVerifyRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostTransitVerifyNameUrlalgorithm

> PostTransitVerifyNameUrlalgorithm(ctx, name, transitMountPath, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()

Verify a signature or HMAC for input data created using the named key

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The key to use
	urlalgorithm := "urlalgorithm_example" // string | Hash algorithm to use (POST URL parameter)
	transitMountPath := "transitMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.PostTransitVerifyNameUrlalgorithm(context.Background(), name, transitMountPath, urlalgorithm, transitVerifyRequest)
	if err != nil {
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
**transitMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

