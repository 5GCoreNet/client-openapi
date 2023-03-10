/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ServiceAccessInformation

import (
	"encoding/json"
	"fmt"
)

// SdfMethod struct for SdfMethod
type SdfMethod struct {
	SdfMethodAnyOf *SdfMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SdfMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SdfMethodAnyOf
	err = json.Unmarshal(data, &dst.SdfMethodAnyOf);
	if err == nil {
		jsonSdfMethodAnyOf, _ := json.Marshal(dst.SdfMethodAnyOf)
		if string(jsonSdfMethodAnyOf) == "{}" { // empty struct
			dst.SdfMethodAnyOf = nil
		} else {
			return nil // data stored in dst.SdfMethodAnyOf, return on the first match
		}
	} else {
		dst.SdfMethodAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SdfMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SdfMethod) MarshalJSON() ([]byte, error) {
	if src.SdfMethodAnyOf != nil {
		return json.Marshal(&src.SdfMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSdfMethod struct {
	value *SdfMethod
	isSet bool
}

func (v NullableSdfMethod) Get() *SdfMethod {
	return v.value
}

func (v *NullableSdfMethod) Set(val *SdfMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSdfMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSdfMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdfMethod(val *SdfMethod) *NullableSdfMethod {
	return &NullableSdfMethod{value: val, isSet: true}
}

func (v NullableSdfMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdfMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


