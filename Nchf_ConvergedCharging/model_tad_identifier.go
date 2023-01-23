/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// TADIdentifier struct for TADIdentifier
type TADIdentifier struct {
	TADIdentifierAnyOf *TADIdentifierAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TADIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TADIdentifierAnyOf
	err = json.Unmarshal(data, &dst.TADIdentifierAnyOf);
	if err == nil {
		jsonTADIdentifierAnyOf, _ := json.Marshal(dst.TADIdentifierAnyOf)
		if string(jsonTADIdentifierAnyOf) == "{}" { // empty struct
			dst.TADIdentifierAnyOf = nil
		} else {
			return nil // data stored in dst.TADIdentifierAnyOf, return on the first match
		}
	} else {
		dst.TADIdentifierAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TADIdentifier)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TADIdentifier) MarshalJSON() ([]byte, error) {
	if src.TADIdentifierAnyOf != nil {
		return json.Marshal(&src.TADIdentifierAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTADIdentifier struct {
	value *TADIdentifier
	isSet bool
}

func (v NullableTADIdentifier) Get() *TADIdentifier {
	return v.value
}

func (v *NullableTADIdentifier) Set(val *TADIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTADIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTADIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTADIdentifier(val *TADIdentifier) *NullableTADIdentifier {
	return &NullableTADIdentifier{value: val, isSet: true}
}

func (v NullableTADIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTADIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


