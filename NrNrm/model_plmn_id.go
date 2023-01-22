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

// PlmnId struct for PlmnId
type PlmnId struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
}

// NewPlmnId instantiates a new PlmnId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnId() *PlmnId {
	this := PlmnId{}
	return &this
}

// NewPlmnIdWithDefaults instantiates a new PlmnId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdWithDefaults() *PlmnId {
	this := PlmnId{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *PlmnId) GetMcc() string {
	if o == nil || isNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMccOk() (*string, bool) {
	if o == nil || isNil(o.Mcc) {
    return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *PlmnId) HasMcc() bool {
	if o != nil && !isNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *PlmnId) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *PlmnId) GetMnc() string {
	if o == nil || isNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMncOk() (*string, bool) {
	if o == nil || isNil(o.Mnc) {
    return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *PlmnId) HasMnc() bool {
	if o != nil && !isNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *PlmnId) SetMnc(v string) {
	o.Mnc = &v
}

func (o PlmnId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !isNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnId struct {
	value *PlmnId
	isSet bool
}

func (v NullablePlmnId) Get() *PlmnId {
	return v.value
}

func (v *NullablePlmnId) Set(val *PlmnId) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnId) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnId(val *PlmnId) *NullablePlmnId {
	return &NullablePlmnId{value: val, isSet: true}
}

func (v NullablePlmnId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


