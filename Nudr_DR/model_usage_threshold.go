/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// UsageThreshold Represents a usage threshold.
type UsageThreshold struct {
	// Unsigned integer identifying a period of time in units of seconds.
	Duration *int32 `json:"duration,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume *int64 `json:"totalVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
}

// NewUsageThreshold instantiates a new UsageThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageThreshold() *UsageThreshold {
	this := UsageThreshold{}
	return &this
}

// NewUsageThresholdWithDefaults instantiates a new UsageThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageThresholdWithDefaults() *UsageThreshold {
	this := UsageThreshold{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UsageThreshold) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageThreshold) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UsageThreshold) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *UsageThreshold) SetDuration(v int32) {
	o.Duration = &v
}

// GetTotalVolume returns the TotalVolume field value if set, zero value otherwise.
func (o *UsageThreshold) GetTotalVolume() int64 {
	if o == nil || isNil(o.TotalVolume) {
		var ret int64
		return ret
	}
	return *o.TotalVolume
}

// GetTotalVolumeOk returns a tuple with the TotalVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageThreshold) GetTotalVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.TotalVolume) {
    return nil, false
	}
	return o.TotalVolume, true
}

// HasTotalVolume returns a boolean if a field has been set.
func (o *UsageThreshold) HasTotalVolume() bool {
	if o != nil && !isNil(o.TotalVolume) {
		return true
	}

	return false
}

// SetTotalVolume gets a reference to the given int64 and assigns it to the TotalVolume field.
func (o *UsageThreshold) SetTotalVolume(v int64) {
	o.TotalVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *UsageThreshold) GetDownlinkVolume() int64 {
	if o == nil || isNil(o.DownlinkVolume) {
		var ret int64
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageThreshold) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.DownlinkVolume) {
    return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *UsageThreshold) HasDownlinkVolume() bool {
	if o != nil && !isNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int64 and assigns it to the DownlinkVolume field.
func (o *UsageThreshold) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = &v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *UsageThreshold) GetUplinkVolume() int64 {
	if o == nil || isNil(o.UplinkVolume) {
		var ret int64
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageThreshold) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.UplinkVolume) {
    return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *UsageThreshold) HasUplinkVolume() bool {
	if o != nil && !isNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int64 and assigns it to the UplinkVolume field.
func (o *UsageThreshold) SetUplinkVolume(v int64) {
	o.UplinkVolume = &v
}

func (o UsageThreshold) MarshalJSON() ([]byte, error) {
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

type NullableUsageThreshold struct {
	value *UsageThreshold
	isSet bool
}

func (v NullableUsageThreshold) Get() *UsageThreshold {
	return v.value
}

func (v *NullableUsageThreshold) Set(val *UsageThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageThreshold(val *UsageThreshold) *NullableUsageThreshold {
	return &NullableUsageThreshold{value: val, isSet: true}
}

func (v NullableUsageThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


