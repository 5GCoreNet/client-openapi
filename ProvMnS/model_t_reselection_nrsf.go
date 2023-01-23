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

// TReselectionNRSf the model 'TReselectionNRSf'
type TReselectionNRSf int32

// List of TReselectionNRSf
const (
	_25 TReselectionNRSf = 25
	_50 TReselectionNRSf = 50
	_75 TReselectionNRSf = 75
	_100 TReselectionNRSf = 100
)

// All allowed values of TReselectionNRSf enum
var AllowedTReselectionNRSfEnumValues = []TReselectionNRSf{
	25,
	50,
	75,
	100,
}

func (v *TReselectionNRSf) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TReselectionNRSf(value)
	for _, existing := range AllowedTReselectionNRSfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TReselectionNRSf", value)
}

// NewTReselectionNRSfFromValue returns a pointer to a valid TReselectionNRSf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTReselectionNRSfFromValue(v int32) (*TReselectionNRSf, error) {
	ev := TReselectionNRSf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TReselectionNRSf: valid values are %v", v, AllowedTReselectionNRSfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TReselectionNRSf) IsValid() bool {
	for _, existing := range AllowedTReselectionNRSfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TReselectionNRSf value
func (v TReselectionNRSf) Ptr() *TReselectionNRSf {
	return &v
}

type NullableTReselectionNRSf struct {
	value *TReselectionNRSf
	isSet bool
}

func (v NullableTReselectionNRSf) Get() *TReselectionNRSf {
	return v.value
}

func (v *NullableTReselectionNRSf) Set(val *TReselectionNRSf) {
	v.value = val
	v.isSet = true
}

func (v NullableTReselectionNRSf) IsSet() bool {
	return v.isSet
}

func (v *NullableTReselectionNRSf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTReselectionNRSf(val *TReselectionNRSf) *NullableTReselectionNRSf {
	return &NullableTReselectionNRSf{value: val, isSet: true}
}

func (v NullableTReselectionNRSf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTReselectionNRSf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

