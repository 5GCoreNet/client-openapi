/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// UpConfidentiality indicates whether UP confidentiality protection is required, preferred or not needed for all the traffic on the PDU Session. It shall comply with the provisions defined in table 5.4.3.5-1. 
type UpConfidentiality struct {
	UpIntegrityAnyOf *UpIntegrityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpConfidentiality) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpIntegrityAnyOf
	err = json.Unmarshal(data, &dst.UpIntegrityAnyOf);
	if err == nil {
		jsonUpIntegrityAnyOf, _ := json.Marshal(dst.UpIntegrityAnyOf)
		if string(jsonUpIntegrityAnyOf) == "{}" { // empty struct
			dst.UpIntegrityAnyOf = nil
		} else {
			return nil // data stored in dst.UpIntegrityAnyOf, return on the first match
		}
	} else {
		dst.UpIntegrityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UpConfidentiality)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpConfidentiality) MarshalJSON() ([]byte, error) {
	if src.UpIntegrityAnyOf != nil {
		return json.Marshal(&src.UpIntegrityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUpConfidentiality struct {
	value *UpConfidentiality
	isSet bool
}

func (v NullableUpConfidentiality) Get() *UpConfidentiality {
	return v.value
}

func (v *NullableUpConfidentiality) Set(val *UpConfidentiality) {
	v.value = val
	v.isSet = true
}

func (v NullableUpConfidentiality) IsSet() bool {
	return v.isSet
}

func (v *NullableUpConfidentiality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpConfidentiality(val *UpConfidentiality) *NullableUpConfidentiality {
	return &NullableUpConfidentiality{value: val, isSet: true}
}

func (v NullableUpConfidentiality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpConfidentiality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


