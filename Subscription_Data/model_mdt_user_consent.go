/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// MdtUserConsent struct for MdtUserConsent
type MdtUserConsent struct {
	MdtUserConsentAnyOf *MdtUserConsentAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MdtUserConsent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MdtUserConsentAnyOf
	err = json.Unmarshal(data, &dst.MdtUserConsentAnyOf);
	if err == nil {
		jsonMdtUserConsentAnyOf, _ := json.Marshal(dst.MdtUserConsentAnyOf)
		if string(jsonMdtUserConsentAnyOf) == "{}" { // empty struct
			dst.MdtUserConsentAnyOf = nil
		} else {
			return nil // data stored in dst.MdtUserConsentAnyOf, return on the first match
		}
	} else {
		dst.MdtUserConsentAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MdtUserConsent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MdtUserConsent) MarshalJSON() ([]byte, error) {
	if src.MdtUserConsentAnyOf != nil {
		return json.Marshal(&src.MdtUserConsentAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMdtUserConsent struct {
	value *MdtUserConsent
	isSet bool
}

func (v NullableMdtUserConsent) Get() *MdtUserConsent {
	return v.value
}

func (v *NullableMdtUserConsent) Set(val *MdtUserConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableMdtUserConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableMdtUserConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdtUserConsent(val *MdtUserConsent) *NullableMdtUserConsent {
	return &NullableMdtUserConsent{value: val, isSet: true}
}

func (v NullableMdtUserConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdtUserConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


