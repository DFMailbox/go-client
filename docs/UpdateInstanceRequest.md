# UpdateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | **NullableString** | An optional base64 URL encoded ed25519 public key | 

## Methods

### NewUpdateInstanceRequest

`func NewUpdateInstanceRequest(instance NullableString, ) *UpdateInstanceRequest`

NewUpdateInstanceRequest instantiates a new UpdateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceRequestWithDefaults

`func NewUpdateInstanceRequestWithDefaults() *UpdateInstanceRequest`

NewUpdateInstanceRequestWithDefaults instantiates a new UpdateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *UpdateInstanceRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateInstanceRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateInstanceRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.


### SetInstanceNil

`func (o *UpdateInstanceRequest) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *UpdateInstanceRequest) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


