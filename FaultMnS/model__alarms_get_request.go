/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
	"fmt"
)

// AlarmsGetRequest - struct for AlarmsGetRequest
type AlarmsGetRequest struct {
	MapmapOfStringMergePatchAcknowledgeAlarm *map[string]MergePatchAcknowledgeAlarm
	MapmapOfStringMergePatchClearAlarm *map[string]MergePatchClearAlarm
}

// map[string]MergePatchAcknowledgeAlarmAsAlarmsGetRequest is a convenience function that returns map[string]MergePatchAcknowledgeAlarm wrapped in AlarmsGetRequest
func MapmapOfStringMergePatchAcknowledgeAlarmAsAlarmsGetRequest(v *map[string]MergePatchAcknowledgeAlarm) AlarmsGetRequest {
	return AlarmsGetRequest{
		MapmapOfStringMergePatchAcknowledgeAlarm: v,
	}
}

// map[string]MergePatchClearAlarmAsAlarmsGetRequest is a convenience function that returns map[string]MergePatchClearAlarm wrapped in AlarmsGetRequest
func MapmapOfStringMergePatchClearAlarmAsAlarmsGetRequest(v *map[string]MergePatchClearAlarm) AlarmsGetRequest {
	return AlarmsGetRequest{
		MapmapOfStringMergePatchClearAlarm: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AlarmsGetRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MapmapOfStringMergePatchAcknowledgeAlarm
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringMergePatchAcknowledgeAlarm)
	if err == nil {
		jsonMapmapOfStringMergePatchAcknowledgeAlarm, _ := json.Marshal(dst.MapmapOfStringMergePatchAcknowledgeAlarm)
		if string(jsonMapmapOfStringMergePatchAcknowledgeAlarm) == "{}" { // empty struct
			dst.MapmapOfStringMergePatchAcknowledgeAlarm = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringMergePatchAcknowledgeAlarm = nil
	}

	// try to unmarshal data into MapmapOfStringMergePatchClearAlarm
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringMergePatchClearAlarm)
	if err == nil {
		jsonMapmapOfStringMergePatchClearAlarm, _ := json.Marshal(dst.MapmapOfStringMergePatchClearAlarm)
		if string(jsonMapmapOfStringMergePatchClearAlarm) == "{}" { // empty struct
			dst.MapmapOfStringMergePatchClearAlarm = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringMergePatchClearAlarm = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringMergePatchAcknowledgeAlarm = nil
		dst.MapmapOfStringMergePatchClearAlarm = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AlarmsGetRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AlarmsGetRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AlarmsGetRequest) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringMergePatchAcknowledgeAlarm != nil {
		return json.Marshal(&src.MapmapOfStringMergePatchAcknowledgeAlarm)
	}

	if src.MapmapOfStringMergePatchClearAlarm != nil {
		return json.Marshal(&src.MapmapOfStringMergePatchClearAlarm)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AlarmsGetRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringMergePatchAcknowledgeAlarm != nil {
		return obj.MapmapOfStringMergePatchAcknowledgeAlarm
	}

	if obj.MapmapOfStringMergePatchClearAlarm != nil {
		return obj.MapmapOfStringMergePatchClearAlarm
	}

	// all schemas are nil
	return nil
}

type NullableAlarmsGetRequest struct {
	value *AlarmsGetRequest
	isSet bool
}

func (v NullableAlarmsGetRequest) Get() *AlarmsGetRequest {
	return v.value
}

func (v *NullableAlarmsGetRequest) Set(val *AlarmsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmsGetRequest(val *AlarmsGetRequest) *NullableAlarmsGetRequest {
	return &NullableAlarmsGetRequest{value: val, isSet: true}
}

func (v NullableAlarmsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


