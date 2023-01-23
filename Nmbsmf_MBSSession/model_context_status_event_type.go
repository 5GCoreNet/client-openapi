/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// ContextStatusEventType Context Status Event Type
type ContextStatusEventType struct {
	ContextStatusEventTypeAnyOf *ContextStatusEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ContextStatusEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ContextStatusEventTypeAnyOf
	err = json.Unmarshal(data, &dst.ContextStatusEventTypeAnyOf);
	if err == nil {
		jsonContextStatusEventTypeAnyOf, _ := json.Marshal(dst.ContextStatusEventTypeAnyOf)
		if string(jsonContextStatusEventTypeAnyOf) == "{}" { // empty struct
			dst.ContextStatusEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ContextStatusEventTypeAnyOf, return on the first match
		}
	} else {
		dst.ContextStatusEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ContextStatusEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ContextStatusEventType) MarshalJSON() ([]byte, error) {
	if src.ContextStatusEventTypeAnyOf != nil {
		return json.Marshal(&src.ContextStatusEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableContextStatusEventType struct {
	value *ContextStatusEventType
	isSet bool
}

func (v NullableContextStatusEventType) Get() *ContextStatusEventType {
	return v.value
}

func (v *NullableContextStatusEventType) Set(val *ContextStatusEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusEventType(val *ContextStatusEventType) *NullableContextStatusEventType {
	return &NullableContextStatusEventType{value: val, isSet: true}
}

func (v NullableContextStatusEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


