/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// NefInfo Information of an NEF NF Instance
type NefInfo struct {
	// Identity of the NEF
	NefId *string `json:"nefId,omitempty"`
	PfdData *PfdData `json:"pfdData,omitempty"`
	AfEeData *AfEventExposureData `json:"afEeData,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	ServedFqdnList []string `json:"servedFqdnList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	DnaiList []string `json:"dnaiList,omitempty"`
	UnTrustAfInfoList []UnTrustAfInfo `json:"unTrustAfInfoList,omitempty"`
	UasNfFunctionalityInd *bool `json:"uasNfFunctionalityInd,omitempty"`
}

// NewNefInfo instantiates a new NefInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefInfo() *NefInfo {
	this := NefInfo{}
	var uasNfFunctionalityInd bool = false
	this.UasNfFunctionalityInd = &uasNfFunctionalityInd
	return &this
}

// NewNefInfoWithDefaults instantiates a new NefInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefInfoWithDefaults() *NefInfo {
	this := NefInfo{}
	var uasNfFunctionalityInd bool = false
	this.UasNfFunctionalityInd = &uasNfFunctionalityInd
	return &this
}

// GetNefId returns the NefId field value if set, zero value otherwise.
func (o *NefInfo) GetNefId() string {
	if o == nil || isNil(o.NefId) {
		var ret string
		return ret
	}
	return *o.NefId
}

// GetNefIdOk returns a tuple with the NefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetNefIdOk() (*string, bool) {
	if o == nil || isNil(o.NefId) {
    return nil, false
	}
	return o.NefId, true
}

// HasNefId returns a boolean if a field has been set.
func (o *NefInfo) HasNefId() bool {
	if o != nil && !isNil(o.NefId) {
		return true
	}

	return false
}

// SetNefId gets a reference to the given string and assigns it to the NefId field.
func (o *NefInfo) SetNefId(v string) {
	o.NefId = &v
}

// GetPfdData returns the PfdData field value if set, zero value otherwise.
func (o *NefInfo) GetPfdData() PfdData {
	if o == nil || isNil(o.PfdData) {
		var ret PfdData
		return ret
	}
	return *o.PfdData
}

// GetPfdDataOk returns a tuple with the PfdData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetPfdDataOk() (*PfdData, bool) {
	if o == nil || isNil(o.PfdData) {
    return nil, false
	}
	return o.PfdData, true
}

// HasPfdData returns a boolean if a field has been set.
func (o *NefInfo) HasPfdData() bool {
	if o != nil && !isNil(o.PfdData) {
		return true
	}

	return false
}

// SetPfdData gets a reference to the given PfdData and assigns it to the PfdData field.
func (o *NefInfo) SetPfdData(v PfdData) {
	o.PfdData = &v
}

// GetAfEeData returns the AfEeData field value if set, zero value otherwise.
func (o *NefInfo) GetAfEeData() AfEventExposureData {
	if o == nil || isNil(o.AfEeData) {
		var ret AfEventExposureData
		return ret
	}
	return *o.AfEeData
}

// GetAfEeDataOk returns a tuple with the AfEeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetAfEeDataOk() (*AfEventExposureData, bool) {
	if o == nil || isNil(o.AfEeData) {
    return nil, false
	}
	return o.AfEeData, true
}

// HasAfEeData returns a boolean if a field has been set.
func (o *NefInfo) HasAfEeData() bool {
	if o != nil && !isNil(o.AfEeData) {
		return true
	}

	return false
}

// SetAfEeData gets a reference to the given AfEventExposureData and assigns it to the AfEeData field.
func (o *NefInfo) SetAfEeData(v AfEventExposureData) {
	o.AfEeData = &v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *NefInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || isNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.GpsiRanges) {
    return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *NefInfo) HasGpsiRanges() bool {
	if o != nil && !isNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *NefInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *NefInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || isNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.ExternalGroupIdentifiersRanges) {
    return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *NefInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !isNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *NefInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetServedFqdnList returns the ServedFqdnList field value if set, zero value otherwise.
func (o *NefInfo) GetServedFqdnList() []string {
	if o == nil || isNil(o.ServedFqdnList) {
		var ret []string
		return ret
	}
	return o.ServedFqdnList
}

// GetServedFqdnListOk returns a tuple with the ServedFqdnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetServedFqdnListOk() ([]string, bool) {
	if o == nil || isNil(o.ServedFqdnList) {
    return nil, false
	}
	return o.ServedFqdnList, true
}

