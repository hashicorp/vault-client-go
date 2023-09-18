# TokenCreateAgainstRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name to associate with this token | [optional] 
**EntityAlias** | Pointer to **string** | Name of the entity alias to associate with this token | [optional] 
**ExplicitMaxTtl** | Pointer to **string** | Explicit Max TTL of this token | [optional] 
**Id** | Pointer to **string** | Value for the token | [optional] 
**Lease** | Pointer to **string** | Use &#x27;ttl&#x27; instead | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Arbitrary key&#x3D;value metadata to associate with the token | [optional] 
**NoDefaultPolicy** | Pointer to **bool** | Do not include default policy for this token | [optional] 
**NoParent** | Pointer to **bool** | Create the token with no parent | [optional] 
**NumUses** | Pointer to **int32** | Max number of uses for this token | [optional] 
**Period** | Pointer to **string** | Renew period | [optional] 
**Policies** | Pointer to **[]string** | List of policies for the token | [optional] 
**Renewable** | Pointer to **bool** | Allow token to be renewed past its initial TTL up to system/mount maximum TTL | [optional] [default to true]
**Ttl** | Pointer to **string** | Time to live for this token | [optional] 
**Type** | Pointer to **string** | Token type | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


