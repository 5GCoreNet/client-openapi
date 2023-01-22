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

// ExternalAmfFunctionSingleAllOf struct for ExternalAmfFunctionSingleAllOf
type ExternalAmfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewExternalAmfFunctionSingleAllOf instantiates a new ExternalAmfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAmfFunctionSingleAllOf() *ExternalAmfFunctionSingleAllOf {
	this := ExternalAmfFunctionSingleAllOf{}
	return &this
}

// NewExternalAmfFunctionSingleAllOfWithDefaults instantiates a new ExternalAmfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAmfFunctionSingleAllOfWithDefaults() *ExternalAmfFunctionSingleAllOf {
	this := ExternalAmfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalAmfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAmfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalAmfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalAmfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o ExternalAmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAmfFunctionSingleAllOf struct {
	value *ExternalAmfFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalAmfFunctionSingleAllOf) Get() *ExternalAmfFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalAmfFunctionSingleAllOf) Set(val *ExternalAmfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAmfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAmfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAmfFunctionSingleAllOf(val *ExternalAmfFunctionSingleAllOf) *NullableExternalAmfFunctionSingleAllOf {
	return &NullableExternalAmfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalAmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAmfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


