/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// PlmnInfo struct for PlmnInfo
type PlmnInfo struct {
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewPlmnInfo instantiates a new PlmnInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnInfo() *PlmnInfo {
	this := PlmnInfo{}
	return &this
}

// NewPlmnInfoWithDefaults instantiates a new PlmnInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnInfoWithDefaults() *PlmnInfo {
	this := PlmnInfo{}
	return &this
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *PlmnInfo) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
    return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *PlmnInfo) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *PlmnInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *PlmnInfo) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *PlmnInfo) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *PlmnInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o PlmnInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnInfo struct {
	value *PlmnInfo
	isSet bool
}

func (v NullablePlmnInfo) Get() *PlmnInfo {
	return v.value
}

func (v *NullablePlmnInfo) Set(val *PlmnInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnInfo(val *PlmnInfo) *NullablePlmnInfo {
	return &NullablePlmnInfo{value: val, isSet: true}
}

func (v NullablePlmnInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


