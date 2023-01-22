/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package UAE_C2OperationModeManagement

import (
	"encoding/json"
	"fmt"
)

// C2CommModeSwitching Possible values are: - DIRECT_TO_NETWORK_ASSISTED_C2: Indicates the C2 Communication Mode switching from Direct C2 Communication mode to Network-Assisted C2 Communication mode. - NETWORK_ASSISTED_TO_DIRECT_C2: Indicates the C2 Communication Mode switching from Network-Assisted C2 Communication mode to Direct C2 Communication mode. - DIRECT_TO_UTM_NAVIGATED_C2: Indicates the C2 Communication Mode switching from Direct C2 Communication mode to UTM-Navigated C2 communication mode. - NETWORK_ASSISTED_TO_UTM_NAVIGATED_C2: Indicates the C2 Communication Mode switching from Network-Assisted C2 Communication mode to UTM-Navigated C2 communication mode. 
type C2CommModeSwitching struct {
	C2CommModeSwitchingAnyOf *C2CommModeSwitchingAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *C2CommModeSwitching) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into C2CommModeSwitchingAnyOf
	err = json.Unmarshal(data, &dst.C2CommModeSwitchingAnyOf);
	if err == nil {
		jsonC2CommModeSwitchingAnyOf, _ := json.Marshal(dst.C2CommModeSwitchingAnyOf)
		if string(jsonC2CommModeSwitchingAnyOf) == "{}" { // empty struct
			dst.C2CommModeSwitchingAnyOf = nil
		} else {
			return nil // data stored in dst.C2CommModeSwitchingAnyOf, return on the first match
		}
	} else {
		dst.C2CommModeSwitchingAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(C2CommModeSwitching)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *C2CommModeSwitching) MarshalJSON() ([]byte, error) {
	if src.C2CommModeSwitchingAnyOf != nil {
		return json.Marshal(&src.C2CommModeSwitchingAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableC2CommModeSwitching struct {
	value *C2CommModeSwitching
	isSet bool
}

func (v NullableC2CommModeSwitching) Get() *C2CommModeSwitching {
	return v.value
}

func (v *NullableC2CommModeSwitching) Set(val *C2CommModeSwitching) {
	v.value = val
	v.isSet = true
}

func (v NullableC2CommModeSwitching) IsSet() bool {
	return v.isSet
}

func (v *NullableC2CommModeSwitching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC2CommModeSwitching(val *C2CommModeSwitching) *NullableC2CommModeSwitching {
	return &NullableC2CommModeSwitching{value: val, isSet: true}
}

func (v NullableC2CommModeSwitching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC2CommModeSwitching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


