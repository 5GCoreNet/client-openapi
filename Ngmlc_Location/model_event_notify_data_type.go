/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// EventNotifyDataType Contains the type of event that triggers event notification
type EventNotifyDataType struct {
	EventNotifyDataTypeAnyOf *EventNotifyDataTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventNotifyDataType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventNotifyDataTypeAnyOf
	err = json.Unmarshal(data, &dst.EventNotifyDataTypeAnyOf);
	if err == nil {
		jsonEventNotifyDataTypeAnyOf, _ := json.Marshal(dst.EventNotifyDataTypeAnyOf)
		if string(jsonEventNotifyDataTypeAnyOf) == "{}" { // empty struct
			dst.EventNotifyDataTypeAnyOf = nil
		} else {
			return nil // data stored in dst.EventNotifyDataTypeAnyOf, return on the first match
		}
	} else {
		dst.EventNotifyDataTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EventNotifyDataType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventNotifyDataType) MarshalJSON() ([]byte, error) {
	if src.EventNotifyDataTypeAnyOf != nil {
		return json.Marshal(&src.EventNotifyDataTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventNotifyDataType struct {
	value *EventNotifyDataType
	isSet bool
}

func (v NullableEventNotifyDataType) Get() *EventNotifyDataType {
	return v.value
}

func (v *NullableEventNotifyDataType) Set(val *EventNotifyDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotifyDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotifyDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotifyDataType(val *EventNotifyDataType) *NullableEventNotifyDataType {
	return &NullableEventNotifyDataType{value: val, isSet: true}
}

func (v NullableEventNotifyDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotifyDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


