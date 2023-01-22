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

// UnrelatedClass struct for UnrelatedClass
type UnrelatedClass struct {
	DefaultUnrelatedClass DefaultUnrelatedClass `json:"defaultUnrelatedClass"`
	ExternalUnrelatedClass *ExternalUnrelatedClass `json:"externalUnrelatedClass,omitempty"`
	ServiceTypeUnrelatedClasses []ServiceTypeUnrelatedClass `json:"serviceTypeUnrelatedClasses,omitempty"`
}

// NewUnrelatedClass instantiates a new UnrelatedClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnrelatedClass(defaultUnrelatedClass DefaultUnrelatedClass) *UnrelatedClass {
	this := UnrelatedClass{}
	this.DefaultUnrelatedClass = defaultUnrelatedClass
	return &this
}

// NewUnrelatedClassWithDefaults instantiates a new UnrelatedClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnrelatedClassWithDefaults() *UnrelatedClass {
	this := UnrelatedClass{}
	return &this
}

// GetDefaultUnrelatedClass returns the DefaultUnrelatedClass field value
func (o *UnrelatedClass) GetDefaultUnrelatedClass() DefaultUnrelatedClass {
	if o == nil {
		var ret DefaultUnrelatedClass
		return ret
	}

	return o.DefaultUnrelatedClass
}

// GetDefaultUnrelatedClassOk returns a tuple with the DefaultUnrelatedClass field value
// and a boolean to check if the value has been set.
func (o *UnrelatedClass) GetDefaultUnrelatedClassOk() (*DefaultUnrelatedClass, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultUnrelatedClass, true
}

// SetDefaultUnrelatedClass sets field value
func (o *UnrelatedClass) SetDefaultUnrelatedClass(v DefaultUnrelatedClass) {
	o.DefaultUnrelatedClass = v
}

// GetExternalUnrelatedClass returns the ExternalUnrelatedClass field value if set, zero value otherwise.
func (o *UnrelatedClass) GetExternalUnrelatedClass() ExternalUnrelatedClass {
	if o == nil || isNil(o.ExternalUnrelatedClass) {
		var ret ExternalUnrelatedClass
		return ret
	}
	return *o.ExternalUnrelatedClass
}

// GetExternalUnrelatedClassOk returns a tuple with the ExternalUnrelatedClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnrelatedClass) GetExternalUnrelatedClassOk() (*ExternalUnrelatedClass, bool) {
	if o == nil || isNil(o.ExternalUnrelatedClass) {
    return nil, false
	}
	return o.ExternalUnrelatedClass, true
}

// HasExternalUnrelatedClass returns a boolean if a field has been set.
func (o *UnrelatedClass) HasExternalUnrelatedClass() bool {
	if o != nil && !isNil(o.ExternalUnrelatedClass) {
		return true
	}

	return false
}

// SetExternalUnrelatedClass gets a reference to the given ExternalUnrelatedClass and assigns it to the ExternalUnrelatedClass field.
func (o *UnrelatedClass) SetExternalUnrelatedClass(v ExternalUnrelatedClass) {
	o.ExternalUnrelatedClass = &v
}

// GetServiceTypeUnrelatedClasses returns the ServiceTypeUnrelatedClasses field value if set, zero value otherwise.
func (o *UnrelatedClass) GetServiceTypeUnrelatedClasses() []ServiceTypeUnrelatedClass {
	if o == nil || isNil(o.ServiceTypeUnrelatedClasses) {
		var ret []ServiceTypeUnrelatedClass
		return ret
	}
	return o.ServiceTypeUnrelatedClasses
}

// GetServiceTypeUnrelatedClassesOk returns a tuple with the ServiceTypeUnrelatedClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnrelatedClass) GetServiceTypeUnrelatedClassesOk() ([]ServiceTypeUnrelatedClass, bool) {
	if o == nil || isNil(o.ServiceTypeUnrelatedClasses) {
    return nil, false
	}
	return o.ServiceTypeUnrelatedClasses, true
}

// HasServiceTypeUnrelatedClasses returns a boolean if a field has been set.
func (o *UnrelatedClass) HasServiceTypeUnrelatedClasses() bool {
	if o != nil && !isNil(o.ServiceTypeUnrelatedClasses) {
		return true
	}

	return false
}

// SetServiceTypeUnrelatedClasses gets a reference to the given []ServiceTypeUnrelatedClass and assigns it to the ServiceTypeUnrelatedClasses field.
func (o *UnrelatedClass) SetServiceTypeUnrelatedClasses(v []ServiceTypeUnrelatedClass) {
	o.ServiceTypeUnrelatedClasses = v
}

func (o UnrelatedClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["defaultUnrelatedClass"] = o.DefaultUnrelatedClass
	}
	if !isNil(o.ExternalUnrelatedClass) {
		toSerialize["externalUnrelatedClass"] = o.ExternalUnrelatedClass
	}
	if !isNil(o.ServiceTypeUnrelatedClasses) {
		toSerialize["serviceTypeUnrelatedClasses"] = o.ServiceTypeUnrelatedClasses
	}
	return json.Marshal(toSerialize)
}

type NullableUnrelatedClass struct {
	value *UnrelatedClass
	isSet bool
}

func (v NullableUnrelatedClass) Get() *UnrelatedClass {
	return v.value
}

func (v *NullableUnrelatedClass) Set(val *UnrelatedClass) {
	v.value = val
	v.isSet = true
}

func (v NullableUnrelatedClass) IsSet() bool {
	return v.isSet
}

func (v *NullableUnrelatedClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnrelatedClass(val *UnrelatedClass) *NullableUnrelatedClass {
	return &NullableUnrelatedClass{value: val, isSet: true}
}

func (v NullableUnrelatedClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnrelatedClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


