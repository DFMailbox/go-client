# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**ExpectedAddress** | **string** | An address pointing to another DFMailbox instance | 
**ExpectedKey** | Pointer to **string** | A base64 URL encoded ed25519 public key | [optional] 

## Methods

### NewInlineObject5

`func NewInlineObject5(type_ interface{}, title interface{}, status int32, detail string, expectedAddress string, ) *InlineObject5`

NewInlineObject5 instantiates a new InlineObject5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject5WithDefaults

`func NewInlineObject5WithDefaults() *InlineObject5`

NewInlineObject5WithDefaults instantiates a new InlineObject5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject5) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject5) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject5) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InlineObject5) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InlineObject5) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InlineObject5) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineObject5) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineObject5) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InlineObject5) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InlineObject5) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InlineObject5) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject5) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject5) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *InlineObject5) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *InlineObject5) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *InlineObject5) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetExpectedAddress

`func (o *InlineObject5) GetExpectedAddress() string`

GetExpectedAddress returns the ExpectedAddress field if non-nil, zero value otherwise.

### GetExpectedAddressOk

`func (o *InlineObject5) GetExpectedAddressOk() (*string, bool)`

GetExpectedAddressOk returns a tuple with the ExpectedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedAddress

`func (o *InlineObject5) SetExpectedAddress(v string)`

SetExpectedAddress sets ExpectedAddress field to given value.


### GetExpectedKey

`func (o *InlineObject5) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *InlineObject5) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *InlineObject5) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.

### HasExpectedKey

`func (o *InlineObject5) HasExpectedKey() bool`

HasExpectedKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


