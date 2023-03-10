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

// EPN2SingleAllOf struct for EPN2SingleAllOf
type EPN2SingleAllOf struct {
	Attributes *EPRPAttr `json:"attributes,omitempty"`
}

// NewEPN2SingleAllOf instantiates a new EPN2SingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN2SingleAllOf() *EPN2SingleAllOf {
	this := EPN2SingleAllOf{}
	return &this
}

// NewEPN2SingleAllOfWithDefaults instantiates a new EPN2SingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN2SingleAllOfWithDefaults() *EPN2SingleAllOf {
	this := EPN2SingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN2SingleAllOf) GetAttributes() EPRPAttr {
	if o == nil || isNil(o.Attributes) {
		var ret EPRPAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN2SingleAllOf) GetAttributesOk() (*EPRPAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN2SingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPRPAttr and assigns it to the Attributes field.
func (o *EPN2SingleAllOf) SetAttributes(v EPRPAttr) {
	o.Attributes = &v
}

func (o EPN2SingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEPN2SingleAllOf struct {
	value *EPN2SingleAllOf
	isSet bool
}

func (v NullableEPN2SingleAllOf) Get() *EPN2SingleAllOf {
	return v.value
}

func (v *NullableEPN2SingleAllOf) Set(val *EPN2SingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN2SingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN2SingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN2SingleAllOf(val *EPN2SingleAllOf) *NullableEPN2SingleAllOf {
	return &NullableEPN2SingleAllOf{value: val, isSet: true}
}

func (v NullableEPN2SingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN2SingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


