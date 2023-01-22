/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// HoState Handover state. Possible values are - NONE - PREPARING - PREPARED - COMPLETED - CANCELLED 
type HoState struct {
	HoStateAnyOf *HoStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *HoState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into HoStateAnyOf
	err = json.Unmarshal(data, &dst.HoStateAnyOf);
	if err == nil {
		jsonHoStateAnyOf, _ := json.Marshal(dst.HoStateAnyOf)
		if string(jsonHoStateAnyOf) == "{}" { // empty struct
			dst.HoStateAnyOf = nil
		} else {
			return nil // data stored in dst.HoStateAnyOf, return on the first match
		}
	} else {
		dst.HoStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(HoState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *HoState) MarshalJSON() ([]byte, error) {
	if src.HoStateAnyOf != nil {
		return json.Marshal(&src.HoStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableHoState struct {
	value *HoState
	isSet bool
}

func (v NullableHoState) Get() *HoState {
	return v.value
}

func (v *NullableHoState) Set(val *HoState) {
	v.value = val
	v.isSet = true
}

func (v NullableHoState) IsSet() bool {
	return v.isSet
}

func (v *NullableHoState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoState(val *HoState) *NullableHoState {
	return &NullableHoState{value: val, isSet: true}
}

func (v NullableHoState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


