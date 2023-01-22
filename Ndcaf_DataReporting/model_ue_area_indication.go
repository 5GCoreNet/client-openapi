/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndcaf_DataReporting

import (
	"encoding/json"
	"fmt"
)

// UeAreaIndication - Indicates area (country or international area) where UE is located
type UeAreaIndication struct {
	Interface{} *interface{}
}

// interface{}AsUeAreaIndication is a convenience function that returns interface{} wrapped in UeAreaIndication
func Interface{}AsUeAreaIndication(v *interface{}) UeAreaIndication {
	return UeAreaIndication{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UeAreaIndication) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UeAreaIndication)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UeAreaIndication)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UeAreaIndication) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UeAreaIndication) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableUeAreaIndication struct {
	value *UeAreaIndication
	isSet bool
}

func (v NullableUeAreaIndication) Get() *UeAreaIndication {
	return v.value
}

func (v *NullableUeAreaIndication) Set(val *UeAreaIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableUeAreaIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableUeAreaIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeAreaIndication(val *UeAreaIndication) *NullableUeAreaIndication {
	return &NullableUeAreaIndication{value: val, isSet: true}
}

func (v NullableUeAreaIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeAreaIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

