/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// ThresholdInfoThresholdValue - struct for ThresholdInfoThresholdValue
type ThresholdInfoThresholdValue struct {
	Float32 *float32
	Int32 *int32
}

// float32AsThresholdInfoThresholdValue is a convenience function that returns float32 wrapped in ThresholdInfoThresholdValue
func Float32AsThresholdInfoThresholdValue(v *float32) ThresholdInfoThresholdValue {
	return ThresholdInfoThresholdValue{
		Float32: v,
	}
}

// int32AsThresholdInfoThresholdValue is a convenience function that returns int32 wrapped in ThresholdInfoThresholdValue
func Int32AsThresholdInfoThresholdValue(v *int32) ThresholdInfoThresholdValue {
	return ThresholdInfoThresholdValue{
		Int32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThresholdInfoThresholdValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ThresholdInfoThresholdValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThresholdInfoThresholdValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThresholdInfoThresholdValue) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThresholdInfoThresholdValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableThresholdInfoThresholdValue struct {
	value *ThresholdInfoThresholdValue
	isSet bool
}

func (v NullableThresholdInfoThresholdValue) Get() *ThresholdInfoThresholdValue {
	return v.value
}

func (v *NullableThresholdInfoThresholdValue) Set(val *ThresholdInfoThresholdValue) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdInfoThresholdValue) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdInfoThresholdValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdInfoThresholdValue(val *ThresholdInfoThresholdValue) *NullableThresholdInfoThresholdValue {
	return &NullableThresholdInfoThresholdValue{value: val, isSet: true}
}

func (v NullableThresholdInfoThresholdValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdInfoThresholdValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


