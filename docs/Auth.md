# Auth

Method | HTTP request | Description
------------- | ------------- | -------------
[**AWSConfigDeleteCertificate**](Auth.md#AWSConfigDeleteCertificate) | **Delete** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigDeleteClient**](Auth.md#AWSConfigDeleteClient) | **Delete** /auth/{aws_mount_path}/config/client | 
[**AWSConfigDeleteIdentityAccessList**](Auth.md#AWSConfigDeleteIdentityAccessList) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigDeleteIdentityWhiteList**](Auth.md#AWSConfigDeleteIdentityWhiteList) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigDeleteRoleTagBlackList**](Auth.md#AWSConfigDeleteRoleTagBlackList) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigDeleteRoleTagDenyList**](Auth.md#AWSConfigDeleteRoleTagDenyList) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigDeleteSecurityTokenServiceAccount**](Auth.md#AWSConfigDeleteSecurityTokenServiceAccount) | **Delete** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSConfigListCertificates**](Auth.md#AWSConfigListCertificates) | **Get** /auth/{aws_mount_path}/config/certificates | 
[**AWSConfigListSecurityTokenService**](Auth.md#AWSConfigListSecurityTokenService) | **Get** /auth/{aws_mount_path}/config/sts | 
[**AWSConfigReadCertificate**](Auth.md#AWSConfigReadCertificate) | **Get** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigReadClient**](Auth.md#AWSConfigReadClient) | **Get** /auth/{aws_mount_path}/config/client | 
[**AWSConfigReadIdentity**](Auth.md#AWSConfigReadIdentity) | **Get** /auth/{aws_mount_path}/config/identity | 
[**AWSConfigReadIdentityAccessList**](Auth.md#AWSConfigReadIdentityAccessList) | **Get** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigReadIdentityWhiteList**](Auth.md#AWSConfigReadIdentityWhiteList) | **Get** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigReadRoleTagBlackList**](Auth.md#AWSConfigReadRoleTagBlackList) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigReadRoleTagDenyList**](Auth.md#AWSConfigReadRoleTagDenyList) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigReadSecurityTokenServiceAccount**](Auth.md#AWSConfigReadSecurityTokenServiceAccount) | **Get** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSConfigRotateRoot**](Auth.md#AWSConfigRotateRoot) | **Post** /auth/{aws_mount_path}/config/rotate-root | 
[**AWSConfigWriteCertificate**](Auth.md#AWSConfigWriteCertificate) | **Post** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigWriteClient**](Auth.md#AWSConfigWriteClient) | **Post** /auth/{aws_mount_path}/config/client | 
[**AWSConfigWriteIdentity**](Auth.md#AWSConfigWriteIdentity) | **Post** /auth/{aws_mount_path}/config/identity | 
[**AWSConfigWriteIdentityAccessList**](Auth.md#AWSConfigWriteIdentityAccessList) | **Post** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigWriteIdentityWhiteList**](Auth.md#AWSConfigWriteIdentityWhiteList) | **Post** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigWriteRoleTagBlackList**](Auth.md#AWSConfigWriteRoleTagBlackList) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigWriteRoleTagDenyList**](Auth.md#AWSConfigWriteRoleTagDenyList) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigWriteSecurityTokenServiceAccount**](Auth.md#AWSConfigWriteSecurityTokenServiceAccount) | **Post** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSDeleteAuthRole**](Auth.md#AWSDeleteAuthRole) | **Delete** /auth/{aws_mount_path}/role/{role} | 
[**AWSDeleteIdentityAccessListFor**](Auth.md#AWSDeleteIdentityAccessListFor) | **Delete** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AWSDeleteIdentityWhiteListFor**](Auth.md#AWSDeleteIdentityWhiteListFor) | **Delete** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AWSDeleteRoleTagBlackListFor**](Auth.md#AWSDeleteRoleTagBlackListFor) | **Delete** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSDeleteRoleTagDenyListFor**](Auth.md#AWSDeleteRoleTagDenyListFor) | **Delete** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSListAuthRoles**](Auth.md#AWSListAuthRoles) | **Get** /auth/{aws_mount_path}/role | 
[**AWSListAuthRoles2**](Auth.md#AWSListAuthRoles2) | **Get** /auth/{aws_mount_path}/roles | 
[**AWSListIdentityAccessList**](Auth.md#AWSListIdentityAccessList) | **Get** /auth/{aws_mount_path}/identity-accesslist | 
[**AWSListIdentityWhiteList**](Auth.md#AWSListIdentityWhiteList) | **Get** /auth/{aws_mount_path}/identity-whitelist | 
[**AWSListRoleTagBlackList**](Auth.md#AWSListRoleTagBlackList) | **Get** /auth/{aws_mount_path}/roletag-blacklist | 
[**AWSListRoleTagDenyList**](Auth.md#AWSListRoleTagDenyList) | **Get** /auth/{aws_mount_path}/roletag-denylist | 
[**AWSLogin**](Auth.md#AWSLogin) | **Post** /auth/{aws_mount_path}/login | 
[**AWSReadAuthRole**](Auth.md#AWSReadAuthRole) | **Get** /auth/{aws_mount_path}/role/{role} | 
[**AWSReadIdentityAccessListFor**](Auth.md#AWSReadIdentityAccessListFor) | **Get** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AWSReadIdentityWhiteListFor**](Auth.md#AWSReadIdentityWhiteListFor) | **Get** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AWSReadRoleTagBlackListFor**](Auth.md#AWSReadRoleTagBlackListFor) | **Get** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSReadRoleTagDenyListFor**](Auth.md#AWSReadRoleTagDenyListFor) | **Get** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSWriteAuthRole**](Auth.md#AWSWriteAuthRole) | **Post** /auth/{aws_mount_path}/role/{role} | 
[**AWSWriteAuthRoleTag**](Auth.md#AWSWriteAuthRoleTag) | **Post** /auth/{aws_mount_path}/role/{role}/tag | 
[**AWSWriteIdentityAccessListTidySettings**](Auth.md#AWSWriteIdentityAccessListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/identity-accesslist | 
[**AWSWriteIdentityWhiteListTidySettings**](Auth.md#AWSWriteIdentityWhiteListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/identity-whitelist | 
[**AWSWriteRoleTagBlackListFor**](Auth.md#AWSWriteRoleTagBlackListFor) | **Post** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSWriteRoleTagBlackListTidySettings**](Auth.md#AWSWriteRoleTagBlackListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/roletag-blacklist | 
[**AWSWriteRoleTagDenyListFor**](Auth.md#AWSWriteRoleTagDenyListFor) | **Post** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSWriteRoleTagDenyListTidySettings**](Auth.md#AWSWriteRoleTagDenyListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/roletag-denylist | 
[**AliCloudDeleteAuthRole**](Auth.md#AliCloudDeleteAuthRole) | **Delete** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudListAuthRoles**](Auth.md#AliCloudListAuthRoles) | **Get** /auth/{alicloud_mount_path}/role | Lists all the roles that are registered with Vault.
[**AliCloudListAuthRoles2**](Auth.md#AliCloudListAuthRoles2) | **Get** /auth/{alicloud_mount_path}/roles | Lists all the roles that are registered with Vault.
[**AliCloudLogin**](Auth.md#AliCloudLogin) | **Post** /auth/{alicloud_mount_path}/login | Authenticates an RAM entity with Vault.
[**AliCloudReadAuthRole**](Auth.md#AliCloudReadAuthRole) | **Get** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudWriteAuthRole**](Auth.md#AliCloudWriteAuthRole) | **Post** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AppRoleDeleteBindSecretID**](Auth.md#AppRoleDeleteBindSecretID) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleDeleteBoundCIDRList**](Auth.md#AppRoleDeleteBoundCIDRList) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleDeletePeriod**](Auth.md#AppRoleDeletePeriod) | **Delete** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleDeletePolicies**](Auth.md#AppRoleDeletePolicies) | **Delete** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleDeleteRole**](Auth.md#AppRoleDeleteRole) | **Delete** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleDeleteSecretIDAccessorDestroy**](Auth.md#AppRoleDeleteSecretIDAccessorDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**AppRoleDeleteSecretIDBoundCIDRs**](Auth.md#AppRoleDeleteSecretIDBoundCIDRs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleDeleteSecretIDDestroy**](Auth.md#AppRoleDeleteSecretIDDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | 
[**AppRoleDeleteSecretIDNumUses**](Auth.md#AppRoleDeleteSecretIDNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleDeleteSecretIDTTL**](Auth.md#AppRoleDeleteSecretIDTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleDeleteTokenBoundCIDRs**](Auth.md#AppRoleDeleteTokenBoundCIDRs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleDeleteTokenMaxTTL**](Auth.md#AppRoleDeleteTokenMaxTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleDeleteTokenNumUses**](Auth.md#AppRoleDeleteTokenNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleDeleteTokenTTL**](Auth.md#AppRoleDeleteTokenTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleListRoles**](Auth.md#AppRoleListRoles) | **Get** /auth/{approle_mount_path}/role | 
[**AppRoleListSecretID**](Auth.md#AppRoleListSecretID) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleLogin**](Auth.md#AppRoleLogin) | **Post** /auth/{approle_mount_path}/login | 
[**AppRoleReadBindSecretID**](Auth.md#AppRoleReadBindSecretID) | **Get** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleReadBoundCIDRList**](Auth.md#AppRoleReadBoundCIDRList) | **Get** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleReadLocalSecretIDs**](Auth.md#AppRoleReadLocalSecretIDs) | **Get** /auth/{approle_mount_path}/role/{role_name}/local-secret-ids | 
[**AppRoleReadPeriod**](Auth.md#AppRoleReadPeriod) | **Get** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleReadPolicies**](Auth.md#AppRoleReadPolicies) | **Get** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleReadRole**](Auth.md#AppRoleReadRole) | **Get** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleReadRoleID**](Auth.md#AppRoleReadRoleID) | **Get** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleReadSecretIDBoundCIDRs**](Auth.md#AppRoleReadSecretIDBoundCIDRs) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleReadSecretIDNumUses**](Auth.md#AppRoleReadSecretIDNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleReadSecretIDTTL**](Auth.md#AppRoleReadSecretIDTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleReadTokenBoundCIDRs**](Auth.md#AppRoleReadTokenBoundCIDRs) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleReadTokenMaxTTL**](Auth.md#AppRoleReadTokenMaxTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleReadTokenNumUses**](Auth.md#AppRoleReadTokenNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleReadTokenTTL**](Auth.md#AppRoleReadTokenTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleTidySecretID**](Auth.md#AppRoleTidySecretID) | **Post** /auth/{approle_mount_path}/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
[**AppRoleWriteBindSecretID**](Auth.md#AppRoleWriteBindSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleWriteBoundCIDRList**](Auth.md#AppRoleWriteBoundCIDRList) | **Post** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleWriteCustomSecretID**](Auth.md#AppRoleWriteCustomSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/custom-secret-id | 
[**AppRoleWritePeriod**](Auth.md#AppRoleWritePeriod) | **Post** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleWritePolicies**](Auth.md#AppRoleWritePolicies) | **Post** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleWriteRole**](Auth.md#AppRoleWriteRole) | **Post** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleWriteRoleID**](Auth.md#AppRoleWriteRoleID) | **Post** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleWriteSecretID**](Auth.md#AppRoleWriteSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleWriteSecretIDAccessorDestroy**](Auth.md#AppRoleWriteSecretIDAccessorDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**AppRoleWriteSecretIDAccessorLookup**](Auth.md#AppRoleWriteSecretIDAccessorLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup | 
[**AppRoleWriteSecretIDBoundCIDRs**](Auth.md#AppRoleWriteSecretIDBoundCIDRs) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleWriteSecretIDDestroy**](Auth.md#AppRoleWriteSecretIDDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | 
[**AppRoleWriteSecretIDLookup**](Auth.md#AppRoleWriteSecretIDLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/lookup | 
[**AppRoleWriteSecretIDNumUses**](Auth.md#AppRoleWriteSecretIDNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleWriteSecretIDTTL**](Auth.md#AppRoleWriteSecretIDTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleWriteTokenBoundCIDRs**](Auth.md#AppRoleWriteTokenBoundCIDRs) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleWriteTokenMaxTTL**](Auth.md#AppRoleWriteTokenMaxTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleWriteTokenNumUses**](Auth.md#AppRoleWriteTokenNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleWriteTokenTTL**](Auth.md#AppRoleWriteTokenTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AzureDeleteAuthConfig**](Auth.md#AzureDeleteAuthConfig) | **Delete** /auth/{azure_mount_path}/config | 
[**AzureDeleteAuthRole**](Auth.md#AzureDeleteAuthRole) | **Delete** /auth/{azure_mount_path}/role/{name} | 
[**AzureListAuthRoles**](Auth.md#AzureListAuthRoles) | **Get** /auth/{azure_mount_path}/role | 
[**AzureLogin**](Auth.md#AzureLogin) | **Post** /auth/{azure_mount_path}/login | 
[**AzureReadAuthConfig**](Auth.md#AzureReadAuthConfig) | **Get** /auth/{azure_mount_path}/config | 
[**AzureReadAuthRole**](Auth.md#AzureReadAuthRole) | **Get** /auth/{azure_mount_path}/role/{name} | 
[**AzureWriteAuthConfig**](Auth.md#AzureWriteAuthConfig) | **Post** /auth/{azure_mount_path}/config | 
[**AzureWriteAuthRole**](Auth.md#AzureWriteAuthRole) | **Post** /auth/{azure_mount_path}/role/{name} | 
[**CentrifyLogin**](Auth.md#CentrifyLogin) | **Post** /auth/{centrify_mount_path}/login | Log in with a username and password.
[**CentrifyReadConfig**](Auth.md#CentrifyReadConfig) | **Get** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**CentrifyWriteConfig**](Auth.md#CentrifyWriteConfig) | **Post** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**CertificatesDelete**](Auth.md#CertificatesDelete) | **Delete** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesDeleteCRL**](Auth.md#CertificatesDeleteCRL) | **Delete** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesList**](Auth.md#CertificatesList) | **Get** /auth/{cert_mount_path}/certs | Manage trusted certificates used for authentication.
[**CertificatesListCRLs**](Auth.md#CertificatesListCRLs) | **Get** /auth/{cert_mount_path}/crls | 
[**CertificatesLogin**](Auth.md#CertificatesLogin) | **Post** /auth/{cert_mount_path}/login | 
[**CertificatesRead**](Auth.md#CertificatesRead) | **Get** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesReadCRL**](Auth.md#CertificatesReadCRL) | **Get** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesReadConfig**](Auth.md#CertificatesReadConfig) | **Get** /auth/{cert_mount_path}/config | 
[**CertificatesWrite**](Auth.md#CertificatesWrite) | **Post** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesWriteCRL**](Auth.md#CertificatesWriteCRL) | **Post** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesWriteConfig**](Auth.md#CertificatesWriteConfig) | **Post** /auth/{cert_mount_path}/config | 
[**CloudFoundryDeleteConfig**](Auth.md#CloudFoundryDeleteConfig) | **Delete** /auth/{cf_mount_path}/config | 
[**CloudFoundryDeleteRole**](Auth.md#CloudFoundryDeleteRole) | **Delete** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryListRoles**](Auth.md#CloudFoundryListRoles) | **Get** /auth/{cf_mount_path}/roles | 
[**CloudFoundryLogin**](Auth.md#CloudFoundryLogin) | **Post** /auth/{cf_mount_path}/login | 
[**CloudFoundryReadConfig**](Auth.md#CloudFoundryReadConfig) | **Get** /auth/{cf_mount_path}/config | 
[**CloudFoundryReadRole**](Auth.md#CloudFoundryReadRole) | **Get** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryWriteConfig**](Auth.md#CloudFoundryWriteConfig) | **Post** /auth/{cf_mount_path}/config | 
[**CloudFoundryWriteRole**](Auth.md#CloudFoundryWriteRole) | **Post** /auth/{cf_mount_path}/roles/{role} | 
[**GitHubDeleteMapTeam**](Auth.md#GitHubDeleteMapTeam) | **Delete** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubDeleteMapUser**](Auth.md#GitHubDeleteMapUser) | **Delete** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GitHubLogin**](Auth.md#GitHubLogin) | **Post** /auth/{github_mount_path}/login | 
[**GitHubReadConfig**](Auth.md#GitHubReadConfig) | **Get** /auth/{github_mount_path}/config | 
[**GitHubReadMapTeam**](Auth.md#GitHubReadMapTeam) | **Get** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubReadMapTeams**](Auth.md#GitHubReadMapTeams) | **Get** /auth/{github_mount_path}/map/teams | Read mappings for teams
[**GitHubReadMapUser**](Auth.md#GitHubReadMapUser) | **Get** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GitHubReadMapUsers**](Auth.md#GitHubReadMapUsers) | **Get** /auth/{github_mount_path}/map/users | Read mappings for users
[**GitHubWriteConfig**](Auth.md#GitHubWriteConfig) | **Post** /auth/{github_mount_path}/config | 
[**GitHubWriteMapTeam**](Auth.md#GitHubWriteMapTeam) | **Post** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubWriteMapUser**](Auth.md#GitHubWriteMapUser) | **Post** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GoogleCloudDeleteRole**](Auth.md#GoogleCloudDeleteRole) | **Delete** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudListRoles**](Auth.md#GoogleCloudListRoles) | **Get** /auth/{gcp_mount_path}/role | Lists all the roles that are registered with Vault.
[**GoogleCloudListRoles2**](Auth.md#GoogleCloudListRoles2) | **Get** /auth/{gcp_mount_path}/roles | Lists all the roles that are registered with Vault.
[**GoogleCloudLogin**](Auth.md#GoogleCloudLogin) | **Post** /auth/{gcp_mount_path}/login | 
[**GoogleCloudReadAuthConfig**](Auth.md#GoogleCloudReadAuthConfig) | **Get** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GoogleCloudReadRole**](Auth.md#GoogleCloudReadRole) | **Get** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudWriteAuthConfig**](Auth.md#GoogleCloudWriteAuthConfig) | **Post** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GoogleCloudWriteRole**](Auth.md#GoogleCloudWriteRole) | **Post** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudWriteRoleLabels**](Auth.md#GoogleCloudWriteRoleLabels) | **Post** /auth/{gcp_mount_path}/role/{name}/labels | Add or remove labels for an existing &#x27;gce&#x27; role
[**GoogleCloudWriteRoleServiceAccounts**](Auth.md#GoogleCloudWriteRoleServiceAccounts) | **Post** /auth/{gcp_mount_path}/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**JWTDeleteRole**](Auth.md#JWTDeleteRole) | **Delete** /auth/{jwt_mount_path}/role/{name} | Delete an existing role.
[**JWTListRoles**](Auth.md#JWTListRoles) | **Get** /auth/{jwt_mount_path}/role | Lists all the roles registered with the backend.
[**JWTLogin**](Auth.md#JWTLogin) | **Post** /auth/{jwt_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**JWTReadConfig**](Auth.md#JWTReadConfig) | **Get** /auth/{jwt_mount_path}/config | Read the current JWT authentication backend configuration.
[**JWTReadOIDCCallback**](Auth.md#JWTReadOIDCCallback) | **Get** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**JWTReadRole**](Auth.md#JWTReadRole) | **Get** /auth/{jwt_mount_path}/role/{name} | Read an existing role.
[**JWTWriteConfig**](Auth.md#JWTWriteConfig) | **Post** /auth/{jwt_mount_path}/config | Configure the JWT authentication backend.
[**JWTWriteOIDCAuthURL**](Auth.md#JWTWriteOIDCAuthURL) | **Post** /auth/{jwt_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**JWTWriteOIDCCallback**](Auth.md#JWTWriteOIDCCallback) | **Post** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**JWTWriteRole**](Auth.md#JWTWriteRole) | **Post** /auth/{jwt_mount_path}/role/{name} | Register an role with the backend.
[**KerberosDeleteGroup**](Auth.md#KerberosDeleteGroup) | **Delete** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosListGroups**](Auth.md#KerberosListGroups) | **Get** /auth/{kerberos_mount_path}/groups | 
[**KerberosLogin**](Auth.md#KerberosLogin) | **Post** /auth/{kerberos_mount_path}/login | 
[**KerberosReadConfig**](Auth.md#KerberosReadConfig) | **Get** /auth/{kerberos_mount_path}/config | 
[**KerberosReadGroup**](Auth.md#KerberosReadGroup) | **Get** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosReadLDAPConfig**](Auth.md#KerberosReadLDAPConfig) | **Get** /auth/{kerberos_mount_path}/config/ldap | 
[**KerberosWriteConfig**](Auth.md#KerberosWriteConfig) | **Post** /auth/{kerberos_mount_path}/config | 
[**KerberosWriteGroup**](Auth.md#KerberosWriteGroup) | **Post** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosWriteLDAPConfig**](Auth.md#KerberosWriteLDAPConfig) | **Post** /auth/{kerberos_mount_path}/config/ldap | 
[**KubernetesDeleteAuthRole**](Auth.md#KubernetesDeleteAuthRole) | **Delete** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesListAuthRoles**](Auth.md#KubernetesListAuthRoles) | **Get** /auth/{kubernetes_mount_path}/role | Lists all the roles registered with the backend.
[**KubernetesLogin**](Auth.md#KubernetesLogin) | **Post** /auth/{kubernetes_mount_path}/login | Authenticates Kubernetes service accounts with Vault.
[**KubernetesReadAuthConfig**](Auth.md#KubernetesReadAuthConfig) | **Get** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**KubernetesReadAuthRole**](Auth.md#KubernetesReadAuthRole) | **Get** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesWriteAuthConfig**](Auth.md#KubernetesWriteAuthConfig) | **Post** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**KubernetesWriteAuthRole**](Auth.md#KubernetesWriteAuthRole) | **Post** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**LDAPDeleteGroup**](Auth.md#LDAPDeleteGroup) | **Delete** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPDeleteUser**](Auth.md#LDAPDeleteUser) | **Delete** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LDAPListGroups**](Auth.md#LDAPListGroups) | **Get** /auth/{ldap_mount_path}/groups | Manage additional groups for users allowed to authenticate.
[**LDAPListUsers**](Auth.md#LDAPListUsers) | **Get** /auth/{ldap_mount_path}/users | Manage users allowed to authenticate.
[**LDAPLogin**](Auth.md#LDAPLogin) | **Post** /auth/{ldap_mount_path}/login/{username} | Log in with a username and password.
[**LDAPReadAuthConfig**](Auth.md#LDAPReadAuthConfig) | **Get** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**LDAPReadGroup**](Auth.md#LDAPReadGroup) | **Get** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPReadUser**](Auth.md#LDAPReadUser) | **Get** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LDAPWriteAuthConfig**](Auth.md#LDAPWriteAuthConfig) | **Post** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**LDAPWriteGroup**](Auth.md#LDAPWriteGroup) | **Post** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPWriteUser**](Auth.md#LDAPWriteUser) | **Post** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**OCIDeleteConfig**](Auth.md#OCIDeleteConfig) | **Delete** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIDeleteRole**](Auth.md#OCIDeleteRole) | **Delete** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OCIListRoles**](Auth.md#OCIListRoles) | **Get** /auth/{oci_mount_path}/role | Lists all the roles that are registered with Vault.
[**OCILoginWithRole**](Auth.md#OCILoginWithRole) | **Post** /auth/{oci_mount_path}/login/{role} | Authenticates to Vault using OCI credentials
[**OCIReadConfig**](Auth.md#OCIReadConfig) | **Get** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIReadRole**](Auth.md#OCIReadRole) | **Get** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OCIWriteConfig**](Auth.md#OCIWriteConfig) | **Post** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIWriteRole**](Auth.md#OCIWriteRole) | **Post** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OIDCDeleteAuthRole**](Auth.md#OIDCDeleteAuthRole) | **Delete** /auth/{oidc_mount_path}/role/{name} | Delete an existing role.
[**OIDCListAuthRoles**](Auth.md#OIDCListAuthRoles) | **Get** /auth/{oidc_mount_path}/role | Lists all the roles registered with the backend.
[**OIDCLogin**](Auth.md#OIDCLogin) | **Post** /auth/{oidc_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**OIDCReadAuthConfig**](Auth.md#OIDCReadAuthConfig) | **Get** /auth/{oidc_mount_path}/config | Read the current JWT authentication backend configuration.
[**OIDCReadAuthRole**](Auth.md#OIDCReadAuthRole) | **Get** /auth/{oidc_mount_path}/role/{name} | Read an existing role.
[**OIDCReadCallback**](Auth.md#OIDCReadCallback) | **Get** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**OIDCWriteAuthConfig**](Auth.md#OIDCWriteAuthConfig) | **Post** /auth/{oidc_mount_path}/config | Configure the JWT authentication backend.
[**OIDCWriteAuthRole**](Auth.md#OIDCWriteAuthRole) | **Post** /auth/{oidc_mount_path}/role/{name} | Register an role with the backend.
[**OIDCWriteAuthURL**](Auth.md#OIDCWriteAuthURL) | **Post** /auth/{oidc_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**OIDCWriteCallback**](Auth.md#OIDCWriteCallback) | **Post** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**OktaDeleteGroup**](Auth.md#OktaDeleteGroup) | **Delete** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaDeleteUser**](Auth.md#OktaDeleteUser) | **Delete** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaListGroups**](Auth.md#OktaListGroups) | **Get** /auth/{okta_mount_path}/groups | Manage users allowed to authenticate.
[**OktaListUsers**](Auth.md#OktaListUsers) | **Get** /auth/{okta_mount_path}/users | Manage additional groups for users allowed to authenticate.
[**OktaLogin**](Auth.md#OktaLogin) | **Post** /auth/{okta_mount_path}/login/{username} | Log in with a username and password.
[**OktaReadConfig**](Auth.md#OktaReadConfig) | **Get** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**OktaReadGroup**](Auth.md#OktaReadGroup) | **Get** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaReadUser**](Auth.md#OktaReadUser) | **Get** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaVerify**](Auth.md#OktaVerify) | **Get** /auth/{okta_mount_path}/verify/{nonce} | 
[**OktaWriteConfig**](Auth.md#OktaWriteConfig) | **Post** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**OktaWriteGroup**](Auth.md#OktaWriteGroup) | **Post** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaWriteUser**](Auth.md#OktaWriteUser) | **Post** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**RadiusDeleteUser**](Auth.md#RadiusDeleteUser) | **Delete** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusListUsers**](Auth.md#RadiusListUsers) | **Get** /auth/{radius_mount_path}/users | Manage users allowed to authenticate.
[**RadiusLogin**](Auth.md#RadiusLogin) | **Post** /auth/{radius_mount_path}/login | Log in with a username and password.
[**RadiusLoginWithUsername**](Auth.md#RadiusLoginWithUsername) | **Post** /auth/{radius_mount_path}/login/{urlusername} | Log in with a username and password.
[**RadiusReadConfig**](Auth.md#RadiusReadConfig) | **Get** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**RadiusReadUser**](Auth.md#RadiusReadUser) | **Get** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusWriteConfig**](Auth.md#RadiusWriteConfig) | **Post** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**RadiusWriteUser**](Auth.md#RadiusWriteUser) | **Post** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**TokenDeleteRole**](Auth.md#TokenDeleteRole) | **Delete** /auth/{token_mount_path}/roles/{role_name} | 
[**TokenListAccessors**](Auth.md#TokenListAccessors) | **Get** /auth/{token_mount_path}/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#x27;sudo&#x27; capability in addition to &#x27;list&#x27;.
[**TokenListRoles**](Auth.md#TokenListRoles) | **Get** /auth/{token_mount_path}/roles | This endpoint lists configured roles.
[**TokenReadLookup**](Auth.md#TokenReadLookup) | **Get** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**TokenReadLookupSelf**](Auth.md#TokenReadLookupSelf) | **Get** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**TokenReadRole**](Auth.md#TokenReadRole) | **Get** /auth/{token_mount_path}/roles/{role_name} | 
[**TokenRenew**](Auth.md#TokenRenew) | **Post** /auth/{token_mount_path}/renew | This endpoint will renew the given token and prevent expiration.
[**TokenRenewAccessor**](Auth.md#TokenRenewAccessor) | **Post** /auth/{token_mount_path}/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenRenewSelf**](Auth.md#TokenRenewSelf) | **Post** /auth/{token_mount_path}/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**TokenRevoke**](Auth.md#TokenRevoke) | **Post** /auth/{token_mount_path}/revoke | This endpoint will delete the given token and all of its child tokens.
[**TokenRevokeAccessor**](Auth.md#TokenRevokeAccessor) | **Post** /auth/{token_mount_path}/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**TokenRevokeOrphan**](Auth.md#TokenRevokeOrphan) | **Post** /auth/{token_mount_path}/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**TokenRevokeSelf**](Auth.md#TokenRevokeSelf) | **Post** /auth/{token_mount_path}/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**TokenTidy**](Auth.md#TokenTidy) | **Post** /auth/{token_mount_path}/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**TokenWriteCreate**](Auth.md#TokenWriteCreate) | **Post** /auth/{token_mount_path}/create | The token create path is used to create new tokens.
[**TokenWriteCreateOrphan**](Auth.md#TokenWriteCreateOrphan) | **Post** /auth/{token_mount_path}/create-orphan | The token create path is used to create new orphan tokens.
[**TokenWriteCreateWithRole**](Auth.md#TokenWriteCreateWithRole) | **Post** /auth/{token_mount_path}/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**TokenWriteLookup**](Auth.md#TokenWriteLookup) | **Post** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**TokenWriteLookupAccessor**](Auth.md#TokenWriteLookupAccessor) | **Post** /auth/{token_mount_path}/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenWriteLookupSelf**](Auth.md#TokenWriteLookupSelf) | **Post** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**TokenWriteRole**](Auth.md#TokenWriteRole) | **Post** /auth/{token_mount_path}/roles/{role_name} | 
[**UserpassDeleteUser**](Auth.md#UserpassDeleteUser) | **Delete** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassListUsers**](Auth.md#UserpassListUsers) | **Get** /auth/{userpass_mount_path}/users | Manage users allowed to authenticate.
[**UserpassLogin**](Auth.md#UserpassLogin) | **Post** /auth/{userpass_mount_path}/login/{username} | Log in with a username and password.
[**UserpassReadUser**](Auth.md#UserpassReadUser) | **Get** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassWriteUser**](Auth.md#UserpassWriteUser) | **Post** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassWriteUserPassword**](Auth.md#UserpassWriteUserPassword) | **Post** /auth/{userpass_mount_path}/users/{username}/password | Reset user&#x27;s password.
[**UserpassWriteUserPolicies**](Auth.md#UserpassWriteUserPolicies) | **Post** /auth/{userpass_mount_path}/users/{username}/policies | Update the policies associated with the username.





## AWSConfigDeleteCertificate

> AWSConfigDeleteCertificate(ctx, awsMountPath, certName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.


	resp, err := client.Auth.AWSConfigDeleteCertificate(
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteClient

> AWSConfigDeleteClient(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigDeleteClient(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteIdentityAccessList

> AWSConfigDeleteIdentityAccessList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigDeleteIdentityAccessList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteIdentityWhiteList

> AWSConfigDeleteIdentityWhiteList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigDeleteIdentityWhiteList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteRoleTagBlackList

> AWSConfigDeleteRoleTagBlackList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigDeleteRoleTagBlackList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteRoleTagDenyList

> AWSConfigDeleteRoleTagDenyList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigDeleteRoleTagDenyList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigDeleteSecurityTokenServiceAccount

> AWSConfigDeleteSecurityTokenServiceAccount(ctx, accountId, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.


	resp, err := client.Auth.AWSConfigDeleteSecurityTokenServiceAccount(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigListCertificates

> AWSConfigListCertificates(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigListCertificates(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigListSecurityTokenService

> AWSConfigListSecurityTokenService(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigListSecurityTokenService(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadCertificate

> AWSConfigReadCertificate(ctx, awsMountPath, certName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.


	resp, err := client.Auth.AWSConfigReadCertificate(
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadClient

> AWSConfigReadClient(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadClient(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadIdentity

> AWSConfigReadIdentity(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadIdentity(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadIdentityAccessList

> AWSConfigReadIdentityAccessList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadIdentityAccessList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadIdentityWhiteList

> AWSConfigReadIdentityWhiteList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadIdentityWhiteList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadRoleTagBlackList

> AWSConfigReadRoleTagBlackList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadRoleTagBlackList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadRoleTagDenyList

> AWSConfigReadRoleTagDenyList(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigReadRoleTagDenyList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigReadSecurityTokenServiceAccount

> AWSConfigReadSecurityTokenServiceAccount(ctx, accountId, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.


	resp, err := client.Auth.AWSConfigReadSecurityTokenServiceAccount(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigRotateRoot

> AWSConfigRotateRoot(ctx, awsMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSConfigRotateRoot(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteCertificate

> AWSConfigWriteCertificate(ctx, awsMountPath, certName).AWSConfigWriteCertificateRequest(aWSConfigWriteCertificateRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	certName := "certName_example" // string | Name of the certificate.

	request := schema.NewAWSConfigWriteCertificateRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteCertificate(
		context.Background(),
		certName,
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
**certName** | **string** | Name of the certificate. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSConfigWriteCertificateRequest** | [**AWSConfigWriteCertificateRequest**](AWSConfigWriteCertificateRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteClient

> AWSConfigWriteClient(ctx, awsMountPath).AWSConfigWriteClientRequest(aWSConfigWriteClientRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteClientRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteClient(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteClientRequest** | [**AWSConfigWriteClientRequest**](AWSConfigWriteClientRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteIdentity

> AWSConfigWriteIdentity(ctx, awsMountPath).AWSConfigWriteIdentityRequest(aWSConfigWriteIdentityRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteIdentityRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteIdentity(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteIdentityRequest** | [**AWSConfigWriteIdentityRequest**](AWSConfigWriteIdentityRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteIdentityAccessList

> AWSConfigWriteIdentityAccessList(ctx, awsMountPath).AWSConfigWriteIdentityAccessListRequest(aWSConfigWriteIdentityAccessListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteIdentityAccessListRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteIdentityAccessList(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteIdentityAccessListRequest** | [**AWSConfigWriteIdentityAccessListRequest**](AWSConfigWriteIdentityAccessListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteIdentityWhiteList

> AWSConfigWriteIdentityWhiteList(ctx, awsMountPath).AWSConfigWriteIdentityWhiteListRequest(aWSConfigWriteIdentityWhiteListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteIdentityWhiteListRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteIdentityWhiteList(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteIdentityWhiteListRequest** | [**AWSConfigWriteIdentityWhiteListRequest**](AWSConfigWriteIdentityWhiteListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteRoleTagBlackList

> AWSConfigWriteRoleTagBlackList(ctx, awsMountPath).AWSConfigWriteRoleTagBlackListRequest(aWSConfigWriteRoleTagBlackListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteRoleTagBlackListRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteRoleTagBlackList(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteRoleTagBlackListRequest** | [**AWSConfigWriteRoleTagBlackListRequest**](AWSConfigWriteRoleTagBlackListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteRoleTagDenyList

> AWSConfigWriteRoleTagDenyList(ctx, awsMountPath).AWSConfigWriteRoleTagDenyListRequest(aWSConfigWriteRoleTagDenyListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSConfigWriteRoleTagDenyListRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteRoleTagDenyList(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSConfigWriteRoleTagDenyListRequest** | [**AWSConfigWriteRoleTagDenyListRequest**](AWSConfigWriteRoleTagDenyListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSConfigWriteSecurityTokenServiceAccount

> AWSConfigWriteSecurityTokenServiceAccount(ctx, accountId, awsMountPath).AWSConfigWriteSecurityTokenServiceAccountRequest(aWSConfigWriteSecurityTokenServiceAccountRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountId := "accountId_example" // string | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.

	request := schema.NewAWSConfigWriteSecurityTokenServiceAccountRequestWithDefaults()

	resp, err := client.Auth.AWSConfigWriteSecurityTokenServiceAccount(
		context.Background(),
		accountId,
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
**accountId** | **string** | AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSConfigWriteSecurityTokenServiceAccountRequest** | [**AWSConfigWriteSecurityTokenServiceAccountRequest**](AWSConfigWriteSecurityTokenServiceAccountRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteAuthRole

> AWSDeleteAuthRole(ctx, awsMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.


	resp, err := client.Auth.AWSDeleteAuthRole(
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteIdentityAccessListFor

> AWSDeleteIdentityAccessListFor(ctx, awsMountPath, instanceId).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.


	resp, err := client.Auth.AWSDeleteIdentityAccessListFor(
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteIdentityWhiteListFor

> AWSDeleteIdentityWhiteListFor(ctx, awsMountPath, instanceId).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.


	resp, err := client.Auth.AWSDeleteIdentityWhiteListFor(
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteRoleTagBlackListFor

> AWSDeleteRoleTagBlackListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSDeleteRoleTagBlackListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSDeleteRoleTagDenyListFor

> AWSDeleteRoleTagDenyListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSDeleteRoleTagDenyListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListAuthRoles

> AWSListAuthRoles(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListAuthRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListAuthRoles2

> AWSListAuthRoles2(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListAuthRoles2(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListIdentityAccessList

> AWSListIdentityAccessList(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListIdentityAccessList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListIdentityWhiteList

> AWSListIdentityWhiteList(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListIdentityWhiteList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListRoleTagBlackList

> AWSListRoleTagBlackList(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListRoleTagBlackList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSListRoleTagDenyList

> AWSListRoleTagDenyList(ctx, awsMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AWSListRoleTagDenyList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSLogin

> AWSLogin(ctx, awsMountPath).AWSLoginRequest(aWSLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSLoginRequestWithDefaults()

	resp, err := client.Auth.AWSLogin(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSLoginRequest** | [**AWSLoginRequest**](AWSLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadAuthRole

> AWSReadAuthRole(ctx, awsMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.


	resp, err := client.Auth.AWSReadAuthRole(
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadIdentityAccessListFor

> AWSReadIdentityAccessListFor(ctx, awsMountPath, instanceId).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.


	resp, err := client.Auth.AWSReadIdentityAccessListFor(
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadIdentityWhiteListFor

> AWSReadIdentityWhiteListFor(ctx, awsMountPath, instanceId).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	instanceId := "instanceId_example" // string | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.


	resp, err := client.Auth.AWSReadIdentityWhiteListFor(
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
**instanceId** | **string** | EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadRoleTagBlackListFor

> AWSReadRoleTagBlackListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSReadRoleTagBlackListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSReadRoleTagDenyListFor

> AWSReadRoleTagDenyListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSReadRoleTagDenyListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteAuthRole

> AWSWriteAuthRole(ctx, awsMountPath, role).AWSWriteAuthRoleRequest(aWSWriteAuthRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	request := schema.NewAWSWriteAuthRoleRequestWithDefaults()

	resp, err := client.Auth.AWSWriteAuthRole(
		context.Background(),
		role,
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSWriteAuthRoleRequest** | [**AWSWriteAuthRoleRequest**](AWSWriteAuthRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteAuthRoleTag

> AWSWriteAuthRoleTag(ctx, awsMountPath, role).AWSWriteAuthRoleTagRequest(aWSWriteAuthRoleTagRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	request := schema.NewAWSWriteAuthRoleTagRequestWithDefaults()

	resp, err := client.Auth.AWSWriteAuthRoleTag(
		context.Background(),
		role,
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSWriteAuthRoleTagRequest** | [**AWSWriteAuthRoleTagRequest**](AWSWriteAuthRoleTagRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteIdentityAccessListTidySettings

> AWSWriteIdentityAccessListTidySettings(ctx, awsMountPath).AWSWriteIdentityAccessListTidySettingsRequest(aWSWriteIdentityAccessListTidySettingsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSWriteIdentityAccessListTidySettingsRequestWithDefaults()

	resp, err := client.Auth.AWSWriteIdentityAccessListTidySettings(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSWriteIdentityAccessListTidySettingsRequest** | [**AWSWriteIdentityAccessListTidySettingsRequest**](AWSWriteIdentityAccessListTidySettingsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteIdentityWhiteListTidySettings

> AWSWriteIdentityWhiteListTidySettings(ctx, awsMountPath).AWSWriteIdentityWhiteListTidySettingsRequest(aWSWriteIdentityWhiteListTidySettingsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSWriteIdentityWhiteListTidySettingsRequestWithDefaults()

	resp, err := client.Auth.AWSWriteIdentityWhiteListTidySettings(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSWriteIdentityWhiteListTidySettingsRequest** | [**AWSWriteIdentityWhiteListTidySettingsRequest**](AWSWriteIdentityWhiteListTidySettingsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteRoleTagBlackListFor

> AWSWriteRoleTagBlackListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSWriteRoleTagBlackListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteRoleTagBlackListTidySettings

> AWSWriteRoleTagBlackListTidySettings(ctx, awsMountPath).AWSWriteRoleTagBlackListTidySettingsRequest(aWSWriteRoleTagBlackListTidySettingsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSWriteRoleTagBlackListTidySettingsRequestWithDefaults()

	resp, err := client.Auth.AWSWriteRoleTagBlackListTidySettings(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSWriteRoleTagBlackListTidySettingsRequest** | [**AWSWriteRoleTagBlackListTidySettingsRequest**](AWSWriteRoleTagBlackListTidySettingsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteRoleTagDenyListFor

> AWSWriteRoleTagDenyListFor(ctx, awsMountPath, roleTag).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleTag := "roleTag_example" // string | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.


	resp, err := client.Auth.AWSWriteRoleTagDenyListFor(
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
**roleTag** | **string** | Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AWSWriteRoleTagDenyListTidySettings

> AWSWriteRoleTagDenyListTidySettings(ctx, awsMountPath).AWSWriteRoleTagDenyListTidySettingsRequest(aWSWriteRoleTagDenyListTidySettingsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults()

	resp, err := client.Auth.AWSWriteRoleTagDenyListTidySettings(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aWSWriteRoleTagDenyListTidySettingsRequest** | [**AWSWriteRoleTagDenyListTidySettingsRequest**](AWSWriteRoleTagDenyListTidySettingsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudDeleteAuthRole

> AliCloudDeleteAuthRole(ctx, alicloudMountPath, role).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.


	resp, err := client.Auth.AliCloudDeleteAuthRole(
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudListAuthRoles

> AliCloudListAuthRoles(ctx, alicloudMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AliCloudListAuthRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudListAuthRoles2

> AliCloudListAuthRoles2(ctx, alicloudMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AliCloudListAuthRoles2(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudLogin

> AliCloudLogin(ctx, alicloudMountPath).AliCloudLoginRequest(aliCloudLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAliCloudLoginRequestWithDefaults()

	resp, err := client.Auth.AliCloudLogin(
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

> AliCloudReadAuthRole(ctx, alicloudMountPath, role).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role as it should appear in Vault.


	resp, err := client.Auth.AliCloudReadAuthRole(
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
**role** | **string** | The name of the role as it should appear in Vault. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliCloudWriteAuthRole

> AliCloudWriteAuthRole(ctx, alicloudMountPath, role).AliCloudWriteAuthRoleRequest(aliCloudWriteAuthRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## AppRoleDeleteBindSecretID

> AppRoleDeleteBindSecretID(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteBindSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteBoundCIDRList

> AppRoleDeleteBoundCIDRList(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteBoundCIDRList(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeletePeriod

> AppRoleDeletePeriod(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeletePeriod(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeletePolicies

> AppRoleDeletePolicies(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeletePolicies(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteRole

> AppRoleDeleteRole(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteRole(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIDAccessorDestroy

> AppRoleDeleteSecretIDAccessorDestroy(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteSecretIDAccessorDestroy(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIDBoundCIDRs

> AppRoleDeleteSecretIDBoundCIDRs(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteSecretIDBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIDDestroy

> AppRoleDeleteSecretIDDestroy(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteSecretIDDestroy(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIDNumUses

> AppRoleDeleteSecretIDNumUses(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteSecretIDNumUses(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteSecretIDTTL

> AppRoleDeleteSecretIDTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteSecretIDTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenBoundCIDRs

> AppRoleDeleteTokenBoundCIDRs(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteTokenBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenMaxTTL

> AppRoleDeleteTokenMaxTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteTokenMaxTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenNumUses

> AppRoleDeleteTokenNumUses(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteTokenNumUses(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleDeleteTokenTTL

> AppRoleDeleteTokenTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleDeleteTokenTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleListRoles

> AppRoleListRolesResponse AppRoleListRoles(ctx, approleMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AppRoleListRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


[**AppRoleListRolesResponse**](AppRoleListRolesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleListSecretID

> AppRoleListSecretIDResponse AppRoleListSecretID(ctx, approleMountPath, roleName).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleListSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 


[**AppRoleListSecretIDResponse**](AppRoleListSecretIDResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleLogin

> AppRoleLogin(ctx, approleMountPath).AppRoleLoginRequest(appRoleLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAppRoleLoginRequestWithDefaults()

	resp, err := client.Auth.AppRoleLogin(
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



## AppRoleReadBindSecretID

> AppRoleReadBindSecretIDResponse AppRoleReadBindSecretID(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadBindSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadBindSecretIDResponse**](AppRoleReadBindSecretIDResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadBoundCIDRList

> AppRoleReadBoundCIDRListResponse AppRoleReadBoundCIDRList(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadBoundCIDRList(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadBoundCIDRListResponse**](AppRoleReadBoundCIDRListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadLocalSecretIDs

> AppRoleReadLocalSecretIDsResponse AppRoleReadLocalSecretIDs(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadLocalSecretIDs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadLocalSecretIDsResponse**](AppRoleReadLocalSecretIDsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadPeriod

> AppRoleReadPeriodResponse AppRoleReadPeriod(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadPeriod(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadPeriodResponse**](AppRoleReadPeriodResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadPolicies

> AppRoleReadPoliciesResponse AppRoleReadPolicies(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadPolicies(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadPoliciesResponse**](AppRoleReadPoliciesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadRole

> AppRoleReadRoleResponse AppRoleReadRole(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadRole(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadRoleResponse**](AppRoleReadRoleResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadRoleID

> AppRoleReadRoleIDResponse AppRoleReadRoleID(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadRoleID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadRoleIDResponse**](AppRoleReadRoleIDResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIDBoundCIDRs

> AppRoleReadSecretIDBoundCIDRsResponse AppRoleReadSecretIDBoundCIDRs(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadSecretIDBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIDBoundCIDRsResponse**](AppRoleReadSecretIDBoundCIDRsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIDNumUses

> AppRoleReadSecretIDNumUsesResponse AppRoleReadSecretIDNumUses(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadSecretIDNumUses(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIDNumUsesResponse**](AppRoleReadSecretIDNumUsesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadSecretIDTTL

> AppRoleReadSecretIDTTLResponse AppRoleReadSecretIDTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadSecretIDTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadSecretIDTTLResponse**](AppRoleReadSecretIDTTLResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenBoundCIDRs

> AppRoleReadTokenBoundCIDRsResponse AppRoleReadTokenBoundCIDRs(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadTokenBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenBoundCIDRsResponse**](AppRoleReadTokenBoundCIDRsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenMaxTTL

> AppRoleReadTokenMaxTTLResponse AppRoleReadTokenMaxTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadTokenMaxTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenMaxTTLResponse**](AppRoleReadTokenMaxTTLResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenNumUses

> AppRoleReadTokenNumUsesResponse AppRoleReadTokenNumUses(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadTokenNumUses(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenNumUsesResponse**](AppRoleReadTokenNumUsesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleReadTokenTTL

> AppRoleReadTokenTTLResponse AppRoleReadTokenTTL(ctx, approleMountPath, roleName).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.


	resp, err := client.Auth.AppRoleReadTokenTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**AppRoleReadTokenTTLResponse**](AppRoleReadTokenTTLResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleTidySecretID

> AppRoleTidySecretID(ctx, approleMountPath).Execute()

Trigger the clean-up of expired SecretID entries.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AppRoleTidySecretID(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteBindSecretID

> AppRoleWriteBindSecretID(ctx, approleMountPath, roleName).AppRoleWriteBindSecretIDRequest(appRoleWriteBindSecretIDRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteBindSecretIDRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteBindSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteBindSecretIDRequest** | [**AppRoleWriteBindSecretIDRequest**](AppRoleWriteBindSecretIDRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteBoundCIDRList

> AppRoleWriteBoundCIDRList(ctx, approleMountPath, roleName).AppRoleWriteBoundCIDRListRequest(appRoleWriteBoundCIDRListRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteBoundCIDRListRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteBoundCIDRList(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteBoundCIDRListRequest** | [**AppRoleWriteBoundCIDRListRequest**](AppRoleWriteBoundCIDRListRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteCustomSecretID

> AppRoleWriteCustomSecretIDResponse AppRoleWriteCustomSecretID(ctx, approleMountPath, roleName).AppRoleWriteCustomSecretIDRequest(appRoleWriteCustomSecretIDRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteCustomSecretIDRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteCustomSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteCustomSecretIDRequest** | [**AppRoleWriteCustomSecretIDRequest**](AppRoleWriteCustomSecretIDRequest.md) |  | 


[**AppRoleWriteCustomSecretIDResponse**](AppRoleWriteCustomSecretIDResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWritePeriod

> AppRoleWritePeriod(ctx, approleMountPath, roleName).AppRoleWritePeriodRequest(appRoleWritePeriodRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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

> AppRoleWritePolicies(ctx, approleMountPath, roleName).AppRoleWritePoliciesRequest(appRoleWritePoliciesRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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

> AppRoleWriteRole(ctx, approleMountPath, roleName).AppRoleWriteRoleRequest(appRoleWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## AppRoleWriteRoleID

> AppRoleWriteRoleID(ctx, approleMountPath, roleName).AppRoleWriteRoleIDRequest(appRoleWriteRoleIDRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteRoleIDRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteRoleID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteRoleIDRequest** | [**AppRoleWriteRoleIDRequest**](AppRoleWriteRoleIDRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretID

> AppRoleWriteSecretIDResponse AppRoleWriteSecretID(ctx, approleMountPath, roleName).AppRoleWriteSecretIDRequest(appRoleWriteSecretIDRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretID(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDRequest** | [**AppRoleWriteSecretIDRequest**](AppRoleWriteSecretIDRequest.md) |  | 


[**AppRoleWriteSecretIDResponse**](AppRoleWriteSecretIDResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDAccessorDestroy

> AppRoleWriteSecretIDAccessorDestroy(ctx, approleMountPath, roleName).AppRoleWriteSecretIDAccessorDestroyRequest(appRoleWriteSecretIDAccessorDestroyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDAccessorDestroyRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDAccessorDestroy(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDAccessorDestroyRequest** | [**AppRoleWriteSecretIDAccessorDestroyRequest**](AppRoleWriteSecretIDAccessorDestroyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDAccessorLookup

> AppRoleWriteSecretIDAccessorLookupResponse AppRoleWriteSecretIDAccessorLookup(ctx, approleMountPath, roleName).AppRoleWriteSecretIDAccessorLookupRequest(appRoleWriteSecretIDAccessorLookupRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDAccessorLookupRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDAccessorLookup(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDAccessorLookupRequest** | [**AppRoleWriteSecretIDAccessorLookupRequest**](AppRoleWriteSecretIDAccessorLookupRequest.md) |  | 


[**AppRoleWriteSecretIDAccessorLookupResponse**](AppRoleWriteSecretIDAccessorLookupResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDBoundCIDRs

> AppRoleWriteSecretIDBoundCIDRs(ctx, approleMountPath, roleName).AppRoleWriteSecretIDBoundCIDRsRequest(appRoleWriteSecretIDBoundCIDRsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDBoundCIDRsRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDBoundCIDRsRequest** | [**AppRoleWriteSecretIDBoundCIDRsRequest**](AppRoleWriteSecretIDBoundCIDRsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDDestroy

> AppRoleWriteSecretIDDestroy(ctx, approleMountPath, roleName).AppRoleWriteSecretIDDestroyRequest(appRoleWriteSecretIDDestroyRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDDestroyRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDDestroy(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDDestroyRequest** | [**AppRoleWriteSecretIDDestroyRequest**](AppRoleWriteSecretIDDestroyRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDLookup

> AppRoleWriteSecretIDLookupResponse AppRoleWriteSecretIDLookup(ctx, approleMountPath, roleName).AppRoleWriteSecretIDLookupRequest(appRoleWriteSecretIDLookupRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDLookupRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDLookup(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDLookupRequest** | [**AppRoleWriteSecretIDLookupRequest**](AppRoleWriteSecretIDLookupRequest.md) |  | 


[**AppRoleWriteSecretIDLookupResponse**](AppRoleWriteSecretIDLookupResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDNumUses

> AppRoleWriteSecretIDNumUses(ctx, approleMountPath, roleName).AppRoleWriteSecretIDNumUsesRequest(appRoleWriteSecretIDNumUsesRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDNumUsesRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDNumUses(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDNumUsesRequest** | [**AppRoleWriteSecretIDNumUsesRequest**](AppRoleWriteSecretIDNumUsesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteSecretIDTTL

> AppRoleWriteSecretIDTTL(ctx, approleMountPath, roleName).AppRoleWriteSecretIDTTLRequest(appRoleWriteSecretIDTTLRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteSecretIDTTLRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteSecretIDTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteSecretIDTTLRequest** | [**AppRoleWriteSecretIDTTLRequest**](AppRoleWriteSecretIDTTLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenBoundCIDRs

> AppRoleWriteTokenBoundCIDRs(ctx, approleMountPath, roleName).AppRoleWriteTokenBoundCIDRsRequest(appRoleWriteTokenBoundCIDRsRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteTokenBoundCIDRsRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteTokenBoundCIDRs(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteTokenBoundCIDRsRequest** | [**AppRoleWriteTokenBoundCIDRsRequest**](AppRoleWriteTokenBoundCIDRsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenMaxTTL

> AppRoleWriteTokenMaxTTL(ctx, approleMountPath, roleName).AppRoleWriteTokenMaxTTLRequest(appRoleWriteTokenMaxTTLRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteTokenMaxTTLRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteTokenMaxTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteTokenMaxTTLRequest** | [**AppRoleWriteTokenMaxTTLRequest**](AppRoleWriteTokenMaxTTLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AppRoleWriteTokenNumUses

> AppRoleWriteTokenNumUses(ctx, approleMountPath, roleName).AppRoleWriteTokenNumUsesRequest(appRoleWriteTokenNumUsesRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## AppRoleWriteTokenTTL

> AppRoleWriteTokenTTL(ctx, approleMountPath, roleName).AppRoleWriteTokenTTLRequest(appRoleWriteTokenTTLRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role. Must be less than 4096 bytes.

	request := schema.NewAppRoleWriteTokenTTLRequestWithDefaults()

	resp, err := client.Auth.AppRoleWriteTokenTTL(
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
**roleName** | **string** | Name of the role. Must be less than 4096 bytes. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleWriteTokenTTLRequest** | [**AppRoleWriteTokenTTLRequest**](AppRoleWriteTokenTTLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteAuthConfig

> AzureDeleteAuthConfig(ctx, azureMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AzureDeleteAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureDeleteAuthRole

> AzureDeleteAuthRole(ctx, azureMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.AzureDeleteAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureListAuthRoles

> AzureListAuthRoles(ctx, azureMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AzureListAuthRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureLogin

> AzureLogin(ctx, azureMountPath).AzureLoginRequest(azureLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAzureLoginRequestWithDefaults()

	resp, err := client.Auth.AzureLogin(
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



## AzureReadAuthConfig

> AzureReadAuthConfig(ctx, azureMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.AzureReadAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureReadAuthRole

> AzureReadAuthRole(ctx, azureMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.AzureReadAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteAuthConfig

> AzureWriteAuthConfig(ctx, azureMountPath).AzureWriteAuthConfigRequest(azureWriteAuthConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewAzureWriteAuthConfigRequestWithDefaults()

	resp, err := client.Auth.AzureWriteAuthConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureWriteAuthConfigRequest** | [**AzureWriteAuthConfigRequest**](AzureWriteAuthConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AzureWriteAuthRole

> AzureWriteAuthRole(ctx, azureMountPath, name).AzureWriteAuthRoleRequest(azureWriteAuthRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## CentrifyLogin

> CentrifyLogin(ctx, centrifyMountPath).CentrifyLoginRequest(centrifyLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCentrifyLoginRequestWithDefaults()

	resp, err := client.Auth.CentrifyLogin(
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



## CentrifyReadConfig

> CentrifyReadConfig(ctx, centrifyMountPath).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CentrifyReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CentrifyWriteConfig

> CentrifyWriteConfig(ctx, centrifyMountPath).CentrifyWriteConfigRequest(centrifyWriteConfigRequest).Execute()

This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCentrifyWriteConfigRequestWithDefaults()

	resp, err := client.Auth.CentrifyWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **centrifyWriteConfigRequest** | [**CentrifyWriteConfigRequest**](CentrifyWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesDelete

> CertificatesDelete(ctx, certMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate


	resp, err := client.Auth.CertificatesDelete(
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesDeleteCRL

> CertificatesDeleteCRL(ctx, certMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate


	resp, err := client.Auth.CertificatesDeleteCRL(
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesList

> CertificatesList(ctx, certMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CertificatesList(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesListCRLs

> CertificatesListCRLs(ctx, certMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CertificatesListCRLs(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesLogin

> CertificatesLogin(ctx, certMountPath).CertificatesLoginRequest(certificatesLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCertificatesLoginRequestWithDefaults()

	resp, err := client.Auth.CertificatesLogin(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificatesLoginRequest** | [**CertificatesLoginRequest**](CertificatesLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesRead

> CertificatesRead(ctx, certMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate


	resp, err := client.Auth.CertificatesRead(
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesReadCRL

> CertificatesReadCRL(ctx, certMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate


	resp, err := client.Auth.CertificatesReadCRL(
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesReadConfig

> CertificatesReadConfig(ctx, certMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CertificatesReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesWrite

> CertificatesWrite(ctx, certMountPath, name).CertificatesWriteRequest(certificatesWriteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	request := schema.NewCertificatesWriteRequestWithDefaults()

	resp, err := client.Auth.CertificatesWrite(
		context.Background(),
		name,
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificatesWriteRequest** | [**CertificatesWriteRequest**](CertificatesWriteRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesWriteCRL

> CertificatesWriteCRL(ctx, certMountPath, name).CertificatesWriteCRLRequest(certificatesWriteCRLRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the certificate

	request := schema.NewCertificatesWriteCRLRequestWithDefaults()

	resp, err := client.Auth.CertificatesWriteCRL(
		context.Background(),
		name,
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
**name** | **string** | The name of the certificate | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificatesWriteCRLRequest** | [**CertificatesWriteCRLRequest**](CertificatesWriteCRLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CertificatesWriteConfig

> CertificatesWriteConfig(ctx, certMountPath).CertificatesWriteConfigRequest(certificatesWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCertificatesWriteConfigRequestWithDefaults()

	resp, err := client.Auth.CertificatesWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificatesWriteConfigRequest** | [**CertificatesWriteConfigRequest**](CertificatesWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryDeleteConfig

> CloudFoundryDeleteConfig(ctx, cfMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CloudFoundryDeleteConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryDeleteRole

> CloudFoundryDeleteRole(ctx, cfMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.


	resp, err := client.Auth.CloudFoundryDeleteRole(
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
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryListRoles

> CloudFoundryListRoles(ctx, cfMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CloudFoundryListRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryLogin

> CloudFoundryLogin(ctx, cfMountPath).CloudFoundryLoginRequest(cloudFoundryLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCloudFoundryLoginRequestWithDefaults()

	resp, err := client.Auth.CloudFoundryLogin(
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



## CloudFoundryReadConfig

> CloudFoundryReadConfig(ctx, cfMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.CloudFoundryReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryReadRole

> CloudFoundryReadRole(ctx, cfMountPath, role).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | The name of the role.


	resp, err := client.Auth.CloudFoundryReadRole(
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
**role** | **string** | The name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryWriteConfig

> CloudFoundryWriteConfig(ctx, cfMountPath).CloudFoundryWriteConfigRequest(cloudFoundryWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewCloudFoundryWriteConfigRequestWithDefaults()

	resp, err := client.Auth.CloudFoundryWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFoundryWriteConfigRequest** | [**CloudFoundryWriteConfigRequest**](CloudFoundryWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CloudFoundryWriteRole

> CloudFoundryWriteRole(ctx, cfMountPath, role).CloudFoundryWriteRoleRequest(cloudFoundryWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## GitHubDeleteMapTeam

> GitHubDeleteMapTeam(ctx, githubMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping


	resp, err := client.Auth.GitHubDeleteMapTeam(
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubDeleteMapUser

> GitHubDeleteMapUser(ctx, githubMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping


	resp, err := client.Auth.GitHubDeleteMapUser(
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubLogin

> GitHubLogin(ctx, githubMountPath).GitHubLoginRequest(gitHubLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGitHubLoginRequestWithDefaults()

	resp, err := client.Auth.GitHubLogin(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitHubLoginRequest** | [**GitHubLoginRequest**](GitHubLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadConfig

> GitHubReadConfig(ctx, githubMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.GitHubReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadMapTeam

> GitHubReadMapTeam(ctx, githubMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping


	resp, err := client.Auth.GitHubReadMapTeam(
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadMapTeams

> GitHubReadMapTeams(ctx, githubMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	list := "list_example" // string | Return a list if `true`
	resp, err := client.Auth.GitHubReadMapTeams(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadMapUser

> GitHubReadMapUser(ctx, githubMountPath, key).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping


	resp, err := client.Auth.GitHubReadMapUser(
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadMapUsers

> GitHubReadMapUsers(ctx, githubMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	list := "list_example" // string | Return a list if `true`
	resp, err := client.Auth.GitHubReadMapUsers(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubWriteConfig

> GitHubWriteConfig(ctx, githubMountPath).GitHubWriteConfigRequest(gitHubWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGitHubWriteConfigRequestWithDefaults()

	resp, err := client.Auth.GitHubWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitHubWriteConfigRequest** | [**GitHubWriteConfigRequest**](GitHubWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubWriteMapTeam

> GitHubWriteMapTeam(ctx, githubMountPath, key).GitHubWriteMapTeamRequest(gitHubWriteMapTeamRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the teams mapping

	request := schema.NewGitHubWriteMapTeamRequestWithDefaults()

	resp, err := client.Auth.GitHubWriteMapTeam(
		context.Background(),
		key,
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
**key** | **string** | Key for the teams mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitHubWriteMapTeamRequest** | [**GitHubWriteMapTeamRequest**](GitHubWriteMapTeamRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubWriteMapUser

> GitHubWriteMapUser(ctx, githubMountPath, key).GitHubWriteMapUserRequest(gitHubWriteMapUserRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	key := "key_example" // string | Key for the users mapping

	request := schema.NewGitHubWriteMapUserRequestWithDefaults()

	resp, err := client.Auth.GitHubWriteMapUser(
		context.Background(),
		key,
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
**key** | **string** | Key for the users mapping | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitHubWriteMapUserRequest** | [**GitHubWriteMapUserRequest**](GitHubWriteMapUserRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudDeleteRole

> GoogleCloudDeleteRole(ctx, gcpMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.GoogleCloudDeleteRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRoles

> GoogleCloudListRoles(ctx, gcpMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.GoogleCloudListRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudListRoles2

> GoogleCloudListRoles2(ctx, gcpMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.GoogleCloudListRoles2(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudLogin

> GoogleCloudLogin(ctx, gcpMountPath).GoogleCloudLoginRequest(googleCloudLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGoogleCloudLoginRequestWithDefaults()

	resp, err := client.Auth.GoogleCloudLogin(
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



## GoogleCloudReadAuthConfig

> GoogleCloudReadAuthConfig(ctx, gcpMountPath).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.GoogleCloudReadAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudReadRole

> GoogleCloudReadRole(ctx, gcpMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.GoogleCloudReadRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteAuthConfig

> GoogleCloudWriteAuthConfig(ctx, gcpMountPath).GoogleCloudWriteAuthConfigRequest(googleCloudWriteAuthConfigRequest).Execute()

Configure credentials used to query the GCP IAM API to verify authenticating service accounts

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewGoogleCloudWriteAuthConfigRequestWithDefaults()

	resp, err := client.Auth.GoogleCloudWriteAuthConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleCloudWriteAuthConfigRequest** | [**GoogleCloudWriteAuthConfigRequest**](GoogleCloudWriteAuthConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRole

> GoogleCloudWriteRole(ctx, gcpMountPath, name).GoogleCloudWriteRoleRequest(googleCloudWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## GoogleCloudWriteRoleLabels

> GoogleCloudWriteRoleLabels(ctx, gcpMountPath, name).GoogleCloudWriteRoleLabelsRequest(googleCloudWriteRoleLabelsRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewGoogleCloudWriteRoleLabelsRequestWithDefaults()

	resp, err := client.Auth.GoogleCloudWriteRoleLabels(
		context.Background(),
		name,
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteRoleLabelsRequest** | [**GoogleCloudWriteRoleLabelsRequest**](GoogleCloudWriteRoleLabelsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GoogleCloudWriteRoleServiceAccounts

> GoogleCloudWriteRoleServiceAccounts(ctx, gcpMountPath, name).GoogleCloudWriteRoleServiceAccountsRequest(googleCloudWriteRoleServiceAccountsRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults()

	resp, err := client.Auth.GoogleCloudWriteRoleServiceAccounts(
		context.Background(),
		name,
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudWriteRoleServiceAccountsRequest** | [**GoogleCloudWriteRoleServiceAccountsRequest**](GoogleCloudWriteRoleServiceAccountsRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTDeleteRole

> JWTDeleteRole(ctx, jwtMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.JWTDeleteRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTListRoles

> JWTListRoles(ctx, jwtMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.JWTListRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTLogin

> JWTLogin(ctx, jwtMountPath).JWTLoginRequest(jWTLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewJWTLoginRequestWithDefaults()

	resp, err := client.Auth.JWTLogin(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTLoginRequest** | [**JWTLoginRequest**](JWTLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTReadConfig

> JWTReadConfig(ctx, jwtMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.JWTReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTReadOIDCCallback

> JWTReadOIDCCallback(ctx, jwtMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.JWTReadOIDCCallback(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTReadRole

> JWTReadRole(ctx, jwtMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.JWTReadRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTWriteConfig

> JWTWriteConfig(ctx, jwtMountPath).JWTWriteConfigRequest(jWTWriteConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewJWTWriteConfigRequestWithDefaults()

	resp, err := client.Auth.JWTWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTWriteConfigRequest** | [**JWTWriteConfigRequest**](JWTWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTWriteOIDCAuthURL

> JWTWriteOIDCAuthURL(ctx, jwtMountPath).JWTWriteOIDCAuthURLRequest(jWTWriteOIDCAuthURLRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewJWTWriteOIDCAuthURLRequestWithDefaults()

	resp, err := client.Auth.JWTWriteOIDCAuthURL(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTWriteOIDCAuthURLRequest** | [**JWTWriteOIDCAuthURLRequest**](JWTWriteOIDCAuthURLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTWriteOIDCCallback

> JWTWriteOIDCCallback(ctx, jwtMountPath).JWTWriteOIDCCallbackRequest(jWTWriteOIDCCallbackRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewJWTWriteOIDCCallbackRequestWithDefaults()

	resp, err := client.Auth.JWTWriteOIDCCallback(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jWTWriteOIDCCallbackRequest** | [**JWTWriteOIDCCallbackRequest**](JWTWriteOIDCCallbackRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## JWTWriteRole

> JWTWriteRole(ctx, jwtMountPath, name).JWTWriteRoleRequest(jWTWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewJWTWriteRoleRequestWithDefaults()

	resp, err := client.Auth.JWTWriteRole(
		context.Background(),
		name,
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jWTWriteRoleRequest** | [**JWTWriteRoleRequest**](JWTWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosDeleteGroup

> KerberosDeleteGroup(ctx, kerberosMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.


	resp, err := client.Auth.KerberosDeleteGroup(
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosListGroups

> KerberosListGroups(ctx, kerberosMountPath).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.KerberosListGroups(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosLogin

> KerberosLogin(ctx, kerberosMountPath).KerberosLoginRequest(kerberosLoginRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKerberosLoginRequestWithDefaults()

	resp, err := client.Auth.KerberosLogin(
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



## KerberosReadConfig

> KerberosReadConfig(ctx, kerberosMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.KerberosReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosReadGroup

> KerberosReadGroup(ctx, kerberosMountPath, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.


	resp, err := client.Auth.KerberosReadGroup(
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosReadLDAPConfig

> KerberosReadLDAPConfig(ctx, kerberosMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.KerberosReadLDAPConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosWriteConfig

> KerberosWriteConfig(ctx, kerberosMountPath).KerberosWriteConfigRequest(kerberosWriteConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKerberosWriteConfigRequestWithDefaults()

	resp, err := client.Auth.KerberosWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosWriteConfigRequest** | [**KerberosWriteConfigRequest**](KerberosWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KerberosWriteGroup

> KerberosWriteGroup(ctx, kerberosMountPath, name).KerberosWriteGroupRequest(kerberosWriteGroupRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## KerberosWriteLDAPConfig

> KerberosWriteLDAPConfig(ctx, kerberosMountPath).KerberosWriteLDAPConfigRequest(kerberosWriteLDAPConfigRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKerberosWriteLDAPConfigRequestWithDefaults()

	resp, err := client.Auth.KerberosWriteLDAPConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kerberosWriteLDAPConfigRequest** | [**KerberosWriteLDAPConfigRequest**](KerberosWriteLDAPConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesDeleteAuthRole

> KubernetesDeleteAuthRole(ctx, kubernetesMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.KubernetesDeleteAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesListAuthRoles

> KubernetesListAuthRoles(ctx, kubernetesMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.KubernetesListAuthRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesLogin

> KubernetesLogin(ctx, kubernetesMountPath).KubernetesLoginRequest(kubernetesLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKubernetesLoginRequestWithDefaults()

	resp, err := client.Auth.KubernetesLogin(
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



## KubernetesReadAuthConfig

> KubernetesReadAuthConfig(ctx, kubernetesMountPath).Execute()

Configures the JWT Public Key and Kubernetes API information.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.KubernetesReadAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesReadAuthRole

> KubernetesReadAuthRole(ctx, kubernetesMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.KubernetesReadAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteAuthConfig

> KubernetesWriteAuthConfig(ctx, kubernetesMountPath).KubernetesWriteAuthConfigRequest(kubernetesWriteAuthConfigRequest).Execute()

Configures the JWT Public Key and Kubernetes API information.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewKubernetesWriteAuthConfigRequestWithDefaults()

	resp, err := client.Auth.KubernetesWriteAuthConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesWriteAuthConfigRequest** | [**KubernetesWriteAuthConfigRequest**](KubernetesWriteAuthConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## KubernetesWriteAuthRole

> KubernetesWriteAuthRole(ctx, kubernetesMountPath, name).KubernetesWriteAuthRoleRequest(kubernetesWriteAuthRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## LDAPDeleteGroup

> LDAPDeleteGroup(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.


	resp, err := client.Auth.LDAPDeleteGroup(
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPDeleteUser

> LDAPDeleteUser(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.


	resp, err := client.Auth.LDAPDeleteUser(
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPListGroups

> LDAPListGroups(ctx, ldapMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.LDAPListGroups(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPListUsers

> LDAPListUsers(ctx, ldapMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.LDAPListUsers(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPLogin

> LDAPLogin(ctx, ldapMountPath, username).LDAPLoginRequest(lDAPLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | DN (distinguished name) to be used for login.

	request := schema.NewLDAPLoginRequestWithDefaults()

	resp, err := client.Auth.LDAPLogin(
		context.Background(),
		username,
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
**username** | **string** | DN (distinguished name) to be used for login. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPLoginRequest** | [**LDAPLoginRequest**](LDAPLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadAuthConfig

> LDAPReadAuthConfig(ctx, ldapMountPath).Execute()

Configure the LDAP server to connect to, along with its options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.LDAPReadAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadGroup

> LDAPReadGroup(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.


	resp, err := client.Auth.LDAPReadGroup(
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPReadUser

> LDAPReadUser(ctx, ldapMountPath, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.


	resp, err := client.Auth.LDAPReadUser(
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteAuthConfig

> LDAPWriteAuthConfig(ctx, ldapMountPath).LDAPWriteAuthConfigRequest(lDAPWriteAuthConfigRequest).Execute()

Configure the LDAP server to connect to, along with its options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewLDAPWriteAuthConfigRequestWithDefaults()

	resp, err := client.Auth.LDAPWriteAuthConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPWriteAuthConfigRequest** | [**LDAPWriteAuthConfigRequest**](LDAPWriteAuthConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteGroup

> LDAPWriteGroup(ctx, ldapMountPath, name).LDAPWriteGroupRequest(lDAPWriteGroupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP group.

	request := schema.NewLDAPWriteGroupRequestWithDefaults()

	resp, err := client.Auth.LDAPWriteGroup(
		context.Background(),
		name,
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
**name** | **string** | Name of the LDAP group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPWriteGroupRequest** | [**LDAPWriteGroupRequest**](LDAPWriteGroupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LDAPWriteUser

> LDAPWriteUser(ctx, ldapMountPath, name).LDAPWriteUserRequest(lDAPWriteUserRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the LDAP user.

	request := schema.NewLDAPWriteUserRequestWithDefaults()

	resp, err := client.Auth.LDAPWriteUser(
		context.Background(),
		name,
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
**name** | **string** | Name of the LDAP user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPWriteUserRequest** | [**LDAPWriteUserRequest**](LDAPWriteUserRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIDeleteConfig

> OCIDeleteConfig(ctx, ociMountPath).Execute()

Manages the configuration for the Vault Auth Plugin.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OCIDeleteConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIDeleteRole

> OCIDeleteRole(ctx, ociMountPath, role).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.


	resp, err := client.Auth.OCIDeleteRole(
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIListRoles

> OCIListRoles(ctx, ociMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OCIListRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCILoginWithRole

> OCILoginWithRole(ctx, ociMountPath, role).OCILoginWithRoleRequest(oCILoginWithRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	request := schema.NewOCILoginWithRoleRequestWithDefaults()

	resp, err := client.Auth.OCILoginWithRole(
		context.Background(),
		role,
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCILoginWithRoleRequest** | [**OCILoginWithRoleRequest**](OCILoginWithRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIReadConfig

> OCIReadConfig(ctx, ociMountPath).Execute()

Manages the configuration for the Vault Auth Plugin.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OCIReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIReadRole

> OCIReadRole(ctx, ociMountPath, role).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.


	resp, err := client.Auth.OCIReadRole(
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIWriteConfig

> OCIWriteConfig(ctx, ociMountPath).OCIWriteConfigRequest(oCIWriteConfigRequest).Execute()

Manages the configuration for the Vault Auth Plugin.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOCIWriteConfigRequestWithDefaults()

	resp, err := client.Auth.OCIWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCIWriteConfigRequest** | [**OCIWriteConfigRequest**](OCIWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OCIWriteRole

> OCIWriteRole(ctx, ociMountPath, role).OCIWriteRoleRequest(oCIWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	role := "role_example" // string | Name of the role.

	request := schema.NewOCIWriteRoleRequestWithDefaults()

	resp, err := client.Auth.OCIWriteRole(
		context.Background(),
		role,
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
**role** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCIWriteRoleRequest** | [**OCIWriteRoleRequest**](OCIWriteRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteAuthRole

> OIDCDeleteAuthRole(ctx, name, oidcMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.OIDCDeleteAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListAuthRoles

> OIDCListAuthRoles(ctx, oidcMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OIDCListAuthRoles(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCLogin

> OIDCLogin(ctx, oidcMountPath).OIDCLoginRequest(oIDCLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOIDCLoginRequestWithDefaults()

	resp, err := client.Auth.OIDCLogin(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCLoginRequest** | [**OIDCLoginRequest**](OIDCLoginRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadAuthConfig

> OIDCReadAuthConfig(ctx, oidcMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OIDCReadAuthConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadAuthRole

> OIDCReadAuthRole(ctx, name, oidcMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.


	resp, err := client.Auth.OIDCReadAuthRole(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadCallback

> OIDCReadCallback(ctx, oidcMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OIDCReadCallback(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteAuthConfig

> OIDCWriteAuthConfig(ctx, oidcMountPath).OIDCWriteAuthConfigRequest(oIDCWriteAuthConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOIDCWriteAuthConfigRequestWithDefaults()

	resp, err := client.Auth.OIDCWriteAuthConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCWriteAuthConfigRequest** | [**OIDCWriteAuthConfigRequest**](OIDCWriteAuthConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteAuthRole

> OIDCWriteAuthRole(ctx, name, oidcMountPath).OIDCWriteAuthRoleRequest(oIDCWriteAuthRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role.

	request := schema.NewOIDCWriteAuthRoleRequestWithDefaults()

	resp, err := client.Auth.OIDCWriteAuthRole(
		context.Background(),
		name,
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
**name** | **string** | Name of the role. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteAuthRoleRequest** | [**OIDCWriteAuthRoleRequest**](OIDCWriteAuthRoleRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteAuthURL

> OIDCWriteAuthURL(ctx, oidcMountPath).OIDCWriteAuthURLRequest(oIDCWriteAuthURLRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOIDCWriteAuthURLRequestWithDefaults()

	resp, err := client.Auth.OIDCWriteAuthURL(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCWriteAuthURLRequest** | [**OIDCWriteAuthURLRequest**](OIDCWriteAuthURLRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteCallback

> OIDCWriteCallback(ctx, oidcMountPath).OIDCWriteCallbackRequest(oIDCWriteCallbackRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOIDCWriteCallbackRequestWithDefaults()

	resp, err := client.Auth.OIDCWriteCallback(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCWriteCallbackRequest** | [**OIDCWriteCallbackRequest**](OIDCWriteCallbackRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaDeleteGroup

> OktaDeleteGroup(ctx, name, oktaMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.


	resp, err := client.Auth.OktaDeleteGroup(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaDeleteUser

> OktaDeleteUser(ctx, name, oktaMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.


	resp, err := client.Auth.OktaDeleteUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaListGroups

> OktaListGroups(ctx, oktaMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OktaListGroups(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaListUsers

> OktaListUsers(ctx, oktaMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OktaListUsers(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaLogin

> OktaLogin(ctx, oktaMountPath, username).OktaLoginRequest(oktaLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## OktaReadConfig

> OktaReadConfig(ctx, oktaMountPath).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.OktaReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaReadGroup

> OktaReadGroup(ctx, name, oktaMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the Okta group.


	resp, err := client.Auth.OktaReadGroup(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaReadUser

> OktaReadUser(ctx, name, oktaMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the user.


	resp, err := client.Auth.OktaReadUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaVerify

> OktaVerify(ctx, nonce, oktaMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	nonce := "nonce_example" // string | Nonce provided during a login request to retrieve the number verification challenge for the matching request.


	resp, err := client.Auth.OktaVerify(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaWriteConfig

> OktaWriteConfig(ctx, oktaMountPath).OktaWriteConfigRequest(oktaWriteConfigRequest).Execute()

This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewOktaWriteConfigRequestWithDefaults()

	resp, err := client.Auth.OktaWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oktaWriteConfigRequest** | [**OktaWriteConfigRequest**](OktaWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OktaWriteGroup

> OktaWriteGroup(ctx, name, oktaMountPath).OktaWriteGroupRequest(oktaWriteGroupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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

> OktaWriteUser(ctx, name, oktaMountPath).OktaWriteUserRequest(oktaWriteUserRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## RadiusDeleteUser

> RadiusDeleteUser(ctx, name, radiusMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.


	resp, err := client.Auth.RadiusDeleteUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusListUsers

> RadiusListUsers(ctx, radiusMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.RadiusListUsers(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusLogin

> RadiusLogin(ctx, radiusMountPath).RadiusLoginRequest(radiusLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewRadiusLoginRequestWithDefaults()

	resp, err := client.Auth.RadiusLogin(
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

> RadiusLoginWithUsername(ctx, radiusMountPath, urlusername).RadiusLoginWithUsernameRequest(radiusLoginWithUsernameRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## RadiusReadConfig

> RadiusReadConfig(ctx, radiusMountPath).Execute()

Configure the RADIUS server to connect to, along with its options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.RadiusReadConfig(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusReadUser

> RadiusReadUser(ctx, name, radiusMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the RADIUS user.


	resp, err := client.Auth.RadiusReadUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusWriteConfig

> RadiusWriteConfig(ctx, radiusMountPath).RadiusWriteConfigRequest(radiusWriteConfigRequest).Execute()

Configure the RADIUS server to connect to, along with its options.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewRadiusWriteConfigRequestWithDefaults()

	resp, err := client.Auth.RadiusWriteConfig(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusWriteConfigRequest** | [**RadiusWriteConfigRequest**](RadiusWriteConfigRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RadiusWriteUser

> RadiusWriteUser(ctx, name, radiusMountPath).RadiusWriteUserRequest(radiusWriteUserRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## TokenDeleteRole

> TokenDeleteRole(ctx, roleName, tokenMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> TokenListAccessors(ctx, tokenMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## TokenListRoles

> TokenListRoles(ctx, tokenMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## TokenReadLookup

> TokenReadLookup(ctx, tokenMountPath).Execute()

This endpoint will lookup a token and its properties.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.TokenReadLookup(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenReadLookupSelf

> TokenReadLookupSelf(ctx, tokenMountPath).Execute()

This endpoint will lookup a token and its properties.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.TokenReadLookupSelf(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenReadRole

> TokenReadRole(ctx, roleName, tokenMountPath).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> TokenRenew(ctx, tokenMountPath).TokenRenewRequest(tokenRenewRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewRequest** | [**TokenRenewRequest**](TokenRenewRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRenewAccessor

> TokenRenewAccessor(ctx, tokenMountPath).TokenRenewAccessorRequest(tokenRenewAccessorRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewAccessorRequest** | [**TokenRenewAccessorRequest**](TokenRenewAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRenewSelf

> TokenRenewSelf(ctx, tokenMountPath).TokenRenewSelfRequest(tokenRenewSelfRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRenewSelfRequest** | [**TokenRenewSelfRequest**](TokenRenewSelfRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRevoke

> TokenRevoke(ctx, tokenMountPath).TokenRevokeRequest(tokenRevokeRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRevokeAccessor

> TokenRevokeAccessor(ctx, tokenMountPath).TokenRevokeAccessorRequest(tokenRevokeAccessorRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeAccessorRequest** | [**TokenRevokeAccessorRequest**](TokenRevokeAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRevokeOrphan

> TokenRevokeOrphan(ctx, tokenMountPath).TokenRevokeOrphanRequest(tokenRevokeOrphanRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRevokeOrphanRequest** | [**TokenRevokeOrphanRequest**](TokenRevokeOrphanRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenRevokeSelf

> TokenRevokeSelf(ctx, tokenMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenTidy

> TokenTidy(ctx, tokenMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteCreate

> TokenWriteCreate(ctx, tokenMountPath).TokenWriteCreateRequest(tokenWriteCreateRequest).Format(format).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTokenWriteCreateRequestWithDefaults()

	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenWriteCreate(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenWriteCreateRequest** | [**TokenWriteCreateRequest**](TokenWriteCreateRequest.md) |  | 
 **format** | **string** | Return json formatted output | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteCreateOrphan

> TokenWriteCreateOrphan(ctx, tokenMountPath).TokenWriteCreateOrphanRequest(tokenWriteCreateOrphanRequest).Format(format).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTokenWriteCreateOrphanRequestWithDefaults()

	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenWriteCreateOrphan(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenWriteCreateOrphanRequest** | [**TokenWriteCreateOrphanRequest**](TokenWriteCreateOrphanRequest.md) |  | 
 **format** | **string** | Return json formatted output | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteCreateWithRole

> TokenWriteCreateWithRole(ctx, roleName, tokenMountPath).TokenWriteCreateWithRoleRequest(tokenWriteCreateWithRoleRequest).Format(format).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	roleName := "roleName_example" // string | Name of the role

	request := schema.NewTokenWriteCreateWithRoleRequestWithDefaults()

	format := "format_example" // string | Return json formatted output
	resp, err := client.Auth.TokenWriteCreateWithRole(
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

 **tokenWriteCreateWithRoleRequest** | [**TokenWriteCreateWithRoleRequest**](TokenWriteCreateWithRoleRequest.md) |  | 
 **format** | **string** | Return json formatted output | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteLookup

> TokenWriteLookup(ctx, tokenMountPath).TokenWriteLookupRequest(tokenWriteLookupRequest).Execute()

This endpoint will lookup a token and its properties.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTokenWriteLookupRequestWithDefaults()

	resp, err := client.Auth.TokenWriteLookup(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenWriteLookupRequest** | [**TokenWriteLookupRequest**](TokenWriteLookupRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteLookupAccessor

> TokenWriteLookupAccessor(ctx, tokenMountPath).TokenWriteLookupAccessorRequest(tokenWriteLookupAccessorRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTokenWriteLookupAccessorRequestWithDefaults()

	resp, err := client.Auth.TokenWriteLookupAccessor(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenWriteLookupAccessorRequest** | [**TokenWriteLookupAccessorRequest**](TokenWriteLookupAccessorRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteLookupSelf

> TokenWriteLookupSelf(ctx, tokenMountPath).TokenWriteLookupSelfRequest(tokenWriteLookupSelfRequest).Execute()

This endpoint will lookup a token and its properties.

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}


	request := schema.NewTokenWriteLookupSelfRequestWithDefaults()

	resp, err := client.Auth.TokenWriteLookupSelf(
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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenWriteLookupSelfRequest** | [**TokenWriteLookupSelfRequest**](TokenWriteLookupSelfRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## TokenWriteRole

> TokenWriteRole(ctx, roleName, tokenMountPath).TokenWriteRoleRequest(tokenWriteRoleRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> UserpassDeleteUser(ctx, username, userpassMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.


	resp, err := client.Auth.UserpassDeleteUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassListUsers

> UserpassListUsers(ctx, userpassMountPath).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}



	resp, err := client.Auth.UserpassListUsers(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassLogin

> UserpassLogin(ctx, username, userpassMountPath).UserpassLoginRequest(userpassLoginRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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

> UserpassReadUser(ctx, username, userpassMountPath).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.


	resp, err := client.Auth.UserpassReadUser(
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

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassWriteUser

> UserpassWriteUser(ctx, username, userpassMountPath).UserpassWriteUserRequest(userpassWriteUserRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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
	)
	if err != nil {
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



## UserpassWriteUserPassword

> UserpassWriteUserPassword(ctx, username, userpassMountPath).UserpassWriteUserPasswordRequest(userpassWriteUserPasswordRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	request := schema.NewUserpassWriteUserPasswordRequestWithDefaults()

	resp, err := client.Auth.UserpassWriteUserPassword(
		context.Background(),
		username,
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassWriteUserPasswordRequest** | [**UserpassWriteUserPasswordRequest**](UserpassWriteUserPasswordRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UserpassWriteUserPolicies

> UserpassWriteUserPolicies(ctx, username, userpassMountPath).UserpassWriteUserPoliciesRequest(userpassWriteUserPoliciesRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	username := "username_example" // string | Username for this user.

	request := schema.NewUserpassWriteUserPoliciesRequestWithDefaults()

	resp, err := client.Auth.UserpassWriteUserPolicies(
		context.Background(),
		username,
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
**username** | **string** | Username for this user. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userpassWriteUserPoliciesRequest** | [**UserpassWriteUserPoliciesRequest**](UserpassWriteUserPoliciesRequest.md) |  | 


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



