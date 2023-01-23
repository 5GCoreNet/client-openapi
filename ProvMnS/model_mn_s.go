/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// MnS - struct for MnS
type MnS struct {
	MnSOneOf *MnSOneOf
	MnSOneOf1 *MnSOneOf1
}

// MnSOneOfAsMnS is a convenience function that returns MnSOneOf wrapped in MnS
func MnSOneOfAsMnS(v *MnSOneOf) MnS {
	return MnS{
		MnSOneOf: v,
	}
}

// MnSOneOf1AsMnS is a convenience function that returns MnSOneOf1 wrapped in MnS
func MnSOneOf1AsMnS(v *MnSOneOf1) MnS {
	return MnS{
		MnSOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MnS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MnSOneOf
	err = newStrictDecoder(data).Decode(&dst.MnSOneOf)
	if err == nil {
		jsonMnSOneOf, _ := json.Marshal(dst.MnSOneOf)
		if string(jsonMnSOneOf) == "{}" { // empty struct
			dst.MnSOneOf = nil
		} else {
			match++
		}
	} else {
		dst.MnSOneOf = nil
	}

	// try to unmarshal data into MnSOneOf1
	err = newStrictDecoder(data).Decode(&dst.MnSOneOf1)
	if err == nil {
		jsonMnSOneOf1, _ := json.Marshal(dst.MnSOneOf1)
		if string(jsonMnSOneOf1) == "{}" { // empty struct
			dst.MnSOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.MnSOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MnSOneOf = nil
		dst.MnSOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MnS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MnS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MnS) MarshalJSON() ([]byte, error) {
	if src.MnSOneOf != nil {
		return json.Marshal(&src.MnSOneOf)
	}

	if src.MnSOneOf1 != nil {
		return json.Marshal(&src.MnSOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MnS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MnSOneOf != nil {
		return obj.MnSOneOf
	}

	if obj.MnSOneOf1 != nil {
		return obj.MnSOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableMnS struct {
	value *MnS
	isSet bool
}

func (v NullableMnS) Get() *MnS {
	return v.value
}

func (v *NullableMnS) Set(val *MnS) {
	v.value = val
	v.isSet = true
}

func (v NullableMnS) IsSet() bool {
	return v.isSet
}

func (v *NullableMnS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnS(val *MnS) *NullableMnS {
	return &NullableMnS{value: val, isSet: true}
}

func (v NullableMnS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


