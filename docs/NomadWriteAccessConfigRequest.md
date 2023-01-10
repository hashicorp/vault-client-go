# NomadWriteAccessConfigRequest

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

### NewNomadWriteAccessConfigRequest

`func NewNomadWriteAccessConfigRequest() *NomadWriteAccessConfigRequest`

NewNomadWriteAccessConfigRequest instantiates a new NomadWriteAccessConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadWriteAccessConfigRequestWithDefaults

`func NewNomadWriteAccessConfigRequestWithDefaults() *NomadWriteAccessConfigRequest`

NewNomadWriteAccessConfigRequestWithDefaults instantiates a new NomadWriteAccessConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NomadWriteAccessConfigRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NomadWriteAccessConfigRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NomadWriteAccessConfigRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NomadWriteAccessConfigRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *NomadWriteAccessConfigRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *NomadWriteAccessConfigRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *NomadWriteAccessConfigRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *NomadWriteAccessConfigRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *NomadWriteAccessConfigRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *NomadWriteAccessConfigRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *NomadWriteAccessConfigRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *NomadWriteAccessConfigRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *NomadWriteAccessConfigRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *NomadWriteAccessConfigRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *NomadWriteAccessConfigRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *NomadWriteAccessConfigRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetMaxTokenNameLength

`func (o *NomadWriteAccessConfigRequest) GetMaxTokenNameLength() int32`

GetMaxTokenNameLength returns the MaxTokenNameLength field if non-nil, zero value otherwise.

### GetMaxTokenNameLengthOk

`func (o *NomadWriteAccessConfigRequest) GetMaxTokenNameLengthOk() (*int32, bool)`

GetMaxTokenNameLengthOk returns a tuple with the MaxTokenNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenNameLength

`func (o *NomadWriteAccessConfigRequest) SetMaxTokenNameLength(v int32)`

SetMaxTokenNameLength sets MaxTokenNameLength field to given value.

### HasMaxTokenNameLength

`func (o *NomadWriteAccessConfigRequest) HasMaxTokenNameLength() bool`

HasMaxTokenNameLength returns a boolean if a field has been set.

### GetToken

`func (o *NomadWriteAccessConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NomadWriteAccessConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NomadWriteAccessConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *NomadWriteAccessConfigRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


