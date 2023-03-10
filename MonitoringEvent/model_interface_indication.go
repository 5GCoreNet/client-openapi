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

// InterfaceIndication Possible values are - EXPOSURE_FUNCTION: SCEF is used for the PDN connection towards the SCS/AS. - PDN_GATEWAY: PDN gateway is used for the PDN connection towards the SCS/AS. 
type InterfaceIndication struct {
	InterfaceIndicationAnyOf *InterfaceIndicationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InterfaceIndication) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InterfaceIndicationAnyOf
	err = json.Unmarshal(data, &dst.InterfaceIndicationAnyOf);
	if err == nil {
		jsonInterfaceIndicationAnyOf, _ := json.Marshal(dst.InterfaceIndicationAnyOf)
		if string(jsonInterfaceIndicationAnyOf) == "{}" { // empty struct
			dst.InterfaceIndicationAnyOf = nil
		} else {
			return nil // data stored in dst.InterfaceIndicationAnyOf, return on the first match
		}
	} else {
		dst.InterfaceIndicationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(InterfaceIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InterfaceIndication) MarshalJSON() ([]byte, error) {
	if src.InterfaceIndicationAnyOf != nil {
		return json.Marshal(&src.InterfaceIndicationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInterfaceIndication struct {
	value *InterfaceIndication
	isSet bool
}

func (v NullableInterfaceIndication) Get() *InterfaceIndication {
	return v.value
}

func (v *NullableInterfaceIndication) Set(val *InterfaceIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceIndication(val *InterfaceIndication) *NullableInterfaceIndication {
	return &NullableInterfaceIndication{value: val, isSet: true}
}

func (v NullableInterfaceIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


