# LookupInstanceAddress200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**NullableAddressKeyPair**](NullableAddressKeyPair.md) |  | 
**Instances** | **interface{}** |  | 

## Methods

### NewLookupInstanceAddress200Response

`func NewLookupInstanceAddress200Response(instance NullableAddressKeyPair, instances interface{}, ) *LookupInstanceAddress200Response`

NewLookupInstanceAddress200Response instantiates a new LookupInstanceAddress200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupInstanceAddress200ResponseWithDefaults

`func NewLookupInstanceAddress200ResponseWithDefaults() *LookupInstanceAddress200Response`

NewLookupInstanceAddress200ResponseWithDefaults instantiates a new LookupInstanceAddress200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *LookupInstanceAddress200Response) GetInstance() NullableAddressKeyPair`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *LookupInstanceAddress200Response) GetInstanceOk() (*NullableAddressKeyPair, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *LookupInstanceAddress200Response) SetInstance(v NullableAddressKeyPair)`

SetInstance sets Instance field to given value.


### GetInstances

`func (o *LookupInstanceAddress200Response) GetInstances() interface{}`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *LookupInstanceAddress200Response) GetInstancesOk() (*interface{}, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *LookupInstanceAddress200Response) SetInstances(v interface{})`

SetInstances sets Instances field to given value.


### SetInstancesNil

`func (o *LookupInstanceAddress200Response) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *LookupInstanceAddress200Response) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


