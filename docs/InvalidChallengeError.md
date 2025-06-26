# InvalidChallengeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**OffendingChallenge** | **string** | A UUID (universally unique identifier) | 

## Methods

### NewInvalidChallengeError

`func NewInvalidChallengeError(type_ interface{}, title interface{}, status int32, offendingChallenge string, ) *InvalidChallengeError`

NewInvalidChallengeError instantiates a new InvalidChallengeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidChallengeErrorWithDefaults

`func NewInvalidChallengeErrorWithDefaults() *InvalidChallengeError`

NewInvalidChallengeErrorWithDefaults instantiates a new InvalidChallengeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvalidChallengeError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvalidChallengeError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvalidChallengeError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InvalidChallengeError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InvalidChallengeError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *InvalidChallengeError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvalidChallengeError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvalidChallengeError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *InvalidChallengeError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvalidChallengeError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *InvalidChallengeError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvalidChallengeError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvalidChallengeError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetOffendingChallenge

`func (o *InvalidChallengeError) GetOffendingChallenge() string`

GetOffendingChallenge returns the OffendingChallenge field if non-nil, zero value otherwise.

### GetOffendingChallengeOk

`func (o *InvalidChallengeError) GetOffendingChallengeOk() (*string, bool)`

GetOffendingChallengeOk returns a tuple with the OffendingChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffendingChallenge

`func (o *InvalidChallengeError) SetOffendingChallenge(v string)`

SetOffendingChallenge sets OffendingChallenge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


