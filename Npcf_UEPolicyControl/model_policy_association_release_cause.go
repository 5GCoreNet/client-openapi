/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_UEPolicyControl

import (
	"encoding/json"
	"fmt"
)

// PolicyAssociationReleaseCause Possible values are: - UNSPECIFIED: This value is used for unspecified reasons. - UE_SUBSCRIPTION: This value is used to indicate that the policy association needs to be   terminated because the subscription of UE has changed (e.g. was removed). - INSUFFICIENT_RES: This value is used to indicate that the server is overloaded and needs   to abort the policy association. 
type PolicyAssociationReleaseCause struct {
	PolicyAssociationReleaseCauseAnyOf *PolicyAssociationReleaseCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyAssociationReleaseCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PolicyAssociationReleaseCauseAnyOf
	err = json.Unmarshal(data, &dst.PolicyAssociationReleaseCauseAnyOf);
	if err == nil {
		jsonPolicyAssociationReleaseCauseAnyOf, _ := json.Marshal(dst.PolicyAssociationReleaseCauseAnyOf)
		if string(jsonPolicyAssociationReleaseCauseAnyOf) == "{}" { // empty struct
			dst.PolicyAssociationReleaseCauseAnyOf = nil
		} else {
			return nil // data stored in dst.PolicyAssociationReleaseCauseAnyOf, return on the first match
		}
	} else {
		dst.PolicyAssociationReleaseCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyAssociationReleaseCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyAssociationReleaseCause) MarshalJSON() ([]byte, error) {
	if src.PolicyAssociationReleaseCauseAnyOf != nil {
		return json.Marshal(&src.PolicyAssociationReleaseCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyAssociationReleaseCause struct {
	value *PolicyAssociationReleaseCause
	isSet bool
}

func (v NullablePolicyAssociationReleaseCause) Get() *PolicyAssociationReleaseCause {
	return v.value
}

func (v *NullablePolicyAssociationReleaseCause) Set(val *PolicyAssociationReleaseCause) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAssociationReleaseCause) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAssociationReleaseCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAssociationReleaseCause(val *PolicyAssociationReleaseCause) *NullablePolicyAssociationReleaseCause {
	return &NullablePolicyAssociationReleaseCause{value: val, isSet: true}
}

func (v NullablePolicyAssociationReleaseCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAssociationReleaseCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


