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

// ManagementOperationStatus struct for ManagementOperationStatus
type ManagementOperationStatus struct {
	ManagementOperationStatusAnyOf *ManagementOperationStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ManagementOperationStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ManagementOperationStatusAnyOf
	err = json.Unmarshal(data, &dst.ManagementOperationStatusAnyOf);
	if err == nil {
		jsonManagementOperationStatusAnyOf, _ := json.Marshal(dst.ManagementOperationStatusAnyOf)
		if string(jsonManagementOperationStatusAnyOf) == "{}" { // empty struct
			dst.ManagementOperationStatusAnyOf = nil
		} else {
			return nil // data stored in dst.ManagementOperationStatusAnyOf, return on the first match
		}
	} else {
		dst.ManagementOperationStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ManagementOperationStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ManagementOperationStatus) MarshalJSON() ([]byte, error) {
	if src.ManagementOperationStatusAnyOf != nil {
		return json.Marshal(&src.ManagementOperationStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableManagementOperationStatus struct {
	value *ManagementOperationStatus
	isSet bool
}

func (v NullableManagementOperationStatus) Get() *ManagementOperationStatus {
	return v.value
}

func (v *NullableManagementOperationStatus) Set(val *ManagementOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementOperationStatus(val *ManagementOperationStatus) *NullableManagementOperationStatus {
	return &NullableManagementOperationStatus{value: val, isSet: true}
}

func (v NullableManagementOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


