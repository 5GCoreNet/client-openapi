/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
	"fmt"
)

// TerminationMode Possible values are: - TIME_TRIGGERED: Time-triggered termination mode. - EVENT_TRIGGERED_NUM_REPORTS_REACHED: Event-triggered termination number of reports reached mode. - EVENT_TRIGGERED_MEAS_THR_REACHED: The event-triggered termination measurement index threshold reached mode. - USER_TRIGGERED: User-triggered termination mode. 
type TerminationMode struct {
	TerminationModeAnyOf *TerminationModeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TerminationMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TerminationModeAnyOf
	err = json.Unmarshal(data, &dst.TerminationModeAnyOf);
	if err == nil {
		jsonTerminationModeAnyOf, _ := json.Marshal(dst.TerminationModeAnyOf)
		if string(jsonTerminationModeAnyOf) == "{}" { // empty struct
			dst.TerminationModeAnyOf = nil
		} else {
			return nil // data stored in dst.TerminationModeAnyOf, return on the first match
		}
	} else {
		dst.TerminationModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TerminationMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TerminationMode) MarshalJSON() ([]byte, error) {
	if src.TerminationModeAnyOf != nil {
		return json.Marshal(&src.TerminationModeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTerminationMode struct {
	value *TerminationMode
	isSet bool
}

func (v NullableTerminationMode) Get() *TerminationMode {
	return v.value
}

func (v *NullableTerminationMode) Set(val *TerminationMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationMode(val *TerminationMode) *NullableTerminationMode {
	return &NullableTerminationMode{value: val, isSet: true}
}

func (v NullableTerminationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


