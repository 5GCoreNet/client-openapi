/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// SmContextReleasedData Data within Release SM Context Response
type SmContextReleasedData struct {
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus *ApnRateStatus `json:"apnRateStatus,omitempty"`
}

// NewSmContextReleasedData instantiates a new SmContextReleasedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextReleasedData() *SmContextReleasedData {
	this := SmContextReleasedData{}
	return &this
}

// NewSmContextReleasedDataWithDefaults instantiates a new SmContextReleasedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextReleasedDataWithDefaults() *SmContextReleasedData {
	this := SmContextReleasedData{}
	return &this
}

// GetSmallDataRateStatus returns the SmallDataRateStatus field value if set, zero value otherwise.
func (o *SmContextReleasedData) GetSmallDataRateStatus() SmallDataRateStatus {
	if o == nil || isNil(o.SmallDataRateStatus) {
		var ret SmallDataRateStatus
		return ret
	}
	return *o.SmallDataRateStatus
}

// GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleasedData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool) {
	if o == nil || isNil(o.SmallDataRateStatus) {
    return nil, false
	}
	return o.SmallDataRateStatus, true
}

// HasSmallDataRateStatus returns a boolean if a field has been set.
func (o *SmContextReleasedData) HasSmallDataRateStatus() bool {
	if o != nil && !isNil(o.SmallDataRateStatus) {
		return true
	}

	return false
}

// SetSmallDataRateStatus gets a reference to the given SmallDataRateStatus and assigns it to the SmallDataRateStatus field.
func (o *SmContextReleasedData) SetSmallDataRateStatus(v SmallDataRateStatus) {
	o.SmallDataRateStatus = &v
}

// GetApnRateStatus returns the ApnRateStatus field value if set, zero value otherwise.
func (o *SmContextReleasedData) GetApnRateStatus() ApnRateStatus {
	if o == nil || isNil(o.ApnRateStatus) {
		var ret ApnRateStatus
		return ret
	}
	return *o.ApnRateStatus
}

// GetApnRateStatusOk returns a tuple with the ApnRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleasedData) GetApnRateStatusOk() (*ApnRateStatus, bool) {
	if o == nil || isNil(o.ApnRateStatus) {
    return nil, false
	}
	return o.ApnRateStatus, true
}

// HasApnRateStatus returns a boolean if a field has been set.
func (o *SmContextReleasedData) HasApnRateStatus() bool {
	if o != nil && !isNil(o.ApnRateStatus) {
		return true
	}

	return false
}

// SetApnRateStatus gets a reference to the given ApnRateStatus and assigns it to the ApnRateStatus field.
func (o *SmContextReleasedData) SetApnRateStatus(v ApnRateStatus) {
	o.ApnRateStatus = &v
}

func (o SmContextReleasedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SmallDataRateStatus) {
		toSerialize["smallDataRateStatus"] = o.SmallDataRateStatus
	}
	if !isNil(o.ApnRateStatus) {
		toSerialize["apnRateStatus"] = o.ApnRateStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextReleasedData struct {
	value *SmContextReleasedData
	isSet bool
}

func (v NullableSmContextReleasedData) Get() *SmContextReleasedData {
	return v.value
}

func (v *NullableSmContextReleasedData) Set(val *SmContextReleasedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextReleasedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextReleasedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextReleasedData(val *SmContextReleasedData) *NullableSmContextReleasedData {
	return &NullableSmContextReleasedData{value: val, isSet: true}
}

func (v NullableSmContextReleasedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextReleasedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


