/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// MatchingDirection Possible values are: - ASCENDING: Threshold is crossed in ascending direction. - DESCENDING: Threshold is crossed in descending direction. - CROSSED: Threshold is crossed either in ascending or descending direction. 
type MatchingDirection struct {
	MatchingDirectionAnyOf *MatchingDirectionAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MatchingDirection) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MatchingDirectionAnyOf
	err = json.Unmarshal(data, &dst.MatchingDirectionAnyOf);
	if err == nil {
		jsonMatchingDirectionAnyOf, _ := json.Marshal(dst.MatchingDirectionAnyOf)
		if string(jsonMatchingDirectionAnyOf) == "{}" { // empty struct
			dst.MatchingDirectionAnyOf = nil
		} else {
			return nil // data stored in dst.MatchingDirectionAnyOf, return on the first match
		}
	} else {
		dst.MatchingDirectionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MatchingDirection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MatchingDirection) MarshalJSON() ([]byte, error) {
	if src.MatchingDirectionAnyOf != nil {
		return json.Marshal(&src.MatchingDirectionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMatchingDirection struct {
	value *MatchingDirection
	isSet bool
}

func (v NullableMatchingDirection) Get() *MatchingDirection {
	return v.value
}

func (v *NullableMatchingDirection) Set(val *MatchingDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingDirection(val *MatchingDirection) *NullableMatchingDirection {
	return &NullableMatchingDirection{value: val, isSet: true}
}

func (v NullableMatchingDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


