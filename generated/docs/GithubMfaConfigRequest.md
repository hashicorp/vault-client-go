# GithubMfaConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Enables MFA with given backend (available: duo) | [optional] 

## Methods

### NewGithubMfaConfigRequest

`func NewGithubMfaConfigRequest() *GithubMfaConfigRequest`

NewGithubMfaConfigRequest instantiates a new GithubMfaConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubMfaConfigRequestWithDefaults

`func NewGithubMfaConfigRequestWithDefaults() *GithubMfaConfigRequest`

NewGithubMfaConfigRequestWithDefaults instantiates a new GithubMfaConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GithubMfaConfigRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubMfaConfigRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubMfaConfigRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GithubMfaConfigRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


