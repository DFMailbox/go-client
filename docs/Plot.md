# Plot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlotId** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**Owner** | **string** | A UUID (universally unique identifier) | 
**PublicKey** | **NullableString** | A base64 URL encoded ed25519 public key | 
**Address** | **NullableString** | An &#x60;Address&#x60;, if null, means it has the private key has been compromised | 
**MailboxMsgId** | **int64** | A mailbox item id that refers to an item in any mailbox | 

## Methods

### NewPlot

`func NewPlot(plotId int32, owner string, publicKey NullableString, address NullableString, mailboxMsgId int64, ) *Plot`

NewPlot instantiates a new Plot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlotWithDefaults

`func NewPlotWithDefaults() *Plot`

NewPlotWithDefaults instantiates a new Plot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlotId

`func (o *Plot) GetPlotId() int32`

GetPlotId returns the PlotId field if non-nil, zero value otherwise.

### GetPlotIdOk

`func (o *Plot) GetPlotIdOk() (*int32, bool)`

GetPlotIdOk returns a tuple with the PlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotId

`func (o *Plot) SetPlotId(v int32)`

SetPlotId sets PlotId field to given value.


### GetOwner

`func (o *Plot) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Plot) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Plot) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPublicKey

`func (o *Plot) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Plot) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Plot) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *Plot) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *Plot) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetAddress

`func (o *Plot) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Plot) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Plot) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *Plot) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Plot) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetMailboxMsgId

`func (o *Plot) GetMailboxMsgId() int64`

GetMailboxMsgId returns the MailboxMsgId field if non-nil, zero value otherwise.

### GetMailboxMsgIdOk

`func (o *Plot) GetMailboxMsgIdOk() (*int64, bool)`

GetMailboxMsgIdOk returns a tuple with the MailboxMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxMsgId

`func (o *Plot) SetMailboxMsgId(v int64)`

SetMailboxMsgId sets MailboxMsgId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


