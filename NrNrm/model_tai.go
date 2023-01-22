/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// Tai struct for Tai
type Tai struct {
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	NrTac *int32 `json:"nrTac,omitempty"`
}

// NewTai instantiates a new Tai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTai() *Tai {
	this := Tai{}
	return &this
}

// NewTaiWithDefaults instantiates a new Tai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiWithDefaults() *Tai {
	this := Tai{}
	return &this
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *Tai) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
    return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *Tai) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *Tai) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetNrTac returns the NrTac field value if set, zero value otherwise.
func (o *Tai) GetNrTac() int32 {
	if o == nil || isNil(o.NrTac) {
		var ret int32
		return ret
	}
	return *o.NrTac
}

// GetNrTacOk returns a tuple with the NrTac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetNrTacOk() (*int32, bool) {
	if o == nil || isNil(o.NrTac) {
    return nil, false
	}
	return o.NrTac, true
}

// HasNrTac returns a boolean if a field has been set.
func (o *Tai) HasNrTac() bool {
	if o != nil && !isNil(o.NrTac) {
		return true
	}

	return false
}

// SetNrTac gets a reference to the given int32 and assigns it to the NrTac field.
func (o *Tai) SetNrTac(v int32) {
	o.NrTac = &v
}

func (o Tai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.NrTac) {
		toSerialize["nrTac"] = o.NrTac
	}
	return json.Marshal(toSerialize)
}

type NullableTai struct {
	value *Tai
	isSet bool
}

func (v NullableTai) Get() *Tai {
	return v.value
}

func (v *NullableTai) Set(val *Tai) {
	v.value = val
	v.isSet = true
}

func (v NullableTai) IsSet() bool {
	return v.isSet
}

func (v *NullableTai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTai(val *Tai) *NullableTai {
	return &NullableTai{value: val, isSet: true}
}

func (v NullableTai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


