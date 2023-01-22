/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_SMContext

import (
	"encoding/json"
	"time"
)

// SmallDataRateStatus It indicates theSmall Data Rate Control Status
type SmallDataRateStatus struct {
	// When present, it shall contain the number of packets the UE is allowed to send uplink in the given time unit for the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainPacketsUl *int32 `json:"remainPacketsUl,omitempty"`
	// When present it shall contain the number of packets the AF is allowed to send downlink in the given time unit for the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainPacketsDl *int32 `json:"remainPacketsDl,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// When present, it shall indicate number of additional exception reports the UE is allowed to send uplink in the given time  unit for  the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainExReportsUl *int32 `json:"remainExReportsUl,omitempty"`
	// When present, it shall indicate number of additional exception reports the AF is allowed to send downlink  in the given time unit for the given PDU session (see clause 5.31.14.3 in 3GPP TS 23.501 
	RemainExReportsDl *int32 `json:"remainExReportsDl,omitempty"`
}

// NewSmallDataRateStatus instantiates a new SmallDataRateStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmallDataRateStatus() *SmallDataRateStatus {
	this := SmallDataRateStatus{}
	return &this
}

// NewSmallDataRateStatusWithDefaults instantiates a new SmallDataRateStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmallDataRateStatusWithDefaults() *SmallDataRateStatus {
	this := SmallDataRateStatus{}
	return &this
}

// GetRemainPacketsUl returns the RemainPacketsUl field value if set, zero value otherwise.
func (o *SmallDataRateStatus) GetRemainPacketsUl() int32 {
	if o == nil || isNil(o.RemainPacketsUl) {
		var ret int32
		return ret
	}
	return *o.RemainPacketsUl
}

// GetRemainPacketsUlOk returns a tuple with the RemainPacketsUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmallDataRateStatus) GetRemainPacketsUlOk() (*int32, bool) {
	if o == nil || isNil(o.RemainPacketsUl) {
    return nil, false
	}
	return o.RemainPacketsUl, true
}

// HasRemainPacketsUl returns a boolean if a field has been set.
func (o *SmallDataRateStatus) HasRemainPacketsUl() bool {
	if o != nil && !isNil(o.RemainPacketsUl) {
		return true
	}

	return false
}

// SetRemainPacketsUl gets a reference to the given int32 and assigns it to the RemainPacketsUl field.
func (o *SmallDataRateStatus) SetRemainPacketsUl(v int32) {
	o.RemainPacketsUl = &v
}

// GetRemainPacketsDl returns the RemainPacketsDl field value if set, zero value otherwise.
func (o *SmallDataRateStatus) GetRemainPacketsDl() int32 {
	if o == nil || isNil(o.RemainPacketsDl) {
		var ret int32
		return ret
	}
	return *o.RemainPacketsDl
}

// GetRemainPacketsDlOk returns a tuple with the RemainPacketsDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmallDataRateStatus) GetRemainPacketsDlOk() (*int32, bool) {
	if o == nil || isNil(o.RemainPacketsDl) {
    return nil, false
	}
	return o.RemainPacketsDl, true
}

// HasRemainPacketsDl returns a boolean if a field has been set.
func (o *SmallDataRateStatus) HasRemainPacketsDl() bool {
	if o != nil && !isNil(o.RemainPacketsDl) {
		return true
	}

	return false
}

// SetRemainPacketsDl gets a reference to the given int32 and assigns it to the RemainPacketsDl field.
func (o *SmallDataRateStatus) SetRemainPacketsDl(v int32) {
	o.RemainPacketsDl = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *SmallDataRateStatus) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmallDataRateStatus) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *SmallDataRateStatus) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *SmallDataRateStatus) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetRemainExReportsUl returns the RemainExReportsUl field value if set, zero value otherwise.
func (o *SmallDataRateStatus) GetRemainExReportsUl() int32 {
	if o == nil || isNil(o.RemainExReportsUl) {
		var ret int32
		return ret
	}
	return *o.RemainExReportsUl
}

// GetRemainExReportsUlOk returns a tuple with the RemainExReportsUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmallDataRateStatus) GetRemainExReportsUlOk() (*int32, bool) {
	if o == nil || isNil(o.RemainExReportsUl) {
    return nil, false
	}
	return o.RemainExReportsUl, true
}

// HasRemainExReportsUl returns a boolean if a field has been set.
func (o *SmallDataRateStatus) HasRemainExReportsUl() bool {
	if o != nil && !isNil(o.RemainExReportsUl) {
		return true
	}

	return false
}

// SetRemainExReportsUl gets a reference to the given int32 and assigns it to the RemainExReportsUl field.
func (o *SmallDataRateStatus) SetRemainExReportsUl(v int32) {
	o.RemainExReportsUl = &v
}

// GetRemainExReportsDl returns the RemainExReportsDl field value if set, zero value otherwise.
func (o *SmallDataRateStatus) GetRemainExReportsDl() int32 {
	if o == nil || isNil(o.RemainExReportsDl) {
		var ret int32
		return ret
	}
	return *o.RemainExReportsDl
}

// GetRemainExReportsDlOk returns a tuple with the RemainExReportsDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmallDataRateStatus) GetRemainExReportsDlOk() (*int32, bool) {
	if o == nil || isNil(o.RemainExReportsDl) {
    return nil, false
	}
	return o.RemainExReportsDl, true
}

// HasRemainExReportsDl returns a boolean if a field has been set.
func (o *SmallDataRateStatus) HasRemainExReportsDl() bool {
	if o != nil && !isNil(o.RemainExReportsDl) {
		return true
	}

	return false
}

// SetRemainExReportsDl gets a reference to the given int32 and assigns it to the RemainExReportsDl field.
func (o *SmallDataRateStatus) SetRemainExReportsDl(v int32) {
	o.RemainExReportsDl = &v
}

func (o SmallDataRateStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainPacketsUl) {
		toSerialize["remainPacketsUl"] = o.RemainPacketsUl
	}
	if !isNil(o.RemainPacketsDl) {
		toSerialize["remainPacketsDl"] = o.RemainPacketsDl
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !isNil(o.RemainExReportsUl) {
		toSerialize["remainExReportsUl"] = o.RemainExReportsUl
	}
	if !isNil(o.RemainExReportsDl) {
		toSerialize["remainExReportsDl"] = o.RemainExReportsDl
	}
	return json.Marshal(toSerialize)
}

type NullableSmallDataRateStatus struct {
	value *SmallDataRateStatus
	isSet bool
}

func (v NullableSmallDataRateStatus) Get() *SmallDataRateStatus {
	return v.value
}

func (v *NullableSmallDataRateStatus) Set(val *SmallDataRateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSmallDataRateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSmallDataRateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmallDataRateStatus(val *SmallDataRateStatus) *NullableSmallDataRateStatus {
	return &NullableSmallDataRateStatus{value: val, isSet: true}
}

func (v NullableSmallDataRateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmallDataRateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


