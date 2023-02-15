# Auth

Method | HTTP request | Description
------------- | ------------- | -------------
[**AWSConfigDeleteCertificate**](AuthApi.md#AWSConfigDeleteCertificate) | **Delete** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigDeleteClient**](AuthApi.md#AWSConfigDeleteClient) | **Delete** /auth/{aws_mount_path}/config/client | 
[**AWSConfigDeleteIdentityAccessList**](AuthApi.md#AWSConfigDeleteIdentityAccessList) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigDeleteIdentityWhiteList**](AuthApi.md#AWSConfigDeleteIdentityWhiteList) | **Delete** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigDeleteRoleTagBlackList**](AuthApi.md#AWSConfigDeleteRoleTagBlackList) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigDeleteRoleTagDenyList**](AuthApi.md#AWSConfigDeleteRoleTagDenyList) | **Delete** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigDeleteSecurityTokenServiceAccount**](AuthApi.md#AWSConfigDeleteSecurityTokenServiceAccount) | **Delete** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSConfigListCertificates**](AuthApi.md#AWSConfigListCertificates) | **Get** /auth/{aws_mount_path}/config/certificates | 
[**AWSConfigListSecurityTokenService**](AuthApi.md#AWSConfigListSecurityTokenService) | **Get** /auth/{aws_mount_path}/config/sts | 
[**AWSConfigReadCertificate**](AuthApi.md#AWSConfigReadCertificate) | **Get** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigReadClient**](AuthApi.md#AWSConfigReadClient) | **Get** /auth/{aws_mount_path}/config/client | 
[**AWSConfigReadIdentity**](AuthApi.md#AWSConfigReadIdentity) | **Get** /auth/{aws_mount_path}/config/identity | 
[**AWSConfigReadIdentityAccessList**](AuthApi.md#AWSConfigReadIdentityAccessList) | **Get** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigReadIdentityWhiteList**](AuthApi.md#AWSConfigReadIdentityWhiteList) | **Get** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigReadRoleTagBlackList**](AuthApi.md#AWSConfigReadRoleTagBlackList) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigReadRoleTagDenyList**](AuthApi.md#AWSConfigReadRoleTagDenyList) | **Get** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigReadSecurityTokenServiceAccount**](AuthApi.md#AWSConfigReadSecurityTokenServiceAccount) | **Get** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSConfigRotateRoot**](AuthApi.md#AWSConfigRotateRoot) | **Post** /auth/{aws_mount_path}/config/rotate-root | 
[**AWSConfigWriteCertificate**](AuthApi.md#AWSConfigWriteCertificate) | **Post** /auth/{aws_mount_path}/config/certificate/{cert_name} | 
[**AWSConfigWriteClient**](AuthApi.md#AWSConfigWriteClient) | **Post** /auth/{aws_mount_path}/config/client | 
[**AWSConfigWriteIdentity**](AuthApi.md#AWSConfigWriteIdentity) | **Post** /auth/{aws_mount_path}/config/identity | 
[**AWSConfigWriteIdentityAccessList**](AuthApi.md#AWSConfigWriteIdentityAccessList) | **Post** /auth/{aws_mount_path}/config/tidy/identity-accesslist | 
[**AWSConfigWriteIdentityWhiteList**](AuthApi.md#AWSConfigWriteIdentityWhiteList) | **Post** /auth/{aws_mount_path}/config/tidy/identity-whitelist | 
[**AWSConfigWriteRoleTagBlackList**](AuthApi.md#AWSConfigWriteRoleTagBlackList) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-blacklist | 
[**AWSConfigWriteRoleTagDenyList**](AuthApi.md#AWSConfigWriteRoleTagDenyList) | **Post** /auth/{aws_mount_path}/config/tidy/roletag-denylist | 
[**AWSConfigWriteSecurityTokenServiceAccount**](AuthApi.md#AWSConfigWriteSecurityTokenServiceAccount) | **Post** /auth/{aws_mount_path}/config/sts/{account_id} | 
[**AWSDeleteAuthRole**](AuthApi.md#AWSDeleteAuthRole) | **Delete** /auth/{aws_mount_path}/role/{role} | 
[**AWSDeleteIdentityAccessListFor**](AuthApi.md#AWSDeleteIdentityAccessListFor) | **Delete** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AWSDeleteIdentityWhiteListFor**](AuthApi.md#AWSDeleteIdentityWhiteListFor) | **Delete** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AWSDeleteRoleTagBlackListFor**](AuthApi.md#AWSDeleteRoleTagBlackListFor) | **Delete** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSDeleteRoleTagDenyListFor**](AuthApi.md#AWSDeleteRoleTagDenyListFor) | **Delete** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSListAuthRoles**](AuthApi.md#AWSListAuthRoles) | **Get** /auth/{aws_mount_path}/role | 
[**AWSListAuthRoles2**](AuthApi.md#AWSListAuthRoles2) | **Get** /auth/{aws_mount_path}/roles | 
[**AWSListIdentityAccessList**](AuthApi.md#AWSListIdentityAccessList) | **Get** /auth/{aws_mount_path}/identity-accesslist | 
[**AWSListIdentityWhiteList**](AuthApi.md#AWSListIdentityWhiteList) | **Get** /auth/{aws_mount_path}/identity-whitelist | 
[**AWSListRoleTagBlackList**](AuthApi.md#AWSListRoleTagBlackList) | **Get** /auth/{aws_mount_path}/roletag-blacklist | 
[**AWSListRoleTagDenyList**](AuthApi.md#AWSListRoleTagDenyList) | **Get** /auth/{aws_mount_path}/roletag-denylist | 
[**AWSLogin**](AuthApi.md#AWSLogin) | **Post** /auth/{aws_mount_path}/login | 
[**AWSReadAuthRole**](AuthApi.md#AWSReadAuthRole) | **Get** /auth/{aws_mount_path}/role/{role} | 
[**AWSReadIdentityAccessListFor**](AuthApi.md#AWSReadIdentityAccessListFor) | **Get** /auth/{aws_mount_path}/identity-accesslist/{instance_id} | 
[**AWSReadIdentityWhiteListFor**](AuthApi.md#AWSReadIdentityWhiteListFor) | **Get** /auth/{aws_mount_path}/identity-whitelist/{instance_id} | 
[**AWSReadRoleTagBlackListFor**](AuthApi.md#AWSReadRoleTagBlackListFor) | **Get** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSReadRoleTagDenyListFor**](AuthApi.md#AWSReadRoleTagDenyListFor) | **Get** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSWriteAuthRole**](AuthApi.md#AWSWriteAuthRole) | **Post** /auth/{aws_mount_path}/role/{role} | 
[**AWSWriteAuthRoleTag**](AuthApi.md#AWSWriteAuthRoleTag) | **Post** /auth/{aws_mount_path}/role/{role}/tag | 
[**AWSWriteIdentityAccessListTidySettings**](AuthApi.md#AWSWriteIdentityAccessListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/identity-accesslist | 
[**AWSWriteIdentityWhiteListTidySettings**](AuthApi.md#AWSWriteIdentityWhiteListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/identity-whitelist | 
[**AWSWriteRoleTagBlackListFor**](AuthApi.md#AWSWriteRoleTagBlackListFor) | **Post** /auth/{aws_mount_path}/roletag-blacklist/{role_tag} | 
[**AWSWriteRoleTagBlackListTidySettings**](AuthApi.md#AWSWriteRoleTagBlackListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/roletag-blacklist | 
[**AWSWriteRoleTagDenyListFor**](AuthApi.md#AWSWriteRoleTagDenyListFor) | **Post** /auth/{aws_mount_path}/roletag-denylist/{role_tag} | 
[**AWSWriteRoleTagDenyListTidySettings**](AuthApi.md#AWSWriteRoleTagDenyListTidySettings) | **Post** /auth/{aws_mount_path}/tidy/roletag-denylist | 
[**AliCloudDeleteAuthRole**](AuthApi.md#AliCloudDeleteAuthRole) | **Delete** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudListAuthRoles**](AuthApi.md#AliCloudListAuthRoles) | **Get** /auth/{alicloud_mount_path}/role | Lists all the roles that are registered with Vault.
[**AliCloudListAuthRoles2**](AuthApi.md#AliCloudListAuthRoles2) | **Get** /auth/{alicloud_mount_path}/roles | Lists all the roles that are registered with Vault.
[**AliCloudLogin**](AuthApi.md#AliCloudLogin) | **Post** /auth/{alicloud_mount_path}/login | Authenticates an RAM entity with Vault.
[**AliCloudReadAuthRole**](AuthApi.md#AliCloudReadAuthRole) | **Get** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AliCloudWriteAuthRole**](AuthApi.md#AliCloudWriteAuthRole) | **Post** /auth/{alicloud_mount_path}/role/{role} | Create a role and associate policies to it.
[**AppRoleDeleteBindSecretID**](AuthApi.md#AppRoleDeleteBindSecretID) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleDeleteBoundCIDRList**](AuthApi.md#AppRoleDeleteBoundCIDRList) | **Delete** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleDeletePeriod**](AuthApi.md#AppRoleDeletePeriod) | **Delete** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleDeletePolicies**](AuthApi.md#AppRoleDeletePolicies) | **Delete** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleDeleteRole**](AuthApi.md#AppRoleDeleteRole) | **Delete** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleDeleteSecretIDAccessorDestroy**](AuthApi.md#AppRoleDeleteSecretIDAccessorDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**AppRoleDeleteSecretIDBoundCIDRs**](AuthApi.md#AppRoleDeleteSecretIDBoundCIDRs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleDeleteSecretIDDestroy**](AuthApi.md#AppRoleDeleteSecretIDDestroy) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | 
[**AppRoleDeleteSecretIDNumUses**](AuthApi.md#AppRoleDeleteSecretIDNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleDeleteSecretIDTTL**](AuthApi.md#AppRoleDeleteSecretIDTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleDeleteTokenBoundCIDRs**](AuthApi.md#AppRoleDeleteTokenBoundCIDRs) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleDeleteTokenMaxTTL**](AuthApi.md#AppRoleDeleteTokenMaxTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleDeleteTokenNumUses**](AuthApi.md#AppRoleDeleteTokenNumUses) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleDeleteTokenTTL**](AuthApi.md#AppRoleDeleteTokenTTL) | **Delete** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleListRoles**](AuthApi.md#AppRoleListRoles) | **Get** /auth/{approle_mount_path}/role | 
[**AppRoleListSecretID**](AuthApi.md#AppRoleListSecretID) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleLogin**](AuthApi.md#AppRoleLogin) | **Post** /auth/{approle_mount_path}/login | 
[**AppRoleReadBindSecretID**](AuthApi.md#AppRoleReadBindSecretID) | **Get** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleReadBoundCIDRList**](AuthApi.md#AppRoleReadBoundCIDRList) | **Get** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleReadLocalSecretIDs**](AuthApi.md#AppRoleReadLocalSecretIDs) | **Get** /auth/{approle_mount_path}/role/{role_name}/local-secret-ids | 
[**AppRoleReadPeriod**](AuthApi.md#AppRoleReadPeriod) | **Get** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleReadPolicies**](AuthApi.md#AppRoleReadPolicies) | **Get** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleReadRole**](AuthApi.md#AppRoleReadRole) | **Get** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleReadRoleID**](AuthApi.md#AppRoleReadRoleID) | **Get** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleReadSecretIDBoundCIDRs**](AuthApi.md#AppRoleReadSecretIDBoundCIDRs) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleReadSecretIDNumUses**](AuthApi.md#AppRoleReadSecretIDNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleReadSecretIDTTL**](AuthApi.md#AppRoleReadSecretIDTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleReadTokenBoundCIDRs**](AuthApi.md#AppRoleReadTokenBoundCIDRs) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleReadTokenMaxTTL**](AuthApi.md#AppRoleReadTokenMaxTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleReadTokenNumUses**](AuthApi.md#AppRoleReadTokenNumUses) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleReadTokenTTL**](AuthApi.md#AppRoleReadTokenTTL) | **Get** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AppRoleTidySecretID**](AuthApi.md#AppRoleTidySecretID) | **Post** /auth/{approle_mount_path}/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
[**AppRoleWriteBindSecretID**](AuthApi.md#AppRoleWriteBindSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/bind-secret-id | 
[**AppRoleWriteBoundCIDRList**](AuthApi.md#AppRoleWriteBoundCIDRList) | **Post** /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list | 
[**AppRoleWriteCustomSecretID**](AuthApi.md#AppRoleWriteCustomSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/custom-secret-id | 
[**AppRoleWritePeriod**](AuthApi.md#AppRoleWritePeriod) | **Post** /auth/{approle_mount_path}/role/{role_name}/period | 
[**AppRoleWritePolicies**](AuthApi.md#AppRoleWritePolicies) | **Post** /auth/{approle_mount_path}/role/{role_name}/policies | 
[**AppRoleWriteRole**](AuthApi.md#AppRoleWriteRole) | **Post** /auth/{approle_mount_path}/role/{role_name} | 
[**AppRoleWriteRoleID**](AuthApi.md#AppRoleWriteRoleID) | **Post** /auth/{approle_mount_path}/role/{role_name}/role-id | 
[**AppRoleWriteSecretID**](AuthApi.md#AppRoleWriteSecretID) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id | 
[**AppRoleWriteSecretIDAccessorDestroy**](AuthApi.md#AppRoleWriteSecretIDAccessorDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy | 
[**AppRoleWriteSecretIDAccessorLookup**](AuthApi.md#AppRoleWriteSecretIDAccessorLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup | 
[**AppRoleWriteSecretIDBoundCIDRs**](AuthApi.md#AppRoleWriteSecretIDBoundCIDRs) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs | 
[**AppRoleWriteSecretIDDestroy**](AuthApi.md#AppRoleWriteSecretIDDestroy) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy | 
[**AppRoleWriteSecretIDLookup**](AuthApi.md#AppRoleWriteSecretIDLookup) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id/lookup | 
[**AppRoleWriteSecretIDNumUses**](AuthApi.md#AppRoleWriteSecretIDNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses | 
[**AppRoleWriteSecretIDTTL**](AuthApi.md#AppRoleWriteSecretIDTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl | 
[**AppRoleWriteTokenBoundCIDRs**](AuthApi.md#AppRoleWriteTokenBoundCIDRs) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs | 
[**AppRoleWriteTokenMaxTTL**](AuthApi.md#AppRoleWriteTokenMaxTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-max-ttl | 
[**AppRoleWriteTokenNumUses**](AuthApi.md#AppRoleWriteTokenNumUses) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-num-uses | 
[**AppRoleWriteTokenTTL**](AuthApi.md#AppRoleWriteTokenTTL) | **Post** /auth/{approle_mount_path}/role/{role_name}/token-ttl | 
[**AzureDeleteAuthConfig**](AuthApi.md#AzureDeleteAuthConfig) | **Delete** /auth/{azure_mount_path}/config | 
[**AzureDeleteAuthRole**](AuthApi.md#AzureDeleteAuthRole) | **Delete** /auth/{azure_mount_path}/role/{name} | 
[**AzureListAuthRoles**](AuthApi.md#AzureListAuthRoles) | **Get** /auth/{azure_mount_path}/role | 
[**AzureLogin**](AuthApi.md#AzureLogin) | **Post** /auth/{azure_mount_path}/login | 
[**AzureReadAuthConfig**](AuthApi.md#AzureReadAuthConfig) | **Get** /auth/{azure_mount_path}/config | 
[**AzureReadAuthRole**](AuthApi.md#AzureReadAuthRole) | **Get** /auth/{azure_mount_path}/role/{name} | 
[**AzureWriteAuthConfig**](AuthApi.md#AzureWriteAuthConfig) | **Post** /auth/{azure_mount_path}/config | 
[**AzureWriteAuthRole**](AuthApi.md#AzureWriteAuthRole) | **Post** /auth/{azure_mount_path}/role/{name} | 
[**CentrifyLogin**](AuthApi.md#CentrifyLogin) | **Post** /auth/{centrify_mount_path}/login | Log in with a username and password.
[**CentrifyReadConfig**](AuthApi.md#CentrifyReadConfig) | **Get** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**CentrifyWriteConfig**](AuthApi.md#CentrifyWriteConfig) | **Post** /auth/{centrify_mount_path}/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
[**CertificatesDelete**](AuthApi.md#CertificatesDelete) | **Delete** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesDeleteCRL**](AuthApi.md#CertificatesDeleteCRL) | **Delete** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesList**](AuthApi.md#CertificatesList) | **Get** /auth/{cert_mount_path}/certs | Manage trusted certificates used for authentication.
[**CertificatesListCRLs**](AuthApi.md#CertificatesListCRLs) | **Get** /auth/{cert_mount_path}/crls | 
[**CertificatesLogin**](AuthApi.md#CertificatesLogin) | **Post** /auth/{cert_mount_path}/login | 
[**CertificatesRead**](AuthApi.md#CertificatesRead) | **Get** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesReadCRL**](AuthApi.md#CertificatesReadCRL) | **Get** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesReadConfig**](AuthApi.md#CertificatesReadConfig) | **Get** /auth/{cert_mount_path}/config | 
[**CertificatesWrite**](AuthApi.md#CertificatesWrite) | **Post** /auth/{cert_mount_path}/certs/{name} | Manage trusted certificates used for authentication.
[**CertificatesWriteCRL**](AuthApi.md#CertificatesWriteCRL) | **Post** /auth/{cert_mount_path}/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
[**CertificatesWriteConfig**](AuthApi.md#CertificatesWriteConfig) | **Post** /auth/{cert_mount_path}/config | 
[**CloudFoundryDeleteConfig**](AuthApi.md#CloudFoundryDeleteConfig) | **Delete** /auth/{cf_mount_path}/config | 
[**CloudFoundryDeleteRole**](AuthApi.md#CloudFoundryDeleteRole) | **Delete** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryListRoles**](AuthApi.md#CloudFoundryListRoles) | **Get** /auth/{cf_mount_path}/roles | 
[**CloudFoundryLogin**](AuthApi.md#CloudFoundryLogin) | **Post** /auth/{cf_mount_path}/login | 
[**CloudFoundryReadConfig**](AuthApi.md#CloudFoundryReadConfig) | **Get** /auth/{cf_mount_path}/config | 
[**CloudFoundryReadRole**](AuthApi.md#CloudFoundryReadRole) | **Get** /auth/{cf_mount_path}/roles/{role} | 
[**CloudFoundryWriteConfig**](AuthApi.md#CloudFoundryWriteConfig) | **Post** /auth/{cf_mount_path}/config | 
[**CloudFoundryWriteRole**](AuthApi.md#CloudFoundryWriteRole) | **Post** /auth/{cf_mount_path}/roles/{role} | 
[**GitHubDeleteMapTeam**](AuthApi.md#GitHubDeleteMapTeam) | **Delete** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubDeleteMapUser**](AuthApi.md#GitHubDeleteMapUser) | **Delete** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GitHubLogin**](AuthApi.md#GitHubLogin) | **Post** /auth/{github_mount_path}/login | 
[**GitHubReadConfig**](AuthApi.md#GitHubReadConfig) | **Get** /auth/{github_mount_path}/config | 
[**GitHubReadMapTeam**](AuthApi.md#GitHubReadMapTeam) | **Get** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubReadMapTeams**](AuthApi.md#GitHubReadMapTeams) | **Get** /auth/{github_mount_path}/map/teams | Read mappings for teams
[**GitHubReadMapUser**](AuthApi.md#GitHubReadMapUser) | **Get** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GitHubReadMapUsers**](AuthApi.md#GitHubReadMapUsers) | **Get** /auth/{github_mount_path}/map/users | Read mappings for users
[**GitHubWriteConfig**](AuthApi.md#GitHubWriteConfig) | **Post** /auth/{github_mount_path}/config | 
[**GitHubWriteMapTeam**](AuthApi.md#GitHubWriteMapTeam) | **Post** /auth/{github_mount_path}/map/teams/{key} | Read/write/delete a single teams mapping
[**GitHubWriteMapUser**](AuthApi.md#GitHubWriteMapUser) | **Post** /auth/{github_mount_path}/map/users/{key} | Read/write/delete a single users mapping
[**GoogleCloudDeleteRole**](AuthApi.md#GoogleCloudDeleteRole) | **Delete** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudListRoles**](AuthApi.md#GoogleCloudListRoles) | **Get** /auth/{gcp_mount_path}/role | Lists all the roles that are registered with Vault.
[**GoogleCloudListRoles2**](AuthApi.md#GoogleCloudListRoles2) | **Get** /auth/{gcp_mount_path}/roles | Lists all the roles that are registered with Vault.
[**GoogleCloudLogin**](AuthApi.md#GoogleCloudLogin) | **Post** /auth/{gcp_mount_path}/login | 
[**GoogleCloudReadAuthConfig**](AuthApi.md#GoogleCloudReadAuthConfig) | **Get** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GoogleCloudReadRole**](AuthApi.md#GoogleCloudReadRole) | **Get** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudWriteAuthConfig**](AuthApi.md#GoogleCloudWriteAuthConfig) | **Post** /auth/{gcp_mount_path}/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
[**GoogleCloudWriteRole**](AuthApi.md#GoogleCloudWriteRole) | **Post** /auth/{gcp_mount_path}/role/{name} | Create a GCP role with associated policies and required attributes.
[**GoogleCloudWriteRoleLabels**](AuthApi.md#GoogleCloudWriteRoleLabels) | **Post** /auth/{gcp_mount_path}/role/{name}/labels | Add or remove labels for an existing &#x27;gce&#x27; role
[**GoogleCloudWriteRoleServiceAccounts**](AuthApi.md#GoogleCloudWriteRoleServiceAccounts) | **Post** /auth/{gcp_mount_path}/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
[**JWTDeleteRole**](AuthApi.md#JWTDeleteRole) | **Delete** /auth/{jwt_mount_path}/role/{name} | Delete an existing role.
[**JWTListRoles**](AuthApi.md#JWTListRoles) | **Get** /auth/{jwt_mount_path}/role | Lists all the roles registered with the backend.
[**JWTLogin**](AuthApi.md#JWTLogin) | **Post** /auth/{jwt_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**JWTReadConfig**](AuthApi.md#JWTReadConfig) | **Get** /auth/{jwt_mount_path}/config | Read the current JWT authentication backend configuration.
[**JWTReadOIDCCallback**](AuthApi.md#JWTReadOIDCCallback) | **Get** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**JWTReadRole**](AuthApi.md#JWTReadRole) | **Get** /auth/{jwt_mount_path}/role/{name} | Read an existing role.
[**JWTWriteConfig**](AuthApi.md#JWTWriteConfig) | **Post** /auth/{jwt_mount_path}/config | Configure the JWT authentication backend.
[**JWTWriteOIDCAuthURL**](AuthApi.md#JWTWriteOIDCAuthURL) | **Post** /auth/{jwt_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**JWTWriteOIDCCallback**](AuthApi.md#JWTWriteOIDCCallback) | **Post** /auth/{jwt_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**JWTWriteRole**](AuthApi.md#JWTWriteRole) | **Post** /auth/{jwt_mount_path}/role/{name} | Register an role with the backend.
[**KerberosDeleteGroup**](AuthApi.md#KerberosDeleteGroup) | **Delete** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosListGroups**](AuthApi.md#KerberosListGroups) | **Get** /auth/{kerberos_mount_path}/groups | 
[**KerberosLogin**](AuthApi.md#KerberosLogin) | **Post** /auth/{kerberos_mount_path}/login | 
[**KerberosReadConfig**](AuthApi.md#KerberosReadConfig) | **Get** /auth/{kerberos_mount_path}/config | 
[**KerberosReadGroup**](AuthApi.md#KerberosReadGroup) | **Get** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosReadLDAPConfig**](AuthApi.md#KerberosReadLDAPConfig) | **Get** /auth/{kerberos_mount_path}/config/ldap | 
[**KerberosWriteConfig**](AuthApi.md#KerberosWriteConfig) | **Post** /auth/{kerberos_mount_path}/config | 
[**KerberosWriteGroup**](AuthApi.md#KerberosWriteGroup) | **Post** /auth/{kerberos_mount_path}/groups/{name} | 
[**KerberosWriteLDAPConfig**](AuthApi.md#KerberosWriteLDAPConfig) | **Post** /auth/{kerberos_mount_path}/config/ldap | 
[**KubernetesDeleteAuthRole**](AuthApi.md#KubernetesDeleteAuthRole) | **Delete** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesListAuthRoles**](AuthApi.md#KubernetesListAuthRoles) | **Get** /auth/{kubernetes_mount_path}/role | Lists all the roles registered with the backend.
[**KubernetesLogin**](AuthApi.md#KubernetesLogin) | **Post** /auth/{kubernetes_mount_path}/login | Authenticates Kubernetes service accounts with Vault.
[**KubernetesReadAuthConfig**](AuthApi.md#KubernetesReadAuthConfig) | **Get** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**KubernetesReadAuthRole**](AuthApi.md#KubernetesReadAuthRole) | **Get** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**KubernetesWriteAuthConfig**](AuthApi.md#KubernetesWriteAuthConfig) | **Post** /auth/{kubernetes_mount_path}/config | Configures the JWT Public Key and Kubernetes API information.
[**KubernetesWriteAuthRole**](AuthApi.md#KubernetesWriteAuthRole) | **Post** /auth/{kubernetes_mount_path}/role/{name} | Register an role with the backend.
[**LDAPDeleteGroup**](AuthApi.md#LDAPDeleteGroup) | **Delete** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPDeleteUser**](AuthApi.md#LDAPDeleteUser) | **Delete** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LDAPListGroups**](AuthApi.md#LDAPListGroups) | **Get** /auth/{ldap_mount_path}/groups | Manage additional groups for users allowed to authenticate.
[**LDAPListUsers**](AuthApi.md#LDAPListUsers) | **Get** /auth/{ldap_mount_path}/users | Manage users allowed to authenticate.
[**LDAPLogin**](AuthApi.md#LDAPLogin) | **Post** /auth/{ldap_mount_path}/login/{username} | Log in with a username and password.
[**LDAPReadAuthConfig**](AuthApi.md#LDAPReadAuthConfig) | **Get** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**LDAPReadGroup**](AuthApi.md#LDAPReadGroup) | **Get** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPReadUser**](AuthApi.md#LDAPReadUser) | **Get** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**LDAPWriteAuthConfig**](AuthApi.md#LDAPWriteAuthConfig) | **Post** /auth/{ldap_mount_path}/config | Configure the LDAP server to connect to, along with its options.
[**LDAPWriteGroup**](AuthApi.md#LDAPWriteGroup) | **Post** /auth/{ldap_mount_path}/groups/{name} | Manage additional groups for users allowed to authenticate.
[**LDAPWriteUser**](AuthApi.md#LDAPWriteUser) | **Post** /auth/{ldap_mount_path}/users/{name} | Manage users allowed to authenticate.
[**OCIDeleteConfig**](AuthApi.md#OCIDeleteConfig) | **Delete** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIDeleteRole**](AuthApi.md#OCIDeleteRole) | **Delete** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OCIListRoles**](AuthApi.md#OCIListRoles) | **Get** /auth/{oci_mount_path}/role | Lists all the roles that are registered with Vault.
[**OCILoginWithRole**](AuthApi.md#OCILoginWithRole) | **Post** /auth/{oci_mount_path}/login/{role} | Authenticates to Vault using OCI credentials
[**OCIReadConfig**](AuthApi.md#OCIReadConfig) | **Get** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIReadRole**](AuthApi.md#OCIReadRole) | **Get** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OCIWriteConfig**](AuthApi.md#OCIWriteConfig) | **Post** /auth/{oci_mount_path}/config | Manages the configuration for the Vault Auth Plugin.
[**OCIWriteRole**](AuthApi.md#OCIWriteRole) | **Post** /auth/{oci_mount_path}/role/{role} | Create a role and associate policies to it.
[**OIDCDeleteAuthRole**](AuthApi.md#OIDCDeleteAuthRole) | **Delete** /auth/{oidc_mount_path}/role/{name} | Delete an existing role.
[**OIDCListAuthRoles**](AuthApi.md#OIDCListAuthRoles) | **Get** /auth/{oidc_mount_path}/role | Lists all the roles registered with the backend.
[**OIDCLogin**](AuthApi.md#OIDCLogin) | **Post** /auth/{oidc_mount_path}/login | Authenticates to Vault using a JWT (or OIDC) token.
[**OIDCReadAuthConfig**](AuthApi.md#OIDCReadAuthConfig) | **Get** /auth/{oidc_mount_path}/config | Read the current JWT authentication backend configuration.
[**OIDCReadAuthRole**](AuthApi.md#OIDCReadAuthRole) | **Get** /auth/{oidc_mount_path}/role/{name} | Read an existing role.
[**OIDCReadCallback**](AuthApi.md#OIDCReadCallback) | **Get** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to complete an OIDC login.
[**OIDCWriteAuthConfig**](AuthApi.md#OIDCWriteAuthConfig) | **Post** /auth/{oidc_mount_path}/config | Configure the JWT authentication backend.
[**OIDCWriteAuthRole**](AuthApi.md#OIDCWriteAuthRole) | **Post** /auth/{oidc_mount_path}/role/{name} | Register an role with the backend.
[**OIDCWriteAuthURL**](AuthApi.md#OIDCWriteAuthURL) | **Post** /auth/{oidc_mount_path}/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
[**OIDCWriteCallback**](AuthApi.md#OIDCWriteCallback) | **Post** /auth/{oidc_mount_path}/oidc/callback | Callback endpoint to handle form_posts.
[**OktaDeleteGroup**](AuthApi.md#OktaDeleteGroup) | **Delete** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaDeleteUser**](AuthApi.md#OktaDeleteUser) | **Delete** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaListGroups**](AuthApi.md#OktaListGroups) | **Get** /auth/{okta_mount_path}/groups | Manage users allowed to authenticate.
[**OktaListUsers**](AuthApi.md#OktaListUsers) | **Get** /auth/{okta_mount_path}/users | Manage additional groups for users allowed to authenticate.
[**OktaLogin**](AuthApi.md#OktaLogin) | **Post** /auth/{okta_mount_path}/login/{username} | Log in with a username and password.
[**OktaReadConfig**](AuthApi.md#OktaReadConfig) | **Get** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**OktaReadGroup**](AuthApi.md#OktaReadGroup) | **Get** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaReadUser**](AuthApi.md#OktaReadUser) | **Get** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**OktaVerify**](AuthApi.md#OktaVerify) | **Get** /auth/{okta_mount_path}/verify/{nonce} | 
[**OktaWriteConfig**](AuthApi.md#OktaWriteConfig) | **Post** /auth/{okta_mount_path}/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
[**OktaWriteGroup**](AuthApi.md#OktaWriteGroup) | **Post** /auth/{okta_mount_path}/groups/{name} | Manage users allowed to authenticate.
[**OktaWriteUser**](AuthApi.md#OktaWriteUser) | **Post** /auth/{okta_mount_path}/users/{name} | Manage additional groups for users allowed to authenticate.
[**RadiusDeleteUser**](AuthApi.md#RadiusDeleteUser) | **Delete** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusListUsers**](AuthApi.md#RadiusListUsers) | **Get** /auth/{radius_mount_path}/users | Manage users allowed to authenticate.
[**RadiusLogin**](AuthApi.md#RadiusLogin) | **Post** /auth/{radius_mount_path}/login | Log in with a username and password.
[**RadiusLoginWithUsername**](AuthApi.md#RadiusLoginWithUsername) | **Post** /auth/{radius_mount_path}/login/{urlusername} | Log in with a username and password.
[**RadiusReadConfig**](AuthApi.md#RadiusReadConfig) | **Get** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**RadiusReadUser**](AuthApi.md#RadiusReadUser) | **Get** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**RadiusWriteConfig**](AuthApi.md#RadiusWriteConfig) | **Post** /auth/{radius_mount_path}/config | Configure the RADIUS server to connect to, along with its options.
[**RadiusWriteUser**](AuthApi.md#RadiusWriteUser) | **Post** /auth/{radius_mount_path}/users/{name} | Manage users allowed to authenticate.
[**TokenDeleteRole**](AuthApi.md#TokenDeleteRole) | **Delete** /auth/{token_mount_path}/roles/{role_name} | 
[**TokenListAccessors**](AuthApi.md#TokenListAccessors) | **Get** /auth/{token_mount_path}/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#x27;sudo&#x27; capability in addition to &#x27;list&#x27;.
[**TokenListRoles**](AuthApi.md#TokenListRoles) | **Get** /auth/{token_mount_path}/roles | This endpoint lists configured roles.
[**TokenReadLookup**](AuthApi.md#TokenReadLookup) | **Get** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**TokenReadLookupSelf**](AuthApi.md#TokenReadLookupSelf) | **Get** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**TokenReadRole**](AuthApi.md#TokenReadRole) | **Get** /auth/{token_mount_path}/roles/{role_name} | 
[**TokenRenew**](AuthApi.md#TokenRenew) | **Post** /auth/{token_mount_path}/renew | This endpoint will renew the given token and prevent expiration.
[**TokenRenewAccessor**](AuthApi.md#TokenRenewAccessor) | **Post** /auth/{token_mount_path}/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenRenewSelf**](AuthApi.md#TokenRenewSelf) | **Post** /auth/{token_mount_path}/renew-self | This endpoint will renew the token used to call it and prevent expiration.
[**TokenRevoke**](AuthApi.md#TokenRevoke) | **Post** /auth/{token_mount_path}/revoke | This endpoint will delete the given token and all of its child tokens.
[**TokenRevokeAccessor**](AuthApi.md#TokenRevokeAccessor) | **Post** /auth/{token_mount_path}/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
[**TokenRevokeOrphan**](AuthApi.md#TokenRevokeOrphan) | **Post** /auth/{token_mount_path}/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
[**TokenRevokeSelf**](AuthApi.md#TokenRevokeSelf) | **Post** /auth/{token_mount_path}/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
[**TokenTidy**](AuthApi.md#TokenTidy) | **Post** /auth/{token_mount_path}/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**TokenWriteCreate**](AuthApi.md#TokenWriteCreate) | **Post** /auth/{token_mount_path}/create | The token create path is used to create new tokens.
[**TokenWriteCreateOrphan**](AuthApi.md#TokenWriteCreateOrphan) | **Post** /auth/{token_mount_path}/create-orphan | The token create path is used to create new orphan tokens.
[**TokenWriteCreateWithRole**](AuthApi.md#TokenWriteCreateWithRole) | **Post** /auth/{token_mount_path}/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
[**TokenWriteLookup**](AuthApi.md#TokenWriteLookup) | **Post** /auth/{token_mount_path}/lookup | This endpoint will lookup a token and its properties.
[**TokenWriteLookupAccessor**](AuthApi.md#TokenWriteLookupAccessor) | **Post** /auth/{token_mount_path}/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
[**TokenWriteLookupSelf**](AuthApi.md#TokenWriteLookupSelf) | **Post** /auth/{token_mount_path}/lookup-self | This endpoint will lookup a token and its properties.
[**TokenWriteRole**](AuthApi.md#TokenWriteRole) | **Post** /auth/{token_mount_path}/roles/{role_name} | 
[**UserpassDeleteUser**](AuthApi.md#UserpassDeleteUser) | **Delete** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassListUsers**](AuthApi.md#UserpassListUsers) | **Get** /auth/{userpass_mount_path}/users | Manage users allowed to authenticate.
[**UserpassLogin**](AuthApi.md#UserpassLogin) | **Post** /auth/{userpass_mount_path}/login/{username} | Log in with a username and password.
[**UserpassReadUser**](AuthApi.md#UserpassReadUser) | **Get** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassWriteUser**](AuthApi.md#UserpassWriteUser) | **Post** /auth/{userpass_mount_path}/users/{username} | Manage users allowed to authenticate.
[**UserpassWriteUserPassword**](AuthApi.md#UserpassWriteUserPassword) | **Post** /auth/{userpass_mount_path}/users/{username}/password | Reset user&#x27;s password.
[**UserpassWriteUserPolicies**](AuthApi.md#UserpassWriteUserPolicies) | **Post** /auth/{userpass_mount_path}/users/{username}/policies | Update the policies associated with the username.





## AWSConfigDeleteCertificate



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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

	resp, err := client.Auth.GitHubReadMapTeams(
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

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubReadMapUser

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

	resp, err := client.Auth.GitHubReadMapUsers(
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

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GitHubWriteConfig



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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



### Example

```go
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



