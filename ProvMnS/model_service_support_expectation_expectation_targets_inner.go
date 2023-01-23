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

// ServiceSupportExpectationExpectationTargetsInner - struct for ServiceSupportExpectationExpectationTargetsInner
type ServiceSupportExpectationExpectationTargetsInner struct {
	ActivityFactorTarget *ActivityFactorTarget
	DLLatencyTarget *DLLatencyTarget
	DLThptPerUETarget *DLThptPerUETarget
	ExpectationTarget *ExpectationTarget
	MaxNumberofUEsTarget *MaxNumberofUEsTarget
	UESpeedTarget *UESpeedTarget
	ULLatencyTarget *ULLatencyTarget
	ULThptPerUETarget *ULThptPerUETarget
}

// ActivityFactorTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns ActivityFactorTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func ActivityFactorTargetAsServiceSupportExpectationExpectationTargetsInner(v *ActivityFactorTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		ActivityFactorTarget: v,
	}
}

// DLLatencyTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns DLLatencyTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func DLLatencyTargetAsServiceSupportExpectationExpectationTargetsInner(v *DLLatencyTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		DLLatencyTarget: v,
	}
}

// DLThptPerUETargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns DLThptPerUETarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func DLThptPerUETargetAsServiceSupportExpectationExpectationTargetsInner(v *DLThptPerUETarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		DLThptPerUETarget: v,
	}
}

// ExpectationTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns ExpectationTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func ExpectationTargetAsServiceSupportExpectationExpectationTargetsInner(v *ExpectationTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		ExpectationTarget: v,
	}
}

// MaxNumberofUEsTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns MaxNumberofUEsTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func MaxNumberofUEsTargetAsServiceSupportExpectationExpectationTargetsInner(v *MaxNumberofUEsTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		MaxNumberofUEsTarget: v,
	}
}

// UESpeedTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns UESpeedTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func UESpeedTargetAsServiceSupportExpectationExpectationTargetsInner(v *UESpeedTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		UESpeedTarget: v,
	}
}

// ULLatencyTargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns ULLatencyTarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func ULLatencyTargetAsServiceSupportExpectationExpectationTargetsInner(v *ULLatencyTarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		ULLatencyTarget: v,
	}
}

