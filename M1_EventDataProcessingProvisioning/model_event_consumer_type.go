/*
M1_EventDataProcessingProvisioning

5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_EventDataProcessingProvisioning

import (
	"encoding/json"
	"fmt"
)

// EventConsumerType The type of event consumer.
type EventConsumerType struct {
	EventConsumerTypeAnyOf *EventConsumerTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventConsumerType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventConsumerTypeAnyOf
	err = json.Unmarshal(data, &dst.EventConsumerTypeAnyOf);
	if err == nil {
		jsonEventConsumerTypeAnyOf, _ := json.Marshal(dst.EventConsumerTypeAnyOf)
		if string(jsonEventConsumerTypeAnyOf) == "{}" { // empty struct
			dst.EventConsumerTypeAnyOf = nil
		} else {
			return nil // data stored in dst.EventConsumerTypeAnyOf, return on the first match
		}
	} else {
		dst.EventConsumerTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EventConsumerType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventConsumerType) MarshalJSON() ([]byte, error) {
	if src.EventConsumerTypeAnyOf != nil {
		return json.Marshal(&src.EventConsumerTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventConsumerType struct {
	value *EventConsumerType
	isSet bool
}

func (v NullableEventConsumerType) Get() *EventConsumerType {
	return v.value
}

func (v *NullableEventConsumerType) Set(val *EventConsumerType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventConsumerType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventConsumerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventConsumerType(val *EventConsumerType) *NullableEventConsumerType {
	return &NullableEventConsumerType{value: val, isSet: true}
}

func (v NullableEventConsumerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventConsumerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


