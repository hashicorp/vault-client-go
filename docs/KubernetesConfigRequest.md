# KubernetesConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableIssValidation** | Pointer to **bool** | Disable JWT issuer validation (Deprecated, will be removed in a future release) | [optional] [default to true]
**DisableLocalCaJwt** | Pointer to **bool** | Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod | [optional] [default to false]
**Issuer** | Pointer to **string** | Optional JWT issuer. If no issuer is specified, then this plugin will use kubernetes.io/serviceaccount as the default issuer. (Deprecated, will be removed in a future release) | [optional] 
**KubernetesCaCert** | Pointer to **string** | PEM encoded CA cert for use by the TLS client used to talk with the API. | [optional] 
**KubernetesHost** | Pointer to **string** | Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server. | [optional] 
**PemKeys** | Pointer to **[]string** | Optional list of PEM-formated public keys or certificates used to verify the signatures of kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys. | [optional] 
**TokenReviewerJwt** | Pointer to **string** | A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API. | [optional] 

## Methods

### NewKubernetesConfigRequest

`func NewKubernetesConfigRequest() *KubernetesConfigRequest`

NewKubernetesConfigRequest instantiates a new KubernetesConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigRequestWithDefaults

`func NewKubernetesConfigRequestWithDefaults() *KubernetesConfigRequest`

NewKubernetesConfigRequestWithDefaults instantiates a new KubernetesConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableIssValidation

`func (o *KubernetesConfigRequest) GetDisableIssValidation() bool`

GetDisableIssValidation returns the DisableIssValidation field if non-nil, zero value otherwise.

### GetDisableIssValidationOk

`func (o *KubernetesConfigRequest) GetDisableIssValidationOk() (*bool, bool)`

GetDisableIssValidationOk returns a tuple with the DisableIssValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssValidation

`func (o *KubernetesConfigRequest) SetDisableIssValidation(v bool)`

SetDisableIssValidation sets DisableIssValidation field to given value.

### HasDisableIssValidation

`func (o *KubernetesConfigRequest) HasDisableIssValidation() bool`

HasDisableIssValidation returns a boolean if a field has been set.

### GetDisableLocalCaJwt

`func (o *KubernetesConfigRequest) GetDisableLocalCaJwt() bool`

GetDisableLocalCaJwt returns the DisableLocalCaJwt field if non-nil, zero value otherwise.

### GetDisableLocalCaJwtOk

`func (o *KubernetesConfigRequest) GetDisableLocalCaJwtOk() (*bool, bool)`

GetDisableLocalCaJwtOk returns a tuple with the DisableLocalCaJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocalCaJwt

`func (o *KubernetesConfigRequest) SetDisableLocalCaJwt(v bool)`

SetDisableLocalCaJwt sets DisableLocalCaJwt field to given value.

### HasDisableLocalCaJwt

`func (o *KubernetesConfigRequest) HasDisableLocalCaJwt() bool`

HasDisableLocalCaJwt returns a boolean if a field has been set.

### GetIssuer

`func (o *KubernetesConfigRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *KubernetesConfigRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *KubernetesConfigRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *KubernetesConfigRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKubernetesCaCert

`func (o *KubernetesConfigRequest) GetKubernetesCaCert() string`

GetKubernetesCaCert returns the KubernetesCaCert field if non-nil, zero value otherwise.

### GetKubernetesCaCertOk

`func (o *KubernetesConfigRequest) GetKubernetesCaCertOk() (*string, bool)`

GetKubernetesCaCertOk returns a tuple with the KubernetesCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCaCert

`func (o *KubernetesConfigRequest) SetKubernetesCaCert(v string)`

SetKubernetesCaCert sets KubernetesCaCert field to given value.

### HasKubernetesCaCert

`func (o *KubernetesConfigRequest) HasKubernetesCaCert() bool`

HasKubernetesCaCert returns a boolean if a field has been set.

### GetKubernetesHost

`func (o *KubernetesConfigRequest) GetKubernetesHost() string`

GetKubernetesHost returns the KubernetesHost field if non-nil, zero value otherwise.

### GetKubernetesHostOk

`func (o *KubernetesConfigRequest) GetKubernetesHostOk() (*string, bool)`

GetKubernetesHostOk returns a tuple with the KubernetesHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHost

`func (o *KubernetesConfigRequest) SetKubernetesHost(v string)`

SetKubernetesHost sets KubernetesHost field to given value.

### HasKubernetesHost

`func (o *KubernetesConfigRequest) HasKubernetesHost() bool`

HasKubernetesHost returns a boolean if a field has been set.

### GetPemKeys

`func (o *KubernetesConfigRequest) GetPemKeys() []string`

GetPemKeys returns the PemKeys field if non-nil, zero value otherwise.

### GetPemKeysOk

`func (o *KubernetesConfigRequest) GetPemKeysOk() (*[]string, bool)`

GetPemKeysOk returns a tuple with the PemKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemKeys

`func (o *KubernetesConfigRequest) SetPemKeys(v []string)`

SetPemKeys sets PemKeys field to given value.

### HasPemKeys

`func (o *KubernetesConfigRequest) HasPemKeys() bool`

HasPemKeys returns a boolean if a field has been set.

### GetTokenReviewerJwt

`func (o *KubernetesConfigRequest) GetTokenReviewerJwt() string`

GetTokenReviewerJwt returns the TokenReviewerJwt field if non-nil, zero value otherwise.

### GetTokenReviewerJwtOk

`func (o *KubernetesConfigRequest) GetTokenReviewerJwtOk() (*string, bool)`

GetTokenReviewerJwtOk returns a tuple with the TokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReviewerJwt

`func (o *KubernetesConfigRequest) SetTokenReviewerJwt(v string)`

SetTokenReviewerJwt sets TokenReviewerJwt field to given value.

### HasTokenReviewerJwt

`func (o *KubernetesConfigRequest) HasTokenReviewerJwt() bool`

HasTokenReviewerJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


