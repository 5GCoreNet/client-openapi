/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// UpConfidentialityRm indicates whether UP integrity protection is required, preferred or not needed for all the  traffic on the PDU Session. It shall comply with the provisions defined in table 5.4.3.4-1, but with the OpenAPI 'nullable: true' property.  
type UpConfidentialityRm struct {
	NullValue *NullValue
	UpConfidentiality *UpConfidentiality
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpConfidentialityRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into UpConfidentiality
	err = json.Unmarshal(data, &dst.UpConfidentiality);
	if err == nil {
		jsonUpConfidentiality, _ := json.Marshal(dst.UpConfidentiality)
		if string(jsonUpConfidentiality) == "{}" { // empty struct
			dst.UpConfidentiality = nil
		} else {
			return nil // data stored in dst.UpConfidentiality, return on the first match
		}
	} else {
		dst.UpConfidentiality = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpConfidentialityRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpConfidentialityRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.UpConfidentiality != nil {
		return json.Marshal(&src.UpConfidentiality)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUpConfidentialityRm struct {
	value *UpConfidentialityRm
	isSet bool
}

func (v NullableUpConfidentialityRm) Get() *UpConfidentialityRm {
	return v.value
}

func (v *NullableUpConfidentialityRm) Set(val *UpConfidentialityRm) {
	v.value = val
	v.isSet = true
}

func (v NullableUpConfidentialityRm) IsSet() bool {
	return v.isSet
}

func (v *NullableUpConfidentialityRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpConfidentialityRm(val *UpConfidentialityRm) *NullableUpConfidentialityRm {
	return &NullableUpConfidentialityRm{value: val, isSet: true}
}

func (v NullableUpConfidentialityRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpConfidentialityRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


