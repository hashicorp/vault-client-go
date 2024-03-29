# GoogleCloudConfigureAuthRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | Google credentials JSON that Vault will use to verify users against GCP APIs. If not specified, will use application default credentials | [optional] 
**CustomEndpoint** | Pointer to **map[string]interface{}** | Specifies overrides for various Google API Service Endpoints used in requests. | [optional] 
**GceAlias** | Pointer to **string** | Indicates what value to use when generating an alias for GCE authentications. | [optional] [default to "role_id"]
**GceMetadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#x27;default&#x27;, includes: instance_creation_timestamp, instance_id, instance_name, project_id, project_number, role, service_account_id, service_account_email, zone. Not editing this field means the &#x27;default&#x27; fields are included. Explicitly setting this field to empty overrides the &#x27;default&#x27; and means no metadata will be included. If not using &#x27;default&#x27;, explicit fields must be sent like: &#x27;field1,field2&#x27;. | [optional] [default to ["default"]]
**GoogleCertsEndpoint** | Pointer to **string** | Deprecated. This field does nothing and be removed in a future release | [optional] 
**IamAlias** | Pointer to **string** | Indicates what value to use when generating an alias for IAM authentications. | [optional] [default to "role_id"]
**IamMetadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#x27;default&#x27;, includes: project_id, role, service_account_id, service_account_email. Not editing this field means the &#x27;default&#x27; fields are included. Explicitly setting this field to empty overrides the &#x27;default&#x27; and means no metadata will be included. If not using &#x27;default&#x27;, explicit fields must be sent like: &#x27;field1,field2&#x27;. | [optional] [default to ["default"]]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


