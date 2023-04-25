# System

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditingCalculateHash**](SystemApi.md#AuditingCalculateHash) | **Post** /sys/audit-hash/{path} | 
[**AuditingDisableDevice**](SystemApi.md#AuditingDisableDevice) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
[**AuditingDisableRequestHeader**](SystemApi.md#AuditingDisableRequestHeader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
[**AuditingEnableDevice**](SystemApi.md#AuditingEnableDevice) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
[**AuditingEnableRequestHeader**](SystemApi.md#AuditingEnableRequestHeader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
[**AuditingListEnabledDevices**](SystemApi.md#AuditingListEnabledDevices) | **Get** /sys/audit | List the enabled audit devices.
[**AuditingListRequestHeaders**](SystemApi.md#AuditingListRequestHeaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
[**AuditingReadRequestHeaderInformation**](SystemApi.md#AuditingReadRequestHeaderInformation) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
[**AuthDisableMethod**](SystemApi.md#AuthDisableMethod) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
[**AuthEnableMethod**](SystemApi.md#AuthEnableMethod) | **Post** /sys/auth/{path} | Enables a new auth method.
[**AuthListEnabledMethods**](SystemApi.md#AuthListEnabledMethods) | **Get** /sys/auth | 
[**AuthReadConfiguration**](SystemApi.md#AuthReadConfiguration) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
[**AuthReadTuningInformation**](SystemApi.md#AuthReadTuningInformation) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#x27;s configuration.
[**AuthTuneConfigurationParameters**](SystemApi.md#AuthTuneConfigurationParameters) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
[**CollectHostInformation**](SystemApi.md#CollectHostInformation) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
[**CollectInFlightRequestInformation**](SystemApi.md#CollectInFlightRequestInformation) | **Get** /sys/in-flight-req | reports in-flight requests
[**CorsConfigure**](SystemApi.md#CorsConfigure) | **Post** /sys/config/cors | Configure the CORS settings.
[**CorsDeleteConfiguration**](SystemApi.md#CorsDeleteConfiguration) | **Delete** /sys/config/cors | Remove any CORS settings.
[**CorsReadConfiguration**](SystemApi.md#CorsReadConfiguration) | **Get** /sys/config/cors | Return the current CORS settings.
[**EncryptionKeyConfigureRotation**](SystemApi.md#EncryptionKeyConfigureRotation) | **Post** /sys/rotate/config | 
[**EncryptionKeyReadRotationConfiguration**](SystemApi.md#EncryptionKeyReadRotationConfiguration) | **Get** /sys/rotate/config | 
[**EncryptionKeyRotate**](SystemApi.md#EncryptionKeyRotate) | **Post** /sys/rotate | 
[**EncryptionKeyStatus**](SystemApi.md#EncryptionKeyStatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
[**GenerateHash**](SystemApi.md#GenerateHash) | **Post** /sys/tools/hash | 
[**GenerateHashWithAlgorithm**](SystemApi.md#GenerateHashWithAlgorithm) | **Post** /sys/tools/hash/{urlalgorithm} | 
[**GenerateRandom**](SystemApi.md#GenerateRandom) | **Post** /sys/tools/random | 
[**GenerateRandomWithBytes**](SystemApi.md#GenerateRandomWithBytes) | **Post** /sys/tools/random/{urlbytes} | 
[**GenerateRandomWithSource**](SystemApi.md#GenerateRandomWithSource) | **Post** /sys/tools/random/{source} | 
[**GenerateRandomWithSourceAndBytes**](SystemApi.md#GenerateRandomWithSourceAndBytes) | **Post** /sys/tools/random/{source}/{urlbytes} | 
[**HaStatus**](SystemApi.md#HaStatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
[**Initialize**](SystemApi.md#Initialize) | **Post** /sys/init | Initialize a new Vault.
[**InternalClientActivityConfigure**](SystemApi.md#InternalClientActivityConfigure) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
[**InternalClientActivityExport**](SystemApi.md#InternalClientActivityExport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
[**InternalClientActivityReadConfiguration**](SystemApi.md#InternalClientActivityReadConfiguration) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
[**InternalClientActivityReportCounts**](SystemApi.md#InternalClientActivityReportCounts) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
[**InternalClientActivityReportCountsThisMonth**](SystemApi.md#InternalClientActivityReportCountsThisMonth) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
[**InternalCountEntities**](SystemApi.md#InternalCountEntities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
[**InternalCountRequests**](SystemApi.md#InternalCountRequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
[**InternalCountTokens**](SystemApi.md#InternalCountTokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
[**InternalGenerateOpenApiDocument**](SystemApi.md#InternalGenerateOpenApiDocument) | **Get** /sys/internal/specs/openapi | 
[**InternalGenerateOpenApiDocument2**](SystemApi.md#InternalGenerateOpenApiDocument2) | **Post** /sys/internal/specs/openapi | 
[**InternalInspectRouter**](SystemApi.md#InternalInspectRouter) | **Get** /sys/internal/inspect/router/{tag} | Expose the route entry and mount entry tables present in the router
[**InternalUiListEnabledFeatureFlags**](SystemApi.md#InternalUiListEnabledFeatureFlags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
[**InternalUiListEnabledVisibleMounts**](SystemApi.md#InternalUiListEnabledVisibleMounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
[**InternalUiListNamespaces**](SystemApi.md#InternalUiListNamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
[**InternalUiReadMountInformation**](SystemApi.md#InternalUiReadMountInformation) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
[**InternalUiReadResultantAcl**](SystemApi.md#InternalUiReadResultantAcl) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
[**LeaderStatus**](SystemApi.md#LeaderStatus) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
[**LeasesCount**](SystemApi.md#LeasesCount) | **Get** /sys/leases/count | 
[**LeasesForceRevokeLeaseWithPrefix**](SystemApi.md#LeasesForceRevokeLeaseWithPrefix) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**LeasesForceRevokeLeaseWithPrefix2**](SystemApi.md#LeasesForceRevokeLeaseWithPrefix2) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**LeasesList**](SystemApi.md#LeasesList) | **Get** /sys/leases | 
[**LeasesLookUp**](SystemApi.md#LeasesLookUp) | **Get** /sys/leases/lookup/ | 
[**LeasesLookUpWithPrefix**](SystemApi.md#LeasesLookUpWithPrefix) | **Get** /sys/leases/lookup/{prefix} | 
[**LeasesReadLease**](SystemApi.md#LeasesReadLease) | **Post** /sys/leases/lookup | 
[**LeasesRenewLease**](SystemApi.md#LeasesRenewLease) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
[**LeasesRenewLease2**](SystemApi.md#LeasesRenewLease2) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
[**LeasesRenewLeaseWithId**](SystemApi.md#LeasesRenewLeaseWithId) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**LeasesRenewLeaseWithId2**](SystemApi.md#LeasesRenewLeaseWithId2) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**LeasesRevokeLease**](SystemApi.md#LeasesRevokeLease) | **Post** /sys/leases/revoke | Revokes a lease immediately.
[**LeasesRevokeLease2**](SystemApi.md#LeasesRevokeLease2) | **Post** /sys/revoke | Revokes a lease immediately.
[**LeasesRevokeLeaseWithId**](SystemApi.md#LeasesRevokeLeaseWithId) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
[**LeasesRevokeLeaseWithId2**](SystemApi.md#LeasesRevokeLeaseWithId2) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
[**LeasesRevokeLeaseWithPrefix**](SystemApi.md#LeasesRevokeLeaseWithPrefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**LeasesRevokeLeaseWithPrefix2**](SystemApi.md#LeasesRevokeLeaseWithPrefix2) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**LeasesTidy**](SystemApi.md#LeasesTidy) | **Post** /sys/leases/tidy | 
[**ListExperimentalFeatures**](SystemApi.md#ListExperimentalFeatures) | **Get** /sys/experiments | Returns the available and enabled experiments
[**LockedUsersList**](SystemApi.md#LockedUsersList) | **Get** /sys/locked-users | Report the locked user count metrics, for this namespace and all child namespaces.
[**LockedUsersUnlock**](SystemApi.md#LockedUsersUnlock) | **Post** /sys/locked-users/{mount_accessor}/unlock/{alias_identifier} | Unlocks the user with given mount_accessor and alias_identifier
[**LoggersReadVerbosityLevel**](SystemApi.md#LoggersReadVerbosityLevel) | **Get** /sys/loggers | Read the log level for all existing loggers.
[**LoggersReadVerbosityLevelFor**](SystemApi.md#LoggersReadVerbosityLevelFor) | **Get** /sys/loggers/{name} | Read the log level for a single logger.
[**LoggersRevertVerbosityLevel**](SystemApi.md#LoggersRevertVerbosityLevel) | **Delete** /sys/loggers | Revert the all loggers to use log level provided in config.
[**LoggersRevertVerbosityLevelFor**](SystemApi.md#LoggersRevertVerbosityLevelFor) | **Delete** /sys/loggers/{name} | Revert a single logger to use log level provided in config.
[**LoggersUpdateVerbosityLevel**](SystemApi.md#LoggersUpdateVerbosityLevel) | **Post** /sys/loggers | Modify the log level for all existing loggers.
[**LoggersUpdateVerbosityLevelFor**](SystemApi.md#LoggersUpdateVerbosityLevelFor) | **Post** /sys/loggers/{name} | Modify the log level of a single logger.
[**Metrics**](SystemApi.md#Metrics) | **Get** /sys/metrics | 
[**MfaValidate**](SystemApi.md#MfaValidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
[**Monitor**](SystemApi.md#Monitor) | **Get** /sys/monitor | 
[**MountsDisableSecretsEngine**](SystemApi.md#MountsDisableSecretsEngine) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
[**MountsEnableSecretsEngine**](SystemApi.md#MountsEnableSecretsEngine) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
[**MountsListSecretsEngines**](SystemApi.md#MountsListSecretsEngines) | **Get** /sys/mounts | 
[**MountsReadConfiguration**](SystemApi.md#MountsReadConfiguration) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
[**MountsReadTuningInformation**](SystemApi.md#MountsReadTuningInformation) | **Get** /sys/mounts/{path}/tune | 
[**MountsTuneConfigurationParameters**](SystemApi.md#MountsTuneConfigurationParameters) | **Post** /sys/mounts/{path}/tune | 
[**PluginsCatalogListPlugins**](SystemApi.md#PluginsCatalogListPlugins) | **Get** /sys/plugins/catalog | 
[**PluginsCatalogListPluginsWithType**](SystemApi.md#PluginsCatalogListPluginsWithType) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
[**PluginsCatalogReadPluginConfiguration**](SystemApi.md#PluginsCatalogReadPluginConfiguration) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
[**PluginsCatalogReadPluginConfigurationWithType**](SystemApi.md#PluginsCatalogReadPluginConfigurationWithType) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
[**PluginsCatalogRegisterPlugin**](SystemApi.md#PluginsCatalogRegisterPlugin) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PluginsCatalogRegisterPluginWithType**](SystemApi.md#PluginsCatalogRegisterPluginWithType) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PluginsCatalogRemovePlugin**](SystemApi.md#PluginsCatalogRemovePlugin) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
[**PluginsCatalogRemovePluginWithType**](SystemApi.md#PluginsCatalogRemovePluginWithType) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
[**PluginsReloadBackends**](SystemApi.md#PluginsReloadBackends) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
[**PoliciesDeleteAclPolicy**](SystemApi.md#PoliciesDeleteAclPolicy) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
[**PoliciesDeleteAclPolicy2**](SystemApi.md#PoliciesDeleteAclPolicy2) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
[**PoliciesDeletePasswordPolicy**](SystemApi.md#PoliciesDeletePasswordPolicy) | **Delete** /sys/policies/password/{name} | Delete a password policy.
[**PoliciesGeneratePasswordFromPasswordPolicy**](SystemApi.md#PoliciesGeneratePasswordFromPasswordPolicy) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
[**PoliciesList**](SystemApi.md#PoliciesList) | **Get** /sys/policy | 
[**PoliciesListAclPolicies**](SystemApi.md#PoliciesListAclPolicies) | **Get** /sys/policies/acl | 
[**PoliciesListPasswordPolicies**](SystemApi.md#PoliciesListPasswordPolicies) | **Get** /sys/policies/password | List the existing password policies.
[**PoliciesReadAclPolicy**](SystemApi.md#PoliciesReadAclPolicy) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
[**PoliciesReadAclPolicy2**](SystemApi.md#PoliciesReadAclPolicy2) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
[**PoliciesReadPasswordPolicy**](SystemApi.md#PoliciesReadPasswordPolicy) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
[**PoliciesWriteAclPolicy**](SystemApi.md#PoliciesWriteAclPolicy) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
[**PoliciesWriteAclPolicy2**](SystemApi.md#PoliciesWriteAclPolicy2) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
[**PoliciesWritePasswordPolicy**](SystemApi.md#PoliciesWritePasswordPolicy) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
[**PprofBlocking**](SystemApi.md#PprofBlocking) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
[**PprofCommandLine**](SystemApi.md#PprofCommandLine) | **Get** /sys/pprof/cmdline | Returns the running program&#x27;s command line.
[**PprofCpuProfile**](SystemApi.md#PprofCpuProfile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
[**PprofExecutionTrace**](SystemApi.md#PprofExecutionTrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
[**PprofGoroutines**](SystemApi.md#PprofGoroutines) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
[**PprofIndex**](SystemApi.md#PprofIndex) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
[**PprofMemoryAllocations**](SystemApi.md#PprofMemoryAllocations) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
[**PprofMemoryAllocationsLive**](SystemApi.md#PprofMemoryAllocationsLive) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
[**PprofMutexes**](SystemApi.md#PprofMutexes) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
[**PprofSymbols**](SystemApi.md#PprofSymbols) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
[**PprofThreadCreations**](SystemApi.md#PprofThreadCreations) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
[**QueryTokenAccessorCapabilities**](SystemApi.md#QueryTokenAccessorCapabilities) | **Post** /sys/capabilities-accessor | 
[**QueryTokenCapabilities**](SystemApi.md#QueryTokenCapabilities) | **Post** /sys/capabilities | 
[**QueryTokenSelfCapabilities**](SystemApi.md#QueryTokenSelfCapabilities) | **Post** /sys/capabilities-self | 
[**RateLimitQuotasConfigure**](SystemApi.md#RateLimitQuotasConfigure) | **Post** /sys/quotas/config | 
[**RateLimitQuotasDelete**](SystemApi.md#RateLimitQuotasDelete) | **Delete** /sys/quotas/rate-limit/{name} | 
[**RateLimitQuotasList**](SystemApi.md#RateLimitQuotasList) | **Get** /sys/quotas/rate-limit | 
[**RateLimitQuotasRead**](SystemApi.md#RateLimitQuotasRead) | **Get** /sys/quotas/rate-limit/{name} | 
[**RateLimitQuotasReadConfiguration**](SystemApi.md#RateLimitQuotasReadConfiguration) | **Get** /sys/quotas/config | 
[**RateLimitQuotasWrite**](SystemApi.md#RateLimitQuotasWrite) | **Post** /sys/quotas/rate-limit/{name} | 
[**ReadHealthStatus**](SystemApi.md#ReadHealthStatus) | **Get** /sys/health | Returns the health status of Vault.
[**ReadInitializationStatus**](SystemApi.md#ReadInitializationStatus) | **Get** /sys/init | Returns the initialization status of Vault.
[**ReadSanitizedConfigurationState**](SystemApi.md#ReadSanitizedConfigurationState) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
[**ReadWrappingProperties**](SystemApi.md#ReadWrappingProperties) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
[**ReadWrappingProperties2**](SystemApi.md#ReadWrappingProperties2) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#x27;s token.
[**RekeyAttemptCancel**](SystemApi.md#RekeyAttemptCancel) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
[**RekeyAttemptInitialize**](SystemApi.md#RekeyAttemptInitialize) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
[**RekeyAttemptReadProgress**](SystemApi.md#RekeyAttemptReadProgress) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
[**RekeyAttemptUpdate**](SystemApi.md#RekeyAttemptUpdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
[**RekeyDeleteBackupKey**](SystemApi.md#RekeyDeleteBackupKey) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
[**RekeyDeleteBackupRecoveryKey**](SystemApi.md#RekeyDeleteBackupRecoveryKey) | **Delete** /sys/rekey/recovery-key-backup | 
[**RekeyReadBackupKey**](SystemApi.md#RekeyReadBackupKey) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
[**RekeyReadBackupRecoveryKey**](SystemApi.md#RekeyReadBackupRecoveryKey) | **Get** /sys/rekey/recovery-key-backup | 
[**RekeyVerificationCancel**](SystemApi.md#RekeyVerificationCancel) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
[**RekeyVerificationReadProgress**](SystemApi.md#RekeyVerificationReadProgress) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
[**RekeyVerificationUpdate**](SystemApi.md#RekeyVerificationUpdate) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
[**ReloadSubsystem**](SystemApi.md#ReloadSubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
[**Remount**](SystemApi.md#Remount) | **Post** /sys/remount | Initiate a mount migration
[**RemountStatus**](SystemApi.md#RemountStatus) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
[**ReplicationStatus**](SystemApi.md#ReplicationStatus) | **Get** /sys/replication/status | 
[**Rewrap**](SystemApi.md#Rewrap) | **Post** /sys/wrapping/rewrap | 
[**RootTokenGenerationCancel**](SystemApi.md#RootTokenGenerationCancel) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
[**RootTokenGenerationCancel2**](SystemApi.md#RootTokenGenerationCancel2) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
[**RootTokenGenerationInitialize**](SystemApi.md#RootTokenGenerationInitialize) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
[**RootTokenGenerationInitialize2**](SystemApi.md#RootTokenGenerationInitialize2) | **Post** /sys/generate-root | Initializes a new root generation attempt.
[**RootTokenGenerationReadProgress**](SystemApi.md#RootTokenGenerationReadProgress) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
[**RootTokenGenerationReadProgress2**](SystemApi.md#RootTokenGenerationReadProgress2) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
[**RootTokenGenerationUpdate**](SystemApi.md#RootTokenGenerationUpdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
[**Seal**](SystemApi.md#Seal) | **Post** /sys/seal | Seal the Vault.
[**SealStatus**](SystemApi.md#SealStatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
[**StepDownLeader**](SystemApi.md#StepDownLeader) | **Post** /sys/step-down | Cause the node to give up active status.
[**UiHeadersConfigure**](SystemApi.md#UiHeadersConfigure) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
[**UiHeadersDeleteConfiguration**](SystemApi.md#UiHeadersDeleteConfiguration) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
[**UiHeadersList**](SystemApi.md#UiHeadersList) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
[**UiHeadersReadConfiguration**](SystemApi.md#UiHeadersReadConfiguration) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#x27;s configuration
[**Unseal**](SystemApi.md#Unseal) | **Post** /sys/unseal | Unseal the Vault.
[**Unwrap**](SystemApi.md#Unwrap) | **Post** /sys/wrapping/unwrap | 
[**VersionHistory**](SystemApi.md#VersionHistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
[**Wrap**](SystemApi.md#Wrap) | **Post** /sys/wrapping/wrap | 





## AuditingCalculateHash



### Example

```go
package main

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

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	request := schema.NewAuditingCalculateHashRequestWithDefaults()
	resp, err := client.System.AuditingCalculateHash(
		context.Background(),
		path,
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
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditingCalculateHashRequest** | [**AuditingCalculateHashRequest**](AuditingCalculateHashRequest.md) |  | 

[**AuditingCalculateHashResponse**](AuditingCalculateHashResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingDisableDevice

Disable the audit device at the given path.

### Example

```go
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

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	resp, err := client.System.AuditingDisableDevice(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingDisableRequestHeader

Disable auditing of the given request header.

### Example

```go
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

	header := "header_example" // string | 
	resp, err := client.System.AuditingDisableRequestHeader(
		context.Background(),
		header,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**header** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingEnableDevice

Enable a new audit device at the supplied path.

### Example

```go
package main

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

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	request := schema.NewAuditingEnableDeviceRequestWithDefaults()
	resp, err := client.System.AuditingEnableDevice(
		context.Background(),
		path,
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
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditingEnableDeviceRequest** | [**AuditingEnableDeviceRequest**](AuditingEnableDeviceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingEnableRequestHeader

Enable auditing of a header.

### Example

```go
package main

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

	header := "header_example" // string | 
	request := schema.NewAuditingEnableRequestHeaderRequestWithDefaults()
	resp, err := client.System.AuditingEnableRequestHeader(
		context.Background(),
		header,
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
**header** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditingEnableRequestHeaderRequest** | [**AuditingEnableRequestHeaderRequest**](AuditingEnableRequestHeaderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingListEnabledDevices

List the enabled audit devices.

### Example

```go
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

	resp, err := client.System.AuditingListEnabledDevices(
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



## AuditingListRequestHeaders

List the request headers that are configured to be audited.

### Example

```go
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

	resp, err := client.System.AuditingListRequestHeaders(
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



[**AuditingListRequestHeadersResponse**](AuditingListRequestHeadersResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuditingReadRequestHeaderInformation

List the information for the given request header.

### Example

```go
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

	header := "header_example" // string | 
	resp, err := client.System.AuditingReadRequestHeaderInformation(
		context.Background(),
		header,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**header** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuthDisableMethod

Disable the auth method at the given auth path

### Example

```go
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

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	resp, err := client.System.AuthDisableMethod(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuthEnableMethod

Enables a new auth method.



### Example

```go
package main

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

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	request := schema.NewAuthEnableMethodRequestWithDefaults()
	resp, err := client.System.AuthEnableMethod(
		context.Background(),
		path,
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
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authEnableMethodRequest** | [**AuthEnableMethodRequest**](AuthEnableMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuthListEnabledMethods



### Example

```go
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

	resp, err := client.System.AuthListEnabledMethods(
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



## AuthReadConfiguration

Read the configuration of the auth engine at the given path.

### Example

```go
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

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	resp, err := client.System.AuthReadConfiguration(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**AuthReadConfigurationResponse**](AuthReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuthReadTuningInformation

Reads the given auth path's configuration.



### Example

```go
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

	path := "path_example" // string | Tune the configuration parameters for an auth path.
	resp, err := client.System.AuthReadTuningInformation(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**AuthReadTuningInformationResponse**](AuthReadTuningInformationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## AuthTuneConfigurationParameters

Tune configuration parameters for a given auth path.



### Example

```go
package main

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

	path := "path_example" // string | Tune the configuration parameters for an auth path.
	request := schema.NewAuthTuneConfigurationParametersRequestWithDefaults()
	resp, err := client.System.AuthTuneConfigurationParameters(
		context.Background(),
		path,
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
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authTuneConfigurationParametersRequest** | [**AuthTuneConfigurationParametersRequest**](AuthTuneConfigurationParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CollectHostInformation

Information about the host instance that this Vault server is running on.



### Example

```go
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

	resp, err := client.System.CollectHostInformation(
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



[**CollectHostInformationResponse**](CollectHostInformationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## CollectInFlightRequestInformation

reports in-flight requests



### Example

```go
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

	resp, err := client.System.CollectInFlightRequestInformation(
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



## CorsConfigure

Configure the CORS settings.

### Example

```go
package main

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

	request := schema.NewCorsConfigureRequestWithDefaults()
	resp, err := client.System.CorsConfigure(
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
 **corsConfigureRequest** | [**CorsConfigureRequest**](CorsConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## CorsDeleteConfiguration

Remove any CORS settings.

### Example

```go
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

	resp, err := client.System.CorsDeleteConfiguration(
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



## CorsReadConfiguration

Return the current CORS settings.

### Example

```go
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

	resp, err := client.System.CorsReadConfiguration(
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



[**CorsReadConfigurationResponse**](CorsReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## EncryptionKeyConfigureRotation



### Example

```go
package main

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

	request := schema.NewEncryptionKeyConfigureRotationRequestWithDefaults()
	resp, err := client.System.EncryptionKeyConfigureRotation(
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
 **encryptionKeyConfigureRotationRequest** | [**EncryptionKeyConfigureRotationRequest**](EncryptionKeyConfigureRotationRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EncryptionKeyReadRotationConfiguration



### Example

```go
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

	resp, err := client.System.EncryptionKeyReadRotationConfiguration(
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



[**EncryptionKeyReadRotationConfigurationResponse**](EncryptionKeyReadRotationConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## EncryptionKeyRotate



### Example

```go
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

	resp, err := client.System.EncryptionKeyRotate(
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



## EncryptionKeyStatus

Provides information about the backend encryption key.

### Example

```go
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

	resp, err := client.System.EncryptionKeyStatus(
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



## GenerateHash



### Example

```go
package main

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

	request := schema.NewGenerateHashRequestWithDefaults()
	resp, err := client.System.GenerateHash(
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
 **generateHashRequest** | [**GenerateHashRequest**](GenerateHashRequest.md) |  | 

[**GenerateHashResponse**](GenerateHashResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## GenerateHashWithAlgorithm



### Example

```go
package main

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

	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	request := schema.NewGenerateHashWithAlgorithmRequestWithDefaults()
	resp, err := client.System.GenerateHashWithAlgorithm(
		context.Background(),
		urlalgorithm,
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
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateHashWithAlgorithmRequest** | [**GenerateHashWithAlgorithmRequest**](GenerateHashWithAlgorithmRequest.md) |  | 

[**GenerateHashWithAlgorithmResponse**](GenerateHashWithAlgorithmResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## GenerateRandom



### Example

```go
package main

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

	request := schema.NewGenerateRandomRequestWithDefaults()
	resp, err := client.System.GenerateRandom(
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
 **generateRandomRequest** | [**GenerateRandomRequest**](GenerateRandomRequest.md) |  | 

[**GenerateRandomResponse**](GenerateRandomResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## GenerateRandomWithBytes



### Example

```go
package main

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

	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	request := schema.NewGenerateRandomWithBytesRequestWithDefaults()
	resp, err := client.System.GenerateRandomWithBytes(
		context.Background(),
		urlbytes,
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
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateRandomWithBytesRequest** | [**GenerateRandomWithBytesRequest**](GenerateRandomWithBytesRequest.md) |  | 

[**GenerateRandomWithBytesResponse**](GenerateRandomWithBytesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## GenerateRandomWithSource



### Example

```go
package main

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

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	request := schema.NewGenerateRandomWithSourceRequestWithDefaults()
	resp, err := client.System.GenerateRandomWithSource(
		context.Background(),
		source,
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
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateRandomWithSourceRequest** | [**GenerateRandomWithSourceRequest**](GenerateRandomWithSourceRequest.md) |  | 

[**GenerateRandomWithSourceResponse**](GenerateRandomWithSourceResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## GenerateRandomWithSourceAndBytes



### Example

```go
package main

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

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	request := schema.NewGenerateRandomWithSourceAndBytesRequestWithDefaults()
	resp, err := client.System.GenerateRandomWithSourceAndBytes(
		context.Background(),
		source,
		urlbytes,
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
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **generateRandomWithSourceAndBytesRequest** | [**GenerateRandomWithSourceAndBytesRequest**](GenerateRandomWithSourceAndBytesRequest.md) |  | 

[**GenerateRandomWithSourceAndBytesResponse**](GenerateRandomWithSourceAndBytesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## HaStatus

Check the HA status of a Vault cluster

### Example

```go
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

	resp, err := client.System.HaStatus(
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



[**HaStatusResponse**](HaStatusResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## Initialize

Initialize a new Vault.



### Example

```go
package main

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

	request := schema.NewInitializeRequestWithDefaults()
	resp, err := client.System.Initialize(
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
 **initializeRequest** | [**InitializeRequest**](InitializeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalClientActivityConfigure

Enable or disable collection of client count, set retention period, or set default reporting period.

### Example

```go
package main

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

	request := schema.NewInternalClientActivityConfigureRequestWithDefaults()
	resp, err := client.System.InternalClientActivityConfigure(
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
 **internalClientActivityConfigureRequest** | [**InternalClientActivityConfigureRequest**](InternalClientActivityConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalClientActivityExport

Report the client count metrics, for this namespace and all child namespaces.

### Example

```go
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

	resp, err := client.System.InternalClientActivityExport(
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



## InternalClientActivityReadConfiguration

Read the client count tracking configuration.

### Example

```go
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

	resp, err := client.System.InternalClientActivityReadConfiguration(
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



## InternalClientActivityReportCounts

Report the client count metrics, for this namespace and all child namespaces.

### Example

```go
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

	resp, err := client.System.InternalClientActivityReportCounts(
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



## InternalClientActivityReportCountsThisMonth

Report the number of clients for this month, for this namespace and all child namespaces.

### Example

```go
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

	resp, err := client.System.InternalClientActivityReportCountsThisMonth(
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



## InternalCountEntities

Backwards compatibility is not guaranteed for this API

### Example

```go
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

	resp, err := client.System.InternalCountEntities(
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



[**InternalCountEntitiesResponse**](InternalCountEntitiesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalCountRequests

Backwards compatibility is not guaranteed for this API

### Example

```go
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

	resp, err := client.System.InternalCountRequests(
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



## InternalCountTokens

Backwards compatibility is not guaranteed for this API

### Example

```go
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

	resp, err := client.System.InternalCountTokens(
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



[**InternalCountTokensResponse**](InternalCountTokensResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalGenerateOpenApiDocument



### Example

```go
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

	genericMountPaths := true // bool | Use generic mount paths (defaults to false)
	resp, err := client.System.InternalGenerateOpenApiDocument(
		context.Background(),
		genericMountPaths,
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
 **genericMountPaths** | **bool** | Use generic mount paths | [default to false]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalGenerateOpenApiDocument2



### Example

```go
package main

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

	request := schema.NewInternalGenerateOpenApiDocument2RequestWithDefaults()
	genericMountPaths := true // bool | Use generic mount paths (defaults to false)
	resp, err := client.System.InternalGenerateOpenApiDocument2(
		context.Background(),
		request,
		genericMountPaths,
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
 **internalGenerateOpenApiDocument2Request** | [**InternalGenerateOpenApiDocument2Request**](InternalGenerateOpenApiDocument2Request.md) |  | 
 **genericMountPaths** | **bool** | Use generic mount paths | [default to false]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalInspectRouter

Expose the route entry and mount entry tables present in the router

### Example

```go
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

	tag := "tag_example" // string | Name of subtree being observed
	resp, err := client.System.InternalInspectRouter(
		context.Background(),
		tag,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**tag** | **string** | Name of subtree being observed | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalUiListEnabledFeatureFlags

Lists enabled feature flags.

### Example

```go
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

	resp, err := client.System.InternalUiListEnabledFeatureFlags(
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



[**InternalUiListEnabledFeatureFlagsResponse**](InternalUiListEnabledFeatureFlagsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalUiListEnabledVisibleMounts

Lists all enabled and visible auth and secrets mounts.

### Example

```go
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

	resp, err := client.System.InternalUiListEnabledVisibleMounts(
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



[**InternalUiListEnabledVisibleMountsResponse**](InternalUiListEnabledVisibleMountsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalUiListNamespaces

Backwards compatibility is not guaranteed for this API

### Example

```go
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

	resp, err := client.System.InternalUiListNamespaces(
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



[**InternalUiListNamespacesResponse**](InternalUiListNamespacesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalUiReadMountInformation

Return information about the given mount.

### Example

```go
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

	path := "path_example" // string | The path of the mount.
	resp, err := client.System.InternalUiReadMountInformation(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path of the mount. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**InternalUiReadMountInformationResponse**](InternalUiReadMountInformationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## InternalUiReadResultantAcl

Backwards compatibility is not guaranteed for this API

### Example

```go
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

	resp, err := client.System.InternalUiReadResultantAcl(
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



[**InternalUiReadResultantAclResponse**](InternalUiReadResultantAclResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeaderStatus

Returns the high availability status and current leader instance of Vault.

### Example

```go
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

	resp, err := client.System.LeaderStatus(
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



[**LeaderStatusResponse**](LeaderStatusResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesCount



### Example

```go
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

	resp, err := client.System.LeasesCount(
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



[**LeasesCountResponse**](LeasesCountResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesForceRevokeLeaseWithPrefix

Revokes all secrets or tokens generated under a given prefix immediately



### Example

```go
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

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	resp, err := client.System.LeasesForceRevokeLeaseWithPrefix(
		context.Background(),
		prefix,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesForceRevokeLeaseWithPrefix2

Revokes all secrets or tokens generated under a given prefix immediately



### Example

```go
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

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	resp, err := client.System.LeasesForceRevokeLeaseWithPrefix2(
		context.Background(),
		prefix,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesList



### Example

```go
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

	resp, err := client.System.LeasesList(
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



[**LeasesListResponse**](LeasesListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesLookUp



### Example

```go
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

	resp, err := client.System.LeasesLookUp(
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

[**LeasesLookUpResponse**](LeasesLookUpResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesLookUpWithPrefix



### Example

```go
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

	prefix := "prefix_example" // string | The path to list leases under. Example: \"aws/creds/deploy\"
	resp, err := client.System.LeasesLookUpWithPrefix(
		context.Background(),
		prefix,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**prefix** | **string** | The path to list leases under. Example: \&quot;aws/creds/deploy\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**LeasesLookUpWithPrefixResponse**](LeasesLookUpWithPrefixResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesReadLease



### Example

```go
package main

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

	request := schema.NewLeasesReadLeaseRequestWithDefaults()
	resp, err := client.System.LeasesReadLease(
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
 **leasesReadLeaseRequest** | [**LeasesReadLeaseRequest**](LeasesReadLeaseRequest.md) |  | 

[**LeasesReadLeaseResponse**](LeasesReadLeaseResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRenewLease

Renews a lease, requesting to extend the lease.

### Example

```go
package main

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

	request := schema.NewLeasesRenewLeaseRequestWithDefaults()
	resp, err := client.System.LeasesRenewLease(
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
 **leasesRenewLeaseRequest** | [**LeasesRenewLeaseRequest**](LeasesRenewLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRenewLease2

Renews a lease, requesting to extend the lease.

### Example

```go
package main

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

	request := schema.NewLeasesRenewLease2RequestWithDefaults()
	resp, err := client.System.LeasesRenewLease2(
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
 **leasesRenewLease2Request** | [**LeasesRenewLease2Request**](LeasesRenewLease2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRenewLeaseWithId

Renews a lease, requesting to extend the lease.

### Example

```go
package main

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

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewLeasesRenewLeaseWithIdRequestWithDefaults()
	resp, err := client.System.LeasesRenewLeaseWithId(
		context.Background(),
		urlLeaseId,
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
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRenewLeaseWithIdRequest** | [**LeasesRenewLeaseWithIdRequest**](LeasesRenewLeaseWithIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRenewLeaseWithId2

Renews a lease, requesting to extend the lease.

### Example

```go
package main

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

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewLeasesRenewLeaseWithId2RequestWithDefaults()
	resp, err := client.System.LeasesRenewLeaseWithId2(
		context.Background(),
		urlLeaseId,
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
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRenewLeaseWithId2Request** | [**LeasesRenewLeaseWithId2Request**](LeasesRenewLeaseWithId2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLease

Revokes a lease immediately.

### Example

```go
package main

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

	request := schema.NewLeasesRevokeLeaseRequestWithDefaults()
	resp, err := client.System.LeasesRevokeLease(
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
 **leasesRevokeLeaseRequest** | [**LeasesRevokeLeaseRequest**](LeasesRevokeLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLease2

Revokes a lease immediately.

### Example

```go
package main

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

	request := schema.NewLeasesRevokeLease2RequestWithDefaults()
	resp, err := client.System.LeasesRevokeLease2(
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
 **leasesRevokeLease2Request** | [**LeasesRevokeLease2Request**](LeasesRevokeLease2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLeaseWithId

Revokes a lease immediately.

### Example

```go
package main

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

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewLeasesRevokeLeaseWithIdRequestWithDefaults()
	resp, err := client.System.LeasesRevokeLeaseWithId(
		context.Background(),
		urlLeaseId,
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
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRevokeLeaseWithIdRequest** | [**LeasesRevokeLeaseWithIdRequest**](LeasesRevokeLeaseWithIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLeaseWithId2

Revokes a lease immediately.

### Example

```go
package main

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

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewLeasesRevokeLeaseWithId2RequestWithDefaults()
	resp, err := client.System.LeasesRevokeLeaseWithId2(
		context.Background(),
		urlLeaseId,
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
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRevokeLeaseWithId2Request** | [**LeasesRevokeLeaseWithId2Request**](LeasesRevokeLeaseWithId2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLeaseWithPrefix

Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.

### Example

```go
package main

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

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	request := schema.NewLeasesRevokeLeaseWithPrefixRequestWithDefaults()
	resp, err := client.System.LeasesRevokeLeaseWithPrefix(
		context.Background(),
		prefix,
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
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRevokeLeaseWithPrefixRequest** | [**LeasesRevokeLeaseWithPrefixRequest**](LeasesRevokeLeaseWithPrefixRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesRevokeLeaseWithPrefix2

Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.

### Example

```go
package main

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

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	request := schema.NewLeasesRevokeLeaseWithPrefix2RequestWithDefaults()
	resp, err := client.System.LeasesRevokeLeaseWithPrefix2(
		context.Background(),
		prefix,
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
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leasesRevokeLeaseWithPrefix2Request** | [**LeasesRevokeLeaseWithPrefix2Request**](LeasesRevokeLeaseWithPrefix2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LeasesTidy



### Example

```go
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

	resp, err := client.System.LeasesTidy(
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



## ListExperimentalFeatures

Returns the available and enabled experiments

### Example

```go
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

	resp, err := client.System.ListExperimentalFeatures(
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



## LockedUsersList

Report the locked user count metrics, for this namespace and all child namespaces.

### Example

```go
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

	resp, err := client.System.LockedUsersList(
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



## LockedUsersUnlock

Unlocks the user with given mount_accessor and alias_identifier

### Example

```go
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

	aliasIdentifier := "aliasIdentifier_example" // string | It is the name of the alias (user). For example, if the alias belongs to userpass backend, the name should be a valid username within userpass auth method. If the alias belongs to an approle auth method, the name should be a valid RoleID
	mountAccessor := "mountAccessor_example" // string | MountAccessor is the identifier of the mount entry to which the user belongs
	resp, err := client.System.LockedUsersUnlock(
		context.Background(),
		aliasIdentifier,
		mountAccessor,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**aliasIdentifier** | **string** | It is the name of the alias (user). For example, if the alias belongs to userpass backend, the name should be a valid username within userpass auth method. If the alias belongs to an approle auth method, the name should be a valid RoleID | 
**mountAccessor** | **string** | MountAccessor is the identifier of the mount entry to which the user belongs | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LoggersReadVerbosityLevel

Read the log level for all existing loggers.

### Example

```go
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

	resp, err := client.System.LoggersReadVerbosityLevel(
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



## LoggersReadVerbosityLevelFor

Read the log level for a single logger.

### Example

```go
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

	name := "name_example" // string | The name of the logger to be modified.
	resp, err := client.System.LoggersReadVerbosityLevelFor(
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
**name** | **string** | The name of the logger to be modified. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LoggersRevertVerbosityLevel

Revert the all loggers to use log level provided in config.

### Example

```go
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

	resp, err := client.System.LoggersRevertVerbosityLevel(
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



## LoggersRevertVerbosityLevelFor

Revert a single logger to use log level provided in config.

### Example

```go
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

	name := "name_example" // string | The name of the logger to be modified.
	resp, err := client.System.LoggersRevertVerbosityLevelFor(
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
**name** | **string** | The name of the logger to be modified. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LoggersUpdateVerbosityLevel

Modify the log level for all existing loggers.

### Example

```go
package main

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

	request := schema.NewLoggersUpdateVerbosityLevelRequestWithDefaults()
	resp, err := client.System.LoggersUpdateVerbosityLevel(
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
 **loggersUpdateVerbosityLevelRequest** | [**LoggersUpdateVerbosityLevelRequest**](LoggersUpdateVerbosityLevelRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## LoggersUpdateVerbosityLevelFor

Modify the log level of a single logger.

### Example

```go
package main

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

	name := "name_example" // string | The name of the logger to be modified.
	request := schema.NewLoggersUpdateVerbosityLevelForRequestWithDefaults()
	resp, err := client.System.LoggersUpdateVerbosityLevelFor(
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
**name** | **string** | The name of the logger to be modified. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loggersUpdateVerbosityLevelForRequest** | [**LoggersUpdateVerbosityLevelForRequest**](LoggersUpdateVerbosityLevelForRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Metrics



### Example

```go
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

	format := "format_example" // string | Format to export metrics into. Currently accepts only \"prometheus\".
	resp, err := client.System.Metrics(
		context.Background(),
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
 **format** | **string** | Format to export metrics into. Currently accepts only \&quot;prometheus\&quot;. | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MfaValidate

Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token

### Example

```go
package main

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

	request := schema.NewMfaValidateRequestWithDefaults()
	resp, err := client.System.MfaValidate(
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
 **mfaValidateRequest** | [**MfaValidateRequest**](MfaValidateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Monitor



### Example

```go
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

	logFormat := "logFormat_example" // string | Output format of logs. Supported values are \"standard\" and \"json\". The default is \"standard\". (defaults to "standard")
	logLevel := "logLevel_example" // string | Log level to view system logs at. Currently supported values are \"trace\", \"debug\", \"info\", \"warn\", \"error\".
	resp, err := client.System.Monitor(
		context.Background(),
		logFormat,
		logLevel,
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
 **logFormat** | **string** | Output format of logs. Supported values are \&quot;standard\&quot; and \&quot;json\&quot;. The default is \&quot;standard\&quot;. | [default to &quot;standard&quot;]
 **logLevel** | **string** | Log level to view system logs at. Currently supported values are \&quot;trace\&quot;, \&quot;debug\&quot;, \&quot;info\&quot;, \&quot;warn\&quot;, \&quot;error\&quot;. | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MountsDisableSecretsEngine

Disable the mount point specified at the given path.

### Example

```go
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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.MountsDisableSecretsEngine(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MountsEnableSecretsEngine

Enable a new secrets engine at the given path.

### Example

```go
package main

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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	request := schema.NewMountsEnableSecretsEngineRequestWithDefaults()
	resp, err := client.System.MountsEnableSecretsEngine(
		context.Background(),
		path,
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
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mountsEnableSecretsEngineRequest** | [**MountsEnableSecretsEngineRequest**](MountsEnableSecretsEngineRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MountsListSecretsEngines



### Example

```go
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

	resp, err := client.System.MountsListSecretsEngines(
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



## MountsReadConfiguration

Read the configuration of the secret engine at the given path.

### Example

```go
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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.MountsReadConfiguration(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**MountsReadConfigurationResponse**](MountsReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## MountsReadTuningInformation



### Example

```go
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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.MountsReadTuningInformation(
		context.Background(),
		path,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**MountsReadTuningInformationResponse**](MountsReadTuningInformationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## MountsTuneConfigurationParameters



### Example

```go
package main

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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	request := schema.NewMountsTuneConfigurationParametersRequestWithDefaults()
	resp, err := client.System.MountsTuneConfigurationParameters(
		context.Background(),
		path,
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
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mountsTuneConfigurationParametersRequest** | [**MountsTuneConfigurationParametersRequest**](MountsTuneConfigurationParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogListPlugins



### Example

```go
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

	resp, err := client.System.PluginsCatalogListPlugins(
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



[**PluginsCatalogListPluginsResponse**](PluginsCatalogListPluginsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogListPluginsWithType

List the plugins in the catalog.

### Example

```go
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

	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.PluginsCatalogListPluginsWithType(
		context.Background(),
		type_,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**PluginsCatalogListPluginsWithTypeResponse**](PluginsCatalogListPluginsWithTypeResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogReadPluginConfiguration

Return the configuration data for the plugin with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the plugin
	resp, err := client.System.PluginsCatalogReadPluginConfiguration(
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
**name** | **string** | The name of the plugin | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PluginsCatalogReadPluginConfigurationResponse**](PluginsCatalogReadPluginConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogReadPluginConfigurationWithType

Return the configuration data for the plugin with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.PluginsCatalogReadPluginConfigurationWithType(
		context.Background(),
		name,
		type_,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



[**PluginsCatalogReadPluginConfigurationWithTypeResponse**](PluginsCatalogReadPluginConfigurationWithTypeResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogRegisterPlugin

Register a new plugin, or updates an existing one with the supplied name.

### Example

```go
package main

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

	name := "name_example" // string | The name of the plugin
	request := schema.NewPluginsCatalogRegisterPluginRequestWithDefaults()
	resp, err := client.System.PluginsCatalogRegisterPlugin(
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
**name** | **string** | The name of the plugin | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pluginsCatalogRegisterPluginRequest** | [**PluginsCatalogRegisterPluginRequest**](PluginsCatalogRegisterPluginRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogRegisterPluginWithType

Register a new plugin, or updates an existing one with the supplied name.

### Example

```go
package main

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

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	request := schema.NewPluginsCatalogRegisterPluginWithTypeRequestWithDefaults()
	resp, err := client.System.PluginsCatalogRegisterPluginWithType(
		context.Background(),
		name,
		type_,
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
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pluginsCatalogRegisterPluginWithTypeRequest** | [**PluginsCatalogRegisterPluginWithTypeRequest**](PluginsCatalogRegisterPluginWithTypeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogRemovePlugin

Remove the plugin with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the plugin
	resp, err := client.System.PluginsCatalogRemovePlugin(
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
**name** | **string** | The name of the plugin | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsCatalogRemovePluginWithType

Remove the plugin with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.PluginsCatalogRemovePluginWithType(
		context.Background(),
		name,
		type_,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PluginsReloadBackends

Reload mounted plugin backends.



### Example

```go
package main

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

	request := schema.NewPluginsReloadBackendsRequestWithDefaults()
	resp, err := client.System.PluginsReloadBackends(
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
 **pluginsReloadBackendsRequest** | [**PluginsReloadBackendsRequest**](PluginsReloadBackendsRequest.md) |  | 

[**PluginsReloadBackendsResponse**](PluginsReloadBackendsResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesDeleteAclPolicy

Delete the ACL policy with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.PoliciesDeleteAclPolicy(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesDeleteAclPolicy2

Delete the policy with the given name.

### Example

```go
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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.PoliciesDeleteAclPolicy2(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesDeletePasswordPolicy

Delete a password policy.

### Example

```go
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

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.PoliciesDeletePasswordPolicy(
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
**name** | **string** | The name of the password policy. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesGeneratePasswordFromPasswordPolicy

Generate a password from an existing password policy.

### Example

```go
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

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.PoliciesGeneratePasswordFromPasswordPolicy(
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
**name** | **string** | The name of the password policy. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PoliciesGeneratePasswordFromPasswordPolicyResponse**](PoliciesGeneratePasswordFromPasswordPolicyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesList



### Example

```go
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

	resp, err := client.System.PoliciesList(
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
 **list** | **string** | Return a list if &#x60;true&#x60; | 

[**PoliciesListResponse**](PoliciesListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesListAclPolicies



### Example

```go
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

	resp, err := client.System.PoliciesListAclPolicies(
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

[**PoliciesListAclPoliciesResponse**](PoliciesListAclPoliciesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesListPasswordPolicies

List the existing password policies.

### Example

```go
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

	resp, err := client.System.PoliciesListPasswordPolicies(
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

[**PoliciesListPasswordPoliciesResponse**](PoliciesListPasswordPoliciesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesReadAclPolicy

Retrieve information about the named ACL policy.

### Example

```go
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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.PoliciesReadAclPolicy(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PoliciesReadAclPolicyResponse**](PoliciesReadAclPolicyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesReadAclPolicy2

Retrieve the policy body for the named policy.

### Example

```go
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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.PoliciesReadAclPolicy2(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PoliciesReadAclPolicy2Response**](PoliciesReadAclPolicy2Response.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesReadPasswordPolicy

Retrieve an existing password policy.

### Example

```go
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

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.PoliciesReadPasswordPolicy(
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
**name** | **string** | The name of the password policy. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**PoliciesReadPasswordPolicyResponse**](PoliciesReadPasswordPolicyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesWriteAclPolicy

Add a new or update an existing ACL policy.

### Example

```go
package main

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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	request := schema.NewPoliciesWriteAclPolicyRequestWithDefaults()
	resp, err := client.System.PoliciesWriteAclPolicy(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policiesWriteAclPolicyRequest** | [**PoliciesWriteAclPolicyRequest**](PoliciesWriteAclPolicyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesWriteAclPolicy2

Add a new or update an existing policy.

### Example

```go
package main

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

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	request := schema.NewPoliciesWriteAclPolicy2RequestWithDefaults()
	resp, err := client.System.PoliciesWriteAclPolicy2(
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
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policiesWriteAclPolicy2Request** | [**PoliciesWriteAclPolicy2Request**](PoliciesWriteAclPolicy2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PoliciesWritePasswordPolicy

Add a new or update an existing password policy.

### Example

```go
package main

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

	name := "name_example" // string | The name of the password policy.
	request := schema.NewPoliciesWritePasswordPolicyRequestWithDefaults()
	resp, err := client.System.PoliciesWritePasswordPolicy(
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
**name** | **string** | The name of the password policy. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policiesWritePasswordPolicyRequest** | [**PoliciesWritePasswordPolicyRequest**](PoliciesWritePasswordPolicyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PprofBlocking

Returns stack traces that led to blocking on synchronization primitives



### Example

```go
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

	resp, err := client.System.PprofBlocking(
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



## PprofCommandLine

Returns the running program's command line.



### Example

```go
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

	resp, err := client.System.PprofCommandLine(
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



## PprofCpuProfile

Returns a pprof-formatted cpu profile payload.



### Example

```go
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

	resp, err := client.System.PprofCpuProfile(
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



## PprofExecutionTrace

Returns the execution trace in binary form.



### Example

```go
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

	resp, err := client.System.PprofExecutionTrace(
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



## PprofGoroutines

Returns stack traces of all current goroutines.



### Example

```go
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

	resp, err := client.System.PprofGoroutines(
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



## PprofIndex

Returns an HTML page listing the available profiles.



### Example

```go
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

	resp, err := client.System.PprofIndex(
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



## PprofMemoryAllocations

Returns a sampling of all past memory allocations.



### Example

```go
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

	resp, err := client.System.PprofMemoryAllocations(
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



## PprofMemoryAllocationsLive

Returns a sampling of memory allocations of live object.



### Example

```go
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

	resp, err := client.System.PprofMemoryAllocationsLive(
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



## PprofMutexes

Returns stack traces of holders of contended mutexes



### Example

```go
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

	resp, err := client.System.PprofMutexes(
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



## PprofSymbols

Returns the program counters listed in the request.



### Example

```go
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

	resp, err := client.System.PprofSymbols(
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



## PprofThreadCreations

Returns stack traces that led to the creation of new OS threads



### Example

```go
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

	resp, err := client.System.PprofThreadCreations(
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



## QueryTokenAccessorCapabilities



### Example

```go
package main

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

	request := schema.NewQueryTokenAccessorCapabilitiesRequestWithDefaults()
	resp, err := client.System.QueryTokenAccessorCapabilities(
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
 **queryTokenAccessorCapabilitiesRequest** | [**QueryTokenAccessorCapabilitiesRequest**](QueryTokenAccessorCapabilitiesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## QueryTokenCapabilities



### Example

```go
package main

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

	request := schema.NewQueryTokenCapabilitiesRequestWithDefaults()
	resp, err := client.System.QueryTokenCapabilities(
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
 **queryTokenCapabilitiesRequest** | [**QueryTokenCapabilitiesRequest**](QueryTokenCapabilitiesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## QueryTokenSelfCapabilities



### Example

```go
package main

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

	request := schema.NewQueryTokenSelfCapabilitiesRequestWithDefaults()
	resp, err := client.System.QueryTokenSelfCapabilities(
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
 **queryTokenSelfCapabilitiesRequest** | [**QueryTokenSelfCapabilitiesRequest**](QueryTokenSelfCapabilitiesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasConfigure



### Example

```go
package main

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

	request := schema.NewRateLimitQuotasConfigureRequestWithDefaults()
	resp, err := client.System.RateLimitQuotasConfigure(
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
 **rateLimitQuotasConfigureRequest** | [**RateLimitQuotasConfigureRequest**](RateLimitQuotasConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasDelete



### Example

```go
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

	name := "name_example" // string | Name of the quota rule.
	resp, err := client.System.RateLimitQuotasDelete(
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
**name** | **string** | Name of the quota rule. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasList



### Example

```go
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

	resp, err := client.System.RateLimitQuotasList(
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

[**RateLimitQuotasListResponse**](RateLimitQuotasListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasRead



### Example

```go
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

	name := "name_example" // string | Name of the quota rule.
	resp, err := client.System.RateLimitQuotasRead(
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
**name** | **string** | Name of the quota rule. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**RateLimitQuotasReadResponse**](RateLimitQuotasReadResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasReadConfiguration



### Example

```go
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

	resp, err := client.System.RateLimitQuotasReadConfiguration(
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



[**RateLimitQuotasReadConfigurationResponse**](RateLimitQuotasReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RateLimitQuotasWrite



### Example

```go
package main

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

	name := "name_example" // string | Name of the quota rule.
	request := schema.NewRateLimitQuotasWriteRequestWithDefaults()
	resp, err := client.System.RateLimitQuotasWrite(
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
**name** | **string** | Name of the quota rule. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateLimitQuotasWriteRequest** | [**RateLimitQuotasWriteRequest**](RateLimitQuotasWriteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadHealthStatus

Returns the health status of Vault.

### Example

```go
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

	resp, err := client.System.ReadHealthStatus(
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



## ReadInitializationStatus

Returns the initialization status of Vault.

### Example

```go
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

	resp, err := client.System.ReadInitializationStatus(
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



## ReadSanitizedConfigurationState

Return a sanitized version of the Vault server configuration.



### Example

```go
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

	resp, err := client.System.ReadSanitizedConfigurationState(
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



## ReadWrappingProperties

Look up wrapping properties for the given token.

### Example

```go
package main

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

	request := schema.NewReadWrappingPropertiesRequestWithDefaults()
	resp, err := client.System.ReadWrappingProperties(
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
 **readWrappingPropertiesRequest** | [**ReadWrappingPropertiesRequest**](ReadWrappingPropertiesRequest.md) |  | 

[**ReadWrappingPropertiesResponse**](ReadWrappingPropertiesResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadWrappingProperties2

Look up wrapping properties for the requester's token.

### Example

```go
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

	resp, err := client.System.ReadWrappingProperties2(
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



[**ReadWrappingProperties2Response**](ReadWrappingProperties2Response.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyAttemptCancel

Cancels any in-progress rekey.



### Example

```go
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

	resp, err := client.System.RekeyAttemptCancel(
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



## RekeyAttemptInitialize

Initializes a new rekey attempt.



### Example

```go
package main

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

	request := schema.NewRekeyAttemptInitializeRequestWithDefaults()
	resp, err := client.System.RekeyAttemptInitialize(
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
 **rekeyAttemptInitializeRequest** | [**RekeyAttemptInitializeRequest**](RekeyAttemptInitializeRequest.md) |  | 

[**RekeyAttemptInitializeResponse**](RekeyAttemptInitializeResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyAttemptReadProgress

Reads the configuration and progress of the current rekey attempt.

### Example

```go
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

	resp, err := client.System.RekeyAttemptReadProgress(
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



[**RekeyAttemptReadProgressResponse**](RekeyAttemptReadProgressResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyAttemptUpdate

Enter a single unseal key share to progress the rekey of the Vault.

### Example

```go
package main

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

	request := schema.NewRekeyAttemptUpdateRequestWithDefaults()
	resp, err := client.System.RekeyAttemptUpdate(
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
 **rekeyAttemptUpdateRequest** | [**RekeyAttemptUpdateRequest**](RekeyAttemptUpdateRequest.md) |  | 

[**RekeyAttemptUpdateResponse**](RekeyAttemptUpdateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyDeleteBackupKey

Delete the backup copy of PGP-encrypted unseal keys.

### Example

```go
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

	resp, err := client.System.RekeyDeleteBackupKey(
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



## RekeyDeleteBackupRecoveryKey



### Example

```go
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

	resp, err := client.System.RekeyDeleteBackupRecoveryKey(
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



## RekeyReadBackupKey

Return the backup copy of PGP-encrypted unseal keys.

### Example

```go
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

	resp, err := client.System.RekeyReadBackupKey(
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



[**RekeyReadBackupKeyResponse**](RekeyReadBackupKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyReadBackupRecoveryKey



### Example

```go
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

	resp, err := client.System.RekeyReadBackupRecoveryKey(
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



[**RekeyReadBackupRecoveryKeyResponse**](RekeyReadBackupRecoveryKeyResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyVerificationCancel

Cancel any in-progress rekey verification operation.



### Example

```go
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

	resp, err := client.System.RekeyVerificationCancel(
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



[**RekeyVerificationCancelResponse**](RekeyVerificationCancelResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyVerificationReadProgress

Read the configuration and progress of the current rekey verification attempt.

### Example

```go
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

	resp, err := client.System.RekeyVerificationReadProgress(
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



[**RekeyVerificationReadProgressResponse**](RekeyVerificationReadProgressResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RekeyVerificationUpdate

Enter a single new key share to progress the rekey verification operation.

### Example

```go
package main

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

	request := schema.NewRekeyVerificationUpdateRequestWithDefaults()
	resp, err := client.System.RekeyVerificationUpdate(
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
 **rekeyVerificationUpdateRequest** | [**RekeyVerificationUpdateRequest**](RekeyVerificationUpdateRequest.md) |  | 

[**RekeyVerificationUpdateResponse**](RekeyVerificationUpdateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReloadSubsystem

Reload the given subsystem

### Example

```go
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

	subsystem := "subsystem_example" // string | 
	resp, err := client.System.ReloadSubsystem(
		context.Background(),
		subsystem,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**subsystem** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Remount

Initiate a mount migration

### Example

```go
package main

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

	request := schema.NewRemountRequestWithDefaults()
	resp, err := client.System.Remount(
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
 **remountRequest** | [**RemountRequest**](RemountRequest.md) |  | 

[**RemountResponse**](RemountResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RemountStatus

Check status of a mount migration

### Example

```go
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

	migrationId := "migrationId_example" // string | The ID of the migration operation
	resp, err := client.System.RemountStatus(
		context.Background(),
		migrationId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**migrationId** | **string** | The ID of the migration operation | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**RemountStatusResponse**](RemountStatusResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReplicationStatus



### Example

```go
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

	resp, err := client.System.ReplicationStatus(
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



## Rewrap



### Example

```go
package main

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

	request := schema.NewRewrapRequestWithDefaults()
	resp, err := client.System.Rewrap(
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
 **rewrapRequest** | [**RewrapRequest**](RewrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RootTokenGenerationCancel

Cancels any in-progress root generation attempt.

### Example

```go
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

	resp, err := client.System.RootTokenGenerationCancel(
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



## RootTokenGenerationCancel2

Cancels any in-progress root generation attempt.

### Example

```go
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

	resp, err := client.System.RootTokenGenerationCancel2(
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



## RootTokenGenerationInitialize

Initializes a new root generation attempt.



### Example

```go
package main

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

	request := schema.NewRootTokenGenerationInitializeRequestWithDefaults()
	resp, err := client.System.RootTokenGenerationInitialize(
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
 **rootTokenGenerationInitializeRequest** | [**RootTokenGenerationInitializeRequest**](RootTokenGenerationInitializeRequest.md) |  | 

[**RootTokenGenerationInitializeResponse**](RootTokenGenerationInitializeResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RootTokenGenerationInitialize2

Initializes a new root generation attempt.



### Example

```go
package main

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

	request := schema.NewRootTokenGenerationInitialize2RequestWithDefaults()
	resp, err := client.System.RootTokenGenerationInitialize2(
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
 **rootTokenGenerationInitialize2Request** | [**RootTokenGenerationInitialize2Request**](RootTokenGenerationInitialize2Request.md) |  | 

[**RootTokenGenerationInitialize2Response**](RootTokenGenerationInitialize2Response.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RootTokenGenerationReadProgress

Read the configuration and progress of the current root generation attempt.

### Example

```go
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

	resp, err := client.System.RootTokenGenerationReadProgress(
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



[**RootTokenGenerationReadProgressResponse**](RootTokenGenerationReadProgressResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RootTokenGenerationReadProgress2

Read the configuration and progress of the current root generation attempt.

### Example

```go
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

	resp, err := client.System.RootTokenGenerationReadProgress2(
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



[**RootTokenGenerationReadProgress2Response**](RootTokenGenerationReadProgress2Response.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## RootTokenGenerationUpdate

Enter a single unseal key share to progress the root generation attempt.



### Example

```go
package main

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

	request := schema.NewRootTokenGenerationUpdateRequestWithDefaults()
	resp, err := client.System.RootTokenGenerationUpdate(
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
 **rootTokenGenerationUpdateRequest** | [**RootTokenGenerationUpdateRequest**](RootTokenGenerationUpdateRequest.md) |  | 

[**RootTokenGenerationUpdateResponse**](RootTokenGenerationUpdateResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## Seal

Seal the Vault.

### Example

```go
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

	resp, err := client.System.Seal(
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



## SealStatus

Check the seal status of a Vault.

### Example

```go
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

	resp, err := client.System.SealStatus(
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



[**SealStatusResponse**](SealStatusResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## StepDownLeader

Cause the node to give up active status.



### Example

```go
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

	resp, err := client.System.StepDownLeader(
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



## UiHeadersConfigure

Configure the values to be returned for the UI header.

### Example

```go
package main

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

	header := "header_example" // string | The name of the header.
	request := schema.NewUiHeadersConfigureRequestWithDefaults()
	resp, err := client.System.UiHeadersConfigure(
		context.Background(),
		header,
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
**header** | **string** | The name of the header. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiHeadersConfigureRequest** | [**UiHeadersConfigureRequest**](UiHeadersConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UiHeadersDeleteConfiguration

Remove a UI header.

### Example

```go
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

	header := "header_example" // string | The name of the header.
	resp, err := client.System.UiHeadersDeleteConfiguration(
		context.Background(),
		header,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**header** | **string** | The name of the header. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## UiHeadersList

Return a list of configured UI headers.

### Example

```go
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

	resp, err := client.System.UiHeadersList(
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

[**UiHeadersListResponse**](UiHeadersListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## UiHeadersReadConfiguration

Return the given UI header's configuration

### Example

```go
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

	header := "header_example" // string | The name of the header.
	resp, err := client.System.UiHeadersReadConfiguration(
		context.Background(),
		header,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**header** | **string** | The name of the header. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


[**UiHeadersReadConfigurationResponse**](UiHeadersReadConfigurationResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## Unseal

Unseal the Vault.

### Example

```go
package main

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

	request := schema.NewUnsealRequestWithDefaults()
	resp, err := client.System.Unseal(
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
 **unsealRequest** | [**UnsealRequest**](UnsealRequest.md) |  | 

[**UnsealResponse**](UnsealResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## Unwrap



### Example

```go
package main

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

	request := schema.NewUnwrapRequestWithDefaults()
	resp, err := client.System.Unwrap(
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
 **unwrapRequest** | [**UnwrapRequest**](UnwrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## VersionHistory

Returns map of historical version change entries

### Example

```go
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

	resp, err := client.System.VersionHistory(
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

[**VersionHistoryResponse**](VersionHistoryResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)



## Wrap



### Example

```go
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

	resp, err := client.System.Wrap(
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



