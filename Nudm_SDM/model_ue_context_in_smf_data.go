/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// UeContextInSmfData struct for UeContextInSmfData
type UeContextInSmfData struct {
	// A map (list of key-value pairs where PduSessionId serves as key) of PduSessions
	PduSessions *map[string]PduSession `json:"pduSessions,omitempty"`
	PgwInfo []PgwInfo `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo `json:"emergencyInfo,omitempty"`
}

// NewUeContextInSmfData instantiates a new UeContextInSmfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInSmfData() *UeContextInSmfData {
	this := UeContextInSmfData{}
	return &this
}

// NewUeContextInSmfDataWithDefaults instantiates a new UeContextInSmfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInSmfDataWithDefaults() *UeContextInSmfData {
	this := UeContextInSmfData{}
	return &this
}

// GetPduSessions returns the PduSessions field value if set, zero value otherwise.
func (o *UeContextInSmfData) GetPduSessions() map[string]PduSession {
	if o == nil || isNil(o.PduSessions) {
		var ret map[string]PduSession
		return ret
	}
	return *o.PduSessions
}

// GetPduSessionsOk returns a tuple with the PduSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfData) GetPduSessionsOk() (*map[string]PduSession, bool) {
	if o == nil || isNil(o.PduSessions) {
    return nil, false
	}
	return o.PduSessions, true
}

// HasPduSessions returns a boolean if a field has been set.
func (o *UeContextInSmfData) HasPduSessions() bool {
	if o != nil && !isNil(o.PduSessions) {
		return true
	}

	return false
}

// SetPduSessions gets a reference to the given map[string]PduSession and assigns it to the PduSessions field.
func (o *UeContextInSmfData) SetPduSessions(v map[string]PduSession) {
	o.PduSessions = &v
}

// GetPgwInfo returns the PgwInfo field value if set, zero value otherwise.
func (o *UeContextInSmfData) GetPgwInfo() []PgwInfo {
	if o == nil || isNil(o.PgwInfo) {
		var ret []PgwInfo
		return ret
	}
	return o.PgwInfo
}

// GetPgwInfoOk returns a tuple with the PgwInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfData) GetPgwInfoOk() ([]PgwInfo, bool) {
	if o == nil || isNil(o.PgwInfo) {
    return nil, false
	}
	return o.PgwInfo, true
}

// HasPgwInfo returns a boolean if a field has been set.
func (o *UeContextInSmfData) HasPgwInfo() bool {
	if o != nil && !isNil(o.PgwInfo) {
		return true
	}

	return false
}

// SetPgwInfo gets a reference to the given []PgwInfo and assigns it to the PgwInfo field.
func (o *UeContextInSmfData) SetPgwInfo(v []PgwInfo) {
	o.PgwInfo = v
}

// GetEmergencyInfo returns the EmergencyInfo field value if set, zero value otherwise.
func (o *UeContextInSmfData) GetEmergencyInfo() EmergencyInfo {
	if o == nil || isNil(o.EmergencyInfo) {
		var ret EmergencyInfo
		return ret
	}
	return *o.EmergencyInfo
}

// GetEmergencyInfoOk returns a tuple with the EmergencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfData) GetEmergencyInfoOk() (*EmergencyInfo, bool) {
	if o == nil || isNil(o.EmergencyInfo) {
    return nil, false
	}
	return o.EmergencyInfo, true
}

// HasEmergencyInfo returns a boolean if a field has been set.
func (o *UeContextInSmfData) HasEmergencyInfo() bool {
	if o != nil && !isNil(o.EmergencyInfo) {
		return true
	}

	return false
}

// SetEmergencyInfo gets a reference to the given EmergencyInfo and assigns it to the EmergencyInfo field.
func (o *UeContextInSmfData) SetEmergencyInfo(v EmergencyInfo) {
	o.EmergencyInfo = &v
}

func (o UeContextInSmfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PduSessions) {
		toSerialize["pduSessions"] = o.PduSessions
	}
	if !isNil(o.PgwInfo) {
		toSerialize["pgwInfo"] = o.PgwInfo
	}
	if !isNil(o.EmergencyInfo) {
		toSerialize["emergencyInfo"] = o.EmergencyInfo
	}
	return json.Marshal(toSerialize)
}

type NullableUeContextInSmfData struct {
	value *UeContextInSmfData
	isSet bool
}

func (v NullableUeContextInSmfData) Get() *UeContextInSmfData {
	return v.value
}

func (v *NullableUeContextInSmfData) Set(val *UeContextInSmfData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInSmfData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInSmfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInSmfData(val *UeContextInSmfData) *NullableUeContextInSmfData {
	return &NullableUeContextInSmfData{value: val, isSet: true}
}

func (v NullableUeContextInSmfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInSmfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


