# CrossInstanceSendError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **interface{}** |  | 
**ActualKey** | **string** | A base64 URL encoded ed25519 public key | 

## Methods

### NewCrossInstanceSendError

`func NewCrossInstanceSendError(error_ interface{}, actualKey string, ) *CrossInstanceSendError`

NewCrossInstanceSendError instantiates a new CrossInstanceSendError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossInstanceSendErrorWithDefaults

`func NewCrossInstanceSendErrorWithDefaults() *CrossInstanceSendError`

NewCrossInstanceSendErrorWithDefaults instantiates a new CrossInstanceSendError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CrossInstanceSendError) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CrossInstanceSendError) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CrossInstanceSendError) SetError(v interface{})`

SetError sets Error field to given value.


### SetErrorNil

`func (o *CrossInstanceSendError) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CrossInstanceSendError) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetActualKey

`func (o *CrossInstanceSendError) GetActualKey() string`

GetActualKey returns the ActualKey field if non-nil, zero value otherwise.

### GetActualKeyOk

`func (o *CrossInstanceSendError) GetActualKeyOk() (*string, bool)`

GetActualKeyOk returns a tuple with the ActualKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKey

`func (o *CrossInstanceSendError) SetActualKey(v string)`

SetActualKey sets ActualKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


