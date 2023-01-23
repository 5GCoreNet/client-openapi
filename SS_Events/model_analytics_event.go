/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
	"fmt"
)

// AnalyticsEvent Possible values are: - UE_MOBILITY: The AF requests to be notified about analytics information of UE mobility. - UE_COMM: The AF requests to be notified about analytics information of UE communication. - ABNORMAL_BEHAVIOR: The AF requests to be notified about analytics information of UE's abnormal behavior. - CONGESTION: The AF requests to be notified about analytics information of user data congestion information.  - NETWORK_PERFORMANCE: The AF requests to be notified about analytics information of network performance.  - QOS_SUSTAINABILITY: The AF requests to be notified about analytics information of QoS sustainability. - DISPERSION: The AF requests to be notified about analytics information of Dispersion analytics. - DN_PERFORMANCE: The AF requests to be notified about analytics information of DN performance. - SERVICE_EXPERIENCE: The AF requests to be notified about analytics information of service experience. 
type AnalyticsEvent struct {
	AnalyticsEventAnyOf *AnalyticsEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnalyticsEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnalyticsEventAnyOf
	err = json.Unmarshal(data, &dst.AnalyticsEventAnyOf);
	if err == nil {
		jsonAnalyticsEventAnyOf, _ := json.Marshal(dst.AnalyticsEventAnyOf)
		if string(jsonAnalyticsEventAnyOf) == "{}" { // empty struct
			dst.AnalyticsEventAnyOf = nil
		} else {
			return nil // data stored in dst.AnalyticsEventAnyOf, return on the first match
		}
	} else {
		dst.AnalyticsEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AnalyticsEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnalyticsEvent) MarshalJSON() ([]byte, error) {
	if src.AnalyticsEventAnyOf != nil {
		return json.Marshal(&src.AnalyticsEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnalyticsEvent struct {
	value *AnalyticsEvent
	isSet bool
}

func (v NullableAnalyticsEvent) Get() *AnalyticsEvent {
	return v.value
}

func (v *NullableAnalyticsEvent) Set(val *AnalyticsEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsEvent(val *AnalyticsEvent) *NullableAnalyticsEvent {
	return &NullableAnalyticsEvent{value: val, isSet: true}
}

func (v NullableAnalyticsEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


