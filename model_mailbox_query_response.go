/*
DFMailbox API

DFMailbox is a decentralized way to send messages to other DiamondFire plots.

API version: 0.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// MailboxQueryResponse - A response for any request that queries a mailbox
type MailboxQueryResponse struct {
	MailboxQueryResponseOneOf *MailboxQueryResponseOneOf
	MailboxQueryResponseOneOf1 *MailboxQueryResponseOneOf1
	MailboxQueryResponseOneOf2 *MailboxQueryResponseOneOf2
	MailboxQueryResponseOneOf3 *MailboxQueryResponseOneOf3
}

// MailboxQueryResponseOneOfAsMailboxQueryResponse is a convenience function that returns MailboxQueryResponseOneOf wrapped in MailboxQueryResponse
func MailboxQueryResponseOneOfAsMailboxQueryResponse(v *MailboxQueryResponseOneOf) MailboxQueryResponse {
	return MailboxQueryResponse{
		MailboxQueryResponseOneOf: v,
	}
}

// MailboxQueryResponseOneOf1AsMailboxQueryResponse is a convenience function that returns MailboxQueryResponseOneOf1 wrapped in MailboxQueryResponse
func MailboxQueryResponseOneOf1AsMailboxQueryResponse(v *MailboxQueryResponseOneOf1) MailboxQueryResponse {
	return MailboxQueryResponse{
		MailboxQueryResponseOneOf1: v,
	}
}

// MailboxQueryResponseOneOf2AsMailboxQueryResponse is a convenience function that returns MailboxQueryResponseOneOf2 wrapped in MailboxQueryResponse
func MailboxQueryResponseOneOf2AsMailboxQueryResponse(v *MailboxQueryResponseOneOf2) MailboxQueryResponse {
	return MailboxQueryResponse{
		MailboxQueryResponseOneOf2: v,
	}
}

// MailboxQueryResponseOneOf3AsMailboxQueryResponse is a convenience function that returns MailboxQueryResponseOneOf3 wrapped in MailboxQueryResponse
func MailboxQueryResponseOneOf3AsMailboxQueryResponse(v *MailboxQueryResponseOneOf3) MailboxQueryResponse {
	return MailboxQueryResponse{
		MailboxQueryResponseOneOf3: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MailboxQueryResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MailboxQueryResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.MailboxQueryResponseOneOf)
	if err == nil {
		jsonMailboxQueryResponseOneOf, _ := json.Marshal(dst.MailboxQueryResponseOneOf)
		if string(jsonMailboxQueryResponseOneOf) == "{}" { // empty struct
			dst.MailboxQueryResponseOneOf = nil
		} else {
			if err = validator.Validate(dst.MailboxQueryResponseOneOf); err != nil {
				dst.MailboxQueryResponseOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.MailboxQueryResponseOneOf = nil
	}

	// try to unmarshal data into MailboxQueryResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.MailboxQueryResponseOneOf1)
	if err == nil {
		jsonMailboxQueryResponseOneOf1, _ := json.Marshal(dst.MailboxQueryResponseOneOf1)
		if string(jsonMailboxQueryResponseOneOf1) == "{}" { // empty struct
			dst.MailboxQueryResponseOneOf1 = nil
		} else {
			if err = validator.Validate(dst.MailboxQueryResponseOneOf1); err != nil {
				dst.MailboxQueryResponseOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.MailboxQueryResponseOneOf1 = nil
	}

	// try to unmarshal data into MailboxQueryResponseOneOf2
	err = newStrictDecoder(data).Decode(&dst.MailboxQueryResponseOneOf2)
	if err == nil {
		jsonMailboxQueryResponseOneOf2, _ := json.Marshal(dst.MailboxQueryResponseOneOf2)
		if string(jsonMailboxQueryResponseOneOf2) == "{}" { // empty struct
			dst.MailboxQueryResponseOneOf2 = nil
		} else {
			if err = validator.Validate(dst.MailboxQueryResponseOneOf2); err != nil {
				dst.MailboxQueryResponseOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.MailboxQueryResponseOneOf2 = nil
	}

	// try to unmarshal data into MailboxQueryResponseOneOf3
	err = newStrictDecoder(data).Decode(&dst.MailboxQueryResponseOneOf3)
	if err == nil {
		jsonMailboxQueryResponseOneOf3, _ := json.Marshal(dst.MailboxQueryResponseOneOf3)
		if string(jsonMailboxQueryResponseOneOf3) == "{}" { // empty struct
			dst.MailboxQueryResponseOneOf3 = nil
		} else {
			if err = validator.Validate(dst.MailboxQueryResponseOneOf3); err != nil {
				dst.MailboxQueryResponseOneOf3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.MailboxQueryResponseOneOf3 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MailboxQueryResponseOneOf = nil
		dst.MailboxQueryResponseOneOf1 = nil
		dst.MailboxQueryResponseOneOf2 = nil
		dst.MailboxQueryResponseOneOf3 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MailboxQueryResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MailboxQueryResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MailboxQueryResponse) MarshalJSON() ([]byte, error) {
	if src.MailboxQueryResponseOneOf != nil {
		return json.Marshal(&src.MailboxQueryResponseOneOf)
	}

	if src.MailboxQueryResponseOneOf1 != nil {
		return json.Marshal(&src.MailboxQueryResponseOneOf1)
	}

	if src.MailboxQueryResponseOneOf2 != nil {
		return json.Marshal(&src.MailboxQueryResponseOneOf2)
	}

	if src.MailboxQueryResponseOneOf3 != nil {
		return json.Marshal(&src.MailboxQueryResponseOneOf3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MailboxQueryResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MailboxQueryResponseOneOf != nil {
		return obj.MailboxQueryResponseOneOf
	}

	if obj.MailboxQueryResponseOneOf1 != nil {
		return obj.MailboxQueryResponseOneOf1
	}

	if obj.MailboxQueryResponseOneOf2 != nil {
		return obj.MailboxQueryResponseOneOf2
	}

	if obj.MailboxQueryResponseOneOf3 != nil {
		return obj.MailboxQueryResponseOneOf3
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MailboxQueryResponse) GetActualInstanceValue() (interface{}) {
	if obj.MailboxQueryResponseOneOf != nil {
		return *obj.MailboxQueryResponseOneOf
	}

	if obj.MailboxQueryResponseOneOf1 != nil {
		return *obj.MailboxQueryResponseOneOf1
	}

	if obj.MailboxQueryResponseOneOf2 != nil {
		return *obj.MailboxQueryResponseOneOf2
	}

	if obj.MailboxQueryResponseOneOf3 != nil {
		return *obj.MailboxQueryResponseOneOf3
	}

	// all schemas are nil
	return nil
}

type NullableMailboxQueryResponse struct {
	value *MailboxQueryResponse
	isSet bool
}

func (v NullableMailboxQueryResponse) Get() *MailboxQueryResponse {
	return v.value
}

func (v *NullableMailboxQueryResponse) Set(val *MailboxQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMailboxQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMailboxQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailboxQueryResponse(val *MailboxQueryResponse) *NullableMailboxQueryResponse {
	return &NullableMailboxQueryResponse{value: val, isSet: true}
}

func (v NullableMailboxQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailboxQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


