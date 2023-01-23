/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// ScheduledCommunicationType Possible values are: -DOWNLINK_ONLY: Downlink only -UPLINK_ONLY: Uplink only -BIDIRECTIONA: Bi-directional 
type ScheduledCommunicationType struct {
	ScheduledCommunicationTypeAnyOf *ScheduledCommunicationTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScheduledCommunicationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ScheduledCommunicationTypeAnyOf
	err = json.Unmarshal(data, &dst.ScheduledCommunicationTypeAnyOf);
	if err == nil {
		jsonScheduledCommunicationTypeAnyOf, _ := json.Marshal(dst.ScheduledCommunicationTypeAnyOf)
		if string(jsonScheduledCommunicationTypeAnyOf) == "{}" { // empty struct
			dst.ScheduledCommunicationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ScheduledCommunicationTypeAnyOf, return on the first match
		}
	} else {
		dst.ScheduledCommunicationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ScheduledCommunicationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScheduledCommunicationType) MarshalJSON() ([]byte, error) {
	if src.ScheduledCommunicationTypeAnyOf != nil {
		return json.Marshal(&src.ScheduledCommunicationTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScheduledCommunicationType struct {
	value *ScheduledCommunicationType
	isSet bool
}

func (v NullableScheduledCommunicationType) Get() *ScheduledCommunicationType {
	return v.value
}

func (v *NullableScheduledCommunicationType) Set(val *ScheduledCommunicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledCommunicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledCommunicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledCommunicationType(val *ScheduledCommunicationType) *NullableScheduledCommunicationType {
	return &NullableScheduledCommunicationType{value: val, isSet: true}
}

func (v NullableScheduledCommunicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledCommunicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


