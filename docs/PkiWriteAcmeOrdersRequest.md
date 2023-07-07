# PkiWriteAcmeOrdersRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **string** | ACME request &#x27;payload&#x27; value | [optional] 
**Protected** | Pointer to **string** | ACME request &#x27;protected&#x27; value | [optional] 
**Signature** | Pointer to **string** | ACME request &#x27;signature&#x27; value | [optional] 



## Methods


### NewPkiWriteAcmeOrdersRequest

`func NewPkiWriteAcmeOrdersRequest() *PkiWriteAcmeOrdersRequest`

NewPkiWriteAcmeOrdersRequest instantiates a new PkiWriteAcmeOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiWriteAcmeOrdersRequestWithDefaults

`func NewPkiWriteAcmeOrdersRequestWithDefaults() *PkiWriteAcmeOrdersRequest`

NewPkiWriteAcmeOrdersRequestWithDefaults instantiates a new PkiWriteAcmeOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPayload

`func (o *PkiWriteAcmeOrdersRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PkiWriteAcmeOrdersRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PkiWriteAcmeOrdersRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### HasPayload

`func (o *PkiWriteAcmeOrdersRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.




### GetProtected

`func (o *PkiWriteAcmeOrdersRequest) GetProtected() string`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PkiWriteAcmeOrdersRequest) GetProtectedOk() (*string, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PkiWriteAcmeOrdersRequest) SetProtected(v string)`

SetProtected sets Protected field to given value.


### HasProtected

`func (o *PkiWriteAcmeOrdersRequest) HasProtected() bool`

HasProtected returns a boolean if a field has been set.




### GetSignature

`func (o *PkiWriteAcmeOrdersRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PkiWriteAcmeOrdersRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PkiWriteAcmeOrdersRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### HasSignature

`func (o *PkiWriteAcmeOrdersRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


