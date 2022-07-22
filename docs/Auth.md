# \Auth

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthAlicloudRoleRole**](Auth.md#DeleteAuthAlicloudRoleRole) | **Delete** /auth/alicloud/role/{role} | Create a role and associate policies to it.
[**DeleteAuthAppIdMapAppIdKey**](Auth.md#DeleteAuthAppIdMapAppIdKey) | **Delete** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
[**DeleteAuthAppIdMapUserIdKey**](Auth.md#DeleteAuthAppIdMapUserIdKey) | **Delete** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
[**DeleteAuthApproleRoleRoleName**](Auth.md#DeleteAuthApproleRoleRoleName) | **Delete** /auth/approle/role/{role_name} | Register an role with the backend.
[**DeleteAuthApproleRoleRoleNameBindSecretId**](Auth.md#DeleteAuthApproleRoleRoleNameBindSecretId) | **Delete** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**DeleteAuthApproleRoleRoleNameBoundCidrList**](Auth.md#DeleteAuthApproleRoleRoleNameBoundCidrList) | **Delete** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteAuthApproleRoleRoleNamePeriod**](Auth.md#DeleteAuthApproleRoleRoleNamePeriod) | **Delete** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**DeleteAuthApproleRoleRoleNamePolicies**](Auth.md#DeleteAuthApproleRoleRoleNamePolicies) | **Delete** /auth/approle/role/{role_name}/policies | Policies of the role.
[**DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy) | **Delete** /auth/approle/role/{role_name}/secret-id-accessor/destroy | 
[**DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Delete** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**DeleteAuthApproleRoleRoleNameSecretIdDestroy**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdDestroy) | **Delete** /auth/approle/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**DeleteAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdNumUses) | **Delete** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**DeleteAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#DeleteAuthApproleRoleRoleNameSecretIdTtl) | **Delete** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
[**DeleteAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#DeleteAuthApproleRoleRoleNameTokenBoundCidrs) | **Delete** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**DeleteAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#DeleteAuthApproleRoleRoleNameTokenMaxTtl) | **Delete** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**DeleteAuthApproleRoleRoleNameTokenNumUses**](Auth.md#DeleteAuthApproleRoleRoleNameTokenNumUses) | **Delete** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**DeleteAuthApproleRoleRoleNameTokenTtl**](Auth.md#DeleteAuthApproleRoleRoleNameTokenTtl) | **Delete** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**DeleteAuthAwsConfigCertificateCertName**](Auth.md#DeleteAuthAwsConfigCertificateCertName) | **Delete** /auth/aws/config/certificate/{cert_name} | 
[**DeleteAuthAwsConfigClient**](Auth.md#DeleteAuthAwsConfigClient) | **Delete** /auth/aws/config/client | 
[**DeleteAuthAwsConfigStsAccountId**](Auth.md#DeleteAuthAwsConfigStsAccountId) | **Delete** /auth/aws/config/sts/{account_id} | 
[**DeleteAuthAwsConfigTidyIdentityAccesslist**](Auth.md#DeleteAuthAwsConfigTidyIdentityAccesslist) | **Delete** /auth/aws/config/tidy/identity-accesslist | 
[**DeleteAuthAwsConfigTidyIdentityWhitelist**](Auth.md#DeleteAuthAwsConfigTidyIdentityWhitelist) | **Delete** /auth/aws/config/tidy/identity-whitelist | 
[**DeleteAuthAwsConfigTidyRoletagBlacklist**](Auth.md#DeleteAuthAwsConfigTidyRoletagBlacklist) | **Delete** /auth/aws/config/tidy/roletag-blacklist | 
[**DeleteAuthAwsConfigTidyRoletagDenylist**](Auth.md#DeleteAuthAwsConfigTidyRoletagDenylist) | **Delete** /auth/aws/config/tidy/roletag-denylist | 
[**DeleteAuthAwsIdentityAccesslistInstanceId**](Auth.md#DeleteAuthAwsIdentityAccesslistInstanceId) | **Delete** /auth/aws/identity-accesslist/{instance_id} | 
[**DeleteAuthAwsIdentityWhitelistInstanceId**](Auth.md#DeleteAuthAwsIdentityWhitelistInstanceId) | **Delete** /auth/aws/identity-whitelist/{instance_id} | 
[**DeleteAuthAwsRoleRole**](Auth.md#DeleteAuthAwsRoleRole) | **Delete** /auth/aws/role/{role} | 
[**DeleteAuthAwsRoletagBlacklistRoleTag**](Auth.md#DeleteAuthAwsRoletagBlacklistRoleTag) | **Delete** /auth/aws/roletag-blacklist/{role_tag} | 
[**DeleteAuthAwsRoletagDenylistRoleTag**](Auth.md#DeleteAuthAwsRoletagDenylistRoleTag) | **Delete** /auth/aws/roletag-denylist/{role_tag} | 
[**DeleteAuthAzureConfig**](Auth.md#DeleteAuthAzureConfig) | **Delete** /auth/azure/config | 
[**DeleteAuthAzureRoleName**](Auth.md#DeleteAuthAzureRoleName) | **Delete** /auth/azure/role/{name} | 
[**DeleteAuthCertCertsName**](Auth.md#DeleteAuthCertCertsName) | **Delete** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
[**DeleteAuthCertCrlsName**](Auth.md#DeleteAuthCertCrlsName) | **Delete** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**DeleteAuthCfConfig**](Auth.md#DeleteAuthCfConfig) | **Delete** /auth/cf/config | 
[**DeleteAuthCfRolesRole**](Auth.md#DeleteAuthCfRolesRole) | **Delete** /auth/cf/roles/{role} | 
[**DeleteAuthGcpRoleName**](Auth.md#DeleteAuthGcpRoleName) | **Delete** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
[**DeleteAuthGithubMapTeamsKey**](Auth.md#DeleteAuthGithubMapTeamsKey) | **Delete** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
[**DeleteAuthGithubMapUsersKey**](Auth.md#DeleteAuthGithubMapUsersKey) | **Delete** /auth/github/map/users/{key} | Read/write/delete a single users mapping
[**DeleteAuthJwtRoleName**](Auth.md#DeleteAuthJwtRoleName) | **Delete** /auth/jwt/role/{name} | Delete an existing role.
[**DeleteAuthKerberosGroupsName**](Auth.md#DeleteAuthKerberosGroupsName) | **Delete** /auth/kerberos/groups/{name} | 
[**DeleteAuthKubernetesRoleName**](Auth.md#DeleteAuthKubernetesRoleName) | **Delete** /auth/kubernetes/role/{name} | Register an role with the backend.
[**DeleteAuthLdapGroupsName**](Auth.md#DeleteAuthLdapGroupsName) | **Delete** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
[**DeleteAuthLdapUsersName**](Auth.md#DeleteAuthLdapUsersName) | **Delete** /auth/ldap/users/{name} | Manage users allowed to authenticate.
[**DeleteAuthOciConfig**](Auth.md#DeleteAuthOciConfig) | **Delete** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
[**DeleteAuthOciRoleRole**](Auth.md#DeleteAuthOciRoleRole) | **Delete** /auth/oci/role/{role} | Create a role and associate policies to it.
[**DeleteAuthOidcRoleName**](Auth.md#DeleteAuthOidcRoleName) | **Delete** /auth/oidc/role/{name} | Delete an existing role.
[**DeleteAuthOktaGroupsName**](Auth.md#DeleteAuthOktaGroupsName) | **Delete** /auth/okta/groups/{name} | Manage users allowed to authenticate.
[**DeleteAuthOktaUsersName**](Auth.md#DeleteAuthOktaUsersName) | **Delete** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
[**DeleteAuthRadiusUsersName**](Auth.md#DeleteAuthRadiusUsersName) | **Delete** /auth/radius/users/{name} | Manage users allowed to authenticate.
[**DeleteAuthTokenRolesRoleName**](Auth.md#DeleteAuthTokenRolesRoleName) | **Delete** /auth/token/roles/{role_name} | 
[**DeleteAuthUserpassUsersUsername**](Auth.md#DeleteAuthUserpassUsersUsername) | **Delete** /auth/userpass/users/{username} | Manage users allowed to authenticate.
[**GetAuthAlicloudRole**](Auth.md#GetAuthAlicloudRole) | **Get** /auth/alicloud/role | Lists all the roles that are registered with Vault.
[**GetAuthAlicloudRoleRole**](Auth.md#GetAuthAlicloudRoleRole) | **Get** /auth/alicloud/role/{role} | Create a role and associate policies to it.
[**GetAuthAlicloudRoles**](Auth.md#GetAuthAlicloudRoles) | **Get** /auth/alicloud/roles | Lists all the roles that are registered with Vault.
[**GetAuthAppIdMapAppId**](Auth.md#GetAuthAppIdMapAppId) | **Get** /auth/app-id/map/app-id | Read mappings for app-id
[**GetAuthAppIdMapAppIdKey**](Auth.md#GetAuthAppIdMapAppIdKey) | **Get** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
[**GetAuthAppIdMapUserId**](Auth.md#GetAuthAppIdMapUserId) | **Get** /auth/app-id/map/user-id | Read mappings for user-id
[**GetAuthAppIdMapUserIdKey**](Auth.md#GetAuthAppIdMapUserIdKey) | **Get** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
[**GetAuthApproleRole**](Auth.md#GetAuthApproleRole) | **Get** /auth/approle/role | Lists all the roles registered with the backend.
[**GetAuthApproleRoleRoleName**](Auth.md#GetAuthApproleRoleRoleName) | **Get** /auth/approle/role/{role_name} | Register an role with the backend.
[**GetAuthApproleRoleRoleNameBindSecretId**](Auth.md#GetAuthApproleRoleRoleNameBindSecretId) | **Get** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**GetAuthApproleRoleRoleNameBoundCidrList**](Auth.md#GetAuthApproleRoleRoleNameBoundCidrList) | **Get** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**GetAuthApproleRoleRoleNameLocalSecretIds**](Auth.md#GetAuthApproleRoleRoleNameLocalSecretIds) | **Get** /auth/approle/role/{role_name}/local-secret-ids | Enables cluster local secret IDs
[**GetAuthApproleRoleRoleNamePeriod**](Auth.md#GetAuthApproleRoleRoleNamePeriod) | **Get** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**GetAuthApproleRoleRoleNamePolicies**](Auth.md#GetAuthApproleRoleRoleNamePolicies) | **Get** /auth/approle/role/{role_name}/policies | Policies of the role.
[**GetAuthApproleRoleRoleNameRoleId**](Auth.md#GetAuthApproleRoleRoleNameRoleId) | **Get** /auth/approle/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**GetAuthApproleRoleRoleNameSecretId**](Auth.md#GetAuthApproleRoleRoleNameSecretId) | **Get** /auth/approle/role/{role_name}/secret-id | Generate a SecretID against this role.
[**GetAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#GetAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Get** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**GetAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#GetAuthApproleRoleRoleNameSecretIdNumUses) | **Get** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**GetAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#GetAuthApproleRoleRoleNameSecretIdTtl) | **Get** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
[**GetAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#GetAuthApproleRoleRoleNameTokenBoundCidrs) | **Get** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**GetAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#GetAuthApproleRoleRoleNameTokenMaxTtl) | **Get** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**GetAuthApproleRoleRoleNameTokenNumUses**](Auth.md#GetAuthApproleRoleRoleNameTokenNumUses) | **Get** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**GetAuthApproleRoleRoleNameTokenTtl**](Auth.md#GetAuthApproleRoleRoleNameTokenTtl) | **Get** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**GetAuthAwsConfigCertificateCertName**](Auth.md#GetAuthAwsConfigCertificateCertName) | **Get** /auth/aws/config/certificate/{cert_name} | 
[**GetAuthAwsConfigCertificates**](Auth.md#GetAuthAwsConfigCertificates) | **Get** /auth/aws/config/certificates | 
[**GetAuthAwsConfigClient**](Auth.md#GetAuthAwsConfigClient) | **Get** /auth/aws/config/client | 
[**GetAuthAwsConfigIdentity**](Auth.md#GetAuthAwsConfigIdentity) | **Get** /auth/aws/config/identity | 
[**GetAuthAwsConfigSts**](Auth.md#GetAuthAwsConfigSts) | **Get** /auth/aws/config/sts | 
[**GetAuthAwsConfigStsAccountId**](Auth.md#GetAuthAwsConfigStsAccountId) | **Get** /auth/aws/config/sts/{account_id} | 
[**GetAuthAwsConfigTidyIdentityAccesslist**](Auth.md#GetAuthAwsConfigTidyIdentityAccesslist) | **Get** /auth/aws/config/tidy/identity-accesslist | 
[**GetAuthAwsConfigTidyIdentityWhitelist**](Auth.md#GetAuthAwsConfigTidyIdentityWhitelist) | **Get** /auth/aws/config/tidy/identity-whitelist | 
[**GetAuthAwsConfigTidyRoletagBlacklist**](Auth.md#GetAuthAwsConfigTidyRoletagBlacklist) | **Get** /auth/aws/config/tidy/roletag-blacklist | 
[**GetAuthAwsConfigTidyRoletagDenylist**](Auth.md#GetAuthAwsConfigTidyRoletagDenylist) | **Get** /auth/aws/config/tidy/roletag-denylist | 
[**GetAuthAwsIdentityAccesslist**](Auth.md#GetAuthAwsIdentityAccesslist) | **Get** /auth/aws/identity-accesslist | 
[**GetAuthAwsIdentityAccesslistInstanceId**](Auth.md#GetAuthAwsIdentityAccesslistInstanceId) | **Get** /auth/aws/identity-accesslist/{instance_id} | 
[**GetAuthAwsIdentityWhitelist**](Auth.md#GetAuthAwsIdentityWhitelist) | **Get** /auth/aws/identity-whitelist | 
[**GetAuthAwsIdentityWhitelistInstanceId**](Auth.md#GetAuthAwsIdentityWhitelistInstanceId) | **Get** /auth/aws/identity-whitelist/{instance_id} | 
[**GetAuthAwsRole**](Auth.md#GetAuthAwsRole) | **Get** /auth/aws/role | 
[**GetAuthAwsRoleRole**](Auth.md#GetAuthAwsRoleRole) | **Get** /auth/aws/role/{role} | 
[**GetAuthAwsRoles**](Auth.md#GetAuthAwsRoles) | **Get** /auth/aws/roles | 
[**GetAuthAwsRoletagBlacklist**](Auth.md#GetAuthAwsRoletagBlacklist) | **Get** /auth/aws/roletag-blacklist | 
[**GetAuthAwsRoletagBlacklistRoleTag**](Auth.md#GetAuthAwsRoletagBlacklistRoleTag) | **Get** /auth/aws/roletag-blacklist/{role_tag} | 
[**GetAuthAwsRoletagDenylist**](Auth.md#GetAuthAwsRoletagDenylist) | **Get** /auth/aws/roletag-denylist | 
[**GetAuthAwsRoletagDenylistRoleTag**](Auth.md#GetAuthAwsRoletagDenylistRoleTag) | **Get** /auth/aws/roletag-denylist/{role_tag} | 
[**GetAuthAzureConfig**](Auth.md#GetAuthAzureConfig) | **Get** /auth/azure/config | 
[**GetAuthAzureRole**](Auth.md#GetAuthAzureRole) | **Get** /auth/azure/role | 
[**GetAuthAzureRoleName**](Auth.md#GetAuthAzureRoleName) | **Get** /auth/azure/role/{name} | 
[**GetAuthCentrifyConfig**](Auth.md#GetAuthCentrifyConfig) | **Get** /auth/centrify/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**GetAuthCertCerts**](Auth.md#GetAuthCertCerts) | **Get** /auth/cert/certs | Manage trusted certificates used for authentication.
[**GetAuthCertCertsName**](Auth.md#GetAuthCertCertsName) | **Get** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
[**GetAuthCertCrlsName**](Auth.md#GetAuthCertCrlsName) | **Get** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**GetAuthCfConfig**](Auth.md#GetAuthCfConfig) | **Get** /auth/cf/config | 
[**GetAuthCfRoles**](Auth.md#GetAuthCfRoles) | **Get** /auth/cf/roles | 
[**GetAuthCfRolesRole**](Auth.md#GetAuthCfRolesRole) | **Get** /auth/cf/roles/{role} | 
[**GetAuthGcpConfig**](Auth.md#GetAuthGcpConfig) | **Get** /auth/gcp/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GetAuthGcpRole**](Auth.md#GetAuthGcpRole) | **Get** /auth/gcp/role | Lists all the roles that are registered with Vault.
[**GetAuthGcpRoleName**](Auth.md#GetAuthGcpRoleName) | **Get** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
[**GetAuthGcpRoles**](Auth.md#GetAuthGcpRoles) | **Get** /auth/gcp/roles | Lists all the roles that are registered with Vault.
[**GetAuthGithubConfig**](Auth.md#GetAuthGithubConfig) | **Get** /auth/github/config | 
[**GetAuthGithubMapTeams**](Auth.md#GetAuthGithubMapTeams) | **Get** /auth/github/map/teams | Read mappings for teams
[**GetAuthGithubMapTeamsKey**](Auth.md#GetAuthGithubMapTeamsKey) | **Get** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
[**GetAuthGithubMapUsers**](Auth.md#GetAuthGithubMapUsers) | **Get** /auth/github/map/users | Read mappings for users
[**GetAuthGithubMapUsersKey**](Auth.md#GetAuthGithubMapUsersKey) | **Get** /auth/github/map/users/{key} | Read/write/delete a single users mapping
[**GetAuthJwtConfig**](Auth.md#GetAuthJwtConfig) | **Get** /auth/jwt/config | Read the current JWT authentication backend configuration.
[**GetAuthJwtOidcCallback**](Auth.md#GetAuthJwtOidcCallback) | **Get** /auth/jwt/oidc/callback | Callback endpoint to complete an OIDC login.
[**GetAuthJwtRole**](Auth.md#GetAuthJwtRole) | **Get** /auth/jwt/role | Lists all the roles registered with the backend.
[**GetAuthJwtRoleName**](Auth.md#GetAuthJwtRoleName) | **Get** /auth/jwt/role/{name} | Read an existing role.
[**GetAuthKerberosConfig**](Auth.md#GetAuthKerberosConfig) | **Get** /auth/kerberos/config | 
[**GetAuthKerberosConfigLdap**](Auth.md#GetAuthKerberosConfigLdap) | **Get** /auth/kerberos/config/ldap | 
[**GetAuthKerberosGroups**](Auth.md#GetAuthKerberosGroups) | **Get** /auth/kerberos/groups | 
[**GetAuthKerberosGroupsName**](Auth.md#GetAuthKerberosGroupsName) | **Get** /auth/kerberos/groups/{name} | 
[**GetAuthKerberosLogin**](Auth.md#GetAuthKerberosLogin) | **Get** /auth/kerberos/login | 
[**GetAuthKubernetesConfig**](Auth.md#GetAuthKubernetesConfig) | **Get** /auth/kubernetes/config | Configures the JWT Public Key and Kubernetes API information.
[**GetAuthKubernetesRole**](Auth.md#GetAuthKubernetesRole) | **Get** /auth/kubernetes/role | Lists all the roles registered with the backend.
[**GetAuthKubernetesRoleName**](Auth.md#GetAuthKubernetesRoleName) | **Get** /auth/kubernetes/role/{name} | Register an role with the backend.
[**GetAuthLdapConfig**](Auth.md#GetAuthLdapConfig) | **Get** /auth/ldap/config | Configure the LDAP server to connect to, along with its options.
[**GetAuthLdapGroups**](Auth.md#GetAuthLdapGroups) | **Get** /auth/ldap/groups | Manage additional groups for users allowed to authenticate.
[**GetAuthLdapGroupsName**](Auth.md#GetAuthLdapGroupsName) | **Get** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
[**GetAuthLdapUsers**](Auth.md#GetAuthLdapUsers) | **Get** /auth/ldap/users | Manage users allowed to authenticate.
[**GetAuthLdapUsersName**](Auth.md#GetAuthLdapUsersName) | **Get** /auth/ldap/users/{name} | Manage users allowed to authenticate.
[**GetAuthOciConfig**](Auth.md#GetAuthOciConfig) | **Get** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
[**GetAuthOciRole**](Auth.md#GetAuthOciRole) | **Get** /auth/oci/role | Lists all the roles that are registered with Vault.
[**GetAuthOciRoleRole**](Auth.md#GetAuthOciRoleRole) | **Get** /auth/oci/role/{role} | Create a role and associate policies to it.
[**GetAuthOidcConfig**](Auth.md#GetAuthOidcConfig) | **Get** /auth/oidc/config | Read the current JWT authentication backend configuration.
[**GetAuthOidcOidcCallback**](Auth.md#GetAuthOidcOidcCallback) | **Get** /auth/oidc/oidc/callback | Callback endpoint to complete an OIDC login.
[**GetAuthOidcRole**](Auth.md#GetAuthOidcRole) | **Get** /auth/oidc/role | Lists all the roles registered with the backend.
[**GetAuthOidcRoleName**](Auth.md#GetAuthOidcRoleName) | **Get** /auth/oidc/role/{name} | Read an existing role.
[**GetAuthOktaConfig**](Auth.md#GetAuthOktaConfig) | **Get** /auth/okta/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**GetAuthOktaGroups**](Auth.md#GetAuthOktaGroups) | **Get** /auth/okta/groups | Manage users allowed to authenticate.
[**GetAuthOktaGroupsName**](Auth.md#GetAuthOktaGroupsName) | **Get** /auth/okta/groups/{name} | Manage users allowed to authenticate.
[**GetAuthOktaUsers**](Auth.md#GetAuthOktaUsers) | **Get** /auth/okta/users | Manage additional groups for users allowed to authenticate.
[**GetAuthOktaUsersName**](Auth.md#GetAuthOktaUsersName) | **Get** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
[**GetAuthOktaVerifyNonce**](Auth.md#GetAuthOktaVerifyNonce) | **Get** /auth/okta/verify/{nonce} | 
[**GetAuthRadiusConfig**](Auth.md#GetAuthRadiusConfig) | **Get** /auth/radius/config | Configure the RADIUS server to connect to, along with its options.
[**GetAuthRadiusUsers**](Auth.md#GetAuthRadiusUsers) | **Get** /auth/radius/users | Manage users allowed to authenticate.
[**GetAuthRadiusUsersName**](Auth.md#GetAuthRadiusUsersName) | **Get** /auth/radius/users/{name} | Manage users allowed to authenticate.
[**GetAuthTokenAccessors**](Auth.md#GetAuthTokenAccessors) | **Get** /auth/token/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#39;sudo&#39; capability in addition to &#39;list&#39;.
[**GetAuthTokenLookup**](Auth.md#GetAuthTokenLookup) | **Get** /auth/token/lookup | This endpoint will lookup a token and its properties.
[**GetAuthTokenLookupSelf**](Auth.md#GetAuthTokenLookupSelf) | **Get** /auth/token/lookup-self | This endpoint will lookup a token and its properties.
[**GetAuthTokenRoles**](Auth.md#GetAuthTokenRoles) | **Get** /auth/token/roles | This endpoint lists configured roles.
[**GetAuthTokenRolesRoleName**](Auth.md#GetAuthTokenRolesRoleName) | **Get** /auth/token/roles/{role_name} | 
[**GetAuthUserpassUsers**](Auth.md#GetAuthUserpassUsers) | **Get** /auth/userpass/users | Manage users allowed to authenticate.
[**GetAuthUserpassUsersUsername**](Auth.md#GetAuthUserpassUsersUsername) | **Get** /auth/userpass/users/{username} | Manage users allowed to authenticate.
[**PostAuthAlicloudLogin**](Auth.md#PostAuthAlicloudLogin) | **Post** /auth/alicloud/login | Authenticates an RAM entity with Vault.
[**PostAuthAlicloudRoleRole**](Auth.md#PostAuthAlicloudRoleRole) | **Post** /auth/alicloud/role/{role} | Create a role and associate policies to it.
[**PostAuthAppIdLogin**](Auth.md#PostAuthAppIdLogin) | **Post** /auth/app-id/login | Log in with an App ID and User ID.
[**PostAuthAppIdLoginAppId**](Auth.md#PostAuthAppIdLoginAppId) | **Post** /auth/app-id/login/{app_id} | Log in with an App ID and User ID.
[**PostAuthAppIdMapAppIdKey**](Auth.md#PostAuthAppIdMapAppIdKey) | **Post** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
[**PostAuthAppIdMapUserIdKey**](Auth.md#PostAuthAppIdMapUserIdKey) | **Post** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
[**PostAuthApproleLogin**](Auth.md#PostAuthApproleLogin) | **Post** /auth/approle/login | 
[**PostAuthApproleRoleRoleName**](Auth.md#PostAuthApproleRoleRoleName) | **Post** /auth/approle/role/{role_name} | Register an role with the backend.
[**PostAuthApproleRoleRoleNameBindSecretId**](Auth.md#PostAuthApproleRoleRoleNameBindSecretId) | **Post** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
[**PostAuthApproleRoleRoleNameBoundCidrList**](Auth.md#PostAuthApproleRoleRoleNameBoundCidrList) | **Post** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**PostAuthApproleRoleRoleNameCustomSecretId**](Auth.md#PostAuthApproleRoleRoleNameCustomSecretId) | **Post** /auth/approle/role/{role_name}/custom-secret-id | Assign a SecretID of choice against the role.
[**PostAuthApproleRoleRoleNamePeriod**](Auth.md#PostAuthApproleRoleRoleNamePeriod) | **Post** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
[**PostAuthApproleRoleRoleNamePolicies**](Auth.md#PostAuthApproleRoleRoleNamePolicies) | **Post** /auth/approle/role/{role_name}/policies | Policies of the role.
[**PostAuthApproleRoleRoleNameRoleId**](Auth.md#PostAuthApproleRoleRoleNameRoleId) | **Post** /auth/approle/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
[**PostAuthApproleRoleRoleNameSecretId**](Auth.md#PostAuthApproleRoleRoleNameSecretId) | **Post** /auth/approle/role/{role_name}/secret-id | Generate a SecretID against this role.
[**PostAuthApproleRoleRoleNameSecretIdAccessorDestroy**](Auth.md#PostAuthApproleRoleRoleNameSecretIdAccessorDestroy) | **Post** /auth/approle/role/{role_name}/secret-id-accessor/destroy | 
[**PostAuthApproleRoleRoleNameSecretIdAccessorLookup**](Auth.md#PostAuthApproleRoleRoleNameSecretIdAccessorLookup) | **Post** /auth/approle/role/{role_name}/secret-id-accessor/lookup | 
[**PostAuthApproleRoleRoleNameSecretIdBoundCidrs**](Auth.md#PostAuthApproleRoleRoleNameSecretIdBoundCidrs) | **Post** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
[**PostAuthApproleRoleRoleNameSecretIdDestroy**](Auth.md#PostAuthApproleRoleRoleNameSecretIdDestroy) | **Post** /auth/approle/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
[**PostAuthApproleRoleRoleNameSecretIdLookup**](Auth.md#PostAuthApproleRoleRoleNameSecretIdLookup) | **Post** /auth/approle/role/{role_name}/secret-id/lookup | Read the properties of an issued secret_id
[**PostAuthApproleRoleRoleNameSecretIdNumUses**](Auth.md#PostAuthApproleRoleRoleNameSecretIdNumUses) | **Post** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
[**PostAuthApproleRoleRoleNameSecretIdTtl**](Auth.md#PostAuthApproleRoleRoleNameSecretIdTtl) | **Post** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
[**PostAuthApproleRoleRoleNameTokenBoundCidrs**](Auth.md#PostAuthApproleRoleRoleNameTokenBoundCidrs) | **Post** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
[**PostAuthApproleRoleRoleNameTokenMaxTtl**](Auth.md#PostAuthApproleRoleRoleNameTokenMaxTtl) | **Post** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
[**PostAuthApproleRoleRoleNameTokenNumUses**](Auth.md#PostAuthApproleRoleRoleNameTokenNumUses) | **Post** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
[**PostAuthApproleRoleRoleNameTokenTtl**](Auth.md#PostAuthApproleRoleRoleNameTokenTtl) | **Post** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
[**PostAuthApproleTidySecretId**](Auth.md#PostAuthApproleTidySecretId) | **Post** /auth/approle/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
[**PostAuthAwsConfigCertificateCertName**](Auth.md#PostAuthAwsConfigCertificateCertName) | **Post** /auth/aws/config/certificate/{cert_name} | 
[**PostAuthAwsConfigClient**](Auth.md#PostAuthAwsConfigClient) | **Post** /auth/aws/config/client | 
[**PostAuthAwsConfigIdentity**](Auth.md#PostAuthAwsConfigIdentity) | **Post** /auth/aws/config/identity | 
[**PostAuthAwsConfigRotateRoot**](Auth.md#PostAuthAwsConfigRotateRoot) | **Post** /auth/aws/config/rotate-root | 
[**PostAuthAwsConfigStsAccountId**](Auth.md#PostAuthAwsConfigStsAccountId) | **Post** /auth/aws/config/sts/{account_id} | 
[**PostAuthAwsConfigTidyIdentityAccesslist**](Auth.md#PostAuthAwsConfigTidyIdentityAccesslist) | **Post** /auth/aws/config/tidy/identity-accesslist | 
[**PostAuthAwsConfigTidyIdentityWhitelist**](Auth.md#PostAuthAwsConfigTidyIdentityWhitelist) | **Post** /auth/aws/config/tidy/identity-whitelist | 
[**PostAuthAwsConfigTidyRoletagBlacklist**](Auth.md#PostAuthAwsConfigTidyRoletagBlacklist) | **Post** /auth/aws/config/tidy/roletag-blacklist | 
[**PostAuthAwsConfigTidyRoletagDenylist**](Auth.md#PostAuthAwsConfigTidyRoletagDenylist) | **Post** /auth/aws/config/tidy/roletag-denylist | 
[**PostAuthAwsLogin**](Auth.md#PostAuthAwsLogin) | **Post** /auth/aws/login | 
[**PostAuthAwsRoleRole**](Auth.md#PostAuthAwsRoleRole) | **Post** /auth/aws/role/{role} | 
[**PostAuthAwsRoleRoleTag**](Auth.md#PostAuthAwsRoleRoleTag) | **Post** /auth/aws/role/{role}/tag | 
[**PostAuthAwsRoletagBlacklistRoleTag**](Auth.md#PostAuthAwsRoletagBlacklistRoleTag) | **Post** /auth/aws/roletag-blacklist/{role_tag} | 
[**PostAuthAwsRoletagDenylistRoleTag**](Auth.md#PostAuthAwsRoletagDenylistRoleTag) | **Post** /auth/aws/roletag-denylist/{role_tag} | 
[**PostAuthAwsTidyIdentityAccesslist**](Auth.md#PostAuthAwsTidyIdentityAccesslist) | **Post** /auth/aws/tidy/identity-accesslist | 
[**PostAuthAwsTidyIdentityWhitelist**](Auth.md#PostAuthAwsTidyIdentityWhitelist) | **Post** /auth/aws/tidy/identity-whitelist | 
[**PostAuthAwsTidyRoletagBlacklist**](Auth.md#PostAuthAwsTidyRoletagBlacklist) | **Post** /auth/aws/tidy/roletag-blacklist | 
[**PostAuthAwsTidyRoletagDenylist**](Auth.md#PostAuthAwsTidyRoletagDenylist) | **Post** /auth/aws/tidy/roletag-denylist | 
[**PostAuthAzureConfig**](Auth.md#PostAuthAzureConfig) | **Post** /auth/azure/config | 
[**PostAuthAzureLogin**](Auth.md#PostAuthAzureLogin) | **Post** /auth/azure/login | 
[**PostAuthAzureRoleName**](Auth.md#PostAuthAzureRoleName) | **Post** /auth/azure/role/{name} | 
[**PostAuthCentrifyConfig**](Auth.md#PostAuthCentrifyConfig) | **Post** /auth/centrify/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**PostAuthCentrifyLogin**](Auth.md#PostAuthCentrifyLogin) | **Post** /auth/centrify/login | Log in with a username and password.
[**PostAuthCertCertsName**](Auth.md#PostAuthCertCertsName) | **Post** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
[**PostAuthCertConfig**](Auth.md#PostAuthCertConfig) | **Post** /auth/cert/config | 
[**PostAuthCertCrlsName**](Auth.md#PostAuthCertCrlsName) | **Post** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**PostAuthCertLogin**](Auth.md#PostAuthCertLogin) | **Post** /auth/cert/login | 
[**PostAuthCfConfig**](Auth.md#PostAuthCfConfig) | **Post** /auth/cf/config | 
[**PostAuthCfLogin**](Auth.md#PostAuthCfLogin) | **Post** /auth/cf/login | 
[**PostAuthCfRolesRole**](Auth.md#PostAuthCfRolesRole) | **Post** /auth/cf/roles/{role} | 
[**PostAuthGcpConfig**](Auth.md#PostAuthGcpConfig) | **Post** /auth/gcp/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**PostAuthGcpLogin**](Auth.md#PostAuthGcpLogin) | **Post** /auth/gcp/login | 
[**PostAuthGcpRoleName**](Auth.md#PostAuthGcpRoleName) | **Post** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
[**PostAuthGcpRoleNameLabels**](Auth.md#PostAuthGcpRoleNameLabels) | **Post** /auth/gcp/role/{name}/labels | Add or remove labels for an existing &#39;gce&#39; role
[**PostAuthGcpRoleNameServiceAccounts**](Auth.md#PostAuthGcpRoleNameServiceAccounts) | **Post** /auth/gcp/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**PostAuthGithubConfig**](Auth.md#PostAuthGithubConfig) | **Post** /auth/github/config | 
[**PostAuthGithubLogin**](Auth.md#PostAuthGithubLogin) | **Post** /auth/github/login | 
[**PostAuthGithubMapTeamsKey**](Auth.md#PostAuthGithubMapTeamsKey) | **Post** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
[**PostAuthGithubMapUsersKey**](Auth.md#PostAuthGithubMapUsersKey) | **Post** /auth/github/map/users/{key} | Read/write/delete a single users mapping
[**PostAuthJwtConfig**](Auth.md#PostAuthJwtConfig) | **Post** /auth/jwt/config | Configure the JWT authentication backend.
[**PostAuthJwtLogin**](Auth.md#PostAuthJwtLogin) | **Post** /auth/jwt/login | Authenticates to Vault using a JWT (or OIDC) token.
[**PostAuthJwtOidcAuthUrl**](Auth.md#PostAuthJwtOidcAuthUrl) | **Post** /auth/jwt/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**PostAuthJwtOidcCallback**](Auth.md#PostAuthJwtOidcCallback) | **Post** /auth/jwt/oidc/callback | Callback endpoint to handle form_posts.
[**PostAuthJwtRoleName**](Auth.md#PostAuthJwtRoleName) | **Post** /auth/jwt/role/{name} | Register an role with the backend.
[**PostAuthKerberosConfig**](Auth.md#PostAuthKerberosConfig) | **Post** /auth/kerberos/config | 
[**PostAuthKerberosConfigLdap**](Auth.md#PostAuthKerberosConfigLdap) | **Post** /auth/kerberos/config/ldap | 
[**PostAuthKerberosGroupsName**](Auth.md#PostAuthKerberosGroupsName) | **Post** /auth/kerberos/groups/{name} | 
[**PostAuthKerberosLogin**](Auth.md#PostAuthKerberosLogin) | **Post** /auth/kerberos/login | 
[**PostAuthKubernetesConfig**](Auth.md#PostAuthKubernetesConfig) | **Post** /auth/kubernetes/config | Configures the JWT Public Key and Kubernetes API information.
[**PostAuthKubernetesLogin**](Auth.md#PostAuthKubernetesLogin) | **Post** /auth/kubernetes/login | Authenticates Kubernetes service accounts with Vault.
[**PostAuthKubernetesRoleName**](Auth.md#PostAuthKubernetesRoleName) | **Post** /auth/kubernetes/role/{name} | Register an role with the backend.
[**PostAuthLdapConfig**](Auth.md#PostAuthLdapConfig) | **Post** /auth/ldap/config | Configure the LDAP server to connect to, along with its options.
[**PostAuthLdapGroupsName**](Auth.md#PostAuthLdapGroupsName) | **Post** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
[**PostAuthLdapLoginUsername**](Auth.md#PostAuthLdapLoginUsername) | **Post** /auth/ldap/login/{username} | Log in with a username and password.
[**PostAuthLdapUsersName**](Auth.md#PostAuthLdapUsersName) | **Post** /auth/ldap/users/{name} | Manage users allowed to authenticate.
[**PostAuthOciConfig**](Auth.md#PostAuthOciConfig) | **Post** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
[**PostAuthOciLoginRole**](Auth.md#PostAuthOciLoginRole) | **Post** /auth/oci/login/{role} | Authenticates to Vault using OCI credentials
[**PostAuthOciRoleRole**](Auth.md#PostAuthOciRoleRole) | **Post** /auth/oci/role/{role} | Create a role and associate policies to it.
[**PostAuthOidcConfig**](Auth.md#PostAuthOidcConfig) | **Post** /auth/oidc/config | Configure the JWT authentication backend.
[**PostAuthOidcLogin**](Auth.md#PostAuthOidcLogin) | **Post** /auth/oidc/login | Authenticates to Vault using a JWT (or OIDC) token.
[**PostAuthOidcOidcAuthUrl**](Auth.md#PostAuthOidcOidcAuthUrl) | **Post** /auth/oidc/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**PostAuthOidcOidcCallback**](Auth.md#PostAuthOidcOidcCallback) | **Post** /auth/oidc/oidc/callback | Callback endpoint to handle form_posts.
[**PostAuthOidcRoleName**](Auth.md#PostAuthOidcRoleName) | **Post** /auth/oidc/role/{name} | Register an role with the backend.
[**PostAuthOktaConfig**](Auth.md#PostAuthOktaConfig) | **Post** /auth/okta/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**PostAuthOktaGroupsName**](Auth.md#PostAuthOktaGroupsName) | **Post** /auth/okta/groups/{name} | Manage users allowed to authenticate.
[**PostAuthOktaLoginUsername**](Auth.md#PostAuthOktaLoginUsername) | **Post** /auth/okta/login/{username} | Log in with a username and password.
[**PostAuthOktaUsersName**](Auth.md#PostAuthOktaUsersName) | **Post** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
[**PostAuthRadiusConfig**](Auth.md#PostAuthRadiusConfig) | **Post** /auth/radius/config | Configure the RADIUS server to connect to, along with its options.
[**PostAuthRadiusLogin**](Auth.md#PostAuthRadiusLogin) | **Post** /auth/radius/login | Log in with a username and password.
[**PostAuthRadiusLoginUrlusername**](Auth.md#PostAuthRadiusLoginUrlusername) | **Post** /auth/radius/login/{urlusername} | Log in with a username and password.
[**PostAuthRadiusUsersName**](Auth.md#PostAuthRadiusUsersName) | **Post** /auth/radius/users/{name} | Manage users allowed to authenticate.
[**PostAuthTokenCreate**](Auth.md#PostAuthTokenCreate) | **Post** /auth/token/create | The token create path is used to create new tokens.
[**PostAuthTokenCreateOrphan**](Auth.md#PostAuthTokenCreateOrphan) | **Post** /auth/token/create-orphan | The token create path is used to create new orphan tokens.
[**PostAuthTokenCreateRoleName**](Auth.md#PostAuthTokenCreateRoleName) | **Post** /auth/token/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**PostAuthTokenLookup**](Auth.md#PostAuthTokenLookup) | **Post** /auth/token/lookup | This endpoint will lookup a token and its properties.
[**PostAuthTokenLookupAccessor**](Auth.md#PostAuthTokenLookupAccessor) | **Post** /auth/token/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**PostAuthTokenLookupSelf**](Auth.md#PostAuthTokenLookupSelf) | **Post** /auth/token/lookup-self | This endpoint will lookup a token and its properties.
[**PostAuthTokenRenew**](Auth.md#PostAuthTokenRenew) | **Post** /auth/token/renew | This endpoint will renew the given token and prevent expiration.
[**PostAuthTokenRenewAccessor**](Auth.md#PostAuthTokenRenewAccessor) | **Post** /auth/token/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**PostAuthTokenRenewSelf**](Auth.md#PostAuthTokenRenewSelf) | **Post** /auth/token/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**PostAuthTokenRevoke**](Auth.md#PostAuthTokenRevoke) | **Post** /auth/token/revoke | This endpoint will delete the given token and all of its child tokens.
[**PostAuthTokenRevokeAccessor**](Auth.md#PostAuthTokenRevokeAccessor) | **Post** /auth/token/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**PostAuthTokenRevokeOrphan**](Auth.md#PostAuthTokenRevokeOrphan) | **Post** /auth/token/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**PostAuthTokenRevokeSelf**](Auth.md#PostAuthTokenRevokeSelf) | **Post** /auth/token/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**PostAuthTokenRolesRoleName**](Auth.md#PostAuthTokenRolesRoleName) | **Post** /auth/token/roles/{role_name} | 
[**PostAuthTokenTidy**](Auth.md#PostAuthTokenTidy) | **Post** /auth/token/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**PostAuthUserpassLoginUsername**](Auth.md#PostAuthUserpassLoginUsername) | **Post** /auth/userpass/login/{username} | Log in with a username and password.
[**PostAuthUserpassUsersUsername**](Auth.md#PostAuthUserpassUsersUsername) | **Post** /auth/userpass/users/{username} | Manage users allowed to authenticate.
[**PostAuthUserpassUsersUsernamePassword**](Auth.md#PostAuthUserpassUsersUsernamePassword) | **Post** /auth/userpass/users/{username}/password | Reset user&#39;s password.
[**PostAuthUserpassUsersUsernamePolicies**](Auth.md#PostAuthUserpassUsersUsernamePolicies) | **Post** /auth/userpass/users/{username}/policies | Update the policies associated with the username.



## DeleteAuthAlicloudRoleRole

> DeleteAuthAlicloudRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role as it should appear in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAlicloudRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAlicloudRoleRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAlicloudRoleRoleRequest struct via the builder pattern


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


## DeleteAuthAppIdMapAppIdKey

> DeleteAuthAppIdMapAppIdKey(ctx, key).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the app-id mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAppIdMapAppIdKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAppIdMapAppIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the app-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAppIdMapAppIdKeyRequest struct via the builder pattern


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


## DeleteAuthAppIdMapUserIdKey

> DeleteAuthAppIdMapUserIdKey(ctx, key).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the user-id mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAppIdMapUserIdKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAppIdMapUserIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the user-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAppIdMapUserIdKeyRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleName

> DeleteAuthApproleRoleRoleName(ctx, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleName(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameBindSecretId

> DeleteAuthApproleRoleRoleNameBindSecretId(ctx, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameBindSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameBindSecretIdRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameBoundCidrList

> DeleteAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameBoundCidrList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameBoundCidrListRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNamePeriod

> DeleteAuthApproleRoleRoleNamePeriod(ctx, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNamePeriod(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNamePeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNamePeriodRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNamePolicies

> DeleteAuthApproleRoleRoleNamePolicies(ctx, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNamePolicies(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNamePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNamePoliciesRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy

> DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, roleName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameSecretIdAccessorDestroyRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs

> DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameSecretIdBoundCidrsRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameSecretIdDestroy

> DeleteAuthApproleRoleRoleNameSecretIdDestroy(ctx, roleName).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameSecretIdDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameSecretIdDestroyRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameSecretIdNumUses

> DeleteAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameSecretIdNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameSecretIdNumUsesRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameSecretIdTtl

> DeleteAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameSecretIdTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameSecretIdTtlRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameTokenBoundCidrs

> DeleteAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameTokenBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameTokenBoundCidrsRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameTokenMaxTtl

> DeleteAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameTokenMaxTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameTokenMaxTtlRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameTokenNumUses

> DeleteAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameTokenNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameTokenNumUsesRequest struct via the builder pattern


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


## DeleteAuthApproleRoleRoleNameTokenTtl

> DeleteAuthApproleRoleRoleNameTokenTtl(ctx, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthApproleRoleRoleNameTokenTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthApproleRoleRoleNameTokenTtlRequest struct via the builder pattern


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


## DeleteAuthAwsConfigCertificateCertName

> DeleteAuthAwsConfigCertificateCertName(ctx, certName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certName := "certName_example" // string | Name of the certificate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigCertificateCertName(context.Background(), certName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigCertificateCertName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certName** | **string** | Name of the certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigCertificateCertNameRequest struct via the builder pattern


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


## DeleteAuthAwsConfigClient

> DeleteAuthAwsConfigClient(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigClientRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigStsAccountId

> DeleteAuthAwsConfigStsAccountId(ctx, accountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigStsAccountId(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigStsAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigStsAccountIdRequest struct via the builder pattern


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


## DeleteAuthAwsConfigTidyIdentityAccesslist

> DeleteAuthAwsConfigTidyIdentityAccesslist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigTidyIdentityAccesslist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigTidyIdentityAccesslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigTidyIdentityAccesslistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigTidyIdentityWhitelist

> DeleteAuthAwsConfigTidyIdentityWhitelist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigTidyIdentityWhitelist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigTidyIdentityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigTidyIdentityWhitelistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigTidyRoletagBlacklist

> DeleteAuthAwsConfigTidyRoletagBlacklist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigTidyRoletagBlacklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigTidyRoletagBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigTidyRoletagBlacklistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigTidyRoletagDenylist

> DeleteAuthAwsConfigTidyRoletagDenylist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsConfigTidyRoletagDenylist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsConfigTidyRoletagDenylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsConfigTidyRoletagDenylistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAwsIdentityAccesslistInstanceId

> DeleteAuthAwsIdentityAccesslistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsIdentityAccesslistInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsIdentityAccesslistInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsIdentityAccesslistInstanceIdRequest struct via the builder pattern


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


## DeleteAuthAwsIdentityWhitelistInstanceId

> DeleteAuthAwsIdentityWhitelistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsIdentityWhitelistInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsIdentityWhitelistInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsIdentityWhitelistInstanceIdRequest struct via the builder pattern


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


## DeleteAuthAwsRoleRole

> DeleteAuthAwsRoleRole(ctx, role).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAwsRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthAwsRoleRoleRequest struct via the builder pattern


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


## DeleteAuthAwsRoletagBlacklistRoleTag

> DeleteAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsRoletagBlacklistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsRoletagBlacklistRoleTagRequest struct via the builder pattern


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


## DeleteAuthAwsRoletagDenylistRoleTag

> DeleteAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAwsRoletagDenylistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAwsRoletagDenylistRoleTagRequest struct via the builder pattern


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


## DeleteAuthAzureConfig

> DeleteAuthAzureConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthAzureConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthAzureRoleName

> DeleteAuthAzureRoleName(ctx, name).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthAzureRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthAzureRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthAzureRoleNameRequest struct via the builder pattern


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


## DeleteAuthCertCertsName

> DeleteAuthCertCertsName(ctx, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthCertCertsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthCertCertsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCertCertsNameRequest struct via the builder pattern


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


## DeleteAuthCertCrlsName

> DeleteAuthCertCrlsName(ctx, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthCertCrlsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthCertCrlsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCertCrlsNameRequest struct via the builder pattern


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


## DeleteAuthCfConfig

> DeleteAuthCfConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthCfConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthCfConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCfConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthCfRolesRole

> DeleteAuthCfRolesRole(ctx, role).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthCfRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthCfRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCfRolesRoleRequest struct via the builder pattern


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


## DeleteAuthGcpRoleName

> DeleteAuthGcpRoleName(ctx, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthGcpRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthGcpRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthGcpRoleNameRequest struct via the builder pattern


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


## DeleteAuthGithubMapTeamsKey

> DeleteAuthGithubMapTeamsKey(ctx, key).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the teams mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthGithubMapTeamsKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthGithubMapTeamsKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the teams mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthGithubMapTeamsKeyRequest struct via the builder pattern


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


## DeleteAuthGithubMapUsersKey

> DeleteAuthGithubMapUsersKey(ctx, key).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the users mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthGithubMapUsersKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthGithubMapUsersKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the users mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthGithubMapUsersKeyRequest struct via the builder pattern


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


## DeleteAuthJwtRoleName

> DeleteAuthJwtRoleName(ctx, name).Execute()

Delete an existing role.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthJwtRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthJwtRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthJwtRoleNameRequest struct via the builder pattern


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


## DeleteAuthKerberosGroupsName

> DeleteAuthKerberosGroupsName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthKerberosGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthKerberosGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthKerberosGroupsNameRequest struct via the builder pattern


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


## DeleteAuthKubernetesRoleName

> DeleteAuthKubernetesRoleName(ctx, name).Execute()

Register an role with the backend.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthKubernetesRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthKubernetesRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthKubernetesRoleNameRequest struct via the builder pattern


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


## DeleteAuthLdapGroupsName

> DeleteAuthLdapGroupsName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthLdapGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthLdapGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthLdapGroupsNameRequest struct via the builder pattern


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


## DeleteAuthLdapUsersName

> DeleteAuthLdapUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthLdapUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthLdapUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthLdapUsersNameRequest struct via the builder pattern


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


## DeleteAuthOciConfig

> DeleteAuthOciConfig(ctx).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthOciConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthOciConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthOciConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthOciRoleRole

> DeleteAuthOciRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthOciRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthOciRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthOciRoleRoleRequest struct via the builder pattern


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


## DeleteAuthOidcRoleName

> DeleteAuthOidcRoleName(ctx, name).Execute()

Delete an existing role.

### Example

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
    resp, r, err := apiClient.Auth.DeleteAuthOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthOidcRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthOidcRoleNameRequest struct via the builder pattern


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


## DeleteAuthOktaGroupsName

> DeleteAuthOktaGroupsName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Okta group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthOktaGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthOktaGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Okta group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthOktaGroupsNameRequest struct via the builder pattern


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


## DeleteAuthOktaUsersName

> DeleteAuthOktaUsersName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthOktaUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthOktaUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthOktaUsersNameRequest struct via the builder pattern


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


## DeleteAuthRadiusUsersName

> DeleteAuthRadiusUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the RADIUS user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthRadiusUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthRadiusUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthRadiusUsersNameRequest struct via the builder pattern


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


## DeleteAuthTokenRolesRoleName

> DeleteAuthTokenRolesRoleName(ctx, roleName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthTokenRolesRoleName(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthTokenRolesRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthTokenRolesRoleNameRequest struct via the builder pattern


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


## DeleteAuthUserpassUsersUsername

> DeleteAuthUserpassUsersUsername(ctx, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username for this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.DeleteAuthUserpassUsersUsername(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.DeleteAuthUserpassUsersUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username for this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthUserpassUsersUsernameRequest struct via the builder pattern


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


## GetAuthAlicloudRole

> GetAuthAlicloudRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthAlicloudRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAlicloudRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAlicloudRoleRequest struct via the builder pattern


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


## GetAuthAlicloudRoleRole

> GetAuthAlicloudRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role as it should appear in Vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAlicloudRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAlicloudRoleRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAlicloudRoleRoleRequest struct via the builder pattern


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


## GetAuthAlicloudRoles

> GetAuthAlicloudRoles(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthAlicloudRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAlicloudRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAlicloudRolesRequest struct via the builder pattern


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


## GetAuthAppIdMapAppId

> GetAuthAppIdMapAppId(ctx).List(list).Execute()

Read mappings for app-id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAppIdMapAppId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAppIdMapAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAppIdMapAppIdRequest struct via the builder pattern


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


## GetAuthAppIdMapAppIdKey

> GetAuthAppIdMapAppIdKey(ctx, key).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the app-id mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAppIdMapAppIdKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAppIdMapAppIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the app-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAppIdMapAppIdKeyRequest struct via the builder pattern


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


## GetAuthAppIdMapUserId

> GetAuthAppIdMapUserId(ctx).List(list).Execute()

Read mappings for user-id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAppIdMapUserId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAppIdMapUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAppIdMapUserIdRequest struct via the builder pattern


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


## GetAuthAppIdMapUserIdKey

> GetAuthAppIdMapUserIdKey(ctx, key).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the user-id mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAppIdMapUserIdKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAppIdMapUserIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the user-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAppIdMapUserIdKeyRequest struct via the builder pattern


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


## GetAuthApproleRole

> GetAuthApproleRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthApproleRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleName

> GetAuthApproleRoleRoleName(ctx, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleName(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameBindSecretId

> GetAuthApproleRoleRoleNameBindSecretId(ctx, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameBindSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameBindSecretIdRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameBoundCidrList

> GetAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameBoundCidrList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameBoundCidrListRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameLocalSecretIds

> GetAuthApproleRoleRoleNameLocalSecretIds(ctx, roleName).Execute()

Enables cluster local secret IDs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameLocalSecretIds(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameLocalSecretIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameLocalSecretIdsRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNamePeriod

> GetAuthApproleRoleRoleNamePeriod(ctx, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNamePeriod(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNamePeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNamePeriodRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNamePolicies

> GetAuthApproleRoleRoleNamePolicies(ctx, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNamePolicies(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNamePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNamePoliciesRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameRoleId

> GetAuthApproleRoleRoleNameRoleId(ctx, roleName).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameRoleId(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameRoleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameRoleIdRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameSecretId

> GetAuthApproleRoleRoleNameSecretId(ctx, roleName).List(list).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameSecretId(context.Background(), roleName).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameSecretIdRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameSecretIdBoundCidrs

> GetAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameSecretIdBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameSecretIdBoundCidrsRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameSecretIdNumUses

> GetAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameSecretIdNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameSecretIdNumUsesRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameSecretIdTtl

> GetAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameSecretIdTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameSecretIdTtlRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameTokenBoundCidrs

> GetAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameTokenBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameTokenBoundCidrsRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameTokenMaxTtl

> GetAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameTokenMaxTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameTokenMaxTtlRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameTokenNumUses

> GetAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameTokenNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameTokenNumUsesRequest struct via the builder pattern


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


## GetAuthApproleRoleRoleNameTokenTtl

> GetAuthApproleRoleRoleNameTokenTtl(ctx, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthApproleRoleRoleNameTokenTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthApproleRoleRoleNameTokenTtlRequest struct via the builder pattern


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


## GetAuthAwsConfigCertificateCertName

> GetAuthAwsConfigCertificateCertName(ctx, certName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certName := "certName_example" // string | Name of the certificate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsConfigCertificateCertName(context.Background(), certName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigCertificateCertName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certName** | **string** | Name of the certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigCertificateCertNameRequest struct via the builder pattern


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


## GetAuthAwsConfigCertificates

> GetAuthAwsConfigCertificates(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigCertificates(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigCertificatesRequest struct via the builder pattern


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


## GetAuthAwsConfigClient

> GetAuthAwsConfigClient(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigClientRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsConfigIdentity

> GetAuthAwsConfigIdentity(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigIdentity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigIdentityRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsConfigSts

> GetAuthAwsConfigSts(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigSts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigStsRequest struct via the builder pattern


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


## GetAuthAwsConfigStsAccountId

> GetAuthAwsConfigStsAccountId(ctx, accountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsConfigStsAccountId(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigStsAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigStsAccountIdRequest struct via the builder pattern


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


## GetAuthAwsConfigTidyIdentityAccesslist

> GetAuthAwsConfigTidyIdentityAccesslist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigTidyIdentityAccesslist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigTidyIdentityAccesslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigTidyIdentityAccesslistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsConfigTidyIdentityWhitelist

> GetAuthAwsConfigTidyIdentityWhitelist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigTidyIdentityWhitelist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigTidyIdentityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigTidyIdentityWhitelistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsConfigTidyRoletagBlacklist

> GetAuthAwsConfigTidyRoletagBlacklist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigTidyRoletagBlacklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigTidyRoletagBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigTidyRoletagBlacklistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsConfigTidyRoletagDenylist

> GetAuthAwsConfigTidyRoletagDenylist(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsConfigTidyRoletagDenylist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsConfigTidyRoletagDenylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsConfigTidyRoletagDenylistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAwsIdentityAccesslist

> GetAuthAwsIdentityAccesslist(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsIdentityAccesslist(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsIdentityAccesslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsIdentityAccesslistRequest struct via the builder pattern


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


## GetAuthAwsIdentityAccesslistInstanceId

> GetAuthAwsIdentityAccesslistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsIdentityAccesslistInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsIdentityAccesslistInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsIdentityAccesslistInstanceIdRequest struct via the builder pattern


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


## GetAuthAwsIdentityWhitelist

> GetAuthAwsIdentityWhitelist(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsIdentityWhitelist(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsIdentityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsIdentityWhitelistRequest struct via the builder pattern


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


## GetAuthAwsIdentityWhitelistInstanceId

> GetAuthAwsIdentityWhitelistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsIdentityWhitelistInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsIdentityWhitelistInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsIdentityWhitelistInstanceIdRequest struct via the builder pattern


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


## GetAuthAwsRole

> GetAuthAwsRole(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRoleRequest struct via the builder pattern


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


## GetAuthAwsRoleRole

> GetAuthAwsRoleRole(ctx, role).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthAwsRoleRoleRequest struct via the builder pattern


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


## GetAuthAwsRoles

> GetAuthAwsRoles(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRolesRequest struct via the builder pattern


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


## GetAuthAwsRoletagBlacklist

> GetAuthAwsRoletagBlacklist(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsRoletagBlacklist(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoletagBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRoletagBlacklistRequest struct via the builder pattern


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


## GetAuthAwsRoletagBlacklistRoleTag

> GetAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoletagBlacklistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRoletagBlacklistRoleTagRequest struct via the builder pattern


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


## GetAuthAwsRoletagDenylist

> GetAuthAwsRoletagDenylist(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAwsRoletagDenylist(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoletagDenylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRoletagDenylistRequest struct via the builder pattern


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


## GetAuthAwsRoletagDenylistRoleTag

> GetAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAwsRoletagDenylistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAwsRoletagDenylistRoleTagRequest struct via the builder pattern


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


## GetAuthAzureConfig

> GetAuthAzureConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAzureConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAzureConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAzureRole

> GetAuthAzureRole(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAzureRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAzureRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAzureRoleRequest struct via the builder pattern


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


## GetAuthAzureRoleName

> GetAuthAzureRoleName(ctx, name).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthAzureRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthAzureRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthAzureRoleNameRequest struct via the builder pattern


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


## GetAuthCentrifyConfig

> GetAuthCentrifyConfig(ctx).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthCentrifyConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCentrifyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCentrifyConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthCertCerts

> GetAuthCertCerts(ctx).List(list).Execute()

Manage trusted certificates used for authentication.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthCertCerts(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCertCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCertCertsRequest struct via the builder pattern


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


## GetAuthCertCertsName

> GetAuthCertCertsName(ctx, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthCertCertsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCertCertsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCertCertsNameRequest struct via the builder pattern


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


## GetAuthCertCrlsName

> GetAuthCertCrlsName(ctx, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthCertCrlsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCertCrlsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCertCrlsNameRequest struct via the builder pattern


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


## GetAuthCfConfig

> GetAuthCfConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthCfConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCfConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCfConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthCfRoles

> GetAuthCfRoles(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthCfRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCfRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCfRolesRequest struct via the builder pattern


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


## GetAuthCfRolesRole

> GetAuthCfRolesRole(ctx, role).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthCfRolesRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthCfRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCfRolesRoleRequest struct via the builder pattern


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


## GetAuthGcpConfig

> GetAuthGcpConfig(ctx).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

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
    resp, r, err := apiClient.Auth.GetAuthGcpConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGcpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGcpConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthGcpRole

> GetAuthGcpRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthGcpRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGcpRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGcpRoleRequest struct via the builder pattern


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


## GetAuthGcpRoleName

> GetAuthGcpRoleName(ctx, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthGcpRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGcpRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthGcpRoleNameRequest struct via the builder pattern


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


## GetAuthGcpRoles

> GetAuthGcpRoles(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthGcpRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGcpRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGcpRolesRequest struct via the builder pattern


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


## GetAuthGithubConfig

> GetAuthGithubConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthGithubConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGithubConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGithubConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthGithubMapTeams

> GetAuthGithubMapTeams(ctx).List(list).Execute()

Read mappings for teams

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthGithubMapTeams(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGithubMapTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGithubMapTeamsRequest struct via the builder pattern


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


## GetAuthGithubMapTeamsKey

> GetAuthGithubMapTeamsKey(ctx, key).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the teams mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthGithubMapTeamsKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGithubMapTeamsKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the teams mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGithubMapTeamsKeyRequest struct via the builder pattern


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


## GetAuthGithubMapUsers

> GetAuthGithubMapUsers(ctx).List(list).Execute()

Read mappings for users

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthGithubMapUsers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGithubMapUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGithubMapUsersRequest struct via the builder pattern


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


## GetAuthGithubMapUsersKey

> GetAuthGithubMapUsersKey(ctx, key).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the users mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthGithubMapUsersKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthGithubMapUsersKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the users mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthGithubMapUsersKeyRequest struct via the builder pattern


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


## GetAuthJwtConfig

> GetAuthJwtConfig(ctx).Execute()

Read the current JWT authentication backend configuration.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthJwtConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthJwtConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthJwtConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthJwtOidcCallback

> GetAuthJwtOidcCallback(ctx).Execute()

Callback endpoint to complete an OIDC login.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthJwtOidcCallback(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthJwtOidcCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthJwtOidcCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthJwtRole

> GetAuthJwtRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.



### Example

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
    resp, r, err := apiClient.Auth.GetAuthJwtRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthJwtRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthJwtRoleRequest struct via the builder pattern


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


## GetAuthJwtRoleName

> GetAuthJwtRoleName(ctx, name).Execute()

Read an existing role.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthJwtRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthJwtRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthJwtRoleNameRequest struct via the builder pattern


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


## GetAuthKerberosConfig

> GetAuthKerberosConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthKerberosConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKerberosConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKerberosConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthKerberosConfigLdap

> GetAuthKerberosConfigLdap(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthKerberosConfigLdap(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKerberosConfigLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKerberosConfigLdapRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthKerberosGroups

> GetAuthKerberosGroups(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthKerberosGroups(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKerberosGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKerberosGroupsRequest struct via the builder pattern


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


## GetAuthKerberosGroupsName

> GetAuthKerberosGroupsName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthKerberosGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKerberosGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKerberosGroupsNameRequest struct via the builder pattern


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


## GetAuthKerberosLogin

> GetAuthKerberosLogin(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.GetAuthKerberosLogin(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKerberosLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKerberosLoginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthKubernetesConfig

> GetAuthKubernetesConfig(ctx).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthKubernetesConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKubernetesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKubernetesConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthKubernetesRole

> GetAuthKubernetesRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthKubernetesRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKubernetesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKubernetesRoleRequest struct via the builder pattern


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


## GetAuthKubernetesRoleName

> GetAuthKubernetesRoleName(ctx, name).Execute()

Register an role with the backend.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthKubernetesRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthKubernetesRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthKubernetesRoleNameRequest struct via the builder pattern


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


## GetAuthLdapConfig

> GetAuthLdapConfig(ctx).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthLdapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthLdapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLdapConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthLdapGroups

> GetAuthLdapGroups(ctx).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthLdapGroups(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthLdapGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLdapGroupsRequest struct via the builder pattern


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


## GetAuthLdapGroupsName

> GetAuthLdapGroupsName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthLdapGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthLdapGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLdapGroupsNameRequest struct via the builder pattern


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


## GetAuthLdapUsers

> GetAuthLdapUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthLdapUsers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthLdapUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLdapUsersRequest struct via the builder pattern


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


## GetAuthLdapUsersName

> GetAuthLdapUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthLdapUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthLdapUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLdapUsersNameRequest struct via the builder pattern


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


## GetAuthOciConfig

> GetAuthOciConfig(ctx).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOciConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOciConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOciConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOciRole

> GetAuthOciRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOciRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOciRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOciRoleRequest struct via the builder pattern


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


## GetAuthOciRoleRole

> GetAuthOciRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOciRoleRole(context.Background(), role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOciRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthOciRoleRoleRequest struct via the builder pattern


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


## GetAuthOidcConfig

> GetAuthOidcConfig(ctx).Execute()

Read the current JWT authentication backend configuration.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOidcConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOidcConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOidcOidcCallback

> GetAuthOidcOidcCallback(ctx).Execute()

Callback endpoint to complete an OIDC login.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOidcOidcCallback(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOidcOidcCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOidcOidcCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOidcRole

> GetAuthOidcRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.



### Example

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
    resp, r, err := apiClient.Auth.GetAuthOidcRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOidcRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOidcRoleRequest struct via the builder pattern


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


## GetAuthOidcRoleName

> GetAuthOidcRoleName(ctx, name).Execute()

Read an existing role.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOidcRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAuthOidcRoleNameRequest struct via the builder pattern


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


## GetAuthOktaConfig

> GetAuthOktaConfig(ctx).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOktaConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOktaGroups

> GetAuthOktaGroups(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOktaGroups(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaGroupsRequest struct via the builder pattern


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


## GetAuthOktaGroupsName

> GetAuthOktaGroupsName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Okta group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthOktaGroupsName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Okta group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaGroupsNameRequest struct via the builder pattern


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


## GetAuthOktaUsers

> GetAuthOktaUsers(ctx).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthOktaUsers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaUsersRequest struct via the builder pattern


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


## GetAuthOktaUsersName

> GetAuthOktaUsersName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthOktaUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaUsersNameRequest struct via the builder pattern


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


## GetAuthOktaVerifyNonce

> GetAuthOktaVerifyNonce(ctx, nonce).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nonce := "nonce_example" // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthOktaVerifyNonce(context.Background(), nonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthOktaVerifyNonce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nonce** | **string** | Nonce provided during a login request to retrieve the number verification challenge for the matching request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOktaVerifyNonceRequest struct via the builder pattern


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


## GetAuthRadiusConfig

> GetAuthRadiusConfig(ctx).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthRadiusConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthRadiusConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthRadiusConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthRadiusUsers

> GetAuthRadiusUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthRadiusUsers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthRadiusUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthRadiusUsersRequest struct via the builder pattern


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


## GetAuthRadiusUsersName

> GetAuthRadiusUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the RADIUS user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthRadiusUsersName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthRadiusUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthRadiusUsersNameRequest struct via the builder pattern


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


## GetAuthTokenAccessors

> GetAuthTokenAccessors(ctx).List(list).Execute()

List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthTokenAccessors(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthTokenAccessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenAccessorsRequest struct via the builder pattern


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


## GetAuthTokenLookup

> GetAuthTokenLookup(ctx).Execute()

This endpoint will lookup a token and its properties.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthTokenLookup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthTokenLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenLookupRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthTokenLookupSelf

> GetAuthTokenLookupSelf(ctx).Execute()

This endpoint will lookup a token and its properties.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthTokenLookupSelf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthTokenLookupSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenLookupSelfRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthTokenRoles

> GetAuthTokenRoles(ctx).List(list).Execute()

This endpoint lists configured roles.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthTokenRoles(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthTokenRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenRolesRequest struct via the builder pattern


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


## GetAuthTokenRolesRoleName

> GetAuthTokenRolesRoleName(ctx, roleName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthTokenRolesRoleName(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthTokenRolesRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenRolesRoleNameRequest struct via the builder pattern


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


## GetAuthUserpassUsers

> GetAuthUserpassUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

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
    resp, r, err := apiClient.Auth.GetAuthUserpassUsers(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthUserpassUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthUserpassUsersRequest struct via the builder pattern


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


## GetAuthUserpassUsersUsername

> GetAuthUserpassUsersUsername(ctx, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username for this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.GetAuthUserpassUsersUsername(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.GetAuthUserpassUsersUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username for this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthUserpassUsersUsernameRequest struct via the builder pattern


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


## PostAuthAlicloudLogin

> PostAuthAlicloudLogin(ctx).AlicloudLoginRequest(alicloudLoginRequest).Execute()

Authenticates an RAM entity with Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    alicloudLoginRequest := *openapiclient.NewAlicloudLoginRequest() // AlicloudLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAlicloudLogin(context.Background()).AlicloudLoginRequest(alicloudLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAlicloudLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAlicloudLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alicloudLoginRequest** | [**AlicloudLoginRequest**](AlicloudLoginRequest.md) |  | 

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


## PostAuthAlicloudRoleRole

> PostAuthAlicloudRoleRole(ctx, role).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role as it should appear in Vault.
    alicloudRoleRequest := *openapiclient.NewAlicloudRoleRequest() // AlicloudRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAlicloudRoleRole(context.Background(), role).AlicloudRoleRequest(alicloudRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAlicloudRoleRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAlicloudRoleRoleRequest struct via the builder pattern


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


## PostAuthAppIdLogin

> PostAuthAppIdLogin(ctx).AppIdLoginRequest(appIdLoginRequest).Execute()

Log in with an App ID and User ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appIdLoginRequest := *openapiclient.NewAppIdLoginRequest() // AppIdLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAppIdLogin(context.Background()).AppIdLoginRequest(appIdLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAppIdLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAppIdLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appIdLoginRequest** | [**AppIdLoginRequest**](AppIdLoginRequest.md) |  | 

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


## PostAuthAppIdLoginAppId

> PostAuthAppIdLoginAppId(ctx, appId).AppIdLoginRequest(appIdLoginRequest).Execute()

Log in with an App ID and User ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | The unique app ID
    appIdLoginRequest := *openapiclient.NewAppIdLoginRequest() // AppIdLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAppIdLoginAppId(context.Background(), appId).AppIdLoginRequest(appIdLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAppIdLoginAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The unique app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAppIdLoginAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdLoginRequest** | [**AppIdLoginRequest**](AppIdLoginRequest.md) |  | 

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


## PostAuthAppIdMapAppIdKey

> PostAuthAppIdMapAppIdKey(ctx, key).AppIdMapAppIdRequest(appIdMapAppIdRequest).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the app-id mapping
    appIdMapAppIdRequest := *openapiclient.NewAppIdMapAppIdRequest() // AppIdMapAppIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAppIdMapAppIdKey(context.Background(), key).AppIdMapAppIdRequest(appIdMapAppIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAppIdMapAppIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the app-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAppIdMapAppIdKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdMapAppIdRequest** | [**AppIdMapAppIdRequest**](AppIdMapAppIdRequest.md) |  | 

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


## PostAuthAppIdMapUserIdKey

> PostAuthAppIdMapUserIdKey(ctx, key).AppIdMapUserIdRequest(appIdMapUserIdRequest).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the user-id mapping
    appIdMapUserIdRequest := *openapiclient.NewAppIdMapUserIdRequest() // AppIdMapUserIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAppIdMapUserIdKey(context.Background(), key).AppIdMapUserIdRequest(appIdMapUserIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAppIdMapUserIdKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the user-id mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAppIdMapUserIdKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdMapUserIdRequest** | [**AppIdMapUserIdRequest**](AppIdMapUserIdRequest.md) |  | 

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


## PostAuthApproleLogin

> PostAuthApproleLogin(ctx).ApproleLoginRequest(approleLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    approleLoginRequest := *openapiclient.NewApproleLoginRequest() // ApproleLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleLogin(context.Background()).ApproleLoginRequest(approleLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approleLoginRequest** | [**ApproleLoginRequest**](ApproleLoginRequest.md) |  | 

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


## PostAuthApproleRoleRoleName

> PostAuthApproleRoleRoleName(ctx, roleName).ApproleRoleRequest(approleRoleRequest).Execute()

Register an role with the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleRequest := *openapiclient.NewApproleRoleRequest() // ApproleRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleName(context.Background(), roleName).ApproleRoleRequest(approleRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRequest** | [**ApproleRoleRequest**](ApproleRoleRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameBindSecretId

> PostAuthApproleRoleRoleNameBindSecretId(ctx, roleName).ApproleRoleBindSecretIdRequest(approleRoleBindSecretIdRequest).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleBindSecretIdRequest := *openapiclient.NewApproleRoleBindSecretIdRequest() // ApproleRoleBindSecretIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName).ApproleRoleBindSecretIdRequest(approleRoleBindSecretIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameBindSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameBindSecretIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBindSecretIdRequest** | [**ApproleRoleBindSecretIdRequest**](ApproleRoleBindSecretIdRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameBoundCidrList

> PostAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).ApproleRoleBoundCidrListRequest(approleRoleBoundCidrListRequest).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleBoundCidrListRequest := *openapiclient.NewApproleRoleBoundCidrListRequest() // ApproleRoleBoundCidrListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName).ApproleRoleBoundCidrListRequest(approleRoleBoundCidrListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameBoundCidrList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameBoundCidrListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBoundCidrListRequest** | [**ApproleRoleBoundCidrListRequest**](ApproleRoleBoundCidrListRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameCustomSecretId

> PostAuthApproleRoleRoleNameCustomSecretId(ctx, roleName).ApproleRoleCustomSecretIdRequest(approleRoleCustomSecretIdRequest).Execute()

Assign a SecretID of choice against the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleCustomSecretIdRequest := *openapiclient.NewApproleRoleCustomSecretIdRequest() // ApproleRoleCustomSecretIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameCustomSecretId(context.Background(), roleName).ApproleRoleCustomSecretIdRequest(approleRoleCustomSecretIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameCustomSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameCustomSecretIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleCustomSecretIdRequest** | [**ApproleRoleCustomSecretIdRequest**](ApproleRoleCustomSecretIdRequest.md) |  | 

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


## PostAuthApproleRoleRoleNamePeriod

> PostAuthApproleRoleRoleNamePeriod(ctx, roleName).ApproleRolePeriodRequest(approleRolePeriodRequest).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRolePeriodRequest := *openapiclient.NewApproleRolePeriodRequest() // ApproleRolePeriodRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNamePeriod(context.Background(), roleName).ApproleRolePeriodRequest(approleRolePeriodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNamePeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNamePeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePeriodRequest** | [**ApproleRolePeriodRequest**](ApproleRolePeriodRequest.md) |  | 

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


## PostAuthApproleRoleRoleNamePolicies

> PostAuthApproleRoleRoleNamePolicies(ctx, roleName).ApproleRolePoliciesRequest(approleRolePoliciesRequest).Execute()

Policies of the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRolePoliciesRequest := *openapiclient.NewApproleRolePoliciesRequest() // ApproleRolePoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNamePolicies(context.Background(), roleName).ApproleRolePoliciesRequest(approleRolePoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNamePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNamePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePoliciesRequest** | [**ApproleRolePoliciesRequest**](ApproleRolePoliciesRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameRoleId

> PostAuthApproleRoleRoleNameRoleId(ctx, roleName).ApproleRoleRoleIdRequest(approleRoleRoleIdRequest).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleRoleIdRequest := *openapiclient.NewApproleRoleRoleIdRequest() // ApproleRoleRoleIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameRoleId(context.Background(), roleName).ApproleRoleRoleIdRequest(approleRoleRoleIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameRoleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameRoleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRoleIdRequest** | [**ApproleRoleRoleIdRequest**](ApproleRoleRoleIdRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretId

> PostAuthApproleRoleRoleNameSecretId(ctx, roleName).ApproleRoleSecretIdRequest(approleRoleSecretIdRequest).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdRequest := *openapiclient.NewApproleRoleSecretIdRequest() // ApproleRoleSecretIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretId(context.Background(), roleName).ApproleRoleSecretIdRequest(approleRoleSecretIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdRequest** | [**ApproleRoleSecretIdRequest**](ApproleRoleSecretIdRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdAccessorDestroy

> PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, roleName).ApproleRoleSecretIdAccessorDestroyRequest(approleRoleSecretIdAccessorDestroyRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdAccessorDestroyRequest := *openapiclient.NewApproleRoleSecretIdAccessorDestroyRequest() // ApproleRoleSecretIdAccessorDestroyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), roleName).ApproleRoleSecretIdAccessorDestroyRequest(approleRoleSecretIdAccessorDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdAccessorDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdAccessorDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorDestroyRequest** | [**ApproleRoleSecretIdAccessorDestroyRequest**](ApproleRoleSecretIdAccessorDestroyRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdAccessorLookup

> PostAuthApproleRoleRoleNameSecretIdAccessorLookup(ctx, roleName).ApproleRoleSecretIdAccessorLookupRequest(approleRoleSecretIdAccessorLookupRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdAccessorLookupRequest := *openapiclient.NewApproleRoleSecretIdAccessorLookupRequest() // ApproleRoleSecretIdAccessorLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorLookup(context.Background(), roleName).ApproleRoleSecretIdAccessorLookupRequest(approleRoleSecretIdAccessorLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdAccessorLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdAccessorLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorLookupRequest** | [**ApproleRoleSecretIdAccessorLookupRequest**](ApproleRoleSecretIdAccessorLookupRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdBoundCidrs

> PostAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).ApproleRoleSecretIdBoundCidrsRequest(approleRoleSecretIdBoundCidrsRequest).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdBoundCidrsRequest := *openapiclient.NewApproleRoleSecretIdBoundCidrsRequest() // ApproleRoleSecretIdBoundCidrsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName).ApproleRoleSecretIdBoundCidrsRequest(approleRoleSecretIdBoundCidrsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdBoundCidrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdBoundCidrsRequest** | [**ApproleRoleSecretIdBoundCidrsRequest**](ApproleRoleSecretIdBoundCidrsRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdDestroy

> PostAuthApproleRoleRoleNameSecretIdDestroy(ctx, roleName).ApproleRoleSecretIdDestroyRequest(approleRoleSecretIdDestroyRequest).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdDestroyRequest := *openapiclient.NewApproleRoleSecretIdDestroyRequest() // ApproleRoleSecretIdDestroyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), roleName).ApproleRoleSecretIdDestroyRequest(approleRoleSecretIdDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdDestroyRequest** | [**ApproleRoleSecretIdDestroyRequest**](ApproleRoleSecretIdDestroyRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdLookup

> PostAuthApproleRoleRoleNameSecretIdLookup(ctx, roleName).ApproleRoleSecretIdLookupRequest(approleRoleSecretIdLookupRequest).Execute()

Read the properties of an issued secret_id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdLookupRequest := *openapiclient.NewApproleRoleSecretIdLookupRequest() // ApproleRoleSecretIdLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdLookup(context.Background(), roleName).ApproleRoleSecretIdLookupRequest(approleRoleSecretIdLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdLookupRequest** | [**ApproleRoleSecretIdLookupRequest**](ApproleRoleSecretIdLookupRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdNumUses

> PostAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).ApproleRoleSecretIdNumUsesRequest(approleRoleSecretIdNumUsesRequest).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdNumUsesRequest := *openapiclient.NewApproleRoleSecretIdNumUsesRequest() // ApproleRoleSecretIdNumUsesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName).ApproleRoleSecretIdNumUsesRequest(approleRoleSecretIdNumUsesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdNumUsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdNumUsesRequest** | [**ApproleRoleSecretIdNumUsesRequest**](ApproleRoleSecretIdNumUsesRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameSecretIdTtl

> PostAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).ApproleRoleSecretIdTtlRequest(approleRoleSecretIdTtlRequest).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleSecretIdTtlRequest := *openapiclient.NewApproleRoleSecretIdTtlRequest() // ApproleRoleSecretIdTtlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName).ApproleRoleSecretIdTtlRequest(approleRoleSecretIdTtlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameSecretIdTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameSecretIdTtlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdTtlRequest** | [**ApproleRoleSecretIdTtlRequest**](ApproleRoleSecretIdTtlRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameTokenBoundCidrs

> PostAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).ApproleRoleTokenBoundCidrsRequest(approleRoleTokenBoundCidrsRequest).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleTokenBoundCidrsRequest := *openapiclient.NewApproleRoleTokenBoundCidrsRequest() // ApproleRoleTokenBoundCidrsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName).ApproleRoleTokenBoundCidrsRequest(approleRoleTokenBoundCidrsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameTokenBoundCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameTokenBoundCidrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenBoundCidrsRequest** | [**ApproleRoleTokenBoundCidrsRequest**](ApproleRoleTokenBoundCidrsRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameTokenMaxTtl

> PostAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).ApproleRoleTokenMaxTtlRequest(approleRoleTokenMaxTtlRequest).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleTokenMaxTtlRequest := *openapiclient.NewApproleRoleTokenMaxTtlRequest() // ApproleRoleTokenMaxTtlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName).ApproleRoleTokenMaxTtlRequest(approleRoleTokenMaxTtlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameTokenMaxTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameTokenMaxTtlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenMaxTtlRequest** | [**ApproleRoleTokenMaxTtlRequest**](ApproleRoleTokenMaxTtlRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameTokenNumUses

> PostAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).ApproleRoleTokenNumUsesRequest(approleRoleTokenNumUsesRequest).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleTokenNumUsesRequest := *openapiclient.NewApproleRoleTokenNumUsesRequest() // ApproleRoleTokenNumUsesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName).ApproleRoleTokenNumUsesRequest(approleRoleTokenNumUsesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameTokenNumUses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameTokenNumUsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenNumUsesRequest** | [**ApproleRoleTokenNumUsesRequest**](ApproleRoleTokenNumUsesRequest.md) |  | 

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


## PostAuthApproleRoleRoleNameTokenTtl

> PostAuthApproleRoleRoleNameTokenTtl(ctx, roleName).ApproleRoleTokenTtlRequest(approleRoleTokenTtlRequest).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role.
    approleRoleTokenTtlRequest := *openapiclient.NewApproleRoleTokenTtlRequest() // ApproleRoleTokenTtlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName).ApproleRoleTokenTtlRequest(approleRoleTokenTtlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleRoleRoleNameTokenTtl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleRoleRoleNameTokenTtlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenTtlRequest** | [**ApproleRoleTokenTtlRequest**](ApproleRoleTokenTtlRequest.md) |  | 

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


## PostAuthApproleTidySecretId

> PostAuthApproleTidySecretId(ctx).Execute()

Trigger the clean-up of expired SecretID entries.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthApproleTidySecretId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthApproleTidySecretId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthApproleTidySecretIdRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthAwsConfigCertificateCertName

> PostAuthAwsConfigCertificateCertName(ctx, certName).AwsConfigCertificateRequest(awsConfigCertificateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certName := "certName_example" // string | Name of the certificate.
    awsConfigCertificateRequest := *openapiclient.NewAwsConfigCertificateRequest() // AwsConfigCertificateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigCertificateCertName(context.Background(), certName).AwsConfigCertificateRequest(awsConfigCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigCertificateCertName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certName** | **string** | Name of the certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigCertificateCertNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigCertificateRequest** | [**AwsConfigCertificateRequest**](AwsConfigCertificateRequest.md) |  | 

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


## PostAuthAwsConfigClient

> PostAuthAwsConfigClient(ctx).AwsConfigClientRequest(awsConfigClientRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigClientRequest := *openapiclient.NewAwsConfigClientRequest() // AwsConfigClientRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigClient(context.Background()).AwsConfigClientRequest(awsConfigClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigClientRequest** | [**AwsConfigClientRequest**](AwsConfigClientRequest.md) |  | 

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


## PostAuthAwsConfigIdentity

> PostAuthAwsConfigIdentity(ctx).AwsConfigIdentityRequest(awsConfigIdentityRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigIdentityRequest := *openapiclient.NewAwsConfigIdentityRequest() // AwsConfigIdentityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigIdentity(context.Background()).AwsConfigIdentityRequest(awsConfigIdentityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigIdentityRequest** | [**AwsConfigIdentityRequest**](AwsConfigIdentityRequest.md) |  | 

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


## PostAuthAwsConfigRotateRoot

> PostAuthAwsConfigRotateRoot(ctx).Execute()



### Example

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
    resp, r, err := apiClient.Auth.PostAuthAwsConfigRotateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigRotateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigRotateRootRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthAwsConfigStsAccountId

> PostAuthAwsConfigStsAccountId(ctx, accountId).AwsConfigStsRequest(awsConfigStsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
    awsConfigStsRequest := *openapiclient.NewAwsConfigStsRequest() // AwsConfigStsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigStsAccountId(context.Background(), accountId).AwsConfigStsRequest(awsConfigStsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigStsAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigStsAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigStsRequest** | [**AwsConfigStsRequest**](AwsConfigStsRequest.md) |  | 

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


## PostAuthAwsConfigTidyIdentityAccesslist

> PostAuthAwsConfigTidyIdentityAccesslist(ctx).AwsConfigTidyIdentityAccesslistRequest(awsConfigTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigTidyIdentityAccesslistRequest := *openapiclient.NewAwsConfigTidyIdentityAccesslistRequest() // AwsConfigTidyIdentityAccesslistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigTidyIdentityAccesslist(context.Background()).AwsConfigTidyIdentityAccesslistRequest(awsConfigTidyIdentityAccesslistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigTidyIdentityAccesslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigTidyIdentityAccesslistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigTidyIdentityAccesslistRequest** | [**AwsConfigTidyIdentityAccesslistRequest**](AwsConfigTidyIdentityAccesslistRequest.md) |  | 

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


## PostAuthAwsConfigTidyIdentityWhitelist

> PostAuthAwsConfigTidyIdentityWhitelist(ctx).AwsConfigTidyIdentityWhitelistRequest(awsConfigTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigTidyIdentityWhitelistRequest := *openapiclient.NewAwsConfigTidyIdentityWhitelistRequest() // AwsConfigTidyIdentityWhitelistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigTidyIdentityWhitelist(context.Background()).AwsConfigTidyIdentityWhitelistRequest(awsConfigTidyIdentityWhitelistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigTidyIdentityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigTidyIdentityWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigTidyIdentityWhitelistRequest** | [**AwsConfigTidyIdentityWhitelistRequest**](AwsConfigTidyIdentityWhitelistRequest.md) |  | 

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


## PostAuthAwsConfigTidyRoletagBlacklist

> PostAuthAwsConfigTidyRoletagBlacklist(ctx).AwsConfigTidyRoletagBlacklistRequest(awsConfigTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigTidyRoletagBlacklistRequest := *openapiclient.NewAwsConfigTidyRoletagBlacklistRequest() // AwsConfigTidyRoletagBlacklistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigTidyRoletagBlacklist(context.Background()).AwsConfigTidyRoletagBlacklistRequest(awsConfigTidyRoletagBlacklistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigTidyRoletagBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigTidyRoletagBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigTidyRoletagBlacklistRequest** | [**AwsConfigTidyRoletagBlacklistRequest**](AwsConfigTidyRoletagBlacklistRequest.md) |  | 

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


## PostAuthAwsConfigTidyRoletagDenylist

> PostAuthAwsConfigTidyRoletagDenylist(ctx).AwsConfigTidyRoletagDenylistRequest(awsConfigTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsConfigTidyRoletagDenylistRequest := *openapiclient.NewAwsConfigTidyRoletagDenylistRequest() // AwsConfigTidyRoletagDenylistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsConfigTidyRoletagDenylist(context.Background()).AwsConfigTidyRoletagDenylistRequest(awsConfigTidyRoletagDenylistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsConfigTidyRoletagDenylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsConfigTidyRoletagDenylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsConfigTidyRoletagDenylistRequest** | [**AwsConfigTidyRoletagDenylistRequest**](AwsConfigTidyRoletagDenylistRequest.md) |  | 

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


## PostAuthAwsLogin

> PostAuthAwsLogin(ctx).AwsLoginRequest(awsLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsLoginRequest := *openapiclient.NewAwsLoginRequest() // AwsLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsLogin(context.Background()).AwsLoginRequest(awsLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsLoginRequest** | [**AwsLoginRequest**](AwsLoginRequest.md) |  | 

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


## PostAuthAwsRoleRole

> PostAuthAwsRoleRole(ctx, role).AwsRoleRequest(awsRoleRequest).Execute()



### Example

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
    awsRoleRequest := *openapiclient.NewAwsRoleRequest() // AwsRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsRoleRole(context.Background(), role).AwsRoleRequest(awsRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthAwsRoleRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRoleRequest** | [**AwsRoleRequest**](AwsRoleRequest.md) |  | 

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


## PostAuthAwsRoleRoleTag

> PostAuthAwsRoleRoleTag(ctx, role).AwsRoleTagRequest(awsRoleTagRequest).Execute()



### Example

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
    awsRoleTagRequest := *openapiclient.NewAwsRoleTagRequest() // AwsRoleTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsRoleRoleTag(context.Background(), role).AwsRoleTagRequest(awsRoleTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsRoleRoleTag``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthAwsRoleRoleTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsRoleTagRequest** | [**AwsRoleTagRequest**](AwsRoleTagRequest.md) |  | 

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


## PostAuthAwsRoletagBlacklistRoleTag

> PostAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsRoletagBlacklistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsRoletagBlacklistRoleTagRequest struct via the builder pattern


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


## PostAuthAwsRoletagDenylistRoleTag

> PostAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsRoletagDenylistRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsRoletagDenylistRoleTagRequest struct via the builder pattern


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


## PostAuthAwsTidyIdentityAccesslist

> PostAuthAwsTidyIdentityAccesslist(ctx).AwsTidyIdentityAccesslistRequest(awsTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsTidyIdentityAccesslistRequest := *openapiclient.NewAwsTidyIdentityAccesslistRequest() // AwsTidyIdentityAccesslistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsTidyIdentityAccesslist(context.Background()).AwsTidyIdentityAccesslistRequest(awsTidyIdentityAccesslistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsTidyIdentityAccesslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsTidyIdentityAccesslistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsTidyIdentityAccesslistRequest** | [**AwsTidyIdentityAccesslistRequest**](AwsTidyIdentityAccesslistRequest.md) |  | 

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


## PostAuthAwsTidyIdentityWhitelist

> PostAuthAwsTidyIdentityWhitelist(ctx).AwsTidyIdentityWhitelistRequest(awsTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsTidyIdentityWhitelistRequest := *openapiclient.NewAwsTidyIdentityWhitelistRequest() // AwsTidyIdentityWhitelistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsTidyIdentityWhitelist(context.Background()).AwsTidyIdentityWhitelistRequest(awsTidyIdentityWhitelistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsTidyIdentityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsTidyIdentityWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsTidyIdentityWhitelistRequest** | [**AwsTidyIdentityWhitelistRequest**](AwsTidyIdentityWhitelistRequest.md) |  | 

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


## PostAuthAwsTidyRoletagBlacklist

> PostAuthAwsTidyRoletagBlacklist(ctx).AwsTidyRoletagBlacklistRequest(awsTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsTidyRoletagBlacklistRequest := *openapiclient.NewAwsTidyRoletagBlacklistRequest() // AwsTidyRoletagBlacklistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsTidyRoletagBlacklist(context.Background()).AwsTidyRoletagBlacklistRequest(awsTidyRoletagBlacklistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsTidyRoletagBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsTidyRoletagBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsTidyRoletagBlacklistRequest** | [**AwsTidyRoletagBlacklistRequest**](AwsTidyRoletagBlacklistRequest.md) |  | 

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


## PostAuthAwsTidyRoletagDenylist

> PostAuthAwsTidyRoletagDenylist(ctx).AwsTidyRoletagDenylistRequest(awsTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsTidyRoletagDenylistRequest := *openapiclient.NewAwsTidyRoletagDenylistRequest() // AwsTidyRoletagDenylistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAwsTidyRoletagDenylist(context.Background()).AwsTidyRoletagDenylistRequest(awsTidyRoletagDenylistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAwsTidyRoletagDenylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAwsTidyRoletagDenylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsTidyRoletagDenylistRequest** | [**AwsTidyRoletagDenylistRequest**](AwsTidyRoletagDenylistRequest.md) |  | 

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


## PostAuthAzureConfig

> PostAuthAzureConfig(ctx).AzureConfigRequest(azureConfigRequest).Execute()



### Example

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
    resp, r, err := apiClient.Auth.PostAuthAzureConfig(context.Background()).AzureConfigRequest(azureConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAzureConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAzureConfigRequest struct via the builder pattern


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


## PostAuthAzureLogin

> PostAuthAzureLogin(ctx).AzureLoginRequest(azureLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    azureLoginRequest := *openapiclient.NewAzureLoginRequest() // AzureLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAzureLogin(context.Background()).AzureLoginRequest(azureLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAzureLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthAzureLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureLoginRequest** | [**AzureLoginRequest**](AzureLoginRequest.md) |  | 

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


## PostAuthAzureRoleName

> PostAuthAzureRoleName(ctx, name).AzureRoleRequest(azureRoleRequest).Execute()



### Example

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
    azureRoleRequest := *openapiclient.NewAzureRoleRequest() // AzureRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthAzureRoleName(context.Background(), name).AzureRoleRequest(azureRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthAzureRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthAzureRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureRoleRequest** | [**AzureRoleRequest**](AzureRoleRequest.md) |  | 

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


## PostAuthCentrifyConfig

> PostAuthCentrifyConfig(ctx).CentrifyConfigRequest(centrifyConfigRequest).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    centrifyConfigRequest := *openapiclient.NewCentrifyConfigRequest() // CentrifyConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCentrifyConfig(context.Background()).CentrifyConfigRequest(centrifyConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCentrifyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCentrifyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **centrifyConfigRequest** | [**CentrifyConfigRequest**](CentrifyConfigRequest.md) |  | 

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


## PostAuthCentrifyLogin

> PostAuthCentrifyLogin(ctx).CentrifyLoginRequest(centrifyLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    centrifyLoginRequest := *openapiclient.NewCentrifyLoginRequest() // CentrifyLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCentrifyLogin(context.Background()).CentrifyLoginRequest(centrifyLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCentrifyLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCentrifyLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **centrifyLoginRequest** | [**CentrifyLoginRequest**](CentrifyLoginRequest.md) |  | 

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


## PostAuthCertCertsName

> PostAuthCertCertsName(ctx, name).CertCertsRequest(certCertsRequest).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate
    certCertsRequest := *openapiclient.NewCertCertsRequest() // CertCertsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCertCertsName(context.Background(), name).CertCertsRequest(certCertsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCertCertsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCertCertsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCertsRequest** | [**CertCertsRequest**](CertCertsRequest.md) |  | 

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


## PostAuthCertConfig

> PostAuthCertConfig(ctx).CertConfigRequest(certConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certConfigRequest := *openapiclient.NewCertConfigRequest() // CertConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCertConfig(context.Background()).CertConfigRequest(certConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certConfigRequest** | [**CertConfigRequest**](CertConfigRequest.md) |  | 

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


## PostAuthCertCrlsName

> PostAuthCertCrlsName(ctx, name).CertCrlsRequest(certCrlsRequest).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the certificate
    certCrlsRequest := *openapiclient.NewCertCrlsRequest() // CertCrlsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCertCrlsName(context.Background(), name).CertCrlsRequest(certCrlsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCertCrlsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCertCrlsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCrlsRequest** | [**CertCrlsRequest**](CertCrlsRequest.md) |  | 

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


## PostAuthCertLogin

> PostAuthCertLogin(ctx).CertLoginRequest(certLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    certLoginRequest := *openapiclient.NewCertLoginRequest() // CertLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCertLogin(context.Background()).CertLoginRequest(certLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCertLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCertLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certLoginRequest** | [**CertLoginRequest**](CertLoginRequest.md) |  | 

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


## PostAuthCfConfig

> PostAuthCfConfig(ctx).CfConfigRequest(cfConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cfConfigRequest := *openapiclient.NewCfConfigRequest() // CfConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCfConfig(context.Background()).CfConfigRequest(cfConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCfConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCfConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cfConfigRequest** | [**CfConfigRequest**](CfConfigRequest.md) |  | 

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


## PostAuthCfLogin

> PostAuthCfLogin(ctx).CfLoginRequest(cfLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cfLoginRequest := *openapiclient.NewCfLoginRequest("CfInstanceCert_example", "Role_example", "Signature_example", "SigningTime_example") // CfLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCfLogin(context.Background()).CfLoginRequest(cfLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCfLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCfLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cfLoginRequest** | [**CfLoginRequest**](CfLoginRequest.md) |  | 

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


## PostAuthCfRolesRole

> PostAuthCfRolesRole(ctx, role).CfRolesRequest(cfRolesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    role := "role_example" // string | The name of the role.
    cfRolesRequest := *openapiclient.NewCfRolesRequest() // CfRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthCfRolesRole(context.Background(), role).CfRolesRequest(cfRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthCfRolesRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**role** | **string** | The name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthCfRolesRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cfRolesRequest** | [**CfRolesRequest**](CfRolesRequest.md) |  | 

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


## PostAuthGcpConfig

> PostAuthGcpConfig(ctx).GcpConfigRequest(gcpConfigRequest).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

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
    resp, r, err := apiClient.Auth.PostAuthGcpConfig(context.Background()).GcpConfigRequest(gcpConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGcpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGcpConfigRequest struct via the builder pattern


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


## PostAuthGcpLogin

> PostAuthGcpLogin(ctx).GcpLoginRequest(gcpLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gcpLoginRequest := *openapiclient.NewGcpLoginRequest() // GcpLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGcpLogin(context.Background()).GcpLoginRequest(gcpLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGcpLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGcpLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpLoginRequest** | [**GcpLoginRequest**](GcpLoginRequest.md) |  | 

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


## PostAuthGcpRoleName

> PostAuthGcpRoleName(ctx, name).GcpRoleRequest(gcpRoleRequest).Execute()

Create a GCP role with associated policies and required attributes.

### Example

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
    gcpRoleRequest := *openapiclient.NewGcpRoleRequest() // GcpRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGcpRoleName(context.Background(), name).GcpRoleRequest(gcpRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGcpRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthGcpRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRoleRequest** | [**GcpRoleRequest**](GcpRoleRequest.md) |  | 

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


## PostAuthGcpRoleNameLabels

> PostAuthGcpRoleNameLabels(ctx, name).GcpRoleLabelsRequest(gcpRoleLabelsRequest).Execute()

Add or remove labels for an existing 'gce' role

### Example

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
    gcpRoleLabelsRequest := *openapiclient.NewGcpRoleLabelsRequest() // GcpRoleLabelsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGcpRoleNameLabels(context.Background(), name).GcpRoleLabelsRequest(gcpRoleLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGcpRoleNameLabels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthGcpRoleNameLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRoleLabelsRequest** | [**GcpRoleLabelsRequest**](GcpRoleLabelsRequest.md) |  | 

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


## PostAuthGcpRoleNameServiceAccounts

> PostAuthGcpRoleNameServiceAccounts(ctx, name).GcpRoleServiceAccountsRequest(gcpRoleServiceAccountsRequest).Execute()

Add or remove service accounts for an existing `iam` role

### Example

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
    gcpRoleServiceAccountsRequest := *openapiclient.NewGcpRoleServiceAccountsRequest() // GcpRoleServiceAccountsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGcpRoleNameServiceAccounts(context.Background(), name).GcpRoleServiceAccountsRequest(gcpRoleServiceAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGcpRoleNameServiceAccounts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthGcpRoleNameServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpRoleServiceAccountsRequest** | [**GcpRoleServiceAccountsRequest**](GcpRoleServiceAccountsRequest.md) |  | 

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


## PostAuthGithubConfig

> PostAuthGithubConfig(ctx).GithubConfigRequest(githubConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubConfigRequest := *openapiclient.NewGithubConfigRequest("Organization_example") // GithubConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGithubConfig(context.Background()).GithubConfigRequest(githubConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGithubConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGithubConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubConfigRequest** | [**GithubConfigRequest**](GithubConfigRequest.md) |  | 

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


## PostAuthGithubLogin

> PostAuthGithubLogin(ctx).GithubLoginRequest(githubLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubLoginRequest := *openapiclient.NewGithubLoginRequest() // GithubLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGithubLogin(context.Background()).GithubLoginRequest(githubLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGithubLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGithubLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubLoginRequest** | [**GithubLoginRequest**](GithubLoginRequest.md) |  | 

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


## PostAuthGithubMapTeamsKey

> PostAuthGithubMapTeamsKey(ctx, key).GithubMapTeamsRequest(githubMapTeamsRequest).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the teams mapping
    githubMapTeamsRequest := *openapiclient.NewGithubMapTeamsRequest() // GithubMapTeamsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGithubMapTeamsKey(context.Background(), key).GithubMapTeamsRequest(githubMapTeamsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGithubMapTeamsKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the teams mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGithubMapTeamsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapTeamsRequest** | [**GithubMapTeamsRequest**](GithubMapTeamsRequest.md) |  | 

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


## PostAuthGithubMapUsersKey

> PostAuthGithubMapUsersKey(ctx, key).GithubMapUsersRequest(githubMapUsersRequest).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key for the users mapping
    githubMapUsersRequest := *openapiclient.NewGithubMapUsersRequest() // GithubMapUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthGithubMapUsersKey(context.Background(), key).GithubMapUsersRequest(githubMapUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthGithubMapUsersKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key for the users mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthGithubMapUsersKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapUsersRequest** | [**GithubMapUsersRequest**](GithubMapUsersRequest.md) |  | 

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


## PostAuthJwtConfig

> PostAuthJwtConfig(ctx).JwtConfigRequest(jwtConfigRequest).Execute()

Configure the JWT authentication backend.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jwtConfigRequest := *openapiclient.NewJwtConfigRequest() // JwtConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthJwtConfig(context.Background()).JwtConfigRequest(jwtConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthJwtConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthJwtConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtConfigRequest** | [**JwtConfigRequest**](JwtConfigRequest.md) |  | 

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


## PostAuthJwtLogin

> PostAuthJwtLogin(ctx).JwtLoginRequest(jwtLoginRequest).Execute()

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jwtLoginRequest := *openapiclient.NewJwtLoginRequest() // JwtLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthJwtLogin(context.Background()).JwtLoginRequest(jwtLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthJwtLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthJwtLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtLoginRequest** | [**JwtLoginRequest**](JwtLoginRequest.md) |  | 

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


## PostAuthJwtOidcAuthUrl

> PostAuthJwtOidcAuthUrl(ctx).JwtOidcAuthUrlRequest(jwtOidcAuthUrlRequest).Execute()

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jwtOidcAuthUrlRequest := *openapiclient.NewJwtOidcAuthUrlRequest() // JwtOidcAuthUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthJwtOidcAuthUrl(context.Background()).JwtOidcAuthUrlRequest(jwtOidcAuthUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthJwtOidcAuthUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthJwtOidcAuthUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtOidcAuthUrlRequest** | [**JwtOidcAuthUrlRequest**](JwtOidcAuthUrlRequest.md) |  | 

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


## PostAuthJwtOidcCallback

> PostAuthJwtOidcCallback(ctx).JwtOidcCallbackRequest(jwtOidcCallbackRequest).Execute()

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jwtOidcCallbackRequest := *openapiclient.NewJwtOidcCallbackRequest() // JwtOidcCallbackRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthJwtOidcCallback(context.Background()).JwtOidcCallbackRequest(jwtOidcCallbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthJwtOidcCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthJwtOidcCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jwtOidcCallbackRequest** | [**JwtOidcCallbackRequest**](JwtOidcCallbackRequest.md) |  | 

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


## PostAuthJwtRoleName

> PostAuthJwtRoleName(ctx, name).JwtRoleRequest(jwtRoleRequest).Execute()

Register an role with the backend.



### Example

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
    jwtRoleRequest := *openapiclient.NewJwtRoleRequest() // JwtRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthJwtRoleName(context.Background(), name).JwtRoleRequest(jwtRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthJwtRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthJwtRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtRoleRequest** | [**JwtRoleRequest**](JwtRoleRequest.md) |  | 

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


## PostAuthKerberosConfig

> PostAuthKerberosConfig(ctx).KerberosConfigRequest(kerberosConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kerberosConfigRequest := *openapiclient.NewKerberosConfigRequest() // KerberosConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKerberosConfig(context.Background()).KerberosConfigRequest(kerberosConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKerberosConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKerberosConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosConfigRequest** | [**KerberosConfigRequest**](KerberosConfigRequest.md) |  | 

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


## PostAuthKerberosConfigLdap

> PostAuthKerberosConfigLdap(ctx).KerberosConfigLdapRequest(kerberosConfigLdapRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kerberosConfigLdapRequest := *openapiclient.NewKerberosConfigLdapRequest() // KerberosConfigLdapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKerberosConfigLdap(context.Background()).KerberosConfigLdapRequest(kerberosConfigLdapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKerberosConfigLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKerberosConfigLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosConfigLdapRequest** | [**KerberosConfigLdapRequest**](KerberosConfigLdapRequest.md) |  | 

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


## PostAuthKerberosGroupsName

> PostAuthKerberosGroupsName(ctx, name).KerberosGroupsRequest(kerberosGroupsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.
    kerberosGroupsRequest := *openapiclient.NewKerberosGroupsRequest() // KerberosGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKerberosGroupsName(context.Background(), name).KerberosGroupsRequest(kerberosGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKerberosGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKerberosGroupsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosGroupsRequest** | [**KerberosGroupsRequest**](KerberosGroupsRequest.md) |  | 

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


## PostAuthKerberosLogin

> PostAuthKerberosLogin(ctx).KerberosLoginRequest(kerberosLoginRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kerberosLoginRequest := *openapiclient.NewKerberosLoginRequest() // KerberosLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKerberosLogin(context.Background()).KerberosLoginRequest(kerberosLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKerberosLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKerberosLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosLoginRequest** | [**KerberosLoginRequest**](KerberosLoginRequest.md) |  | 

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


## PostAuthKubernetesConfig

> PostAuthKubernetesConfig(ctx).KubernetesConfigRequest(kubernetesConfigRequest).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthKubernetesConfig(context.Background()).KubernetesConfigRequest(kubernetesConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKubernetesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKubernetesConfigRequest struct via the builder pattern


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


## PostAuthKubernetesLogin

> PostAuthKubernetesLogin(ctx).KubernetesLoginRequest(kubernetesLoginRequest).Execute()

Authenticates Kubernetes service accounts with Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kubernetesLoginRequest := *openapiclient.NewKubernetesLoginRequest() // KubernetesLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKubernetesLogin(context.Background()).KubernetesLoginRequest(kubernetesLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKubernetesLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKubernetesLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesLoginRequest** | [**KubernetesLoginRequest**](KubernetesLoginRequest.md) |  | 

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


## PostAuthKubernetesRoleName

> PostAuthKubernetesRoleName(ctx, name).KubernetesRoleRequest(kubernetesRoleRequest).Execute()

Register an role with the backend.

### Example

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
    kubernetesRoleRequest := *openapiclient.NewKubernetesRoleRequest() // KubernetesRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthKubernetesRoleName(context.Background(), name).KubernetesRoleRequest(kubernetesRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthKubernetesRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthKubernetesRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesRoleRequest** | [**KubernetesRoleRequest**](KubernetesRoleRequest.md) |  | 

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


## PostAuthLdapConfig

> PostAuthLdapConfig(ctx).LdapConfigRequest(ldapConfigRequest).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ldapConfigRequest := *openapiclient.NewLdapConfigRequest() // LdapConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthLdapConfig(context.Background()).LdapConfigRequest(ldapConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthLdapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthLdapConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigRequest** | [**LdapConfigRequest**](LdapConfigRequest.md) |  | 

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


## PostAuthLdapGroupsName

> PostAuthLdapGroupsName(ctx, name).LdapGroupsRequest(ldapGroupsRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP group.
    ldapGroupsRequest := *openapiclient.NewLdapGroupsRequest() // LdapGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthLdapGroupsName(context.Background(), name).LdapGroupsRequest(ldapGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthLdapGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthLdapGroupsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapGroupsRequest** | [**LdapGroupsRequest**](LdapGroupsRequest.md) |  | 

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


## PostAuthLdapLoginUsername

> PostAuthLdapLoginUsername(ctx, username).LdapLoginRequest(ldapLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | DN (distinguished name) to be used for login.
    ldapLoginRequest := *openapiclient.NewLdapLoginRequest() // LdapLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthLdapLoginUsername(context.Background(), username).LdapLoginRequest(ldapLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthLdapLoginUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | DN (distinguished name) to be used for login. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthLdapLoginUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLoginRequest** | [**LdapLoginRequest**](LdapLoginRequest.md) |  | 

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


## PostAuthLdapUsersName

> PostAuthLdapUsersName(ctx, name).LdapUsersRequest(ldapUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the LDAP user.
    ldapUsersRequest := *openapiclient.NewLdapUsersRequest() // LdapUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthLdapUsersName(context.Background(), name).LdapUsersRequest(ldapUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthLdapUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthLdapUsersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapUsersRequest** | [**LdapUsersRequest**](LdapUsersRequest.md) |  | 

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


## PostAuthOciConfig

> PostAuthOciConfig(ctx).OciConfigRequest(ociConfigRequest).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ociConfigRequest := *openapiclient.NewOciConfigRequest() // OciConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOciConfig(context.Background()).OciConfigRequest(ociConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOciConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOciConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ociConfigRequest** | [**OciConfigRequest**](OciConfigRequest.md) |  | 

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


## PostAuthOciLoginRole

> PostAuthOciLoginRole(ctx, role).OciLoginRequest(ociLoginRequest).Execute()

Authenticates to Vault using OCI credentials

### Example

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
    ociLoginRequest := *openapiclient.NewOciLoginRequest() // OciLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOciLoginRole(context.Background(), role).OciLoginRequest(ociLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOciLoginRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthOciLoginRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ociLoginRequest** | [**OciLoginRequest**](OciLoginRequest.md) |  | 

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


## PostAuthOciRoleRole

> PostAuthOciRoleRole(ctx, role).OciRoleRequest(ociRoleRequest).Execute()

Create a role and associate policies to it.

### Example

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
    ociRoleRequest := *openapiclient.NewOciRoleRequest() // OciRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOciRoleRole(context.Background(), role).OciRoleRequest(ociRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOciRoleRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthOciRoleRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ociRoleRequest** | [**OciRoleRequest**](OciRoleRequest.md) |  | 

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


## PostAuthOidcConfig

> PostAuthOidcConfig(ctx).OidcConfigRequest(oidcConfigRequest).Execute()

Configure the JWT authentication backend.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oidcConfigRequest := *openapiclient.NewOidcConfigRequest() // OidcConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOidcConfig(context.Background()).OidcConfigRequest(oidcConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOidcConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcConfigRequest** | [**OidcConfigRequest**](OidcConfigRequest.md) |  | 

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


## PostAuthOidcLogin

> PostAuthOidcLogin(ctx).OidcLoginRequest(oidcLoginRequest).Execute()

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oidcLoginRequest := *openapiclient.NewOidcLoginRequest() // OidcLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOidcLogin(context.Background()).OidcLoginRequest(oidcLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOidcLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOidcLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcLoginRequest** | [**OidcLoginRequest**](OidcLoginRequest.md) |  | 

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


## PostAuthOidcOidcAuthUrl

> PostAuthOidcOidcAuthUrl(ctx).OidcOidcAuthUrlRequest(oidcOidcAuthUrlRequest).Execute()

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oidcOidcAuthUrlRequest := *openapiclient.NewOidcOidcAuthUrlRequest() // OidcOidcAuthUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOidcOidcAuthUrl(context.Background()).OidcOidcAuthUrlRequest(oidcOidcAuthUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOidcOidcAuthUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOidcOidcAuthUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcOidcAuthUrlRequest** | [**OidcOidcAuthUrlRequest**](OidcOidcAuthUrlRequest.md) |  | 

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


## PostAuthOidcOidcCallback

> PostAuthOidcOidcCallback(ctx).OidcOidcCallbackRequest(oidcOidcCallbackRequest).Execute()

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oidcOidcCallbackRequest := *openapiclient.NewOidcOidcCallbackRequest() // OidcOidcCallbackRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOidcOidcCallback(context.Background()).OidcOidcCallbackRequest(oidcOidcCallbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOidcOidcCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOidcOidcCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcOidcCallbackRequest** | [**OidcOidcCallbackRequest**](OidcOidcCallbackRequest.md) |  | 

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


## PostAuthOidcRoleName

> PostAuthOidcRoleName(ctx, name).OidcRoleRequest(oidcRoleRequest).Execute()

Register an role with the backend.



### Example

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
    oidcRoleRequest := *openapiclient.NewOidcRoleRequest() // OidcRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOidcRoleName(context.Background(), name).OidcRoleRequest(oidcRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOidcRoleName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAuthOidcRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcRoleRequest** | [**OidcRoleRequest**](OidcRoleRequest.md) |  | 

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


## PostAuthOktaConfig

> PostAuthOktaConfig(ctx).OktaConfigRequest(oktaConfigRequest).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oktaConfigRequest := *openapiclient.NewOktaConfigRequest() // OktaConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOktaConfig(context.Background()).OktaConfigRequest(oktaConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOktaConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOktaConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oktaConfigRequest** | [**OktaConfigRequest**](OktaConfigRequest.md) |  | 

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


## PostAuthOktaGroupsName

> PostAuthOktaGroupsName(ctx, name).OktaGroupsRequest(oktaGroupsRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the Okta group.
    oktaGroupsRequest := *openapiclient.NewOktaGroupsRequest() // OktaGroupsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOktaGroupsName(context.Background(), name).OktaGroupsRequest(oktaGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOktaGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Okta group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOktaGroupsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaGroupsRequest** | [**OktaGroupsRequest**](OktaGroupsRequest.md) |  | 

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


## PostAuthOktaLoginUsername

> PostAuthOktaLoginUsername(ctx, username).OktaLoginRequest(oktaLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username to be used for login.
    oktaLoginRequest := *openapiclient.NewOktaLoginRequest() // OktaLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOktaLoginUsername(context.Background(), username).OktaLoginRequest(oktaLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOktaLoginUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username to be used for login. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOktaLoginUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaLoginRequest** | [**OktaLoginRequest**](OktaLoginRequest.md) |  | 

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


## PostAuthOktaUsersName

> PostAuthOktaUsersName(ctx, name).OktaUsersRequest(oktaUsersRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the user.
    oktaUsersRequest := *openapiclient.NewOktaUsersRequest() // OktaUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthOktaUsersName(context.Background(), name).OktaUsersRequest(oktaUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthOktaUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthOktaUsersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaUsersRequest** | [**OktaUsersRequest**](OktaUsersRequest.md) |  | 

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


## PostAuthRadiusConfig

> PostAuthRadiusConfig(ctx).RadiusConfigRequest(radiusConfigRequest).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    radiusConfigRequest := *openapiclient.NewRadiusConfigRequest() // RadiusConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthRadiusConfig(context.Background()).RadiusConfigRequest(radiusConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthRadiusConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthRadiusConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusConfigRequest** | [**RadiusConfigRequest**](RadiusConfigRequest.md) |  | 

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


## PostAuthRadiusLogin

> PostAuthRadiusLogin(ctx).RadiusLoginRequest(radiusLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    radiusLoginRequest := *openapiclient.NewRadiusLoginRequest() // RadiusLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthRadiusLogin(context.Background()).RadiusLoginRequest(radiusLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthRadiusLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthRadiusLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

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


## PostAuthRadiusLoginUrlusername

> PostAuthRadiusLoginUrlusername(ctx, urlusername).RadiusLoginRequest(radiusLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlusername := "urlusername_example" // string | Username to be used for login. (URL parameter)
    radiusLoginRequest := *openapiclient.NewRadiusLoginRequest() // RadiusLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthRadiusLoginUrlusername(context.Background(), urlusername).RadiusLoginRequest(radiusLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthRadiusLoginUrlusername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlusername** | **string** | Username to be used for login. (URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthRadiusLoginUrlusernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

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


## PostAuthRadiusUsersName

> PostAuthRadiusUsersName(ctx, name).RadiusUsersRequest(radiusUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the RADIUS user.
    radiusUsersRequest := *openapiclient.NewRadiusUsersRequest() // RadiusUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthRadiusUsersName(context.Background(), name).RadiusUsersRequest(radiusUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthRadiusUsersName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthRadiusUsersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusUsersRequest** | [**RadiusUsersRequest**](RadiusUsersRequest.md) |  | 

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


## PostAuthTokenCreate

> PostAuthTokenCreate(ctx).Execute()

The token create path is used to create new tokens.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthTokenCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthTokenCreateOrphan

> PostAuthTokenCreateOrphan(ctx).Execute()

The token create path is used to create new orphan tokens.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthTokenCreateOrphan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenCreateOrphan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenCreateOrphanRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthTokenCreateRoleName

> PostAuthTokenCreateRoleName(ctx, roleName).Execute()

This token create path is used to create new tokens adhering to the given role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenCreateRoleName(context.Background(), roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenCreateRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenCreateRoleNameRequest struct via the builder pattern


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


## PostAuthTokenLookup

> PostAuthTokenLookup(ctx).TokenLookupRequest(tokenLookupRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenLookupRequest := *openapiclient.NewTokenLookupRequest() // TokenLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenLookup(context.Background()).TokenLookupRequest(tokenLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupRequest** | [**TokenLookupRequest**](TokenLookupRequest.md) |  | 

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


## PostAuthTokenLookupAccessor

> PostAuthTokenLookupAccessor(ctx).TokenLookupAccessorRequest(tokenLookupAccessorRequest).Execute()

This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenLookupAccessorRequest := *openapiclient.NewTokenLookupAccessorRequest() // TokenLookupAccessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenLookupAccessor(context.Background()).TokenLookupAccessorRequest(tokenLookupAccessorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenLookupAccessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenLookupAccessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupAccessorRequest** | [**TokenLookupAccessorRequest**](TokenLookupAccessorRequest.md) |  | 

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


## PostAuthTokenLookupSelf

> PostAuthTokenLookupSelf(ctx).TokenLookupSelfRequest(tokenLookupSelfRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenLookupSelfRequest := *openapiclient.NewTokenLookupSelfRequest() // TokenLookupSelfRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenLookupSelf(context.Background()).TokenLookupSelfRequest(tokenLookupSelfRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenLookupSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenLookupSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookupSelfRequest** | [**TokenLookupSelfRequest**](TokenLookupSelfRequest.md) |  | 

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


## PostAuthTokenRenew

> PostAuthTokenRenew(ctx).TokenRenewRequest(tokenRenewRequest).Execute()

This endpoint will renew the given token and prevent expiration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRenewRequest := *openapiclient.NewTokenRenewRequest() // TokenRenewRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRenew(context.Background()).TokenRenewRequest(tokenRenewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewRequest** | [**TokenRenewRequest**](TokenRenewRequest.md) |  | 

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


## PostAuthTokenRenewAccessor

> PostAuthTokenRenewAccessor(ctx).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()

This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRenewAccessorRequest := *openapiclient.NewTokenRenewAccessorRequest() // TokenRenewAccessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRenewAccessor(context.Background()).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRenewAccessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRenewAccessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewAccessorRequest** | [**TokenRenewAccessorRequest**](TokenRenewAccessorRequest.md) |  | 

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


## PostAuthTokenRenewSelf

> PostAuthTokenRenewSelf(ctx).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()

This endpoint will renew the token used to call it and prevent expiration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRenewSelfRequest := *openapiclient.NewTokenRenewSelfRequest() // TokenRenewSelfRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRenewSelf(context.Background()).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRenewSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRenewSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewSelfRequest** | [**TokenRenewSelfRequest**](TokenRenewSelfRequest.md) |  | 

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


## PostAuthTokenRevoke

> PostAuthTokenRevoke(ctx).TokenRevokeRequest(tokenRevokeRequest).Execute()

This endpoint will delete the given token and all of its child tokens.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRevokeRequest := *openapiclient.NewTokenRevokeRequest() // TokenRevokeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRevoke(context.Background()).TokenRevokeRequest(tokenRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

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


## PostAuthTokenRevokeAccessor

> PostAuthTokenRevokeAccessor(ctx).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()

This endpoint will delete the token associated with the accessor and all of its child tokens.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRevokeAccessorRequest := *openapiclient.NewTokenRevokeAccessorRequest() // TokenRevokeAccessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRevokeAccessor(context.Background()).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRevokeAccessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRevokeAccessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeAccessorRequest** | [**TokenRevokeAccessorRequest**](TokenRevokeAccessorRequest.md) |  | 

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


## PostAuthTokenRevokeOrphan

> PostAuthTokenRevokeOrphan(ctx).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()

This endpoint will delete the token and orphan its child tokens.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRevokeOrphanRequest := *openapiclient.NewTokenRevokeOrphanRequest() // TokenRevokeOrphanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRevokeOrphan(context.Background()).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRevokeOrphan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRevokeOrphanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeOrphanRequest** | [**TokenRevokeOrphanRequest**](TokenRevokeOrphanRequest.md) |  | 

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


## PostAuthTokenRevokeSelf

> PostAuthTokenRevokeSelf(ctx).Execute()

This endpoint will delete the token used to call it and all of its child tokens.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthTokenRevokeSelf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRevokeSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRevokeSelfRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthTokenRolesRoleName

> PostAuthTokenRolesRoleName(ctx, roleName).TokenRolesRequest(tokenRolesRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleName := "roleName_example" // string | Name of the role
    tokenRolesRequest := *openapiclient.NewTokenRolesRequest() // TokenRolesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthTokenRolesRoleName(context.Background(), roleName).TokenRolesRequest(tokenRolesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenRolesRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRolesRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRolesRequest** | [**TokenRolesRequest**](TokenRolesRequest.md) |  | 

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


## PostAuthTokenTidy

> PostAuthTokenTidy(ctx).Execute()

This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.

### Example

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
    resp, r, err := apiClient.Auth.PostAuthTokenTidy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthTokenTidy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenTidyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthUserpassLoginUsername

> PostAuthUserpassLoginUsername(ctx, username).UserpassLoginRequest(userpassLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username of the user.
    userpassLoginRequest := *openapiclient.NewUserpassLoginRequest() // UserpassLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthUserpassLoginUsername(context.Background(), username).UserpassLoginRequest(userpassLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthUserpassLoginUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthUserpassLoginUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassLoginRequest** | [**UserpassLoginRequest**](UserpassLoginRequest.md) |  | 

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


## PostAuthUserpassUsersUsername

> PostAuthUserpassUsersUsername(ctx, username).UserpassUsersRequest(userpassUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username for this user.
    userpassUsersRequest := *openapiclient.NewUserpassUsersRequest() // UserpassUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthUserpassUsersUsername(context.Background(), username).UserpassUsersRequest(userpassUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthUserpassUsersUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username for this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthUserpassUsersUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersRequest** | [**UserpassUsersRequest**](UserpassUsersRequest.md) |  | 

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


## PostAuthUserpassUsersUsernamePassword

> PostAuthUserpassUsersUsernamePassword(ctx, username).UserpassUsersPasswordRequest(userpassUsersPasswordRequest).Execute()

Reset user's password.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username for this user.
    userpassUsersPasswordRequest := *openapiclient.NewUserpassUsersPasswordRequest() // UserpassUsersPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthUserpassUsersUsernamePassword(context.Background(), username).UserpassUsersPasswordRequest(userpassUsersPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthUserpassUsersUsernamePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username for this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthUserpassUsersUsernamePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPasswordRequest** | [**UserpassUsersPasswordRequest**](UserpassUsersPasswordRequest.md) |  | 

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


## PostAuthUserpassUsersUsernamePolicies

> PostAuthUserpassUsersUsernamePolicies(ctx, username).UserpassUsersPoliciesRequest(userpassUsersPoliciesRequest).Execute()

Update the policies associated with the username.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username for this user.
    userpassUsersPoliciesRequest := *openapiclient.NewUserpassUsersPoliciesRequest() // UserpassUsersPoliciesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Auth.PostAuthUserpassUsersUsernamePolicies(context.Background(), username).UserpassUsersPoliciesRequest(userpassUsersPoliciesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.PostAuthUserpassUsersUsernamePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username for this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthUserpassUsersUsernamePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPoliciesRequest** | [**UserpassUsersPoliciesRequest**](UserpassUsersPoliciesRequest.md) |  | 

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

