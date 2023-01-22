/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
	"fmt"
)

// SliceSimultaneousUse the model 'SliceSimultaneousUse'
type SliceSimultaneousUse string

// List of SliceSimultaneousUse
const (
	ZERO SliceSimultaneousUse = "ZERO"
	ONE SliceSimultaneousUse = "ONE"
	TWO SliceSimultaneousUse = "TWO"
	THREE SliceSimultaneousUse = "THREE"
	FOUR SliceSimultaneousUse = "FOUR"
)

// All allowed values of SliceSimultaneousUse enum
var AllowedSliceSimultaneousUseEnumValues = []SliceSimultaneousUse{
	"ZERO",
	"ONE",
	"TWO",
	"THREE",
	"FOUR",
}

func (v *SliceSimultaneousUse) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SliceSimultaneousUse(value)
	for _, existing := range AllowedSliceSimultaneousUseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SliceSimultaneousUse", value)
}

// NewSliceSimultaneousUseFromValue returns a pointer to a valid SliceSimultaneousUse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSliceSimultaneousUseFromValue(v string) (*SliceSimultaneousUse, error) {
	ev := SliceSimultaneousUse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SliceSimultaneousUse: valid values are %v", v, AllowedSliceSimultaneousUseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SliceSimultaneousUse) IsValid() bool {
	for _, existing := range AllowedSliceSimultaneousUseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SliceSimultaneousUse value
func (v SliceSimultaneousUse) Ptr() *SliceSimultaneousUse {
	return &v
}

type NullableSliceSimultaneousUse struct {
	value *SliceSimultaneousUse
	isSet bool
}

func (v NullableSliceSimultaneousUse) Get() *SliceSimultaneousUse {
	return v.value
}

func (v *NullableSliceSimultaneousUse) Set(val *SliceSimultaneousUse) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceSimultaneousUse) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceSimultaneousUse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceSimultaneousUse(val *SliceSimultaneousUse) *NullableSliceSimultaneousUse {
	return &NullableSliceSimultaneousUse{value: val, isSet: true}
}

func (v NullableSliceSimultaneousUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceSimultaneousUse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

