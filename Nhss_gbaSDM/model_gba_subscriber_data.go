/*
Nhss_gbaSDM

Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaSDM

import (
	"encoding/json"
)

// GbaSubscriberData GBA subscriber data of the UE; it includes the GBA User Security Settings (GUSS) data 
type GbaSubscriberData struct {
	Guss *Guss `json:"guss,omitempty"`
}

// NewGbaSubscriberData instantiates a new GbaSubscriberData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGbaSubscriberData() *GbaSubscriberData {
	this := GbaSubscriberData{}
	return &this
}

// NewGbaSubscriberDataWithDefaults instantiates a new GbaSubscriberData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGbaSubscriberDataWithDefaults() *GbaSubscriberData {
	this := GbaSubscriberData{}
	return &this
}

// GetGuss returns the Guss field value if set, zero value otherwise.
func (o *GbaSubscriberData) GetGuss() Guss {
	if o == nil || isNil(o.Guss) {
		var ret Guss
		return ret
	}
	return *o.Guss
}

// GetGussOk returns a tuple with the Guss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GbaSubscriberData) GetGussOk() (*Guss, bool) {
	if o == nil || isNil(o.Guss) {
    return nil, false
	}
	return o.Guss, true
}

// HasGuss returns a boolean if a field has been set.
func (o *GbaSubscriberData) HasGuss() bool {
	if o != nil && !isNil(o.Guss) {
		return true
	}

	return false
}

// SetGuss gets a reference to the given Guss and assigns it to the Guss field.
func (o *GbaSubscriberData) SetGuss(v Guss) {
	o.Guss = &v
}

func (o GbaSubscriberData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Guss) {
		toSerialize["guss"] = o.Guss
	}
	return json.Marshal(toSerialize)
}

type NullableGbaSubscriberData struct {
	value *GbaSubscriberData
	isSet bool
}

func (v NullableGbaSubscriberData) Get() *GbaSubscriberData {
	return v.value
}

func (v *NullableGbaSubscriberData) Set(val *GbaSubscriberData) {
	v.value = val
	v.isSet = true
}

func (v NullableGbaSubscriberData) IsSet() bool {
	return v.isSet
}

func (v *NullableGbaSubscriberData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbaSubscriberData(val *GbaSubscriberData) *NullableGbaSubscriberData {
	return &NullableGbaSubscriberData{value: val, isSet: true}
}

func (v NullableGbaSubscriberData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbaSubscriberData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


