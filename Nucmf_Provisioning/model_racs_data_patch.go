/*
Nucmf_Provisioning

UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_Provisioning

import (
	"encoding/json"
)

// RacsDataPatch Represents a modification of a UE radio capability data provided by the NF service consumer. 
type RacsDataPatch struct {
	// Identifies the configuration related to manufacturer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key. 
	RacsConfigs *map[string]RacsConfigurationRm `json:"racsConfigs,omitempty"`
}

// NewRacsDataPatch instantiates a new RacsDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacsDataPatch() *RacsDataPatch {
	this := RacsDataPatch{}
	return &this
}

// NewRacsDataPatchWithDefaults instantiates a new RacsDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacsDataPatchWithDefaults() *RacsDataPatch {
	this := RacsDataPatch{}
	return &this
}

// GetRacsConfigs returns the RacsConfigs field value if set, zero value otherwise.
func (o *RacsDataPatch) GetRacsConfigs() map[string]RacsConfigurationRm {
	if o == nil || isNil(o.RacsConfigs) {
		var ret map[string]RacsConfigurationRm
		return ret
	}
	return *o.RacsConfigs
}

// GetRacsConfigsOk returns a tuple with the RacsConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsDataPatch) GetRacsConfigsOk() (*map[string]RacsConfigurationRm, bool) {
	if o == nil || isNil(o.RacsConfigs) {
    return nil, false
	}
	return o.RacsConfigs, true
}

// HasRacsConfigs returns a boolean if a field has been set.
func (o *RacsDataPatch) HasRacsConfigs() bool {
	if o != nil && !isNil(o.RacsConfigs) {
		return true
	}

	return false
}

// SetRacsConfigs gets a reference to the given map[string]RacsConfigurationRm and assigns it to the RacsConfigs field.
func (o *RacsDataPatch) SetRacsConfigs(v map[string]RacsConfigurationRm) {
	o.RacsConfigs = &v
}

func (o RacsDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RacsConfigs) {
		toSerialize["racsConfigs"] = o.RacsConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableRacsDataPatch struct {
	value *RacsDataPatch
	isSet bool
}

func (v NullableRacsDataPatch) Get() *RacsDataPatch {
	return v.value
}

func (v *NullableRacsDataPatch) Set(val *RacsDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsDataPatch(val *RacsDataPatch) *NullableRacsDataPatch {
	return &NullableRacsDataPatch{value: val, isSet: true}
}

func (v NullableRacsDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


