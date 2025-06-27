# MaybeAddressKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | An nullable address pointing to another DFMailbox instance. Null, in this case meaning the instance is compromised. | 
**PublicKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewMaybeAddressKeyPair

`func NewMaybeAddressKeyPair(address NullableString, publicKey string, ) *MaybeAddressKeyPair`

NewMaybeAddressKeyPair instantiates a new MaybeAddressKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaybeAddressKeyPairWithDefaults

`func NewMaybeAddressKeyPairWithDefaults() *MaybeAddressKeyPair`

NewMaybeAddressKeyPairWithDefaults instantiates a new MaybeAddressKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MaybeAddressKeyPair) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MaybeAddressKeyPair) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MaybeAddressKeyPair) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *MaybeAddressKeyPair) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MaybeAddressKeyPair) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPublicKey

`func (o *MaybeAddressKeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *MaybeAddressKeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *MaybeAddressKeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


