/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// Category the model 'Category'
type Category string

// List of Category
const (
	CHARACTER Category = "CHARACTER"
	SCALABILITY Category = "SCALABILITY"
)

// All allowed values of Category enum
var AllowedCategoryEnumValues = []Category{
	"CHARACTER",
	"SCALABILITY",
}

func (v *Category) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Category(value)
	for _, existing := range AllowedCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Category", value)
}

// NewCategoryFromValue returns a pointer to a valid Category
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCategoryFromValue(v string) (*Category, error) {
	ev := Category(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Category: valid values are %v", v, AllowedCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Category) IsValid() bool {
	for _, existing := range AllowedCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Category value
func (v Category) Ptr() *Category {
	return &v
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

