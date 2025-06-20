# MailboxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | **int64** | A mailbox item id that refers to an item in any mailbox | 
**Timestamp** | **int32** | A unix timestamp | 
**PlotOrigin** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**Val** | **interface{}** |  | 

## Methods

### NewMailboxItem

`func NewMailboxItem(msgId int64, timestamp int32, plotOrigin int32, val interface{}, ) *MailboxItem`

NewMailboxItem instantiates a new MailboxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxItemWithDefaults

`func NewMailboxItemWithDefaults() *MailboxItem`

NewMailboxItemWithDefaults instantiates a new MailboxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgId

`func (o *MailboxItem) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MailboxItem) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MailboxItem) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.


### GetTimestamp

`func (o *MailboxItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MailboxItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MailboxItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetPlotOrigin

`func (o *MailboxItem) GetPlotOrigin() int32`

GetPlotOrigin returns the PlotOrigin field if non-nil, zero value otherwise.

### GetPlotOriginOk

`func (o *MailboxItem) GetPlotOriginOk() (*int32, bool)`

GetPlotOriginOk returns a tuple with the PlotOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotOrigin

`func (o *MailboxItem) SetPlotOrigin(v int32)`

SetPlotOrigin sets PlotOrigin field to given value.


### GetVal

`func (o *MailboxItem) GetVal() interface{}`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *MailboxItem) GetValOk() (*interface{}, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *MailboxItem) SetVal(v interface{})`

SetVal sets Val field to given value.


### SetValNil

`func (o *MailboxItem) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *MailboxItem) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


