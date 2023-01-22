/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_MBSBroadcast

import (
	"encoding/json"
)

// ContextUpdateRspData Data within ContextUpdate Response
type ContextUpdateRspData struct {
	N2MbsSmInfoList []N2MbsSmInfo `json:"n2MbsSmInfoList,omitempty"`
	OperationStatus *OperationStatus `json:"operationStatus,omitempty"`
}

// NewContextUpdateRspData instantiates a new ContextUpdateRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdateRspData() *ContextUpdateRspData {
	this := ContextUpdateRspData{}
	return &this
}

// NewContextUpdateRspDataWithDefaults instantiates a new ContextUpdateRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdateRspDataWithDefaults() *ContextUpdateRspData {
	this := ContextUpdateRspData{}
	return &this
}

// GetN2MbsSmInfoList returns the N2MbsSmInfoList field value if set, zero value otherwise.
func (o *ContextUpdateRspData) GetN2MbsSmInfoList() []N2MbsSmInfo {
	if o == nil || isNil(o.N2MbsSmInfoList) {
		var ret []N2MbsSmInfo
		return ret
	}
	return o.N2MbsSmInfoList
}

// GetN2MbsSmInfoListOk returns a tuple with the N2MbsSmInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRspData) GetN2MbsSmInfoListOk() ([]N2MbsSmInfo, bool) {
	if o == nil || isNil(o.N2MbsSmInfoList) {
    return nil, false
	}
	return o.N2MbsSmInfoList, true
}

// HasN2MbsSmInfoList returns a boolean if a field has been set.
func (o *ContextUpdateRspData) HasN2MbsSmInfoList() bool {
	if o != nil && !isNil(o.N2MbsSmInfoList) {
		return true
	}

	return false
}

// SetN2MbsSmInfoList gets a reference to the given []N2MbsSmInfo and assigns it to the N2MbsSmInfoList field.
func (o *ContextUpdateRspData) SetN2MbsSmInfoList(v []N2MbsSmInfo) {
	o.N2MbsSmInfoList = v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *ContextUpdateRspData) GetOperationStatus() OperationStatus {
	if o == nil || isNil(o.OperationStatus) {
		var ret OperationStatus
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRspData) GetOperationStatusOk() (*OperationStatus, bool) {
	if o == nil || isNil(o.OperationStatus) {
    return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *ContextUpdateRspData) HasOperationStatus() bool {
	if o != nil && !isNil(o.OperationStatus) {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given OperationStatus and assigns it to the OperationStatus field.
func (o *ContextUpdateRspData) SetOperationStatus(v OperationStatus) {
	o.OperationStatus = &v
}

func (o ContextUpdateRspData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.N2MbsSmInfoList) {
		toSerialize["n2MbsSmInfoList"] = o.N2MbsSmInfoList
	}
	if !isNil(o.OperationStatus) {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableContextUpdateRspData struct {
	value *ContextUpdateRspData
	isSet bool
}

func (v NullableContextUpdateRspData) Get() *ContextUpdateRspData {
	return v.value
}

func (v *NullableContextUpdateRspData) Set(val *ContextUpdateRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdateRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdateRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdateRspData(val *ContextUpdateRspData) *NullableContextUpdateRspData {
	return &NullableContextUpdateRspData{value: val, isSet: true}
}

func (v NullableContextUpdateRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdateRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


