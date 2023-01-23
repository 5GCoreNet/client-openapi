/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AppAmContextExpRespData It represents a response to a modification or creation request of an Individual Application AM resource. It may contain the notification of the already met events 
type AppAmContextExpRespData struct {
	AmEventsNotification *AmEventsNotification
	AppAmContextData *AppAmContextData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppAmContextExpRespData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AmEventsNotification
	err = json.Unmarshal(data, &dst.AmEventsNotification);
	if err == nil {
		jsonAmEventsNotification, _ := json.Marshal(dst.AmEventsNotification)
		if string(jsonAmEventsNotification) == "{}" { // empty struct
			dst.AmEventsNotification = nil
		} else {
			return nil // data stored in dst.AmEventsNotification, return on the first match
		}
	} else {
		dst.AmEventsNotification = nil
	}

	// try to unmarshal JSON data into AppAmContextData
	err = json.Unmarshal(data, &dst.AppAmContextData);
	if err == nil {
		jsonAppAmContextData, _ := json.Marshal(dst.AppAmContextData)
		if string(jsonAppAmContextData) == "{}" { // empty struct
			dst.AppAmContextData = nil
		} else {
			return nil // data stored in dst.AppAmContextData, return on the first match
		}
	} else {
		dst.AppAmContextData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AppAmContextExpRespData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppAmContextExpRespData) MarshalJSON() ([]byte, error) {
	if src.AmEventsNotification != nil {
		return json.Marshal(&src.AmEventsNotification)
	}

	if src.AppAmContextData != nil {
		return json.Marshal(&src.AppAmContextData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppAmContextExpRespData struct {
	value *AppAmContextExpRespData
	isSet bool
}

func (v NullableAppAmContextExpRespData) Get() *AppAmContextExpRespData {
	return v.value
}

func (v *NullableAppAmContextExpRespData) Set(val *AppAmContextExpRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAmContextExpRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAmContextExpRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAmContextExpRespData(val *AppAmContextExpRespData) *NullableAppAmContextExpRespData {
	return &NullableAppAmContextExpRespData{value: val, isSet: true}
}

func (v NullableAppAmContextExpRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAmContextExpRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


