# System

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateAuditHash**](SystemApi.md#CalculateAuditHash) | **Post** /sys/audit-hash/{path} | The hash of the given string via the given audit backend
[**DeleteAuditDevice**](SystemApi.md#DeleteAuditDevice) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
[**DeleteAuthMethod**](SystemApi.md#DeleteAuthMethod) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
[**DeleteConfigAuditingRequestHeader**](SystemApi.md#DeleteConfigAuditingRequestHeader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
[**DeleteConfigCORS**](SystemApi.md#DeleteConfigCORS) | **Delete** /sys/config/cors | Remove any CORS settings.
[**DeleteConfigUIHeader**](SystemApi.md#DeleteConfigUIHeader) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
[**DeleteGenerateRoot**](SystemApi.md#DeleteGenerateRoot) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
[**DeleteGenerateRootAttempt**](SystemApi.md#DeleteGenerateRootAttempt) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
[**DeleteLogger**](SystemApi.md#DeleteLogger) | **Delete** /sys/loggers/{name} | Revert a single logger to use log level provided in config.
[**DeleteLoggers**](SystemApi.md#DeleteLoggers) | **Delete** /sys/loggers | Revert the all loggers to use log level provided in config.
[**DeleteMount**](SystemApi.md#DeleteMount) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
[**DeletePluginsCatalogByTypeByName**](SystemApi.md#DeletePluginsCatalogByTypeByName) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
[**DeletePoliciesACL**](SystemApi.md#DeletePoliciesACL) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
[**DeletePoliciesPassword**](SystemApi.md#DeletePoliciesPassword) | **Delete** /sys/policies/password/{name} | Delete a password policy.
[**DeletePolicy**](SystemApi.md#DeletePolicy) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
[**DeleteQuotasRateLimit**](SystemApi.md#DeleteQuotasRateLimit) | **Delete** /sys/quotas/rate-limit/{name} | 
[**DeleteRaw**](SystemApi.md#DeleteRaw) | **Delete** /sys/raw | Delete the key with given path.
[**DeleteRawPath**](SystemApi.md#DeleteRawPath) | **Delete** /sys/raw/{path} | Delete the key with given path.
[**DeleteRekeyBackup**](SystemApi.md#DeleteRekeyBackup) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
[**DeleteRekeyInit**](SystemApi.md#DeleteRekeyInit) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
[**DeleteRekeyRecoveryKeyBackup**](SystemApi.md#DeleteRekeyRecoveryKeyBackup) | **Delete** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**DeleteRekeyVerify**](SystemApi.md#DeleteRekeyVerify) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
[**ListConfigUIHeaders**](SystemApi.md#ListConfigUIHeaders) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
[**ListLeasesLookupPrefix**](SystemApi.md#ListLeasesLookupPrefix) | **Get** /sys/leases/lookup/{prefix} | Returns a list of lease ids.
[**ListPluginsCatalogByType**](SystemApi.md#ListPluginsCatalogByType) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
[**ListPoliciesACL**](SystemApi.md#ListPoliciesACL) | **Get** /sys/policies/acl | List the configured access control policies.
[**ListPoliciesPassword**](SystemApi.md#ListPoliciesPassword) | **Get** /sys/policies/password | List the existing password policies.
[**ListQuotasRateLimits**](SystemApi.md#ListQuotasRateLimits) | **Get** /sys/quotas/rate-limit | 
[**ListVersionHistory**](SystemApi.md#ListVersionHistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
[**MFAValidate**](SystemApi.md#MFAValidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
[**Monitor**](SystemApi.md#Monitor) | **Get** /sys/monitor | 
[**PprofRead**](SystemApi.md#PprofRead) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
[**PprofReadAllocs**](SystemApi.md#PprofReadAllocs) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
[**PprofReadBlock**](SystemApi.md#PprofReadBlock) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
[**PprofReadCmdline**](SystemApi.md#PprofReadCmdline) | **Get** /sys/pprof/cmdline | Returns the running program&#x27;s command line.
[**PprofReadGoroutine**](SystemApi.md#PprofReadGoroutine) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
[**PprofReadHeap**](SystemApi.md#PprofReadHeap) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
[**PprofReadMutex**](SystemApi.md#PprofReadMutex) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
[**PprofReadProfile**](SystemApi.md#PprofReadProfile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
[**PprofReadSymbol**](SystemApi.md#PprofReadSymbol) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
[**PprofReadThreadcreate**](SystemApi.md#PprofReadThreadcreate) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
[**PprofReadTrace**](SystemApi.md#PprofReadTrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
[**ReadAuditDevices**](SystemApi.md#ReadAuditDevices) | **Get** /sys/audit | List the enabled audit devices.
[**ReadAuthMethod**](SystemApi.md#ReadAuthMethod) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
[**ReadAuthMethodTune**](SystemApi.md#ReadAuthMethodTune) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#x27;s configuration.
[**ReadAuthMethods**](SystemApi.md#ReadAuthMethods) | **Get** /sys/auth | List the currently enabled credential backends.
[**ReadConfigAuditingRequestHeader**](SystemApi.md#ReadConfigAuditingRequestHeader) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
[**ReadConfigAuditingRequestHeaders**](SystemApi.md#ReadConfigAuditingRequestHeaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
[**ReadConfigCORS**](SystemApi.md#ReadConfigCORS) | **Get** /sys/config/cors | Return the current CORS settings.
[**ReadConfigStateSanitized**](SystemApi.md#ReadConfigStateSanitized) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
[**ReadConfigUIHeader**](SystemApi.md#ReadConfigUIHeader) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#x27;s configuration
[**ReadGenerateRoot**](SystemApi.md#ReadGenerateRoot) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
[**ReadGenerateRootAttempt**](SystemApi.md#ReadGenerateRootAttempt) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
[**ReadHAStatus**](SystemApi.md#ReadHAStatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
[**ReadHealth**](SystemApi.md#ReadHealth) | **Get** /sys/health | Returns the health status of Vault.
[**ReadHostInfo**](SystemApi.md#ReadHostInfo) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
[**ReadInFlightRequests**](SystemApi.md#ReadInFlightRequests) | **Get** /sys/in-flight-req | reports in-flight requests
[**ReadInit**](SystemApi.md#ReadInit) | **Get** /sys/init | Returns the initialization status of Vault.
[**ReadInternalCountersActivity**](SystemApi.md#ReadInternalCountersActivity) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
[**ReadInternalCountersActivityExport**](SystemApi.md#ReadInternalCountersActivityExport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
[**ReadInternalCountersActivityMonthly**](SystemApi.md#ReadInternalCountersActivityMonthly) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
[**ReadInternalCountersConfig**](SystemApi.md#ReadInternalCountersConfig) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
[**ReadInternalCountersEntities**](SystemApi.md#ReadInternalCountersEntities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
[**ReadInternalCountersRequests**](SystemApi.md#ReadInternalCountersRequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
[**ReadInternalCountersTokens**](SystemApi.md#ReadInternalCountersTokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
[**ReadInternalInspectRouter**](SystemApi.md#ReadInternalInspectRouter) | **Get** /sys/internal/inspect/router/{tag} | Expose the route entry and mount entry tables present in the router
[**ReadInternalSpecsOpenAPI**](SystemApi.md#ReadInternalSpecsOpenAPI) | **Get** /sys/internal/specs/openapi | Generate an OpenAPI 3 document of all mounted paths.
[**ReadInternalUIFeatureFlags**](SystemApi.md#ReadInternalUIFeatureFlags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
[**ReadInternalUIMount**](SystemApi.md#ReadInternalUIMount) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
[**ReadInternalUIMounts**](SystemApi.md#ReadInternalUIMounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
[**ReadInternalUINamespaces**](SystemApi.md#ReadInternalUINamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
[**ReadInternalUIResultantACL**](SystemApi.md#ReadInternalUIResultantACL) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
[**ReadKeyStatus**](SystemApi.md#ReadKeyStatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
[**ReadLeader**](SystemApi.md#ReadLeader) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
[**ReadLeases**](SystemApi.md#ReadLeases) | **Get** /sys/leases | List leases associated with this Vault cluster
[**ReadLeasesCount**](SystemApi.md#ReadLeasesCount) | **Get** /sys/leases/count | Count of leases associated with this Vault cluster
[**ReadLogger**](SystemApi.md#ReadLogger) | **Get** /sys/loggers/{name} | Read the log level for a single logger.
[**ReadLoggers**](SystemApi.md#ReadLoggers) | **Get** /sys/loggers | Read the log level for all existing loggers.
[**ReadMetrics**](SystemApi.md#ReadMetrics) | **Get** /sys/metrics | Export the metrics aggregated for telemetry purpose.
[**ReadMount**](SystemApi.md#ReadMount) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
[**ReadMounts**](SystemApi.md#ReadMounts) | **Get** /sys/mounts | List the currently mounted backends.
[**ReadMountsConfig**](SystemApi.md#ReadMountsConfig) | **Get** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**ReadPluginsCatalog**](SystemApi.md#ReadPluginsCatalog) | **Get** /sys/plugins/catalog | Lists all the plugins known to Vault
[**ReadPluginsCatalogByTypeByName**](SystemApi.md#ReadPluginsCatalogByTypeByName) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
[**ReadPolicies**](SystemApi.md#ReadPolicies) | **Get** /sys/policy | List the configured access control policies.
[**ReadPoliciesACL**](SystemApi.md#ReadPoliciesACL) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
[**ReadPoliciesPassword**](SystemApi.md#ReadPoliciesPassword) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
[**ReadPoliciesPasswordGenerate**](SystemApi.md#ReadPoliciesPasswordGenerate) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
[**ReadPolicy**](SystemApi.md#ReadPolicy) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
[**ReadQuotasConfig**](SystemApi.md#ReadQuotasConfig) | **Get** /sys/quotas/config | 
[**ReadQuotasRateLimit**](SystemApi.md#ReadQuotasRateLimit) | **Get** /sys/quotas/rate-limit/{name} | 
[**ReadRaw**](SystemApi.md#ReadRaw) | **Get** /sys/raw | Read the value of the key at the given path.
[**ReadRawPath**](SystemApi.md#ReadRawPath) | **Get** /sys/raw/{path} | Read the value of the key at the given path.
[**ReadRekeyBackup**](SystemApi.md#ReadRekeyBackup) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
[**ReadRekeyInit**](SystemApi.md#ReadRekeyInit) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
[**ReadRekeyRecoveryKeyBackup**](SystemApi.md#ReadRekeyRecoveryKeyBackup) | **Get** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
[**ReadRekeyVerify**](SystemApi.md#ReadRekeyVerify) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
[**ReadRemountStatus**](SystemApi.md#ReadRemountStatus) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
[**ReadReplicationStatus**](SystemApi.md#ReadReplicationStatus) | **Get** /sys/replication/status | 
[**ReadRotateConfig**](SystemApi.md#ReadRotateConfig) | **Get** /sys/rotate/config | 
[**ReadSealStatus**](SystemApi.md#ReadSealStatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
[**Remount**](SystemApi.md#Remount) | **Post** /sys/remount | Initiate a mount migration
[**Renew**](SystemApi.md#Renew) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
[**RenewFor**](SystemApi.md#RenewFor) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**Revoke**](SystemApi.md#Revoke) | **Post** /sys/revoke | Revokes a lease immediately.
[**RevokeForce**](SystemApi.md#RevokeForce) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**RevokeLease**](SystemApi.md#RevokeLease) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
[**RevokePrefix**](SystemApi.md#RevokePrefix) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**Rotate**](SystemApi.md#Rotate) | **Post** /sys/rotate | Rotates the backend encryption key used to persist data.
[**Seal**](SystemApi.md#Seal) | **Post** /sys/seal | Seal the Vault.
[**StepDownLeader**](SystemApi.md#StepDownLeader) | **Post** /sys/step-down | Cause the node to give up active status.
[**SysDeletePluginsCatalogName**](SystemApi.md#SysDeletePluginsCatalogName) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
[**SysListLeasesLookup**](SystemApi.md#SysListLeasesLookup) | **Get** /sys/leases/lookup/ | Returns a list of lease ids.
[**SysReadPluginsCatalogName**](SystemApi.md#SysReadPluginsCatalogName) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
[**SysWriteLockedusersMountAccessorUnlockAliasIdentifier**](SystemApi.md#SysWriteLockedusersMountAccessorUnlockAliasIdentifier) | **Post** /sys/lockedusers/{mount_accessor}/unlock/{alias_identifier} | Unlocks the user with given mount_accessor and alias_identifier
[**SysWritePluginsCatalogName**](SystemApi.md#SysWritePluginsCatalogName) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**SysWriteToolsRandomUrlbytes**](SystemApi.md#SysWriteToolsRandomUrlbytes) | **Post** /sys/tools/random/{urlbytes} | Generate random bytes
[**ToolsGenerateRandom**](SystemApi.md#ToolsGenerateRandom) | **Post** /sys/tools/random | Generate random bytes
[**ToolsGenerateRandomSource**](SystemApi.md#ToolsGenerateRandomSource) | **Post** /sys/tools/random/{source} | Generate random bytes
[**ToolsGenerateRandomSourceBytes**](SystemApi.md#ToolsGenerateRandomSourceBytes) | **Post** /sys/tools/random/{source}/{urlbytes} | Generate random bytes
[**ToolsHash**](SystemApi.md#ToolsHash) | **Post** /sys/tools/hash | Generate a hash sum for input data
[**ToolsHashWith**](SystemApi.md#ToolsHashWith) | **Post** /sys/tools/hash/{urlalgorithm} | Generate a hash sum for input data
[**Unseal**](SystemApi.md#Unseal) | **Post** /sys/unseal | Unseal the Vault.
[**WrappingReadLookup**](SystemApi.md#WrappingReadLookup) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#x27;s token.
[**WrappingRewrap**](SystemApi.md#WrappingRewrap) | **Post** /sys/wrapping/rewrap | Rotates a response-wrapped token.
[**WrappingUnwrap**](SystemApi.md#WrappingUnwrap) | **Post** /sys/wrapping/unwrap | Unwraps a response-wrapped token.
[**WrappingWrap**](SystemApi.md#WrappingWrap) | **Post** /sys/wrapping/wrap | Response-wraps an arbitrary JSON object.
[**WrappingWriteLookup**](SystemApi.md#WrappingWriteLookup) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
[**WriteAuditDevice**](SystemApi.md#WriteAuditDevice) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
[**WriteAuthMethod**](SystemApi.md#WriteAuthMethod) | **Post** /sys/auth/{path} | Enables a new auth method.
[**WriteAuthMethodTune**](SystemApi.md#WriteAuthMethodTune) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
[**WriteCapabilities**](SystemApi.md#WriteCapabilities) | **Post** /sys/capabilities | Fetches the capabilities of the given token on the given path.
[**WriteCapabilitiesAccessor**](SystemApi.md#WriteCapabilitiesAccessor) | **Post** /sys/capabilities-accessor | Fetches the capabilities of the token associated with the given token, on the given path.
[**WriteCapabilitiesSelf**](SystemApi.md#WriteCapabilitiesSelf) | **Post** /sys/capabilities-self | Fetches the capabilities of the given token on the given path.
[**WriteConfigAuditingRequestHeader**](SystemApi.md#WriteConfigAuditingRequestHeader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
[**WriteConfigCORS**](SystemApi.md#WriteConfigCORS) | **Post** /sys/config/cors | Configure the CORS settings.
[**WriteConfigReloadSubsystem**](SystemApi.md#WriteConfigReloadSubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
[**WriteConfigUIHeader**](SystemApi.md#WriteConfigUIHeader) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
[**WriteGenerateRoot**](SystemApi.md#WriteGenerateRoot) | **Post** /sys/generate-root | Initializes a new root generation attempt.
[**WriteGenerateRootAttempt**](SystemApi.md#WriteGenerateRootAttempt) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
[**WriteGenerateRootUpdate**](SystemApi.md#WriteGenerateRootUpdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
[**WriteInit**](SystemApi.md#WriteInit) | **Post** /sys/init | Initialize a new Vault.
[**WriteInternalCountersConfig**](SystemApi.md#WriteInternalCountersConfig) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
[**WriteLeasesLookup**](SystemApi.md#WriteLeasesLookup) | **Post** /sys/leases/lookup | Retrieve lease metadata.
[**WriteLeasesRenew**](SystemApi.md#WriteLeasesRenew) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
[**WriteLeasesRenew2**](SystemApi.md#WriteLeasesRenew2) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
[**WriteLeasesRevoke**](SystemApi.md#WriteLeasesRevoke) | **Post** /sys/leases/revoke | Revokes a lease immediately.
[**WriteLeasesRevoke2**](SystemApi.md#WriteLeasesRevoke2) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
[**WriteLeasesRevokeForce**](SystemApi.md#WriteLeasesRevokeForce) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
[**WriteLeasesRevokePrefix**](SystemApi.md#WriteLeasesRevokePrefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#x27; path property) generated under a given prefix immediately.
[**WriteLeasesTidy**](SystemApi.md#WriteLeasesTidy) | **Post** /sys/leases/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
[**WriteLogger**](SystemApi.md#WriteLogger) | **Post** /sys/loggers/{name} | Modify the log level of a single logger.
[**WriteLoggers**](SystemApi.md#WriteLoggers) | **Post** /sys/loggers | Modify the log level for all existing loggers.
[**WriteMount**](SystemApi.md#WriteMount) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
[**WriteMountsConfig**](SystemApi.md#WriteMountsConfig) | **Post** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
[**WritePluginsCatalogByTypeByName**](SystemApi.md#WritePluginsCatalogByTypeByName) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
[**WritePluginsReloadBackend**](SystemApi.md#WritePluginsReloadBackend) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
[**WritePoliciesACL**](SystemApi.md#WritePoliciesACL) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
[**WritePoliciesPassword**](SystemApi.md#WritePoliciesPassword) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
[**WritePolicy**](SystemApi.md#WritePolicy) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
[**WriteQuotasConfig**](SystemApi.md#WriteQuotasConfig) | **Post** /sys/quotas/config | 
[**WriteQuotasRateLimit**](SystemApi.md#WriteQuotasRateLimit) | **Post** /sys/quotas/rate-limit/{name} | 
[**WriteRaw**](SystemApi.md#WriteRaw) | **Post** /sys/raw | Update the value of the key at the given path.
[**WriteRawPath**](SystemApi.md#WriteRawPath) | **Post** /sys/raw/{path} | Update the value of the key at the given path.
[**WriteRekeyInit**](SystemApi.md#WriteRekeyInit) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
[**WriteRekeyUpdate**](SystemApi.md#WriteRekeyUpdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
[**WriteRekeyVerify**](SystemApi.md#WriteRekeyVerify) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
[**WriteRotateConfig**](SystemApi.md#WriteRotateConfig) | **Post** /sys/rotate/config | 





## CalculateAuditHash

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
		vault.WithAddress("http://127.0.0.1:8200"),
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



### Example

```go
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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



### Example

```go
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



### Example

```go
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



## PprofRead

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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadPolicies(
		context.Background(),
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



### Example

```go
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



### Example

```go
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
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.System.ReadRaw(
		context.Background(),
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
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	path := "path_example" // string | 
	resp, err := client.System.ReadRawPath(
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

 **list** | **string** | Return a list if &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## ReadRekeyBackup

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
		vault.WithAddress("http://127.0.0.1:8200"),
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



### Example

```go
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



### Example

```go
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## Renew

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
		vault.WithAddress("http://127.0.0.1:8200"),
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



## SysDeletePluginsCatalogName

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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## WrappingReadLookup

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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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



### Example

```go
package main

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



### Example

```go
package main

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
		vault.WithAddress("http://127.0.0.1:8200"),
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
		vault.WithAddress("http://127.0.0.1:8200"),
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



### Example

```go
package main

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



