/*
TS 29.122 Common Data Types

Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// Event Possible values are - SESSION_TERMINATION: Indicates that Rx session is terminated. - LOSS_OF_BEARER : Indicates a loss of a bearer. - RECOVERY_OF_BEARER: Indicates a recovery of a bearer. - RELEASE_OF_BEARER: Indicates a release of a bearer. - USAGE_REPORT: Indicates the usage report event.  - FAILED_RESOURCES_ALLOCATION: Indicates the resource allocation is failed. - SUCCESSFUL_RESOURCES_ALLOCATION: Indicates the resource allocation is successful. 
type Event struct {
	EventAnyOf *EventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Event) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventAnyOf
	err = json.Unmarshal(data, &dst.EventAnyOf);
	if err == nil {
		jsonEventAnyOf, _ := json.Marshal(dst.EventAnyOf)
		if string(jsonEventAnyOf) == "{}" { // empty struct
			dst.EventAnyOf = nil
		} else {
			return nil // data stored in dst.EventAnyOf, return on the first match
		}
	} else {
		dst.EventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Event)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Event) MarshalJSON() ([]byte, error) {
	if src.EventAnyOf != nil {
		return json.Marshal(&src.EventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


