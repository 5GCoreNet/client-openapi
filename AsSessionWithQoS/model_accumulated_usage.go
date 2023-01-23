/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AsSessionWithQoS

import (
	"encoding/json"
)

// AccumulatedUsage Represents an accumulated usage.
type AccumulatedUsage struct {
	// Unsigned integer identifying a period of time in units of seconds.
	Duration *int32 `json:"duration,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume *int64 `json:"totalVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
}

// NewAccumulatedUsage instantiates a new AccumulatedUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccumulatedUsage() *AccumulatedUsage {
	this := AccumulatedUsage{}
	return &this
}

// NewAccumulatedUsageWithDefaults instantiates a new AccumulatedUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccumulatedUsageWithDefaults() *AccumulatedUsage {
	this := AccumulatedUsage{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AccumulatedUsage) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccumulatedUsage) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AccumulatedUsage) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AccumulatedUsage) SetDuration(v int32) {
	o.Duration = &v
}

// GetTotalVolume returns the TotalVolume field value if set, zero value otherwise.
func (o *AccumulatedUsage) GetTotalVolume() int64 {
	if o == nil || isNil(o.TotalVolume) {
		var ret int64
		return ret
	}
	return *o.TotalVolume
}

// GetTotalVolumeOk returns a tuple with the TotalVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccumulatedUsage) GetTotalVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.TotalVolume) {
    return nil, false
	}
	return o.TotalVolume, true
}

// HasTotalVolume returns a boolean if a field has been set.
func (o *AccumulatedUsage) HasTotalVolume() bool {
	if o != nil && !isNil(o.TotalVolume) {
		return true
	}

	return false
}

// SetTotalVolume gets a reference to the given int64 and assigns it to the TotalVolume field.
func (o *AccumulatedUsage) SetTotalVolume(v int64) {
	o.TotalVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *AccumulatedUsage) GetDownlinkVolume() int64 {
	if o == nil || isNil(o.DownlinkVolume) {
		var ret int64
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccumulatedUsage) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.DownlinkVolume) {
    return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *AccumulatedUsage) HasDownlinkVolume() bool {
	if o != nil && !isNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int64 and assigns it to the DownlinkVolume field.
func (o *AccumulatedUsage) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = &v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *AccumulatedUsage) GetUplinkVolume() int64 {
	if o == nil || isNil(o.UplinkVolume) {
		var ret int64
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccumulatedUsage) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.UplinkVolume) {
    return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *AccumulatedUsage) HasUplinkVolume() bool {
	if o != nil && !isNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int64 and assigns it to the UplinkVolume field.
func (o *AccumulatedUsage) SetUplinkVolume(v int64) {
	o.UplinkVolume = &v
}

func (o AccumulatedUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.TotalVolume) {
		toSerialize["totalVolume"] = o.TotalVolume
	}
	if !isNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	if !isNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	return json.Marshal(toSerialize)
}

type NullableAccumulatedUsage struct {
	value *AccumulatedUsage
	isSet bool
}

func (v NullableAccumulatedUsage) Get() *AccumulatedUsage {
	return v.value
}

func (v *NullableAccumulatedUsage) Set(val *AccumulatedUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableAccumulatedUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableAccumulatedUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccumulatedUsage(val *AccumulatedUsage) *NullableAccumulatedUsage {
	return &NullableAccumulatedUsage{value: val, isSet: true}
}

func (v NullableAccumulatedUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccumulatedUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


