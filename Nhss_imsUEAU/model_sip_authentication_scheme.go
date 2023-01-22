/*
Nhss_imsUEAU

Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUEAU

import (
	"encoding/json"
	"fmt"
)

// SipAuthenticationScheme Authentication scheme to be used in the SIP authentication request message
type SipAuthenticationScheme struct {
	SipAuthenticationSchemeAnyOf *SipAuthenticationSchemeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SipAuthenticationScheme) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SipAuthenticationSchemeAnyOf
	err = json.Unmarshal(data, &dst.SipAuthenticationSchemeAnyOf);
	if err == nil {
		jsonSipAuthenticationSchemeAnyOf, _ := json.Marshal(dst.SipAuthenticationSchemeAnyOf)
		if string(jsonSipAuthenticationSchemeAnyOf) == "{}" { // empty struct
			dst.SipAuthenticationSchemeAnyOf = nil
		} else {
			return nil // data stored in dst.SipAuthenticationSchemeAnyOf, return on the first match
		}
	} else {
		dst.SipAuthenticationSchemeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SipAuthenticationScheme)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SipAuthenticationScheme) MarshalJSON() ([]byte, error) {
	if src.SipAuthenticationSchemeAnyOf != nil {
		return json.Marshal(&src.SipAuthenticationSchemeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSipAuthenticationScheme struct {
	value *SipAuthenticationScheme
	isSet bool
}

func (v NullableSipAuthenticationScheme) Get() *SipAuthenticationScheme {
	return v.value
}

func (v *NullableSipAuthenticationScheme) Set(val *SipAuthenticationScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableSipAuthenticationScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableSipAuthenticationScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipAuthenticationScheme(val *SipAuthenticationScheme) *NullableSipAuthenticationScheme {
	return &NullableSipAuthenticationScheme{value: val, isSet: true}
}

func (v NullableSipAuthenticationScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipAuthenticationScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

