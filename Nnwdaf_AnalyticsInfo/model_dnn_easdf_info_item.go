/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// DnnEasdfInfoItem Set of parameters supported by EASDF for a given DNN
type DnnEasdfInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
	DnaiList []string `json:"dnaiList,omitempty"`
}

// NewDnnEasdfInfoItem instantiates a new DnnEasdfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnEasdfInfoItem(dnn DnnSmfInfoItemDnn) *DnnEasdfInfoItem {
	this := DnnEasdfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnEasdfInfoItemWithDefaults instantiates a new DnnEasdfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnEasdfInfoItemWithDefaults() *DnnEasdfInfoItem {
	this := DnnEasdfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnEasdfInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnEasdfInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnEasdfInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnEasdfInfoItem) GetDnaiList() []string {
	if o == nil || isNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnEasdfInfoItem) GetDnaiListOk() ([]string, bool) {
	if o == nil || isNil(o.DnaiList) {
    return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *DnnEasdfInfoItem) HasDnaiList() bool {
	if o != nil && !isNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *DnnEasdfInfoItem) SetDnaiList(v []string) {
	o.DnaiList = v
}

func (o DnnEasdfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	return json.Marshal(toSerialize)
}

type NullableDnnEasdfInfoItem struct {
	value *DnnEasdfInfoItem
	isSet bool
}

func (v NullableDnnEasdfInfoItem) Get() *DnnEasdfInfoItem {
	return v.value
}

func (v *NullableDnnEasdfInfoItem) Set(val *DnnEasdfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnEasdfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnEasdfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnEasdfInfoItem(val *DnnEasdfInfoItem) *NullableDnnEasdfInfoItem {
	return &NullableDnnEasdfInfoItem{value: val, isSet: true}
}

func (v NullableDnnEasdfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnEasdfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


