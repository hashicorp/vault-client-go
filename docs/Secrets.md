# Secrets

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAwsRolesName**](Secrets.md#DeleteAwsRolesName) | **Delete** /{mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**DeleteCubbyholePath**](Secrets.md#DeleteCubbyholePath) | **Delete** /{mount_path}/{path} | Deletes the secret at the specified location.
[**DeleteGcpRolesetName**](Secrets.md#DeleteGcpRolesetName) | **Delete** /{mount_path}/roleset/{name} | 
[**DeleteGcpStaticAccountName**](Secrets.md#DeleteGcpStaticAccountName) | **Delete** /{mount_path}/static-account/{name} | 
[**DeleteGcpkmsKeysDeregisterKey**](Secrets.md#DeleteGcpkmsKeysDeregisterKey) | **Delete** /{mount_path}/keys/deregister/{key} | Deregister an existing key in Vault
[**DeleteGcpkmsKeysKey**](Secrets.md#DeleteGcpkmsKeysKey) | **Delete** /{mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**DeleteGcpkmsKeysTrimKey**](Secrets.md#DeleteGcpkmsKeysTrimKey) | **Delete** /{mount_path}/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**DeleteLdapLibraryName**](Secrets.md#DeleteLdapLibraryName) | **Delete** /{mount_path}/library/{name} | Delete a library set.
[**DeleteLdapRoleName**](Secrets.md#DeleteLdapRoleName) | **Delete** /{mount_path}/role/{name} | 
[**DeleteLdapStaticRoleName**](Secrets.md#DeleteLdapStaticRoleName) | **Delete** /{mount_path}/static-role/{name} | 
[**DeleteNomadConfigAccess**](Secrets.md#DeleteNomadConfigAccess) | **Delete** /{mount_path}/config/access | 
[**DeletePkiIssuerRefDerPem**](Secrets.md#DeletePkiIssuerRefDerPem) | **Delete** /{mount_path}/{issuer_ref}/der|/pem | 
[**DeletePkiJson**](Secrets.md#DeletePkiJson) | **Delete** /{mount_path}//json | 
[**DeletePkiKeyKeyRef**](Secrets.md#DeletePkiKeyKeyRef) | **Delete** /{mount_path}/key/{key_ref} | 
[**DeletePkiRoot**](Secrets.md#DeletePkiRoot) | **Delete** /{mount_path}/root | 
[**DeleteSecretDataPath**](Secrets.md#DeleteSecretDataPath) | **Delete** /{mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**DeleteSecretMetadataPath**](Secrets.md#DeleteSecretMetadataPath) | **Delete** /{mount_path}/metadata/{path} | Configures settings for the KV store
[**DeleteSshConfigZeroaddress**](Secrets.md#DeleteSshConfigZeroaddress) | **Delete** /{mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**DeleteSshKeysKeyName**](Secrets.md#DeleteSshKeysKeyName) | **Delete** /{mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**DeleteSshRolesRole**](Secrets.md#DeleteSshRolesRole) | **Delete** /{mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**DeleteTotpKeysName**](Secrets.md#DeleteTotpKeysName) | **Delete** /{mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**ListAwsRoles**](Secrets.md#ListAwsRoles) | **Get** /{mount_path}/roles | List the existing roles in this backend
[**ListGcpRolesets**](Secrets.md#ListGcpRolesets) | **Get** /{mount_path}/rolesets | 
[**ListGcpStaticAccounts**](Secrets.md#ListGcpStaticAccounts) | **Get** /{mount_path}/static-accounts | 
[**ListLdapLibrary**](Secrets.md#ListLdapLibrary) | **Get** /{mount_path}/library | 
[**ListLdapRole**](Secrets.md#ListLdapRole) | **Get** /{mount_path}/role | 
[**ListLdapStaticRole**](Secrets.md#ListLdapStaticRole) | **Get** /{mount_path}/static-role | 
[**ListPkiCerts**](Secrets.md#ListPkiCerts) | **Get** /{mount_path}/certs | 
[**ListPkiIssuers**](Secrets.md#ListPkiIssuers) | **Get** /{mount_path}/issuers | 
[**ListTotpKeys**](Secrets.md#ListTotpKeys) | **Get** /{mount_path}/keys | Manage the keys that can be created with this backend.
[**ReadAwsConfigLease**](Secrets.md#ReadAwsConfigLease) | **Get** /{mount_path}/config/lease | Configure the default lease information for generated credentials.
[**ReadAwsConfigRoot**](Secrets.md#ReadAwsConfigRoot) | **Get** /{mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**ReadAwsCreds**](Secrets.md#ReadAwsCreds) | **Get** /{mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**ReadAwsRolesName**](Secrets.md#ReadAwsRolesName) | **Get** /{mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**ReadAwsStsName**](Secrets.md#ReadAwsStsName) | **Get** /{mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**ReadAzureCredsRole**](Secrets.md#ReadAzureCredsRole) | **Get** /{mount_path}/creds/{role} | 
[**ReadCubbyholePath**](Secrets.md#ReadCubbyholePath) | **Get** /{mount_path}/{path} | Retrieve the secret at the specified location.
[**ReadGcpConfig**](Secrets.md#ReadGcpConfig) | **Get** /{mount_path}/config | 
[**ReadGcpKeyRoleset**](Secrets.md#ReadGcpKeyRoleset) | **Get** /{mount_path}/key/{roleset} | 
[**ReadGcpRolesetName**](Secrets.md#ReadGcpRolesetName) | **Get** /{mount_path}/roleset/{name} | 
[**ReadGcpRolesetRolesetKey**](Secrets.md#ReadGcpRolesetRolesetKey) | **Get** /{mount_path}/roleset/{roleset}/key | 
[**ReadGcpRolesetRolesetToken**](Secrets.md#ReadGcpRolesetRolesetToken) | **Get** /{mount_path}/roleset/{roleset}/token | 
[**ReadGcpStaticAccountName**](Secrets.md#ReadGcpStaticAccountName) | **Get** /{mount_path}/static-account/{name} | 
[**ReadGcpStaticAccountNameKey**](Secrets.md#ReadGcpStaticAccountNameKey) | **Get** /{mount_path}/static-account/{name}/key | 
[**ReadGcpStaticAccountNameToken**](Secrets.md#ReadGcpStaticAccountNameToken) | **Get** /{mount_path}/static-account/{name}/token | 
[**ReadGcpTokenRoleset**](Secrets.md#ReadGcpTokenRoleset) | **Get** /{mount_path}/token/{roleset} | 
[**ReadGcpkmsKeysConfigKey**](Secrets.md#ReadGcpkmsKeysConfigKey) | **Get** /{mount_path}/keys/config/{key} | Configure the key in Vault
[**ReadGcpkmsKeysKey**](Secrets.md#ReadGcpkmsKeysKey) | **Get** /{mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**ReadGcpkmsPubkeyKey**](Secrets.md#ReadGcpkmsPubkeyKey) | **Get** /{mount_path}/pubkey/{key} | Retrieve the public key associated with the named key
[**ReadLdapCredsName**](Secrets.md#ReadLdapCredsName) | **Get** /{mount_path}/creds/{name} | 
[**ReadLdapLibraryName**](Secrets.md#ReadLdapLibraryName) | **Get** /{mount_path}/library/{name} | Read a library set.
[**ReadLdapLibraryNameStatus**](Secrets.md#ReadLdapLibraryNameStatus) | **Get** /{mount_path}/library/{name}/status | Check the status of the service accounts in a library set.
[**ReadLdapRoleName**](Secrets.md#ReadLdapRoleName) | **Get** /{mount_path}/role/{name} | 
[**ReadLdapStaticCredName**](Secrets.md#ReadLdapStaticCredName) | **Get** /{mount_path}/static-cred/{name} | 
[**ReadLdapStaticRoleName**](Secrets.md#ReadLdapStaticRoleName) | **Get** /{mount_path}/static-role/{name} | 
[**ReadNomadConfigAccess**](Secrets.md#ReadNomadConfigAccess) | **Get** /{mount_path}/config/access | 
[**ReadPkiCa**](Secrets.md#ReadPkiCa) | **Get** /{mount_path}/ca | 
[**ReadPkiCaChain**](Secrets.md#ReadPkiCaChain) | **Get** /{mount_path}/ca_chain | 
[**ReadPkiCaPem**](Secrets.md#ReadPkiCaPem) | **Get** /{mount_path}/ca/pem | 
[**ReadPkiCertCaChain**](Secrets.md#ReadPkiCertCaChain) | **Get** /{mount_path}/cert/ca_chain | 
[**ReadPkiCertSerial**](Secrets.md#ReadPkiCertSerial) | **Get** /{mount_path}/cert/{serial} | 
[**ReadPkiCertSerialRaw**](Secrets.md#ReadPkiCertSerialRaw) | **Get** /{mount_path}/cert/{serial}/raw | 
[**ReadPkiCertSerialRawPem**](Secrets.md#ReadPkiCertSerialRawPem) | **Get** /{mount_path}/cert/{serial}/raw/pem | 
[**ReadPkiConfigAutoTidy**](Secrets.md#ReadPkiConfigAutoTidy) | **Get** /{mount_path}/config/auto-tidy | 
[**ReadPkiConfigCrl**](Secrets.md#ReadPkiConfigCrl) | **Get** /{mount_path}/config/crl | 
[**ReadPkiConfigIssuers**](Secrets.md#ReadPkiConfigIssuers) | **Get** /{mount_path}/config/issuers | 
[**ReadPkiConfigKeys**](Secrets.md#ReadPkiConfigKeys) | **Get** /{mount_path}/config/keys | 
[**ReadPkiConfigUrls**](Secrets.md#ReadPkiConfigUrls) | **Get** /{mount_path}/config/urls | 
[**ReadPkiCrl**](Secrets.md#ReadPkiCrl) | **Get** /{mount_path}/crl | 
[**ReadPkiCrlRotate**](Secrets.md#ReadPkiCrlRotate) | **Get** /{mount_path}/crl/rotate | 
[**ReadPkiCrlRotateDelta**](Secrets.md#ReadPkiCrlRotateDelta) | **Get** /{mount_path}/crl/rotate-delta | 
[**ReadPkiDelta**](Secrets.md#ReadPkiDelta) | **Get** /{mount_path}//delta | 
[**ReadPkiDeltaCrl**](Secrets.md#ReadPkiDeltaCrl) | **Get** /{mount_path}/delta-crl | 
[**ReadPkiDeltaPem**](Secrets.md#ReadPkiDeltaPem) | **Get** /{mount_path}//delta/pem | 
[**ReadPkiDer**](Secrets.md#ReadPkiDer) | **Get** /{mount_path}//der | 
[**ReadPkiIssuerRefCrlPemDerDeltaPem**](Secrets.md#ReadPkiIssuerRefCrlPemDerDeltaPem) | **Get** /{mount_path}/{issuer_ref}/crl/pem|/der|/delta/pem | 
[**ReadPkiIssuerRefDerPem**](Secrets.md#ReadPkiIssuerRefDerPem) | **Get** /{mount_path}/{issuer_ref}/der|/pem | 
[**ReadPkiJson**](Secrets.md#ReadPkiJson) | **Get** /{mount_path}//json | 
[**ReadPkiKeyKeyRef**](Secrets.md#ReadPkiKeyKeyRef) | **Get** /{mount_path}/key/{key_ref} | 
[**ReadPkiOcspReq**](Secrets.md#ReadPkiOcspReq) | **Get** /{mount_path}/ocsp/{req} | 
[**ReadPkiPem**](Secrets.md#ReadPkiPem) | **Get** /{mount_path}//pem | 
[**ReadPkiTidyStatus**](Secrets.md#ReadPkiTidyStatus) | **Get** /{mount_path}/tidy-status | 
[**ReadSecretDataPath**](Secrets.md#ReadSecretDataPath) | **Get** /{mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**ReadSecretMetadataPath**](Secrets.md#ReadSecretMetadataPath) | **Get** /{mount_path}/metadata/{path} | Configures settings for the KV store
[**ReadSecretSubkeysPath**](Secrets.md#ReadSecretSubkeysPath) | **Get** /{mount_path}/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
[**ReadSshConfigZeroaddress**](Secrets.md#ReadSshConfigZeroaddress) | **Get** /{mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**ReadSshPublicKey**](Secrets.md#ReadSshPublicKey) | **Get** /{mount_path}/public_key | Retrieve the public key.
[**ReadSshRolesRole**](Secrets.md#ReadSshRolesRole) | **Get** /{mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**ReadTotpCodeName**](Secrets.md#ReadTotpCodeName) | **Get** /{mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**ReadTotpKeysName**](Secrets.md#ReadTotpKeysName) | **Get** /{mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**ReadTransitBackupName**](Secrets.md#ReadTransitBackupName) | **Get** /{mount_path}/backup/{name} | Backup the named key
[**ReadTransitCacheConfig**](Secrets.md#ReadTransitCacheConfig) | **Get** /{mount_path}/cache-config | Returns the size of the active cache
[**ReadTransitExportTypeName**](Secrets.md#ReadTransitExportTypeName) | **Get** /{mount_path}/export/{type}/{name} | Export named encryption or signing key
[**ReadTransitExportTypeNameVersion**](Secrets.md#ReadTransitExportTypeNameVersion) | **Get** /{mount_path}/export/{type}/{name}/{version} | Export named encryption or signing key
[**ReadTransitWrappingKey**](Secrets.md#ReadTransitWrappingKey) | **Get** /{mount_path}/wrapping_key | Returns the public key to use for wrapping imported keys
[**UpdateAwsConfigLease**](Secrets.md#UpdateAwsConfigLease) | **Post** /{mount_path}/config/lease | Configure the default lease information for generated credentials.
[**UpdateAwsConfigRoot**](Secrets.md#UpdateAwsConfigRoot) | **Post** /{mount_path}/config/root | Configure the root credentials that are used to manage IAM.
[**UpdateAwsCreds**](Secrets.md#UpdateAwsCreds) | **Post** /{mount_path}/creds | Generate AWS credentials from a specific Vault role.
[**UpdateAwsRolesName**](Secrets.md#UpdateAwsRolesName) | **Post** /{mount_path}/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
[**UpdateAwsStsName**](Secrets.md#UpdateAwsStsName) | **Post** /{mount_path}/sts/{name} | Generate AWS credentials from a specific Vault role.
[**UpdateCubbyholePath**](Secrets.md#UpdateCubbyholePath) | **Post** /{mount_path}/{path} | Store a secret at the specified location.
[**UpdateGcpConfig**](Secrets.md#UpdateGcpConfig) | **Post** /{mount_path}/config | 
[**UpdateGcpConfigRotateRoot**](Secrets.md#UpdateGcpConfigRotateRoot) | **Post** /{mount_path}/config/rotate-root | 
[**UpdateGcpKeyRoleset**](Secrets.md#UpdateGcpKeyRoleset) | **Post** /{mount_path}/key/{roleset} | 
[**UpdateGcpRolesetName**](Secrets.md#UpdateGcpRolesetName) | **Post** /{mount_path}/roleset/{name} | 
[**UpdateGcpRolesetNameRotate**](Secrets.md#UpdateGcpRolesetNameRotate) | **Post** /{mount_path}/roleset/{name}/rotate | 
[**UpdateGcpRolesetNameRotateKey**](Secrets.md#UpdateGcpRolesetNameRotateKey) | **Post** /{mount_path}/roleset/{name}/rotate-key | 
[**UpdateGcpRolesetRolesetKey**](Secrets.md#UpdateGcpRolesetRolesetKey) | **Post** /{mount_path}/roleset/{roleset}/key | 
[**UpdateGcpRolesetRolesetToken**](Secrets.md#UpdateGcpRolesetRolesetToken) | **Post** /{mount_path}/roleset/{roleset}/token | 
[**UpdateGcpStaticAccountName**](Secrets.md#UpdateGcpStaticAccountName) | **Post** /{mount_path}/static-account/{name} | 
[**UpdateGcpStaticAccountNameKey**](Secrets.md#UpdateGcpStaticAccountNameKey) | **Post** /{mount_path}/static-account/{name}/key | 
[**UpdateGcpStaticAccountNameRotateKey**](Secrets.md#UpdateGcpStaticAccountNameRotateKey) | **Post** /{mount_path}/static-account/{name}/rotate-key | 
[**UpdateGcpStaticAccountNameToken**](Secrets.md#UpdateGcpStaticAccountNameToken) | **Post** /{mount_path}/static-account/{name}/token | 
[**UpdateGcpTokenRoleset**](Secrets.md#UpdateGcpTokenRoleset) | **Post** /{mount_path}/token/{roleset} | 
[**UpdateGcpkmsDecryptKey**](Secrets.md#UpdateGcpkmsDecryptKey) | **Post** /{mount_path}/decrypt/{key} | Decrypt a ciphertext value using a named key
[**UpdateGcpkmsEncryptKey**](Secrets.md#UpdateGcpkmsEncryptKey) | **Post** /{mount_path}/encrypt/{key} | Encrypt a plaintext value using a named key
[**UpdateGcpkmsKeysConfigKey**](Secrets.md#UpdateGcpkmsKeysConfigKey) | **Post** /{mount_path}/keys/config/{key} | Configure the key in Vault
[**UpdateGcpkmsKeysDeregisterKey**](Secrets.md#UpdateGcpkmsKeysDeregisterKey) | **Post** /{mount_path}/keys/deregister/{key} | Deregister an existing key in Vault
[**UpdateGcpkmsKeysKey**](Secrets.md#UpdateGcpkmsKeysKey) | **Post** /{mount_path}/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
[**UpdateGcpkmsKeysRegisterKey**](Secrets.md#UpdateGcpkmsKeysRegisterKey) | **Post** /{mount_path}/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
[**UpdateGcpkmsKeysRotateKey**](Secrets.md#UpdateGcpkmsKeysRotateKey) | **Post** /{mount_path}/keys/rotate/{key} | Rotate a crypto key to a new primary version
[**UpdateGcpkmsKeysTrimKey**](Secrets.md#UpdateGcpkmsKeysTrimKey) | **Post** /{mount_path}/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
[**UpdateGcpkmsReencryptKey**](Secrets.md#UpdateGcpkmsReencryptKey) | **Post** /{mount_path}/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
[**UpdateGcpkmsSignKey**](Secrets.md#UpdateGcpkmsSignKey) | **Post** /{mount_path}/sign/{key} | Signs a message or digest using a named key
[**UpdateGcpkmsVerifyKey**](Secrets.md#UpdateGcpkmsVerifyKey) | **Post** /{mount_path}/verify/{key} | Verify a signature using a named key
[**UpdateLdapLibraryManageNameCheckIn**](Secrets.md#UpdateLdapLibraryManageNameCheckIn) | **Post** /{mount_path}/library/manage/{name}/check-in | Check service accounts in to the library.
[**UpdateLdapLibraryName**](Secrets.md#UpdateLdapLibraryName) | **Post** /{mount_path}/library/{name} | Update a library set.
[**UpdateLdapLibraryNameCheckIn**](Secrets.md#UpdateLdapLibraryNameCheckIn) | **Post** /{mount_path}/library/{name}/check-in | Check service accounts in to the library.
[**UpdateLdapLibraryNameCheckOut**](Secrets.md#UpdateLdapLibraryNameCheckOut) | **Post** /{mount_path}/library/{name}/check-out | Check a service account out from the library.
[**UpdateLdapRoleName**](Secrets.md#UpdateLdapRoleName) | **Post** /{mount_path}/role/{name} | 
[**UpdateLdapRotateRoleName**](Secrets.md#UpdateLdapRotateRoleName) | **Post** /{mount_path}/rotate-role/{name} | 
[**UpdateLdapRotateRoot**](Secrets.md#UpdateLdapRotateRoot) | **Post** /{mount_path}/rotate-root | 
[**UpdateLdapStaticRoleName**](Secrets.md#UpdateLdapStaticRoleName) | **Post** /{mount_path}/static-role/{name} | 
[**UpdateNomadConfigAccess**](Secrets.md#UpdateNomadConfigAccess) | **Post** /{mount_path}/config/access | 
[**UpdatePkiBundle**](Secrets.md#UpdatePkiBundle) | **Post** /{mount_path}/bundle | 
[**UpdatePkiCert**](Secrets.md#UpdatePkiCert) | **Post** /{mount_path}/cert | 
[**UpdatePkiConfigAutoTidy**](Secrets.md#UpdatePkiConfigAutoTidy) | **Post** /{mount_path}/config/auto-tidy | 
[**UpdatePkiConfigCa**](Secrets.md#UpdatePkiConfigCa) | **Post** /{mount_path}/config/ca | 
[**UpdatePkiConfigCrl**](Secrets.md#UpdatePkiConfigCrl) | **Post** /{mount_path}/config/crl | 
[**UpdatePkiConfigIssuers**](Secrets.md#UpdatePkiConfigIssuers) | **Post** /{mount_path}/config/issuers | 
[**UpdatePkiConfigKeys**](Secrets.md#UpdatePkiConfigKeys) | **Post** /{mount_path}/config/keys | 
[**UpdatePkiConfigUrls**](Secrets.md#UpdatePkiConfigUrls) | **Post** /{mount_path}/config/urls | 
[**UpdatePkiIntermediateCrossSign**](Secrets.md#UpdatePkiIntermediateCrossSign) | **Post** /{mount_path}/intermediate/cross-sign | 
[**UpdatePkiIntermediateGenerateExported**](Secrets.md#UpdatePkiIntermediateGenerateExported) | **Post** /{mount_path}/intermediate/generate/{exported} | 
[**UpdatePkiIntermediateSetSigned**](Secrets.md#UpdatePkiIntermediateSetSigned) | **Post** /{mount_path}/intermediate/set-signed | 
[**UpdatePkiIssueRole**](Secrets.md#UpdatePkiIssueRole) | **Post** /{mount_path}/issue/{role} | 
[**UpdatePkiIssuerIssuerRefIssueRole**](Secrets.md#UpdatePkiIssuerIssuerRefIssueRole) | **Post** /{mount_path}/issuer/{issuer_ref}/issue/{role} | 
[**UpdatePkiIssuerIssuerRefRevoke**](Secrets.md#UpdatePkiIssuerIssuerRefRevoke) | **Post** /{mount_path}/issuer/{issuer_ref}/revoke | 
[**UpdatePkiIssuerIssuerRefSignIntermediate**](Secrets.md#UpdatePkiIssuerIssuerRefSignIntermediate) | **Post** /{mount_path}/issuer/{issuer_ref}/sign-intermediate | 
[**UpdatePkiIssuerIssuerRefSignRole**](Secrets.md#UpdatePkiIssuerIssuerRefSignRole) | **Post** /{mount_path}/issuer/{issuer_ref}/sign/{role} | 
[**UpdatePkiIssuerIssuerRefSignSelfIssued**](Secrets.md#UpdatePkiIssuerIssuerRefSignSelfIssued) | **Post** /{mount_path}/issuer/{issuer_ref}/sign-self-issued | 
[**UpdatePkiIssuerIssuerRefSignVerbatim**](Secrets.md#UpdatePkiIssuerIssuerRefSignVerbatim) | **Post** /{mount_path}/issuer/{issuer_ref}/sign-verbatim | 
[**UpdatePkiIssuerIssuerRefSignVerbatimRole**](Secrets.md#UpdatePkiIssuerIssuerRefSignVerbatimRole) | **Post** /{mount_path}/issuer/{issuer_ref}/sign-verbatim/{role} | 
[**UpdatePkiIssuersGenerateIntermediateExported**](Secrets.md#UpdatePkiIssuersGenerateIntermediateExported) | **Post** /{mount_path}/issuers/generate/intermediate/{exported} | 
[**UpdatePkiIssuersGenerateRootExported**](Secrets.md#UpdatePkiIssuersGenerateRootExported) | **Post** /{mount_path}/issuers/generate/root/{exported} | 
[**UpdatePkiJson**](Secrets.md#UpdatePkiJson) | **Post** /{mount_path}//json | 
[**UpdatePkiKeyKeyRef**](Secrets.md#UpdatePkiKeyKeyRef) | **Post** /{mount_path}/key/{key_ref} | 
[**UpdatePkiKeysImport**](Secrets.md#UpdatePkiKeysImport) | **Post** /{mount_path}/keys/import | 
[**UpdatePkiKms**](Secrets.md#UpdatePkiKms) | **Post** /{mount_path}/kms | 
[**UpdatePkiOcsp**](Secrets.md#UpdatePkiOcsp) | **Post** /{mount_path}/ocsp | 
[**UpdatePkiRevoke**](Secrets.md#UpdatePkiRevoke) | **Post** /{mount_path}/revoke | 
[**UpdatePkiRevokeWithKey**](Secrets.md#UpdatePkiRevokeWithKey) | **Post** /{mount_path}/revoke-with-key | 
[**UpdatePkiRootGenerateExported**](Secrets.md#UpdatePkiRootGenerateExported) | **Post** /{mount_path}/root/generate/{exported} | 
[**UpdatePkiRootReplace**](Secrets.md#UpdatePkiRootReplace) | **Post** /{mount_path}/root/replace | 
[**UpdatePkiRootRotateExported**](Secrets.md#UpdatePkiRootRotateExported) | **Post** /{mount_path}/root/rotate/{exported} | 
[**UpdatePkiRootSignIntermediate**](Secrets.md#UpdatePkiRootSignIntermediate) | **Post** /{mount_path}/root/sign-intermediate | 
[**UpdatePkiRootSignSelfIssued**](Secrets.md#UpdatePkiRootSignSelfIssued) | **Post** /{mount_path}/root/sign-self-issued | 
[**UpdatePkiSignRole**](Secrets.md#UpdatePkiSignRole) | **Post** /{mount_path}/sign/{role} | 
[**UpdatePkiSignVerbatim**](Secrets.md#UpdatePkiSignVerbatim) | **Post** /{mount_path}/sign-verbatim | 
[**UpdatePkiSignVerbatimRole**](Secrets.md#UpdatePkiSignVerbatimRole) | **Post** /{mount_path}/sign-verbatim/{role} | 
[**UpdatePkiTidy**](Secrets.md#UpdatePkiTidy) | **Post** /{mount_path}/tidy | 
[**UpdatePkiTidyCancel**](Secrets.md#UpdatePkiTidyCancel) | **Post** /{mount_path}/tidy-cancel | 
[**UpdateRabbitmqConfigConnection**](Secrets.md#UpdateRabbitmqConfigConnection) | **Post** /{mount_path}/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
[**UpdateSecretDataPath**](Secrets.md#UpdateSecretDataPath) | **Post** /{mount_path}/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
[**UpdateSecretDeletePath**](Secrets.md#UpdateSecretDeletePath) | **Post** /{mount_path}/delete/{path} | Marks one or more versions as deleted in the KV store.
[**UpdateSecretDestroyPath**](Secrets.md#UpdateSecretDestroyPath) | **Post** /{mount_path}/destroy/{path} | Permanently removes one or more versions in the KV store
[**UpdateSecretMetadataPath**](Secrets.md#UpdateSecretMetadataPath) | **Post** /{mount_path}/metadata/{path} | Configures settings for the KV store
[**UpdateSecretUndeletePath**](Secrets.md#UpdateSecretUndeletePath) | **Post** /{mount_path}/undelete/{path} | Undeletes one or more versions from the KV store.
[**UpdateSshConfigZeroaddress**](Secrets.md#UpdateSshConfigZeroaddress) | **Post** /{mount_path}/config/zeroaddress | Assign zero address as default CIDR block for select roles.
[**UpdateSshKeysKeyName**](Secrets.md#UpdateSshKeysKeyName) | **Post** /{mount_path}/keys/{key_name} | Register a shared private key with Vault.
[**UpdateSshLookup**](Secrets.md#UpdateSshLookup) | **Post** /{mount_path}/lookup | List all the roles associated with the given IP address.
[**UpdateSshRolesRole**](Secrets.md#UpdateSshRolesRole) | **Post** /{mount_path}/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
[**UpdateSshVerify**](Secrets.md#UpdateSshVerify) | **Post** /{mount_path}/verify | Validate the OTP provided by Vault SSH Agent.
[**UpdateTotpCodeName**](Secrets.md#UpdateTotpCodeName) | **Post** /{mount_path}/code/{name} | Request time-based one-time use password or validate a password for a certain key .
[**UpdateTotpKeysName**](Secrets.md#UpdateTotpKeysName) | **Post** /{mount_path}/keys/{name} | Manage the keys that can be created with this backend.
[**UpdateTransitCacheConfig**](Secrets.md#UpdateTransitCacheConfig) | **Post** /{mount_path}/cache-config | Configures a new cache of the specified size
[**UpdateTransitDatakeyPlaintextName**](Secrets.md#UpdateTransitDatakeyPlaintextName) | **Post** /{mount_path}/datakey/{plaintext}/{name} | Generate a data key
[**UpdateTransitDecryptName**](Secrets.md#UpdateTransitDecryptName) | **Post** /{mount_path}/decrypt/{name} | Decrypt a ciphertext value using a named key
[**UpdateTransitEncryptName**](Secrets.md#UpdateTransitEncryptName) | **Post** /{mount_path}/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
[**UpdateTransitHash**](Secrets.md#UpdateTransitHash) | **Post** /{mount_path}/hash | Generate a hash sum for input data
[**UpdateTransitHashUrlalgorithm**](Secrets.md#UpdateTransitHashUrlalgorithm) | **Post** /{mount_path}/hash/{urlalgorithm} | Generate a hash sum for input data
[**UpdateTransitHmacName**](Secrets.md#UpdateTransitHmacName) | **Post** /{mount_path}/hmac/{name} | Generate an HMAC for input data using the named key
[**UpdateTransitHmacNameUrlalgorithm**](Secrets.md#UpdateTransitHmacNameUrlalgorithm) | **Post** /{mount_path}/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
[**UpdateTransitKeysNameConfig**](Secrets.md#UpdateTransitKeysNameConfig) | **Post** /{mount_path}/keys/{name}/config | Configure a named encryption key
[**UpdateTransitKeysNameImport**](Secrets.md#UpdateTransitKeysNameImport) | **Post** /{mount_path}/keys/{name}/import | Imports an externally-generated key into a new transit key
[**UpdateTransitKeysNameImportVersion**](Secrets.md#UpdateTransitKeysNameImportVersion) | **Post** /{mount_path}/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
[**UpdateTransitKeysNameRotate**](Secrets.md#UpdateTransitKeysNameRotate) | **Post** /{mount_path}/keys/{name}/rotate | Rotate named encryption key
[**UpdateTransitKeysNameTrim**](Secrets.md#UpdateTransitKeysNameTrim) | **Post** /{mount_path}/keys/{name}/trim | Trim key versions of a named key
[**UpdateTransitRandom**](Secrets.md#UpdateTransitRandom) | **Post** /{mount_path}/random | Generate random bytes
[**UpdateTransitRandomSource**](Secrets.md#UpdateTransitRandomSource) | **Post** /{mount_path}/random/{source} | Generate random bytes
[**UpdateTransitRandomSourceUrlbytes**](Secrets.md#UpdateTransitRandomSourceUrlbytes) | **Post** /{mount_path}/random/{source}/{urlbytes} | Generate random bytes
[**UpdateTransitRandomUrlbytes**](Secrets.md#UpdateTransitRandomUrlbytes) | **Post** /{mount_path}/random/{urlbytes} | Generate random bytes
[**UpdateTransitRestore**](Secrets.md#UpdateTransitRestore) | **Post** /{mount_path}/restore | Restore the named key
[**UpdateTransitRestoreName**](Secrets.md#UpdateTransitRestoreName) | **Post** /{mount_path}/restore/{name} | Restore the named key
[**UpdateTransitRewrapName**](Secrets.md#UpdateTransitRewrapName) | **Post** /{mount_path}/rewrap/{name} | Rewrap ciphertext
[**UpdateTransitSignName**](Secrets.md#UpdateTransitSignName) | **Post** /{mount_path}/sign/{name} | Generate a signature for input data using the named key
[**UpdateTransitSignNameUrlalgorithm**](Secrets.md#UpdateTransitSignNameUrlalgorithm) | **Post** /{mount_path}/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
[**UpdateTransitVerifyName**](Secrets.md#UpdateTransitVerifyName) | **Post** /{mount_path}/verify/{name} | Verify a signature or HMAC for input data created using the named key
[**UpdateTransitVerifyNameUrlalgorithm**](Secrets.md#UpdateTransitVerifyNameUrlalgorithm) | **Post** /{mount_path}/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key



## DeleteAwsRolesName

> DeleteAwsRolesName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.DeleteAwsRolesName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteCubbyholePath

> DeleteCubbyholePath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	resp, err := client.WithToken("my-token").Secrets.DeleteCubbyholePath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpRolesetName

> DeleteGcpRolesetName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpRolesetName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpStaticAccountName

> DeleteGcpStaticAccountName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpStaticAccountName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysDeregisterKey

> DeleteGcpkmsKeysDeregisterKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysDeregisterKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysKey

> DeleteGcpkmsKeysKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGcpkmsKeysTrimKey

> DeleteGcpkmsKeysTrimKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.DeleteGcpkmsKeysTrimKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapLibraryName

> DeleteLdapLibraryName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapLibraryName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapRoleName

> DeleteLdapRoleName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteLdapStaticRoleName

> DeleteLdapStaticRoleName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.DeleteLdapStaticRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteNomadConfigAccess

> DeleteNomadConfigAccess(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.DeleteNomadConfigAccess(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiIssuerRefDerPem

> DeletePkiIssuerRefDerPem(ctx, issuerRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiIssuerRefDerPem(context.Background(), issuerRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiJson

> DeletePkiJson(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiJson(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiKeyKeyRef

> DeletePkiKeyKeyRef(ctx, keyRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiKeyKeyRef(context.Background(), keyRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeletePkiRoot

> DeletePkiRoot(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.DeletePkiRoot(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSecretDataPath

> DeleteSecretDataPath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.DeleteSecretDataPath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSecretMetadataPath

> DeleteSecretMetadataPath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.DeleteSecretMetadataPath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshConfigZeroaddress

> DeleteSshConfigZeroaddress(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshConfigZeroaddress(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshKeysKeyName

> DeleteSshKeysKeyName(ctx, keyName, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshKeysKeyName(context.Background(), keyName, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteSshRolesRole

> DeleteSshRolesRole(ctx, mountPath, role).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.DeleteSshRolesRole(context.Background(), mountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTotpKeysName

> DeleteTotpKeysName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.DeleteTotpKeysName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]
**name** | **string** | Name of the key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListAwsRoles

> ListAwsRoles(ctx, mountPath).List(list).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListAwsRoles(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListGcpRolesets

> ListGcpRolesets(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListGcpRolesets(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListGcpStaticAccounts

> ListGcpStaticAccounts(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListGcpStaticAccounts(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListLdapLibrary

> ListLdapLibrary(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListLdapLibrary(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListLdapRole

> ListLdapRole(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListLdapRole(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListLdapStaticRole

> ListLdapStaticRole(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListLdapStaticRole(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListPkiCerts

> ListPkiCerts(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListPkiCerts(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListPkiIssuers

> ListPkiIssuers(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListPkiIssuers(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListTotpKeys

> ListTotpKeys(ctx, mountPath).List(list).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ListTotpKeys(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsConfigLease

> ReadAwsConfigLease(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.ReadAwsConfigLease(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsConfigRoot

> ReadAwsConfigRoot(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.ReadAwsConfigRoot(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsCreds

> ReadAwsCreds(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.ReadAwsCreds(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsRolesName

> ReadAwsRolesName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.ReadAwsRolesName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsStsName

> ReadAwsStsName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Secrets.ReadAwsStsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAzureCredsRole

> ReadAzureCredsRole(ctx, mountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "azure")

	resp, err := client.WithToken("my-token").Secrets.ReadAzureCredsRole(context.Background(), mountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;azure&quot;]
**role** | **string** | Name of the Vault role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadCubbyholePath

> ReadCubbyholePath(ctx, mountPath, path).List(list).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ReadCubbyholePath(context.Background(), mountPath, path, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpConfig

> ReadGcpConfig(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpConfig(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpKeyRoleset

> ReadGcpKeyRoleset(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpKeyRoleset(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpRolesetName

> ReadGcpRolesetName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpRolesetName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpRolesetRolesetKey

> ReadGcpRolesetRolesetKey(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpRolesetRolesetKey(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpRolesetRolesetToken

> ReadGcpRolesetRolesetToken(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpRolesetRolesetToken(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpStaticAccountName

> ReadGcpStaticAccountName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpStaticAccountName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpStaticAccountNameKey

> ReadGcpStaticAccountNameKey(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpStaticAccountNameKey(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpStaticAccountNameToken

> ReadGcpStaticAccountNameToken(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpStaticAccountNameToken(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpTokenRoleset

> ReadGcpTokenRoleset(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpTokenRoleset(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpkmsKeysConfigKey

> ReadGcpkmsKeysConfigKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpkmsKeysConfigKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpkmsKeysKey

> ReadGcpkmsKeysKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpkmsKeysKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGcpkmsPubkeyKey

> ReadGcpkmsPubkeyKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.ReadGcpkmsPubkeyKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapCredsName

> ReadLdapCredsName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapCredsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the dynamic role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapLibraryName

> ReadLdapLibraryName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapLibraryName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapLibraryNameStatus

> ReadLdapLibraryNameStatus(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapLibraryNameStatus(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapRoleName

> ReadLdapRoleName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapStaticCredName

> ReadLdapStaticCredName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapStaticCredName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the static role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadLdapStaticRoleName

> ReadLdapStaticRoleName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.ReadLdapStaticRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadNomadConfigAccess

> ReadNomadConfigAccess(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	resp, err := client.WithToken("my-token").Secrets.ReadNomadConfigAccess(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCa

> ReadPkiCa(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCa(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCaChain

> ReadPkiCaChain(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCaChain(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCaPem

> ReadPkiCaPem(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCaPem(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCertCaChain

> ReadPkiCertCaChain(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCertCaChain(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCertSerial

> ReadPkiCertSerial(ctx, mountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCertSerial(context.Background(), mountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCertSerialRaw

> ReadPkiCertSerialRaw(ctx, mountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCertSerialRaw(context.Background(), mountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCertSerialRawPem

> ReadPkiCertSerialRawPem(ctx, mountPath, serial).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCertSerialRawPem(context.Background(), mountPath, serial)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**serial** | **string** | Certificate serial number, in colon- or hyphen-separated octal | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiConfigAutoTidy

> ReadPkiConfigAutoTidy(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiConfigAutoTidy(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiConfigCrl

> ReadPkiConfigCrl(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiConfigCrl(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiConfigIssuers

> ReadPkiConfigIssuers(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiConfigIssuers(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiConfigKeys

> ReadPkiConfigKeys(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiConfigKeys(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiConfigUrls

> ReadPkiConfigUrls(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiConfigUrls(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCrl

> ReadPkiCrl(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCrl(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCrlRotate

> ReadPkiCrlRotate(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCrlRotate(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiCrlRotateDelta

> ReadPkiCrlRotateDelta(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiCrlRotateDelta(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiDelta

> ReadPkiDelta(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiDelta(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiDeltaCrl

> ReadPkiDeltaCrl(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiDeltaCrl(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiDeltaPem

> ReadPkiDeltaPem(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiDeltaPem(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiDer

> ReadPkiDer(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiDer(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiIssuerRefCrlPemDerDeltaPem

> ReadPkiIssuerRefCrlPemDerDeltaPem(ctx, issuerRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiIssuerRefCrlPemDerDeltaPem(context.Background(), issuerRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiIssuerRefDerPem

> ReadPkiIssuerRefDerPem(ctx, issuerRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiIssuerRefDerPem(context.Background(), issuerRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiJson

> ReadPkiJson(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiJson(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiKeyKeyRef

> ReadPkiKeyKeyRef(ctx, keyRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiKeyKeyRef(context.Background(), keyRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiOcspReq

> ReadPkiOcspReq(ctx, mountPath, req).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiOcspReq(context.Background(), mountPath, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**req** | **string** | base-64 encoded ocsp request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiPem

> ReadPkiPem(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiPem(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadPkiTidyStatus

> ReadPkiTidyStatus(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.ReadPkiTidyStatus(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSecretDataPath

> ReadSecretDataPath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.ReadSecretDataPath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSecretMetadataPath

> ReadSecretMetadataPath(ctx, mountPath, path).List(list).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.ReadSecretMetadataPath(context.Background(), mountPath, path, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSecretSubkeysPath

> ReadSecretSubkeysPath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	resp, err := client.WithToken("my-token").Secrets.ReadSecretSubkeysPath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSshConfigZeroaddress

> ReadSshConfigZeroaddress(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.ReadSshConfigZeroaddress(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSshPublicKey

> ReadSshPublicKey(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.ReadSshPublicKey(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadSshRolesRole

> ReadSshRolesRole(ctx, mountPath, role).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	resp, err := client.WithToken("my-token").Secrets.ReadSshRolesRole(context.Background(), mountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTotpCodeName

> ReadTotpCodeName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.ReadTotpCodeName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]
**name** | **string** | Name of the key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTotpKeysName

> ReadTotpKeysName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	resp, err := client.WithToken("my-token").Secrets.ReadTotpKeysName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]
**name** | **string** | Name of the key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTransitBackupName

> ReadTransitBackupName(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.ReadTransitBackupName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTransitCacheConfig

> ReadTransitCacheConfig(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.ReadTransitCacheConfig(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTransitExportTypeName

> ReadTransitExportTypeName(ctx, mountPath, name, type_).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.ReadTransitExportTypeName(context.Background(), mountPath, name, type_)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTransitExportTypeNameVersion

> ReadTransitExportTypeNameVersion(ctx, mountPath, name, type_, version).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.ReadTransitExportTypeNameVersion(context.Background(), mountPath, name, type_, version)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 
**type_** | **string** | Type of key to export (encryption-key, signing-key, hmac-key) | 
**version** | **string** | Version of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTransitWrappingKey

> ReadTransitWrappingKey(ctx, mountPath).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.ReadTransitWrappingKey(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigLease

> UpdateAwsConfigLease(ctx, mountPath).AwsConfigLeaseRequest(awsConfigLeaseRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigLeaseRequest := NewAwsConfigLeaseRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateAwsConfigLease(context.Background(), mountPath, awsConfigLeaseRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigLeaseRequest** | [**AwsConfigLeaseRequest**](AwsConfigLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigRoot

> UpdateAwsConfigRoot(ctx, mountPath).AwsConfigRootRequest(awsConfigRootRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigRootRequest := NewAwsConfigRootRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateAwsConfigRoot(context.Background(), mountPath, awsConfigRootRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigRootRequest** | [**AwsConfigRootRequest**](AwsConfigRootRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsCreds

> UpdateAwsCreds(ctx, mountPath).AwsCredsRequest(awsCredsRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsCredsRequest := NewAwsCredsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateAwsCreds(context.Background(), mountPath, awsCredsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredsRequest** | [**AwsCredsRequest**](AwsCredsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsRolesName

> UpdateAwsRolesName(ctx, mountPath, name).AwsRolesRequest(awsRolesRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRolesRequest := NewAwsRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateAwsRolesName(context.Background(), mountPath, name, awsRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the policy | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsRolesRequest** | [**AwsRolesRequest**](AwsRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsStsName

> UpdateAwsStsName(ctx, mountPath, name).AwsStsRequest(awsStsRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsStsRequest := NewAwsStsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateAwsStsName(context.Background(), mountPath, name, awsStsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsStsRequest** | [**AwsStsRequest**](AwsStsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateCubbyholePath

> UpdateCubbyholePath(ctx, mountPath, path).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cubbyhole")

	resp, err := client.WithToken("my-token").Secrets.UpdateCubbyholePath(context.Background(), mountPath, path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cubbyhole&quot;]
**path** | **string** | Specifies the path of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpConfig

> UpdateGcpConfig(ctx, mountPath).GcpConfigRequest(gcpConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpConfigRequest := NewGcpConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpConfig(context.Background(), mountPath, gcpConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpConfigRequest** | [**GcpConfigRequest**](GcpConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpConfigRotateRoot

> UpdateGcpConfigRotateRoot(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpConfigRotateRoot(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpKeyRoleset

> UpdateGcpKeyRoleset(ctx, mountPath, roleset).GcpKeyRequest(gcpKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpKeyRequest := NewGcpKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpKeyRoleset(context.Background(), mountPath, roleset, gcpKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpKeyRequest** | [**GcpKeyRequest**](GcpKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRolesetName

> UpdateGcpRolesetName(ctx, mountPath, name).GcpRolesetRequest(gcpRolesetRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRolesetRequest := NewGcpRolesetRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpRolesetName(context.Background(), mountPath, name, gcpRolesetRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpRolesetRequest** | [**GcpRolesetRequest**](GcpRolesetRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRolesetNameRotate

> UpdateGcpRolesetNameRotate(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpRolesetNameRotate(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRolesetNameRotateKey

> UpdateGcpRolesetNameRotateKey(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpRolesetNameRotateKey(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRolesetRolesetKey

> UpdateGcpRolesetRolesetKey(ctx, mountPath, roleset).GcpRolesetKeyRequest(gcpRolesetKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRolesetKeyRequest := NewGcpRolesetKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpRolesetRolesetKey(context.Background(), mountPath, roleset, gcpRolesetKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpRolesetKeyRequest** | [**GcpRolesetKeyRequest**](GcpRolesetKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRolesetRolesetToken

> UpdateGcpRolesetRolesetToken(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpRolesetRolesetToken(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpStaticAccountName

> UpdateGcpStaticAccountName(ctx, mountPath, name).GcpStaticAccountRequest(gcpStaticAccountRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpStaticAccountRequest := NewGcpStaticAccountRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpStaticAccountName(context.Background(), mountPath, name, gcpStaticAccountRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name to refer to this static account in Vault. Cannot be updated. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpStaticAccountRequest** | [**GcpStaticAccountRequest**](GcpStaticAccountRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpStaticAccountNameKey

> UpdateGcpStaticAccountNameKey(ctx, mountPath, name).GcpStaticAccountKeyRequest(gcpStaticAccountKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpStaticAccountKeyRequest := NewGcpStaticAccountKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpStaticAccountNameKey(context.Background(), mountPath, name, gcpStaticAccountKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpStaticAccountKeyRequest** | [**GcpStaticAccountKeyRequest**](GcpStaticAccountKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpStaticAccountNameRotateKey

> UpdateGcpStaticAccountNameRotateKey(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpStaticAccountNameRotateKey(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Name of the account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpStaticAccountNameToken

> UpdateGcpStaticAccountNameToken(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpStaticAccountNameToken(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**name** | **string** | Required. Name of the static account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpTokenRoleset

> UpdateGcpTokenRoleset(ctx, mountPath, roleset).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpTokenRoleset(context.Background(), mountPath, roleset)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcp&quot;]
**roleset** | **string** | Required. Name of the role set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsDecryptKey

> UpdateGcpkmsDecryptKey(ctx, key, mountPath).GcpkmsDecryptRequest(gcpkmsDecryptRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsDecryptRequest := NewGcpkmsDecryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsDecryptKey(context.Background(), key, mountPath, gcpkmsDecryptRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsDecryptRequest** | [**GcpkmsDecryptRequest**](GcpkmsDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsEncryptKey

> UpdateGcpkmsEncryptKey(ctx, key, mountPath).GcpkmsEncryptRequest(gcpkmsEncryptRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsEncryptRequest := NewGcpkmsEncryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsEncryptKey(context.Background(), key, mountPath, gcpkmsEncryptRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsEncryptRequest** | [**GcpkmsEncryptRequest**](GcpkmsEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysConfigKey

> UpdateGcpkmsKeysConfigKey(ctx, key, mountPath).GcpkmsKeysConfigRequest(gcpkmsKeysConfigRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysConfigRequest := NewGcpkmsKeysConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysConfigKey(context.Background(), key, mountPath, gcpkmsKeysConfigRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsKeysConfigRequest** | [**GcpkmsKeysConfigRequest**](GcpkmsKeysConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysDeregisterKey

> UpdateGcpkmsKeysDeregisterKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysDeregisterKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysKey

> UpdateGcpkmsKeysKey(ctx, key, mountPath).GcpkmsKeysRequest(gcpkmsKeysRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysRequest := NewGcpkmsKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysKey(context.Background(), key, mountPath, gcpkmsKeysRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsKeysRequest** | [**GcpkmsKeysRequest**](GcpkmsKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysRegisterKey

> UpdateGcpkmsKeysRegisterKey(ctx, key, mountPath).GcpkmsKeysRegisterRequest(gcpkmsKeysRegisterRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsKeysRegisterRequest := NewGcpkmsKeysRegisterRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysRegisterKey(context.Background(), key, mountPath, gcpkmsKeysRegisterRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsKeysRegisterRequest** | [**GcpkmsKeysRegisterRequest**](GcpkmsKeysRegisterRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysRotateKey

> UpdateGcpkmsKeysRotateKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysRotateKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsKeysTrimKey

> UpdateGcpkmsKeysTrimKey(ctx, key, mountPath).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsKeysTrimKey(context.Background(), key, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsReencryptKey

> UpdateGcpkmsReencryptKey(ctx, key, mountPath).GcpkmsReencryptRequest(gcpkmsReencryptRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsReencryptRequest := NewGcpkmsReencryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsReencryptKey(context.Background(), key, mountPath, gcpkmsReencryptRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsReencryptRequest** | [**GcpkmsReencryptRequest**](GcpkmsReencryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsSignKey

> UpdateGcpkmsSignKey(ctx, key, mountPath).GcpkmsSignRequest(gcpkmsSignRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsSignRequest := NewGcpkmsSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsSignKey(context.Background(), key, mountPath, gcpkmsSignRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsSignRequest** | [**GcpkmsSignRequest**](GcpkmsSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpkmsVerifyKey

> UpdateGcpkmsVerifyKey(ctx, key, mountPath).GcpkmsVerifyRequest(gcpkmsVerifyRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcpkms")

	gcpkmsVerifyRequest := NewGcpkmsVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateGcpkmsVerifyKey(context.Background(), key, mountPath, gcpkmsVerifyRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;gcpkms&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpkmsVerifyRequest** | [**GcpkmsVerifyRequest**](GcpkmsVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapLibraryManageNameCheckIn

> UpdateLdapLibraryManageNameCheckIn(ctx, mountPath, name).LdapLibraryManageCheckInRequest(ldapLibraryManageCheckInRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryManageCheckInRequest := NewLdapLibraryManageCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapLibraryManageNameCheckIn(context.Background(), mountPath, name, ldapLibraryManageCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryManageCheckInRequest** | [**LdapLibraryManageCheckInRequest**](LdapLibraryManageCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapLibraryName

> UpdateLdapLibraryName(ctx, mountPath, name).LdapLibraryRequest(ldapLibraryRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryRequest := NewLdapLibraryRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapLibraryName(context.Background(), mountPath, name, ldapLibraryRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryRequest** | [**LdapLibraryRequest**](LdapLibraryRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapLibraryNameCheckIn

> UpdateLdapLibraryNameCheckIn(ctx, mountPath, name).LdapLibraryCheckInRequest(ldapLibraryCheckInRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryCheckInRequest := NewLdapLibraryCheckInRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapLibraryNameCheckIn(context.Background(), mountPath, name, ldapLibraryCheckInRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryCheckInRequest** | [**LdapLibraryCheckInRequest**](LdapLibraryCheckInRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapLibraryNameCheckOut

> UpdateLdapLibraryNameCheckOut(ctx, mountPath, name).LdapLibraryCheckOutRequest(ldapLibraryCheckOutRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLibraryCheckOutRequest := NewLdapLibraryCheckOutRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapLibraryNameCheckOut(context.Background(), mountPath, name, ldapLibraryCheckOutRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the set | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapLibraryCheckOutRequest** | [**LdapLibraryCheckOutRequest**](LdapLibraryCheckOutRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapRoleName

> UpdateLdapRoleName(ctx, mountPath, name).LdapRoleRequest(ldapRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapRoleRequest := NewLdapRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapRoleName(context.Background(), mountPath, name, ldapRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role (lowercase) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapRoleRequest** | [**LdapRoleRequest**](LdapRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapRotateRoleName

> UpdateLdapRotateRoleName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.UpdateLdapRotateRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the static role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapRotateRoot

> UpdateLdapRotateRoot(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Secrets.UpdateLdapRotateRoot(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateLdapStaticRoleName

> UpdateLdapStaticRoleName(ctx, mountPath, name).LdapStaticRoleRequest(ldapStaticRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapStaticRoleRequest := NewLdapStaticRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateLdapStaticRoleName(context.Background(), mountPath, name, ldapStaticRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ldap&quot;]
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ldapStaticRoleRequest** | [**LdapStaticRoleRequest**](LdapStaticRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateNomadConfigAccess

> UpdateNomadConfigAccess(ctx, mountPath).NomadConfigAccessRequest(nomadConfigAccessRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "nomad")

	nomadConfigAccessRequest := NewNomadConfigAccessRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateNomadConfigAccess(context.Background(), mountPath, nomadConfigAccessRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;nomad&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nomadConfigAccessRequest** | [**NomadConfigAccessRequest**](NomadConfigAccessRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiBundle

> UpdatePkiBundle(ctx, mountPath).PkiBundleRequest(pkiBundleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiBundleRequest := NewPkiBundleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiBundle(context.Background(), mountPath, pkiBundleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiBundleRequest** | [**PkiBundleRequest**](PkiBundleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiCert

> UpdatePkiCert(ctx, mountPath).PkiCertRequest(pkiCertRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiCertRequest := NewPkiCertRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiCert(context.Background(), mountPath, pkiCertRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiCertRequest** | [**PkiCertRequest**](PkiCertRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigAutoTidy

> UpdatePkiConfigAutoTidy(ctx, mountPath).PkiConfigAutoTidyRequest(pkiConfigAutoTidyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigAutoTidyRequest := NewPkiConfigAutoTidyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigAutoTidy(context.Background(), mountPath, pkiConfigAutoTidyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigAutoTidyRequest** | [**PkiConfigAutoTidyRequest**](PkiConfigAutoTidyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigCa

> UpdatePkiConfigCa(ctx, mountPath).PkiConfigCaRequest(pkiConfigCaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigCaRequest := NewPkiConfigCaRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigCa(context.Background(), mountPath, pkiConfigCaRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigCaRequest** | [**PkiConfigCaRequest**](PkiConfigCaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigCrl

> UpdatePkiConfigCrl(ctx, mountPath).PkiConfigCrlRequest(pkiConfigCrlRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigCrlRequest := NewPkiConfigCrlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigCrl(context.Background(), mountPath, pkiConfigCrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigCrlRequest** | [**PkiConfigCrlRequest**](PkiConfigCrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigIssuers

> UpdatePkiConfigIssuers(ctx, mountPath).PkiConfigIssuersRequest(pkiConfigIssuersRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigIssuersRequest := NewPkiConfigIssuersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigIssuers(context.Background(), mountPath, pkiConfigIssuersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigIssuersRequest** | [**PkiConfigIssuersRequest**](PkiConfigIssuersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigKeys

> UpdatePkiConfigKeys(ctx, mountPath).PkiConfigKeysRequest(pkiConfigKeysRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigKeysRequest := NewPkiConfigKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigKeys(context.Background(), mountPath, pkiConfigKeysRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigKeysRequest** | [**PkiConfigKeysRequest**](PkiConfigKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiConfigUrls

> UpdatePkiConfigUrls(ctx, mountPath).PkiConfigUrlsRequest(pkiConfigUrlsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiConfigUrlsRequest := NewPkiConfigUrlsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiConfigUrls(context.Background(), mountPath, pkiConfigUrlsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiConfigUrlsRequest** | [**PkiConfigUrlsRequest**](PkiConfigUrlsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIntermediateCrossSign

> UpdatePkiIntermediateCrossSign(ctx, mountPath).PkiIntermediateCrossSignRequest(pkiIntermediateCrossSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateCrossSignRequest := NewPkiIntermediateCrossSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIntermediateCrossSign(context.Background(), mountPath, pkiIntermediateCrossSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIntermediateCrossSignRequest** | [**PkiIntermediateCrossSignRequest**](PkiIntermediateCrossSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIntermediateGenerateExported

> UpdatePkiIntermediateGenerateExported(ctx, exported, mountPath).PkiIntermediateGenerateRequest(pkiIntermediateGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateGenerateRequest := NewPkiIntermediateGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIntermediateGenerateExported(context.Background(), exported, mountPath, pkiIntermediateGenerateRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIntermediateGenerateRequest** | [**PkiIntermediateGenerateRequest**](PkiIntermediateGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIntermediateSetSigned

> UpdatePkiIntermediateSetSigned(ctx, mountPath).PkiIntermediateSetSignedRequest(pkiIntermediateSetSignedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIntermediateSetSignedRequest := NewPkiIntermediateSetSignedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIntermediateSetSigned(context.Background(), mountPath, pkiIntermediateSetSignedRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiIntermediateSetSignedRequest** | [**PkiIntermediateSetSignedRequest**](PkiIntermediateSetSignedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssueRole

> UpdatePkiIssueRole(ctx, mountPath, role).PkiIssueRequest(pkiIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssueRequest := NewPkiIssueRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssueRole(context.Background(), mountPath, role, pkiIssueRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssueRequest** | [**PkiIssueRequest**](PkiIssueRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefIssueRole

> UpdatePkiIssuerIssuerRefIssueRole(ctx, issuerRef, mountPath, role).PkiIssuerIssueRequest(pkiIssuerIssueRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerIssueRequest := NewPkiIssuerIssueRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefIssueRole(context.Background(), issuerRef, mountPath, role, pkiIssuerIssueRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerIssueRequest** | [**PkiIssuerIssueRequest**](PkiIssuerIssueRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefRevoke

> UpdatePkiIssuerIssuerRefRevoke(ctx, issuerRef, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefRevoke(context.Background(), issuerRef, mountPath)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefSignIntermediate

> UpdatePkiIssuerIssuerRefSignIntermediate(ctx, issuerRef, mountPath).PkiIssuerSignIntermediateRequest(pkiIssuerSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignIntermediateRequest := NewPkiIssuerSignIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefSignIntermediate(context.Background(), issuerRef, mountPath, pkiIssuerSignIntermediateRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignIntermediateRequest** | [**PkiIssuerSignIntermediateRequest**](PkiIssuerSignIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefSignRole

> UpdatePkiIssuerIssuerRefSignRole(ctx, issuerRef, mountPath, role).PkiIssuerSignRequest(pkiIssuerSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignRequest := NewPkiIssuerSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefSignRole(context.Background(), issuerRef, mountPath, role, pkiIssuerSignRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerSignRequest** | [**PkiIssuerSignRequest**](PkiIssuerSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefSignSelfIssued

> UpdatePkiIssuerIssuerRefSignSelfIssued(ctx, issuerRef, mountPath).PkiIssuerSignSelfIssuedRequest(pkiIssuerSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignSelfIssuedRequest := NewPkiIssuerSignSelfIssuedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefSignSelfIssued(context.Background(), issuerRef, mountPath, pkiIssuerSignSelfIssuedRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignSelfIssuedRequest** | [**PkiIssuerSignSelfIssuedRequest**](PkiIssuerSignSelfIssuedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefSignVerbatim

> UpdatePkiIssuerIssuerRefSignVerbatim(ctx, issuerRef, mountPath).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefSignVerbatim(context.Background(), issuerRef, mountPath, pkiIssuerSignVerbatimRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuerIssuerRefSignVerbatimRole

> UpdatePkiIssuerIssuerRefSignVerbatimRole(ctx, issuerRef, mountPath, role).PkiIssuerSignVerbatimRequest(pkiIssuerSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuerSignVerbatimRequest := NewPkiIssuerSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuerIssuerRefSignVerbatimRole(context.Background(), issuerRef, mountPath, role, pkiIssuerSignVerbatimRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pkiIssuerSignVerbatimRequest** | [**PkiIssuerSignVerbatimRequest**](PkiIssuerSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuersGenerateIntermediateExported

> UpdatePkiIssuersGenerateIntermediateExported(ctx, exported, mountPath).PkiIssuersGenerateIntermediateRequest(pkiIssuersGenerateIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuersGenerateIntermediateRequest := NewPkiIssuersGenerateIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuersGenerateIntermediateExported(context.Background(), exported, mountPath, pkiIssuersGenerateIntermediateRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuersGenerateIntermediateRequest** | [**PkiIssuersGenerateIntermediateRequest**](PkiIssuersGenerateIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiIssuersGenerateRootExported

> UpdatePkiIssuersGenerateRootExported(ctx, exported, mountPath).PkiIssuersGenerateRootRequest(pkiIssuersGenerateRootRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiIssuersGenerateRootRequest := NewPkiIssuersGenerateRootRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiIssuersGenerateRootExported(context.Background(), exported, mountPath, pkiIssuersGenerateRootRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiIssuersGenerateRootRequest** | [**PkiIssuersGenerateRootRequest**](PkiIssuersGenerateRootRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiJson

> UpdatePkiJson(ctx, mountPath).PkiJsonRequest(pkiJsonRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiJsonRequest := NewPkiJsonRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiJson(context.Background(), mountPath, pkiJsonRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiJsonRequest** | [**PkiJsonRequest**](PkiJsonRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiKeyKeyRef

> UpdatePkiKeyKeyRef(ctx, keyRef, mountPath).PkiKeyRequest(pkiKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKeyRequest := NewPkiKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiKeyKeyRef(context.Background(), keyRef, mountPath, pkiKeyRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiKeyRequest** | [**PkiKeyRequest**](PkiKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiKeysImport

> UpdatePkiKeysImport(ctx, mountPath).PkiKeysImportRequest(pkiKeysImportRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKeysImportRequest := NewPkiKeysImportRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiKeysImport(context.Background(), mountPath, pkiKeysImportRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiKeysImportRequest** | [**PkiKeysImportRequest**](PkiKeysImportRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiKms

> UpdatePkiKms(ctx, mountPath).PkiKmsRequest(pkiKmsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiKmsRequest := NewPkiKmsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiKms(context.Background(), mountPath, pkiKmsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiKmsRequest** | [**PkiKmsRequest**](PkiKmsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiOcsp

> UpdatePkiOcsp(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.UpdatePkiOcsp(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRevoke

> UpdatePkiRevoke(ctx, mountPath).PkiRevokeRequest(pkiRevokeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRevokeRequest := NewPkiRevokeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRevoke(context.Background(), mountPath, pkiRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRevokeRequest** | [**PkiRevokeRequest**](PkiRevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRevokeWithKey

> UpdatePkiRevokeWithKey(ctx, mountPath).PkiRevokeWithKeyRequest(pkiRevokeWithKeyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRevokeWithKeyRequest := NewPkiRevokeWithKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRevokeWithKey(context.Background(), mountPath, pkiRevokeWithKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRevokeWithKeyRequest** | [**PkiRevokeWithKeyRequest**](PkiRevokeWithKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRootGenerateExported

> UpdatePkiRootGenerateExported(ctx, exported, mountPath).PkiRootGenerateRequest(pkiRootGenerateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootGenerateRequest := NewPkiRootGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRootGenerateExported(context.Background(), exported, mountPath, pkiRootGenerateRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiRootGenerateRequest** | [**PkiRootGenerateRequest**](PkiRootGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRootReplace

> UpdatePkiRootReplace(ctx, mountPath).PkiRootReplaceRequest(pkiRootReplaceRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootReplaceRequest := NewPkiRootReplaceRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRootReplace(context.Background(), mountPath, pkiRootReplaceRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootReplaceRequest** | [**PkiRootReplaceRequest**](PkiRootReplaceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRootRotateExported

> UpdatePkiRootRotateExported(ctx, exported, mountPath).PkiRootRotateRequest(pkiRootRotateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootRotateRequest := NewPkiRootRotateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRootRotateExported(context.Background(), exported, mountPath, pkiRootRotateRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiRootRotateRequest** | [**PkiRootRotateRequest**](PkiRootRotateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRootSignIntermediate

> UpdatePkiRootSignIntermediate(ctx, mountPath).PkiRootSignIntermediateRequest(pkiRootSignIntermediateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootSignIntermediateRequest := NewPkiRootSignIntermediateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRootSignIntermediate(context.Background(), mountPath, pkiRootSignIntermediateRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootSignIntermediateRequest** | [**PkiRootSignIntermediateRequest**](PkiRootSignIntermediateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiRootSignSelfIssued

> UpdatePkiRootSignSelfIssued(ctx, mountPath).PkiRootSignSelfIssuedRequest(pkiRootSignSelfIssuedRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiRootSignSelfIssuedRequest := NewPkiRootSignSelfIssuedRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiRootSignSelfIssued(context.Background(), mountPath, pkiRootSignSelfIssuedRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiRootSignSelfIssuedRequest** | [**PkiRootSignSelfIssuedRequest**](PkiRootSignSelfIssuedRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiSignRole

> UpdatePkiSignRole(ctx, mountPath, role).PkiSignRequest(pkiSignRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignRequest := NewPkiSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiSignRole(context.Background(), mountPath, role, pkiSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiSignRequest** | [**PkiSignRequest**](PkiSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiSignVerbatim

> UpdatePkiSignVerbatim(ctx, mountPath).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiSignVerbatim(context.Background(), mountPath, pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiSignVerbatimRole

> UpdatePkiSignVerbatimRole(ctx, mountPath, role).PkiSignVerbatimRequest(pkiSignVerbatimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiSignVerbatimRequest := NewPkiSignVerbatimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiSignVerbatimRole(context.Background(), mountPath, role, pkiSignVerbatimRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]
**role** | **string** | The desired role with configuration for this request | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pkiSignVerbatimRequest** | [**PkiSignVerbatimRequest**](PkiSignVerbatimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiTidy

> UpdatePkiTidy(ctx, mountPath).PkiTidyRequest(pkiTidyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	pkiTidyRequest := NewPkiTidyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdatePkiTidy(context.Background(), mountPath, pkiTidyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pkiTidyRequest** | [**PkiTidyRequest**](PkiTidyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdatePkiTidyCancel

> UpdatePkiTidyCancel(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "pki")

	resp, err := client.WithToken("my-token").Secrets.UpdatePkiTidyCancel(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;pki&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateRabbitmqConfigConnection

> UpdateRabbitmqConfigConnection(ctx, mountPath).RabbitmqConfigConnectionRequest(rabbitmqConfigConnectionRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "rabbitmq")

	rabbitmqConfigConnectionRequest := NewRabbitmqConfigConnectionRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateRabbitmqConfigConnection(context.Background(), mountPath, rabbitmqConfigConnectionRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;rabbitmq&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rabbitmqConfigConnectionRequest** | [**RabbitmqConfigConnectionRequest**](RabbitmqConfigConnectionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSecretDataPath

> UpdateSecretDataPath(ctx, mountPath, path).KvDataRequest(kvDataRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDataRequest := NewKvDataRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSecretDataPath(context.Background(), mountPath, path, kvDataRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvDataRequest** | [**KvDataRequest**](KvDataRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSecretDeletePath

> UpdateSecretDeletePath(ctx, mountPath, path).KvDeleteRequest(kvDeleteRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDeleteRequest := NewKvDeleteRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSecretDeletePath(context.Background(), mountPath, path, kvDeleteRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvDeleteRequest** | [**KvDeleteRequest**](KvDeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSecretDestroyPath

> UpdateSecretDestroyPath(ctx, mountPath, path).KvDestroyRequest(kvDestroyRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvDestroyRequest := NewKvDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSecretDestroyPath(context.Background(), mountPath, path, kvDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvDestroyRequest** | [**KvDestroyRequest**](KvDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSecretMetadataPath

> UpdateSecretMetadataPath(ctx, mountPath, path).KvMetadataRequest(kvMetadataRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvMetadataRequest := NewKvMetadataRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSecretMetadataPath(context.Background(), mountPath, path, kvMetadataRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvMetadataRequest** | [**KvMetadataRequest**](KvMetadataRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSecretUndeletePath

> UpdateSecretUndeletePath(ctx, mountPath, path).KvUndeleteRequest(kvUndeleteRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "secret")

	kvUndeleteRequest := NewKvUndeleteRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSecretUndeletePath(context.Background(), mountPath, path, kvUndeleteRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;secret&quot;]
**path** | **string** | Location of the secret. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kvUndeleteRequest** | [**KvUndeleteRequest**](KvUndeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSshConfigZeroaddress

> UpdateSshConfigZeroaddress(ctx, mountPath).SshConfigZeroaddressRequest(sshConfigZeroaddressRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshConfigZeroaddressRequest := NewSshConfigZeroaddressRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSshConfigZeroaddress(context.Background(), mountPath, sshConfigZeroaddressRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshConfigZeroaddressRequest** | [**SshConfigZeroaddressRequest**](SshConfigZeroaddressRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSshKeysKeyName

> UpdateSshKeysKeyName(ctx, keyName, mountPath).SshKeysRequest(sshKeysRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshKeysRequest := NewSshKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSshKeysKeyName(context.Background(), keyName, mountPath, sshKeysRequest)
	if err != nil {
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
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshKeysRequest** | [**SshKeysRequest**](SshKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSshLookup

> UpdateSshLookup(ctx, mountPath).SshLookupRequest(sshLookupRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshLookupRequest := NewSshLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSshLookup(context.Background(), mountPath, sshLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshLookupRequest** | [**SshLookupRequest**](SshLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSshRolesRole

> UpdateSshRolesRole(ctx, mountPath, role).SshRolesRequest(sshRolesRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshRolesRequest := NewSshRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSshRolesRole(context.Background(), mountPath, role, sshRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]
**role** | **string** | [Required for all types] Name of the role being created. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sshRolesRequest** | [**SshRolesRequest**](SshRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateSshVerify

> UpdateSshVerify(ctx, mountPath).SshVerifyRequest(sshVerifyRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ssh")

	sshVerifyRequest := NewSshVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateSshVerify(context.Background(), mountPath, sshVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;ssh&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshVerifyRequest** | [**SshVerifyRequest**](SshVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTotpCodeName

> UpdateTotpCodeName(ctx, mountPath, name).TotpCodeRequest(totpCodeRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	totpCodeRequest := NewTotpCodeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTotpCodeName(context.Background(), mountPath, name, totpCodeRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]
**name** | **string** | Name of the key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **totpCodeRequest** | [**TotpCodeRequest**](TotpCodeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTotpKeysName

> UpdateTotpKeysName(ctx, mountPath, name).TotpKeysRequest(totpKeysRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "totp")

	totpKeysRequest := NewTotpKeysRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTotpKeysName(context.Background(), mountPath, name, totpKeysRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;totp&quot;]
**name** | **string** | Name of the key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **totpKeysRequest** | [**TotpKeysRequest**](TotpKeysRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitCacheConfig

> UpdateTransitCacheConfig(ctx, mountPath).TransitCacheConfigRequest(transitCacheConfigRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitCacheConfigRequest := NewTransitCacheConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitCacheConfig(context.Background(), mountPath, transitCacheConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitCacheConfigRequest** | [**TransitCacheConfigRequest**](TransitCacheConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitDatakeyPlaintextName

> UpdateTransitDatakeyPlaintextName(ctx, mountPath, name, plaintext).TransitDatakeyRequest(transitDatakeyRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitDatakeyRequest := NewTransitDatakeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitDatakeyPlaintextName(context.Background(), mountPath, name, plaintext, transitDatakeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The backend key used for encrypting the data key | 
**plaintext** | **string** | \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitDatakeyRequest** | [**TransitDatakeyRequest**](TransitDatakeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitDecryptName

> UpdateTransitDecryptName(ctx, mountPath, name).TransitDecryptRequest(transitDecryptRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitDecryptRequest := NewTransitDecryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitDecryptName(context.Background(), mountPath, name, transitDecryptRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitDecryptRequest** | [**TransitDecryptRequest**](TransitDecryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitEncryptName

> UpdateTransitEncryptName(ctx, mountPath, name).TransitEncryptRequest(transitEncryptRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitEncryptRequest := NewTransitEncryptRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitEncryptName(context.Background(), mountPath, name, transitEncryptRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitEncryptRequest** | [**TransitEncryptRequest**](TransitEncryptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitHash

> UpdateTransitHash(ctx, mountPath).TransitHashRequest(transitHashRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHashRequest := NewTransitHashRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitHash(context.Background(), mountPath, transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitHashUrlalgorithm

> UpdateTransitHashUrlalgorithm(ctx, mountPath, urlalgorithm).TransitHashRequest(transitHashRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHashRequest := NewTransitHashRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitHashUrlalgorithm(context.Background(), mountPath, urlalgorithm, transitHashRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitHashRequest** | [**TransitHashRequest**](TransitHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitHmacName

> UpdateTransitHmacName(ctx, mountPath, name).TransitHmacRequest(transitHmacRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitHmacName(context.Background(), mountPath, name, transitHmacRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use for the HMAC function | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitHmacNameUrlalgorithm

> UpdateTransitHmacNameUrlalgorithm(ctx, mountPath, name, urlalgorithm).TransitHmacRequest(transitHmacRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitHmacRequest := NewTransitHmacRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitHmacNameUrlalgorithm(context.Background(), mountPath, name, urlalgorithm, transitHmacRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use for the HMAC function | 
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitHmacRequest** | [**TransitHmacRequest**](TransitHmacRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitKeysNameConfig

> UpdateTransitKeysNameConfig(ctx, mountPath, name).TransitKeysConfigRequest(transitKeysConfigRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysConfigRequest := NewTransitKeysConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitKeysNameConfig(context.Background(), mountPath, name, transitKeysConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitKeysConfigRequest** | [**TransitKeysConfigRequest**](TransitKeysConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitKeysNameImport

> UpdateTransitKeysNameImport(ctx, mountPath, name).TransitKeysImportRequest(transitKeysImportRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysImportRequest := NewTransitKeysImportRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitKeysNameImport(context.Background(), mountPath, name, transitKeysImportRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitKeysImportRequest** | [**TransitKeysImportRequest**](TransitKeysImportRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitKeysNameImportVersion

> UpdateTransitKeysNameImportVersion(ctx, mountPath, name).TransitKeysImportVersionRequest(transitKeysImportVersionRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysImportVersionRequest := NewTransitKeysImportVersionRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitKeysNameImportVersion(context.Background(), mountPath, name, transitKeysImportVersionRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitKeysImportVersionRequest** | [**TransitKeysImportVersionRequest**](TransitKeysImportVersionRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitKeysNameRotate

> UpdateTransitKeysNameRotate(ctx, mountPath, name).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	resp, err := client.WithToken("my-token").Secrets.UpdateTransitKeysNameRotate(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitKeysNameTrim

> UpdateTransitKeysNameTrim(ctx, mountPath, name).TransitKeysTrimRequest(transitKeysTrimRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitKeysTrimRequest := NewTransitKeysTrimRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitKeysNameTrim(context.Background(), mountPath, name, transitKeysTrimRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitKeysTrimRequest** | [**TransitKeysTrimRequest**](TransitKeysTrimRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRandom

> UpdateTransitRandom(ctx, mountPath).TransitRandomRequest(transitRandomRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRandom(context.Background(), mountPath, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRandomSource

> UpdateTransitRandomSource(ctx, mountPath, source).TransitRandomRequest(transitRandomRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRandomSource(context.Background(), mountPath, source, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRandomSourceUrlbytes

> UpdateTransitRandomSourceUrlbytes(ctx, mountPath, source, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRandomSourceUrlbytes(context.Background(), mountPath, source, urlbytes, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRandomUrlbytes

> UpdateTransitRandomUrlbytes(ctx, mountPath, urlbytes).TransitRandomRequest(transitRandomRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRandomRequest := NewTransitRandomRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRandomUrlbytes(context.Background(), mountPath, urlbytes, transitRandomRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRandomRequest** | [**TransitRandomRequest**](TransitRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRestore

> UpdateTransitRestore(ctx, mountPath).TransitRestoreRequest(transitRestoreRequest).Execute()

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

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRestore(context.Background(), mountPath, transitRestoreRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRestoreName

> UpdateTransitRestoreName(ctx, mountPath, name).TransitRestoreRequest(transitRestoreRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRestoreRequest := NewTransitRestoreRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRestoreName(context.Background(), mountPath, name, transitRestoreRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | If set, this will be the name of the restored key. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRestoreRequest** | [**TransitRestoreRequest**](TransitRestoreRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitRewrapName

> UpdateTransitRewrapName(ctx, mountPath, name).TransitRewrapRequest(transitRewrapRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitRewrapRequest := NewTransitRewrapRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitRewrapName(context.Background(), mountPath, name, transitRewrapRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitRewrapRequest** | [**TransitRewrapRequest**](TransitRewrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitSignName

> UpdateTransitSignName(ctx, mountPath, name).TransitSignRequest(transitSignRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitSignRequest := NewTransitSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitSignName(context.Background(), mountPath, name, transitSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitSignNameUrlalgorithm

> UpdateTransitSignNameUrlalgorithm(ctx, mountPath, name, urlalgorithm).TransitSignRequest(transitSignRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitSignRequest := NewTransitSignRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitSignNameUrlalgorithm(context.Background(), mountPath, name, urlalgorithm, transitSignRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitSignRequest** | [**TransitSignRequest**](TransitSignRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitVerifyName

> UpdateTransitVerifyName(ctx, mountPath, name).TransitVerifyRequest(transitVerifyRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitVerifyName(context.Background(), mountPath, name, transitVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTransitVerifyNameUrlalgorithm

> UpdateTransitVerifyNameUrlalgorithm(ctx, mountPath, name, urlalgorithm).TransitVerifyRequest(transitVerifyRequest).Execute()

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "transit")

	transitVerifyRequest := NewTransitVerifyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Secrets.UpdateTransitVerifyNameUrlalgorithm(context.Background(), mountPath, name, urlalgorithm, transitVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;transit&quot;]
**name** | **string** | The key to use | 
**urlalgorithm** | **string** | Hash algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitVerifyRequest** | [**TransitVerifyRequest**](TransitVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

