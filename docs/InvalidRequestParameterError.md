# InvalidRequestParameterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**Errors** | **interface{}** |  | 

## Methods

### NewInvalidRequestParameterError

`func NewInvalidRequestParameterError(type_ interface{}, title interface{}, status int32, detail string, errors interface{}, ) *InvalidRequestParameterError`

NewInvalidRequestParameterError instantiates a new InvalidRequestParameterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidRequestParameterErrorWithDefaults

`func NewInvalidRequestParameterErrorWithDefaults() *InvalidRequestParameterError`

NewInvalidRequestParameterErrorWithDefaults instantiates a new InvalidRequestParameterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvalidRequestParameterError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvalidRequestParameterError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvalidRequestParameterError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InvalidRequestParameterError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InvalidRequestParameterError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InvalidRequestParameterError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvalidRequestParameterError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvalidRequestParameterError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InvalidRequestParameterError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvalidRequestParameterError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InvalidRequestParameterError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvalidRequestParameterError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvalidRequestParameterError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *InvalidRequestParameterError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *InvalidRequestParameterError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *InvalidRequestParameterError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *InvalidRequestParameterError) GetErrors() interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InvalidRequestParameterError) GetErrorsOk() (*interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InvalidRequestParameterError) SetErrors(v interface{})`

SetErrors sets Errors field to given value.


### SetErrorsNil

`func (o *InvalidRequestParameterError) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *InvalidRequestParameterError) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


