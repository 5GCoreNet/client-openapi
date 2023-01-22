/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSSession

import (
	"encoding/json"
	"fmt"
)

// MbsSessionActivityStatus Indicates the MBS session's activity status
type MbsSessionActivityStatus struct {
	MbsSessionActivityStatusAnyOf *MbsSessionActivityStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsSessionActivityStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsSessionActivityStatusAnyOf
	err = json.Unmarshal(data, &dst.MbsSessionActivityStatusAnyOf);
	if err == nil {
		jsonMbsSessionActivityStatusAnyOf, _ := json.Marshal(dst.MbsSessionActivityStatusAnyOf)
		if string(jsonMbsSessionActivityStatusAnyOf) == "{}" { // empty struct
			dst.MbsSessionActivityStatusAnyOf = nil
		} else {
			return nil // data stored in dst.MbsSessionActivityStatusAnyOf, return on the first match
		}
	} else {
		dst.MbsSessionActivityStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsSessionActivityStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsSessionActivityStatus) MarshalJSON() ([]byte, error) {
	if src.MbsSessionActivityStatusAnyOf != nil {
		return json.Marshal(&src.MbsSessionActivityStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsSessionActivityStatus struct {
	value *MbsSessionActivityStatus
	isSet bool
}

func (v NullableMbsSessionActivityStatus) Get() *MbsSessionActivityStatus {
	return v.value
}

func (v *NullableMbsSessionActivityStatus) Set(val *MbsSessionActivityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionActivityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionActivityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionActivityStatus(val *MbsSessionActivityStatus) *NullableMbsSessionActivityStatus {
	return &NullableMbsSessionActivityStatus{value: val, isSet: true}
}

func (v NullableMbsSessionActivityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionActivityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


