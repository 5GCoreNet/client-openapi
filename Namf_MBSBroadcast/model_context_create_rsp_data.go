/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
)

// ContextCreateRspData Data within ContextCreate Response
type ContextCreateRspData struct {
	MbsSessionId MbsSessionId `json:"mbsSessionId"`
	N2MbsSmInfoList []N2MbsSmInfo `json:"n2MbsSmInfoList,omitempty"`
	OperationStatus *OperationStatus `json:"operationStatus,omitempty"`
}

// NewContextCreateRspData instantiates a new ContextCreateRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextCreateRspData(mbsSessionId MbsSessionId) *ContextCreateRspData {
	this := ContextCreateRspData{}
	this.MbsSessionId = mbsSessionId
	return &this
}

// NewContextCreateRspDataWithDefaults instantiates a new ContextCreateRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextCreateRspDataWithDefaults() *ContextCreateRspData {
	this := ContextCreateRspData{}
	return &this
}

// GetMbsSessionId returns the MbsSessionId field value
func (o *ContextCreateRspData) GetMbsSessionId() MbsSessionId {
	if o == nil {
		var ret MbsSessionId
		return ret
	}

	return o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value
// and a boolean to check if the value has been set.
func (o *ContextCreateRspData) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsSessionId, true
}

// SetMbsSessionId sets field value
func (o *ContextCreateRspData) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = v
}

// GetN2MbsSmInfoList returns the N2MbsSmInfoList field value if set, zero value otherwise.
func (o *ContextCreateRspData) GetN2MbsSmInfoList() []N2MbsSmInfo {
	if o == nil || isNil(o.N2MbsSmInfoList) {
		var ret []N2MbsSmInfo
		return ret
	}
	return o.N2MbsSmInfoList
}

// GetN2MbsSmInfoListOk returns a tuple with the N2MbsSmInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextCreateRspData) GetN2MbsSmInfoListOk() ([]N2MbsSmInfo, bool) {
	if o == nil || isNil(o.N2MbsSmInfoList) {
    return nil, false
	}
	return o.N2MbsSmInfoList, true
}

// HasN2MbsSmInfoList returns a boolean if a field has been set.
func (o *ContextCreateRspData) HasN2MbsSmInfoList() bool {
	if o != nil && !isNil(o.N2MbsSmInfoList) {
		return true
	}

	return false
}

// SetN2MbsSmInfoList gets a reference to the given []N2MbsSmInfo and assigns it to the N2MbsSmInfoList field.
func (o *ContextCreateRspData) SetN2MbsSmInfoList(v []N2MbsSmInfo) {
	o.N2MbsSmInfoList = v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *ContextCreateRspData) GetOperationStatus() OperationStatus {
	if o == nil || isNil(o.OperationStatus) {
		var ret OperationStatus
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextCreateRspData) GetOperationStatusOk() (*OperationStatus, bool) {
	if o == nil || isNil(o.OperationStatus) {
    return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *ContextCreateRspData) HasOperationStatus() bool {
	if o != nil && !isNil(o.OperationStatus) {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given OperationStatus and assigns it to the OperationStatus field.
func (o *ContextCreateRspData) SetOperationStatus(v OperationStatus) {
	o.OperationStatus = &v
}

func (o ContextCreateRspData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsSessionId"] = o.MbsSessionId
	}
	if !isNil(o.N2MbsSmInfoList) {
		toSerialize["n2MbsSmInfoList"] = o.N2MbsSmInfoList
	}
	if !isNil(o.OperationStatus) {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableContextCreateRspData struct {
	value *ContextCreateRspData
	isSet bool
}

func (v NullableContextCreateRspData) Get() *ContextCreateRspData {
	return v.value
}

func (v *NullableContextCreateRspData) Set(val *ContextCreateRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextCreateRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextCreateRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextCreateRspData(val *ContextCreateRspData) *NullableContextCreateRspData {
	return &NullableContextCreateRspData{value: val, isSet: true}
}

func (v NullableContextCreateRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextCreateRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


