/*
DFMailbox API

DFMailbox is a decentralized way to send messages to other DiamondFire plots.

API version: 0.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetChallenge201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChallenge201Response{}

// GetChallenge201Response struct for GetChallenge201Response
type GetChallenge201Response struct {
	// A UUID (universally unique identifier)
	Challenge *string `json:"challenge,omitempty"`
}

// NewGetChallenge201Response instantiates a new GetChallenge201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChallenge201Response() *GetChallenge201Response {
	this := GetChallenge201Response{}
	return &this
}

// NewGetChallenge201ResponseWithDefaults instantiates a new GetChallenge201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChallenge201ResponseWithDefaults() *GetChallenge201Response {
	this := GetChallenge201Response{}
	return &this
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *GetChallenge201Response) GetChallenge() string {
	if o == nil || IsNil(o.Challenge) {
		var ret string
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChallenge201Response) GetChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *GetChallenge201Response) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given string and assigns it to the Challenge field.
func (o *GetChallenge201Response) SetChallenge(v string) {
	o.Challenge = &v
}

func (o GetChallenge201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChallenge201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	return toSerialize, nil
}

type NullableGetChallenge201Response struct {
	value *GetChallenge201Response
	isSet bool
}

func (v NullableGetChallenge201Response) Get() *GetChallenge201Response {
	return v.value
}

func (v *NullableGetChallenge201Response) Set(val *GetChallenge201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChallenge201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChallenge201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChallenge201Response(val *GetChallenge201Response) *NullableGetChallenge201Response {
	return &NullableGetChallenge201Response{value: val, isSet: true}
}

func (v NullableGetChallenge201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChallenge201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


