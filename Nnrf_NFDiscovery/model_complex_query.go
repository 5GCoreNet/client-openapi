/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// ComplexQuery - The ComplexQuery data type is either a conjunctive normal form or a disjunctive normal form.  The attribute names \"cnfUnits\" and \"dnfUnits\" (see clause 5.2.4.11 and clause 5.2.4.12)  serve as discriminator. 
type ComplexQuery struct {
	Cnf *Cnf
	Dnf *Dnf
}

// CnfAsComplexQuery is a convenience function that returns Cnf wrapped in ComplexQuery
func CnfAsComplexQuery(v *Cnf) ComplexQuery {
	return ComplexQuery{
		Cnf: v,
	}
}

// DnfAsComplexQuery is a convenience function that returns Dnf wrapped in ComplexQuery
func DnfAsComplexQuery(v *Dnf) ComplexQuery {
	return ComplexQuery{
		Dnf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComplexQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Cnf
	err = newStrictDecoder(data).Decode(&dst.Cnf)
	if err == nil {
		jsonCnf, _ := json.Marshal(dst.Cnf)
		if string(jsonCnf) == "{}" { // empty struct
			dst.Cnf = nil
		} else {
			match++
		}
	} else {
		dst.Cnf = nil
	}

	// try to unmarshal data into Dnf
	err = newStrictDecoder(data).Decode(&dst.Dnf)
	if err == nil {
		jsonDnf, _ := json.Marshal(dst.Dnf)
		if string(jsonDnf) == "{}" { // empty struct
			dst.Dnf = nil
		} else {
			match++
		}
	} else {
		dst.Dnf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Cnf = nil
		dst.Dnf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ComplexQuery)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ComplexQuery)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComplexQuery) MarshalJSON() ([]byte, error) {
	if src.Cnf != nil {
		return json.Marshal(&src.Cnf)
	}

	if src.Dnf != nil {
		return json.Marshal(&src.Dnf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComplexQuery) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Cnf != nil {
		return obj.Cnf
	}

	if obj.Dnf != nil {
		return obj.Dnf
	}

	// all schemas are nil
	return nil
}

type NullableComplexQuery struct {
	value *ComplexQuery
	isSet bool
}

func (v NullableComplexQuery) Get() *ComplexQuery {
	return v.value
}

func (v *NullableComplexQuery) Set(val *ComplexQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableComplexQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableComplexQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplexQuery(val *ComplexQuery) *NullableComplexQuery {
	return &NullableComplexQuery{value: val, isSet: true}
}

func (v NullableComplexQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplexQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


