/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// ReservPriority Indicates the reservation priority.
type ReservPriority struct {
	ReservPriorityAnyOf *ReservPriorityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReservPriority) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReservPriorityAnyOf
	err = json.Unmarshal(data, &dst.ReservPriorityAnyOf);
	if err == nil {
		jsonReservPriorityAnyOf, _ := json.Marshal(dst.ReservPriorityAnyOf)
		if string(jsonReservPriorityAnyOf) == "{}" { // empty struct
			dst.ReservPriorityAnyOf = nil
		} else {
			return nil // data stored in dst.ReservPriorityAnyOf, return on the first match
		}
	} else {
		dst.ReservPriorityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReservPriority)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReservPriority) MarshalJSON() ([]byte, error) {
	if src.ReservPriorityAnyOf != nil {
		return json.Marshal(&src.ReservPriorityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReservPriority struct {
	value *ReservPriority
	isSet bool
}

func (v NullableReservPriority) Get() *ReservPriority {
	return v.value
}

func (v *NullableReservPriority) Set(val *ReservPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableReservPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableReservPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservPriority(val *ReservPriority) *NullableReservPriority {
	return &NullableReservPriority{value: val, isSet: true}
}

func (v NullableReservPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


