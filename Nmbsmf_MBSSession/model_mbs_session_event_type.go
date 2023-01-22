/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// MbsSessionEventType MBS Session Event Type
type MbsSessionEventType struct {
	MbsSessionEventTypeAnyOf *MbsSessionEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsSessionEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsSessionEventTypeAnyOf
	err = json.Unmarshal(data, &dst.MbsSessionEventTypeAnyOf);
	if err == nil {
		jsonMbsSessionEventTypeAnyOf, _ := json.Marshal(dst.MbsSessionEventTypeAnyOf)
		if string(jsonMbsSessionEventTypeAnyOf) == "{}" { // empty struct
			dst.MbsSessionEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.MbsSessionEventTypeAnyOf, return on the first match
		}
	} else {
		dst.MbsSessionEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsSessionEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsSessionEventType) MarshalJSON() ([]byte, error) {
	if src.MbsSessionEventTypeAnyOf != nil {
		return json.Marshal(&src.MbsSessionEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsSessionEventType struct {
	value *MbsSessionEventType
	isSet bool
}

func (v NullableMbsSessionEventType) Get() *MbsSessionEventType {
	return v.value
}

func (v *NullableMbsSessionEventType) Set(val *MbsSessionEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionEventType(val *MbsSessionEventType) *NullableMbsSessionEventType {
	return &NullableMbsSessionEventType{value: val, isSet: true}
}

func (v NullableMbsSessionEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


