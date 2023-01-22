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

// PerfData Represents DN performance data.
type PerfData struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgTrafficRate *string `json:"avgTrafficRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxTrafficRate *string `json:"maxTrafficRate,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvePacketDelay *int32 `json:"avePacketDelay,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxPacketDelay *int32 `json:"maxPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPacketLossRate *int32 `json:"avgPacketLossRate,omitempty"`
}

// NewPerfData instantiates a new PerfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfData() *PerfData {
	this := PerfData{}
	return &this
}

// NewPerfDataWithDefaults instantiates a new PerfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfDataWithDefaults() *PerfData {
	this := PerfData{}
	return &this
}

// GetAvgTrafficRate returns the AvgTrafficRate field value if set, zero value otherwise.
func (o *PerfData) GetAvgTrafficRate() string {
	if o == nil || isNil(o.AvgTrafficRate) {
		var ret string
		return ret
	}
	return *o.AvgTrafficRate
}

// GetAvgTrafficRateOk returns a tuple with the AvgTrafficRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfData) GetAvgTrafficRateOk() (*string, bool) {
	if o == nil || isNil(o.AvgTrafficRate) {
    return nil, false
	}
	return o.AvgTrafficRate, true
}

// HasAvgTrafficRate returns a boolean if a field has been set.
func (o *PerfData) HasAvgTrafficRate() bool {
	if o != nil && !isNil(o.AvgTrafficRate) {
		return true
	}

	return false
}

// SetAvgTrafficRate gets a reference to the given string and assigns it to the AvgTrafficRate field.
func (o *PerfData) SetAvgTrafficRate(v string) {
	o.AvgTrafficRate = &v
}

// GetMaxTrafficRate returns the MaxTrafficRate field value if set, zero value otherwise.
func (o *PerfData) GetMaxTrafficRate() string {
	if o == nil || isNil(o.MaxTrafficRate) {
		var ret string
		return ret
	}
	return *o.MaxTrafficRate
}

// GetMaxTrafficRateOk returns a tuple with the MaxTrafficRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfData) GetMaxTrafficRateOk() (*string, bool) {
	if o == nil || isNil(o.MaxTrafficRate) {
    return nil, false
	}
	return o.MaxTrafficRate, true
}

// HasMaxTrafficRate returns a boolean if a field has been set.
func (o *PerfData) HasMaxTrafficRate() bool {
	if o != nil && !isNil(o.MaxTrafficRate) {
		return true
	}

	return false
}

// SetMaxTrafficRate gets a reference to the given string and assigns it to the MaxTrafficRate field.
func (o *PerfData) SetMaxTrafficRate(v string) {
	o.MaxTrafficRate = &v
}

// GetAvePacketDelay returns the AvePacketDelay field value if set, zero value otherwise.
func (o *PerfData) GetAvePacketDelay() int32 {
	if o == nil || isNil(o.AvePacketDelay) {
		var ret int32
		return ret
	}
	return *o.AvePacketDelay
}

// GetAvePacketDelayOk returns a tuple with the AvePacketDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfData) GetAvePacketDelayOk() (*int32, bool) {
	if o == nil || isNil(o.AvePacketDelay) {
    return nil, false
	}
	return o.AvePacketDelay, true
}

// HasAvePacketDelay returns a boolean if a field has been set.
func (o *PerfData) HasAvePacketDelay() bool {
	if o != nil && !isNil(o.AvePacketDelay) {
		return true
	}

	return false
}

// SetAvePacketDelay gets a reference to the given int32 and assigns it to the AvePacketDelay field.
func (o *PerfData) SetAvePacketDelay(v int32) {
	o.AvePacketDelay = &v
}

// GetMaxPacketDelay returns the MaxPacketDelay field value if set, zero value otherwise.
func (o *PerfData) GetMaxPacketDelay() int32 {
	if o == nil || isNil(o.MaxPacketDelay) {
		var ret int32
		return ret
	}
	return *o.MaxPacketDelay
}

// GetMaxPacketDelayOk returns a tuple with the MaxPacketDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfData) GetMaxPacketDelayOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPacketDelay) {
    return nil, false
	}
	return o.MaxPacketDelay, true
}

// HasMaxPacketDelay returns a boolean if a field has been set.
func (o *PerfData) HasMaxPacketDelay() bool {
	if o != nil && !isNil(o.MaxPacketDelay) {
		return true
	}

	return false
}

// SetMaxPacketDelay gets a reference to the given int32 and assigns it to the MaxPacketDelay field.
func (o *PerfData) SetMaxPacketDelay(v int32) {
	o.MaxPacketDelay = &v
}

// GetAvgPacketLossRate returns the AvgPacketLossRate field value if set, zero value otherwise.
func (o *PerfData) GetAvgPacketLossRate() int32 {
	if o == nil || isNil(o.AvgPacketLossRate) {
		var ret int32
		return ret
	}
	return *o.AvgPacketLossRate
}

// GetAvgPacketLossRateOk returns a tuple with the AvgPacketLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfData) GetAvgPacketLossRateOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPacketLossRate) {
    return nil, false
	}
	return o.AvgPacketLossRate, true
}

// HasAvgPacketLossRate returns a boolean if a field has been set.
func (o *PerfData) HasAvgPacketLossRate() bool {
	if o != nil && !isNil(o.AvgPacketLossRate) {
		return true
	}

	return false
}

// SetAvgPacketLossRate gets a reference to the given int32 and assigns it to the AvgPacketLossRate field.
func (o *PerfData) SetAvgPacketLossRate(v int32) {
	o.AvgPacketLossRate = &v
}

func (o PerfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AvgTrafficRate) {
		toSerialize["avgTrafficRate"] = o.AvgTrafficRate
	}
	if !isNil(o.MaxTrafficRate) {
		toSerialize["maxTrafficRate"] = o.MaxTrafficRate
	}
	if !isNil(o.AvePacketDelay) {
		toSerialize["avePacketDelay"] = o.AvePacketDelay
	}
	if !isNil(o.MaxPacketDelay) {
		toSerialize["maxPacketDelay"] = o.MaxPacketDelay
	}
	if !isNil(o.AvgPacketLossRate) {
		toSerialize["avgPacketLossRate"] = o.AvgPacketLossRate
	}
	return json.Marshal(toSerialize)
}

type NullablePerfData struct {
	value *PerfData
	isSet bool
}

func (v NullablePerfData) Get() *PerfData {
	return v.value
}

func (v *NullablePerfData) Set(val *PerfData) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfData) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfData(val *PerfData) *NullablePerfData {
	return &NullablePerfData{value: val, isSet: true}
}

func (v NullablePerfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


