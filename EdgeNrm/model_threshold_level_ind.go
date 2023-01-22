/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EdgeNrm

import (
	"encoding/json"
	"fmt"
)

// ThresholdLevelInd - struct for ThresholdLevelInd
type ThresholdLevelInd struct {
	ThresholdLevelIndOneOf *ThresholdLevelIndOneOf
	ThresholdLevelIndOneOf1 *ThresholdLevelIndOneOf1
}

// ThresholdLevelIndOneOfAsThresholdLevelInd is a convenience function that returns ThresholdLevelIndOneOf wrapped in ThresholdLevelInd
func ThresholdLevelIndOneOfAsThresholdLevelInd(v *ThresholdLevelIndOneOf) ThresholdLevelInd {
	return ThresholdLevelInd{
		ThresholdLevelIndOneOf: v,
	}
}

// ThresholdLevelIndOneOf1AsThresholdLevelInd is a convenience function that returns ThresholdLevelIndOneOf1 wrapped in ThresholdLevelInd
func ThresholdLevelIndOneOf1AsThresholdLevelInd(v *ThresholdLevelIndOneOf1) ThresholdLevelInd {
	return ThresholdLevelInd{
		ThresholdLevelIndOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThresholdLevelInd) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ThresholdLevelIndOneOf
	err = newStrictDecoder(data).Decode(&dst.ThresholdLevelIndOneOf)
	if err == nil {
		jsonThresholdLevelIndOneOf, _ := json.Marshal(dst.ThresholdLevelIndOneOf)
		if string(jsonThresholdLevelIndOneOf) == "{}" { // empty struct
			dst.ThresholdLevelIndOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ThresholdLevelIndOneOf = nil
	}

	// try to unmarshal data into ThresholdLevelIndOneOf1
	err = newStrictDecoder(data).Decode(&dst.ThresholdLevelIndOneOf1)
	if err == nil {
		jsonThresholdLevelIndOneOf1, _ := json.Marshal(dst.ThresholdLevelIndOneOf1)
		if string(jsonThresholdLevelIndOneOf1) == "{}" { // empty struct
			dst.ThresholdLevelIndOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ThresholdLevelIndOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ThresholdLevelIndOneOf = nil
		dst.ThresholdLevelIndOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ThresholdLevelInd)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThresholdLevelInd)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThresholdLevelInd) MarshalJSON() ([]byte, error) {
	if src.ThresholdLevelIndOneOf != nil {
		return json.Marshal(&src.ThresholdLevelIndOneOf)
	}

	if src.ThresholdLevelIndOneOf1 != nil {
		return json.Marshal(&src.ThresholdLevelIndOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThresholdLevelInd) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ThresholdLevelIndOneOf != nil {
		return obj.ThresholdLevelIndOneOf
	}

	if obj.ThresholdLevelIndOneOf1 != nil {
		return obj.ThresholdLevelIndOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableThresholdLevelInd struct {
	value *ThresholdLevelInd
	isSet bool
}

func (v NullableThresholdLevelInd) Get() *ThresholdLevelInd {
	return v.value
}

func (v *NullableThresholdLevelInd) Set(val *ThresholdLevelInd) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdLevelInd) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdLevelInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdLevelInd(val *ThresholdLevelInd) *NullableThresholdLevelInd {
	return &NullableThresholdLevelInd{value: val, isSet: true}
}

func (v NullableThresholdLevelInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdLevelInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


