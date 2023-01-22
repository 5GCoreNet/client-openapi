/*
Npcf_MBSPolicyAuthorization API

MBS Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_MBSPolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// PreemptionCapability The enumeration PreemptionCapability indicates the pre-emption capability of a request on other QoS flows. See clause 5.7.2.2 of 3GPP TS 23.501. It shall comply with the provisions defined in table 5.5.3.1-1. 
type PreemptionCapability struct {
	PreemptionCapabilityAnyOf *PreemptionCapabilityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PreemptionCapability) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PreemptionCapabilityAnyOf
	err = json.Unmarshal(data, &dst.PreemptionCapabilityAnyOf);
	if err == nil {
		jsonPreemptionCapabilityAnyOf, _ := json.Marshal(dst.PreemptionCapabilityAnyOf)
		if string(jsonPreemptionCapabilityAnyOf) == "{}" { // empty struct
			dst.PreemptionCapabilityAnyOf = nil
		} else {
			return nil // data stored in dst.PreemptionCapabilityAnyOf, return on the first match
		}
	} else {
		dst.PreemptionCapabilityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PreemptionCapability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PreemptionCapability) MarshalJSON() ([]byte, error) {
	if src.PreemptionCapabilityAnyOf != nil {
		return json.Marshal(&src.PreemptionCapabilityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePreemptionCapability struct {
	value *PreemptionCapability
	isSet bool
}

func (v NullablePreemptionCapability) Get() *PreemptionCapability {
	return v.value
}

func (v *NullablePreemptionCapability) Set(val *PreemptionCapability) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionCapability) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionCapability(val *PreemptionCapability) *NullablePreemptionCapability {
	return &NullablePreemptionCapability{value: val, isSet: true}
}

func (v NullablePreemptionCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


