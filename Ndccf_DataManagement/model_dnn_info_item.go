/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
)

// DnnInfoItem Set of parameters supported by NF for a given DNN
type DnnInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
}

// NewDnnInfoItem instantiates a new DnnInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnInfoItem(dnn DnnSmfInfoItemDnn) *DnnInfoItem {
	this := DnnInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnInfoItemWithDefaults instantiates a new DnnInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnInfoItemWithDefaults() *DnnInfoItem {
	this := DnnInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

func (o DnnInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	return json.Marshal(toSerialize)
}

type NullableDnnInfoItem struct {
	value *DnnInfoItem
	isSet bool
}

func (v NullableDnnInfoItem) Get() *DnnInfoItem {
	return v.value
}

func (v *NullableDnnInfoItem) Set(val *DnnInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnInfoItem(val *DnnInfoItem) *NullableDnnInfoItem {
	return &NullableDnnInfoItem{value: val, isSet: true}
}

func (v NullableDnnInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


