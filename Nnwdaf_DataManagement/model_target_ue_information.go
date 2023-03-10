/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// TargetUeInformation Identifies the target UE information.
type TargetUeInformation struct {
	AnyUe *bool `json:"anyUe,omitempty"`
	Supis []string `json:"supis,omitempty"`
	Gpsis []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
}

// NewTargetUeInformation instantiates a new TargetUeInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetUeInformation() *TargetUeInformation {
	this := TargetUeInformation{}
	return &this
}

// NewTargetUeInformationWithDefaults instantiates a new TargetUeInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetUeInformationWithDefaults() *TargetUeInformation {
	this := TargetUeInformation{}
	return &this
}

// GetAnyUe returns the AnyUe field value if set, zero value otherwise.
func (o *TargetUeInformation) GetAnyUe() bool {
	if o == nil || isNil(o.AnyUe) {
		var ret bool
		return ret
	}
	return *o.AnyUe
}

// GetAnyUeOk returns a tuple with the AnyUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeInformation) GetAnyUeOk() (*bool, bool) {
	if o == nil || isNil(o.AnyUe) {
    return nil, false
	}
	return o.AnyUe, true
}

// HasAnyUe returns a boolean if a field has been set.
func (o *TargetUeInformation) HasAnyUe() bool {
	if o != nil && !isNil(o.AnyUe) {
		return true
	}

	return false
}

// SetAnyUe gets a reference to the given bool and assigns it to the AnyUe field.
func (o *TargetUeInformation) SetAnyUe(v bool) {
	o.AnyUe = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *TargetUeInformation) GetSupis() []string {
	if o == nil || isNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeInformation) GetSupisOk() ([]string, bool) {
	if o == nil || isNil(o.Supis) {
    return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *TargetUeInformation) HasSupis() bool {
	if o != nil && !isNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *TargetUeInformation) SetSupis(v []string) {
	o.Supis = v
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *TargetUeInformation) GetGpsis() []string {
	if o == nil || isNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeInformation) GetGpsisOk() ([]string, bool) {
	if o == nil || isNil(o.Gpsis) {
    return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *TargetUeInformation) HasGpsis() bool {
	if o != nil && !isNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *TargetUeInformation) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetIntGroupIds returns the IntGroupIds field value if set, zero value otherwise.
func (o *TargetUeInformation) GetIntGroupIds() []string {
	if o == nil || isNil(o.IntGroupIds) {
		var ret []string
		return ret
	}
	return o.IntGroupIds
}

// GetIntGroupIdsOk returns a tuple with the IntGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeInformation) GetIntGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.IntGroupIds) {
    return nil, false
	}
	return o.IntGroupIds, true
}

// HasIntGroupIds returns a boolean if a field has been set.
func (o *TargetUeInformation) HasIntGroupIds() bool {
	if o != nil && !isNil(o.IntGroupIds) {
		return true
	}

	return false
}

// SetIntGroupIds gets a reference to the given []string and assigns it to the IntGroupIds field.
func (o *TargetUeInformation) SetIntGroupIds(v []string) {
	o.IntGroupIds = v
}

func (o TargetUeInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnyUe) {
		toSerialize["anyUe"] = o.AnyUe
	}
	if !isNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !isNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !isNil(o.IntGroupIds) {
		toSerialize["intGroupIds"] = o.IntGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableTargetUeInformation struct {
	value *TargetUeInformation
	isSet bool
}

func (v NullableTargetUeInformation) Get() *TargetUeInformation {
	return v.value
}

func (v *NullableTargetUeInformation) Set(val *TargetUeInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetUeInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetUeInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetUeInformation(val *TargetUeInformation) *NullableTargetUeInformation {
	return &NullableTargetUeInformation{value: val, isSet: true}
}

func (v NullableTargetUeInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetUeInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


