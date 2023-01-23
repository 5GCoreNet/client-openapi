/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"fmt"
)

// ProbableCause - The value of the probable cause may be a specific standardized string, or any vendor provided string. Probable cause strings are not standardized in the present document. They may be added in a future version. Up to then the mapping of the generic probable cause strings \"PROBABLE_CAUSE_001\" to \"PROBABLE_CAUSE_005\" is vendor specific. The value of the probable cause may also be an integer. The mapping of integer values to probable causes is vendor specific.
type ProbableCause struct {
	AnyOfstringstring *AnyOfstringstring
	Int32 *int32
}

// AnyOfstringstringAsProbableCause is a convenience function that returns AnyOfstringstring wrapped in ProbableCause
func AnyOfstringstringAsProbableCause(v *AnyOfstringstring) ProbableCause {
	return ProbableCause{
		AnyOfstringstring: v,
	}
}

// int32AsProbableCause is a convenience function that returns int32 wrapped in ProbableCause
func Int32AsProbableCause(v *int32) ProbableCause {
	return ProbableCause{
		Int32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProbableCause) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnyOfstringstring
	err = newStrictDecoder(data).Decode(&dst.AnyOfstringstring)
	if err == nil {
		jsonAnyOfstringstring, _ := json.Marshal(dst.AnyOfstringstring)
		if string(jsonAnyOfstringstring) == "{}" { // empty struct
			dst.AnyOfstringstring = nil
		} else {
			match++
		}
	} else {
		dst.AnyOfstringstring = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AnyOfstringstring = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProbableCause)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProbableCause)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProbableCause) MarshalJSON() ([]byte, error) {
	if src.AnyOfstringstring != nil {
		return json.Marshal(&src.AnyOfstringstring)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProbableCause) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AnyOfstringstring != nil {
		return obj.AnyOfstringstring
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableProbableCause struct {
	value *ProbableCause
	isSet bool
}

func (v NullableProbableCause) Get() *ProbableCause {
	return v.value
}

func (v *NullableProbableCause) Set(val *ProbableCause) {
	v.value = val
	v.isSet = true
}

func (v NullableProbableCause) IsSet() bool {
	return v.isSet
}

func (v *NullableProbableCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbableCause(val *ProbableCause) *NullableProbableCause {
	return &NullableProbableCause{value: val, isSet: true}
}

func (v NullableProbableCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbableCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


