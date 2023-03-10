/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
	"fmt"
)

// SecFeature Security features supported by the BSF or the NAF
type SecFeature struct {
	SecFeatureAnyOf *SecFeatureAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SecFeature) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SecFeatureAnyOf
	err = json.Unmarshal(data, &dst.SecFeatureAnyOf);
	if err == nil {
		jsonSecFeatureAnyOf, _ := json.Marshal(dst.SecFeatureAnyOf)
		if string(jsonSecFeatureAnyOf) == "{}" { // empty struct
			dst.SecFeatureAnyOf = nil
		} else {
			return nil // data stored in dst.SecFeatureAnyOf, return on the first match
		}
	} else {
		dst.SecFeatureAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SecFeature)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SecFeature) MarshalJSON() ([]byte, error) {
	if src.SecFeatureAnyOf != nil {
		return json.Marshal(&src.SecFeatureAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSecFeature struct {
	value *SecFeature
	isSet bool
}

func (v NullableSecFeature) Get() *SecFeature {
	return v.value
}

func (v *NullableSecFeature) Set(val *SecFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableSecFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableSecFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecFeature(val *SecFeature) *NullableSecFeature {
	return &NullableSecFeature{value: val, isSet: true}
}

func (v NullableSecFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


