/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// N1N2MsgTxfrFailureNotification Data within a N1/N2 Message Transfer Failure Notification request
type N1N2MsgTxfrFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause"`
	// String providing an URI formatted according to RFC 3986.
	N1n2MsgDataUri string `json:"n1n2MsgDataUri"`
}

// NewN1N2MsgTxfrFailureNotification instantiates a new N1N2MsgTxfrFailureNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN1N2MsgTxfrFailureNotification(cause N1N2MessageTransferCause, n1n2MsgDataUri string) *N1N2MsgTxfrFailureNotification {
	this := N1N2MsgTxfrFailureNotification{}
	this.Cause = cause
	this.N1n2MsgDataUri = n1n2MsgDataUri
	return &this
}

// NewN1N2MsgTxfrFailureNotificationWithDefaults instantiates a new N1N2MsgTxfrFailureNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN1N2MsgTxfrFailureNotificationWithDefaults() *N1N2MsgTxfrFailureNotification {
	this := N1N2MsgTxfrFailureNotification{}
	return &this
}

// GetCause returns the Cause field value
func (o *N1N2MsgTxfrFailureNotification) GetCause() N1N2MessageTransferCause {
	if o == nil {
		var ret N1N2MessageTransferCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *N1N2MsgTxfrFailureNotification) GetCauseOk() (*N1N2MessageTransferCause, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *N1N2MsgTxfrFailureNotification) SetCause(v N1N2MessageTransferCause) {
	o.Cause = v
}

// GetN1n2MsgDataUri returns the N1n2MsgDataUri field value
func (o *N1N2MsgTxfrFailureNotification) GetN1n2MsgDataUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N1n2MsgDataUri
}

// GetN1n2MsgDataUriOk returns a tuple with the N1n2MsgDataUri field value
// and a boolean to check if the value has been set.
func (o *N1N2MsgTxfrFailureNotification) GetN1n2MsgDataUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.N1n2MsgDataUri, true
}

// SetN1n2MsgDataUri sets field value
func (o *N1N2MsgTxfrFailureNotification) SetN1n2MsgDataUri(v string) {
	o.N1n2MsgDataUri = v
}

func (o N1N2MsgTxfrFailureNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cause"] = o.Cause
	}
	if true {
		toSerialize["n1n2MsgDataUri"] = o.N1n2MsgDataUri
	}
	return json.Marshal(toSerialize)
}

type NullableN1N2MsgTxfrFailureNotification struct {
	value *N1N2MsgTxfrFailureNotification
	isSet bool
}

func (v NullableN1N2MsgTxfrFailureNotification) Get() *N1N2MsgTxfrFailureNotification {
	return v.value
}

func (v *NullableN1N2MsgTxfrFailureNotification) Set(val *N1N2MsgTxfrFailureNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableN1N2MsgTxfrFailureNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableN1N2MsgTxfrFailureNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1N2MsgTxfrFailureNotification(val *N1N2MsgTxfrFailureNotification) *NullableN1N2MsgTxfrFailureNotification {
	return &NullableN1N2MsgTxfrFailureNotification{value: val, isSet: true}
}

func (v NullableN1N2MsgTxfrFailureNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1N2MsgTxfrFailureNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


