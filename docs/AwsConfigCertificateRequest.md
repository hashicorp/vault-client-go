# AwsConfigCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsPublicCert** | Pointer to **string** | Base64 encoded AWS Public cert required to verify PKCS7 signature of the EC2 instance metadata. | [optional] 
**Type** | Pointer to **string** | Takes the value of either \&quot;pkcs7\&quot; or \&quot;identity\&quot;, indicating the type of document which can be verified using the given certificate. The reason is that the PKCS#7 document will have a DSA digest and the identity signature will have an RSA signature, and accordingly the public certificates to verify those also vary. Defaults to \&quot;pkcs7\&quot;. | [optional] [default to "pkcs7"]

## Methods

### NewAwsConfigCertificateRequest

`func NewAwsConfigCertificateRequest() *AwsConfigCertificateRequest`

NewAwsConfigCertificateRequest instantiates a new AwsConfigCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigCertificateRequestWithDefaults

`func NewAwsConfigCertificateRequestWithDefaults() *AwsConfigCertificateRequest`

NewAwsConfigCertificateRequestWithDefaults instantiates a new AwsConfigCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsPublicCert

`func (o *AwsConfigCertificateRequest) GetAwsPublicCert() string`

GetAwsPublicCert returns the AwsPublicCert field if non-nil, zero value otherwise.

### GetAwsPublicCertOk

`func (o *AwsConfigCertificateRequest) GetAwsPublicCertOk() (*string, bool)`

GetAwsPublicCertOk returns a tuple with the AwsPublicCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPublicCert

`func (o *AwsConfigCertificateRequest) SetAwsPublicCert(v string)`

SetAwsPublicCert sets AwsPublicCert field to given value.

### HasAwsPublicCert

`func (o *AwsConfigCertificateRequest) HasAwsPublicCert() bool`

HasAwsPublicCert returns a boolean if a field has been set.

### GetType

`func (o *AwsConfigCertificateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsConfigCertificateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsConfigCertificateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AwsConfigCertificateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


