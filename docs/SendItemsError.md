# SendItemsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Details** | **string** |  | 
**PlotId** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

## Methods

### NewSendItemsError

`func NewSendItemsError(type_ interface{}, title interface{}, status int32, details string, plotId int32, ) *SendItemsError`

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


### GetDetails

`func (o *SendItemsError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SendItemsError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SendItemsError) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetPlotId

`func (o *SendItemsError) GetPlotId() int32`

GetPlotId returns the PlotId field if non-nil, zero value otherwise.

### GetPlotIdOk

`func (o *SendItemsError) GetPlotIdOk() (*int32, bool)`

GetPlotIdOk returns a tuple with the PlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotId

`func (o *SendItemsError) SetPlotId(v int32)`

SetPlotId sets PlotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


