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

// AccessTransferType struct for AccessTransferType
type AccessTransferType struct {
	AccessTransferTypeAnyOf *AccessTransferTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessTransferType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccessTransferTypeAnyOf
	err = json.Unmarshal(data, &dst.AccessTransferTypeAnyOf);
	if err == nil {
		jsonAccessTransferTypeAnyOf, _ := json.Marshal(dst.AccessTransferTypeAnyOf)
		if string(jsonAccessTransferTypeAnyOf) == "{}" { // empty struct
			dst.AccessTransferTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AccessTransferTypeAnyOf, return on the first match
		}
	} else {
		dst.AccessTransferTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessTransferType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessTransferType) MarshalJSON() ([]byte, error) {
	if src.AccessTransferTypeAnyOf != nil {
		return json.Marshal(&src.AccessTransferTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessTransferType struct {
	value *AccessTransferType
	isSet bool
}

func (v NullableAccessTransferType) Get() *AccessTransferType {
	return v.value
}

func (v *NullableAccessTransferType) Set(val *AccessTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTransferType(val *AccessTransferType) *NullableAccessTransferType {
	return &NullableAccessTransferType{value: val, isSet: true}
}

func (v NullableAccessTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


