/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// SACRepFormat Indicates the NSAC reporting format.
type SACRepFormat struct {
	SACRepFormatAnyOf *SACRepFormatAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SACRepFormat) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SACRepFormatAnyOf
	err = json.Unmarshal(data, &dst.SACRepFormatAnyOf);
	if err == nil {
		jsonSACRepFormatAnyOf, _ := json.Marshal(dst.SACRepFormatAnyOf)
		if string(jsonSACRepFormatAnyOf) == "{}" { // empty struct
			dst.SACRepFormatAnyOf = nil
		} else {
			return nil // data stored in dst.SACRepFormatAnyOf, return on the first match
		}
	} else {
		dst.SACRepFormatAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SACRepFormat)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SACRepFormat) MarshalJSON() ([]byte, error) {
	if src.SACRepFormatAnyOf != nil {
		return json.Marshal(&src.SACRepFormatAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSACRepFormat struct {
	value *SACRepFormat
	isSet bool
}

func (v NullableSACRepFormat) Get() *SACRepFormat {
	return v.value
}

func (v *NullableSACRepFormat) Set(val *SACRepFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableSACRepFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableSACRepFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACRepFormat(val *SACRepFormat) *NullableSACRepFormat {
	return &NullableSACRepFormat{value: val, isSet: true}
}

func (v NullableSACRepFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACRepFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


