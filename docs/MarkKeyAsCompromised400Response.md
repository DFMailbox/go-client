# MarkKeyAsCompromised400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Title** | **interface{}** |  | 
**Status** | **int32** | HTTP status code | 
**OffendingChallenge** | **string** | A UUID (universally unique identifier) | 
**Challenge** | **string** | Base64 representation of the challenge | 

## Methods

### NewMarkKeyAsCompromised400Response

`func NewMarkKeyAsCompromised400Response(type_ interface{}, title interface{}, status int32, offendingChallenge string, challenge string, ) *MarkKeyAsCompromised400Response`

NewMarkKeyAsCompromised400Response instantiates a new MarkKeyAsCompromised400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkKeyAsCompromised400ResponseWithDefaults

`func NewMarkKeyAsCompromised400ResponseWithDefaults() *MarkKeyAsCompromised400Response`

NewMarkKeyAsCompromised400ResponseWithDefaults instantiates a new MarkKeyAsCompromised400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarkKeyAsCompromised400Response) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarkKeyAsCompromised400Response) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarkKeyAsCompromised400Response) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MarkKeyAsCompromised400Response) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MarkKeyAsCompromised400Response) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *MarkKeyAsCompromised400Response) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarkKeyAsCompromised400Response) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarkKeyAsCompromised400Response) SetTitle(v interface{})`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *MarkKeyAsCompromised400Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MarkKeyAsCompromised400Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *MarkKeyAsCompromised400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarkKeyAsCompromised400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarkKeyAsCompromised400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetOffendingChallenge

`func (o *MarkKeyAsCompromised400Response) GetOffendingChallenge() string`

GetOffendingChallenge returns the OffendingChallenge field if non-nil, zero value otherwise.

### GetOffendingChallengeOk

`func (o *MarkKeyAsCompromised400Response) GetOffendingChallengeOk() (*string, bool)`

GetOffendingChallengeOk returns a tuple with the OffendingChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffendingChallenge

`func (o *MarkKeyAsCompromised400Response) SetOffendingChallenge(v string)`

SetOffendingChallenge sets OffendingChallenge field to given value.


### GetChallenge

`func (o *MarkKeyAsCompromised400Response) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *MarkKeyAsCompromised400Response) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *MarkKeyAsCompromised400Response) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


