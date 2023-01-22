/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// AllowedSnssai Contains the authorized S-NSSAI and optional mapped home S-NSSAI and network slice instance information 
type AllowedSnssai struct {
	AllowedSnssai Snssai `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
	MappedHomeSnssai *Snssai `json:"mappedHomeSnssai,omitempty"`
}

// NewAllowedSnssai instantiates a new AllowedSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedSnssai(allowedSnssai Snssai) *AllowedSnssai {
	this := AllowedSnssai{}
	this.AllowedSnssai = allowedSnssai
	return &this
}

// NewAllowedSnssaiWithDefaults instantiates a new AllowedSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedSnssaiWithDefaults() *AllowedSnssai {
	this := AllowedSnssai{}
	return &this
}

// GetAllowedSnssai returns the AllowedSnssai field value
func (o *AllowedSnssai) GetAllowedSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.AllowedSnssai
}

// GetAllowedSnssaiOk returns a tuple with the AllowedSnssai field value
// and a boolean to check if the value has been set.
func (o *AllowedSnssai) GetAllowedSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AllowedSnssai, true
}

// SetAllowedSnssai sets field value
func (o *AllowedSnssai) SetAllowedSnssai(v Snssai) {
	o.AllowedSnssai = v
}

// GetNsiInformationList returns the NsiInformationList field value if set, zero value otherwise.
func (o *AllowedSnssai) GetNsiInformationList() []NsiInformation {
	if o == nil || isNil(o.NsiInformationList) {
		var ret []NsiInformation
		return ret
	}
	return o.NsiInformationList
}

// GetNsiInformationListOk returns a tuple with the NsiInformationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedSnssai) GetNsiInformationListOk() ([]NsiInformation, bool) {
	if o == nil || isNil(o.NsiInformationList) {
    return nil, false
	}
	return o.NsiInformationList, true
}

// HasNsiInformationList returns a boolean if a field has been set.
func (o *AllowedSnssai) HasNsiInformationList() bool {
	if o != nil && !isNil(o.NsiInformationList) {
		return true
	}

	return false
}

// SetNsiInformationList gets a reference to the given []NsiInformation and assigns it to the NsiInformationList field.
func (o *AllowedSnssai) SetNsiInformationList(v []NsiInformation) {
	o.NsiInformationList = v
}

// GetMappedHomeSnssai returns the MappedHomeSnssai field value if set, zero value otherwise.
func (o *AllowedSnssai) GetMappedHomeSnssai() Snssai {
	if o == nil || isNil(o.MappedHomeSnssai) {
		var ret Snssai
		return ret
	}
	return *o.MappedHomeSnssai
}

// GetMappedHomeSnssaiOk returns a tuple with the MappedHomeSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedSnssai) GetMappedHomeSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.MappedHomeSnssai) {
    return nil, false
	}
	return o.MappedHomeSnssai, true
}

// HasMappedHomeSnssai returns a boolean if a field has been set.
func (o *AllowedSnssai) HasMappedHomeSnssai() bool {
	if o != nil && !isNil(o.MappedHomeSnssai) {
		return true
	}

	return false
}

// SetMappedHomeSnssai gets a reference to the given Snssai and assigns it to the MappedHomeSnssai field.
func (o *AllowedSnssai) SetMappedHomeSnssai(v Snssai) {
	o.MappedHomeSnssai = &v
}

func (o AllowedSnssai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedSnssai"] = o.AllowedSnssai
	}
	if !isNil(o.NsiInformationList) {
		toSerialize["nsiInformationList"] = o.NsiInformationList
	}
	if !isNil(o.MappedHomeSnssai) {
		toSerialize["mappedHomeSnssai"] = o.MappedHomeSnssai
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedSnssai struct {
	value *AllowedSnssai
	isSet bool
}

func (v NullableAllowedSnssai) Get() *AllowedSnssai {
	return v.value
}

func (v *NullableAllowedSnssai) Set(val *AllowedSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedSnssai(val *AllowedSnssai) *NullableAllowedSnssai {
	return &NullableAllowedSnssai{value: val, isSet: true}
}

func (v NullableAllowedSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


