# Auth

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApproleRoleRoleName**](Auth.md#DeleteApproleRoleRoleName) | **Delete** /auth/{mount_path}/role/{role_name} | Register an role with the backend.
[**DeleteApproleRoleRoleNameBindSecretId**](Auth.md#DeleteApproleRoleRoleNameBindSecretId) | **Delete** /auth/{mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**DeleteApproleRoleRoleNameBoundCidrList**](Auth.md#DeleteApproleRoleRoleNameBoundCidrList) | **Delete** /auth/{mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteApproleRoleRoleNamePeriod**](Auth.md#DeleteApproleRoleRoleNamePeriod) | **Delete** /auth/{mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**DeleteApproleRoleRoleNamePolicies**](Auth.md#DeleteApproleRoleRoleNamePolicies) | **Delete** /auth/{mount_path}/role/{role_name}/policies | Policies of the role.
[**DeleteApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#DeleteApproleRoleRoleNameSecretIdAccessorDestroy) | **Delete** /auth/{mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**DeleteApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#DeleteApproleRoleRoleNameSecretIdBoundCidrs) | **Delete** /auth/{mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteApproleRoleRoleNameSecretIdDestroy**](Auth.md#DeleteApproleRoleRoleNameSecretIdDestroy) | **Delete** /auth/{mount_path}/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**DeleteApproleRoleRoleNameSecretIdNumUses**](Auth.md#DeleteApproleRoleRoleNameSecretIdNumUses) | **Delete** /auth/{mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**DeleteApproleRoleRoleNameSecretIdTtl**](Auth.md#DeleteApproleRoleRoleNameSecretIdTtl) | **Delete** /auth/{mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**DeleteApproleRoleRoleNameTokenBoundCidrs**](Auth.md#DeleteApproleRoleRoleNameTokenBoundCidrs) | **Delete** /auth/{mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**DeleteApproleRoleRoleNameTokenMaxTtl**](Auth.md#DeleteApproleRoleRoleNameTokenMaxTtl) | **Delete** /auth/{mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**DeleteApproleRoleRoleNameTokenNumUses**](Auth.md#DeleteApproleRoleRoleNameTokenNumUses) | **Delete** /auth/{mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**DeleteApproleRoleRoleNameTokenTtl**](Auth.md#DeleteApproleRoleRoleNameTokenTtl) | **Delete** /auth/{mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**DeleteAwsConfigCertificateCertName**](Auth.md#DeleteAwsConfigCertificateCertName) | **Delete** /auth/{mount_path}/config/certificate/{cert_name} | 
[**DeleteAwsConfigClient**](Auth.md#DeleteAwsConfigClient) | **Delete** /auth/{mount_path}/config/client | 
[**DeleteAwsConfigStsAccountId**](Auth.md#DeleteAwsConfigStsAccountId) | **Delete** /auth/{mount_path}/config/sts/{account_id} | 
[**DeleteAwsConfigTidyIdentityAccesslist**](Auth.md#DeleteAwsConfigTidyIdentityAccesslist) | **Delete** /auth/{mount_path}/config/tidy/identity-accesslist | 
[**DeleteAwsConfigTidyIdentityWhitelist**](Auth.md#DeleteAwsConfigTidyIdentityWhitelist) | **Delete** /auth/{mount_path}/config/tidy/identity-whitelist | 
[**DeleteAwsConfigTidyRoletagBlacklist**](Auth.md#DeleteAwsConfigTidyRoletagBlacklist) | **Delete** /auth/{mount_path}/config/tidy/roletag-blacklist | 
[**DeleteAwsConfigTidyRoletagDenylist**](Auth.md#DeleteAwsConfigTidyRoletagDenylist) | **Delete** /auth/{mount_path}/config/tidy/roletag-denylist | 
[**DeleteAwsIdentityAccesslistInstanceId**](Auth.md#DeleteAwsIdentityAccesslistInstanceId) | **Delete** /auth/{mount_path}/identity-accesslist/{instance_id} | 
[**DeleteAwsIdentityWhitelistInstanceId**](Auth.md#DeleteAwsIdentityWhitelistInstanceId) | **Delete** /auth/{mount_path}/identity-whitelist/{instance_id} | 
[**DeleteAwsRoleRole**](Auth.md#DeleteAwsRoleRole) | **Delete** /auth/{mount_path}/role/{role} | 
[**DeleteAwsRoletagBlacklistRoleTag**](Auth.md#DeleteAwsRoletagBlacklistRoleTag) | **Delete** /auth/{mount_path}/roletag-blacklist/{role_tag} | 
[**DeleteAwsRoletagDenylistRoleTag**](Auth.md#DeleteAwsRoletagDenylistRoleTag) | **Delete** /auth/{mount_path}/roletag-denylist/{role_tag} | 
[**DeleteCertCertsName**](Auth.md#DeleteCertCertsName) | **Delete** /auth/{mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**DeleteCertCrlsName**](Auth.md#DeleteCertCrlsName) | **Delete** /auth/{mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**DeleteCfRolesRole**](Auth.md#DeleteCfRolesRole) | **Delete** /auth/{mount_path}/roles/{role} | 
[**DeleteGithubMapTeamsKey**](Auth.md#DeleteGithubMapTeamsKey) | **Delete** /auth/{mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**DeleteGithubMapUsersKey**](Auth.md#DeleteGithubMapUsersKey) | **Delete** /auth/{mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**DeleteKerberosGroupsName**](Auth.md#DeleteKerberosGroupsName) | **Delete** /auth/{mount_path}/groups/{name} | 
[**DeleteOidcRoleName**](Auth.md#DeleteOidcRoleName) | **Delete** /auth/{mount_path}/role/{name} | Delete an existing role.
[**DeleteRadiusUsersName**](Auth.md#DeleteRadiusUsersName) | **Delete** /auth/{mount_path}/users/{name} | Manage users allowed to authenticate.
[**DeleteTokenRolesRoleName**](Auth.md#DeleteTokenRolesRoleName) | **Delete** /auth/{mount_path}/roles/{role_name} | 
[**DeleteUserpassUsersUsername**](Auth.md#DeleteUserpassUsersUsername) | **Delete** /auth/{mount_path}/users/{username} | Manage users allowed to authenticate.
[**ListApproleRoleRoleNameSecretId**](Auth.md#ListApproleRoleRoleNameSecretId) | **Get** /auth/{mount_path}/role/{role_name}/secret-id | Generate a SecretID against this role.
[**ListAwsConfigCertificates**](Auth.md#ListAwsConfigCertificates) | **Get** /auth/{mount_path}/config/certificates | 
[**ListAwsConfigSts**](Auth.md#ListAwsConfigSts) | **Get** /auth/{mount_path}/config/sts | 
[**ListAwsIdentityAccesslist**](Auth.md#ListAwsIdentityAccesslist) | **Get** /auth/{mount_path}/identity-accesslist | 
[**ListAwsIdentityWhitelist**](Auth.md#ListAwsIdentityWhitelist) | **Get** /auth/{mount_path}/identity-whitelist | 
[**ListAwsRoletagBlacklist**](Auth.md#ListAwsRoletagBlacklist) | **Get** /auth/{mount_path}/roletag-blacklist | 
[**ListAwsRoletagDenylist**](Auth.md#ListAwsRoletagDenylist) | **Get** /auth/{mount_path}/roletag-denylist | 
[**ListCertCerts**](Auth.md#ListCertCerts) | **Get** /auth/{mount_path}/certs | Manage trusted certificates used for authentication.
[**ListGcpRoles**](Auth.md#ListGcpRoles) | **Get** /auth/{mount_path}/roles | Lists all the roles that are registered with Vault.
[**ListKerberosGroups**](Auth.md#ListKerberosGroups) | **Get** /auth/{mount_path}/groups | 
[**ListOidcRole**](Auth.md#ListOidcRole) | **Get** /auth/{mount_path}/role | Lists all the roles registered with the backend.
[**ListTokenAccessors**](Auth.md#ListTokenAccessors) | **Get** /auth/{mount_path}/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#39;sudo&#39; capability in addition to &#39;list&#39;.
[**ListUserpassUsers**](Auth.md#ListUserpassUsers) | **Get** /auth/{mount_path}/users | Manage users allowed to authenticate.
[**ReadApproleRoleRoleName**](Auth.md#ReadApproleRoleRoleName) | **Get** /auth/{mount_path}/role/{role_name} | Register an role with the backend.
[**ReadApproleRoleRoleNameBindSecretId**](Auth.md#ReadApproleRoleRoleNameBindSecretId) | **Get** /auth/{mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**ReadApproleRoleRoleNameBoundCidrList**](Auth.md#ReadApproleRoleRoleNameBoundCidrList) | **Get** /auth/{mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**ReadApproleRoleRoleNameLocalSecretIds**](Auth.md#ReadApproleRoleRoleNameLocalSecretIds) | **Get** /auth/{mount_path}/role/{role_name}/local-secret-ids | Enables cluster local secret IDs
[**ReadApproleRoleRoleNamePeriod**](Auth.md#ReadApproleRoleRoleNamePeriod) | **Get** /auth/{mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**ReadApproleRoleRoleNamePolicies**](Auth.md#ReadApproleRoleRoleNamePolicies) | **Get** /auth/{mount_path}/role/{role_name}/policies | Policies of the role.
[**ReadApproleRoleRoleNameRoleId**](Auth.md#ReadApproleRoleRoleNameRoleId) | **Get** /auth/{mount_path}/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**ReadApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#ReadApproleRoleRoleNameSecretIdBoundCidrs) | **Get** /auth/{mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**ReadApproleRoleRoleNameSecretIdNumUses**](Auth.md#ReadApproleRoleRoleNameSecretIdNumUses) | **Get** /auth/{mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**ReadApproleRoleRoleNameSecretIdTtl**](Auth.md#ReadApproleRoleRoleNameSecretIdTtl) | **Get** /auth/{mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**ReadApproleRoleRoleNameTokenBoundCidrs**](Auth.md#ReadApproleRoleRoleNameTokenBoundCidrs) | **Get** /auth/{mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**ReadApproleRoleRoleNameTokenMaxTtl**](Auth.md#ReadApproleRoleRoleNameTokenMaxTtl) | **Get** /auth/{mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**ReadApproleRoleRoleNameTokenNumUses**](Auth.md#ReadApproleRoleRoleNameTokenNumUses) | **Get** /auth/{mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**ReadApproleRoleRoleNameTokenTtl**](Auth.md#ReadApproleRoleRoleNameTokenTtl) | **Get** /auth/{mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**ReadAwsConfigCertificateCertName**](Auth.md#ReadAwsConfigCertificateCertName) | **Get** /auth/{mount_path}/config/certificate/{cert_name} | 
[**ReadAwsConfigClient**](Auth.md#ReadAwsConfigClient) | **Get** /auth/{mount_path}/config/client | 
[**ReadAwsConfigIdentity**](Auth.md#ReadAwsConfigIdentity) | **Get** /auth/{mount_path}/config/identity | 
[**ReadAwsConfigStsAccountId**](Auth.md#ReadAwsConfigStsAccountId) | **Get** /auth/{mount_path}/config/sts/{account_id} | 
[**ReadAwsConfigTidyIdentityAccesslist**](Auth.md#ReadAwsConfigTidyIdentityAccesslist) | **Get** /auth/{mount_path}/config/tidy/identity-accesslist | 
[**ReadAwsConfigTidyIdentityWhitelist**](Auth.md#ReadAwsConfigTidyIdentityWhitelist) | **Get** /auth/{mount_path}/config/tidy/identity-whitelist | 
[**ReadAwsConfigTidyRoletagBlacklist**](Auth.md#ReadAwsConfigTidyRoletagBlacklist) | **Get** /auth/{mount_path}/config/tidy/roletag-blacklist | 
[**ReadAwsConfigTidyRoletagDenylist**](Auth.md#ReadAwsConfigTidyRoletagDenylist) | **Get** /auth/{mount_path}/config/tidy/roletag-denylist | 
[**ReadAwsIdentityAccesslistInstanceId**](Auth.md#ReadAwsIdentityAccesslistInstanceId) | **Get** /auth/{mount_path}/identity-accesslist/{instance_id} | 
[**ReadAwsIdentityWhitelistInstanceId**](Auth.md#ReadAwsIdentityWhitelistInstanceId) | **Get** /auth/{mount_path}/identity-whitelist/{instance_id} | 
[**ReadAwsRoleRole**](Auth.md#ReadAwsRoleRole) | **Get** /auth/{mount_path}/role/{role} | 
[**ReadAwsRoletagBlacklistRoleTag**](Auth.md#ReadAwsRoletagBlacklistRoleTag) | **Get** /auth/{mount_path}/roletag-blacklist/{role_tag} | 
[**ReadAwsRoletagDenylistRoleTag**](Auth.md#ReadAwsRoletagDenylistRoleTag) | **Get** /auth/{mount_path}/roletag-denylist/{role_tag} | 
[**ReadCertCertsName**](Auth.md#ReadCertCertsName) | **Get** /auth/{mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**ReadCertCrlsName**](Auth.md#ReadCertCrlsName) | **Get** /auth/{mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**ReadCfRolesRole**](Auth.md#ReadCfRolesRole) | **Get** /auth/{mount_path}/roles/{role} | 
[**ReadGithubMapTeams**](Auth.md#ReadGithubMapTeams) | **Get** /auth/{mount_path}/map/teams | Read mappings for teams
[**ReadGithubMapTeamsKey**](Auth.md#ReadGithubMapTeamsKey) | **Get** /auth/{mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**ReadGithubMapUsers**](Auth.md#ReadGithubMapUsers) | **Get** /auth/{mount_path}/map/users | Read mappings for users
[**ReadGithubMapUsersKey**](Auth.md#ReadGithubMapUsersKey) | **Get** /auth/{mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**ReadKerberosConfigLdap**](Auth.md#ReadKerberosConfigLdap) | **Get** /auth/{mount_path}/config/ldap | 
[**ReadKerberosGroupsName**](Auth.md#ReadKerberosGroupsName) | **Get** /auth/{mount_path}/groups/{name} | 
[**ReadOidcConfig**](Auth.md#ReadOidcConfig) | **Get** /auth/{mount_path}/config | Read the current JWT authentication backend configuration.
[**ReadOidcOidcCallback**](Auth.md#ReadOidcOidcCallback) | **Get** /auth/{mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**ReadOidcRoleName**](Auth.md#ReadOidcRoleName) | **Get** /auth/{mount_path}/role/{name} | Read an existing role.
[**ReadOktaVerifyNonce**](Auth.md#ReadOktaVerifyNonce) | **Get** /auth/{mount_path}/verify/{nonce} | 
[**ReadRadiusUsersName**](Auth.md#ReadRadiusUsersName) | **Get** /auth/{mount_path}/users/{name} | Manage users allowed to authenticate.
[**ReadTokenLookup**](Auth.md#ReadTokenLookup) | **Get** /auth/{mount_path}/lookup | This endpoint will lookup a token and its properties.
[**ReadTokenLookupSelf**](Auth.md#ReadTokenLookupSelf) | **Get** /auth/{mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**ReadTokenRolesRoleName**](Auth.md#ReadTokenRolesRoleName) | **Get** /auth/{mount_path}/roles/{role_name} | 
[**ReadUserpassUsersUsername**](Auth.md#ReadUserpassUsersUsername) | **Get** /auth/{mount_path}/users/{username} | Manage users allowed to authenticate.
[**UpdateApproleRoleRoleName**](Auth.md#UpdateApproleRoleRoleName) | **Post** /auth/{mount_path}/role/{role_name} | Register an role with the backend.
[**UpdateApproleRoleRoleNameBindSecretId**](Auth.md#UpdateApproleRoleRoleNameBindSecretId) | **Post** /auth/{mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**UpdateApproleRoleRoleNameBoundCidrList**](Auth.md#UpdateApproleRoleRoleNameBoundCidrList) | **Post** /auth/{mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**UpdateApproleRoleRoleNameCustomSecretId**](Auth.md#UpdateApproleRoleRoleNameCustomSecretId) | **Post** /auth/{mount_path}/role/{role_name}/custom-secret-id | Assign a SecretID of choice against the role.
[**UpdateApproleRoleRoleNamePeriod**](Auth.md#UpdateApproleRoleRoleNamePeriod) | **Post** /auth/{mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**UpdateApproleRoleRoleNamePolicies**](Auth.md#UpdateApproleRoleRoleNamePolicies) | **Post** /auth/{mount_path}/role/{role_name}/policies | Policies of the role.
[**UpdateApproleRoleRoleNameRoleId**](Auth.md#UpdateApproleRoleRoleNameRoleId) | **Post** /auth/{mount_path}/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**UpdateApproleRoleRoleNameSecretId**](Auth.md#UpdateApproleRoleRoleNameSecretId) | **Post** /auth/{mount_path}/role/{role_name}/secret-id | Generate a SecretID against this role.
[**UpdateApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#UpdateApproleRoleRoleNameSecretIdAccessorDestroy) | **Post** /auth/{mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**UpdateApproleRoleRoleNameSecretIdAccessorLookup**](Auth.md#UpdateApproleRoleRoleNameSecretIdAccessorLookup) | **Post** /auth/{mount_path}/role/{role_name}/secret-id-accessor/lookup | 
[**UpdateApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#UpdateApproleRoleRoleNameSecretIdBoundCidrs) | **Post** /auth/{mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**UpdateApproleRoleRoleNameSecretIdDestroy**](Auth.md#UpdateApproleRoleRoleNameSecretIdDestroy) | **Post** /auth/{mount_path}/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**UpdateApproleRoleRoleNameSecretIdLookup**](Auth.md#UpdateApproleRoleRoleNameSecretIdLookup) | **Post** /auth/{mount_path}/role/{role_name}/secret-id/lookup | Read the properties of an issued secret_id
[**UpdateApproleRoleRoleNameSecretIdNumUses**](Auth.md#UpdateApproleRoleRoleNameSecretIdNumUses) | **Post** /auth/{mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**UpdateApproleRoleRoleNameSecretIdTtl**](Auth.md#UpdateApproleRoleRoleNameSecretIdTtl) | **Post** /auth/{mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**UpdateApproleRoleRoleNameTokenBoundCidrs**](Auth.md#UpdateApproleRoleRoleNameTokenBoundCidrs) | **Post** /auth/{mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**UpdateApproleRoleRoleNameTokenMaxTtl**](Auth.md#UpdateApproleRoleRoleNameTokenMaxTtl) | **Post** /auth/{mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**UpdateApproleRoleRoleNameTokenNumUses**](Auth.md#UpdateApproleRoleRoleNameTokenNumUses) | **Post** /auth/{mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**UpdateApproleRoleRoleNameTokenTtl**](Auth.md#UpdateApproleRoleRoleNameTokenTtl) | **Post** /auth/{mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**UpdateApproleTidySecretId**](Auth.md#UpdateApproleTidySecretId) | **Post** /auth/{mount_path}/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
[**UpdateAwsConfigCertificateCertName**](Auth.md#UpdateAwsConfigCertificateCertName) | **Post** /auth/{mount_path}/config/certificate/{cert_name} | 
[**UpdateAwsConfigClient**](Auth.md#UpdateAwsConfigClient) | **Post** /auth/{mount_path}/config/client | 
[**UpdateAwsConfigIdentity**](Auth.md#UpdateAwsConfigIdentity) | **Post** /auth/{mount_path}/config/identity | 
[**UpdateAwsConfigRotateRoot**](Auth.md#UpdateAwsConfigRotateRoot) | **Post** /auth/{mount_path}/config/rotate-root | 
[**UpdateAwsConfigStsAccountId**](Auth.md#UpdateAwsConfigStsAccountId) | **Post** /auth/{mount_path}/config/sts/{account_id} | 
[**UpdateAwsConfigTidyIdentityAccesslist**](Auth.md#UpdateAwsConfigTidyIdentityAccesslist) | **Post** /auth/{mount_path}/config/tidy/identity-accesslist | 
[**UpdateAwsConfigTidyIdentityWhitelist**](Auth.md#UpdateAwsConfigTidyIdentityWhitelist) | **Post** /auth/{mount_path}/config/tidy/identity-whitelist | 
[**UpdateAwsConfigTidyRoletagBlacklist**](Auth.md#UpdateAwsConfigTidyRoletagBlacklist) | **Post** /auth/{mount_path}/config/tidy/roletag-blacklist | 
[**UpdateAwsConfigTidyRoletagDenylist**](Auth.md#UpdateAwsConfigTidyRoletagDenylist) | **Post** /auth/{mount_path}/config/tidy/roletag-denylist | 
[**UpdateAwsRoleRole**](Auth.md#UpdateAwsRoleRole) | **Post** /auth/{mount_path}/role/{role} | 
[**UpdateAwsRoleRoleTag**](Auth.md#UpdateAwsRoleRoleTag) | **Post** /auth/{mount_path}/role/{role}/tag | 
[**UpdateAwsRoletagBlacklistRoleTag**](Auth.md#UpdateAwsRoletagBlacklistRoleTag) | **Post** /auth/{mount_path}/roletag-blacklist/{role_tag} | 
[**UpdateAwsRoletagDenylistRoleTag**](Auth.md#UpdateAwsRoletagDenylistRoleTag) | **Post** /auth/{mount_path}/roletag-denylist/{role_tag} | 
[**UpdateAwsTidyIdentityAccesslist**](Auth.md#UpdateAwsTidyIdentityAccesslist) | **Post** /auth/{mount_path}/tidy/identity-accesslist | 
[**UpdateAwsTidyIdentityWhitelist**](Auth.md#UpdateAwsTidyIdentityWhitelist) | **Post** /auth/{mount_path}/tidy/identity-whitelist | 
[**UpdateAwsTidyRoletagBlacklist**](Auth.md#UpdateAwsTidyRoletagBlacklist) | **Post** /auth/{mount_path}/tidy/roletag-blacklist | 
[**UpdateAwsTidyRoletagDenylist**](Auth.md#UpdateAwsTidyRoletagDenylist) | **Post** /auth/{mount_path}/tidy/roletag-denylist | 
[**UpdateCertCertsName**](Auth.md#UpdateCertCertsName) | **Post** /auth/{mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**UpdateCertCrlsName**](Auth.md#UpdateCertCrlsName) | **Post** /auth/{mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**UpdateCfRolesRole**](Auth.md#UpdateCfRolesRole) | **Post** /auth/{mount_path}/roles/{role} | 
[**UpdateGcpRoleNameLabels**](Auth.md#UpdateGcpRoleNameLabels) | **Post** /auth/{mount_path}/role/{name}/labels | Add or remove labels for an existing &#39;gce&#39; role
[**UpdateGcpRoleNameServiceAccounts**](Auth.md#UpdateGcpRoleNameServiceAccounts) | **Post** /auth/{mount_path}/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**UpdateGithubMapTeamsKey**](Auth.md#UpdateGithubMapTeamsKey) | **Post** /auth/{mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**UpdateGithubMapUsersKey**](Auth.md#UpdateGithubMapUsersKey) | **Post** /auth/{mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**UpdateKerberosConfigLdap**](Auth.md#UpdateKerberosConfigLdap) | **Post** /auth/{mount_path}/config/ldap | 
[**UpdateKerberosGroupsName**](Auth.md#UpdateKerberosGroupsName) | **Post** /auth/{mount_path}/groups/{name} | 
[**UpdateOciLoginRole**](Auth.md#UpdateOciLoginRole) | **Post** /auth/{mount_path}/login/{role} | Authenticates to Vault using OCI credentials
[**UpdateOidcConfig**](Auth.md#UpdateOidcConfig) | **Post** /auth/{mount_path}/config | Configure the JWT authentication backend.
[**UpdateOidcLogin**](Auth.md#UpdateOidcLogin) | **Post** /auth/{mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**UpdateOidcOidcAuthUrl**](Auth.md#UpdateOidcOidcAuthUrl) | **Post** /auth/{mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**UpdateOidcOidcCallback**](Auth.md#UpdateOidcOidcCallback) | **Post** /auth/{mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**UpdateOidcRoleName**](Auth.md#UpdateOidcRoleName) | **Post** /auth/{mount_path}/role/{name} | Register an role with the backend.
[**UpdateRadiusLoginUrlusername**](Auth.md#UpdateRadiusLoginUrlusername) | **Post** /auth/{mount_path}/login/{urlusername} | Log in with a username and password.
[**UpdateRadiusUsersName**](Auth.md#UpdateRadiusUsersName) | **Post** /auth/{mount_path}/users/{name} | Manage users allowed to authenticate.
[**UpdateTokenCreate**](Auth.md#UpdateTokenCreate) | **Post** /auth/{mount_path}/create | The token create path is used to create new tokens.
[**UpdateTokenCreateOrphan**](Auth.md#UpdateTokenCreateOrphan) | **Post** /auth/{mount_path}/create-orphan | The token create path is used to create new orphan tokens.
[**UpdateTokenCreateRoleName**](Auth.md#UpdateTokenCreateRoleName) | **Post** /auth/{mount_path}/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**UpdateTokenLookup**](Auth.md#UpdateTokenLookup) | **Post** /auth/{mount_path}/lookup | This endpoint will lookup a token and its properties.
[**UpdateTokenLookupAccessor**](Auth.md#UpdateTokenLookupAccessor) | **Post** /auth/{mount_path}/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**UpdateTokenLookupSelf**](Auth.md#UpdateTokenLookupSelf) | **Post** /auth/{mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**UpdateTokenRenew**](Auth.md#UpdateTokenRenew) | **Post** /auth/{mount_path}/renew | This endpoint will renew the given token and prevent expiration.
[**UpdateTokenRenewAccessor**](Auth.md#UpdateTokenRenewAccessor) | **Post** /auth/{mount_path}/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**UpdateTokenRenewSelf**](Auth.md#UpdateTokenRenewSelf) | **Post** /auth/{mount_path}/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**UpdateTokenRevoke**](Auth.md#UpdateTokenRevoke) | **Post** /auth/{mount_path}/revoke | This endpoint will delete the given token and all of its child tokens.
[**UpdateTokenRevokeAccessor**](Auth.md#UpdateTokenRevokeAccessor) | **Post** /auth/{mount_path}/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**UpdateTokenRevokeOrphan**](Auth.md#UpdateTokenRevokeOrphan) | **Post** /auth/{mount_path}/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**UpdateTokenRevokeSelf**](Auth.md#UpdateTokenRevokeSelf) | **Post** /auth/{mount_path}/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**UpdateTokenRolesRoleName**](Auth.md#UpdateTokenRolesRoleName) | **Post** /auth/{mount_path}/roles/{role_name} | 
[**UpdateTokenTidy**](Auth.md#UpdateTokenTidy) | **Post** /auth/{mount_path}/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**UpdateUserpassLoginUsername**](Auth.md#UpdateUserpassLoginUsername) | **Post** /auth/{mount_path}/login/{username} | Log in with a username and password.
[**UpdateUserpassUsersUsername**](Auth.md#UpdateUserpassUsersUsername) | **Post** /auth/{mount_path}/users/{username} | Manage users allowed to authenticate.
[**UpdateUserpassUsersUsernamePassword**](Auth.md#UpdateUserpassUsersUsernamePassword) | **Post** /auth/{mount_path}/users/{username}/password | Reset user&#39;s password.
[**UpdateUserpassUsersUsernamePolicies**](Auth.md#UpdateUserpassUsersUsernamePolicies) | **Post** /auth/{mount_path}/users/{username}/policies | Update the policies associated with the username.



## DeleteApproleRoleRoleName

> DeleteApproleRoleRoleName(ctx, mountPath, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleName(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameBindSecretId

> DeleteApproleRoleRoleNameBindSecretId(ctx, mountPath, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameBindSecretId(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameBoundCidrList

> DeleteApproleRoleRoleNameBoundCidrList(ctx, mountPath, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameBoundCidrList(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNamePeriod

> DeleteApproleRoleRoleNamePeriod(ctx, mountPath, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNamePeriod(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNamePolicies

> DeleteApproleRoleRoleNamePolicies(ctx, mountPath, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNamePolicies(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameSecretIdAccessorDestroy

> DeleteApproleRoleRoleNameSecretIdAccessorDestroy(ctx, mountPath, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameSecretIdBoundCidrs

> DeleteApproleRoleRoleNameSecretIdBoundCidrs(ctx, mountPath, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameSecretIdDestroy

> DeleteApproleRoleRoleNameSecretIdDestroy(ctx, mountPath, roleName).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameSecretIdDestroy(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameSecretIdNumUses

> DeleteApproleRoleRoleNameSecretIdNumUses(ctx, mountPath, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameSecretIdNumUses(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameSecretIdTtl

> DeleteApproleRoleRoleNameSecretIdTtl(ctx, mountPath, roleName).Execute()

Duration in seconds of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameSecretIdTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameTokenBoundCidrs

> DeleteApproleRoleRoleNameTokenBoundCidrs(ctx, mountPath, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameTokenBoundCidrs(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameTokenMaxTtl

> DeleteApproleRoleRoleNameTokenMaxTtl(ctx, mountPath, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameTokenMaxTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameTokenNumUses

> DeleteApproleRoleRoleNameTokenNumUses(ctx, mountPath, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameTokenNumUses(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteApproleRoleRoleNameTokenTtl

> DeleteApproleRoleRoleNameTokenTtl(ctx, mountPath, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteApproleRoleRoleNameTokenTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsConfigCertificateCertName

> DeleteAwsConfigCertificateCertName(ctx, certName, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigCertificateCertName(context.Background(), certName, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certName** | **string** | Name of the certificate. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsConfigClient

> DeleteAwsConfigClient(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigClient(context.Background(), mountPath)
	if err != nil {
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


## DeleteAwsConfigStsAccountId

> DeleteAwsConfigStsAccountId(ctx, accountId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigStsAccountId(context.Background(), accountId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsConfigTidyIdentityAccesslist

> DeleteAwsConfigTidyIdentityAccesslist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigTidyIdentityAccesslist(context.Background(), mountPath)
	if err != nil {
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


## DeleteAwsConfigTidyIdentityWhitelist

> DeleteAwsConfigTidyIdentityWhitelist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigTidyIdentityWhitelist(context.Background(), mountPath)
	if err != nil {
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


## DeleteAwsConfigTidyRoletagBlacklist

> DeleteAwsConfigTidyRoletagBlacklist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigTidyRoletagBlacklist(context.Background(), mountPath)
	if err != nil {
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


## DeleteAwsConfigTidyRoletagDenylist

> DeleteAwsConfigTidyRoletagDenylist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAwsConfigTidyRoletagDenylist(context.Background(), mountPath)
	if err != nil {
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


## DeleteAwsIdentityAccesslistInstanceId

> DeleteAwsIdentityAccesslistInstanceId(ctx, instanceId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsIdentityAccesslistInstanceId(context.Background(), instanceId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsIdentityWhitelistInstanceId

> DeleteAwsIdentityWhitelistInstanceId(ctx, instanceId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsIdentityWhitelistInstanceId(context.Background(), instanceId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsRoleRole

> DeleteAwsRoleRole(ctx, mountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsRoleRole(context.Background(), mountPath, role)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsRoletagBlacklistRoleTag

> DeleteAwsRoletagBlacklistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsRoletagBlacklistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAwsRoletagDenylistRoleTag

> DeleteAwsRoletagDenylistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAwsRoletagDenylistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteCertCertsName

> DeleteCertCertsName(ctx, mountPath, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.DeleteCertCertsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteCertCrlsName

> DeleteCertCrlsName(ctx, mountPath, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.DeleteCertCrlsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteCfRolesRole

> DeleteCfRolesRole(ctx, mountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.DeleteCfRolesRole(context.Background(), mountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGithubMapTeamsKey

> DeleteGithubMapTeamsKey(ctx, key, mountPath).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.DeleteGithubMapTeamsKey(context.Background(), key, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the teams mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteGithubMapUsersKey

> DeleteGithubMapUsersKey(ctx, key, mountPath).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.DeleteGithubMapUsersKey(context.Background(), key, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the users mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteKerberosGroupsName

> DeleteKerberosGroupsName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.DeleteKerberosGroupsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteOidcRoleName

> DeleteOidcRoleName(ctx, mountPath, name).Execute()

Delete an existing role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.DeleteOidcRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteRadiusUsersName

> DeleteRadiusUsersName(ctx, mountPath, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	resp, err := client.WithToken("my-token").Auth.DeleteRadiusUsersName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteTokenRolesRoleName

> DeleteTokenRolesRoleName(ctx, mountPath, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.DeleteTokenRolesRoleName(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteUserpassUsersUsername

> DeleteUserpassUsersUsername(ctx, mountPath, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	resp, err := client.WithToken("my-token").Auth.DeleteUserpassUsersUsername(context.Background(), mountPath, username)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListApproleRoleRoleNameSecretId

> ListApproleRoleRoleNameSecretId(ctx, mountPath, roleName).List(list).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListApproleRoleRoleNameSecretId(context.Background(), mountPath, roleName, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListAwsConfigCertificates

> ListAwsConfigCertificates(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsConfigCertificates(context.Background(), mountPath, list)
	if err != nil {
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


## ListAwsConfigSts

> ListAwsConfigSts(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsConfigSts(context.Background(), mountPath, list)
	if err != nil {
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


## ListAwsIdentityAccesslist

> ListAwsIdentityAccesslist(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsIdentityAccesslist(context.Background(), mountPath, list)
	if err != nil {
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


## ListAwsIdentityWhitelist

> ListAwsIdentityWhitelist(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsIdentityWhitelist(context.Background(), mountPath, list)
	if err != nil {
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


## ListAwsRoletagBlacklist

> ListAwsRoletagBlacklist(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsRoletagBlacklist(context.Background(), mountPath, list)
	if err != nil {
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


## ListAwsRoletagDenylist

> ListAwsRoletagDenylist(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListAwsRoletagDenylist(context.Background(), mountPath, list)
	if err != nil {
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


## ListCertCerts

> ListCertCerts(ctx, mountPath).List(list).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListCertCerts(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListGcpRoles

> ListGcpRoles(ctx, mountPath).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.ListGcpRoles(context.Background(), mountPath, list)
	if err != nil {
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


## ListKerberosGroups

> ListKerberosGroups(ctx, mountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListKerberosGroups(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListOidcRole

> ListOidcRole(ctx, mountPath).List(list).Execute()

Lists all the roles registered with the backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListOidcRole(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListTokenAccessors

> ListTokenAccessors(ctx, mountPath).List(list).Execute()

List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListTokenAccessors(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListUserpassUsers

> ListUserpassUsers(ctx, mountPath).List(list).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ListUserpassUsers(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleName

> ReadApproleRoleRoleName(ctx, mountPath, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleName(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameBindSecretId

> ReadApproleRoleRoleNameBindSecretId(ctx, mountPath, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameBindSecretId(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameBoundCidrList

> ReadApproleRoleRoleNameBoundCidrList(ctx, mountPath, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameBoundCidrList(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameLocalSecretIds

> ReadApproleRoleRoleNameLocalSecretIds(ctx, mountPath, roleName).Execute()

Enables cluster local secret IDs

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameLocalSecretIds(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNamePeriod

> ReadApproleRoleRoleNamePeriod(ctx, mountPath, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNamePeriod(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNamePolicies

> ReadApproleRoleRoleNamePolicies(ctx, mountPath, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNamePolicies(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameRoleId

> ReadApproleRoleRoleNameRoleId(ctx, mountPath, roleName).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameRoleId(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameSecretIdBoundCidrs

> ReadApproleRoleRoleNameSecretIdBoundCidrs(ctx, mountPath, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameSecretIdNumUses

> ReadApproleRoleRoleNameSecretIdNumUses(ctx, mountPath, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameSecretIdNumUses(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameSecretIdTtl

> ReadApproleRoleRoleNameSecretIdTtl(ctx, mountPath, roleName).Execute()

Duration in seconds of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameSecretIdTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameTokenBoundCidrs

> ReadApproleRoleRoleNameTokenBoundCidrs(ctx, mountPath, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameTokenBoundCidrs(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameTokenMaxTtl

> ReadApproleRoleRoleNameTokenMaxTtl(ctx, mountPath, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameTokenMaxTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameTokenNumUses

> ReadApproleRoleRoleNameTokenNumUses(ctx, mountPath, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameTokenNumUses(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadApproleRoleRoleNameTokenTtl

> ReadApproleRoleRoleNameTokenTtl(ctx, mountPath, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.ReadApproleRoleRoleNameTokenTtl(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsConfigCertificateCertName

> ReadAwsConfigCertificateCertName(ctx, certName, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigCertificateCertName(context.Background(), certName, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certName** | **string** | Name of the certificate. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsConfigClient

> ReadAwsConfigClient(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigClient(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsConfigIdentity

> ReadAwsConfigIdentity(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigIdentity(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsConfigStsAccountId

> ReadAwsConfigStsAccountId(ctx, accountId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigStsAccountId(context.Background(), accountId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsConfigTidyIdentityAccesslist

> ReadAwsConfigTidyIdentityAccesslist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigTidyIdentityAccesslist(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsConfigTidyIdentityWhitelist

> ReadAwsConfigTidyIdentityWhitelist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigTidyIdentityWhitelist(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsConfigTidyRoletagBlacklist

> ReadAwsConfigTidyRoletagBlacklist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigTidyRoletagBlacklist(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsConfigTidyRoletagDenylist

> ReadAwsConfigTidyRoletagDenylist(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.ReadAwsConfigTidyRoletagDenylist(context.Background(), mountPath)
	if err != nil {
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


## ReadAwsIdentityAccesslistInstanceId

> ReadAwsIdentityAccesslistInstanceId(ctx, instanceId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsIdentityAccesslistInstanceId(context.Background(), instanceId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsIdentityWhitelistInstanceId

> ReadAwsIdentityWhitelistInstanceId(ctx, instanceId, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsIdentityWhitelistInstanceId(context.Background(), instanceId, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsRoleRole

> ReadAwsRoleRole(ctx, mountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsRoleRole(context.Background(), mountPath, role)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsRoletagBlacklistRoleTag

> ReadAwsRoletagBlacklistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsRoletagBlacklistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadAwsRoletagDenylistRoleTag

> ReadAwsRoletagDenylistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.ReadAwsRoletagDenylistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadCertCertsName

> ReadCertCertsName(ctx, mountPath, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.ReadCertCertsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadCertCrlsName

> ReadCertCrlsName(ctx, mountPath, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.ReadCertCrlsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadCfRolesRole

> ReadCfRolesRole(ctx, mountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.ReadCfRolesRole(context.Background(), mountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGithubMapTeams

> ReadGithubMapTeams(ctx, mountPath).List(list).Execute()

Read mappings for teams

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ReadGithubMapTeams(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGithubMapTeamsKey

> ReadGithubMapTeamsKey(ctx, key, mountPath).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.ReadGithubMapTeamsKey(context.Background(), key, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the teams mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGithubMapUsers

> ReadGithubMapUsers(ctx, mountPath).List(list).Execute()

Read mappings for users

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.ReadGithubMapUsers(context.Background(), mountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadGithubMapUsersKey

> ReadGithubMapUsersKey(ctx, key, mountPath).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.ReadGithubMapUsersKey(context.Background(), key, mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the users mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadKerberosConfigLdap

> ReadKerberosConfigLdap(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.ReadKerberosConfigLdap(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadKerberosGroupsName

> ReadKerberosGroupsName(ctx, mountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.ReadKerberosGroupsName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadOidcConfig

> ReadOidcConfig(ctx, mountPath).Execute()

Read the current JWT authentication backend configuration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.ReadOidcConfig(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadOidcOidcCallback

> ReadOidcOidcCallback(ctx, mountPath).Execute()

Callback endpoint to complete an OIDC login.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.ReadOidcOidcCallback(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadOidcRoleName

> ReadOidcRoleName(ctx, mountPath, name).Execute()

Read an existing role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.ReadOidcRoleName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadOktaVerifyNonce

> ReadOktaVerifyNonce(ctx, mountPath, nonce).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	nonce := "nonce_example" // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.ReadOktaVerifyNonce(context.Background(), mountPath, nonce)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]
**nonce** | **string** | Nonce provided during a login request to retrieve the number verification challenge for the matching request. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadRadiusUsersName

> ReadRadiusUsersName(ctx, mountPath, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	resp, err := client.WithToken("my-token").Auth.ReadRadiusUsersName(context.Background(), mountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTokenLookup

> ReadTokenLookup(ctx, mountPath).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.ReadTokenLookup(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTokenLookupSelf

> ReadTokenLookupSelf(ctx, mountPath).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.ReadTokenLookupSelf(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadTokenRolesRoleName

> ReadTokenRolesRoleName(ctx, mountPath, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.ReadTokenRolesRoleName(context.Background(), mountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadUserpassUsersUsername

> ReadUserpassUsersUsername(ctx, mountPath, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	resp, err := client.WithToken("my-token").Auth.ReadUserpassUsersUsername(context.Background(), mountPath, username)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleName

> UpdateApproleRoleRoleName(ctx, mountPath, roleName).ApproleRoleRequest(approleRoleRequest).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleRequest := NewApproleRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleName(context.Background(), mountPath, roleName, approleRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleRequest** | [**ApproleRoleRequest**](ApproleRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameBindSecretId

> UpdateApproleRoleRoleNameBindSecretId(ctx, mountPath, roleName).ApproleRoleBindSecretIdRequest(approleRoleBindSecretIdRequest).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleBindSecretIdRequest := NewApproleRoleBindSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameBindSecretId(context.Background(), mountPath, roleName, approleRoleBindSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleBindSecretIdRequest** | [**ApproleRoleBindSecretIdRequest**](ApproleRoleBindSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameBoundCidrList

> UpdateApproleRoleRoleNameBoundCidrList(ctx, mountPath, roleName).ApproleRoleBoundCidrListRequest(approleRoleBoundCidrListRequest).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleBoundCidrListRequest := NewApproleRoleBoundCidrListRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameBoundCidrList(context.Background(), mountPath, roleName, approleRoleBoundCidrListRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleBoundCidrListRequest** | [**ApproleRoleBoundCidrListRequest**](ApproleRoleBoundCidrListRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameCustomSecretId

> UpdateApproleRoleRoleNameCustomSecretId(ctx, mountPath, roleName).ApproleRoleCustomSecretIdRequest(approleRoleCustomSecretIdRequest).Execute()

Assign a SecretID of choice against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleCustomSecretIdRequest := NewApproleRoleCustomSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameCustomSecretId(context.Background(), mountPath, roleName, approleRoleCustomSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleCustomSecretIdRequest** | [**ApproleRoleCustomSecretIdRequest**](ApproleRoleCustomSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNamePeriod

> UpdateApproleRoleRoleNamePeriod(ctx, mountPath, roleName).ApproleRolePeriodRequest(approleRolePeriodRequest).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRolePeriodRequest := NewApproleRolePeriodRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNamePeriod(context.Background(), mountPath, roleName, approleRolePeriodRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRolePeriodRequest** | [**ApproleRolePeriodRequest**](ApproleRolePeriodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNamePolicies

> UpdateApproleRoleRoleNamePolicies(ctx, mountPath, roleName).ApproleRolePoliciesRequest(approleRolePoliciesRequest).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRolePoliciesRequest := NewApproleRolePoliciesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNamePolicies(context.Background(), mountPath, roleName, approleRolePoliciesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRolePoliciesRequest** | [**ApproleRolePoliciesRequest**](ApproleRolePoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameRoleId

> UpdateApproleRoleRoleNameRoleId(ctx, mountPath, roleName).ApproleRoleRoleIdRequest(approleRoleRoleIdRequest).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleRoleIdRequest := NewApproleRoleRoleIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameRoleId(context.Background(), mountPath, roleName, approleRoleRoleIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleRoleIdRequest** | [**ApproleRoleRoleIdRequest**](ApproleRoleRoleIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretId

> UpdateApproleRoleRoleNameSecretId(ctx, mountPath, roleName).ApproleRoleSecretIdRequest(approleRoleSecretIdRequest).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdRequest := NewApproleRoleSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretId(context.Background(), mountPath, roleName, approleRoleSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdRequest** | [**ApproleRoleSecretIdRequest**](ApproleRoleSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdAccessorDestroy

> UpdateApproleRoleRoleNameSecretIdAccessorDestroy(ctx, mountPath, roleName).ApproleRoleSecretIdAccessorDestroyRequest(approleRoleSecretIdAccessorDestroyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdAccessorDestroyRequest := NewApproleRoleSecretIdAccessorDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), mountPath, roleName, approleRoleSecretIdAccessorDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdAccessorDestroyRequest** | [**ApproleRoleSecretIdAccessorDestroyRequest**](ApproleRoleSecretIdAccessorDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdAccessorLookup

> UpdateApproleRoleRoleNameSecretIdAccessorLookup(ctx, mountPath, roleName).ApproleRoleSecretIdAccessorLookupRequest(approleRoleSecretIdAccessorLookupRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdAccessorLookupRequest := NewApproleRoleSecretIdAccessorLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdAccessorLookup(context.Background(), mountPath, roleName, approleRoleSecretIdAccessorLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdAccessorLookupRequest** | [**ApproleRoleSecretIdAccessorLookupRequest**](ApproleRoleSecretIdAccessorLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdBoundCidrs

> UpdateApproleRoleRoleNameSecretIdBoundCidrs(ctx, mountPath, roleName).ApproleRoleSecretIdBoundCidrsRequest(approleRoleSecretIdBoundCidrsRequest).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdBoundCidrsRequest := NewApproleRoleSecretIdBoundCidrsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), mountPath, roleName, approleRoleSecretIdBoundCidrsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdBoundCidrsRequest** | [**ApproleRoleSecretIdBoundCidrsRequest**](ApproleRoleSecretIdBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdDestroy

> UpdateApproleRoleRoleNameSecretIdDestroy(ctx, mountPath, roleName).ApproleRoleSecretIdDestroyRequest(approleRoleSecretIdDestroyRequest).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdDestroyRequest := NewApproleRoleSecretIdDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdDestroy(context.Background(), mountPath, roleName, approleRoleSecretIdDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdDestroyRequest** | [**ApproleRoleSecretIdDestroyRequest**](ApproleRoleSecretIdDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdLookup

> UpdateApproleRoleRoleNameSecretIdLookup(ctx, mountPath, roleName).ApproleRoleSecretIdLookupRequest(approleRoleSecretIdLookupRequest).Execute()

Read the properties of an issued secret_id

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdLookupRequest := NewApproleRoleSecretIdLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdLookup(context.Background(), mountPath, roleName, approleRoleSecretIdLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdLookupRequest** | [**ApproleRoleSecretIdLookupRequest**](ApproleRoleSecretIdLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdNumUses

> UpdateApproleRoleRoleNameSecretIdNumUses(ctx, mountPath, roleName).ApproleRoleSecretIdNumUsesRequest(approleRoleSecretIdNumUsesRequest).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdNumUsesRequest := NewApproleRoleSecretIdNumUsesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdNumUses(context.Background(), mountPath, roleName, approleRoleSecretIdNumUsesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdNumUsesRequest** | [**ApproleRoleSecretIdNumUsesRequest**](ApproleRoleSecretIdNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameSecretIdTtl

> UpdateApproleRoleRoleNameSecretIdTtl(ctx, mountPath, roleName).ApproleRoleSecretIdTtlRequest(approleRoleSecretIdTtlRequest).Execute()

Duration in seconds of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdTtlRequest := NewApproleRoleSecretIdTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameSecretIdTtl(context.Background(), mountPath, roleName, approleRoleSecretIdTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleSecretIdTtlRequest** | [**ApproleRoleSecretIdTtlRequest**](ApproleRoleSecretIdTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameTokenBoundCidrs

> UpdateApproleRoleRoleNameTokenBoundCidrs(ctx, mountPath, roleName).ApproleRoleTokenBoundCidrsRequest(approleRoleTokenBoundCidrsRequest).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenBoundCidrsRequest := NewApproleRoleTokenBoundCidrsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameTokenBoundCidrs(context.Background(), mountPath, roleName, approleRoleTokenBoundCidrsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleTokenBoundCidrsRequest** | [**ApproleRoleTokenBoundCidrsRequest**](ApproleRoleTokenBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameTokenMaxTtl

> UpdateApproleRoleRoleNameTokenMaxTtl(ctx, mountPath, roleName).ApproleRoleTokenMaxTtlRequest(approleRoleTokenMaxTtlRequest).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenMaxTtlRequest := NewApproleRoleTokenMaxTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameTokenMaxTtl(context.Background(), mountPath, roleName, approleRoleTokenMaxTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleTokenMaxTtlRequest** | [**ApproleRoleTokenMaxTtlRequest**](ApproleRoleTokenMaxTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameTokenNumUses

> UpdateApproleRoleRoleNameTokenNumUses(ctx, mountPath, roleName).ApproleRoleTokenNumUsesRequest(approleRoleTokenNumUsesRequest).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenNumUsesRequest := NewApproleRoleTokenNumUsesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameTokenNumUses(context.Background(), mountPath, roleName, approleRoleTokenNumUsesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleTokenNumUsesRequest** | [**ApproleRoleTokenNumUsesRequest**](ApproleRoleTokenNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleRoleRoleNameTokenTtl

> UpdateApproleRoleRoleNameTokenTtl(ctx, mountPath, roleName).ApproleRoleTokenTtlRequest(approleRoleTokenTtlRequest).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenTtlRequest := NewApproleRoleTokenTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateApproleRoleRoleNameTokenTtl(context.Background(), mountPath, roleName, approleRoleTokenTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approleRoleTokenTtlRequest** | [**ApproleRoleTokenTtlRequest**](ApproleRoleTokenTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateApproleTidySecretId

> UpdateApproleTidySecretId(ctx, mountPath).Execute()

Trigger the clean-up of expired SecretID entries.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.UpdateApproleTidySecretId(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigCertificateCertName

> UpdateAwsConfigCertificateCertName(ctx, certName, mountPath).AwsConfigCertificateRequest(awsConfigCertificateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigCertificateRequest := NewAwsConfigCertificateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigCertificateCertName(context.Background(), certName, mountPath, awsConfigCertificateRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certName** | **string** | Name of the certificate. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsConfigCertificateRequest** | [**AwsConfigCertificateRequest**](AwsConfigCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigClient

> UpdateAwsConfigClient(ctx, mountPath).AwsConfigClientRequest(awsConfigClientRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigClientRequest := NewAwsConfigClientRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigClient(context.Background(), mountPath, awsConfigClientRequest)
	if err != nil {
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

 **awsConfigClientRequest** | [**AwsConfigClientRequest**](AwsConfigClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigIdentity

> UpdateAwsConfigIdentity(ctx, mountPath).AwsConfigIdentityRequest(awsConfigIdentityRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigIdentityRequest := NewAwsConfigIdentityRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigIdentity(context.Background(), mountPath, awsConfigIdentityRequest)
	if err != nil {
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

 **awsConfigIdentityRequest** | [**AwsConfigIdentityRequest**](AwsConfigIdentityRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigRotateRoot

> UpdateAwsConfigRotateRoot(ctx, mountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigRotateRoot(context.Background(), mountPath)
	if err != nil {
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


## UpdateAwsConfigStsAccountId

> UpdateAwsConfigStsAccountId(ctx, accountId, mountPath).AwsConfigStsRequest(awsConfigStsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigStsRequest := NewAwsConfigStsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigStsAccountId(context.Background(), accountId, mountPath, awsConfigStsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsConfigStsRequest** | [**AwsConfigStsRequest**](AwsConfigStsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigTidyIdentityAccesslist

> UpdateAwsConfigTidyIdentityAccesslist(ctx, mountPath).AwsConfigTidyIdentityAccesslistRequest(awsConfigTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyIdentityAccesslistRequest := NewAwsConfigTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigTidyIdentityAccesslist(context.Background(), mountPath, awsConfigTidyIdentityAccesslistRequest)
	if err != nil {
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

 **awsConfigTidyIdentityAccesslistRequest** | [**AwsConfigTidyIdentityAccesslistRequest**](AwsConfigTidyIdentityAccesslistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigTidyIdentityWhitelist

> UpdateAwsConfigTidyIdentityWhitelist(ctx, mountPath).AwsConfigTidyIdentityWhitelistRequest(awsConfigTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyIdentityWhitelistRequest := NewAwsConfigTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigTidyIdentityWhitelist(context.Background(), mountPath, awsConfigTidyIdentityWhitelistRequest)
	if err != nil {
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

 **awsConfigTidyIdentityWhitelistRequest** | [**AwsConfigTidyIdentityWhitelistRequest**](AwsConfigTidyIdentityWhitelistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigTidyRoletagBlacklist

> UpdateAwsConfigTidyRoletagBlacklist(ctx, mountPath).AwsConfigTidyRoletagBlacklistRequest(awsConfigTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyRoletagBlacklistRequest := NewAwsConfigTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigTidyRoletagBlacklist(context.Background(), mountPath, awsConfigTidyRoletagBlacklistRequest)
	if err != nil {
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

 **awsConfigTidyRoletagBlacklistRequest** | [**AwsConfigTidyRoletagBlacklistRequest**](AwsConfigTidyRoletagBlacklistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsConfigTidyRoletagDenylist

> UpdateAwsConfigTidyRoletagDenylist(ctx, mountPath).AwsConfigTidyRoletagDenylistRequest(awsConfigTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyRoletagDenylistRequest := NewAwsConfigTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsConfigTidyRoletagDenylist(context.Background(), mountPath, awsConfigTidyRoletagDenylistRequest)
	if err != nil {
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

 **awsConfigTidyRoletagDenylistRequest** | [**AwsConfigTidyRoletagDenylistRequest**](AwsConfigTidyRoletagDenylistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsRoleRole

> UpdateAwsRoleRole(ctx, mountPath, role).AwsRoleRequest(awsRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRoleRequest := NewAwsRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsRoleRole(context.Background(), mountPath, role, awsRoleRequest)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsRoleRequest** | [**AwsRoleRequest**](AwsRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsRoleRoleTag

> UpdateAwsRoleRoleTag(ctx, mountPath, role).AwsRoleTagRequest(awsRoleTagRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRoleTagRequest := NewAwsRoleTagRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsRoleRoleTag(context.Background(), mountPath, role, awsRoleTagRequest)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsRoleTagRequest** | [**AwsRoleTagRequest**](AwsRoleTagRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsRoletagBlacklistRoleTag

> UpdateAwsRoletagBlacklistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.UpdateAwsRoletagBlacklistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsRoletagDenylistRoleTag

> UpdateAwsRoletagDenylistRoleTag(ctx, mountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.UpdateAwsRoletagDenylistRoleTag(context.Background(), mountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsTidyIdentityAccesslist

> UpdateAwsTidyIdentityAccesslist(ctx, mountPath).AwsTidyIdentityAccesslistRequest(awsTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyIdentityAccesslistRequest := NewAwsTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsTidyIdentityAccesslist(context.Background(), mountPath, awsTidyIdentityAccesslistRequest)
	if err != nil {
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

 **awsTidyIdentityAccesslistRequest** | [**AwsTidyIdentityAccesslistRequest**](AwsTidyIdentityAccesslistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsTidyIdentityWhitelist

> UpdateAwsTidyIdentityWhitelist(ctx, mountPath).AwsTidyIdentityWhitelistRequest(awsTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyIdentityWhitelistRequest := NewAwsTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsTidyIdentityWhitelist(context.Background(), mountPath, awsTidyIdentityWhitelistRequest)
	if err != nil {
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

 **awsTidyIdentityWhitelistRequest** | [**AwsTidyIdentityWhitelistRequest**](AwsTidyIdentityWhitelistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsTidyRoletagBlacklist

> UpdateAwsTidyRoletagBlacklist(ctx, mountPath).AwsTidyRoletagBlacklistRequest(awsTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyRoletagBlacklistRequest := NewAwsTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsTidyRoletagBlacklist(context.Background(), mountPath, awsTidyRoletagBlacklistRequest)
	if err != nil {
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

 **awsTidyRoletagBlacklistRequest** | [**AwsTidyRoletagBlacklistRequest**](AwsTidyRoletagBlacklistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateAwsTidyRoletagDenylist

> UpdateAwsTidyRoletagDenylist(ctx, mountPath).AwsTidyRoletagDenylistRequest(awsTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyRoletagDenylistRequest := NewAwsTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateAwsTidyRoletagDenylist(context.Background(), mountPath, awsTidyRoletagDenylistRequest)
	if err != nil {
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

 **awsTidyRoletagDenylistRequest** | [**AwsTidyRoletagDenylistRequest**](AwsTidyRoletagDenylistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateCertCertsName

> UpdateCertCertsName(ctx, mountPath, name).CertCertsRequest(certCertsRequest).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certCertsRequest := NewCertCertsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateCertCertsName(context.Background(), mountPath, name, certCertsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certCertsRequest** | [**CertCertsRequest**](CertCertsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateCertCrlsName

> UpdateCertCrlsName(ctx, mountPath, name).CertCrlsRequest(certCrlsRequest).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certCrlsRequest := NewCertCrlsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateCertCrlsName(context.Background(), mountPath, name, certCrlsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certCrlsRequest** | [**CertCrlsRequest**](CertCrlsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateCfRolesRole

> UpdateCfRolesRole(ctx, mountPath, role).CfRolesRequest(cfRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	cfRolesRequest := NewCfRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateCfRolesRole(context.Background(), mountPath, role, cfRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cfRolesRequest** | [**CfRolesRequest**](CfRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRoleNameLabels

> UpdateGcpRoleNameLabels(ctx, mountPath, name).GcpRoleLabelsRequest(gcpRoleLabelsRequest).Execute()

Add or remove labels for an existing 'gce' role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	gcpRoleLabelsRequest := NewGcpRoleLabelsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateGcpRoleNameLabels(context.Background(), mountPath, name, gcpRoleLabelsRequest)
	if err != nil {
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


 **gcpRoleLabelsRequest** | [**GcpRoleLabelsRequest**](GcpRoleLabelsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGcpRoleNameServiceAccounts

> UpdateGcpRoleNameServiceAccounts(ctx, mountPath, name).GcpRoleServiceAccountsRequest(gcpRoleServiceAccountsRequest).Execute()

Add or remove service accounts for an existing `iam` role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	gcpRoleServiceAccountsRequest := NewGcpRoleServiceAccountsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateGcpRoleNameServiceAccounts(context.Background(), mountPath, name, gcpRoleServiceAccountsRequest)
	if err != nil {
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


 **gcpRoleServiceAccountsRequest** | [**GcpRoleServiceAccountsRequest**](GcpRoleServiceAccountsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGithubMapTeamsKey

> UpdateGithubMapTeamsKey(ctx, key, mountPath).GithubMapTeamsRequest(githubMapTeamsRequest).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubMapTeamsRequest := NewGithubMapTeamsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateGithubMapTeamsKey(context.Background(), key, mountPath, githubMapTeamsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the teams mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **githubMapTeamsRequest** | [**GithubMapTeamsRequest**](GithubMapTeamsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateGithubMapUsersKey

> UpdateGithubMapUsersKey(ctx, key, mountPath).GithubMapUsersRequest(githubMapUsersRequest).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubMapUsersRequest := NewGithubMapUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateGithubMapUsersKey(context.Background(), key, mountPath, githubMapUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**key** | **string** | Key for the users mapping | 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **githubMapUsersRequest** | [**GithubMapUsersRequest**](GithubMapUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateKerberosConfigLdap

> UpdateKerberosConfigLdap(ctx, mountPath).KerberosConfigLdapRequest(kerberosConfigLdapRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosConfigLdapRequest := NewKerberosConfigLdapRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateKerberosConfigLdap(context.Background(), mountPath, kerberosConfigLdapRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosConfigLdapRequest** | [**KerberosConfigLdapRequest**](KerberosConfigLdapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateKerberosGroupsName

> UpdateKerberosGroupsName(ctx, mountPath, name).KerberosGroupsRequest(kerberosGroupsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosGroupsRequest := NewKerberosGroupsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateKerberosGroupsName(context.Background(), mountPath, name, kerberosGroupsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kerberosGroupsRequest** | [**KerberosGroupsRequest**](KerberosGroupsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOciLoginRole

> UpdateOciLoginRole(ctx, mountPath, role).OciLoginRequest(ociLoginRequest).Execute()

Authenticates to Vault using OCI credentials

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	ociLoginRequest := NewOciLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOciLoginRole(context.Background(), mountPath, role, ociLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ociLoginRequest** | [**OciLoginRequest**](OciLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOidcConfig

> UpdateOidcConfig(ctx, mountPath).OidcConfigRequest(oidcConfigRequest).Execute()

Configure the JWT authentication backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcConfigRequest := NewOidcConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOidcConfig(context.Background(), mountPath, oidcConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcConfigRequest** | [**OidcConfigRequest**](OidcConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOidcLogin

> UpdateOidcLogin(ctx, mountPath).OidcLoginRequest(oidcLoginRequest).Execute()

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcLoginRequest := NewOidcLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOidcLogin(context.Background(), mountPath, oidcLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcLoginRequest** | [**OidcLoginRequest**](OidcLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOidcOidcAuthUrl

> UpdateOidcOidcAuthUrl(ctx, mountPath).OidcOidcAuthUrlRequest(oidcOidcAuthUrlRequest).Execute()

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcOidcAuthUrlRequest := NewOidcOidcAuthUrlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOidcOidcAuthUrl(context.Background(), mountPath, oidcOidcAuthUrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcOidcAuthUrlRequest** | [**OidcOidcAuthUrlRequest**](OidcOidcAuthUrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOidcOidcCallback

> UpdateOidcOidcCallback(ctx, mountPath).OidcOidcCallbackRequest(oidcOidcCallbackRequest).Execute()

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcOidcCallbackRequest := NewOidcOidcCallbackRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOidcOidcCallback(context.Background(), mountPath, oidcOidcCallbackRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcOidcCallbackRequest** | [**OidcOidcCallbackRequest**](OidcOidcCallbackRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateOidcRoleName

> UpdateOidcRoleName(ctx, mountPath, name).OidcRoleRequest(oidcRoleRequest).Execute()

Register an role with the backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcRoleRequest := NewOidcRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateOidcRoleName(context.Background(), mountPath, name, oidcRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oidcRoleRequest** | [**OidcRoleRequest**](OidcRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateRadiusLoginUrlusername

> UpdateRadiusLoginUrlusername(ctx, mountPath, urlusername).RadiusLoginRequest(radiusLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	urlusername := "urlusername_example" // string | Username to be used for login. (URL parameter)
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateRadiusLoginUrlusername(context.Background(), mountPath, urlusername, radiusLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]
**urlusername** | **string** | Username to be used for login. (URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateRadiusUsersName

> UpdateRadiusUsersName(ctx, mountPath, name).RadiusUsersRequest(radiusUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusUsersRequest := NewRadiusUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateRadiusUsersName(context.Background(), mountPath, name, radiusUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **radiusUsersRequest** | [**RadiusUsersRequest**](RadiusUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenCreate

> UpdateTokenCreate(ctx, mountPath).Format(format).TokenCreateRequest(tokenCreateRequest).Execute()

The token create path is used to create new tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	format := NewstringWithDefaults()
	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenCreate(context.Background(), mountPath, format, tokenCreateRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Return json formatted output | 

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenCreateOrphan

> UpdateTokenCreateOrphan(ctx, mountPath).Format(format).TokenCreateOrphanRequest(tokenCreateOrphanRequest).Execute()

The token create path is used to create new orphan tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	format := NewstringWithDefaults()
	tokenCreateOrphanRequest := NewTokenCreateOrphanRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenCreateOrphan(context.Background(), mountPath, format, tokenCreateOrphanRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Return json formatted output | 

 **tokenCreateOrphanRequest** | [**TokenCreateOrphanRequest**](TokenCreateOrphanRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenCreateRoleName

> UpdateTokenCreateRoleName(ctx, mountPath, roleName).Format(format).TokenCreateRequest(tokenCreateRequest).Execute()

This token create path is used to create new tokens adhering to the given role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	format := NewstringWithDefaults()
	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenCreateRoleName(context.Background(), mountPath, roleName, format, tokenCreateRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | Return json formatted output | 

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenLookup

> UpdateTokenLookup(ctx, mountPath).TokenLookupRequest(tokenLookupRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupRequest := NewTokenLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenLookup(context.Background(), mountPath, tokenLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenLookupRequest** | [**TokenLookupRequest**](TokenLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenLookupAccessor

> UpdateTokenLookupAccessor(ctx, mountPath).TokenLookupAccessorRequest(tokenLookupAccessorRequest).Execute()

This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupAccessorRequest := NewTokenLookupAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenLookupAccessor(context.Background(), mountPath, tokenLookupAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenLookupAccessorRequest** | [**TokenLookupAccessorRequest**](TokenLookupAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenLookupSelf

> UpdateTokenLookupSelf(ctx, mountPath).TokenLookupSelfRequest(tokenLookupSelfRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupSelfRequest := NewTokenLookupSelfRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenLookupSelf(context.Background(), mountPath, tokenLookupSelfRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenLookupSelfRequest** | [**TokenLookupSelfRequest**](TokenLookupSelfRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRenew

> UpdateTokenRenew(ctx, mountPath).TokenRenewRequest(tokenRenewRequest).Execute()

This endpoint will renew the given token and prevent expiration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewRequest := NewTokenRenewRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRenew(context.Background(), mountPath, tokenRenewRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRenewRequest** | [**TokenRenewRequest**](TokenRenewRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRenewAccessor

> UpdateTokenRenewAccessor(ctx, mountPath).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()

This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewAccessorRequest := NewTokenRenewAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRenewAccessor(context.Background(), mountPath, tokenRenewAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRenewAccessorRequest** | [**TokenRenewAccessorRequest**](TokenRenewAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRenewSelf

> UpdateTokenRenewSelf(ctx, mountPath).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()

This endpoint will renew the token used to call it and prevent expiration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewSelfRequest := NewTokenRenewSelfRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRenewSelf(context.Background(), mountPath, tokenRenewSelfRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRenewSelfRequest** | [**TokenRenewSelfRequest**](TokenRenewSelfRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRevoke

> UpdateTokenRevoke(ctx, mountPath).TokenRevokeRequest(tokenRevokeRequest).Execute()

This endpoint will delete the given token and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeRequest := NewTokenRevokeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRevoke(context.Background(), mountPath, tokenRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRevokeAccessor

> UpdateTokenRevokeAccessor(ctx, mountPath).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()

This endpoint will delete the token associated with the accessor and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeAccessorRequest := NewTokenRevokeAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRevokeAccessor(context.Background(), mountPath, tokenRevokeAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeAccessorRequest** | [**TokenRevokeAccessorRequest**](TokenRevokeAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRevokeOrphan

> UpdateTokenRevokeOrphan(ctx, mountPath).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()

This endpoint will delete the token and orphan its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeOrphanRequest := NewTokenRevokeOrphanRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRevokeOrphan(context.Background(), mountPath, tokenRevokeOrphanRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeOrphanRequest** | [**TokenRevokeOrphanRequest**](TokenRevokeOrphanRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRevokeSelf

> UpdateTokenRevokeSelf(ctx, mountPath).Execute()

This endpoint will delete the token used to call it and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.UpdateTokenRevokeSelf(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenRolesRoleName

> UpdateTokenRolesRoleName(ctx, mountPath, roleName).TokenRolesRequest(tokenRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRolesRequest := NewTokenRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateTokenRolesRoleName(context.Background(), mountPath, roleName, tokenRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenRolesRequest** | [**TokenRolesRequest**](TokenRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateTokenTidy

> UpdateTokenTidy(ctx, mountPath).Execute()

This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.UpdateTokenTidy(context.Background(), mountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateUserpassLoginUsername

> UpdateUserpassLoginUsername(ctx, mountPath, username).UserpassLoginRequest(userpassLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username of the user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassLoginRequest := NewUserpassLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateUserpassLoginUsername(context.Background(), mountPath, username, userpassLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username of the user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassLoginRequest** | [**UserpassLoginRequest**](UserpassLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateUserpassUsersUsername

> UpdateUserpassUsersUsername(ctx, mountPath, username).UserpassUsersRequest(userpassUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersRequest := NewUserpassUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateUserpassUsersUsername(context.Background(), mountPath, username, userpassUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassUsersRequest** | [**UserpassUsersRequest**](UserpassUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateUserpassUsersUsernamePassword

> UpdateUserpassUsersUsernamePassword(ctx, mountPath, username).UserpassUsersPasswordRequest(userpassUsersPasswordRequest).Execute()

Reset user's password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersPasswordRequest := NewUserpassUsersPasswordRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateUserpassUsersUsernamePassword(context.Background(), mountPath, username, userpassUsersPasswordRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassUsersPasswordRequest** | [**UserpassUsersPasswordRequest**](UserpassUsersPasswordRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateUserpassUsersUsernamePolicies

> UpdateUserpassUsersUsernamePolicies(ctx, mountPath, username).UserpassUsersPoliciesRequest(userpassUsersPoliciesRequest).Execute()

Update the policies associated with the username.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	mountPath := "mountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersPoliciesRequest := NewUserpassUsersPoliciesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.UpdateUserpassUsersUsernamePolicies(context.Background(), mountPath, username, userpassUsersPoliciesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**mountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassUsersPoliciesRequest** | [**UserpassUsersPoliciesRequest**](UserpassUsersPoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

