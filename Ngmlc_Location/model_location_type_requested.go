/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LocationTypeRequested Contains the location type requested by the LCS client
type LocationTypeRequested struct {
	LocationTypeRequestedAnyOf *LocationTypeRequestedAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationTypeRequested) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocationTypeRequestedAnyOf
	err = json.Unmarshal(data, &dst.LocationTypeRequestedAnyOf);
	if err == nil {
		jsonLocationTypeRequestedAnyOf, _ := json.Marshal(dst.LocationTypeRequestedAnyOf)
		if string(jsonLocationTypeRequestedAnyOf) == "{}" { // empty struct
			dst.LocationTypeRequestedAnyOf = nil
		} else {
			return nil // data stored in dst.LocationTypeRequestedAnyOf, return on the first match
		}
	} else {
		dst.LocationTypeRequestedAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationTypeRequested)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationTypeRequested) MarshalJSON() ([]byte, error) {
	if src.LocationTypeRequestedAnyOf != nil {
		return json.Marshal(&src.LocationTypeRequestedAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationTypeRequested struct {
	value *LocationTypeRequested
	isSet bool
}

func (v NullableLocationTypeRequested) Get() *LocationTypeRequested {
	return v.value
}

func (v *NullableLocationTypeRequested) Set(val *LocationTypeRequested) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationTypeRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationTypeRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationTypeRequested(val *LocationTypeRequested) *NullableLocationTypeRequested {
	return &NullableLocationTypeRequested{value: val, isSet: true}
}

func (v NullableLocationTypeRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationTypeRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


