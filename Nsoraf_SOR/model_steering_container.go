/*
Nsoraf_SOR

Nsoraf Steering Of Roaming Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsoraf_SOR

import (
	"encoding/json"
	"fmt"
)

// SteeringContainer - It consists of either a list (array) of SteeringInfo objects or a Secured Packet
type SteeringContainer struct {
	ArrayOfSteeringInfo *[]SteeringInfo
	String *string
}

// []SteeringInfoAsSteeringContainer is a convenience function that returns []SteeringInfo wrapped in SteeringContainer
func ArrayOfSteeringInfoAsSteeringContainer(v *[]SteeringInfo) SteeringContainer {
	return SteeringContainer{
		ArrayOfSteeringInfo: v,
	}
}

// stringAsSteeringContainer is a convenience function that returns string wrapped in SteeringContainer
func StringAsSteeringContainer(v *string) SteeringContainer {
	return SteeringContainer{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SteeringContainer) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfSteeringInfo
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSteeringInfo)
	if err == nil {
		jsonArrayOfSteeringInfo, _ := json.Marshal(dst.ArrayOfSteeringInfo)
		if string(jsonArrayOfSteeringInfo) == "{}" { // empty struct
			dst.ArrayOfSteeringInfo = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfSteeringInfo = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfSteeringInfo = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SteeringContainer)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SteeringContainer)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SteeringContainer) MarshalJSON() ([]byte, error) {
	if src.ArrayOfSteeringInfo != nil {
		return json.Marshal(&src.ArrayOfSteeringInfo)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SteeringContainer) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfSteeringInfo != nil {
		return obj.ArrayOfSteeringInfo
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSteeringContainer struct {
	value *SteeringContainer
	isSet bool
}

func (v NullableSteeringContainer) Get() *SteeringContainer {
	return v.value
}

func (v *NullableSteeringContainer) Set(val *SteeringContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableSteeringContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableSteeringContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteeringContainer(val *SteeringContainer) *NullableSteeringContainer {
	return &NullableSteeringContainer{value: val, isSet: true}
}

func (v NullableSteeringContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteeringContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


