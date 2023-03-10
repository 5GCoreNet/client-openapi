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

// RadioResourcesId struct for RadioResourcesId
type RadioResourcesId struct {
	RadioResourcesIdAnyOf *RadioResourcesIdAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RadioResourcesId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RadioResourcesIdAnyOf
	err = json.Unmarshal(data, &dst.RadioResourcesIdAnyOf);
	if err == nil {
		jsonRadioResourcesIdAnyOf, _ := json.Marshal(dst.RadioResourcesIdAnyOf)
		if string(jsonRadioResourcesIdAnyOf) == "{}" { // empty struct
			dst.RadioResourcesIdAnyOf = nil
		} else {
			return nil // data stored in dst.RadioResourcesIdAnyOf, return on the first match
		}
	} else {
		dst.RadioResourcesIdAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RadioResourcesId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RadioResourcesId) MarshalJSON() ([]byte, error) {
	if src.RadioResourcesIdAnyOf != nil {
		return json.Marshal(&src.RadioResourcesIdAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRadioResourcesId struct {
	value *RadioResourcesId
	isSet bool
}

func (v NullableRadioResourcesId) Get() *RadioResourcesId {
	return v.value
}

func (v *NullableRadioResourcesId) Set(val *RadioResourcesId) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioResourcesId) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioResourcesId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioResourcesId(val *RadioResourcesId) *NullableRadioResourcesId {
	return &NullableRadioResourcesId{value: val, isSet: true}
}

func (v NullableRadioResourcesId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioResourcesId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


