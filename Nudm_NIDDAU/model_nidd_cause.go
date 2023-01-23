/*
Nudm_NIDDAU

Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_NIDDAU

import (
	"encoding/json"
	"fmt"
)

// NiddCause struct for NiddCause
type NiddCause struct {
	NiddCauseAnyOf *NiddCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NiddCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NiddCauseAnyOf
	err = json.Unmarshal(data, &dst.NiddCauseAnyOf);
	if err == nil {
		jsonNiddCauseAnyOf, _ := json.Marshal(dst.NiddCauseAnyOf)
		if string(jsonNiddCauseAnyOf) == "{}" { // empty struct
			dst.NiddCauseAnyOf = nil
		} else {
			return nil // data stored in dst.NiddCauseAnyOf, return on the first match
		}
	} else {
		dst.NiddCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NiddCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NiddCause) MarshalJSON() ([]byte, error) {
	if src.NiddCauseAnyOf != nil {
		return json.Marshal(&src.NiddCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNiddCause struct {
	value *NiddCause
	isSet bool
}

func (v NullableNiddCause) Get() *NiddCause {
	return v.value
}

func (v *NullableNiddCause) Set(val *NiddCause) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddCause) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddCause(val *NiddCause) *NullableNiddCause {
	return &NullableNiddCause{value: val, isSet: true}
}

func (v NullableNiddCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


