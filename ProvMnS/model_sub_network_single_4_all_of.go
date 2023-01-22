/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// SubNetworkSingle4AllOf struct for SubNetworkSingle4AllOf
type SubNetworkSingle4AllOf struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
	MDAFunction []MDAFunctionSingle `json:"MDAFunction,omitempty"`
	MDAReport []MDAReport `json:"MDAReport,omitempty"`
}

// NewSubNetworkSingle4AllOf instantiates a new SubNetworkSingle4AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle4AllOf() *SubNetworkSingle4AllOf {
	this := SubNetworkSingle4AllOf{}
	return &this
}

// NewSubNetworkSingle4AllOfWithDefaults instantiates a new SubNetworkSingle4AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle4AllOfWithDefaults() *SubNetworkSingle4AllOf {
	this := SubNetworkSingle4AllOf{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle4AllOf) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingle4AllOf) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetMDAFunction returns the MDAFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetMDAFunction() []MDAFunctionSingle {
	if o == nil || isNil(o.MDAFunction) {
		var ret []MDAFunctionSingle
		return ret
	}
	return o.MDAFunction
}

// GetMDAFunctionOk returns a tuple with the MDAFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetMDAFunctionOk() ([]MDAFunctionSingle, bool) {
	if o == nil || isNil(o.MDAFunction) {
    return nil, false
	}
	return o.MDAFunction, true
}

// HasMDAFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasMDAFunction() bool {
	if o != nil && !isNil(o.MDAFunction) {
		return true
	}

	return false
}

// SetMDAFunction gets a reference to the given []MDAFunctionSingle and assigns it to the MDAFunction field.
func (o *SubNetworkSingle4AllOf) SetMDAFunction(v []MDAFunctionSingle) {
	o.MDAFunction = v
}

// GetMDAReport returns the MDAReport field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetMDAReport() []MDAReport {
	if o == nil || isNil(o.MDAReport) {
		var ret []MDAReport
		return ret
	}
	return o.MDAReport
}

// GetMDAReportOk returns a tuple with the MDAReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetMDAReportOk() ([]MDAReport, bool) {
	if o == nil || isNil(o.MDAReport) {
    return nil, false
	}
	return o.MDAReport, true
}

// HasMDAReport returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasMDAReport() bool {
	if o != nil && !isNil(o.MDAReport) {
		return true
	}

	return false
}

// SetMDAReport gets a reference to the given []MDAReport and assigns it to the MDAReport field.
func (o *SubNetworkSingle4AllOf) SetMDAReport(v []MDAReport) {
	o.MDAReport = v
}

func (o SubNetworkSingle4AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !isNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	if !isNil(o.MDAFunction) {
		toSerialize["MDAFunction"] = o.MDAFunction
	}
	if !isNil(o.MDAReport) {
		toSerialize["MDAReport"] = o.MDAReport
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle4AllOf struct {
	value *SubNetworkSingle4AllOf
	isSet bool
}

func (v NullableSubNetworkSingle4AllOf) Get() *SubNetworkSingle4AllOf {
	return v.value
}

func (v *NullableSubNetworkSingle4AllOf) Set(val *SubNetworkSingle4AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle4AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle4AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle4AllOf(val *SubNetworkSingle4AllOf) *NullableSubNetworkSingle4AllOf {
	return &NullableSubNetworkSingle4AllOf{value: val, isSet: true}
}

func (v NullableSubNetworkSingle4AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle4AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

