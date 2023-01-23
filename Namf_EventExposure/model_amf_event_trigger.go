/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// AmfEventTrigger Describes how AMF should generate the report for the event
type AmfEventTrigger struct {
	AmfEventTriggerAnyOf *AmfEventTriggerAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmfEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AmfEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.AmfEventTriggerAnyOf);
	if err == nil {
		jsonAmfEventTriggerAnyOf, _ := json.Marshal(dst.AmfEventTriggerAnyOf)
		if string(jsonAmfEventTriggerAnyOf) == "{}" { // empty struct
			dst.AmfEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.AmfEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.AmfEventTriggerAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AmfEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmfEventTrigger) MarshalJSON() ([]byte, error) {
	if src.AmfEventTriggerAnyOf != nil {
		return json.Marshal(&src.AmfEventTriggerAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmfEventTrigger struct {
	value *AmfEventTrigger
	isSet bool
}

func (v NullableAmfEventTrigger) Get() *AmfEventTrigger {
	return v.value
}

func (v *NullableAmfEventTrigger) Set(val *AmfEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventTrigger(val *AmfEventTrigger) *NullableAmfEventTrigger {
	return &NullableAmfEventTrigger{value: val, isSet: true}
}

func (v NullableAmfEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


