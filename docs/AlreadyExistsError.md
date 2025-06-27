# AlreadyExistsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Detail** | **string** |  | 
**Status** | **int32** | HTTP status code | 

## Methods

### NewAlreadyExistsError

`func NewAlreadyExistsError(type_ interface{}, title interface{}, detail string, status int32, ) *AlreadyExistsError`

NewAlreadyExistsError instantiates a new AlreadyExistsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlreadyExistsErrorWithDefaults

`func NewAlreadyExistsErrorWithDefaults() *AlreadyExistsError`

NewAlreadyExistsErrorWithDefaults instantiates a new AlreadyExistsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlreadyExistsError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlreadyExistsError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlreadyExistsError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AlreadyExistsError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AlreadyExistsError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *AlreadyExistsError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlreadyExistsError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlreadyExistsError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *AlreadyExistsError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlreadyExistsError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDetail

`func (o *AlreadyExistsError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AlreadyExistsError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AlreadyExistsError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStatus

`func (o *AlreadyExistsError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlreadyExistsError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlreadyExistsError) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


