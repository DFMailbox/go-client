# ChallengeFailedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**ChallengeBytes** | Pointer to **string** | Base64 representation of the challenge ((ascii address bytes) + (uuid bytes)) | [optional] 

## Methods

### NewChallengeFailedError

`func NewChallengeFailedError(type_ interface{}, title interface{}, status int32, ) *ChallengeFailedError`

NewChallengeFailedError instantiates a new ChallengeFailedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeFailedErrorWithDefaults

`func NewChallengeFailedErrorWithDefaults() *ChallengeFailedError`

NewChallengeFailedErrorWithDefaults instantiates a new ChallengeFailedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeFailedError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeFailedError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeFailedError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ChallengeFailedError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChallengeFailedError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ChallengeFailedError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChallengeFailedError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChallengeFailedError) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ChallengeFailedError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ChallengeFailedError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *ChallengeFailedError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeFailedError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeFailedError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetChallengeBytes

`func (o *ChallengeFailedError) GetChallengeBytes() string`

GetChallengeBytes returns the ChallengeBytes field if non-nil, zero value otherwise.

### GetChallengeBytesOk

`func (o *ChallengeFailedError) GetChallengeBytesOk() (*string, bool)`

GetChallengeBytesOk returns a tuple with the ChallengeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeBytes

`func (o *ChallengeFailedError) SetChallengeBytes(v string)`

SetChallengeBytes sets ChallengeBytes field to given value.

### HasChallengeBytes

`func (o *ChallengeFailedError) HasChallengeBytes() bool`

HasChallengeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


