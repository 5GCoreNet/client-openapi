/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
	"fmt"
)

// NFInfo - struct for NFInfo
type NFInfo struct {
	AusfInfo *AusfInfo
}

// AusfInfoAsNFInfo is a convenience function that returns AusfInfo wrapped in NFInfo
func AusfInfoAsNFInfo(v *AusfInfo) NFInfo {
	return NFInfo{
		AusfInfo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NFInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AusfInfo
	err = newStrictDecoder(data).Decode(&dst.AusfInfo)
	if err == nil {
		jsonAusfInfo, _ := json.Marshal(dst.AusfInfo)
		if string(jsonAusfInfo) == "{}" { // empty struct
			dst.AusfInfo = nil
		} else {
			match++
		}
	} else {
		dst.AusfInfo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AusfInfo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NFInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NFInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NFInfo) MarshalJSON() ([]byte, error) {
	if src.AusfInfo != nil {
		return json.Marshal(&src.AusfInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NFInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AusfInfo != nil {
		return obj.AusfInfo
	}

	// all schemas are nil
	return nil
}

type NullableNFInfo struct {
	value *NFInfo
	isSet bool
}

func (v NullableNFInfo) Get() *NFInfo {
	return v.value
}

func (v *NullableNFInfo) Set(val *NFInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNFInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNFInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFInfo(val *NFInfo) *NullableNFInfo {
	return &NullableNFInfo{value: val, isSet: true}
}

func (v NullableNFInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


