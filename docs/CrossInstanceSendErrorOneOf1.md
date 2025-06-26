# CrossInstanceSendErrorOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Details** | **string** |  | 
**Sender** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**ActualKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewCrossInstanceSendErrorOneOf1

`func NewCrossInstanceSendErrorOneOf1(type_ interface{}, title interface{}, status int32, details string, sender int32, actualKey string, ) *CrossInstanceSendErrorOneOf1`

NewCrossInstanceSendErrorOneOf1 instantiates a new CrossInstanceSendErrorOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossInstanceSendErrorOneOf1WithDefaults

`func NewCrossInstanceSendErrorOneOf1WithDefaults() *CrossInstanceSendErrorOneOf1`

NewCrossInstanceSendErrorOneOf1WithDefaults instantiates a new CrossInstanceSendErrorOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossInstanceSendErrorOneOf1) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossInstanceSendErrorOneOf1) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossInstanceSendErrorOneOf1) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CrossInstanceSendErrorOneOf1) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CrossInstanceSendErrorOneOf1) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *CrossInstanceSendErrorOneOf1) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CrossInstanceSendErrorOneOf1) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CrossInstanceSendErrorOneOf1) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CrossInstanceSendErrorOneOf1) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CrossInstanceSendErrorOneOf1) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *CrossInstanceSendErrorOneOf1) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossInstanceSendErrorOneOf1) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossInstanceSendErrorOneOf1) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *CrossInstanceSendErrorOneOf1) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CrossInstanceSendErrorOneOf1) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CrossInstanceSendErrorOneOf1) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetSender

`func (o *CrossInstanceSendErrorOneOf1) GetSender() int32`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CrossInstanceSendErrorOneOf1) GetSenderOk() (*int32, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CrossInstanceSendErrorOneOf1) SetSender(v int32)`

SetSender sets Sender field to given value.


### GetActualKey

`func (o *CrossInstanceSendErrorOneOf1) GetActualKey() string`

GetActualKey returns the ActualKey field if non-nil, zero value otherwise.

### GetActualKeyOk

`func (o *CrossInstanceSendErrorOneOf1) GetActualKeyOk() (*string, bool)`

GetActualKeyOk returns a tuple with the ActualKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKey

`func (o *CrossInstanceSendErrorOneOf1) SetActualKey(v string)`

SetActualKey sets ActualKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


