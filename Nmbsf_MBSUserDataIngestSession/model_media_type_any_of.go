/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// MediaTypeAnyOf the model 'MediaTypeAnyOf'
type MediaTypeAnyOf string

// List of MediaType_anyOf
const (
	AUDIO MediaTypeAnyOf = "AUDIO"
	VIDEO MediaTypeAnyOf = "VIDEO"
	DATA MediaTypeAnyOf = "DATA"
	APPLICATION MediaTypeAnyOf = "APPLICATION"
	CONTROL MediaTypeAnyOf = "CONTROL"
	TEXT MediaTypeAnyOf = "TEXT"
	MESSAGE MediaTypeAnyOf = "MESSAGE"
	OTHER MediaTypeAnyOf = "OTHER"
)

// All allowed values of MediaTypeAnyOf enum
var AllowedMediaTypeAnyOfEnumValues = []MediaTypeAnyOf{
	"AUDIO",
	"VIDEO",
	"DATA",
	"APPLICATION",
	"CONTROL",
	"TEXT",
	"MESSAGE",
	"OTHER",
}

func (v *MediaTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaTypeAnyOf(value)
	for _, existing := range AllowedMediaTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaTypeAnyOf", value)
}

// NewMediaTypeAnyOfFromValue returns a pointer to a valid MediaTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaTypeAnyOfFromValue(v string) (*MediaTypeAnyOf, error) {
	ev := MediaTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaTypeAnyOf: valid values are %v", v, AllowedMediaTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedMediaTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaType_anyOf value
func (v MediaTypeAnyOf) Ptr() *MediaTypeAnyOf {
	return &v
}

type NullableMediaTypeAnyOf struct {
	value *MediaTypeAnyOf
	isSet bool
}

func (v NullableMediaTypeAnyOf) Get() *MediaTypeAnyOf {
	return v.value
}

func (v *NullableMediaTypeAnyOf) Set(val *MediaTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaTypeAnyOf(val *MediaTypeAnyOf) *NullableMediaTypeAnyOf {
	return &NullableMediaTypeAnyOf{value: val, isSet: true}
}

func (v NullableMediaTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

