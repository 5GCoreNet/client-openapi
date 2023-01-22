/*
M1_ProvisioningSessions

5GMS AF M1 Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_ProvisioningSessions

import (
	"encoding/json"
	"fmt"
)

// ProvisioningSessionType struct for ProvisioningSessionType
type ProvisioningSessionType struct {
	ProvisioningSessionTypeAnyOf *ProvisioningSessionTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProvisioningSessionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ProvisioningSessionTypeAnyOf
	err = json.Unmarshal(data, &dst.ProvisioningSessionTypeAnyOf);
	if err == nil {
		jsonProvisioningSessionTypeAnyOf, _ := json.Marshal(dst.ProvisioningSessionTypeAnyOf)
		if string(jsonProvisioningSessionTypeAnyOf) == "{}" { // empty struct
			dst.ProvisioningSessionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ProvisioningSessionTypeAnyOf, return on the first match
		}
	} else {
		dst.ProvisioningSessionTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ProvisioningSessionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProvisioningSessionType) MarshalJSON() ([]byte, error) {
	if src.ProvisioningSessionTypeAnyOf != nil {
		return json.Marshal(&src.ProvisioningSessionTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProvisioningSessionType struct {
	value *ProvisioningSessionType
	isSet bool
}

func (v NullableProvisioningSessionType) Get() *ProvisioningSessionType {
	return v.value
}

func (v *NullableProvisioningSessionType) Set(val *ProvisioningSessionType) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSessionType) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSessionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSessionType(val *ProvisioningSessionType) *NullableProvisioningSessionType {
	return &NullableProvisioningSessionType{value: val, isSet: true}
}

func (v NullableProvisioningSessionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSessionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


