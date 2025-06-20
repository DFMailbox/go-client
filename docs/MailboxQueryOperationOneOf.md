# MailboxQueryOperationOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**After** | **int64** | A mailbox item id that refers to an item in any mailbox | 
**Limit** | **NullableInt32** |  | 

## Methods

### NewMailboxQueryOperationOneOf

`func NewMailboxQueryOperationOneOf(type_ interface{}, after int64, limit NullableInt32, ) *MailboxQueryOperationOneOf`

NewMailboxQueryOperationOneOf instantiates a new MailboxQueryOperationOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxQueryOperationOneOfWithDefaults

`func NewMailboxQueryOperationOneOfWithDefaults() *MailboxQueryOperationOneOf`

NewMailboxQueryOperationOneOfWithDefaults instantiates a new MailboxQueryOperationOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MailboxQueryOperationOneOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailboxQueryOperationOneOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailboxQueryOperationOneOf) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MailboxQueryOperationOneOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MailboxQueryOperationOneOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAfter

`func (o *MailboxQueryOperationOneOf) GetAfter() int64`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *MailboxQueryOperationOneOf) GetAfterOk() (*int64, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *MailboxQueryOperationOneOf) SetAfter(v int64)`

SetAfter sets After field to given value.


### GetLimit

`func (o *MailboxQueryOperationOneOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MailboxQueryOperationOneOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MailboxQueryOperationOneOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *MailboxQueryOperationOneOf) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *MailboxQueryOperationOneOf) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


