# UpdateInstance409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**Detail** | **string** |  | 
**PublicKey** | **string** | A base64 URL encoded ed25519 public key | 
**InstanceKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewUpdateInstance409Response

`func NewUpdateInstance409Response(type_ interface{}, title interface{}, status int32, detail string, publicKey string, instanceKey string, ) *UpdateInstance409Response`

NewUpdateInstance409Response instantiates a new UpdateInstance409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstance409ResponseWithDefaults

`func NewUpdateInstance409ResponseWithDefaults() *UpdateInstance409Response`

NewUpdateInstance409ResponseWithDefaults instantiates a new UpdateInstance409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateInstance409Response) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateInstance409Response) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateInstance409Response) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *UpdateInstance409Response) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateInstance409Response) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *UpdateInstance409Response) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateInstance409Response) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateInstance409Response) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *UpdateInstance409Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateInstance409Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *UpdateInstance409Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateInstance409Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateInstance409Response) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *UpdateInstance409Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UpdateInstance409Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UpdateInstance409Response) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetPublicKey

`func (o *UpdateInstance409Response) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UpdateInstance409Response) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UpdateInstance409Response) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetInstanceKey

`func (o *UpdateInstance409Response) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *UpdateInstance409Response) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *UpdateInstance409Response) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


