/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// AccessStateTransitionType Access State Transition Type.
type AccessStateTransitionType struct {
	AccessStateTransitionTypeAnyOf *AccessStateTransitionTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessStateTransitionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccessStateTransitionTypeAnyOf
	err = json.Unmarshal(data, &dst.AccessStateTransitionTypeAnyOf);
	if err == nil {
		jsonAccessStateTransitionTypeAnyOf, _ := json.Marshal(dst.AccessStateTransitionTypeAnyOf)
		if string(jsonAccessStateTransitionTypeAnyOf) == "{}" { // empty struct
			dst.AccessStateTransitionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AccessStateTransitionTypeAnyOf, return on the first match
		}
	} else {
		dst.AccessStateTransitionTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessStateTransitionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessStateTransitionType) MarshalJSON() ([]byte, error) {
	if src.AccessStateTransitionTypeAnyOf != nil {
		return json.Marshal(&src.AccessStateTransitionTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessStateTransitionType struct {
	value *AccessStateTransitionType
	isSet bool
}

func (v NullableAccessStateTransitionType) Get() *AccessStateTransitionType {
	return v.value
}

func (v *NullableAccessStateTransitionType) Set(val *AccessStateTransitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessStateTransitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessStateTransitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessStateTransitionType(val *AccessStateTransitionType) *NullableAccessStateTransitionType {
	return &NullableAccessStateTransitionType{value: val, isSet: true}
}

func (v NullableAccessStateTransitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessStateTransitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