// HasServedFqdnList returns a boolean if a field has been set.
func (o *NefInfo) HasServedFqdnList() bool {
	if o != nil && !isNil(o.ServedFqdnList) {
		return true
	}

	return false
}

// SetServedFqdnList gets a reference to the given []string and assigns it to the ServedFqdnList field.
func (o *NefInfo) SetServedFqdnList(v []string) {
	o.ServedFqdnList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NefInfo) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NefInfo) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NefInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NefInfo) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NefInfo) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NefInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *NefInfo) GetDnaiList() []string {
	if o == nil || isNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetDnaiListOk() ([]string, bool) {
	if o == nil || isNil(o.DnaiList) {
    return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *NefInfo) HasDnaiList() bool {
	if o != nil && !isNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *NefInfo) SetDnaiList(v []string) {
	o.DnaiList = v
}

// GetUnTrustAfInfoList returns the UnTrustAfInfoList field value if set, zero value otherwise.
func (o *NefInfo) GetUnTrustAfInfoList() []UnTrustAfInfo {
	if o == nil || isNil(o.UnTrustAfInfoList) {
		var ret []UnTrustAfInfo
		return ret
	}
	return o.UnTrustAfInfoList
}

// GetUnTrustAfInfoListOk returns a tuple with the UnTrustAfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetUnTrustAfInfoListOk() ([]UnTrustAfInfo, bool) {
	if o == nil || isNil(o.UnTrustAfInfoList) {
    return nil, false
	}
	return o.UnTrustAfInfoList, true
}

// HasUnTrustAfInfoList returns a boolean if a field has been set.
func (o *NefInfo) HasUnTrustAfInfoList() bool {
	if o != nil && !isNil(o.UnTrustAfInfoList) {
		return true
	}

	return false
}

// SetUnTrustAfInfoList gets a reference to the given []UnTrustAfInfo and assigns it to the UnTrustAfInfoList field.
func (o *NefInfo) SetUnTrustAfInfoList(v []UnTrustAfInfo) {
	o.UnTrustAfInfoList = v
}

// GetUasNfFunctionalityInd returns the UasNfFunctionalityInd field value if set, zero value otherwise.
func (o *NefInfo) GetUasNfFunctionalityInd() bool {
	if o == nil || isNil(o.UasNfFunctionalityInd) {
		var ret bool
		return ret
	}
	return *o.UasNfFunctionalityInd
}

// GetUasNfFunctionalityIndOk returns a tuple with the UasNfFunctionalityInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefInfo) GetUasNfFunctionalityIndOk() (*bool, bool) {
	if o == nil || isNil(o.UasNfFunctionalityInd) {
    return nil, false
	}
	return o.UasNfFunctionalityInd, true
}

// HasUasNfFunctionalityInd returns a boolean if a field has been set.
func (o *NefInfo) HasUasNfFunctionalityInd() bool {
	if o != nil && !isNil(o.UasNfFunctionalityInd) {
		return true
	}

	return false
}

// SetUasNfFunctionalityInd gets a reference to the given bool and assigns it to the UasNfFunctionalityInd field.
func (o *NefInfo) SetUasNfFunctionalityInd(v bool) {
	o.UasNfFunctionalityInd = &v
}

func (o NefInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NefId) {
		toSerialize["nefId"] = o.NefId
	}
	if !isNil(o.PfdData) {
		toSerialize["pfdData"] = o.PfdData
	}
	if !isNil(o.AfEeData) {
		toSerialize["afEeData"] = o.AfEeData
	}
	if !isNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !isNil(o.ExternalGroupIdentifiersRanges) {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if !isNil(o.ServedFqdnList) {
		toSerialize["servedFqdnList"] = o.ServedFqdnList
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if !isNil(o.UnTrustAfInfoList) {
		toSerialize["unTrustAfInfoList"] = o.UnTrustAfInfoList
	}
	if !isNil(o.UasNfFunctionalityInd) {
		toSerialize["uasNfFunctionalityInd"] = o.UasNfFunctionalityInd
	}
	return json.Marshal(toSerialize)
}

type NullableNefInfo struct {
	value *NefInfo
	isSet bool
}

func (v NullableNefInfo) Get() *NefInfo {
	return v.value
}

func (v *NullableNefInfo) Set(val *NefInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNefInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNefInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefInfo(val *NefInfo) *NullableNefInfo {
	return &NullableNefInfo{value: val, isSet: true}
}

func (v NullableNefInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


