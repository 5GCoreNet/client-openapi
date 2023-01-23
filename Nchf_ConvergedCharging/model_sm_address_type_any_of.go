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

// SMAddressTypeAnyOf the model 'SMAddressTypeAnyOf'
type SMAddressTypeAnyOf string

// List of SMAddressType_anyOf
const (
	EMAIL_ADDRESS SMAddressTypeAnyOf = "EMAIL_ADDRESS"
	MSISDN SMAddressTypeAnyOf = "MSISDN"
	IPV4_ADDRESS SMAddressTypeAnyOf = "IPV4_ADDRESS"
	IPV6_ADDRESS SMAddressTypeAnyOf = "IPV6_ADDRESS"
	NUMERIC_SHORTCODE SMAddressTypeAnyOf = "NUMERIC_SHORTCODE"
	ALPHANUMERIC_SHORTCODE SMAddressTypeAnyOf = "ALPHANUMERIC_SHORTCODE"
	OTHER SMAddressTypeAnyOf = "OTHER"
	IMSI SMAddressTypeAnyOf = "IMSI"
)

// All allowed values of SMAddressTypeAnyOf enum
var AllowedSMAddressTypeAnyOfEnumValues = []SMAddressTypeAnyOf{
	"EMAIL_ADDRESS",
	"MSISDN",
	"IPV4_ADDRESS",
	"IPV6_ADDRESS",
	"NUMERIC_SHORTCODE",
	"ALPHANUMERIC_SHORTCODE",
	"OTHER",
	"IMSI",
}

func (v *SMAddressTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SMAddressTypeAnyOf(value)
	for _, existing := range AllowedSMAddressTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SMAddressTypeAnyOf", value)
}

// NewSMAddressTypeAnyOfFromValue returns a pointer to a valid SMAddressTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSMAddressTypeAnyOfFromValue(v string) (*SMAddressTypeAnyOf, error) {
	ev := SMAddressTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SMAddressTypeAnyOf: valid values are %v", v, AllowedSMAddressTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SMAddressTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedSMAddressTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SMAddressType_anyOf value
func (v SMAddressTypeAnyOf) Ptr() *SMAddressTypeAnyOf {
	return &v
}

type NullableSMAddressTypeAnyOf struct {
	value *SMAddressTypeAnyOf
	isSet bool
}

func (v NullableSMAddressTypeAnyOf) Get() *SMAddressTypeAnyOf {
	return v.value
}

func (v *NullableSMAddressTypeAnyOf) Set(val *SMAddressTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSMAddressTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSMAddressTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMAddressTypeAnyOf(val *SMAddressTypeAnyOf) *NullableSMAddressTypeAnyOf {
	return &NullableSMAddressTypeAnyOf{value: val, isSet: true}
}

func (v NullableSMAddressTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMAddressTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

