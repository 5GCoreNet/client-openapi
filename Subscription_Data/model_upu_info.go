/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"time"
)

// UpuInfo struct for UpuInfo
type UpuInfo struct {
	UpuDataList []UpuData1 `json:"upuDataList,omitempty"`
	UpuRegInd *bool `json:"upuRegInd,omitempty"`
	// Contains the indication of whether the acknowledgement from UE is needed.
	UpuAckInd *bool `json:"upuAckInd,omitempty"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIausf *string `json:"upuMacIausf,omitempty"`
	// CounterUPU.
	CounterUpu *string `json:"counterUpu,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	// string with format 'bytes' as defined in OpenAPI
	UpuTransparentContainer *string `json:"upuTransparentContainer,omitempty"`
}

// NewUpuInfo instantiates a new UpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpuInfo(provisioningTime time.Time) *UpuInfo {
	this := UpuInfo{}
	this.ProvisioningTime = provisioningTime
	return &this
}

// NewUpuInfoWithDefaults instantiates a new UpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpuInfoWithDefaults() *UpuInfo {
	this := UpuInfo{}
	return &this
}

// GetUpuDataList returns the UpuDataList field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuDataList() []UpuData1 {
	if o == nil || isNil(o.UpuDataList) {
		var ret []UpuData1
		return ret
	}
	return o.UpuDataList
}

// GetUpuDataListOk returns a tuple with the UpuDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuDataListOk() ([]UpuData1, bool) {
	if o == nil || isNil(o.UpuDataList) {
    return nil, false
	}
	return o.UpuDataList, true
}

// HasUpuDataList returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuDataList() bool {
	if o != nil && !isNil(o.UpuDataList) {
		return true
	}

	return false
}

// SetUpuDataList gets a reference to the given []UpuData1 and assigns it to the UpuDataList field.
func (o *UpuInfo) SetUpuDataList(v []UpuData1) {
	o.UpuDataList = v
}

// GetUpuRegInd returns the UpuRegInd field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuRegInd() bool {
	if o == nil || isNil(o.UpuRegInd) {
		var ret bool
		return ret
	}
	return *o.UpuRegInd
}

// GetUpuRegIndOk returns a tuple with the UpuRegInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuRegIndOk() (*bool, bool) {
	if o == nil || isNil(o.UpuRegInd) {
    return nil, false
	}
	return o.UpuRegInd, true
}

// HasUpuRegInd returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuRegInd() bool {
	if o != nil && !isNil(o.UpuRegInd) {
		return true
	}

	return false
}

// SetUpuRegInd gets a reference to the given bool and assigns it to the UpuRegInd field.
func (o *UpuInfo) SetUpuRegInd(v bool) {
	o.UpuRegInd = &v
}

// GetUpuAckInd returns the UpuAckInd field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuAckInd() bool {
	if o == nil || isNil(o.UpuAckInd) {
		var ret bool
		return ret
	}
	return *o.UpuAckInd
}

// GetUpuAckIndOk returns a tuple with the UpuAckInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuAckIndOk() (*bool, bool) {
	if o == nil || isNil(o.UpuAckInd) {
    return nil, false
	}
	return o.UpuAckInd, true
}

// HasUpuAckInd returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuAckInd() bool {
	if o != nil && !isNil(o.UpuAckInd) {
		return true
	}

	return false
}

// SetUpuAckInd gets a reference to the given bool and assigns it to the UpuAckInd field.
func (o *UpuInfo) SetUpuAckInd(v bool) {
	o.UpuAckInd = &v
}

// GetUpuMacIausf returns the UpuMacIausf field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuMacIausf() string {
	if o == nil || isNil(o.UpuMacIausf) {
		var ret string
		return ret
	}
	return *o.UpuMacIausf
}

// GetUpuMacIausfOk returns a tuple with the UpuMacIausf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuMacIausfOk() (*string, bool) {
	if o == nil || isNil(o.UpuMacIausf) {
    return nil, false
	}
	return o.UpuMacIausf, true
}

// HasUpuMacIausf returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuMacIausf() bool {
	if o != nil && !isNil(o.UpuMacIausf) {
		return true
	}

	return false
}

// SetUpuMacIausf gets a reference to the given string and assigns it to the UpuMacIausf field.
func (o *UpuInfo) SetUpuMacIausf(v string) {
	o.UpuMacIausf = &v
}

// GetCounterUpu returns the CounterUpu field value if set, zero value otherwise.
func (o *UpuInfo) GetCounterUpu() string {
	if o == nil || isNil(o.CounterUpu) {
		var ret string
		return ret
	}
	return *o.CounterUpu
}

// GetCounterUpuOk returns a tuple with the CounterUpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetCounterUpuOk() (*string, bool) {
	if o == nil || isNil(o.CounterUpu) {
    return nil, false
	}
	return o.CounterUpu, true
}

// HasCounterUpu returns a boolean if a field has been set.
func (o *UpuInfo) HasCounterUpu() bool {
	if o != nil && !isNil(o.CounterUpu) {
		return true
	}

	return false
}

// SetCounterUpu gets a reference to the given string and assigns it to the CounterUpu field.
func (o *UpuInfo) SetCounterUpu(v string) {
	o.CounterUpu = &v
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *UpuInfo) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *UpuInfo) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetUpuTransparentContainer returns the UpuTransparentContainer field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuTransparentContainer() string {
	if o == nil || isNil(o.UpuTransparentContainer) {
		var ret string
		return ret
	}
	return *o.UpuTransparentContainer
}

// GetUpuTransparentContainerOk returns a tuple with the UpuTransparentContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuTransparentContainerOk() (*string, bool) {
	if o == nil || isNil(o.UpuTransparentContainer) {
    return nil, false
	}
	return o.UpuTransparentContainer, true
}

// HasUpuTransparentContainer returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuTransparentContainer() bool {
	if o != nil && !isNil(o.UpuTransparentContainer) {
		return true
	}

	return false
}

// SetUpuTransparentContainer gets a reference to the given string and assigns it to the UpuTransparentContainer field.
func (o *UpuInfo) SetUpuTransparentContainer(v string) {
	o.UpuTransparentContainer = &v
}

func (o UpuInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpuDataList) {
		toSerialize["upuDataList"] = o.UpuDataList
	}
	if !isNil(o.UpuRegInd) {
		toSerialize["upuRegInd"] = o.UpuRegInd
	}
	if !isNil(o.UpuAckInd) {
		toSerialize["upuAckInd"] = o.UpuAckInd
	}
	if !isNil(o.UpuMacIausf) {
		toSerialize["upuMacIausf"] = o.UpuMacIausf
	}
	if !isNil(o.CounterUpu) {
		toSerialize["counterUpu"] = o.CounterUpu
	}
	if true {
		toSerialize["provisioningTime"] = o.ProvisioningTime
	}
	if !isNil(o.UpuTransparentContainer) {
		toSerialize["upuTransparentContainer"] = o.UpuTransparentContainer
	}
	return json.Marshal(toSerialize)
}

type NullableUpuInfo struct {
	value *UpuInfo
	isSet bool
}

func (v NullableUpuInfo) Get() *UpuInfo {
	return v.value
}

func (v *NullableUpuInfo) Set(val *UpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpuInfo(val *UpuInfo) *NullableUpuInfo {
	return &NullableUpuInfo{value: val, isSet: true}
}

func (v NullableUpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


