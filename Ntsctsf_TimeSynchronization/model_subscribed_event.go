/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// SubscribedEvent Possible values are: - AVAILABILITY_FOR_TIME_SYNC_SERVICE: The UE is availability for time synchronization service. 
type SubscribedEvent struct {
	SubscribedEventAnyOf *SubscribedEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SubscribedEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubscribedEventAnyOf
	err = json.Unmarshal(data, &dst.SubscribedEventAnyOf);
	if err == nil {
		jsonSubscribedEventAnyOf, _ := json.Marshal(dst.SubscribedEventAnyOf)
		if string(jsonSubscribedEventAnyOf) == "{}" { // empty struct
			dst.SubscribedEventAnyOf = nil
		} else {
			return nil // data stored in dst.SubscribedEventAnyOf, return on the first match
		}
	} else {
		dst.SubscribedEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SubscribedEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SubscribedEvent) MarshalJSON() ([]byte, error) {
	if src.SubscribedEventAnyOf != nil {
		return json.Marshal(&src.SubscribedEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubscribedEvent struct {
	value *SubscribedEvent
	isSet bool
}

func (v NullableSubscribedEvent) Get() *SubscribedEvent {
	return v.value
}

func (v *NullableSubscribedEvent) Set(val *SubscribedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedEvent(val *SubscribedEvent) *NullableSubscribedEvent {
	return &NullableSubscribedEvent{value: val, isSet: true}
}

func (v NullableSubscribedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


