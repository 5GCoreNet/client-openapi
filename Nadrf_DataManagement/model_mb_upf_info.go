/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
)

// MbUpfInfo Information of an MB-UPF NF Instance
type MbUpfInfo struct {
	SNssaiMbUpfInfoList []SnssaiUpfInfoItem `json:"sNssaiMbUpfInfoList"`
	MbSmfServingArea []string `json:"mbSmfServingArea,omitempty"`
	InterfaceMbUpfInfoList []InterfaceUpfInfoItem `json:"interfaceMbUpfInfoList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	SupportedPfcpFeatures *string `json:"supportedPfcpFeatures,omitempty"`
}

// NewMbUpfInfo instantiates a new MbUpfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbUpfInfo(sNssaiMbUpfInfoList []SnssaiUpfInfoItem) *MbUpfInfo {
	this := MbUpfInfo{}
	this.SNssaiMbUpfInfoList = sNssaiMbUpfInfoList
	return &this
}

// NewMbUpfInfoWithDefaults instantiates a new MbUpfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbUpfInfoWithDefaults() *MbUpfInfo {
	this := MbUpfInfo{}
	return &this
}

// GetSNssaiMbUpfInfoList returns the SNssaiMbUpfInfoList field value
func (o *MbUpfInfo) GetSNssaiMbUpfInfoList() []SnssaiUpfInfoItem {
	if o == nil {
		var ret []SnssaiUpfInfoItem
		return ret
	}

	return o.SNssaiMbUpfInfoList
}

// GetSNssaiMbUpfInfoListOk returns a tuple with the SNssaiMbUpfInfoList field value
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetSNssaiMbUpfInfoListOk() ([]SnssaiUpfInfoItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.SNssaiMbUpfInfoList, true
}

// SetSNssaiMbUpfInfoList sets field value
func (o *MbUpfInfo) SetSNssaiMbUpfInfoList(v []SnssaiUpfInfoItem) {
	o.SNssaiMbUpfInfoList = v
}

// GetMbSmfServingArea returns the MbSmfServingArea field value if set, zero value otherwise.
func (o *MbUpfInfo) GetMbSmfServingArea() []string {
	if o == nil || isNil(o.MbSmfServingArea) {
		var ret []string
		return ret
	}
	return o.MbSmfServingArea
}

// GetMbSmfServingAreaOk returns a tuple with the MbSmfServingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetMbSmfServingAreaOk() ([]string, bool) {
	if o == nil || isNil(o.MbSmfServingArea) {
    return nil, false
	}
	return o.MbSmfServingArea, true
}

// HasMbSmfServingArea returns a boolean if a field has been set.
func (o *MbUpfInfo) HasMbSmfServingArea() bool {
	if o != nil && !isNil(o.MbSmfServingArea) {
		return true
	}

	return false
}

// SetMbSmfServingArea gets a reference to the given []string and assigns it to the MbSmfServingArea field.
func (o *MbUpfInfo) SetMbSmfServingArea(v []string) {
	o.MbSmfServingArea = v
}

// GetInterfaceMbUpfInfoList returns the InterfaceMbUpfInfoList field value if set, zero value otherwise.
func (o *MbUpfInfo) GetInterfaceMbUpfInfoList() []InterfaceUpfInfoItem {
	if o == nil || isNil(o.InterfaceMbUpfInfoList) {
		var ret []InterfaceUpfInfoItem
		return ret
	}
	return o.InterfaceMbUpfInfoList
}

// GetInterfaceMbUpfInfoListOk returns a tuple with the InterfaceMbUpfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetInterfaceMbUpfInfoListOk() ([]InterfaceUpfInfoItem, bool) {
	if o == nil || isNil(o.InterfaceMbUpfInfoList) {
    return nil, false
	}
	return o.InterfaceMbUpfInfoList, true
}

// HasInterfaceMbUpfInfoList returns a boolean if a field has been set.
func (o *MbUpfInfo) HasInterfaceMbUpfInfoList() bool {
	if o != nil && !isNil(o.InterfaceMbUpfInfoList) {
		return true
	}

	return false
}

// SetInterfaceMbUpfInfoList gets a reference to the given []InterfaceUpfInfoItem and assigns it to the InterfaceMbUpfInfoList field.
func (o *MbUpfInfo) SetInterfaceMbUpfInfoList(v []InterfaceUpfInfoItem) {
	o.InterfaceMbUpfInfoList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *MbUpfInfo) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *MbUpfInfo) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *MbUpfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *MbUpfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *MbUpfInfo) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *MbUpfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *MbUpfInfo) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *MbUpfInfo) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *MbUpfInfo) SetPriority(v int32) {
	o.Priority = &v
}

// GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field value if set, zero value otherwise.
func (o *MbUpfInfo) GetSupportedPfcpFeatures() string {
	if o == nil || isNil(o.SupportedPfcpFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedPfcpFeatures
}

// GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbUpfInfo) GetSupportedPfcpFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedPfcpFeatures) {
    return nil, false
	}
	return o.SupportedPfcpFeatures, true
}

// HasSupportedPfcpFeatures returns a boolean if a field has been set.
func (o *MbUpfInfo) HasSupportedPfcpFeatures() bool {
	if o != nil && !isNil(o.SupportedPfcpFeatures) {
		return true
	}

	return false
}

// SetSupportedPfcpFeatures gets a reference to the given string and assigns it to the SupportedPfcpFeatures field.
func (o *MbUpfInfo) SetSupportedPfcpFeatures(v string) {
	o.SupportedPfcpFeatures = &v
}

func (o MbUpfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssaiMbUpfInfoList"] = o.SNssaiMbUpfInfoList
	}
	if !isNil(o.MbSmfServingArea) {
		toSerialize["mbSmfServingArea"] = o.MbSmfServingArea
	}
	if !isNil(o.InterfaceMbUpfInfoList) {
		toSerialize["interfaceMbUpfInfoList"] = o.InterfaceMbUpfInfoList
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.SupportedPfcpFeatures) {
		toSerialize["supportedPfcpFeatures"] = o.SupportedPfcpFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableMbUpfInfo struct {
	value *MbUpfInfo
	isSet bool
}

func (v NullableMbUpfInfo) Get() *MbUpfInfo {
	return v.value
}

func (v *NullableMbUpfInfo) Set(val *MbUpfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbUpfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbUpfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbUpfInfo(val *MbUpfInfo) *NullableMbUpfInfo {
	return &NullableMbUpfInfo{value: val, isSet: true}
}

func (v NullableMbUpfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbUpfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


