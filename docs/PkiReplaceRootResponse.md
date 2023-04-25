# PkiReplaceRootResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** | Reference (name or identifier) to the default issuer. | [optional] 
**DefaultFollowsLatestIssuer** | Pointer to **bool** | Whether the default issuer should automatically follow the latest generated or imported issuer. Defaults to false. | [optional] 



## Methods


### NewPkiReplaceRootResponse

`func NewPkiReplaceRootResponse() *PkiReplaceRootResponse`

NewPkiReplaceRootResponse instantiates a new PkiReplaceRootResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReplaceRootResponseWithDefaults

`func NewPkiReplaceRootResponseWithDefaults() *PkiReplaceRootResponse`

NewPkiReplaceRootResponseWithDefaults instantiates a new PkiReplaceRootResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDefault

`func (o *PkiReplaceRootResponse) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PkiReplaceRootResponse) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PkiReplaceRootResponse) SetDefault(v string)`

SetDefault sets Default field to given value.


### HasDefault

`func (o *PkiReplaceRootResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.




### GetDefaultFollowsLatestIssuer

`func (o *PkiReplaceRootResponse) GetDefaultFollowsLatestIssuer() bool`

GetDefaultFollowsLatestIssuer returns the DefaultFollowsLatestIssuer field if non-nil, zero value otherwise.

### GetDefaultFollowsLatestIssuerOk

`func (o *PkiReplaceRootResponse) GetDefaultFollowsLatestIssuerOk() (*bool, bool)`

GetDefaultFollowsLatestIssuerOk returns a tuple with the DefaultFollowsLatestIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFollowsLatestIssuer

`func (o *PkiReplaceRootResponse) SetDefaultFollowsLatestIssuer(v bool)`

SetDefaultFollowsLatestIssuer sets DefaultFollowsLatestIssuer field to given value.


### HasDefaultFollowsLatestIssuer

`func (o *PkiReplaceRootResponse) HasDefaultFollowsLatestIssuer() bool`

HasDefaultFollowsLatestIssuer returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


