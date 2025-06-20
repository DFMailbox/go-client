# FetchMailboxResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **interface{}** |  | 
**Until** | **int64** | A mailbox item id that refers to an item in any mailbox | 
**CurrentId** | **int64** | A mailbox item id that refers to an item in any mailbox | 

## Methods

### NewFetchMailboxResult

`func NewFetchMailboxResult(items interface{}, until int64, currentId int64, ) *FetchMailboxResult`

NewFetchMailboxResult instantiates a new FetchMailboxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchMailboxResultWithDefaults

`func NewFetchMailboxResultWithDefaults() *FetchMailboxResult`

NewFetchMailboxResultWithDefaults instantiates a new FetchMailboxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FetchMailboxResult) GetItems() interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FetchMailboxResult) GetItemsOk() (*interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FetchMailboxResult) SetItems(v interface{})`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *FetchMailboxResult) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *FetchMailboxResult) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetUntil

`func (o *FetchMailboxResult) GetUntil() int64`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *FetchMailboxResult) GetUntilOk() (*int64, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *FetchMailboxResult) SetUntil(v int64)`

SetUntil sets Until field to given value.


### GetCurrentId

`func (o *FetchMailboxResult) GetCurrentId() int64`

GetCurrentId returns the CurrentId field if non-nil, zero value otherwise.

### GetCurrentIdOk

`func (o *FetchMailboxResult) GetCurrentIdOk() (*int64, bool)`

GetCurrentIdOk returns a tuple with the CurrentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentId

`func (o *FetchMailboxResult) SetCurrentId(v int64)`

SetCurrentId sets CurrentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


