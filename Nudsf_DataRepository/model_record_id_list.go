/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_DataRepository

import (
	"encoding/json"
)

// RecordIdList List of Record IDs
type RecordIdList struct {
	RecordIdList []string `json:"recordIdList"`
}

// NewRecordIdList instantiates a new RecordIdList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordIdList(recordIdList []string) *RecordIdList {
	this := RecordIdList{}
	this.RecordIdList = recordIdList
	return &this
}

// NewRecordIdListWithDefaults instantiates a new RecordIdList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordIdListWithDefaults() *RecordIdList {
	this := RecordIdList{}
	return &this
}

// GetRecordIdList returns the RecordIdList field value
func (o *RecordIdList) GetRecordIdList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecordIdList
}

// GetRecordIdListOk returns a tuple with the RecordIdList field value
// and a boolean to check if the value has been set.
func (o *RecordIdList) GetRecordIdListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RecordIdList, true
}

// SetRecordIdList sets field value
func (o *RecordIdList) SetRecordIdList(v []string) {
	o.RecordIdList = v
}

func (o RecordIdList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recordIdList"] = o.RecordIdList
	}
	return json.Marshal(toSerialize)
}

type NullableRecordIdList struct {
	value *RecordIdList
	isSet bool
}

func (v NullableRecordIdList) Get() *RecordIdList {
	return v.value
}

func (v *NullableRecordIdList) Set(val *RecordIdList) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordIdList) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordIdList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordIdList(val *RecordIdList) *NullableRecordIdList {
	return &NullableRecordIdList{value: val, isSet: true}
}

func (v NullableRecordIdList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordIdList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


