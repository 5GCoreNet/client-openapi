/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
	"fmt"
)

// EventConsumerTypeAnyOf the model 'EventConsumerTypeAnyOf'
type EventConsumerTypeAnyOf string

// List of EventConsumerType_anyOf
const (
	NWDAF EventConsumerTypeAnyOf = "NWDAF"
	EVENT_CONSUMER_AF EventConsumerTypeAnyOf = "EVENT_CONSUMER_AF"
	NEF EventConsumerTypeAnyOf = "NEF"
)

// All allowed values of EventConsumerTypeAnyOf enum
var AllowedEventConsumerTypeAnyOfEnumValues = []EventConsumerTypeAnyOf{
	"NWDAF",
	"EVENT_CONSUMER_AF",
	"NEF",
}

func (v *EventConsumerTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventConsumerTypeAnyOf(value)
	for _, existing := range AllowedEventConsumerTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventConsumerTypeAnyOf", value)
}

// NewEventConsumerTypeAnyOfFromValue returns a pointer to a valid EventConsumerTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventConsumerTypeAnyOfFromValue(v string) (*EventConsumerTypeAnyOf, error) {
	ev := EventConsumerTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventConsumerTypeAnyOf: valid values are %v", v, AllowedEventConsumerTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventConsumerTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedEventConsumerTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventConsumerType_anyOf value
func (v EventConsumerTypeAnyOf) Ptr() *EventConsumerTypeAnyOf {
	return &v
}

type NullableEventConsumerTypeAnyOf struct {
	value *EventConsumerTypeAnyOf
	isSet bool
}

func (v NullableEventConsumerTypeAnyOf) Get() *EventConsumerTypeAnyOf {
	return v.value
}

func (v *NullableEventConsumerTypeAnyOf) Set(val *EventConsumerTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventConsumerTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventConsumerTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventConsumerTypeAnyOf(val *EventConsumerTypeAnyOf) *NullableEventConsumerTypeAnyOf {
	return &NullableEventConsumerTypeAnyOf{value: val, isSet: true}
}

func (v NullableEventConsumerTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventConsumerTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

