/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// Ecgi1 Contains the ECGI (E-UTRAN Cell Global Identity), as described in 3GPP 23.003
type Ecgi1 struct {
	PlmnId PlmnId1 `json:"plmnId"`
	// 28-bit string identifying an E-UTRA Cell Id as specified in clause 9.3.1.9 of  3GPP TS 38.413, in hexadecimal representation. Each character in the string shall take a  value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most  significant character representing the 4 most significant bits of the Cell Id shall appear  first in the string, and the character representing the 4 least significant bit of the  Cell Id shall appear last in the string.  
	EutraCellId string `json:"eutraCellId"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewEcgi1 instantiates a new Ecgi1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcgi1(plmnId PlmnId1, eutraCellId string) *Ecgi1 {
	this := Ecgi1{}
	this.PlmnId = plmnId
	this.EutraCellId = eutraCellId
	return &this
}

// NewEcgi1WithDefaults instantiates a new Ecgi1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcgi1WithDefaults() *Ecgi1 {
	this := Ecgi1{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Ecgi1) GetPlmnId() PlmnId1 {
	if o == nil {
		var ret PlmnId1
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Ecgi1) GetPlmnIdOk() (*PlmnId1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Ecgi1) SetPlmnId(v PlmnId1) {
	o.PlmnId = v
}

// GetEutraCellId returns the EutraCellId field value
func (o *Ecgi1) GetEutraCellId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EutraCellId
}

// GetEutraCellIdOk returns a tuple with the EutraCellId field value
// and a boolean to check if the value has been set.
func (o *Ecgi1) GetEutraCellIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EutraCellId, true
}

// SetEutraCellId sets field value
func (o *Ecgi1) SetEutraCellId(v string) {
	o.EutraCellId = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *Ecgi1) GetNid() string {
	if o == nil || isNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ecgi1) GetNidOk() (*string, bool) {
	if o == nil || isNil(o.Nid) {
    return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *Ecgi1) HasNid() bool {
	if o != nil && !isNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *Ecgi1) SetNid(v string) {
	o.Nid = &v
}

func (o Ecgi1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["eutraCellId"] = o.EutraCellId
	}
	if !isNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return json.Marshal(toSerialize)
}

type NullableEcgi1 struct {
	value *Ecgi1
	isSet bool
}

func (v NullableEcgi1) Get() *Ecgi1 {
	return v.value
}

func (v *NullableEcgi1) Set(val *Ecgi1) {
	v.value = val
	v.isSet = true
}

func (v NullableEcgi1) IsSet() bool {
	return v.isSet
}

func (v *NullableEcgi1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcgi1(val *Ecgi1) *NullableEcgi1 {
	return &NullableEcgi1{value: val, isSet: true}
}

func (v NullableEcgi1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcgi1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


