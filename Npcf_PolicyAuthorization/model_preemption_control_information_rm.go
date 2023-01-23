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

// PreemptionControlInformationRm This data type is defined in the same way as the PreemptionControlInformation data type, but with the OpenAPI nullable property set to true.
type PreemptionControlInformationRm struct {
	NullValue *NullValue
	PreemptionControlInformation *PreemptionControlInformation
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PreemptionControlInformationRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into PreemptionControlInformation
	err = json.Unmarshal(data, &dst.PreemptionControlInformation);
	if err == nil {
		jsonPreemptionControlInformation, _ := json.Marshal(dst.PreemptionControlInformation)
		if string(jsonPreemptionControlInformation) == "{}" { // empty struct
			dst.PreemptionControlInformation = nil
		} else {
			return nil // data stored in dst.PreemptionControlInformation, return on the first match
		}
	} else {
		dst.PreemptionControlInformation = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PreemptionControlInformationRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PreemptionControlInformationRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.PreemptionControlInformation != nil {
		return json.Marshal(&src.PreemptionControlInformation)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePreemptionControlInformationRm struct {
	value *PreemptionControlInformationRm
	isSet bool
}

func (v NullablePreemptionControlInformationRm) Get() *PreemptionControlInformationRm {
	return v.value
}

func (v *NullablePreemptionControlInformationRm) Set(val *PreemptionControlInformationRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionControlInformationRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionControlInformationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionControlInformationRm(val *PreemptionControlInformationRm) *NullablePreemptionControlInformationRm {
	return &NullablePreemptionControlInformationRm{value: val, isSet: true}
}

func (v NullablePreemptionControlInformationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionControlInformationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


