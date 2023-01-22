/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// ServiceAreaId Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	// Location Area Code.
	Lac string `json:"lac"`
	// Service Area Code.
	Sac string `json:"sac"`
}

// NewServiceAreaId instantiates a new ServiceAreaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAreaId(plmnId PlmnId, lac string, sac string) *ServiceAreaId {
	this := ServiceAreaId{}
	this.PlmnId = plmnId
	this.Lac = lac
	this.Sac = sac
	return &this
}

// NewServiceAreaIdWithDefaults instantiates a new ServiceAreaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaIdWithDefaults() *ServiceAreaId {
	this := ServiceAreaId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *ServiceAreaId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *ServiceAreaId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *ServiceAreaId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetLacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *ServiceAreaId) SetLac(v string) {
	o.Lac = v
}

// GetSac returns the Sac field value
func (o *ServiceAreaId) GetSac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sac
}

// GetSacOk returns a tuple with the Sac field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetSacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sac, true
}

// SetSac sets field value
func (o *ServiceAreaId) SetSac(v string) {
	o.Sac = v
}

func (o ServiceAreaId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["lac"] = o.Lac
	}
	if true {
		toSerialize["sac"] = o.Sac
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAreaId struct {
	value *ServiceAreaId
	isSet bool
}

func (v NullableServiceAreaId) Get() *ServiceAreaId {
	return v.value
}

func (v *NullableServiceAreaId) Set(val *ServiceAreaId) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAreaId) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAreaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAreaId(val *ServiceAreaId) *NullableServiceAreaId {
	return &NullableServiceAreaId{value: val, isSet: true}
}

func (v NullableServiceAreaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAreaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


