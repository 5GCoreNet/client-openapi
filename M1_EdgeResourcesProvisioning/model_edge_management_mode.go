/*
M1_EdgeResourcesProvisioning

5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_EdgeResourcesProvisioning

import (
	"encoding/json"
	"fmt"
)

// EdgeManagementMode The management mode of an EAS.
type EdgeManagementMode struct {
	EdgeManagementModeAnyOf *EdgeManagementModeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EdgeManagementMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EdgeManagementModeAnyOf
	err = json.Unmarshal(data, &dst.EdgeManagementModeAnyOf);
	if err == nil {
		jsonEdgeManagementModeAnyOf, _ := json.Marshal(dst.EdgeManagementModeAnyOf)
		if string(jsonEdgeManagementModeAnyOf) == "{}" { // empty struct
			dst.EdgeManagementModeAnyOf = nil
		} else {
			return nil // data stored in dst.EdgeManagementModeAnyOf, return on the first match
		}
	} else {
		dst.EdgeManagementModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EdgeManagementMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EdgeManagementMode) MarshalJSON() ([]byte, error) {
	if src.EdgeManagementModeAnyOf != nil {
		return json.Marshal(&src.EdgeManagementModeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEdgeManagementMode struct {
	value *EdgeManagementMode
	isSet bool
}

func (v NullableEdgeManagementMode) Get() *EdgeManagementMode {
	return v.value
}

func (v *NullableEdgeManagementMode) Set(val *EdgeManagementMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeManagementMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeManagementMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeManagementMode(val *EdgeManagementMode) *NullableEdgeManagementMode {
	return &NullableEdgeManagementMode{value: val, isSet: true}
}

func (v NullableEdgeManagementMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeManagementMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


