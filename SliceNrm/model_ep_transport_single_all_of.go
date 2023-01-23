/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// EPTransportSingleAllOf struct for EPTransportSingleAllOf
type EPTransportSingleAllOf struct {
	Attributes *EPTransportSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPTransportSingleAllOf instantiates a new EPTransportSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPTransportSingleAllOf() *EPTransportSingleAllOf {
	this := EPTransportSingleAllOf{}
	return &this
}

// NewEPTransportSingleAllOfWithDefaults instantiates a new EPTransportSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPTransportSingleAllOfWithDefaults() *EPTransportSingleAllOf {
	this := EPTransportSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPTransportSingleAllOf) GetAttributes() EPTransportSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret EPTransportSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOf) GetAttributesOk() (*EPTransportSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPTransportSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPTransportSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPTransportSingleAllOf) SetAttributes(v EPTransportSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPTransportSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEPTransportSingleAllOf struct {
	value *EPTransportSingleAllOf
	isSet bool
}

func (v NullableEPTransportSingleAllOf) Get() *EPTransportSingleAllOf {
	return v.value
}

func (v *NullableEPTransportSingleAllOf) Set(val *EPTransportSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPTransportSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPTransportSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPTransportSingleAllOf(val *EPTransportSingleAllOf) *NullableEPTransportSingleAllOf {
	return &NullableEPTransportSingleAllOf{value: val, isSet: true}
}

func (v NullableEPTransportSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPTransportSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


