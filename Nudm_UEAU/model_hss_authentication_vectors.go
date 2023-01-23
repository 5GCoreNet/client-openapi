/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// HssAuthenticationVectors - struct for HssAuthenticationVectors
type HssAuthenticationVectors struct {
	ArrayOfAvEapAkaPrime *[]AvEapAkaPrime
	ArrayOfAvEpsAka *[]AvEpsAka
	ArrayOfAvImsGbaEapAka *[]AvImsGbaEapAka
}

// []AvEapAkaPrimeAsHssAuthenticationVectors is a convenience function that returns []AvEapAkaPrime wrapped in HssAuthenticationVectors
func ArrayOfAvEapAkaPrimeAsHssAuthenticationVectors(v *[]AvEapAkaPrime) HssAuthenticationVectors {
	return HssAuthenticationVectors{
		ArrayOfAvEapAkaPrime: v,
	}
}

// []AvEpsAkaAsHssAuthenticationVectors is a convenience function that returns []AvEpsAka wrapped in HssAuthenticationVectors
func ArrayOfAvEpsAkaAsHssAuthenticationVectors(v *[]AvEpsAka) HssAuthenticationVectors {
	return HssAuthenticationVectors{
		ArrayOfAvEpsAka: v,
	}
}

// []AvImsGbaEapAkaAsHssAuthenticationVectors is a convenience function that returns []AvImsGbaEapAka wrapped in HssAuthenticationVectors
func ArrayOfAvImsGbaEapAkaAsHssAuthenticationVectors(v *[]AvImsGbaEapAka) HssAuthenticationVectors {
	return HssAuthenticationVectors{
		ArrayOfAvImsGbaEapAka: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HssAuthenticationVectors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfAvEapAkaPrime
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAvEapAkaPrime)
	if err == nil {
		jsonArrayOfAvEapAkaPrime, _ := json.Marshal(dst.ArrayOfAvEapAkaPrime)
		if string(jsonArrayOfAvEapAkaPrime) == "{}" { // empty struct
			dst.ArrayOfAvEapAkaPrime = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAvEapAkaPrime = nil
	}

	// try to unmarshal data into ArrayOfAvEpsAka
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAvEpsAka)
	if err == nil {
		jsonArrayOfAvEpsAka, _ := json.Marshal(dst.ArrayOfAvEpsAka)
		if string(jsonArrayOfAvEpsAka) == "{}" { // empty struct
			dst.ArrayOfAvEpsAka = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAvEpsAka = nil
	}

	// try to unmarshal data into ArrayOfAvImsGbaEapAka
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAvImsGbaEapAka)
	if err == nil {
		jsonArrayOfAvImsGbaEapAka, _ := json.Marshal(dst.ArrayOfAvImsGbaEapAka)
		if string(jsonArrayOfAvImsGbaEapAka) == "{}" { // empty struct
			dst.ArrayOfAvImsGbaEapAka = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAvImsGbaEapAka = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfAvEapAkaPrime = nil
		dst.ArrayOfAvEpsAka = nil
		dst.ArrayOfAvImsGbaEapAka = nil

		return fmt.Errorf("data matches more than one schema in oneOf(HssAuthenticationVectors)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HssAuthenticationVectors)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HssAuthenticationVectors) MarshalJSON() ([]byte, error) {
	if src.ArrayOfAvEapAkaPrime != nil {
		return json.Marshal(&src.ArrayOfAvEapAkaPrime)
	}

	if src.ArrayOfAvEpsAka != nil {
		return json.Marshal(&src.ArrayOfAvEpsAka)
	}

	if src.ArrayOfAvImsGbaEapAka != nil {
		return json.Marshal(&src.ArrayOfAvImsGbaEapAka)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HssAuthenticationVectors) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfAvEapAkaPrime != nil {
		return obj.ArrayOfAvEapAkaPrime
	}

	if obj.ArrayOfAvEpsAka != nil {
		return obj.ArrayOfAvEpsAka
	}

	if obj.ArrayOfAvImsGbaEapAka != nil {
		return obj.ArrayOfAvImsGbaEapAka
	}

	// all schemas are nil
	return nil
}

type NullableHssAuthenticationVectors struct {
	value *HssAuthenticationVectors
	isSet bool
}

func (v NullableHssAuthenticationVectors) Get() *HssAuthenticationVectors {
	return v.value
}

func (v *NullableHssAuthenticationVectors) Set(val *HssAuthenticationVectors) {
	v.value = val
	v.isSet = true
}

func (v NullableHssAuthenticationVectors) IsSet() bool {
	return v.isSet
}

func (v *NullableHssAuthenticationVectors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssAuthenticationVectors(val *HssAuthenticationVectors) *NullableHssAuthenticationVectors {
	return &NullableHssAuthenticationVectors{value: val, isSet: true}
}

func (v NullableHssAuthenticationVectors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssAuthenticationVectors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


