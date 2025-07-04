/*
DFMailbox API

DFMailbox is a decentralized way to send messages to other DiamondFire plots.

API version: 0.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnknownInstanceError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnknownInstanceError{}

// UnknownInstanceError Unknown instance error
type UnknownInstanceError struct {
	Type interface{} `json:"type"`
	Title interface{} `json:"title"`
	// HTTP status code
	Status int32 `json:"status"`
	Detail string `json:"detail"`
	// A base64 URL encoded ed25519 public key
	PublicKey string `json:"public_key"`
}

type _UnknownInstanceError UnknownInstanceError

// NewUnknownInstanceError instantiates a new UnknownInstanceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnknownInstanceError(type_ interface{}, title interface{}, status int32, detail string, publicKey string) *UnknownInstanceError {
	this := UnknownInstanceError{}
	this.Type = type_
	this.Title = title
	this.Status = status
	this.Detail = detail
	this.PublicKey = publicKey
	return &this
}

// NewUnknownInstanceErrorWithDefaults instantiates a new UnknownInstanceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnknownInstanceErrorWithDefaults() *UnknownInstanceError {
	this := UnknownInstanceError{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UnknownInstanceError) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnknownInstanceError) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UnknownInstanceError) SetType(v interface{}) {
	o.Type = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UnknownInstanceError) GetTitle() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnknownInstanceError) GetTitleOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UnknownInstanceError) SetTitle(v interface{}) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *UnknownInstanceError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UnknownInstanceError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UnknownInstanceError) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *UnknownInstanceError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *UnknownInstanceError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *UnknownInstanceError) SetDetail(v string) {
	o.Detail = v
}

// GetPublicKey returns the PublicKey field value
func (o *UnknownInstanceError) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *UnknownInstanceError) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *UnknownInstanceError) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o UnknownInstanceError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnknownInstanceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["status"] = o.Status
	toSerialize["detail"] = o.Detail
	toSerialize["public_key"] = o.PublicKey
	return toSerialize, nil
}

func (o *UnknownInstanceError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"title",
		"status",
		"detail",
		"public_key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnknownInstanceError := _UnknownInstanceError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnknownInstanceError)

	if err != nil {
		return err
	}

	*o = UnknownInstanceError(varUnknownInstanceError)

	return err
}

type NullableUnknownInstanceError struct {
	value *UnknownInstanceError
	isSet bool
}

func (v NullableUnknownInstanceError) Get() *UnknownInstanceError {
	return v.value
}

func (v *NullableUnknownInstanceError) Set(val *UnknownInstanceError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnknownInstanceError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnknownInstanceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnknownInstanceError(val *UnknownInstanceError) *NullableUnknownInstanceError {
	return &NullableUnknownInstanceError{value: val, isSet: true}
}

func (v NullableUnknownInstanceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnknownInstanceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


