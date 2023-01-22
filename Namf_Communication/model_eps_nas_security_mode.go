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

// EpsNasSecurityMode Indicates the EPS NAS Security Mode
type EpsNasSecurityMode struct {
	IntegrityAlgorithm EpsNasIntegrityAlgorithm `json:"integrityAlgorithm"`
	CipheringAlgorithm EpsNasCipheringAlgorithm `json:"cipheringAlgorithm"`
}

// NewEpsNasSecurityMode instantiates a new EpsNasSecurityMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpsNasSecurityMode(integrityAlgorithm EpsNasIntegrityAlgorithm, cipheringAlgorithm EpsNasCipheringAlgorithm) *EpsNasSecurityMode {
	this := EpsNasSecurityMode{}
	this.IntegrityAlgorithm = integrityAlgorithm
	this.CipheringAlgorithm = cipheringAlgorithm
	return &this
}

// NewEpsNasSecurityModeWithDefaults instantiates a new EpsNasSecurityMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpsNasSecurityModeWithDefaults() *EpsNasSecurityMode {
	this := EpsNasSecurityMode{}
	return &this
}

// GetIntegrityAlgorithm returns the IntegrityAlgorithm field value
func (o *EpsNasSecurityMode) GetIntegrityAlgorithm() EpsNasIntegrityAlgorithm {
	if o == nil {
		var ret EpsNasIntegrityAlgorithm
		return ret
	}

	return o.IntegrityAlgorithm
}

// GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field value
// and a boolean to check if the value has been set.
func (o *EpsNasSecurityMode) GetIntegrityAlgorithmOk() (*EpsNasIntegrityAlgorithm, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IntegrityAlgorithm, true
}

// SetIntegrityAlgorithm sets field value
func (o *EpsNasSecurityMode) SetIntegrityAlgorithm(v EpsNasIntegrityAlgorithm) {
	o.IntegrityAlgorithm = v
}

// GetCipheringAlgorithm returns the CipheringAlgorithm field value
func (o *EpsNasSecurityMode) GetCipheringAlgorithm() EpsNasCipheringAlgorithm {
	if o == nil {
		var ret EpsNasCipheringAlgorithm
		return ret
	}

	return o.CipheringAlgorithm
}

// GetCipheringAlgorithmOk returns a tuple with the CipheringAlgorithm field value
// and a boolean to check if the value has been set.
func (o *EpsNasSecurityMode) GetCipheringAlgorithmOk() (*EpsNasCipheringAlgorithm, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CipheringAlgorithm, true
}

// SetCipheringAlgorithm sets field value
func (o *EpsNasSecurityMode) SetCipheringAlgorithm(v EpsNasCipheringAlgorithm) {
	o.CipheringAlgorithm = v
}

func (o EpsNasSecurityMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integrityAlgorithm"] = o.IntegrityAlgorithm
	}
	if true {
		toSerialize["cipheringAlgorithm"] = o.CipheringAlgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableEpsNasSecurityMode struct {
	value *EpsNasSecurityMode
	isSet bool
}

func (v NullableEpsNasSecurityMode) Get() *EpsNasSecurityMode {
	return v.value
}

func (v *NullableEpsNasSecurityMode) Set(val *EpsNasSecurityMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsNasSecurityMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsNasSecurityMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsNasSecurityMode(val *EpsNasSecurityMode) *NullableEpsNasSecurityMode {
	return &NullableEpsNasSecurityMode{value: val, isSet: true}
}

func (v NullableEpsNasSecurityMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsNasSecurityMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


