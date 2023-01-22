/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_ACREvents

import (
	"encoding/json"
	"fmt"
)

// EASCategory Possible values are: - UAS: Category of EAS is for Uncrewed Aerial Services. - V2X: Category of EAS is for V2X Services. - OTHER: Any other type of EAS category. 
type EASCategory struct {
	EASCategoryAnyOf *EASCategoryAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EASCategory) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EASCategoryAnyOf
	err = json.Unmarshal(data, &dst.EASCategoryAnyOf);
	if err == nil {
		jsonEASCategoryAnyOf, _ := json.Marshal(dst.EASCategoryAnyOf)
		if string(jsonEASCategoryAnyOf) == "{}" { // empty struct
			dst.EASCategoryAnyOf = nil
		} else {
			return nil // data stored in dst.EASCategoryAnyOf, return on the first match
		}
	} else {
		dst.EASCategoryAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EASCategory)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EASCategory) MarshalJSON() ([]byte, error) {
	if src.EASCategoryAnyOf != nil {
		return json.Marshal(&src.EASCategoryAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEASCategory struct {
	value *EASCategory
	isSet bool
}

func (v NullableEASCategory) Get() *EASCategory {
	return v.value
}

func (v *NullableEASCategory) Set(val *EASCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEASCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEASCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASCategory(val *EASCategory) *NullableEASCategory {
	return &NullableEASCategory{value: val, isSet: true}
}

func (v NullableEASCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


