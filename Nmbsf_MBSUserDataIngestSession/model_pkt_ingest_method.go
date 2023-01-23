/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// PktIngestMethod Packet Ingest Method
type PktIngestMethod struct {
	PktIngestMethodAnyOf *PktIngestMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PktIngestMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PktIngestMethodAnyOf
	err = json.Unmarshal(data, &dst.PktIngestMethodAnyOf);
	if err == nil {
		jsonPktIngestMethodAnyOf, _ := json.Marshal(dst.PktIngestMethodAnyOf)
		if string(jsonPktIngestMethodAnyOf) == "{}" { // empty struct
			dst.PktIngestMethodAnyOf = nil
		} else {
			return nil // data stored in dst.PktIngestMethodAnyOf, return on the first match
		}
	} else {
		dst.PktIngestMethodAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PktIngestMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PktIngestMethod) MarshalJSON() ([]byte, error) {
	if src.PktIngestMethodAnyOf != nil {
		return json.Marshal(&src.PktIngestMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePktIngestMethod struct {
	value *PktIngestMethod
	isSet bool
}

func (v NullablePktIngestMethod) Get() *PktIngestMethod {
	return v.value
}

func (v *NullablePktIngestMethod) Set(val *PktIngestMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePktIngestMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePktIngestMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktIngestMethod(val *PktIngestMethod) *NullablePktIngestMethod {
	return &NullablePktIngestMethod{value: val, isSet: true}
}

func (v NullablePktIngestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktIngestMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


