# MailboxQueryResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Result** | [**FetchMailboxResult**](FetchMailboxResult.md) |  | 

## Methods

### NewMailboxQueryResponseOneOf

`func NewMailboxQueryResponseOneOf(type_ interface{}, result FetchMailboxResult, ) *MailboxQueryResponseOneOf`

NewMailboxQueryResponseOneOf instantiates a new MailboxQueryResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxQueryResponseOneOfWithDefaults

`func NewMailboxQueryResponseOneOfWithDefaults() *MailboxQueryResponseOneOf`

NewMailboxQueryResponseOneOfWithDefaults instantiates a new MailboxQueryResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MailboxQueryResponseOneOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailboxQueryResponseOneOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailboxQueryResponseOneOf) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MailboxQueryResponseOneOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MailboxQueryResponseOneOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetResult

`func (o *MailboxQueryResponseOneOf) GetResult() FetchMailboxResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MailboxQueryResponseOneOf) GetResultOk() (*FetchMailboxResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MailboxQueryResponseOneOf) SetResult(v FetchMailboxResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


