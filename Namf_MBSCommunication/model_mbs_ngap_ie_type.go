/*
Namf_MBSCommunication

AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSCommunication

import (
	"encoding/json"
	"fmt"
)

// MbsNgapIeType NGAP Information Element Type
type MbsNgapIeType struct {
	MbsNgapIeTypeAnyOf *MbsNgapIeTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsNgapIeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsNgapIeTypeAnyOf
	err = json.Unmarshal(data, &dst.MbsNgapIeTypeAnyOf);
	if err == nil {
		jsonMbsNgapIeTypeAnyOf, _ := json.Marshal(dst.MbsNgapIeTypeAnyOf)
		if string(jsonMbsNgapIeTypeAnyOf) == "{}" { // empty struct
			dst.MbsNgapIeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.MbsNgapIeTypeAnyOf, return on the first match
		}
	} else {
		dst.MbsNgapIeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsNgapIeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsNgapIeType) MarshalJSON() ([]byte, error) {
	if src.MbsNgapIeTypeAnyOf != nil {
		return json.Marshal(&src.MbsNgapIeTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsNgapIeType struct {
	value *MbsNgapIeType
	isSet bool
}

func (v NullableMbsNgapIeType) Get() *MbsNgapIeType {
	return v.value
}

func (v *NullableMbsNgapIeType) Set(val *MbsNgapIeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsNgapIeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsNgapIeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsNgapIeType(val *MbsNgapIeType) *NullableMbsNgapIeType {
	return &NullableMbsNgapIeType{value: val, isSet: true}
}

func (v NullableMbsNgapIeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsNgapIeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


