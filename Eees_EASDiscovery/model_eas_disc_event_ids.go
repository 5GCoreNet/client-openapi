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

// EASDiscEventIDs Possible values are - EAS_AVAILABILITY_CHANGE: Represents the EAS availability change event. - EAS_DYNAMIC_INFO_CHANGE: Represents the EAS dynamic information change event. 
type EASDiscEventIDs struct {
	EASDiscEventIDsAnyOf *EASDiscEventIDsAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EASDiscEventIDs) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EASDiscEventIDsAnyOf
	err = json.Unmarshal(data, &dst.EASDiscEventIDsAnyOf);
	if err == nil {
		jsonEASDiscEventIDsAnyOf, _ := json.Marshal(dst.EASDiscEventIDsAnyOf)
		if string(jsonEASDiscEventIDsAnyOf) == "{}" { // empty struct
			dst.EASDiscEventIDsAnyOf = nil
		} else {
			return nil // data stored in dst.EASDiscEventIDsAnyOf, return on the first match
		}
	} else {
		dst.EASDiscEventIDsAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EASDiscEventIDs)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EASDiscEventIDs) MarshalJSON() ([]byte, error) {
	if src.EASDiscEventIDsAnyOf != nil {
		return json.Marshal(&src.EASDiscEventIDsAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEASDiscEventIDs struct {
	value *EASDiscEventIDs
	isSet bool
}

func (v NullableEASDiscEventIDs) Get() *EASDiscEventIDs {
	return v.value
}

func (v *NullableEASDiscEventIDs) Set(val *EASDiscEventIDs) {
	v.value = val
	v.isSet = true
}

func (v NullableEASDiscEventIDs) IsSet() bool {
	return v.isSet
}

func (v *NullableEASDiscEventIDs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASDiscEventIDs(val *EASDiscEventIDs) *NullableEASDiscEventIDs {
	return &NullableEASDiscEventIDs{value: val, isSet: true}
}

func (v NullableEASDiscEventIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASDiscEventIDs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


