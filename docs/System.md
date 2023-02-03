# System

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateAuditHash**](System.md#CalculateAuditHash) | **Post** /sys/audit-hash/{path} | The hash of the given string via the given audit backend
[**DeleteAuditDevice**](System.md#DeleteAuditDevice) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
[**DeleteAuthMethod**](System.md#DeleteAuthMethod) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
[**DeleteConfigAuditingRequestHeader**](System.md#DeleteConfigAuditingRequestHeader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
[**DeleteConfigCORS**](System.md#DeleteConfigCORS) | **Delete** /sys/config/cors | Remove any CORS settings.
[**DeleteConfigUIHeader**](System.md#DeleteConfigUIHeader) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
[**DeleteGenerateRoot**](System.md#DeleteGenerateRoot) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
[**DeleteGenerateRootAttempt**](System.md#DeleteGenerateRootAttempt) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
[**DeleteLogger**](System.md#DeleteLogger) | **Delete** /sys/loggers/{name} | Revert a single logger to use log level provided in config.
[**DeleteLoggers**](System.md#DeleteLoggers) | **Delete** /sys/loggers | Revert the all loggers to use log level provided in config.
[**DeleteMount**](System.md#DeleteMount) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
[**DeletePluginsCatalogByTypeByName**](System.md#DeletePluginsCatalogByTypeByName) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
[**DeletePoliciesACL**](System.md#DeletePoliciesACL) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
[**DeletePoliciesPassword**](System.md#DeletePoliciesPassword) | **Delete** /sys/policies/password/{name} | Delete a password policy.
[**DeletePolicy**](System.md#DeletePolicy) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
[**DeleteQuotasRateLimit**](System.md#DeleteQuotasRateLimit) | **Delete** /sys/quotas/rate-limit/{name} | 
[**DeleteRaw**](System.md#DeleteRaw) | **Delete** /sys/raw | Delete the key with given path.
[**DeleteRawPath**](System.md#DeleteRawPath) | **Delete** /sys/raw/{path} | Delete the key with given path.
[**DeleteRekeyBackup**](System.md#DeleteRekeyBackup) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
[**DeleteRekeyInit**](System.md#DeleteRekeyInit) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
[**DeleteRekeyRecoveryKeyBackup**](System.md#DeleteRekeyRecoveryKeyBackup) | **Delete** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**DeleteRekeyVerify**](System.md#DeleteRekeyVerify) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
[**ListConfigUIHeaders**](System.md#ListConfigUIHeaders) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
[**ListLeasesLookupPrefix**](System.md#ListLeasesLookupPrefix) | **Get** /sys/leases/lookup/{prefix} | Returns a list of lease ids.
[**ListPluginsCatalogByType**](System.md#ListPluginsCatalogByType) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
[**ListPoliciesACL**](System.md#ListPoliciesACL) | **Get** /sys/policies/acl | List the configured access control policies.
[**ListPoliciesPassword**](System.md#ListPoliciesPassword) | **Get** /sys/policies/password | List the existing password policies.
[**ListQuotasRateLimits**](System.md#ListQuotasRateLimits) | **Get** /sys/quotas/rate-limit | 
[**ListVersionHistory**](System.md#ListVersionHistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
[**MFAValidate**](System.md#MFAValidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
[**Monitor**](System.md#Monitor) | **Get** /sys/monitor | 
[**PprofRead**](System.md#PprofRead) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
[**PprofReadAllocs**](System.md#PprofReadAllocs) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
[**PprofReadBlock**](System.md#PprofReadBlock) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
[**PprofReadCmdline**](System.md#PprofReadCmdline) | **Get** /sys/pprof/cmdline | Returns the running program&#x27;s command line.
[**PprofReadGoroutine**](System.md#PprofReadGoroutine) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
[**PprofReadHeap**](System.md#PprofReadHeap) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
[**PprofReadMutex**](System.md#PprofReadMutex) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
[**PprofReadProfile**](System.md#PprofReadProfile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
[**PprofReadSymbol**](System.md#PprofReadSymbol) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
[**PprofReadThreadcreate**](System.md#PprofReadThreadcreate) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
[**PprofReadTrace**](System.md#PprofReadTrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
[**ReadAuditDevices**](System.md#ReadAuditDevices) | **Get** /sys/audit | List the enabled audit devices.
[**ReadAuthMethod**](System.md#ReadAuthMethod) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
[**ReadAuthMethodTune**](System.md#ReadAuthMethodTune) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#x27;s configuration.
[**ReadAuthMethods**](System.md#ReadAuthMethods) | **Get** /sys/auth | List the currently enabled credential backends.
[**ReadConfigAuditingRequestHeader**](System.md#ReadConfigAuditingRequestHeader) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
[**ReadConfigAuditingRequestHeaders**](System.md#ReadConfigAuditingRequestHeaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
[**ReadConfigCORS**](System.md#ReadConfigCORS) | **Get** /sys/config/cors | Return the current CORS settings.
[**ReadConfigStateSanitized**](System.md#ReadConfigStateSanitized) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
[**ReadConfigUIHeader**](System.md#ReadConfigUIHeader) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#x27;s configuration
[**ReadGenerateRoot**](System.md#ReadGenerateRoot) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
[**ReadGenerateRootAttempt**](System.md#ReadGenerateRootAttempt) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
[**ReadHAStatus**](System.md#ReadHAStatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
[**ReadHealth**](System.md#ReadHealth) | **Get** /sys/health | Returns the health status of Vault.
[**ReadHostInfo**](System.md#ReadHostInfo) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
[**ReadInFlightRequests**](System.md#ReadInFlightRequests) | **Get** /sys/in-flight-req | reports in-flight requests
[**ReadInit**](System.md#ReadInit) | **Get** /sys/init | Returns the initialization status of Vault.
[**ReadInternalCountersActivity**](System.md#ReadInternalCountersActivity) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
[**ReadInternalCountersActivityExport**](System.md#ReadInternalCountersActivityExport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
[**ReadInternalCountersActivityMonthly**](System.md#ReadInternalCountersActivityMonthly) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
[**ReadInternalCountersConfig**](System.md#ReadInternalCountersConfig) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
[**ReadInternalCountersEntities**](System.md#ReadInternalCountersEntities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
[**ReadInternalCountersRequests**](System.md#ReadInternalCountersRequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
[**ReadInternalCountersTokens**](System.md#ReadInternalCountersTokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
[**ReadInternalInspectRouter**](System.md#ReadInternalInspectRouter) | **Get** /sys/internal/inspect/router/{tag} | Expose the route entry and mount entry tables present in the router
[**ReadInternalSpecsOpenAPI**](System.md#ReadInternalSpecsOpenAPI) | **Get** /sys/internal/specs/openapi | Generate an OpenAPI 3 document of all mounted paths.
[**ReadInternalUIFeatureFlags**](System.md#ReadInternalUIFeatureFlags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
[**ReadInternalUIMount**](System.md#ReadInternalUIMount) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
[**ReadInternalUIMounts**](System.md#ReadInternalUIMounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
[**ReadInternalUINamespaces**](System.md#ReadInternalUINamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
[**ReadInternalUIResultantACL**](System.md#ReadInternalUIResultantACL) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
[**ReadKeyStatus**](System.md#ReadKeyStatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
[**ReadLeader**](System.md#ReadLeader) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
[**ReadLeases**](System.md#ReadLeases) | **Get** /sys/leases | List leases associated with this Vault cluster
[**ReadLeasesCount**](System.md#ReadLeasesCount) | **Get** /sys/leases/count | Count of leases associated with this Vault cluster
[**ReadLogger**](System.md#ReadLogger) | **Get** /sys/loggers/{name} | Read the log level for a single logger.
[**ReadLoggers**](System.md#ReadLoggers) | **Get** /sys/loggers | Read the log level for all existing loggers.
[**ReadMetrics**](System.md#ReadMetrics) | **Get** /sys/metrics | Export the metrics aggregated for telemetry purpose.
[**ReadMount**](System.md#ReadMount) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
[**ReadMounts**](System.md#ReadMounts) | **Get** /sys/mounts | List the currently mounted backends.
[**ReadMountsConfig**](System.md#ReadMountsConfig) | **Get** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**ReadPluginsCatalog**](System.md#ReadPluginsCatalog) | **Get** /sys/plugins/catalog | Lists all the plugins known to Vault
[**ReadPluginsCatalogByTypeByName**](System.md#ReadPluginsCatalogByTypeByName) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
[**ReadPolicies**](System.md#ReadPolicies) | **Get** /sys/policy | List the configured access control policies.
[**ReadPoliciesACL**](System.md#ReadPoliciesACL) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
[**ReadPoliciesPassword**](System.md#ReadPoliciesPassword) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
[**ReadPoliciesPasswordGenerate**](System.md#ReadPoliciesPasswordGenerate) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
[**ReadPolicy**](System.md#ReadPolicy) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
[**ReadQuotasConfig**](System.md#ReadQuotasConfig) | **Get** /sys/quotas/config | 
[**ReadQuotasRateLimit**](System.md#ReadQuotasRateLimit) | **Get** /sys/quotas/rate-limit/{name} | 
[**ReadRaw**](System.md#ReadRaw) | **Get** /sys/raw | Read the value of the key at the given path.
[**ReadRawPath**](System.md#ReadRawPath) | **Get** /sys/raw/{path} | Read the value of the key at the given path.
[**ReadRekeyBackup**](System.md#ReadRekeyBackup) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
[**ReadRekeyInit**](System.md#ReadRekeyInit) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
[**ReadRekeyRecoveryKeyBackup**](System.md#ReadRekeyRecoveryKeyBackup) | **Get** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**ReadRekeyVerify**](System.md#ReadRekeyVerify) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
[**ReadRemountStatus**](System.md#ReadRemountStatus) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
[**ReadReplicationStatus**](System.md#ReadReplicationStatus) | **Get** /sys/replication/status | 
[**ReadRotateConfig**](System.md#ReadRotateConfig) | **Get** /sys/rotate/config | 
[**ReadSealStatus**](System.md#ReadSealStatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
[**Remount**](System.md#Remount) | **Post** /sys/remount | Initiate a mount migration
[**Renew**](System.md#Renew) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
[**RenewFor**](System.md#RenewFor) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**Revoke**](System.md#Revoke) | **Post** /sys/revoke | Revokes a lease immediately.
[**RevokeForce**](System.md#RevokeForce) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**RevokeLease**](System.md#RevokeLease) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
[**RevokePrefix**](System.md#RevokePrefix) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**Rotate**](System.md#Rotate) | **Post** /sys/rotate | Rotates the backend encryption key used to persist data.
[**Seal**](System.md#Seal) | **Post** /sys/seal | Seal the Vault.
[**StepDownLeader**](System.md#StepDownLeader) | **Post** /sys/step-down | Cause the node to give up active status.
[**SysDeletePluginsCatalogName**](System.md#SysDeletePluginsCatalogName) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
[**SysListLeasesLookup**](System.md#SysListLeasesLookup) | **Get** /sys/leases/lookup/ | Returns a list of lease ids.
[**SysReadPluginsCatalogName**](System.md#SysReadPluginsCatalogName) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
[**SysWriteLockedusersMountAccessorUnlockAliasIdentifier**](System.md#SysWriteLockedusersMountAccessorUnlockAliasIdentifier) | **Post** /sys/lockedusers/{mount_accessor}/unlock/{alias_identifier} | Unlocks the user with given mount_accessor and alias_identifier
[**SysWritePluginsCatalogName**](System.md#SysWritePluginsCatalogName) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**SysWriteToolsRandomUrlbytes**](System.md#SysWriteToolsRandomUrlbytes) | **Post** /sys/tools/random/{urlbytes} | Generate random bytes
[**ToolsGenerateRandom**](System.md#ToolsGenerateRandom) | **Post** /sys/tools/random | Generate random bytes
[**ToolsGenerateRandomSource**](System.md#ToolsGenerateRandomSource) | **Post** /sys/tools/random/{source} | Generate random bytes
[**ToolsGenerateRandomSourceBytes**](System.md#ToolsGenerateRandomSourceBytes) | **Post** /sys/tools/random/{source}/{urlbytes} | Generate random bytes
[**ToolsHash**](System.md#ToolsHash) | **Post** /sys/tools/hash | Generate a hash sum for input data
[**ToolsHashWith**](System.md#ToolsHashWith) | **Post** /sys/tools/hash/{urlalgorithm} | Generate a hash sum for input data
[**Unseal**](System.md#Unseal) | **Post** /sys/unseal | Unseal the Vault.
[**WrappingReadLookup**](System.md#WrappingReadLookup) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#x27;s token.
[**WrappingRewrap**](System.md#WrappingRewrap) | **Post** /sys/wrapping/rewrap | Rotates a response-wrapped token.
[**WrappingUnwrap**](System.md#WrappingUnwrap) | **Post** /sys/wrapping/unwrap | Unwraps a response-wrapped token.
[**WrappingWrap**](System.md#WrappingWrap) | **Post** /sys/wrapping/wrap | Response-wraps an arbitrary JSON object.
[**WrappingWriteLookup**](System.md#WrappingWriteLookup) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
[**WriteAuditDevice**](System.md#WriteAuditDevice) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
[**WriteAuthMethod**](System.md#WriteAuthMethod) | **Post** /sys/auth/{path} | Enables a new auth method.
[**WriteAuthMethodTune**](System.md#WriteAuthMethodTune) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
[**WriteCapabilities**](System.md#WriteCapabilities) | **Post** /sys/capabilities | Fetches the capabilities of the given token on the given path.
[**WriteCapabilitiesAccessor**](System.md#WriteCapabilitiesAccessor) | **Post** /sys/capabilities-accessor | Fetches the capabilities of the token associated with the given token, on the given path.
[**WriteCapabilitiesSelf**](System.md#WriteCapabilitiesSelf) | **Post** /sys/capabilities-self | Fetches the capabilities of the given token on the given path.
[**WriteConfigAuditingRequestHeader**](System.md#WriteConfigAuditingRequestHeader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
[**WriteConfigCORS**](System.md#WriteConfigCORS) | **Post** /sys/config/cors | Configure the CORS settings.
[**WriteConfigReloadSubsystem**](System.md#WriteConfigReloadSubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
[**WriteConfigUIHeader**](System.md#WriteConfigUIHeader) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
[**WriteGenerateRoot**](System.md#WriteGenerateRoot) | **Post** /sys/generate-root | Initializes a new root generation attempt.
[**WriteGenerateRootAttempt**](System.md#WriteGenerateRootAttempt) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
[**WriteGenerateRootUpdate**](System.md#WriteGenerateRootUpdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
[**WriteInit**](System.md#WriteInit) | **Post** /sys/init | Initialize a new Vault.
[**WriteInternalCountersConfig**](System.md#WriteInternalCountersConfig) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
[**WriteLeasesLookup**](System.md#WriteLeasesLookup) | **Post** /sys/leases/lookup | Retrieve lease metadata.
[**WriteLeasesRenew**](System.md#WriteLeasesRenew) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
[**WriteLeasesRenew2**](System.md#WriteLeasesRenew2) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**WriteLeasesRevoke**](System.md#WriteLeasesRevoke) | **Post** /sys/leases/revoke | Revokes a lease immediately.
[**WriteLeasesRevoke2**](System.md#WriteLeasesRevoke2) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
[**WriteLeasesRevokeForce**](System.md#WriteLeasesRevokeForce) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**WriteLeasesRevokePrefix**](System.md#WriteLeasesRevokePrefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**WriteLeasesTidy**](System.md#WriteLeasesTidy) | **Post** /sys/leases/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**WriteLogger**](System.md#WriteLogger) | **Post** /sys/loggers/{name} | Modify the log level of a single logger.
[**WriteLoggers**](System.md#WriteLoggers) | **Post** /sys/loggers | Modify the log level for all existing loggers.
[**WriteMount**](System.md#WriteMount) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
[**WriteMountsConfig**](System.md#WriteMountsConfig) | **Post** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**WritePluginsCatalogByTypeByName**](System.md#WritePluginsCatalogByTypeByName) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**WritePluginsReloadBackend**](System.md#WritePluginsReloadBackend) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
[**WritePoliciesACL**](System.md#WritePoliciesACL) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
[**WritePoliciesPassword**](System.md#WritePoliciesPassword) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
[**WritePolicy**](System.md#WritePolicy) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
[**WriteQuotasConfig**](System.md#WriteQuotasConfig) | **Post** /sys/quotas/config | 
[**WriteQuotasRateLimit**](System.md#WriteQuotasRateLimit) | **Post** /sys/quotas/rate-limit/{name} | 
[**WriteRaw**](System.md#WriteRaw) | **Post** /sys/raw | Update the value of the key at the given path.
[**WriteRawPath**](System.md#WriteRawPath) | **Post** /sys/raw/{path} | Update the value of the key at the given path.
[**WriteRekeyInit**](System.md#WriteRekeyInit) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
[**WriteRekeyUpdate**](System.md#WriteRekeyUpdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
[**WriteRekeyVerify**](System.md#WriteRekeyVerify) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
[**WriteRotateConfig**](System.md#WriteRotateConfig) | **Post** /sys/rotate/config | 





## CalculateAuditHash

> CalculateAuditHash(ctx, path).CalculateAuditHashRequest(calculateAuditHashRequest).Execute()

The hash of the given string via the given audit backend

### Example

```go
package main

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

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	request := schema.NewCalculateAuditHashRequestWithDefaults()
	resp, err := client.System.CalculateAuditHash(
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

 **calculateAuditHashRequest** | [**CalculateAuditHashRequest**](CalculateAuditHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DeleteAuditDevice

> DeleteAuditDevice(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	resp, err := client.System.DeleteAuditDevice(
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



## DeleteAuthMethod

> DeleteAuthMethod(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	resp, err := client.System.DeleteAuthMethod(
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



## DeleteConfigAuditingRequestHeader

> DeleteConfigAuditingRequestHeader(ctx, header).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | 
	resp, err := client.System.DeleteConfigAuditingRequestHeader(
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



## DeleteConfigCORS

> DeleteConfigCORS(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteConfigCORS(
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



## DeleteConfigUIHeader

> DeleteConfigUIHeader(ctx, header).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | The name of the header.
	resp, err := client.System.DeleteConfigUIHeader(
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



## DeleteGenerateRoot

> DeleteGenerateRoot(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteGenerateRoot(
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



## DeleteGenerateRootAttempt

> DeleteGenerateRootAttempt(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteGenerateRootAttempt(
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



## DeleteLogger

> DeleteLogger(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the logger to be modified.
	resp, err := client.System.DeleteLogger(
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



## DeleteLoggers

> DeleteLoggers(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteLoggers(
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



## DeleteMount

> DeleteMount(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.DeleteMount(
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



## DeletePluginsCatalogByTypeByName

> DeletePluginsCatalogByTypeByName(ctx, name, type_).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.DeletePluginsCatalogByTypeByName(
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



## DeletePoliciesACL

> DeletePoliciesACL(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.DeletePoliciesACL(
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



## DeletePoliciesPassword

> DeletePoliciesPassword(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.DeletePoliciesPassword(
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



## DeletePolicy

> DeletePolicy(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.DeletePolicy(
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



## DeleteQuotasRateLimit

> DeleteQuotasRateLimit(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the quota rule.
	resp, err := client.System.DeleteQuotasRateLimit(
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



## DeleteRaw

> DeleteRaw(ctx).Execute()

Delete the key with given path.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteRaw(
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



## DeleteRawPath

> DeleteRawPath(ctx, path).Execute()

Delete the key with given path.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | 
	resp, err := client.System.DeleteRawPath(
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
**path** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## DeleteRekeyBackup

> DeleteRekeyBackup(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteRekeyBackup(
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



## DeleteRekeyInit

> DeleteRekeyInit(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteRekeyInit(
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



## DeleteRekeyRecoveryKeyBackup

> DeleteRekeyRecoveryKeyBackup(ctx).Execute()

Allows fetching or deleting the backup of the rotated unseal keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteRekeyRecoveryKeyBackup(
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



## DeleteRekeyVerify

> DeleteRekeyVerify(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.DeleteRekeyVerify(
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



## ListConfigUIHeaders

> ListConfigUIHeaders(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ListConfigUIHeaders(
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



## ListLeasesLookupPrefix

> ListLeasesLookupPrefix(ctx, prefix).List(list).Execute()

Returns a list of lease ids.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "prefix_example" // string | The path to list leases under. Example: \"aws/creds/deploy\"
	resp, err := client.System.ListLeasesLookupPrefix(
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ListPluginsCatalogByType

> ListPluginsCatalogByType(ctx, type_).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.ListPluginsCatalogByType(
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ListPoliciesACL

> ListPoliciesACL(ctx).List(list).Execute()

List the configured access control policies.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ListPoliciesACL(
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



## ListPoliciesPassword

> ListPoliciesPassword(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ListPoliciesPassword(
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



## ListQuotasRateLimits

> ListQuotasRateLimits(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ListQuotasRateLimits(
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



## ListVersionHistory

> ListVersionHistory(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ListVersionHistory(
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



## MFAValidate

> MFAValidate(ctx).MFAValidateRequest(mFAValidateRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewMFAValidateRequestWithDefaults()
	resp, err := client.System.MFAValidate(
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
 **mFAValidateRequest** | [**MFAValidateRequest**](MFAValidateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Monitor

> Monitor(ctx).LogFormat(logFormat).LogLevel(logLevel).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## PprofRead

> PprofRead(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofRead(
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



## PprofReadAllocs

> PprofReadAllocs(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadAllocs(
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



## PprofReadBlock

> PprofReadBlock(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadBlock(
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



## PprofReadCmdline

> PprofReadCmdline(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadCmdline(
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



## PprofReadGoroutine

> PprofReadGoroutine(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadGoroutine(
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



## PprofReadHeap

> PprofReadHeap(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadHeap(
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



## PprofReadMutex

> PprofReadMutex(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadMutex(
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



## PprofReadProfile

> PprofReadProfile(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadProfile(
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



## PprofReadSymbol

> PprofReadSymbol(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadSymbol(
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



## PprofReadThreadcreate

> PprofReadThreadcreate(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadThreadcreate(
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



## PprofReadTrace

> PprofReadTrace(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.PprofReadTrace(
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



## ReadAuditDevices

> ReadAuditDevices(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadAuditDevices(
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



## ReadAuthMethod

> ReadAuthMethod(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	resp, err := client.System.ReadAuthMethod(
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



## ReadAuthMethodTune

> ReadAuthMethodTune(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Tune the configuration parameters for an auth path.
	resp, err := client.System.ReadAuthMethodTune(
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


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadAuthMethods

> ReadAuthMethods(ctx).Execute()

List the currently enabled credential backends.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadAuthMethods(
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



## ReadConfigAuditingRequestHeader

> ReadConfigAuditingRequestHeader(ctx, header).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | 
	resp, err := client.System.ReadConfigAuditingRequestHeader(
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



## ReadConfigAuditingRequestHeaders

> ReadConfigAuditingRequestHeaders(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadConfigAuditingRequestHeaders(
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



## ReadConfigCORS

> ReadConfigCORS(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadConfigCORS(
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



## ReadConfigStateSanitized

> ReadConfigStateSanitized(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadConfigStateSanitized(
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



## ReadConfigUIHeader

> ReadConfigUIHeader(ctx, header).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | The name of the header.
	resp, err := client.System.ReadConfigUIHeader(
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



## ReadGenerateRoot

> ReadGenerateRoot(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadGenerateRoot(
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



## ReadGenerateRootAttempt

> ReadGenerateRootAttempt(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadGenerateRootAttempt(
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



## ReadHAStatus

> ReadHAStatus(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadHAStatus(
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



## ReadHealth

> ReadHealth(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadHealth(
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



## ReadHostInfo

> ReadHostInfo(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadHostInfo(
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



## ReadInFlightRequests

> ReadInFlightRequests(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInFlightRequests(
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



## ReadInit

> ReadInit(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInit(
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



## ReadInternalCountersActivity

> ReadInternalCountersActivity(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersActivity(
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



## ReadInternalCountersActivityExport

> ReadInternalCountersActivityExport(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersActivityExport(
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



## ReadInternalCountersActivityMonthly

> ReadInternalCountersActivityMonthly(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersActivityMonthly(
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



## ReadInternalCountersConfig

> ReadInternalCountersConfig(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersConfig(
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



## ReadInternalCountersEntities

> ReadInternalCountersEntities(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersEntities(
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



## ReadInternalCountersRequests

> ReadInternalCountersRequests(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersRequests(
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



## ReadInternalCountersTokens

> ReadInternalCountersTokens(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalCountersTokens(
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



## ReadInternalInspectRouter

> ReadInternalInspectRouter(ctx, tag).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	tag := "tag_example" // string | Name of subtree being observed
	resp, err := client.System.ReadInternalInspectRouter(
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



## ReadInternalSpecsOpenAPI

> ReadInternalSpecsOpenAPI(ctx).Execute()

Generate an OpenAPI 3 document of all mounted paths.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalSpecsOpenAPI(
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



## ReadInternalUIFeatureFlags

> ReadInternalUIFeatureFlags(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalUIFeatureFlags(
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



## ReadInternalUIMount

> ReadInternalUIMount(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path of the mount.
	resp, err := client.System.ReadInternalUIMount(
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


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadInternalUIMounts

> ReadInternalUIMounts(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalUIMounts(
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



## ReadInternalUINamespaces

> ReadInternalUINamespaces(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalUINamespaces(
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



## ReadInternalUIResultantACL

> ReadInternalUIResultantACL(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadInternalUIResultantACL(
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



## ReadKeyStatus

> ReadKeyStatus(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadKeyStatus(
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



## ReadLeader

> ReadLeader(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadLeader(
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



## ReadLeases

> ReadLeases(ctx).Execute()

List leases associated with this Vault cluster

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadLeases(
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



## ReadLeasesCount

> ReadLeasesCount(ctx).Execute()

Count of leases associated with this Vault cluster

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadLeasesCount(
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



## ReadLogger

> ReadLogger(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the logger to be modified.
	resp, err := client.System.ReadLogger(
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



## ReadLoggers

> ReadLoggers(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadLoggers(
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



## ReadMetrics

> ReadMetrics(ctx).Format(format).Execute()

Export the metrics aggregated for telemetry purpose.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	format := "format_example" // string | Format to export metrics into. Currently accepts only \"prometheus\".
	resp, err := client.System.ReadMetrics(
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



## ReadMount

> ReadMount(ctx, path).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.ReadMount(
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



## ReadMounts

> ReadMounts(ctx).Execute()

List the currently mounted backends.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadMounts(
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



## ReadMountsConfig

> ReadMountsConfig(ctx, path).Execute()

Tune backend configuration parameters for this mount.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	resp, err := client.System.ReadMountsConfig(
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



## ReadPluginsCatalog

> ReadPluginsCatalog(ctx).Execute()

Lists all the plugins known to Vault

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadPluginsCatalog(
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



## ReadPluginsCatalogByTypeByName

> ReadPluginsCatalogByTypeByName(ctx, name, type_).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	resp, err := client.System.ReadPluginsCatalogByTypeByName(
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



## ReadPolicies

> ReadPolicies(ctx).List(list).Execute()

List the configured access control policies.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.System.ReadPolicies(
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadPoliciesACL

> ReadPoliciesACL(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.ReadPoliciesACL(
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



## ReadPoliciesPassword

> ReadPoliciesPassword(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.ReadPoliciesPassword(
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



## ReadPoliciesPasswordGenerate

> ReadPoliciesPasswordGenerate(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the password policy.
	resp, err := client.System.ReadPoliciesPasswordGenerate(
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



## ReadPolicy

> ReadPolicy(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	resp, err := client.System.ReadPolicy(
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



## ReadQuotasConfig

> ReadQuotasConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadQuotasConfig(
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



## ReadQuotasRateLimit

> ReadQuotasRateLimit(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the quota rule.
	resp, err := client.System.ReadQuotasRateLimit(
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



## ReadRaw

> ReadRaw(ctx).List(list).Execute()

Read the value of the key at the given path.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

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
	resp, err := client.System.ReadRaw(
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



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadRawPath

> ReadRawPath(ctx, path).List(list).Execute()

Read the value of the key at the given path.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | 
	list := "list_example" // string | Return a list if `true`
	resp, err := client.System.ReadRawPath(
		context.Background(),
		path,
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
**path** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadRekeyBackup

> ReadRekeyBackup(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRekeyBackup(
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



## ReadRekeyInit

> ReadRekeyInit(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRekeyInit(
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



## ReadRekeyRecoveryKeyBackup

> ReadRekeyRecoveryKeyBackup(ctx).Execute()

Allows fetching or deleting the backup of the rotated unseal keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRekeyRecoveryKeyBackup(
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



## ReadRekeyVerify

> ReadRekeyVerify(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRekeyVerify(
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



## ReadRemountStatus

> ReadRemountStatus(ctx, migrationId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	migrationId := "migrationId_example" // string | The ID of the migration operation
	resp, err := client.System.ReadRemountStatus(
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


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadReplicationStatus

> ReadReplicationStatus(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadReplicationStatus(
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



## ReadRotateConfig

> ReadRotateConfig(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRotateConfig(
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



## ReadSealStatus

> ReadSealStatus(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadSealStatus(
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



## Remount

> Remount(ctx).RemountRequest(remountRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Renew

> Renew(ctx).RenewRequest(renewRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewRenewRequestWithDefaults()
	resp, err := client.System.Renew(
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
 **renewRequest** | [**RenewRequest**](RenewRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RenewFor

> RenewFor(ctx, urlLeaseId).RenewForRequest(renewForRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewRenewForRequestWithDefaults()
	resp, err := client.System.RenewFor(
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

 **renewForRequest** | [**RenewForRequest**](RenewForRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Revoke

> Revoke(ctx).RevokeRequest(revokeRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewRevokeRequestWithDefaults()
	resp, err := client.System.Revoke(
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
 **revokeRequest** | [**RevokeRequest**](RevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RevokeForce

> RevokeForce(ctx, prefix).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	resp, err := client.System.RevokeForce(
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



## RevokeLease

> RevokeLease(ctx, urlLeaseId).RevokeLeaseRequest(revokeLeaseRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewRevokeLeaseRequestWithDefaults()
	resp, err := client.System.RevokeLease(
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

 **revokeLeaseRequest** | [**RevokeLeaseRequest**](RevokeLeaseRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## RevokePrefix

> RevokePrefix(ctx, prefix).RevokePrefixRequest(revokePrefixRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	request := schema.NewRevokePrefixRequestWithDefaults()
	resp, err := client.System.RevokePrefix(
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

 **revokePrefixRequest** | [**RevokePrefixRequest**](RevokePrefixRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Rotate

> Rotate(ctx).Execute()

Rotates the backend encryption key used to persist data.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.Rotate(
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



## Seal

> Seal(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## StepDownLeader

> StepDownLeader(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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



## SysDeletePluginsCatalogName

> SysDeletePluginsCatalogName(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	resp, err := client.System.SysDeletePluginsCatalogName(
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



## SysListLeasesLookup

> SysListLeasesLookup(ctx).List(list).Execute()

Returns a list of lease ids.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.SysListLeasesLookup(
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



## SysReadPluginsCatalogName

> SysReadPluginsCatalogName(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	resp, err := client.System.SysReadPluginsCatalogName(
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



## SysWriteLockedusersMountAccessorUnlockAliasIdentifier

> SysWriteLockedusersMountAccessorUnlockAliasIdentifier(ctx, aliasIdentifier, mountAccessor).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	aliasIdentifier := "aliasIdentifier_example" // string | It is the name of the alias (user). For example, if the alias belongs to userpass backend, the name should be a valid username within userpass auth method. If the alias belongs to an approle auth method, the name should be a valid RoleID
	mountAccessor := "mountAccessor_example" // string | MountAccessor is the identifier of the mount entry to which the user belongs
	resp, err := client.System.SysWriteLockedusersMountAccessorUnlockAliasIdentifier(
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



## SysWritePluginsCatalogName

> SysWritePluginsCatalogName(ctx, name).SysWritePluginsCatalogNameRequest(sysWritePluginsCatalogNameRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	request := schema.NewSysWritePluginsCatalogNameRequestWithDefaults()
	resp, err := client.System.SysWritePluginsCatalogName(
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

 **sysWritePluginsCatalogNameRequest** | [**SysWritePluginsCatalogNameRequest**](SysWritePluginsCatalogNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## SysWriteToolsRandomUrlbytes

> SysWriteToolsRandomUrlbytes(ctx, urlbytes).SysWriteToolsRandomUrlbytesRequest(sysWriteToolsRandomUrlbytesRequest).Execute()

Generate random bytes

### Example

```go
package main

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

	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	request := schema.NewSysWriteToolsRandomUrlbytesRequestWithDefaults()
	resp, err := client.System.SysWriteToolsRandomUrlbytes(
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

 **sysWriteToolsRandomUrlbytesRequest** | [**SysWriteToolsRandomUrlbytesRequest**](SysWriteToolsRandomUrlbytesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ToolsGenerateRandom

> ToolsGenerateRandom(ctx).ToolsGenerateRandomRequest(toolsGenerateRandomRequest).Execute()

Generate random bytes

### Example

```go
package main

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

	request := schema.NewToolsGenerateRandomRequestWithDefaults()
	resp, err := client.System.ToolsGenerateRandom(
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
 **toolsGenerateRandomRequest** | [**ToolsGenerateRandomRequest**](ToolsGenerateRandomRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ToolsGenerateRandomSource

> ToolsGenerateRandomSource(ctx, source).ToolsGenerateRandomSourceRequest(toolsGenerateRandomSourceRequest).Execute()

Generate random bytes

### Example

```go
package main

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

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	request := schema.NewToolsGenerateRandomSourceRequestWithDefaults()
	resp, err := client.System.ToolsGenerateRandomSource(
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

 **toolsGenerateRandomSourceRequest** | [**ToolsGenerateRandomSourceRequest**](ToolsGenerateRandomSourceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ToolsGenerateRandomSourceBytes

> ToolsGenerateRandomSourceBytes(ctx, source, urlbytes).ToolsGenerateRandomSourceBytesRequest(toolsGenerateRandomSourceBytesRequest).Execute()

Generate random bytes

### Example

```go
package main

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

	source := "source_example" // string | Which system to source random data from, ether \"platform\", \"seal\", or \"all\". (defaults to "platform")
	urlbytes := "urlbytes_example" // string | The number of bytes to generate (POST URL parameter)
	request := schema.NewToolsGenerateRandomSourceBytesRequestWithDefaults()
	resp, err := client.System.ToolsGenerateRandomSourceBytes(
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


 **toolsGenerateRandomSourceBytesRequest** | [**ToolsGenerateRandomSourceBytesRequest**](ToolsGenerateRandomSourceBytesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ToolsHash

> ToolsHash(ctx).ToolsHashRequest(toolsHashRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

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

	request := schema.NewToolsHashRequestWithDefaults()
	resp, err := client.System.ToolsHash(
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
 **toolsHashRequest** | [**ToolsHashRequest**](ToolsHashRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ToolsHashWith

> ToolsHashWith(ctx, urlalgorithm).ToolsHashWithRequest(toolsHashWithRequest).Execute()

Generate a hash sum for input data

### Example

```go
package main

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

	urlalgorithm := "urlalgorithm_example" // string | Algorithm to use (POST URL parameter)
	request := schema.NewToolsHashWithRequestWithDefaults()
	resp, err := client.System.ToolsHashWith(
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

 **toolsHashWithRequest** | [**ToolsHashWithRequest**](ToolsHashWithRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Unseal

> Unseal(ctx).UnsealRequest(unsealRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WrappingReadLookup

> WrappingReadLookup(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.WrappingReadLookup(
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



## WrappingRewrap

> WrappingRewrap(ctx).WrappingRewrapRequest(wrappingRewrapRequest).Execute()

Rotates a response-wrapped token.

### Example

```go
package main

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

	request := schema.NewWrappingRewrapRequestWithDefaults()
	resp, err := client.System.WrappingRewrap(
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
 **wrappingRewrapRequest** | [**WrappingRewrapRequest**](WrappingRewrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WrappingUnwrap

> WrappingUnwrap(ctx).WrappingUnwrapRequest(wrappingUnwrapRequest).Execute()

Unwraps a response-wrapped token.

### Example

```go
package main

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

	request := schema.NewWrappingUnwrapRequestWithDefaults()
	resp, err := client.System.WrappingUnwrap(
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
 **wrappingUnwrapRequest** | [**WrappingUnwrapRequest**](WrappingUnwrapRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WrappingWrap

> WrappingWrap(ctx).Execute()

Response-wraps an arbitrary JSON object.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.WrappingWrap(
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



## WrappingWriteLookup

> WrappingWriteLookup(ctx).WrappingWriteLookupRequest(wrappingWriteLookupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWrappingWriteLookupRequestWithDefaults()
	resp, err := client.System.WrappingWriteLookup(
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
 **wrappingWriteLookupRequest** | [**WrappingWriteLookupRequest**](WrappingWriteLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteAuditDevice

> WriteAuditDevice(ctx, path).WriteAuditDeviceRequest(writeAuditDeviceRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The name of the backend. Cannot be delimited. Example: \"mysql\"
	request := schema.NewWriteAuditDeviceRequestWithDefaults()
	resp, err := client.System.WriteAuditDevice(
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

 **writeAuditDeviceRequest** | [**WriteAuditDeviceRequest**](WriteAuditDeviceRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteAuthMethod

> WriteAuthMethod(ctx, path).WriteAuthMethodRequest(writeAuthMethodRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Cannot be delimited. Example: \"user\"
	request := schema.NewWriteAuthMethodRequestWithDefaults()
	resp, err := client.System.WriteAuthMethod(
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

 **writeAuthMethodRequest** | [**WriteAuthMethodRequest**](WriteAuthMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteAuthMethodTune

> WriteAuthMethodTune(ctx, path).WriteAuthMethodTuneRequest(writeAuthMethodTuneRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | Tune the configuration parameters for an auth path.
	request := schema.NewWriteAuthMethodTuneRequestWithDefaults()
	resp, err := client.System.WriteAuthMethodTune(
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

 **writeAuthMethodTuneRequest** | [**WriteAuthMethodTuneRequest**](WriteAuthMethodTuneRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteCapabilities

> WriteCapabilities(ctx).WriteCapabilitiesRequest(writeCapabilitiesRequest).Execute()

Fetches the capabilities of the given token on the given path.

### Example

```go
package main

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

	request := schema.NewWriteCapabilitiesRequestWithDefaults()
	resp, err := client.System.WriteCapabilities(
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
 **writeCapabilitiesRequest** | [**WriteCapabilitiesRequest**](WriteCapabilitiesRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteCapabilitiesAccessor

> WriteCapabilitiesAccessor(ctx).WriteCapabilitiesAccessorRequest(writeCapabilitiesAccessorRequest).Execute()

Fetches the capabilities of the token associated with the given token, on the given path.

### Example

```go
package main

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

	request := schema.NewWriteCapabilitiesAccessorRequestWithDefaults()
	resp, err := client.System.WriteCapabilitiesAccessor(
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
 **writeCapabilitiesAccessorRequest** | [**WriteCapabilitiesAccessorRequest**](WriteCapabilitiesAccessorRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteCapabilitiesSelf

> WriteCapabilitiesSelf(ctx).WriteCapabilitiesSelfRequest(writeCapabilitiesSelfRequest).Execute()

Fetches the capabilities of the given token on the given path.

### Example

```go
package main

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

	request := schema.NewWriteCapabilitiesSelfRequestWithDefaults()
	resp, err := client.System.WriteCapabilitiesSelf(
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
 **writeCapabilitiesSelfRequest** | [**WriteCapabilitiesSelfRequest**](WriteCapabilitiesSelfRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteConfigAuditingRequestHeader

> WriteConfigAuditingRequestHeader(ctx, header).WriteConfigAuditingRequestHeaderRequest(writeConfigAuditingRequestHeaderRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | 
	request := schema.NewWriteConfigAuditingRequestHeaderRequestWithDefaults()
	resp, err := client.System.WriteConfigAuditingRequestHeader(
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

 **writeConfigAuditingRequestHeaderRequest** | [**WriteConfigAuditingRequestHeaderRequest**](WriteConfigAuditingRequestHeaderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteConfigCORS

> WriteConfigCORS(ctx).WriteConfigCORSRequest(writeConfigCORSRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteConfigCORSRequestWithDefaults()
	resp, err := client.System.WriteConfigCORS(
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
 **writeConfigCORSRequest** | [**WriteConfigCORSRequest**](WriteConfigCORSRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteConfigReloadSubsystem

> WriteConfigReloadSubsystem(ctx, subsystem).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	subsystem := "subsystem_example" // string | 
	resp, err := client.System.WriteConfigReloadSubsystem(
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



## WriteConfigUIHeader

> WriteConfigUIHeader(ctx, header).WriteConfigUIHeaderRequest(writeConfigUIHeaderRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	header := "header_example" // string | The name of the header.
	request := schema.NewWriteConfigUIHeaderRequestWithDefaults()
	resp, err := client.System.WriteConfigUIHeader(
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

 **writeConfigUIHeaderRequest** | [**WriteConfigUIHeaderRequest**](WriteConfigUIHeaderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteGenerateRoot

> WriteGenerateRoot(ctx).WriteGenerateRootRequest(writeGenerateRootRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteGenerateRootRequestWithDefaults()
	resp, err := client.System.WriteGenerateRoot(
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
 **writeGenerateRootRequest** | [**WriteGenerateRootRequest**](WriteGenerateRootRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteGenerateRootAttempt

> WriteGenerateRootAttempt(ctx).WriteGenerateRootAttemptRequest(writeGenerateRootAttemptRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteGenerateRootAttemptRequestWithDefaults()
	resp, err := client.System.WriteGenerateRootAttempt(
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
 **writeGenerateRootAttemptRequest** | [**WriteGenerateRootAttemptRequest**](WriteGenerateRootAttemptRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteGenerateRootUpdate

> WriteGenerateRootUpdate(ctx).WriteGenerateRootUpdateRequest(writeGenerateRootUpdateRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteGenerateRootUpdateRequestWithDefaults()
	resp, err := client.System.WriteGenerateRootUpdate(
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
 **writeGenerateRootUpdateRequest** | [**WriteGenerateRootUpdateRequest**](WriteGenerateRootUpdateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteInit

> WriteInit(ctx).WriteInitRequest(writeInitRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteInitRequestWithDefaults()
	resp, err := client.System.WriteInit(
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
 **writeInitRequest** | [**WriteInitRequest**](WriteInitRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteInternalCountersConfig

> WriteInternalCountersConfig(ctx).WriteInternalCountersConfigRequest(writeInternalCountersConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteInternalCountersConfigRequestWithDefaults()
	resp, err := client.System.WriteInternalCountersConfig(
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
 **writeInternalCountersConfigRequest** | [**WriteInternalCountersConfigRequest**](WriteInternalCountersConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesLookup

> WriteLeasesLookup(ctx).WriteLeasesLookupRequest(writeLeasesLookupRequest).Execute()

Retrieve lease metadata.

### Example

```go
package main

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

	request := schema.NewWriteLeasesLookupRequestWithDefaults()
	resp, err := client.System.WriteLeasesLookup(
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
 **writeLeasesLookupRequest** | [**WriteLeasesLookupRequest**](WriteLeasesLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesRenew

> WriteLeasesRenew(ctx).WriteLeasesRenewRequest(writeLeasesRenewRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteLeasesRenewRequestWithDefaults()
	resp, err := client.System.WriteLeasesRenew(
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
 **writeLeasesRenewRequest** | [**WriteLeasesRenewRequest**](WriteLeasesRenewRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesRenew2

> WriteLeasesRenew2(ctx, urlLeaseId).WriteLeasesRenew2Request(writeLeasesRenew2Request).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewWriteLeasesRenew2RequestWithDefaults()
	resp, err := client.System.WriteLeasesRenew2(
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

 **writeLeasesRenew2Request** | [**WriteLeasesRenew2Request**](WriteLeasesRenew2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesRevoke

> WriteLeasesRevoke(ctx).WriteLeasesRevokeRequest(writeLeasesRevokeRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteLeasesRevokeRequestWithDefaults()
	resp, err := client.System.WriteLeasesRevoke(
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
 **writeLeasesRevokeRequest** | [**WriteLeasesRevokeRequest**](WriteLeasesRevokeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesRevoke2

> WriteLeasesRevoke2(ctx, urlLeaseId).WriteLeasesRevoke2Request(writeLeasesRevoke2Request).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	urlLeaseId := "urlLeaseId_example" // string | The lease identifier to renew. This is included with a lease.
	request := schema.NewWriteLeasesRevoke2RequestWithDefaults()
	resp, err := client.System.WriteLeasesRevoke2(
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

 **writeLeasesRevoke2Request** | [**WriteLeasesRevoke2Request**](WriteLeasesRevoke2Request.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesRevokeForce

> WriteLeasesRevokeForce(ctx, prefix).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	resp, err := client.System.WriteLeasesRevokeForce(
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



## WriteLeasesRevokePrefix

> WriteLeasesRevokePrefix(ctx, prefix).WriteLeasesRevokePrefixRequest(writeLeasesRevokePrefixRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prefix := "prefix_example" // string | The path to revoke keys under. Example: \"prod/aws/ops\"
	request := schema.NewWriteLeasesRevokePrefixRequestWithDefaults()
	resp, err := client.System.WriteLeasesRevokePrefix(
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

 **writeLeasesRevokePrefixRequest** | [**WriteLeasesRevokePrefixRequest**](WriteLeasesRevokePrefixRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLeasesTidy

> WriteLeasesTidy(ctx).Execute()

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

	resp, err := client.System.WriteLeasesTidy(
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



## WriteLogger

> WriteLogger(ctx, name).WriteLoggerRequest(writeLoggerRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the logger to be modified.
	request := schema.NewWriteLoggerRequestWithDefaults()
	resp, err := client.System.WriteLogger(
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

 **writeLoggerRequest** | [**WriteLoggerRequest**](WriteLoggerRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteLoggers

> WriteLoggers(ctx).WriteLoggersRequest(writeLoggersRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteLoggersRequestWithDefaults()
	resp, err := client.System.WriteLoggers(
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
 **writeLoggersRequest** | [**WriteLoggersRequest**](WriteLoggersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteMount

> WriteMount(ctx, path).WriteMountRequest(writeMountRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	request := schema.NewWriteMountRequestWithDefaults()
	resp, err := client.System.WriteMount(
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

 **writeMountRequest** | [**WriteMountRequest**](WriteMountRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteMountsConfig

> WriteMountsConfig(ctx, path).WriteMountsConfigRequest(writeMountsConfigRequest).Execute()

Tune backend configuration parameters for this mount.

### Example

```go
package main

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

	path := "path_example" // string | The path to mount to. Example: \"aws/east\"
	request := schema.NewWriteMountsConfigRequestWithDefaults()
	resp, err := client.System.WriteMountsConfig(
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

 **writeMountsConfigRequest** | [**WriteMountsConfigRequest**](WriteMountsConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WritePluginsCatalogByTypeByName

> WritePluginsCatalogByTypeByName(ctx, name, type_).WritePluginsCatalogByTypeByNameRequest(writePluginsCatalogByTypeByNameRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the plugin
	type_ := "type__example" // string | The type of the plugin, may be auth, secret, or database
	request := schema.NewWritePluginsCatalogByTypeByNameRequestWithDefaults()
	resp, err := client.System.WritePluginsCatalogByTypeByName(
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


 **writePluginsCatalogByTypeByNameRequest** | [**WritePluginsCatalogByTypeByNameRequest**](WritePluginsCatalogByTypeByNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WritePluginsReloadBackend

> WritePluginsReloadBackend(ctx).WritePluginsReloadBackendRequest(writePluginsReloadBackendRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWritePluginsReloadBackendRequestWithDefaults()
	resp, err := client.System.WritePluginsReloadBackend(
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
 **writePluginsReloadBackendRequest** | [**WritePluginsReloadBackendRequest**](WritePluginsReloadBackendRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WritePoliciesACL

> WritePoliciesACL(ctx, name).WritePoliciesACLRequest(writePoliciesACLRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	request := schema.NewWritePoliciesACLRequestWithDefaults()
	resp, err := client.System.WritePoliciesACL(
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

 **writePoliciesACLRequest** | [**WritePoliciesACLRequest**](WritePoliciesACLRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WritePoliciesPassword

> WritePoliciesPassword(ctx, name).WritePoliciesPasswordRequest(writePoliciesPasswordRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the password policy.
	request := schema.NewWritePoliciesPasswordRequestWithDefaults()
	resp, err := client.System.WritePoliciesPassword(
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

 **writePoliciesPasswordRequest** | [**WritePoliciesPasswordRequest**](WritePoliciesPasswordRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WritePolicy

> WritePolicy(ctx, name).WritePolicyRequest(writePolicyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | The name of the policy. Example: \"ops\"
	request := schema.NewWritePolicyRequestWithDefaults()
	resp, err := client.System.WritePolicy(
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

 **writePolicyRequest** | [**WritePolicyRequest**](WritePolicyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteQuotasConfig

> WriteQuotasConfig(ctx).WriteQuotasConfigRequest(writeQuotasConfigRequest).Execute()



### Example

```go
package main

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

	request := schema.NewWriteQuotasConfigRequestWithDefaults()
	resp, err := client.System.WriteQuotasConfig(
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
 **writeQuotasConfigRequest** | [**WriteQuotasConfigRequest**](WriteQuotasConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteQuotasRateLimit

> WriteQuotasRateLimit(ctx, name).WriteQuotasRateLimitRequest(writeQuotasRateLimitRequest).Execute()



### Example

```go
package main

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

	name := "name_example" // string | Name of the quota rule.
	request := schema.NewWriteQuotasRateLimitRequestWithDefaults()
	resp, err := client.System.WriteQuotasRateLimit(
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

 **writeQuotasRateLimitRequest** | [**WriteQuotasRateLimitRequest**](WriteQuotasRateLimitRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRaw

> WriteRaw(ctx).WriteRawRequest(writeRawRequest).Execute()

Update the value of the key at the given path.

### Example

```go
package main

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

	request := schema.NewWriteRawRequestWithDefaults()
	resp, err := client.System.WriteRaw(
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
 **writeRawRequest** | [**WriteRawRequest**](WriteRawRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRawPath

> WriteRawPath(ctx, path).WriteRawPathRequest(writeRawPathRequest).Execute()

Update the value of the key at the given path.

### Example

```go
package main

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

	path := "path_example" // string | 
	request := schema.NewWriteRawPathRequestWithDefaults()
	resp, err := client.System.WriteRawPath(
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
**path** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writeRawPathRequest** | [**WriteRawPathRequest**](WriteRawPathRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRekeyInit

> WriteRekeyInit(ctx).WriteRekeyInitRequest(writeRekeyInitRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteRekeyInitRequestWithDefaults()
	resp, err := client.System.WriteRekeyInit(
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
 **writeRekeyInitRequest** | [**WriteRekeyInitRequest**](WriteRekeyInitRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRekeyUpdate

> WriteRekeyUpdate(ctx).WriteRekeyUpdateRequest(writeRekeyUpdateRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteRekeyUpdateRequestWithDefaults()
	resp, err := client.System.WriteRekeyUpdate(
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
 **writeRekeyUpdateRequest** | [**WriteRekeyUpdateRequest**](WriteRekeyUpdateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRekeyVerify

> WriteRekeyVerify(ctx).WriteRekeyVerifyRequest(writeRekeyVerifyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewWriteRekeyVerifyRequestWithDefaults()
	resp, err := client.System.WriteRekeyVerify(
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
 **writeRekeyVerifyRequest** | [**WriteRekeyVerifyRequest**](WriteRekeyVerifyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WriteRotateConfig

> WriteRotateConfig(ctx).WriteRotateConfigRequest(writeRotateConfigRequest).Execute()



### Example

```go
package main

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

	request := schema.NewWriteRotateConfigRequestWithDefaults()
	resp, err := client.System.WriteRotateConfig(
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
 **writeRotateConfigRequest** | [**WriteRotateConfigRequest**](WriteRotateConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



