/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// LmfFunctionSingleAllOf struct for LmfFunctionSingleAllOf
type LmfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewLmfFunctionSingleAllOf instantiates a new LmfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLmfFunctionSingleAllOf() *LmfFunctionSingleAllOf {
	this := LmfFunctionSingleAllOf{}
	return &this
}

// NewLmfFunctionSingleAllOfWithDefaults instantiates a new LmfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLmfFunctionSingleAllOfWithDefaults() *LmfFunctionSingleAllOf {
	this := LmfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *LmfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LmfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *LmfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o LmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableLmfFunctionSingleAllOf struct {
	value *LmfFunctionSingleAllOf
	isSet bool
}

func (v NullableLmfFunctionSingleAllOf) Get() *LmfFunctionSingleAllOf {
	return v.value
}

func (v *NullableLmfFunctionSingleAllOf) Set(val *LmfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLmfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLmfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLmfFunctionSingleAllOf(val *LmfFunctionSingleAllOf) *NullableLmfFunctionSingleAllOf {
	return &NullableLmfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableLmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLmfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