// ULThptPerUETargetAsServiceSupportExpectationExpectationTargetsInner is a convenience function that returns ULThptPerUETarget wrapped in ServiceSupportExpectationExpectationTargetsInner
func ULThptPerUETargetAsServiceSupportExpectationExpectationTargetsInner(v *ULThptPerUETarget) ServiceSupportExpectationExpectationTargetsInner {
	return ServiceSupportExpectationExpectationTargetsInner{
		ULThptPerUETarget: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServiceSupportExpectationExpectationTargetsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActivityFactorTarget
	err = newStrictDecoder(data).Decode(&dst.ActivityFactorTarget)
	if err == nil {
		jsonActivityFactorTarget, _ := json.Marshal(dst.ActivityFactorTarget)
		if string(jsonActivityFactorTarget) == "{}" { // empty struct
			dst.ActivityFactorTarget = nil
		} else {
			match++
		}
	} else {
		dst.ActivityFactorTarget = nil
	}

	// try to unmarshal data into DLLatencyTarget
	err = newStrictDecoder(data).Decode(&dst.DLLatencyTarget)
	if err == nil {
		jsonDLLatencyTarget, _ := json.Marshal(dst.DLLatencyTarget)
		if string(jsonDLLatencyTarget) == "{}" { // empty struct
			dst.DLLatencyTarget = nil
		} else {
			match++
		}
	} else {
		dst.DLLatencyTarget = nil
	}

	// try to unmarshal data into DLThptPerUETarget
	err = newStrictDecoder(data).Decode(&dst.DLThptPerUETarget)
	if err == nil {
		jsonDLThptPerUETarget, _ := json.Marshal(dst.DLThptPerUETarget)
		if string(jsonDLThptPerUETarget) == "{}" { // empty struct
			dst.DLThptPerUETarget = nil
		} else {
			match++
		}
	} else {
		dst.DLThptPerUETarget = nil
	}

	// try to unmarshal data into ExpectationTarget
	err = newStrictDecoder(data).Decode(&dst.ExpectationTarget)
	if err == nil {
		jsonExpectationTarget, _ := json.Marshal(dst.ExpectationTarget)
		if string(jsonExpectationTarget) == "{}" { // empty struct
			dst.ExpectationTarget = nil
		} else {
			match++
		}
	} else {
		dst.ExpectationTarget = nil
	}

	// try to unmarshal data into MaxNumberofUEsTarget
	err = newStrictDecoder(data).Decode(&dst.MaxNumberofUEsTarget)
	if err == nil {
		jsonMaxNumberofUEsTarget, _ := json.Marshal(dst.MaxNumberofUEsTarget)
		if string(jsonMaxNumberofUEsTarget) == "{}" { // empty struct
			dst.MaxNumberofUEsTarget = nil
		} else {
			match++
		}
	} else {
		dst.MaxNumberofUEsTarget = nil
	}

	// try to unmarshal data into UESpeedTarget
	err = newStrictDecoder(data).Decode(&dst.UESpeedTarget)
	if err == nil {
		jsonUESpeedTarget, _ := json.Marshal(dst.UESpeedTarget)
		if string(jsonUESpeedTarget) == "{}" { // empty struct
			dst.UESpeedTarget = nil
		} else {
			match++
		}
	} else {
		dst.UESpeedTarget = nil
	}

	// try to unmarshal data into ULLatencyTarget
	err = newStrictDecoder(data).Decode(&dst.ULLatencyTarget)
	if err == nil {
		jsonULLatencyTarget, _ := json.Marshal(dst.ULLatencyTarget)
		if string(jsonULLatencyTarget) == "{}" { // empty struct
			dst.ULLatencyTarget = nil
		} else {
			match++
		}
	} else {
		dst.ULLatencyTarget = nil
	}

	// try to unmarshal data into ULThptPerUETarget
	err = newStrictDecoder(data).Decode(&dst.ULThptPerUETarget)
	if err == nil {
		jsonULThptPerUETarget, _ := json.Marshal(dst.ULThptPerUETarget)
		if string(jsonULThptPerUETarget) == "{}" { // empty struct
			dst.ULThptPerUETarget = nil
		} else {
			match++
		}
	} else {
		dst.ULThptPerUETarget = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActivityFactorTarget = nil
		dst.DLLatencyTarget = nil
		dst.DLThptPerUETarget = nil
		dst.ExpectationTarget = nil
		dst.MaxNumberofUEsTarget = nil
		dst.UESpeedTarget = nil
		dst.ULLatencyTarget = nil
		dst.ULThptPerUETarget = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServiceSupportExpectationExpectationTargetsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServiceSupportExpectationExpectationTargetsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServiceSupportExpectationExpectationTargetsInner) MarshalJSON() ([]byte, error) {
	if src.ActivityFactorTarget != nil {
		return json.Marshal(&src.ActivityFactorTarget)
	}

	if src.DLLatencyTarget != nil {
		return json.Marshal(&src.DLLatencyTarget)
	}

	if src.DLThptPerUETarget != nil {
		return json.Marshal(&src.DLThptPerUETarget)
	}

	if src.ExpectationTarget != nil {
		return json.Marshal(&src.ExpectationTarget)
	}

	if src.MaxNumberofUEsTarget != nil {
		return json.Marshal(&src.MaxNumberofUEsTarget)
	}

	if src.UESpeedTarget != nil {
		return json.Marshal(&src.UESpeedTarget)
	}

	if src.ULLatencyTarget != nil {
		return json.Marshal(&src.ULLatencyTarget)
	}

	if src.ULThptPerUETarget != nil {
		return json.Marshal(&src.ULThptPerUETarget)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServiceSupportExpectationExpectationTargetsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActivityFactorTarget != nil {
		return obj.ActivityFactorTarget
	}

	if obj.DLLatencyTarget != nil {
		return obj.DLLatencyTarget
	}

	if obj.DLThptPerUETarget != nil {
		return obj.DLThptPerUETarget
	}

	if obj.ExpectationTarget != nil {
		return obj.ExpectationTarget
	}

	if obj.MaxNumberofUEsTarget != nil {
		return obj.MaxNumberofUEsTarget
	}

	if obj.UESpeedTarget != nil {
		return obj.UESpeedTarget
	}

	if obj.ULLatencyTarget != nil {
		return obj.ULLatencyTarget
	}

	if obj.ULThptPerUETarget != nil {
		return obj.ULThptPerUETarget
	}

	// all schemas are nil
	return nil
}

type NullableServiceSupportExpectationExpectationTargetsInner struct {
	value *ServiceSupportExpectationExpectationTargetsInner
	isSet bool
}

func (v NullableServiceSupportExpectationExpectationTargetsInner) Get() *ServiceSupportExpectationExpectationTargetsInner {
	return v.value
}

func (v *NullableServiceSupportExpectationExpectationTargetsInner) Set(val *ServiceSupportExpectationExpectationTargetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSupportExpectationExpectationTargetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSupportExpectationExpectationTargetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSupportExpectationExpectationTargetsInner(val *ServiceSupportExpectationExpectationTargetsInner) *NullableServiceSupportExpectationExpectationTargetsInner {
	return &NullableServiceSupportExpectationExpectationTargetsInner{value: val, isSet: true}
}

func (v NullableServiceSupportExpectationExpectationTargetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSupportExpectationExpectationTargetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


