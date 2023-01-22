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

// SipDigestQop Quality of Protection for the SIP Digest authentication scheme
type SipDigestQop struct {
	SipDigestQopAnyOf *SipDigestQopAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SipDigestQop) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SipDigestQopAnyOf
	err = json.Unmarshal(data, &dst.SipDigestQopAnyOf);
	if err == nil {
		jsonSipDigestQopAnyOf, _ := json.Marshal(dst.SipDigestQopAnyOf)
		if string(jsonSipDigestQopAnyOf) == "{}" { // empty struct
			dst.SipDigestQopAnyOf = nil
		} else {
			return nil // data stored in dst.SipDigestQopAnyOf, return on the first match
		}
	} else {
		dst.SipDigestQopAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SipDigestQop)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SipDigestQop) MarshalJSON() ([]byte, error) {
	if src.SipDigestQopAnyOf != nil {
		return json.Marshal(&src.SipDigestQopAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSipDigestQop struct {
	value *SipDigestQop
	isSet bool
}

func (v NullableSipDigestQop) Get() *SipDigestQop {
	return v.value
}

func (v *NullableSipDigestQop) Set(val *SipDigestQop) {
	v.value = val
	v.isSet = true
}

func (v NullableSipDigestQop) IsSet() bool {
	return v.isSet
}

func (v *NullableSipDigestQop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipDigestQop(val *SipDigestQop) *NullableSipDigestQop {
	return &NullableSipDigestQop{value: val, isSet: true}
}

func (v NullableSipDigestQop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipDigestQop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


