# InstanceUnreachableError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**Address** | **string** | An address pointing to another DFMailbox instance | 

## Methods

### NewInstanceUnreachableError

`func NewInstanceUnreachableError(type_ interface{}, title interface{}, status int32, detail string, address string, ) *InstanceUnreachableError`

NewInstanceUnreachableError instantiates a new InstanceUnreachableError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUnreachableErrorWithDefaults

`func NewInstanceUnreachableErrorWithDefaults() *InstanceUnreachableError`

NewInstanceUnreachableErrorWithDefaults instantiates a new InstanceUnreachableError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InstanceUnreachableError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceUnreachableError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceUnreachableError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InstanceUnreachableError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceUnreachableError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InstanceUnreachableError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InstanceUnreachableError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InstanceUnreachableError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InstanceUnreachableError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InstanceUnreachableError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InstanceUnreachableError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceUnreachableError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceUnreachableError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *InstanceUnreachableError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *InstanceUnreachableError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *InstanceUnreachableError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetAddress

`func (o *InstanceUnreachableError) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InstanceUnreachableError) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InstanceUnreachableError) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


