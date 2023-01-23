/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// PeriodicCommunicationIndicator Indicates the Periodic Communication Indicator
type PeriodicCommunicationIndicator struct {
	PeriodicCommunicationIndicatorAnyOf *PeriodicCommunicationIndicatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PeriodicCommunicationIndicator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PeriodicCommunicationIndicatorAnyOf
	err = json.Unmarshal(data, &dst.PeriodicCommunicationIndicatorAnyOf);
	if err == nil {
		jsonPeriodicCommunicationIndicatorAnyOf, _ := json.Marshal(dst.PeriodicCommunicationIndicatorAnyOf)
		if string(jsonPeriodicCommunicationIndicatorAnyOf) == "{}" { // empty struct
			dst.PeriodicCommunicationIndicatorAnyOf = nil
		} else {
			return nil // data stored in dst.PeriodicCommunicationIndicatorAnyOf, return on the first match
		}
	} else {
		dst.PeriodicCommunicationIndicatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PeriodicCommunicationIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PeriodicCommunicationIndicator) MarshalJSON() ([]byte, error) {
	if src.PeriodicCommunicationIndicatorAnyOf != nil {
		return json.Marshal(&src.PeriodicCommunicationIndicatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePeriodicCommunicationIndicator struct {
	value *PeriodicCommunicationIndicator
	isSet bool
}

func (v NullablePeriodicCommunicationIndicator) Get() *PeriodicCommunicationIndicator {
	return v.value
}

func (v *NullablePeriodicCommunicationIndicator) Set(val *PeriodicCommunicationIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicCommunicationIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicCommunicationIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicCommunicationIndicator(val *PeriodicCommunicationIndicator) *NullablePeriodicCommunicationIndicator {
	return &NullablePeriodicCommunicationIndicator{value: val, isSet: true}
}

func (v NullablePeriodicCommunicationIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicCommunicationIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


