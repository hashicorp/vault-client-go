# PkiWriteAcmeOrderOrderIdFinalizeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **string** | ACME request &#x27;payload&#x27; value | [optional] 
**Protected** | Pointer to **string** | ACME request &#x27;protected&#x27; value | [optional] 
**Signature** | Pointer to **string** | ACME request &#x27;signature&#x27; value | [optional] 



## Methods


### NewPkiWriteAcmeOrderOrderIdFinalizeRequest

`func NewPkiWriteAcmeOrderOrderIdFinalizeRequest() *PkiWriteAcmeOrderOrderIdFinalizeRequest`

NewPkiWriteAcmeOrderOrderIdFinalizeRequest instantiates a new PkiWriteAcmeOrderOrderIdFinalizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiWriteAcmeOrderOrderIdFinalizeRequestWithDefaults

`func NewPkiWriteAcmeOrderOrderIdFinalizeRequestWithDefaults() *PkiWriteAcmeOrderOrderIdFinalizeRequest`

NewPkiWriteAcmeOrderOrderIdFinalizeRequestWithDefaults instantiates a new PkiWriteAcmeOrderOrderIdFinalizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPayload

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### HasPayload

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.




### GetProtected

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetProtected() string`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetProtectedOk() (*string, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) SetProtected(v string)`

SetProtected sets Protected field to given value.


### HasProtected

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) HasProtected() bool`

HasProtected returns a boolean if a field has been set.




### GetSignature

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### HasSignature

`func (o *PkiWriteAcmeOrderOrderIdFinalizeRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


