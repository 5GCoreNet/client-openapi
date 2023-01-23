/*
3gpp-eas-deployment

API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EASDeployment

import (
	"encoding/json"
	"fmt"
)

// MatchingOperator the matching operation.
type MatchingOperator struct {
	MatchingOperatorAnyOf *MatchingOperatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MatchingOperator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MatchingOperatorAnyOf
	err = json.Unmarshal(data, &dst.MatchingOperatorAnyOf);
	if err == nil {
		jsonMatchingOperatorAnyOf, _ := json.Marshal(dst.MatchingOperatorAnyOf)
		if string(jsonMatchingOperatorAnyOf) == "{}" { // empty struct
			dst.MatchingOperatorAnyOf = nil
		} else {
			return nil // data stored in dst.MatchingOperatorAnyOf, return on the first match
		}
	} else {
		dst.MatchingOperatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MatchingOperator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MatchingOperator) MarshalJSON() ([]byte, error) {
	if src.MatchingOperatorAnyOf != nil {
		return json.Marshal(&src.MatchingOperatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMatchingOperator struct {
	value *MatchingOperator
	isSet bool
}

func (v NullableMatchingOperator) Get() *MatchingOperator {
	return v.value
}

func (v *NullableMatchingOperator) Set(val *MatchingOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingOperator(val *MatchingOperator) *NullableMatchingOperator {
	return &NullableMatchingOperator{value: val, isSet: true}
}

func (v NullableMatchingOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


