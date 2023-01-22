/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// NrV2xAuth Contains NR V2X services authorized information.
type NrV2xAuth struct {
	VehicleUeAuth *UeAuth `json:"vehicleUeAuth,omitempty"`
	PedestrianUeAuth *UeAuth `json:"pedestrianUeAuth,omitempty"`
}

// NewNrV2xAuth instantiates a new NrV2xAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrV2xAuth() *NrV2xAuth {
	this := NrV2xAuth{}
	return &this
}

// NewNrV2xAuthWithDefaults instantiates a new NrV2xAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrV2xAuthWithDefaults() *NrV2xAuth {
	this := NrV2xAuth{}
	return &this
}

// GetVehicleUeAuth returns the VehicleUeAuth field value if set, zero value otherwise.
func (o *NrV2xAuth) GetVehicleUeAuth() UeAuth {
	if o == nil || isNil(o.VehicleUeAuth) {
		var ret UeAuth
		return ret
	}
	return *o.VehicleUeAuth
}

// GetVehicleUeAuthOk returns a tuple with the VehicleUeAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrV2xAuth) GetVehicleUeAuthOk() (*UeAuth, bool) {
	if o == nil || isNil(o.VehicleUeAuth) {
    return nil, false
	}
	return o.VehicleUeAuth, true
}

// HasVehicleUeAuth returns a boolean if a field has been set.
func (o *NrV2xAuth) HasVehicleUeAuth() bool {
	if o != nil && !isNil(o.VehicleUeAuth) {
		return true
	}

	return false
}

// SetVehicleUeAuth gets a reference to the given UeAuth and assigns it to the VehicleUeAuth field.
func (o *NrV2xAuth) SetVehicleUeAuth(v UeAuth) {
	o.VehicleUeAuth = &v
}

// GetPedestrianUeAuth returns the PedestrianUeAuth field value if set, zero value otherwise.
func (o *NrV2xAuth) GetPedestrianUeAuth() UeAuth {
	if o == nil || isNil(o.PedestrianUeAuth) {
		var ret UeAuth
		return ret
	}
	return *o.PedestrianUeAuth
}

// GetPedestrianUeAuthOk returns a tuple with the PedestrianUeAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrV2xAuth) GetPedestrianUeAuthOk() (*UeAuth, bool) {
	if o == nil || isNil(o.PedestrianUeAuth) {
    return nil, false
	}
	return o.PedestrianUeAuth, true
}

// HasPedestrianUeAuth returns a boolean if a field has been set.
func (o *NrV2xAuth) HasPedestrianUeAuth() bool {
	if o != nil && !isNil(o.PedestrianUeAuth) {
		return true
	}

	return false
}

// SetPedestrianUeAuth gets a reference to the given UeAuth and assigns it to the PedestrianUeAuth field.
func (o *NrV2xAuth) SetPedestrianUeAuth(v UeAuth) {
	o.PedestrianUeAuth = &v
}

func (o NrV2xAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VehicleUeAuth) {
		toSerialize["vehicleUeAuth"] = o.VehicleUeAuth
	}
	if !isNil(o.PedestrianUeAuth) {
		toSerialize["pedestrianUeAuth"] = o.PedestrianUeAuth
	}
	return json.Marshal(toSerialize)
}

type NullableNrV2xAuth struct {
	value *NrV2xAuth
	isSet bool
}

func (v NullableNrV2xAuth) Get() *NrV2xAuth {
	return v.value
}

func (v *NullableNrV2xAuth) Set(val *NrV2xAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableNrV2xAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableNrV2xAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrV2xAuth(val *NrV2xAuth) *NullableNrV2xAuth {
	return &NullableNrV2xAuth{value: val, isSet: true}
}

func (v NullableNrV2xAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrV2xAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


