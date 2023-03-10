/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// RimRSGlobalSingleAllOf struct for RimRSGlobalSingleAllOf
type RimRSGlobalSingleAllOf struct {
	Attributes *RimRSGlobalSingleAllOfAttributes `json:"attributes,omitempty"`
	RimRSSet []RimRSSetSingle `json:"RimRSSet,omitempty"`
}

// NewRimRSGlobalSingleAllOf instantiates a new RimRSGlobalSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSGlobalSingleAllOf() *RimRSGlobalSingleAllOf {
	this := RimRSGlobalSingleAllOf{}
	return &this
}

// NewRimRSGlobalSingleAllOfWithDefaults instantiates a new RimRSGlobalSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSGlobalSingleAllOfWithDefaults() *RimRSGlobalSingleAllOf {
	this := RimRSGlobalSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RimRSGlobalSingleAllOf) GetAttributes() RimRSGlobalSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret RimRSGlobalSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingleAllOf) GetAttributesOk() (*RimRSGlobalSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RimRSGlobalSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RimRSGlobalSingleAllOfAttributes and assigns it to the Attributes field.
func (o *RimRSGlobalSingleAllOf) SetAttributes(v RimRSGlobalSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetRimRSSet returns the RimRSSet field value if set, zero value otherwise.
func (o *RimRSGlobalSingleAllOf) GetRimRSSet() []RimRSSetSingle {
	if o == nil || isNil(o.RimRSSet) {
		var ret []RimRSSetSingle
		return ret
	}
	return o.RimRSSet
}

// GetRimRSSetOk returns a tuple with the RimRSSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingleAllOf) GetRimRSSetOk() ([]RimRSSetSingle, bool) {
	if o == nil || isNil(o.RimRSSet) {
    return nil, false
	}
	return o.RimRSSet, true
}

// HasRimRSSet returns a boolean if a field has been set.
func (o *RimRSGlobalSingleAllOf) HasRimRSSet() bool {
	if o != nil && !isNil(o.RimRSSet) {
		return true
	}

	return false
}

// SetRimRSSet gets a reference to the given []RimRSSetSingle and assigns it to the RimRSSet field.
func (o *RimRSGlobalSingleAllOf) SetRimRSSet(v []RimRSSetSingle) {
	o.RimRSSet = v
}

func (o RimRSGlobalSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.RimRSSet) {
		toSerialize["RimRSSet"] = o.RimRSSet
	}
	return json.Marshal(toSerialize)
}

type NullableRimRSGlobalSingleAllOf struct {
	value *RimRSGlobalSingleAllOf
	isSet bool
}

func (v NullableRimRSGlobalSingleAllOf) Get() *RimRSGlobalSingleAllOf {
	return v.value
}

func (v *NullableRimRSGlobalSingleAllOf) Set(val *RimRSGlobalSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSGlobalSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSGlobalSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSGlobalSingleAllOf(val *RimRSGlobalSingleAllOf) *NullableRimRSGlobalSingleAllOf {
	return &NullableRimRSGlobalSingleAllOf{value: val, isSet: true}
}

func (v NullableRimRSGlobalSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSGlobalSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


