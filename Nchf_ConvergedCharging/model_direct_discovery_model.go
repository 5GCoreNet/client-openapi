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

// DirectDiscoveryModel struct for DirectDiscoveryModel
type DirectDiscoveryModel struct {
	DirectDiscoveryModelAnyOf *DirectDiscoveryModelAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DirectDiscoveryModel) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DirectDiscoveryModelAnyOf
	err = json.Unmarshal(data, &dst.DirectDiscoveryModelAnyOf);
	if err == nil {
		jsonDirectDiscoveryModelAnyOf, _ := json.Marshal(dst.DirectDiscoveryModelAnyOf)
		if string(jsonDirectDiscoveryModelAnyOf) == "{}" { // empty struct
			dst.DirectDiscoveryModelAnyOf = nil
		} else {
			return nil // data stored in dst.DirectDiscoveryModelAnyOf, return on the first match
		}
	} else {
		dst.DirectDiscoveryModelAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DirectDiscoveryModel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DirectDiscoveryModel) MarshalJSON() ([]byte, error) {
	if src.DirectDiscoveryModelAnyOf != nil {
		return json.Marshal(&src.DirectDiscoveryModelAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDirectDiscoveryModel struct {
	value *DirectDiscoveryModel
	isSet bool
}

func (v NullableDirectDiscoveryModel) Get() *DirectDiscoveryModel {
	return v.value
}

func (v *NullableDirectDiscoveryModel) Set(val *DirectDiscoveryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDiscoveryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDiscoveryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDiscoveryModel(val *DirectDiscoveryModel) *NullableDirectDiscoveryModel {
	return &NullableDirectDiscoveryModel{value: val, isSet: true}
}

func (v NullableDirectDiscoveryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDiscoveryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


