# ConsulWriteAccessConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Consul server address | [optional] 
**CaCert** | Pointer to **string** | CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded. | [optional] 
**ClientCert** | Pointer to **string** | Client certificate used for Consul&#39;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key. | [optional] 
**ClientKey** | Pointer to **string** | Client key used for Consul&#39;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert. | [optional] 
**Scheme** | Pointer to **string** | URI scheme for the Consul address | [optional] [default to "http"]
**Token** | Pointer to **string** | Token for API calls | [optional] 

## Methods

### NewConsulWriteAccessConfigRequest

`func NewConsulWriteAccessConfigRequest() *ConsulWriteAccessConfigRequest`

NewConsulWriteAccessConfigRequest instantiates a new ConsulWriteAccessConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulWriteAccessConfigRequestWithDefaults

`func NewConsulWriteAccessConfigRequestWithDefaults() *ConsulWriteAccessConfigRequest`

NewConsulWriteAccessConfigRequestWithDefaults instantiates a new ConsulWriteAccessConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConsulWriteAccessConfigRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsulWriteAccessConfigRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsulWriteAccessConfigRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsulWriteAccessConfigRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *ConsulWriteAccessConfigRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *ConsulWriteAccessConfigRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *ConsulWriteAccessConfigRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *ConsulWriteAccessConfigRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *ConsulWriteAccessConfigRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ConsulWriteAccessConfigRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ConsulWriteAccessConfigRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ConsulWriteAccessConfigRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *ConsulWriteAccessConfigRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *ConsulWriteAccessConfigRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *ConsulWriteAccessConfigRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *ConsulWriteAccessConfigRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetScheme

`func (o *ConsulWriteAccessConfigRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ConsulWriteAccessConfigRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ConsulWriteAccessConfigRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ConsulWriteAccessConfigRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetToken

`func (o *ConsulWriteAccessConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConsulWriteAccessConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConsulWriteAccessConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConsulWriteAccessConfigRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


