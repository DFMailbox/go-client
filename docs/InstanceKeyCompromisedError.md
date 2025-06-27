# InstanceKeyCompromisedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**InstanceKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewInstanceKeyCompromisedError

`func NewInstanceKeyCompromisedError(type_ interface{}, title interface{}, status int32, detail string, instanceKey string, ) *InstanceKeyCompromisedError`

NewInstanceKeyCompromisedError instantiates a new InstanceKeyCompromisedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceKeyCompromisedErrorWithDefaults

`func NewInstanceKeyCompromisedErrorWithDefaults() *InstanceKeyCompromisedError`

NewInstanceKeyCompromisedErrorWithDefaults instantiates a new InstanceKeyCompromisedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InstanceKeyCompromisedError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceKeyCompromisedError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceKeyCompromisedError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InstanceKeyCompromisedError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceKeyCompromisedError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InstanceKeyCompromisedError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InstanceKeyCompromisedError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InstanceKeyCompromisedError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InstanceKeyCompromisedError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InstanceKeyCompromisedError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InstanceKeyCompromisedError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceKeyCompromisedError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceKeyCompromisedError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *InstanceKeyCompromisedError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *InstanceKeyCompromisedError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *InstanceKeyCompromisedError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstanceKey

`func (o *InstanceKeyCompromisedError) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *InstanceKeyCompromisedError) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *InstanceKeyCompromisedError) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


