# \SystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSysAuditPath**](SystemApi.md#DeleteSysAuditPath) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
[**DeleteSysAuthPath**](SystemApi.md#DeleteSysAuthPath) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
[**DeleteSysConfigAuditingRequestHeadersHeader**](SystemApi.md#DeleteSysConfigAuditingRequestHeadersHeader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
[**DeleteSysConfigCors**](SystemApi.md#DeleteSysConfigCors) | **Delete** /sys/config/cors | Remove any CORS settings.
[**DeleteSysConfigUiHeadersHeader**](SystemApi.md#DeleteSysConfigUiHeadersHeader) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
[**DeleteSysGenerateRoot**](SystemApi.md#DeleteSysGenerateRoot) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
[**DeleteSysGenerateRootAttempt**](SystemApi.md#DeleteSysGenerateRootAttempt) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
[**DeleteSysMountsPath**](SystemApi.md#DeleteSysMountsPath) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
[**DeleteSysPluginsCatalogName**](SystemApi.md#DeleteSysPluginsCatalogName) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
[**DeleteSysPluginsCatalogTypeName**](SystemApi.md#DeleteSysPluginsCatalogTypeName) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
[**DeleteSysPoliciesAclName**](SystemApi.md#DeleteSysPoliciesAclName) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
[**DeleteSysPoliciesPasswordName**](SystemApi.md#DeleteSysPoliciesPasswordName) | **Delete** /sys/policies/password/{name} | Delete a password policy.
[**DeleteSysPolicyName**](SystemApi.md#DeleteSysPolicyName) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
[**DeleteSysQuotasRateLimitName**](SystemApi.md#DeleteSysQuotasRateLimitName) | **Delete** /sys/quotas/rate-limit/{name} | 
[**DeleteSysRaw**](SystemApi.md#DeleteSysRaw) | **Delete** /sys/raw | Delete the key with given path.
[**DeleteSysRawPath**](SystemApi.md#DeleteSysRawPath) | **Delete** /sys/raw/{path} | Delete the key with given path.
[**DeleteSysRekeyBackup**](SystemApi.md#DeleteSysRekeyBackup) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
[**DeleteSysRekeyInit**](SystemApi.md#DeleteSysRekeyInit) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
[**DeleteSysRekeyRecoveryKeyBackup**](SystemApi.md#DeleteSysRekeyRecoveryKeyBackup) | **Delete** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**DeleteSysRekeyVerify**](SystemApi.md#DeleteSysRekeyVerify) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
[**GetSysAudit**](SystemApi.md#GetSysAudit) | **Get** /sys/audit | List the enabled audit devices.
[**GetSysAuth**](SystemApi.md#GetSysAuth) | **Get** /sys/auth | List the currently enabled credential backends.
[**GetSysAuthPath**](SystemApi.md#GetSysAuthPath) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
[**GetSysAuthPathTune**](SystemApi.md#GetSysAuthPathTune) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#39;s configuration.
[**GetSysConfigAuditingRequestHeaders**](SystemApi.md#GetSysConfigAuditingRequestHeaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
[**GetSysConfigAuditingRequestHeadersHeader**](SystemApi.md#GetSysConfigAuditingRequestHeadersHeader) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
[**GetSysConfigCors**](SystemApi.md#GetSysConfigCors) | **Get** /sys/config/cors | Return the current CORS settings.
[**GetSysConfigStateSanitized**](SystemApi.md#GetSysConfigStateSanitized) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
[**GetSysConfigUiHeaders**](SystemApi.md#GetSysConfigUiHeaders) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
[**GetSysConfigUiHeadersHeader**](SystemApi.md#GetSysConfigUiHeadersHeader) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#39;s configuration
[**GetSysGenerateRoot**](SystemApi.md#GetSysGenerateRoot) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
[**GetSysGenerateRootAttempt**](SystemApi.md#GetSysGenerateRootAttempt) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
[**GetSysHaStatus**](SystemApi.md#GetSysHaStatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
[**GetSysHealth**](SystemApi.md#GetSysHealth) | **Get** /sys/health | Returns the health status of Vault.
[**GetSysHostInfo**](SystemApi.md#GetSysHostInfo) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
[**GetSysInFlightReq**](SystemApi.md#GetSysInFlightReq) | **Get** /sys/in-flight-req | reports in-flight requests
[**GetSysInit**](SystemApi.md#GetSysInit) | **Get** /sys/init | Returns the initialization status of Vault.
[**GetSysInternalCountersActivity**](SystemApi.md#GetSysInternalCountersActivity) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
[**GetSysInternalCountersActivityExport**](SystemApi.md#GetSysInternalCountersActivityExport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
[**GetSysInternalCountersActivityMonthly**](SystemApi.md#GetSysInternalCountersActivityMonthly) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
[**GetSysInternalCountersConfig**](SystemApi.md#GetSysInternalCountersConfig) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
[**GetSysInternalCountersEntities**](SystemApi.md#GetSysInternalCountersEntities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
[**GetSysInternalCountersRequests**](SystemApi.md#GetSysInternalCountersRequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
[**GetSysInternalCountersTokens**](SystemApi.md#GetSysInternalCountersTokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
[**GetSysInternalSpecsOpenapi**](SystemApi.md#GetSysInternalSpecsOpenapi) | **Get** /sys/internal/specs/openapi | Generate an OpenAPI 3 document of all mounted paths.
[**GetSysInternalUiFeatureFlags**](SystemApi.md#GetSysInternalUiFeatureFlags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
[**GetSysInternalUiMounts**](SystemApi.md#GetSysInternalUiMounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
[**GetSysInternalUiMountsPath**](SystemApi.md#GetSysInternalUiMountsPath) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
[**GetSysInternalUiNamespaces**](SystemApi.md#GetSysInternalUiNamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
[**GetSysInternalUiResultantAcl**](SystemApi.md#GetSysInternalUiResultantAcl) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
[**GetSysKeyStatus**](SystemApi.md#GetSysKeyStatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
[**GetSysLeader**](SystemApi.md#GetSysLeader) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
[**GetSysLeases**](SystemApi.md#GetSysLeases) | **Get** /sys/leases | List leases associated with this Vault cluster
[**GetSysLeasesCount**](SystemApi.md#GetSysLeasesCount) | **Get** /sys/leases/count | Count of leases associated with this Vault cluster
[**GetSysLeasesLookup**](SystemApi.md#GetSysLeasesLookup) | **Get** /sys/leases/lookup/ | Returns a list of lease ids.
[**GetSysLeasesLookupPrefix**](SystemApi.md#GetSysLeasesLookupPrefix) | **Get** /sys/leases/lookup/{prefix} | Returns a list of lease ids.
[**GetSysMetrics**](SystemApi.md#GetSysMetrics) | **Get** /sys/metrics | Export the metrics aggregated for telemetry purpose.
[**GetSysMonitor**](SystemApi.md#GetSysMonitor) | **Get** /sys/monitor | 
[**GetSysMounts**](SystemApi.md#GetSysMounts) | **Get** /sys/mounts | List the currently mounted backends.
[**GetSysMountsPath**](SystemApi.md#GetSysMountsPath) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
[**GetSysMountsPathTune**](SystemApi.md#GetSysMountsPathTune) | **Get** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**GetSysPluginsCatalog**](SystemApi.md#GetSysPluginsCatalog) | **Get** /sys/plugins/catalog | Lists all the plugins known to Vault
[**GetSysPluginsCatalogName**](SystemApi.md#GetSysPluginsCatalogName) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
[**GetSysPluginsCatalogType**](SystemApi.md#GetSysPluginsCatalogType) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
[**GetSysPluginsCatalogTypeName**](SystemApi.md#GetSysPluginsCatalogTypeName) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
[**GetSysPoliciesAcl**](SystemApi.md#GetSysPoliciesAcl) | **Get** /sys/policies/acl | List the configured access control policies.
[**GetSysPoliciesAclName**](SystemApi.md#GetSysPoliciesAclName) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
[**GetSysPoliciesPassword**](SystemApi.md#GetSysPoliciesPassword) | **Get** /sys/policies/password | List the existing password policies.
[**GetSysPoliciesPasswordName**](SystemApi.md#GetSysPoliciesPasswordName) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
[**GetSysPoliciesPasswordNameGenerate**](SystemApi.md#GetSysPoliciesPasswordNameGenerate) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
[**GetSysPolicy**](SystemApi.md#GetSysPolicy) | **Get** /sys/policy | List the configured access control policies.
[**GetSysPolicyName**](SystemApi.md#GetSysPolicyName) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
[**GetSysPprof**](SystemApi.md#GetSysPprof) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
[**GetSysPprofAllocs**](SystemApi.md#GetSysPprofAllocs) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
[**GetSysPprofBlock**](SystemApi.md#GetSysPprofBlock) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
[**GetSysPprofCmdline**](SystemApi.md#GetSysPprofCmdline) | **Get** /sys/pprof/cmdline | Returns the running program&#39;s command line.
[**GetSysPprofGoroutine**](SystemApi.md#GetSysPprofGoroutine) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
[**GetSysPprofHeap**](SystemApi.md#GetSysPprofHeap) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
[**GetSysPprofMutex**](SystemApi.md#GetSysPprofMutex) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
[**GetSysPprofProfile**](SystemApi.md#GetSysPprofProfile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
[**GetSysPprofSymbol**](SystemApi.md#GetSysPprofSymbol) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
[**GetSysPprofThreadcreate**](SystemApi.md#GetSysPprofThreadcreate) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
[**GetSysPprofTrace**](SystemApi.md#GetSysPprofTrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
[**GetSysQuotasConfig**](SystemApi.md#GetSysQuotasConfig) | **Get** /sys/quotas/config | 
[**GetSysQuotasRateLimit**](SystemApi.md#GetSysQuotasRateLimit) | **Get** /sys/quotas/rate-limit | 
[**GetSysQuotasRateLimitName**](SystemApi.md#GetSysQuotasRateLimitName) | **Get** /sys/quotas/rate-limit/{name} | 
[**GetSysRaw**](SystemApi.md#GetSysRaw) | **Get** /sys/raw | Read the value of the key at the given path.
[**GetSysRawPath**](SystemApi.md#GetSysRawPath) | **Get** /sys/raw/{path} | Read the value of the key at the given path.
[**GetSysRekeyBackup**](SystemApi.md#GetSysRekeyBackup) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
[**GetSysRekeyInit**](SystemApi.md#GetSysRekeyInit) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
[**GetSysRekeyRecoveryKeyBackup**](SystemApi.md#GetSysRekeyRecoveryKeyBackup) | **Get** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**GetSysRekeyVerify**](SystemApi.md#GetSysRekeyVerify) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
[**GetSysRemountStatusMigrationId**](SystemApi.md#GetSysRemountStatusMigrationId) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
[**GetSysReplicationStatus**](SystemApi.md#GetSysReplicationStatus) | **Get** /sys/replication/status | 
[**GetSysRotateConfig**](SystemApi.md#GetSysRotateConfig) | **Get** /sys/rotate/config | 
[**GetSysSealStatus**](SystemApi.md#GetSysSealStatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
[**GetSysVersionHistory**](SystemApi.md#GetSysVersionHistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
[**GetSysWrappingLookup**](SystemApi.md#GetSysWrappingLookup) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#39;s token.
[**PostSysAuditHashPath**](SystemApi.md#PostSysAuditHashPath) | **Post** /sys/audit-hash/{path} | The hash of the given string via the given audit backend
[**PostSysAuditPath**](SystemApi.md#PostSysAuditPath) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
[**PostSysAuthPath**](SystemApi.md#PostSysAuthPath) | **Post** /sys/auth/{path} | Enables a new auth method.
[**PostSysAuthPathTune**](SystemApi.md#PostSysAuthPathTune) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
[**PostSysCapabilities**](SystemApi.md#PostSysCapabilities) | **Post** /sys/capabilities | Fetches the capabilities of the given token on the given path.
[**PostSysCapabilitiesAccessor**](SystemApi.md#PostSysCapabilitiesAccessor) | **Post** /sys/capabilities-accessor | Fetches the capabilities of the token associated with the given token, on the given path.
[**PostSysCapabilitiesSelf**](SystemApi.md#PostSysCapabilitiesSelf) | **Post** /sys/capabilities-self | Fetches the capabilities of the given token on the given path.
[**PostSysConfigAuditingRequestHeadersHeader**](SystemApi.md#PostSysConfigAuditingRequestHeadersHeader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
[**PostSysConfigCors**](SystemApi.md#PostSysConfigCors) | **Post** /sys/config/cors | Configure the CORS settings.
[**PostSysConfigReloadSubsystem**](SystemApi.md#PostSysConfigReloadSubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
[**PostSysConfigUiHeadersHeader**](SystemApi.md#PostSysConfigUiHeadersHeader) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
[**PostSysGenerateRoot**](SystemApi.md#PostSysGenerateRoot) | **Post** /sys/generate-root | Initializes a new root generation attempt.
[**PostSysGenerateRootAttempt**](SystemApi.md#PostSysGenerateRootAttempt) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
[**PostSysGenerateRootUpdate**](SystemApi.md#PostSysGenerateRootUpdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
[**PostSysInit**](SystemApi.md#PostSysInit) | **Post** /sys/init | Initialize a new Vault.
[**PostSysInternalCountersConfig**](SystemApi.md#PostSysInternalCountersConfig) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
[**PostSysLeasesLookup**](SystemApi.md#PostSysLeasesLookup) | **Post** /sys/leases/lookup | Retrieve lease metadata.
[**PostSysLeasesRenew**](SystemApi.md#PostSysLeasesRenew) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
[**PostSysLeasesRenewUrlLeaseId**](SystemApi.md#PostSysLeasesRenewUrlLeaseId) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**PostSysLeasesRevoke**](SystemApi.md#PostSysLeasesRevoke) | **Post** /sys/leases/revoke | Revokes a lease immediately.
[**PostSysLeasesRevokeForcePrefix**](SystemApi.md#PostSysLeasesRevokeForcePrefix) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**PostSysLeasesRevokePrefixPrefix**](SystemApi.md#PostSysLeasesRevokePrefixPrefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
[**PostSysLeasesRevokeUrlLeaseId**](SystemApi.md#PostSysLeasesRevokeUrlLeaseId) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
[**PostSysLeasesTidy**](SystemApi.md#PostSysLeasesTidy) | **Post** /sys/leases/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**PostSysMfaValidate**](SystemApi.md#PostSysMfaValidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
[**PostSysMountsPath**](SystemApi.md#PostSysMountsPath) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
[**PostSysMountsPathTune**](SystemApi.md#PostSysMountsPathTune) | **Post** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**PostSysPluginsCatalogName**](SystemApi.md#PostSysPluginsCatalogName) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PostSysPluginsCatalogTypeName**](SystemApi.md#PostSysPluginsCatalogTypeName) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PostSysPluginsReloadBackend**](SystemApi.md#PostSysPluginsReloadBackend) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
[**PostSysPoliciesAclName**](SystemApi.md#PostSysPoliciesAclName) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
[**PostSysPoliciesPasswordName**](SystemApi.md#PostSysPoliciesPasswordName) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
[**PostSysPolicyName**](SystemApi.md#PostSysPolicyName) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
[**PostSysQuotasConfig**](SystemApi.md#PostSysQuotasConfig) | **Post** /sys/quotas/config | 
[**PostSysQuotasRateLimitName**](SystemApi.md#PostSysQuotasRateLimitName) | **Post** /sys/quotas/rate-limit/{name} | 
[**PostSysRaw**](SystemApi.md#PostSysRaw) | **Post** /sys/raw | Update the value of the key at the given path.
[**PostSysRawPath**](SystemApi.md#PostSysRawPath) | **Post** /sys/raw/{path} | Update the value of the key at the given path.
[**PostSysRekeyInit**](SystemApi.md#PostSysRekeyInit) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
[**PostSysRekeyUpdate**](SystemApi.md#PostSysRekeyUpdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
[**PostSysRekeyVerify**](SystemApi.md#PostSysRekeyVerify) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
[**PostSysRemount**](SystemApi.md#PostSysRemount) | **Post** /sys/remount | Initiate a mount migration
[**PostSysRenew**](SystemApi.md#PostSysRenew) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
[**PostSysRenewUrlLeaseId**](SystemApi.md#PostSysRenewUrlLeaseId) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**PostSysRevoke**](SystemApi.md#PostSysRevoke) | **Post** /sys/revoke | Revokes a lease immediately.
[**PostSysRevokeForcePrefix**](SystemApi.md#PostSysRevokeForcePrefix) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**PostSysRevokePrefixPrefix**](SystemApi.md#PostSysRevokePrefixPrefix) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
[**PostSysRevokeUrlLeaseId**](SystemApi.md#PostSysRevokeUrlLeaseId) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
[**PostSysRotate**](SystemApi.md#PostSysRotate) | **Post** /sys/rotate | Rotates the backend encryption key used to persist data.
[**PostSysRotateConfig**](SystemApi.md#PostSysRotateConfig) | **Post** /sys/rotate/config | 
[**PostSysSeal**](SystemApi.md#PostSysSeal) | **Post** /sys/seal | Seal the Vault.
[**PostSysStepDown**](SystemApi.md#PostSysStepDown) | **Post** /sys/step-down | Cause the node to give up active status.
[**PostSysToolsHash**](SystemApi.md#PostSysToolsHash) | **Post** /sys/tools/hash | Generate a hash sum for input data
[**PostSysToolsHashUrlalgorithm**](SystemApi.md#PostSysToolsHashUrlalgorithm) | **Post** /sys/tools/hash/{urlalgorithm} | Generate a hash sum for input data
[**PostSysToolsRandom**](SystemApi.md#PostSysToolsRandom) | **Post** /sys/tools/random | Generate random bytes
[**PostSysToolsRandomSource**](SystemApi.md#PostSysToolsRandomSource) | **Post** /sys/tools/random/{source} | Generate random bytes
[**PostSysToolsRandomSourceUrlbytes**](SystemApi.md#PostSysToolsRandomSourceUrlbytes) | **Post** /sys/tools/random/{source}/{urlbytes} | Generate random bytes
[**PostSysToolsRandomUrlbytes**](SystemApi.md#PostSysToolsRandomUrlbytes) | **Post** /sys/tools/random/{urlbytes} | Generate random bytes
[**PostSysUnseal**](SystemApi.md#PostSysUnseal) | **Post** /sys/unseal | Unseal the Vault.
[**PostSysWrappingLookup**](SystemApi.md#PostSysWrappingLookup) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
[**PostSysWrappingRewrap**](SystemApi.md#PostSysWrappingRewrap) | **Post** /sys/wrapping/rewrap | Rotates a response-wrapped token.
[**PostSysWrappingUnwrap**](SystemApi.md#PostSysWrappingUnwrap) | **Post** /sys/wrapping/unwrap | Unwraps a response-wrapped token.
[**PostSysWrappingWrap**](SystemApi.md#PostSysWrappingWrap) | **Post** /sys/wrapping/wrap | Response-wraps an arbitrary JSON object.



## DeleteSysAuditPath

> DeleteSysAuditPath(ctx, path).Execute()

Disable the audit device at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysAuditPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysAuditPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysAuditPathRequest struct via the builder pattern


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


## DeleteSysAuthPath

> DeleteSysAuthPath(ctx, path).Execute()

Disable the auth method at the given auth path

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysAuthPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysAuthPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysAuthPathRequest struct via the builder pattern


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


## DeleteSysConfigAuditingRequestHeadersHeader

> DeleteSysConfigAuditingRequestHeadersHeader(ctx, header).Execute()

Disable auditing of the given request header.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysConfigAuditingRequestHeadersHeader(context.Background(), header).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysConfigAuditingRequestHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


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


## DeleteSysConfigCors

> DeleteSysConfigCors(ctx).Execute()

Remove any CORS settings.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysConfigCors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysConfigCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigCorsRequest struct via the builder pattern


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


## DeleteSysConfigUiHeadersHeader

> DeleteSysConfigUiHeadersHeader(ctx, header).Execute()

Remove a UI header.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | The name of the header.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysConfigUiHeadersHeader(context.Background(), header).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysConfigUiHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigUiHeadersHeaderRequest struct via the builder pattern


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


## DeleteSysGenerateRoot

> DeleteSysGenerateRoot(ctx).Execute()

Cancels any in-progress root generation attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysGenerateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysGenerateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysGenerateRootRequest struct via the builder pattern


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


## DeleteSysGenerateRootAttempt

> DeleteSysGenerateRootAttempt(ctx).Execute()

Cancels any in-progress root generation attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysGenerateRootAttempt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysGenerateRootAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysGenerateRootAttemptRequest struct via the builder pattern


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


## DeleteSysMountsPath

> DeleteSysMountsPath(ctx, path).Execute()

Disable the mount point specified at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Example: \"aws/east\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysMountsPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysMountsPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysMountsPathRequest struct via the builder pattern


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


## DeleteSysPluginsCatalogName

> DeleteSysPluginsCatalogName(ctx, name).Execute()

Remove the plugin with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysPluginsCatalogName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysPluginsCatalogName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPluginsCatalogNameRequest struct via the builder pattern


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


## DeleteSysPluginsCatalogTypeName

> DeleteSysPluginsCatalogTypeName(ctx, name, type_).Execute()

Remove the plugin with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin
    type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysPluginsCatalogTypeName(context.Background(), name, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysPluginsCatalogTypeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPluginsCatalogTypeNameRequest struct via the builder pattern


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


## DeleteSysPoliciesAclName

> DeleteSysPoliciesAclName(ctx, name).Execute()

Delete the ACL policy with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysPoliciesAclName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysPoliciesAclName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPoliciesAclNameRequest struct via the builder pattern


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


## DeleteSysPoliciesPasswordName

> DeleteSysPoliciesPasswordName(ctx, name).Execute()

Delete a password policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the password policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysPoliciesPasswordName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysPoliciesPasswordName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPoliciesPasswordNameRequest struct via the builder pattern


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


## DeleteSysPolicyName

> DeleteSysPolicyName(ctx, name).Execute()

Delete the policy with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysPolicyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysPolicyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPolicyNameRequest struct via the builder pattern


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


## DeleteSysQuotasRateLimitName

> DeleteSysQuotasRateLimitName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the quota rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysQuotasRateLimitName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysQuotasRateLimitName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysQuotasRateLimitNameRequest struct via the builder pattern


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


## DeleteSysRaw

> DeleteSysRaw(ctx).Execute()

Delete the key with given path.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysRaw(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRawRequest struct via the builder pattern


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


## DeleteSysRawPath

> DeleteSysRawPath(ctx, path).Execute()

Delete the key with given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteSysRawPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRawPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRawPathRequest struct via the builder pattern


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


## DeleteSysRekeyBackup

> DeleteSysRekeyBackup(ctx).Execute()

Delete the backup copy of PGP-encrypted unseal keys.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysRekeyBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRekeyBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyBackupRequest struct via the builder pattern


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


## DeleteSysRekeyInit

> DeleteSysRekeyInit(ctx).Execute()

Cancels any in-progress rekey.



### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysRekeyInit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRekeyInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyInitRequest struct via the builder pattern


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


## DeleteSysRekeyRecoveryKeyBackup

> DeleteSysRekeyRecoveryKeyBackup(ctx).Execute()

Allows fetching or deleting the backup of the rotated unseal keys.

### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysRekeyRecoveryKeyBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRekeyRecoveryKeyBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyRecoveryKeyBackupRequest struct via the builder pattern


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


## DeleteSysRekeyVerify

> DeleteSysRekeyVerify(ctx).Execute()

Cancel any in-progress rekey verification operation.



### Example

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
    resp, r, err := apiClient.SystemApi.DeleteSysRekeyVerify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteSysRekeyVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyVerifyRequest struct via the builder pattern


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


## GetSysAudit

> GetSysAudit(ctx).Execute()

List the enabled audit devices.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysAudit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuditRequest struct via the builder pattern


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


## GetSysAuth

> GetSysAuth(ctx).Execute()

List the currently enabled credential backends.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysAuth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthRequest struct via the builder pattern


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


## GetSysAuthPath

> GetSysAuthPath(ctx, path).Execute()

Read the configuration of the auth engine at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysAuthPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysAuthPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthPathRequest struct via the builder pattern


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


## GetSysAuthPathTune

> GetSysAuthPathTune(ctx, path).Execute()

Reads the given auth path's configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Tune the configuration parameters for an auth path.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysAuthPathTune(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysAuthPathTune``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthPathTuneRequest struct via the builder pattern


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


## GetSysConfigAuditingRequestHeaders

> GetSysConfigAuditingRequestHeaders(ctx).Execute()

List the request headers that are configured to be audited.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysConfigAuditingRequestHeaders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigAuditingRequestHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigAuditingRequestHeadersRequest struct via the builder pattern


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


## GetSysConfigAuditingRequestHeadersHeader

> GetSysConfigAuditingRequestHeadersHeader(ctx, header).Execute()

List the information for the given request header.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysConfigAuditingRequestHeadersHeader(context.Background(), header).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigAuditingRequestHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


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


## GetSysConfigCors

> GetSysConfigCors(ctx).Execute()

Return the current CORS settings.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysConfigCors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigCorsRequest struct via the builder pattern


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


## GetSysConfigStateSanitized

> GetSysConfigStateSanitized(ctx).Execute()

Return a sanitized version of the Vault server configuration.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysConfigStateSanitized(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigStateSanitized``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigStateSanitizedRequest struct via the builder pattern


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


## GetSysConfigUiHeaders

> GetSysConfigUiHeaders(ctx).List(list).Execute()

Return a list of configured UI headers.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysConfigUiHeaders(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigUiHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigUiHeadersRequest struct via the builder pattern


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


## GetSysConfigUiHeadersHeader

> GetSysConfigUiHeadersHeader(ctx, header).Execute()

Return the given UI header's configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | The name of the header.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysConfigUiHeadersHeader(context.Background(), header).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysConfigUiHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigUiHeadersHeaderRequest struct via the builder pattern


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


## GetSysGenerateRoot

> GetSysGenerateRoot(ctx).Execute()

Read the configuration and progress of the current root generation attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysGenerateRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysGenerateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysGenerateRootRequest struct via the builder pattern


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


## GetSysGenerateRootAttempt

> GetSysGenerateRootAttempt(ctx).Execute()

Read the configuration and progress of the current root generation attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysGenerateRootAttempt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysGenerateRootAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysGenerateRootAttemptRequest struct via the builder pattern


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


## GetSysHaStatus

> GetSysHaStatus(ctx).Execute()

Check the HA status of a Vault cluster

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysHaStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysHaStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHaStatusRequest struct via the builder pattern


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


## GetSysHealth

> GetSysHealth(ctx).Execute()

Returns the health status of Vault.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysHealth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHealthRequest struct via the builder pattern


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


## GetSysHostInfo

> GetSysHostInfo(ctx).Execute()

Information about the host instance that this Vault server is running on.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysHostInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysHostInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHostInfoRequest struct via the builder pattern


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


## GetSysInFlightReq

> GetSysInFlightReq(ctx).Execute()

reports in-flight requests



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInFlightReq(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInFlightReq``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInFlightReqRequest struct via the builder pattern


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


## GetSysInit

> GetSysInit(ctx).Execute()

Returns the initialization status of Vault.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInitRequest struct via the builder pattern


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


## GetSysInternalCountersActivity

> GetSysInternalCountersActivity(ctx).Execute()

Report the client count metrics, for this namespace and all child namespaces.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersActivity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityRequest struct via the builder pattern


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


## GetSysInternalCountersActivityExport

> GetSysInternalCountersActivityExport(ctx).Execute()

Report the client count metrics, for this namespace and all child namespaces.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersActivityExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersActivityExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityExportRequest struct via the builder pattern


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


## GetSysInternalCountersActivityMonthly

> GetSysInternalCountersActivityMonthly(ctx).Execute()

Report the number of clients for this month, for this namespace and all child namespaces.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersActivityMonthly(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersActivityMonthly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityMonthlyRequest struct via the builder pattern


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


## GetSysInternalCountersConfig

> GetSysInternalCountersConfig(ctx).Execute()

Read the client count tracking configuration.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersConfigRequest struct via the builder pattern


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


## GetSysInternalCountersEntities

> GetSysInternalCountersEntities(ctx).Execute()

Backwards compatibility is not guaranteed for this API

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersEntities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersEntitiesRequest struct via the builder pattern


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


## GetSysInternalCountersRequests

> GetSysInternalCountersRequests(ctx).Execute()

Backwards compatibility is not guaranteed for this API

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersRequestsRequest struct via the builder pattern


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


## GetSysInternalCountersTokens

> GetSysInternalCountersTokens(ctx).Execute()

Backwards compatibility is not guaranteed for this API

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalCountersTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalCountersTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersTokensRequest struct via the builder pattern


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


## GetSysInternalSpecsOpenapi

> GetSysInternalSpecsOpenapi(ctx).Execute()

Generate an OpenAPI 3 document of all mounted paths.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalSpecsOpenapi(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalSpecsOpenapi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalSpecsOpenapiRequest struct via the builder pattern


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


## GetSysInternalUiFeatureFlags

> GetSysInternalUiFeatureFlags(ctx).Execute()

Lists enabled feature flags.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalUiFeatureFlags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalUiFeatureFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiFeatureFlagsRequest struct via the builder pattern


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


## GetSysInternalUiMounts

> GetSysInternalUiMounts(ctx).Execute()

Lists all enabled and visible auth and secrets mounts.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalUiMounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalUiMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiMountsRequest struct via the builder pattern


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


## GetSysInternalUiMountsPath

> GetSysInternalUiMountsPath(ctx, path).Execute()

Return information about the given mount.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path of the mount.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysInternalUiMountsPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalUiMountsPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path of the mount. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiMountsPathRequest struct via the builder pattern


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


## GetSysInternalUiNamespaces

> GetSysInternalUiNamespaces(ctx).Execute()

Backwards compatibility is not guaranteed for this API

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalUiNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalUiNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiNamespacesRequest struct via the builder pattern


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


## GetSysInternalUiResultantAcl

> GetSysInternalUiResultantAcl(ctx).Execute()

Backwards compatibility is not guaranteed for this API

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysInternalUiResultantAcl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysInternalUiResultantAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiResultantAclRequest struct via the builder pattern


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


## GetSysKeyStatus

> GetSysKeyStatus(ctx).Execute()

Provides information about the backend encryption key.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysKeyStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysKeyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysKeyStatusRequest struct via the builder pattern


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


## GetSysLeader

> GetSysLeader(ctx).Execute()

Returns the high availability status and current leader instance of Vault.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysLeader(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysLeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeaderRequest struct via the builder pattern


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


## GetSysLeases

> GetSysLeases(ctx).Execute()

List leases associated with this Vault cluster

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysLeases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysLeases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesRequest struct via the builder pattern


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


## GetSysLeasesCount

> GetSysLeasesCount(ctx).Execute()

Count of leases associated with this Vault cluster

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysLeasesCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysLeasesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesCountRequest struct via the builder pattern


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


## GetSysLeasesLookup

> GetSysLeasesLookup(ctx).List(list).Execute()

Returns a list of lease ids.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysLeasesLookup(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysLeasesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesLookupRequest struct via the builder pattern


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


## GetSysLeasesLookupPrefix

> GetSysLeasesLookupPrefix(ctx, prefix).List(list).Execute()

Returns a list of lease ids.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The path to list leases under. Example: \"aws/creds/deploy\"
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysLeasesLookupPrefix(context.Background(), prefix).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysLeasesLookupPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to list leases under. Example: \&quot;aws/creds/deploy\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesLookupPrefixRequest struct via the builder pattern


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


## GetSysMetrics

> GetSysMetrics(ctx).Format(format).Execute()

Export the metrics aggregated for telemetry purpose.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    format := "format_example" // string | Format to export metrics into. Currently accepts only \"prometheus\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysMetrics(context.Background()).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format to export metrics into. Currently accepts only \&quot;prometheus\&quot;. | 

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


## GetSysMonitor

> GetSysMonitor(ctx).LogFormat(logFormat).LogLevel(logLevel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    logFormat := "logFormat_example" // string | Output format of logs. Supported values are \"standard\" and \"json\". The default is \"standard\". (optional) (default to "standard")
    logLevel := "logLevel_example" // string | Log level to view system logs at. Currently supported values are \"trace\", \"debug\", \"info\", \"warn\", \"error\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysMonitor(context.Background()).LogFormat(logFormat).LogLevel(logLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logFormat** | **string** | Output format of logs. Supported values are \&quot;standard\&quot; and \&quot;json\&quot;. The default is \&quot;standard\&quot;. | [default to &quot;standard&quot;]
 **logLevel** | **string** | Log level to view system logs at. Currently supported values are \&quot;trace\&quot;, \&quot;debug\&quot;, \&quot;info\&quot;, \&quot;warn\&quot;, \&quot;error\&quot;. | 

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


## GetSysMounts

> GetSysMounts(ctx).Execute()

List the currently mounted backends.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysMounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsRequest struct via the builder pattern


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


## GetSysMountsPath

> GetSysMountsPath(ctx, path).Execute()

Read the configuration of the secret engine at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Example: \"aws/east\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysMountsPath(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysMountsPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsPathRequest struct via the builder pattern


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


## GetSysMountsPathTune

> GetSysMountsPathTune(ctx, path).Execute()

Tune backend configuration parameters for this mount.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Example: \"aws/east\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysMountsPathTune(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysMountsPathTune``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsPathTuneRequest struct via the builder pattern


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


## GetSysPluginsCatalog

> GetSysPluginsCatalog(ctx).Execute()

Lists all the plugins known to Vault

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPluginsCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPluginsCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogRequest struct via the builder pattern


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


## GetSysPluginsCatalogName

> GetSysPluginsCatalogName(ctx, name).Execute()

Return the configuration data for the plugin with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPluginsCatalogName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPluginsCatalogName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogNameRequest struct via the builder pattern


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


## GetSysPluginsCatalogType

> GetSysPluginsCatalogType(ctx, type_).List(list).Execute()

List the plugins in the catalog.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPluginsCatalogType(context.Background(), type_).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPluginsCatalogType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogTypeRequest struct via the builder pattern


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


## GetSysPluginsCatalogTypeName

> GetSysPluginsCatalogTypeName(ctx, name, type_).Execute()

Return the configuration data for the plugin with the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin
    type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPluginsCatalogTypeName(context.Background(), name, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPluginsCatalogTypeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogTypeNameRequest struct via the builder pattern


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


## GetSysPoliciesAcl

> GetSysPoliciesAcl(ctx).List(list).Execute()

List the configured access control policies.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPoliciesAcl(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPoliciesAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesAclRequest struct via the builder pattern


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


## GetSysPoliciesAclName

> GetSysPoliciesAclName(ctx, name).Execute()

Retrieve information about the named ACL policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPoliciesAclName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPoliciesAclName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesAclNameRequest struct via the builder pattern


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


## GetSysPoliciesPassword

> GetSysPoliciesPassword(ctx).List(list).Execute()

List the existing password policies.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPoliciesPassword(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPoliciesPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesPasswordRequest struct via the builder pattern


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


## GetSysPoliciesPasswordName

> GetSysPoliciesPasswordName(ctx, name).Execute()

Retrieve an existing password policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the password policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPoliciesPasswordName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPoliciesPasswordName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesPasswordNameRequest struct via the builder pattern


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


## GetSysPoliciesPasswordNameGenerate

> GetSysPoliciesPasswordNameGenerate(ctx, name).Execute()

Generate a password from an existing password policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the password policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPoliciesPasswordNameGenerate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPoliciesPasswordNameGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesPasswordNameGenerateRequest struct via the builder pattern


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


## GetSysPolicy

> GetSysPolicy(ctx).List(list).Execute()

List the configured access control policies.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPolicy(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPolicyRequest struct via the builder pattern


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


## GetSysPolicyName

> GetSysPolicyName(ctx, name).Execute()

Retrieve the policy body for the named policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysPolicyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPolicyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPolicyNameRequest struct via the builder pattern


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


## GetSysPprof

> GetSysPprof(ctx).Execute()

Returns an HTML page listing the available profiles.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprof(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprof``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofRequest struct via the builder pattern


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


## GetSysPprofAllocs

> GetSysPprofAllocs(ctx).Execute()

Returns a sampling of all past memory allocations.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofAllocs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofAllocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofAllocsRequest struct via the builder pattern


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


## GetSysPprofBlock

> GetSysPprofBlock(ctx).Execute()

Returns stack traces that led to blocking on synchronization primitives



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofBlock(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofBlockRequest struct via the builder pattern


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


## GetSysPprofCmdline

> GetSysPprofCmdline(ctx).Execute()

Returns the running program's command line.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofCmdline(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofCmdline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofCmdlineRequest struct via the builder pattern


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


## GetSysPprofGoroutine

> GetSysPprofGoroutine(ctx).Execute()

Returns stack traces of all current goroutines.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofGoroutine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofGoroutine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofGoroutineRequest struct via the builder pattern


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


## GetSysPprofHeap

> GetSysPprofHeap(ctx).Execute()

Returns a sampling of memory allocations of live object.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofHeap(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofHeap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofHeapRequest struct via the builder pattern


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


## GetSysPprofMutex

> GetSysPprofMutex(ctx).Execute()

Returns stack traces of holders of contended mutexes



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofMutex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofMutex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofMutexRequest struct via the builder pattern


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


## GetSysPprofProfile

> GetSysPprofProfile(ctx).Execute()

Returns a pprof-formatted cpu profile payload.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofProfileRequest struct via the builder pattern


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


## GetSysPprofSymbol

> GetSysPprofSymbol(ctx).Execute()

Returns the program counters listed in the request.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofSymbol(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofSymbol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofSymbolRequest struct via the builder pattern


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


## GetSysPprofThreadcreate

> GetSysPprofThreadcreate(ctx).Execute()

Returns stack traces that led to the creation of new OS threads



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofThreadcreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofThreadcreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofThreadcreateRequest struct via the builder pattern


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


## GetSysPprofTrace

> GetSysPprofTrace(ctx).Execute()

Returns the execution trace in binary form.



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysPprofTrace(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysPprofTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofTraceRequest struct via the builder pattern


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


## GetSysQuotasConfig

> GetSysQuotasConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysQuotasConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysQuotasConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysQuotasConfigRequest struct via the builder pattern


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


## GetSysQuotasRateLimit

> GetSysQuotasRateLimit(ctx).List(list).Execute()



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysQuotasRateLimit(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysQuotasRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysQuotasRateLimitRequest struct via the builder pattern


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


## GetSysQuotasRateLimitName

> GetSysQuotasRateLimitName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the quota rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysQuotasRateLimitName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysQuotasRateLimitName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysQuotasRateLimitNameRequest struct via the builder pattern


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


## GetSysRaw

> GetSysRaw(ctx).List(list).Execute()

Read the value of the key at the given path.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRaw(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRawRequest struct via the builder pattern


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


## GetSysRawPath

> GetSysRawPath(ctx, path).List(list).Execute()

Read the value of the key at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | 
    list := "list_example" // string | Return a list if `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysRawPath(context.Background(), path).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRawPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRawPathRequest struct via the builder pattern


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


## GetSysRekeyBackup

> GetSysRekeyBackup(ctx).Execute()

Return the backup copy of PGP-encrypted unseal keys.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRekeyBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRekeyBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyBackupRequest struct via the builder pattern


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


## GetSysRekeyInit

> GetSysRekeyInit(ctx).Execute()

Reads the configuration and progress of the current rekey attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRekeyInit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRekeyInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyInitRequest struct via the builder pattern


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


## GetSysRekeyRecoveryKeyBackup

> GetSysRekeyRecoveryKeyBackup(ctx).Execute()

Allows fetching or deleting the backup of the rotated unseal keys.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRekeyRecoveryKeyBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRekeyRecoveryKeyBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyRecoveryKeyBackupRequest struct via the builder pattern


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


## GetSysRekeyVerify

> GetSysRekeyVerify(ctx).Execute()

Read the configuration and progress of the current rekey verification attempt.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRekeyVerify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRekeyVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyVerifyRequest struct via the builder pattern


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


## GetSysRemountStatusMigrationId

> GetSysRemountStatusMigrationId(ctx, migrationId).Execute()

Check status of a mount migration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    migrationId := "migrationId_example" // string | The ID of the migration operation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetSysRemountStatusMigrationId(context.Background(), migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRemountStatusMigrationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **string** | The ID of the migration operation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRemountStatusMigrationIdRequest struct via the builder pattern


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


## GetSysReplicationStatus

> GetSysReplicationStatus(ctx).Execute()



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysReplicationStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysReplicationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysReplicationStatusRequest struct via the builder pattern


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


## GetSysRotateConfig

> GetSysRotateConfig(ctx).Execute()



### Example

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
    resp, r, err := apiClient.SystemApi.GetSysRotateConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysRotateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRotateConfigRequest struct via the builder pattern


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


## GetSysSealStatus

> GetSysSealStatus(ctx).Execute()

Check the seal status of a Vault.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysSealStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysSealStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysSealStatusRequest struct via the builder pattern


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


## GetSysVersionHistory

> GetSysVersionHistory(ctx).List(list).Execute()

Returns map of historical version change entries

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysVersionHistory(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysVersionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSysVersionHistoryRequest struct via the builder pattern


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


## GetSysWrappingLookup

> GetSysWrappingLookup(ctx).Execute()

Look up wrapping properties for the requester's token.

### Example

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
    resp, r, err := apiClient.SystemApi.GetSysWrappingLookup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSysWrappingLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysWrappingLookupRequest struct via the builder pattern


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


## PostSysAuditHashPath

> PostSysAuditHashPath(ctx, path).SystemAuditHashRequest(systemAuditHashRequest).Execute()

The hash of the given string via the given audit backend

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
    systemAuditHashRequest := *openapiclient.NewSystemAuditHashRequest() // SystemAuditHashRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysAuditHashPath(context.Background(), path).SystemAuditHashRequest(systemAuditHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysAuditHashPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuditHashPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuditHashRequest** | [**SystemAuditHashRequest**](SystemAuditHashRequest.md) |  | 

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


## PostSysAuditPath

> PostSysAuditPath(ctx, path).SystemAuditRequest(systemAuditRequest).Execute()

Enable a new audit device at the supplied path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
    systemAuditRequest := *openapiclient.NewSystemAuditRequest() // SystemAuditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysAuditPath(context.Background(), path).SystemAuditRequest(systemAuditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysAuditPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuditPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuditRequest** | [**SystemAuditRequest**](SystemAuditRequest.md) |  | 

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


## PostSysAuthPath

> PostSysAuthPath(ctx, path).SystemAuthRequest(systemAuthRequest).Execute()

Enables a new auth method.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
    systemAuthRequest := *openapiclient.NewSystemAuthRequest() // SystemAuthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysAuthPath(context.Background(), path).SystemAuthRequest(systemAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysAuthPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuthPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuthRequest** | [**SystemAuthRequest**](SystemAuthRequest.md) |  | 

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


## PostSysAuthPathTune

> PostSysAuthPathTune(ctx, path).SystemAuthTuneRequest(systemAuthTuneRequest).Execute()

Tune configuration parameters for a given auth path.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Tune the configuration parameters for an auth path.
    systemAuthTuneRequest := *openapiclient.NewSystemAuthTuneRequest() // SystemAuthTuneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysAuthPathTune(context.Background(), path).SystemAuthTuneRequest(systemAuthTuneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysAuthPathTune``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuthPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuthTuneRequest** | [**SystemAuthTuneRequest**](SystemAuthTuneRequest.md) |  | 

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


## PostSysCapabilities

> PostSysCapabilities(ctx).SystemCapabilitiesRequest(systemCapabilitiesRequest).Execute()

Fetches the capabilities of the given token on the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemCapabilitiesRequest := *openapiclient.NewSystemCapabilitiesRequest() // SystemCapabilitiesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysCapabilities(context.Background()).SystemCapabilitiesRequest(systemCapabilitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemCapabilitiesRequest** | [**SystemCapabilitiesRequest**](SystemCapabilitiesRequest.md) |  | 

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


## PostSysCapabilitiesAccessor

> PostSysCapabilitiesAccessor(ctx).SystemCapabilitiesAccessorRequest(systemCapabilitiesAccessorRequest).Execute()

Fetches the capabilities of the token associated with the given token, on the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemCapabilitiesAccessorRequest := *openapiclient.NewSystemCapabilitiesAccessorRequest() // SystemCapabilitiesAccessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysCapabilitiesAccessor(context.Background()).SystemCapabilitiesAccessorRequest(systemCapabilitiesAccessorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysCapabilitiesAccessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysCapabilitiesAccessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemCapabilitiesAccessorRequest** | [**SystemCapabilitiesAccessorRequest**](SystemCapabilitiesAccessorRequest.md) |  | 

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


## PostSysCapabilitiesSelf

> PostSysCapabilitiesSelf(ctx).SystemCapabilitiesSelfRequest(systemCapabilitiesSelfRequest).Execute()

Fetches the capabilities of the given token on the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemCapabilitiesSelfRequest := *openapiclient.NewSystemCapabilitiesSelfRequest() // SystemCapabilitiesSelfRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysCapabilitiesSelf(context.Background()).SystemCapabilitiesSelfRequest(systemCapabilitiesSelfRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysCapabilitiesSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysCapabilitiesSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemCapabilitiesSelfRequest** | [**SystemCapabilitiesSelfRequest**](SystemCapabilitiesSelfRequest.md) |  | 

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


## PostSysConfigAuditingRequestHeadersHeader

> PostSysConfigAuditingRequestHeadersHeader(ctx, header).SystemConfigAuditingRequestHeadersRequest(systemConfigAuditingRequestHeadersRequest).Execute()

Enable auditing of a header.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | 
    systemConfigAuditingRequestHeadersRequest := *openapiclient.NewSystemConfigAuditingRequestHeadersRequest() // SystemConfigAuditingRequestHeadersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysConfigAuditingRequestHeadersHeader(context.Background(), header).SystemConfigAuditingRequestHeadersRequest(systemConfigAuditingRequestHeadersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysConfigAuditingRequestHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemConfigAuditingRequestHeadersRequest** | [**SystemConfigAuditingRequestHeadersRequest**](SystemConfigAuditingRequestHeadersRequest.md) |  | 

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


## PostSysConfigCors

> PostSysConfigCors(ctx).SystemConfigCorsRequest(systemConfigCorsRequest).Execute()

Configure the CORS settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemConfigCorsRequest := *openapiclient.NewSystemConfigCorsRequest() // SystemConfigCorsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysConfigCors(context.Background()).SystemConfigCorsRequest(systemConfigCorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysConfigCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemConfigCorsRequest** | [**SystemConfigCorsRequest**](SystemConfigCorsRequest.md) |  | 

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


## PostSysConfigReloadSubsystem

> PostSysConfigReloadSubsystem(ctx, subsystem).Execute()

Reload the given subsystem

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subsystem := "subsystem_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysConfigReloadSubsystem(context.Background(), subsystem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysConfigReloadSubsystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsystem** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigReloadSubsystemRequest struct via the builder pattern


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


## PostSysConfigUiHeadersHeader

> PostSysConfigUiHeadersHeader(ctx, header).SystemConfigUiHeadersRequest(systemConfigUiHeadersRequest).Execute()

Configure the values to be returned for the UI header.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    header := "header_example" // string | The name of the header.
    systemConfigUiHeadersRequest := *openapiclient.NewSystemConfigUiHeadersRequest() // SystemConfigUiHeadersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysConfigUiHeadersHeader(context.Background(), header).SystemConfigUiHeadersRequest(systemConfigUiHeadersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysConfigUiHeadersHeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigUiHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemConfigUiHeadersRequest** | [**SystemConfigUiHeadersRequest**](SystemConfigUiHeadersRequest.md) |  | 

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


## PostSysGenerateRoot

> PostSysGenerateRoot(ctx).SystemGenerateRootRequest(systemGenerateRootRequest).Execute()

Initializes a new root generation attempt.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemGenerateRootRequest := *openapiclient.NewSystemGenerateRootRequest() // SystemGenerateRootRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysGenerateRoot(context.Background()).SystemGenerateRootRequest(systemGenerateRootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysGenerateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysGenerateRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemGenerateRootRequest** | [**SystemGenerateRootRequest**](SystemGenerateRootRequest.md) |  | 

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


## PostSysGenerateRootAttempt

> PostSysGenerateRootAttempt(ctx).SystemGenerateRootAttemptRequest(systemGenerateRootAttemptRequest).Execute()

Initializes a new root generation attempt.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemGenerateRootAttemptRequest := *openapiclient.NewSystemGenerateRootAttemptRequest() // SystemGenerateRootAttemptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysGenerateRootAttempt(context.Background()).SystemGenerateRootAttemptRequest(systemGenerateRootAttemptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysGenerateRootAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysGenerateRootAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemGenerateRootAttemptRequest** | [**SystemGenerateRootAttemptRequest**](SystemGenerateRootAttemptRequest.md) |  | 

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


## PostSysGenerateRootUpdate

> PostSysGenerateRootUpdate(ctx).SystemGenerateRootUpdateRequest(systemGenerateRootUpdateRequest).Execute()

Enter a single unseal key share to progress the root generation attempt.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemGenerateRootUpdateRequest := *openapiclient.NewSystemGenerateRootUpdateRequest() // SystemGenerateRootUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysGenerateRootUpdate(context.Background()).SystemGenerateRootUpdateRequest(systemGenerateRootUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysGenerateRootUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysGenerateRootUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemGenerateRootUpdateRequest** | [**SystemGenerateRootUpdateRequest**](SystemGenerateRootUpdateRequest.md) |  | 

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


## PostSysInit

> PostSysInit(ctx).SystemInitRequest(systemInitRequest).Execute()

Initialize a new Vault.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemInitRequest := *openapiclient.NewSystemInitRequest() // SystemInitRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysInit(context.Background()).SystemInitRequest(systemInitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemInitRequest** | [**SystemInitRequest**](SystemInitRequest.md) |  | 

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


## PostSysInternalCountersConfig

> PostSysInternalCountersConfig(ctx).SystemInternalCountersConfigRequest(systemInternalCountersConfigRequest).Execute()

Enable or disable collection of client count, set retention period, or set default reporting period.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemInternalCountersConfigRequest := *openapiclient.NewSystemInternalCountersConfigRequest() // SystemInternalCountersConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysInternalCountersConfig(context.Background()).SystemInternalCountersConfigRequest(systemInternalCountersConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysInternalCountersConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysInternalCountersConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemInternalCountersConfigRequest** | [**SystemInternalCountersConfigRequest**](SystemInternalCountersConfigRequest.md) |  | 

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


## PostSysLeasesLookup

> PostSysLeasesLookup(ctx).SystemLeasesLookupRequest(systemLeasesLookupRequest).Execute()

Retrieve lease metadata.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemLeasesLookupRequest := *openapiclient.NewSystemLeasesLookupRequest() // SystemLeasesLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesLookup(context.Background()).SystemLeasesLookupRequest(systemLeasesLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemLeasesLookupRequest** | [**SystemLeasesLookupRequest**](SystemLeasesLookupRequest.md) |  | 

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


## PostSysLeasesRenew

> PostSysLeasesRenew(ctx).SystemLeasesRenewRequest(systemLeasesRenewRequest).Execute()

Renews a lease, requesting to extend the lease.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemLeasesRenewRequest := *openapiclient.NewSystemLeasesRenewRequest() // SystemLeasesRenewRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRenew(context.Background()).SystemLeasesRenewRequest(systemLeasesRenewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemLeasesRenewRequest** | [**SystemLeasesRenewRequest**](SystemLeasesRenewRequest.md) |  | 

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


## PostSysLeasesRenewUrlLeaseId

> PostSysLeasesRenewUrlLeaseId(ctx, urlLeaseId).SystemLeasesRenewLeaseRequest(systemLeasesRenewLeaseRequest).Execute()

Renews a lease, requesting to extend the lease.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
    systemLeasesRenewLeaseRequest := *openapiclient.NewSystemLeasesRenewLeaseRequest() // SystemLeasesRenewLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRenewUrlLeaseId(context.Background(), urlLeaseId).SystemLeasesRenewLeaseRequest(systemLeasesRenewLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRenewUrlLeaseId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRenewUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRenewLeaseRequest** | [**SystemLeasesRenewLeaseRequest**](SystemLeasesRenewLeaseRequest.md) |  | 

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


## PostSysLeasesRevoke

> PostSysLeasesRevoke(ctx).SystemLeasesRevokeRequest(systemLeasesRevokeRequest).Execute()

Revokes a lease immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemLeasesRevokeRequest := *openapiclient.NewSystemLeasesRevokeRequest() // SystemLeasesRevokeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRevoke(context.Background()).SystemLeasesRevokeRequest(systemLeasesRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemLeasesRevokeRequest** | [**SystemLeasesRevokeRequest**](SystemLeasesRevokeRequest.md) |  | 

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


## PostSysLeasesRevokeForcePrefix

> PostSysLeasesRevokeForcePrefix(ctx, prefix).Execute()

Revokes all secrets or tokens generated under a given prefix immediately



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRevokeForcePrefix(context.Background(), prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRevokeForcePrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokeForcePrefixRequest struct via the builder pattern


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


## PostSysLeasesRevokePrefixPrefix

> PostSysLeasesRevokePrefixPrefix(ctx, prefix).SystemLeasesRevokePrefixRequest(systemLeasesRevokePrefixRequest).Execute()

Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
    systemLeasesRevokePrefixRequest := *openapiclient.NewSystemLeasesRevokePrefixRequest() // SystemLeasesRevokePrefixRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRevokePrefixPrefix(context.Background(), prefix).SystemLeasesRevokePrefixRequest(systemLeasesRevokePrefixRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRevokePrefixPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokePrefixPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRevokePrefixRequest** | [**SystemLeasesRevokePrefixRequest**](SystemLeasesRevokePrefixRequest.md) |  | 

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


## PostSysLeasesRevokeUrlLeaseId

> PostSysLeasesRevokeUrlLeaseId(ctx, urlLeaseId).SystemLeasesRevokeLeaseRequest(systemLeasesRevokeLeaseRequest).Execute()

Revokes a lease immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
    systemLeasesRevokeLeaseRequest := *openapiclient.NewSystemLeasesRevokeLeaseRequest() // SystemLeasesRevokeLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysLeasesRevokeUrlLeaseId(context.Background(), urlLeaseId).SystemLeasesRevokeLeaseRequest(systemLeasesRevokeLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesRevokeUrlLeaseId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokeUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRevokeLeaseRequest** | [**SystemLeasesRevokeLeaseRequest**](SystemLeasesRevokeLeaseRequest.md) |  | 

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


## PostSysLeasesTidy

> PostSysLeasesTidy(ctx).Execute()

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
    resp, r, err := apiClient.SystemApi.PostSysLeasesTidy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysLeasesTidy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesTidyRequest struct via the builder pattern


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


## PostSysMfaValidate

> PostSysMfaValidate(ctx).SystemMfaValidateRequest(systemMfaValidateRequest).Execute()

Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemMfaValidateRequest := *openapiclient.NewSystemMfaValidateRequest(map[string]interface{}(123), "MfaRequestId_example") // SystemMfaValidateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysMfaValidate(context.Background()).SystemMfaValidateRequest(systemMfaValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysMfaValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysMfaValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemMfaValidateRequest** | [**SystemMfaValidateRequest**](SystemMfaValidateRequest.md) |  | 

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


## PostSysMountsPath

> PostSysMountsPath(ctx, path).SystemMountsRequest(systemMountsRequest).Execute()

Enable a new secrets engine at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Example: \"aws/east\"
    systemMountsRequest := *openapiclient.NewSystemMountsRequest() // SystemMountsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysMountsPath(context.Background(), path).SystemMountsRequest(systemMountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysMountsPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysMountsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemMountsRequest** | [**SystemMountsRequest**](SystemMountsRequest.md) |  | 

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


## PostSysMountsPathTune

> PostSysMountsPathTune(ctx, path).SystemMountsTuneRequest(systemMountsTuneRequest).Execute()

Tune backend configuration parameters for this mount.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | The path to mount to. Example: \"aws/east\"
    systemMountsTuneRequest := *openapiclient.NewSystemMountsTuneRequest() // SystemMountsTuneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysMountsPathTune(context.Background(), path).SystemMountsTuneRequest(systemMountsTuneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysMountsPathTune``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysMountsPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemMountsTuneRequest** | [**SystemMountsTuneRequest**](SystemMountsTuneRequest.md) |  | 

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


## PostSysPluginsCatalogName

> PostSysPluginsCatalogName(ctx, name).SystemPluginsCatalogRequest(systemPluginsCatalogRequest).Execute()

Register a new plugin, or updates an existing one with the supplied name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin
    systemPluginsCatalogRequest := *openapiclient.NewSystemPluginsCatalogRequest() // SystemPluginsCatalogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPluginsCatalogName(context.Background(), name).SystemPluginsCatalogRequest(systemPluginsCatalogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPluginsCatalogName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPluginsCatalogNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPluginsCatalogRequest** | [**SystemPluginsCatalogRequest**](SystemPluginsCatalogRequest.md) |  | 

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


## PostSysPluginsCatalogTypeName

> PostSysPluginsCatalogTypeName(ctx, name, type_).SystemPluginsCatalogRequest(systemPluginsCatalogRequest).Execute()

Register a new plugin, or updates an existing one with the supplied name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the plugin
    type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
    systemPluginsCatalogRequest := *openapiclient.NewSystemPluginsCatalogRequest() // SystemPluginsCatalogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPluginsCatalogTypeName(context.Background(), name, type_).SystemPluginsCatalogRequest(systemPluginsCatalogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPluginsCatalogTypeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPluginsCatalogTypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **systemPluginsCatalogRequest** | [**SystemPluginsCatalogRequest**](SystemPluginsCatalogRequest.md) |  | 

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


## PostSysPluginsReloadBackend

> PostSysPluginsReloadBackend(ctx).SystemPluginsReloadBackendRequest(systemPluginsReloadBackendRequest).Execute()

Reload mounted plugin backends.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemPluginsReloadBackendRequest := *openapiclient.NewSystemPluginsReloadBackendRequest() // SystemPluginsReloadBackendRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPluginsReloadBackend(context.Background()).SystemPluginsReloadBackendRequest(systemPluginsReloadBackendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPluginsReloadBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPluginsReloadBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemPluginsReloadBackendRequest** | [**SystemPluginsReloadBackendRequest**](SystemPluginsReloadBackendRequest.md) |  | 

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


## PostSysPoliciesAclName

> PostSysPoliciesAclName(ctx, name).SystemPoliciesAclRequest(systemPoliciesAclRequest).Execute()

Add a new or update an existing ACL policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"
    systemPoliciesAclRequest := *openapiclient.NewSystemPoliciesAclRequest() // SystemPoliciesAclRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPoliciesAclName(context.Background(), name).SystemPoliciesAclRequest(systemPoliciesAclRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPoliciesAclName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPoliciesAclNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPoliciesAclRequest** | [**SystemPoliciesAclRequest**](SystemPoliciesAclRequest.md) |  | 

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


## PostSysPoliciesPasswordName

> PostSysPoliciesPasswordName(ctx, name).SystemPoliciesPasswordRequest(systemPoliciesPasswordRequest).Execute()

Add a new or update an existing password policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the password policy.
    systemPoliciesPasswordRequest := *openapiclient.NewSystemPoliciesPasswordRequest() // SystemPoliciesPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPoliciesPasswordName(context.Background(), name).SystemPoliciesPasswordRequest(systemPoliciesPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPoliciesPasswordName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPoliciesPasswordNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPoliciesPasswordRequest** | [**SystemPoliciesPasswordRequest**](SystemPoliciesPasswordRequest.md) |  | 

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


## PostSysPolicyName

> PostSysPolicyName(ctx, name).SystemPolicyRequest(systemPolicyRequest).Execute()

Add a new or update an existing policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the policy. Example: \"ops\"
    systemPolicyRequest := *openapiclient.NewSystemPolicyRequest() // SystemPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysPolicyName(context.Background(), name).SystemPolicyRequest(systemPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysPolicyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPolicyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPolicyRequest** | [**SystemPolicyRequest**](SystemPolicyRequest.md) |  | 

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


## PostSysQuotasConfig

> PostSysQuotasConfig(ctx).SystemQuotasConfigRequest(systemQuotasConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemQuotasConfigRequest := *openapiclient.NewSystemQuotasConfigRequest() // SystemQuotasConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysQuotasConfig(context.Background()).SystemQuotasConfigRequest(systemQuotasConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysQuotasConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysQuotasConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemQuotasConfigRequest** | [**SystemQuotasConfigRequest**](SystemQuotasConfigRequest.md) |  | 

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


## PostSysQuotasRateLimitName

> PostSysQuotasRateLimitName(ctx, name).SystemQuotasRateLimitRequest(systemQuotasRateLimitRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the quota rule.
    systemQuotasRateLimitRequest := *openapiclient.NewSystemQuotasRateLimitRequest() // SystemQuotasRateLimitRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysQuotasRateLimitName(context.Background(), name).SystemQuotasRateLimitRequest(systemQuotasRateLimitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysQuotasRateLimitName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysQuotasRateLimitNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemQuotasRateLimitRequest** | [**SystemQuotasRateLimitRequest**](SystemQuotasRateLimitRequest.md) |  | 

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


## PostSysRaw

> PostSysRaw(ctx).SystemRawRequest(systemRawRequest).Execute()

Update the value of the key at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRawRequest := *openapiclient.NewSystemRawRequest() // SystemRawRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRaw(context.Background()).SystemRawRequest(systemRawRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRawRequest** | [**SystemRawRequest**](SystemRawRequest.md) |  | 

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


## PostSysRawPath

> PostSysRawPath(ctx, path).SystemRawRequest(systemRawRequest).Execute()

Update the value of the key at the given path.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | 
    systemRawRequest := *openapiclient.NewSystemRawRequest() // SystemRawRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRawPath(context.Background(), path).SystemRawRequest(systemRawRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRawPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRawPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRawRequest** | [**SystemRawRequest**](SystemRawRequest.md) |  | 

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


## PostSysRekeyInit

> PostSysRekeyInit(ctx).SystemRekeyInitRequest(systemRekeyInitRequest).Execute()

Initializes a new rekey attempt.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRekeyInitRequest := *openapiclient.NewSystemRekeyInitRequest() // SystemRekeyInitRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRekeyInit(context.Background()).SystemRekeyInitRequest(systemRekeyInitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRekeyInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRekeyInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRekeyInitRequest** | [**SystemRekeyInitRequest**](SystemRekeyInitRequest.md) |  | 

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


## PostSysRekeyUpdate

> PostSysRekeyUpdate(ctx).SystemRekeyUpdateRequest(systemRekeyUpdateRequest).Execute()

Enter a single unseal key share to progress the rekey of the Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRekeyUpdateRequest := *openapiclient.NewSystemRekeyUpdateRequest() // SystemRekeyUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRekeyUpdate(context.Background()).SystemRekeyUpdateRequest(systemRekeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRekeyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRekeyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRekeyUpdateRequest** | [**SystemRekeyUpdateRequest**](SystemRekeyUpdateRequest.md) |  | 

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


## PostSysRekeyVerify

> PostSysRekeyVerify(ctx).SystemRekeyVerifyRequest(systemRekeyVerifyRequest).Execute()

Enter a single new key share to progress the rekey verification operation.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRekeyVerifyRequest := *openapiclient.NewSystemRekeyVerifyRequest() // SystemRekeyVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRekeyVerify(context.Background()).SystemRekeyVerifyRequest(systemRekeyVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRekeyVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRekeyVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRekeyVerifyRequest** | [**SystemRekeyVerifyRequest**](SystemRekeyVerifyRequest.md) |  | 

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


## PostSysRemount

> PostSysRemount(ctx).SystemRemountRequest(systemRemountRequest).Execute()

Initiate a mount migration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRemountRequest := *openapiclient.NewSystemRemountRequest() // SystemRemountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRemount(context.Background()).SystemRemountRequest(systemRemountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRemount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRemountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRemountRequest** | [**SystemRemountRequest**](SystemRemountRequest.md) |  | 

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


## PostSysRenew

> PostSysRenew(ctx).SystemRenewRequest(systemRenewRequest).Execute()

Renews a lease, requesting to extend the lease.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRenewRequest := *openapiclient.NewSystemRenewRequest() // SystemRenewRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRenew(context.Background()).SystemRenewRequest(systemRenewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRenewRequest** | [**SystemRenewRequest**](SystemRenewRequest.md) |  | 

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


## PostSysRenewUrlLeaseId

> PostSysRenewUrlLeaseId(ctx, urlLeaseId).SystemRenewLeaseRequest(systemRenewLeaseRequest).Execute()

Renews a lease, requesting to extend the lease.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
    systemRenewLeaseRequest := *openapiclient.NewSystemRenewLeaseRequest() // SystemRenewLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRenewUrlLeaseId(context.Background(), urlLeaseId).SystemRenewLeaseRequest(systemRenewLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRenewUrlLeaseId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRenewUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRenewLeaseRequest** | [**SystemRenewLeaseRequest**](SystemRenewLeaseRequest.md) |  | 

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


## PostSysRevoke

> PostSysRevoke(ctx).SystemRevokeRequest(systemRevokeRequest).Execute()

Revokes a lease immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRevokeRequest := *openapiclient.NewSystemRevokeRequest() // SystemRevokeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRevoke(context.Background()).SystemRevokeRequest(systemRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRevokeRequest** | [**SystemRevokeRequest**](SystemRevokeRequest.md) |  | 

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


## PostSysRevokeForcePrefix

> PostSysRevokeForcePrefix(ctx, prefix).Execute()

Revokes all secrets or tokens generated under a given prefix immediately



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRevokeForcePrefix(context.Background(), prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRevokeForcePrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokeForcePrefixRequest struct via the builder pattern


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


## PostSysRevokePrefixPrefix

> PostSysRevokePrefixPrefix(ctx, prefix).SystemRevokePrefixRequest(systemRevokePrefixRequest).Execute()

Revokes all secrets (via a lease ID prefix) or tokens (via the tokens' path property) generated under a given prefix immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
    systemRevokePrefixRequest := *openapiclient.NewSystemRevokePrefixRequest() // SystemRevokePrefixRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRevokePrefixPrefix(context.Background(), prefix).SystemRevokePrefixRequest(systemRevokePrefixRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRevokePrefixPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokePrefixPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRevokePrefixRequest** | [**SystemRevokePrefixRequest**](SystemRevokePrefixRequest.md) |  | 

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


## PostSysRevokeUrlLeaseId

> PostSysRevokeUrlLeaseId(ctx, urlLeaseId).SystemRevokeLeaseRequest(systemRevokeLeaseRequest).Execute()

Revokes a lease immediately.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
    systemRevokeLeaseRequest := *openapiclient.NewSystemRevokeLeaseRequest() // SystemRevokeLeaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRevokeUrlLeaseId(context.Background(), urlLeaseId).SystemRevokeLeaseRequest(systemRevokeLeaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRevokeUrlLeaseId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokeUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRevokeLeaseRequest** | [**SystemRevokeLeaseRequest**](SystemRevokeLeaseRequest.md) |  | 

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


## PostSysRotate

> PostSysRotate(ctx).Execute()

Rotates the backend encryption key used to persist data.

### Example

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
    resp, r, err := apiClient.SystemApi.PostSysRotate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRotateRequest struct via the builder pattern


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


## PostSysRotateConfig

> PostSysRotateConfig(ctx).SystemRotateConfigRequest(systemRotateConfigRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemRotateConfigRequest := *openapiclient.NewSystemRotateConfigRequest() // SystemRotateConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysRotateConfig(context.Background()).SystemRotateConfigRequest(systemRotateConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysRotateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRotateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemRotateConfigRequest** | [**SystemRotateConfigRequest**](SystemRotateConfigRequest.md) |  | 

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


## PostSysSeal

> PostSysSeal(ctx).Execute()

Seal the Vault.

### Example

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
    resp, r, err := apiClient.SystemApi.PostSysSeal(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysSeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysSealRequest struct via the builder pattern


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


## PostSysStepDown

> PostSysStepDown(ctx).Execute()

Cause the node to give up active status.



### Example

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
    resp, r, err := apiClient.SystemApi.PostSysStepDown(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysStepDown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysStepDownRequest struct via the builder pattern


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


## PostSysToolsHash

> PostSysToolsHash(ctx).SystemToolsHashRequest(systemToolsHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemToolsHashRequest := *openapiclient.NewSystemToolsHashRequest() // SystemToolsHashRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsHash(context.Background()).SystemToolsHashRequest(systemToolsHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemToolsHashRequest** | [**SystemToolsHashRequest**](SystemToolsHashRequest.md) |  | 

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


## PostSysToolsHashUrlalgorithm

> PostSysToolsHashUrlalgorithm(ctx, urlalgorithm).SystemToolsHashRequest(systemToolsHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
    systemToolsHashRequest := *openapiclient.NewSystemToolsHashRequest() // SystemToolsHashRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsHashUrlalgorithm(context.Background(), urlalgorithm).SystemToolsHashRequest(systemToolsHashRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsHashUrlalgorithm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsHashUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsHashRequest** | [**SystemToolsHashRequest**](SystemToolsHashRequest.md) |  | 

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


## PostSysToolsRandom

> PostSysToolsRandom(ctx).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemToolsRandomRequest := *openapiclient.NewSystemToolsRandomRequest() // SystemToolsRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsRandom(context.Background()).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsRandom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

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


## PostSysToolsRandomSource

> PostSysToolsRandomSource(ctx, source).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (default to "platform")
    systemToolsRandomRequest := *openapiclient.NewSystemToolsRandomRequest() // SystemToolsRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsRandomSource(context.Background(), source).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsRandomSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

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


## PostSysToolsRandomSourceUrlbytes

> PostSysToolsRandomSourceUrlbytes(ctx, source, urlbytes).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (default to "platform")
    urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
    systemToolsRandomRequest := *openapiclient.NewSystemToolsRandomRequest() // SystemToolsRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsRandomSourceUrlbytes(context.Background(), source, urlbytes).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsRandomSourceUrlbytes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomSourceUrlbytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

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


## PostSysToolsRandomUrlbytes

> PostSysToolsRandomUrlbytes(ctx, urlbytes).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
    systemToolsRandomRequest := *openapiclient.NewSystemToolsRandomRequest() // SystemToolsRandomRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysToolsRandomUrlbytes(context.Background(), urlbytes).SystemToolsRandomRequest(systemToolsRandomRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysToolsRandomUrlbytes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomUrlbytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

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


## PostSysUnseal

> PostSysUnseal(ctx).SystemUnsealRequest(systemUnsealRequest).Execute()

Unseal the Vault.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemUnsealRequest := *openapiclient.NewSystemUnsealRequest() // SystemUnsealRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysUnseal(context.Background()).SystemUnsealRequest(systemUnsealRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysUnseal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysUnsealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemUnsealRequest** | [**SystemUnsealRequest**](SystemUnsealRequest.md) |  | 

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


## PostSysWrappingLookup

> PostSysWrappingLookup(ctx).SystemWrappingLookupRequest(systemWrappingLookupRequest).Execute()

Look up wrapping properties for the given token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemWrappingLookupRequest := *openapiclient.NewSystemWrappingLookupRequest() // SystemWrappingLookupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysWrappingLookup(context.Background()).SystemWrappingLookupRequest(systemWrappingLookupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysWrappingLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysWrappingLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemWrappingLookupRequest** | [**SystemWrappingLookupRequest**](SystemWrappingLookupRequest.md) |  | 

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


## PostSysWrappingRewrap

> PostSysWrappingRewrap(ctx).SystemWrappingRewrapRequest(systemWrappingRewrapRequest).Execute()

Rotates a response-wrapped token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemWrappingRewrapRequest := *openapiclient.NewSystemWrappingRewrapRequest() // SystemWrappingRewrapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysWrappingRewrap(context.Background()).SystemWrappingRewrapRequest(systemWrappingRewrapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysWrappingRewrap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysWrappingRewrapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemWrappingRewrapRequest** | [**SystemWrappingRewrapRequest**](SystemWrappingRewrapRequest.md) |  | 

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


## PostSysWrappingUnwrap

> PostSysWrappingUnwrap(ctx).SystemWrappingUnwrapRequest(systemWrappingUnwrapRequest).Execute()

Unwraps a response-wrapped token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemWrappingUnwrapRequest := *openapiclient.NewSystemWrappingUnwrapRequest() // SystemWrappingUnwrapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PostSysWrappingUnwrap(context.Background()).SystemWrappingUnwrapRequest(systemWrappingUnwrapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysWrappingUnwrap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSysWrappingUnwrapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemWrappingUnwrapRequest** | [**SystemWrappingUnwrapRequest**](SystemWrappingUnwrapRequest.md) |  | 

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


## PostSysWrappingWrap

> PostSysWrappingWrap(ctx).Execute()

Response-wraps an arbitrary JSON object.

### Example

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
    resp, r, err := apiClient.SystemApi.PostSysWrappingWrap(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSysWrappingWrap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysWrappingWrapRequest struct via the builder pattern


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

