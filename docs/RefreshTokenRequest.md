# RefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | **string** | A UUID (universally unique identifier) | 
**Signature** | **string** | A base64 encoded ed25519 signature | 

## Methods

### NewRefreshTokenRequest

`func NewRefreshTokenRequest(challenge string, signature string, ) *RefreshTokenRequest`

NewRefreshTokenRequest instantiates a new RefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenRequestWithDefaults

`func NewRefreshTokenRequestWithDefaults() *RefreshTokenRequest`

NewRefreshTokenRequestWithDefaults instantiates a new RefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *RefreshTokenRequest) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *RefreshTokenRequest) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *RefreshTokenRequest) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetSignature

`func (o *RefreshTokenRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RefreshTokenRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RefreshTokenRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


