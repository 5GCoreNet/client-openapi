/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// EventsSubscPutData Identifies the events the application subscribes to within an Events Subscription sub-resource data. It may contain the notification of the already met events. 
type EventsSubscPutData struct {
	EventsNotification *EventsNotification
	EventsSubscReqData *EventsSubscReqData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventsSubscPutData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventsNotification
	err = json.Unmarshal(data, &dst.EventsNotification);
	if err == nil {
		jsonEventsNotification, _ := json.Marshal(dst.EventsNotification)
		if string(jsonEventsNotification) == "{}" { // empty struct
			dst.EventsNotification = nil
		} else {
			return nil // data stored in dst.EventsNotification, return on the first match
		}
	} else {
		dst.EventsNotification = nil
	}

	// try to unmarshal JSON data into EventsSubscReqData
	err = json.Unmarshal(data, &dst.EventsSubscReqData);
	if err == nil {
		jsonEventsSubscReqData, _ := json.Marshal(dst.EventsSubscReqData)
		if string(jsonEventsSubscReqData) == "{}" { // empty struct
			dst.EventsSubscReqData = nil
		} else {
			return nil // data stored in dst.EventsSubscReqData, return on the first match
		}
	} else {
		dst.EventsSubscReqData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EventsSubscPutData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventsSubscPutData) MarshalJSON() ([]byte, error) {
	if src.EventsNotification != nil {
		return json.Marshal(&src.EventsNotification)
	}

	if src.EventsSubscReqData != nil {
		return json.Marshal(&src.EventsSubscReqData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventsSubscPutData struct {
	value *EventsSubscPutData
	isSet bool
}

func (v NullableEventsSubscPutData) Get() *EventsSubscPutData {
	return v.value
}

func (v *NullableEventsSubscPutData) Set(val *EventsSubscPutData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsSubscPutData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsSubscPutData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsSubscPutData(val *EventsSubscPutData) *NullableEventsSubscPutData {
	return &NullableEventsSubscPutData{value: val, isSet: true}
}

func (v NullableEventsSubscPutData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsSubscPutData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


