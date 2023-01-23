/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// NfGroupListCond Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {
	ConditionType string `json:"conditionType"`
	NfType string `json:"nfType"`
	NfGroupIdList []string `json:"nfGroupIdList"`
}

// NewNfGroupListCond instantiates a new NfGroupListCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfGroupListCond(conditionType string, nfType string, nfGroupIdList []string) *NfGroupListCond {
	this := NfGroupListCond{}
	this.ConditionType = conditionType
	this.NfType = nfType
	this.NfGroupIdList = nfGroupIdList
	return &this
}

// NewNfGroupListCondWithDefaults instantiates a new NfGroupListCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfGroupListCondWithDefaults() *NfGroupListCond {
	this := NfGroupListCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *NfGroupListCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *NfGroupListCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *NfGroupListCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetNfType returns the NfType field value
func (o *NfGroupListCond) GetNfType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *NfGroupListCond) GetNfTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *NfGroupListCond) SetNfType(v string) {
	o.NfType = v
}

// GetNfGroupIdList returns the NfGroupIdList field value
func (o *NfGroupListCond) GetNfGroupIdList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NfGroupIdList
}

// GetNfGroupIdListOk returns a tuple with the NfGroupIdList field value
// and a boolean to check if the value has been set.
func (o *NfGroupListCond) GetNfGroupIdListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NfGroupIdList, true
}

// SetNfGroupIdList sets field value
func (o *NfGroupListCond) SetNfGroupIdList(v []string) {
	o.NfGroupIdList = v
}

func (o NfGroupListCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conditionType"] = o.ConditionType
	}
	if true {
		toSerialize["nfType"] = o.NfType
	}
	if true {
		toSerialize["nfGroupIdList"] = o.NfGroupIdList
	}
	return json.Marshal(toSerialize)
}

type NullableNfGroupListCond struct {
	value *NfGroupListCond
	isSet bool
}

func (v NullableNfGroupListCond) Get() *NfGroupListCond {
	return v.value
}

func (v *NullableNfGroupListCond) Set(val *NfGroupListCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfGroupListCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfGroupListCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfGroupListCond(val *NfGroupListCond) *NullableNfGroupListCond {
	return &NullableNfGroupListCond{value: val, isSet: true}
}

func (v NullableNfGroupListCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfGroupListCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


