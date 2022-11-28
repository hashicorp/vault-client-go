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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.

	resp, err := client.Auth.DeleteAuthAlicloudRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleName(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameBindSecretId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameBoundCidrList(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNamePeriod(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNamePolicies(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdDestroy(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenMaxTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.

	resp, err := client.Auth.DeleteAuthAwsConfigCertificateCertName(
		context.Background(),
		certName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAwsConfigClient(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

	resp, err := client.Auth.DeleteAuthAwsConfigStsAccountId(
		context.Background(),
		accountId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAwsConfigTidyIdentityAccesslist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAwsConfigTidyIdentityWhitelist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAwsConfigTidyRoletagBlacklist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAwsConfigTidyRoletagDenylist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

	resp, err := client.Auth.DeleteAuthAwsIdentityAccesslistInstanceId(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

	resp, err := client.Auth.DeleteAuthAwsIdentityWhitelistInstanceId(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthAwsRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.DeleteAuthAwsRoletagBlacklistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.DeleteAuthAwsRoletagDenylistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthAzureConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthAzureRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	resp, err := client.Auth.DeleteAuthCertCertsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	resp, err := client.Auth.DeleteAuthCertCrlsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthCfConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.

	resp, err := client.Auth.DeleteAuthCfRolesRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthGcpRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping

	resp, err := client.Auth.DeleteAuthGithubMapTeamsKey(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping

	resp, err := client.Auth.DeleteAuthGithubMapUsersKey(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthJwtRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	resp, err := client.Auth.DeleteAuthKerberosGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthKubernetesRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	resp, err := client.Auth.DeleteAuthLdapGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.

	resp, err := client.Auth.DeleteAuthLdapUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.DeleteAuthOciConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthOciRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.DeleteAuthOidcRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.

	resp, err := client.Auth.DeleteAuthOktaGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.

	resp, err := client.Auth.DeleteAuthOktaUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.

	resp, err := client.Auth.DeleteAuthRadiusUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role

	resp, err := client.Auth.DeleteAuthTokenRolesRoleName(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	resp, err := client.Auth.DeleteAuthUserpassUsersUsername(
		context.Background(),
		username,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAlicloudRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.

	resp, err := client.Auth.GetAuthAlicloudRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAlicloudRoles(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthApproleRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleName(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameBindSecretId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameBoundCidrList(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameLocalSecretIds(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNamePeriod(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNamePolicies(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameRoleId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretId(
		context.Background(),
		roleName,
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenMaxTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.

	resp, err := client.Auth.GetAuthAwsConfigCertificateCertName(
		context.Background(),
		certName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsConfigCertificates(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigClient(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigIdentity(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsConfigSts(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

	resp, err := client.Auth.GetAuthAwsConfigStsAccountId(
		context.Background(),
		accountId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigTidyIdentityAccesslist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigTidyIdentityWhitelist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigTidyRoletagBlacklist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAwsConfigTidyRoletagDenylist(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsIdentityAccesslist(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

	resp, err := client.Auth.GetAuthAwsIdentityAccesslistInstanceId(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsIdentityWhitelist(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

	resp, err := client.Auth.GetAuthAwsIdentityWhitelistInstanceId(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthAwsRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsRoles(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsRoletagBlacklist(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.GetAuthAwsRoletagBlacklistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAwsRoletagDenylist(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.GetAuthAwsRoletagDenylistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthAzureConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthAzureRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthAzureRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthCentrifyConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthCertCerts(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	resp, err := client.Auth.GetAuthCertCertsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthCertConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	resp, err := client.Auth.GetAuthCertCrlsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthCfConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthCfRoles(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.

	resp, err := client.Auth.GetAuthCfRolesRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthGcpConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthGcpRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthGcpRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthGcpRoles(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthGithubConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthGithubMapTeams(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping

	resp, err := client.Auth.GetAuthGithubMapTeamsKey(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthGithubMapUsers(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping

	resp, err := client.Auth.GetAuthGithubMapUsersKey(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthJwtConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthJwtOidcCallback(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthJwtRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthJwtRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthKerberosConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthKerberosConfigLdap(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthKerberosGroups(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	resp, err := client.Auth.GetAuthKerberosGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthKerberosLogin(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthKubernetesConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthKubernetesRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthKubernetesRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthLdapConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthLdapGroups(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	resp, err := client.Auth.GetAuthLdapGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthLdapUsers(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.

	resp, err := client.Auth.GetAuthLdapUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthOciConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthOciRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthOciRoleRole(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthOidcConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthOidcOidcCallback(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthOidcRole(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	resp, err := client.Auth.GetAuthOidcRoleName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthOktaConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthOktaGroups(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.

	resp, err := client.Auth.GetAuthOktaGroupsName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthOktaUsers(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.

	resp, err := client.Auth.GetAuthOktaUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nonce := "nonce_example" // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.

	resp, err := client.Auth.GetAuthOktaVerifyNonce(
		context.Background(),
		nonce,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthRadiusConfig(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthRadiusUsers(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.

	resp, err := client.Auth.GetAuthRadiusUsersName(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthTokenAccessors(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthTokenLookup(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.GetAuthTokenLookupSelf(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthTokenRoles(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role

	resp, err := client.Auth.GetAuthTokenRolesRoleName(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	list := NewstringWithDefaults()
	resp, err := client.Auth.GetAuthUserpassUsers(
		context.Background(),
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	resp, err := client.Auth.GetAuthUserpassUsersUsername(
		context.Background(),
		username,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	alicloudLoginRequest := NewAlicloudLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthAlicloudLogin(
		context.Background(),
		alicloudLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.

	alicloudRoleRequest := NewAlicloudRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthAlicloudRoleRole(
		context.Background(),
		role,
		alicloudRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	approleLoginRequest := NewApproleLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleLogin(
		context.Background(),
		approleLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleRequest := NewApproleRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleName(
		context.Background(),
		roleName,
		approleRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleBindSecretIdRequest := NewApproleRoleBindSecretIdRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameBindSecretId(
		context.Background(),
		roleName,
		approleRoleBindSecretIdRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleBoundCidrListRequest := NewApproleRoleBoundCidrListRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameBoundCidrList(
		context.Background(),
		roleName,
		approleRoleBoundCidrListRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleCustomSecretIdRequest := NewApproleRoleCustomSecretIdRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameCustomSecretId(
		context.Background(),
		roleName,
		approleRoleCustomSecretIdRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRolePeriodRequest := NewApproleRolePeriodRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNamePeriod(
		context.Background(),
		roleName,
		approleRolePeriodRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRolePoliciesRequest := NewApproleRolePoliciesRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNamePolicies(
		context.Background(),
		roleName,
		approleRolePoliciesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleRoleIdRequest := NewApproleRoleRoleIdRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameRoleId(
		context.Background(),
		roleName,
		approleRoleRoleIdRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdRequest := NewApproleRoleSecretIdRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretId(
		context.Background(),
		roleName,
		approleRoleSecretIdRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdAccessorDestroyRequest := NewApproleRoleSecretIdAccessorDestroyRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(
		context.Background(),
		roleName,
		approleRoleSecretIdAccessorDestroyRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdAccessorLookupRequest := NewApproleRoleSecretIdAccessorLookupRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorLookup(
		context.Background(),
		roleName,
		approleRoleSecretIdAccessorLookupRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdBoundCidrsRequest := NewApproleRoleSecretIdBoundCidrsRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdBoundCidrs(
		context.Background(),
		roleName,
		approleRoleSecretIdBoundCidrsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdDestroyRequest := NewApproleRoleSecretIdDestroyRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdDestroy(
		context.Background(),
		roleName,
		approleRoleSecretIdDestroyRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdLookupRequest := NewApproleRoleSecretIdLookupRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdLookup(
		context.Background(),
		roleName,
		approleRoleSecretIdLookupRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdNumUsesRequest := NewApproleRoleSecretIdNumUsesRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdNumUses(
		context.Background(),
		roleName,
		approleRoleSecretIdNumUsesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleSecretIdTtlRequest := NewApproleRoleSecretIdTtlRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdTtl(
		context.Background(),
		roleName,
		approleRoleSecretIdTtlRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleTokenBoundCidrsRequest := NewApproleRoleTokenBoundCidrsRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenBoundCidrs(
		context.Background(),
		roleName,
		approleRoleTokenBoundCidrsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleTokenMaxTtlRequest := NewApproleRoleTokenMaxTtlRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenMaxTtl(
		context.Background(),
		roleName,
		approleRoleTokenMaxTtlRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleTokenNumUsesRequest := NewApproleRoleTokenNumUsesRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenNumUses(
		context.Background(),
		roleName,
		approleRoleTokenNumUsesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	approleRoleTokenTtlRequest := NewApproleRoleTokenTtlRequestWithDefaults()
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenTtl(
		context.Background(),
		roleName,
		approleRoleTokenTtlRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.PostAuthApproleTidySecretId(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.

	awsConfigCertificateRequest := NewAwsConfigCertificateRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigCertificateCertName(
		context.Background(),
		certName,
		awsConfigCertificateRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigClientRequest := NewAwsConfigClientRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigClient(
		context.Background(),
		awsConfigClientRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigIdentityRequest := NewAwsConfigIdentityRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigIdentity(
		context.Background(),
		awsConfigIdentityRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.PostAuthAwsConfigRotateRoot(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

	awsConfigStsRequest := NewAwsConfigStsRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigStsAccountId(
		context.Background(),
		accountId,
		awsConfigStsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigTidyIdentityAccesslistRequest := NewAwsConfigTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigTidyIdentityAccesslist(
		context.Background(),
		awsConfigTidyIdentityAccesslistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigTidyIdentityWhitelistRequest := NewAwsConfigTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigTidyIdentityWhitelist(
		context.Background(),
		awsConfigTidyIdentityWhitelistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigTidyRoletagBlacklistRequest := NewAwsConfigTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigTidyRoletagBlacklist(
		context.Background(),
		awsConfigTidyRoletagBlacklistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsConfigTidyRoletagDenylistRequest := NewAwsConfigTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsConfigTidyRoletagDenylist(
		context.Background(),
		awsConfigTidyRoletagDenylistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsLoginRequest := NewAwsLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsLogin(
		context.Background(),
		awsLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	awsRoleRequest := NewAwsRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsRoleRole(
		context.Background(),
		role,
		awsRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	awsRoleTagRequest := NewAwsRoleTagRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsRoleRoleTag(
		context.Background(),
		role,
		awsRoleTagRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.PostAuthAwsRoletagBlacklistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

	resp, err := client.Auth.PostAuthAwsRoletagDenylistRoleTag(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsTidyIdentityAccesslistRequest := NewAwsTidyIdentityAccesslistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsTidyIdentityAccesslist(
		context.Background(),
		awsTidyIdentityAccesslistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsTidyIdentityWhitelistRequest := NewAwsTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsTidyIdentityWhitelist(
		context.Background(),
		awsTidyIdentityWhitelistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsTidyRoletagBlacklistRequest := NewAwsTidyRoletagBlacklistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsTidyRoletagBlacklist(
		context.Background(),
		awsTidyRoletagBlacklistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	awsTidyRoletagDenylistRequest := NewAwsTidyRoletagDenylistRequestWithDefaults()
	resp, err := client.Auth.PostAuthAwsTidyRoletagDenylist(
		context.Background(),
		awsTidyRoletagDenylistRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	azureConfigRequest := NewAzureConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthAzureConfig(
		context.Background(),
		azureConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	azureLoginRequest := NewAzureLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthAzureLogin(
		context.Background(),
		azureLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	azureRoleRequest := NewAzureRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthAzureRoleName(
		context.Background(),
		name,
		azureRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	centrifyConfigRequest := NewCentrifyConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthCentrifyConfig(
		context.Background(),
		centrifyConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	centrifyLoginRequest := NewCentrifyLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthCentrifyLogin(
		context.Background(),
		centrifyLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	certCertsRequest := NewCertCertsRequestWithDefaults()
	resp, err := client.Auth.PostAuthCertCertsName(
		context.Background(),
		name,
		certCertsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	certConfigRequest := NewCertConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthCertConfig(
		context.Background(),
		certConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	certCrlsRequest := NewCertCrlsRequestWithDefaults()
	resp, err := client.Auth.PostAuthCertCrlsName(
		context.Background(),
		name,
		certCrlsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	certLoginRequest := NewCertLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthCertLogin(
		context.Background(),
		certLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	cfConfigRequest := NewCfConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthCfConfig(
		context.Background(),
		cfConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	cfLoginRequest := NewCfLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthCfLogin(
		context.Background(),
		cfLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.

	cfRolesRequest := NewCfRolesRequestWithDefaults()
	resp, err := client.Auth.PostAuthCfRolesRole(
		context.Background(),
		role,
		cfRolesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	gcpConfigRequest := NewGcpConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthGcpConfig(
		context.Background(),
		gcpConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	gcpLoginRequest := NewGcpLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthGcpLogin(
		context.Background(),
		gcpLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	gcpRoleRequest := NewGcpRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthGcpRoleName(
		context.Background(),
		name,
		gcpRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	gcpRoleLabelsRequest := NewGcpRoleLabelsRequestWithDefaults()
	resp, err := client.Auth.PostAuthGcpRoleNameLabels(
		context.Background(),
		name,
		gcpRoleLabelsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	gcpRoleServiceAccountsRequest := NewGcpRoleServiceAccountsRequestWithDefaults()
	resp, err := client.Auth.PostAuthGcpRoleNameServiceAccounts(
		context.Background(),
		name,
		gcpRoleServiceAccountsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	githubConfigRequest := NewGithubConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthGithubConfig(
		context.Background(),
		githubConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	githubLoginRequest := NewGithubLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthGithubLogin(
		context.Background(),
		githubLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping

	githubMapTeamsRequest := NewGithubMapTeamsRequestWithDefaults()
	resp, err := client.Auth.PostAuthGithubMapTeamsKey(
		context.Background(),
		key,
		githubMapTeamsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping

	githubMapUsersRequest := NewGithubMapUsersRequestWithDefaults()
	resp, err := client.Auth.PostAuthGithubMapUsersKey(
		context.Background(),
		key,
		githubMapUsersRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	jwtConfigRequest := NewJwtConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthJwtConfig(
		context.Background(),
		jwtConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	jwtLoginRequest := NewJwtLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthJwtLogin(
		context.Background(),
		jwtLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	jwtOidcAuthUrlRequest := NewJwtOidcAuthUrlRequestWithDefaults()
	resp, err := client.Auth.PostAuthJwtOidcAuthUrl(
		context.Background(),
		jwtOidcAuthUrlRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	jwtOidcCallbackRequest := NewJwtOidcCallbackRequestWithDefaults()
	resp, err := client.Auth.PostAuthJwtOidcCallback(
		context.Background(),
		jwtOidcCallbackRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	jwtRoleRequest := NewJwtRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthJwtRoleName(
		context.Background(),
		name,
		jwtRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	kerberosConfigRequest := NewKerberosConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthKerberosConfig(
		context.Background(),
		kerberosConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	kerberosConfigLdapRequest := NewKerberosConfigLdapRequestWithDefaults()
	resp, err := client.Auth.PostAuthKerberosConfigLdap(
		context.Background(),
		kerberosConfigLdapRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	kerberosGroupsRequest := NewKerberosGroupsRequestWithDefaults()
	resp, err := client.Auth.PostAuthKerberosGroupsName(
		context.Background(),
		name,
		kerberosGroupsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	kerberosLoginRequest := NewKerberosLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthKerberosLogin(
		context.Background(),
		kerberosLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	kubernetesConfigRequest := NewKubernetesConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthKubernetesConfig(
		context.Background(),
		kubernetesConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	kubernetesLoginRequest := NewKubernetesLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthKubernetesLogin(
		context.Background(),
		kubernetesLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	kubernetesRoleRequest := NewKubernetesRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthKubernetesRoleName(
		context.Background(),
		name,
		kubernetesRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	ldapConfigRequest := NewLdapConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthLdapConfig(
		context.Background(),
		ldapConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	ldapGroupsRequest := NewLdapGroupsRequestWithDefaults()
	resp, err := client.Auth.PostAuthLdapGroupsName(
		context.Background(),
		name,
		ldapGroupsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | DN (distinguished name) to be used for login.

	ldapLoginRequest := NewLdapLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthLdapLoginUsername(
		context.Background(),
		username,
		ldapLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.

	ldapUsersRequest := NewLdapUsersRequestWithDefaults()
	resp, err := client.Auth.PostAuthLdapUsersName(
		context.Background(),
		name,
		ldapUsersRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	ociConfigRequest := NewOciConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthOciConfig(
		context.Background(),
		ociConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	ociLoginRequest := NewOciLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthOciLoginRole(
		context.Background(),
		role,
		ociLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	ociRoleRequest := NewOciRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthOciRoleRole(
		context.Background(),
		role,
		ociRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	oidcConfigRequest := NewOidcConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthOidcConfig(
		context.Background(),
		oidcConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	oidcLoginRequest := NewOidcLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthOidcLogin(
		context.Background(),
		oidcLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	oidcOidcAuthUrlRequest := NewOidcOidcAuthUrlRequestWithDefaults()
	resp, err := client.Auth.PostAuthOidcOidcAuthUrl(
		context.Background(),
		oidcOidcAuthUrlRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	oidcOidcCallbackRequest := NewOidcOidcCallbackRequestWithDefaults()
	resp, err := client.Auth.PostAuthOidcOidcCallback(
		context.Background(),
		oidcOidcCallbackRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	oidcRoleRequest := NewOidcRoleRequestWithDefaults()
	resp, err := client.Auth.PostAuthOidcRoleName(
		context.Background(),
		name,
		oidcRoleRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	oktaConfigRequest := NewOktaConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthOktaConfig(
		context.Background(),
		oktaConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.

	oktaGroupsRequest := NewOktaGroupsRequestWithDefaults()
	resp, err := client.Auth.PostAuthOktaGroupsName(
		context.Background(),
		name,
		oktaGroupsRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username to be used for login.

	oktaLoginRequest := NewOktaLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthOktaLoginUsername(
		context.Background(),
		username,
		oktaLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.

	oktaUsersRequest := NewOktaUsersRequestWithDefaults()
	resp, err := client.Auth.PostAuthOktaUsersName(
		context.Background(),
		name,
		oktaUsersRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	radiusConfigRequest := NewRadiusConfigRequestWithDefaults()
	resp, err := client.Auth.PostAuthRadiusConfig(
		context.Background(),
		radiusConfigRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthRadiusLogin(
		context.Background(),
		radiusLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlusername := "urlusername_example" // string | Username to be used for login. (URL parameter)

	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthRadiusLoginUrlusername(
		context.Background(),
		urlusername,
		radiusLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.

	radiusUsersRequest := NewRadiusUsersRequestWithDefaults()
	resp, err := client.Auth.PostAuthRadiusUsersName(
		context.Background(),
		name,
		radiusUsersRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.Auth.PostAuthTokenCreate(
		context.Background(),
		tokenCreateRequest,
		format,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenCreateOrphanRequest := NewTokenCreateOrphanRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.Auth.PostAuthTokenCreateOrphan(
		context.Background(),
		tokenCreateOrphanRequest,
		format,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role

	tokenCreateRequest := NewTokenCreateRequestWithDefaults()
	format := NewstringWithDefaults()
	resp, err := client.Auth.PostAuthTokenCreateRoleName(
		context.Background(),
		roleName,
		tokenCreateRequest,
		format,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenLookupRequest := NewTokenLookupRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenLookup(
		context.Background(),
		tokenLookupRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenLookupAccessorRequest := NewTokenLookupAccessorRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenLookupAccessor(
		context.Background(),
		tokenLookupAccessorRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenLookupSelfRequest := NewTokenLookupSelfRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenLookupSelf(
		context.Background(),
		tokenLookupSelfRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRenewRequest := NewTokenRenewRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRenew(
		context.Background(),
		tokenRenewRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRenewAccessorRequest := NewTokenRenewAccessorRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRenewAccessor(
		context.Background(),
		tokenRenewAccessorRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRenewSelfRequest := NewTokenRenewSelfRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRenewSelf(
		context.Background(),
		tokenRenewSelfRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRevokeRequest := NewTokenRevokeRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRevoke(
		context.Background(),
		tokenRevokeRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRevokeAccessorRequest := NewTokenRevokeAccessorRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRevokeAccessor(
		context.Background(),
		tokenRevokeAccessorRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	tokenRevokeOrphanRequest := NewTokenRevokeOrphanRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRevokeOrphan(
		context.Background(),
		tokenRevokeOrphanRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.PostAuthTokenRevokeSelf(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role

	tokenRolesRequest := NewTokenRolesRequestWithDefaults()
	resp, err := client.Auth.PostAuthTokenRolesRoleName(
		context.Background(),
		roleName,
		tokenRolesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	resp, err := client.Auth.PostAuthTokenTidy(
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username of the user.

	userpassLoginRequest := NewUserpassLoginRequestWithDefaults()
	resp, err := client.Auth.PostAuthUserpassLoginUsername(
		context.Background(),
		username,
		userpassLoginRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	userpassUsersRequest := NewUserpassUsersRequestWithDefaults()
	resp, err := client.Auth.PostAuthUserpassUsersUsername(
		context.Background(),
		username,
		userpassUsersRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	userpassUsersPasswordRequest := NewUserpassUsersPasswordRequestWithDefaults()
	resp, err := client.Auth.PostAuthUserpassUsersUsernamePassword(
		context.Background(),
		username,
		userpassUsersPasswordRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	userpassUsersPoliciesRequest := NewUserpassUsersPoliciesRequestWithDefaults()
	resp, err := client.Auth.PostAuthUserpassUsersUsernamePolicies(
		context.Background(),
		username,
		userpassUsersPoliciesRequest,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

