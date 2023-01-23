/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// ServiceExperienceType Possible values are: - VOICE: Indicates that the service experience analytics is for voice service. - VIDEO: Indicates that the service experience analytics is for video service. - OTHER: Indicates that the service experience analytics is for other service. 
type ServiceExperienceType struct {
	ServiceExperienceTypeAnyOf *ServiceExperienceTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceExperienceType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ServiceExperienceTypeAnyOf
	err = json.Unmarshal(data, &dst.ServiceExperienceTypeAnyOf);
	if err == nil {
		jsonServiceExperienceTypeAnyOf, _ := json.Marshal(dst.ServiceExperienceTypeAnyOf)
		if string(jsonServiceExperienceTypeAnyOf) == "{}" { // empty struct
			dst.ServiceExperienceTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ServiceExperienceTypeAnyOf, return on the first match
		}
	} else {
		dst.ServiceExperienceTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceExperienceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceExperienceType) MarshalJSON() ([]byte, error) {
	if src.ServiceExperienceTypeAnyOf != nil {
		return json.Marshal(&src.ServiceExperienceTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceExperienceType struct {
	value *ServiceExperienceType
	isSet bool
}

func (v NullableServiceExperienceType) Get() *ServiceExperienceType {
	return v.value
}

func (v *NullableServiceExperienceType) Set(val *ServiceExperienceType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceType(val *ServiceExperienceType) *NullableServiceExperienceType {
	return &NullableServiceExperienceType{value: val, isSet: true}
}

func (v NullableServiceExperienceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


