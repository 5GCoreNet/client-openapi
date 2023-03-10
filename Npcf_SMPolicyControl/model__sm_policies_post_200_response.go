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

// SmPoliciesPost200Response - struct for SmPoliciesPost200Response
type SmPoliciesPost200Response struct {
	UeCampingRep *UeCampingRep
	ArrayOfPartialSuccessReport *[]PartialSuccessReport
	ArrayOfPolicyDecisionFailureCode *[]PolicyDecisionFailureCode
}

// UeCampingRepAsSmPoliciesPost200Response is a convenience function that returns UeCampingRep wrapped in SmPoliciesPost200Response
func UeCampingRepAsSmPoliciesPost200Response(v *UeCampingRep) SmPoliciesPost200Response {
	return SmPoliciesPost200Response{
		UeCampingRep: v,
	}
}

// []PartialSuccessReportAsSmPoliciesPost200Response is a convenience function that returns []PartialSuccessReport wrapped in SmPoliciesPost200Response
func ArrayOfPartialSuccessReportAsSmPoliciesPost200Response(v *[]PartialSuccessReport) SmPoliciesPost200Response {
	return SmPoliciesPost200Response{
		ArrayOfPartialSuccessReport: v,
	}
}

// []PolicyDecisionFailureCodeAsSmPoliciesPost200Response is a convenience function that returns []PolicyDecisionFailureCode wrapped in SmPoliciesPost200Response
func ArrayOfPolicyDecisionFailureCodeAsSmPoliciesPost200Response(v *[]PolicyDecisionFailureCode) SmPoliciesPost200Response {
	return SmPoliciesPost200Response{
		ArrayOfPolicyDecisionFailureCode: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SmPoliciesPost200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UeCampingRep
	err = newStrictDecoder(data).Decode(&dst.UeCampingRep)
	if err == nil {
		jsonUeCampingRep, _ := json.Marshal(dst.UeCampingRep)
		if string(jsonUeCampingRep) == "{}" { // empty struct
			dst.UeCampingRep = nil
		} else {
			match++
		}
	} else {
		dst.UeCampingRep = nil
	}

	// try to unmarshal data into ArrayOfPartialSuccessReport
	err = newStrictDecoder(data).Decode(&dst.ArrayOfPartialSuccessReport)
	if err == nil {
		jsonArrayOfPartialSuccessReport, _ := json.Marshal(dst.ArrayOfPartialSuccessReport)
		if string(jsonArrayOfPartialSuccessReport) == "{}" { // empty struct
			dst.ArrayOfPartialSuccessReport = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfPartialSuccessReport = nil
	}

	// try to unmarshal data into ArrayOfPolicyDecisionFailureCode
	err = newStrictDecoder(data).Decode(&dst.ArrayOfPolicyDecisionFailureCode)
	if err == nil {
		jsonArrayOfPolicyDecisionFailureCode, _ := json.Marshal(dst.ArrayOfPolicyDecisionFailureCode)
		if string(jsonArrayOfPolicyDecisionFailureCode) == "{}" { // empty struct
			dst.ArrayOfPolicyDecisionFailureCode = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfPolicyDecisionFailureCode = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UeCampingRep = nil
		dst.ArrayOfPartialSuccessReport = nil
		dst.ArrayOfPolicyDecisionFailureCode = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SmPoliciesPost200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SmPoliciesPost200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SmPoliciesPost200Response) MarshalJSON() ([]byte, error) {
	if src.UeCampingRep != nil {
		return json.Marshal(&src.UeCampingRep)
	}

	if src.ArrayOfPartialSuccessReport != nil {
		return json.Marshal(&src.ArrayOfPartialSuccessReport)
	}

	if src.ArrayOfPolicyDecisionFailureCode != nil {
		return json.Marshal(&src.ArrayOfPolicyDecisionFailureCode)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SmPoliciesPost200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UeCampingRep != nil {
		return obj.UeCampingRep
	}

	if obj.ArrayOfPartialSuccessReport != nil {
		return obj.ArrayOfPartialSuccessReport
	}

	if obj.ArrayOfPolicyDecisionFailureCode != nil {
		return obj.ArrayOfPolicyDecisionFailureCode
	}

	// all schemas are nil
	return nil
}

type NullableSmPoliciesPost200Response struct {
	value *SmPoliciesPost200Response
	isSet bool
}

func (v NullableSmPoliciesPost200Response) Get() *SmPoliciesPost200Response {
	return v.value
}

func (v *NullableSmPoliciesPost200Response) Set(val *SmPoliciesPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPoliciesPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPoliciesPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPoliciesPost200Response(val *SmPoliciesPost200Response) *NullableSmPoliciesPost200Response {
	return &NullableSmPoliciesPost200Response{value: val, isSet: true}
}

func (v NullableSmPoliciesPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPoliciesPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


