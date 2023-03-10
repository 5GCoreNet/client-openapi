/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// MonitoringType Possible values are - LOSS_OF_CONNECTIVITY: The SCS/AS requests to be notified when the 3GPP network detects that the UE is no longer reachable for signalling or user plane communication - UE_REACHABILITY: The SCS/AS requests to be notified when the UE becomes reachable for sending either SMS or downlink data to the UE - LOCATION_REPORTING: The SCS/AS requests to be notified of the current location or the last known location of the UE - CHANGE_OF_IMSI_IMEI_ASSOCIATION: The SCS/AS requests to be notified when the association of an ME (IMEI(SV)) that uses a specific subscription (IMSI) is changed - ROAMING_STATUS: The SCS/AS queries the UE's current roaming status and requests to get notified when the status changes - COMMUNICATION_FAILURE: The SCS/AS requests to be notified of communication failure events - AVAILABILITY_AFTER_DDN_FAILURE: The SCS/AS requests to be notified when the UE has become available after a DDN failure - NUMBER_OF_UES_IN_AN_AREA: The SCS/AS requests to be notified the number of UEs in a given geographic area - PDN_CONNECTIVITY_STATUS: The SCS/AS requests to be notified when the 3GPP network detects that the UE’s PDN connection is set up or torn down - DOWNLINK_DATA_DELIVERY_STATUS: The AF requests to be notified when the 3GPP network detects that the downlink data delivery status is changed. - API_SUPPORT_CAPABILITY: The SCS/AS requests to be notified of the availability of support of service APIs. - NUM_OF_REGD_UES: The AF requests to be notified of the current number of registered UEs for a network slice. - NUM_OF_ESTD_PDU_SESSIONS: The AF requests to be notified of the current number of established PDU Sessions for a network slice. - AREA_OF_INTEREST: The SCS/AS requests to be notified when the UAV moves in or out of the geographic area. 
type MonitoringType struct {
	MonitoringTypeAnyOf *MonitoringTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MonitoringType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MonitoringTypeAnyOf
	err = json.Unmarshal(data, &dst.MonitoringTypeAnyOf);
	if err == nil {
		jsonMonitoringTypeAnyOf, _ := json.Marshal(dst.MonitoringTypeAnyOf)
		if string(jsonMonitoringTypeAnyOf) == "{}" { // empty struct
			dst.MonitoringTypeAnyOf = nil
		} else {
			return nil // data stored in dst.MonitoringTypeAnyOf, return on the first match
		}
	} else {
		dst.MonitoringTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MonitoringType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MonitoringType) MarshalJSON() ([]byte, error) {
	if src.MonitoringTypeAnyOf != nil {
		return json.Marshal(&src.MonitoringTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMonitoringType struct {
	value *MonitoringType
	isSet bool
}

func (v NullableMonitoringType) Get() *MonitoringType {
	return v.value
}

func (v *NullableMonitoringType) Set(val *MonitoringType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringType(val *MonitoringType) *NullableMonitoringType {
	return &NullableMonitoringType{value: val, isSet: true}
}

func (v NullableMonitoringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


