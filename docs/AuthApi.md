# Auth

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliCloudDeleteAuthRole**](AuthApi.md#AliCloudDeleteAuthRole) | **Delete** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudListAuthRoles**](AuthApi.md#AliCloudListAuthRoles) | **Get** /auth/{alicloud_mount_path}/role | Lists all the roles that are registered with Vault.
[**AliCloudLogin**](AuthApi.md#AliCloudLogin) | **Post** /auth/{alicloud_mount_path}/login | Authenticates an RAM entity with Vault.
[**AliCloudReadAuthRole**](AuthApi.md#AliCloudReadAuthRole) | **Get** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudWriteAuthRole**](AuthApi.md#AliCloudWriteAuthRole) | **Post** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AppRoleDeleteBindSecretId**](AuthApi.md#AppRoleDeleteBindSecretId) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleDeleteBoundCidrList**](AuthApi.md#AppRoleDeleteBoundCidrList) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleDeletePeriod**](AuthApi.md#AppRoleDeletePeriod) | **Delete** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleDeletePolicies**](AuthApi.md#AppRoleDeletePolicies) | **Delete** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleDeleteRole**](AuthApi.md#AppRoleDeleteRole) | **Delete** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleDeleteSecretIdBoundCidrs**](AuthApi.md#AppRoleDeleteSecretIdBoundCidrs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleDeleteSecretIdNumUses**](AuthApi.md#AppRoleDeleteSecretIdNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleDeleteSecretIdTtl**](AuthApi.md#AppRoleDeleteSecretIdTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleDeleteTokenBoundCidrs**](AuthApi.md#AppRoleDeleteTokenBoundCidrs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleDeleteTokenMaxTtl**](AuthApi.md#AppRoleDeleteTokenMaxTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleDeleteTokenNumUses**](AuthApi.md#AppRoleDeleteTokenNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleDeleteTokenTtl**](AuthApi.md#AppRoleDeleteTokenTtl) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleDestroySecretId**](AuthApi.md#AppRoleDestroySecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | 
[**AppRoleDestroySecretIdByAccessor**](AuthApi.md#AppRoleDestroySecretIdByAccessor) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**AppRoleListRoles**](AuthApi.md#AppRoleListRoles) | **Get** /auth/{approle_mount_path}/role | 
[**AppRoleListSecretIds**](AuthApi.md#AppRoleListSecretIds) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleLogin**](AuthApi.md#AppRoleLogin) | **Post** /auth/{approle_mount_path}/login | 
[**AppRoleLookUpSecretId**](AuthApi.md#AppRoleLookUpSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/lookup | 
[**AppRoleLookUpSecretIdByAccessor**](AuthApi.md#AppRoleLookUpSecretIdByAccessor) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup | 
[**AppRoleReadBindSecretId**](AuthApi.md#AppRoleReadBindSecretId) | **Get** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleReadBoundCidrList**](AuthApi.md#AppRoleReadBoundCidrList) | **Get** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleReadLocalSecretIds**](AuthApi.md#AppRoleReadLocalSecretIds) | **Get** /auth/{approle_mount_path}/role/{role_name}/local-secret-ids | 
[**AppRoleReadPeriod**](AuthApi.md#AppRoleReadPeriod) | **Get** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleReadPolicies**](AuthApi.md#AppRoleReadPolicies) | **Get** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleReadRole**](AuthApi.md#AppRoleReadRole) | **Get** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleReadRoleId**](AuthApi.md#AppRoleReadRoleId) | **Get** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleReadSecretIdBoundCidrs**](AuthApi.md#AppRoleReadSecretIdBoundCidrs) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleReadSecretIdNumUses**](AuthApi.md#AppRoleReadSecretIdNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleReadSecretIdTtl**](AuthApi.md#AppRoleReadSecretIdTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleReadTokenBoundCidrs**](AuthApi.md#AppRoleReadTokenBoundCidrs) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleReadTokenMaxTtl**](AuthApi.md#AppRoleReadTokenMaxTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleReadTokenNumUses**](AuthApi.md#AppRoleReadTokenNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleReadTokenTtl**](AuthApi.md#AppRoleReadTokenTtl) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleTidySecretId**](AuthApi.md#AppRoleTidySecretId) | **Post** /auth/{approle_mount_path}/tidy/secret-id | 
[**AppRoleWriteBindSecretId**](AuthApi.md#AppRoleWriteBindSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleWriteBoundCidrList**](AuthApi.md#AppRoleWriteBoundCidrList) | **Post** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleWriteCustomSecretId**](AuthApi.md#AppRoleWriteCustomSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/custom-secret-id | 
[**AppRoleWritePeriod**](AuthApi.md#AppRoleWritePeriod) | **Post** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleWritePolicies**](AuthApi.md#AppRoleWritePolicies) | **Post** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleWriteRole**](AuthApi.md#AppRoleWriteRole) | **Post** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleWriteRoleId**](AuthApi.md#AppRoleWriteRoleId) | **Post** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleWriteSecretId**](AuthApi.md#AppRoleWriteSecretId) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleWriteSecretIdBoundCidrs**](AuthApi.md#AppRoleWriteSecretIdBoundCidrs) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleWriteSecretIdNumUses**](AuthApi.md#AppRoleWriteSecretIdNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleWriteSecretIdTtl**](AuthApi.md#AppRoleWriteSecretIdTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleWriteTokenBoundCidrs**](AuthApi.md#AppRoleWriteTokenBoundCidrs) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleWriteTokenMaxTtl**](AuthApi.md#AppRoleWriteTokenMaxTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleWriteTokenNumUses**](AuthApi.md#AppRoleWriteTokenNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleWriteTokenTtl**](AuthApi.md#AppRoleWriteTokenTtl) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AwsConfigureCertificate**](AuthApi.md#AwsConfigureCertificate) | **Post** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AwsConfigureClient**](AuthApi.md#AwsConfigureClient) | **Post** /auth/{aws_mount_path}/config/client | 
[**AwsConfigureIdentityAccessListTidyOperation**](AuthApi.md#AwsConfigureIdentityAccessListTidyOperation) | **Post** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AwsConfigureIdentityIntegration**](AuthApi.md#AwsConfigureIdentityIntegration) | **Post** /auth/{aws_mount_path}/config/identity | 
[**AwsConfigureIdentityWhitelistTidyOperation**](AuthApi.md#AwsConfigureIdentityWhitelistTidyOperation) | **Post** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AwsConfigureRoleTagBlacklistTidyOperation**](AuthApi.md#AwsConfigureRoleTagBlacklistTidyOperation) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AwsConfigureRoleTagDenyListTidyOperation**](AuthApi.md#AwsConfigureRoleTagDenyListTidyOperation) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AwsDeleteAuthRole**](AuthApi.md#AwsDeleteAuthRole) | **Delete** /auth/{aws_mount_path}/role/{role} | 
[**AwsDeleteCertificateConfiguration**](AuthApi.md#AwsDeleteCertificateConfiguration) | **Delete** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AwsDeleteClientConfiguration**](AuthApi.md#AwsDeleteClientConfiguration) | **Delete** /auth/{aws_mount_path}/config/client | 
[**AwsDeleteIdentityAccessList**](AuthApi.md#AwsDeleteIdentityAccessList) | **Delete** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AwsDeleteIdentityAccessListTidySettings**](AuthApi.md#AwsDeleteIdentityAccessListTidySettings) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AwsDeleteIdentityWhitelist**](AuthApi.md#AwsDeleteIdentityWhitelist) | **Delete** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AwsDeleteIdentityWhitelistTidySettings**](AuthApi.md#AwsDeleteIdentityWhitelistTidySettings) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AwsDeleteRoleTagBlacklist**](AuthApi.md#AwsDeleteRoleTagBlacklist) | **Delete** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AwsDeleteRoleTagBlacklistTidySettings**](AuthApi.md#AwsDeleteRoleTagBlacklistTidySettings) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AwsDeleteRoleTagDenyList**](AuthApi.md#AwsDeleteRoleTagDenyList) | **Delete** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AwsDeleteRoleTagDenyListTidySettings**](AuthApi.md#AwsDeleteRoleTagDenyListTidySettings) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AwsDeleteStsRole**](AuthApi.md#AwsDeleteStsRole) | **Delete** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AwsListAuthRoles**](AuthApi.md#AwsListAuthRoles) | **Get** /auth/{aws_mount_path}/role | 
[**AwsListCertificateConfigurations**](AuthApi.md#AwsListCertificateConfigurations) | **Get** /auth/{aws_mount_path}/config/certificates | 
[**AwsListIdentityAccessList**](AuthApi.md#AwsListIdentityAccessList) | **Get** /auth/{aws_mount_path}/identity-accesslist | 
[**AwsListIdentityWhitelist**](AuthApi.md#AwsListIdentityWhitelist) | **Get** /auth/{aws_mount_path}/identity-whitelist | 
[**AwsListRoleTagBlacklists**](AuthApi.md#AwsListRoleTagBlacklists) | **Get** /auth/{aws_mount_path}/roletag-blacklist | 
[**AwsListRoleTagDenyLists**](AuthApi.md#AwsListRoleTagDenyLists) | **Get** /auth/{aws_mount_path}/roletag-denylist | 
[**AwsListStsRoleRelationships**](AuthApi.md#AwsListStsRoleRelationships) | **Get** /auth/{aws_mount_path}/config/sts | 
[**AwsLogin**](AuthApi.md#AwsLogin) | **Post** /auth/{aws_mount_path}/login | 
[**AwsReadAuthRole**](AuthApi.md#AwsReadAuthRole) | **Get** /auth/{aws_mount_path}/role/{role} | 
[**AwsReadCertificateConfiguration**](AuthApi.md#AwsReadCertificateConfiguration) | **Get** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AwsReadClientConfiguration**](AuthApi.md#AwsReadClientConfiguration) | **Get** /auth/{aws_mount_path}/config/client | 
[**AwsReadIdentityAccessList**](AuthApi.md#AwsReadIdentityAccessList) | **Get** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AwsReadIdentityAccessListTidySettings**](AuthApi.md#AwsReadIdentityAccessListTidySettings) | **Get** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AwsReadIdentityIntegrationConfiguration**](AuthApi.md#AwsReadIdentityIntegrationConfiguration) | **Get** /auth/{aws_mount_path}/config/identity | 
[**AwsReadIdentityWhitelist**](AuthApi.md#AwsReadIdentityWhitelist) | **Get** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AwsReadIdentityWhitelistTidySettings**](AuthApi.md#AwsReadIdentityWhitelistTidySettings) | **Get** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AwsReadRoleTagBlacklist**](AuthApi.md#AwsReadRoleTagBlacklist) | **Get** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AwsReadRoleTagBlacklistTidySettings**](AuthApi.md#AwsReadRoleTagBlacklistTidySettings) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AwsReadRoleTagDenyList**](AuthApi.md#AwsReadRoleTagDenyList) | **Get** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AwsReadRoleTagDenyListTidySettings**](AuthApi.md#AwsReadRoleTagDenyListTidySettings) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AwsReadStsRole**](AuthApi.md#AwsReadStsRole) | **Get** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AwsRotateRootCredentials**](AuthApi.md#AwsRotateRootCredentials) | **Post** /auth/{aws_mount_path}/config/rotate-root | 
[**AwsTidyIdentityAccessList**](AuthApi.md#AwsTidyIdentityAccessList) | **Post** /auth/{aws_mount_path}/tidy/identity-accesslist | 
[**AwsTidyIdentityWhitelist**](AuthApi.md#AwsTidyIdentityWhitelist) | **Post** /auth/{aws_mount_path}/tidy/identity-whitelist | 
[**AwsTidyRoleTagBlacklist**](AuthApi.md#AwsTidyRoleTagBlacklist) | **Post** /auth/{aws_mount_path}/tidy/roletag-blacklist | 
[**AwsTidyRoleTagDenyList**](AuthApi.md#AwsTidyRoleTagDenyList) | **Post** /auth/{aws_mount_path}/tidy/roletag-denylist | 
[**AwsWriteAuthRole**](AuthApi.md#AwsWriteAuthRole) | **Post** /auth/{aws_mount_path}/role/{role} | 
[**AwsWriteRoleTag**](AuthApi.md#AwsWriteRoleTag) | **Post** /auth/{aws_mount_path}/role/{role}/tag | 
[**AwsWriteRoleTagBlacklist**](AuthApi.md#AwsWriteRoleTagBlacklist) | **Post** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AwsWriteRoleTagDenyList**](AuthApi.md#AwsWriteRoleTagDenyList) | **Post** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AwsWriteStsRole**](AuthApi.md#AwsWriteStsRole) | **Post** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AzureConfigureAuth**](AuthApi.md#AzureConfigureAuth) | **Post** /auth/{azure_mount_path}/config | 
[**AzureDeleteAuthConfiguration**](AuthApi.md#AzureDeleteAuthConfiguration) | **Delete** /auth/{azure_mount_path}/config | 
[**AzureDeleteAuthRole**](AuthApi.md#AzureDeleteAuthRole) | **Delete** /auth/{azure_mount_path}/role/{name} | 
[**AzureListAuthRoles**](AuthApi.md#AzureListAuthRoles) | **Get** /auth/{azure_mount_path}/role | 
[**AzureLogin**](AuthApi.md#AzureLogin) | **Post** /auth/{azure_mount_path}/login | 
[**AzureReadAuthConfiguration**](AuthApi.md#AzureReadAuthConfiguration) | **Get** /auth/{azure_mount_path}/config | 
[**AzureReadAuthRole**](AuthApi.md#AzureReadAuthRole) | **Get** /auth/{azure_mount_path}/role/{name} | 
[**AzureRotateRootCredentials**](AuthApi.md#AzureRotateRootCredentials) | **Post** /auth/{azure_mount_path}/rotate-root | 
[**AzureWriteAuthRole**](AuthApi.md#AzureWriteAuthRole) | **Post** /auth/{azure_mount_path}/role/{name} | 
[**CentrifyConfigure**](AuthApi.md#CentrifyConfigure) | **Post** /auth/{centrify_mount_path}/config | 
[**CentrifyLogin**](AuthApi.md#CentrifyLogin) | **Post** /auth/{centrify_mount_path}/login | Log in with a username and password.
[**CentrifyReadConfiguration**](AuthApi.md#CentrifyReadConfiguration) | **Get** /auth/{centrify_mount_path}/config | 
[**CertConfigure**](AuthApi.md#CertConfigure) | **Post** /auth/{cert_mount_path}/config | 
[**CertDeleteCertificate**](AuthApi.md#CertDeleteCertificate) | **Delete** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertDeleteCrl**](AuthApi.md#CertDeleteCrl) | **Delete** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertListCertificates**](AuthApi.md#CertListCertificates) | **Get** /auth/{cert_mount_path}/certs | Manage trusted certificates used for authentication.
[**CertListCrls**](AuthApi.md#CertListCrls) | **Get** /auth/{cert_mount_path}/crls | 
[**CertLogin**](AuthApi.md#CertLogin) | **Post** /auth/{cert_mount_path}/login | 
[**CertReadCertificate**](AuthApi.md#CertReadCertificate) | **Get** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertReadConfiguration**](AuthApi.md#CertReadConfiguration) | **Get** /auth/{cert_mount_path}/config | 
[**CertReadCrl**](AuthApi.md#CertReadCrl) | **Get** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertWriteCertificate**](AuthApi.md#CertWriteCertificate) | **Post** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertWriteCrl**](AuthApi.md#CertWriteCrl) | **Post** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CloudFoundryConfigure**](AuthApi.md#CloudFoundryConfigure) | **Post** /auth/{cf_mount_path}/config | 
[**CloudFoundryDeleteConfiguration**](AuthApi.md#CloudFoundryDeleteConfiguration) | **Delete** /auth/{cf_mount_path}/config | 
[**CloudFoundryDeleteRole**](AuthApi.md#CloudFoundryDeleteRole) | **Delete** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryListRoles**](AuthApi.md#CloudFoundryListRoles) | **Get** /auth/{cf_mount_path}/roles | 
[**CloudFoundryLogin**](AuthApi.md#CloudFoundryLogin) | **Post** /auth/{cf_mount_path}/login | 
[**CloudFoundryReadConfiguration**](AuthApi.md#CloudFoundryReadConfiguration) | **Get** /auth/{cf_mount_path}/config | 
[**CloudFoundryReadRole**](AuthApi.md#CloudFoundryReadRole) | **Get** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryWriteRole**](AuthApi.md#CloudFoundryWriteRole) | **Post** /auth/{cf_mount_path}/roles/{role} | 
[**GithubConfigure**](AuthApi.md#GithubConfigure) | **Post** /auth/{github_mount_path}/config | 
[**GithubDeleteTeamMapping**](AuthApi.md#GithubDeleteTeamMapping) | **Delete** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GithubDeleteUserMapping**](AuthApi.md#GithubDeleteUserMapping) | **Delete** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GithubLogin**](AuthApi.md#GithubLogin) | **Post** /auth/{github_mount_path}/login | 
[**GithubReadConfiguration**](AuthApi.md#GithubReadConfiguration) | **Get** /auth/{github_mount_path}/config | 
[**GithubReadTeamMapping**](AuthApi.md#GithubReadTeamMapping) | **Get** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GithubReadTeams**](AuthApi.md#GithubReadTeams) | **Get** /auth/{github_mount_path}/map/teams | Read mappings for teams
[**GithubReadUserMapping**](AuthApi.md#GithubReadUserMapping) | **Get** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GithubReadUsers**](AuthApi.md#GithubReadUsers) | **Get** /auth/{github_mount_path}/map/users | Read mappings for users
[**GithubWriteTeamMapping**](AuthApi.md#GithubWriteTeamMapping) | **Post** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GithubWriteUserMapping**](AuthApi.md#GithubWriteUserMapping) | **Post** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GoogleCloudConfigureAuth**](AuthApi.md#GoogleCloudConfigureAuth) | **Post** /auth/{gcp_mount_path}/config | 
[**GoogleCloudDeleteRole**](AuthApi.md#GoogleCloudDeleteRole) | **Delete** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudEditLabelsForRole**](AuthApi.md#GoogleCloudEditLabelsForRole) | **Post** /auth/{gcp_mount_path}/role/{name}/labels | Add or remove labels for an existing &#x27;gce&#x27; role
[**GoogleCloudEditServiceAccountsForRole**](AuthApi.md#GoogleCloudEditServiceAccountsForRole) | **Post** /auth/{gcp_mount_path}/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**GoogleCloudListRoles**](AuthApi.md#GoogleCloudListRoles) | **Get** /auth/{gcp_mount_path}/role | Lists all the roles that are registered with Vault.
[**GoogleCloudLogin**](AuthApi.md#GoogleCloudLogin) | **Post** /auth/{gcp_mount_path}/login | 
[**GoogleCloudReadAuthConfiguration**](AuthApi.md#GoogleCloudReadAuthConfiguration) | **Get** /auth/{gcp_mount_path}/config | 
[**GoogleCloudReadRole**](AuthApi.md#GoogleCloudReadRole) | **Get** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudWriteRole**](AuthApi.md#GoogleCloudWriteRole) | **Post** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**JwtConfigure**](AuthApi.md#JwtConfigure) | **Post** /auth/{jwt_mount_path}/config | Configure the JWT authentication backend.
[**JwtDeleteRole**](AuthApi.md#JwtDeleteRole) | **Delete** /auth/{jwt_mount_path}/role/{name} | Delete an existing role.
[**JwtListRoles**](AuthApi.md#JwtListRoles) | **Get** /auth/{jwt_mount_path}/role | Lists all the roles registered with the backend.
[**JwtLogin**](AuthApi.md#JwtLogin) | **Post** /auth/{jwt_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**JwtOidcCallback**](AuthApi.md#JwtOidcCallback) | **Get** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**JwtOidcCallbackWithParameters**](AuthApi.md#JwtOidcCallbackWithParameters) | **Post** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**JwtOidcRequestAuthorizationUrl**](AuthApi.md#JwtOidcRequestAuthorizationUrl) | **Post** /auth/{jwt_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**JwtReadConfiguration**](AuthApi.md#JwtReadConfiguration) | **Get** /auth/{jwt_mount_path}/config | Read the current JWT authentication backend configuration.
[**JwtReadRole**](AuthApi.md#JwtReadRole) | **Get** /auth/{jwt_mount_path}/role/{name} | Read an existing role.
[**JwtWriteRole**](AuthApi.md#JwtWriteRole) | **Post** /auth/{jwt_mount_path}/role/{name} | Register an role with the backend.
[**KerberosConfigure**](AuthApi.md#KerberosConfigure) | **Post** /auth/{kerberos_mount_path}/config | 
[**KerberosConfigureLdap**](AuthApi.md#KerberosConfigureLdap) | **Post** /auth/{kerberos_mount_path}/config/ldap | 
[**KerberosDeleteGroup**](AuthApi.md#KerberosDeleteGroup) | **Delete** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosListGroups**](AuthApi.md#KerberosListGroups) | **Get** /auth/{kerberos_mount_path}/groups | 
[**KerberosLogin**](AuthApi.md#KerberosLogin) | **Post** /auth/{kerberos_mount_path}/login | 
[**KerberosReadConfiguration**](AuthApi.md#KerberosReadConfiguration) | **Get** /auth/{kerberos_mount_path}/config | 
[**KerberosReadGroup**](AuthApi.md#KerberosReadGroup) | **Get** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosReadLdapConfiguration**](AuthApi.md#KerberosReadLdapConfiguration) | **Get** /auth/{kerberos_mount_path}/config/ldap | 
[**KerberosWriteGroup**](AuthApi.md#KerberosWriteGroup) | **Post** /auth/{kerberos_mount_path}/groups/{name} | 
[**KubernetesConfigureAuth**](AuthApi.md#KubernetesConfigureAuth) | **Post** /auth/{kubernetes_mount_path}/config | 
[**KubernetesDeleteAuthRole**](AuthApi.md#KubernetesDeleteAuthRole) | **Delete** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesListAuthRoles**](AuthApi.md#KubernetesListAuthRoles) | **Get** /auth/{kubernetes_mount_path}/role | Lists all the roles registered with the backend.
[**KubernetesLogin**](AuthApi.md#KubernetesLogin) | **Post** /auth/{kubernetes_mount_path}/login | Authenticates Kubernetes service accounts with Vault.
[**KubernetesReadAuthConfiguration**](AuthApi.md#KubernetesReadAuthConfiguration) | **Get** /auth/{kubernetes_mount_path}/config | 
[**KubernetesReadAuthRole**](AuthApi.md#KubernetesReadAuthRole) | **Get** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesWriteAuthRole**](AuthApi.md#KubernetesWriteAuthRole) | **Post** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**LdapConfigureAuth**](AuthApi.md#LdapConfigureAuth) | **Post** /auth/{ldap_mount_path}/config | 
[**LdapDeleteGroup**](AuthApi.md#LdapDeleteGroup) | **Delete** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LdapDeleteUser**](AuthApi.md#LdapDeleteUser) | **Delete** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LdapListGroups**](AuthApi.md#LdapListGroups) | **Get** /auth/{ldap_mount_path}/groups | Manage additional groups for users allowed to authenticate.
[**LdapListUsers**](AuthApi.md#LdapListUsers) | **Get** /auth/{ldap_mount_path}/users | Manage users allowed to authenticate.
[**LdapLogin**](AuthApi.md#LdapLogin) | **Post** /auth/{ldap_mount_path}/login/{username} | Log in with a username and password.
[**LdapReadAuthConfiguration**](AuthApi.md#LdapReadAuthConfiguration) | **Get** /auth/{ldap_mount_path}/config | 
[**LdapReadGroup**](AuthApi.md#LdapReadGroup) | **Get** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LdapReadUser**](AuthApi.md#LdapReadUser) | **Get** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LdapWriteGroup**](AuthApi.md#LdapWriteGroup) | **Post** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LdapWriteUser**](AuthApi.md#LdapWriteUser) | **Post** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**OciConfigure**](AuthApi.md#OciConfigure) | **Post** /auth/{oci_mount_path}/config | 
[**OciDeleteConfiguration**](AuthApi.md#OciDeleteConfiguration) | **Delete** /auth/{oci_mount_path}/config | 
[**OciDeleteRole**](AuthApi.md#OciDeleteRole) | **Delete** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OciListRoles**](AuthApi.md#OciListRoles) | **Get** /auth/{oci_mount_path}/role | Lists all the roles that are registered with Vault.
[**OciLogin**](AuthApi.md#OciLogin) | **Post** /auth/{oci_mount_path}/login/{role} | Authenticates to Vault using OCI credentials
[**OciReadConfiguration**](AuthApi.md#OciReadConfiguration) | **Get** /auth/{oci_mount_path}/config | 
[**OciReadRole**](AuthApi.md#OciReadRole) | **Get** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OciWriteRole**](AuthApi.md#OciWriteRole) | **Post** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OktaConfigure**](AuthApi.md#OktaConfigure) | **Post** /auth/{okta_mount_path}/config | 
[**OktaDeleteGroup**](AuthApi.md#OktaDeleteGroup) | **Delete** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaDeleteUser**](AuthApi.md#OktaDeleteUser) | **Delete** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaListGroups**](AuthApi.md#OktaListGroups) | **Get** /auth/{okta_mount_path}/groups | Manage users allowed to authenticate.
[**OktaListUsers**](AuthApi.md#OktaListUsers) | **Get** /auth/{okta_mount_path}/users | Manage additional groups for users allowed to authenticate.
[**OktaLogin**](AuthApi.md#OktaLogin) | **Post** /auth/{okta_mount_path}/login/{username} | Log in with a username and password.
[**OktaReadConfiguration**](AuthApi.md#OktaReadConfiguration) | **Get** /auth/{okta_mount_path}/config | 
[**OktaReadGroup**](AuthApi.md#OktaReadGroup) | **Get** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaReadUser**](AuthApi.md#OktaReadUser) | **Get** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaVerify**](AuthApi.md#OktaVerify) | **Get** /auth/{okta_mount_path}/verify/{nonce} | 
[**OktaWriteGroup**](AuthApi.md#OktaWriteGroup) | **Post** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaWriteUser**](AuthApi.md#OktaWriteUser) | **Post** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**RadiusConfigure**](AuthApi.md#RadiusConfigure) | **Post** /auth/{radius_mount_path}/config | 
[**RadiusDeleteUser**](AuthApi.md#RadiusDeleteUser) | **Delete** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusListUsers**](AuthApi.md#RadiusListUsers) | **Get** /auth/{radius_mount_path}/users | Manage users allowed to authenticate.
[**RadiusLogin**](AuthApi.md#RadiusLogin) | **Post** /auth/{radius_mount_path}/login | Log in with a username and password.
[**RadiusLoginWithUsername**](AuthApi.md#RadiusLoginWithUsername) | **Post** /auth/{radius_mount_path}/login/{urlusername} | Log in with a username and password.
[**RadiusReadConfiguration**](AuthApi.md#RadiusReadConfiguration) | **Get** /auth/{radius_mount_path}/config | 
[**RadiusReadUser**](AuthApi.md#RadiusReadUser) | **Get** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusWriteUser**](AuthApi.md#RadiusWriteUser) | **Post** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**TokenCreate**](AuthApi.md#TokenCreate) | **Post** /auth/token/create | The token create path is used to create new tokens.
[**TokenCreateAgainstRole**](AuthApi.md#TokenCreateAgainstRole) | **Post** /auth/token/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**TokenCreateOrphan**](AuthApi.md#TokenCreateOrphan) | **Post** /auth/token/create-orphan | The token create path is used to create new orphan tokens.
[**TokenDeleteRole**](AuthApi.md#TokenDeleteRole) | **Delete** /auth/token/roles/{role_name} | 
[**TokenListAccessors**](AuthApi.md#TokenListAccessors) | **Get** /auth/token/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#x27;sudo&#x27; capability in addition to &#x27;list&#x27;.
[**TokenListRoles**](AuthApi.md#TokenListRoles) | **Get** /auth/token/roles | This endpoint lists configured roles.
[**TokenLookUp**](AuthApi.md#TokenLookUp) | **Post** /auth/token/lookup | 
[**TokenLookUpAccessor**](AuthApi.md#TokenLookUpAccessor) | **Post** /auth/token/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenLookUpSelf**](AuthApi.md#TokenLookUpSelf) | **Get** /auth/token/lookup-self | 
[**TokenReadRole**](AuthApi.md#TokenReadRole) | **Get** /auth/token/roles/{role_name} | 
[**TokenRenew**](AuthApi.md#TokenRenew) | **Post** /auth/token/renew | This endpoint will renew the given token and prevent expiration.
[**TokenRenewAccessor**](AuthApi.md#TokenRenewAccessor) | **Post** /auth/token/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenRenewSelf**](AuthApi.md#TokenRenewSelf) | **Post** /auth/token/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**TokenRevoke**](AuthApi.md#TokenRevoke) | **Post** /auth/token/revoke | This endpoint will delete the given token and all of its child tokens.
[**TokenRevokeAccessor**](AuthApi.md#TokenRevokeAccessor) | **Post** /auth/token/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**TokenRevokeOrphan**](AuthApi.md#TokenRevokeOrphan) | **Post** /auth/token/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**TokenRevokeSelf**](AuthApi.md#TokenRevokeSelf) | **Post** /auth/token/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**TokenTidy**](AuthApi.md#TokenTidy) | **Post** /auth/token/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**TokenWriteRole**](AuthApi.md#TokenWriteRole) | **Post** /auth/token/roles/{role_name} | 
[**UserpassDeleteUser**](AuthApi.md#UserpassDeleteUser) | **Delete** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassListUsers**](AuthApi.md#UserpassListUsers) | **Get** /auth/{userpass_mount_path}/users | Manage users allowed to authenticate.
[**UserpassLogin**](AuthApi.md#UserpassLogin) | **Post** /auth/{userpass_mount_path}/login/{username} | Log in with a username and password.
[**UserpassReadUser**](AuthApi.md#UserpassReadUser) | **Get** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassResetPassword**](AuthApi.md#UserpassResetPassword) | **Post** /auth/{userpass_mount_path}/users/{username}/password | Reset user&#x27;s password.
[**UserpassUpdatePolicies**](AuthApi.md#UserpassUpdatePolicies) | **Post** /auth/{userpass_mount_path}/users/{username}/policies | Update the policies associated with the username.
[**UserpassWriteUser**](AuthApi.md#UserpassWriteUser) | **Post** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.





## AliCloudDeleteAuthRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	resp, err := client.Auth.AliCloudDeleteAuthRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("alicloud"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AliCloudListAuthRoles

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AliCloudListAuthRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("alicloud"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudLogin

Authenticates an RAM entity with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAliCloudLoginRequestWithDefaults()
	resp, err := client.Auth.AliCloudLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("alicloud"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliCloudLoginRequest** | [**AliCloudLoginRequest**](AliCloudLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudReadAuthRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	resp, err := client.Auth.AliCloudReadAuthRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("alicloud"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AliCloudWriteAuthRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.
	request := schema.NewAliCloudWriteAuthRoleRequestWithDefaults()
	resp, err := client.Auth.AliCloudWriteAuthRole(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("alicloud"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **aliCloudWriteAuthRoleRequest** | [**AliCloudWriteAuthRoleRequest**](AliCloudWriteAuthRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteBindSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteBindSecretId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteBoundCidrList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteBoundCidrList(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeletePeriod



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeletePeriod(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeletePolicies



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeletePolicies(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteRole(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIdBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteSecretIdBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIdNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteSecretIdNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIdTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteSecretIdTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteTokenBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenMaxTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteTokenMaxTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteTokenNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleDeleteTokenTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDestroySecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleDestroySecretIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleDestroySecretId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleDestroySecretIdRequest** | [**AppRoleDestroySecretIdRequest**](AppRoleDestroySecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDestroySecretIdByAccessor



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleDestroySecretIdByAccessorRequestWithDefaults()
	resp, err := client.Auth.AppRoleDestroySecretIdByAccessor(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleDestroySecretIdByAccessorRequest** | [**AppRoleDestroySecretIdByAccessorRequest**](AppRoleDestroySecretIdByAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AppRoleListRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**AppRoleListRolesResponse**](AppRoleListRolesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleListSecretIds



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleListSecretIds(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**AppRoleListSecretIdsResponse**](AppRoleListSecretIdsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAppRoleLoginRequestWithDefaults()
	resp, err := client.Auth.AppRoleLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleLoginRequest** | [**AppRoleLoginRequest**](AppRoleLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleLookUpSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleLookUpSecretIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleLookUpSecretId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleLookUpSecretIdRequest** | [**AppRoleLookUpSecretIdRequest**](AppRoleLookUpSecretIdRequest.md) |  | 

[**AppRoleLookUpSecretIdResponse**](AppRoleLookUpSecretIdResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleLookUpSecretIdByAccessor



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleLookUpSecretIdByAccessorRequestWithDefaults()
	resp, err := client.Auth.AppRoleLookUpSecretIdByAccessor(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleLookUpSecretIdByAccessorRequest** | [**AppRoleLookUpSecretIdByAccessorRequest**](AppRoleLookUpSecretIdByAccessorRequest.md) |  | 

[**AppRoleLookUpSecretIdByAccessorResponse**](AppRoleLookUpSecretIdByAccessorResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadBindSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadBindSecretId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadBindSecretIdResponse**](AppRoleReadBindSecretIdResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadBoundCidrList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadBoundCidrList(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadBoundCidrListResponse**](AppRoleReadBoundCidrListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadLocalSecretIds



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadLocalSecretIds(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadLocalSecretIdsResponse**](AppRoleReadLocalSecretIdsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadPeriod



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadPeriod(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadPeriodResponse**](AppRoleReadPeriodResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadPolicies



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadPolicies(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadPoliciesResponse**](AppRoleReadPoliciesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadRole(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadRoleResponse**](AppRoleReadRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadRoleId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadRoleId(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadRoleIdResponse**](AppRoleReadRoleIdResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIdBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadSecretIdBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIdBoundCidrsResponse**](AppRoleReadSecretIdBoundCidrsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIdNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadSecretIdNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIdNumUsesResponse**](AppRoleReadSecretIdNumUsesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIdTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadSecretIdTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIdTtlResponse**](AppRoleReadSecretIdTtlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadTokenBoundCidrs(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenBoundCidrsResponse**](AppRoleReadTokenBoundCidrsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenMaxTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadTokenMaxTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenMaxTtlResponse**](AppRoleReadTokenMaxTtlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadTokenNumUses(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenNumUsesResponse**](AppRoleReadTokenNumUsesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	resp, err := client.Auth.AppRoleReadTokenTtl(
		context.Background(),
		roleName,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenTtlResponse**](AppRoleReadTokenTtlResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleTidySecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AppRoleTidySecretId(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteBindSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteBindSecretIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteBindSecretId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteBindSecretIdRequest** | [**AppRoleWriteBindSecretIdRequest**](AppRoleWriteBindSecretIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteBoundCidrList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteBoundCidrListRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteBoundCidrList(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteBoundCidrListRequest** | [**AppRoleWriteBoundCidrListRequest**](AppRoleWriteBoundCidrListRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteCustomSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteCustomSecretIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteCustomSecretId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteCustomSecretIdRequest** | [**AppRoleWriteCustomSecretIdRequest**](AppRoleWriteCustomSecretIdRequest.md) |  | 

[**AppRoleWriteCustomSecretIdResponse**](AppRoleWriteCustomSecretIdResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWritePeriod



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWritePeriodRequestWithDefaults()
	resp, err := client.Auth.AppRoleWritePeriod(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWritePeriodRequest** | [**AppRoleWritePeriodRequest**](AppRoleWritePeriodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWritePolicies



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWritePoliciesRequestWithDefaults()
	resp, err := client.Auth.AppRoleWritePolicies(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWritePoliciesRequest** | [**AppRoleWritePoliciesRequest**](AppRoleWritePoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteRoleRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteRole(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteRoleRequest** | [**AppRoleWriteRoleRequest**](AppRoleWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteRoleId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteRoleIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteRoleId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteRoleIdRequest** | [**AppRoleWriteRoleIdRequest**](AppRoleWriteRoleIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretId



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteSecretIdRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteSecretId(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteSecretIdRequest** | [**AppRoleWriteSecretIdRequest**](AppRoleWriteSecretIdRequest.md) |  | 

[**AppRoleWriteSecretIdResponse**](AppRoleWriteSecretIdResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIdBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteSecretIdBoundCidrsRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteSecretIdBoundCidrs(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteSecretIdBoundCidrsRequest** | [**AppRoleWriteSecretIdBoundCidrsRequest**](AppRoleWriteSecretIdBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIdNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteSecretIdNumUsesRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteSecretIdNumUses(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteSecretIdNumUsesRequest** | [**AppRoleWriteSecretIdNumUsesRequest**](AppRoleWriteSecretIdNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIdTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteSecretIdTtlRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteSecretIdTtl(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteSecretIdTtlRequest** | [**AppRoleWriteSecretIdTtlRequest**](AppRoleWriteSecretIdTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenBoundCidrs



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteTokenBoundCidrsRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteTokenBoundCidrs(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteTokenBoundCidrsRequest** | [**AppRoleWriteTokenBoundCidrsRequest**](AppRoleWriteTokenBoundCidrsRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenMaxTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteTokenMaxTtlRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteTokenMaxTtl(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteTokenMaxTtlRequest** | [**AppRoleWriteTokenMaxTtlRequest**](AppRoleWriteTokenMaxTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenNumUses



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteTokenNumUsesRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteTokenNumUses(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteTokenNumUsesRequest** | [**AppRoleWriteTokenNumUsesRequest**](AppRoleWriteTokenNumUsesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenTtl



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.
	request := schema.NewAppRoleWriteTokenTtlRequestWithDefaults()
	resp, err := client.Auth.AppRoleWriteTokenTtl(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("approle"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRoleWriteTokenTtlRequest** | [**AppRoleWriteTokenTtlRequest**](AppRoleWriteTokenTtlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureCertificate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	request := schema.NewAwsConfigureCertificateRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureCertificate(
		context.Background(),
		certName,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsConfigureCertificateRequest** | [**AwsConfigureCertificateRequest**](AwsConfigureCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureClientRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureClient(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureClientRequest** | [**AwsConfigureClientRequest**](AwsConfigureClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureIdentityAccessListTidyOperation



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureIdentityAccessListTidyOperation(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureIdentityAccessListTidyOperationRequest** | [**AwsConfigureIdentityAccessListTidyOperationRequest**](AwsConfigureIdentityAccessListTidyOperationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureIdentityIntegration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureIdentityIntegrationRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureIdentityIntegration(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureIdentityIntegrationRequest** | [**AwsConfigureIdentityIntegrationRequest**](AwsConfigureIdentityIntegrationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureIdentityWhitelistTidyOperation



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureIdentityWhitelistTidyOperationRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureIdentityWhitelistTidyOperation(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureIdentityWhitelistTidyOperationRequest** | [**AwsConfigureIdentityWhitelistTidyOperationRequest**](AwsConfigureIdentityWhitelistTidyOperationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureRoleTagBlacklistTidyOperation



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureRoleTagBlacklistTidyOperationRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureRoleTagBlacklistTidyOperation(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureRoleTagBlacklistTidyOperationRequest** | [**AwsConfigureRoleTagBlacklistTidyOperationRequest**](AwsConfigureRoleTagBlacklistTidyOperationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsConfigureRoleTagDenyListTidyOperation



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsConfigureRoleTagDenyListTidyOperationRequestWithDefaults()
	resp, err := client.Auth.AwsConfigureRoleTagDenyListTidyOperation(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsConfigureRoleTagDenyListTidyOperationRequest** | [**AwsConfigureRoleTagDenyListTidyOperationRequest**](AwsConfigureRoleTagDenyListTidyOperationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	resp, err := client.Auth.AwsDeleteAuthRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsDeleteCertificateConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	resp, err := client.Auth.AwsDeleteCertificateConfiguration(
		context.Background(),
		certName,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteClientConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsDeleteClientConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteIdentityAccessList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	resp, err := client.Auth.AwsDeleteIdentityAccessList(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteIdentityAccessListTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsDeleteIdentityAccessListTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteIdentityWhitelist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	resp, err := client.Auth.AwsDeleteIdentityWhitelist(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteIdentityWhitelistTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsDeleteIdentityWhitelistTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteRoleTagBlacklist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsDeleteRoleTagBlacklist(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsDeleteRoleTagBlacklistTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsDeleteRoleTagBlacklistTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteRoleTagDenyList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsDeleteRoleTagDenyList(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsDeleteRoleTagDenyListTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsDeleteRoleTagDenyListTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsDeleteStsRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	resp, err := client.Auth.AwsDeleteStsRole(
		context.Background(),
		accountId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListAuthRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListAuthRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListCertificateConfigurations



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListCertificateConfigurations(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListIdentityAccessList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListIdentityAccessList(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListIdentityWhitelist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListIdentityWhitelist(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListRoleTagBlacklists



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListRoleTagBlacklists(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListRoleTagDenyLists



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListRoleTagDenyLists(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsListStsRoleRelationships



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsListStsRoleRelationships(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsLoginRequestWithDefaults()
	resp, err := client.Auth.AwsLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsLoginRequest** | [**AwsLoginRequest**](AwsLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	resp, err := client.Auth.AwsReadAuthRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsReadCertificateConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.
	resp, err := client.Auth.AwsReadCertificateConfiguration(
		context.Background(),
		certName,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadClientConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadClientConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadIdentityAccessList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	resp, err := client.Auth.AwsReadIdentityAccessList(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadIdentityAccessListTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadIdentityAccessListTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadIdentityIntegrationConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadIdentityIntegrationConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadIdentityWhitelist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
	resp, err := client.Auth.AwsReadIdentityWhitelist(
		context.Background(),
		instanceId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadIdentityWhitelistTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadIdentityWhitelistTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadRoleTagBlacklist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsReadRoleTagBlacklist(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsReadRoleTagBlacklistTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadRoleTagBlacklistTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadRoleTagDenyList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsReadRoleTagDenyList(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsReadRoleTagDenyListTidySettings



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsReadRoleTagDenyListTidySettings(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsReadStsRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	resp, err := client.Auth.AwsReadStsRole(
		context.Background(),
		accountId,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsRotateRootCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AwsRotateRootCredentials(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsTidyIdentityAccessList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsTidyIdentityAccessListRequestWithDefaults()
	resp, err := client.Auth.AwsTidyIdentityAccessList(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsTidyIdentityAccessListRequest** | [**AwsTidyIdentityAccessListRequest**](AwsTidyIdentityAccessListRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsTidyIdentityWhitelist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsTidyIdentityWhitelistRequestWithDefaults()
	resp, err := client.Auth.AwsTidyIdentityWhitelist(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsTidyIdentityWhitelistRequest** | [**AwsTidyIdentityWhitelistRequest**](AwsTidyIdentityWhitelistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsTidyRoleTagBlacklist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsTidyRoleTagBlacklistRequestWithDefaults()
	resp, err := client.Auth.AwsTidyRoleTagBlacklist(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsTidyRoleTagBlacklistRequest** | [**AwsTidyRoleTagBlacklistRequest**](AwsTidyRoleTagBlacklistRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsTidyRoleTagDenyList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAwsTidyRoleTagDenyListRequestWithDefaults()
	resp, err := client.Auth.AwsTidyRoleTagDenyList(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsTidyRoleTagDenyListRequest** | [**AwsTidyRoleTagDenyListRequest**](AwsTidyRoleTagDenyListRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsWriteAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	request := schema.NewAwsWriteAuthRoleRequestWithDefaults()
	resp, err := client.Auth.AwsWriteAuthRole(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **awsWriteAuthRoleRequest** | [**AwsWriteAuthRoleRequest**](AwsWriteAuthRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsWriteRoleTag



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	request := schema.NewAwsWriteRoleTagRequestWithDefaults()
	resp, err := client.Auth.AwsWriteRoleTag(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **awsWriteRoleTagRequest** | [**AwsWriteRoleTagRequest**](AwsWriteRoleTagRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AwsWriteRoleTagBlacklist



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsWriteRoleTagBlacklist(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsWriteRoleTagDenyList



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
	resp, err := client.Auth.AwsWriteRoleTagDenyList(
		context.Background(),
		roleTag,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## AwsWriteStsRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
	request := schema.NewAwsWriteStsRoleRequestWithDefaults()
	resp, err := client.Auth.AwsWriteStsRole(
		context.Background(),
		accountId,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("aws"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsWriteStsRoleRequest** | [**AwsWriteStsRoleRequest**](AwsWriteStsRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureConfigureAuth



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAzureConfigureAuthRequestWithDefaults()
	resp, err := client.Auth.AzureConfigureAuth(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureConfigureAuthRequest** | [**AzureConfigureAuthRequest**](AzureConfigureAuthRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteAuthConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AzureDeleteAuthConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.AzureDeleteAuthRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureListAuthRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AzureListAuthRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAzureLoginRequestWithDefaults()
	resp, err := client.Auth.AzureLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureLoginRequest** | [**AzureLoginRequest**](AzureLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadAuthConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AzureReadAuthConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.AzureReadAuthRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureRotateRootCredentials



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.AzureRotateRootCredentials(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteAuthRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewAzureWriteAuthRoleRequestWithDefaults()
	resp, err := client.Auth.AzureWriteAuthRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("azure"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureWriteAuthRoleRequest** | [**AzureWriteAuthRoleRequest**](AzureWriteAuthRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CentrifyConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCentrifyConfigureRequestWithDefaults()
	resp, err := client.Auth.CentrifyConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("centrify"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **centrifyConfigureRequest** | [**CentrifyConfigureRequest**](CentrifyConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CentrifyLogin

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCentrifyLoginRequestWithDefaults()
	resp, err := client.Auth.CentrifyLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("centrify"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **centrifyLoginRequest** | [**CentrifyLoginRequest**](CentrifyLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CentrifyReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CentrifyReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("centrify"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCertConfigureRequestWithDefaults()
	resp, err := client.Auth.CertConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certConfigureRequest** | [**CertConfigureRequest**](CertConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertDeleteCertificate

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	resp, err := client.Auth.CertDeleteCertificate(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CertDeleteCrl

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	resp, err := client.Auth.CertDeleteCrl(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CertListCertificates

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CertListCertificates(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertListCrls



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CertListCrls(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCertLoginRequestWithDefaults()
	resp, err := client.Auth.CertLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certLoginRequest** | [**CertLoginRequest**](CertLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertReadCertificate

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	resp, err := client.Auth.CertReadCertificate(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CertReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CertReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertReadCrl

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	resp, err := client.Auth.CertReadCrl(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CertWriteCertificate

Manage trusted certificates used for authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	request := schema.NewCertWriteCertificateRequestWithDefaults()
	resp, err := client.Auth.CertWriteCertificate(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **certWriteCertificateRequest** | [**CertWriteCertificateRequest**](CertWriteCertificateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertWriteCrl

Manage Certificate Revocation Lists checked during authentication.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate
	request := schema.NewCertWriteCrlRequestWithDefaults()
	resp, err := client.Auth.CertWriteCrl(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cert"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **certWriteCrlRequest** | [**CertWriteCrlRequest**](CertWriteCrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCloudFoundryConfigureRequestWithDefaults()
	resp, err := client.Auth.CloudFoundryConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudFoundryConfigureRequest** | [**CloudFoundryConfigureRequest**](CloudFoundryConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CloudFoundryDeleteConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	resp, err := client.Auth.CloudFoundryDeleteRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CloudFoundryListRoles



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CloudFoundryListRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewCloudFoundryLoginRequestWithDefaults()
	resp, err := client.Auth.CloudFoundryLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudFoundryLoginRequest** | [**CloudFoundryLoginRequest**](CloudFoundryLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.CloudFoundryReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	resp, err := client.Auth.CloudFoundryReadRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## CloudFoundryWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.
	request := schema.NewCloudFoundryWriteRoleRequestWithDefaults()
	resp, err := client.Auth.CloudFoundryWriteRole(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("cf"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **cloudFoundryWriteRoleRequest** | [**CloudFoundryWriteRoleRequest**](CloudFoundryWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGithubConfigureRequestWithDefaults()
	resp, err := client.Auth.GithubConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubConfigureRequest** | [**GithubConfigureRequest**](GithubConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubDeleteTeamMapping

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	resp, err := client.Auth.GithubDeleteTeamMapping(
		context.Background(),
		key,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubDeleteUserMapping

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	resp, err := client.Auth.GithubDeleteUserMapping(
		context.Background(),
		key,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGithubLoginRequestWithDefaults()
	resp, err := client.Auth.GithubLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubLoginRequest** | [**GithubLoginRequest**](GithubLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.GithubReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubReadTeamMapping

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	resp, err := client.Auth.GithubReadTeamMapping(
		context.Background(),
		key,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubReadTeams

Read mappings for teams

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.GithubReadTeams(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubReadUserMapping

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	resp, err := client.Auth.GithubReadUserMapping(
		context.Background(),
		key,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubReadUsers

Read mappings for users

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.GithubReadUsers(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubWriteTeamMapping

Read/write/delete a single teams mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping
	request := schema.NewGithubWriteTeamMappingRequestWithDefaults()
	resp, err := client.Auth.GithubWriteTeamMapping(
		context.Background(),
		key,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **githubWriteTeamMappingRequest** | [**GithubWriteTeamMappingRequest**](GithubWriteTeamMappingRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GithubWriteUserMapping

Read/write/delete a single users mapping

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping
	request := schema.NewGithubWriteUserMappingRequestWithDefaults()
	resp, err := client.Auth.GithubWriteUserMapping(
		context.Background(),
		key,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("github"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **githubWriteUserMappingRequest** | [**GithubWriteUserMappingRequest**](GithubWriteUserMappingRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudConfigureAuth



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGoogleCloudConfigureAuthRequestWithDefaults()
	resp, err := client.Auth.GoogleCloudConfigureAuth(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudConfigureAuthRequest** | [**GoogleCloudConfigureAuthRequest**](GoogleCloudConfigureAuthRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteRole

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.GoogleCloudDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudEditLabelsForRole

Add or remove labels for an existing 'gce' role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewGoogleCloudEditLabelsForRoleRequestWithDefaults()
	resp, err := client.Auth.GoogleCloudEditLabelsForRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudEditLabelsForRoleRequest** | [**GoogleCloudEditLabelsForRoleRequest**](GoogleCloudEditLabelsForRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudEditServiceAccountsForRole

Add or remove service accounts for an existing `iam` role

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewGoogleCloudEditServiceAccountsForRoleRequestWithDefaults()
	resp, err := client.Auth.GoogleCloudEditServiceAccountsForRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudEditServiceAccountsForRoleRequest** | [**GoogleCloudEditServiceAccountsForRoleRequest**](GoogleCloudEditServiceAccountsForRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRoles

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.GoogleCloudListRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGoogleCloudLoginRequestWithDefaults()
	resp, err := client.Auth.GoogleCloudLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudLoginRequest** | [**GoogleCloudLoginRequest**](GoogleCloudLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadAuthConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.GoogleCloudReadAuthConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRole

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.GoogleCloudReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRole

Create a GCP role with associated policies and required attributes.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewGoogleCloudWriteRoleRequestWithDefaults()
	resp, err := client.Auth.GoogleCloudWriteRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("gcp"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **googleCloudWriteRoleRequest** | [**GoogleCloudWriteRoleRequest**](GoogleCloudWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtConfigure

Configure the JWT authentication backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewJwtConfigureRequestWithDefaults()
	resp, err := client.Auth.JwtConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtConfigureRequest** | [**JwtConfigureRequest**](JwtConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtDeleteRole

Delete an existing role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.JwtDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtListRoles

Lists all the roles registered with the backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.JwtListRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtLogin

Authenticates to Vault using a JWT (or OIDC) token.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewJwtLoginRequestWithDefaults()
	resp, err := client.Auth.JwtLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtLoginRequest** | [**JwtLoginRequest**](JwtLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtOidcCallback

Callback endpoint to complete an OIDC login.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.JwtOidcCallback(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtOidcCallbackWithParameters

Callback endpoint to handle form_posts.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewJwtOidcCallbackWithParametersRequestWithDefaults()
	resp, err := client.Auth.JwtOidcCallbackWithParameters(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtOidcCallbackWithParametersRequest** | [**JwtOidcCallbackWithParametersRequest**](JwtOidcCallbackWithParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtOidcRequestAuthorizationUrl

Request an authorization URL to start an OIDC login flow.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewJwtOidcRequestAuthorizationUrlRequestWithDefaults()
	resp, err := client.Auth.JwtOidcRequestAuthorizationUrl(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jwtOidcRequestAuthorizationUrlRequest** | [**JwtOidcRequestAuthorizationUrlRequest**](JwtOidcRequestAuthorizationUrlRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtReadConfiguration

Read the current JWT authentication backend configuration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.JwtReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtReadRole

Read an existing role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.JwtReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JwtWriteRole

Register an role with the backend.



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewJwtWriteRoleRequestWithDefaults()
	resp, err := client.Auth.JwtWriteRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("jwt"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jwtWriteRoleRequest** | [**JwtWriteRoleRequest**](JwtWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewKerberosConfigureRequestWithDefaults()
	resp, err := client.Auth.KerberosConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosConfigureRequest** | [**KerberosConfigureRequest**](KerberosConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosConfigureLdap



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewKerberosConfigureLdapRequestWithDefaults()
	resp, err := client.Auth.KerberosConfigureLdap(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosConfigureLdapRequest** | [**KerberosConfigureLdapRequest**](KerberosConfigureLdapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosDeleteGroup



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	resp, err := client.Auth.KerberosDeleteGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## KerberosListGroups



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.KerberosListGroups(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosLogin



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewKerberosLoginRequestWithDefaults()
	resp, err := client.Auth.KerberosLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kerberosLoginRequest** | [**KerberosLoginRequest**](KerberosLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.KerberosReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosReadGroup



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	resp, err := client.Auth.KerberosReadGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## KerberosReadLdapConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.KerberosReadLdapConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosWriteGroup



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	request := schema.NewKerberosWriteGroupRequestWithDefaults()
	resp, err := client.Auth.KerberosWriteGroup(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kerberos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **kerberosWriteGroupRequest** | [**KerberosWriteGroupRequest**](KerberosWriteGroupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesConfigureAuth



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewKubernetesConfigureAuthRequestWithDefaults()
	resp, err := client.Auth.KubernetesConfigureAuth(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesConfigureAuthRequest** | [**KubernetesConfigureAuthRequest**](KubernetesConfigureAuthRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteAuthRole

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.KubernetesDeleteAuthRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesListAuthRoles

Lists all the roles registered with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.KubernetesListAuthRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesLogin

Authenticates Kubernetes service accounts with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewKubernetesLoginRequestWithDefaults()
	resp, err := client.Auth.KubernetesLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesLoginRequest** | [**KubernetesLoginRequest**](KubernetesLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadAuthConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.KubernetesReadAuthConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadAuthRole

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	resp, err := client.Auth.KubernetesReadAuthRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteAuthRole

Register an role with the backend.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.
	request := schema.NewKubernetesWriteAuthRoleRequestWithDefaults()
	resp, err := client.Auth.KubernetesWriteAuthRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("kubernetes"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kubernetesWriteAuthRoleRequest** | [**KubernetesWriteAuthRoleRequest**](KubernetesWriteAuthRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapConfigureAuth



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewLdapConfigureAuthRequestWithDefaults()
	resp, err := client.Auth.LdapConfigureAuth(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapConfigureAuthRequest** | [**LdapConfigureAuthRequest**](LdapConfigureAuthRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapDeleteGroup

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	resp, err := client.Auth.LdapDeleteGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## LdapDeleteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.
	resp, err := client.Auth.LdapDeleteUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## LdapListGroups

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.LdapListGroups(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapListUsers

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.LdapListUsers(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapLogin

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | DN (distinguished name) to be used for login.
	request := schema.NewLdapLoginRequestWithDefaults()
	resp, err := client.Auth.LdapLogin(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## LdapReadAuthConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.LdapReadAuthConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapReadGroup

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	resp, err := client.Auth.LdapReadGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## LdapReadUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.
	resp, err := client.Auth.LdapReadUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## LdapWriteGroup

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.
	request := schema.NewLdapWriteGroupRequestWithDefaults()
	resp, err := client.Auth.LdapWriteGroup(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **ldapWriteGroupRequest** | [**LdapWriteGroupRequest**](LdapWriteGroupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LdapWriteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.
	request := schema.NewLdapWriteUserRequestWithDefaults()
	resp, err := client.Auth.LdapWriteUser(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("ldap"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **ldapWriteUserRequest** | [**LdapWriteUserRequest**](LdapWriteUserRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OciConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewOciConfigureRequestWithDefaults()
	resp, err := client.Auth.OciConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ociConfigureRequest** | [**OciConfigureRequest**](OciConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OciDeleteConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OciDeleteConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OciDeleteRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	resp, err := client.Auth.OciDeleteRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## OciListRoles

Lists all the roles that are registered with Vault.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OciListRoles(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OciLogin

Authenticates to Vault using OCI credentials

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	request := schema.NewOciLoginRequestWithDefaults()
	resp, err := client.Auth.OciLogin(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## OciReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OciReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OciReadRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	resp, err := client.Auth.OciReadRole(
		context.Background(),
		role,
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## OciWriteRole

Create a role and associate policies to it.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.
	request := schema.NewOciWriteRoleRequestWithDefaults()
	resp, err := client.Auth.OciWriteRole(
		context.Background(),
		role,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("oci"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **ociWriteRoleRequest** | [**OciWriteRoleRequest**](OciWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewOktaConfigureRequestWithDefaults()
	resp, err := client.Auth.OktaConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaConfigureRequest** | [**OktaConfigureRequest**](OktaConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaDeleteGroup

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.
	resp, err := client.Auth.OktaDeleteGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaDeleteUser

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	resp, err := client.Auth.OktaDeleteUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaListGroups

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OktaListGroups(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaListUsers

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OktaListUsers(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaLogin

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username to be used for login.
	request := schema.NewOktaLoginRequestWithDefaults()
	resp, err := client.Auth.OktaLogin(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## OktaReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.OktaReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaReadGroup

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.
	resp, err := client.Auth.OktaReadGroup(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaReadUser

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	resp, err := client.Auth.OktaReadUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaVerify



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nonce := "nonce_example" // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.
	resp, err := client.Auth.OktaVerify(
		context.Background(),
		nonce,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaWriteGroup

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.
	request := schema.NewOktaWriteGroupRequestWithDefaults()
	resp, err := client.Auth.OktaWriteGroup(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oktaWriteGroupRequest** | [**OktaWriteGroupRequest**](OktaWriteGroupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaWriteUser

Manage additional groups for users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.
	request := schema.NewOktaWriteUserRequestWithDefaults()
	resp, err := client.Auth.OktaWriteUser(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("okta"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oktaWriteUserRequest** | [**OktaWriteUserRequest**](OktaWriteUserRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewRadiusConfigureRequestWithDefaults()
	resp, err := client.Auth.RadiusConfigure(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusConfigureRequest** | [**RadiusConfigureRequest**](RadiusConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusDeleteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	resp, err := client.Auth.RadiusDeleteUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusListUsers

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.RadiusListUsers(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusLogin

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewRadiusLoginRequestWithDefaults()
	resp, err := client.Auth.RadiusLogin(
		context.Background(),
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusLoginRequest** | [**RadiusLoginRequest**](RadiusLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusLoginWithUsername

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlusername := "urlusername_example" // string | Username to be used for login. (URL parameter)
	request := schema.NewRadiusLoginWithUsernameRequestWithDefaults()
	resp, err := client.Auth.RadiusLoginWithUsername(
		context.Background(),
		urlusername,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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


 **radiusLoginWithUsernameRequest** | [**RadiusLoginWithUsernameRequest**](RadiusLoginWithUsernameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.RadiusReadConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusReadUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	resp, err := client.Auth.RadiusReadUser(
		context.Background(),
		name,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusWriteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.
	request := schema.NewRadiusWriteUserRequestWithDefaults()
	resp, err := client.Auth.RadiusWriteUser(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("radius"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **radiusWriteUserRequest** | [**RadiusWriteUserRequest**](RadiusWriteUserRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenCreate

The token create path is used to create new tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenCreateRequestWithDefaults()
	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenCreate(
		context.Background(),
		request,
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 
 **format** | **string** | Return json formatted output | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenCreateAgainstRole

This token create path is used to create new tokens adhering to the given role.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	request := schema.NewTokenCreateAgainstRoleRequestWithDefaults()
	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenCreateAgainstRole(
		context.Background(),
		roleName,
		request,
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateAgainstRoleRequest** | [**TokenCreateAgainstRoleRequest**](TokenCreateAgainstRoleRequest.md) |  | 
 **format** | **string** | Return json formatted output | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenCreateOrphan

The token create path is used to create new orphan tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenCreateOrphanRequestWithDefaults()
	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenCreateOrphan(
		context.Background(),
		request,
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreateOrphanRequest** | [**TokenCreateOrphanRequest**](TokenCreateOrphanRequest.md) |  | 
 **format** | **string** | Return json formatted output | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenDeleteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	resp, err := client.Auth.TokenDeleteRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenListAccessors

List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.TokenListAccessors(
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenListRoles

This endpoint lists configured roles.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.TokenListRoles(
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenLookUp



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenLookUpRequestWithDefaults()
	resp, err := client.Auth.TokenLookUp(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookUpRequest** | [**TokenLookUpRequest**](TokenLookUpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenLookUpAccessor

This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenLookUpAccessorRequestWithDefaults()
	resp, err := client.Auth.TokenLookUpAccessor(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLookUpAccessorRequest** | [**TokenLookUpAccessorRequest**](TokenLookUpAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenLookUpSelf



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.TokenLookUpSelf(
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

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenReadRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	resp, err := client.Auth.TokenReadRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRenew

This endpoint will renew the given token and prevent expiration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRenewRequestWithDefaults()
	resp, err := client.Auth.TokenRenew(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRenewAccessor

This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRenewAccessorRequestWithDefaults()
	resp, err := client.Auth.TokenRenewAccessor(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRenewSelf

This endpoint will renew the token used to call it and prevent expiration.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRenewSelfRequestWithDefaults()
	resp, err := client.Auth.TokenRenewSelf(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRevoke

This endpoint will delete the given token and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRevokeRequestWithDefaults()
	resp, err := client.Auth.TokenRevoke(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRevokeAccessor

This endpoint will delete the token associated with the accessor and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRevokeAccessorRequestWithDefaults()
	resp, err := client.Auth.TokenRevokeAccessor(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRevokeOrphan

This endpoint will delete the token and orphan its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewTokenRevokeOrphanRequestWithDefaults()
	resp, err := client.Auth.TokenRevokeOrphan(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
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



## TokenRevokeSelf

This endpoint will delete the token used to call it and all of its child tokens.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.TokenRevokeSelf(
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

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenTidy

This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.TokenTidy(
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

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteRole



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role
	request := schema.NewTokenWriteRoleRequestWithDefaults()
	resp, err := client.Auth.TokenWriteRole(
		context.Background(),
		roleName,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenWriteRoleRequest** | [**TokenWriteRoleRequest**](TokenWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassDeleteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	resp, err := client.Auth.UserpassDeleteUser(
		context.Background(),
		username,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassListUsers

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Auth.UserpassListUsers(
		context.Background(),
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassLogin

Log in with a username and password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username of the user.
	request := schema.NewUserpassLoginRequestWithDefaults()
	resp, err := client.Auth.UserpassLogin(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassLoginRequest** | [**UserpassLoginRequest**](UserpassLoginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassReadUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	resp, err := client.Auth.UserpassReadUser(
		context.Background(),
		username,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassResetPassword

Reset user's password.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	request := schema.NewUserpassResetPasswordRequestWithDefaults()
	resp, err := client.Auth.UserpassResetPassword(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassResetPasswordRequest** | [**UserpassResetPasswordRequest**](UserpassResetPasswordRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassUpdatePolicies

Update the policies associated with the username.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	request := schema.NewUserpassUpdatePoliciesRequestWithDefaults()
	resp, err := client.Auth.UserpassUpdatePolicies(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassUpdatePoliciesRequest** | [**UserpassUpdatePoliciesRequest**](UserpassUpdatePoliciesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassWriteUser

Manage users allowed to authenticate.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.
	request := schema.NewUserpassWriteUserRequestWithDefaults()
	resp, err := client.Auth.UserpassWriteUser(
		context.Background(),
		username,
		request,
		vault.WithToken("my-token"),
		vault.WithMountPath("userpass"),
	)
	if err != nil {
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userpassWriteUserRequest** | [**UserpassWriteUserRequest**](UserpassWriteUserRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



