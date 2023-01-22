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

// ReflectiveQoSAttribute The enumeration ReflectiveQosAttribute indicates whether certain traffic of the QoS flow may be subject to Reflective QoS (see clause 5.7.2.3 of 3GPP TS 23.501). It shall comply with the provisions defined in table 5.5.3.3-1.  
type ReflectiveQoSAttribute struct {
	ReflectiveQoSAttributeAnyOf *ReflectiveQoSAttributeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReflectiveQoSAttribute) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReflectiveQoSAttributeAnyOf
	err = json.Unmarshal(data, &dst.ReflectiveQoSAttributeAnyOf);
	if err == nil {
		jsonReflectiveQoSAttributeAnyOf, _ := json.Marshal(dst.ReflectiveQoSAttributeAnyOf)
		if string(jsonReflectiveQoSAttributeAnyOf) == "{}" { // empty struct
			dst.ReflectiveQoSAttributeAnyOf = nil
		} else {
			return nil // data stored in dst.ReflectiveQoSAttributeAnyOf, return on the first match
		}
	} else {
		dst.ReflectiveQoSAttributeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReflectiveQoSAttribute)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReflectiveQoSAttribute) MarshalJSON() ([]byte, error) {
	if src.ReflectiveQoSAttributeAnyOf != nil {
		return json.Marshal(&src.ReflectiveQoSAttributeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReflectiveQoSAttribute struct {
	value *ReflectiveQoSAttribute
	isSet bool
}

func (v NullableReflectiveQoSAttribute) Get() *ReflectiveQoSAttribute {
	return v.value
}

func (v *NullableReflectiveQoSAttribute) Set(val *ReflectiveQoSAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableReflectiveQoSAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableReflectiveQoSAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReflectiveQoSAttribute(val *ReflectiveQoSAttribute) *NullableReflectiveQoSAttribute {
	return &NullableReflectiveQoSAttribute{value: val, isSet: true}
}

func (v NullableReflectiveQoSAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReflectiveQoSAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


