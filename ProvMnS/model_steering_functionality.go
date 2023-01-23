/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// SteeringFunctionality Possible values are   - MPTCP: Indicates that PCF authorizes the MPTCP functionality to support traffic steering, switching and splitting.   - ATSSS_LL: Indicates that PCF authorizes the ATSSS-LL functionality to support traffic steering, switching and splitting. 
type SteeringFunctionality struct {
	SteeringFunctionalityAnyOf *SteeringFunctionalityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SteeringFunctionality) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SteeringFunctionalityAnyOf
	err = json.Unmarshal(data, &dst.SteeringFunctionalityAnyOf);
	if err == nil {
		jsonSteeringFunctionalityAnyOf, _ := json.Marshal(dst.SteeringFunctionalityAnyOf)
		if string(jsonSteeringFunctionalityAnyOf) == "{}" { // empty struct
			dst.SteeringFunctionalityAnyOf = nil
		} else {
			return nil // data stored in dst.SteeringFunctionalityAnyOf, return on the first match
		}
	} else {
		dst.SteeringFunctionalityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SteeringFunctionality)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SteeringFunctionality) MarshalJSON() ([]byte, error) {
	if src.SteeringFunctionalityAnyOf != nil {
		return json.Marshal(&src.SteeringFunctionalityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSteeringFunctionality struct {
	value *SteeringFunctionality
	isSet bool
}

func (v NullableSteeringFunctionality) Get() *SteeringFunctionality {
	return v.value
}

func (v *NullableSteeringFunctionality) Set(val *SteeringFunctionality) {
	v.value = val
	v.isSet = true
}

func (v NullableSteeringFunctionality) IsSet() bool {
	return v.isSet
}

func (v *NullableSteeringFunctionality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteeringFunctionality(val *SteeringFunctionality) *NullableSteeringFunctionality {
	return &NullableSteeringFunctionality{value: val, isSet: true}
}

func (v NullableSteeringFunctionality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteeringFunctionality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


