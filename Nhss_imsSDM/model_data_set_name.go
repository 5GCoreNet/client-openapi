/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// DataSetName A set of data within the IMS User Profile
type DataSetName struct {
	DataSetNameAnyOf *DataSetNameAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataSetName) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DataSetNameAnyOf
	err = json.Unmarshal(data, &dst.DataSetNameAnyOf);
	if err == nil {
		jsonDataSetNameAnyOf, _ := json.Marshal(dst.DataSetNameAnyOf)
		if string(jsonDataSetNameAnyOf) == "{}" { // empty struct
			dst.DataSetNameAnyOf = nil
		} else {
			return nil // data stored in dst.DataSetNameAnyOf, return on the first match
		}
	} else {
		dst.DataSetNameAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DataSetName)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataSetName) MarshalJSON() ([]byte, error) {
	if src.DataSetNameAnyOf != nil {
		return json.Marshal(&src.DataSetNameAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataSetName struct {
	value *DataSetName
	isSet bool
}

func (v NullableDataSetName) Get() *DataSetName {
	return v.value
}

func (v *NullableDataSetName) Set(val *DataSetName) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetName) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetName(val *DataSetName) *NullableDataSetName {
	return &NullableDataSetName{value: val, isSet: true}
}

func (v NullableDataSetName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


