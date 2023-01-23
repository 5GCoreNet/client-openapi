/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// DnnSmfInfoItem Set of parameters supported by SMF for a given DNN
type DnnSmfInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
	DnaiList []DnnSmfInfoItemDnaiListInner `json:"dnaiList,omitempty"`
}

// NewDnnSmfInfoItem instantiates a new DnnSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnSmfInfoItem(dnn DnnSmfInfoItemDnn) *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnSmfInfoItemWithDefaults instantiates a new DnnSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnSmfInfoItemWithDefaults() *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnSmfInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnSmfInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnSmfInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnSmfInfoItem) GetDnaiList() []DnnSmfInfoItemDnaiListInner {
	if o == nil || isNil(o.DnaiList) {
		var ret []DnnSmfInfoItemDnaiListInner
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnSmfInfoItem) GetDnaiListOk() ([]DnnSmfInfoItemDnaiListInner, bool) {
	if o == nil || isNil(o.DnaiList) {
    return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *DnnSmfInfoItem) HasDnaiList() bool {
	if o != nil && !isNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []DnnSmfInfoItemDnaiListInner and assigns it to the DnaiList field.
func (o *DnnSmfInfoItem) SetDnaiList(v []DnnSmfInfoItemDnaiListInner) {
	o.DnaiList = v
}

func (o DnnSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	return json.Marshal(toSerialize)
}

type NullableDnnSmfInfoItem struct {
	value *DnnSmfInfoItem
	isSet bool
}

func (v NullableDnnSmfInfoItem) Get() *DnnSmfInfoItem {
	return v.value
}

func (v *NullableDnnSmfInfoItem) Set(val *DnnSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnSmfInfoItem(val *DnnSmfInfoItem) *NullableDnnSmfInfoItem {
	return &NullableDnnSmfInfoItem{value: val, isSet: true}
}

func (v NullableDnnSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


