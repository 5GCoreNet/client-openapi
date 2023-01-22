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

// PlmnOperatorClass struct for PlmnOperatorClass
type PlmnOperatorClass struct {
	LcsClientClass LcsClientClass `json:"lcsClientClass"`
	LcsClientIds []string `json:"lcsClientIds"`
}

// NewPlmnOperatorClass instantiates a new PlmnOperatorClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnOperatorClass(lcsClientClass LcsClientClass, lcsClientIds []string) *PlmnOperatorClass {
	this := PlmnOperatorClass{}
	this.LcsClientClass = lcsClientClass
	this.LcsClientIds = lcsClientIds
	return &this
}

// NewPlmnOperatorClassWithDefaults instantiates a new PlmnOperatorClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnOperatorClassWithDefaults() *PlmnOperatorClass {
	this := PlmnOperatorClass{}
	return &this
}

// GetLcsClientClass returns the LcsClientClass field value
func (o *PlmnOperatorClass) GetLcsClientClass() LcsClientClass {
	if o == nil {
		var ret LcsClientClass
		return ret
	}

	return o.LcsClientClass
}

// GetLcsClientClassOk returns a tuple with the LcsClientClass field value
// and a boolean to check if the value has been set.
func (o *PlmnOperatorClass) GetLcsClientClassOk() (*LcsClientClass, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LcsClientClass, true
}

// SetLcsClientClass sets field value
func (o *PlmnOperatorClass) SetLcsClientClass(v LcsClientClass) {
	o.LcsClientClass = v
}

// GetLcsClientIds returns the LcsClientIds field value
func (o *PlmnOperatorClass) GetLcsClientIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LcsClientIds
}

// GetLcsClientIdsOk returns a tuple with the LcsClientIds field value
// and a boolean to check if the value has been set.
func (o *PlmnOperatorClass) GetLcsClientIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LcsClientIds, true
}

// SetLcsClientIds sets field value
func (o *PlmnOperatorClass) SetLcsClientIds(v []string) {
	o.LcsClientIds = v
}

func (o PlmnOperatorClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lcsClientClass"] = o.LcsClientClass
	}
	if true {
		toSerialize["lcsClientIds"] = o.LcsClientIds
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnOperatorClass struct {
	value *PlmnOperatorClass
	isSet bool
}

func (v NullablePlmnOperatorClass) Get() *PlmnOperatorClass {
	return v.value
}

func (v *NullablePlmnOperatorClass) Set(val *PlmnOperatorClass) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnOperatorClass) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnOperatorClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnOperatorClass(val *PlmnOperatorClass) *NullablePlmnOperatorClass {
	return &NullablePlmnOperatorClass{value: val, isSet: true}
}

func (v NullablePlmnOperatorClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnOperatorClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


