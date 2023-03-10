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

// NgeirFunctionSingleAllOf struct for NgeirFunctionSingleAllOf
type NgeirFunctionSingleAllOf struct {
	EPN17 []EPN17Single `json:"EP_N17,omitempty"`
}

// NewNgeirFunctionSingleAllOf instantiates a new NgeirFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgeirFunctionSingleAllOf() *NgeirFunctionSingleAllOf {
	this := NgeirFunctionSingleAllOf{}
	return &this
}

// NewNgeirFunctionSingleAllOfWithDefaults instantiates a new NgeirFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgeirFunctionSingleAllOfWithDefaults() *NgeirFunctionSingleAllOf {
	this := NgeirFunctionSingleAllOf{}
	return &this
}

// GetEPN17 returns the EPN17 field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOf) GetEPN17() []EPN17Single {
	if o == nil || isNil(o.EPN17) {
		var ret []EPN17Single
		return ret
	}
	return o.EPN17
}

// GetEPN17Ok returns a tuple with the EPN17 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOf) GetEPN17Ok() ([]EPN17Single, bool) {
	if o == nil || isNil(o.EPN17) {
    return nil, false
	}
	return o.EPN17, true
}

// HasEPN17 returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOf) HasEPN17() bool {
	if o != nil && !isNil(o.EPN17) {
		return true
	}

	return false
}

// SetEPN17 gets a reference to the given []EPN17Single and assigns it to the EPN17 field.
func (o *NgeirFunctionSingleAllOf) SetEPN17(v []EPN17Single) {
	o.EPN17 = v
}

func (o NgeirFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EPN17) {
		toSerialize["EP_N17"] = o.EPN17
	}
	return json.Marshal(toSerialize)
}

type NullableNgeirFunctionSingleAllOf struct {
	value *NgeirFunctionSingleAllOf
	isSet bool
}

func (v NullableNgeirFunctionSingleAllOf) Get() *NgeirFunctionSingleAllOf {
	return v.value
}

func (v *NullableNgeirFunctionSingleAllOf) Set(val *NgeirFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgeirFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgeirFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgeirFunctionSingleAllOf(val *NgeirFunctionSingleAllOf) *NullableNgeirFunctionSingleAllOf {
	return &NullableNgeirFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableNgeirFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgeirFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


