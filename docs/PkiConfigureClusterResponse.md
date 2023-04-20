# PkiConfigureClusterResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiaPath** | Pointer to **string** | Optional URI to this mount&#x27;s AIA distribution point; may refer to an external non-Vault responder. This is for resolving AIA URLs and providing the {{cluster_aia_path}} template parameter and will not be used for other purposes. As such, unlike path above, this could safely be an insecure transit mechanism (like HTTP without TLS). For example: http://cdn.example.com/pr1/pki | [optional] 
**Path** | Pointer to **string** | Canonical URI to this mount on this performance replication cluster&#x27;s external address. This is for resolving AIA URLs and providing the {{cluster_path}} template parameter but might be used for other purposes in the future. This should only point back to this particular PR replica and should not ever point to another PR cluster. It may point to any node in the PR replica, including standby nodes, and need not always point to the active node. For example: https://pr1.vault.example.com:8200/v1/pki | [optional] 



## Methods


### NewPkiConfigureClusterResponse

`func NewPkiConfigureClusterResponse() *PkiConfigureClusterResponse`

NewPkiConfigureClusterResponse instantiates a new PkiConfigureClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureClusterResponseWithDefaults

`func NewPkiConfigureClusterResponseWithDefaults() *PkiConfigureClusterResponse`

NewPkiConfigureClusterResponseWithDefaults instantiates a new PkiConfigureClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAiaPath

`func (o *PkiConfigureClusterResponse) GetAiaPath() string`

GetAiaPath returns the AiaPath field if non-nil, zero value otherwise.

### GetAiaPathOk

`func (o *PkiConfigureClusterResponse) GetAiaPathOk() (*string, bool)`

GetAiaPathOk returns a tuple with the AiaPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaPath

`func (o *PkiConfigureClusterResponse) SetAiaPath(v string)`

SetAiaPath sets AiaPath field to given value.


### HasAiaPath

`func (o *PkiConfigureClusterResponse) HasAiaPath() bool`

HasAiaPath returns a boolean if a field has been set.




### GetPath

`func (o *PkiConfigureClusterResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PkiConfigureClusterResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PkiConfigureClusterResponse) SetPath(v string)`

SetPath sets Path field to given value.


### HasPath

`func (o *PkiConfigureClusterResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


