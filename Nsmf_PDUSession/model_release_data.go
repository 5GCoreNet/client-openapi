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

// ReleaseData Data within Release Request
type ReleaseData struct {
	Cause *Cause `json:"cause,omitempty"`
	NgApCause *NgApCause `json:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue *int32 `json:"5gMmCauseValue,omitempty"`
	UeLocation *UserLocation `json:"ueLocation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	SecondaryRatUsageInfo []SecondaryRatUsageInfo `json:"secondaryRatUsageInfo,omitempty"`
	N4Info *N4Information `json:"n4Info,omitempty"`
	N4InfoExt1 *N4Information `json:"n4InfoExt1,omitempty"`
	N4InfoExt2 *N4Information `json:"n4InfoExt2,omitempty"`
}

// NewReleaseData instantiates a new ReleaseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseData() *ReleaseData {
	this := ReleaseData{}
	return &this
}

// NewReleaseDataWithDefaults instantiates a new ReleaseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseDataWithDefaults() *ReleaseData {
	this := ReleaseData{}
	return &this
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ReleaseData) GetCause() Cause {
	if o == nil || isNil(o.Cause) {
		var ret Cause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetCauseOk() (*Cause, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ReleaseData) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given Cause and assigns it to the Cause field.
func (o *ReleaseData) SetCause(v Cause) {
	o.Cause = &v
}

// GetNgApCause returns the NgApCause field value if set, zero value otherwise.
func (o *ReleaseData) GetNgApCause() NgApCause {
	if o == nil || isNil(o.NgApCause) {
		var ret NgApCause
		return ret
	}
	return *o.NgApCause
}

// GetNgApCauseOk returns a tuple with the NgApCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetNgApCauseOk() (*NgApCause, bool) {
	if o == nil || isNil(o.NgApCause) {
    return nil, false
	}
	return o.NgApCause, true
}

// HasNgApCause returns a boolean if a field has been set.
func (o *ReleaseData) HasNgApCause() bool {
	if o != nil && !isNil(o.NgApCause) {
		return true
	}

	return false
}

// SetNgApCause gets a reference to the given NgApCause and assigns it to the NgApCause field.
func (o *ReleaseData) SetNgApCause(v NgApCause) {
	o.NgApCause = &v
}

// GetVar5gMmCauseValue returns the Var5gMmCauseValue field value if set, zero value otherwise.
func (o *ReleaseData) GetVar5gMmCauseValue() int32 {
	if o == nil || isNil(o.Var5gMmCauseValue) {
		var ret int32
		return ret
	}
	return *o.Var5gMmCauseValue
}

// GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetVar5gMmCauseValueOk() (*int32, bool) {
	if o == nil || isNil(o.Var5gMmCauseValue) {
    return nil, false
	}
	return o.Var5gMmCauseValue, true
}

// HasVar5gMmCauseValue returns a boolean if a field has been set.
func (o *ReleaseData) HasVar5gMmCauseValue() bool {
	if o != nil && !isNil(o.Var5gMmCauseValue) {
		return true
	}

	return false
}

// SetVar5gMmCauseValue gets a reference to the given int32 and assigns it to the Var5gMmCauseValue field.
func (o *ReleaseData) SetVar5gMmCauseValue(v int32) {
	o.Var5gMmCauseValue = &v
}

// GetUeLocation returns the UeLocation field value if set, zero value otherwise.
func (o *ReleaseData) GetUeLocation() UserLocation {
	if o == nil || isNil(o.UeLocation) {
		var ret UserLocation
		return ret
	}
	return *o.UeLocation
}

// GetUeLocationOk returns a tuple with the UeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetUeLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UeLocation) {
    return nil, false
	}
	return o.UeLocation, true
}

// HasUeLocation returns a boolean if a field has been set.
func (o *ReleaseData) HasUeLocation() bool {
	if o != nil && !isNil(o.UeLocation) {
		return true
	}

	return false
}

// SetUeLocation gets a reference to the given UserLocation and assigns it to the UeLocation field.
func (o *ReleaseData) SetUeLocation(v UserLocation) {
	o.UeLocation = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *ReleaseData) GetUeTimeZone() string {
	if o == nil || isNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.UeTimeZone) {
    return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *ReleaseData) HasUeTimeZone() bool {
	if o != nil && !isNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *ReleaseData) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetAddUeLocation returns the AddUeLocation field value if set, zero value otherwise.
func (o *ReleaseData) GetAddUeLocation() UserLocation {
	if o == nil || isNil(o.AddUeLocation) {
		var ret UserLocation
		return ret
	}
	return *o.AddUeLocation
}

// GetAddUeLocationOk returns a tuple with the AddUeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetAddUeLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.AddUeLocation) {
    return nil, false
	}
	return o.AddUeLocation, true
}

// HasAddUeLocation returns a boolean if a field has been set.
func (o *ReleaseData) HasAddUeLocation() bool {
	if o != nil && !isNil(o.AddUeLocation) {
		return true
	}

	return false
}

// SetAddUeLocation gets a reference to the given UserLocation and assigns it to the AddUeLocation field.
func (o *ReleaseData) SetAddUeLocation(v UserLocation) {
	o.AddUeLocation = &v
}

// GetSecondaryRatUsageReport returns the SecondaryRatUsageReport field value if set, zero value otherwise.
func (o *ReleaseData) GetSecondaryRatUsageReport() []SecondaryRatUsageReport {
	if o == nil || isNil(o.SecondaryRatUsageReport) {
		var ret []SecondaryRatUsageReport
		return ret
	}
	return o.SecondaryRatUsageReport
}

