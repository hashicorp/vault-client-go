# NomadConfigureAccessRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Nomad server address | [optional] 
**CaCert** | Pointer to **string** | CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded. | [optional] 
**ClientCert** | Pointer to **string** | Client certificate used for Nomad&#x27;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key. | [optional] 
**ClientKey** | Pointer to **string** | Client key used for Nomad&#x27;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert. | [optional] 
**MaxTokenNameLength** | Pointer to **int32** | Max length for name of generated Nomad tokens | [optional] 
**Token** | Pointer to **string** | Token for API calls | [optional] 



## Methods


### NewNomadConfigureAccessRequest

`func NewNomadConfigureAccessRequest() *NomadConfigureAccessRequest`

NewNomadConfigureAccessRequest instantiates a new NomadConfigureAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadConfigureAccessRequestWithDefaults

`func NewNomadConfigureAccessRequestWithDefaults() *NomadConfigureAccessRequest`

NewNomadConfigureAccessRequestWithDefaults instantiates a new NomadConfigureAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAddress

`func (o *NomadConfigureAccessRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NomadConfigureAccessRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NomadConfigureAccessRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### HasAddress

`func (o *NomadConfigureAccessRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.




### GetCaCert

`func (o *NomadConfigureAccessRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *NomadConfigureAccessRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *NomadConfigureAccessRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.


### HasCaCert

`func (o *NomadConfigureAccessRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.




### GetClientCert

`func (o *NomadConfigureAccessRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *NomadConfigureAccessRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *NomadConfigureAccessRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.


### HasClientCert

`func (o *NomadConfigureAccessRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.




### GetClientKey

`func (o *NomadConfigureAccessRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *NomadConfigureAccessRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *NomadConfigureAccessRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### HasClientKey

`func (o *NomadConfigureAccessRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.




### GetMaxTokenNameLength

`func (o *NomadConfigureAccessRequest) GetMaxTokenNameLength() int32`

GetMaxTokenNameLength returns the MaxTokenNameLength field if non-nil, zero value otherwise.

### GetMaxTokenNameLengthOk

`func (o *NomadConfigureAccessRequest) GetMaxTokenNameLengthOk() (*int32, bool)`

GetMaxTokenNameLengthOk returns a tuple with the MaxTokenNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenNameLength

`func (o *NomadConfigureAccessRequest) SetMaxTokenNameLength(v int32)`

SetMaxTokenNameLength sets MaxTokenNameLength field to given value.


### HasMaxTokenNameLength

`func (o *NomadConfigureAccessRequest) HasMaxTokenNameLength() bool`

HasMaxTokenNameLength returns a boolean if a field has been set.




### GetToken

`func (o *NomadConfigureAccessRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NomadConfigureAccessRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NomadConfigureAccessRequest) SetToken(v string)`

SetToken sets Token field to given value.


### HasToken

`func (o *NomadConfigureAccessRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


