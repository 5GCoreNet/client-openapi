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

// RRMPolicyRatioSingleAllOf struct for RRMPolicyRatioSingleAllOf
type RRMPolicyRatioSingleAllOf struct {
	Attributes *RrmPolicyAttr `json:"attributes,omitempty"`
}

// NewRRMPolicyRatioSingleAllOf instantiates a new RRMPolicyRatioSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRRMPolicyRatioSingleAllOf() *RRMPolicyRatioSingleAllOf {
	this := RRMPolicyRatioSingleAllOf{}
	return &this
}

// NewRRMPolicyRatioSingleAllOfWithDefaults instantiates a new RRMPolicyRatioSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRRMPolicyRatioSingleAllOfWithDefaults() *RRMPolicyRatioSingleAllOf {
	this := RRMPolicyRatioSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RRMPolicyRatioSingleAllOf) GetAttributes() RrmPolicyAttr {
	if o == nil || isNil(o.Attributes) {
		var ret RrmPolicyAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RRMPolicyRatioSingleAllOf) GetAttributesOk() (*RrmPolicyAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RRMPolicyRatioSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RrmPolicyAttr and assigns it to the Attributes field.
func (o *RRMPolicyRatioSingleAllOf) SetAttributes(v RrmPolicyAttr) {
	o.Attributes = &v
}

func (o RRMPolicyRatioSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableRRMPolicyRatioSingleAllOf struct {
	value *RRMPolicyRatioSingleAllOf
	isSet bool
}

func (v NullableRRMPolicyRatioSingleAllOf) Get() *RRMPolicyRatioSingleAllOf {
	return v.value
}

func (v *NullableRRMPolicyRatioSingleAllOf) Set(val *RRMPolicyRatioSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRRMPolicyRatioSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRRMPolicyRatioSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRRMPolicyRatioSingleAllOf(val *RRMPolicyRatioSingleAllOf) *NullableRRMPolicyRatioSingleAllOf {
	return &NullableRRMPolicyRatioSingleAllOf{value: val, isSet: true}
}

func (v NullableRRMPolicyRatioSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRRMPolicyRatioSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