// GetSecondaryRatUsageReportOk returns a tuple with the SecondaryRatUsageReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetSecondaryRatUsageReportOk() ([]SecondaryRatUsageReport, bool) {
	if o == nil || isNil(o.SecondaryRatUsageReport) {
    return nil, false
	}
	return o.SecondaryRatUsageReport, true
}

// HasSecondaryRatUsageReport returns a boolean if a field has been set.
func (o *ReleaseData) HasSecondaryRatUsageReport() bool {
	if o != nil && !isNil(o.SecondaryRatUsageReport) {
		return true
	}

	return false
}

// SetSecondaryRatUsageReport gets a reference to the given []SecondaryRatUsageReport and assigns it to the SecondaryRatUsageReport field.
func (o *ReleaseData) SetSecondaryRatUsageReport(v []SecondaryRatUsageReport) {
	o.SecondaryRatUsageReport = v
}

// GetSecondaryRatUsageInfo returns the SecondaryRatUsageInfo field value if set, zero value otherwise.
func (o *ReleaseData) GetSecondaryRatUsageInfo() []SecondaryRatUsageInfo {
	if o == nil || isNil(o.SecondaryRatUsageInfo) {
		var ret []SecondaryRatUsageInfo
		return ret
	}
	return o.SecondaryRatUsageInfo
}

// GetSecondaryRatUsageInfoOk returns a tuple with the SecondaryRatUsageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetSecondaryRatUsageInfoOk() ([]SecondaryRatUsageInfo, bool) {
	if o == nil || isNil(o.SecondaryRatUsageInfo) {
    return nil, false
	}
	return o.SecondaryRatUsageInfo, true
}

// HasSecondaryRatUsageInfo returns a boolean if a field has been set.
func (o *ReleaseData) HasSecondaryRatUsageInfo() bool {
	if o != nil && !isNil(o.SecondaryRatUsageInfo) {
		return true
	}

	return false
}

// SetSecondaryRatUsageInfo gets a reference to the given []SecondaryRatUsageInfo and assigns it to the SecondaryRatUsageInfo field.
func (o *ReleaseData) SetSecondaryRatUsageInfo(v []SecondaryRatUsageInfo) {
	o.SecondaryRatUsageInfo = v
}

// GetN4Info returns the N4Info field value if set, zero value otherwise.
func (o *ReleaseData) GetN4Info() N4Information {
	if o == nil || isNil(o.N4Info) {
		var ret N4Information
		return ret
	}
	return *o.N4Info
}

// GetN4InfoOk returns a tuple with the N4Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetN4InfoOk() (*N4Information, bool) {
	if o == nil || isNil(o.N4Info) {
    return nil, false
	}
	return o.N4Info, true
}

// HasN4Info returns a boolean if a field has been set.
func (o *ReleaseData) HasN4Info() bool {
	if o != nil && !isNil(o.N4Info) {
		return true
	}

	return false
}

// SetN4Info gets a reference to the given N4Information and assigns it to the N4Info field.
func (o *ReleaseData) SetN4Info(v N4Information) {
	o.N4Info = &v
}

// GetN4InfoExt1 returns the N4InfoExt1 field value if set, zero value otherwise.
func (o *ReleaseData) GetN4InfoExt1() N4Information {
	if o == nil || isNil(o.N4InfoExt1) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt1
}

// GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetN4InfoExt1Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt1) {
    return nil, false
	}
	return o.N4InfoExt1, true
}

// HasN4InfoExt1 returns a boolean if a field has been set.
func (o *ReleaseData) HasN4InfoExt1() bool {
	if o != nil && !isNil(o.N4InfoExt1) {
		return true
	}

	return false
}

// SetN4InfoExt1 gets a reference to the given N4Information and assigns it to the N4InfoExt1 field.
func (o *ReleaseData) SetN4InfoExt1(v N4Information) {
	o.N4InfoExt1 = &v
}

// GetN4InfoExt2 returns the N4InfoExt2 field value if set, zero value otherwise.
func (o *ReleaseData) GetN4InfoExt2() N4Information {
	if o == nil || isNil(o.N4InfoExt2) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt2
}

// GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseData) GetN4InfoExt2Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt2) {
    return nil, false
	}
	return o.N4InfoExt2, true
}

// HasN4InfoExt2 returns a boolean if a field has been set.
func (o *ReleaseData) HasN4InfoExt2() bool {
	if o != nil && !isNil(o.N4InfoExt2) {
		return true
	}

	return false
}

// SetN4InfoExt2 gets a reference to the given N4Information and assigns it to the N4InfoExt2 field.
func (o *ReleaseData) SetN4InfoExt2(v N4Information) {
	o.N4InfoExt2 = &v
}

func (o ReleaseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !isNil(o.NgApCause) {
		toSerialize["ngApCause"] = o.NgApCause
	}
	if !isNil(o.Var5gMmCauseValue) {
		toSerialize["5gMmCauseValue"] = o.Var5gMmCauseValue
	}
	if !isNil(o.UeLocation) {
		toSerialize["ueLocation"] = o.UeLocation
	}
	if !isNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !isNil(o.AddUeLocation) {
		toSerialize["addUeLocation"] = o.AddUeLocation
	}
	if !isNil(o.SecondaryRatUsageReport) {
		toSerialize["secondaryRatUsageReport"] = o.SecondaryRatUsageReport
	}
	if !isNil(o.SecondaryRatUsageInfo) {
		toSerialize["secondaryRatUsageInfo"] = o.SecondaryRatUsageInfo
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

type NullableReleaseData struct {
	value *ReleaseData
	isSet bool
}

func (v NullableReleaseData) Get() *ReleaseData {
	return v.value
}

func (v *NullableReleaseData) Set(val *ReleaseData) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseData) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseData(val *ReleaseData) *NullableReleaseData {
	return &NullableReleaseData{value: val, isSet: true}
}

func (v NullableReleaseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


