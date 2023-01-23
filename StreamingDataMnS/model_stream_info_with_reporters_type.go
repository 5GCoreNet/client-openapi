/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
)

// StreamInfoWithReportersType Reporting stream meta-data with added information about reporters.
type StreamInfoWithReportersType struct {
	StreamInfo *StreamInfoType `json:"streamInfo,omitempty"`
	Reporters []ProducerIdType `json:"reporters,omitempty"`
}

// NewStreamInfoWithReportersType instantiates a new StreamInfoWithReportersType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamInfoWithReportersType() *StreamInfoWithReportersType {
	this := StreamInfoWithReportersType{}
	return &this
}

// NewStreamInfoWithReportersTypeWithDefaults instantiates a new StreamInfoWithReportersType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamInfoWithReportersTypeWithDefaults() *StreamInfoWithReportersType {
	this := StreamInfoWithReportersType{}
	return &this
}

// GetStreamInfo returns the StreamInfo field value if set, zero value otherwise.
func (o *StreamInfoWithReportersType) GetStreamInfo() StreamInfoType {
	if o == nil || isNil(o.StreamInfo) {
		var ret StreamInfoType
		return ret
	}
	return *o.StreamInfo
}

// GetStreamInfoOk returns a tuple with the StreamInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamInfoWithReportersType) GetStreamInfoOk() (*StreamInfoType, bool) {
	if o == nil || isNil(o.StreamInfo) {
    return nil, false
	}
	return o.StreamInfo, true
}

// HasStreamInfo returns a boolean if a field has been set.
func (o *StreamInfoWithReportersType) HasStreamInfo() bool {
	if o != nil && !isNil(o.StreamInfo) {
		return true
	}

	return false
}

// SetStreamInfo gets a reference to the given StreamInfoType and assigns it to the StreamInfo field.
func (o *StreamInfoWithReportersType) SetStreamInfo(v StreamInfoType) {
	o.StreamInfo = &v
}

// GetReporters returns the Reporters field value if set, zero value otherwise.
func (o *StreamInfoWithReportersType) GetReporters() []ProducerIdType {
	if o == nil || isNil(o.Reporters) {
		var ret []ProducerIdType
		return ret
	}
	return o.Reporters
}

// GetReportersOk returns a tuple with the Reporters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamInfoWithReportersType) GetReportersOk() ([]ProducerIdType, bool) {
	if o == nil || isNil(o.Reporters) {
    return nil, false
	}
	return o.Reporters, true
}

// HasReporters returns a boolean if a field has been set.
func (o *StreamInfoWithReportersType) HasReporters() bool {
	if o != nil && !isNil(o.Reporters) {
		return true
	}

	return false
}

// SetReporters gets a reference to the given []ProducerIdType and assigns it to the Reporters field.
func (o *StreamInfoWithReportersType) SetReporters(v []ProducerIdType) {
	o.Reporters = v
}

func (o StreamInfoWithReportersType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StreamInfo) {
		toSerialize["streamInfo"] = o.StreamInfo
	}
	if !isNil(o.Reporters) {
		toSerialize["reporters"] = o.Reporters
	}
	return json.Marshal(toSerialize)
}

type NullableStreamInfoWithReportersType struct {
	value *StreamInfoWithReportersType
	isSet bool
}

func (v NullableStreamInfoWithReportersType) Get() *StreamInfoWithReportersType {
	return v.value
}

func (v *NullableStreamInfoWithReportersType) Set(val *StreamInfoWithReportersType) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamInfoWithReportersType) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamInfoWithReportersType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamInfoWithReportersType(val *StreamInfoWithReportersType) *NullableStreamInfoWithReportersType {
	return &NullableStreamInfoWithReportersType{value: val, isSet: true}
}

func (v NullableStreamInfoWithReportersType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamInfoWithReportersType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


