/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// AssociationType struct for AssociationType
type AssociationType struct {
	AssociationTypeAnyOf *AssociationTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AssociationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AssociationTypeAnyOf
	err = json.Unmarshal(data, &dst.AssociationTypeAnyOf);
	if err == nil {
		jsonAssociationTypeAnyOf, _ := json.Marshal(dst.AssociationTypeAnyOf)
		if string(jsonAssociationTypeAnyOf) == "{}" { // empty struct
			dst.AssociationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AssociationTypeAnyOf, return on the first match
		}
	} else {
		dst.AssociationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AssociationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AssociationType) MarshalJSON() ([]byte, error) {
	if src.AssociationTypeAnyOf != nil {
		return json.Marshal(&src.AssociationTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAssociationType struct {
	value *AssociationType
	isSet bool
}

func (v NullableAssociationType) Get() *AssociationType {
	return v.value
}

func (v *NullableAssociationType) Set(val *AssociationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationType(val *AssociationType) *NullableAssociationType {
	return &NullableAssociationType{value: val, isSet: true}
}

func (v NullableAssociationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

