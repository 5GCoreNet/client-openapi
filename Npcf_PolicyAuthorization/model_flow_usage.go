/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// FlowUsage Describes the flow usage of the flows described by a media subcomponent.
type FlowUsage struct {
	FlowUsageAnyOf *FlowUsageAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FlowUsage) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FlowUsageAnyOf
	err = json.Unmarshal(data, &dst.FlowUsageAnyOf);
	if err == nil {
		jsonFlowUsageAnyOf, _ := json.Marshal(dst.FlowUsageAnyOf)
		if string(jsonFlowUsageAnyOf) == "{}" { // empty struct
			dst.FlowUsageAnyOf = nil
		} else {
			return nil // data stored in dst.FlowUsageAnyOf, return on the first match
		}
	} else {
		dst.FlowUsageAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FlowUsage)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FlowUsage) MarshalJSON() ([]byte, error) {
	if src.FlowUsageAnyOf != nil {
		return json.Marshal(&src.FlowUsageAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFlowUsage struct {
	value *FlowUsage
	isSet bool
}

func (v NullableFlowUsage) Get() *FlowUsage {
	return v.value
}

func (v *NullableFlowUsage) Set(val *FlowUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowUsage(val *FlowUsage) *NullableFlowUsage {
	return &NullableFlowUsage{value: val, isSet: true}
}

func (v NullableFlowUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


