/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// ScscfSelectionAssistanceInformation Information used by the I-CSCF to select an S-CSCF for the UE
type ScscfSelectionAssistanceInformation struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScscfSelectionAssistanceInformation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ScscfSelectionAssistanceInformation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScscfSelectionAssistanceInformation) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScscfSelectionAssistanceInformation struct {
	value *ScscfSelectionAssistanceInformation
	isSet bool
}

func (v NullableScscfSelectionAssistanceInformation) Get() *ScscfSelectionAssistanceInformation {
	return v.value
}

func (v *NullableScscfSelectionAssistanceInformation) Set(val *ScscfSelectionAssistanceInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableScscfSelectionAssistanceInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableScscfSelectionAssistanceInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScscfSelectionAssistanceInformation(val *ScscfSelectionAssistanceInformation) *NullableScscfSelectionAssistanceInformation {
	return &NullableScscfSelectionAssistanceInformation{value: val, isSet: true}
}

func (v NullableScscfSelectionAssistanceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScscfSelectionAssistanceInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


