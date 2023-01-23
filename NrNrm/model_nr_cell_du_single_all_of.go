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

// NrCellDuSingleAllOf struct for NrCellDuSingleAllOf
type NrCellDuSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewNrCellDuSingleAllOf instantiates a new NrCellDuSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellDuSingleAllOf() *NrCellDuSingleAllOf {
	this := NrCellDuSingleAllOf{}
	return &this
}

// NewNrCellDuSingleAllOfWithDefaults instantiates a new NrCellDuSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellDuSingleAllOfWithDefaults() *NrCellDuSingleAllOf {
	this := NrCellDuSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NrCellDuSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellDuSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NrCellDuSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *NrCellDuSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o NrCellDuSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableNrCellDuSingleAllOf struct {
	value *NrCellDuSingleAllOf
	isSet bool
}

func (v NullableNrCellDuSingleAllOf) Get() *NrCellDuSingleAllOf {
	return v.value
}

func (v *NullableNrCellDuSingleAllOf) Set(val *NrCellDuSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellDuSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellDuSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellDuSingleAllOf(val *NrCellDuSingleAllOf) *NullableNrCellDuSingleAllOf {
	return &NullableNrCellDuSingleAllOf{value: val, isSet: true}
}

func (v NullableNrCellDuSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellDuSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


