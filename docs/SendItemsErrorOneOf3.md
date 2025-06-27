# SendItemsErrorOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**Sender** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**Receiver** | Pointer to **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | [optional] 

## Methods

### NewSendItemsErrorOneOf3

`func NewSendItemsErrorOneOf3(type_ interface{}, title interface{}, status int32, detail string, sender int32, ) *SendItemsErrorOneOf3`

NewSendItemsErrorOneOf3 instantiates a new SendItemsErrorOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendItemsErrorOneOf3WithDefaults

`func NewSendItemsErrorOneOf3WithDefaults() *SendItemsErrorOneOf3`

NewSendItemsErrorOneOf3WithDefaults instantiates a new SendItemsErrorOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendItemsErrorOneOf3) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendItemsErrorOneOf3) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendItemsErrorOneOf3) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SendItemsErrorOneOf3) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SendItemsErrorOneOf3) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SendItemsErrorOneOf3) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendItemsErrorOneOf3) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendItemsErrorOneOf3) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SendItemsErrorOneOf3) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SendItemsErrorOneOf3) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *SendItemsErrorOneOf3) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendItemsErrorOneOf3) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendItemsErrorOneOf3) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *SendItemsErrorOneOf3) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SendItemsErrorOneOf3) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SendItemsErrorOneOf3) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSender

`func (o *SendItemsErrorOneOf3) GetSender() int32`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SendItemsErrorOneOf3) GetSenderOk() (*int32, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SendItemsErrorOneOf3) SetSender(v int32)`

SetSender sets Sender field to given value.


### GetReceiver

`func (o *SendItemsErrorOneOf3) GetReceiver() int32`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SendItemsErrorOneOf3) GetReceiverOk() (*int32, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SendItemsErrorOneOf3) SetReceiver(v int32)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *SendItemsErrorOneOf3) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


