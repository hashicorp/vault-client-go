# RabbitMqConfigureConnectionRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUri** | Pointer to **string** | RabbitMQ Management URI | [optional] 
**Password** | Pointer to **string** | Password of the provided RabbitMQ management user | [optional] 
**PasswordPolicy** | Pointer to **string** | Name of the password policy to use to generate passwords for dynamic credentials. | [optional] 
**Username** | Pointer to **string** | Username of a RabbitMQ management administrator | [optional] 
**UsernameTemplate** | Pointer to **string** | Template describing how dynamic usernames are generated. | [optional] 
**VerifyConnection** | Pointer to **bool** | If set, connection_uri is verified by actually connecting to the RabbitMQ management API | [optional] [default to true]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


