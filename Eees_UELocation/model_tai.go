/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
)

// Tai Contains the tracking area identity as described in 3GPP 23.003
type Tai struct {
	PlmnId PlmnId `json:"plmnId"`
	// 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.  
	Tac string `json:"tac"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewTai instantiates a new Tai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTai(plmnId PlmnId, tac string) *Tai {
	this := Tai{}
	this.PlmnId = plmnId
	this.Tac = tac
	return &this
}

// NewTaiWithDefaults instantiates a new Tai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiWithDefaults() *Tai {
	this := Tai{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Tai) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Tai) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Tai) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetTac returns the Tac field value
func (o *Tai) GetTac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tac
}

// GetTacOk returns a tuple with the Tac field value
// and a boolean to check if the value has been set.
func (o *Tai) GetTacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tac, true
}

// SetTac sets field value
func (o *Tai) SetTac(v string) {
	o.Tac = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *Tai) GetNid() string {
	if o == nil || isNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetNidOk() (*string, bool) {
	if o == nil || isNil(o.Nid) {
    return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *Tai) HasNid() bool {
	if o != nil && !isNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *Tai) SetNid(v string) {
	o.Nid = &v
}

func (o Tai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["tac"] = o.Tac
	}
	if !isNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return json.Marshal(toSerialize)
}

type NullableTai struct {
	value *Tai
	isSet bool
}

func (v NullableTai) Get() *Tai {
	return v.value
}

func (v *NullableTai) Set(val *Tai) {
	v.value = val
	v.isSet = true
}

func (v NullableTai) IsSet() bool {
	return v.isSet
}

func (v *NullableTai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTai(val *Tai) *NullableTai {
	return &NullableTai{value: val, isSet: true}
}

func (v NullableTai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


