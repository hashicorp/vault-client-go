# PkiIssuersImportCertResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingIssuers** | Pointer to **[]string** | Existing issuers specified as part of the import bundle of this request | [optional] 
**ExistingKeys** | Pointer to **[]string** | Existing keys specified as part of the import bundle of this request | [optional] 
**ImportedIssuers** | Pointer to **[]string** | Net-new issuers imported as a part of this request | [optional] 
**ImportedKeys** | Pointer to **[]string** | Net-new keys imported as a part of this request | [optional] 
**Mapping** | Pointer to **map[string]interface{}** | A mapping of issuer_id to key_id for all issuers included in this request | [optional] 



## Methods


### NewPkiIssuersImportCertResponse

`func NewPkiIssuersImportCertResponse() *PkiIssuersImportCertResponse`

NewPkiIssuersImportCertResponse instantiates a new PkiIssuersImportCertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuersImportCertResponseWithDefaults

`func NewPkiIssuersImportCertResponseWithDefaults() *PkiIssuersImportCertResponse`

NewPkiIssuersImportCertResponseWithDefaults instantiates a new PkiIssuersImportCertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetExistingIssuers

`func (o *PkiIssuersImportCertResponse) GetExistingIssuers() []string`

GetExistingIssuers returns the ExistingIssuers field if non-nil, zero value otherwise.

### GetExistingIssuersOk

`func (o *PkiIssuersImportCertResponse) GetExistingIssuersOk() (*[]string, bool)`

GetExistingIssuersOk returns a tuple with the ExistingIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingIssuers

`func (o *PkiIssuersImportCertResponse) SetExistingIssuers(v []string)`

SetExistingIssuers sets ExistingIssuers field to given value.


### HasExistingIssuers

`func (o *PkiIssuersImportCertResponse) HasExistingIssuers() bool`

HasExistingIssuers returns a boolean if a field has been set.




### GetExistingKeys

`func (o *PkiIssuersImportCertResponse) GetExistingKeys() []string`

GetExistingKeys returns the ExistingKeys field if non-nil, zero value otherwise.

### GetExistingKeysOk

`func (o *PkiIssuersImportCertResponse) GetExistingKeysOk() (*[]string, bool)`

GetExistingKeysOk returns a tuple with the ExistingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingKeys

`func (o *PkiIssuersImportCertResponse) SetExistingKeys(v []string)`

SetExistingKeys sets ExistingKeys field to given value.


### HasExistingKeys

`func (o *PkiIssuersImportCertResponse) HasExistingKeys() bool`

HasExistingKeys returns a boolean if a field has been set.




### GetImportedIssuers

`func (o *PkiIssuersImportCertResponse) GetImportedIssuers() []string`

GetImportedIssuers returns the ImportedIssuers field if non-nil, zero value otherwise.

### GetImportedIssuersOk

`func (o *PkiIssuersImportCertResponse) GetImportedIssuersOk() (*[]string, bool)`

GetImportedIssuersOk returns a tuple with the ImportedIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedIssuers

`func (o *PkiIssuersImportCertResponse) SetImportedIssuers(v []string)`

SetImportedIssuers sets ImportedIssuers field to given value.


### HasImportedIssuers

`func (o *PkiIssuersImportCertResponse) HasImportedIssuers() bool`

HasImportedIssuers returns a boolean if a field has been set.




### GetImportedKeys

`func (o *PkiIssuersImportCertResponse) GetImportedKeys() []string`

GetImportedKeys returns the ImportedKeys field if non-nil, zero value otherwise.

### GetImportedKeysOk

`func (o *PkiIssuersImportCertResponse) GetImportedKeysOk() (*[]string, bool)`

GetImportedKeysOk returns a tuple with the ImportedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedKeys

`func (o *PkiIssuersImportCertResponse) SetImportedKeys(v []string)`

SetImportedKeys sets ImportedKeys field to given value.


### HasImportedKeys

`func (o *PkiIssuersImportCertResponse) HasImportedKeys() bool`

HasImportedKeys returns a boolean if a field has been set.




### GetMapping

`func (o *PkiIssuersImportCertResponse) GetMapping() map[string]interface{}`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PkiIssuersImportCertResponse) GetMappingOk() (*map[string]interface{}, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PkiIssuersImportCertResponse) SetMapping(v map[string]interface{})`

SetMapping sets Mapping field to given value.


### HasMapping

`func (o *PkiIssuersImportCertResponse) HasMapping() bool`

HasMapping returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


