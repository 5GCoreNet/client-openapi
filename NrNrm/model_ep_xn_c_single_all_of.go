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

// EPXnCSingleAllOf struct for EPXnCSingleAllOf
type EPXnCSingleAllOf struct {
	Attributes *EPRPAttr `json:"attributes,omitempty"`
}

// NewEPXnCSingleAllOf instantiates a new EPXnCSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPXnCSingleAllOf() *EPXnCSingleAllOf {
	this := EPXnCSingleAllOf{}
	return &this
}

// NewEPXnCSingleAllOfWithDefaults instantiates a new EPXnCSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPXnCSingleAllOfWithDefaults() *EPXnCSingleAllOf {
	this := EPXnCSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPXnCSingleAllOf) GetAttributes() EPRPAttr {
	if o == nil || isNil(o.Attributes) {
		var ret EPRPAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPXnCSingleAllOf) GetAttributesOk() (*EPRPAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPXnCSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPRPAttr and assigns it to the Attributes field.
func (o *EPXnCSingleAllOf) SetAttributes(v EPRPAttr) {
	o.Attributes = &v
}

func (o EPXnCSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEPXnCSingleAllOf struct {
	value *EPXnCSingleAllOf
	isSet bool
}

func (v NullableEPXnCSingleAllOf) Get() *EPXnCSingleAllOf {
	return v.value
}

func (v *NullableEPXnCSingleAllOf) Set(val *EPXnCSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPXnCSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPXnCSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPXnCSingleAllOf(val *EPXnCSingleAllOf) *NullableEPXnCSingleAllOf {
	return &NullableEPXnCSingleAllOf{value: val, isSet: true}
}

func (v NullableEPXnCSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPXnCSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


