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

// ResourcesMdaNrm - struct for ResourcesMdaNrm
type ResourcesMdaNrm struct {
	MDAFunctionSingle *MDAFunctionSingle
	MDAReport *MDAReport
	MDARequestSingle *MDARequestSingle
	ManagedElementSingle3 *ManagedElementSingle3
	SubNetworkSingle4 *SubNetworkSingle4
}

// MDAFunctionSingleAsResourcesMdaNrm is a convenience function that returns MDAFunctionSingle wrapped in ResourcesMdaNrm
func MDAFunctionSingleAsResourcesMdaNrm(v *MDAFunctionSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDAFunctionSingle: v,
	}
}

// MDAReportAsResourcesMdaNrm is a convenience function that returns MDAReport wrapped in ResourcesMdaNrm
func MDAReportAsResourcesMdaNrm(v *MDAReport) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDAReport: v,
	}
}

// MDARequestSingleAsResourcesMdaNrm is a convenience function that returns MDARequestSingle wrapped in ResourcesMdaNrm
func MDARequestSingleAsResourcesMdaNrm(v *MDARequestSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDARequestSingle: v,
	}
}

// ManagedElementSingle3AsResourcesMdaNrm is a convenience function that returns ManagedElementSingle3 wrapped in ResourcesMdaNrm
func ManagedElementSingle3AsResourcesMdaNrm(v *ManagedElementSingle3) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		ManagedElementSingle3: v,
	}
}

// SubNetworkSingle4AsResourcesMdaNrm is a convenience function that returns SubNetworkSingle4 wrapped in ResourcesMdaNrm
func SubNetworkSingle4AsResourcesMdaNrm(v *SubNetworkSingle4) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		SubNetworkSingle4: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcesMdaNrm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MDAFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.MDAFunctionSingle)
	if err == nil {
		jsonMDAFunctionSingle, _ := json.Marshal(dst.MDAFunctionSingle)
		if string(jsonMDAFunctionSingle) == "{}" { // empty struct
			dst.MDAFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.MDAFunctionSingle = nil
	}

	// try to unmarshal data into MDAReport
	err = newStrictDecoder(data).Decode(&dst.MDAReport)
	if err == nil {
		jsonMDAReport, _ := json.Marshal(dst.MDAReport)
		if string(jsonMDAReport) == "{}" { // empty struct
			dst.MDAReport = nil
		} else {
			match++
		}
	} else {
		dst.MDAReport = nil
	}

	// try to unmarshal data into MDARequestSingle
	err = newStrictDecoder(data).Decode(&dst.MDARequestSingle)
	if err == nil {
		jsonMDARequestSingle, _ := json.Marshal(dst.MDARequestSingle)
		if string(jsonMDARequestSingle) == "{}" { // empty struct
			dst.MDARequestSingle = nil
		} else {
			match++
		}
	} else {
		dst.MDARequestSingle = nil
	}

	// try to unmarshal data into ManagedElementSingle3
	err = newStrictDecoder(data).Decode(&dst.ManagedElementSingle3)
	if err == nil {
		jsonManagedElementSingle3, _ := json.Marshal(dst.ManagedElementSingle3)
		if string(jsonManagedElementSingle3) == "{}" { // empty struct
			dst.ManagedElementSingle3 = nil
		} else {
			match++
		}
	} else {
		dst.ManagedElementSingle3 = nil
	}

	// try to unmarshal data into SubNetworkSingle4
	err = newStrictDecoder(data).Decode(&dst.SubNetworkSingle4)
	if err == nil {
		jsonSubNetworkSingle4, _ := json.Marshal(dst.SubNetworkSingle4)
		if string(jsonSubNetworkSingle4) == "{}" { // empty struct
			dst.SubNetworkSingle4 = nil
		} else {
			match++
		}
	} else {
		dst.SubNetworkSingle4 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MDAFunctionSingle = nil
		dst.MDAReport = nil
		dst.MDARequestSingle = nil
		dst.ManagedElementSingle3 = nil
		dst.SubNetworkSingle4 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResourcesMdaNrm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResourcesMdaNrm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcesMdaNrm) MarshalJSON() ([]byte, error) {
	if src.MDAFunctionSingle != nil {
		return json.Marshal(&src.MDAFunctionSingle)
	}

	if src.MDAReport != nil {
		return json.Marshal(&src.MDAReport)
	}

	if src.MDARequestSingle != nil {
		return json.Marshal(&src.MDARequestSingle)
	}

	if src.ManagedElementSingle3 != nil {
		return json.Marshal(&src.ManagedElementSingle3)
	}

	if src.SubNetworkSingle4 != nil {
		return json.Marshal(&src.SubNetworkSingle4)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResourcesMdaNrm) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MDAFunctionSingle != nil {
		return obj.MDAFunctionSingle
	}

	if obj.MDAReport != nil {
		return obj.MDAReport
	}

	if obj.MDARequestSingle != nil {
		return obj.MDARequestSingle
	}

	if obj.ManagedElementSingle3 != nil {
		return obj.ManagedElementSingle3
	}

	if obj.SubNetworkSingle4 != nil {
		return obj.SubNetworkSingle4
	}

	// all schemas are nil
	return nil
}

type NullableResourcesMdaNrm struct {
	value *ResourcesMdaNrm
	isSet bool
}

func (v NullableResourcesMdaNrm) Get() *ResourcesMdaNrm {
	return v.value
}

func (v *NullableResourcesMdaNrm) Set(val *ResourcesMdaNrm) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesMdaNrm) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesMdaNrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesMdaNrm(val *ResourcesMdaNrm) *NullableResourcesMdaNrm {
	return &NullableResourcesMdaNrm{value: val, isSet: true}
}

func (v NullableResourcesMdaNrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesMdaNrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


