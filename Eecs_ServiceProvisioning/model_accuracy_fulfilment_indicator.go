/*
Eecs_ServiceProvisioning

API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eecs_ServiceProvisioning

import (
	"encoding/json"
	"fmt"
)

// AccuracyFulfilmentIndicator Indicates fulfilment of requested accuracy.
type AccuracyFulfilmentIndicator struct {
	AccuracyFulfilmentIndicatorAnyOf *AccuracyFulfilmentIndicatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccuracyFulfilmentIndicator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccuracyFulfilmentIndicatorAnyOf
	err = json.Unmarshal(data, &dst.AccuracyFulfilmentIndicatorAnyOf);
	if err == nil {
		jsonAccuracyFulfilmentIndicatorAnyOf, _ := json.Marshal(dst.AccuracyFulfilmentIndicatorAnyOf)
		if string(jsonAccuracyFulfilmentIndicatorAnyOf) == "{}" { // empty struct
			dst.AccuracyFulfilmentIndicatorAnyOf = nil
		} else {
			return nil // data stored in dst.AccuracyFulfilmentIndicatorAnyOf, return on the first match
		}
	} else {
		dst.AccuracyFulfilmentIndicatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccuracyFulfilmentIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccuracyFulfilmentIndicator) MarshalJSON() ([]byte, error) {
	if src.AccuracyFulfilmentIndicatorAnyOf != nil {
		return json.Marshal(&src.AccuracyFulfilmentIndicatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccuracyFulfilmentIndicator struct {
	value *AccuracyFulfilmentIndicator
	isSet bool
}

func (v NullableAccuracyFulfilmentIndicator) Get() *AccuracyFulfilmentIndicator {
	return v.value
}

func (v *NullableAccuracyFulfilmentIndicator) Set(val *AccuracyFulfilmentIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracyFulfilmentIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracyFulfilmentIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracyFulfilmentIndicator(val *AccuracyFulfilmentIndicator) *NullableAccuracyFulfilmentIndicator {
	return &NullableAccuracyFulfilmentIndicator{value: val, isSet: true}
}

func (v NullableAccuracyFulfilmentIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracyFulfilmentIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


