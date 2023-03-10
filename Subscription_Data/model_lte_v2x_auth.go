/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// LteV2xAuth Contains LTE V2X services authorized information.
type LteV2xAuth struct {
	VehicleUeAuth *UeAuth `json:"vehicleUeAuth,omitempty"`
	PedestrianUeAuth *UeAuth `json:"pedestrianUeAuth,omitempty"`
}

// NewLteV2xAuth instantiates a new LteV2xAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLteV2xAuth() *LteV2xAuth {
	this := LteV2xAuth{}
	return &this
}

// NewLteV2xAuthWithDefaults instantiates a new LteV2xAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLteV2xAuthWithDefaults() *LteV2xAuth {
	this := LteV2xAuth{}
	return &this
}

// GetVehicleUeAuth returns the VehicleUeAuth field value if set, zero value otherwise.
func (o *LteV2xAuth) GetVehicleUeAuth() UeAuth {
	if o == nil || isNil(o.VehicleUeAuth) {
		var ret UeAuth
		return ret
	}
	return *o.VehicleUeAuth
}

// GetVehicleUeAuthOk returns a tuple with the VehicleUeAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LteV2xAuth) GetVehicleUeAuthOk() (*UeAuth, bool) {
	if o == nil || isNil(o.VehicleUeAuth) {
    return nil, false
	}
	return o.VehicleUeAuth, true
}

// HasVehicleUeAuth returns a boolean if a field has been set.
func (o *LteV2xAuth) HasVehicleUeAuth() bool {
	if o != nil && !isNil(o.VehicleUeAuth) {
		return true
	}

	return false
}

// SetVehicleUeAuth gets a reference to the given UeAuth and assigns it to the VehicleUeAuth field.
func (o *LteV2xAuth) SetVehicleUeAuth(v UeAuth) {
	o.VehicleUeAuth = &v
}

// GetPedestrianUeAuth returns the PedestrianUeAuth field value if set, zero value otherwise.
func (o *LteV2xAuth) GetPedestrianUeAuth() UeAuth {
	if o == nil || isNil(o.PedestrianUeAuth) {
		var ret UeAuth
		return ret
	}
	return *o.PedestrianUeAuth
}

// GetPedestrianUeAuthOk returns a tuple with the PedestrianUeAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LteV2xAuth) GetPedestrianUeAuthOk() (*UeAuth, bool) {
	if o == nil || isNil(o.PedestrianUeAuth) {
    return nil, false
	}
	return o.PedestrianUeAuth, true
}

// HasPedestrianUeAuth returns a boolean if a field has been set.
func (o *LteV2xAuth) HasPedestrianUeAuth() bool {
	if o != nil && !isNil(o.PedestrianUeAuth) {
		return true
	}

	return false
}

// SetPedestrianUeAuth gets a reference to the given UeAuth and assigns it to the PedestrianUeAuth field.
func (o *LteV2xAuth) SetPedestrianUeAuth(v UeAuth) {
	o.PedestrianUeAuth = &v
}

func (o LteV2xAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VehicleUeAuth) {
		toSerialize["vehicleUeAuth"] = o.VehicleUeAuth
	}
	if !isNil(o.PedestrianUeAuth) {
		toSerialize["pedestrianUeAuth"] = o.PedestrianUeAuth
	}
	return json.Marshal(toSerialize)
}

type NullableLteV2xAuth struct {
	value *LteV2xAuth
	isSet bool
}

func (v NullableLteV2xAuth) Get() *LteV2xAuth {
	return v.value
}

func (v *NullableLteV2xAuth) Set(val *LteV2xAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableLteV2xAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableLteV2xAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLteV2xAuth(val *LteV2xAuth) *NullableLteV2xAuth {
	return &NullableLteV2xAuth{value: val, isSet: true}
}

func (v NullableLteV2xAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLteV2xAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


