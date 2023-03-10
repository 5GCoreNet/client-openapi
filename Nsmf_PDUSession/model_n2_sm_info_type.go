/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// N2SmInfoType N2 SM Information Type. Possible values are - PDU_RES_SETUP_REQ - PDU_RES_SETUP_RSP - PDU_RES_SETUP_FAIL - PDU_RES_REL_CMD - PDU_RES_REL_RSP - PDU_RES_MOD_REQ - PDU_RES_MOD_RSP - PDU_RES_MOD_FAIL - PDU_RES_NTY - PDU_RES_NTY_REL - PDU_RES_MOD_IND - PDU_RES_MOD_CFM - PATH_SWITCH_REQ - PATH_SWITCH_SETUP_FAIL - PATH_SWITCH_REQ_ACK - PATH_SWITCH_REQ_FAIL - HANDOVER_REQUIRED - HANDOVER_CMD - HANDOVER_PREP_FAIL - HANDOVER_REQ_ACK - HANDOVER_RES_ALLOC_FAIL - SECONDARY_RAT_USAGE - PDU_RES_MOD_IND_FAIL - UE_CONTEXT_RESUME_REQ - UE_CONTEXT_RESUME_RSP - UE_CONTEXT_SUSPEND_REQ 
type N2SmInfoType struct {
	N2SmInfoTypeAnyOf *N2SmInfoTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2SmInfoType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into N2SmInfoTypeAnyOf
	err = json.Unmarshal(data, &dst.N2SmInfoTypeAnyOf);
	if err == nil {
		jsonN2SmInfoTypeAnyOf, _ := json.Marshal(dst.N2SmInfoTypeAnyOf)
		if string(jsonN2SmInfoTypeAnyOf) == "{}" { // empty struct
			dst.N2SmInfoTypeAnyOf = nil
		} else {
			return nil // data stored in dst.N2SmInfoTypeAnyOf, return on the first match
		}
	} else {
		dst.N2SmInfoTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2SmInfoType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2SmInfoType) MarshalJSON() ([]byte, error) {
	if src.N2SmInfoTypeAnyOf != nil {
		return json.Marshal(&src.N2SmInfoTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2SmInfoType struct {
	value *N2SmInfoType
	isSet bool
}

func (v NullableN2SmInfoType) Get() *N2SmInfoType {
	return v.value
}

func (v *NullableN2SmInfoType) Set(val *N2SmInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableN2SmInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableN2SmInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2SmInfoType(val *N2SmInfoType) *NullableN2SmInfoType {
	return &NullableN2SmInfoType{value: val, isSet: true}
}

func (v NullableN2SmInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2SmInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


