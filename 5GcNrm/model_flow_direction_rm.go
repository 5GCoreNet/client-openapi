/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// FlowDirectionRm This data type is defined in the same way as the \"FlowDirection\" data type, with the only difference that it allows null value.
type FlowDirectionRm struct {
	FlowDirection *FlowDirection
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FlowDirectionRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FlowDirection
	err = json.Unmarshal(data, &dst.FlowDirection);
	if err == nil {
		jsonFlowDirection, _ := json.Marshal(dst.FlowDirection)
		if string(jsonFlowDirection) == "{}" { // empty struct
			dst.FlowDirection = nil
		} else {
			return nil // data stored in dst.FlowDirection, return on the first match
		}
	} else {
		dst.FlowDirection = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FlowDirectionRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FlowDirectionRm) MarshalJSON() ([]byte, error) {
	if src.FlowDirection != nil {
		return json.Marshal(&src.FlowDirection)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFlowDirectionRm struct {
	value *FlowDirectionRm
	isSet bool
}

func (v NullableFlowDirectionRm) Get() *FlowDirectionRm {
	return v.value
}

func (v *NullableFlowDirectionRm) Set(val *FlowDirectionRm) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDirectionRm) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDirectionRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDirectionRm(val *FlowDirectionRm) *NullableFlowDirectionRm {
	return &NullableFlowDirectionRm{value: val, isSet: true}
}

func (v NullableFlowDirectionRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDirectionRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


