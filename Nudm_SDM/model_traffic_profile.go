/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// TrafficProfile Possible values are: - SINGLE_TRANS_UL: Uplink single packet transmission. - SINGLE_TRANS_DL: Downlink single packet transmission. - DUAL_TRANS_UL_FIRST: Dual packet transmission, firstly uplink packet transmission   with subsequent downlink packet transmission. - DUAL_TRANS_DL_FIRST: Dual packet transmission, firstly downlink packet transmission   with subsequent uplink packet transmission.  
type TrafficProfile struct {
	TrafficProfileAnyOf *TrafficProfileAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficProfile) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TrafficProfileAnyOf
	err = json.Unmarshal(data, &dst.TrafficProfileAnyOf);
	if err == nil {
		jsonTrafficProfileAnyOf, _ := json.Marshal(dst.TrafficProfileAnyOf)
		if string(jsonTrafficProfileAnyOf) == "{}" { // empty struct
			dst.TrafficProfileAnyOf = nil
		} else {
			return nil // data stored in dst.TrafficProfileAnyOf, return on the first match
		}
	} else {
		dst.TrafficProfileAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficProfile)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficProfile) MarshalJSON() ([]byte, error) {
	if src.TrafficProfileAnyOf != nil {
		return json.Marshal(&src.TrafficProfileAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficProfile struct {
	value *TrafficProfile
	isSet bool
}

func (v NullableTrafficProfile) Get() *TrafficProfile {
	return v.value
}

func (v *NullableTrafficProfile) Set(val *TrafficProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficProfile(val *TrafficProfile) *NullableTrafficProfile {
	return &NullableTrafficProfile{value: val, isSet: true}
}

func (v NullableTrafficProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


