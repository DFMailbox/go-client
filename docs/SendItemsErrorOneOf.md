# SendItemsErrorOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **interface{}** |  | [optional] 
**ExpectedKey** | Pointer to **string** | A base64 URL encoded ed25519 public key | [optional] 

## Methods

### NewSendItemsErrorOneOf

`func NewSendItemsErrorOneOf() *SendItemsErrorOneOf`

NewSendItemsErrorOneOf instantiates a new SendItemsErrorOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendItemsErrorOneOfWithDefaults

`func NewSendItemsErrorOneOfWithDefaults() *SendItemsErrorOneOf`

NewSendItemsErrorOneOfWithDefaults instantiates a new SendItemsErrorOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SendItemsErrorOneOf) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SendItemsErrorOneOf) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SendItemsErrorOneOf) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *SendItemsErrorOneOf) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SendItemsErrorOneOf) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SendItemsErrorOneOf) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExpectedKey

`func (o *SendItemsErrorOneOf) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *SendItemsErrorOneOf) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *SendItemsErrorOneOf) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.

### HasExpectedKey

`func (o *SendItemsErrorOneOf) HasExpectedKey() bool`

HasExpectedKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


