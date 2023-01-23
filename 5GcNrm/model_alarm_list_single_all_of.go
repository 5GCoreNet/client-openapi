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

// AlarmListSingleAllOf struct for AlarmListSingleAllOf
type AlarmListSingleAllOf struct {
	Attributes *AlarmListSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAlarmListSingleAllOf instantiates a new AlarmListSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmListSingleAllOf() *AlarmListSingleAllOf {
	this := AlarmListSingleAllOf{}
	return &this
}

// NewAlarmListSingleAllOfWithDefaults instantiates a new AlarmListSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmListSingleAllOfWithDefaults() *AlarmListSingleAllOf {
	this := AlarmListSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AlarmListSingleAllOf) GetAttributes() AlarmListSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret AlarmListSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOf) GetAttributesOk() (*AlarmListSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AlarmListSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AlarmListSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AlarmListSingleAllOf) SetAttributes(v AlarmListSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AlarmListSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmListSingleAllOf struct {
	value *AlarmListSingleAllOf
	isSet bool
}

func (v NullableAlarmListSingleAllOf) Get() *AlarmListSingleAllOf {
	return v.value
}

func (v *NullableAlarmListSingleAllOf) Set(val *AlarmListSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmListSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmListSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmListSingleAllOf(val *AlarmListSingleAllOf) *NullableAlarmListSingleAllOf {
	return &NullableAlarmListSingleAllOf{value: val, isSet: true}
}

func (v NullableAlarmListSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmListSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


