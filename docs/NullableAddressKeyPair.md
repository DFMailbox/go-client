# NullableAddressKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An address pointing to another DFMailbox instance | 
**PublicKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewNullableAddressKeyPair

`func NewNullableAddressKeyPair(address string, publicKey string, ) *NullableAddressKeyPair`

NewNullableAddressKeyPair instantiates a new NullableAddressKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableAddressKeyPairWithDefaults

`func NewNullableAddressKeyPairWithDefaults() *NullableAddressKeyPair`

NewNullableAddressKeyPairWithDefaults instantiates a new NullableAddressKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NullableAddressKeyPair) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NullableAddressKeyPair) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NullableAddressKeyPair) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPublicKey

`func (o *NullableAddressKeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *NullableAddressKeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *NullableAddressKeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


