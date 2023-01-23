/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRmmNrMdt The enumeration CollectionPeriodRmmNrMdt defines Collection period for RRM measurements NR for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.19-1 
type CollectionPeriodRmmNrMdt struct {
	CollectionPeriodRmmNrMdtAnyOf *CollectionPeriodRmmNrMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollectionPeriodRmmNrMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CollectionPeriodRmmNrMdtAnyOf
	err = json.Unmarshal(data, &dst.CollectionPeriodRmmNrMdtAnyOf);
	if err == nil {
		jsonCollectionPeriodRmmNrMdtAnyOf, _ := json.Marshal(dst.CollectionPeriodRmmNrMdtAnyOf)
		if string(jsonCollectionPeriodRmmNrMdtAnyOf) == "{}" { // empty struct
			dst.CollectionPeriodRmmNrMdtAnyOf = nil
		} else {
			return nil // data stored in dst.CollectionPeriodRmmNrMdtAnyOf, return on the first match
		}
	} else {
		dst.CollectionPeriodRmmNrMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CollectionPeriodRmmNrMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollectionPeriodRmmNrMdt) MarshalJSON() ([]byte, error) {
	if src.CollectionPeriodRmmNrMdtAnyOf != nil {
		return json.Marshal(&src.CollectionPeriodRmmNrMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollectionPeriodRmmNrMdt struct {
	value *CollectionPeriodRmmNrMdt
	isSet bool
}

func (v NullableCollectionPeriodRmmNrMdt) Get() *CollectionPeriodRmmNrMdt {
	return v.value
}

func (v *NullableCollectionPeriodRmmNrMdt) Set(val *CollectionPeriodRmmNrMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRmmNrMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRmmNrMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRmmNrMdt(val *CollectionPeriodRmmNrMdt) *NullableCollectionPeriodRmmNrMdt {
	return &NullableCollectionPeriodRmmNrMdt{value: val, isSet: true}
}

func (v NullableCollectionPeriodRmmNrMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRmmNrMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


