/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
	"fmt"
)

// EventThresholdTypeEventThresholdRSRP - struct for EventThresholdTypeEventThresholdRSRP
type EventThresholdTypeEventThresholdRSRP struct {
	Int32 *int32
}

// int32AsEventThresholdTypeEventThresholdRSRP is a convenience function that returns int32 wrapped in EventThresholdTypeEventThresholdRSRP
func Int32AsEventThresholdTypeEventThresholdRSRP(v *int32) EventThresholdTypeEventThresholdRSRP {
	return EventThresholdTypeEventThresholdRSRP{
		Int32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EventThresholdTypeEventThresholdRSRP) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EventThresholdTypeEventThresholdRSRP)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EventThresholdTypeEventThresholdRSRP)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventThresholdTypeEventThresholdRSRP) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EventThresholdTypeEventThresholdRSRP) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableEventThresholdTypeEventThresholdRSRP struct {
	value *EventThresholdTypeEventThresholdRSRP
	isSet bool
}

func (v NullableEventThresholdTypeEventThresholdRSRP) Get() *EventThresholdTypeEventThresholdRSRP {
	return v.value
}

func (v *NullableEventThresholdTypeEventThresholdRSRP) Set(val *EventThresholdTypeEventThresholdRSRP) {
	v.value = val
	v.isSet = true
}

func (v NullableEventThresholdTypeEventThresholdRSRP) IsSet() bool {
	return v.isSet
}

func (v *NullableEventThresholdTypeEventThresholdRSRP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventThresholdTypeEventThresholdRSRP(val *EventThresholdTypeEventThresholdRSRP) *NullableEventThresholdTypeEventThresholdRSRP {
	return &NullableEventThresholdTypeEventThresholdRSRP{value: val, isSet: true}
}

func (v NullableEventThresholdTypeEventThresholdRSRP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventThresholdTypeEventThresholdRSRP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


