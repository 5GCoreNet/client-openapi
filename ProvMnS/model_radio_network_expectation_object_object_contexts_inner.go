/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// RadioNetworkExpectationObjectObjectContextsInner - struct for RadioNetworkExpectationObjectObjectContextsInner
type RadioNetworkExpectationObjectObjectContextsInner struct {
	CoverageAreaPolygonContext *CoverageAreaPolygonContext
	CoverageTACContext *CoverageTACContext
	NRFqBandContext *NRFqBandContext
	ObjectContext *ObjectContext
	PLMNContext *PLMNContext
	RATContext *RATContext
}

// CoverageAreaPolygonContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns CoverageAreaPolygonContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func CoverageAreaPolygonContextAsRadioNetworkExpectationObjectObjectContextsInner(v *CoverageAreaPolygonContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		CoverageAreaPolygonContext: v,
	}
}

// CoverageTACContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns CoverageTACContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func CoverageTACContextAsRadioNetworkExpectationObjectObjectContextsInner(v *CoverageTACContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		CoverageTACContext: v,
	}
}

// NRFqBandContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns NRFqBandContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func NRFqBandContextAsRadioNetworkExpectationObjectObjectContextsInner(v *NRFqBandContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		NRFqBandContext: v,
	}
}

// ObjectContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns ObjectContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func ObjectContextAsRadioNetworkExpectationObjectObjectContextsInner(v *ObjectContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		ObjectContext: v,
	}
}

// PLMNContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns PLMNContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func PLMNContextAsRadioNetworkExpectationObjectObjectContextsInner(v *PLMNContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		PLMNContext: v,
	}
}

// RATContextAsRadioNetworkExpectationObjectObjectContextsInner is a convenience function that returns RATContext wrapped in RadioNetworkExpectationObjectObjectContextsInner
func RATContextAsRadioNetworkExpectationObjectObjectContextsInner(v *RATContext) RadioNetworkExpectationObjectObjectContextsInner {
	return RadioNetworkExpectationObjectObjectContextsInner{
		RATContext: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RadioNetworkExpectationObjectObjectContextsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CoverageAreaPolygonContext
	err = newStrictDecoder(data).Decode(&dst.CoverageAreaPolygonContext)
	if err == nil {
		jsonCoverageAreaPolygonContext, _ := json.Marshal(dst.CoverageAreaPolygonContext)
		if string(jsonCoverageAreaPolygonContext) == "{}" { // empty struct
			dst.CoverageAreaPolygonContext = nil
		} else {
			match++
		}
	} else {
		dst.CoverageAreaPolygonContext = nil
	}

	// try to unmarshal data into CoverageTACContext
	err = newStrictDecoder(data).Decode(&dst.CoverageTACContext)
	if err == nil {
		jsonCoverageTACContext, _ := json.Marshal(dst.CoverageTACContext)
		if string(jsonCoverageTACContext) == "{}" { // empty struct
			dst.CoverageTACContext = nil
		} else {
			match++
		}
	} else {
		dst.CoverageTACContext = nil
	}

	// try to unmarshal data into NRFqBandContext
	err = newStrictDecoder(data).Decode(&dst.NRFqBandContext)
	if err == nil {
		jsonNRFqBandContext, _ := json.Marshal(dst.NRFqBandContext)
		if string(jsonNRFqBandContext) == "{}" { // empty struct
			dst.NRFqBandContext = nil
		} else {
			match++
		}
	} else {
		dst.NRFqBandContext = nil
	}

	// try to unmarshal data into ObjectContext
	err = newStrictDecoder(data).Decode(&dst.ObjectContext)
	if err == nil {
		jsonObjectContext, _ := json.Marshal(dst.ObjectContext)
		if string(jsonObjectContext) == "{}" { // empty struct
			dst.ObjectContext = nil
		} else {
			match++
		}
	} else {
		dst.ObjectContext = nil
	}

	// try to unmarshal data into PLMNContext
	err = newStrictDecoder(data).Decode(&dst.PLMNContext)
	if err == nil {
		jsonPLMNContext, _ := json.Marshal(dst.PLMNContext)
		if string(jsonPLMNContext) == "{}" { // empty struct
			dst.PLMNContext = nil
		} else {
			match++
		}
	} else {
		dst.PLMNContext = nil
	}

	// try to unmarshal data into RATContext
	err = newStrictDecoder(data).Decode(&dst.RATContext)
	if err == nil {
		jsonRATContext, _ := json.Marshal(dst.RATContext)
		if string(jsonRATContext) == "{}" { // empty struct
			dst.RATContext = nil
		} else {
			match++
		}
	} else {
		dst.RATContext = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CoverageAreaPolygonContext = nil
		dst.CoverageTACContext = nil
		dst.NRFqBandContext = nil
		dst.ObjectContext = nil
		dst.PLMNContext = nil
		dst.RATContext = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RadioNetworkExpectationObjectObjectContextsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RadioNetworkExpectationObjectObjectContextsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RadioNetworkExpectationObjectObjectContextsInner) MarshalJSON() ([]byte, error) {
	if src.CoverageAreaPolygonContext != nil {
		return json.Marshal(&src.CoverageAreaPolygonContext)
	}

	if src.CoverageTACContext != nil {
		return json.Marshal(&src.CoverageTACContext)
	}

	if src.NRFqBandContext != nil {
		return json.Marshal(&src.NRFqBandContext)
	}

	if src.ObjectContext != nil {
		return json.Marshal(&src.ObjectContext)
	}

	if src.PLMNContext != nil {
		return json.Marshal(&src.PLMNContext)
	}

	if src.RATContext != nil {
		return json.Marshal(&src.RATContext)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RadioNetworkExpectationObjectObjectContextsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CoverageAreaPolygonContext != nil {
		return obj.CoverageAreaPolygonContext
	}

	if obj.CoverageTACContext != nil {
		return obj.CoverageTACContext
	}

	if obj.NRFqBandContext != nil {
		return obj.NRFqBandContext
	}

	if obj.ObjectContext != nil {
		return obj.ObjectContext
	}

	if obj.PLMNContext != nil {
		return obj.PLMNContext
	}

	if obj.RATContext != nil {
		return obj.RATContext
	}

	// all schemas are nil
	return nil
}

type NullableRadioNetworkExpectationObjectObjectContextsInner struct {
	value *RadioNetworkExpectationObjectObjectContextsInner
	isSet bool
}

func (v NullableRadioNetworkExpectationObjectObjectContextsInner) Get() *RadioNetworkExpectationObjectObjectContextsInner {
	return v.value
}

func (v *NullableRadioNetworkExpectationObjectObjectContextsInner) Set(val *RadioNetworkExpectationObjectObjectContextsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioNetworkExpectationObjectObjectContextsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioNetworkExpectationObjectObjectContextsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioNetworkExpectationObjectObjectContextsInner(val *RadioNetworkExpectationObjectObjectContextsInner) *NullableRadioNetworkExpectationObjectObjectContextsInner {
	return &NullableRadioNetworkExpectationObjectObjectContextsInner{value: val, isSet: true}
}

func (v NullableRadioNetworkExpectationObjectObjectContextsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioNetworkExpectationObjectObjectContextsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


