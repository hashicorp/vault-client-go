# PkiJsonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerName** | Pointer to **string** | Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value &#39;default&#39; | [optional] 
**IssuerRef** | Pointer to **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [optional] [default to "default"]
**LeafNotAfterBehavior** | Pointer to **string** | Behavior of leaf&#39;s NotAfter fields: \&quot;err\&quot; to error if the computed NotAfter date exceeds that of this issuer; \&quot;truncate\&quot; to silently truncate to that of this issuer; or \&quot;permit\&quot; to allow this issuance to succeed (with NotAfter exceeding that of an issuer). Note that not all values will results in certificates that can be validated through the entire validity period. It is suggested to use \&quot;truncate\&quot; for intermediate CAs and \&quot;permit\&quot; only for root CAs. | [optional] [default to "err"]
**ManualChain** | Pointer to **[]string** | Chain of issuer references to use to build this issuer&#39;s computed CAChain field, when non-empty. | [optional] 
**Usage** | Pointer to **[]string** | Comma-separated list (or string slice) of usages for this issuer; valid values are \&quot;read-only\&quot;, \&quot;issuing-certificates\&quot;, and \&quot;crl-signing\&quot;. Multiple values may be specified. Read-only is implicit and always set. | [optional] [default to ["read-only","issuing-certificates","crl-signing"]]

## Methods

### NewPkiJsonRequest

`func NewPkiJsonRequest() *PkiJsonRequest`

NewPkiJsonRequest instantiates a new PkiJsonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiJsonRequestWithDefaults

`func NewPkiJsonRequestWithDefaults() *PkiJsonRequest`

NewPkiJsonRequestWithDefaults instantiates a new PkiJsonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerName

`func (o *PkiJsonRequest) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiJsonRequest) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiJsonRequest) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *PkiJsonRequest) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetIssuerRef

`func (o *PkiJsonRequest) GetIssuerRef() string`

GetIssuerRef returns the IssuerRef field if non-nil, zero value otherwise.

### GetIssuerRefOk

`func (o *PkiJsonRequest) GetIssuerRefOk() (*string, bool)`

GetIssuerRefOk returns a tuple with the IssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRef

`func (o *PkiJsonRequest) SetIssuerRef(v string)`

SetIssuerRef sets IssuerRef field to given value.

### HasIssuerRef

`func (o *PkiJsonRequest) HasIssuerRef() bool`

HasIssuerRef returns a boolean if a field has been set.

### GetLeafNotAfterBehavior

`func (o *PkiJsonRequest) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiJsonRequest) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiJsonRequest) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.

### HasLeafNotAfterBehavior

`func (o *PkiJsonRequest) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.

### GetManualChain

`func (o *PkiJsonRequest) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiJsonRequest) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiJsonRequest) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.

### HasManualChain

`func (o *PkiJsonRequest) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.

### GetUsage

`func (o *PkiJsonRequest) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiJsonRequest) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiJsonRequest) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PkiJsonRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


