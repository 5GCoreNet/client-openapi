/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// SummarizationAttribute Possible values are: - SPACING: Average and variance of the time interval separating two consecutive occurrences of the same event and parameter value, or periodicity for periodic reporting. - DURATION: Average and variance of the time for which the parameter value applies. - OCCURRENCES: Number of countable occurrences for the parameter. - AVG_VAR: Average and variance of the parameter. - FREQ_VAL: Most and least frequent values. - MIN_MAX: Maximum and minimum parameter values. 
type SummarizationAttribute struct {
	SummarizationAttributeAnyOf *SummarizationAttributeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SummarizationAttribute) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SummarizationAttributeAnyOf
	err = json.Unmarshal(data, &dst.SummarizationAttributeAnyOf);
	if err == nil {
		jsonSummarizationAttributeAnyOf, _ := json.Marshal(dst.SummarizationAttributeAnyOf)
		if string(jsonSummarizationAttributeAnyOf) == "{}" { // empty struct
			dst.SummarizationAttributeAnyOf = nil
		} else {
			return nil // data stored in dst.SummarizationAttributeAnyOf, return on the first match
		}
	} else {
		dst.SummarizationAttributeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SummarizationAttribute)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SummarizationAttribute) MarshalJSON() ([]byte, error) {
	if src.SummarizationAttributeAnyOf != nil {
		return json.Marshal(&src.SummarizationAttributeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSummarizationAttribute struct {
	value *SummarizationAttribute
	isSet bool
}

func (v NullableSummarizationAttribute) Get() *SummarizationAttribute {
	return v.value
}

func (v *NullableSummarizationAttribute) Set(val *SummarizationAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSummarizationAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSummarizationAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummarizationAttribute(val *SummarizationAttribute) *NullableSummarizationAttribute {
	return &NullableSummarizationAttribute{value: val, isSet: true}
}

func (v NullableSummarizationAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummarizationAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


