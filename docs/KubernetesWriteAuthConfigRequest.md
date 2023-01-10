# KubernetesWriteAuthConfigRequest

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

### NewKubernetesWriteAuthConfigRequest

`func NewKubernetesWriteAuthConfigRequest() *KubernetesWriteAuthConfigRequest`

NewKubernetesWriteAuthConfigRequest instantiates a new KubernetesWriteAuthConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesWriteAuthConfigRequestWithDefaults

`func NewKubernetesWriteAuthConfigRequestWithDefaults() *KubernetesWriteAuthConfigRequest`

NewKubernetesWriteAuthConfigRequestWithDefaults instantiates a new KubernetesWriteAuthConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableIssValidation

`func (o *KubernetesWriteAuthConfigRequest) GetDisableIssValidation() bool`

GetDisableIssValidation returns the DisableIssValidation field if non-nil, zero value otherwise.

### GetDisableIssValidationOk

`func (o *KubernetesWriteAuthConfigRequest) GetDisableIssValidationOk() (*bool, bool)`

GetDisableIssValidationOk returns a tuple with the DisableIssValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssValidation

`func (o *KubernetesWriteAuthConfigRequest) SetDisableIssValidation(v bool)`

SetDisableIssValidation sets DisableIssValidation field to given value.

### HasDisableIssValidation

`func (o *KubernetesWriteAuthConfigRequest) HasDisableIssValidation() bool`

HasDisableIssValidation returns a boolean if a field has been set.

### GetDisableLocalCaJwt

`func (o *KubernetesWriteAuthConfigRequest) GetDisableLocalCaJwt() bool`

GetDisableLocalCaJwt returns the DisableLocalCaJwt field if non-nil, zero value otherwise.

### GetDisableLocalCaJwtOk

`func (o *KubernetesWriteAuthConfigRequest) GetDisableLocalCaJwtOk() (*bool, bool)`

GetDisableLocalCaJwtOk returns a tuple with the DisableLocalCaJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocalCaJwt

`func (o *KubernetesWriteAuthConfigRequest) SetDisableLocalCaJwt(v bool)`

SetDisableLocalCaJwt sets DisableLocalCaJwt field to given value.

### HasDisableLocalCaJwt

`func (o *KubernetesWriteAuthConfigRequest) HasDisableLocalCaJwt() bool`

HasDisableLocalCaJwt returns a boolean if a field has been set.

### GetIssuer

`func (o *KubernetesWriteAuthConfigRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *KubernetesWriteAuthConfigRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *KubernetesWriteAuthConfigRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *KubernetesWriteAuthConfigRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKubernetesCaCert

`func (o *KubernetesWriteAuthConfigRequest) GetKubernetesCaCert() string`

GetKubernetesCaCert returns the KubernetesCaCert field if non-nil, zero value otherwise.

### GetKubernetesCaCertOk

`func (o *KubernetesWriteAuthConfigRequest) GetKubernetesCaCertOk() (*string, bool)`

GetKubernetesCaCertOk returns a tuple with the KubernetesCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCaCert

`func (o *KubernetesWriteAuthConfigRequest) SetKubernetesCaCert(v string)`

SetKubernetesCaCert sets KubernetesCaCert field to given value.

### HasKubernetesCaCert

`func (o *KubernetesWriteAuthConfigRequest) HasKubernetesCaCert() bool`

HasKubernetesCaCert returns a boolean if a field has been set.

### GetKubernetesHost

`func (o *KubernetesWriteAuthConfigRequest) GetKubernetesHost() string`

GetKubernetesHost returns the KubernetesHost field if non-nil, zero value otherwise.

### GetKubernetesHostOk

`func (o *KubernetesWriteAuthConfigRequest) GetKubernetesHostOk() (*string, bool)`

GetKubernetesHostOk returns a tuple with the KubernetesHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHost

`func (o *KubernetesWriteAuthConfigRequest) SetKubernetesHost(v string)`

SetKubernetesHost sets KubernetesHost field to given value.

### HasKubernetesHost

`func (o *KubernetesWriteAuthConfigRequest) HasKubernetesHost() bool`

HasKubernetesHost returns a boolean if a field has been set.

### GetPemKeys

`func (o *KubernetesWriteAuthConfigRequest) GetPemKeys() []string`

GetPemKeys returns the PemKeys field if non-nil, zero value otherwise.

### GetPemKeysOk

`func (o *KubernetesWriteAuthConfigRequest) GetPemKeysOk() (*[]string, bool)`

GetPemKeysOk returns a tuple with the PemKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemKeys

`func (o *KubernetesWriteAuthConfigRequest) SetPemKeys(v []string)`

SetPemKeys sets PemKeys field to given value.

### HasPemKeys

`func (o *KubernetesWriteAuthConfigRequest) HasPemKeys() bool`

HasPemKeys returns a boolean if a field has been set.

### GetTokenReviewerJwt

`func (o *KubernetesWriteAuthConfigRequest) GetTokenReviewerJwt() string`

GetTokenReviewerJwt returns the TokenReviewerJwt field if non-nil, zero value otherwise.

### GetTokenReviewerJwtOk

`func (o *KubernetesWriteAuthConfigRequest) GetTokenReviewerJwtOk() (*string, bool)`

GetTokenReviewerJwtOk returns a tuple with the TokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReviewerJwt

`func (o *KubernetesWriteAuthConfigRequest) SetTokenReviewerJwt(v string)`

SetTokenReviewerJwt sets TokenReviewerJwt field to given value.

### HasTokenReviewerJwt

`func (o *KubernetesWriteAuthConfigRequest) HasTokenReviewerJwt() bool`

HasTokenReviewerJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


