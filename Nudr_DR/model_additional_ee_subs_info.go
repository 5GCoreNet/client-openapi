/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// AdditionalEeSubsInfo struct for AdditionalEeSubsInfo
type AdditionalEeSubsInfo struct {
	AmfSubscriptionInfoList []AmfSubscriptionInfo `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo *SmfSubscriptionInfo `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
}

// NewAdditionalEeSubsInfo instantiates a new AdditionalEeSubsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalEeSubsInfo() *AdditionalEeSubsInfo {
	this := AdditionalEeSubsInfo{}
	return &this
}

// NewAdditionalEeSubsInfoWithDefaults instantiates a new AdditionalEeSubsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalEeSubsInfoWithDefaults() *AdditionalEeSubsInfo {
	this := AdditionalEeSubsInfo{}
	return &this
}

// GetAmfSubscriptionInfoList returns the AmfSubscriptionInfoList field value if set, zero value otherwise.
func (o *AdditionalEeSubsInfo) GetAmfSubscriptionInfoList() []AmfSubscriptionInfo {
	if o == nil || isNil(o.AmfSubscriptionInfoList) {
		var ret []AmfSubscriptionInfo
		return ret
	}
	return o.AmfSubscriptionInfoList
}

// GetAmfSubscriptionInfoListOk returns a tuple with the AmfSubscriptionInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalEeSubsInfo) GetAmfSubscriptionInfoListOk() ([]AmfSubscriptionInfo, bool) {
	if o == nil || isNil(o.AmfSubscriptionInfoList) {
    return nil, false
	}
	return o.AmfSubscriptionInfoList, true
}

// HasAmfSubscriptionInfoList returns a boolean if a field has been set.
func (o *AdditionalEeSubsInfo) HasAmfSubscriptionInfoList() bool {
	if o != nil && !isNil(o.AmfSubscriptionInfoList) {
		return true
	}

	return false
}

// SetAmfSubscriptionInfoList gets a reference to the given []AmfSubscriptionInfo and assigns it to the AmfSubscriptionInfoList field.
func (o *AdditionalEeSubsInfo) SetAmfSubscriptionInfoList(v []AmfSubscriptionInfo) {
	o.AmfSubscriptionInfoList = v
}

// GetSmfSubscriptionInfo returns the SmfSubscriptionInfo field value if set, zero value otherwise.
func (o *AdditionalEeSubsInfo) GetSmfSubscriptionInfo() SmfSubscriptionInfo {
	if o == nil || isNil(o.SmfSubscriptionInfo) {
		var ret SmfSubscriptionInfo
		return ret
	}
	return *o.SmfSubscriptionInfo
}

// GetSmfSubscriptionInfoOk returns a tuple with the SmfSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalEeSubsInfo) GetSmfSubscriptionInfoOk() (*SmfSubscriptionInfo, bool) {
	if o == nil || isNil(o.SmfSubscriptionInfo) {
    return nil, false
	}
	return o.SmfSubscriptionInfo, true
}

// HasSmfSubscriptionInfo returns a boolean if a field has been set.
func (o *AdditionalEeSubsInfo) HasSmfSubscriptionInfo() bool {
	if o != nil && !isNil(o.SmfSubscriptionInfo) {
		return true
	}

	return false
}

// SetSmfSubscriptionInfo gets a reference to the given SmfSubscriptionInfo and assigns it to the SmfSubscriptionInfo field.
func (o *AdditionalEeSubsInfo) SetSmfSubscriptionInfo(v SmfSubscriptionInfo) {
	o.SmfSubscriptionInfo = &v
}

// GetHssSubscriptionInfo returns the HssSubscriptionInfo field value if set, zero value otherwise.
func (o *AdditionalEeSubsInfo) GetHssSubscriptionInfo() HssSubscriptionInfo {
	if o == nil || isNil(o.HssSubscriptionInfo) {
		var ret HssSubscriptionInfo
		return ret
	}
	return *o.HssSubscriptionInfo
}

// GetHssSubscriptionInfoOk returns a tuple with the HssSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalEeSubsInfo) GetHssSubscriptionInfoOk() (*HssSubscriptionInfo, bool) {
	if o == nil || isNil(o.HssSubscriptionInfo) {
    return nil, false
	}
	return o.HssSubscriptionInfo, true
}

// HasHssSubscriptionInfo returns a boolean if a field has been set.
func (o *AdditionalEeSubsInfo) HasHssSubscriptionInfo() bool {
	if o != nil && !isNil(o.HssSubscriptionInfo) {
		return true
	}

	return false
}

// SetHssSubscriptionInfo gets a reference to the given HssSubscriptionInfo and assigns it to the HssSubscriptionInfo field.
func (o *AdditionalEeSubsInfo) SetHssSubscriptionInfo(v HssSubscriptionInfo) {
	o.HssSubscriptionInfo = &v
}

func (o AdditionalEeSubsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmfSubscriptionInfoList) {
		toSerialize["amfSubscriptionInfoList"] = o.AmfSubscriptionInfoList
	}
	if !isNil(o.SmfSubscriptionInfo) {
		toSerialize["smfSubscriptionInfo"] = o.SmfSubscriptionInfo
	}
	if !isNil(o.HssSubscriptionInfo) {
		toSerialize["hssSubscriptionInfo"] = o.HssSubscriptionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalEeSubsInfo struct {
	value *AdditionalEeSubsInfo
	isSet bool
}

func (v NullableAdditionalEeSubsInfo) Get() *AdditionalEeSubsInfo {
	return v.value
}

func (v *NullableAdditionalEeSubsInfo) Set(val *AdditionalEeSubsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalEeSubsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalEeSubsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalEeSubsInfo(val *AdditionalEeSubsInfo) *NullableAdditionalEeSubsInfo {
	return &NullableAdditionalEeSubsInfo{value: val, isSet: true}
}

func (v NullableAdditionalEeSubsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalEeSubsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


