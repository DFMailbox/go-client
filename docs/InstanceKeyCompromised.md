# InstanceKeyCompromised

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Details** | **string** |  | 
**InstanceKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewInstanceKeyCompromised

`func NewInstanceKeyCompromised(type_ interface{}, title interface{}, status int32, details string, instanceKey string, ) *InstanceKeyCompromised`

NewInstanceKeyCompromised instantiates a new InstanceKeyCompromised object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceKeyCompromisedWithDefaults

`func NewInstanceKeyCompromisedWithDefaults() *InstanceKeyCompromised`

NewInstanceKeyCompromisedWithDefaults instantiates a new InstanceKeyCompromised object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InstanceKeyCompromised) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceKeyCompromised) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceKeyCompromised) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InstanceKeyCompromised) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceKeyCompromised) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InstanceKeyCompromised) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InstanceKeyCompromised) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InstanceKeyCompromised) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InstanceKeyCompromised) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InstanceKeyCompromised) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InstanceKeyCompromised) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceKeyCompromised) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceKeyCompromised) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *InstanceKeyCompromised) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InstanceKeyCompromised) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InstanceKeyCompromised) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetInstanceKey

`func (o *InstanceKeyCompromised) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *InstanceKeyCompromised) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *InstanceKeyCompromised) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


