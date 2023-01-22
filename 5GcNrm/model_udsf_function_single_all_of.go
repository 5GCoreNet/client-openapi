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

// UdsfFunctionSingleAllOf struct for UdsfFunctionSingleAllOf
type UdsfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewUdsfFunctionSingleAllOf instantiates a new UdsfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdsfFunctionSingleAllOf() *UdsfFunctionSingleAllOf {
	this := UdsfFunctionSingleAllOf{}
	return &this
}

// NewUdsfFunctionSingleAllOfWithDefaults instantiates a new UdsfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdsfFunctionSingleAllOfWithDefaults() *UdsfFunctionSingleAllOf {
	this := UdsfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *UdsfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o UdsfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableUdsfFunctionSingleAllOf struct {
	value *UdsfFunctionSingleAllOf
	isSet bool
}

func (v NullableUdsfFunctionSingleAllOf) Get() *UdsfFunctionSingleAllOf {
	return v.value
}

func (v *NullableUdsfFunctionSingleAllOf) Set(val *UdsfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUdsfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUdsfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdsfFunctionSingleAllOf(val *UdsfFunctionSingleAllOf) *NullableUdsfFunctionSingleAllOf {
	return &NullableUdsfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableUdsfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdsfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


