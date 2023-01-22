/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// AdrfDataType Possible values are: - HISTORICAL_ANALYTICS: Indicates that historical analytics are stored in the ADRF. - HISTORICAL_DATA: Indicates that historical data are stored in the ADRF. 
type AdrfDataType struct {
	AdrfDataTypeAnyOf *AdrfDataTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AdrfDataType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AdrfDataTypeAnyOf
	err = json.Unmarshal(data, &dst.AdrfDataTypeAnyOf);
	if err == nil {
		jsonAdrfDataTypeAnyOf, _ := json.Marshal(dst.AdrfDataTypeAnyOf)
		if string(jsonAdrfDataTypeAnyOf) == "{}" { // empty struct
			dst.AdrfDataTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AdrfDataTypeAnyOf, return on the first match
		}
	} else {
		dst.AdrfDataTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AdrfDataType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AdrfDataType) MarshalJSON() ([]byte, error) {
	if src.AdrfDataTypeAnyOf != nil {
		return json.Marshal(&src.AdrfDataTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAdrfDataType struct {
	value *AdrfDataType
	isSet bool
}

func (v NullableAdrfDataType) Get() *AdrfDataType {
	return v.value
}

func (v *NullableAdrfDataType) Set(val *AdrfDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdrfDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdrfDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdrfDataType(val *AdrfDataType) *NullableAdrfDataType {
	return &NullableAdrfDataType{value: val, isSet: true}
}

func (v NullableAdrfDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdrfDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


