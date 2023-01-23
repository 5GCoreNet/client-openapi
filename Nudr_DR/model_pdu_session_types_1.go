/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// PduSessionTypes1 struct for PduSessionTypes1
type PduSessionTypes1 struct {
	DefaultSessionType *PduSessionType `json:"defaultSessionType,omitempty"`
	AllowedSessionTypes []PduSessionType `json:"allowedSessionTypes,omitempty"`
}

// NewPduSessionTypes1 instantiates a new PduSessionTypes1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionTypes1() *PduSessionTypes1 {
	this := PduSessionTypes1{}
	return &this
}

// NewPduSessionTypes1WithDefaults instantiates a new PduSessionTypes1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionTypes1WithDefaults() *PduSessionTypes1 {
	this := PduSessionTypes1{}
	return &this
}

// GetDefaultSessionType returns the DefaultSessionType field value if set, zero value otherwise.
func (o *PduSessionTypes1) GetDefaultSessionType() PduSessionType {
	if o == nil || isNil(o.DefaultSessionType) {
		var ret PduSessionType
		return ret
	}
	return *o.DefaultSessionType
}

// GetDefaultSessionTypeOk returns a tuple with the DefaultSessionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes1) GetDefaultSessionTypeOk() (*PduSessionType, bool) {
	if o == nil || isNil(o.DefaultSessionType) {
    return nil, false
	}
	return o.DefaultSessionType, true
}

// HasDefaultSessionType returns a boolean if a field has been set.
func (o *PduSessionTypes1) HasDefaultSessionType() bool {
	if o != nil && !isNil(o.DefaultSessionType) {
		return true
	}

	return false
}

// SetDefaultSessionType gets a reference to the given PduSessionType and assigns it to the DefaultSessionType field.
func (o *PduSessionTypes1) SetDefaultSessionType(v PduSessionType) {
	o.DefaultSessionType = &v
}

// GetAllowedSessionTypes returns the AllowedSessionTypes field value if set, zero value otherwise.
func (o *PduSessionTypes1) GetAllowedSessionTypes() []PduSessionType {
	if o == nil || isNil(o.AllowedSessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.AllowedSessionTypes
}

// GetAllowedSessionTypesOk returns a tuple with the AllowedSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes1) GetAllowedSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || isNil(o.AllowedSessionTypes) {
    return nil, false
	}
	return o.AllowedSessionTypes, true
}

// HasAllowedSessionTypes returns a boolean if a field has been set.
func (o *PduSessionTypes1) HasAllowedSessionTypes() bool {
	if o != nil && !isNil(o.AllowedSessionTypes) {
		return true
	}

	return false
}

// SetAllowedSessionTypes gets a reference to the given []PduSessionType and assigns it to the AllowedSessionTypes field.
func (o *PduSessionTypes1) SetAllowedSessionTypes(v []PduSessionType) {
	o.AllowedSessionTypes = v
}

func (o PduSessionTypes1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSessionType) {
		toSerialize["defaultSessionType"] = o.DefaultSessionType
	}
	if !isNil(o.AllowedSessionTypes) {
		toSerialize["allowedSessionTypes"] = o.AllowedSessionTypes
	}
	return json.Marshal(toSerialize)
}

type NullablePduSessionTypes1 struct {
	value *PduSessionTypes1
	isSet bool
}

func (v NullablePduSessionTypes1) Get() *PduSessionTypes1 {
	return v.value
}

func (v *NullablePduSessionTypes1) Set(val *PduSessionTypes1) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionTypes1) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionTypes1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionTypes1(val *PduSessionTypes1) *NullablePduSessionTypes1 {
	return &NullablePduSessionTypes1{value: val, isSet: true}
}

func (v NullablePduSessionTypes1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionTypes1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


