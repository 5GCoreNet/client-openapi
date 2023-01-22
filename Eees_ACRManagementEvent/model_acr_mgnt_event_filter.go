/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_ACRManagementEvent

import (
	"encoding/json"
	"fmt"
)

// AcrMgntEventFilter Possible values are: - INTRA_EDN_MOBILITY: Indicates that the ACR Management Event filter is intra-EDN mobility. - INTER_EDN_MOBILITY: Indicates that the ACR Management Event filter is inter-EDN mobility. 
type AcrMgntEventFilter struct {
	AcrMgntEventFilterAnyOf *AcrMgntEventFilterAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AcrMgntEventFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AcrMgntEventFilterAnyOf
	err = json.Unmarshal(data, &dst.AcrMgntEventFilterAnyOf);
	if err == nil {
		jsonAcrMgntEventFilterAnyOf, _ := json.Marshal(dst.AcrMgntEventFilterAnyOf)
		if string(jsonAcrMgntEventFilterAnyOf) == "{}" { // empty struct
			dst.AcrMgntEventFilterAnyOf = nil
		} else {
			return nil // data stored in dst.AcrMgntEventFilterAnyOf, return on the first match
		}
	} else {
		dst.AcrMgntEventFilterAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AcrMgntEventFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AcrMgntEventFilter) MarshalJSON() ([]byte, error) {
	if src.AcrMgntEventFilterAnyOf != nil {
		return json.Marshal(&src.AcrMgntEventFilterAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAcrMgntEventFilter struct {
	value *AcrMgntEventFilter
	isSet bool
}

func (v NullableAcrMgntEventFilter) Get() *AcrMgntEventFilter {
	return v.value
}

func (v *NullableAcrMgntEventFilter) Set(val *AcrMgntEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrMgntEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrMgntEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrMgntEventFilter(val *AcrMgntEventFilter) *NullableAcrMgntEventFilter {
	return &NullableAcrMgntEventFilter{value: val, isSet: true}
}

func (v NullableAcrMgntEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrMgntEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


