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

// PdnConnectionStatus Possible values are - CREATED: The PDN connection is created. - RELEASED: The PDN connection is released. 
type PdnConnectionStatus struct {
	PdnConnectionStatusAnyOf *PdnConnectionStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PdnConnectionStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PdnConnectionStatusAnyOf
	err = json.Unmarshal(data, &dst.PdnConnectionStatusAnyOf);
	if err == nil {
		jsonPdnConnectionStatusAnyOf, _ := json.Marshal(dst.PdnConnectionStatusAnyOf)
		if string(jsonPdnConnectionStatusAnyOf) == "{}" { // empty struct
			dst.PdnConnectionStatusAnyOf = nil
		} else {
			return nil // data stored in dst.PdnConnectionStatusAnyOf, return on the first match
		}
	} else {
		dst.PdnConnectionStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PdnConnectionStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PdnConnectionStatus) MarshalJSON() ([]byte, error) {
	if src.PdnConnectionStatusAnyOf != nil {
		return json.Marshal(&src.PdnConnectionStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePdnConnectionStatus struct {
	value *PdnConnectionStatus
	isSet bool
}

func (v NullablePdnConnectionStatus) Get() *PdnConnectionStatus {
	return v.value
}

func (v *NullablePdnConnectionStatus) Set(val *PdnConnectionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnConnectionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnConnectionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnConnectionStatus(val *PdnConnectionStatus) *NullablePdnConnectionStatus {
	return &NullablePdnConnectionStatus{value: val, isSet: true}
}

func (v NullablePdnConnectionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnConnectionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


