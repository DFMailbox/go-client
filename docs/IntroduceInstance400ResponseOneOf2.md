# IntroduceInstance400ResponseOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **interface{}** |  | 
**ExpectedKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewIntroduceInstance400ResponseOneOf2

`func NewIntroduceInstance400ResponseOneOf2(error_ interface{}, expectedKey string, ) *IntroduceInstance400ResponseOneOf2`

NewIntroduceInstance400ResponseOneOf2 instantiates a new IntroduceInstance400ResponseOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntroduceInstance400ResponseOneOf2WithDefaults

`func NewIntroduceInstance400ResponseOneOf2WithDefaults() *IntroduceInstance400ResponseOneOf2`

NewIntroduceInstance400ResponseOneOf2WithDefaults instantiates a new IntroduceInstance400ResponseOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IntroduceInstance400ResponseOneOf2) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IntroduceInstance400ResponseOneOf2) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IntroduceInstance400ResponseOneOf2) SetError(v interface{})`

SetError sets Error field to given value.


### SetErrorNil

`func (o *IntroduceInstance400ResponseOneOf2) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *IntroduceInstance400ResponseOneOf2) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExpectedKey

`func (o *IntroduceInstance400ResponseOneOf2) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *IntroduceInstance400ResponseOneOf2) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *IntroduceInstance400ResponseOneOf2) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


