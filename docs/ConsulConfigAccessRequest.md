# ConsulConfigAccessRequest

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

### NewConsulConfigAccessRequest

`func NewConsulConfigAccessRequest() *ConsulConfigAccessRequest`

NewConsulConfigAccessRequest instantiates a new ConsulConfigAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulConfigAccessRequestWithDefaults

`func NewConsulConfigAccessRequestWithDefaults() *ConsulConfigAccessRequest`

NewConsulConfigAccessRequestWithDefaults instantiates a new ConsulConfigAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConsulConfigAccessRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsulConfigAccessRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsulConfigAccessRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsulConfigAccessRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *ConsulConfigAccessRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *ConsulConfigAccessRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *ConsulConfigAccessRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *ConsulConfigAccessRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *ConsulConfigAccessRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ConsulConfigAccessRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ConsulConfigAccessRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ConsulConfigAccessRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *ConsulConfigAccessRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *ConsulConfigAccessRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *ConsulConfigAccessRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *ConsulConfigAccessRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetScheme

`func (o *ConsulConfigAccessRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ConsulConfigAccessRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ConsulConfigAccessRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ConsulConfigAccessRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetToken

`func (o *ConsulConfigAccessRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConsulConfigAccessRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConsulConfigAccessRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConsulConfigAccessRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


