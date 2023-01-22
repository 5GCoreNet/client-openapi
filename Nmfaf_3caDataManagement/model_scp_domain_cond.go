/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// ScpDomainCond Subscription to a set of NF or SCP or SEPP instances belonging to certain SCP domains 
type ScpDomainCond struct {
	ScpDomains []string `json:"scpDomains"`
	NfTypeList []NFType `json:"nfTypeList,omitempty"`
}

// NewScpDomainCond instantiates a new ScpDomainCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpDomainCond(scpDomains []string) *ScpDomainCond {
	this := ScpDomainCond{}
	this.ScpDomains = scpDomains
	return &this
}

// NewScpDomainCondWithDefaults instantiates a new ScpDomainCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpDomainCondWithDefaults() *ScpDomainCond {
	this := ScpDomainCond{}
	return &this
}

// GetScpDomains returns the ScpDomains field value
func (o *ScpDomainCond) GetScpDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScpDomains
}

// GetScpDomainsOk returns a tuple with the ScpDomains field value
// and a boolean to check if the value has been set.
func (o *ScpDomainCond) GetScpDomainsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ScpDomains, true
}

// SetScpDomains sets field value
func (o *ScpDomainCond) SetScpDomains(v []string) {
	o.ScpDomains = v
}

// GetNfTypeList returns the NfTypeList field value if set, zero value otherwise.
func (o *ScpDomainCond) GetNfTypeList() []NFType {
	if o == nil || isNil(o.NfTypeList) {
		var ret []NFType
		return ret
	}
	return o.NfTypeList
}

// GetNfTypeListOk returns a tuple with the NfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpDomainCond) GetNfTypeListOk() ([]NFType, bool) {
	if o == nil || isNil(o.NfTypeList) {
    return nil, false
	}
	return o.NfTypeList, true
}

// HasNfTypeList returns a boolean if a field has been set.
func (o *ScpDomainCond) HasNfTypeList() bool {
	if o != nil && !isNil(o.NfTypeList) {
		return true
	}

	return false
}

// SetNfTypeList gets a reference to the given []NFType and assigns it to the NfTypeList field.
func (o *ScpDomainCond) SetNfTypeList(v []NFType) {
	o.NfTypeList = v
}

func (o ScpDomainCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scpDomains"] = o.ScpDomains
	}
	if !isNil(o.NfTypeList) {
		toSerialize["nfTypeList"] = o.NfTypeList
	}
	return json.Marshal(toSerialize)
}

type NullableScpDomainCond struct {
	value *ScpDomainCond
	isSet bool
}

func (v NullableScpDomainCond) Get() *ScpDomainCond {
	return v.value
}

func (v *NullableScpDomainCond) Set(val *ScpDomainCond) {
	v.value = val
	v.isSet = true
}

func (v NullableScpDomainCond) IsSet() bool {
	return v.isSet
}

func (v *NullableScpDomainCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpDomainCond(val *ScpDomainCond) *NullableScpDomainCond {
	return &NullableScpDomainCond{value: val, isSet: true}
}

func (v NullableScpDomainCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpDomainCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


