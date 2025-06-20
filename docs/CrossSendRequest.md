# CrossSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**To** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 
**Data** | **interface{}** |  | 

## Methods

### NewCrossSendRequest

`func NewCrossSendRequest(from int32, to int32, data interface{}, ) *CrossSendRequest`

NewCrossSendRequest instantiates a new CrossSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossSendRequestWithDefaults

`func NewCrossSendRequestWithDefaults() *CrossSendRequest`

NewCrossSendRequestWithDefaults instantiates a new CrossSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CrossSendRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CrossSendRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CrossSendRequest) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *CrossSendRequest) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CrossSendRequest) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CrossSendRequest) SetTo(v int32)`

SetTo sets To field to given value.


### GetData

`func (o *CrossSendRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CrossSendRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CrossSendRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *CrossSendRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CrossSendRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


