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

// SMAddresseeType struct for SMAddresseeType
type SMAddresseeType struct {
	SMAddresseeTypeAnyOf *SMAddresseeTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SMAddresseeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SMAddresseeTypeAnyOf
	err = json.Unmarshal(data, &dst.SMAddresseeTypeAnyOf);
	if err == nil {
		jsonSMAddresseeTypeAnyOf, _ := json.Marshal(dst.SMAddresseeTypeAnyOf)
		if string(jsonSMAddresseeTypeAnyOf) == "{}" { // empty struct
			dst.SMAddresseeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.SMAddresseeTypeAnyOf, return on the first match
		}
	} else {
		dst.SMAddresseeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SMAddresseeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SMAddresseeType) MarshalJSON() ([]byte, error) {
	if src.SMAddresseeTypeAnyOf != nil {
		return json.Marshal(&src.SMAddresseeTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSMAddresseeType struct {
	value *SMAddresseeType
	isSet bool
}

func (v NullableSMAddresseeType) Get() *SMAddresseeType {
	return v.value
}

func (v *NullableSMAddresseeType) Set(val *SMAddresseeType) {
	v.value = val
	v.isSet = true
}

func (v NullableSMAddresseeType) IsSet() bool {
	return v.isSet
}

func (v *NullableSMAddresseeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMAddresseeType(val *SMAddresseeType) *NullableSMAddresseeType {
	return &NullableSMAddresseeType{value: val, isSet: true}
}

func (v NullableSMAddresseeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMAddresseeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


