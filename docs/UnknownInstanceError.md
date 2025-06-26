# UnknownInstanceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Details** | **string** |  | 
**PublicKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewUnknownInstanceError

`func NewUnknownInstanceError(type_ interface{}, title interface{}, status int32, details string, publicKey string, ) *UnknownInstanceError`

NewUnknownInstanceError instantiates a new UnknownInstanceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnknownInstanceErrorWithDefaults

`func NewUnknownInstanceErrorWithDefaults() *UnknownInstanceError`

NewUnknownInstanceErrorWithDefaults instantiates a new UnknownInstanceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnknownInstanceError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnknownInstanceError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnknownInstanceError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *UnknownInstanceError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UnknownInstanceError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *UnknownInstanceError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UnknownInstanceError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UnknownInstanceError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *UnknownInstanceError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UnknownInstanceError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *UnknownInstanceError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnknownInstanceError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnknownInstanceError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *UnknownInstanceError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UnknownInstanceError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UnknownInstanceError) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetPublicKey

`func (o *UnknownInstanceError) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UnknownInstanceError) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UnknownInstanceError) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


