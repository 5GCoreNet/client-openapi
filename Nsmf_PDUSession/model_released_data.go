/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// ReleasedData Data within Release Response
type ReleasedData struct {
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus *ApnRateStatus `json:"apnRateStatus,omitempty"`
	N4Info *N4Information `json:"n4Info,omitempty"`
	N4InfoExt1 *N4Information `json:"n4InfoExt1,omitempty"`
	N4InfoExt2 *N4Information `json:"n4InfoExt2,omitempty"`
}

// NewReleasedData instantiates a new ReleasedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleasedData() *ReleasedData {
	this := ReleasedData{}
	return &this
}

// NewReleasedDataWithDefaults instantiates a new ReleasedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleasedDataWithDefaults() *ReleasedData {
	this := ReleasedData{}
	return &this
}

// GetSmallDataRateStatus returns the SmallDataRateStatus field value if set, zero value otherwise.
func (o *ReleasedData) GetSmallDataRateStatus() SmallDataRateStatus {
	if o == nil || isNil(o.SmallDataRateStatus) {
		var ret SmallDataRateStatus
		return ret
	}
	return *o.SmallDataRateStatus
}

// GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasedData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool) {
	if o == nil || isNil(o.SmallDataRateStatus) {
    return nil, false
	}
	return o.SmallDataRateStatus, true
}

// HasSmallDataRateStatus returns a boolean if a field has been set.
func (o *ReleasedData) HasSmallDataRateStatus() bool {
	if o != nil && !isNil(o.SmallDataRateStatus) {
		return true
	}

	return false
}

// SetSmallDataRateStatus gets a reference to the given SmallDataRateStatus and assigns it to the SmallDataRateStatus field.
func (o *ReleasedData) SetSmallDataRateStatus(v SmallDataRateStatus) {
	o.SmallDataRateStatus = &v
}

// GetApnRateStatus returns the ApnRateStatus field value if set, zero value otherwise.
func (o *ReleasedData) GetApnRateStatus() ApnRateStatus {
	if o == nil || isNil(o.ApnRateStatus) {
		var ret ApnRateStatus
		return ret
	}
	return *o.ApnRateStatus
}

// GetApnRateStatusOk returns a tuple with the ApnRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasedData) GetApnRateStatusOk() (*ApnRateStatus, bool) {
	if o == nil || isNil(o.ApnRateStatus) {
    return nil, false
	}
	return o.ApnRateStatus, true
}

// HasApnRateStatus returns a boolean if a field has been set.
func (o *ReleasedData) HasApnRateStatus() bool {
	if o != nil && !isNil(o.ApnRateStatus) {
		return true
	}

	return false
}

// SetApnRateStatus gets a reference to the given ApnRateStatus and assigns it to the ApnRateStatus field.
func (o *ReleasedData) SetApnRateStatus(v ApnRateStatus) {
	o.ApnRateStatus = &v
}

// GetN4Info returns the N4Info field value if set, zero value otherwise.
func (o *ReleasedData) GetN4Info() N4Information {
	if o == nil || isNil(o.N4Info) {
		var ret N4Information
		return ret
	}
	return *o.N4Info
}

// GetN4InfoOk returns a tuple with the N4Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasedData) GetN4InfoOk() (*N4Information, bool) {
	if o == nil || isNil(o.N4Info) {
    return nil, false
	}
	return o.N4Info, true
}

// HasN4Info returns a boolean if a field has been set.
func (o *ReleasedData) HasN4Info() bool {
	if o != nil && !isNil(o.N4Info) {
		return true
	}

	return false
}

// SetN4Info gets a reference to the given N4Information and assigns it to the N4Info field.
func (o *ReleasedData) SetN4Info(v N4Information) {
	o.N4Info = &v
}

// GetN4InfoExt1 returns the N4InfoExt1 field value if set, zero value otherwise.
func (o *ReleasedData) GetN4InfoExt1() N4Information {
	if o == nil || isNil(o.N4InfoExt1) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt1
}

// GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasedData) GetN4InfoExt1Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt1) {
    return nil, false
	}
	return o.N4InfoExt1, true
}

// HasN4InfoExt1 returns a boolean if a field has been set.
func (o *ReleasedData) HasN4InfoExt1() bool {
	if o != nil && !isNil(o.N4InfoExt1) {
		return true
	}

	return false
}

// SetN4InfoExt1 gets a reference to the given N4Information and assigns it to the N4InfoExt1 field.
func (o *ReleasedData) SetN4InfoExt1(v N4Information) {
	o.N4InfoExt1 = &v
}

// GetN4InfoExt2 returns the N4InfoExt2 field value if set, zero value otherwise.
func (o *ReleasedData) GetN4InfoExt2() N4Information {
	if o == nil || isNil(o.N4InfoExt2) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt2
}

// GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasedData) GetN4InfoExt2Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt2) {
    return nil, false
	}
	return o.N4InfoExt2, true
}

// HasN4InfoExt2 returns a boolean if a field has been set.
func (o *ReleasedData) HasN4InfoExt2() bool {
	if o != nil && !isNil(o.N4InfoExt2) {
		return true
	}

	return false
}

// SetN4InfoExt2 gets a reference to the given N4Information and assigns it to the N4InfoExt2 field.
func (o *ReleasedData) SetN4InfoExt2(v N4Information) {
	o.N4InfoExt2 = &v
}

func (o ReleasedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SmallDataRateStatus) {
		toSerialize["smallDataRateStatus"] = o.SmallDataRateStatus
	}
	if !isNil(o.ApnRateStatus) {
		toSerialize["apnRateStatus"] = o.ApnRateStatus
	}
	if !isNil(o.N4Info) {
		toSerialize["n4Info"] = o.N4Info
	}
	if !isNil(o.N4InfoExt1) {
		toSerialize["n4InfoExt1"] = o.N4InfoExt1
	}
	if !isNil(o.N4InfoExt2) {
		toSerialize["n4InfoExt2"] = o.N4InfoExt2
	}
	return json.Marshal(toSerialize)
}

type NullableReleasedData struct {
	value *ReleasedData
	isSet bool
}

func (v NullableReleasedData) Get() *ReleasedData {
	return v.value
}

func (v *NullableReleasedData) Set(val *ReleasedData) {
	v.value = val
	v.isSet = true
}

func (v NullableReleasedData) IsSet() bool {
	return v.isSet
}

func (v *NullableReleasedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleasedData(val *ReleasedData) *NullableReleasedData {
	return &NullableReleasedData{value: val, isSet: true}
}

func (v NullableReleasedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleasedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


