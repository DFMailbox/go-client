# UnregisteredPlotError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**PlotId** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

## Methods

### NewUnregisteredPlotError

`func NewUnregisteredPlotError(type_ interface{}, title interface{}, status int32, detail string, plotId int32, ) *UnregisteredPlotError`

NewUnregisteredPlotError instantiates a new UnregisteredPlotError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisteredPlotErrorWithDefaults

`func NewUnregisteredPlotErrorWithDefaults() *UnregisteredPlotError`

NewUnregisteredPlotErrorWithDefaults instantiates a new UnregisteredPlotError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnregisteredPlotError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnregisteredPlotError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnregisteredPlotError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *UnregisteredPlotError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UnregisteredPlotError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *UnregisteredPlotError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UnregisteredPlotError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UnregisteredPlotError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *UnregisteredPlotError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UnregisteredPlotError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *UnregisteredPlotError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnregisteredPlotError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnregisteredPlotError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *UnregisteredPlotError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UnregisteredPlotError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UnregisteredPlotError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetPlotId

`func (o *UnregisteredPlotError) GetPlotId() int32`

GetPlotId returns the PlotId field if non-nil, zero value otherwise.

### GetPlotIdOk

`func (o *UnregisteredPlotError) GetPlotIdOk() (*int32, bool)`

GetPlotIdOk returns a tuple with the PlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotId

`func (o *UnregisteredPlotError) SetPlotId(v int32)`

SetPlotId sets PlotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


