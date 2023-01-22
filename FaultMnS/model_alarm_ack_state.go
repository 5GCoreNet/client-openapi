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

// AlarmAckState the model 'AlarmAckState'
type AlarmAckState string

// List of AlarmAckState
const (
	ALARMS AlarmAckState = "ALL_ALARMS"
	ACTIVE_ALARMS AlarmAckState = "ALL_ACTIVE_ALARMS"
	ACTIVE_AND_ACKNOWLEDGED_ALARMS AlarmAckState = "ALL_ACTIVE_AND_ACKNOWLEDGED_ALARMS"
	ACTIVE_AND_UNACKNOWLEDGED_ALARMS AlarmAckState = "ALL_ACTIVE_AND_UNACKNOWLEDGED_ALARMS"
	CLEARED_AND_UNACKNOWLEDGED_ALARMS AlarmAckState = "ALL_CLEARED_AND_UNACKNOWLEDGED_ALARMS"
	UNACKNOWLEDGED_ALARMS AlarmAckState = "ALL_UNACKNOWLEDGED_ALARMS"
)

// All allowed values of AlarmAckState enum
var AllowedAlarmAckStateEnumValues = []AlarmAckState{
	"ALL_ALARMS",
	"ALL_ACTIVE_ALARMS",
	"ALL_ACTIVE_AND_ACKNOWLEDGED_ALARMS",
	"ALL_ACTIVE_AND_UNACKNOWLEDGED_ALARMS",
	"ALL_CLEARED_AND_UNACKNOWLEDGED_ALARMS",
	"ALL_UNACKNOWLEDGED_ALARMS",
}

func (v *AlarmAckState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlarmAckState(value)
	for _, existing := range AllowedAlarmAckStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlarmAckState", value)
}

// NewAlarmAckStateFromValue returns a pointer to a valid AlarmAckState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlarmAckStateFromValue(v string) (*AlarmAckState, error) {
	ev := AlarmAckState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlarmAckState: valid values are %v", v, AllowedAlarmAckStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlarmAckState) IsValid() bool {
	for _, existing := range AllowedAlarmAckStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlarmAckState value
func (v AlarmAckState) Ptr() *AlarmAckState {
	return &v
}

type NullableAlarmAckState struct {
	value *AlarmAckState
	isSet bool
}

func (v NullableAlarmAckState) Get() *AlarmAckState {
	return v.value
}

func (v *NullableAlarmAckState) Set(val *AlarmAckState) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmAckState) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmAckState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmAckState(val *AlarmAckState) *NullableAlarmAckState {
	return &NullableAlarmAckState{value: val, isSet: true}
}

func (v NullableAlarmAckState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmAckState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

