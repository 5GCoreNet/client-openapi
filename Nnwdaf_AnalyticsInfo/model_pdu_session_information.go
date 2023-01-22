/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// PduSessionInformation Represents the PDU session related information.
type PduSessionInformation struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessId *int32 `json:"pduSessId,omitempty"`
	SessInfo *PduSessionInfo `json:"sessInfo,omitempty"`
}

// NewPduSessionInformation instantiates a new PduSessionInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionInformation() *PduSessionInformation {
	this := PduSessionInformation{}
	return &this
}

// NewPduSessionInformationWithDefaults instantiates a new PduSessionInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionInformationWithDefaults() *PduSessionInformation {
	this := PduSessionInformation{}
	return &this
}

// GetPduSessId returns the PduSessId field value if set, zero value otherwise.
func (o *PduSessionInformation) GetPduSessId() int32 {
	if o == nil || isNil(o.PduSessId) {
		var ret int32
		return ret
	}
	return *o.PduSessId
}

// GetPduSessIdOk returns a tuple with the PduSessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionInformation) GetPduSessIdOk() (*int32, bool) {
	if o == nil || isNil(o.PduSessId) {
    return nil, false
	}
	return o.PduSessId, true
}

// HasPduSessId returns a boolean if a field has been set.
func (o *PduSessionInformation) HasPduSessId() bool {
	if o != nil && !isNil(o.PduSessId) {
		return true
	}

	return false
}

// SetPduSessId gets a reference to the given int32 and assigns it to the PduSessId field.
func (o *PduSessionInformation) SetPduSessId(v int32) {
	o.PduSessId = &v
}

// GetSessInfo returns the SessInfo field value if set, zero value otherwise.
func (o *PduSessionInformation) GetSessInfo() PduSessionInfo {
	if o == nil || isNil(o.SessInfo) {
		var ret PduSessionInfo
		return ret
	}
	return *o.SessInfo
}

// GetSessInfoOk returns a tuple with the SessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionInformation) GetSessInfoOk() (*PduSessionInfo, bool) {
	if o == nil || isNil(o.SessInfo) {
    return nil, false
	}
	return o.SessInfo, true
}

// HasSessInfo returns a boolean if a field has been set.
func (o *PduSessionInformation) HasSessInfo() bool {
	if o != nil && !isNil(o.SessInfo) {
		return true
	}

	return false
}

// SetSessInfo gets a reference to the given PduSessionInfo and assigns it to the SessInfo field.
func (o *PduSessionInformation) SetSessInfo(v PduSessionInfo) {
	o.SessInfo = &v
}

func (o PduSessionInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PduSessId) {
		toSerialize["pduSessId"] = o.PduSessId
	}
	if !isNil(o.SessInfo) {
		toSerialize["sessInfo"] = o.SessInfo
	}
	return json.Marshal(toSerialize)
}

type NullablePduSessionInformation struct {
	value *PduSessionInformation
	isSet bool
}

func (v NullablePduSessionInformation) Get() *PduSessionInformation {
	return v.value
}

func (v *NullablePduSessionInformation) Set(val *PduSessionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionInformation(val *PduSessionInformation) *NullablePduSessionInformation {
	return &NullablePduSessionInformation{value: val, isSet: true}
}

func (v NullablePduSessionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


