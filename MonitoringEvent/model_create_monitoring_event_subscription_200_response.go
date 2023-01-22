/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// CreateMonitoringEventSubscription200Response - struct for CreateMonitoringEventSubscription200Response
type CreateMonitoringEventSubscription200Response struct {
	MonitoringEventReport *MonitoringEventReport
	MonitoringEventReports *MonitoringEventReports
}

// MonitoringEventReportAsCreateMonitoringEventSubscription200Response is a convenience function that returns MonitoringEventReport wrapped in CreateMonitoringEventSubscription200Response
func MonitoringEventReportAsCreateMonitoringEventSubscription200Response(v *MonitoringEventReport) CreateMonitoringEventSubscription200Response {
	return CreateMonitoringEventSubscription200Response{
		MonitoringEventReport: v,
	}
}

// MonitoringEventReportsAsCreateMonitoringEventSubscription200Response is a convenience function that returns MonitoringEventReports wrapped in CreateMonitoringEventSubscription200Response
func MonitoringEventReportsAsCreateMonitoringEventSubscription200Response(v *MonitoringEventReports) CreateMonitoringEventSubscription200Response {
	return CreateMonitoringEventSubscription200Response{
		MonitoringEventReports: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMonitoringEventSubscription200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitoringEventReport
	err = newStrictDecoder(data).Decode(&dst.MonitoringEventReport)
	if err == nil {
		jsonMonitoringEventReport, _ := json.Marshal(dst.MonitoringEventReport)
		if string(jsonMonitoringEventReport) == "{}" { // empty struct
			dst.MonitoringEventReport = nil
		} else {
			match++
		}
	} else {
		dst.MonitoringEventReport = nil
	}

	// try to unmarshal data into MonitoringEventReports
	err = newStrictDecoder(data).Decode(&dst.MonitoringEventReports)
	if err == nil {
		jsonMonitoringEventReports, _ := json.Marshal(dst.MonitoringEventReports)
		if string(jsonMonitoringEventReports) == "{}" { // empty struct
			dst.MonitoringEventReports = nil
		} else {
			match++
		}
	} else {
		dst.MonitoringEventReports = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MonitoringEventReport = nil
		dst.MonitoringEventReports = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateMonitoringEventSubscription200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateMonitoringEventSubscription200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateMonitoringEventSubscription200Response) MarshalJSON() ([]byte, error) {
	if src.MonitoringEventReport != nil {
		return json.Marshal(&src.MonitoringEventReport)
	}

	if src.MonitoringEventReports != nil {
		return json.Marshal(&src.MonitoringEventReports)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateMonitoringEventSubscription200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MonitoringEventReport != nil {
		return obj.MonitoringEventReport
	}

	if obj.MonitoringEventReports != nil {
		return obj.MonitoringEventReports
	}

	// all schemas are nil
	return nil
}

type NullableCreateMonitoringEventSubscription200Response struct {
	value *CreateMonitoringEventSubscription200Response
	isSet bool
}

func (v NullableCreateMonitoringEventSubscription200Response) Get() *CreateMonitoringEventSubscription200Response {
	return v.value
}

func (v *NullableCreateMonitoringEventSubscription200Response) Set(val *CreateMonitoringEventSubscription200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMonitoringEventSubscription200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMonitoringEventSubscription200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMonitoringEventSubscription200Response(val *CreateMonitoringEventSubscription200Response) *NullableCreateMonitoringEventSubscription200Response {
	return &NullableCreateMonitoringEventSubscription200Response{value: val, isSet: true}
}

func (v NullableCreateMonitoringEventSubscription200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMonitoringEventSubscription200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

