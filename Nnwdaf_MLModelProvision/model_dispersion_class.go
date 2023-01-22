/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// DispersionClass - Possible values are: - FIXED: Dispersion class as fixed UE its data or transaction usage at a location or a slice, is higher than its class threshold set for its all data or transaction usage. - CAMPER: Dispersion class as camper UE, its data or transaction usage at a location or a slice, is higher than its class threshold and lower than the fixed class threshold set for its all data or transaction usage.. - TRAVELLER: Dispersion class as traveller UE, its data or transaction usage at a location or a slice, is lower than the camper class threshold set for its all data or transaction usage. - TOP_HEAVY: Dispersion class as Top_Heavy UE, who's dispersion percentile rating at a location or a slice, is higher than its class threshold. 
type DispersionClass struct {
	DispersionClassOneOf *DispersionClassOneOf
	String *string
}

// DispersionClassOneOfAsDispersionClass is a convenience function that returns DispersionClassOneOf wrapped in DispersionClass
func DispersionClassOneOfAsDispersionClass(v *DispersionClassOneOf) DispersionClass {
	return DispersionClass{
		DispersionClassOneOf: v,
	}
}

// stringAsDispersionClass is a convenience function that returns string wrapped in DispersionClass
func StringAsDispersionClass(v *string) DispersionClass {
	return DispersionClass{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DispersionClass) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DispersionClassOneOf
	err = newStrictDecoder(data).Decode(&dst.DispersionClassOneOf)
	if err == nil {
		jsonDispersionClassOneOf, _ := json.Marshal(dst.DispersionClassOneOf)
		if string(jsonDispersionClassOneOf) == "{}" { // empty struct
			dst.DispersionClassOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DispersionClassOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DispersionClassOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DispersionClass)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DispersionClass)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DispersionClass) MarshalJSON() ([]byte, error) {
	if src.DispersionClassOneOf != nil {
		return json.Marshal(&src.DispersionClassOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DispersionClass) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DispersionClassOneOf != nil {
		return obj.DispersionClassOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDispersionClass struct {
	value *DispersionClass
	isSet bool
}

func (v NullableDispersionClass) Get() *DispersionClass {
	return v.value
}

func (v *NullableDispersionClass) Set(val *DispersionClass) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionClass) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionClass(val *DispersionClass) *NullableDispersionClass {
	return &NullableDispersionClass{value: val, isSet: true}
}

func (v NullableDispersionClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


