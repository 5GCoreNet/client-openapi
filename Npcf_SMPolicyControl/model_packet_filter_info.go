/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
)

// PacketFilterInfo Contains the information from a single packet filter sent from the SMF to the PCF.
type PacketFilterInfo struct {
	// An identifier of packet filter.
	PackFiltId *string `json:"packFiltId,omitempty"`
	// Defines a packet filter for an IP flow.
	PackFiltCont *string `json:"packFiltCont,omitempty"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass *string `json:"tosTrafficClass,omitempty"`
	// The security parameter index of the IPSec packet.
	Spi *string `json:"spi,omitempty"`
	// The Ipv6 flow label header field.
	FlowLabel *string `json:"flowLabel,omitempty"`
	FlowDirection *FlowDirection `json:"flowDirection,omitempty"`
}

// NewPacketFilterInfo instantiates a new PacketFilterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPacketFilterInfo() *PacketFilterInfo {
	this := PacketFilterInfo{}
	return &this
}

// NewPacketFilterInfoWithDefaults instantiates a new PacketFilterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPacketFilterInfoWithDefaults() *PacketFilterInfo {
	this := PacketFilterInfo{}
	return &this
}

// GetPackFiltId returns the PackFiltId field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetPackFiltId() string {
	if o == nil || isNil(o.PackFiltId) {
		var ret string
		return ret
	}
	return *o.PackFiltId
}

// GetPackFiltIdOk returns a tuple with the PackFiltId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetPackFiltIdOk() (*string, bool) {
	if o == nil || isNil(o.PackFiltId) {
    return nil, false
	}
	return o.PackFiltId, true
}

// HasPackFiltId returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasPackFiltId() bool {
	if o != nil && !isNil(o.PackFiltId) {
		return true
	}

	return false
}

// SetPackFiltId gets a reference to the given string and assigns it to the PackFiltId field.
func (o *PacketFilterInfo) SetPackFiltId(v string) {
	o.PackFiltId = &v
}

// GetPackFiltCont returns the PackFiltCont field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetPackFiltCont() string {
	if o == nil || isNil(o.PackFiltCont) {
		var ret string
		return ret
	}
	return *o.PackFiltCont
}

// GetPackFiltContOk returns a tuple with the PackFiltCont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetPackFiltContOk() (*string, bool) {
	if o == nil || isNil(o.PackFiltCont) {
    return nil, false
	}
	return o.PackFiltCont, true
}

// HasPackFiltCont returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasPackFiltCont() bool {
	if o != nil && !isNil(o.PackFiltCont) {
		return true
	}

	return false
}

// SetPackFiltCont gets a reference to the given string and assigns it to the PackFiltCont field.
func (o *PacketFilterInfo) SetPackFiltCont(v string) {
	o.PackFiltCont = &v
}

// GetTosTrafficClass returns the TosTrafficClass field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetTosTrafficClass() string {
	if o == nil || isNil(o.TosTrafficClass) {
		var ret string
		return ret
	}
	return *o.TosTrafficClass
}

// GetTosTrafficClassOk returns a tuple with the TosTrafficClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetTosTrafficClassOk() (*string, bool) {
	if o == nil || isNil(o.TosTrafficClass) {
    return nil, false
	}
	return o.TosTrafficClass, true
}

// HasTosTrafficClass returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasTosTrafficClass() bool {
	if o != nil && !isNil(o.TosTrafficClass) {
		return true
	}

	return false
}

// SetTosTrafficClass gets a reference to the given string and assigns it to the TosTrafficClass field.
func (o *PacketFilterInfo) SetTosTrafficClass(v string) {
	o.TosTrafficClass = &v
}

// GetSpi returns the Spi field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetSpi() string {
	if o == nil || isNil(o.Spi) {
		var ret string
		return ret
	}
	return *o.Spi
}

// GetSpiOk returns a tuple with the Spi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetSpiOk() (*string, bool) {
	if o == nil || isNil(o.Spi) {
    return nil, false
	}
	return o.Spi, true
}

// HasSpi returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasSpi() bool {
	if o != nil && !isNil(o.Spi) {
		return true
	}

	return false
}

// SetSpi gets a reference to the given string and assigns it to the Spi field.
func (o *PacketFilterInfo) SetSpi(v string) {
	o.Spi = &v
}

// GetFlowLabel returns the FlowLabel field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetFlowLabel() string {
	if o == nil || isNil(o.FlowLabel) {
		var ret string
		return ret
	}
	return *o.FlowLabel
}

// GetFlowLabelOk returns a tuple with the FlowLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetFlowLabelOk() (*string, bool) {
	if o == nil || isNil(o.FlowLabel) {
    return nil, false
	}
	return o.FlowLabel, true
}

// HasFlowLabel returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasFlowLabel() bool {
	if o != nil && !isNil(o.FlowLabel) {
		return true
	}

	return false
}

// SetFlowLabel gets a reference to the given string and assigns it to the FlowLabel field.
func (o *PacketFilterInfo) SetFlowLabel(v string) {
	o.FlowLabel = &v
}

// GetFlowDirection returns the FlowDirection field value if set, zero value otherwise.
func (o *PacketFilterInfo) GetFlowDirection() FlowDirection {
	if o == nil || isNil(o.FlowDirection) {
		var ret FlowDirection
		return ret
	}
	return *o.FlowDirection
}

// GetFlowDirectionOk returns a tuple with the FlowDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketFilterInfo) GetFlowDirectionOk() (*FlowDirection, bool) {
	if o == nil || isNil(o.FlowDirection) {
    return nil, false
	}
	return o.FlowDirection, true
}

// HasFlowDirection returns a boolean if a field has been set.
func (o *PacketFilterInfo) HasFlowDirection() bool {
	if o != nil && !isNil(o.FlowDirection) {
		return true
	}

	return false
}

// SetFlowDirection gets a reference to the given FlowDirection and assigns it to the FlowDirection field.
func (o *PacketFilterInfo) SetFlowDirection(v FlowDirection) {
	o.FlowDirection = &v
}

func (o PacketFilterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PackFiltId) {
		toSerialize["packFiltId"] = o.PackFiltId
	}
	if !isNil(o.PackFiltCont) {
		toSerialize["packFiltCont"] = o.PackFiltCont
	}
	if !isNil(o.TosTrafficClass) {
		toSerialize["tosTrafficClass"] = o.TosTrafficClass
	}
	if !isNil(o.Spi) {
		toSerialize["spi"] = o.Spi
	}
	if !isNil(o.FlowLabel) {
		toSerialize["flowLabel"] = o.FlowLabel
	}
	if !isNil(o.FlowDirection) {
		toSerialize["flowDirection"] = o.FlowDirection
	}
	return json.Marshal(toSerialize)
}

type NullablePacketFilterInfo struct {
	value *PacketFilterInfo
	isSet bool
}

func (v NullablePacketFilterInfo) Get() *PacketFilterInfo {
	return v.value
}

func (v *NullablePacketFilterInfo) Set(val *PacketFilterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePacketFilterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePacketFilterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePacketFilterInfo(val *PacketFilterInfo) *NullablePacketFilterInfo {
	return &NullablePacketFilterInfo{value: val, isSet: true}
}

func (v NullablePacketFilterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePacketFilterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

