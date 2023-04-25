# MongoDbAtlasConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | MongoDB Atlas Programmatic Private Key | 
**PublicKey** | **string** | MongoDB Atlas Programmatic Public Key | 



## Methods


### NewMongoDbAtlasConfigureRequest

`func NewMongoDbAtlasConfigureRequest(privateKey string, publicKey string, ) *MongoDbAtlasConfigureRequest`

NewMongoDbAtlasConfigureRequest instantiates a new MongoDbAtlasConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDbAtlasConfigureRequestWithDefaults

`func NewMongoDbAtlasConfigureRequestWithDefaults() *MongoDbAtlasConfigureRequest`

NewMongoDbAtlasConfigureRequestWithDefaults instantiates a new MongoDbAtlasConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPrivateKey

`func (o *MongoDbAtlasConfigureRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *MongoDbAtlasConfigureRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *MongoDbAtlasConfigureRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.





### GetPublicKey

`func (o *MongoDbAtlasConfigureRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *MongoDbAtlasConfigureRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *MongoDbAtlasConfigureRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.










[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


