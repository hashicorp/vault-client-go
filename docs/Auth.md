# Auth

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthAlicloudRoleRole**](Auth.md#DeleteAuthAlicloudRoleRole) | **Delete** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**DeleteAuthApproleRoleRoleName**](Auth.md#DeleteAuthApproleRoleRoleName) | **Delete** /auth/{approle_mount_path}/role/{role_name} | Register an role with the backend.
[**DeleteAuthApproleRoleRoleNameBindSecretId**](Auth.md#DeleteAuthApproleRoleRoleNameBindSecretId) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**DeleteAuthApproleRoleRoleNameBoundCidrList**](Auth.md#DeleteAuthApproleRoleRoleNameBoundCidrList) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteAuthApproleRoleRoleNamePeriod**](Auth.md#DeleteAuthApproleRoleRoleNamePeriod) | **Delete** /auth/{approle_mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**DeleteAuthApproleRoleRoleNamePolicies**](Auth.md#DeleteAuthApproleRoleRoleNamePolicies) | **Delete** /auth/{approle_mount_path}/role/{role_name}/policies | Policies of the role.
[**DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteAuthApproleRoleRoleNameSecretIdDestroy**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**DeleteAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**DeleteAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**DeleteAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#DeleteAuthApproleRoleRoleNameTokenBoundCidrs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**DeleteAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#DeleteAuthApproleRoleRoleNameTokenMaxTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**DeleteAuthApproleRoleRoleNameTokenNumUses**](Auth.md#DeleteAuthApproleRoleRoleNameTokenNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**DeleteAuthApproleRoleRoleNameTokenTtl**](Auth.md#DeleteAuthApproleRoleRoleNameTokenTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**DeleteAuthAwsConfigCertificateCertName**](Auth.md#DeleteAuthAwsConfigCertificateCertName) | **Delete** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**DeleteAuthAwsConfigClient**](Auth.md#DeleteAuthAwsConfigClient) | **Delete** /auth/{aws_mount_path}/config/client | 
[**DeleteAuthAwsConfigStsAccountId**](Auth.md#DeleteAuthAwsConfigStsAccountId) | **Delete** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**DeleteAuthAwsConfigTidyIdentityAccesslist**](Auth.md#DeleteAuthAwsConfigTidyIdentityAccesslist) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**DeleteAuthAwsConfigTidyIdentityWhitelist**](Auth.md#DeleteAuthAwsConfigTidyIdentityWhitelist) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**DeleteAuthAwsConfigTidyRoletagBlacklist**](Auth.md#DeleteAuthAwsConfigTidyRoletagBlacklist) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**DeleteAuthAwsConfigTidyRoletagDenylist**](Auth.md#DeleteAuthAwsConfigTidyRoletagDenylist) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**DeleteAuthAwsIdentityAccesslistInstanceId**](Auth.md#DeleteAuthAwsIdentityAccesslistInstanceId) | **Delete** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**DeleteAuthAwsIdentityWhitelistInstanceId**](Auth.md#DeleteAuthAwsIdentityWhitelistInstanceId) | **Delete** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**DeleteAuthAwsRoleRole**](Auth.md#DeleteAuthAwsRoleRole) | **Delete** /auth/{aws_mount_path}/role/{role} | 
[**DeleteAuthAwsRoletagBlacklistRoleTag**](Auth.md#DeleteAuthAwsRoletagBlacklistRoleTag) | **Delete** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**DeleteAuthAwsRoletagDenylistRoleTag**](Auth.md#DeleteAuthAwsRoletagDenylistRoleTag) | **Delete** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**DeleteAuthAzureConfig**](Auth.md#DeleteAuthAzureConfig) | **Delete** /auth/{azure_mount_path}/config | 
[**DeleteAuthAzureRoleName**](Auth.md#DeleteAuthAzureRoleName) | **Delete** /auth/{azure_mount_path}/role/{name} | 
[**DeleteAuthCertCertsName**](Auth.md#DeleteAuthCertCertsName) | **Delete** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**DeleteAuthCertCrlsName**](Auth.md#DeleteAuthCertCrlsName) | **Delete** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**DeleteAuthCfConfig**](Auth.md#DeleteAuthCfConfig) | **Delete** /auth/{cf_mount_path}/config | 
[**DeleteAuthCfRolesRole**](Auth.md#DeleteAuthCfRolesRole) | **Delete** /auth/{cf_mount_path}/roles/{role} | 
[**DeleteAuthGcpRoleName**](Auth.md#DeleteAuthGcpRoleName) | **Delete** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**DeleteAuthGithubMapTeamsKey**](Auth.md#DeleteAuthGithubMapTeamsKey) | **Delete** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**DeleteAuthGithubMapUsersKey**](Auth.md#DeleteAuthGithubMapUsersKey) | **Delete** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**DeleteAuthJwtRoleName**](Auth.md#DeleteAuthJwtRoleName) | **Delete** /auth/{jwt_mount_path}/role/{name} | Delete an existing role.
[**DeleteAuthKerberosGroupsName**](Auth.md#DeleteAuthKerberosGroupsName) | **Delete** /auth/{kerberos_mount_path}/groups/{name} | 
[**DeleteAuthKubernetesRoleName**](Auth.md#DeleteAuthKubernetesRoleName) | **Delete** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**DeleteAuthLdapGroupsName**](Auth.md#DeleteAuthLdapGroupsName) | **Delete** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**DeleteAuthLdapUsersName**](Auth.md#DeleteAuthLdapUsersName) | **Delete** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**DeleteAuthOciConfig**](Auth.md#DeleteAuthOciConfig) | **Delete** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**DeleteAuthOciRoleRole**](Auth.md#DeleteAuthOciRoleRole) | **Delete** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**DeleteAuthOidcRoleName**](Auth.md#DeleteAuthOidcRoleName) | **Delete** /auth/{oidc_mount_path}/role/{name} | Delete an existing role.
[**DeleteAuthOktaGroupsName**](Auth.md#DeleteAuthOktaGroupsName) | **Delete** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**DeleteAuthOktaUsersName**](Auth.md#DeleteAuthOktaUsersName) | **Delete** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**DeleteAuthRadiusUsersName**](Auth.md#DeleteAuthRadiusUsersName) | **Delete** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**DeleteAuthTokenRolesRoleName**](Auth.md#DeleteAuthTokenRolesRoleName) | **Delete** /auth/{token_mount_path}/roles/{role_name} | 
[**DeleteAuthUserpassUsersUsername**](Auth.md#DeleteAuthUserpassUsersUsername) | **Delete** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**GetAuthAlicloudRole**](Auth.md#GetAuthAlicloudRole) | **Get** /auth/{alicloud_mount_path}/role | Lists all the roles that are registered with Vault.
[**GetAuthAlicloudRoleRole**](Auth.md#GetAuthAlicloudRoleRole) | **Get** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**GetAuthAlicloudRoles**](Auth.md#GetAuthAlicloudRoles) | **Get** /auth/{alicloud_mount_path}/roles | Lists all the roles that are registered with Vault.
[**GetAuthApproleRole**](Auth.md#GetAuthApproleRole) | **Get** /auth/{approle_mount_path}/role | Lists all the roles registered with the backend.
[**GetAuthApproleRoleRoleName**](Auth.md#GetAuthApproleRoleRoleName) | **Get** /auth/{approle_mount_path}/role/{role_name} | Register an role with the backend.
[**GetAuthApproleRoleRoleNameBindSecretId**](Auth.md#GetAuthApproleRoleRoleNameBindSecretId) | **Get** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**GetAuthApproleRoleRoleNameBoundCidrList**](Auth.md#GetAuthApproleRoleRoleNameBoundCidrList) | **Get** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**GetAuthApproleRoleRoleNameLocalSecretIds**](Auth.md#GetAuthApproleRoleRoleNameLocalSecretIds) | **Get** /auth/{approle_mount_path}/role/{role_name}/local-secret-ids | Enables cluster local secret IDs
[**GetAuthApproleRoleRoleNamePeriod**](Auth.md#GetAuthApproleRoleRoleNamePeriod) | **Get** /auth/{approle_mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**GetAuthApproleRoleRoleNamePolicies**](Auth.md#GetAuthApproleRoleRoleNamePolicies) | **Get** /auth/{approle_mount_path}/role/{role_name}/policies | Policies of the role.
[**GetAuthApproleRoleRoleNameRoleId**](Auth.md#GetAuthApproleRoleRoleNameRoleId) | **Get** /auth/{approle_mount_path}/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**GetAuthApproleRoleRoleNameSecretId**](Auth.md#GetAuthApproleRoleRoleNameSecretId) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id | Generate a SecretID against this role.
[**GetAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#GetAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**GetAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#GetAuthApproleRoleRoleNameSecretIdNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**GetAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#GetAuthApproleRoleRoleNameSecretIdTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**GetAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#GetAuthApproleRoleRoleNameTokenBoundCidrs) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**GetAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#GetAuthApproleRoleRoleNameTokenMaxTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**GetAuthApproleRoleRoleNameTokenNumUses**](Auth.md#GetAuthApproleRoleRoleNameTokenNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**GetAuthApproleRoleRoleNameTokenTtl**](Auth.md#GetAuthApproleRoleRoleNameTokenTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**GetAuthAwsConfigCertificateCertName**](Auth.md#GetAuthAwsConfigCertificateCertName) | **Get** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**GetAuthAwsConfigCertificates**](Auth.md#GetAuthAwsConfigCertificates) | **Get** /auth/{aws_mount_path}/config/certificates | 
[**GetAuthAwsConfigClient**](Auth.md#GetAuthAwsConfigClient) | **Get** /auth/{aws_mount_path}/config/client | 
[**GetAuthAwsConfigIdentity**](Auth.md#GetAuthAwsConfigIdentity) | **Get** /auth/{aws_mount_path}/config/identity | 
[**GetAuthAwsConfigSts**](Auth.md#GetAuthAwsConfigSts) | **Get** /auth/{aws_mount_path}/config/sts | 
[**GetAuthAwsConfigStsAccountId**](Auth.md#GetAuthAwsConfigStsAccountId) | **Get** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**GetAuthAwsConfigTidyIdentityAccesslist**](Auth.md#GetAuthAwsConfigTidyIdentityAccesslist) | **Get** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**GetAuthAwsConfigTidyIdentityWhitelist**](Auth.md#GetAuthAwsConfigTidyIdentityWhitelist) | **Get** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**GetAuthAwsConfigTidyRoletagBlacklist**](Auth.md#GetAuthAwsConfigTidyRoletagBlacklist) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**GetAuthAwsConfigTidyRoletagDenylist**](Auth.md#GetAuthAwsConfigTidyRoletagDenylist) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**GetAuthAwsIdentityAccesslist**](Auth.md#GetAuthAwsIdentityAccesslist) | **Get** /auth/{aws_mount_path}/identity-accesslist | 
[**GetAuthAwsIdentityAccesslistInstanceId**](Auth.md#GetAuthAwsIdentityAccesslistInstanceId) | **Get** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**GetAuthAwsIdentityWhitelist**](Auth.md#GetAuthAwsIdentityWhitelist) | **Get** /auth/{aws_mount_path}/identity-whitelist | 
[**GetAuthAwsIdentityWhitelistInstanceId**](Auth.md#GetAuthAwsIdentityWhitelistInstanceId) | **Get** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**GetAuthAwsRole**](Auth.md#GetAuthAwsRole) | **Get** /auth/{aws_mount_path}/role | 
[**GetAuthAwsRoleRole**](Auth.md#GetAuthAwsRoleRole) | **Get** /auth/{aws_mount_path}/role/{role} | 
[**GetAuthAwsRoles**](Auth.md#GetAuthAwsRoles) | **Get** /auth/{aws_mount_path}/roles | 
[**GetAuthAwsRoletagBlacklist**](Auth.md#GetAuthAwsRoletagBlacklist) | **Get** /auth/{aws_mount_path}/roletag-blacklist | 
[**GetAuthAwsRoletagBlacklistRoleTag**](Auth.md#GetAuthAwsRoletagBlacklistRoleTag) | **Get** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**GetAuthAwsRoletagDenylist**](Auth.md#GetAuthAwsRoletagDenylist) | **Get** /auth/{aws_mount_path}/roletag-denylist | 
[**GetAuthAwsRoletagDenylistRoleTag**](Auth.md#GetAuthAwsRoletagDenylistRoleTag) | **Get** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**GetAuthAzureConfig**](Auth.md#GetAuthAzureConfig) | **Get** /auth/{azure_mount_path}/config | 
[**GetAuthAzureRole**](Auth.md#GetAuthAzureRole) | **Get** /auth/{azure_mount_path}/role | 
[**GetAuthAzureRoleName**](Auth.md#GetAuthAzureRoleName) | **Get** /auth/{azure_mount_path}/role/{name} | 
[**GetAuthCentrifyConfig**](Auth.md#GetAuthCentrifyConfig) | **Get** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**GetAuthCertCerts**](Auth.md#GetAuthCertCerts) | **Get** /auth/{cert_mount_path}/certs | Manage trusted certificates used for authentication.
[**GetAuthCertCertsName**](Auth.md#GetAuthCertCertsName) | **Get** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**GetAuthCertConfig**](Auth.md#GetAuthCertConfig) | **Get** /auth/{cert_mount_path}/config | 
[**GetAuthCertCrlsName**](Auth.md#GetAuthCertCrlsName) | **Get** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**GetAuthCfConfig**](Auth.md#GetAuthCfConfig) | **Get** /auth/{cf_mount_path}/config | 
[**GetAuthCfRoles**](Auth.md#GetAuthCfRoles) | **Get** /auth/{cf_mount_path}/roles | 
[**GetAuthCfRolesRole**](Auth.md#GetAuthCfRolesRole) | **Get** /auth/{cf_mount_path}/roles/{role} | 
[**GetAuthGcpConfig**](Auth.md#GetAuthGcpConfig) | **Get** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GetAuthGcpRole**](Auth.md#GetAuthGcpRole) | **Get** /auth/{gcp_mount_path}/role | Lists all the roles that are registered with Vault.
[**GetAuthGcpRoleName**](Auth.md#GetAuthGcpRoleName) | **Get** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GetAuthGcpRoles**](Auth.md#GetAuthGcpRoles) | **Get** /auth/{gcp_mount_path}/roles | Lists all the roles that are registered with Vault.
[**GetAuthGithubConfig**](Auth.md#GetAuthGithubConfig) | **Get** /auth/{github_mount_path}/config | 
[**GetAuthGithubMapTeams**](Auth.md#GetAuthGithubMapTeams) | **Get** /auth/{github_mount_path}/map/teams | Read mappings for teams
[**GetAuthGithubMapTeamsKey**](Auth.md#GetAuthGithubMapTeamsKey) | **Get** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GetAuthGithubMapUsers**](Auth.md#GetAuthGithubMapUsers) | **Get** /auth/{github_mount_path}/map/users | Read mappings for users
[**GetAuthGithubMapUsersKey**](Auth.md#GetAuthGithubMapUsersKey) | **Get** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GetAuthJwtConfig**](Auth.md#GetAuthJwtConfig) | **Get** /auth/{jwt_mount_path}/config | Read the current JWT authentication backend configuration.
[**GetAuthJwtOidcCallback**](Auth.md#GetAuthJwtOidcCallback) | **Get** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**GetAuthJwtRole**](Auth.md#GetAuthJwtRole) | **Get** /auth/{jwt_mount_path}/role | Lists all the roles registered with the backend.
[**GetAuthJwtRoleName**](Auth.md#GetAuthJwtRoleName) | **Get** /auth/{jwt_mount_path}/role/{name} | Read an existing role.
[**GetAuthKerberosConfig**](Auth.md#GetAuthKerberosConfig) | **Get** /auth/{kerberos_mount_path}/config | 
[**GetAuthKerberosConfigLdap**](Auth.md#GetAuthKerberosConfigLdap) | **Get** /auth/{kerberos_mount_path}/config/ldap | 
[**GetAuthKerberosGroups**](Auth.md#GetAuthKerberosGroups) | **Get** /auth/{kerberos_mount_path}/groups | 
[**GetAuthKerberosGroupsName**](Auth.md#GetAuthKerberosGroupsName) | **Get** /auth/{kerberos_mount_path}/groups/{name} | 
[**GetAuthKerberosLogin**](Auth.md#GetAuthKerberosLogin) | **Get** /auth/{kerberos_mount_path}/login | 
[**GetAuthKubernetesConfig**](Auth.md#GetAuthKubernetesConfig) | **Get** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**GetAuthKubernetesRole**](Auth.md#GetAuthKubernetesRole) | **Get** /auth/{kubernetes_mount_path}/role | Lists all the roles registered with the backend.
[**GetAuthKubernetesRoleName**](Auth.md#GetAuthKubernetesRoleName) | **Get** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**GetAuthLdapConfig**](Auth.md#GetAuthLdapConfig) | **Get** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**GetAuthLdapGroups**](Auth.md#GetAuthLdapGroups) | **Get** /auth/{ldap_mount_path}/groups | Manage additional groups for users allowed to authenticate.
[**GetAuthLdapGroupsName**](Auth.md#GetAuthLdapGroupsName) | **Get** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**GetAuthLdapUsers**](Auth.md#GetAuthLdapUsers) | **Get** /auth/{ldap_mount_path}/users | Manage users allowed to authenticate.
[**GetAuthLdapUsersName**](Auth.md#GetAuthLdapUsersName) | **Get** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**GetAuthOciConfig**](Auth.md#GetAuthOciConfig) | **Get** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**GetAuthOciRole**](Auth.md#GetAuthOciRole) | **Get** /auth/{oci_mount_path}/role | Lists all the roles that are registered with Vault.
[**GetAuthOciRoleRole**](Auth.md#GetAuthOciRoleRole) | **Get** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**GetAuthOidcConfig**](Auth.md#GetAuthOidcConfig) | **Get** /auth/{oidc_mount_path}/config | Read the current JWT authentication backend configuration.
[**GetAuthOidcOidcCallback**](Auth.md#GetAuthOidcOidcCallback) | **Get** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**GetAuthOidcRole**](Auth.md#GetAuthOidcRole) | **Get** /auth/{oidc_mount_path}/role | Lists all the roles registered with the backend.
[**GetAuthOidcRoleName**](Auth.md#GetAuthOidcRoleName) | **Get** /auth/{oidc_mount_path}/role/{name} | Read an existing role.
[**GetAuthOktaConfig**](Auth.md#GetAuthOktaConfig) | **Get** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**GetAuthOktaGroups**](Auth.md#GetAuthOktaGroups) | **Get** /auth/{okta_mount_path}/groups | Manage users allowed to authenticate.
[**GetAuthOktaGroupsName**](Auth.md#GetAuthOktaGroupsName) | **Get** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**GetAuthOktaUsers**](Auth.md#GetAuthOktaUsers) | **Get** /auth/{okta_mount_path}/users | Manage additional groups for users allowed to authenticate.
[**GetAuthOktaUsersName**](Auth.md#GetAuthOktaUsersName) | **Get** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**GetAuthOktaVerifyNonce**](Auth.md#GetAuthOktaVerifyNonce) | **Get** /auth/{okta_mount_path}/verify/{nonce} | 
[**GetAuthRadiusConfig**](Auth.md#GetAuthRadiusConfig) | **Get** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**GetAuthRadiusUsers**](Auth.md#GetAuthRadiusUsers) | **Get** /auth/{radius_mount_path}/users | Manage users allowed to authenticate.
[**GetAuthRadiusUsersName**](Auth.md#GetAuthRadiusUsersName) | **Get** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**GetAuthTokenAccessors**](Auth.md#GetAuthTokenAccessors) | **Get** /auth/{token_mount_path}/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#39;sudo&#39; capability in addition to &#39;list&#39;.
[**GetAuthTokenLookup**](Auth.md#GetAuthTokenLookup) | **Get** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**GetAuthTokenLookupSelf**](Auth.md#GetAuthTokenLookupSelf) | **Get** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**GetAuthTokenRoles**](Auth.md#GetAuthTokenRoles) | **Get** /auth/{token_mount_path}/roles | This endpoint lists configured roles.
[**GetAuthTokenRolesRoleName**](Auth.md#GetAuthTokenRolesRoleName) | **Get** /auth/{token_mount_path}/roles/{role_name} | 
[**GetAuthUserpassUsers**](Auth.md#GetAuthUserpassUsers) | **Get** /auth/{userpass_mount_path}/users | Manage users allowed to authenticate.
[**GetAuthUserpassUsersUsername**](Auth.md#GetAuthUserpassUsersUsername) | **Get** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**PostAuthAlicloudLogin**](Auth.md#PostAuthAlicloudLogin) | **Post** /auth/{alicloud_mount_path}/login | Authenticates an RAM entity with Vault.
[**PostAuthAlicloudRoleRole**](Auth.md#PostAuthAlicloudRoleRole) | **Post** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**PostAuthApproleLogin**](Auth.md#PostAuthApproleLogin) | **Post** /auth/{approle_mount_path}/login | 
[**PostAuthApproleRoleRoleName**](Auth.md#PostAuthApproleRoleRoleName) | **Post** /auth/{approle_mount_path}/role/{role_name} | Register an role with the backend.
[**PostAuthApproleRoleRoleNameBindSecretId**](Auth.md#PostAuthApproleRoleRoleNameBindSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**PostAuthApproleRoleRoleNameBoundCidrList**](Auth.md#PostAuthApproleRoleRoleNameBoundCidrList) | **Post** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**PostAuthApproleRoleRoleNameCustomSecretId**](Auth.md#PostAuthApproleRoleRoleNameCustomSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/custom-secret-id | Assign a SecretID of choice against the role.
[**PostAuthApproleRoleRoleNamePeriod**](Auth.md#PostAuthApproleRoleRoleNamePeriod) | **Post** /auth/{approle_mount_path}/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**PostAuthApproleRoleRoleNamePolicies**](Auth.md#PostAuthApproleRoleRoleNamePolicies) | **Post** /auth/{approle_mount_path}/role/{role_name}/policies | Policies of the role.
[**PostAuthApproleRoleRoleNameRoleId**](Auth.md#PostAuthApproleRoleRoleNameRoleId) | **Post** /auth/{approle_mount_path}/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**PostAuthApproleRoleRoleNameSecretId**](Auth.md#PostAuthApproleRoleRoleNameSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id | Generate a SecretID against this role.
[**PostAuthApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#PostAuthApproleRoleRoleNameSecretIdAccessorDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**PostAuthApproleRoleRoleNameSecretIdAccessorLookup**](Auth.md#PostAuthApproleRoleRoleNameSecretIdAccessorLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup | 
[**PostAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#PostAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**PostAuthApproleRoleRoleNameSecretIdDestroy**](Auth.md#PostAuthApproleRoleRoleNameSecretIdDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**PostAuthApproleRoleRoleNameSecretIdLookup**](Auth.md#PostAuthApproleRoleRoleNameSecretIdLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/lookup | Read the properties of an issued secret_id
[**PostAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#PostAuthApproleRoleRoleNameSecretIdNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**PostAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#PostAuthApproleRoleRoleNameSecretIdTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | Duration in seconds of the SecretID generated against the role.
[**PostAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#PostAuthApproleRoleRoleNameTokenBoundCidrs) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**PostAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#PostAuthApproleRoleRoleNameTokenMaxTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**PostAuthApproleRoleRoleNameTokenNumUses**](Auth.md#PostAuthApproleRoleRoleNameTokenNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**PostAuthApproleRoleRoleNameTokenTtl**](Auth.md#PostAuthApproleRoleRoleNameTokenTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**PostAuthApproleTidySecretId**](Auth.md#PostAuthApproleTidySecretId) | **Post** /auth/{approle_mount_path}/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
[**PostAuthAwsConfigCertificateCertName**](Auth.md#PostAuthAwsConfigCertificateCertName) | **Post** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**PostAuthAwsConfigClient**](Auth.md#PostAuthAwsConfigClient) | **Post** /auth/{aws_mount_path}/config/client | 
[**PostAuthAwsConfigIdentity**](Auth.md#PostAuthAwsConfigIdentity) | **Post** /auth/{aws_mount_path}/config/identity | 
[**PostAuthAwsConfigRotateRoot**](Auth.md#PostAuthAwsConfigRotateRoot) | **Post** /auth/{aws_mount_path}/config/rotate-root | 
[**PostAuthAwsConfigStsAccountId**](Auth.md#PostAuthAwsConfigStsAccountId) | **Post** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**PostAuthAwsConfigTidyIdentityAccesslist**](Auth.md#PostAuthAwsConfigTidyIdentityAccesslist) | **Post** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**PostAuthAwsConfigTidyIdentityWhitelist**](Auth.md#PostAuthAwsConfigTidyIdentityWhitelist) | **Post** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**PostAuthAwsConfigTidyRoletagBlacklist**](Auth.md#PostAuthAwsConfigTidyRoletagBlacklist) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**PostAuthAwsConfigTidyRoletagDenylist**](Auth.md#PostAuthAwsConfigTidyRoletagDenylist) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**PostAuthAwsLogin**](Auth.md#PostAuthAwsLogin) | **Post** /auth/{aws_mount_path}/login | 
[**PostAuthAwsRoleRole**](Auth.md#PostAuthAwsRoleRole) | **Post** /auth/{aws_mount_path}/role/{role} | 
[**PostAuthAwsRoleRoleTag**](Auth.md#PostAuthAwsRoleRoleTag) | **Post** /auth/{aws_mount_path}/role/{role}/tag | 
[**PostAuthAwsRoletagBlacklistRoleTag**](Auth.md#PostAuthAwsRoletagBlacklistRoleTag) | **Post** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**PostAuthAwsRoletagDenylistRoleTag**](Auth.md#PostAuthAwsRoletagDenylistRoleTag) | **Post** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**PostAuthAwsTidyIdentityAccesslist**](Auth.md#PostAuthAwsTidyIdentityAccesslist) | **Post** /auth/{aws_mount_path}/tidy/identity-accesslist | 
[**PostAuthAwsTidyIdentityWhitelist**](Auth.md#PostAuthAwsTidyIdentityWhitelist) | **Post** /auth/{aws_mount_path}/tidy/identity-whitelist | 
[**PostAuthAwsTidyRoletagBlacklist**](Auth.md#PostAuthAwsTidyRoletagBlacklist) | **Post** /auth/{aws_mount_path}/tidy/roletag-blacklist | 
[**PostAuthAwsTidyRoletagDenylist**](Auth.md#PostAuthAwsTidyRoletagDenylist) | **Post** /auth/{aws_mount_path}/tidy/roletag-denylist | 
[**PostAuthAzureConfig**](Auth.md#PostAuthAzureConfig) | **Post** /auth/{azure_mount_path}/config | 
[**PostAuthAzureLogin**](Auth.md#PostAuthAzureLogin) | **Post** /auth/{azure_mount_path}/login | 
[**PostAuthAzureRoleName**](Auth.md#PostAuthAzureRoleName) | **Post** /auth/{azure_mount_path}/role/{name} | 
[**PostAuthCentrifyConfig**](Auth.md#PostAuthCentrifyConfig) | **Post** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**PostAuthCentrifyLogin**](Auth.md#PostAuthCentrifyLogin) | **Post** /auth/{centrify_mount_path}/login | Log in with a username and password.
[**PostAuthCertCertsName**](Auth.md#PostAuthCertCertsName) | **Post** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**PostAuthCertConfig**](Auth.md#PostAuthCertConfig) | **Post** /auth/{cert_mount_path}/config | 
[**PostAuthCertCrlsName**](Auth.md#PostAuthCertCrlsName) | **Post** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**PostAuthCertLogin**](Auth.md#PostAuthCertLogin) | **Post** /auth/{cert_mount_path}/login | 
[**PostAuthCfConfig**](Auth.md#PostAuthCfConfig) | **Post** /auth/{cf_mount_path}/config | 
[**PostAuthCfLogin**](Auth.md#PostAuthCfLogin) | **Post** /auth/{cf_mount_path}/login | 
[**PostAuthCfRolesRole**](Auth.md#PostAuthCfRolesRole) | **Post** /auth/{cf_mount_path}/roles/{role} | 
[**PostAuthGcpConfig**](Auth.md#PostAuthGcpConfig) | **Post** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**PostAuthGcpLogin**](Auth.md#PostAuthGcpLogin) | **Post** /auth/{gcp_mount_path}/login | 
[**PostAuthGcpRoleName**](Auth.md#PostAuthGcpRoleName) | **Post** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**PostAuthGcpRoleNameLabels**](Auth.md#PostAuthGcpRoleNameLabels) | **Post** /auth/{gcp_mount_path}/role/{name}/labels | Add or remove labels for an existing &#39;gce&#39; role
[**PostAuthGcpRoleNameServiceAccounts**](Auth.md#PostAuthGcpRoleNameServiceAccounts) | **Post** /auth/{gcp_mount_path}/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**PostAuthGithubConfig**](Auth.md#PostAuthGithubConfig) | **Post** /auth/{github_mount_path}/config | 
[**PostAuthGithubLogin**](Auth.md#PostAuthGithubLogin) | **Post** /auth/{github_mount_path}/login | 
[**PostAuthGithubMapTeamsKey**](Auth.md#PostAuthGithubMapTeamsKey) | **Post** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**PostAuthGithubMapUsersKey**](Auth.md#PostAuthGithubMapUsersKey) | **Post** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**PostAuthJwtConfig**](Auth.md#PostAuthJwtConfig) | **Post** /auth/{jwt_mount_path}/config | Configure the JWT authentication backend.
[**PostAuthJwtLogin**](Auth.md#PostAuthJwtLogin) | **Post** /auth/{jwt_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**PostAuthJwtOidcAuthUrl**](Auth.md#PostAuthJwtOidcAuthUrl) | **Post** /auth/{jwt_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**PostAuthJwtOidcCallback**](Auth.md#PostAuthJwtOidcCallback) | **Post** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**PostAuthJwtRoleName**](Auth.md#PostAuthJwtRoleName) | **Post** /auth/{jwt_mount_path}/role/{name} | Register an role with the backend.
[**PostAuthKerberosConfig**](Auth.md#PostAuthKerberosConfig) | **Post** /auth/{kerberos_mount_path}/config | 
[**PostAuthKerberosConfigLdap**](Auth.md#PostAuthKerberosConfigLdap) | **Post** /auth/{kerberos_mount_path}/config/ldap | 
[**PostAuthKerberosGroupsName**](Auth.md#PostAuthKerberosGroupsName) | **Post** /auth/{kerberos_mount_path}/groups/{name} | 
[**PostAuthKerberosLogin**](Auth.md#PostAuthKerberosLogin) | **Post** /auth/{kerberos_mount_path}/login | 
[**PostAuthKubernetesConfig**](Auth.md#PostAuthKubernetesConfig) | **Post** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**PostAuthKubernetesLogin**](Auth.md#PostAuthKubernetesLogin) | **Post** /auth/{kubernetes_mount_path}/login | Authenticates Kubernetes service accounts with Vault.
[**PostAuthKubernetesRoleName**](Auth.md#PostAuthKubernetesRoleName) | **Post** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**PostAuthLdapConfig**](Auth.md#PostAuthLdapConfig) | **Post** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**PostAuthLdapGroupsName**](Auth.md#PostAuthLdapGroupsName) | **Post** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**PostAuthLdapLoginUsername**](Auth.md#PostAuthLdapLoginUsername) | **Post** /auth/{ldap_mount_path}/login/{username} | Log in with a username and password.
[**PostAuthLdapUsersName**](Auth.md#PostAuthLdapUsersName) | **Post** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**PostAuthOciConfig**](Auth.md#PostAuthOciConfig) | **Post** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**PostAuthOciLoginRole**](Auth.md#PostAuthOciLoginRole) | **Post** /auth/{oci_mount_path}/login/{role} | Authenticates to Vault using OCI credentials
[**PostAuthOciRoleRole**](Auth.md#PostAuthOciRoleRole) | **Post** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**PostAuthOidcConfig**](Auth.md#PostAuthOidcConfig) | **Post** /auth/{oidc_mount_path}/config | Configure the JWT authentication backend.
[**PostAuthOidcLogin**](Auth.md#PostAuthOidcLogin) | **Post** /auth/{oidc_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**PostAuthOidcOidcAuthUrl**](Auth.md#PostAuthOidcOidcAuthUrl) | **Post** /auth/{oidc_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**PostAuthOidcOidcCallback**](Auth.md#PostAuthOidcOidcCallback) | **Post** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**PostAuthOidcRoleName**](Auth.md#PostAuthOidcRoleName) | **Post** /auth/{oidc_mount_path}/role/{name} | Register an role with the backend.
[**PostAuthOktaConfig**](Auth.md#PostAuthOktaConfig) | **Post** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**PostAuthOktaGroupsName**](Auth.md#PostAuthOktaGroupsName) | **Post** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**PostAuthOktaLoginUsername**](Auth.md#PostAuthOktaLoginUsername) | **Post** /auth/{okta_mount_path}/login/{username} | Log in with a username and password.
[**PostAuthOktaUsersName**](Auth.md#PostAuthOktaUsersName) | **Post** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**PostAuthRadiusConfig**](Auth.md#PostAuthRadiusConfig) | **Post** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**PostAuthRadiusLogin**](Auth.md#PostAuthRadiusLogin) | **Post** /auth/{radius_mount_path}/login | Log in with a username and password.
[**PostAuthRadiusLoginUrlusername**](Auth.md#PostAuthRadiusLoginUrlusername) | **Post** /auth/{radius_mount_path}/login/{urlusername} | Log in with a username and password.
[**PostAuthRadiusUsersName**](Auth.md#PostAuthRadiusUsersName) | **Post** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**PostAuthTokenCreate**](Auth.md#PostAuthTokenCreate) | **Post** /auth/{token_mount_path}/create | The token create path is used to create new tokens.
[**PostAuthTokenCreateOrphan**](Auth.md#PostAuthTokenCreateOrphan) | **Post** /auth/{token_mount_path}/create-orphan | The token create path is used to create new orphan tokens.
[**PostAuthTokenCreateRoleName**](Auth.md#PostAuthTokenCreateRoleName) | **Post** /auth/{token_mount_path}/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**PostAuthTokenLookup**](Auth.md#PostAuthTokenLookup) | **Post** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**PostAuthTokenLookupAccessor**](Auth.md#PostAuthTokenLookupAccessor) | **Post** /auth/{token_mount_path}/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**PostAuthTokenLookupSelf**](Auth.md#PostAuthTokenLookupSelf) | **Post** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**PostAuthTokenRenew**](Auth.md#PostAuthTokenRenew) | **Post** /auth/{token_mount_path}/renew | This endpoint will renew the given token and prevent expiration.
[**PostAuthTokenRenewAccessor**](Auth.md#PostAuthTokenRenewAccessor) | **Post** /auth/{token_mount_path}/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**PostAuthTokenRenewSelf**](Auth.md#PostAuthTokenRenewSelf) | **Post** /auth/{token_mount_path}/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**PostAuthTokenRevoke**](Auth.md#PostAuthTokenRevoke) | **Post** /auth/{token_mount_path}/revoke | This endpoint will delete the given token and all of its child tokens.
[**PostAuthTokenRevokeAccessor**](Auth.md#PostAuthTokenRevokeAccessor) | **Post** /auth/{token_mount_path}/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**PostAuthTokenRevokeOrphan**](Auth.md#PostAuthTokenRevokeOrphan) | **Post** /auth/{token_mount_path}/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**PostAuthTokenRevokeSelf**](Auth.md#PostAuthTokenRevokeSelf) | **Post** /auth/{token_mount_path}/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**PostAuthTokenRolesRoleName**](Auth.md#PostAuthTokenRolesRoleName) | **Post** /auth/{token_mount_path}/roles/{role_name} | 
[**PostAuthTokenTidy**](Auth.md#PostAuthTokenTidy) | **Post** /auth/{token_mount_path}/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**PostAuthUserpassLoginUsername**](Auth.md#PostAuthUserpassLoginUsername) | **Post** /auth/{userpass_mount_path}/login/{username} | Log in with a username and password.
[**PostAuthUserpassUsersUsername**](Auth.md#PostAuthUserpassUsersUsername) | **Post** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**PostAuthUserpassUsersUsernamePassword**](Auth.md#PostAuthUserpassUsersUsernamePassword) | **Post** /auth/{userpass_mount_path}/users/{username}/password | Reset user&#39;s password.
[**PostAuthUserpassUsersUsernamePolicies**](Auth.md#PostAuthUserpassUsersUsernamePolicies) | **Post** /auth/{userpass_mount_path}/users/{username}/policies | Update the policies associated with the username.



## DeleteAuthAlicloudRoleRole

> DeleteAuthAlicloudRoleRole(ctx, alicloudMountPath, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAlicloudRoleRole(context.Background(), alicloudMountPath, role)
	if err != nil {
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleName

> DeleteAuthApproleRoleRoleName(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleName(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameBindSecretId

> DeleteAuthApproleRoleRoleNameBindSecretId(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameBindSecretId(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameBoundCidrList

> DeleteAuthApproleRoleRoleNameBoundCidrList(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameBoundCidrList(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNamePeriod

> DeleteAuthApproleRoleRoleNamePeriod(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNamePeriod(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNamePolicies

> DeleteAuthApproleRoleRoleNamePolicies(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNamePolicies(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy

> DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, approleMountPath, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs

> DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdDestroy

> DeleteAuthApproleRoleRoleNameSecretIdDestroy(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdNumUses

> DeleteAuthApproleRoleRoleNameSecretIdNumUses(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdTtl

> DeleteAuthApproleRoleRoleNameSecretIdTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameSecretIdTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenBoundCidrs

> DeleteAuthApproleRoleRoleNameTokenBoundCidrs(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenMaxTtl

> DeleteAuthApproleRoleRoleNameTokenMaxTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenNumUses

> DeleteAuthApproleRoleRoleNameTokenNumUses(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameTokenNumUses(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenTtl

> DeleteAuthApproleRoleRoleNameTokenTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthApproleRoleRoleNameTokenTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigCertificateCertName

> DeleteAuthAwsConfigCertificateCertName(ctx, awsMountPath, certName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigCertificateCertName(context.Background(), awsMountPath, certName)
	if err != nil {
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigClient

> DeleteAuthAwsConfigClient(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigClient(context.Background(), awsMountPath)
	if err != nil {
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


## DeleteAuthAwsConfigStsAccountId

> DeleteAuthAwsConfigStsAccountId(ctx, accountId, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigStsAccountId(context.Background(), accountId, awsMountPath)
	if err != nil {
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
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigTidyIdentityAccesslist

> DeleteAuthAwsConfigTidyIdentityAccesslist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigTidyIdentityAccesslist(context.Background(), awsMountPath)
	if err != nil {
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


## DeleteAuthAwsConfigTidyIdentityWhitelist

> DeleteAuthAwsConfigTidyIdentityWhitelist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigTidyIdentityWhitelist(context.Background(), awsMountPath)
	if err != nil {
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


## DeleteAuthAwsConfigTidyRoletagBlacklist

> DeleteAuthAwsConfigTidyRoletagBlacklist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigTidyRoletagBlacklist(context.Background(), awsMountPath)
	if err != nil {
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


## DeleteAuthAwsConfigTidyRoletagDenylist

> DeleteAuthAwsConfigTidyRoletagDenylist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsConfigTidyRoletagDenylist(context.Background(), awsMountPath)
	if err != nil {
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


## DeleteAuthAwsIdentityAccesslistInstanceId

> DeleteAuthAwsIdentityAccesslistInstanceId(ctx, awsMountPath, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsIdentityAccesslistInstanceId(context.Background(), awsMountPath, instanceId)
	if err != nil {
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsIdentityWhitelistInstanceId

> DeleteAuthAwsIdentityWhitelistInstanceId(ctx, awsMountPath, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsIdentityWhitelistInstanceId(context.Background(), awsMountPath, instanceId)
	if err != nil {
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsRoleRole

> DeleteAuthAwsRoleRole(ctx, awsMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsRoleRole(context.Background(), awsMountPath, role)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsRoletagBlacklistRoleTag

> DeleteAuthAwsRoletagBlacklistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsRoletagBlacklistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsRoletagDenylistRoleTag

> DeleteAuthAwsRoletagDenylistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAwsRoletagDenylistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAzureConfig

> DeleteAuthAzureConfig(ctx, azureMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAzureConfig(context.Background(), azureMountPath)
	if err != nil {
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


## DeleteAuthAzureRoleName

> DeleteAuthAzureRoleName(ctx, azureMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthAzureRoleName(context.Background(), azureMountPath, name)
	if err != nil {
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


## DeleteAuthCertCertsName

> DeleteAuthCertCertsName(ctx, certMountPath, name).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthCertCertsName(context.Background(), certMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthCertCrlsName

> DeleteAuthCertCrlsName(ctx, certMountPath, name).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthCertCrlsName(context.Background(), certMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthCfConfig

> DeleteAuthCfConfig(ctx, cfMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthCfConfig(context.Background(), cfMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthCfRolesRole

> DeleteAuthCfRolesRole(ctx, cfMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthCfRolesRole(context.Background(), cfMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthGcpRoleName

> DeleteAuthGcpRoleName(ctx, gcpMountPath, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.DeleteAuthGcpRoleName(context.Background(), gcpMountPath, name)
	if err != nil {
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


## DeleteAuthGithubMapTeamsKey

> DeleteAuthGithubMapTeamsKey(ctx, githubMountPath, key).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthGithubMapTeamsKey(context.Background(), githubMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthGithubMapUsersKey

> DeleteAuthGithubMapUsersKey(ctx, githubMountPath, key).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthGithubMapUsersKey(context.Background(), githubMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthJwtRoleName

> DeleteAuthJwtRoleName(ctx, jwtMountPath, name).Execute()

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
	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthJwtRoleName(context.Background(), jwtMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthKerberosGroupsName

> DeleteAuthKerberosGroupsName(ctx, kerberosMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthKerberosGroupsName(context.Background(), kerberosMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthKubernetesRoleName

> DeleteAuthKubernetesRoleName(ctx, kubernetesMountPath, name).Execute()

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
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthKubernetesRoleName(context.Background(), kubernetesMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthLdapGroupsName

> DeleteAuthLdapGroupsName(ctx, ldapMountPath, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthLdapGroupsName(context.Background(), ldapMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthLdapUsersName

> DeleteAuthLdapUsersName(ctx, ldapMountPath, name).Execute()

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

	name := "name_example" // string | Name of the LDAP user.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthLdapUsersName(context.Background(), ldapMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOciConfig

> DeleteAuthOciConfig(ctx, ociMountPath).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthOciConfig(context.Background(), ociMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOciRoleRole

> DeleteAuthOciRoleRole(ctx, ociMountPath, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthOciRoleRole(context.Background(), ociMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOidcRoleName

> DeleteAuthOidcRoleName(ctx, name, oidcMountPath).Execute()

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
	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthOidcRoleName(context.Background(), name, oidcMountPath)
	if err != nil {
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
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOktaGroupsName

> DeleteAuthOktaGroupsName(ctx, name, oktaMountPath).Execute()

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

	name := "name_example" // string | Name of the Okta group.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthOktaGroupsName(context.Background(), name, oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Okta group. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOktaUsersName

> DeleteAuthOktaUsersName(ctx, name, oktaMountPath).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthOktaUsersName(context.Background(), name, oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the user. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthRadiusUsersName

> DeleteAuthRadiusUsersName(ctx, name, radiusMountPath).Execute()

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
	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthRadiusUsersName(context.Background(), name, radiusMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the RADIUS user. | 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthTokenRolesRoleName

> DeleteAuthTokenRolesRoleName(ctx, roleName, tokenMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthTokenRolesRoleName(context.Background(), roleName, tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role | 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthUserpassUsersUsername

> DeleteAuthUserpassUsersUsername(ctx, username, userpassMountPath).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	resp, err := client.WithToken("my-token").Auth.DeleteAuthUserpassUsersUsername(context.Background(), username, userpassMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username for this user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAlicloudRole

> GetAuthAlicloudRole(ctx, alicloudMountPath).List(list).Execute()

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

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthAlicloudRole(context.Background(), alicloudMountPath, list)
	if err != nil {
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


## GetAuthAlicloudRoleRole

> GetAuthAlicloudRoleRole(ctx, alicloudMountPath, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	resp, err := client.WithToken("my-token").Auth.GetAuthAlicloudRoleRole(context.Background(), alicloudMountPath, role)
	if err != nil {
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAlicloudRoles

> GetAuthAlicloudRoles(ctx, alicloudMountPath).List(list).Execute()

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

	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthAlicloudRoles(context.Background(), alicloudMountPath, list)
	if err != nil {
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


## GetAuthApproleRole

> GetAuthApproleRole(ctx, approleMountPath).List(list).Execute()

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

	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRole(context.Background(), approleMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleName

> GetAuthApproleRoleRoleName(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleName(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameBindSecretId

> GetAuthApproleRoleRoleNameBindSecretId(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameBindSecretId(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameBoundCidrList

> GetAuthApproleRoleRoleNameBoundCidrList(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameBoundCidrList(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameLocalSecretIds

> GetAuthApproleRoleRoleNameLocalSecretIds(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameLocalSecretIds(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNamePeriod

> GetAuthApproleRoleRoleNamePeriod(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNamePeriod(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNamePolicies

> GetAuthApproleRoleRoleNamePolicies(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNamePolicies(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameRoleId

> GetAuthApproleRoleRoleNameRoleId(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameRoleId(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretId

> GetAuthApproleRoleRoleNameSecretId(ctx, approleMountPath, roleName).List(list).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameSecretId(context.Background(), approleMountPath, roleName, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdBoundCidrs

> GetAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdNumUses

> GetAuthApproleRoleRoleNameSecretIdNumUses(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdTtl

> GetAuthApproleRoleRoleNameSecretIdTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameSecretIdTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenBoundCidrs

> GetAuthApproleRoleRoleNameTokenBoundCidrs(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenMaxTtl

> GetAuthApproleRoleRoleNameTokenMaxTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenNumUses

> GetAuthApproleRoleRoleNameTokenNumUses(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameTokenNumUses(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenTtl

> GetAuthApproleRoleRoleNameTokenTtl(ctx, approleMountPath, roleName).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.GetAuthApproleRoleRoleNameTokenTtl(context.Background(), approleMountPath, roleName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigCertificateCertName

> GetAuthAwsConfigCertificateCertName(ctx, awsMountPath, certName).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigCertificateCertName(context.Background(), awsMountPath, certName)
	if err != nil {
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigCertificates

> GetAuthAwsConfigCertificates(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigCertificates(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsConfigClient

> GetAuthAwsConfigClient(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigClient(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsConfigIdentity

> GetAuthAwsConfigIdentity(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigIdentity(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsConfigSts

> GetAuthAwsConfigSts(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigSts(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsConfigStsAccountId

> GetAuthAwsConfigStsAccountId(ctx, accountId, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigStsAccountId(context.Background(), accountId, awsMountPath)
	if err != nil {
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
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigTidyIdentityAccesslist

> GetAuthAwsConfigTidyIdentityAccesslist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigTidyIdentityAccesslist(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsConfigTidyIdentityWhitelist

> GetAuthAwsConfigTidyIdentityWhitelist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigTidyIdentityWhitelist(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsConfigTidyRoletagBlacklist

> GetAuthAwsConfigTidyRoletagBlacklist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigTidyRoletagBlacklist(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsConfigTidyRoletagDenylist

> GetAuthAwsConfigTidyRoletagDenylist(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsConfigTidyRoletagDenylist(context.Background(), awsMountPath)
	if err != nil {
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


## GetAuthAwsIdentityAccesslist

> GetAuthAwsIdentityAccesslist(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsIdentityAccesslist(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsIdentityAccesslistInstanceId

> GetAuthAwsIdentityAccesslistInstanceId(ctx, awsMountPath, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsIdentityAccesslistInstanceId(context.Background(), awsMountPath, instanceId)
	if err != nil {
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsIdentityWhitelist

> GetAuthAwsIdentityWhitelist(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsIdentityWhitelist(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsIdentityWhitelistInstanceId

> GetAuthAwsIdentityWhitelistInstanceId(ctx, awsMountPath, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsIdentityWhitelistInstanceId(context.Background(), awsMountPath, instanceId)
	if err != nil {
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsRole

> GetAuthAwsRole(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRole(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsRoleRole

> GetAuthAwsRoleRole(ctx, awsMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoleRole(context.Background(), awsMountPath, role)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsRoles

> GetAuthAwsRoles(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoles(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsRoletagBlacklist

> GetAuthAwsRoletagBlacklist(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoletagBlacklist(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsRoletagBlacklistRoleTag

> GetAuthAwsRoletagBlacklistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoletagBlacklistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsRoletagDenylist

> GetAuthAwsRoletagDenylist(ctx, awsMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoletagDenylist(context.Background(), awsMountPath, list)
	if err != nil {
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


## GetAuthAwsRoletagDenylistRoleTag

> GetAuthAwsRoletagDenylistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.GetAuthAwsRoletagDenylistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAzureConfig

> GetAuthAzureConfig(ctx, azureMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAzureConfig(context.Background(), azureMountPath)
	if err != nil {
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


## GetAuthAzureRole

> GetAuthAzureRole(ctx, azureMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthAzureRole(context.Background(), azureMountPath, list)
	if err != nil {
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


## GetAuthAzureRoleName

> GetAuthAzureRoleName(ctx, azureMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthAzureRoleName(context.Background(), azureMountPath, name)
	if err != nil {
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


## GetAuthCentrifyConfig

> GetAuthCentrifyConfig(ctx, centrifyMountPath).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	centrifyMountPath := "centrifyMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "centrify")

	resp, err := client.WithToken("my-token").Auth.GetAuthCentrifyConfig(context.Background(), centrifyMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**centrifyMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;centrify&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCertCerts

> GetAuthCertCerts(ctx, certMountPath).List(list).Execute()

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

	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthCertCerts(context.Background(), certMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCertCertsName

> GetAuthCertCertsName(ctx, certMountPath, name).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.GetAuthCertCertsName(context.Background(), certMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCertConfig

> GetAuthCertConfig(ctx, certMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.GetAuthCertConfig(context.Background(), certMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCertCrlsName

> GetAuthCertCrlsName(ctx, certMountPath, name).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	resp, err := client.WithToken("my-token").Auth.GetAuthCertCrlsName(context.Background(), certMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCfConfig

> GetAuthCfConfig(ctx, cfMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.GetAuthCfConfig(context.Background(), cfMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCfRoles

> GetAuthCfRoles(ctx, cfMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthCfRoles(context.Background(), cfMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCfRolesRole

> GetAuthCfRolesRole(ctx, cfMountPath, role).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	resp, err := client.WithToken("my-token").Auth.GetAuthCfRolesRole(context.Background(), cfMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGcpConfig

> GetAuthGcpConfig(ctx, gcpMountPath).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthGcpConfig(context.Background(), gcpMountPath)
	if err != nil {
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


## GetAuthGcpRole

> GetAuthGcpRole(ctx, gcpMountPath).List(list).Execute()

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

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthGcpRole(context.Background(), gcpMountPath, list)
	if err != nil {
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


## GetAuthGcpRoleName

> GetAuthGcpRoleName(ctx, gcpMountPath, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthGcpRoleName(context.Background(), gcpMountPath, name)
	if err != nil {
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


## GetAuthGcpRoles

> GetAuthGcpRoles(ctx, gcpMountPath).List(list).Execute()

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

	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthGcpRoles(context.Background(), gcpMountPath, list)
	if err != nil {
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


## GetAuthGithubConfig

> GetAuthGithubConfig(ctx, githubMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.GetAuthGithubConfig(context.Background(), githubMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapTeams

> GetAuthGithubMapTeams(ctx, githubMountPath).List(list).Execute()

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

	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthGithubMapTeams(context.Background(), githubMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | [default to &quot;false&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapTeamsKey

> GetAuthGithubMapTeamsKey(ctx, githubMountPath, key).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.GetAuthGithubMapTeamsKey(context.Background(), githubMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapUsers

> GetAuthGithubMapUsers(ctx, githubMountPath).List(list).Execute()

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

	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthGithubMapUsers(context.Background(), githubMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | [default to &quot;false&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapUsersKey

> GetAuthGithubMapUsersKey(ctx, githubMountPath, key).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	resp, err := client.WithToken("my-token").Auth.GetAuthGithubMapUsersKey(context.Background(), githubMountPath, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthJwtConfig

> GetAuthJwtConfig(ctx, jwtMountPath).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	resp, err := client.WithToken("my-token").Auth.GetAuthJwtConfig(context.Background(), jwtMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthJwtOidcCallback

> GetAuthJwtOidcCallback(ctx, jwtMountPath).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	resp, err := client.WithToken("my-token").Auth.GetAuthJwtOidcCallback(context.Background(), jwtMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthJwtRole

> GetAuthJwtRole(ctx, jwtMountPath).List(list).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthJwtRole(context.Background(), jwtMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthJwtRoleName

> GetAuthJwtRoleName(ctx, jwtMountPath, name).Execute()

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
	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	resp, err := client.WithToken("my-token").Auth.GetAuthJwtRoleName(context.Background(), jwtMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosConfig

> GetAuthKerberosConfig(ctx, kerberosMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.GetAuthKerberosConfig(context.Background(), kerberosMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosConfigLdap

> GetAuthKerberosConfigLdap(ctx, kerberosMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.GetAuthKerberosConfigLdap(context.Background(), kerberosMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosGroups

> GetAuthKerberosGroups(ctx, kerberosMountPath).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthKerberosGroups(context.Background(), kerberosMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosGroupsName

> GetAuthKerberosGroupsName(ctx, kerberosMountPath, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.GetAuthKerberosGroupsName(context.Background(), kerberosMountPath, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosLogin

> GetAuthKerberosLogin(ctx, kerberosMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	resp, err := client.WithToken("my-token").Auth.GetAuthKerberosLogin(context.Background(), kerberosMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKubernetesConfig

> GetAuthKubernetesConfig(ctx, kubernetesMountPath).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthKubernetesConfig(context.Background(), kubernetesMountPath)
	if err != nil {
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


## GetAuthKubernetesRole

> GetAuthKubernetesRole(ctx, kubernetesMountPath).List(list).Execute()

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

	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthKubernetesRole(context.Background(), kubernetesMountPath, list)
	if err != nil {
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


## GetAuthKubernetesRoleName

> GetAuthKubernetesRoleName(ctx, kubernetesMountPath, name).Execute()

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
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	resp, err := client.WithToken("my-token").Auth.GetAuthKubernetesRoleName(context.Background(), kubernetesMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthLdapConfig

> GetAuthLdapConfig(ctx, ldapMountPath).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.GetAuthLdapConfig(context.Background(), ldapMountPath)
	if err != nil {
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


## GetAuthLdapGroups

> GetAuthLdapGroups(ctx, ldapMountPath).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.GetAuthLdapGroups(context.Background(), ldapMountPath, list)
	if err != nil {
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


## GetAuthLdapGroupsName

> GetAuthLdapGroupsName(ctx, ldapMountPath, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Auth.GetAuthLdapGroupsName(context.Background(), ldapMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthLdapUsers

> GetAuthLdapUsers(ctx, ldapMountPath).List(list).Execute()

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

	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthLdapUsers(context.Background(), ldapMountPath, list)
	if err != nil {
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


## GetAuthLdapUsersName

> GetAuthLdapUsersName(ctx, ldapMountPath, name).Execute()

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

	name := "name_example" // string | Name of the LDAP user.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	resp, err := client.WithToken("my-token").Auth.GetAuthLdapUsersName(context.Background(), ldapMountPath, name)
	if err != nil {
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOciConfig

> GetAuthOciConfig(ctx, ociMountPath).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	resp, err := client.WithToken("my-token").Auth.GetAuthOciConfig(context.Background(), ociMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOciRole

> GetAuthOciRole(ctx, ociMountPath).List(list).Execute()

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

	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthOciRole(context.Background(), ociMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOciRoleRole

> GetAuthOciRoleRole(ctx, ociMountPath, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	resp, err := client.WithToken("my-token").Auth.GetAuthOciRoleRole(context.Background(), ociMountPath, role)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOidcConfig

> GetAuthOidcConfig(ctx, oidcMountPath).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.GetAuthOidcConfig(context.Background(), oidcMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOidcOidcCallback

> GetAuthOidcOidcCallback(ctx, oidcMountPath).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.GetAuthOidcOidcCallback(context.Background(), oidcMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOidcRole

> GetAuthOidcRole(ctx, oidcMountPath).List(list).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthOidcRole(context.Background(), oidcMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOidcRoleName

> GetAuthOidcRoleName(ctx, name, oidcMountPath).Execute()

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
	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	resp, err := client.WithToken("my-token").Auth.GetAuthOidcRoleName(context.Background(), name, oidcMountPath)
	if err != nil {
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
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaConfig

> GetAuthOktaConfig(ctx, oktaMountPath).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.GetAuthOktaConfig(context.Background(), oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaGroups

> GetAuthOktaGroups(ctx, oktaMountPath).List(list).Execute()

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

	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthOktaGroups(context.Background(), oktaMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaGroupsName

> GetAuthOktaGroupsName(ctx, name, oktaMountPath).Execute()

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

	name := "name_example" // string | Name of the Okta group.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.GetAuthOktaGroupsName(context.Background(), name, oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Okta group. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaUsers

> GetAuthOktaUsers(ctx, oktaMountPath).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthOktaUsers(context.Background(), oktaMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaUsersName

> GetAuthOktaUsersName(ctx, name, oktaMountPath).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.GetAuthOktaUsersName(context.Background(), name, oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the user. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaVerifyNonce

> GetAuthOktaVerifyNonce(ctx, nonce, oktaMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	resp, err := client.WithToken("my-token").Auth.GetAuthOktaVerifyNonce(context.Background(), nonce, oktaMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**nonce** | **string** | Nonce provided during a login request to retrieve the number verification challenge for the matching request. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthRadiusConfig

> GetAuthRadiusConfig(ctx, radiusMountPath).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	resp, err := client.WithToken("my-token").Auth.GetAuthRadiusConfig(context.Background(), radiusMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthRadiusUsers

> GetAuthRadiusUsers(ctx, radiusMountPath).List(list).Execute()

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

	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthRadiusUsers(context.Background(), radiusMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthRadiusUsersName

> GetAuthRadiusUsersName(ctx, name, radiusMountPath).Execute()

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
	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	resp, err := client.WithToken("my-token").Auth.GetAuthRadiusUsersName(context.Background(), name, radiusMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the RADIUS user. | 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenAccessors

> GetAuthTokenAccessors(ctx, tokenMountPath).List(list).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthTokenAccessors(context.Background(), tokenMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenLookup

> GetAuthTokenLookup(ctx, tokenMountPath).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.GetAuthTokenLookup(context.Background(), tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenLookupSelf

> GetAuthTokenLookupSelf(ctx, tokenMountPath).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.GetAuthTokenLookupSelf(context.Background(), tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenRoles

> GetAuthTokenRoles(ctx, tokenMountPath).List(list).Execute()

This endpoint lists configured roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthTokenRoles(context.Background(), tokenMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenRolesRoleName

> GetAuthTokenRolesRoleName(ctx, roleName, tokenMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.GetAuthTokenRolesRoleName(context.Background(), roleName, tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role | 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthUserpassUsers

> GetAuthUserpassUsers(ctx, userpassMountPath).List(list).Execute()

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

	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.GetAuthUserpassUsers(context.Background(), userpassMountPath, list)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthUserpassUsersUsername

> GetAuthUserpassUsersUsername(ctx, username, userpassMountPath).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	resp, err := client.WithToken("my-token").Auth.GetAuthUserpassUsersUsername(context.Background(), username, userpassMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username for this user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAlicloudLogin

> PostAuthAlicloudLogin(ctx, alicloudMountPath).AlicloudLoginRequest(alicloudLoginRequest).Execute()

Authenticates an RAM entity with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	alicloudLoginRequest := NewAlicloudLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAlicloudLogin(context.Background(), alicloudMountPath, alicloudLoginRequest)
	if err != nil {
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
 **alicloudLoginRequest** | [**AlicloudLoginRequest**](AlicloudLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAlicloudRoleRole

> PostAuthAlicloudRoleRole(ctx, alicloudMountPath, role).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	alicloudMountPath := "alicloudMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "alicloud")

	alicloudRoleRequest := NewAlicloudRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAlicloudRoleRole(context.Background(), alicloudMountPath, role, alicloudRoleRequest)
	if err != nil {
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alicloudRoleRequest** | [**AlicloudRoleRequest**](AlicloudRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleLogin

> PostAuthApproleLogin(ctx, approleMountPath).ApproleLoginRequest(approleLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleLoginRequest := NewApproleLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleLogin(context.Background(), approleMountPath, approleLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approleLoginRequest** | [**ApproleLoginRequest**](ApproleLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleName

> PostAuthApproleRoleRoleName(ctx, approleMountPath, roleName).ApproleRoleRequest(approleRoleRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleRequest := NewApproleRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleName(context.Background(), approleMountPath, roleName, approleRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRequest** | [**ApproleRoleRequest**](ApproleRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameBindSecretId

> PostAuthApproleRoleRoleNameBindSecretId(ctx, approleMountPath, roleName).ApproleRoleBindSecretIdRequest(approleRoleBindSecretIdRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleBindSecretIdRequest := NewApproleRoleBindSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameBindSecretId(context.Background(), approleMountPath, roleName, approleRoleBindSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBindSecretIdRequest** | [**ApproleRoleBindSecretIdRequest**](ApproleRoleBindSecretIdRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameBoundCidrList

> PostAuthApproleRoleRoleNameBoundCidrList(ctx, approleMountPath, roleName).ApproleRoleBoundCidrListRequest(approleRoleBoundCidrListRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleBoundCidrListRequest := NewApproleRoleBoundCidrListRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameBoundCidrList(context.Background(), approleMountPath, roleName, approleRoleBoundCidrListRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBoundCidrListRequest** | [**ApproleRoleBoundCidrListRequest**](ApproleRoleBoundCidrListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameCustomSecretId

> PostAuthApproleRoleRoleNameCustomSecretId(ctx, approleMountPath, roleName).ApproleRoleCustomSecretIdRequest(approleRoleCustomSecretIdRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleCustomSecretIdRequest := NewApproleRoleCustomSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameCustomSecretId(context.Background(), approleMountPath, roleName, approleRoleCustomSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleCustomSecretIdRequest** | [**ApproleRoleCustomSecretIdRequest**](ApproleRoleCustomSecretIdRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNamePeriod

> PostAuthApproleRoleRoleNamePeriod(ctx, approleMountPath, roleName).ApproleRolePeriodRequest(approleRolePeriodRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRolePeriodRequest := NewApproleRolePeriodRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNamePeriod(context.Background(), approleMountPath, roleName, approleRolePeriodRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePeriodRequest** | [**ApproleRolePeriodRequest**](ApproleRolePeriodRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNamePolicies

> PostAuthApproleRoleRoleNamePolicies(ctx, approleMountPath, roleName).ApproleRolePoliciesRequest(approleRolePoliciesRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRolePoliciesRequest := NewApproleRolePoliciesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNamePolicies(context.Background(), approleMountPath, roleName, approleRolePoliciesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePoliciesRequest** | [**ApproleRolePoliciesRequest**](ApproleRolePoliciesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameRoleId

> PostAuthApproleRoleRoleNameRoleId(ctx, approleMountPath, roleName).ApproleRoleRoleIdRequest(approleRoleRoleIdRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleRoleIdRequest := NewApproleRoleRoleIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameRoleId(context.Background(), approleMountPath, roleName, approleRoleRoleIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRoleIdRequest** | [**ApproleRoleRoleIdRequest**](ApproleRoleRoleIdRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretId

> PostAuthApproleRoleRoleNameSecretId(ctx, approleMountPath, roleName).ApproleRoleSecretIdRequest(approleRoleSecretIdRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdRequest := NewApproleRoleSecretIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretId(context.Background(), approleMountPath, roleName, approleRoleSecretIdRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdRequest** | [**ApproleRoleSecretIdRequest**](ApproleRoleSecretIdRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdAccessorDestroy

> PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, approleMountPath, roleName).ApproleRoleSecretIdAccessorDestroyRequest(approleRoleSecretIdAccessorDestroyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdAccessorDestroyRequest := NewApproleRoleSecretIdAccessorDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), approleMountPath, roleName, approleRoleSecretIdAccessorDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorDestroyRequest** | [**ApproleRoleSecretIdAccessorDestroyRequest**](ApproleRoleSecretIdAccessorDestroyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdAccessorLookup

> PostAuthApproleRoleRoleNameSecretIdAccessorLookup(ctx, approleMountPath, roleName).ApproleRoleSecretIdAccessorLookupRequest(approleRoleSecretIdAccessorLookupRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdAccessorLookupRequest := NewApproleRoleSecretIdAccessorLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdAccessorLookup(context.Background(), approleMountPath, roleName, approleRoleSecretIdAccessorLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorLookupRequest** | [**ApproleRoleSecretIdAccessorLookupRequest**](ApproleRoleSecretIdAccessorLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdBoundCidrs

> PostAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, approleMountPath, roleName).ApproleRoleSecretIdBoundCidrsRequest(approleRoleSecretIdBoundCidrsRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdBoundCidrsRequest := NewApproleRoleSecretIdBoundCidrsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), approleMountPath, roleName, approleRoleSecretIdBoundCidrsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdBoundCidrsRequest** | [**ApproleRoleSecretIdBoundCidrsRequest**](ApproleRoleSecretIdBoundCidrsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdDestroy

> PostAuthApproleRoleRoleNameSecretIdDestroy(ctx, approleMountPath, roleName).ApproleRoleSecretIdDestroyRequest(approleRoleSecretIdDestroyRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdDestroyRequest := NewApproleRoleSecretIdDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), approleMountPath, roleName, approleRoleSecretIdDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdDestroyRequest** | [**ApproleRoleSecretIdDestroyRequest**](ApproleRoleSecretIdDestroyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdLookup

> PostAuthApproleRoleRoleNameSecretIdLookup(ctx, approleMountPath, roleName).ApproleRoleSecretIdLookupRequest(approleRoleSecretIdLookupRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdLookupRequest := NewApproleRoleSecretIdLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdLookup(context.Background(), approleMountPath, roleName, approleRoleSecretIdLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdLookupRequest** | [**ApproleRoleSecretIdLookupRequest**](ApproleRoleSecretIdLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdNumUses

> PostAuthApproleRoleRoleNameSecretIdNumUses(ctx, approleMountPath, roleName).ApproleRoleSecretIdNumUsesRequest(approleRoleSecretIdNumUsesRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdNumUsesRequest := NewApproleRoleSecretIdNumUsesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), approleMountPath, roleName, approleRoleSecretIdNumUsesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdNumUsesRequest** | [**ApproleRoleSecretIdNumUsesRequest**](ApproleRoleSecretIdNumUsesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdTtl

> PostAuthApproleRoleRoleNameSecretIdTtl(ctx, approleMountPath, roleName).ApproleRoleSecretIdTtlRequest(approleRoleSecretIdTtlRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleSecretIdTtlRequest := NewApproleRoleSecretIdTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameSecretIdTtl(context.Background(), approleMountPath, roleName, approleRoleSecretIdTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdTtlRequest** | [**ApproleRoleSecretIdTtlRequest**](ApproleRoleSecretIdTtlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenBoundCidrs

> PostAuthApproleRoleRoleNameTokenBoundCidrs(ctx, approleMountPath, roleName).ApproleRoleTokenBoundCidrsRequest(approleRoleTokenBoundCidrsRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenBoundCidrsRequest := NewApproleRoleTokenBoundCidrsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), approleMountPath, roleName, approleRoleTokenBoundCidrsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenBoundCidrsRequest** | [**ApproleRoleTokenBoundCidrsRequest**](ApproleRoleTokenBoundCidrsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenMaxTtl

> PostAuthApproleRoleRoleNameTokenMaxTtl(ctx, approleMountPath, roleName).ApproleRoleTokenMaxTtlRequest(approleRoleTokenMaxTtlRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenMaxTtlRequest := NewApproleRoleTokenMaxTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), approleMountPath, roleName, approleRoleTokenMaxTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenMaxTtlRequest** | [**ApproleRoleTokenMaxTtlRequest**](ApproleRoleTokenMaxTtlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenNumUses

> PostAuthApproleRoleRoleNameTokenNumUses(ctx, approleMountPath, roleName).ApproleRoleTokenNumUsesRequest(approleRoleTokenNumUsesRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenNumUsesRequest := NewApproleRoleTokenNumUsesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameTokenNumUses(context.Background(), approleMountPath, roleName, approleRoleTokenNumUsesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenNumUsesRequest** | [**ApproleRoleTokenNumUsesRequest**](ApproleRoleTokenNumUsesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenTtl

> PostAuthApproleRoleRoleNameTokenTtl(ctx, approleMountPath, roleName).ApproleRoleTokenTtlRequest(approleRoleTokenTtlRequest).Execute()

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
	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	approleRoleTokenTtlRequest := NewApproleRoleTokenTtlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthApproleRoleRoleNameTokenTtl(context.Background(), approleMountPath, roleName, approleRoleTokenTtlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenTtlRequest** | [**ApproleRoleTokenTtlRequest**](ApproleRoleTokenTtlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleTidySecretId

> PostAuthApproleTidySecretId(ctx, approleMountPath).Execute()

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

	approleMountPath := "approleMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "approle")

	resp, err := client.WithToken("my-token").Auth.PostAuthApproleTidySecretId(context.Background(), approleMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**approleMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;approle&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigCertificateCertName

> PostAuthAwsConfigCertificateCertName(ctx, awsMountPath, certName).AwsConfigCertificateRequest(awsConfigCertificateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigCertificateRequest := NewAwsConfigCertificateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigCertificateCertName(context.Background(), awsMountPath, certName, awsConfigCertificateRequest)
	if err != nil {
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigCertificateRequest** | [**AwsConfigCertificateRequest**](AwsConfigCertificateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigClient

> PostAuthAwsConfigClient(ctx, awsMountPath).AwsConfigClientRequest(awsConfigClientRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigClientRequest := NewAwsConfigClientRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigClient(context.Background(), awsMountPath, awsConfigClientRequest)
	if err != nil {
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
 **awsConfigClientRequest** | [**AwsConfigClientRequest**](AwsConfigClientRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigIdentity

> PostAuthAwsConfigIdentity(ctx, awsMountPath).AwsConfigIdentityRequest(awsConfigIdentityRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigIdentityRequest := NewAwsConfigIdentityRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigIdentity(context.Background(), awsMountPath, awsConfigIdentityRequest)
	if err != nil {
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
 **awsConfigIdentityRequest** | [**AwsConfigIdentityRequest**](AwsConfigIdentityRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigRotateRoot

> PostAuthAwsConfigRotateRoot(ctx, awsMountPath).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigRotateRoot(context.Background(), awsMountPath)
	if err != nil {
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


## PostAuthAwsConfigStsAccountId

> PostAuthAwsConfigStsAccountId(ctx, accountId, awsMountPath).AwsConfigStsRequest(awsConfigStsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsConfigStsRequest := NewAwsConfigStsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigStsAccountId(context.Background(), accountId, awsMountPath, awsConfigStsRequest)
	if err != nil {
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
**awsMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;aws&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigStsRequest** | [**AwsConfigStsRequest**](AwsConfigStsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyIdentityAccesslist

> PostAuthAwsConfigTidyIdentityAccesslist(ctx, awsMountPath).AwsConfigTidyIdentityAccesslistRequest(awsConfigTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyIdentityAccesslistRequest := NewAwsConfigTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigTidyIdentityAccesslist(context.Background(), awsMountPath, awsConfigTidyIdentityAccesslistRequest)
	if err != nil {
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
 **awsConfigTidyIdentityAccesslistRequest** | [**AwsConfigTidyIdentityAccesslistRequest**](AwsConfigTidyIdentityAccesslistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyIdentityWhitelist

> PostAuthAwsConfigTidyIdentityWhitelist(ctx, awsMountPath).AwsConfigTidyIdentityWhitelistRequest(awsConfigTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyIdentityWhitelistRequest := NewAwsConfigTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigTidyIdentityWhitelist(context.Background(), awsMountPath, awsConfigTidyIdentityWhitelistRequest)
	if err != nil {
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
 **awsConfigTidyIdentityWhitelistRequest** | [**AwsConfigTidyIdentityWhitelistRequest**](AwsConfigTidyIdentityWhitelistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyRoletagBlacklist

> PostAuthAwsConfigTidyRoletagBlacklist(ctx, awsMountPath).AwsConfigTidyRoletagBlacklistRequest(awsConfigTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyRoletagBlacklistRequest := NewAwsConfigTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigTidyRoletagBlacklist(context.Background(), awsMountPath, awsConfigTidyRoletagBlacklistRequest)
	if err != nil {
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
 **awsConfigTidyRoletagBlacklistRequest** | [**AwsConfigTidyRoletagBlacklistRequest**](AwsConfigTidyRoletagBlacklistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyRoletagDenylist

> PostAuthAwsConfigTidyRoletagDenylist(ctx, awsMountPath).AwsConfigTidyRoletagDenylistRequest(awsConfigTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsConfigTidyRoletagDenylistRequest := NewAwsConfigTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsConfigTidyRoletagDenylist(context.Background(), awsMountPath, awsConfigTidyRoletagDenylistRequest)
	if err != nil {
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
 **awsConfigTidyRoletagDenylistRequest** | [**AwsConfigTidyRoletagDenylistRequest**](AwsConfigTidyRoletagDenylistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsLogin

> PostAuthAwsLogin(ctx, awsMountPath).AwsLoginRequest(awsLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsLoginRequest := NewAwsLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsLogin(context.Background(), awsMountPath, awsLoginRequest)
	if err != nil {
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
 **awsLoginRequest** | [**AwsLoginRequest**](AwsLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoleRole

> PostAuthAwsRoleRole(ctx, awsMountPath, role).AwsRoleRequest(awsRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRoleRequest := NewAwsRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsRoleRole(context.Background(), awsMountPath, role, awsRoleRequest)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRoleRequest** | [**AwsRoleRequest**](AwsRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoleRoleTag

> PostAuthAwsRoleRoleTag(ctx, awsMountPath, role).AwsRoleTagRequest(awsRoleTagRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	awsRoleTagRequest := NewAwsRoleTagRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsRoleRoleTag(context.Background(), awsMountPath, role, awsRoleTagRequest)
	if err != nil {
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRoleTagRequest** | [**AwsRoleTagRequest**](AwsRoleTagRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoletagBlacklistRoleTag

> PostAuthAwsRoletagBlacklistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.PostAuthAwsRoletagBlacklistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoletagDenylistRoleTag

> PostAuthAwsRoletagDenylistRoleTag(ctx, awsMountPath, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	awsMountPath := "awsMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "aws")

	resp, err := client.WithToken("my-token").Auth.PostAuthAwsRoletagDenylistRoleTag(context.Background(), awsMountPath, roleTag)
	if err != nil {
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyIdentityAccesslist

> PostAuthAwsTidyIdentityAccesslist(ctx, awsMountPath).AwsTidyIdentityAccesslistRequest(awsTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyIdentityAccesslistRequest := NewAwsTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsTidyIdentityAccesslist(context.Background(), awsMountPath, awsTidyIdentityAccesslistRequest)
	if err != nil {
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
 **awsTidyIdentityAccesslistRequest** | [**AwsTidyIdentityAccesslistRequest**](AwsTidyIdentityAccesslistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyIdentityWhitelist

> PostAuthAwsTidyIdentityWhitelist(ctx, awsMountPath).AwsTidyIdentityWhitelistRequest(awsTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyIdentityWhitelistRequest := NewAwsTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsTidyIdentityWhitelist(context.Background(), awsMountPath, awsTidyIdentityWhitelistRequest)
	if err != nil {
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
 **awsTidyIdentityWhitelistRequest** | [**AwsTidyIdentityWhitelistRequest**](AwsTidyIdentityWhitelistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyRoletagBlacklist

> PostAuthAwsTidyRoletagBlacklist(ctx, awsMountPath).AwsTidyRoletagBlacklistRequest(awsTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyRoletagBlacklistRequest := NewAwsTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsTidyRoletagBlacklist(context.Background(), awsMountPath, awsTidyRoletagBlacklistRequest)
	if err != nil {
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
 **awsTidyRoletagBlacklistRequest** | [**AwsTidyRoletagBlacklistRequest**](AwsTidyRoletagBlacklistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyRoletagDenylist

> PostAuthAwsTidyRoletagDenylist(ctx, awsMountPath).AwsTidyRoletagDenylistRequest(awsTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	awsTidyRoletagDenylistRequest := NewAwsTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAwsTidyRoletagDenylist(context.Background(), awsMountPath, awsTidyRoletagDenylistRequest)
	if err != nil {
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
 **awsTidyRoletagDenylistRequest** | [**AwsTidyRoletagDenylistRequest**](AwsTidyRoletagDenylistRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAzureConfig

> PostAuthAzureConfig(ctx, azureMountPath).AzureConfigRequest(azureConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.PostAuthAzureConfig(context.Background(), azureMountPath, azureConfigRequest)
	if err != nil {
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


## PostAuthAzureLogin

> PostAuthAzureLogin(ctx, azureMountPath).AzureLoginRequest(azureLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	azureLoginRequest := NewAzureLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAzureLogin(context.Background(), azureMountPath, azureLoginRequest)
	if err != nil {
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
 **azureLoginRequest** | [**AzureLoginRequest**](AzureLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAzureRoleName

> PostAuthAzureRoleName(ctx, azureMountPath, name).AzureRoleRequest(azureRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	azureRoleRequest := NewAzureRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthAzureRoleName(context.Background(), azureMountPath, name, azureRoleRequest)
	if err != nil {
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

 **azureRoleRequest** | [**AzureRoleRequest**](AzureRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCentrifyConfig

> PostAuthCentrifyConfig(ctx, centrifyMountPath).CentrifyConfigRequest(centrifyConfigRequest).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	centrifyMountPath := "centrifyMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "centrify")

	centrifyConfigRequest := NewCentrifyConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCentrifyConfig(context.Background(), centrifyMountPath, centrifyConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**centrifyMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;centrify&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **centrifyConfigRequest** | [**CentrifyConfigRequest**](CentrifyConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCentrifyLogin

> PostAuthCentrifyLogin(ctx, centrifyMountPath).CentrifyLoginRequest(centrifyLoginRequest).Execute()

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

	centrifyMountPath := "centrifyMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "centrify")

	centrifyLoginRequest := NewCentrifyLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCentrifyLogin(context.Background(), centrifyMountPath, centrifyLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**centrifyMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;centrify&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **centrifyLoginRequest** | [**CentrifyLoginRequest**](CentrifyLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertCertsName

> PostAuthCertCertsName(ctx, certMountPath, name).CertCertsRequest(certCertsRequest).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certCertsRequest := NewCertCertsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCertCertsName(context.Background(), certMountPath, name, certCertsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCertsRequest** | [**CertCertsRequest**](CertCertsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertConfig

> PostAuthCertConfig(ctx, certMountPath).CertConfigRequest(certConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certConfigRequest := NewCertConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCertConfig(context.Background(), certMountPath, certConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certConfigRequest** | [**CertConfigRequest**](CertConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertCrlsName

> PostAuthCertCrlsName(ctx, certMountPath, name).CertCrlsRequest(certCrlsRequest).Execute()

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
	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certCrlsRequest := NewCertCrlsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCertCrlsName(context.Background(), certMountPath, name, certCrlsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCrlsRequest** | [**CertCrlsRequest**](CertCrlsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertLogin

> PostAuthCertLogin(ctx, certMountPath).CertLoginRequest(certLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	certMountPath := "certMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cert")

	certLoginRequest := NewCertLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCertLogin(context.Background(), certMountPath, certLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**certMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cert&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certLoginRequest** | [**CertLoginRequest**](CertLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfConfig

> PostAuthCfConfig(ctx, cfMountPath).CfConfigRequest(cfConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	cfConfigRequest := NewCfConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCfConfig(context.Background(), cfMountPath, cfConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cfConfigRequest** | [**CfConfigRequest**](CfConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfLogin

> PostAuthCfLogin(ctx, cfMountPath).CfLoginRequest(cfLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	cfLoginRequest := NewCfLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCfLogin(context.Background(), cfMountPath, cfLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cfLoginRequest** | [**CfLoginRequest**](CfLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfRolesRole

> PostAuthCfRolesRole(ctx, cfMountPath, role).CfRolesRequest(cfRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	cfMountPath := "cfMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "cf")

	cfRolesRequest := NewCfRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthCfRolesRole(context.Background(), cfMountPath, role, cfRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**cfMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;cf&quot;]
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cfRolesRequest** | [**CfRolesRequest**](CfRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpConfig

> PostAuthGcpConfig(ctx, gcpMountPath).GcpConfigRequest(gcpConfigRequest).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.PostAuthGcpConfig(context.Background(), gcpMountPath, gcpConfigRequest)
	if err != nil {
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


## PostAuthGcpLogin

> PostAuthGcpLogin(ctx, gcpMountPath).GcpLoginRequest(gcpLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	gcpLoginRequest := NewGcpLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGcpLogin(context.Background(), gcpMountPath, gcpLoginRequest)
	if err != nil {
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
 **gcpLoginRequest** | [**GcpLoginRequest**](GcpLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleName

> PostAuthGcpRoleName(ctx, gcpMountPath, name).GcpRoleRequest(gcpRoleRequest).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	gcpRoleRequest := NewGcpRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGcpRoleName(context.Background(), gcpMountPath, name, gcpRoleRequest)
	if err != nil {
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

 **gcpRoleRequest** | [**GcpRoleRequest**](GcpRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleNameLabels

> PostAuthGcpRoleNameLabels(ctx, gcpMountPath, name).GcpRoleLabelsRequest(gcpRoleLabelsRequest).Execute()

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
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRoleLabelsRequest := NewGcpRoleLabelsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGcpRoleNameLabels(context.Background(), gcpMountPath, name, gcpRoleLabelsRequest)
	if err != nil {
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

 **gcpRoleLabelsRequest** | [**GcpRoleLabelsRequest**](GcpRoleLabelsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleNameServiceAccounts

> PostAuthGcpRoleNameServiceAccounts(ctx, gcpMountPath, name).GcpRoleServiceAccountsRequest(gcpRoleServiceAccountsRequest).Execute()

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
	gcpMountPath := "gcpMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "gcp")

	gcpRoleServiceAccountsRequest := NewGcpRoleServiceAccountsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGcpRoleNameServiceAccounts(context.Background(), gcpMountPath, name, gcpRoleServiceAccountsRequest)
	if err != nil {
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

 **gcpRoleServiceAccountsRequest** | [**GcpRoleServiceAccountsRequest**](GcpRoleServiceAccountsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubConfig

> PostAuthGithubConfig(ctx, githubMountPath).GithubConfigRequest(githubConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubConfigRequest := NewGithubConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGithubConfig(context.Background(), githubMountPath, githubConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubConfigRequest** | [**GithubConfigRequest**](GithubConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubLogin

> PostAuthGithubLogin(ctx, githubMountPath).GithubLoginRequest(githubLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubLoginRequest := NewGithubLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGithubLogin(context.Background(), githubMountPath, githubLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubLoginRequest** | [**GithubLoginRequest**](GithubLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubMapTeamsKey

> PostAuthGithubMapTeamsKey(ctx, githubMountPath, key).GithubMapTeamsRequest(githubMapTeamsRequest).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubMapTeamsRequest := NewGithubMapTeamsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGithubMapTeamsKey(context.Background(), githubMountPath, key, githubMapTeamsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapTeamsRequest** | [**GithubMapTeamsRequest**](GithubMapTeamsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubMapUsersKey

> PostAuthGithubMapUsersKey(ctx, githubMountPath, key).GithubMapUsersRequest(githubMapUsersRequest).Execute()

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
	githubMountPath := "githubMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "github")

	githubMapUsersRequest := NewGithubMapUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthGithubMapUsersKey(context.Background(), githubMountPath, key, githubMapUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**githubMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;github&quot;]
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapUsersRequest** | [**GithubMapUsersRequest**](GithubMapUsersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtConfig

> PostAuthJwtConfig(ctx, jwtMountPath).JwtConfigRequest(jwtConfigRequest).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	jwtConfigRequest := NewJwtConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthJwtConfig(context.Background(), jwtMountPath, jwtConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtConfigRequest** | [**JwtConfigRequest**](JwtConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtLogin

> PostAuthJwtLogin(ctx, jwtMountPath).JwtLoginRequest(jwtLoginRequest).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	jwtLoginRequest := NewJwtLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthJwtLogin(context.Background(), jwtMountPath, jwtLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtLoginRequest** | [**JwtLoginRequest**](JwtLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtOidcAuthUrl

> PostAuthJwtOidcAuthUrl(ctx, jwtMountPath).JwtOidcAuthUrlRequest(jwtOidcAuthUrlRequest).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	jwtOidcAuthUrlRequest := NewJwtOidcAuthUrlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthJwtOidcAuthUrl(context.Background(), jwtMountPath, jwtOidcAuthUrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtOidcAuthUrlRequest** | [**JwtOidcAuthUrlRequest**](JwtOidcAuthUrlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtOidcCallback

> PostAuthJwtOidcCallback(ctx, jwtMountPath).JwtOidcCallbackRequest(jwtOidcCallbackRequest).Execute()

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

	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	jwtOidcCallbackRequest := NewJwtOidcCallbackRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthJwtOidcCallback(context.Background(), jwtMountPath, jwtOidcCallbackRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtOidcCallbackRequest** | [**JwtOidcCallbackRequest**](JwtOidcCallbackRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtRoleName

> PostAuthJwtRoleName(ctx, jwtMountPath, name).JwtRoleRequest(jwtRoleRequest).Execute()

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
	jwtMountPath := "jwtMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "jwt")

	jwtRoleRequest := NewJwtRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthJwtRoleName(context.Background(), jwtMountPath, name, jwtRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**jwtMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;jwt&quot;]
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtRoleRequest** | [**JwtRoleRequest**](JwtRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosConfig

> PostAuthKerberosConfig(ctx, kerberosMountPath).KerberosConfigRequest(kerberosConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosConfigRequest := NewKerberosConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKerberosConfig(context.Background(), kerberosMountPath, kerberosConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosConfigRequest** | [**KerberosConfigRequest**](KerberosConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosConfigLdap

> PostAuthKerberosConfigLdap(ctx, kerberosMountPath).KerberosConfigLdapRequest(kerberosConfigLdapRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosConfigLdapRequest := NewKerberosConfigLdapRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKerberosConfigLdap(context.Background(), kerberosMountPath, kerberosConfigLdapRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosConfigLdapRequest** | [**KerberosConfigLdapRequest**](KerberosConfigLdapRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosGroupsName

> PostAuthKerberosGroupsName(ctx, kerberosMountPath, name).KerberosGroupsRequest(kerberosGroupsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosGroupsRequest := NewKerberosGroupsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKerberosGroupsName(context.Background(), kerberosMountPath, name, kerberosGroupsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosGroupsRequest** | [**KerberosGroupsRequest**](KerberosGroupsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosLogin

> PostAuthKerberosLogin(ctx, kerberosMountPath).KerberosLoginRequest(kerberosLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	kerberosMountPath := "kerberosMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kerberos")

	kerberosLoginRequest := NewKerberosLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKerberosLogin(context.Background(), kerberosMountPath, kerberosLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**kerberosMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;kerberos&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosLoginRequest** | [**KerberosLoginRequest**](KerberosLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKubernetesConfig

> PostAuthKubernetesConfig(ctx, kubernetesMountPath).KubernetesConfigRequest(kubernetesConfigRequest).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.PostAuthKubernetesConfig(context.Background(), kubernetesMountPath, kubernetesConfigRequest)
	if err != nil {
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


## PostAuthKubernetesLogin

> PostAuthKubernetesLogin(ctx, kubernetesMountPath).KubernetesLoginRequest(kubernetesLoginRequest).Execute()

Authenticates Kubernetes service accounts with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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

	kubernetesLoginRequest := NewKubernetesLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKubernetesLogin(context.Background(), kubernetesMountPath, kubernetesLoginRequest)
	if err != nil {
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
 **kubernetesLoginRequest** | [**KubernetesLoginRequest**](KubernetesLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKubernetesRoleName

> PostAuthKubernetesRoleName(ctx, kubernetesMountPath, name).KubernetesRoleRequest(kubernetesRoleRequest).Execute()

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
	kubernetesMountPath := "kubernetesMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "kubernetes")

	kubernetesRoleRequest := NewKubernetesRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthKubernetesRoleName(context.Background(), kubernetesMountPath, name, kubernetesRoleRequest)
	if err != nil {
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesRoleRequest** | [**KubernetesRoleRequest**](KubernetesRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapConfig

> PostAuthLdapConfig(ctx, ldapMountPath).LdapConfigRequest(ldapConfigRequest).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.WithToken("my-token").Auth.PostAuthLdapConfig(context.Background(), ldapMountPath, ldapConfigRequest)
	if err != nil {
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


## PostAuthLdapGroupsName

> PostAuthLdapGroupsName(ctx, ldapMountPath, name).LdapGroupsRequest(ldapGroupsRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapGroupsRequest := NewLdapGroupsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthLdapGroupsName(context.Background(), ldapMountPath, name, ldapGroupsRequest)
	if err != nil {
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapGroupsRequest** | [**LdapGroupsRequest**](LdapGroupsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapLoginUsername

> PostAuthLdapLoginUsername(ctx, ldapMountPath, username).LdapLoginRequest(ldapLoginRequest).Execute()

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

	username := "username_example" // string | DN (distinguished name) to be used for login.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapLoginRequest := NewLdapLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthLdapLoginUsername(context.Background(), ldapMountPath, username, ldapLoginRequest)
	if err != nil {
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
**username** | **string** | DN (distinguished name) to be used for login. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLoginRequest** | [**LdapLoginRequest**](LdapLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapUsersName

> PostAuthLdapUsersName(ctx, ldapMountPath, name).LdapUsersRequest(ldapUsersRequest).Execute()

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

	name := "name_example" // string | Name of the LDAP user.
	ldapMountPath := "ldapMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "ldap")

	ldapUsersRequest := NewLdapUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthLdapUsersName(context.Background(), ldapMountPath, name, ldapUsersRequest)
	if err != nil {
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapUsersRequest** | [**LdapUsersRequest**](LdapUsersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciConfig

> PostAuthOciConfig(ctx, ociMountPath).OciConfigRequest(ociConfigRequest).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	ociConfigRequest := NewOciConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOciConfig(context.Background(), ociMountPath, ociConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ociConfigRequest** | [**OciConfigRequest**](OciConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciLoginRole

> PostAuthOciLoginRole(ctx, ociMountPath, role).OciLoginRequest(ociLoginRequest).Execute()

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
	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	ociLoginRequest := NewOciLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOciLoginRole(context.Background(), ociMountPath, role, ociLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ociLoginRequest** | [**OciLoginRequest**](OciLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciRoleRole

> PostAuthOciRoleRole(ctx, ociMountPath, role).OciRoleRequest(ociRoleRequest).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	ociMountPath := "ociMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oci")

	ociRoleRequest := NewOciRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOciRoleRole(context.Background(), ociMountPath, role, ociRoleRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**ociMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oci&quot;]
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ociRoleRequest** | [**OciRoleRequest**](OciRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcConfig

> PostAuthOidcConfig(ctx, oidcMountPath).OidcConfigRequest(oidcConfigRequest).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcConfigRequest := NewOidcConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOidcConfig(context.Background(), oidcMountPath, oidcConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcConfigRequest** | [**OidcConfigRequest**](OidcConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcLogin

> PostAuthOidcLogin(ctx, oidcMountPath).OidcLoginRequest(oidcLoginRequest).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcLoginRequest := NewOidcLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOidcLogin(context.Background(), oidcMountPath, oidcLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcLoginRequest** | [**OidcLoginRequest**](OidcLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcOidcAuthUrl

> PostAuthOidcOidcAuthUrl(ctx, oidcMountPath).OidcOidcAuthUrlRequest(oidcOidcAuthUrlRequest).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcOidcAuthUrlRequest := NewOidcOidcAuthUrlRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOidcOidcAuthUrl(context.Background(), oidcMountPath, oidcOidcAuthUrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcOidcAuthUrlRequest** | [**OidcOidcAuthUrlRequest**](OidcOidcAuthUrlRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcOidcCallback

> PostAuthOidcOidcCallback(ctx, oidcMountPath).OidcOidcCallbackRequest(oidcOidcCallbackRequest).Execute()

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

	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcOidcCallbackRequest := NewOidcOidcCallbackRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOidcOidcCallback(context.Background(), oidcMountPath, oidcOidcCallbackRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcOidcCallbackRequest** | [**OidcOidcCallbackRequest**](OidcOidcCallbackRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcRoleName

> PostAuthOidcRoleName(ctx, name, oidcMountPath).OidcRoleRequest(oidcRoleRequest).Execute()

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
	oidcMountPath := "oidcMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "oidc")

	oidcRoleRequest := NewOidcRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOidcRoleName(context.Background(), name, oidcMountPath, oidcRoleRequest)
	if err != nil {
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
**oidcMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;oidc&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcRoleRequest** | [**OidcRoleRequest**](OidcRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaConfig

> PostAuthOktaConfig(ctx, oktaMountPath).OktaConfigRequest(oktaConfigRequest).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	oktaConfigRequest := NewOktaConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOktaConfig(context.Background(), oktaMountPath, oktaConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oktaConfigRequest** | [**OktaConfigRequest**](OktaConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaGroupsName

> PostAuthOktaGroupsName(ctx, name, oktaMountPath).OktaGroupsRequest(oktaGroupsRequest).Execute()

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

	name := "name_example" // string | Name of the Okta group.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	oktaGroupsRequest := NewOktaGroupsRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOktaGroupsName(context.Background(), name, oktaMountPath, oktaGroupsRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the Okta group. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaGroupsRequest** | [**OktaGroupsRequest**](OktaGroupsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaLoginUsername

> PostAuthOktaLoginUsername(ctx, oktaMountPath, username).OktaLoginRequest(oktaLoginRequest).Execute()

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

	username := "username_example" // string | Username to be used for login.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	oktaLoginRequest := NewOktaLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOktaLoginUsername(context.Background(), oktaMountPath, username, oktaLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]
**username** | **string** | Username to be used for login. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaLoginRequest** | [**OktaLoginRequest**](OktaLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaUsersName

> PostAuthOktaUsersName(ctx, name, oktaMountPath).OktaUsersRequest(oktaUsersRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	oktaMountPath := "oktaMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "okta")

	oktaUsersRequest := NewOktaUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthOktaUsersName(context.Background(), name, oktaMountPath, oktaUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the user. | 
**oktaMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;okta&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaUsersRequest** | [**OktaUsersRequest**](OktaUsersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusConfig

> PostAuthRadiusConfig(ctx, radiusMountPath).RadiusConfigRequest(radiusConfigRequest).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusConfigRequest := NewRadiusConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthRadiusConfig(context.Background(), radiusMountPath, radiusConfigRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusConfigRequest** | [**RadiusConfigRequest**](RadiusConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusLogin

> PostAuthRadiusLogin(ctx, radiusMountPath).RadiusLoginRequest(radiusLoginRequest).Execute()

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

	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthRadiusLogin(context.Background(), radiusMountPath, radiusLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusLoginUrlusername

> PostAuthRadiusLoginUrlusername(ctx, radiusMountPath, urlusername).RadiusLoginRequest(radiusLoginRequest).Execute()

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
	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthRadiusLoginUrlusername(context.Background(), radiusMountPath, urlusername, radiusLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]
**urlusername** | **string** | Username to be used for login. (URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusUsersName

> PostAuthRadiusUsersName(ctx, name, radiusMountPath).RadiusUsersRequest(radiusUsersRequest).Execute()

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
	radiusMountPath := "radiusMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "radius")

	radiusUsersRequest := NewRadiusUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthRadiusUsersName(context.Background(), name, radiusMountPath, radiusUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the RADIUS user. | 
**radiusMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;radius&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusUsersRequest** | [**RadiusUsersRequest**](RadiusUsersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenCreate

> PostAuthTokenCreate(ctx, tokenMountPath).TokenCreateRequest(tokenCreateRequest).Format(format).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenCreate(context.Background(), tokenMountPath, tokenCreateRequest, format)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 
 **format** | **string** | Return json formatted output | [default to &quot;json&quot;]


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenCreateOrphan

> PostAuthTokenCreateOrphan(ctx, tokenMountPath).TokenCreateOrphanRequest(tokenCreateOrphanRequest).Format(format).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenCreateOrphanRequest := NewTokenCreateOrphanRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenCreateOrphan(context.Background(), tokenMountPath, tokenCreateOrphanRequest, format)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreateOrphanRequest** | [**TokenCreateOrphanRequest**](TokenCreateOrphanRequest.md) |  | 
 **format** | **string** | Return json formatted output | [default to &quot;json&quot;]


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenCreateRoleName

> PostAuthTokenCreateRoleName(ctx, roleName, tokenMountPath).TokenCreateRequest(tokenCreateRequest).Format(format).Execute()

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
	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenCreateRoleName(context.Background(), roleName, tokenMountPath, tokenCreateRequest, format)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role | 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 
 **format** | **string** | Return json formatted output | [default to &quot;json&quot;]


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookup

> PostAuthTokenLookup(ctx, tokenMountPath).TokenLookupRequest(tokenLookupRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupRequest := NewTokenLookupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenLookup(context.Background(), tokenMountPath, tokenLookupRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupRequest** | [**TokenLookupRequest**](TokenLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookupAccessor

> PostAuthTokenLookupAccessor(ctx, tokenMountPath).TokenLookupAccessorRequest(tokenLookupAccessorRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupAccessorRequest := NewTokenLookupAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenLookupAccessor(context.Background(), tokenMountPath, tokenLookupAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupAccessorRequest** | [**TokenLookupAccessorRequest**](TokenLookupAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookupSelf

> PostAuthTokenLookupSelf(ctx, tokenMountPath).TokenLookupSelfRequest(tokenLookupSelfRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenLookupSelfRequest := NewTokenLookupSelfRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenLookupSelf(context.Background(), tokenMountPath, tokenLookupSelfRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupSelfRequest** | [**TokenLookupSelfRequest**](TokenLookupSelfRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenew

> PostAuthTokenRenew(ctx, tokenMountPath).TokenRenewRequest(tokenRenewRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewRequest := NewTokenRenewRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRenew(context.Background(), tokenMountPath, tokenRenewRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewRequest** | [**TokenRenewRequest**](TokenRenewRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenewAccessor

> PostAuthTokenRenewAccessor(ctx, tokenMountPath).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewAccessorRequest := NewTokenRenewAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRenewAccessor(context.Background(), tokenMountPath, tokenRenewAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewAccessorRequest** | [**TokenRenewAccessorRequest**](TokenRenewAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenewSelf

> PostAuthTokenRenewSelf(ctx, tokenMountPath).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRenewSelfRequest := NewTokenRenewSelfRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRenewSelf(context.Background(), tokenMountPath, tokenRenewSelfRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewSelfRequest** | [**TokenRenewSelfRequest**](TokenRenewSelfRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevoke

> PostAuthTokenRevoke(ctx, tokenMountPath).TokenRevokeRequest(tokenRevokeRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeRequest := NewTokenRevokeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRevoke(context.Background(), tokenMountPath, tokenRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeAccessor

> PostAuthTokenRevokeAccessor(ctx, tokenMountPath).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeAccessorRequest := NewTokenRevokeAccessorRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRevokeAccessor(context.Background(), tokenMountPath, tokenRevokeAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeAccessorRequest** | [**TokenRevokeAccessorRequest**](TokenRevokeAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeOrphan

> PostAuthTokenRevokeOrphan(ctx, tokenMountPath).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRevokeOrphanRequest := NewTokenRevokeOrphanRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRevokeOrphan(context.Background(), tokenMountPath, tokenRevokeOrphanRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeOrphanRequest** | [**TokenRevokeOrphanRequest**](TokenRevokeOrphanRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeSelf

> PostAuthTokenRevokeSelf(ctx, tokenMountPath).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRevokeSelf(context.Background(), tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRolesRoleName

> PostAuthTokenRolesRoleName(ctx, roleName, tokenMountPath).TokenRolesRequest(tokenRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	tokenRolesRequest := NewTokenRolesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthTokenRolesRoleName(context.Background(), roleName, tokenMountPath, tokenRolesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role | 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRolesRequest** | [**TokenRolesRequest**](TokenRolesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenTidy

> PostAuthTokenTidy(ctx, tokenMountPath).Execute()

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

	tokenMountPath := "tokenMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "token")

	resp, err := client.WithToken("my-token").Auth.PostAuthTokenTidy(context.Background(), tokenMountPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tokenMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;token&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassLoginUsername

> PostAuthUserpassLoginUsername(ctx, username, userpassMountPath).UserpassLoginRequest(userpassLoginRequest).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassLoginRequest := NewUserpassLoginRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthUserpassLoginUsername(context.Background(), username, userpassMountPath, userpassLoginRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username of the user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassLoginRequest** | [**UserpassLoginRequest**](UserpassLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsername

> PostAuthUserpassUsersUsername(ctx, username, userpassMountPath).UserpassUsersRequest(userpassUsersRequest).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersRequest := NewUserpassUsersRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthUserpassUsersUsername(context.Background(), username, userpassMountPath, userpassUsersRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username for this user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersRequest** | [**UserpassUsersRequest**](UserpassUsersRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsernamePassword

> PostAuthUserpassUsersUsernamePassword(ctx, username, userpassMountPath).UserpassUsersPasswordRequest(userpassUsersPasswordRequest).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersPasswordRequest := NewUserpassUsersPasswordRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthUserpassUsersUsernamePassword(context.Background(), username, userpassMountPath, userpassUsersPasswordRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username for this user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPasswordRequest** | [**UserpassUsersPasswordRequest**](UserpassUsersPasswordRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsernamePolicies

> PostAuthUserpassUsersUsernamePolicies(ctx, username, userpassMountPath).UserpassUsersPoliciesRequest(userpassUsersPoliciesRequest).Execute()

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
	userpassMountPath := "userpassMountPath_example" // string | Path where the backend was mounted; the endpoint path will be offset by the mount path (defaults to "userpass")

	userpassUsersPoliciesRequest := NewUserpassUsersPoliciesRequestWithDefaults()
	resp, err := client.WithToken("my-token").Auth.PostAuthUserpassUsersUsernamePolicies(context.Background(), username, userpassMountPath, userpassUsersPoliciesRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**username** | **string** | Username for this user. | 
**userpassMountPath** | **string** | Path where the backend was mounted; the endpoint path will be offset by the mount path | [default to &quot;userpass&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPoliciesRequest** | [**UserpassUsersPoliciesRequest**](UserpassUsersPoliciesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

