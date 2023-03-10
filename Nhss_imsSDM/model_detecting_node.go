/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// DetectingNode Represents the type of serving node that detected the reachability of the UE
type DetectingNode struct {
	DetectingNodeAnyOf *DetectingNodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DetectingNode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DetectingNodeAnyOf
	err = json.Unmarshal(data, &dst.DetectingNodeAnyOf);
	if err == nil {
		jsonDetectingNodeAnyOf, _ := json.Marshal(dst.DetectingNodeAnyOf)
		if string(jsonDetectingNodeAnyOf) == "{}" { // empty struct
			dst.DetectingNodeAnyOf = nil
		} else {
			return nil // data stored in dst.DetectingNodeAnyOf, return on the first match
		}
	} else {
		dst.DetectingNodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DetectingNode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DetectingNode) MarshalJSON() ([]byte, error) {
	if src.DetectingNodeAnyOf != nil {
		return json.Marshal(&src.DetectingNodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDetectingNode struct {
	value *DetectingNode
	isSet bool
}

func (v NullableDetectingNode) Get() *DetectingNode {
	return v.value
}

func (v *NullableDetectingNode) Set(val *DetectingNode) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectingNode) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectingNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectingNode(val *DetectingNode) *NullableDetectingNode {
	return &NullableDetectingNode{value: val, isSet: true}
}

func (v NullableDetectingNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectingNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


