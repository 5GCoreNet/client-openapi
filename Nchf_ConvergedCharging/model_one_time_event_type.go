/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// OneTimeEventType struct for OneTimeEventType
type OneTimeEventType struct {
	OneTimeEventTypeAnyOf *OneTimeEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OneTimeEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OneTimeEventTypeAnyOf
	err = json.Unmarshal(data, &dst.OneTimeEventTypeAnyOf);
	if err == nil {
		jsonOneTimeEventTypeAnyOf, _ := json.Marshal(dst.OneTimeEventTypeAnyOf)
		if string(jsonOneTimeEventTypeAnyOf) == "{}" { // empty struct
			dst.OneTimeEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.OneTimeEventTypeAnyOf, return on the first match
		}
	} else {
		dst.OneTimeEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(OneTimeEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OneTimeEventType) MarshalJSON() ([]byte, error) {
	if src.OneTimeEventTypeAnyOf != nil {
		return json.Marshal(&src.OneTimeEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOneTimeEventType struct {
	value *OneTimeEventType
	isSet bool
}

func (v NullableOneTimeEventType) Get() *OneTimeEventType {
	return v.value
}

func (v *NullableOneTimeEventType) Set(val *OneTimeEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableOneTimeEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableOneTimeEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneTimeEventType(val *OneTimeEventType) *NullableOneTimeEventType {
	return &NullableOneTimeEventType{value: val, isSet: true}
}

func (v NullableOneTimeEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneTimeEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


