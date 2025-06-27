# LookupInstanceAddress200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An address pointing to another DFMailbox instance | 
**Addresses** | **interface{}** |  | 

## Methods

### NewLookupInstanceAddress200Response

`func NewLookupInstanceAddress200Response(address string, addresses interface{}, ) *LookupInstanceAddress200Response`

NewLookupInstanceAddress200Response instantiates a new LookupInstanceAddress200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupInstanceAddress200ResponseWithDefaults

`func NewLookupInstanceAddress200ResponseWithDefaults() *LookupInstanceAddress200Response`

NewLookupInstanceAddress200ResponseWithDefaults instantiates a new LookupInstanceAddress200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LookupInstanceAddress200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LookupInstanceAddress200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LookupInstanceAddress200Response) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddresses

`func (o *LookupInstanceAddress200Response) GetAddresses() interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *LookupInstanceAddress200Response) GetAddressesOk() (*interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *LookupInstanceAddress200Response) SetAddresses(v interface{})`

SetAddresses sets Addresses field to given value.


### SetAddressesNil

`func (o *LookupInstanceAddress200Response) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *LookupInstanceAddress200Response) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


