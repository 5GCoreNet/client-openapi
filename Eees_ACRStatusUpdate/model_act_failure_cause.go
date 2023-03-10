/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
	"fmt"
)

// ACTFailureCause Possible values are: - ACR_CANCELLATION: Indicates that the ACT failed due to the cancellation of the ACR. - OTHER: Indicates that the ACT failed for other reasons. 
type ACTFailureCause struct {
	ACTFailureCauseAnyOf *ACTFailureCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ACTFailureCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ACTFailureCauseAnyOf
	err = json.Unmarshal(data, &dst.ACTFailureCauseAnyOf);
	if err == nil {
		jsonACTFailureCauseAnyOf, _ := json.Marshal(dst.ACTFailureCauseAnyOf)
		if string(jsonACTFailureCauseAnyOf) == "{}" { // empty struct
			dst.ACTFailureCauseAnyOf = nil
		} else {
			return nil // data stored in dst.ACTFailureCauseAnyOf, return on the first match
		}
	} else {
		dst.ACTFailureCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ACTFailureCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ACTFailureCause) MarshalJSON() ([]byte, error) {
	if src.ACTFailureCauseAnyOf != nil {
		return json.Marshal(&src.ACTFailureCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableACTFailureCause struct {
	value *ACTFailureCause
	isSet bool
}

func (v NullableACTFailureCause) Get() *ACTFailureCause {
	return v.value
}

func (v *NullableACTFailureCause) Set(val *ACTFailureCause) {
	v.value = val
	v.isSet = true
}

func (v NullableACTFailureCause) IsSet() bool {
	return v.isSet
}

func (v *NullableACTFailureCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACTFailureCause(val *ACTFailureCause) *NullableACTFailureCause {
	return &NullableACTFailureCause{value: val, isSet: true}
}

func (v NullableACTFailureCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACTFailureCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


