# PkiConfigClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Canonical URI to this mount on this performance replication cluster&#39;s external address. This is for resolving AIA URLs and providing the {{cluster_path}} template parameter but might be used for other purposes in the future. This should only point back to this particular PR replica and should not ever point to another PR cluster. It may point to any node in the PR replica, including standby nodes, and need not always point to the active node. For example: https://pr1.vault.example.com:8200/v1/pki | [optional] 

## Methods

### NewPkiConfigClusterRequest

`func NewPkiConfigClusterRequest() *PkiConfigClusterRequest`

NewPkiConfigClusterRequest instantiates a new PkiConfigClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigClusterRequestWithDefaults

`func NewPkiConfigClusterRequestWithDefaults() *PkiConfigClusterRequest`

NewPkiConfigClusterRequestWithDefaults instantiates a new PkiConfigClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PkiConfigClusterRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PkiConfigClusterRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PkiConfigClusterRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PkiConfigClusterRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


