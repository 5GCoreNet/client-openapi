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

// HeartbeatControlSingleAllOf struct for HeartbeatControlSingleAllOf
type HeartbeatControlSingleAllOf struct {
	Attributes *HeartbeatControlSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewHeartbeatControlSingleAllOf instantiates a new HeartbeatControlSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeartbeatControlSingleAllOf() *HeartbeatControlSingleAllOf {
	this := HeartbeatControlSingleAllOf{}
	return &this
}

// NewHeartbeatControlSingleAllOfWithDefaults instantiates a new HeartbeatControlSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeartbeatControlSingleAllOfWithDefaults() *HeartbeatControlSingleAllOf {
	this := HeartbeatControlSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *HeartbeatControlSingleAllOf) GetAttributes() HeartbeatControlSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret HeartbeatControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatControlSingleAllOf) GetAttributesOk() (*HeartbeatControlSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *HeartbeatControlSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given HeartbeatControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *HeartbeatControlSingleAllOf) SetAttributes(v HeartbeatControlSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o HeartbeatControlSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableHeartbeatControlSingleAllOf struct {
	value *HeartbeatControlSingleAllOf
	isSet bool
}

func (v NullableHeartbeatControlSingleAllOf) Get() *HeartbeatControlSingleAllOf {
	return v.value
}

func (v *NullableHeartbeatControlSingleAllOf) Set(val *HeartbeatControlSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartbeatControlSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartbeatControlSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartbeatControlSingleAllOf(val *HeartbeatControlSingleAllOf) *NullableHeartbeatControlSingleAllOf {
	return &NullableHeartbeatControlSingleAllOf{value: val, isSet: true}
}

func (v NullableHeartbeatControlSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartbeatControlSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


