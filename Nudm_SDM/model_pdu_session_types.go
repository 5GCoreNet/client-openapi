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

// PduSessionTypes struct for PduSessionTypes
type PduSessionTypes struct {
	DefaultSessionType *PduSessionType `json:"defaultSessionType,omitempty"`
	AllowedSessionTypes []PduSessionType `json:"allowedSessionTypes,omitempty"`
}

// NewPduSessionTypes instantiates a new PduSessionTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionTypes() *PduSessionTypes {
	this := PduSessionTypes{}
	return &this
}

// NewPduSessionTypesWithDefaults instantiates a new PduSessionTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionTypesWithDefaults() *PduSessionTypes {
	this := PduSessionTypes{}
	return &this
}

// GetDefaultSessionType returns the DefaultSessionType field value if set, zero value otherwise.
func (o *PduSessionTypes) GetDefaultSessionType() PduSessionType {
	if o == nil || isNil(o.DefaultSessionType) {
		var ret PduSessionType
		return ret
	}
	return *o.DefaultSessionType
}

// GetDefaultSessionTypeOk returns a tuple with the DefaultSessionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes) GetDefaultSessionTypeOk() (*PduSessionType, bool) {
	if o == nil || isNil(o.DefaultSessionType) {
    return nil, false
	}
	return o.DefaultSessionType, true
}

// HasDefaultSessionType returns a boolean if a field has been set.
func (o *PduSessionTypes) HasDefaultSessionType() bool {
	if o != nil && !isNil(o.DefaultSessionType) {
		return true
	}

	return false
}

// SetDefaultSessionType gets a reference to the given PduSessionType and assigns it to the DefaultSessionType field.
func (o *PduSessionTypes) SetDefaultSessionType(v PduSessionType) {
	o.DefaultSessionType = &v
}

// GetAllowedSessionTypes returns the AllowedSessionTypes field value if set, zero value otherwise.
func (o *PduSessionTypes) GetAllowedSessionTypes() []PduSessionType {
	if o == nil || isNil(o.AllowedSessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.AllowedSessionTypes
}

// GetAllowedSessionTypesOk returns a tuple with the AllowedSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes) GetAllowedSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || isNil(o.AllowedSessionTypes) {
    return nil, false
	}
	return o.AllowedSessionTypes, true
}

// HasAllowedSessionTypes returns a boolean if a field has been set.
func (o *PduSessionTypes) HasAllowedSessionTypes() bool {
	if o != nil && !isNil(o.AllowedSessionTypes) {
		return true
	}

	return false
}

// SetAllowedSessionTypes gets a reference to the given []PduSessionType and assigns it to the AllowedSessionTypes field.
func (o *PduSessionTypes) SetAllowedSessionTypes(v []PduSessionType) {
	o.AllowedSessionTypes = v
}

func (o PduSessionTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSessionType) {
		toSerialize["defaultSessionType"] = o.DefaultSessionType
	}
	if !isNil(o.AllowedSessionTypes) {
		toSerialize["allowedSessionTypes"] = o.AllowedSessionTypes
	}
	return json.Marshal(toSerialize)
}

type NullablePduSessionTypes struct {
	value *PduSessionTypes
	isSet bool
}

func (v NullablePduSessionTypes) Get() *PduSessionTypes {
	return v.value
}

func (v *NullablePduSessionTypes) Set(val *PduSessionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionTypes(val *PduSessionTypes) *NullablePduSessionTypes {
	return &NullablePduSessionTypes{value: val, isSet: true}
}

func (v NullablePduSessionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


