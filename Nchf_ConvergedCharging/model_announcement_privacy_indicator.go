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

// AnnouncementPrivacyIndicator struct for AnnouncementPrivacyIndicator
type AnnouncementPrivacyIndicator struct {
	AnnouncementPrivacyIndicatorAnyOf *AnnouncementPrivacyIndicatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnnouncementPrivacyIndicator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnnouncementPrivacyIndicatorAnyOf
	err = json.Unmarshal(data, &dst.AnnouncementPrivacyIndicatorAnyOf);
	if err == nil {
		jsonAnnouncementPrivacyIndicatorAnyOf, _ := json.Marshal(dst.AnnouncementPrivacyIndicatorAnyOf)
		if string(jsonAnnouncementPrivacyIndicatorAnyOf) == "{}" { // empty struct
			dst.AnnouncementPrivacyIndicatorAnyOf = nil
		} else {
			return nil // data stored in dst.AnnouncementPrivacyIndicatorAnyOf, return on the first match
		}
	} else {
		dst.AnnouncementPrivacyIndicatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AnnouncementPrivacyIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnnouncementPrivacyIndicator) MarshalJSON() ([]byte, error) {
	if src.AnnouncementPrivacyIndicatorAnyOf != nil {
		return json.Marshal(&src.AnnouncementPrivacyIndicatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnnouncementPrivacyIndicator struct {
	value *AnnouncementPrivacyIndicator
	isSet bool
}

func (v NullableAnnouncementPrivacyIndicator) Get() *AnnouncementPrivacyIndicator {
	return v.value
}

func (v *NullableAnnouncementPrivacyIndicator) Set(val *AnnouncementPrivacyIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnouncementPrivacyIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnouncementPrivacyIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnouncementPrivacyIndicator(val *AnnouncementPrivacyIndicator) *NullableAnnouncementPrivacyIndicator {
	return &NullableAnnouncementPrivacyIndicator{value: val, isSet: true}
}

func (v NullableAnnouncementPrivacyIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnouncementPrivacyIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


