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

// CmInfoReport struct for CmInfoReport
type CmInfoReport struct {
	OldCmInfoList []CmInfo `json:"oldCmInfoList,omitempty"`
	NewCmInfoList []CmInfo `json:"newCmInfoList"`
}

// NewCmInfoReport instantiates a new CmInfoReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmInfoReport(newCmInfoList []CmInfo) *CmInfoReport {
	this := CmInfoReport{}
	this.NewCmInfoList = newCmInfoList
	return &this
}

// NewCmInfoReportWithDefaults instantiates a new CmInfoReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmInfoReportWithDefaults() *CmInfoReport {
	this := CmInfoReport{}
	return &this
}

// GetOldCmInfoList returns the OldCmInfoList field value if set, zero value otherwise.
func (o *CmInfoReport) GetOldCmInfoList() []CmInfo {
	if o == nil || isNil(o.OldCmInfoList) {
		var ret []CmInfo
		return ret
	}
	return o.OldCmInfoList
}

// GetOldCmInfoListOk returns a tuple with the OldCmInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmInfoReport) GetOldCmInfoListOk() ([]CmInfo, bool) {
	if o == nil || isNil(o.OldCmInfoList) {
    return nil, false
	}
	return o.OldCmInfoList, true
}

// HasOldCmInfoList returns a boolean if a field has been set.
func (o *CmInfoReport) HasOldCmInfoList() bool {
	if o != nil && !isNil(o.OldCmInfoList) {
		return true
	}

	return false
}

// SetOldCmInfoList gets a reference to the given []CmInfo and assigns it to the OldCmInfoList field.
func (o *CmInfoReport) SetOldCmInfoList(v []CmInfo) {
	o.OldCmInfoList = v
}

// GetNewCmInfoList returns the NewCmInfoList field value
func (o *CmInfoReport) GetNewCmInfoList() []CmInfo {
	if o == nil {
		var ret []CmInfo
		return ret
	}

	return o.NewCmInfoList
}

// GetNewCmInfoListOk returns a tuple with the NewCmInfoList field value
// and a boolean to check if the value has been set.
func (o *CmInfoReport) GetNewCmInfoListOk() ([]CmInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.NewCmInfoList, true
}

// SetNewCmInfoList sets field value
func (o *CmInfoReport) SetNewCmInfoList(v []CmInfo) {
	o.NewCmInfoList = v
}

func (o CmInfoReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OldCmInfoList) {
		toSerialize["oldCmInfoList"] = o.OldCmInfoList
	}
	if true {
		toSerialize["newCmInfoList"] = o.NewCmInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableCmInfoReport struct {
	value *CmInfoReport
	isSet bool
}

func (v NullableCmInfoReport) Get() *CmInfoReport {
	return v.value
}

func (v *NullableCmInfoReport) Set(val *CmInfoReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCmInfoReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCmInfoReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmInfoReport(val *CmInfoReport) *NullableCmInfoReport {
	return &NullableCmInfoReport{value: val, isSet: true}
}

func (v NullableCmInfoReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmInfoReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


