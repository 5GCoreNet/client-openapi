/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// HssAuthType struct for HssAuthType
type HssAuthType struct {
	HssAuthTypeAnyOf *HssAuthTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *HssAuthType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into HssAuthTypeAnyOf
	err = json.Unmarshal(data, &dst.HssAuthTypeAnyOf);
	if err == nil {
		jsonHssAuthTypeAnyOf, _ := json.Marshal(dst.HssAuthTypeAnyOf)
		if string(jsonHssAuthTypeAnyOf) == "{}" { // empty struct
			dst.HssAuthTypeAnyOf = nil
		} else {
			return nil // data stored in dst.HssAuthTypeAnyOf, return on the first match
		}
	} else {
		dst.HssAuthTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(HssAuthType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *HssAuthType) MarshalJSON() ([]byte, error) {
	if src.HssAuthTypeAnyOf != nil {
		return json.Marshal(&src.HssAuthTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableHssAuthType struct {
	value *HssAuthType
	isSet bool
}

func (v NullableHssAuthType) Get() *HssAuthType {
	return v.value
}

func (v *NullableHssAuthType) Set(val *HssAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableHssAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableHssAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssAuthType(val *HssAuthType) *NullableHssAuthType {
	return &NullableHssAuthType{value: val, isSet: true}
}

func (v NullableHssAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


