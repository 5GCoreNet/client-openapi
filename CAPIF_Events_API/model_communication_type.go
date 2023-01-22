/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Events_API

import (
	"encoding/json"
	"fmt"
)

// CommunicationType Possible values are: - REQUEST_RESPONSE: The communication is of the type request-response - SUBSCRIBE_NOTIFY: The communication is of the type subscribe-notify 
type CommunicationType struct {
	CommunicationTypeAnyOf *CommunicationTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CommunicationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CommunicationTypeAnyOf
	err = json.Unmarshal(data, &dst.CommunicationTypeAnyOf);
	if err == nil {
		jsonCommunicationTypeAnyOf, _ := json.Marshal(dst.CommunicationTypeAnyOf)
		if string(jsonCommunicationTypeAnyOf) == "{}" { // empty struct
			dst.CommunicationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.CommunicationTypeAnyOf, return on the first match
		}
	} else {
		dst.CommunicationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CommunicationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CommunicationType) MarshalJSON() ([]byte, error) {
	if src.CommunicationTypeAnyOf != nil {
		return json.Marshal(&src.CommunicationTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCommunicationType struct {
	value *CommunicationType
	isSet bool
}

func (v NullableCommunicationType) Get() *CommunicationType {
	return v.value
}

func (v *NullableCommunicationType) Set(val *CommunicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationType(val *CommunicationType) *NullableCommunicationType {
	return &NullableCommunicationType{value: val, isSet: true}
}

func (v NullableCommunicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


