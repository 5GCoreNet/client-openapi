/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// ProseDirectAllowed Indicates the 5G ProSe Direct services that can be authorised to use in the given PLMN for the UE.
type ProseDirectAllowed struct {
	ProseDirectAllowedAnyOf *ProseDirectAllowedAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProseDirectAllowed) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ProseDirectAllowedAnyOf
	err = json.Unmarshal(data, &dst.ProseDirectAllowedAnyOf);
	if err == nil {
		jsonProseDirectAllowedAnyOf, _ := json.Marshal(dst.ProseDirectAllowedAnyOf)
		if string(jsonProseDirectAllowedAnyOf) == "{}" { // empty struct
			dst.ProseDirectAllowedAnyOf = nil
		} else {
			return nil // data stored in dst.ProseDirectAllowedAnyOf, return on the first match
		}
	} else {
		dst.ProseDirectAllowedAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ProseDirectAllowed)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProseDirectAllowed) MarshalJSON() ([]byte, error) {
	if src.ProseDirectAllowedAnyOf != nil {
		return json.Marshal(&src.ProseDirectAllowedAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProseDirectAllowed struct {
	value *ProseDirectAllowed
	isSet bool
}

func (v NullableProseDirectAllowed) Get() *ProseDirectAllowed {
	return v.value
}

func (v *NullableProseDirectAllowed) Set(val *ProseDirectAllowed) {
	v.value = val
	v.isSet = true
}

func (v NullableProseDirectAllowed) IsSet() bool {
	return v.isSet
}

func (v *NullableProseDirectAllowed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseDirectAllowed(val *ProseDirectAllowed) *NullableProseDirectAllowed {
	return &NullableProseDirectAllowed{value: val, isSet: true}
}

func (v NullableProseDirectAllowed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseDirectAllowed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


