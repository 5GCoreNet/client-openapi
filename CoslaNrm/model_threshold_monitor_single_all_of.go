/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// ThresholdMonitorSingleAllOf struct for ThresholdMonitorSingleAllOf
type ThresholdMonitorSingleAllOf struct {
	Attributes *ThresholdMonitorSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewThresholdMonitorSingleAllOf instantiates a new ThresholdMonitorSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdMonitorSingleAllOf() *ThresholdMonitorSingleAllOf {
	this := ThresholdMonitorSingleAllOf{}
	return &this
}

// NewThresholdMonitorSingleAllOfWithDefaults instantiates a new ThresholdMonitorSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdMonitorSingleAllOfWithDefaults() *ThresholdMonitorSingleAllOf {
	this := ThresholdMonitorSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOf) GetAttributes() ThresholdMonitorSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret ThresholdMonitorSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOf) GetAttributesOk() (*ThresholdMonitorSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ThresholdMonitorSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ThresholdMonitorSingleAllOf) SetAttributes(v ThresholdMonitorSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ThresholdMonitorSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableThresholdMonitorSingleAllOf struct {
	value *ThresholdMonitorSingleAllOf
	isSet bool
}

func (v NullableThresholdMonitorSingleAllOf) Get() *ThresholdMonitorSingleAllOf {
	return v.value
}

func (v *NullableThresholdMonitorSingleAllOf) Set(val *ThresholdMonitorSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdMonitorSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdMonitorSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdMonitorSingleAllOf(val *ThresholdMonitorSingleAllOf) *NullableThresholdMonitorSingleAllOf {
	return &NullableThresholdMonitorSingleAllOf{value: val, isSet: true}
}

func (v NullableThresholdMonitorSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdMonitorSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


