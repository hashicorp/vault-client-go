# PKIIssuerSignRevocationListRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlNumber** | Pointer to **int32** | The sequence number to be written within the CRL Number extension. | [optional] 
**DeltaCrlBaseNumber** | Pointer to **int32** | Using a zero or greater value specifies the base CRL revision number to encode within a Delta CRL indicator extension, otherwise the extension will not be added. | [optional] [default to -1]
**Extensions** | Pointer to **[]map[string]interface{}** | A list of maps containing extensions with keys id (string), critical (bool), value (string) | [optional] 
**Format** | Pointer to **string** | The format of the combined CRL, can be \&quot;pem\&quot; or \&quot;der\&quot;. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**NextUpdate** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours. | [optional] [default to "72h"]
**RevokedCerts** | Pointer to **[]map[string]interface{}** | A list of maps containing the keys serial_number (string), revocation_time (string), and extensions (map with keys id (string), critical (bool), value (string)) | [optional] 



## Methods


### NewPKIIssuerSignRevocationListRequest

`func NewPKIIssuerSignRevocationListRequest() *PKIIssuerSignRevocationListRequest`

NewPKIIssuerSignRevocationListRequest instantiates a new PKIIssuerSignRevocationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIIssuerSignRevocationListRequestWithDefaults

`func NewPKIIssuerSignRevocationListRequestWithDefaults() *PKIIssuerSignRevocationListRequest`

NewPKIIssuerSignRevocationListRequestWithDefaults instantiates a new PKIIssuerSignRevocationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCrlNumber

`func (o *PKIIssuerSignRevocationListRequest) GetCrlNumber() int32`

GetCrlNumber returns the CrlNumber field if non-nil, zero value otherwise.

### GetCrlNumberOk

`func (o *PKIIssuerSignRevocationListRequest) GetCrlNumberOk() (*int32, bool)`

GetCrlNumberOk returns a tuple with the CrlNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlNumber

`func (o *PKIIssuerSignRevocationListRequest) SetCrlNumber(v int32)`

SetCrlNumber sets CrlNumber field to given value.


### HasCrlNumber

`func (o *PKIIssuerSignRevocationListRequest) HasCrlNumber() bool`

HasCrlNumber returns a boolean if a field has been set.




### GetDeltaCrlBaseNumber

`func (o *PKIIssuerSignRevocationListRequest) GetDeltaCrlBaseNumber() int32`

GetDeltaCrlBaseNumber returns the DeltaCrlBaseNumber field if non-nil, zero value otherwise.

### GetDeltaCrlBaseNumberOk

`func (o *PKIIssuerSignRevocationListRequest) GetDeltaCrlBaseNumberOk() (*int32, bool)`

GetDeltaCrlBaseNumberOk returns a tuple with the DeltaCrlBaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCrlBaseNumber

`func (o *PKIIssuerSignRevocationListRequest) SetDeltaCrlBaseNumber(v int32)`

SetDeltaCrlBaseNumber sets DeltaCrlBaseNumber field to given value.


### HasDeltaCrlBaseNumber

`func (o *PKIIssuerSignRevocationListRequest) HasDeltaCrlBaseNumber() bool`

HasDeltaCrlBaseNumber returns a boolean if a field has been set.




### GetExtensions

`func (o *PKIIssuerSignRevocationListRequest) GetExtensions() []map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *PKIIssuerSignRevocationListRequest) GetExtensionsOk() (*[]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *PKIIssuerSignRevocationListRequest) SetExtensions(v []map[string]interface{})`

SetExtensions sets Extensions field to given value.


### HasExtensions

`func (o *PKIIssuerSignRevocationListRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.




### GetFormat

`func (o *PKIIssuerSignRevocationListRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PKIIssuerSignRevocationListRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PKIIssuerSignRevocationListRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *PKIIssuerSignRevocationListRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetNextUpdate

`func (o *PKIIssuerSignRevocationListRequest) GetNextUpdate() string`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *PKIIssuerSignRevocationListRequest) GetNextUpdateOk() (*string, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *PKIIssuerSignRevocationListRequest) SetNextUpdate(v string)`

SetNextUpdate sets NextUpdate field to given value.


### HasNextUpdate

`func (o *PKIIssuerSignRevocationListRequest) HasNextUpdate() bool`

HasNextUpdate returns a boolean if a field has been set.




### GetRevokedCerts

`func (o *PKIIssuerSignRevocationListRequest) GetRevokedCerts() []map[string]interface{}`

GetRevokedCerts returns the RevokedCerts field if non-nil, zero value otherwise.

### GetRevokedCertsOk

`func (o *PKIIssuerSignRevocationListRequest) GetRevokedCertsOk() (*[]map[string]interface{}, bool)`

GetRevokedCertsOk returns a tuple with the RevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCerts

`func (o *PKIIssuerSignRevocationListRequest) SetRevokedCerts(v []map[string]interface{})`

SetRevokedCerts sets RevokedCerts field to given value.


### HasRevokedCerts

`func (o *PKIIssuerSignRevocationListRequest) HasRevokedCerts() bool`

HasRevokedCerts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


