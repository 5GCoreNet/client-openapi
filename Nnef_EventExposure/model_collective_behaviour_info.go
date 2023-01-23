/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EventExposure

import (
	"encoding/json"
	"fmt"
)

// CollectiveBehaviourInfo - Contains the collective behaviour information to be reported to the subscriber.
type CollectiveBehaviourInfo struct {
	Interface{} *interface{}
}

// interface{}AsCollectiveBehaviourInfo is a convenience function that returns interface{} wrapped in CollectiveBehaviourInfo
func Interface{}AsCollectiveBehaviourInfo(v *interface{}) CollectiveBehaviourInfo {
	return CollectiveBehaviourInfo{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CollectiveBehaviourInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CollectiveBehaviourInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CollectiveBehaviourInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CollectiveBehaviourInfo) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CollectiveBehaviourInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableCollectiveBehaviourInfo struct {
	value *CollectiveBehaviourInfo
	isSet bool
}

func (v NullableCollectiveBehaviourInfo) Get() *CollectiveBehaviourInfo {
	return v.value
}

func (v *NullableCollectiveBehaviourInfo) Set(val *CollectiveBehaviourInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectiveBehaviourInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectiveBehaviourInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectiveBehaviourInfo(val *CollectiveBehaviourInfo) *NullableCollectiveBehaviourInfo {
	return &NullableCollectiveBehaviourInfo{value: val, isSet: true}
}

func (v NullableCollectiveBehaviourInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectiveBehaviourInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


