/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRmmLteMdt The enumeration CollectionPeriodRmmLteMdt defines Collection period for RRM measurements LTE for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.15-1. 
type CollectionPeriodRmmLteMdt struct {
	CollectionPeriodRmmLteMdtAnyOf *CollectionPeriodRmmLteMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollectionPeriodRmmLteMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CollectionPeriodRmmLteMdtAnyOf
	err = json.Unmarshal(data, &dst.CollectionPeriodRmmLteMdtAnyOf);
	if err == nil {
		jsonCollectionPeriodRmmLteMdtAnyOf, _ := json.Marshal(dst.CollectionPeriodRmmLteMdtAnyOf)
		if string(jsonCollectionPeriodRmmLteMdtAnyOf) == "{}" { // empty struct
			dst.CollectionPeriodRmmLteMdtAnyOf = nil
		} else {
			return nil // data stored in dst.CollectionPeriodRmmLteMdtAnyOf, return on the first match
		}
	} else {
		dst.CollectionPeriodRmmLteMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CollectionPeriodRmmLteMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollectionPeriodRmmLteMdt) MarshalJSON() ([]byte, error) {
	if src.CollectionPeriodRmmLteMdtAnyOf != nil {
		return json.Marshal(&src.CollectionPeriodRmmLteMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollectionPeriodRmmLteMdt struct {
	value *CollectionPeriodRmmLteMdt
	isSet bool
}

func (v NullableCollectionPeriodRmmLteMdt) Get() *CollectionPeriodRmmLteMdt {
	return v.value
}

func (v *NullableCollectionPeriodRmmLteMdt) Set(val *CollectionPeriodRmmLteMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRmmLteMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRmmLteMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRmmLteMdt(val *CollectionPeriodRmmLteMdt) *NullableCollectionPeriodRmmLteMdt {
	return &NullableCollectionPeriodRmmLteMdt{value: val, isSet: true}
}

func (v NullableCollectionPeriodRmmLteMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRmmLteMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


