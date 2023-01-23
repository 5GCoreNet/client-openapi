/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N32_Handshake

import (
	"encoding/json"
)

// N32fErrorDetail Details about the N32f error
type N32fErrorDetail struct {
	Attribute string `json:"attribute"`
	MsgReconstructFailReason FailureReason `json:"msgReconstructFailReason"`
}

// NewN32fErrorDetail instantiates a new N32fErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN32fErrorDetail(attribute string, msgReconstructFailReason FailureReason) *N32fErrorDetail {
	this := N32fErrorDetail{}
	this.Attribute = attribute
	this.MsgReconstructFailReason = msgReconstructFailReason
	return &this
}

// NewN32fErrorDetailWithDefaults instantiates a new N32fErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN32fErrorDetailWithDefaults() *N32fErrorDetail {
	this := N32fErrorDetail{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *N32fErrorDetail) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *N32fErrorDetail) GetAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *N32fErrorDetail) SetAttribute(v string) {
	o.Attribute = v
}

// GetMsgReconstructFailReason returns the MsgReconstructFailReason field value
func (o *N32fErrorDetail) GetMsgReconstructFailReason() FailureReason {
	if o == nil {
		var ret FailureReason
		return ret
	}

	return o.MsgReconstructFailReason
}

// GetMsgReconstructFailReasonOk returns a tuple with the MsgReconstructFailReason field value
// and a boolean to check if the value has been set.
func (o *N32fErrorDetail) GetMsgReconstructFailReasonOk() (*FailureReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MsgReconstructFailReason, true
}

// SetMsgReconstructFailReason sets field value
func (o *N32fErrorDetail) SetMsgReconstructFailReason(v FailureReason) {
	o.MsgReconstructFailReason = v
}

func (o N32fErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["msgReconstructFailReason"] = o.MsgReconstructFailReason
	}
	return json.Marshal(toSerialize)
}

type NullableN32fErrorDetail struct {
	value *N32fErrorDetail
	isSet bool
}

func (v NullableN32fErrorDetail) Get() *N32fErrorDetail {
	return v.value
}

func (v *NullableN32fErrorDetail) Set(val *N32fErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableN32fErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableN32fErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN32fErrorDetail(val *N32fErrorDetail) *NullableN32fErrorDetail {
	return &NullableN32fErrorDetail{value: val, isSet: true}
}

func (v NullableN32fErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN32fErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


