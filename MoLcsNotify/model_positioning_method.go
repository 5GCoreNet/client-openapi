/*
3gpp-mo-lcs-notify

API for UE updated location information notification.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MoLcsNotify

import (
	"encoding/json"
	"fmt"
)

// PositioningMethod Indicates supported positioning methods.
type PositioningMethod struct {
	PositioningMethodAnyOf *PositioningMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PositioningMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PositioningMethodAnyOf
	err = json.Unmarshal(data, &dst.PositioningMethodAnyOf);
	if err == nil {
		jsonPositioningMethodAnyOf, _ := json.Marshal(dst.PositioningMethodAnyOf)
		if string(jsonPositioningMethodAnyOf) == "{}" { // empty struct
			dst.PositioningMethodAnyOf = nil
		} else {
			return nil // data stored in dst.PositioningMethodAnyOf, return on the first match
		}
	} else {
		dst.PositioningMethodAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PositioningMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PositioningMethod) MarshalJSON() ([]byte, error) {
	if src.PositioningMethodAnyOf != nil {
		return json.Marshal(&src.PositioningMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePositioningMethod struct {
	value *PositioningMethod
	isSet bool
}

func (v NullablePositioningMethod) Get() *PositioningMethod {
	return v.value
}

func (v *NullablePositioningMethod) Set(val *PositioningMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethod(val *PositioningMethod) *NullablePositioningMethod {
	return &NullablePositioningMethod{value: val, isSet: true}
}

func (v NullablePositioningMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


