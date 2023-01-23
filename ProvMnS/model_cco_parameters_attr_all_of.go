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

// CCOParametersAttrAllOf struct for CCOParametersAttrAllOf
type CCOParametersAttrAllOf struct {
	Attributes *CCOParametersAttrAllOfAttributes `json:"attributes,omitempty"`
}

// NewCCOParametersAttrAllOf instantiates a new CCOParametersAttrAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOParametersAttrAllOf() *CCOParametersAttrAllOf {
	this := CCOParametersAttrAllOf{}
	return &this
}

// NewCCOParametersAttrAllOfWithDefaults instantiates a new CCOParametersAttrAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOParametersAttrAllOfWithDefaults() *CCOParametersAttrAllOf {
	this := CCOParametersAttrAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOf) GetAttributes() CCOParametersAttrAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret CCOParametersAttrAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOf) GetAttributesOk() (*CCOParametersAttrAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CCOParametersAttrAllOfAttributes and assigns it to the Attributes field.
func (o *CCOParametersAttrAllOf) SetAttributes(v CCOParametersAttrAllOfAttributes) {
	o.Attributes = &v
}

func (o CCOParametersAttrAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCCOParametersAttrAllOf struct {
	value *CCOParametersAttrAllOf
	isSet bool
}

func (v NullableCCOParametersAttrAllOf) Get() *CCOParametersAttrAllOf {
	return v.value
}

func (v *NullableCCOParametersAttrAllOf) Set(val *CCOParametersAttrAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOParametersAttrAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOParametersAttrAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOParametersAttrAllOf(val *CCOParametersAttrAllOf) *NullableCCOParametersAttrAllOf {
	return &NullableCCOParametersAttrAllOf{value: val, isSet: true}
}

func (v NullableCCOParametersAttrAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOParametersAttrAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


