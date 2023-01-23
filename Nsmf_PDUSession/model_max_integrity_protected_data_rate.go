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

// MaxIntegrityProtectedDataRate Maximum Integrity Protected Data Rate. Possible values are   - 64_KBPS   - MAX_UE_RATE 
type MaxIntegrityProtectedDataRate struct {
	MaxIntegrityProtectedDataRateAnyOf *MaxIntegrityProtectedDataRateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MaxIntegrityProtectedDataRate) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MaxIntegrityProtectedDataRateAnyOf
	err = json.Unmarshal(data, &dst.MaxIntegrityProtectedDataRateAnyOf);
	if err == nil {
		jsonMaxIntegrityProtectedDataRateAnyOf, _ := json.Marshal(dst.MaxIntegrityProtectedDataRateAnyOf)
		if string(jsonMaxIntegrityProtectedDataRateAnyOf) == "{}" { // empty struct
			dst.MaxIntegrityProtectedDataRateAnyOf = nil
		} else {
			return nil // data stored in dst.MaxIntegrityProtectedDataRateAnyOf, return on the first match
		}
	} else {
		dst.MaxIntegrityProtectedDataRateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MaxIntegrityProtectedDataRate)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MaxIntegrityProtectedDataRate) MarshalJSON() ([]byte, error) {
	if src.MaxIntegrityProtectedDataRateAnyOf != nil {
		return json.Marshal(&src.MaxIntegrityProtectedDataRateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMaxIntegrityProtectedDataRate struct {
	value *MaxIntegrityProtectedDataRate
	isSet bool
}

func (v NullableMaxIntegrityProtectedDataRate) Get() *MaxIntegrityProtectedDataRate {
	return v.value
}

func (v *NullableMaxIntegrityProtectedDataRate) Set(val *MaxIntegrityProtectedDataRate) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxIntegrityProtectedDataRate) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxIntegrityProtectedDataRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxIntegrityProtectedDataRate(val *MaxIntegrityProtectedDataRate) *NullableMaxIntegrityProtectedDataRate {
	return &NullableMaxIntegrityProtectedDataRate{value: val, isSet: true}
}

func (v NullableMaxIntegrityProtectedDataRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxIntegrityProtectedDataRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


