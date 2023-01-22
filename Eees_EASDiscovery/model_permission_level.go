/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EASDiscovery

import (
	"encoding/json"
	"fmt"
)

// PermissionLevel Possible values are: - TRIAL: Level of service permission supported is TRIAL. - GOLD: Level of service permission supported is GOLD. - SILVER: Level of service permission supported is SILVER. - OTHER: Any other level of service permissions supported. 
type PermissionLevel struct {
	PermissionLevelAnyOf *PermissionLevelAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PermissionLevel) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PermissionLevelAnyOf
	err = json.Unmarshal(data, &dst.PermissionLevelAnyOf);
	if err == nil {
		jsonPermissionLevelAnyOf, _ := json.Marshal(dst.PermissionLevelAnyOf)
		if string(jsonPermissionLevelAnyOf) == "{}" { // empty struct
			dst.PermissionLevelAnyOf = nil
		} else {
			return nil // data stored in dst.PermissionLevelAnyOf, return on the first match
		}
	} else {
		dst.PermissionLevelAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PermissionLevel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PermissionLevel) MarshalJSON() ([]byte, error) {
	if src.PermissionLevelAnyOf != nil {
		return json.Marshal(&src.PermissionLevelAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePermissionLevel struct {
	value *PermissionLevel
	isSet bool
}

func (v NullablePermissionLevel) Get() *PermissionLevel {
	return v.value
}

func (v *NullablePermissionLevel) Set(val *PermissionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLevel(val *PermissionLevel) *NullablePermissionLevel {
	return &NullablePermissionLevel{value: val, isSet: true}
}

func (v NullablePermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


