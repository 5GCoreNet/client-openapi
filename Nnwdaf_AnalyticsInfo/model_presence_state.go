/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// PresenceState Possible values are: -IN_AREA: Indicates that the UE is inside or enters the presence reporting area. -OUT_OF_AREA: Indicates that the UE is outside or leaves the presence reporting area -UNKNOW: Indicates it is unknown whether the UE is in the presence reporting area or not -INACTIVE: Indicates that the presence reporting area is inactive in the serving node.  
type PresenceState struct {
	PresenceStateAnyOf *PresenceStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PresenceState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PresenceStateAnyOf
	err = json.Unmarshal(data, &dst.PresenceStateAnyOf);
	if err == nil {
		jsonPresenceStateAnyOf, _ := json.Marshal(dst.PresenceStateAnyOf)
		if string(jsonPresenceStateAnyOf) == "{}" { // empty struct
			dst.PresenceStateAnyOf = nil
		} else {
			return nil // data stored in dst.PresenceStateAnyOf, return on the first match
		}
	} else {
		dst.PresenceStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PresenceState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PresenceState) MarshalJSON() ([]byte, error) {
	if src.PresenceStateAnyOf != nil {
		return json.Marshal(&src.PresenceStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePresenceState struct {
	value *PresenceState
	isSet bool
}

func (v NullablePresenceState) Get() *PresenceState {
	return v.value
}

func (v *NullablePresenceState) Set(val *PresenceState) {
	v.value = val
	v.isSet = true
}

func (v NullablePresenceState) IsSet() bool {
	return v.isSet
}

func (v *NullablePresenceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresenceState(val *PresenceState) *NullablePresenceState {
	return &NullablePresenceState{value: val, isSet: true}
}

func (v NullablePresenceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresenceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


