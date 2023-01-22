/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_ContentHostingProvisioning

import (
	"encoding/json"
	"fmt"
)

// DistributionNetworkType Type of distribution network.
type DistributionNetworkType struct {
	DistributionNetworkTypeAnyOf *DistributionNetworkTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DistributionNetworkType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DistributionNetworkTypeAnyOf
	err = json.Unmarshal(data, &dst.DistributionNetworkTypeAnyOf);
	if err == nil {
		jsonDistributionNetworkTypeAnyOf, _ := json.Marshal(dst.DistributionNetworkTypeAnyOf)
		if string(jsonDistributionNetworkTypeAnyOf) == "{}" { // empty struct
			dst.DistributionNetworkTypeAnyOf = nil
		} else {
			return nil // data stored in dst.DistributionNetworkTypeAnyOf, return on the first match
		}
	} else {
		dst.DistributionNetworkTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DistributionNetworkType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DistributionNetworkType) MarshalJSON() ([]byte, error) {
	if src.DistributionNetworkTypeAnyOf != nil {
		return json.Marshal(&src.DistributionNetworkTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDistributionNetworkType struct {
	value *DistributionNetworkType
	isSet bool
}

func (v NullableDistributionNetworkType) Get() *DistributionNetworkType {
	return v.value
}

func (v *NullableDistributionNetworkType) Set(val *DistributionNetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionNetworkType(val *DistributionNetworkType) *NullableDistributionNetworkType {
	return &NullableDistributionNetworkType{value: val, isSet: true}
}

func (v NullableDistributionNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


