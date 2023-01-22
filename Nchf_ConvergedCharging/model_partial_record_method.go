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

// PartialRecordMethod struct for PartialRecordMethod
type PartialRecordMethod struct {
	PartialRecordMethodAnyOf *PartialRecordMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PartialRecordMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PartialRecordMethodAnyOf
	err = json.Unmarshal(data, &dst.PartialRecordMethodAnyOf);
	if err == nil {
		jsonPartialRecordMethodAnyOf, _ := json.Marshal(dst.PartialRecordMethodAnyOf)
		if string(jsonPartialRecordMethodAnyOf) == "{}" { // empty struct
			dst.PartialRecordMethodAnyOf = nil
		} else {
			return nil // data stored in dst.PartialRecordMethodAnyOf, return on the first match
		}
	} else {
		dst.PartialRecordMethodAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PartialRecordMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PartialRecordMethod) MarshalJSON() ([]byte, error) {
	if src.PartialRecordMethodAnyOf != nil {
		return json.Marshal(&src.PartialRecordMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePartialRecordMethod struct {
	value *PartialRecordMethod
	isSet bool
}

func (v NullablePartialRecordMethod) Get() *PartialRecordMethod {
	return v.value
}

func (v *NullablePartialRecordMethod) Set(val *PartialRecordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialRecordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialRecordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialRecordMethod(val *PartialRecordMethod) *NullablePartialRecordMethod {
	return &NullablePartialRecordMethod{value: val, isSet: true}
}

func (v NullablePartialRecordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialRecordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


