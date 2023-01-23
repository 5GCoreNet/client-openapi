/*
Namf_MT

AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MT

import (
	"encoding/json"
)

// UeInfo list of UEs requested to be made reachable for the MBS Session
type UeInfo struct {
	UeList []string `json:"ueList"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
}

// NewUeInfo instantiates a new UeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeInfo(ueList []string) *UeInfo {
	this := UeInfo{}
	this.UeList = ueList
	return &this
}

// NewUeInfoWithDefaults instantiates a new UeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeInfoWithDefaults() *UeInfo {
	this := UeInfo{}
	return &this
}

// GetUeList returns the UeList field value
func (o *UeInfo) GetUeList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UeList
}

// GetUeListOk returns a tuple with the UeList field value
// and a boolean to check if the value has been set.
func (o *UeInfo) GetUeListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UeList, true
}

// SetUeList sets field value
func (o *UeInfo) SetUeList(v []string) {
	o.UeList = v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *UeInfo) GetPduSessionId() int32 {
	if o == nil || isNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInfo) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || isNil(o.PduSessionId) {
    return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *UeInfo) HasPduSessionId() bool {
	if o != nil && !isNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *UeInfo) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

func (o UeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ueList"] = o.UeList
	}
	if !isNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	return json.Marshal(toSerialize)
}

type NullableUeInfo struct {
	value *UeInfo
	isSet bool
}

func (v NullableUeInfo) Get() *UeInfo {
	return v.value
}

func (v *NullableUeInfo) Set(val *UeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeInfo(val *UeInfo) *NullableUeInfo {
	return &NullableUeInfo{value: val, isSet: true}
}

func (v NullableUeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


