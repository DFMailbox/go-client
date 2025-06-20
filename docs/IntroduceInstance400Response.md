# IntroduceInstance400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **interface{}** |  | 
**ErrorMessage** | **string** | A non-standard error message explaining in human terms | 
**ExpectedAddress** | **string** | An address pointing to another DFMailbox instance | 
**ExpectedKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewIntroduceInstance400Response

`func NewIntroduceInstance400Response(error_ interface{}, errorMessage string, expectedAddress string, expectedKey string, ) *IntroduceInstance400Response`

NewIntroduceInstance400Response instantiates a new IntroduceInstance400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntroduceInstance400ResponseWithDefaults

`func NewIntroduceInstance400ResponseWithDefaults() *IntroduceInstance400Response`

NewIntroduceInstance400ResponseWithDefaults instantiates a new IntroduceInstance400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IntroduceInstance400Response) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IntroduceInstance400Response) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IntroduceInstance400Response) SetError(v interface{})`

SetError sets Error field to given value.


### SetErrorNil

`func (o *IntroduceInstance400Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *IntroduceInstance400Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorMessage

`func (o *IntroduceInstance400Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *IntroduceInstance400Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *IntroduceInstance400Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetExpectedAddress

`func (o *IntroduceInstance400Response) GetExpectedAddress() string`

GetExpectedAddress returns the ExpectedAddress field if non-nil, zero value otherwise.

### GetExpectedAddressOk

`func (o *IntroduceInstance400Response) GetExpectedAddressOk() (*string, bool)`

GetExpectedAddressOk returns a tuple with the ExpectedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedAddress

`func (o *IntroduceInstance400Response) SetExpectedAddress(v string)`

SetExpectedAddress sets ExpectedAddress field to given value.


### GetExpectedKey

`func (o *IntroduceInstance400Response) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *IntroduceInstance400Response) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *IntroduceInstance400Response) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


