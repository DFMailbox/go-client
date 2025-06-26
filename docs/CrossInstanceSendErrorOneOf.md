# CrossInstanceSendErrorOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Details** | **string** |  | 
**Sender** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

## Methods

### NewCrossInstanceSendErrorOneOf

`func NewCrossInstanceSendErrorOneOf(type_ interface{}, title interface{}, status int32, details string, sender int32, ) *CrossInstanceSendErrorOneOf`

NewCrossInstanceSendErrorOneOf instantiates a new CrossInstanceSendErrorOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossInstanceSendErrorOneOfWithDefaults

`func NewCrossInstanceSendErrorOneOfWithDefaults() *CrossInstanceSendErrorOneOf`

NewCrossInstanceSendErrorOneOfWithDefaults instantiates a new CrossInstanceSendErrorOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossInstanceSendErrorOneOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossInstanceSendErrorOneOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossInstanceSendErrorOneOf) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CrossInstanceSendErrorOneOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CrossInstanceSendErrorOneOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *CrossInstanceSendErrorOneOf) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CrossInstanceSendErrorOneOf) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CrossInstanceSendErrorOneOf) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CrossInstanceSendErrorOneOf) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CrossInstanceSendErrorOneOf) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *CrossInstanceSendErrorOneOf) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossInstanceSendErrorOneOf) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossInstanceSendErrorOneOf) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *CrossInstanceSendErrorOneOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CrossInstanceSendErrorOneOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CrossInstanceSendErrorOneOf) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetSender

`func (o *CrossInstanceSendErrorOneOf) GetSender() int32`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CrossInstanceSendErrorOneOf) GetSenderOk() (*int32, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CrossInstanceSendErrorOneOf) SetSender(v int32)`

SetSender sets Sender field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


