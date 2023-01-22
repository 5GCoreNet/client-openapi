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

// ObjAcquisitionMethod Object Acquisition Method
type ObjAcquisitionMethod struct {
	ObjAcquisitionMethodAnyOf *ObjAcquisitionMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ObjAcquisitionMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ObjAcquisitionMethodAnyOf
	err = json.Unmarshal(data, &dst.ObjAcquisitionMethodAnyOf);
	if err == nil {
		jsonObjAcquisitionMethodAnyOf, _ := json.Marshal(dst.ObjAcquisitionMethodAnyOf)
		if string(jsonObjAcquisitionMethodAnyOf) == "{}" { // empty struct
			dst.ObjAcquisitionMethodAnyOf = nil
		} else {
			return nil // data stored in dst.ObjAcquisitionMethodAnyOf, return on the first match
		}
	} else {
		dst.ObjAcquisitionMethodAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ObjAcquisitionMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ObjAcquisitionMethod) MarshalJSON() ([]byte, error) {
	if src.ObjAcquisitionMethodAnyOf != nil {
		return json.Marshal(&src.ObjAcquisitionMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableObjAcquisitionMethod struct {
	value *ObjAcquisitionMethod
	isSet bool
}

func (v NullableObjAcquisitionMethod) Get() *ObjAcquisitionMethod {
	return v.value
}

func (v *NullableObjAcquisitionMethod) Set(val *ObjAcquisitionMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableObjAcquisitionMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableObjAcquisitionMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjAcquisitionMethod(val *ObjAcquisitionMethod) *NullableObjAcquisitionMethod {
	return &NullableObjAcquisitionMethod{value: val, isSet: true}
}

func (v NullableObjAcquisitionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjAcquisitionMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


