# MongoDBAtlasWriteConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | MongoDB Atlas Programmatic Private Key | 
**PublicKey** | **string** | MongoDB Atlas Programmatic Public Key | 

## Methods

### NewMongoDBAtlasWriteConfigRequest

`func NewMongoDBAtlasWriteConfigRequest(privateKey string, publicKey string, ) *MongoDBAtlasWriteConfigRequest`

NewMongoDBAtlasWriteConfigRequest instantiates a new MongoDBAtlasWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBAtlasWriteConfigRequestWithDefaults

`func NewMongoDBAtlasWriteConfigRequestWithDefaults() *MongoDBAtlasWriteConfigRequest`

NewMongoDBAtlasWriteConfigRequestWithDefaults instantiates a new MongoDBAtlasWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *MongoDBAtlasWriteConfigRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *MongoDBAtlasWriteConfigRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *MongoDBAtlasWriteConfigRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPublicKey

`func (o *MongoDBAtlasWriteConfigRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *MongoDBAtlasWriteConfigRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *MongoDBAtlasWriteConfigRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


