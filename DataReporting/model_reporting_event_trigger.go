/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
	"fmt"
)

// ReportingEventTrigger The type of event that triggers reporting by a data collection client to the Data Collection AF.
type ReportingEventTrigger struct {
	ReportingEventTriggerAnyOf *ReportingEventTriggerAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportingEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.ReportingEventTriggerAnyOf);
	if err == nil {
		jsonReportingEventTriggerAnyOf, _ := json.Marshal(dst.ReportingEventTriggerAnyOf)
		if string(jsonReportingEventTriggerAnyOf) == "{}" { // empty struct
			dst.ReportingEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.ReportingEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.ReportingEventTriggerAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingEventTrigger) MarshalJSON() ([]byte, error) {
	if src.ReportingEventTriggerAnyOf != nil {
		return json.Marshal(&src.ReportingEventTriggerAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingEventTrigger struct {
	value *ReportingEventTrigger
	isSet bool
}

func (v NullableReportingEventTrigger) Get() *ReportingEventTrigger {
	return v.value
}

func (v *NullableReportingEventTrigger) Set(val *ReportingEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingEventTrigger(val *ReportingEventTrigger) *NullableReportingEventTrigger {
	return &NullableReportingEventTrigger{value: val, isSet: true}
}

func (v NullableReportingEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


