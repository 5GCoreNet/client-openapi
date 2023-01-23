/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
	"fmt"
)

// MbsFailureCode Possible values are: - NF_MALFUNCTION: Indicates that the MBS PCC rule could not be successfully installed due to MB-SMF/MB-UPF malfunction. - NF_RESOURCES_UNAVAILABLE: Indicates that the MBS PCC rule could not be successfully installed due to resources unavailable at MB-SMF/MB-UPF. - RESOURCE_ALLOCATION_FAILURE: Indicates that the MBS PCC rule could not be successfully installed or maintained since the associated MBS QoS flow establishment/modification failed or the associated MBS QoS flow was released. - MBS_QOS_VALIDATION_FAILURE: Indicates that MBS QoS validation has failed. - NO_MBS_QOS_FLOW: Indicates that there is no MBS QoS flow to which the MB-SMF can bind the MBS PCC rule(s). - MBS_QOS_DECISION_ERROR: Indicates failure in the provisioning of MBS QoS Decision data. - MBS_POLICY_PARAM_ERROR: Indicates that the information related to the provisioned MBS policy parameter(s) is incorrect, incomplete or inconsistent. 
type MbsFailureCode struct {
	MbsFailureCodeAnyOf *MbsFailureCodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsFailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsFailureCodeAnyOf
	err = json.Unmarshal(data, &dst.MbsFailureCodeAnyOf);
	if err == nil {
		jsonMbsFailureCodeAnyOf, _ := json.Marshal(dst.MbsFailureCodeAnyOf)
		if string(jsonMbsFailureCodeAnyOf) == "{}" { // empty struct
			dst.MbsFailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.MbsFailureCodeAnyOf, return on the first match
		}
	} else {
		dst.MbsFailureCodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsFailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsFailureCode) MarshalJSON() ([]byte, error) {
	if src.MbsFailureCodeAnyOf != nil {
		return json.Marshal(&src.MbsFailureCodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsFailureCode struct {
	value *MbsFailureCode
	isSet bool
}

func (v NullableMbsFailureCode) Get() *MbsFailureCode {
	return v.value
}

func (v *NullableMbsFailureCode) Set(val *MbsFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsFailureCode(val *MbsFailureCode) *NullableMbsFailureCode {
	return &NullableMbsFailureCode{value: val, isSet: true}
}

func (v NullableMbsFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


