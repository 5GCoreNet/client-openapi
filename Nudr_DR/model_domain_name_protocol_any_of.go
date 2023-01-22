/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// DomainNameProtocolAnyOf the model 'DomainNameProtocolAnyOf'
type DomainNameProtocolAnyOf string

// List of DomainNameProtocol_anyOf
const (
	DNS_QNAME DomainNameProtocolAnyOf = "DNS_QNAME"
	TLS_SNI DomainNameProtocolAnyOf = "TLS_SNI"
	TLS_SAN DomainNameProtocolAnyOf = "TLS_SAN"
	TSL_SCN DomainNameProtocolAnyOf = "TSL_SCN"
)

// All allowed values of DomainNameProtocolAnyOf enum
var AllowedDomainNameProtocolAnyOfEnumValues = []DomainNameProtocolAnyOf{
	"DNS_QNAME",
	"TLS_SNI",
	"TLS_SAN",
	"TSL_SCN",
}

func (v *DomainNameProtocolAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DomainNameProtocolAnyOf(value)
	for _, existing := range AllowedDomainNameProtocolAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DomainNameProtocolAnyOf", value)
}

// NewDomainNameProtocolAnyOfFromValue returns a pointer to a valid DomainNameProtocolAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDomainNameProtocolAnyOfFromValue(v string) (*DomainNameProtocolAnyOf, error) {
	ev := DomainNameProtocolAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DomainNameProtocolAnyOf: valid values are %v", v, AllowedDomainNameProtocolAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DomainNameProtocolAnyOf) IsValid() bool {
	for _, existing := range AllowedDomainNameProtocolAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DomainNameProtocol_anyOf value
func (v DomainNameProtocolAnyOf) Ptr() *DomainNameProtocolAnyOf {
	return &v
}

type NullableDomainNameProtocolAnyOf struct {
	value *DomainNameProtocolAnyOf
	isSet bool
}

func (v NullableDomainNameProtocolAnyOf) Get() *DomainNameProtocolAnyOf {
	return v.value
}

func (v *NullableDomainNameProtocolAnyOf) Set(val *DomainNameProtocolAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainNameProtocolAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainNameProtocolAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainNameProtocolAnyOf(val *DomainNameProtocolAnyOf) *NullableDomainNameProtocolAnyOf {
	return &NullableDomainNameProtocolAnyOf{value: val, isSet: true}
}

func (v NullableDomainNameProtocolAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainNameProtocolAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

