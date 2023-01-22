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

// PdnEstablishmentOptions Possible values are - WAIT_FOR_UE: wait for the UE to establish the PDN connection  - INDICATE_ERROR: respond with an error cause - SEND_TRIGGER: send a device trigger 
type PdnEstablishmentOptions struct {
	PdnEstablishmentOptionsAnyOf *PdnEstablishmentOptionsAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PdnEstablishmentOptions) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PdnEstablishmentOptionsAnyOf
	err = json.Unmarshal(data, &dst.PdnEstablishmentOptionsAnyOf);
	if err == nil {
		jsonPdnEstablishmentOptionsAnyOf, _ := json.Marshal(dst.PdnEstablishmentOptionsAnyOf)
		if string(jsonPdnEstablishmentOptionsAnyOf) == "{}" { // empty struct
			dst.PdnEstablishmentOptionsAnyOf = nil
		} else {
			return nil // data stored in dst.PdnEstablishmentOptionsAnyOf, return on the first match
		}
	} else {
		dst.PdnEstablishmentOptionsAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PdnEstablishmentOptions)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PdnEstablishmentOptions) MarshalJSON() ([]byte, error) {
	if src.PdnEstablishmentOptionsAnyOf != nil {
		return json.Marshal(&src.PdnEstablishmentOptionsAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePdnEstablishmentOptions struct {
	value *PdnEstablishmentOptions
	isSet bool
}

func (v NullablePdnEstablishmentOptions) Get() *PdnEstablishmentOptions {
	return v.value
}

func (v *NullablePdnEstablishmentOptions) Set(val *PdnEstablishmentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnEstablishmentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnEstablishmentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnEstablishmentOptions(val *PdnEstablishmentOptions) *NullablePdnEstablishmentOptions {
	return &NullablePdnEstablishmentOptions{value: val, isSet: true}
}

func (v NullablePdnEstablishmentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnEstablishmentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


