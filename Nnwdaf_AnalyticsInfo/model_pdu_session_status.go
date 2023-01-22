/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// PduSessionStatus Possible values are: - ACTIVATED: PDU Session status is activated. - DEACTIVATED: PDU Session status is deactivated. 
type PduSessionStatus struct {
	PduSessionStatusAnyOf *PduSessionStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PduSessionStatusAnyOf
	err = json.Unmarshal(data, &dst.PduSessionStatusAnyOf);
	if err == nil {
		jsonPduSessionStatusAnyOf, _ := json.Marshal(dst.PduSessionStatusAnyOf)
		if string(jsonPduSessionStatusAnyOf) == "{}" { // empty struct
			dst.PduSessionStatusAnyOf = nil
		} else {
			return nil // data stored in dst.PduSessionStatusAnyOf, return on the first match
		}
	} else {
		dst.PduSessionStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionStatus) MarshalJSON() ([]byte, error) {
	if src.PduSessionStatusAnyOf != nil {
		return json.Marshal(&src.PduSessionStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionStatus struct {
	value *PduSessionStatus
	isSet bool
}

func (v NullablePduSessionStatus) Get() *PduSessionStatus {
	return v.value
}

func (v *NullablePduSessionStatus) Set(val *PduSessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionStatus(val *PduSessionStatus) *NullablePduSessionStatus {
	return &NullablePduSessionStatus{value: val, isSet: true}
}

func (v NullablePduSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


