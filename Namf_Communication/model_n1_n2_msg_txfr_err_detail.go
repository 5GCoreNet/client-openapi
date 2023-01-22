/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// N1N2MsgTxfrErrDetail N1/N2 Message Transfer Error Details
type N1N2MsgTxfrErrDetail struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RetryAfter *int32 `json:"retryAfter,omitempty"`
	HighestPrioArp *Arp `json:"highestPrioArp,omitempty"`
	// indicating a time in seconds.
	MaxWaitingTime *int32 `json:"maxWaitingTime,omitempty"`
}

// NewN1N2MsgTxfrErrDetail instantiates a new N1N2MsgTxfrErrDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN1N2MsgTxfrErrDetail() *N1N2MsgTxfrErrDetail {
	this := N1N2MsgTxfrErrDetail{}
	return &this
}

// NewN1N2MsgTxfrErrDetailWithDefaults instantiates a new N1N2MsgTxfrErrDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN1N2MsgTxfrErrDetailWithDefaults() *N1N2MsgTxfrErrDetail {
	this := N1N2MsgTxfrErrDetail{}
	return &this
}

// GetRetryAfter returns the RetryAfter field value if set, zero value otherwise.
func (o *N1N2MsgTxfrErrDetail) GetRetryAfter() int32 {
	if o == nil || isNil(o.RetryAfter) {
		var ret int32
		return ret
	}
	return *o.RetryAfter
}

// GetRetryAfterOk returns a tuple with the RetryAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MsgTxfrErrDetail) GetRetryAfterOk() (*int32, bool) {
	if o == nil || isNil(o.RetryAfter) {
    return nil, false
	}
	return o.RetryAfter, true
}

// HasRetryAfter returns a boolean if a field has been set.
func (o *N1N2MsgTxfrErrDetail) HasRetryAfter() bool {
	if o != nil && !isNil(o.RetryAfter) {
		return true
	}

	return false
}

// SetRetryAfter gets a reference to the given int32 and assigns it to the RetryAfter field.
func (o *N1N2MsgTxfrErrDetail) SetRetryAfter(v int32) {
	o.RetryAfter = &v
}

// GetHighestPrioArp returns the HighestPrioArp field value if set, zero value otherwise.
func (o *N1N2MsgTxfrErrDetail) GetHighestPrioArp() Arp {
	if o == nil || isNil(o.HighestPrioArp) {
		var ret Arp
		return ret
	}
	return *o.HighestPrioArp
}

// GetHighestPrioArpOk returns a tuple with the HighestPrioArp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MsgTxfrErrDetail) GetHighestPrioArpOk() (*Arp, bool) {
	if o == nil || isNil(o.HighestPrioArp) {
    return nil, false
	}
	return o.HighestPrioArp, true
}

// HasHighestPrioArp returns a boolean if a field has been set.
func (o *N1N2MsgTxfrErrDetail) HasHighestPrioArp() bool {
	if o != nil && !isNil(o.HighestPrioArp) {
		return true
	}

	return false
}

// SetHighestPrioArp gets a reference to the given Arp and assigns it to the HighestPrioArp field.
func (o *N1N2MsgTxfrErrDetail) SetHighestPrioArp(v Arp) {
	o.HighestPrioArp = &v
}

// GetMaxWaitingTime returns the MaxWaitingTime field value if set, zero value otherwise.
func (o *N1N2MsgTxfrErrDetail) GetMaxWaitingTime() int32 {
	if o == nil || isNil(o.MaxWaitingTime) {
		var ret int32
		return ret
	}
	return *o.MaxWaitingTime
}

// GetMaxWaitingTimeOk returns a tuple with the MaxWaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MsgTxfrErrDetail) GetMaxWaitingTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxWaitingTime) {
    return nil, false
	}
	return o.MaxWaitingTime, true
}

// HasMaxWaitingTime returns a boolean if a field has been set.
func (o *N1N2MsgTxfrErrDetail) HasMaxWaitingTime() bool {
	if o != nil && !isNil(o.MaxWaitingTime) {
		return true
	}

	return false
}

// SetMaxWaitingTime gets a reference to the given int32 and assigns it to the MaxWaitingTime field.
func (o *N1N2MsgTxfrErrDetail) SetMaxWaitingTime(v int32) {
	o.MaxWaitingTime = &v
}

func (o N1N2MsgTxfrErrDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RetryAfter) {
		toSerialize["retryAfter"] = o.RetryAfter
	}
	if !isNil(o.HighestPrioArp) {
		toSerialize["highestPrioArp"] = o.HighestPrioArp
	}
	if !isNil(o.MaxWaitingTime) {
		toSerialize["maxWaitingTime"] = o.MaxWaitingTime
	}
	return json.Marshal(toSerialize)
}

type NullableN1N2MsgTxfrErrDetail struct {
	value *N1N2MsgTxfrErrDetail
	isSet bool
}

func (v NullableN1N2MsgTxfrErrDetail) Get() *N1N2MsgTxfrErrDetail {
	return v.value
}

func (v *NullableN1N2MsgTxfrErrDetail) Set(val *N1N2MsgTxfrErrDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableN1N2MsgTxfrErrDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableN1N2MsgTxfrErrDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1N2MsgTxfrErrDetail(val *N1N2MsgTxfrErrDetail) *NullableN1N2MsgTxfrErrDetail {
	return &NullableN1N2MsgTxfrErrDetail{value: val, isSet: true}
}

func (v NullableN1N2MsgTxfrErrDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1N2MsgTxfrErrDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


