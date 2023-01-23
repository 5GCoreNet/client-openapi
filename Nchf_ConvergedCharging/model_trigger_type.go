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

// TriggerType struct for TriggerType
type TriggerType struct {
	TriggerTypeAnyOf *TriggerTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TriggerType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TriggerTypeAnyOf
	err = json.Unmarshal(data, &dst.TriggerTypeAnyOf);
	if err == nil {
		jsonTriggerTypeAnyOf, _ := json.Marshal(dst.TriggerTypeAnyOf)
		if string(jsonTriggerTypeAnyOf) == "{}" { // empty struct
			dst.TriggerTypeAnyOf = nil
		} else {
			return nil // data stored in dst.TriggerTypeAnyOf, return on the first match
		}
	} else {
		dst.TriggerTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TriggerType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TriggerType) MarshalJSON() ([]byte, error) {
	if src.TriggerTypeAnyOf != nil {
		return json.Marshal(&src.TriggerTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTriggerType struct {
	value *TriggerType
	isSet bool
}

func (v NullableTriggerType) Get() *TriggerType {
	return v.value
}

func (v *NullableTriggerType) Set(val *TriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerType(val *TriggerType) *NullableTriggerType {
	return &NullableTriggerType{value: val, isSet: true}
}

func (v NullableTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


