/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// RoleOfIMSNode struct for RoleOfIMSNode
type RoleOfIMSNode struct {
	RoleOfIMSNodeAnyOf *RoleOfIMSNodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RoleOfIMSNode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RoleOfIMSNodeAnyOf
	err = json.Unmarshal(data, &dst.RoleOfIMSNodeAnyOf);
	if err == nil {
		jsonRoleOfIMSNodeAnyOf, _ := json.Marshal(dst.RoleOfIMSNodeAnyOf)
		if string(jsonRoleOfIMSNodeAnyOf) == "{}" { // empty struct
			dst.RoleOfIMSNodeAnyOf = nil
		} else {
			return nil // data stored in dst.RoleOfIMSNodeAnyOf, return on the first match
		}
	} else {
		dst.RoleOfIMSNodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RoleOfIMSNode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RoleOfIMSNode) MarshalJSON() ([]byte, error) {
	if src.RoleOfIMSNodeAnyOf != nil {
		return json.Marshal(&src.RoleOfIMSNodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRoleOfIMSNode struct {
	value *RoleOfIMSNode
	isSet bool
}

func (v NullableRoleOfIMSNode) Get() *RoleOfIMSNode {
	return v.value
}

func (v *NullableRoleOfIMSNode) Set(val *RoleOfIMSNode) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleOfIMSNode) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleOfIMSNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleOfIMSNode(val *RoleOfIMSNode) *NullableRoleOfIMSNode {
	return &NullableRoleOfIMSNode{value: val, isSet: true}
}

func (v NullableRoleOfIMSNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleOfIMSNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


