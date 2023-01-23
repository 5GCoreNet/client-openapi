/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"fmt"
)

// ServiceAnnoucementMode Possible values are: - NRM: NRM server performs the service announcement. - VAL: VAL server performs the service announcement. 
type ServiceAnnoucementMode struct {
	ServiceAnnoucementModeAnyOf *ServiceAnnoucementModeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceAnnoucementMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ServiceAnnoucementModeAnyOf
	err = json.Unmarshal(data, &dst.ServiceAnnoucementModeAnyOf);
	if err == nil {
		jsonServiceAnnoucementModeAnyOf, _ := json.Marshal(dst.ServiceAnnoucementModeAnyOf)
		if string(jsonServiceAnnoucementModeAnyOf) == "{}" { // empty struct
			dst.ServiceAnnoucementModeAnyOf = nil
		} else {
			return nil // data stored in dst.ServiceAnnoucementModeAnyOf, return on the first match
		}
	} else {
		dst.ServiceAnnoucementModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceAnnoucementMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceAnnoucementMode) MarshalJSON() ([]byte, error) {
	if src.ServiceAnnoucementModeAnyOf != nil {
		return json.Marshal(&src.ServiceAnnoucementModeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceAnnoucementMode struct {
	value *ServiceAnnoucementMode
	isSet bool
}

func (v NullableServiceAnnoucementMode) Get() *ServiceAnnoucementMode {
	return v.value
}

func (v *NullableServiceAnnoucementMode) Set(val *ServiceAnnoucementMode) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAnnoucementMode) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAnnoucementMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAnnoucementMode(val *ServiceAnnoucementMode) *NullableServiceAnnoucementMode {
	return &NullableServiceAnnoucementMode{value: val, isSet: true}
}

func (v NullableServiceAnnoucementMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAnnoucementMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


