# SendItemsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**InstanceKey** | **string** | A base64 URL encoded ed25519 public key | 
**ChallengeBytes** | Pointer to **string** | Base64 representation of the challenge ((ascii address bytes) + (uuid bytes)) | [optional] 
**Sender** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**ActualKey** | **string** | A base64 URL encoded ed25519 public key | 
**ExpectedKey** | **string** | A base64 URL encoded ed25519 public key | 
**Receiver** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

## Methods

### NewSendItemsError

`func NewSendItemsError(type_ interface{}, title interface{}, status int32, detail string, instanceKey string, sender int32, actualKey string, expectedKey string, receiver int32, ) *SendItemsError`

NewSendItemsError instantiates a new SendItemsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendItemsErrorWithDefaults

`func NewSendItemsErrorWithDefaults() *SendItemsError`

NewSendItemsErrorWithDefaults instantiates a new SendItemsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendItemsError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendItemsError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendItemsError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SendItemsError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SendItemsError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SendItemsError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendItemsError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendItemsError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SendItemsError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SendItemsError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *SendItemsError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendItemsError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendItemsError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *SendItemsError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SendItemsError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SendItemsError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstanceKey

`func (o *SendItemsError) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *SendItemsError) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *SendItemsError) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.


### GetChallengeBytes

`func (o *SendItemsError) GetChallengeBytes() string`

GetChallengeBytes returns the ChallengeBytes field if non-nil, zero value otherwise.

### GetChallengeBytesOk

`func (o *SendItemsError) GetChallengeBytesOk() (*string, bool)`

GetChallengeBytesOk returns a tuple with the ChallengeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeBytes

`func (o *SendItemsError) SetChallengeBytes(v string)`

SetChallengeBytes sets ChallengeBytes field to given value.

### HasChallengeBytes

`func (o *SendItemsError) HasChallengeBytes() bool`

HasChallengeBytes returns a boolean if a field has been set.

### GetSender

`func (o *SendItemsError) GetSender() int32`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *SendItemsError) GetSenderOk() (*int32, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *SendItemsError) SetSender(v int32)`

SetSender sets Sender field to given value.


### GetActualKey

`func (o *SendItemsError) GetActualKey() string`

GetActualKey returns the ActualKey field if non-nil, zero value otherwise.

### GetActualKeyOk

`func (o *SendItemsError) GetActualKeyOk() (*string, bool)`

GetActualKeyOk returns a tuple with the ActualKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKey

`func (o *SendItemsError) SetActualKey(v string)`

SetActualKey sets ActualKey field to given value.


### GetExpectedKey

`func (o *SendItemsError) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *SendItemsError) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *SendItemsError) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.


### GetReceiver

`func (o *SendItemsError) GetReceiver() int32`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SendItemsError) GetReceiverOk() (*int32, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SendItemsError) SetReceiver(v int32)`

SetReceiver sets Receiver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


