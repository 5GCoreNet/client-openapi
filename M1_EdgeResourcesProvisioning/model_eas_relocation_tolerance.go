/*
M1_EdgeResourcesProvisioning

5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_EdgeResourcesProvisioning

import (
	"encoding/json"
	"fmt"
)

// EASRelocationTolerance struct for EASRelocationTolerance
type EASRelocationTolerance struct {
	EASRelocationToleranceAnyOf *EASRelocationToleranceAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EASRelocationTolerance) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EASRelocationToleranceAnyOf
	err = json.Unmarshal(data, &dst.EASRelocationToleranceAnyOf);
	if err == nil {
		jsonEASRelocationToleranceAnyOf, _ := json.Marshal(dst.EASRelocationToleranceAnyOf)
		if string(jsonEASRelocationToleranceAnyOf) == "{}" { // empty struct
			dst.EASRelocationToleranceAnyOf = nil
		} else {
			return nil // data stored in dst.EASRelocationToleranceAnyOf, return on the first match
		}
	} else {
		dst.EASRelocationToleranceAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(EASRelocationTolerance)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EASRelocationTolerance) MarshalJSON() ([]byte, error) {
	if src.EASRelocationToleranceAnyOf != nil {
		return json.Marshal(&src.EASRelocationToleranceAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEASRelocationTolerance struct {
	value *EASRelocationTolerance
	isSet bool
}

func (v NullableEASRelocationTolerance) Get() *EASRelocationTolerance {
	return v.value
}

func (v *NullableEASRelocationTolerance) Set(val *EASRelocationTolerance) {
	v.value = val
	v.isSet = true
}

func (v NullableEASRelocationTolerance) IsSet() bool {
	return v.isSet
}

func (v *NullableEASRelocationTolerance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASRelocationTolerance(val *EASRelocationTolerance) *NullableEASRelocationTolerance {
	return &NullableEASRelocationTolerance{value: val, isSet: true}
}

func (v NullableEASRelocationTolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASRelocationTolerance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


