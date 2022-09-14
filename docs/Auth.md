# Auth

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
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role as it should appear in Vault.
	
	resp, err := client.Auth.DeleteAuthAlicloudRoleRole(context.Background(), role)
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAppIdMapAppIdKey

> DeleteAuthAppIdMapAppIdKey(ctx, key).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the app-id mapping
	
	resp, err := client.Auth.DeleteAuthAppIdMapAppIdKey(context.Background(), key)
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
**key** | **string** | Key for the app-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAppIdMapUserIdKey

> DeleteAuthAppIdMapUserIdKey(ctx, key).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the user-id mapping
	
	resp, err := client.Auth.DeleteAuthAppIdMapUserIdKey(context.Background(), key)
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
**key** | **string** | Key for the user-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleName

> DeleteAuthApproleRoleRoleName(ctx, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleName(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameBindSecretId

> DeleteAuthApproleRoleRoleNameBindSecretId(ctx, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameBoundCidrList

> DeleteAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNamePeriod

> DeleteAuthApproleRoleRoleNamePeriod(ctx, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNamePeriod(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNamePolicies

> DeleteAuthApproleRoleRoleNamePolicies(ctx, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNamePolicies(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy

> DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs

> DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdDestroy

> DeleteAuthApproleRoleRoleNameSecretIdDestroy(ctx, roleName).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdNumUses

> DeleteAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameSecretIdTtl

> DeleteAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenBoundCidrs

> DeleteAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenMaxTtl

> DeleteAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenNumUses

> DeleteAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthApproleRoleRoleNameTokenTtl

> DeleteAuthApproleRoleRoleNameTokenTtl(ctx, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.DeleteAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigCertificateCertName

> DeleteAuthAwsConfigCertificateCertName(ctx, certName).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	certName :=  // string | Name of the certificate.
	
	resp, err := client.Auth.DeleteAuthAwsConfigCertificateCertName(context.Background(), certName)
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigClient

> DeleteAuthAwsConfigClient(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAwsConfigClient(context.Background())
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


## DeleteAuthAwsConfigStsAccountId

> DeleteAuthAwsConfigStsAccountId(ctx, accountId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	accountId :=  // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	
	resp, err := client.Auth.DeleteAuthAwsConfigStsAccountId(context.Background(), accountId)
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
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsConfigTidyIdentityAccesslist

> DeleteAuthAwsConfigTidyIdentityAccesslist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAwsConfigTidyIdentityAccesslist(context.Background())
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


## DeleteAuthAwsConfigTidyIdentityWhitelist

> DeleteAuthAwsConfigTidyIdentityWhitelist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAwsConfigTidyIdentityWhitelist(context.Background())
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


## DeleteAuthAwsConfigTidyRoletagBlacklist

> DeleteAuthAwsConfigTidyRoletagBlacklist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAwsConfigTidyRoletagBlacklist(context.Background())
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


## DeleteAuthAwsConfigTidyRoletagDenylist

> DeleteAuthAwsConfigTidyRoletagDenylist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAwsConfigTidyRoletagDenylist(context.Background())
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


## DeleteAuthAwsIdentityAccesslistInstanceId

> DeleteAuthAwsIdentityAccesslistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	instanceId :=  // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	
	resp, err := client.Auth.DeleteAuthAwsIdentityAccesslistInstanceId(context.Background(), instanceId)
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsIdentityWhitelistInstanceId

> DeleteAuthAwsIdentityWhitelistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	instanceId :=  // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	
	resp, err := client.Auth.DeleteAuthAwsIdentityWhitelistInstanceId(context.Background(), instanceId)
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsRoleRole

> DeleteAuthAwsRoleRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthAwsRoleRole(context.Background(), role)
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


## DeleteAuthAwsRoletagBlacklistRoleTag

> DeleteAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.DeleteAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAwsRoletagDenylistRoleTag

> DeleteAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.DeleteAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthAzureConfig

> DeleteAuthAzureConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthAzureConfig(context.Background())
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


## DeleteAuthAzureRoleName

> DeleteAuthAzureRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthAzureRoleName(context.Background(), name)
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


## DeleteAuthCertCertsName

> DeleteAuthCertCertsName(ctx, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	resp, err := client.Auth.DeleteAuthCertCertsName(context.Background(), name)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthCertCrlsName

> DeleteAuthCertCrlsName(ctx, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	resp, err := client.Auth.DeleteAuthCertCrlsName(context.Background(), name)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthCfConfig

> DeleteAuthCfConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthCfConfig(context.Background())
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


## DeleteAuthCfRolesRole

> DeleteAuthCfRolesRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role.
	
	resp, err := client.Auth.DeleteAuthCfRolesRole(context.Background(), role)
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
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthGcpRoleName

> DeleteAuthGcpRoleName(ctx, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthGcpRoleName(context.Background(), name)
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


## DeleteAuthGithubMapTeamsKey

> DeleteAuthGithubMapTeamsKey(ctx, key).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the teams mapping
	
	resp, err := client.Auth.DeleteAuthGithubMapTeamsKey(context.Background(), key)
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthGithubMapUsersKey

> DeleteAuthGithubMapUsersKey(ctx, key).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the users mapping
	
	resp, err := client.Auth.DeleteAuthGithubMapUsersKey(context.Background(), key)
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthJwtRoleName

> DeleteAuthJwtRoleName(ctx, name).Execute()

Delete an existing role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthJwtRoleName(context.Background(), name)
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


## DeleteAuthKerberosGroupsName

> DeleteAuthKerberosGroupsName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	resp, err := client.Auth.DeleteAuthKerberosGroupsName(context.Background(), name)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthKubernetesRoleName

> DeleteAuthKubernetesRoleName(ctx, name).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthKubernetesRoleName(context.Background(), name)
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


## DeleteAuthLdapGroupsName

> DeleteAuthLdapGroupsName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	resp, err := client.Auth.DeleteAuthLdapGroupsName(context.Background(), name)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthLdapUsersName

> DeleteAuthLdapUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP user.
	
	resp, err := client.Auth.DeleteAuthLdapUsersName(context.Background(), name)
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOciConfig

> DeleteAuthOciConfig(ctx).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.DeleteAuthOciConfig(context.Background())
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


## DeleteAuthOciRoleRole

> DeleteAuthOciRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthOciRoleRole(context.Background(), role)
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


## DeleteAuthOidcRoleName

> DeleteAuthOidcRoleName(ctx, name).Execute()

Delete an existing role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.DeleteAuthOidcRoleName(context.Background(), name)
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


## DeleteAuthOktaGroupsName

> DeleteAuthOktaGroupsName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Okta group.
	
	resp, err := client.Auth.DeleteAuthOktaGroupsName(context.Background(), name)
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
**name** | **string** | Name of the Okta group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthOktaUsersName

> DeleteAuthOktaUsersName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the user.
	
	resp, err := client.Auth.DeleteAuthOktaUsersName(context.Background(), name)
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
**name** | **string** | Name of the user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthRadiusUsersName

> DeleteAuthRadiusUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the RADIUS user.
	
	resp, err := client.Auth.DeleteAuthRadiusUsersName(context.Background(), name)
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
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthTokenRolesRoleName

> DeleteAuthTokenRolesRoleName(ctx, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role
	
	resp, err := client.Auth.DeleteAuthTokenRolesRoleName(context.Background(), roleName)
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
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## DeleteAuthUserpassUsersUsername

> DeleteAuthUserpassUsersUsername(ctx, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username for this user.
	
	resp, err := client.Auth.DeleteAuthUserpassUsersUsername(context.Background(), username)
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAlicloudRole

> GetAuthAlicloudRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAlicloudRole(context.Background(), list)
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


## GetAuthAlicloudRoleRole

> GetAuthAlicloudRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role as it should appear in Vault.
	
	resp, err := client.Auth.GetAuthAlicloudRoleRole(context.Background(), role)
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAlicloudRoles

> GetAuthAlicloudRoles(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAlicloudRoles(context.Background(), list)
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


## GetAuthAppIdMapAppId

> GetAuthAppIdMapAppId(ctx).List(list).Execute()

Read mappings for app-id

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAppIdMapAppId(context.Background(), list)
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
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAppIdMapAppIdKey

> GetAuthAppIdMapAppIdKey(ctx, key).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the app-id mapping
	
	resp, err := client.Auth.GetAuthAppIdMapAppIdKey(context.Background(), key)
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
**key** | **string** | Key for the app-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAppIdMapUserId

> GetAuthAppIdMapUserId(ctx).List(list).Execute()

Read mappings for user-id

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAppIdMapUserId(context.Background(), list)
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
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAppIdMapUserIdKey

> GetAuthAppIdMapUserIdKey(ctx, key).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the user-id mapping
	
	resp, err := client.Auth.GetAuthAppIdMapUserIdKey(context.Background(), key)
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
**key** | **string** | Key for the user-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRole

> GetAuthApproleRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthApproleRole(context.Background(), list)
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


## GetAuthApproleRoleRoleName

> GetAuthApproleRoleRoleName(ctx, roleName).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleName(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameBindSecretId

> GetAuthApproleRoleRoleNameBindSecretId(ctx, roleName).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameBoundCidrList

> GetAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameLocalSecretIds

> GetAuthApproleRoleRoleNameLocalSecretIds(ctx, roleName).Execute()

Enables cluster local secret IDs

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameLocalSecretIds(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNamePeriod

> GetAuthApproleRoleRoleNamePeriod(ctx, roleName).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNamePeriod(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNamePolicies

> GetAuthApproleRoleRoleNamePolicies(ctx, roleName).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNamePolicies(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameRoleId

> GetAuthApproleRoleRoleNameRoleId(ctx, roleName).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameRoleId(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretId

> GetAuthApproleRoleRoleNameSecretId(ctx, roleName).List(list).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	list := NewstringWithDefaults()
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretId(context.Background(), roleName, list)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdBoundCidrs

> GetAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdNumUses

> GetAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameSecretIdTtl

> GetAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenBoundCidrs

> GetAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenMaxTtl

> GetAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenNumUses

> GetAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthApproleRoleRoleNameTokenTtl

> GetAuthApproleRoleRoleNameTokenTtl(ctx, roleName).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	resp, err := client.Auth.GetAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigCertificateCertName

> GetAuthAwsConfigCertificateCertName(ctx, certName).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	certName :=  // string | Name of the certificate.
	
	resp, err := client.Auth.GetAuthAwsConfigCertificateCertName(context.Background(), certName)
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigCertificates

> GetAuthAwsConfigCertificates(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsConfigCertificates(context.Background(), list)
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


## GetAuthAwsConfigClient

> GetAuthAwsConfigClient(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigClient(context.Background())
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


## GetAuthAwsConfigIdentity

> GetAuthAwsConfigIdentity(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigIdentity(context.Background())
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


## GetAuthAwsConfigSts

> GetAuthAwsConfigSts(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsConfigSts(context.Background(), list)
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


## GetAuthAwsConfigStsAccountId

> GetAuthAwsConfigStsAccountId(ctx, accountId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	accountId :=  // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	
	resp, err := client.Auth.GetAuthAwsConfigStsAccountId(context.Background(), accountId)
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
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsConfigTidyIdentityAccesslist

> GetAuthAwsConfigTidyIdentityAccesslist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigTidyIdentityAccesslist(context.Background())
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


## GetAuthAwsConfigTidyIdentityWhitelist

> GetAuthAwsConfigTidyIdentityWhitelist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigTidyIdentityWhitelist(context.Background())
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


## GetAuthAwsConfigTidyRoletagBlacklist

> GetAuthAwsConfigTidyRoletagBlacklist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigTidyRoletagBlacklist(context.Background())
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


## GetAuthAwsConfigTidyRoletagDenylist

> GetAuthAwsConfigTidyRoletagDenylist(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAwsConfigTidyRoletagDenylist(context.Background())
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


## GetAuthAwsIdentityAccesslist

> GetAuthAwsIdentityAccesslist(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsIdentityAccesslist(context.Background(), list)
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


## GetAuthAwsIdentityAccesslistInstanceId

> GetAuthAwsIdentityAccesslistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	instanceId :=  // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	
	resp, err := client.Auth.GetAuthAwsIdentityAccesslistInstanceId(context.Background(), instanceId)
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsIdentityWhitelist

> GetAuthAwsIdentityWhitelist(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsIdentityWhitelist(context.Background(), list)
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


## GetAuthAwsIdentityWhitelistInstanceId

> GetAuthAwsIdentityWhitelistInstanceId(ctx, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	instanceId :=  // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	
	resp, err := client.Auth.GetAuthAwsIdentityWhitelistInstanceId(context.Background(), instanceId)
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsRole

> GetAuthAwsRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsRole(context.Background(), list)
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


## GetAuthAwsRoleRole

> GetAuthAwsRoleRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsRoleRole(context.Background(), role)
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


## GetAuthAwsRoles

> GetAuthAwsRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsRoles(context.Background(), list)
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


## GetAuthAwsRoletagBlacklist

> GetAuthAwsRoletagBlacklist(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsRoletagBlacklist(context.Background(), list)
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


## GetAuthAwsRoletagBlacklistRoleTag

> GetAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.GetAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAwsRoletagDenylist

> GetAuthAwsRoletagDenylist(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAwsRoletagDenylist(context.Background(), list)
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


## GetAuthAwsRoletagDenylistRoleTag

> GetAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.GetAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthAzureConfig

> GetAuthAzureConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthAzureConfig(context.Background())
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


## GetAuthAzureRole

> GetAuthAzureRole(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAzureRole(context.Background(), list)
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


## GetAuthAzureRoleName

> GetAuthAzureRoleName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthAzureRoleName(context.Background(), name)
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


## GetAuthCentrifyConfig

> GetAuthCentrifyConfig(ctx).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthCentrifyConfig(context.Background())
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


## GetAuthCertCerts

> GetAuthCertCerts(ctx).List(list).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthCertCerts(context.Background(), list)
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


## GetAuthCertCertsName

> GetAuthCertCertsName(ctx, name).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	resp, err := client.Auth.GetAuthCertCertsName(context.Background(), name)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCertCrlsName

> GetAuthCertCrlsName(ctx, name).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	resp, err := client.Auth.GetAuthCertCrlsName(context.Background(), name)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthCfConfig

> GetAuthCfConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthCfConfig(context.Background())
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


## GetAuthCfRoles

> GetAuthCfRoles(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthCfRoles(context.Background(), list)
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


## GetAuthCfRolesRole

> GetAuthCfRolesRole(ctx, role).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role.
	
	resp, err := client.Auth.GetAuthCfRolesRole(context.Background(), role)
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
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGcpConfig

> GetAuthGcpConfig(ctx).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthGcpConfig(context.Background())
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


## GetAuthGcpRole

> GetAuthGcpRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthGcpRole(context.Background(), list)
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


## GetAuthGcpRoleName

> GetAuthGcpRoleName(ctx, name).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthGcpRoleName(context.Background(), name)
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


## GetAuthGcpRoles

> GetAuthGcpRoles(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthGcpRoles(context.Background(), list)
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


## GetAuthGithubConfig

> GetAuthGithubConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthGithubConfig(context.Background())
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


## GetAuthGithubMapTeams

> GetAuthGithubMapTeams(ctx).List(list).Execute()

Read mappings for teams

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthGithubMapTeams(context.Background(), list)
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
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapTeamsKey

> GetAuthGithubMapTeamsKey(ctx, key).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the teams mapping
	
	resp, err := client.Auth.GetAuthGithubMapTeamsKey(context.Background(), key)
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapUsers

> GetAuthGithubMapUsers(ctx).List(list).Execute()

Read mappings for users

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthGithubMapUsers(context.Background(), list)
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
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthGithubMapUsersKey

> GetAuthGithubMapUsersKey(ctx, key).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the users mapping
	
	resp, err := client.Auth.GetAuthGithubMapUsersKey(context.Background(), key)
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthJwtConfig

> GetAuthJwtConfig(ctx).Execute()

Read the current JWT authentication backend configuration.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthJwtConfig(context.Background())
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


## GetAuthJwtOidcCallback

> GetAuthJwtOidcCallback(ctx).Execute()

Callback endpoint to complete an OIDC login.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthJwtOidcCallback(context.Background())
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


## GetAuthJwtRole

> GetAuthJwtRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthJwtRole(context.Background(), list)
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


## GetAuthJwtRoleName

> GetAuthJwtRoleName(ctx, name).Execute()

Read an existing role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthJwtRoleName(context.Background(), name)
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


## GetAuthKerberosConfig

> GetAuthKerberosConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthKerberosConfig(context.Background())
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


## GetAuthKerberosConfigLdap

> GetAuthKerberosConfigLdap(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthKerberosConfigLdap(context.Background())
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


## GetAuthKerberosGroups

> GetAuthKerberosGroups(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthKerberosGroups(context.Background(), list)
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


## GetAuthKerberosGroupsName

> GetAuthKerberosGroupsName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	resp, err := client.Auth.GetAuthKerberosGroupsName(context.Background(), name)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthKerberosLogin

> GetAuthKerberosLogin(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthKerberosLogin(context.Background())
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


## GetAuthKubernetesConfig

> GetAuthKubernetesConfig(ctx).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthKubernetesConfig(context.Background())
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


## GetAuthKubernetesRole

> GetAuthKubernetesRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthKubernetesRole(context.Background(), list)
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


## GetAuthKubernetesRoleName

> GetAuthKubernetesRoleName(ctx, name).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthKubernetesRoleName(context.Background(), name)
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


## GetAuthLdapConfig

> GetAuthLdapConfig(ctx).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthLdapConfig(context.Background())
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


## GetAuthLdapGroups

> GetAuthLdapGroups(ctx).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthLdapGroups(context.Background(), list)
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


## GetAuthLdapGroupsName

> GetAuthLdapGroupsName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	resp, err := client.Auth.GetAuthLdapGroupsName(context.Background(), name)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthLdapUsers

> GetAuthLdapUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthLdapUsers(context.Background(), list)
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


## GetAuthLdapUsersName

> GetAuthLdapUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP user.
	
	resp, err := client.Auth.GetAuthLdapUsersName(context.Background(), name)
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOciConfig

> GetAuthOciConfig(ctx).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthOciConfig(context.Background())
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


## GetAuthOciRole

> GetAuthOciRole(ctx).List(list).Execute()

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOciRole(context.Background(), list)
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


## GetAuthOciRoleRole

> GetAuthOciRoleRole(ctx, role).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOciRoleRole(context.Background(), role)
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


## GetAuthOidcConfig

> GetAuthOidcConfig(ctx).Execute()

Read the current JWT authentication backend configuration.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthOidcConfig(context.Background())
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


## GetAuthOidcOidcCallback

> GetAuthOidcOidcCallback(ctx).Execute()

Callback endpoint to complete an OIDC login.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthOidcOidcCallback(context.Background())
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


## GetAuthOidcRole

> GetAuthOidcRole(ctx).List(list).Execute()

Lists all the roles registered with the backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOidcRole(context.Background(), list)
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


## GetAuthOidcRoleName

> GetAuthOidcRoleName(ctx, name).Execute()

Read an existing role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOidcRoleName(context.Background(), name)
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


## GetAuthOktaConfig

> GetAuthOktaConfig(ctx).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthOktaConfig(context.Background())
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


## GetAuthOktaGroups

> GetAuthOktaGroups(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOktaGroups(context.Background(), list)
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


## GetAuthOktaGroupsName

> GetAuthOktaGroupsName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Okta group.
	
	resp, err := client.Auth.GetAuthOktaGroupsName(context.Background(), name)
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
**name** | **string** | Name of the Okta group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaUsers

> GetAuthOktaUsers(ctx).List(list).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthOktaUsers(context.Background(), list)
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


## GetAuthOktaUsersName

> GetAuthOktaUsersName(ctx, name).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the user.
	
	resp, err := client.Auth.GetAuthOktaUsersName(context.Background(), name)
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
**name** | **string** | Name of the user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthOktaVerifyNonce

> GetAuthOktaVerifyNonce(ctx, nonce).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	nonce :=  // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.
	
	resp, err := client.Auth.GetAuthOktaVerifyNonce(context.Background(), nonce)
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
**nonce** | **string** | Nonce provided during a login request to retrieve the number verification challenge for the matching request. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthRadiusConfig

> GetAuthRadiusConfig(ctx).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthRadiusConfig(context.Background())
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


## GetAuthRadiusUsers

> GetAuthRadiusUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthRadiusUsers(context.Background(), list)
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


## GetAuthRadiusUsersName

> GetAuthRadiusUsersName(ctx, name).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the RADIUS user.
	
	resp, err := client.Auth.GetAuthRadiusUsersName(context.Background(), name)
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
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthTokenAccessors

> GetAuthTokenAccessors(ctx).List(list).Execute()

List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthTokenAccessors(context.Background(), list)
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


## GetAuthTokenLookup

> GetAuthTokenLookup(ctx).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthTokenLookup(context.Background())
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


## GetAuthTokenLookupSelf

> GetAuthTokenLookupSelf(ctx).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.GetAuthTokenLookupSelf(context.Background())
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


## GetAuthTokenRoles

> GetAuthTokenRoles(ctx).List(list).Execute()

This endpoint lists configured roles.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthTokenRoles(context.Background(), list)
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


## GetAuthTokenRolesRoleName

> GetAuthTokenRolesRoleName(ctx, roleName).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role
	
	resp, err := client.Auth.GetAuthTokenRolesRoleName(context.Background(), roleName)
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
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAuthUserpassUsers

> GetAuthUserpassUsers(ctx).List(list).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.GetAuthUserpassUsers(context.Background(), list)
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


## GetAuthUserpassUsersUsername

> GetAuthUserpassUsersUsername(ctx, username).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username for this user.
	
	resp, err := client.Auth.GetAuthUserpassUsersUsername(context.Background(), username)
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAlicloudLogin

> PostAuthAlicloudLogin(ctx).AlicloudLoginRequest(alicloudLoginRequest).Execute()

Authenticates an RAM entity with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	alicloudLoginRequest := NewAlicloudLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAlicloudLogin(context.Background(), alicloudLoginRequest)
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
 **alicloudLoginRequest** | [**AlicloudLoginRequest**](AlicloudLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAlicloudRoleRole

> PostAuthAlicloudRoleRole(ctx, role).AlicloudRoleRequest(alicloudRoleRequest).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role as it should appear in Vault.
	
	alicloudRoleRequest := NewAlicloudRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAlicloudRoleRole(context.Background(), role, alicloudRoleRequest)
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alicloudRoleRequest** | [**AlicloudRoleRequest**](AlicloudRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAppIdLogin

> PostAuthAppIdLogin(ctx).AppIdLoginRequest(appIdLoginRequest).Execute()

Log in with an App ID and User ID.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	appIdLoginRequest := NewAppIdLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAppIdLogin(context.Background(), appIdLoginRequest)
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
 **appIdLoginRequest** | [**AppIdLoginRequest**](AppIdLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAppIdLoginAppId

> PostAuthAppIdLoginAppId(ctx, appId).AppIdLoginRequest(appIdLoginRequest).Execute()

Log in with an App ID and User ID.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	appId :=  // string | The unique app ID
	
	appIdLoginRequest := NewAppIdLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAppIdLoginAppId(context.Background(), appId, appIdLoginRequest)
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
**appId** | **string** | The unique app ID | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdLoginRequest** | [**AppIdLoginRequest**](AppIdLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAppIdMapAppIdKey

> PostAuthAppIdMapAppIdKey(ctx, key).AppIdMapAppIdRequest(appIdMapAppIdRequest).Execute()

Read/write/delete a single app-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the app-id mapping
	
	appIdMapAppIdRequest := NewAppIdMapAppIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAppIdMapAppIdKey(context.Background(), key, appIdMapAppIdRequest)
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
**key** | **string** | Key for the app-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdMapAppIdRequest** | [**AppIdMapAppIdRequest**](AppIdMapAppIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAppIdMapUserIdKey

> PostAuthAppIdMapUserIdKey(ctx, key).AppIdMapUserIdRequest(appIdMapUserIdRequest).Execute()

Read/write/delete a single user-id mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the user-id mapping
	
	appIdMapUserIdRequest := NewAppIdMapUserIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAppIdMapUserIdKey(context.Background(), key, appIdMapUserIdRequest)
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
**key** | **string** | Key for the user-id mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appIdMapUserIdRequest** | [**AppIdMapUserIdRequest**](AppIdMapUserIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleLogin

> PostAuthApproleLogin(ctx).ApproleLoginRequest(approleLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	approleLoginRequest := NewApproleLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleLogin(context.Background(), approleLoginRequest)
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
 **approleLoginRequest** | [**ApproleLoginRequest**](ApproleLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleName

> PostAuthApproleRoleRoleName(ctx, roleName).ApproleRoleRequest(approleRoleRequest).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleRequest := NewApproleRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleName(context.Background(), roleName, approleRoleRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRequest** | [**ApproleRoleRequest**](ApproleRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameBindSecretId

> PostAuthApproleRoleRoleNameBindSecretId(ctx, roleName).ApproleRoleBindSecretIdRequest(approleRoleBindSecretIdRequest).Execute()

Impose secret_id to be presented during login using this role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleBindSecretIdRequest := NewApproleRoleBindSecretIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameBindSecretId(context.Background(), roleName, approleRoleBindSecretIdRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBindSecretIdRequest** | [**ApproleRoleBindSecretIdRequest**](ApproleRoleBindSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameBoundCidrList

> PostAuthApproleRoleRoleNameBoundCidrList(ctx, roleName).ApproleRoleBoundCidrListRequest(approleRoleBoundCidrListRequest).Execute()

Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleBoundCidrListRequest := NewApproleRoleBoundCidrListRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameBoundCidrList(context.Background(), roleName, approleRoleBoundCidrListRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleBoundCidrListRequest** | [**ApproleRoleBoundCidrListRequest**](ApproleRoleBoundCidrListRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameCustomSecretId

> PostAuthApproleRoleRoleNameCustomSecretId(ctx, roleName).ApproleRoleCustomSecretIdRequest(approleRoleCustomSecretIdRequest).Execute()

Assign a SecretID of choice against the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleCustomSecretIdRequest := NewApproleRoleCustomSecretIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameCustomSecretId(context.Background(), roleName, approleRoleCustomSecretIdRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleCustomSecretIdRequest** | [**ApproleRoleCustomSecretIdRequest**](ApproleRoleCustomSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNamePeriod

> PostAuthApproleRoleRoleNamePeriod(ctx, roleName).ApproleRolePeriodRequest(approleRolePeriodRequest).Execute()

Updates the value of 'period' on the role

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRolePeriodRequest := NewApproleRolePeriodRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNamePeriod(context.Background(), roleName, approleRolePeriodRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePeriodRequest** | [**ApproleRolePeriodRequest**](ApproleRolePeriodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNamePolicies

> PostAuthApproleRoleRoleNamePolicies(ctx, roleName).ApproleRolePoliciesRequest(approleRolePoliciesRequest).Execute()

Policies of the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRolePoliciesRequest := NewApproleRolePoliciesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNamePolicies(context.Background(), roleName, approleRolePoliciesRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRolePoliciesRequest** | [**ApproleRolePoliciesRequest**](ApproleRolePoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameRoleId

> PostAuthApproleRoleRoleNameRoleId(ctx, roleName).ApproleRoleRoleIdRequest(approleRoleRoleIdRequest).Execute()

Returns the 'role_id' of the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleRoleIdRequest := NewApproleRoleRoleIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameRoleId(context.Background(), roleName, approleRoleRoleIdRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleRoleIdRequest** | [**ApproleRoleRoleIdRequest**](ApproleRoleRoleIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretId

> PostAuthApproleRoleRoleNameSecretId(ctx, roleName).ApproleRoleSecretIdRequest(approleRoleSecretIdRequest).Execute()

Generate a SecretID against this role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdRequest := NewApproleRoleSecretIdRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretId(context.Background(), roleName, approleRoleSecretIdRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdRequest** | [**ApproleRoleSecretIdRequest**](ApproleRoleSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdAccessorDestroy

> PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(ctx, roleName).ApproleRoleSecretIdAccessorDestroyRequest(approleRoleSecretIdAccessorDestroyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdAccessorDestroyRequest := NewApproleRoleSecretIdAccessorDestroyRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorDestroy(context.Background(), roleName, approleRoleSecretIdAccessorDestroyRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorDestroyRequest** | [**ApproleRoleSecretIdAccessorDestroyRequest**](ApproleRoleSecretIdAccessorDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdAccessorLookup

> PostAuthApproleRoleRoleNameSecretIdAccessorLookup(ctx, roleName).ApproleRoleSecretIdAccessorLookupRequest(approleRoleSecretIdAccessorLookupRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdAccessorLookupRequest := NewApproleRoleSecretIdAccessorLookupRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdAccessorLookup(context.Background(), roleName, approleRoleSecretIdAccessorLookupRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdAccessorLookupRequest** | [**ApproleRoleSecretIdAccessorLookupRequest**](ApproleRoleSecretIdAccessorLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdBoundCidrs

> PostAuthApproleRoleRoleNameSecretIdBoundCidrs(ctx, roleName).ApproleRoleSecretIdBoundCidrsRequest(approleRoleSecretIdBoundCidrsRequest).Execute()

Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdBoundCidrsRequest := NewApproleRoleSecretIdBoundCidrsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdBoundCidrs(context.Background(), roleName, approleRoleSecretIdBoundCidrsRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdBoundCidrsRequest** | [**ApproleRoleSecretIdBoundCidrsRequest**](ApproleRoleSecretIdBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdDestroy

> PostAuthApproleRoleRoleNameSecretIdDestroy(ctx, roleName).ApproleRoleSecretIdDestroyRequest(approleRoleSecretIdDestroyRequest).Execute()

Invalidate an issued secret_id

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdDestroyRequest := NewApproleRoleSecretIdDestroyRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdDestroy(context.Background(), roleName, approleRoleSecretIdDestroyRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdDestroyRequest** | [**ApproleRoleSecretIdDestroyRequest**](ApproleRoleSecretIdDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdLookup

> PostAuthApproleRoleRoleNameSecretIdLookup(ctx, roleName).ApproleRoleSecretIdLookupRequest(approleRoleSecretIdLookupRequest).Execute()

Read the properties of an issued secret_id

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdLookupRequest := NewApproleRoleSecretIdLookupRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdLookup(context.Background(), roleName, approleRoleSecretIdLookupRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdLookupRequest** | [**ApproleRoleSecretIdLookupRequest**](ApproleRoleSecretIdLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdNumUses

> PostAuthApproleRoleRoleNameSecretIdNumUses(ctx, roleName).ApproleRoleSecretIdNumUsesRequest(approleRoleSecretIdNumUsesRequest).Execute()

Use limit of the SecretID generated against the role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdNumUsesRequest := NewApproleRoleSecretIdNumUsesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdNumUses(context.Background(), roleName, approleRoleSecretIdNumUsesRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdNumUsesRequest** | [**ApproleRoleSecretIdNumUsesRequest**](ApproleRoleSecretIdNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameSecretIdTtl

> PostAuthApproleRoleRoleNameSecretIdTtl(ctx, roleName).ApproleRoleSecretIdTtlRequest(approleRoleSecretIdTtlRequest).Execute()

Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using 'role/<role_name>/secret-id' or 'role/<role_name>/custom-secret-id' endpoints.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleSecretIdTtlRequest := NewApproleRoleSecretIdTtlRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameSecretIdTtl(context.Background(), roleName, approleRoleSecretIdTtlRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleSecretIdTtlRequest** | [**ApproleRoleSecretIdTtlRequest**](ApproleRoleSecretIdTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenBoundCidrs

> PostAuthApproleRoleRoleNameTokenBoundCidrs(ctx, roleName).ApproleRoleTokenBoundCidrsRequest(approleRoleTokenBoundCidrsRequest).Execute()

Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleTokenBoundCidrsRequest := NewApproleRoleTokenBoundCidrsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenBoundCidrs(context.Background(), roleName, approleRoleTokenBoundCidrsRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenBoundCidrsRequest** | [**ApproleRoleTokenBoundCidrsRequest**](ApproleRoleTokenBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenMaxTtl

> PostAuthApproleRoleRoleNameTokenMaxTtl(ctx, roleName).ApproleRoleTokenMaxTtlRequest(approleRoleTokenMaxTtlRequest).Execute()

Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleTokenMaxTtlRequest := NewApproleRoleTokenMaxTtlRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenMaxTtl(context.Background(), roleName, approleRoleTokenMaxTtlRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenMaxTtlRequest** | [**ApproleRoleTokenMaxTtlRequest**](ApproleRoleTokenMaxTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenNumUses

> PostAuthApproleRoleRoleNameTokenNumUses(ctx, roleName).ApproleRoleTokenNumUsesRequest(approleRoleTokenNumUsesRequest).Execute()

Number of times issued tokens can be used

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleTokenNumUsesRequest := NewApproleRoleTokenNumUsesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenNumUses(context.Background(), roleName, approleRoleTokenNumUsesRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenNumUsesRequest** | [**ApproleRoleTokenNumUsesRequest**](ApproleRoleTokenNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleRoleRoleNameTokenTtl

> PostAuthApproleRoleRoleNameTokenTtl(ctx, roleName).ApproleRoleTokenTtlRequest(approleRoleTokenTtlRequest).Execute()

Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role.
	
	approleRoleTokenTtlRequest := NewApproleRoleTokenTtlRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthApproleRoleRoleNameTokenTtl(context.Background(), roleName, approleRoleTokenTtlRequest)
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
**roleName** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approleRoleTokenTtlRequest** | [**ApproleRoleTokenTtlRequest**](ApproleRoleTokenTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthApproleTidySecretId

> PostAuthApproleTidySecretId(ctx).Execute()

Trigger the clean-up of expired SecretID entries.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthApproleTidySecretId(context.Background())
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


## PostAuthAwsConfigCertificateCertName

> PostAuthAwsConfigCertificateCertName(ctx, certName).AwsConfigCertificateRequest(awsConfigCertificateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	certName :=  // string | Name of the certificate.
	
	awsConfigCertificateRequest := NewAwsConfigCertificateRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigCertificateCertName(context.Background(), certName, awsConfigCertificateRequest)
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigCertificateRequest** | [**AwsConfigCertificateRequest**](AwsConfigCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigClient

> PostAuthAwsConfigClient(ctx).AwsConfigClientRequest(awsConfigClientRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigClientRequest := NewAwsConfigClientRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigClient(context.Background(), awsConfigClientRequest)
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
 **awsConfigClientRequest** | [**AwsConfigClientRequest**](AwsConfigClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigIdentity

> PostAuthAwsConfigIdentity(ctx).AwsConfigIdentityRequest(awsConfigIdentityRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigIdentityRequest := NewAwsConfigIdentityRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigIdentity(context.Background(), awsConfigIdentityRequest)
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
 **awsConfigIdentityRequest** | [**AwsConfigIdentityRequest**](AwsConfigIdentityRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigRotateRoot

> PostAuthAwsConfigRotateRoot(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthAwsConfigRotateRoot(context.Background())
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


## PostAuthAwsConfigStsAccountId

> PostAuthAwsConfigStsAccountId(ctx, accountId).AwsConfigStsRequest(awsConfigStsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	accountId :=  // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	
	awsConfigStsRequest := NewAwsConfigStsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigStsAccountId(context.Background(), accountId, awsConfigStsRequest)
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
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigStsRequest** | [**AwsConfigStsRequest**](AwsConfigStsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyIdentityAccesslist

> PostAuthAwsConfigTidyIdentityAccesslist(ctx).AwsConfigTidyIdentityAccesslistRequest(awsConfigTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigTidyIdentityAccesslistRequest := NewAwsConfigTidyIdentityAccesslistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigTidyIdentityAccesslist(context.Background(), awsConfigTidyIdentityAccesslistRequest)
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
 **awsConfigTidyIdentityAccesslistRequest** | [**AwsConfigTidyIdentityAccesslistRequest**](AwsConfigTidyIdentityAccesslistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyIdentityWhitelist

> PostAuthAwsConfigTidyIdentityWhitelist(ctx).AwsConfigTidyIdentityWhitelistRequest(awsConfigTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigTidyIdentityWhitelistRequest := NewAwsConfigTidyIdentityWhitelistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigTidyIdentityWhitelist(context.Background(), awsConfigTidyIdentityWhitelistRequest)
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
 **awsConfigTidyIdentityWhitelistRequest** | [**AwsConfigTidyIdentityWhitelistRequest**](AwsConfigTidyIdentityWhitelistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyRoletagBlacklist

> PostAuthAwsConfigTidyRoletagBlacklist(ctx).AwsConfigTidyRoletagBlacklistRequest(awsConfigTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigTidyRoletagBlacklistRequest := NewAwsConfigTidyRoletagBlacklistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigTidyRoletagBlacklist(context.Background(), awsConfigTidyRoletagBlacklistRequest)
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
 **awsConfigTidyRoletagBlacklistRequest** | [**AwsConfigTidyRoletagBlacklistRequest**](AwsConfigTidyRoletagBlacklistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsConfigTidyRoletagDenylist

> PostAuthAwsConfigTidyRoletagDenylist(ctx).AwsConfigTidyRoletagDenylistRequest(awsConfigTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsConfigTidyRoletagDenylistRequest := NewAwsConfigTidyRoletagDenylistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsConfigTidyRoletagDenylist(context.Background(), awsConfigTidyRoletagDenylistRequest)
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
 **awsConfigTidyRoletagDenylistRequest** | [**AwsConfigTidyRoletagDenylistRequest**](AwsConfigTidyRoletagDenylistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsLogin

> PostAuthAwsLogin(ctx).AwsLoginRequest(awsLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsLoginRequest := NewAwsLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsLogin(context.Background(), awsLoginRequest)
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
 **awsLoginRequest** | [**AwsLoginRequest**](AwsLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoleRole

> PostAuthAwsRoleRole(ctx, role).AwsRoleRequest(awsRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	awsRoleRequest := NewAwsRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsRoleRole(context.Background(), role, awsRoleRequest)
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

 **awsRoleRequest** | [**AwsRoleRequest**](AwsRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoleRoleTag

> PostAuthAwsRoleRoleTag(ctx, role).AwsRoleTagRequest(awsRoleTagRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	awsRoleTagRequest := NewAwsRoleTagRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsRoleRoleTag(context.Background(), role, awsRoleTagRequest)
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

 **awsRoleTagRequest** | [**AwsRoleTagRequest**](AwsRoleTagRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoletagBlacklistRoleTag

> PostAuthAwsRoletagBlacklistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.PostAuthAwsRoletagBlacklistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsRoletagDenylistRoleTag

> PostAuthAwsRoletagDenylistRoleTag(ctx, roleTag).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleTag :=  // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	
	resp, err := client.Auth.PostAuthAwsRoletagDenylistRoleTag(context.Background(), roleTag)
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyIdentityAccesslist

> PostAuthAwsTidyIdentityAccesslist(ctx).AwsTidyIdentityAccesslistRequest(awsTidyIdentityAccesslistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsTidyIdentityAccesslistRequest := NewAwsTidyIdentityAccesslistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsTidyIdentityAccesslist(context.Background(), awsTidyIdentityAccesslistRequest)
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
 **awsTidyIdentityAccesslistRequest** | [**AwsTidyIdentityAccesslistRequest**](AwsTidyIdentityAccesslistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyIdentityWhitelist

> PostAuthAwsTidyIdentityWhitelist(ctx).AwsTidyIdentityWhitelistRequest(awsTidyIdentityWhitelistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsTidyIdentityWhitelistRequest := NewAwsTidyIdentityWhitelistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsTidyIdentityWhitelist(context.Background(), awsTidyIdentityWhitelistRequest)
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
 **awsTidyIdentityWhitelistRequest** | [**AwsTidyIdentityWhitelistRequest**](AwsTidyIdentityWhitelistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyRoletagBlacklist

> PostAuthAwsTidyRoletagBlacklist(ctx).AwsTidyRoletagBlacklistRequest(awsTidyRoletagBlacklistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsTidyRoletagBlacklistRequest := NewAwsTidyRoletagBlacklistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsTidyRoletagBlacklist(context.Background(), awsTidyRoletagBlacklistRequest)
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
 **awsTidyRoletagBlacklistRequest** | [**AwsTidyRoletagBlacklistRequest**](AwsTidyRoletagBlacklistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAwsTidyRoletagDenylist

> PostAuthAwsTidyRoletagDenylist(ctx).AwsTidyRoletagDenylistRequest(awsTidyRoletagDenylistRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	awsTidyRoletagDenylistRequest := NewAwsTidyRoletagDenylistRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAwsTidyRoletagDenylist(context.Background(), awsTidyRoletagDenylistRequest)
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
 **awsTidyRoletagDenylistRequest** | [**AwsTidyRoletagDenylistRequest**](AwsTidyRoletagDenylistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAzureConfig

> PostAuthAzureConfig(ctx).AzureConfigRequest(azureConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.PostAuthAzureConfig(context.Background(), azureConfigRequest)
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


## PostAuthAzureLogin

> PostAuthAzureLogin(ctx).AzureLoginRequest(azureLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	azureLoginRequest := NewAzureLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAzureLogin(context.Background(), azureLoginRequest)
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
 **azureLoginRequest** | [**AzureLoginRequest**](AzureLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthAzureRoleName

> PostAuthAzureRoleName(ctx, name).AzureRoleRequest(azureRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	azureRoleRequest := NewAzureRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthAzureRoleName(context.Background(), name, azureRoleRequest)
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

 **azureRoleRequest** | [**AzureRoleRequest**](AzureRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCentrifyConfig

> PostAuthCentrifyConfig(ctx).CentrifyConfigRequest(centrifyConfigRequest).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	centrifyConfigRequest := NewCentrifyConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCentrifyConfig(context.Background(), centrifyConfigRequest)
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
 **centrifyConfigRequest** | [**CentrifyConfigRequest**](CentrifyConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCentrifyLogin

> PostAuthCentrifyLogin(ctx).CentrifyLoginRequest(centrifyLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	centrifyLoginRequest := NewCentrifyLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCentrifyLogin(context.Background(), centrifyLoginRequest)
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
 **centrifyLoginRequest** | [**CentrifyLoginRequest**](CentrifyLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertCertsName

> PostAuthCertCertsName(ctx, name).CertCertsRequest(certCertsRequest).Execute()

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	certCertsRequest := NewCertCertsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCertCertsName(context.Background(), name, certCertsRequest)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCertsRequest** | [**CertCertsRequest**](CertCertsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertConfig

> PostAuthCertConfig(ctx).CertConfigRequest(certConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	certConfigRequest := NewCertConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCertConfig(context.Background(), certConfigRequest)
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
 **certConfigRequest** | [**CertConfigRequest**](CertConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertCrlsName

> PostAuthCertCrlsName(ctx, name).CertCrlsRequest(certCrlsRequest).Execute()

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the certificate
	
	certCrlsRequest := NewCertCrlsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCertCrlsName(context.Background(), name, certCrlsRequest)
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certCrlsRequest** | [**CertCrlsRequest**](CertCrlsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCertLogin

> PostAuthCertLogin(ctx).CertLoginRequest(certLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	certLoginRequest := NewCertLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCertLogin(context.Background(), certLoginRequest)
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
 **certLoginRequest** | [**CertLoginRequest**](CertLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfConfig

> PostAuthCfConfig(ctx).CfConfigRequest(cfConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	cfConfigRequest := NewCfConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCfConfig(context.Background(), cfConfigRequest)
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
 **cfConfigRequest** | [**CfConfigRequest**](CfConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfLogin

> PostAuthCfLogin(ctx).CfLoginRequest(cfLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	cfLoginRequest := NewCfLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCfLogin(context.Background(), cfLoginRequest)
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
 **cfLoginRequest** | [**CfLoginRequest**](CfLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthCfRolesRole

> PostAuthCfRolesRole(ctx, role).CfRolesRequest(cfRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	role :=  // string | The name of the role.
	
	cfRolesRequest := NewCfRolesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthCfRolesRole(context.Background(), role, cfRolesRequest)
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
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cfRolesRequest** | [**CfRolesRequest**](CfRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpConfig

> PostAuthGcpConfig(ctx).GcpConfigRequest(gcpConfigRequest).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.PostAuthGcpConfig(context.Background(), gcpConfigRequest)
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


## PostAuthGcpLogin

> PostAuthGcpLogin(ctx).GcpLoginRequest(gcpLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	gcpLoginRequest := NewGcpLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGcpLogin(context.Background(), gcpLoginRequest)
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
 **gcpLoginRequest** | [**GcpLoginRequest**](GcpLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleName

> PostAuthGcpRoleName(ctx, name).GcpRoleRequest(gcpRoleRequest).Execute()

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	gcpRoleRequest := NewGcpRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGcpRoleName(context.Background(), name, gcpRoleRequest)
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

 **gcpRoleRequest** | [**GcpRoleRequest**](GcpRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleNameLabels

> PostAuthGcpRoleNameLabels(ctx, name).GcpRoleLabelsRequest(gcpRoleLabelsRequest).Execute()

Add or remove labels for an existing 'gce' role

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	gcpRoleLabelsRequest := NewGcpRoleLabelsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGcpRoleNameLabels(context.Background(), name, gcpRoleLabelsRequest)
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

 **gcpRoleLabelsRequest** | [**GcpRoleLabelsRequest**](GcpRoleLabelsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGcpRoleNameServiceAccounts

> PostAuthGcpRoleNameServiceAccounts(ctx, name).GcpRoleServiceAccountsRequest(gcpRoleServiceAccountsRequest).Execute()

Add or remove service accounts for an existing `iam` role

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	gcpRoleServiceAccountsRequest := NewGcpRoleServiceAccountsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGcpRoleNameServiceAccounts(context.Background(), name, gcpRoleServiceAccountsRequest)
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

 **gcpRoleServiceAccountsRequest** | [**GcpRoleServiceAccountsRequest**](GcpRoleServiceAccountsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubConfig

> PostAuthGithubConfig(ctx).GithubConfigRequest(githubConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	githubConfigRequest := NewGithubConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGithubConfig(context.Background(), githubConfigRequest)
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
 **githubConfigRequest** | [**GithubConfigRequest**](GithubConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubLogin

> PostAuthGithubLogin(ctx).GithubLoginRequest(githubLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	githubLoginRequest := NewGithubLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGithubLogin(context.Background(), githubLoginRequest)
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
 **githubLoginRequest** | [**GithubLoginRequest**](GithubLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubMapTeamsKey

> PostAuthGithubMapTeamsKey(ctx, key).GithubMapTeamsRequest(githubMapTeamsRequest).Execute()

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the teams mapping
	
	githubMapTeamsRequest := NewGithubMapTeamsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGithubMapTeamsKey(context.Background(), key, githubMapTeamsRequest)
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapTeamsRequest** | [**GithubMapTeamsRequest**](GithubMapTeamsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthGithubMapUsersKey

> PostAuthGithubMapUsersKey(ctx, key).GithubMapUsersRequest(githubMapUsersRequest).Execute()

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	key :=  // string | Key for the users mapping
	
	githubMapUsersRequest := NewGithubMapUsersRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthGithubMapUsersKey(context.Background(), key, githubMapUsersRequest)
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubMapUsersRequest** | [**GithubMapUsersRequest**](GithubMapUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtConfig

> PostAuthJwtConfig(ctx).JwtConfigRequest(jwtConfigRequest).Execute()

Configure the JWT authentication backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	jwtConfigRequest := NewJwtConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthJwtConfig(context.Background(), jwtConfigRequest)
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
 **jwtConfigRequest** | [**JwtConfigRequest**](JwtConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtLogin

> PostAuthJwtLogin(ctx).JwtLoginRequest(jwtLoginRequest).Execute()

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	jwtLoginRequest := NewJwtLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthJwtLogin(context.Background(), jwtLoginRequest)
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
 **jwtLoginRequest** | [**JwtLoginRequest**](JwtLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtOidcAuthUrl

> PostAuthJwtOidcAuthUrl(ctx).JwtOidcAuthUrlRequest(jwtOidcAuthUrlRequest).Execute()

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	jwtOidcAuthUrlRequest := NewJwtOidcAuthUrlRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthJwtOidcAuthUrl(context.Background(), jwtOidcAuthUrlRequest)
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
 **jwtOidcAuthUrlRequest** | [**JwtOidcAuthUrlRequest**](JwtOidcAuthUrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtOidcCallback

> PostAuthJwtOidcCallback(ctx).JwtOidcCallbackRequest(jwtOidcCallbackRequest).Execute()

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	jwtOidcCallbackRequest := NewJwtOidcCallbackRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthJwtOidcCallback(context.Background(), jwtOidcCallbackRequest)
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
 **jwtOidcCallbackRequest** | [**JwtOidcCallbackRequest**](JwtOidcCallbackRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthJwtRoleName

> PostAuthJwtRoleName(ctx, name).JwtRoleRequest(jwtRoleRequest).Execute()

Register an role with the backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	jwtRoleRequest := NewJwtRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthJwtRoleName(context.Background(), name, jwtRoleRequest)
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

 **jwtRoleRequest** | [**JwtRoleRequest**](JwtRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosConfig

> PostAuthKerberosConfig(ctx).KerberosConfigRequest(kerberosConfigRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kerberosConfigRequest := NewKerberosConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKerberosConfig(context.Background(), kerberosConfigRequest)
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
 **kerberosConfigRequest** | [**KerberosConfigRequest**](KerberosConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosConfigLdap

> PostAuthKerberosConfigLdap(ctx).KerberosConfigLdapRequest(kerberosConfigLdapRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kerberosConfigLdapRequest := NewKerberosConfigLdapRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKerberosConfigLdap(context.Background(), kerberosConfigLdapRequest)
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
 **kerberosConfigLdapRequest** | [**KerberosConfigLdapRequest**](KerberosConfigLdapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosGroupsName

> PostAuthKerberosGroupsName(ctx, name).KerberosGroupsRequest(kerberosGroupsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	kerberosGroupsRequest := NewKerberosGroupsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKerberosGroupsName(context.Background(), name, kerberosGroupsRequest)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosGroupsRequest** | [**KerberosGroupsRequest**](KerberosGroupsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKerberosLogin

> PostAuthKerberosLogin(ctx).KerberosLoginRequest(kerberosLoginRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kerberosLoginRequest := NewKerberosLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKerberosLogin(context.Background(), kerberosLoginRequest)
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
 **kerberosLoginRequest** | [**KerberosLoginRequest**](KerberosLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKubernetesConfig

> PostAuthKubernetesConfig(ctx).KubernetesConfigRequest(kubernetesConfigRequest).Execute()

Configures the JWT Public Key and Kubernetes API information.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.Auth.PostAuthKubernetesConfig(context.Background(), kubernetesConfigRequest)
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


## PostAuthKubernetesLogin

> PostAuthKubernetesLogin(ctx).KubernetesLoginRequest(kubernetesLoginRequest).Execute()

Authenticates Kubernetes service accounts with Vault.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	kubernetesLoginRequest := NewKubernetesLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKubernetesLogin(context.Background(), kubernetesLoginRequest)
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
 **kubernetesLoginRequest** | [**KubernetesLoginRequest**](KubernetesLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthKubernetesRoleName

> PostAuthKubernetesRoleName(ctx, name).KubernetesRoleRequest(kubernetesRoleRequest).Execute()

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	kubernetesRoleRequest := NewKubernetesRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthKubernetesRoleName(context.Background(), name, kubernetesRoleRequest)
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

 **kubernetesRoleRequest** | [**KubernetesRoleRequest**](KubernetesRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapConfig

> PostAuthLdapConfig(ctx).LdapConfigRequest(ldapConfigRequest).Execute()

Configure the LDAP server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	ldapConfigRequest := NewLdapConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthLdapConfig(context.Background(), ldapConfigRequest)
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
 **ldapConfigRequest** | [**LdapConfigRequest**](LdapConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapGroupsName

> PostAuthLdapGroupsName(ctx, name).LdapGroupsRequest(ldapGroupsRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP group.
	
	ldapGroupsRequest := NewLdapGroupsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthLdapGroupsName(context.Background(), name, ldapGroupsRequest)
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapGroupsRequest** | [**LdapGroupsRequest**](LdapGroupsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapLoginUsername

> PostAuthLdapLoginUsername(ctx, username).LdapLoginRequest(ldapLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | DN (distinguished name) to be used for login.
	
	ldapLoginRequest := NewLdapLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthLdapLoginUsername(context.Background(), username, ldapLoginRequest)
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
**username** | **string** | DN (distinguished name) to be used for login. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapLoginRequest** | [**LdapLoginRequest**](LdapLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthLdapUsersName

> PostAuthLdapUsersName(ctx, name).LdapUsersRequest(ldapUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the LDAP user.
	
	ldapUsersRequest := NewLdapUsersRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthLdapUsersName(context.Background(), name, ldapUsersRequest)
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapUsersRequest** | [**LdapUsersRequest**](LdapUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciConfig

> PostAuthOciConfig(ctx).OciConfigRequest(ociConfigRequest).Execute()

Manages the configuration for the Vault Auth Plugin.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	ociConfigRequest := NewOciConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOciConfig(context.Background(), ociConfigRequest)
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
 **ociConfigRequest** | [**OciConfigRequest**](OciConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciLoginRole

> PostAuthOciLoginRole(ctx, role).OciLoginRequest(ociLoginRequest).Execute()

Authenticates to Vault using OCI credentials

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	ociLoginRequest := NewOciLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOciLoginRole(context.Background(), role, ociLoginRequest)
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

 **ociLoginRequest** | [**OciLoginRequest**](OciLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOciRoleRole

> PostAuthOciRoleRole(ctx, role).OciRoleRequest(ociRoleRequest).Execute()

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	ociRoleRequest := NewOciRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOciRoleRole(context.Background(), role, ociRoleRequest)
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

 **ociRoleRequest** | [**OciRoleRequest**](OciRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcConfig

> PostAuthOidcConfig(ctx).OidcConfigRequest(oidcConfigRequest).Execute()

Configure the JWT authentication backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	oidcConfigRequest := NewOidcConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOidcConfig(context.Background(), oidcConfigRequest)
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
 **oidcConfigRequest** | [**OidcConfigRequest**](OidcConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcLogin

> PostAuthOidcLogin(ctx).OidcLoginRequest(oidcLoginRequest).Execute()

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	oidcLoginRequest := NewOidcLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOidcLogin(context.Background(), oidcLoginRequest)
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
 **oidcLoginRequest** | [**OidcLoginRequest**](OidcLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcOidcAuthUrl

> PostAuthOidcOidcAuthUrl(ctx).OidcOidcAuthUrlRequest(oidcOidcAuthUrlRequest).Execute()

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	oidcOidcAuthUrlRequest := NewOidcOidcAuthUrlRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOidcOidcAuthUrl(context.Background(), oidcOidcAuthUrlRequest)
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
 **oidcOidcAuthUrlRequest** | [**OidcOidcAuthUrlRequest**](OidcOidcAuthUrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcOidcCallback

> PostAuthOidcOidcCallback(ctx).OidcOidcCallbackRequest(oidcOidcCallbackRequest).Execute()

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	oidcOidcCallbackRequest := NewOidcOidcCallbackRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOidcOidcCallback(context.Background(), oidcOidcCallbackRequest)
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
 **oidcOidcCallbackRequest** | [**OidcOidcCallbackRequest**](OidcOidcCallbackRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOidcRoleName

> PostAuthOidcRoleName(ctx, name).OidcRoleRequest(oidcRoleRequest).Execute()

Register an role with the backend.



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
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
	
	oidcRoleRequest := NewOidcRoleRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOidcRoleName(context.Background(), name, oidcRoleRequest)
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

 **oidcRoleRequest** | [**OidcRoleRequest**](OidcRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaConfig

> PostAuthOktaConfig(ctx).OktaConfigRequest(oktaConfigRequest).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	oktaConfigRequest := NewOktaConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOktaConfig(context.Background(), oktaConfigRequest)
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
 **oktaConfigRequest** | [**OktaConfigRequest**](OktaConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaGroupsName

> PostAuthOktaGroupsName(ctx, name).OktaGroupsRequest(oktaGroupsRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the Okta group.
	
	oktaGroupsRequest := NewOktaGroupsRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOktaGroupsName(context.Background(), name, oktaGroupsRequest)
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
**name** | **string** | Name of the Okta group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaGroupsRequest** | [**OktaGroupsRequest**](OktaGroupsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaLoginUsername

> PostAuthOktaLoginUsername(ctx, username).OktaLoginRequest(oktaLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username to be used for login.
	
	oktaLoginRequest := NewOktaLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOktaLoginUsername(context.Background(), username, oktaLoginRequest)
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
**username** | **string** | Username to be used for login. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaLoginRequest** | [**OktaLoginRequest**](OktaLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthOktaUsersName

> PostAuthOktaUsersName(ctx, name).OktaUsersRequest(oktaUsersRequest).Execute()

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the user.
	
	oktaUsersRequest := NewOktaUsersRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthOktaUsersName(context.Background(), name, oktaUsersRequest)
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
**name** | **string** | Name of the user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaUsersRequest** | [**OktaUsersRequest**](OktaUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusConfig

> PostAuthRadiusConfig(ctx).RadiusConfigRequest(radiusConfigRequest).Execute()

Configure the RADIUS server to connect to, along with its options.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	radiusConfigRequest := NewRadiusConfigRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthRadiusConfig(context.Background(), radiusConfigRequest)
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
 **radiusConfigRequest** | [**RadiusConfigRequest**](RadiusConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusLogin

> PostAuthRadiusLogin(ctx).RadiusLoginRequest(radiusLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthRadiusLogin(context.Background(), radiusLoginRequest)
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
 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusLoginUrlusername

> PostAuthRadiusLoginUrlusername(ctx, urlusername).RadiusLoginRequest(radiusLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlusername :=  // string | Username to be used for login. (URL parameter)
	
	radiusLoginRequest := NewRadiusLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthRadiusLoginUrlusername(context.Background(), urlusername, radiusLoginRequest)
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
**urlusername** | **string** | Username to be used for login. (URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthRadiusUsersName

> PostAuthRadiusUsersName(ctx, name).RadiusUsersRequest(radiusUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the RADIUS user.
	
	radiusUsersRequest := NewRadiusUsersRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthRadiusUsersName(context.Background(), name, radiusUsersRequest)
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
**name** | **string** | Name of the RADIUS user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusUsersRequest** | [**RadiusUsersRequest**](RadiusUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenCreate

> PostAuthTokenCreate(ctx).Execute()

The token create path is used to create new tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthTokenCreate(context.Background())
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


## PostAuthTokenCreateOrphan

> PostAuthTokenCreateOrphan(ctx).Execute()

The token create path is used to create new orphan tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthTokenCreateOrphan(context.Background())
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


## PostAuthTokenCreateRoleName

> PostAuthTokenCreateRoleName(ctx, roleName).Execute()

This token create path is used to create new tokens adhering to the given role.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role
	
	resp, err := client.Auth.PostAuthTokenCreateRoleName(context.Background(), roleName)
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
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookup

> PostAuthTokenLookup(ctx).TokenLookupRequest(tokenLookupRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenLookupRequest := NewTokenLookupRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenLookup(context.Background(), tokenLookupRequest)
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
 **tokenLookupRequest** | [**TokenLookupRequest**](TokenLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookupAccessor

> PostAuthTokenLookupAccessor(ctx).TokenLookupAccessorRequest(tokenLookupAccessorRequest).Execute()

This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenLookupAccessorRequest := NewTokenLookupAccessorRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenLookupAccessor(context.Background(), tokenLookupAccessorRequest)
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
 **tokenLookupAccessorRequest** | [**TokenLookupAccessorRequest**](TokenLookupAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenLookupSelf

> PostAuthTokenLookupSelf(ctx).TokenLookupSelfRequest(tokenLookupSelfRequest).Execute()

This endpoint will lookup a token and its properties.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenLookupSelfRequest := NewTokenLookupSelfRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenLookupSelf(context.Background(), tokenLookupSelfRequest)
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
 **tokenLookupSelfRequest** | [**TokenLookupSelfRequest**](TokenLookupSelfRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenew

> PostAuthTokenRenew(ctx).TokenRenewRequest(tokenRenewRequest).Execute()

This endpoint will renew the given token and prevent expiration.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRenewRequest := NewTokenRenewRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRenew(context.Background(), tokenRenewRequest)
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
 **tokenRenewRequest** | [**TokenRenewRequest**](TokenRenewRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenewAccessor

> PostAuthTokenRenewAccessor(ctx).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()

This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRenewAccessorRequest := NewTokenRenewAccessorRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRenewAccessor(context.Background(), tokenRenewAccessorRequest)
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
 **tokenRenewAccessorRequest** | [**TokenRenewAccessorRequest**](TokenRenewAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRenewSelf

> PostAuthTokenRenewSelf(ctx).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()

This endpoint will renew the token used to call it and prevent expiration.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRenewSelfRequest := NewTokenRenewSelfRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRenewSelf(context.Background(), tokenRenewSelfRequest)
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
 **tokenRenewSelfRequest** | [**TokenRenewSelfRequest**](TokenRenewSelfRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevoke

> PostAuthTokenRevoke(ctx).TokenRevokeRequest(tokenRevokeRequest).Execute()

This endpoint will delete the given token and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRevokeRequest := NewTokenRevokeRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRevoke(context.Background(), tokenRevokeRequest)
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
 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeAccessor

> PostAuthTokenRevokeAccessor(ctx).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()

This endpoint will delete the token associated with the accessor and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRevokeAccessorRequest := NewTokenRevokeAccessorRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRevokeAccessor(context.Background(), tokenRevokeAccessorRequest)
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
 **tokenRevokeAccessorRequest** | [**TokenRevokeAccessorRequest**](TokenRevokeAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeOrphan

> PostAuthTokenRevokeOrphan(ctx).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()

This endpoint will delete the token and orphan its child tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	tokenRevokeOrphanRequest := NewTokenRevokeOrphanRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRevokeOrphan(context.Background(), tokenRevokeOrphanRequest)
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
 **tokenRevokeOrphanRequest** | [**TokenRevokeOrphanRequest**](TokenRevokeOrphanRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenRevokeSelf

> PostAuthTokenRevokeSelf(ctx).Execute()

This endpoint will delete the token used to call it and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthTokenRevokeSelf(context.Background())
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


## PostAuthTokenRolesRoleName

> PostAuthTokenRolesRoleName(ctx, roleName).TokenRolesRequest(tokenRolesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	roleName :=  // string | Name of the role
	
	tokenRolesRequest := NewTokenRolesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthTokenRolesRoleName(context.Background(), roleName, tokenRolesRequest)
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
**roleName** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRolesRequest** | [**TokenRolesRequest**](TokenRolesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthTokenTidy

> PostAuthTokenTidy(ctx).Execute()

This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.Auth.PostAuthTokenTidy(context.Background())
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


## PostAuthUserpassLoginUsername

> PostAuthUserpassLoginUsername(ctx, username).UserpassLoginRequest(userpassLoginRequest).Execute()

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username of the user.
	
	userpassLoginRequest := NewUserpassLoginRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthUserpassLoginUsername(context.Background(), username, userpassLoginRequest)
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
**username** | **string** | Username of the user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassLoginRequest** | [**UserpassLoginRequest**](UserpassLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsername

> PostAuthUserpassUsersUsername(ctx, username).UserpassUsersRequest(userpassUsersRequest).Execute()

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username for this user.
	
	userpassUsersRequest := NewUserpassUsersRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthUserpassUsersUsername(context.Background(), username, userpassUsersRequest)
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersRequest** | [**UserpassUsersRequest**](UserpassUsersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsernamePassword

> PostAuthUserpassUsersUsernamePassword(ctx, username).UserpassUsersPasswordRequest(userpassUsersPasswordRequest).Execute()

Reset user's password.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username for this user.
	
	userpassUsersPasswordRequest := NewUserpassUsersPasswordRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthUserpassUsersUsernamePassword(context.Background(), username, userpassUsersPasswordRequest)
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPasswordRequest** | [**UserpassUsersPasswordRequest**](UserpassUsersPasswordRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## PostAuthUserpassUsersUsernamePolicies

> PostAuthUserpassUsersUsernamePolicies(ctx, username).UserpassUsersPoliciesRequest(userpassUsersPoliciesRequest).Execute()

Update the policies associated with the username.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	username :=  // string | Username for this user.
	
	userpassUsersPoliciesRequest := NewUserpassUsersPoliciesRequestWithDefaults()
	
	resp, err := client.Auth.PostAuthUserpassUsersUsernamePolicies(context.Background(), username, userpassUsersPoliciesRequest)
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassUsersPoliciesRequest** | [**UserpassUsersPoliciesRequest**](UserpassUsersPoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

