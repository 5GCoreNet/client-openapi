/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// DnnSmfInfoItem struct for DnnSmfInfoItem
type DnnSmfInfoItem struct {
	Dnn *string `json:"dnn,omitempty"`
	DnaiList []string `json:"dnaiList,omitempty"`
}

// NewDnnSmfInfoItem instantiates a new DnnSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnSmfInfoItem() *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	return &this
}

// NewDnnSmfInfoItemWithDefaults instantiates a new DnnSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnSmfInfoItemWithDefaults() *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *DnnSmfInfoItem) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnSmfInfoItem) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *DnnSmfInfoItem) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *DnnSmfInfoItem) SetDnn(v string) {
	o.Dnn = &v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnSmfInfoItem) GetDnaiList() []string {
	if o == nil || isNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnSmfInfoItem) GetDnaiListOk() ([]string, bool) {
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

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *DnnSmfInfoItem) SetDnaiList(v []string) {
	o.DnaiList = v
}

func (o DnnSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dnn) {
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


