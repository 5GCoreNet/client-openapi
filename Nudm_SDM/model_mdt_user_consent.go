/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// MdtUserConsent struct for MdtUserConsent
type MdtUserConsent struct {
	UserConsentAnyOf *UserConsentAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MdtUserConsent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UserConsentAnyOf
	err = json.Unmarshal(data, &dst.UserConsentAnyOf);
	if err == nil {
		jsonUserConsentAnyOf, _ := json.Marshal(dst.UserConsentAnyOf)
		if string(jsonUserConsentAnyOf) == "{}" { // empty struct
			dst.UserConsentAnyOf = nil
		} else {
			return nil // data stored in dst.UserConsentAnyOf, return on the first match
		}
	} else {
		dst.UserConsentAnyOf = nil
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
	if src.UserConsentAnyOf != nil {
		return json.Marshal(&src.UserConsentAnyOf)
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


