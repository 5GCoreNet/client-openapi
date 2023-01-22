/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// Dynamic5QISetSingleAllOf struct for Dynamic5QISetSingleAllOf
type Dynamic5QISetSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewDynamic5QISetSingleAllOf instantiates a new Dynamic5QISetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamic5QISetSingleAllOf() *Dynamic5QISetSingleAllOf {
	this := Dynamic5QISetSingleAllOf{}
	return &this
}

// NewDynamic5QISetSingleAllOfWithDefaults instantiates a new Dynamic5QISetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamic5QISetSingleAllOfWithDefaults() *Dynamic5QISetSingleAllOf {
	this := Dynamic5QISetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Dynamic5QISetSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5QISetSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Dynamic5QISetSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *Dynamic5QISetSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o Dynamic5QISetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableDynamic5QISetSingleAllOf struct {
	value *Dynamic5QISetSingleAllOf
	isSet bool
}

func (v NullableDynamic5QISetSingleAllOf) Get() *Dynamic5QISetSingleAllOf {
	return v.value
}

func (v *NullableDynamic5QISetSingleAllOf) Set(val *Dynamic5QISetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamic5QISetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamic5QISetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamic5QISetSingleAllOf(val *Dynamic5QISetSingleAllOf) *NullableDynamic5QISetSingleAllOf {
	return &NullableDynamic5QISetSingleAllOf{value: val, isSet: true}
}

func (v NullableDynamic5QISetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamic5QISetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


