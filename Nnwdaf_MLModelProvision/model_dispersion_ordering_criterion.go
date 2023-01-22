/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// DispersionOrderingCriterion Possible values are: - TIME_SLOT_START: Indicates the order of time slot start. - DISPERSION: Indicates the order of data/transaction dispersion. - CLASSIFICATION: Indicates the order of data/transaction classification. - RANKING: Indicates the order of data/transaction ranking. - PERCENTILE_RANKING: Indicates the order of data/transaction percentile ranking. 
type DispersionOrderingCriterion struct {
	DispersionOrderingCriterionAnyOf *DispersionOrderingCriterionAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DispersionOrderingCriterion) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DispersionOrderingCriterionAnyOf
	err = json.Unmarshal(data, &dst.DispersionOrderingCriterionAnyOf);
	if err == nil {
		jsonDispersionOrderingCriterionAnyOf, _ := json.Marshal(dst.DispersionOrderingCriterionAnyOf)
		if string(jsonDispersionOrderingCriterionAnyOf) == "{}" { // empty struct
			dst.DispersionOrderingCriterionAnyOf = nil
		} else {
			return nil // data stored in dst.DispersionOrderingCriterionAnyOf, return on the first match
		}
	} else {
		dst.DispersionOrderingCriterionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DispersionOrderingCriterion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DispersionOrderingCriterion) MarshalJSON() ([]byte, error) {
	if src.DispersionOrderingCriterionAnyOf != nil {
		return json.Marshal(&src.DispersionOrderingCriterionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDispersionOrderingCriterion struct {
	value *DispersionOrderingCriterion
	isSet bool
}

func (v NullableDispersionOrderingCriterion) Get() *DispersionOrderingCriterion {
	return v.value
}

func (v *NullableDispersionOrderingCriterion) Set(val *DispersionOrderingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionOrderingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionOrderingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionOrderingCriterion(val *DispersionOrderingCriterion) *NullableDispersionOrderingCriterion {
	return &NullableDispersionOrderingCriterion{value: val, isSet: true}
}

func (v NullableDispersionOrderingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionOrderingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


