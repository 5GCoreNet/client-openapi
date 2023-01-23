/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// RuleOperation Possible values are - CREATE_PCC_RULE: Indicates to create a new PCC rule to reserve the resource requested by the UE.  - DELETE_PCC_RULE: Indicates to delete a PCC rule corresponding to reserve the resource requested by the UE. - MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS: Indicates to modify the PCC rule by adding new packet filter(s). - MODIFY_ PCC_RULE_AND_REPLACE_PACKET_FILTERS: Indicates to modify the PCC rule by replacing the existing packet filter(s). - MODIFY_ PCC_RULE_AND_DELETE_PACKET_FILTERS: Indicates to modify the PCC rule by deleting the existing packet filter(s). - MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS: Indicates to modify the PCC rule by modifying the QoS of the PCC rule. 
type RuleOperation struct {
	RuleOperationAnyOf *RuleOperationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RuleOperation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RuleOperationAnyOf
	err = json.Unmarshal(data, &dst.RuleOperationAnyOf);
	if err == nil {
		jsonRuleOperationAnyOf, _ := json.Marshal(dst.RuleOperationAnyOf)
		if string(jsonRuleOperationAnyOf) == "{}" { // empty struct
			dst.RuleOperationAnyOf = nil
		} else {
			return nil // data stored in dst.RuleOperationAnyOf, return on the first match
		}
	} else {
		dst.RuleOperationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RuleOperation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RuleOperation) MarshalJSON() ([]byte, error) {
	if src.RuleOperationAnyOf != nil {
		return json.Marshal(&src.RuleOperationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRuleOperation struct {
	value *RuleOperation
	isSet bool
}

func (v NullableRuleOperation) Get() *RuleOperation {
	return v.value
}

func (v *NullableRuleOperation) Set(val *RuleOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleOperation(val *RuleOperation) *NullableRuleOperation {
	return &NullableRuleOperation{value: val, isSet: true}
}

func (v NullableRuleOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


