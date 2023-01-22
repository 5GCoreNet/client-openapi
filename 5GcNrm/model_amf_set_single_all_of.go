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

// AmfSetSingleAllOf struct for AmfSetSingleAllOf
type AmfSetSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewAmfSetSingleAllOf instantiates a new AmfSetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfSetSingleAllOf() *AmfSetSingleAllOf {
	this := AmfSetSingleAllOf{}
	return &this
}

// NewAmfSetSingleAllOfWithDefaults instantiates a new AmfSetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfSetSingleAllOfWithDefaults() *AmfSetSingleAllOf {
	this := AmfSetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AmfSetSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfSetSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AmfSetSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *AmfSetSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o AmfSetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAmfSetSingleAllOf struct {
	value *AmfSetSingleAllOf
	isSet bool
}

func (v NullableAmfSetSingleAllOf) Get() *AmfSetSingleAllOf {
	return v.value
}

func (v *NullableAmfSetSingleAllOf) Set(val *AmfSetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfSetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfSetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfSetSingleAllOf(val *AmfSetSingleAllOf) *NullableAmfSetSingleAllOf {
	return &NullableAmfSetSingleAllOf{value: val, isSet: true}
}

func (v NullableAmfSetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfSetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

