/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package GenericNrm

import (
	"encoding/json"
	"fmt"
)

// ReportingCtrl - struct for ReportingCtrl
type ReportingCtrl struct {
	ReportingCtrlOneOf *ReportingCtrlOneOf
	ReportingCtrlOneOf1 *ReportingCtrlOneOf1
	ReportingCtrlOneOf2 *ReportingCtrlOneOf2
}

// ReportingCtrlOneOfAsReportingCtrl is a convenience function that returns ReportingCtrlOneOf wrapped in ReportingCtrl
func ReportingCtrlOneOfAsReportingCtrl(v *ReportingCtrlOneOf) ReportingCtrl {
	return ReportingCtrl{
		ReportingCtrlOneOf: v,
	}
}

// ReportingCtrlOneOf1AsReportingCtrl is a convenience function that returns ReportingCtrlOneOf1 wrapped in ReportingCtrl
func ReportingCtrlOneOf1AsReportingCtrl(v *ReportingCtrlOneOf1) ReportingCtrl {
	return ReportingCtrl{
		ReportingCtrlOneOf1: v,
	}
}

// ReportingCtrlOneOf2AsReportingCtrl is a convenience function that returns ReportingCtrlOneOf2 wrapped in ReportingCtrl
func ReportingCtrlOneOf2AsReportingCtrl(v *ReportingCtrlOneOf2) ReportingCtrl {
	return ReportingCtrl{
		ReportingCtrlOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReportingCtrl) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ReportingCtrlOneOf
	err = newStrictDecoder(data).Decode(&dst.ReportingCtrlOneOf)
	if err == nil {
		jsonReportingCtrlOneOf, _ := json.Marshal(dst.ReportingCtrlOneOf)
		if string(jsonReportingCtrlOneOf) == "{}" { // empty struct
			dst.ReportingCtrlOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ReportingCtrlOneOf = nil
	}

	// try to unmarshal data into ReportingCtrlOneOf1
	err = newStrictDecoder(data).Decode(&dst.ReportingCtrlOneOf1)
	if err == nil {
		jsonReportingCtrlOneOf1, _ := json.Marshal(dst.ReportingCtrlOneOf1)
		if string(jsonReportingCtrlOneOf1) == "{}" { // empty struct
			dst.ReportingCtrlOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ReportingCtrlOneOf1 = nil
	}

	// try to unmarshal data into ReportingCtrlOneOf2
	err = newStrictDecoder(data).Decode(&dst.ReportingCtrlOneOf2)
	if err == nil {
		jsonReportingCtrlOneOf2, _ := json.Marshal(dst.ReportingCtrlOneOf2)
		if string(jsonReportingCtrlOneOf2) == "{}" { // empty struct
			dst.ReportingCtrlOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.ReportingCtrlOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ReportingCtrlOneOf = nil
		dst.ReportingCtrlOneOf1 = nil
		dst.ReportingCtrlOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ReportingCtrl)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReportingCtrl)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReportingCtrl) MarshalJSON() ([]byte, error) {
	if src.ReportingCtrlOneOf != nil {
		return json.Marshal(&src.ReportingCtrlOneOf)
	}

	if src.ReportingCtrlOneOf1 != nil {
		return json.Marshal(&src.ReportingCtrlOneOf1)
	}

	if src.ReportingCtrlOneOf2 != nil {
		return json.Marshal(&src.ReportingCtrlOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReportingCtrl) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ReportingCtrlOneOf != nil {
		return obj.ReportingCtrlOneOf
	}

	if obj.ReportingCtrlOneOf1 != nil {
		return obj.ReportingCtrlOneOf1
	}

	if obj.ReportingCtrlOneOf2 != nil {
		return obj.ReportingCtrlOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableReportingCtrl struct {
	value *ReportingCtrl
	isSet bool
}

func (v NullableReportingCtrl) Get() *ReportingCtrl {
	return v.value
}

func (v *NullableReportingCtrl) Set(val *ReportingCtrl) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingCtrl) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingCtrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingCtrl(val *ReportingCtrl) *NullableReportingCtrl {
	return &NullableReportingCtrl{value: val, isSet: true}
}

func (v NullableReportingCtrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingCtrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


