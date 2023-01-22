/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// GUAMInfo struct for GUAMInfo
type GUAMInfo struct {
	PLMNId *PlmnId `json:"pLMNId,omitempty"`
	AMFIdentifier *int32 `json:"aMFIdentifier,omitempty"`
}

// NewGUAMInfo instantiates a new GUAMInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGUAMInfo() *GUAMInfo {
	this := GUAMInfo{}
	return &this
}

// NewGUAMInfoWithDefaults instantiates a new GUAMInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGUAMInfoWithDefaults() *GUAMInfo {
	this := GUAMInfo{}
	return &this
}

// GetPLMNId returns the PLMNId field value if set, zero value otherwise.
func (o *GUAMInfo) GetPLMNId() PlmnId {
	if o == nil || isNil(o.PLMNId) {
		var ret PlmnId
		return ret
	}
	return *o.PLMNId
}

// GetPLMNIdOk returns a tuple with the PLMNId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GUAMInfo) GetPLMNIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PLMNId) {
    return nil, false
	}
	return o.PLMNId, true
}

// HasPLMNId returns a boolean if a field has been set.
func (o *GUAMInfo) HasPLMNId() bool {
	if o != nil && !isNil(o.PLMNId) {
		return true
	}

	return false
}

// SetPLMNId gets a reference to the given PlmnId and assigns it to the PLMNId field.
func (o *GUAMInfo) SetPLMNId(v PlmnId) {
	o.PLMNId = &v
}

// GetAMFIdentifier returns the AMFIdentifier field value if set, zero value otherwise.
func (o *GUAMInfo) GetAMFIdentifier() int32 {
	if o == nil || isNil(o.AMFIdentifier) {
		var ret int32
		return ret
	}
	return *o.AMFIdentifier
}

// GetAMFIdentifierOk returns a tuple with the AMFIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GUAMInfo) GetAMFIdentifierOk() (*int32, bool) {
	if o == nil || isNil(o.AMFIdentifier) {
    return nil, false
	}
	return o.AMFIdentifier, true
}

// HasAMFIdentifier returns a boolean if a field has been set.
func (o *GUAMInfo) HasAMFIdentifier() bool {
	if o != nil && !isNil(o.AMFIdentifier) {
		return true
	}

	return false
}

// SetAMFIdentifier gets a reference to the given int32 and assigns it to the AMFIdentifier field.
func (o *GUAMInfo) SetAMFIdentifier(v int32) {
	o.AMFIdentifier = &v
}

func (o GUAMInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PLMNId) {
		toSerialize["pLMNId"] = o.PLMNId
	}
	if !isNil(o.AMFIdentifier) {
		toSerialize["aMFIdentifier"] = o.AMFIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableGUAMInfo struct {
	value *GUAMInfo
	isSet bool
}

func (v NullableGUAMInfo) Get() *GUAMInfo {
	return v.value
}

func (v *NullableGUAMInfo) Set(val *GUAMInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGUAMInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGUAMInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGUAMInfo(val *GUAMInfo) *NullableGUAMInfo {
	return &NullableGUAMInfo{value: val, isSet: true}
}

func (v NullableGUAMInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGUAMInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


