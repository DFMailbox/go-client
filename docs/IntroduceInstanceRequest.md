# IntroduceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | A base64 URL encoded ed25519 public key | 
**Address** | **string** | An address pointing to another DFMailbox instance | 
**Update** | Pointer to **bool** | Whether to update the key or not | [optional] 

## Methods

### NewIntroduceInstanceRequest

`func NewIntroduceInstanceRequest(publicKey string, address string, ) *IntroduceInstanceRequest`

NewIntroduceInstanceRequest instantiates a new IntroduceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntroduceInstanceRequestWithDefaults

`func NewIntroduceInstanceRequestWithDefaults() *IntroduceInstanceRequest`

NewIntroduceInstanceRequestWithDefaults instantiates a new IntroduceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *IntroduceInstanceRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *IntroduceInstanceRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *IntroduceInstanceRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetAddress

`func (o *IntroduceInstanceRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntroduceInstanceRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntroduceInstanceRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetUpdate

`func (o *IntroduceInstanceRequest) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *IntroduceInstanceRequest) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *IntroduceInstanceRequest) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *IntroduceInstanceRequest) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


