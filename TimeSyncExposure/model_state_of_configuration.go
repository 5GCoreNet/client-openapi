/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// StateOfConfiguration Contains the state of the time synchronization configuration.
type StateOfConfiguration struct {
	// When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for NW-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for NW-TT port is inactive. Default value is false. 
	StateOfNwtt *bool `json:"stateOfNwtt,omitempty"`
	// Contains the PTP port states of the DS-TT(s). 
	StateOfDstts []StateOfDstt `json:"stateOfDstts,omitempty"`
}

// NewStateOfConfiguration instantiates a new StateOfConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateOfConfiguration() *StateOfConfiguration {
	this := StateOfConfiguration{}
	return &this
}

// NewStateOfConfigurationWithDefaults instantiates a new StateOfConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateOfConfigurationWithDefaults() *StateOfConfiguration {
	this := StateOfConfiguration{}
	return &this
}

// GetStateOfNwtt returns the StateOfNwtt field value if set, zero value otherwise.
func (o *StateOfConfiguration) GetStateOfNwtt() bool {
	if o == nil || isNil(o.StateOfNwtt) {
		var ret bool
		return ret
	}
	return *o.StateOfNwtt
}

// GetStateOfNwttOk returns a tuple with the StateOfNwtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateOfConfiguration) GetStateOfNwttOk() (*bool, bool) {
	if o == nil || isNil(o.StateOfNwtt) {
    return nil, false
	}
	return o.StateOfNwtt, true
}

// HasStateOfNwtt returns a boolean if a field has been set.
func (o *StateOfConfiguration) HasStateOfNwtt() bool {
	if o != nil && !isNil(o.StateOfNwtt) {
		return true
	}

	return false
}

// SetStateOfNwtt gets a reference to the given bool and assigns it to the StateOfNwtt field.
func (o *StateOfConfiguration) SetStateOfNwtt(v bool) {
	o.StateOfNwtt = &v
}

// GetStateOfDstts returns the StateOfDstts field value if set, zero value otherwise.
func (o *StateOfConfiguration) GetStateOfDstts() []StateOfDstt {
	if o == nil || isNil(o.StateOfDstts) {
		var ret []StateOfDstt
		return ret
	}
	return o.StateOfDstts
}

// GetStateOfDsttsOk returns a tuple with the StateOfDstts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateOfConfiguration) GetStateOfDsttsOk() ([]StateOfDstt, bool) {
	if o == nil || isNil(o.StateOfDstts) {
    return nil, false
	}
	return o.StateOfDstts, true
}

// HasStateOfDstts returns a boolean if a field has been set.
func (o *StateOfConfiguration) HasStateOfDstts() bool {
	if o != nil && !isNil(o.StateOfDstts) {
		return true
	}

	return false
}

// SetStateOfDstts gets a reference to the given []StateOfDstt and assigns it to the StateOfDstts field.
func (o *StateOfConfiguration) SetStateOfDstts(v []StateOfDstt) {
	o.StateOfDstts = v
}

func (o StateOfConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StateOfNwtt) {
		toSerialize["stateOfNwtt"] = o.StateOfNwtt
	}
	if !isNil(o.StateOfDstts) {
		toSerialize["stateOfDstts"] = o.StateOfDstts
	}
	return json.Marshal(toSerialize)
}

type NullableStateOfConfiguration struct {
	value *StateOfConfiguration
	isSet bool
}

func (v NullableStateOfConfiguration) Get() *StateOfConfiguration {
	return v.value
}

func (v *NullableStateOfConfiguration) Set(val *StateOfConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStateOfConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStateOfConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateOfConfiguration(val *StateOfConfiguration) *NullableStateOfConfiguration {
	return &NullableStateOfConfiguration{value: val, isSet: true}
}

func (v NullableStateOfConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateOfConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


