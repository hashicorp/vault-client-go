# PkiIssuerResignCrlsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlNumber** | Pointer to **int32** | The sequence number to be written within the CRL Number extension. | [optional] 
**Crls** | Pointer to **[]string** | A list of PEM encoded CRLs to combine, originally signed by the requested issuer. | [optional] 
**DeltaCrlBaseNumber** | Pointer to **int32** | Using a zero or greater value specifies the base CRL revision number to encode within a Delta CRL indicator extension, otherwise the extension will not be added. | [optional] [default to -1]
**Format** | Pointer to **string** | The format of the combined CRL, can be \&quot;pem\&quot; or \&quot;der\&quot;. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**NextUpdate** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours. | [optional] [default to "72h"]

## Methods

### NewPkiIssuerResignCrlsRequest

`func NewPkiIssuerResignCrlsRequest() *PkiIssuerResignCrlsRequest`

NewPkiIssuerResignCrlsRequest instantiates a new PkiIssuerResignCrlsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuerResignCrlsRequestWithDefaults

`func NewPkiIssuerResignCrlsRequestWithDefaults() *PkiIssuerResignCrlsRequest`

NewPkiIssuerResignCrlsRequestWithDefaults instantiates a new PkiIssuerResignCrlsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrlNumber

`func (o *PkiIssuerResignCrlsRequest) GetCrlNumber() int32`

GetCrlNumber returns the CrlNumber field if non-nil, zero value otherwise.

### GetCrlNumberOk

`func (o *PkiIssuerResignCrlsRequest) GetCrlNumberOk() (*int32, bool)`

GetCrlNumberOk returns a tuple with the CrlNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlNumber

`func (o *PkiIssuerResignCrlsRequest) SetCrlNumber(v int32)`

SetCrlNumber sets CrlNumber field to given value.

### HasCrlNumber

`func (o *PkiIssuerResignCrlsRequest) HasCrlNumber() bool`

HasCrlNumber returns a boolean if a field has been set.

### GetCrls

`func (o *PkiIssuerResignCrlsRequest) GetCrls() []string`

GetCrls returns the Crls field if non-nil, zero value otherwise.

### GetCrlsOk

`func (o *PkiIssuerResignCrlsRequest) GetCrlsOk() (*[]string, bool)`

GetCrlsOk returns a tuple with the Crls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrls

`func (o *PkiIssuerResignCrlsRequest) SetCrls(v []string)`

SetCrls sets Crls field to given value.

### HasCrls

`func (o *PkiIssuerResignCrlsRequest) HasCrls() bool`

HasCrls returns a boolean if a field has been set.

### GetDeltaCrlBaseNumber

`func (o *PkiIssuerResignCrlsRequest) GetDeltaCrlBaseNumber() int32`

GetDeltaCrlBaseNumber returns the DeltaCrlBaseNumber field if non-nil, zero value otherwise.

### GetDeltaCrlBaseNumberOk

`func (o *PkiIssuerResignCrlsRequest) GetDeltaCrlBaseNumberOk() (*int32, bool)`

GetDeltaCrlBaseNumberOk returns a tuple with the DeltaCrlBaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCrlBaseNumber

`func (o *PkiIssuerResignCrlsRequest) SetDeltaCrlBaseNumber(v int32)`

SetDeltaCrlBaseNumber sets DeltaCrlBaseNumber field to given value.

### HasDeltaCrlBaseNumber

`func (o *PkiIssuerResignCrlsRequest) HasDeltaCrlBaseNumber() bool`

HasDeltaCrlBaseNumber returns a boolean if a field has been set.

### GetFormat

`func (o *PkiIssuerResignCrlsRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PkiIssuerResignCrlsRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PkiIssuerResignCrlsRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PkiIssuerResignCrlsRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetNextUpdate

`func (o *PkiIssuerResignCrlsRequest) GetNextUpdate() string`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *PkiIssuerResignCrlsRequest) GetNextUpdateOk() (*string, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *PkiIssuerResignCrlsRequest) SetNextUpdate(v string)`

SetNextUpdate sets NextUpdate field to given value.

### HasNextUpdate

`func (o *PkiIssuerResignCrlsRequest) HasNextUpdate() bool`

HasNextUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


