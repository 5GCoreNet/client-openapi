/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// DatasetStatisticalProperty Possible values are: - UNIFORM_DIST_DATA: Indicates the use of data samples that are uniformly distributed according to the different aspects of the requested analytics. - NO_OUTLIERS: Indicates that the data samples shall disregard data samples that are at the extreme boundaries of the value range. 
type DatasetStatisticalProperty struct {
	DatasetStatisticalPropertyAnyOf *DatasetStatisticalPropertyAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatasetStatisticalProperty) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DatasetStatisticalPropertyAnyOf
	err = json.Unmarshal(data, &dst.DatasetStatisticalPropertyAnyOf);
	if err == nil {
		jsonDatasetStatisticalPropertyAnyOf, _ := json.Marshal(dst.DatasetStatisticalPropertyAnyOf)
		if string(jsonDatasetStatisticalPropertyAnyOf) == "{}" { // empty struct
			dst.DatasetStatisticalPropertyAnyOf = nil
		} else {
			return nil // data stored in dst.DatasetStatisticalPropertyAnyOf, return on the first match
		}
	} else {
		dst.DatasetStatisticalPropertyAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DatasetStatisticalProperty)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatasetStatisticalProperty) MarshalJSON() ([]byte, error) {
	if src.DatasetStatisticalPropertyAnyOf != nil {
		return json.Marshal(&src.DatasetStatisticalPropertyAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatasetStatisticalProperty struct {
	value *DatasetStatisticalProperty
	isSet bool
}

func (v NullableDatasetStatisticalProperty) Get() *DatasetStatisticalProperty {
	return v.value
}

func (v *NullableDatasetStatisticalProperty) Set(val *DatasetStatisticalProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetStatisticalProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetStatisticalProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetStatisticalProperty(val *DatasetStatisticalProperty) *NullableDatasetStatisticalProperty {
	return &NullableDatasetStatisticalProperty{value: val, isSet: true}
}

func (v NullableDatasetStatisticalProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetStatisticalProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


