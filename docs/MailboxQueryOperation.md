# MailboxQueryOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**After** | **int64** | A mailbox item id that refers to an item in any mailbox | 
**Limit** | **int32** |  | 
**Value** | **interface{}** |  | 
**To** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**BeforeAt** | **int64** | A mailbox item id that refers to an item in any mailbox | 

## Methods

### NewMailboxQueryOperation

`func NewMailboxQueryOperation(type_ interface{}, after int64, limit int32, value interface{}, to int32, beforeAt int64, ) *MailboxQueryOperation`

NewMailboxQueryOperation instantiates a new MailboxQueryOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxQueryOperationWithDefaults

`func NewMailboxQueryOperationWithDefaults() *MailboxQueryOperation`

NewMailboxQueryOperationWithDefaults instantiates a new MailboxQueryOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MailboxQueryOperation) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailboxQueryOperation) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailboxQueryOperation) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MailboxQueryOperation) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MailboxQueryOperation) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAfter

`func (o *MailboxQueryOperation) GetAfter() int64`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *MailboxQueryOperation) GetAfterOk() (*int64, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *MailboxQueryOperation) SetAfter(v int64)`

SetAfter sets After field to given value.


### GetLimit

`func (o *MailboxQueryOperation) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MailboxQueryOperation) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MailboxQueryOperation) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetValue

`func (o *MailboxQueryOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MailboxQueryOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MailboxQueryOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *MailboxQueryOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MailboxQueryOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTo

`func (o *MailboxQueryOperation) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MailboxQueryOperation) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MailboxQueryOperation) SetTo(v int32)`

SetTo sets To field to given value.


### GetBeforeAt

`func (o *MailboxQueryOperation) GetBeforeAt() int64`

GetBeforeAt returns the BeforeAt field if non-nil, zero value otherwise.

### GetBeforeAtOk

`func (o *MailboxQueryOperation) GetBeforeAtOk() (*int64, bool)`

GetBeforeAtOk returns a tuple with the BeforeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeAt

`func (o *MailboxQueryOperation) SetBeforeAt(v int64)`

SetBeforeAt sets BeforeAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


