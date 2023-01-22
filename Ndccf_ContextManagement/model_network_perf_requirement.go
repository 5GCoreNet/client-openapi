/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
)

// NetworkPerfRequirement Represents a network performance requirement.
type NetworkPerfRequirement struct {
	NwPerfType NetworkPerfType `json:"nwPerfType"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	RelativeRatio *int32 `json:"relativeRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum *int32 `json:"absoluteNum,omitempty"`
}

// NewNetworkPerfRequirement instantiates a new NetworkPerfRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPerfRequirement(nwPerfType NetworkPerfType) *NetworkPerfRequirement {
	this := NetworkPerfRequirement{}
	this.NwPerfType = nwPerfType
	return &this
}

// NewNetworkPerfRequirementWithDefaults instantiates a new NetworkPerfRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPerfRequirementWithDefaults() *NetworkPerfRequirement {
	this := NetworkPerfRequirement{}
	return &this
}

// GetNwPerfType returns the NwPerfType field value
func (o *NetworkPerfRequirement) GetNwPerfType() NetworkPerfType {
	if o == nil {
		var ret NetworkPerfType
		return ret
	}

	return o.NwPerfType
}

// GetNwPerfTypeOk returns a tuple with the NwPerfType field value
// and a boolean to check if the value has been set.
func (o *NetworkPerfRequirement) GetNwPerfTypeOk() (*NetworkPerfType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NwPerfType, true
}

// SetNwPerfType sets field value
func (o *NetworkPerfRequirement) SetNwPerfType(v NetworkPerfType) {
	o.NwPerfType = v
}

// GetRelativeRatio returns the RelativeRatio field value if set, zero value otherwise.
func (o *NetworkPerfRequirement) GetRelativeRatio() int32 {
	if o == nil || isNil(o.RelativeRatio) {
		var ret int32
		return ret
	}
	return *o.RelativeRatio
}

// GetRelativeRatioOk returns a tuple with the RelativeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPerfRequirement) GetRelativeRatioOk() (*int32, bool) {
	if o == nil || isNil(o.RelativeRatio) {
    return nil, false
	}
	return o.RelativeRatio, true
}

// HasRelativeRatio returns a boolean if a field has been set.
func (o *NetworkPerfRequirement) HasRelativeRatio() bool {
	if o != nil && !isNil(o.RelativeRatio) {
		return true
	}

	return false
}

// SetRelativeRatio gets a reference to the given int32 and assigns it to the RelativeRatio field.
func (o *NetworkPerfRequirement) SetRelativeRatio(v int32) {
	o.RelativeRatio = &v
}

// GetAbsoluteNum returns the AbsoluteNum field value if set, zero value otherwise.
func (o *NetworkPerfRequirement) GetAbsoluteNum() int32 {
	if o == nil || isNil(o.AbsoluteNum) {
		var ret int32
		return ret
	}
	return *o.AbsoluteNum
}

// GetAbsoluteNumOk returns a tuple with the AbsoluteNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPerfRequirement) GetAbsoluteNumOk() (*int32, bool) {
	if o == nil || isNil(o.AbsoluteNum) {
    return nil, false
	}
	return o.AbsoluteNum, true
}

// HasAbsoluteNum returns a boolean if a field has been set.
func (o *NetworkPerfRequirement) HasAbsoluteNum() bool {
	if o != nil && !isNil(o.AbsoluteNum) {
		return true
	}

	return false
}

// SetAbsoluteNum gets a reference to the given int32 and assigns it to the AbsoluteNum field.
func (o *NetworkPerfRequirement) SetAbsoluteNum(v int32) {
	o.AbsoluteNum = &v
}

func (o NetworkPerfRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nwPerfType"] = o.NwPerfType
	}
	if !isNil(o.RelativeRatio) {
		toSerialize["relativeRatio"] = o.RelativeRatio
	}
	if !isNil(o.AbsoluteNum) {
		toSerialize["absoluteNum"] = o.AbsoluteNum
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPerfRequirement struct {
	value *NetworkPerfRequirement
	isSet bool
}

func (v NullableNetworkPerfRequirement) Get() *NetworkPerfRequirement {
	return v.value
}

func (v *NullableNetworkPerfRequirement) Set(val *NetworkPerfRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPerfRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPerfRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPerfRequirement(val *NetworkPerfRequirement) *NullableNetworkPerfRequirement {
	return &NullableNetworkPerfRequirement{value: val, isSet: true}
}

func (v NullableNetworkPerfRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPerfRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


