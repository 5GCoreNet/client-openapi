/*
M1_PolicyTemplatesProvisioning

5GMS AF M1 Policy Templates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_PolicyTemplatesProvisioning

import (
	"encoding/json"
	"fmt"
)

// PolicyTemplateApiType struct for PolicyTemplateApiType
type PolicyTemplateApiType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyTemplateApiType) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyTemplateApiType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyTemplateApiType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyTemplateApiType struct {
	value *PolicyTemplateApiType
	isSet bool
}

func (v NullablePolicyTemplateApiType) Get() *PolicyTemplateApiType {
	return v.value
}

func (v *NullablePolicyTemplateApiType) Set(val *PolicyTemplateApiType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTemplateApiType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTemplateApiType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTemplateApiType(val *PolicyTemplateApiType) *NullablePolicyTemplateApiType {
	return &NullablePolicyTemplateApiType{value: val, isSet: true}
}

func (v NullablePolicyTemplateApiType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTemplateApiType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


