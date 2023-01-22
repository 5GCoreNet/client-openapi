/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// TerminationCause Specifies causes of event reporting termination.
type TerminationCause struct {
	TerminationCauseAnyOf *TerminationCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TerminationCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TerminationCauseAnyOf
	err = json.Unmarshal(data, &dst.TerminationCauseAnyOf);
	if err == nil {
		jsonTerminationCauseAnyOf, _ := json.Marshal(dst.TerminationCauseAnyOf)
		if string(jsonTerminationCauseAnyOf) == "{}" { // empty struct
			dst.TerminationCauseAnyOf = nil
		} else {
			return nil // data stored in dst.TerminationCauseAnyOf, return on the first match
		}
	} else {
		dst.TerminationCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TerminationCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TerminationCause) MarshalJSON() ([]byte, error) {
	if src.TerminationCauseAnyOf != nil {
		return json.Marshal(&src.TerminationCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTerminationCause struct {
	value *TerminationCause
	isSet bool
}

func (v NullableTerminationCause) Get() *TerminationCause {
	return v.value
}

func (v *NullableTerminationCause) Set(val *TerminationCause) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationCause) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationCause(val *TerminationCause) *NullableTerminationCause {
	return &NullableTerminationCause{value: val, isSet: true}
}

func (v NullableTerminationCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


