/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NIDD

import (
	"encoding/json"
	"fmt"
)

// ManageEntity Possible values are - UE: Representing the UE. - AS: Representing the Application Server. 
type ManageEntity struct {
	ManageEntityAnyOf *ManageEntityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ManageEntity) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ManageEntityAnyOf
	err = json.Unmarshal(data, &dst.ManageEntityAnyOf);
	if err == nil {
		jsonManageEntityAnyOf, _ := json.Marshal(dst.ManageEntityAnyOf)
		if string(jsonManageEntityAnyOf) == "{}" { // empty struct
			dst.ManageEntityAnyOf = nil
		} else {
			return nil // data stored in dst.ManageEntityAnyOf, return on the first match
		}
	} else {
		dst.ManageEntityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ManageEntity)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ManageEntity) MarshalJSON() ([]byte, error) {
	if src.ManageEntityAnyOf != nil {
		return json.Marshal(&src.ManageEntityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableManageEntity struct {
	value *ManageEntity
	isSet bool
}

func (v NullableManageEntity) Get() *ManageEntity {
	return v.value
}

func (v *NullableManageEntity) Set(val *ManageEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableManageEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableManageEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageEntity(val *ManageEntity) *NullableManageEntity {
	return &NullableManageEntity{value: val, isSet: true}
}

func (v NullableManageEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


