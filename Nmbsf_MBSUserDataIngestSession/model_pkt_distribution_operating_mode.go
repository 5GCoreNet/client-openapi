/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// PktDistributionOperatingMode Mode of data ingestion for Packet distribution method
type PktDistributionOperatingMode struct {
	PktDistributionOperatingModeAnyOf *PktDistributionOperatingModeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PktDistributionOperatingMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PktDistributionOperatingModeAnyOf
	err = json.Unmarshal(data, &dst.PktDistributionOperatingModeAnyOf);
	if err == nil {
		jsonPktDistributionOperatingModeAnyOf, _ := json.Marshal(dst.PktDistributionOperatingModeAnyOf)
		if string(jsonPktDistributionOperatingModeAnyOf) == "{}" { // empty struct
			dst.PktDistributionOperatingModeAnyOf = nil
		} else {
			return nil // data stored in dst.PktDistributionOperatingModeAnyOf, return on the first match
		}
	} else {
		dst.PktDistributionOperatingModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PktDistributionOperatingMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PktDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	if src.PktDistributionOperatingModeAnyOf != nil {
		return json.Marshal(&src.PktDistributionOperatingModeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePktDistributionOperatingMode struct {
	value *PktDistributionOperatingMode
	isSet bool
}

func (v NullablePktDistributionOperatingMode) Get() *PktDistributionOperatingMode {
	return v.value
}

func (v *NullablePktDistributionOperatingMode) Set(val *PktDistributionOperatingMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePktDistributionOperatingMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePktDistributionOperatingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktDistributionOperatingMode(val *PktDistributionOperatingMode) *NullablePktDistributionOperatingMode {
	return &NullablePktDistributionOperatingMode{value: val, isSet: true}
}

func (v NullablePktDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktDistributionOperatingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

