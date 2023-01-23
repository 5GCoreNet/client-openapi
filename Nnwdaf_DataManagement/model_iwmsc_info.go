/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// IwmscInfo Information of an SMS-IWMSC NF Instance
type IwmscInfo struct {
	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	ScNumber *string `json:"scNumber,omitempty"`
}

// NewIwmscInfo instantiates a new IwmscInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwmscInfo() *IwmscInfo {
	this := IwmscInfo{}
	return &this
}

// NewIwmscInfoWithDefaults instantiates a new IwmscInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwmscInfoWithDefaults() *IwmscInfo {
	this := IwmscInfo{}
	return &this
}

// GetMsisdnRanges returns the MsisdnRanges field value if set, zero value otherwise.
func (o *IwmscInfo) GetMsisdnRanges() []IdentityRange {
	if o == nil || isNil(o.MsisdnRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.MsisdnRanges
}

// GetMsisdnRangesOk returns a tuple with the MsisdnRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwmscInfo) GetMsisdnRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.MsisdnRanges) {
    return nil, false
	}
	return o.MsisdnRanges, true
}

// HasMsisdnRanges returns a boolean if a field has been set.
func (o *IwmscInfo) HasMsisdnRanges() bool {
	if o != nil && !isNil(o.MsisdnRanges) {
		return true
	}

	return false
}

// SetMsisdnRanges gets a reference to the given []IdentityRange and assigns it to the MsisdnRanges field.
func (o *IwmscInfo) SetMsisdnRanges(v []IdentityRange) {
	o.MsisdnRanges = v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *IwmscInfo) GetSupiRanges() []SupiRange {
	if o == nil || isNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwmscInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || isNil(o.SupiRanges) {
    return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *IwmscInfo) HasSupiRanges() bool {
	if o != nil && !isNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *IwmscInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *IwmscInfo) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwmscInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *IwmscInfo) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *IwmscInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetScNumber returns the ScNumber field value if set, zero value otherwise.
func (o *IwmscInfo) GetScNumber() string {
	if o == nil || isNil(o.ScNumber) {
		var ret string
		return ret
	}
	return *o.ScNumber
}

// GetScNumberOk returns a tuple with the ScNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwmscInfo) GetScNumberOk() (*string, bool) {
	if o == nil || isNil(o.ScNumber) {
    return nil, false
	}
	return o.ScNumber, true
}

// HasScNumber returns a boolean if a field has been set.
func (o *IwmscInfo) HasScNumber() bool {
	if o != nil && !isNil(o.ScNumber) {
		return true
	}

	return false
}

// SetScNumber gets a reference to the given string and assigns it to the ScNumber field.
func (o *IwmscInfo) SetScNumber(v string) {
	o.ScNumber = &v
}

func (o IwmscInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MsisdnRanges) {
		toSerialize["msisdnRanges"] = o.MsisdnRanges
	}
	if !isNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.ScNumber) {
		toSerialize["scNumber"] = o.ScNumber
	}
	return json.Marshal(toSerialize)
}

type NullableIwmscInfo struct {
	value *IwmscInfo
	isSet bool
}

func (v NullableIwmscInfo) Get() *IwmscInfo {
	return v.value
}

func (v *NullableIwmscInfo) Set(val *IwmscInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIwmscInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIwmscInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwmscInfo(val *IwmscInfo) *NullableIwmscInfo {
	return &NullableIwmscInfo{value: val, isSet: true}
}

func (v NullableIwmscInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwmscInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


