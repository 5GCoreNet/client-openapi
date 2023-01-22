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

// SmContextReleaseData Data within Release SM Context Request
type SmContextReleaseData struct {
	Cause *Cause `json:"cause,omitempty"`
	NgApCause *NgApCause `json:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue *int32 `json:"5gMmCauseValue,omitempty"`
	UeLocation *UserLocation `json:"ueLocation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`
	VsmfReleaseOnly *bool `json:"vsmfReleaseOnly,omitempty"`
	N2SmInfo *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType *N2SmInfoType `json:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly *bool `json:"ismfReleaseOnly,omitempty"`
}

// NewSmContextReleaseData instantiates a new SmContextReleaseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextReleaseData() *SmContextReleaseData {
	this := SmContextReleaseData{}
	var vsmfReleaseOnly bool = false
	this.VsmfReleaseOnly = &vsmfReleaseOnly
	var ismfReleaseOnly bool = false
	this.IsmfReleaseOnly = &ismfReleaseOnly
	return &this
}

// NewSmContextReleaseDataWithDefaults instantiates a new SmContextReleaseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextReleaseDataWithDefaults() *SmContextReleaseData {
	this := SmContextReleaseData{}
	var vsmfReleaseOnly bool = false
	this.VsmfReleaseOnly = &vsmfReleaseOnly
	var ismfReleaseOnly bool = false
	this.IsmfReleaseOnly = &ismfReleaseOnly
	return &this
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetCause() Cause {
	if o == nil || isNil(o.Cause) {
		var ret Cause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetCauseOk() (*Cause, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given Cause and assigns it to the Cause field.
func (o *SmContextReleaseData) SetCause(v Cause) {
	o.Cause = &v
}

// GetNgApCause returns the NgApCause field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetNgApCause() NgApCause {
	if o == nil || isNil(o.NgApCause) {
		var ret NgApCause
		return ret
	}
	return *o.NgApCause
}

// GetNgApCauseOk returns a tuple with the NgApCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetNgApCauseOk() (*NgApCause, bool) {
	if o == nil || isNil(o.NgApCause) {
    return nil, false
	}
	return o.NgApCause, true
}

// HasNgApCause returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasNgApCause() bool {
	if o != nil && !isNil(o.NgApCause) {
		return true
	}

	return false
}

// SetNgApCause gets a reference to the given NgApCause and assigns it to the NgApCause field.
func (o *SmContextReleaseData) SetNgApCause(v NgApCause) {
	o.NgApCause = &v
}

// GetVar5gMmCauseValue returns the Var5gMmCauseValue field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetVar5gMmCauseValue() int32 {
	if o == nil || isNil(o.Var5gMmCauseValue) {
		var ret int32
		return ret
	}
	return *o.Var5gMmCauseValue
}

// GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetVar5gMmCauseValueOk() (*int32, bool) {
	if o == nil || isNil(o.Var5gMmCauseValue) {
    return nil, false
	}
	return o.Var5gMmCauseValue, true
}

// HasVar5gMmCauseValue returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasVar5gMmCauseValue() bool {
	if o != nil && !isNil(o.Var5gMmCauseValue) {
		return true
	}

	return false
}

// SetVar5gMmCauseValue gets a reference to the given int32 and assigns it to the Var5gMmCauseValue field.
func (o *SmContextReleaseData) SetVar5gMmCauseValue(v int32) {
	o.Var5gMmCauseValue = &v
}

// GetUeLocation returns the UeLocation field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetUeLocation() UserLocation {
	if o == nil || isNil(o.UeLocation) {
		var ret UserLocation
		return ret
	}
	return *o.UeLocation
}

// GetUeLocationOk returns a tuple with the UeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetUeLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UeLocation) {
    return nil, false
	}
	return o.UeLocation, true
}

// HasUeLocation returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasUeLocation() bool {
	if o != nil && !isNil(o.UeLocation) {
		return true
	}

	return false
}

// SetUeLocation gets a reference to the given UserLocation and assigns it to the UeLocation field.
func (o *SmContextReleaseData) SetUeLocation(v UserLocation) {
	o.UeLocation = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetUeTimeZone() string {
	if o == nil || isNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.UeTimeZone) {
    return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasUeTimeZone() bool {
	if o != nil && !isNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *SmContextReleaseData) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetAddUeLocation returns the AddUeLocation field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetAddUeLocation() UserLocation {
	if o == nil || isNil(o.AddUeLocation) {
		var ret UserLocation
		return ret
	}
	return *o.AddUeLocation
}

// GetAddUeLocationOk returns a tuple with the AddUeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetAddUeLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.AddUeLocation) {
    return nil, false
	}
	return o.AddUeLocation, true
}

// HasAddUeLocation returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasAddUeLocation() bool {
	if o != nil && !isNil(o.AddUeLocation) {
		return true
	}

	return false
}

// SetAddUeLocation gets a reference to the given UserLocation and assigns it to the AddUeLocation field.
func (o *SmContextReleaseData) SetAddUeLocation(v UserLocation) {
	o.AddUeLocation = &v
}

// GetVsmfReleaseOnly returns the VsmfReleaseOnly field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetVsmfReleaseOnly() bool {
	if o == nil || isNil(o.VsmfReleaseOnly) {
		var ret bool
		return ret
	}
	return *o.VsmfReleaseOnly
}

// GetVsmfReleaseOnlyOk returns a tuple with the VsmfReleaseOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetVsmfReleaseOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.VsmfReleaseOnly) {
    return nil, false
	}
	return o.VsmfReleaseOnly, true
}

// HasVsmfReleaseOnly returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasVsmfReleaseOnly() bool {
	if o != nil && !isNil(o.VsmfReleaseOnly) {
		return true
	}

	return false
}

// SetVsmfReleaseOnly gets a reference to the given bool and assigns it to the VsmfReleaseOnly field.
func (o *SmContextReleaseData) SetVsmfReleaseOnly(v bool) {
	o.VsmfReleaseOnly = &v
}

// GetN2SmInfo returns the N2SmInfo field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetN2SmInfo() RefToBinaryData {
	if o == nil || isNil(o.N2SmInfo) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N2SmInfo
}

// GetN2SmInfoOk returns a tuple with the N2SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetN2SmInfoOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.N2SmInfo) {
    return nil, false
	}
	return o.N2SmInfo, true
}

// HasN2SmInfo returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasN2SmInfo() bool {
	if o != nil && !isNil(o.N2SmInfo) {
		return true
	}

	return false
}

// SetN2SmInfo gets a reference to the given RefToBinaryData and assigns it to the N2SmInfo field.
func (o *SmContextReleaseData) SetN2SmInfo(v RefToBinaryData) {
	o.N2SmInfo = &v
}

// GetN2SmInfoType returns the N2SmInfoType field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetN2SmInfoType() N2SmInfoType {
	if o == nil || isNil(o.N2SmInfoType) {
		var ret N2SmInfoType
		return ret
	}
	return *o.N2SmInfoType
}

// GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool) {
	if o == nil || isNil(o.N2SmInfoType) {
    return nil, false
	}
	return o.N2SmInfoType, true
}

// HasN2SmInfoType returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasN2SmInfoType() bool {
	if o != nil && !isNil(o.N2SmInfoType) {
		return true
	}

	return false
}

// SetN2SmInfoType gets a reference to the given N2SmInfoType and assigns it to the N2SmInfoType field.
func (o *SmContextReleaseData) SetN2SmInfoType(v N2SmInfoType) {
	o.N2SmInfoType = &v
}

// GetIsmfReleaseOnly returns the IsmfReleaseOnly field value if set, zero value otherwise.
func (o *SmContextReleaseData) GetIsmfReleaseOnly() bool {
	if o == nil || isNil(o.IsmfReleaseOnly) {
		var ret bool
		return ret
	}
	return *o.IsmfReleaseOnly
}

// GetIsmfReleaseOnlyOk returns a tuple with the IsmfReleaseOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetIsmfReleaseOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.IsmfReleaseOnly) {
    return nil, false
	}
	return o.IsmfReleaseOnly, true
}

// HasIsmfReleaseOnly returns a boolean if a field has been set.
func (o *SmContextReleaseData) HasIsmfReleaseOnly() bool {
	if o != nil && !isNil(o.IsmfReleaseOnly) {
		return true
	}

	return false
}

// SetIsmfReleaseOnly gets a reference to the given bool and assigns it to the IsmfReleaseOnly field.
func (o *SmContextReleaseData) SetIsmfReleaseOnly(v bool) {
	o.IsmfReleaseOnly = &v
}

func (o SmContextReleaseData) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.VsmfReleaseOnly) {
		toSerialize["vsmfReleaseOnly"] = o.VsmfReleaseOnly
	}
	if !isNil(o.N2SmInfo) {
		toSerialize["n2SmInfo"] = o.N2SmInfo
	}
	if !isNil(o.N2SmInfoType) {
		toSerialize["n2SmInfoType"] = o.N2SmInfoType
	}
	if !isNil(o.IsmfReleaseOnly) {
		toSerialize["ismfReleaseOnly"] = o.IsmfReleaseOnly
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextReleaseData struct {
	value *SmContextReleaseData
	isSet bool
}

func (v NullableSmContextReleaseData) Get() *SmContextReleaseData {
	return v.value
}

func (v *NullableSmContextReleaseData) Set(val *SmContextReleaseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextReleaseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextReleaseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextReleaseData(val *SmContextReleaseData) *NullableSmContextReleaseData {
	return &NullableSmContextReleaseData{value: val, isSet: true}
}

func (v NullableSmContextReleaseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextReleaseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


