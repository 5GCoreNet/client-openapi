/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
)

// C2LinkQualityThrlds Represents the C2 link quality thresholds.
type C2LinkQualityThrlds struct {
	NrRsrpThrldLow *int32 `json:"nrRsrpThrldLow,omitempty"`
	NrRsrpThrldHigh *int32 `json:"nrRsrpThrldHigh,omitempty"`
	NrRsrqThrldLow *int32 `json:"nrRsrqThrldLow,omitempty"`
	NrRsrqThrldHigh *int32 `json:"nrRsrqThrldHigh,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	PacketLossThrldLow *int32 `json:"packetLossThrldLow,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	PacketLossThrldHigh *int32 `json:"packetLossThrldHigh,omitempty"`
}

// NewC2LinkQualityThrlds instantiates a new C2LinkQualityThrlds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC2LinkQualityThrlds() *C2LinkQualityThrlds {
	this := C2LinkQualityThrlds{}
	return &this
}

// NewC2LinkQualityThrldsWithDefaults instantiates a new C2LinkQualityThrlds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC2LinkQualityThrldsWithDefaults() *C2LinkQualityThrlds {
	this := C2LinkQualityThrlds{}
	return &this
}

// GetNrRsrpThrldLow returns the NrRsrpThrldLow field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetNrRsrpThrldLow() int32 {
	if o == nil || isNil(o.NrRsrpThrldLow) {
		var ret int32
		return ret
	}
	return *o.NrRsrpThrldLow
}

// GetNrRsrpThrldLowOk returns a tuple with the NrRsrpThrldLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetNrRsrpThrldLowOk() (*int32, bool) {
	if o == nil || isNil(o.NrRsrpThrldLow) {
    return nil, false
	}
	return o.NrRsrpThrldLow, true
}

// HasNrRsrpThrldLow returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasNrRsrpThrldLow() bool {
	if o != nil && !isNil(o.NrRsrpThrldLow) {
		return true
	}

	return false
}

// SetNrRsrpThrldLow gets a reference to the given int32 and assigns it to the NrRsrpThrldLow field.
func (o *C2LinkQualityThrlds) SetNrRsrpThrldLow(v int32) {
	o.NrRsrpThrldLow = &v
}

// GetNrRsrpThrldHigh returns the NrRsrpThrldHigh field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetNrRsrpThrldHigh() int32 {
	if o == nil || isNil(o.NrRsrpThrldHigh) {
		var ret int32
		return ret
	}
	return *o.NrRsrpThrldHigh
}

// GetNrRsrpThrldHighOk returns a tuple with the NrRsrpThrldHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetNrRsrpThrldHighOk() (*int32, bool) {
	if o == nil || isNil(o.NrRsrpThrldHigh) {
    return nil, false
	}
	return o.NrRsrpThrldHigh, true
}

// HasNrRsrpThrldHigh returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasNrRsrpThrldHigh() bool {
	if o != nil && !isNil(o.NrRsrpThrldHigh) {
		return true
	}

	return false
}

// SetNrRsrpThrldHigh gets a reference to the given int32 and assigns it to the NrRsrpThrldHigh field.
func (o *C2LinkQualityThrlds) SetNrRsrpThrldHigh(v int32) {
	o.NrRsrpThrldHigh = &v
}

// GetNrRsrqThrldLow returns the NrRsrqThrldLow field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetNrRsrqThrldLow() int32 {
	if o == nil || isNil(o.NrRsrqThrldLow) {
		var ret int32
		return ret
	}
	return *o.NrRsrqThrldLow
}

// GetNrRsrqThrldLowOk returns a tuple with the NrRsrqThrldLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetNrRsrqThrldLowOk() (*int32, bool) {
	if o == nil || isNil(o.NrRsrqThrldLow) {
    return nil, false
	}
	return o.NrRsrqThrldLow, true
}

// HasNrRsrqThrldLow returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasNrRsrqThrldLow() bool {
	if o != nil && !isNil(o.NrRsrqThrldLow) {
		return true
	}

	return false
}

// SetNrRsrqThrldLow gets a reference to the given int32 and assigns it to the NrRsrqThrldLow field.
func (o *C2LinkQualityThrlds) SetNrRsrqThrldLow(v int32) {
	o.NrRsrqThrldLow = &v
}

// GetNrRsrqThrldHigh returns the NrRsrqThrldHigh field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetNrRsrqThrldHigh() int32 {
	if o == nil || isNil(o.NrRsrqThrldHigh) {
		var ret int32
		return ret
	}
	return *o.NrRsrqThrldHigh
}

