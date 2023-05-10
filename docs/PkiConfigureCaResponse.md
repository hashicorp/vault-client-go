# PkiConfigureCaResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportedIssuers** | Pointer to **[]string** | Net-new issuers imported as a part of this request | [optional] 
**ImportedKeys** | Pointer to **[]string** | Net-new keys imported as a part of this request | [optional] 
**Mapping** | Pointer to **map[string]interface{}** | A mapping of issuer_id to key_id for all issuers included in this request | [optional] 



## Methods


### NewPkiConfigureCaResponse

`func NewPkiConfigureCaResponse() *PkiConfigureCaResponse`

NewPkiConfigureCaResponse instantiates a new PkiConfigureCaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureCaResponseWithDefaults

`func NewPkiConfigureCaResponseWithDefaults() *PkiConfigureCaResponse`

NewPkiConfigureCaResponseWithDefaults instantiates a new PkiConfigureCaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetImportedIssuers

`func (o *PkiConfigureCaResponse) GetImportedIssuers() []string`

GetImportedIssuers returns the ImportedIssuers field if non-nil, zero value otherwise.

### GetImportedIssuersOk

`func (o *PkiConfigureCaResponse) GetImportedIssuersOk() (*[]string, bool)`

GetImportedIssuersOk returns a tuple with the ImportedIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedIssuers

`func (o *PkiConfigureCaResponse) SetImportedIssuers(v []string)`

SetImportedIssuers sets ImportedIssuers field to given value.


### HasImportedIssuers

`func (o *PkiConfigureCaResponse) HasImportedIssuers() bool`

HasImportedIssuers returns a boolean if a field has been set.




### GetImportedKeys

`func (o *PkiConfigureCaResponse) GetImportedKeys() []string`

GetImportedKeys returns the ImportedKeys field if non-nil, zero value otherwise.

### GetImportedKeysOk

`func (o *PkiConfigureCaResponse) GetImportedKeysOk() (*[]string, bool)`

GetImportedKeysOk returns a tuple with the ImportedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedKeys

`func (o *PkiConfigureCaResponse) SetImportedKeys(v []string)`

SetImportedKeys sets ImportedKeys field to given value.


### HasImportedKeys

`func (o *PkiConfigureCaResponse) HasImportedKeys() bool`

HasImportedKeys returns a boolean if a field has been set.




### GetMapping

`func (o *PkiConfigureCaResponse) GetMapping() map[string]interface{}`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PkiConfigureCaResponse) GetMappingOk() (*map[string]interface{}, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PkiConfigureCaResponse) SetMapping(v map[string]interface{})`

SetMapping sets Mapping field to given value.


### HasMapping

`func (o *PkiConfigureCaResponse) HasMapping() bool`

HasMapping returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


