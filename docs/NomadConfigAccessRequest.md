# NomadConfigAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Nomad server address | [optional] 
**CaCert** | Pointer to **string** | CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded. | [optional] 
**ClientCert** | Pointer to **string** | Client certificate used for Nomad&#39;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key. | [optional] 
**ClientKey** | Pointer to **string** | Client key used for Nomad&#39;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert. | [optional] 
**MaxTokenNameLength** | Pointer to **int32** | Max length for name of generated Nomad tokens | [optional] 
**Token** | Pointer to **string** | Token for API calls | [optional] 

## Methods

### NewNomadConfigAccessRequest

`func NewNomadConfigAccessRequest() *NomadConfigAccessRequest`

NewNomadConfigAccessRequest instantiates a new NomadConfigAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadConfigAccessRequestWithDefaults

`func NewNomadConfigAccessRequestWithDefaults() *NomadConfigAccessRequest`

NewNomadConfigAccessRequestWithDefaults instantiates a new NomadConfigAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NomadConfigAccessRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NomadConfigAccessRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NomadConfigAccessRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NomadConfigAccessRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *NomadConfigAccessRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *NomadConfigAccessRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *NomadConfigAccessRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *NomadConfigAccessRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *NomadConfigAccessRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *NomadConfigAccessRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *NomadConfigAccessRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *NomadConfigAccessRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *NomadConfigAccessRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *NomadConfigAccessRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *NomadConfigAccessRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *NomadConfigAccessRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetMaxTokenNameLength

`func (o *NomadConfigAccessRequest) GetMaxTokenNameLength() int32`

GetMaxTokenNameLength returns the MaxTokenNameLength field if non-nil, zero value otherwise.

### GetMaxTokenNameLengthOk

`func (o *NomadConfigAccessRequest) GetMaxTokenNameLengthOk() (*int32, bool)`

GetMaxTokenNameLengthOk returns a tuple with the MaxTokenNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenNameLength

`func (o *NomadConfigAccessRequest) SetMaxTokenNameLength(v int32)`

SetMaxTokenNameLength sets MaxTokenNameLength field to given value.

### HasMaxTokenNameLength

`func (o *NomadConfigAccessRequest) HasMaxTokenNameLength() bool`

HasMaxTokenNameLength returns a boolean if a field has been set.

### GetToken

`func (o *NomadConfigAccessRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NomadConfigAccessRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NomadConfigAccessRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *NomadConfigAccessRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


