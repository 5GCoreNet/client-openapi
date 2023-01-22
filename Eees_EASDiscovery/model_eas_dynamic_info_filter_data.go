/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EASDiscovery

import (
	"encoding/json"
)

// EasDynamicInfoFilterData Represents an EAS dynamic information.
type EasDynamicInfoFilterData struct {
	// Represents a unique identifier of the EEC.
	EecId string `json:"eecId"`
	// Notify if EAS status changed.
	EasStatus *bool `json:"easStatus,omitempty"`
	// Notify if list of AC identifiers changed.
	EasAcIds *bool `json:"easAcIds,omitempty"`
	// Notify if EAS description changed.
	EasDesc *bool `json:"easDesc,omitempty"`
	// Notify if EAS endpoint changed.
	EasPt *bool `json:"easPt,omitempty"`
	// NotiNotify if EAS feature changed.
	EasFeature *bool `json:"easFeature,omitempty"`
	// Notify if EAS schedule changed.
	EasSchedule *bool `json:"easSchedule,omitempty"`
	// Notify if EAS service area changed.
	SvcArea *bool `json:"svcArea,omitempty"`
	// Notify if EAS KPIs changed.
	SvcKpi *bool `json:"svcKpi,omitempty"`
	// Notify if EAS supported ACR changed.
	SvcCont *bool `json:"svcCont,omitempty"`
}

// NewEasDynamicInfoFilterData instantiates a new EasDynamicInfoFilterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDynamicInfoFilterData(eecId string) *EasDynamicInfoFilterData {
	this := EasDynamicInfoFilterData{}
	this.EecId = eecId
	return &this
}

// NewEasDynamicInfoFilterDataWithDefaults instantiates a new EasDynamicInfoFilterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDynamicInfoFilterDataWithDefaults() *EasDynamicInfoFilterData {
	this := EasDynamicInfoFilterData{}
	return &this
}

// GetEecId returns the EecId field value
func (o *EasDynamicInfoFilterData) GetEecId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EecId
}

// GetEecIdOk returns a tuple with the EecId field value
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEecIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EecId, true
}

// SetEecId sets field value
func (o *EasDynamicInfoFilterData) SetEecId(v string) {
	o.EecId = v
}

// GetEasStatus returns the EasStatus field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasStatus() bool {
	if o == nil || isNil(o.EasStatus) {
		var ret bool
		return ret
	}
	return *o.EasStatus
}

// GetEasStatusOk returns a tuple with the EasStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasStatusOk() (*bool, bool) {
	if o == nil || isNil(o.EasStatus) {
    return nil, false
	}
	return o.EasStatus, true
}

// HasEasStatus returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasStatus() bool {
	if o != nil && !isNil(o.EasStatus) {
		return true
	}

	return false
}

// SetEasStatus gets a reference to the given bool and assigns it to the EasStatus field.
func (o *EasDynamicInfoFilterData) SetEasStatus(v bool) {
	o.EasStatus = &v
}

// GetEasAcIds returns the EasAcIds field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasAcIds() bool {
	if o == nil || isNil(o.EasAcIds) {
		var ret bool
		return ret
	}
	return *o.EasAcIds
}

// GetEasAcIdsOk returns a tuple with the EasAcIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasAcIdsOk() (*bool, bool) {
	if o == nil || isNil(o.EasAcIds) {
    return nil, false
	}
	return o.EasAcIds, true
}

// HasEasAcIds returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasAcIds() bool {
	if o != nil && !isNil(o.EasAcIds) {
		return true
	}

	return false
}

// SetEasAcIds gets a reference to the given bool and assigns it to the EasAcIds field.
func (o *EasDynamicInfoFilterData) SetEasAcIds(v bool) {
	o.EasAcIds = &v
}

// GetEasDesc returns the EasDesc field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasDesc() bool {
	if o == nil || isNil(o.EasDesc) {
		var ret bool
		return ret
	}
	return *o.EasDesc
}

// GetEasDescOk returns a tuple with the EasDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasDescOk() (*bool, bool) {
	if o == nil || isNil(o.EasDesc) {
    return nil, false
	}
	return o.EasDesc, true
}

// HasEasDesc returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasDesc() bool {
	if o != nil && !isNil(o.EasDesc) {
		return true
	}

	return false
}

// SetEasDesc gets a reference to the given bool and assigns it to the EasDesc field.
func (o *EasDynamicInfoFilterData) SetEasDesc(v bool) {
	o.EasDesc = &v
}

// GetEasPt returns the EasPt field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasPt() bool {
	if o == nil || isNil(o.EasPt) {
		var ret bool
		return ret
	}
	return *o.EasPt
}

// GetEasPtOk returns a tuple with the EasPt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasPtOk() (*bool, bool) {
	if o == nil || isNil(o.EasPt) {
    return nil, false
	}
	return o.EasPt, true
}

// HasEasPt returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasPt() bool {
	if o != nil && !isNil(o.EasPt) {
		return true
	}

	return false
}

// SetEasPt gets a reference to the given bool and assigns it to the EasPt field.
func (o *EasDynamicInfoFilterData) SetEasPt(v bool) {
	o.EasPt = &v
}

// GetEasFeature returns the EasFeature field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasFeature() bool {
	if o == nil || isNil(o.EasFeature) {
		var ret bool
		return ret
	}
	return *o.EasFeature
}

// GetEasFeatureOk returns a tuple with the EasFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasFeatureOk() (*bool, bool) {
	if o == nil || isNil(o.EasFeature) {
    return nil, false
	}
	return o.EasFeature, true
}

