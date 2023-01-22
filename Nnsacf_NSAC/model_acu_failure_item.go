/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnsacf_NSAC

import (
	"encoding/json"
)

// AcuFailureItem struct for AcuFailureItem
type AcuFailureItem struct {
	Snssai Snssai `json:"snssai"`
	Reason *AcuFailureReason `json:"reason,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
}

// NewAcuFailureItem instantiates a new AcuFailureItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcuFailureItem(snssai Snssai) *AcuFailureItem {
	this := AcuFailureItem{}
	this.Snssai = snssai
	return &this
}

// NewAcuFailureItemWithDefaults instantiates a new AcuFailureItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcuFailureItemWithDefaults() *AcuFailureItem {
	this := AcuFailureItem{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *AcuFailureItem) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *AcuFailureItem) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *AcuFailureItem) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AcuFailureItem) GetReason() AcuFailureReason {
	if o == nil || isNil(o.Reason) {
		var ret AcuFailureReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcuFailureItem) GetReasonOk() (*AcuFailureReason, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AcuFailureItem) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given AcuFailureReason and assigns it to the Reason field.
func (o *AcuFailureItem) SetReason(v AcuFailureReason) {
	o.Reason = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *AcuFailureItem) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcuFailureItem) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
    return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *AcuFailureItem) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *AcuFailureItem) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *AcuFailureItem) GetPduSessionId() int32 {
	if o == nil || isNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcuFailureItem) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || isNil(o.PduSessionId) {
    return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *AcuFailureItem) HasPduSessionId() bool {
	if o != nil && !isNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *AcuFailureItem) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

func (o AcuFailureItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	return json.Marshal(toSerialize)
}

type NullableAcuFailureItem struct {
	value *AcuFailureItem
	isSet bool
}

func (v NullableAcuFailureItem) Get() *AcuFailureItem {
	return v.value
}

func (v *NullableAcuFailureItem) Set(val *AcuFailureItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAcuFailureItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAcuFailureItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcuFailureItem(val *AcuFailureItem) *NullableAcuFailureItem {
	return &NullableAcuFailureItem{value: val, isSet: true}
}

func (v NullableAcuFailureItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcuFailureItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


