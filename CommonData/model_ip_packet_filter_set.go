/*
5GMS Common Data Types

5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
)

// IpPacketFilterSet struct for IpPacketFilterSet
type IpPacketFilterSet struct {
	SrcIp *string `json:"srcIp,omitempty"`
	DstIp *string `json:"dstIp,omitempty"`
	Protocol *int32 `json:"protocol,omitempty"`
	SrcPort *int32 `json:"srcPort,omitempty"`
	DstPort *int32 `json:"dstPort,omitempty"`
	ToSTc *string `json:"toSTc,omitempty"`
	FlowLabel *int32 `json:"flowLabel,omitempty"`
	Spi *int32 `json:"spi,omitempty"`
	Direction string `json:"direction"`
}

// NewIpPacketFilterSet instantiates a new IpPacketFilterSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpPacketFilterSet(direction string) *IpPacketFilterSet {
	this := IpPacketFilterSet{}
	this.Direction = direction
	return &this
}

// NewIpPacketFilterSetWithDefaults instantiates a new IpPacketFilterSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPacketFilterSetWithDefaults() *IpPacketFilterSet {
	this := IpPacketFilterSet{}
	return &this
}

// GetSrcIp returns the SrcIp field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetSrcIp() string {
	if o == nil || isNil(o.SrcIp) {
		var ret string
		return ret
	}
	return *o.SrcIp
}

// GetSrcIpOk returns a tuple with the SrcIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetSrcIpOk() (*string, bool) {
	if o == nil || isNil(o.SrcIp) {
    return nil, false
	}
	return o.SrcIp, true
}

// HasSrcIp returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasSrcIp() bool {
	if o != nil && !isNil(o.SrcIp) {
		return true
	}

	return false
}

// SetSrcIp gets a reference to the given string and assigns it to the SrcIp field.
func (o *IpPacketFilterSet) SetSrcIp(v string) {
	o.SrcIp = &v
}

// GetDstIp returns the DstIp field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetDstIp() string {
	if o == nil || isNil(o.DstIp) {
		var ret string
		return ret
	}
	return *o.DstIp
}

// GetDstIpOk returns a tuple with the DstIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetDstIpOk() (*string, bool) {
	if o == nil || isNil(o.DstIp) {
    return nil, false
	}
	return o.DstIp, true
}

// HasDstIp returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasDstIp() bool {
	if o != nil && !isNil(o.DstIp) {
		return true
	}

	return false
}

// SetDstIp gets a reference to the given string and assigns it to the DstIp field.
func (o *IpPacketFilterSet) SetDstIp(v string) {
	o.DstIp = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetProtocol() int32 {
	if o == nil || isNil(o.Protocol) {
		var ret int32
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetProtocolOk() (*int32, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given int32 and assigns it to the Protocol field.
func (o *IpPacketFilterSet) SetProtocol(v int32) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetSrcPort() int32 {
	if o == nil || isNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetSrcPortOk() (*int32, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *IpPacketFilterSet) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetDstPort() int32 {
	if o == nil || isNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetDstPortOk() (*int32, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *IpPacketFilterSet) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetToSTc returns the ToSTc field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetToSTc() string {
	if o == nil || isNil(o.ToSTc) {
		var ret string
		return ret
	}
	return *o.ToSTc
}

// GetToSTcOk returns a tuple with the ToSTc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetToSTcOk() (*string, bool) {
	if o == nil || isNil(o.ToSTc) {
    return nil, false
	}
	return o.ToSTc, true
}

// HasToSTc returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasToSTc() bool {
	if o != nil && !isNil(o.ToSTc) {
		return true
	}

	return false
}

// SetToSTc gets a reference to the given string and assigns it to the ToSTc field.
func (o *IpPacketFilterSet) SetToSTc(v string) {
	o.ToSTc = &v
}

// GetFlowLabel returns the FlowLabel field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetFlowLabel() int32 {
	if o == nil || isNil(o.FlowLabel) {
		var ret int32
		return ret
	}
	return *o.FlowLabel
}

// GetFlowLabelOk returns a tuple with the FlowLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetFlowLabelOk() (*int32, bool) {
	if o == nil || isNil(o.FlowLabel) {
    return nil, false
	}
	return o.FlowLabel, true
}

// HasFlowLabel returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasFlowLabel() bool {
	if o != nil && !isNil(o.FlowLabel) {
		return true
	}

	return false
}

// SetFlowLabel gets a reference to the given int32 and assigns it to the FlowLabel field.
func (o *IpPacketFilterSet) SetFlowLabel(v int32) {
	o.FlowLabel = &v
}

// GetSpi returns the Spi field value if set, zero value otherwise.
func (o *IpPacketFilterSet) GetSpi() int32 {
	if o == nil || isNil(o.Spi) {
		var ret int32
		return ret
	}
	return *o.Spi
}

// GetSpiOk returns a tuple with the Spi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetSpiOk() (*int32, bool) {
	if o == nil || isNil(o.Spi) {
    return nil, false
	}
	return o.Spi, true
}

// HasSpi returns a boolean if a field has been set.
func (o *IpPacketFilterSet) HasSpi() bool {
	if o != nil && !isNil(o.Spi) {
		return true
	}

	return false
}

// SetSpi gets a reference to the given int32 and assigns it to the Spi field.
func (o *IpPacketFilterSet) SetSpi(v int32) {
	o.Spi = &v
}

// GetDirection returns the Direction field value
func (o *IpPacketFilterSet) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *IpPacketFilterSet) GetDirectionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *IpPacketFilterSet) SetDirection(v string) {
	o.Direction = v
}

func (o IpPacketFilterSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SrcIp) {
		toSerialize["srcIp"] = o.SrcIp
	}
	if !isNil(o.DstIp) {
		toSerialize["dstIp"] = o.DstIp
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !isNil(o.ToSTc) {
		toSerialize["toSTc"] = o.ToSTc
	}
	if !isNil(o.FlowLabel) {
		toSerialize["flowLabel"] = o.FlowLabel
	}
	if !isNil(o.Spi) {
		toSerialize["spi"] = o.Spi
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	return json.Marshal(toSerialize)
}

type NullableIpPacketFilterSet struct {
	value *IpPacketFilterSet
	isSet bool
}

func (v NullableIpPacketFilterSet) Get() *IpPacketFilterSet {
	return v.value
}

func (v *NullableIpPacketFilterSet) Set(val *IpPacketFilterSet) {
	v.value = val
	v.isSet = true
}

func (v NullableIpPacketFilterSet) IsSet() bool {
	return v.isSet
}

func (v *NullableIpPacketFilterSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpPacketFilterSet(val *IpPacketFilterSet) *NullableIpPacketFilterSet {
	return &NullableIpPacketFilterSet{value: val, isSet: true}
}

func (v NullableIpPacketFilterSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpPacketFilterSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