// HasEasFeature returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasFeature() bool {
	if o != nil && !isNil(o.EasFeature) {
		return true
	}

	return false
}

// SetEasFeature gets a reference to the given bool and assigns it to the EasFeature field.
func (o *EasDynamicInfoFilterData) SetEasFeature(v bool) {
	o.EasFeature = &v
}

// GetEasSchedule returns the EasSchedule field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetEasSchedule() bool {
	if o == nil || isNil(o.EasSchedule) {
		var ret bool
		return ret
	}
	return *o.EasSchedule
}

// GetEasScheduleOk returns a tuple with the EasSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetEasScheduleOk() (*bool, bool) {
	if o == nil || isNil(o.EasSchedule) {
    return nil, false
	}
	return o.EasSchedule, true
}

// HasEasSchedule returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasEasSchedule() bool {
	if o != nil && !isNil(o.EasSchedule) {
		return true
	}

	return false
}

// SetEasSchedule gets a reference to the given bool and assigns it to the EasSchedule field.
func (o *EasDynamicInfoFilterData) SetEasSchedule(v bool) {
	o.EasSchedule = &v
}

// GetSvcArea returns the SvcArea field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetSvcArea() bool {
	if o == nil || isNil(o.SvcArea) {
		var ret bool
		return ret
	}
	return *o.SvcArea
}

// GetSvcAreaOk returns a tuple with the SvcArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetSvcAreaOk() (*bool, bool) {
	if o == nil || isNil(o.SvcArea) {
    return nil, false
	}
	return o.SvcArea, true
}

// HasSvcArea returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasSvcArea() bool {
	if o != nil && !isNil(o.SvcArea) {
		return true
	}

	return false
}

// SetSvcArea gets a reference to the given bool and assigns it to the SvcArea field.
func (o *EasDynamicInfoFilterData) SetSvcArea(v bool) {
	o.SvcArea = &v
}

// GetSvcKpi returns the SvcKpi field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetSvcKpi() bool {
	if o == nil || isNil(o.SvcKpi) {
		var ret bool
		return ret
	}
	return *o.SvcKpi
}

// GetSvcKpiOk returns a tuple with the SvcKpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetSvcKpiOk() (*bool, bool) {
	if o == nil || isNil(o.SvcKpi) {
    return nil, false
	}
	return o.SvcKpi, true
}

// HasSvcKpi returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasSvcKpi() bool {
	if o != nil && !isNil(o.SvcKpi) {
		return true
	}

	return false
}

// SetSvcKpi gets a reference to the given bool and assigns it to the SvcKpi field.
func (o *EasDynamicInfoFilterData) SetSvcKpi(v bool) {
	o.SvcKpi = &v
}

// GetSvcCont returns the SvcCont field value if set, zero value otherwise.
func (o *EasDynamicInfoFilterData) GetSvcCont() bool {
	if o == nil || isNil(o.SvcCont) {
		var ret bool
		return ret
	}
	return *o.SvcCont
}

// GetSvcContOk returns a tuple with the SvcCont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDynamicInfoFilterData) GetSvcContOk() (*bool, bool) {
	if o == nil || isNil(o.SvcCont) {
    return nil, false
	}
	return o.SvcCont, true
}

// HasSvcCont returns a boolean if a field has been set.
func (o *EasDynamicInfoFilterData) HasSvcCont() bool {
	if o != nil && !isNil(o.SvcCont) {
		return true
	}

	return false
}

// SetSvcCont gets a reference to the given bool and assigns it to the SvcCont field.
func (o *EasDynamicInfoFilterData) SetSvcCont(v bool) {
	o.SvcCont = &v
}

func (o EasDynamicInfoFilterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eecId"] = o.EecId
	}
	if !isNil(o.EasStatus) {
		toSerialize["easStatus"] = o.EasStatus
	}
	if !isNil(o.EasAcIds) {
		toSerialize["easAcIds"] = o.EasAcIds
	}
	if !isNil(o.EasDesc) {
		toSerialize["easDesc"] = o.EasDesc
	}
	if !isNil(o.EasPt) {
		toSerialize["easPt"] = o.EasPt
	}
	if !isNil(o.EasFeature) {
		toSerialize["easFeature"] = o.EasFeature
	}
	if !isNil(o.EasSchedule) {
		toSerialize["easSchedule"] = o.EasSchedule
	}
	if !isNil(o.SvcArea) {
		toSerialize["svcArea"] = o.SvcArea
	}
	if !isNil(o.SvcKpi) {
		toSerialize["svcKpi"] = o.SvcKpi
	}
	if !isNil(o.SvcCont) {
		toSerialize["svcCont"] = o.SvcCont
	}
	return json.Marshal(toSerialize)
}

type NullableEasDynamicInfoFilterData struct {
	value *EasDynamicInfoFilterData
	isSet bool
}

func (v NullableEasDynamicInfoFilterData) Get() *EasDynamicInfoFilterData {
	return v.value
}

func (v *NullableEasDynamicInfoFilterData) Set(val *EasDynamicInfoFilterData) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDynamicInfoFilterData) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDynamicInfoFilterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDynamicInfoFilterData(val *EasDynamicInfoFilterData) *NullableEasDynamicInfoFilterData {
	return &NullableEasDynamicInfoFilterData{value: val, isSet: true}
}

func (v NullableEasDynamicInfoFilterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDynamicInfoFilterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

