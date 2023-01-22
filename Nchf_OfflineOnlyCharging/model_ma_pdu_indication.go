/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// MaPduIndication Contains the MA PDU session indication, i.e., MA PDU Request or MA PDU Network-Upgrade Allowed.
type MaPduIndication struct {
	MaPduIndicationAnyOf *MaPduIndicationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MaPduIndication) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MaPduIndicationAnyOf
	err = json.Unmarshal(data, &dst.MaPduIndicationAnyOf);
	if err == nil {
		jsonMaPduIndicationAnyOf, _ := json.Marshal(dst.MaPduIndicationAnyOf)
		if string(jsonMaPduIndicationAnyOf) == "{}" { // empty struct
			dst.MaPduIndicationAnyOf = nil
		} else {
			return nil // data stored in dst.MaPduIndicationAnyOf, return on the first match
		}
	} else {
		dst.MaPduIndicationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MaPduIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MaPduIndication) MarshalJSON() ([]byte, error) {
	if src.MaPduIndicationAnyOf != nil {
		return json.Marshal(&src.MaPduIndicationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMaPduIndication struct {
	value *MaPduIndication
	isSet bool
}

func (v NullableMaPduIndication) Get() *MaPduIndication {
	return v.value
}

func (v *NullableMaPduIndication) Set(val *MaPduIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableMaPduIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableMaPduIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaPduIndication(val *MaPduIndication) *NullableMaPduIndication {
	return &NullableMaPduIndication{value: val, isSet: true}
}

func (v NullableMaPduIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaPduIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


