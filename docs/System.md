# System

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSysAuditPath**](System.md#DeleteSysAuditPath) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
[**DeleteSysAuthPath**](System.md#DeleteSysAuthPath) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
[**DeleteSysConfigAuditingRequestHeadersHeader**](System.md#DeleteSysConfigAuditingRequestHeadersHeader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
[**DeleteSysConfigCors**](System.md#DeleteSysConfigCors) | **Delete** /sys/config/cors | Remove any CORS settings.
[**DeleteSysConfigUiHeadersHeader**](System.md#DeleteSysConfigUiHeadersHeader) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
[**DeleteSysGenerateRoot**](System.md#DeleteSysGenerateRoot) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
[**DeleteSysGenerateRootAttempt**](System.md#DeleteSysGenerateRootAttempt) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
[**DeleteSysMountsPath**](System.md#DeleteSysMountsPath) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
[**DeleteSysPluginsCatalogName**](System.md#DeleteSysPluginsCatalogName) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
[**DeleteSysPluginsCatalogTypeName**](System.md#DeleteSysPluginsCatalogTypeName) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
[**DeleteSysPoliciesAclName**](System.md#DeleteSysPoliciesAclName) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
[**DeleteSysPoliciesPasswordName**](System.md#DeleteSysPoliciesPasswordName) | **Delete** /sys/policies/password/{name} | Delete a password policy.
[**DeleteSysPolicyName**](System.md#DeleteSysPolicyName) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
[**DeleteSysQuotasRateLimitName**](System.md#DeleteSysQuotasRateLimitName) | **Delete** /sys/quotas/rate-limit/{name} | 
[**DeleteSysRaw**](System.md#DeleteSysRaw) | **Delete** /sys/raw | Delete the key with given path.
[**DeleteSysRawPath**](System.md#DeleteSysRawPath) | **Delete** /sys/raw/{path} | Delete the key with given path.
[**DeleteSysRekeyBackup**](System.md#DeleteSysRekeyBackup) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
[**DeleteSysRekeyInit**](System.md#DeleteSysRekeyInit) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
[**DeleteSysRekeyRecoveryKeyBackup**](System.md#DeleteSysRekeyRecoveryKeyBackup) | **Delete** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**DeleteSysRekeyVerify**](System.md#DeleteSysRekeyVerify) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
[**GetSysAudit**](System.md#GetSysAudit) | **Get** /sys/audit | List the enabled audit devices.
[**GetSysAuth**](System.md#GetSysAuth) | **Get** /sys/auth | List the currently enabled credential backends.
[**GetSysAuthPath**](System.md#GetSysAuthPath) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
[**GetSysAuthPathTune**](System.md#GetSysAuthPathTune) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#39;s configuration.
[**GetSysConfigAuditingRequestHeaders**](System.md#GetSysConfigAuditingRequestHeaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
[**GetSysConfigAuditingRequestHeadersHeader**](System.md#GetSysConfigAuditingRequestHeadersHeader) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
[**GetSysConfigCors**](System.md#GetSysConfigCors) | **Get** /sys/config/cors | Return the current CORS settings.
[**GetSysConfigStateSanitized**](System.md#GetSysConfigStateSanitized) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
[**GetSysConfigUiHeaders**](System.md#GetSysConfigUiHeaders) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
[**GetSysConfigUiHeadersHeader**](System.md#GetSysConfigUiHeadersHeader) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#39;s configuration
[**GetSysGenerateRoot**](System.md#GetSysGenerateRoot) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
[**GetSysGenerateRootAttempt**](System.md#GetSysGenerateRootAttempt) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
[**GetSysHaStatus**](System.md#GetSysHaStatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
[**GetSysHealth**](System.md#GetSysHealth) | **Get** /sys/health | Returns the health status of Vault.
[**GetSysHostInfo**](System.md#GetSysHostInfo) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
[**GetSysInFlightReq**](System.md#GetSysInFlightReq) | **Get** /sys/in-flight-req | reports in-flight requests
[**GetSysInit**](System.md#GetSysInit) | **Get** /sys/init | Returns the initialization status of Vault.
[**GetSysInternalCountersActivity**](System.md#GetSysInternalCountersActivity) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
[**GetSysInternalCountersActivityExport**](System.md#GetSysInternalCountersActivityExport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
[**GetSysInternalCountersActivityMonthly**](System.md#GetSysInternalCountersActivityMonthly) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
[**GetSysInternalCountersConfig**](System.md#GetSysInternalCountersConfig) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
[**GetSysInternalCountersEntities**](System.md#GetSysInternalCountersEntities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
[**GetSysInternalCountersRequests**](System.md#GetSysInternalCountersRequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
[**GetSysInternalCountersTokens**](System.md#GetSysInternalCountersTokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
[**GetSysInternalSpecsOpenapi**](System.md#GetSysInternalSpecsOpenapi) | **Get** /sys/internal/specs/openapi | Generate an OpenAPI 3 document of all mounted paths.
[**GetSysInternalUiFeatureFlags**](System.md#GetSysInternalUiFeatureFlags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
[**GetSysInternalUiMounts**](System.md#GetSysInternalUiMounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
[**GetSysInternalUiMountsPath**](System.md#GetSysInternalUiMountsPath) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
[**GetSysInternalUiNamespaces**](System.md#GetSysInternalUiNamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
[**GetSysInternalUiResultantAcl**](System.md#GetSysInternalUiResultantAcl) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
[**GetSysKeyStatus**](System.md#GetSysKeyStatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
[**GetSysLeader**](System.md#GetSysLeader) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
[**GetSysLeases**](System.md#GetSysLeases) | **Get** /sys/leases | List leases associated with this Vault cluster
[**GetSysLeasesCount**](System.md#GetSysLeasesCount) | **Get** /sys/leases/count | Count of leases associated with this Vault cluster
[**GetSysLeasesLookup**](System.md#GetSysLeasesLookup) | **Get** /sys/leases/lookup/ | Returns a list of lease ids.
[**GetSysLeasesLookupPrefix**](System.md#GetSysLeasesLookupPrefix) | **Get** /sys/leases/lookup/{prefix} | Returns a list of lease ids.
[**GetSysMetrics**](System.md#GetSysMetrics) | **Get** /sys/metrics | Export the metrics aggregated for telemetry purpose.
[**GetSysMonitor**](System.md#GetSysMonitor) | **Get** /sys/monitor | 
[**GetSysMounts**](System.md#GetSysMounts) | **Get** /sys/mounts | List the currently mounted backends.
[**GetSysMountsPath**](System.md#GetSysMountsPath) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
[**GetSysMountsPathTune**](System.md#GetSysMountsPathTune) | **Get** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**GetSysPluginsCatalog**](System.md#GetSysPluginsCatalog) | **Get** /sys/plugins/catalog | Lists all the plugins known to Vault
[**GetSysPluginsCatalogName**](System.md#GetSysPluginsCatalogName) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
[**GetSysPluginsCatalogType**](System.md#GetSysPluginsCatalogType) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
[**GetSysPluginsCatalogTypeName**](System.md#GetSysPluginsCatalogTypeName) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
[**GetSysPoliciesAcl**](System.md#GetSysPoliciesAcl) | **Get** /sys/policies/acl | List the configured access control policies.
[**GetSysPoliciesAclName**](System.md#GetSysPoliciesAclName) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
[**GetSysPoliciesPassword**](System.md#GetSysPoliciesPassword) | **Get** /sys/policies/password | List the existing password policies.
[**GetSysPoliciesPasswordName**](System.md#GetSysPoliciesPasswordName) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
[**GetSysPoliciesPasswordNameGenerate**](System.md#GetSysPoliciesPasswordNameGenerate) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
[**GetSysPolicy**](System.md#GetSysPolicy) | **Get** /sys/policy | List the configured access control policies.
[**GetSysPolicyName**](System.md#GetSysPolicyName) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
[**GetSysPprof**](System.md#GetSysPprof) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
[**GetSysPprofAllocs**](System.md#GetSysPprofAllocs) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
[**GetSysPprofBlock**](System.md#GetSysPprofBlock) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
[**GetSysPprofCmdline**](System.md#GetSysPprofCmdline) | **Get** /sys/pprof/cmdline | Returns the running program&#39;s command line.
[**GetSysPprofGoroutine**](System.md#GetSysPprofGoroutine) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
[**GetSysPprofHeap**](System.md#GetSysPprofHeap) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
[**GetSysPprofMutex**](System.md#GetSysPprofMutex) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
[**GetSysPprofProfile**](System.md#GetSysPprofProfile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
[**GetSysPprofSymbol**](System.md#GetSysPprofSymbol) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
[**GetSysPprofThreadcreate**](System.md#GetSysPprofThreadcreate) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
[**GetSysPprofTrace**](System.md#GetSysPprofTrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
[**GetSysQuotasConfig**](System.md#GetSysQuotasConfig) | **Get** /sys/quotas/config | 
[**GetSysQuotasRateLimit**](System.md#GetSysQuotasRateLimit) | **Get** /sys/quotas/rate-limit | 
[**GetSysQuotasRateLimitName**](System.md#GetSysQuotasRateLimitName) | **Get** /sys/quotas/rate-limit/{name} | 
[**GetSysRaw**](System.md#GetSysRaw) | **Get** /sys/raw | Read the value of the key at the given path.
[**GetSysRawPath**](System.md#GetSysRawPath) | **Get** /sys/raw/{path} | Read the value of the key at the given path.
[**GetSysRekeyBackup**](System.md#GetSysRekeyBackup) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
[**GetSysRekeyInit**](System.md#GetSysRekeyInit) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
[**GetSysRekeyRecoveryKeyBackup**](System.md#GetSysRekeyRecoveryKeyBackup) | **Get** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**GetSysRekeyVerify**](System.md#GetSysRekeyVerify) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
[**GetSysRemountStatusMigrationId**](System.md#GetSysRemountStatusMigrationId) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
[**GetSysReplicationStatus**](System.md#GetSysReplicationStatus) | **Get** /sys/replication/status | 
[**GetSysRotateConfig**](System.md#GetSysRotateConfig) | **Get** /sys/rotate/config | 
[**GetSysSealStatus**](System.md#GetSysSealStatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
[**GetSysVersionHistory**](System.md#GetSysVersionHistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
[**GetSysWrappingLookup**](System.md#GetSysWrappingLookup) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#39;s token.
[**PostSysAuditHashPath**](System.md#PostSysAuditHashPath) | **Post** /sys/audit-hash/{path} | The hash of the given string via the given audit backend
[**PostSysAuditPath**](System.md#PostSysAuditPath) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
[**PostSysAuthPath**](System.md#PostSysAuthPath) | **Post** /sys/auth/{path} | Enables a new auth method.
[**PostSysAuthPathTune**](System.md#PostSysAuthPathTune) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
[**PostSysCapabilities**](System.md#PostSysCapabilities) | **Post** /sys/capabilities | Fetches the capabilities of the given token on the given path.
[**PostSysCapabilitiesAccessor**](System.md#PostSysCapabilitiesAccessor) | **Post** /sys/capabilities-accessor | Fetches the capabilities of the token associated with the given token, on the given path.
[**PostSysCapabilitiesSelf**](System.md#PostSysCapabilitiesSelf) | **Post** /sys/capabilities-self | Fetches the capabilities of the given token on the given path.
[**PostSysConfigAuditingRequestHeadersHeader**](System.md#PostSysConfigAuditingRequestHeadersHeader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
[**PostSysConfigCors**](System.md#PostSysConfigCors) | **Post** /sys/config/cors | Configure the CORS settings.
[**PostSysConfigReloadSubsystem**](System.md#PostSysConfigReloadSubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
[**PostSysConfigUiHeadersHeader**](System.md#PostSysConfigUiHeadersHeader) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
[**PostSysGenerateRoot**](System.md#PostSysGenerateRoot) | **Post** /sys/generate-root | Initializes a new root generation attempt.
[**PostSysGenerateRootAttempt**](System.md#PostSysGenerateRootAttempt) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
[**PostSysGenerateRootUpdate**](System.md#PostSysGenerateRootUpdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
[**PostSysInit**](System.md#PostSysInit) | **Post** /sys/init | Initialize a new Vault.
[**PostSysInternalCountersConfig**](System.md#PostSysInternalCountersConfig) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
[**PostSysLeasesLookup**](System.md#PostSysLeasesLookup) | **Post** /sys/leases/lookup | Retrieve lease metadata.
[**PostSysLeasesRenew**](System.md#PostSysLeasesRenew) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
[**PostSysLeasesRenewUrlLeaseId**](System.md#PostSysLeasesRenewUrlLeaseId) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**PostSysLeasesRevoke**](System.md#PostSysLeasesRevoke) | **Post** /sys/leases/revoke | Revokes a lease immediately.
[**PostSysLeasesRevokeForcePrefix**](System.md#PostSysLeasesRevokeForcePrefix) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**PostSysLeasesRevokePrefixPrefix**](System.md#PostSysLeasesRevokePrefixPrefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
[**PostSysLeasesRevokeUrlLeaseId**](System.md#PostSysLeasesRevokeUrlLeaseId) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
[**PostSysLeasesTidy**](System.md#PostSysLeasesTidy) | **Post** /sys/leases/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**PostSysMfaValidate**](System.md#PostSysMfaValidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
[**PostSysMountsPath**](System.md#PostSysMountsPath) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
[**PostSysMountsPathTune**](System.md#PostSysMountsPathTune) | **Post** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**PostSysPluginsCatalogName**](System.md#PostSysPluginsCatalogName) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PostSysPluginsCatalogTypeName**](System.md#PostSysPluginsCatalogTypeName) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**PostSysPluginsReloadBackend**](System.md#PostSysPluginsReloadBackend) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
[**PostSysPoliciesAclName**](System.md#PostSysPoliciesAclName) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
[**PostSysPoliciesPasswordName**](System.md#PostSysPoliciesPasswordName) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
[**PostSysPolicyName**](System.md#PostSysPolicyName) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
[**PostSysQuotasConfig**](System.md#PostSysQuotasConfig) | **Post** /sys/quotas/config | 
[**PostSysQuotasRateLimitName**](System.md#PostSysQuotasRateLimitName) | **Post** /sys/quotas/rate-limit/{name} | 
[**PostSysRaw**](System.md#PostSysRaw) | **Post** /sys/raw | Update the value of the key at the given path.
[**PostSysRawPath**](System.md#PostSysRawPath) | **Post** /sys/raw/{path} | Update the value of the key at the given path.
[**PostSysRekeyInit**](System.md#PostSysRekeyInit) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
[**PostSysRekeyUpdate**](System.md#PostSysRekeyUpdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
[**PostSysRekeyVerify**](System.md#PostSysRekeyVerify) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
[**PostSysRemount**](System.md#PostSysRemount) | **Post** /sys/remount | Initiate a mount migration
[**PostSysRenew**](System.md#PostSysRenew) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
[**PostSysRenewUrlLeaseId**](System.md#PostSysRenewUrlLeaseId) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**PostSysRevoke**](System.md#PostSysRevoke) | **Post** /sys/revoke | Revokes a lease immediately.
[**PostSysRevokeForcePrefix**](System.md#PostSysRevokeForcePrefix) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**PostSysRevokePrefixPrefix**](System.md#PostSysRevokePrefixPrefix) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
[**PostSysRevokeUrlLeaseId**](System.md#PostSysRevokeUrlLeaseId) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
[**PostSysRotate**](System.md#PostSysRotate) | **Post** /sys/rotate | Rotates the backend encryption key used to persist data.
[**PostSysRotateConfig**](System.md#PostSysRotateConfig) | **Post** /sys/rotate/config | 
[**PostSysSeal**](System.md#PostSysSeal) | **Post** /sys/seal | Seal the Vault.
[**PostSysStepDown**](System.md#PostSysStepDown) | **Post** /sys/step-down | Cause the node to give up active status.
[**PostSysToolsHash**](System.md#PostSysToolsHash) | **Post** /sys/tools/hash | Generate a hash sum for input data
[**PostSysToolsHashUrlalgorithm**](System.md#PostSysToolsHashUrlalgorithm) | **Post** /sys/tools/hash/{urlalgorithm} | Generate a hash sum for input data
[**PostSysToolsRandom**](System.md#PostSysToolsRandom) | **Post** /sys/tools/random | Generate random bytes
[**PostSysToolsRandomSource**](System.md#PostSysToolsRandomSource) | **Post** /sys/tools/random/{source} | Generate random bytes
[**PostSysToolsRandomSourceUrlbytes**](System.md#PostSysToolsRandomSourceUrlbytes) | **Post** /sys/tools/random/{source}/{urlbytes} | Generate random bytes
[**PostSysToolsRandomUrlbytes**](System.md#PostSysToolsRandomUrlbytes) | **Post** /sys/tools/random/{urlbytes} | Generate random bytes
[**PostSysUnseal**](System.md#PostSysUnseal) | **Post** /sys/unseal | Unseal the Vault.
[**PostSysWrappingLookup**](System.md#PostSysWrappingLookup) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
[**PostSysWrappingRewrap**](System.md#PostSysWrappingRewrap) | **Post** /sys/wrapping/rewrap | Rotates a response-wrapped token.
[**PostSysWrappingUnwrap**](System.md#PostSysWrappingUnwrap) | **Post** /sys/wrapping/unwrap | Unwraps a response-wrapped token.
[**PostSysWrappingWrap**](System.md#PostSysWrappingWrap) | **Post** /sys/wrapping/wrap | Response-wraps an arbitrary JSON object.



## DeleteSysAuditPath

> DeleteSysAuditPath(ctx, path).Execute()

Disable the audit device at the given path.

### Example

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	
	resp, err := client.System.DeleteSysAuditPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysAuditPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Cannot be delimited. Example: \"user\"
	
	resp, err := client.System.DeleteSysAuthPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysAuthPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | 
	
	resp, err := client.System.DeleteSysConfigAuditingRequestHeadersHeader(context.Background(), header)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysConfigCors(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigCorsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | The name of the header.
	
	resp, err := client.System.DeleteSysConfigUiHeadersHeader(context.Background(), header)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysConfigUiHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysGenerateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysGenerateRootRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysGenerateRootAttempt(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysGenerateRootAttemptRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Example: \"aws/east\"
	
	resp, err := client.System.DeleteSysMountsPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysMountsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	
	resp, err := client.System.DeleteSysPluginsCatalogName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPluginsCatalogNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	type_ :=  // string | The type of the plugin, may be auth, secret, or database
	
	resp, err := client.System.DeleteSysPluginsCatalogTypeName(context.Background(), name, type_)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPluginsCatalogTypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	resp, err := client.System.DeleteSysPoliciesAclName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPoliciesAclNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the password policy.
	
	resp, err := client.System.DeleteSysPoliciesPasswordName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPoliciesPasswordNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	resp, err := client.System.DeleteSysPolicyName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysPolicyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the quota rule.
	
	resp, err := client.System.DeleteSysQuotasRateLimitName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysQuotasRateLimitNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysRaw(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRawRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | 
	
	resp, err := client.System.DeleteSysRawPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRawPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysRekeyBackup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyBackupRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysRekeyInit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyInitRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysRekeyRecoveryKeyBackup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyRecoveryKeyBackupRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.DeleteSysRekeyVerify(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSysRekeyVerifyRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysAudit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuditRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysAuth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Cannot be delimited. Example: \"user\"
	
	resp, err := client.System.GetSysAuthPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Tune the configuration parameters for an auth path.
	
	resp, err := client.System.GetSysAuthPathTune(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysAuthPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysConfigAuditingRequestHeaders(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigAuditingRequestHeadersRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | 
	
	resp, err := client.System.GetSysConfigAuditingRequestHeadersHeader(context.Background(), header)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysConfigCors(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigCorsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysConfigStateSanitized(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigStateSanitizedRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysConfigUiHeaders(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | The name of the header.
	
	resp, err := client.System.GetSysConfigUiHeadersHeader(context.Background(), header)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysConfigUiHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysGenerateRoot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysGenerateRootRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysGenerateRootAttempt(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysGenerateRootAttemptRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysHaStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHaStatusRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysHealth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHealthRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysHostInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysHostInfoRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInFlightReq(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInFlightReqRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInitRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersActivity(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersActivityExport(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityExportRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersActivityMonthly(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersActivityMonthlyRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersConfigRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersEntities(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersEntitiesRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersRequests(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersRequestsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalCountersTokens(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalCountersTokensRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalSpecsOpenapi(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalSpecsOpenapiRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalUiFeatureFlags(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiFeatureFlagsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalUiMounts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiMountsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path of the mount.
	
	resp, err := client.System.GetSysInternalUiMountsPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path of the mount. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiMountsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalUiNamespaces(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiNamespacesRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysInternalUiResultantAcl(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysInternalUiResultantAclRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysKeyStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysKeyStatusRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysLeader(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeaderRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysLeases(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysLeasesCount(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesCountRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysLeasesLookup(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	prefix :=  // string | The path to list leases under. Example: \"aws/creds/deploy\"
	
	list := NewstringWithDefaults()
	
	resp, err := client.System.GetSysLeasesLookupPrefix(context.Background(), prefix, list)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to list leases under. Example: \&quot;aws/creds/deploy\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysLeasesLookupPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	format := NewstringWithDefaults()
	
	resp, err := client.System.GetSysMetrics(context.Background(), format)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	logFormat := NewstringWithDefaults()
	
	logLevel := NewstringWithDefaults()
	
	resp, err := client.System.GetSysMonitor(context.Background(), logFormat, logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysMounts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Example: \"aws/east\"
	
	resp, err := client.System.GetSysMountsPath(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Example: \"aws/east\"
	
	resp, err := client.System.GetSysMountsPathTune(context.Background(), path)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysMountsPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPluginsCatalog(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	
	resp, err := client.System.GetSysPluginsCatalogName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	type_ :=  // string | The type of the plugin, may be auth, secret, or database
	
	list := NewstringWithDefaults()
	
	resp, err := client.System.GetSysPluginsCatalogType(context.Background(), type_, list)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	type_ :=  // string | The type of the plugin, may be auth, secret, or database
	
	resp, err := client.System.GetSysPluginsCatalogTypeName(context.Background(), name, type_)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 
**type_** | **string** | The type of the plugin, may be auth, secret, or database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPluginsCatalogTypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysPoliciesAcl(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	resp, err := client.System.GetSysPoliciesAclName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesAclNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysPoliciesPassword(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the password policy.
	
	resp, err := client.System.GetSysPoliciesPasswordName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesPasswordNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the password policy.
	
	resp, err := client.System.GetSysPoliciesPasswordNameGenerate(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPoliciesPasswordNameGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysPolicy(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	resp, err := client.System.GetSysPolicyName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPolicyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprof(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofAllocs(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofAllocsRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofBlock(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofBlockRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofCmdline(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofCmdlineRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofGoroutine(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofGoroutineRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofHeap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofHeapRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofMutex(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofMutexRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofProfile(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofProfileRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofSymbol(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofSymbolRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofThreadcreate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofThreadcreateRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysPprofTrace(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysPprofTraceRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysQuotasConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysQuotasConfigRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysQuotasRateLimit(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the quota rule.
	
	resp, err := client.System.GetSysQuotasRateLimitName(context.Background(), name)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysQuotasRateLimitNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysRaw(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | 
	
	list := NewstringWithDefaults()
	
	resp, err := client.System.GetSysRawPath(context.Background(), path, list)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRawPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysRekeyBackup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyBackupRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysRekeyInit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyInitRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysRekeyRecoveryKeyBackup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyRecoveryKeyBackupRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysRekeyVerify(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRekeyVerifyRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	migrationId :=  // string | The ID of the migration operation
	
	resp, err := client.System.GetSysRemountStatusMigrationId(context.Background(), migrationId)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **string** | The ID of the migration operation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRemountStatusMigrationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysReplicationStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysReplicationStatusRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysRotateConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysRotateConfigRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysSealStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysSealStatusRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
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
	
	resp, err := client.System.GetSysVersionHistory(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.GetSysWrappingLookup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSysWrappingLookupRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	
	systemAuditHashRequest := NewSystemAuditHashRequestWithDefaults()
	
	resp, err := client.System.PostSysAuditHashPath(context.Background(), path, systemAuditHashRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuditHashPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuditHashRequest** | [**SystemAuditHashRequest**](SystemAuditHashRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	
	systemAuditRequest := NewSystemAuditRequestWithDefaults()
	
	resp, err := client.System.PostSysAuditPath(context.Background(), path, systemAuditRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The name of the backend. Cannot be delimited. Example: \&quot;mysql\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuditPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuditRequest** | [**SystemAuditRequest**](SystemAuditRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Cannot be delimited. Example: \"user\"
	
	systemAuthRequest := NewSystemAuthRequestWithDefaults()
	
	resp, err := client.System.PostSysAuthPath(context.Background(), path, systemAuthRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Cannot be delimited. Example: \&quot;user\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuthPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuthRequest** | [**SystemAuthRequest**](SystemAuthRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | Tune the configuration parameters for an auth path.
	
	systemAuthTuneRequest := NewSystemAuthTuneRequestWithDefaults()
	
	resp, err := client.System.PostSysAuthPathTune(context.Background(), path, systemAuthTuneRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Tune the configuration parameters for an auth path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysAuthPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAuthTuneRequest** | [**SystemAuthTuneRequest**](SystemAuthTuneRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemCapabilitiesRequest := NewSystemCapabilitiesRequestWithDefaults()
	
	resp, err := client.System.PostSysCapabilities(context.Background(), systemCapabilitiesRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemCapabilitiesAccessorRequest := NewSystemCapabilitiesAccessorRequestWithDefaults()
	
	resp, err := client.System.PostSysCapabilitiesAccessor(context.Background(), systemCapabilitiesAccessorRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemCapabilitiesSelfRequest := NewSystemCapabilitiesSelfRequestWithDefaults()
	
	resp, err := client.System.PostSysCapabilitiesSelf(context.Background(), systemCapabilitiesSelfRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | 
	
	systemConfigAuditingRequestHeadersRequest := NewSystemConfigAuditingRequestHeadersRequestWithDefaults()
	
	resp, err := client.System.PostSysConfigAuditingRequestHeadersHeader(context.Background(), header, systemConfigAuditingRequestHeadersRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigAuditingRequestHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemConfigAuditingRequestHeadersRequest** | [**SystemConfigAuditingRequestHeadersRequest**](SystemConfigAuditingRequestHeadersRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemConfigCorsRequest := NewSystemConfigCorsRequestWithDefaults()
	
	resp, err := client.System.PostSysConfigCors(context.Background(), systemConfigCorsRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	subsystem :=  // string | 
	
	resp, err := client.System.PostSysConfigReloadSubsystem(context.Background(), subsystem)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsystem** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigReloadSubsystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	header :=  // string | The name of the header.
	
	systemConfigUiHeadersRequest := NewSystemConfigUiHeadersRequestWithDefaults()
	
	resp, err := client.System.PostSysConfigUiHeadersHeader(context.Background(), header, systemConfigUiHeadersRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**header** | **string** | The name of the header. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysConfigUiHeadersHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemConfigUiHeadersRequest** | [**SystemConfigUiHeadersRequest**](SystemConfigUiHeadersRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemGenerateRootRequest := NewSystemGenerateRootRequestWithDefaults()
	
	resp, err := client.System.PostSysGenerateRoot(context.Background(), systemGenerateRootRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemGenerateRootAttemptRequest := NewSystemGenerateRootAttemptRequestWithDefaults()
	
	resp, err := client.System.PostSysGenerateRootAttempt(context.Background(), systemGenerateRootAttemptRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemGenerateRootUpdateRequest := NewSystemGenerateRootUpdateRequestWithDefaults()
	
	resp, err := client.System.PostSysGenerateRootUpdate(context.Background(), systemGenerateRootUpdateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemInitRequest := NewSystemInitRequestWithDefaults()
	
	resp, err := client.System.PostSysInit(context.Background(), systemInitRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemInternalCountersConfigRequest := NewSystemInternalCountersConfigRequestWithDefaults()
	
	resp, err := client.System.PostSysInternalCountersConfig(context.Background(), systemInternalCountersConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemLeasesLookupRequest := NewSystemLeasesLookupRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesLookup(context.Background(), systemLeasesLookupRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemLeasesRenewRequest := NewSystemLeasesRenewRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesRenew(context.Background(), systemLeasesRenewRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlLeaseId :=  // string | The lease identifier to renew. This is included with a lease.
	
	systemLeasesRenewLeaseRequest := NewSystemLeasesRenewLeaseRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesRenewUrlLeaseId(context.Background(), urlLeaseId, systemLeasesRenewLeaseRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRenewUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRenewLeaseRequest** | [**SystemLeasesRenewLeaseRequest**](SystemLeasesRenewLeaseRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemLeasesRevokeRequest := NewSystemLeasesRevokeRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesRevoke(context.Background(), systemLeasesRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	prefix :=  // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	
	resp, err := client.System.PostSysLeasesRevokeForcePrefix(context.Background(), prefix)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokeForcePrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	prefix :=  // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	
	systemLeasesRevokePrefixRequest := NewSystemLeasesRevokePrefixRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesRevokePrefixPrefix(context.Background(), prefix, systemLeasesRevokePrefixRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokePrefixPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRevokePrefixRequest** | [**SystemLeasesRevokePrefixRequest**](SystemLeasesRevokePrefixRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlLeaseId :=  // string | The lease identifier to renew. This is included with a lease.
	
	systemLeasesRevokeLeaseRequest := NewSystemLeasesRevokeLeaseRequestWithDefaults()
	
	resp, err := client.System.PostSysLeasesRevokeUrlLeaseId(context.Background(), urlLeaseId, systemLeasesRevokeLeaseRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesRevokeUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemLeasesRevokeLeaseRequest** | [**SystemLeasesRevokeLeaseRequest**](SystemLeasesRevokeLeaseRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.PostSysLeasesTidy(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysLeasesTidyRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemMfaValidateRequest := NewSystemMfaValidateRequestWithDefaults()
	
	resp, err := client.System.PostSysMfaValidate(context.Background(), systemMfaValidateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Example: \"aws/east\"
	
	systemMountsRequest := NewSystemMountsRequestWithDefaults()
	
	resp, err := client.System.PostSysMountsPath(context.Background(), path, systemMountsRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysMountsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemMountsRequest** | [**SystemMountsRequest**](SystemMountsRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | The path to mount to. Example: \"aws/east\"
	
	systemMountsTuneRequest := NewSystemMountsTuneRequestWithDefaults()
	
	resp, err := client.System.PostSysMountsPathTune(context.Background(), path, systemMountsTuneRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to mount to. Example: \&quot;aws/east\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysMountsPathTuneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemMountsTuneRequest** | [**SystemMountsTuneRequest**](SystemMountsTuneRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	
	systemPluginsCatalogRequest := NewSystemPluginsCatalogRequestWithDefaults()
	
	resp, err := client.System.PostSysPluginsCatalogName(context.Background(), name, systemPluginsCatalogRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPluginsCatalogNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPluginsCatalogRequest** | [**SystemPluginsCatalogRequest**](SystemPluginsCatalogRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the plugin
	type_ :=  // string | The type of the plugin, may be auth, secret, or database
	
	systemPluginsCatalogRequest := NewSystemPluginsCatalogRequestWithDefaults()
	
	resp, err := client.System.PostSysPluginsCatalogTypeName(context.Background(), name, type_, systemPluginsCatalogRequest)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemPluginsReloadBackendRequest := NewSystemPluginsReloadBackendRequestWithDefaults()
	
	resp, err := client.System.PostSysPluginsReloadBackend(context.Background(), systemPluginsReloadBackendRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	systemPoliciesAclRequest := NewSystemPoliciesAclRequestWithDefaults()
	
	resp, err := client.System.PostSysPoliciesAclName(context.Background(), name, systemPoliciesAclRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPoliciesAclNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPoliciesAclRequest** | [**SystemPoliciesAclRequest**](SystemPoliciesAclRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the password policy.
	
	systemPoliciesPasswordRequest := NewSystemPoliciesPasswordRequestWithDefaults()
	
	resp, err := client.System.PostSysPoliciesPasswordName(context.Background(), name, systemPoliciesPasswordRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the password policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPoliciesPasswordNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPoliciesPasswordRequest** | [**SystemPoliciesPasswordRequest**](SystemPoliciesPasswordRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | The name of the policy. Example: \"ops\"
	
	systemPolicyRequest := NewSystemPolicyRequestWithDefaults()
	
	resp, err := client.System.PostSysPolicyName(context.Background(), name, systemPolicyRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the policy. Example: \&quot;ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysPolicyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemPolicyRequest** | [**SystemPolicyRequest**](SystemPolicyRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemQuotasConfigRequest := NewSystemQuotasConfigRequestWithDefaults()
	
	resp, err := client.System.PostSysQuotasConfig(context.Background(), systemQuotasConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	name :=  // string | Name of the quota rule.
	
	systemQuotasRateLimitRequest := NewSystemQuotasRateLimitRequestWithDefaults()
	
	resp, err := client.System.PostSysQuotasRateLimitName(context.Background(), name, systemQuotasRateLimitRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the quota rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysQuotasRateLimitNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemQuotasRateLimitRequest** | [**SystemQuotasRateLimitRequest**](SystemQuotasRateLimitRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRawRequest := NewSystemRawRequestWithDefaults()
	
	resp, err := client.System.PostSysRaw(context.Background(), systemRawRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	path :=  // string | 
	
	systemRawRequest := NewSystemRawRequestWithDefaults()
	
	resp, err := client.System.PostSysRawPath(context.Background(), path, systemRawRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRawPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRawRequest** | [**SystemRawRequest**](SystemRawRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRekeyInitRequest := NewSystemRekeyInitRequestWithDefaults()
	
	resp, err := client.System.PostSysRekeyInit(context.Background(), systemRekeyInitRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRekeyUpdateRequest := NewSystemRekeyUpdateRequestWithDefaults()
	
	resp, err := client.System.PostSysRekeyUpdate(context.Background(), systemRekeyUpdateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRekeyVerifyRequest := NewSystemRekeyVerifyRequestWithDefaults()
	
	resp, err := client.System.PostSysRekeyVerify(context.Background(), systemRekeyVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRemountRequest := NewSystemRemountRequestWithDefaults()
	
	resp, err := client.System.PostSysRemount(context.Background(), systemRemountRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRenewRequest := NewSystemRenewRequestWithDefaults()
	
	resp, err := client.System.PostSysRenew(context.Background(), systemRenewRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlLeaseId :=  // string | The lease identifier to renew. This is included with a lease.
	
	systemRenewLeaseRequest := NewSystemRenewLeaseRequestWithDefaults()
	
	resp, err := client.System.PostSysRenewUrlLeaseId(context.Background(), urlLeaseId, systemRenewLeaseRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRenewUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRenewLeaseRequest** | [**SystemRenewLeaseRequest**](SystemRenewLeaseRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRevokeRequest := NewSystemRevokeRequestWithDefaults()
	
	resp, err := client.System.PostSysRevoke(context.Background(), systemRevokeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	prefix :=  // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	
	resp, err := client.System.PostSysRevokeForcePrefix(context.Background(), prefix)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokeForcePrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	prefix :=  // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	
	systemRevokePrefixRequest := NewSystemRevokePrefixRequestWithDefaults()
	
	resp, err := client.System.PostSysRevokePrefixPrefix(context.Background(), prefix, systemRevokePrefixRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | The path to revoke keys under. Example: \&quot;prod/aws/ops\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokePrefixPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRevokePrefixRequest** | [**SystemRevokePrefixRequest**](SystemRevokePrefixRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlLeaseId :=  // string | The lease identifier to renew. This is included with a lease.
	
	systemRevokeLeaseRequest := NewSystemRevokeLeaseRequestWithDefaults()
	
	resp, err := client.System.PostSysRevokeUrlLeaseId(context.Background(), urlLeaseId, systemRevokeLeaseRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlLeaseId** | **string** | The lease identifier to renew. This is included with a lease. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRevokeUrlLeaseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemRevokeLeaseRequest** | [**SystemRevokeLeaseRequest**](SystemRevokeLeaseRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.PostSysRotate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysRotateRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemRotateConfigRequest := NewSystemRotateConfigRequestWithDefaults()
	
	resp, err := client.System.PostSysRotateConfig(context.Background(), systemRotateConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.PostSysSeal(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysSealRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.PostSysStepDown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysStepDownRequest struct via the builder pattern


### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemToolsHashRequest := NewSystemToolsHashRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsHash(context.Background(), systemToolsHashRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlalgorithm :=  // string | Algorithm to use (POST URL parameter)
	
	systemToolsHashRequest := NewSystemToolsHashRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsHashUrlalgorithm(context.Background(), urlalgorithm, systemToolsHashRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlalgorithm** | **string** | Algorithm to use (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsHashUrlalgorithmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsHashRequest** | [**SystemToolsHashRequest**](SystemToolsHashRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemToolsRandomRequest := NewSystemToolsRandomRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsRandom(context.Background(), systemToolsRandomRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	source :=  // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	
	systemToolsRandomRequest := NewSystemToolsRandomRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsRandomSource(context.Background(), source, systemToolsRandomRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [default to &quot;platform&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	source :=  // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes :=  // string | The number of bytes to generate (POST URL parameter)
	
	systemToolsRandomRequest := NewSystemToolsRandomRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsRandomSourceUrlbytes(context.Background(), source, urlbytes, systemToolsRandomRequest)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	urlbytes :=  // string | The number of bytes to generate (POST URL parameter)
	
	systemToolsRandomRequest := NewSystemToolsRandomRequestWithDefaults()
	
	resp, err := client.System.PostSysToolsRandomUrlbytes(context.Background(), urlbytes, systemToolsRandomRequest)
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlbytes** | **string** | The number of bytes to generate (POST URL parameter) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysToolsRandomUrlbytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemToolsRandomRequest** | [**SystemToolsRandomRequest**](SystemToolsRandomRequest.md) |  | 

### Return type

 (empty response body)

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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemUnsealRequest := NewSystemUnsealRequestWithDefaults()
	
	resp, err := client.System.PostSysUnseal(context.Background(), systemUnsealRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemWrappingLookupRequest := NewSystemWrappingLookupRequestWithDefaults()
	
	resp, err := client.System.PostSysWrappingLookup(context.Background(), systemWrappingLookupRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemWrappingRewrapRequest := NewSystemWrappingRewrapRequestWithDefaults()
	
	resp, err := client.System.PostSysWrappingRewrap(context.Background(), systemWrappingRewrapRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	systemWrappingUnwrapRequest := NewSystemWrappingUnwrapRequestWithDefaults()
	
	resp, err := client.System.PostSysWrappingUnwrap(context.Background(), systemWrappingUnwrapRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	"io"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken("my-token")

	
	resp, err := client.System.PostSysWrappingWrap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSysWrappingWrapRequest struct via the builder pattern


### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

