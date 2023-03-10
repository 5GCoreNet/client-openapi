/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// QosFlowTunnel Tunnel Information per QoS Flow
type QosFlowTunnel struct {
	QfiList []int32 `json:"qfiList"`
	TunnelInfo TunnelInfo `json:"tunnelInfo"`
}

// NewQosFlowTunnel instantiates a new QosFlowTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowTunnel(qfiList []int32, tunnelInfo TunnelInfo) *QosFlowTunnel {
	this := QosFlowTunnel{}
	this.QfiList = qfiList
	this.TunnelInfo = tunnelInfo
	return &this
}

// NewQosFlowTunnelWithDefaults instantiates a new QosFlowTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowTunnelWithDefaults() *QosFlowTunnel {
	this := QosFlowTunnel{}
	return &this
}

// GetQfiList returns the QfiList field value
func (o *QosFlowTunnel) GetQfiList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.QfiList
}

// GetQfiListOk returns a tuple with the QfiList field value
// and a boolean to check if the value has been set.
func (o *QosFlowTunnel) GetQfiListOk() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.QfiList, true
}

// SetQfiList sets field value
func (o *QosFlowTunnel) SetQfiList(v []int32) {
	o.QfiList = v
}

// GetTunnelInfo returns the TunnelInfo field value
func (o *QosFlowTunnel) GetTunnelInfo() TunnelInfo {
	if o == nil {
		var ret TunnelInfo
		return ret
	}

	return o.TunnelInfo
}

// GetTunnelInfoOk returns a tuple with the TunnelInfo field value
// and a boolean to check if the value has been set.
func (o *QosFlowTunnel) GetTunnelInfoOk() (*TunnelInfo, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TunnelInfo, true
}

// SetTunnelInfo sets field value
func (o *QosFlowTunnel) SetTunnelInfo(v TunnelInfo) {
	o.TunnelInfo = v
}

func (o QosFlowTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["qfiList"] = o.QfiList
	}
	if true {
		toSerialize["tunnelInfo"] = o.TunnelInfo
	}
	return json.Marshal(toSerialize)
}

type NullableQosFlowTunnel struct {
	value *QosFlowTunnel
	isSet bool
}

func (v NullableQosFlowTunnel) Get() *QosFlowTunnel {
	return v.value
}

func (v *NullableQosFlowTunnel) Set(val *QosFlowTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowTunnel(val *QosFlowTunnel) *NullableQosFlowTunnel {
	return &NullableQosFlowTunnel{value: val, isSet: true}
}

func (v NullableQosFlowTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


