/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// DistributionMethod - Possible values are: - OBJECT: Indicates the Object Distribution Method. - PACKET: Indicates the Packet Distribution Method. 
type DistributionMethod struct {
	DistributionMethodOneOf *DistributionMethodOneOf
	String *string
}

// DistributionMethodOneOfAsDistributionMethod is a convenience function that returns DistributionMethodOneOf wrapped in DistributionMethod
func DistributionMethodOneOfAsDistributionMethod(v *DistributionMethodOneOf) DistributionMethod {
	return DistributionMethod{
		DistributionMethodOneOf: v,
	}
}

// stringAsDistributionMethod is a convenience function that returns string wrapped in DistributionMethod
func StringAsDistributionMethod(v *string) DistributionMethod {
	return DistributionMethod{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DistributionMethod) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DistributionMethodOneOf
	err = newStrictDecoder(data).Decode(&dst.DistributionMethodOneOf)
	if err == nil {
		jsonDistributionMethodOneOf, _ := json.Marshal(dst.DistributionMethodOneOf)
		if string(jsonDistributionMethodOneOf) == "{}" { // empty struct
			dst.DistributionMethodOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DistributionMethodOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DistributionMethodOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DistributionMethod)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DistributionMethod)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DistributionMethod) MarshalJSON() ([]byte, error) {
	if src.DistributionMethodOneOf != nil {
		return json.Marshal(&src.DistributionMethodOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DistributionMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DistributionMethodOneOf != nil {
		return obj.DistributionMethodOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDistributionMethod struct {
	value *DistributionMethod
	isSet bool
}

func (v NullableDistributionMethod) Get() *DistributionMethod {
	return v.value
}

func (v *NullableDistributionMethod) Set(val *DistributionMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionMethod(val *DistributionMethod) *NullableDistributionMethod {
	return &NullableDistributionMethod{value: val, isSet: true}
}

func (v NullableDistributionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


