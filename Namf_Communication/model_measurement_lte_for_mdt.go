/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// MeasurementLteForMdt The enumeration MeasurementLteForMdt defines Measurements used for MDT in LTE in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.5-1. 
type MeasurementLteForMdt struct {
	MeasurementLteForMdtAnyOf *MeasurementLteForMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeasurementLteForMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MeasurementLteForMdtAnyOf
	err = json.Unmarshal(data, &dst.MeasurementLteForMdtAnyOf);
	if err == nil {
		jsonMeasurementLteForMdtAnyOf, _ := json.Marshal(dst.MeasurementLteForMdtAnyOf)
		if string(jsonMeasurementLteForMdtAnyOf) == "{}" { // empty struct
			dst.MeasurementLteForMdtAnyOf = nil
		} else {
			return nil // data stored in dst.MeasurementLteForMdtAnyOf, return on the first match
		}
	} else {
		dst.MeasurementLteForMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MeasurementLteForMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeasurementLteForMdt) MarshalJSON() ([]byte, error) {
	if src.MeasurementLteForMdtAnyOf != nil {
		return json.Marshal(&src.MeasurementLteForMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeasurementLteForMdt struct {
	value *MeasurementLteForMdt
	isSet bool
}

func (v NullableMeasurementLteForMdt) Get() *MeasurementLteForMdt {
	return v.value
}

func (v *NullableMeasurementLteForMdt) Set(val *MeasurementLteForMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementLteForMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementLteForMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementLteForMdt(val *MeasurementLteForMdt) *NullableMeasurementLteForMdt {
	return &NullableMeasurementLteForMdt{value: val, isSet: true}
}

func (v NullableMeasurementLteForMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementLteForMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


