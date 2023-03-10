/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
	"fmt"
)

// ProSeCapability Possible values are: - PROSE_DD: This value is used to indicate that 5G ProSe Direct Discovery is supported   by the UE. - PROSE_DC: This value is used to indicate that 5G ProSe Direct Communication is supported   by the UE. - PROSE_L2_U2N_RELAY: This value is used to indicate that Layer-2 5G ProSe UE-to-Network   Relay is supported by the UE. - PROSE_L3_U2N_RELAY: This value is used to indicate that Layer-3 5G ProSe UE-to-Network   Relay is supported by the UE. - PROSE_L2_REMOTE_UE: This value is used to indicate that Layer-2 5G ProSe Remote UE is   supported by the UE. - PROSE_L3_REMOTE_UE: This value is used to indicate that Layer-3 5G ProSe Remote UE is   supported by the UE. 
type ProSeCapability struct {
	ProSeCapabilityAnyOf *ProSeCapabilityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProSeCapability) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ProSeCapabilityAnyOf
	err = json.Unmarshal(data, &dst.ProSeCapabilityAnyOf);
	if err == nil {
		jsonProSeCapabilityAnyOf, _ := json.Marshal(dst.ProSeCapabilityAnyOf)
		if string(jsonProSeCapabilityAnyOf) == "{}" { // empty struct
			dst.ProSeCapabilityAnyOf = nil
		} else {
			return nil // data stored in dst.ProSeCapabilityAnyOf, return on the first match
		}
	} else {
		dst.ProSeCapabilityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ProSeCapability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProSeCapability) MarshalJSON() ([]byte, error) {
	if src.ProSeCapabilityAnyOf != nil {
		return json.Marshal(&src.ProSeCapabilityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProSeCapability struct {
	value *ProSeCapability
	isSet bool
}

func (v NullableProSeCapability) Get() *ProSeCapability {
	return v.value
}

func (v *NullableProSeCapability) Set(val *ProSeCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeCapability(val *ProSeCapability) *NullableProSeCapability {
	return &NullableProSeCapability{value: val, isSet: true}
}

func (v NullableProSeCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