// GetNrRsrqThrldHighOk returns a tuple with the NrRsrqThrldHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetNrRsrqThrldHighOk() (*int32, bool) {
	if o == nil || isNil(o.NrRsrqThrldHigh) {
    return nil, false
	}
	return o.NrRsrqThrldHigh, true
}

// HasNrRsrqThrldHigh returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasNrRsrqThrldHigh() bool {
	if o != nil && !isNil(o.NrRsrqThrldHigh) {
		return true
	}

	return false
}

// SetNrRsrqThrldHigh gets a reference to the given int32 and assigns it to the NrRsrqThrldHigh field.
func (o *C2LinkQualityThrlds) SetNrRsrqThrldHigh(v int32) {
	o.NrRsrqThrldHigh = &v
}

// GetPacketLossThrldLow returns the PacketLossThrldLow field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetPacketLossThrldLow() int32 {
	if o == nil || isNil(o.PacketLossThrldLow) {
		var ret int32
		return ret
	}
	return *o.PacketLossThrldLow
}

// GetPacketLossThrldLowOk returns a tuple with the PacketLossThrldLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetPacketLossThrldLowOk() (*int32, bool) {
	if o == nil || isNil(o.PacketLossThrldLow) {
    return nil, false
	}
	return o.PacketLossThrldLow, true
}

// HasPacketLossThrldLow returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasPacketLossThrldLow() bool {
	if o != nil && !isNil(o.PacketLossThrldLow) {
		return true
	}

	return false
}

// SetPacketLossThrldLow gets a reference to the given int32 and assigns it to the PacketLossThrldLow field.
func (o *C2LinkQualityThrlds) SetPacketLossThrldLow(v int32) {
	o.PacketLossThrldLow = &v
}

// GetPacketLossThrldHigh returns the PacketLossThrldHigh field value if set, zero value otherwise.
func (o *C2LinkQualityThrlds) GetPacketLossThrldHigh() int32 {
	if o == nil || isNil(o.PacketLossThrldHigh) {
		var ret int32
		return ret
	}
	return *o.PacketLossThrldHigh
}

// GetPacketLossThrldHighOk returns a tuple with the PacketLossThrldHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2LinkQualityThrlds) GetPacketLossThrldHighOk() (*int32, bool) {
	if o == nil || isNil(o.PacketLossThrldHigh) {
    return nil, false
	}
	return o.PacketLossThrldHigh, true
}

// HasPacketLossThrldHigh returns a boolean if a field has been set.
func (o *C2LinkQualityThrlds) HasPacketLossThrldHigh() bool {
	if o != nil && !isNil(o.PacketLossThrldHigh) {
		return true
	}

	return false
}

// SetPacketLossThrldHigh gets a reference to the given int32 and assigns it to the PacketLossThrldHigh field.
func (o *C2LinkQualityThrlds) SetPacketLossThrldHigh(v int32) {
	o.PacketLossThrldHigh = &v
}

func (o C2LinkQualityThrlds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NrRsrpThrldLow) {
		toSerialize["nrRsrpThrldLow"] = o.NrRsrpThrldLow
	}
	if !isNil(o.NrRsrpThrldHigh) {
		toSerialize["nrRsrpThrldHigh"] = o.NrRsrpThrldHigh
	}
	if !isNil(o.NrRsrqThrldLow) {
		toSerialize["nrRsrqThrldLow"] = o.NrRsrqThrldLow
	}
	if !isNil(o.NrRsrqThrldHigh) {
		toSerialize["nrRsrqThrldHigh"] = o.NrRsrqThrldHigh
	}
	if !isNil(o.PacketLossThrldLow) {
		toSerialize["packetLossThrldLow"] = o.PacketLossThrldLow
	}
	if !isNil(o.PacketLossThrldHigh) {
		toSerialize["packetLossThrldHigh"] = o.PacketLossThrldHigh
	}
	return json.Marshal(toSerialize)
}

type NullableC2LinkQualityThrlds struct {
	value *C2LinkQualityThrlds
	isSet bool
}

func (v NullableC2LinkQualityThrlds) Get() *C2LinkQualityThrlds {
	return v.value
}

func (v *NullableC2LinkQualityThrlds) Set(val *C2LinkQualityThrlds) {
	v.value = val
	v.isSet = true
}

func (v NullableC2LinkQualityThrlds) IsSet() bool {
	return v.isSet
}

func (v *NullableC2LinkQualityThrlds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC2LinkQualityThrlds(val *C2LinkQualityThrlds) *NullableC2LinkQualityThrlds {
	return &NullableC2LinkQualityThrlds{value: val, isSet: true}
}

func (v NullableC2LinkQualityThrlds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC2LinkQualityThrlds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


