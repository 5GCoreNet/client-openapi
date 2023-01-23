/*
3gpp-racs-parameter-provisioning

API for provisioning UE radio capability parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_RacsParameterProvisioning

import (
	"encoding/json"
)

// RacsProvisioningDataPatch Represents parameters to request the modification of a UE's radio capability data.
type RacsProvisioningDataPatch struct {
	// Identifies the configuration related to manufactuer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key.
	RacsConfigs *map[string]RacsConfigurationRm `json:"racsConfigs,omitempty"`
}

// NewRacsProvisioningDataPatch instantiates a new RacsProvisioningDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacsProvisioningDataPatch() *RacsProvisioningDataPatch {
	this := RacsProvisioningDataPatch{}
	return &this
}

// NewRacsProvisioningDataPatchWithDefaults instantiates a new RacsProvisioningDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacsProvisioningDataPatchWithDefaults() *RacsProvisioningDataPatch {
	this := RacsProvisioningDataPatch{}
	return &this
}

// GetRacsConfigs returns the RacsConfigs field value if set, zero value otherwise.
func (o *RacsProvisioningDataPatch) GetRacsConfigs() map[string]RacsConfigurationRm {
	if o == nil || isNil(o.RacsConfigs) {
		var ret map[string]RacsConfigurationRm
		return ret
	}
	return *o.RacsConfigs
}

// GetRacsConfigsOk returns a tuple with the RacsConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsProvisioningDataPatch) GetRacsConfigsOk() (*map[string]RacsConfigurationRm, bool) {
	if o == nil || isNil(o.RacsConfigs) {
    return nil, false
	}
	return o.RacsConfigs, true
}

// HasRacsConfigs returns a boolean if a field has been set.
func (o *RacsProvisioningDataPatch) HasRacsConfigs() bool {
	if o != nil && !isNil(o.RacsConfigs) {
		return true
	}

	return false
}

// SetRacsConfigs gets a reference to the given map[string]RacsConfigurationRm and assigns it to the RacsConfigs field.
func (o *RacsProvisioningDataPatch) SetRacsConfigs(v map[string]RacsConfigurationRm) {
	o.RacsConfigs = &v
}

func (o RacsProvisioningDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RacsConfigs) {
		toSerialize["racsConfigs"] = o.RacsConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableRacsProvisioningDataPatch struct {
	value *RacsProvisioningDataPatch
	isSet bool
}

func (v NullableRacsProvisioningDataPatch) Get() *RacsProvisioningDataPatch {
	return v.value
}

func (v *NullableRacsProvisioningDataPatch) Set(val *RacsProvisioningDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsProvisioningDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsProvisioningDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsProvisioningDataPatch(val *RacsProvisioningDataPatch) *NullableRacsProvisioningDataPatch {
	return &NullableRacsProvisioningDataPatch{value: val, isSet: true}
}

func (v NullableRacsProvisioningDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsProvisioningDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


