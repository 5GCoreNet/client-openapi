/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// EventType struct for EventType
type EventType struct {
	EventTypeAnyOf *EventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventTypeAnyOf
	err = json.Unmarshal(data, &dst.EventTypeAnyOf);
	if err == nil {
		jsonEventTypeAnyOf, _ := json.Marshal(dst.EventTypeAnyOf)
		if string(jsonEventTypeAnyOf) == "{}" { // empty struct
			dst.EventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.EventTypeAnyOf, return on the first match
		}
	} else {
		dst.EventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventType) MarshalJSON() ([]byte, error) {
	if src.EventTypeAnyOf != nil {
		return json.Marshal(&src.EventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


