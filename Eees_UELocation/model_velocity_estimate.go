/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"fmt"
)

// VelocityEstimate - Velocity estimate.
type VelocityEstimate struct {
	HorizontalVelocity *HorizontalVelocity
	HorizontalVelocityWithUncertainty *HorizontalVelocityWithUncertainty
	HorizontalWithVerticalVelocity *HorizontalWithVerticalVelocity
	HorizontalWithVerticalVelocityAndUncertainty *HorizontalWithVerticalVelocityAndUncertainty
}

// HorizontalVelocityAsVelocityEstimate is a convenience function that returns HorizontalVelocity wrapped in VelocityEstimate
func HorizontalVelocityAsVelocityEstimate(v *HorizontalVelocity) VelocityEstimate {
	return VelocityEstimate{
		HorizontalVelocity: v,
	}
}

// HorizontalVelocityWithUncertaintyAsVelocityEstimate is a convenience function that returns HorizontalVelocityWithUncertainty wrapped in VelocityEstimate
func HorizontalVelocityWithUncertaintyAsVelocityEstimate(v *HorizontalVelocityWithUncertainty) VelocityEstimate {
	return VelocityEstimate{
		HorizontalVelocityWithUncertainty: v,
	}
}

// HorizontalWithVerticalVelocityAsVelocityEstimate is a convenience function that returns HorizontalWithVerticalVelocity wrapped in VelocityEstimate
func HorizontalWithVerticalVelocityAsVelocityEstimate(v *HorizontalWithVerticalVelocity) VelocityEstimate {
	return VelocityEstimate{
		HorizontalWithVerticalVelocity: v,
	}
}

// HorizontalWithVerticalVelocityAndUncertaintyAsVelocityEstimate is a convenience function that returns HorizontalWithVerticalVelocityAndUncertainty wrapped in VelocityEstimate
func HorizontalWithVerticalVelocityAndUncertaintyAsVelocityEstimate(v *HorizontalWithVerticalVelocityAndUncertainty) VelocityEstimate {
	return VelocityEstimate{
		HorizontalWithVerticalVelocityAndUncertainty: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VelocityEstimate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HorizontalVelocity
	err = newStrictDecoder(data).Decode(&dst.HorizontalVelocity)
	if err == nil {
		jsonHorizontalVelocity, _ := json.Marshal(dst.HorizontalVelocity)
		if string(jsonHorizontalVelocity) == "{}" { // empty struct
			dst.HorizontalVelocity = nil
		} else {
			match++
		}
	} else {
		dst.HorizontalVelocity = nil
	}

	// try to unmarshal data into HorizontalVelocityWithUncertainty
	err = newStrictDecoder(data).Decode(&dst.HorizontalVelocityWithUncertainty)
	if err == nil {
		jsonHorizontalVelocityWithUncertainty, _ := json.Marshal(dst.HorizontalVelocityWithUncertainty)
		if string(jsonHorizontalVelocityWithUncertainty) == "{}" { // empty struct
			dst.HorizontalVelocityWithUncertainty = nil
		} else {
			match++
		}
	} else {
		dst.HorizontalVelocityWithUncertainty = nil
	}

	// try to unmarshal data into HorizontalWithVerticalVelocity
	err = newStrictDecoder(data).Decode(&dst.HorizontalWithVerticalVelocity)
	if err == nil {
		jsonHorizontalWithVerticalVelocity, _ := json.Marshal(dst.HorizontalWithVerticalVelocity)
		if string(jsonHorizontalWithVerticalVelocity) == "{}" { // empty struct
			dst.HorizontalWithVerticalVelocity = nil
		} else {
			match++
		}
	} else {
		dst.HorizontalWithVerticalVelocity = nil
	}

	// try to unmarshal data into HorizontalWithVerticalVelocityAndUncertainty
	err = newStrictDecoder(data).Decode(&dst.HorizontalWithVerticalVelocityAndUncertainty)
	if err == nil {
		jsonHorizontalWithVerticalVelocityAndUncertainty, _ := json.Marshal(dst.HorizontalWithVerticalVelocityAndUncertainty)
		if string(jsonHorizontalWithVerticalVelocityAndUncertainty) == "{}" { // empty struct
			dst.HorizontalWithVerticalVelocityAndUncertainty = nil
		} else {
			match++
		}
	} else {
		dst.HorizontalWithVerticalVelocityAndUncertainty = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HorizontalVelocity = nil
		dst.HorizontalVelocityWithUncertainty = nil
		dst.HorizontalWithVerticalVelocity = nil
		dst.HorizontalWithVerticalVelocityAndUncertainty = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VelocityEstimate)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VelocityEstimate)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VelocityEstimate) MarshalJSON() ([]byte, error) {
	if src.HorizontalVelocity != nil {
		return json.Marshal(&src.HorizontalVelocity)
	}

	if src.HorizontalVelocityWithUncertainty != nil {
		return json.Marshal(&src.HorizontalVelocityWithUncertainty)
	}

	if src.HorizontalWithVerticalVelocity != nil {
		return json.Marshal(&src.HorizontalWithVerticalVelocity)
	}

	if src.HorizontalWithVerticalVelocityAndUncertainty != nil {
		return json.Marshal(&src.HorizontalWithVerticalVelocityAndUncertainty)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VelocityEstimate) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HorizontalVelocity != nil {
		return obj.HorizontalVelocity
	}

	if obj.HorizontalVelocityWithUncertainty != nil {
		return obj.HorizontalVelocityWithUncertainty
	}

	if obj.HorizontalWithVerticalVelocity != nil {
		return obj.HorizontalWithVerticalVelocity
	}

	if obj.HorizontalWithVerticalVelocityAndUncertainty != nil {
		return obj.HorizontalWithVerticalVelocityAndUncertainty
	}

	// all schemas are nil
	return nil
}

type NullableVelocityEstimate struct {
	value *VelocityEstimate
	isSet bool
}

func (v NullableVelocityEstimate) Get() *VelocityEstimate {
	return v.value
}

func (v *NullableVelocityEstimate) Set(val *VelocityEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityEstimate(val *VelocityEstimate) *NullableVelocityEstimate {
	return &NullableVelocityEstimate{value: val, isSet: true}
}

func (v NullableVelocityEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


