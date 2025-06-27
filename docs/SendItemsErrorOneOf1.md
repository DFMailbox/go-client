# SendItemsErrorOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**Sender** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**ActualKey** | **string** | A base64 URL encoded ed25519 public key | 
**ExpectedKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewSendItemsErrorOneOf1

`func NewSendItemsErrorOneOf1(type_ interface{}, title interface{}, status int32, detail string, sender int32, actualKey string, expectedKey string, ) *SendItemsErrorOneOf1`

NewSendItemsErrorOneOf1 instantiates a new SendItemsErrorOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendItemsErrorOneOf1WithDefaults

`func NewSendItemsErrorOneOf1WithDefaults() *SendItemsErrorOneOf1`

NewSendItemsErrorOneOf1WithDefaults instantiates a new SendItemsErrorOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendItemsErrorOneOf1) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendItemsErrorOneOf1) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendItemsErrorOneOf1) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SendItemsErrorOneOf1) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SendItemsErrorOneOf1) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SendItemsErrorOneOf1) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendItemsErrorOneOf1) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendItemsErrorOneOf1) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SendItemsErrorOneOf1) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SendItemsErrorOneOf1) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *SendItemsErrorOneOf1) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendItemsErrorOneOf1) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendItemsErrorOneOf1) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *SendItemsErrorOneOf1) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SendItemsErrorOneOf1) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SendItemsErrorOneOf1) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSender

`func (o *SendItemsErrorOneOf1) GetSender() int32`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SendItemsErrorOneOf1) GetSenderOk() (*int32, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SendItemsErrorOneOf1) SetSender(v int32)`

SetSender sets Sender field to given value.


### GetActualKey

`func (o *SendItemsErrorOneOf1) GetActualKey() string`

GetActualKey returns the ActualKey field if non-nil, zero value otherwise.

### GetActualKeyOk

`func (o *SendItemsErrorOneOf1) GetActualKeyOk() (*string, bool)`

GetActualKeyOk returns a tuple with the ActualKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKey

`func (o *SendItemsErrorOneOf1) SetActualKey(v string)`

SetActualKey sets ActualKey field to given value.


### GetExpectedKey

`func (o *SendItemsErrorOneOf1) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *SendItemsErrorOneOf1) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *SendItemsErrorOneOf1) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


