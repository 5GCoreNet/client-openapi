/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PrivacyCheckRelatedAction struct for PrivacyCheckRelatedAction
type PrivacyCheckRelatedAction struct {
	PrivacyCheckRelatedActionAnyOf *PrivacyCheckRelatedActionAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PrivacyCheckRelatedAction) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PrivacyCheckRelatedActionAnyOf
	err = json.Unmarshal(data, &dst.PrivacyCheckRelatedActionAnyOf);
	if err == nil {
		jsonPrivacyCheckRelatedActionAnyOf, _ := json.Marshal(dst.PrivacyCheckRelatedActionAnyOf)
		if string(jsonPrivacyCheckRelatedActionAnyOf) == "{}" { // empty struct
			dst.PrivacyCheckRelatedActionAnyOf = nil
		} else {
			return nil // data stored in dst.PrivacyCheckRelatedActionAnyOf, return on the first match
		}
	} else {
		dst.PrivacyCheckRelatedActionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PrivacyCheckRelatedAction)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PrivacyCheckRelatedAction) MarshalJSON() ([]byte, error) {
	if src.PrivacyCheckRelatedActionAnyOf != nil {
		return json.Marshal(&src.PrivacyCheckRelatedActionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePrivacyCheckRelatedAction struct {
	value *PrivacyCheckRelatedAction
	isSet bool
}

func (v NullablePrivacyCheckRelatedAction) Get() *PrivacyCheckRelatedAction {
	return v.value
}

func (v *NullablePrivacyCheckRelatedAction) Set(val *PrivacyCheckRelatedAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivacyCheckRelatedAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivacyCheckRelatedAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivacyCheckRelatedAction(val *PrivacyCheckRelatedAction) *NullablePrivacyCheckRelatedAction {
	return &NullablePrivacyCheckRelatedAction{value: val, isSet: true}
}

func (v NullablePrivacyCheckRelatedAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivacyCheckRelatedAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


