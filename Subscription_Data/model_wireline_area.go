/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// WirelineArea One and only one of the \"globLineIds\", \"hfcNIds\", \"areaCodeB\" and \"areaCodeC\" attributes shall be included in a WirelineArea data structure 
type WirelineArea struct {
	GlobalLineIds []string `json:"globalLineIds,omitempty"`
	HfcNIds []string `json:"hfcNIds,omitempty"`
	// Values are operator specific.
	AreaCodeB *string `json:"areaCodeB,omitempty"`
	// Values are operator specific.
	AreaCodeC *string `json:"areaCodeC,omitempty"`
}

// NewWirelineArea instantiates a new WirelineArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelineArea() *WirelineArea {
	this := WirelineArea{}
	return &this
}

// NewWirelineAreaWithDefaults instantiates a new WirelineArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelineAreaWithDefaults() *WirelineArea {
	this := WirelineArea{}
	return &this
}

// GetGlobalLineIds returns the GlobalLineIds field value if set, zero value otherwise.
func (o *WirelineArea) GetGlobalLineIds() []string {
	if o == nil || isNil(o.GlobalLineIds) {
		var ret []string
		return ret
	}
	return o.GlobalLineIds
}

// GetGlobalLineIdsOk returns a tuple with the GlobalLineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineArea) GetGlobalLineIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GlobalLineIds) {
    return nil, false
	}
	return o.GlobalLineIds, true
}

// HasGlobalLineIds returns a boolean if a field has been set.
func (o *WirelineArea) HasGlobalLineIds() bool {
	if o != nil && !isNil(o.GlobalLineIds) {
		return true
	}

	return false
}

// SetGlobalLineIds gets a reference to the given []string and assigns it to the GlobalLineIds field.
func (o *WirelineArea) SetGlobalLineIds(v []string) {
	o.GlobalLineIds = v
}

// GetHfcNIds returns the HfcNIds field value if set, zero value otherwise.
func (o *WirelineArea) GetHfcNIds() []string {
	if o == nil || isNil(o.HfcNIds) {
		var ret []string
		return ret
	}
	return o.HfcNIds
}

// GetHfcNIdsOk returns a tuple with the HfcNIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineArea) GetHfcNIdsOk() ([]string, bool) {
	if o == nil || isNil(o.HfcNIds) {
    return nil, false
	}
	return o.HfcNIds, true
}

// HasHfcNIds returns a boolean if a field has been set.
func (o *WirelineArea) HasHfcNIds() bool {
	if o != nil && !isNil(o.HfcNIds) {
		return true
	}

	return false
}

// SetHfcNIds gets a reference to the given []string and assigns it to the HfcNIds field.
func (o *WirelineArea) SetHfcNIds(v []string) {
	o.HfcNIds = v
}

// GetAreaCodeB returns the AreaCodeB field value if set, zero value otherwise.
func (o *WirelineArea) GetAreaCodeB() string {
	if o == nil || isNil(o.AreaCodeB) {
		var ret string
		return ret
	}
	return *o.AreaCodeB
}

// GetAreaCodeBOk returns a tuple with the AreaCodeB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineArea) GetAreaCodeBOk() (*string, bool) {
	if o == nil || isNil(o.AreaCodeB) {
    return nil, false
	}
	return o.AreaCodeB, true
}

// HasAreaCodeB returns a boolean if a field has been set.
func (o *WirelineArea) HasAreaCodeB() bool {
	if o != nil && !isNil(o.AreaCodeB) {
		return true
	}

	return false
}

// SetAreaCodeB gets a reference to the given string and assigns it to the AreaCodeB field.
func (o *WirelineArea) SetAreaCodeB(v string) {
	o.AreaCodeB = &v
}

// GetAreaCodeC returns the AreaCodeC field value if set, zero value otherwise.
func (o *WirelineArea) GetAreaCodeC() string {
	if o == nil || isNil(o.AreaCodeC) {
		var ret string
		return ret
	}
	return *o.AreaCodeC
}

// GetAreaCodeCOk returns a tuple with the AreaCodeC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineArea) GetAreaCodeCOk() (*string, bool) {
	if o == nil || isNil(o.AreaCodeC) {
    return nil, false
	}
	return o.AreaCodeC, true
}

// HasAreaCodeC returns a boolean if a field has been set.
func (o *WirelineArea) HasAreaCodeC() bool {
	if o != nil && !isNil(o.AreaCodeC) {
		return true
	}

	return false
}

// SetAreaCodeC gets a reference to the given string and assigns it to the AreaCodeC field.
func (o *WirelineArea) SetAreaCodeC(v string) {
	o.AreaCodeC = &v
}

func (o WirelineArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GlobalLineIds) {
		toSerialize["globalLineIds"] = o.GlobalLineIds
	}
	if !isNil(o.HfcNIds) {
		toSerialize["hfcNIds"] = o.HfcNIds
	}
	if !isNil(o.AreaCodeB) {
		toSerialize["areaCodeB"] = o.AreaCodeB
	}
	if !isNil(o.AreaCodeC) {
		toSerialize["areaCodeC"] = o.AreaCodeC
	}
	return json.Marshal(toSerialize)
}

type NullableWirelineArea struct {
	value *WirelineArea
	isSet bool
}

func (v NullableWirelineArea) Get() *WirelineArea {
	return v.value
}

func (v *NullableWirelineArea) Set(val *WirelineArea) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelineArea) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelineArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelineArea(val *WirelineArea) *NullableWirelineArea {
	return &NullableWirelineArea{value: val, isSet: true}
}

func (v NullableWirelineArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelineArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

