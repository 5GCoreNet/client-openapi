/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// MediaType Indicates the media type of a media component.
type MediaType struct {
	MediaTypeAnyOf *MediaTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MediaType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MediaTypeAnyOf
	err = json.Unmarshal(data, &dst.MediaTypeAnyOf);
	if err == nil {
		jsonMediaTypeAnyOf, _ := json.Marshal(dst.MediaTypeAnyOf)
		if string(jsonMediaTypeAnyOf) == "{}" { // empty struct
			dst.MediaTypeAnyOf = nil
		} else {
			return nil // data stored in dst.MediaTypeAnyOf, return on the first match
		}
	} else {
		dst.MediaTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MediaType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MediaType) MarshalJSON() ([]byte, error) {
	if src.MediaTypeAnyOf != nil {
		return json.Marshal(&src.MediaTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMediaType struct {
	value *MediaType
	isSet bool
}

func (v NullableMediaType) Get() *MediaType {
	return v.value
}

func (v *NullableMediaType) Set(val *MediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaType(val *MediaType) *NullableMediaType {
	return &NullableMediaType{value: val, isSet: true}
}

func (v NullableMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


