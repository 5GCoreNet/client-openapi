/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// CoreNetworkTypeRm It contains the Core Network type 5GC or EPC but with the OpenAPI 'nullable: true' property.  
type CoreNetworkTypeRm struct {
	CoreNetworkType *CoreNetworkType
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CoreNetworkTypeRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CoreNetworkType
	err = json.Unmarshal(data, &dst.CoreNetworkType);
	if err == nil {
		jsonCoreNetworkType, _ := json.Marshal(dst.CoreNetworkType)
		if string(jsonCoreNetworkType) == "{}" { // empty struct
			dst.CoreNetworkType = nil
		} else {
			return nil // data stored in dst.CoreNetworkType, return on the first match
		}
	} else {
		dst.CoreNetworkType = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CoreNetworkTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CoreNetworkTypeRm) MarshalJSON() ([]byte, error) {
	if src.CoreNetworkType != nil {
		return json.Marshal(&src.CoreNetworkType)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCoreNetworkTypeRm struct {
	value *CoreNetworkTypeRm
	isSet bool
}

func (v NullableCoreNetworkTypeRm) Get() *CoreNetworkTypeRm {
	return v.value
}

func (v *NullableCoreNetworkTypeRm) Set(val *CoreNetworkTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNetworkTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNetworkTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNetworkTypeRm(val *CoreNetworkTypeRm) *NullableCoreNetworkTypeRm {
	return &NullableCoreNetworkTypeRm{value: val, isSet: true}
}

func (v NullableCoreNetworkTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNetworkTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


