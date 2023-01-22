/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// Tai1 struct for Tai1
type Tai1 struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
	Tac *string `json:"tac,omitempty"`
}

// NewTai1 instantiates a new Tai1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTai1() *Tai1 {
	this := Tai1{}
	return &this
}

// NewTai1WithDefaults instantiates a new Tai1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTai1WithDefaults() *Tai1 {
	this := Tai1{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *Tai1) GetMcc() string {
	if o == nil || isNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai1) GetMccOk() (*string, bool) {
	if o == nil || isNil(o.Mcc) {
    return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *Tai1) HasMcc() bool {
	if o != nil && !isNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *Tai1) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *Tai1) GetMnc() string {
	if o == nil || isNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai1) GetMncOk() (*string, bool) {
	if o == nil || isNil(o.Mnc) {
    return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *Tai1) HasMnc() bool {
	if o != nil && !isNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *Tai1) SetMnc(v string) {
	o.Mnc = &v
}

// GetTac returns the Tac field value if set, zero value otherwise.
func (o *Tai1) GetTac() string {
	if o == nil || isNil(o.Tac) {
		var ret string
		return ret
	}
	return *o.Tac
}

// GetTacOk returns a tuple with the Tac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai1) GetTacOk() (*string, bool) {
	if o == nil || isNil(o.Tac) {
    return nil, false
	}
	return o.Tac, true
}

// HasTac returns a boolean if a field has been set.
func (o *Tai1) HasTac() bool {
	if o != nil && !isNil(o.Tac) {
		return true
	}

	return false
}

// SetTac gets a reference to the given string and assigns it to the Tac field.
func (o *Tai1) SetTac(v string) {
	o.Tac = &v
}

func (o Tai1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !isNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !isNil(o.Tac) {
		toSerialize["tac"] = o.Tac
	}
	return json.Marshal(toSerialize)
}

type NullableTai1 struct {
	value *Tai1
	isSet bool
}

func (v NullableTai1) Get() *Tai1 {
	return v.value
}

func (v *NullableTai1) Set(val *Tai1) {
	v.value = val
	v.isSet = true
}

func (v NullableTai1) IsSet() bool {
	return v.isSet
}

func (v *NullableTai1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTai1(val *Tai1) *NullableTai1 {
	return &NullableTai1{value: val, isSet: true}
}

func (v NullableTai1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTai1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


