# MailboxQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Result** | [**FetchMailboxResult**](FetchMailboxResult.md) |  | 
**MsgId** | **int64** | A mailbox item id that refers to an item in any mailbox | 
**Msg** | [**SendItemsError**](SendItemsError.md) |  | 

## Methods

### NewMailboxQueryResponse

`func NewMailboxQueryResponse(type_ interface{}, result FetchMailboxResult, msgId int64, msg SendItemsError, ) *MailboxQueryResponse`

NewMailboxQueryResponse instantiates a new MailboxQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxQueryResponseWithDefaults

`func NewMailboxQueryResponseWithDefaults() *MailboxQueryResponse`

NewMailboxQueryResponseWithDefaults instantiates a new MailboxQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MailboxQueryResponse) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailboxQueryResponse) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailboxQueryResponse) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MailboxQueryResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MailboxQueryResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetResult

`func (o *MailboxQueryResponse) GetResult() FetchMailboxResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MailboxQueryResponse) GetResultOk() (*FetchMailboxResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MailboxQueryResponse) SetResult(v FetchMailboxResult)`

SetResult sets Result field to given value.


### GetMsgId

`func (o *MailboxQueryResponse) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MailboxQueryResponse) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MailboxQueryResponse) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.


### GetMsg

`func (o *MailboxQueryResponse) GetMsg() SendItemsError`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *MailboxQueryResponse) GetMsgOk() (*SendItemsError, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *MailboxQueryResponse) SetMsg(v SendItemsError)`

SetMsg sets Msg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


