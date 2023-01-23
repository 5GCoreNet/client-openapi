/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// UriScheme HTTP and HTTPS URI scheme.
type UriScheme struct {
	UriSchemeAnyOf *UriSchemeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UriScheme) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UriSchemeAnyOf
	err = json.Unmarshal(data, &dst.UriSchemeAnyOf);
	if err == nil {
		jsonUriSchemeAnyOf, _ := json.Marshal(dst.UriSchemeAnyOf)
		if string(jsonUriSchemeAnyOf) == "{}" { // empty struct
			dst.UriSchemeAnyOf = nil
		} else {
			return nil // data stored in dst.UriSchemeAnyOf, return on the first match
		}
	} else {
		dst.UriSchemeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UriScheme)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UriScheme) MarshalJSON() ([]byte, error) {
	if src.UriSchemeAnyOf != nil {
		return json.Marshal(&src.UriSchemeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUriScheme struct {
	value *UriScheme
	isSet bool
}

func (v NullableUriScheme) Get() *UriScheme {
	return v.value
}

func (v *NullableUriScheme) Set(val *UriScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableUriScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableUriScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUriScheme(val *UriScheme) *NullableUriScheme {
	return &NullableUriScheme{value: val, isSet: true}
}

func (v NullableUriScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUriScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


