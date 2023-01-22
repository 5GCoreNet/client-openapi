/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
	"fmt"
)

// SpecificProblem - struct for SpecificProblem
type SpecificProblem struct {
	Int32 *int32
	String *string
}

// int32AsSpecificProblem is a convenience function that returns int32 wrapped in SpecificProblem
func Int32AsSpecificProblem(v *int32) SpecificProblem {
	return SpecificProblem{
		Int32: v,
	}
}

// stringAsSpecificProblem is a convenience function that returns string wrapped in SpecificProblem
func StringAsSpecificProblem(v *string) SpecificProblem {
	return SpecificProblem{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpecificProblem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
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
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SpecificProblem)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SpecificProblem)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpecificProblem) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpecificProblem) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSpecificProblem struct {
	value *SpecificProblem
	isSet bool
}

func (v NullableSpecificProblem) Get() *SpecificProblem {
	return v.value
}

func (v *NullableSpecificProblem) Set(val *SpecificProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecificProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecificProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecificProblem(val *SpecificProblem) *NullableSpecificProblem {
	return &NullableSpecificProblem{value: val, isSet: true}
}

func (v NullableSpecificProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecificProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


