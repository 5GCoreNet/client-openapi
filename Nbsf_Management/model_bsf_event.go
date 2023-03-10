/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"fmt"
)

// BsfEvent Represents an event to be notified by the BSF.
type BsfEvent struct {
	BsfEventAnyOf *BsfEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BsfEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BsfEventAnyOf
	err = json.Unmarshal(data, &dst.BsfEventAnyOf);
	if err == nil {
		jsonBsfEventAnyOf, _ := json.Marshal(dst.BsfEventAnyOf)
		if string(jsonBsfEventAnyOf) == "{}" { // empty struct
			dst.BsfEventAnyOf = nil
		} else {
			return nil // data stored in dst.BsfEventAnyOf, return on the first match
		}
	} else {
		dst.BsfEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(BsfEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BsfEvent) MarshalJSON() ([]byte, error) {
	if src.BsfEventAnyOf != nil {
		return json.Marshal(&src.BsfEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBsfEvent struct {
	value *BsfEvent
	isSet bool
}

func (v NullableBsfEvent) Get() *BsfEvent {
	return v.value
}

func (v *NullableBsfEvent) Set(val *BsfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfEvent(val *BsfEvent) *NullableBsfEvent {
	return &NullableBsfEvent{value: val, isSet: true}
}

func (v NullableBsfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


