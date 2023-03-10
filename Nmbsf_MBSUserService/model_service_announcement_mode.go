/*
nmbsf-mbs-us

API for MBS User Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserService

import (
	"encoding/json"
	"fmt"
)

// ServiceAnnouncementMode - Possible values are: - VIA_MBS_5: Indicates the MBS User Service Announcement compiled by the MBSF is advertised to the MBSF Client at reference point MBS-5. - VIA_MBS_DISTRIBUTION_SESSION: Indicates the MBS User Service Announcement compiled by the MBSF is advertised to the MBSF Client via the MBS Distribution Session. - PASSED_BACK: Indicates the MBS User Service Announcement compiled by the MBSF is passed back to the MBS Application Provider. 
type ServiceAnnouncementMode struct {
	ServiceAnnouncementModeOneOf *ServiceAnnouncementModeOneOf
	String *string
}

// ServiceAnnouncementModeOneOfAsServiceAnnouncementMode is a convenience function that returns ServiceAnnouncementModeOneOf wrapped in ServiceAnnouncementMode
func ServiceAnnouncementModeOneOfAsServiceAnnouncementMode(v *ServiceAnnouncementModeOneOf) ServiceAnnouncementMode {
	return ServiceAnnouncementMode{
		ServiceAnnouncementModeOneOf: v,
	}
}

// stringAsServiceAnnouncementMode is a convenience function that returns string wrapped in ServiceAnnouncementMode
func StringAsServiceAnnouncementMode(v *string) ServiceAnnouncementMode {
	return ServiceAnnouncementMode{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServiceAnnouncementMode) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceAnnouncementModeOneOf
	err = newStrictDecoder(data).Decode(&dst.ServiceAnnouncementModeOneOf)
	if err == nil {
		jsonServiceAnnouncementModeOneOf, _ := json.Marshal(dst.ServiceAnnouncementModeOneOf)
		if string(jsonServiceAnnouncementModeOneOf) == "{}" { // empty struct
			dst.ServiceAnnouncementModeOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ServiceAnnouncementModeOneOf = nil
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
		dst.ServiceAnnouncementModeOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServiceAnnouncementMode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServiceAnnouncementMode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServiceAnnouncementMode) MarshalJSON() ([]byte, error) {
	if src.ServiceAnnouncementModeOneOf != nil {
		return json.Marshal(&src.ServiceAnnouncementModeOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServiceAnnouncementMode) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ServiceAnnouncementModeOneOf != nil {
		return obj.ServiceAnnouncementModeOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableServiceAnnouncementMode struct {
	value *ServiceAnnouncementMode
	isSet bool
}

func (v NullableServiceAnnouncementMode) Get() *ServiceAnnouncementMode {
	return v.value
}

func (v *NullableServiceAnnouncementMode) Set(val *ServiceAnnouncementMode) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAnnouncementMode) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAnnouncementMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAnnouncementMode(val *ServiceAnnouncementMode) *NullableServiceAnnouncementMode {
	return &NullableServiceAnnouncementMode{value: val, isSet: true}
}

func (v NullableServiceAnnouncementMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAnnouncementMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


