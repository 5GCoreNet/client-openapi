/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// MBSDistSessionAnmt Represents the set of MBS Distribution Session Announcements currently associated with this  MBS User Service Announcement. 
type MBSDistSessionAnmt struct {
	MbsSessionId *MbsSessionId `json:"mbsSessionId,omitempty"`
	// MBS Frequency Selection Area Identifier
	MbsFSAId *string `json:"mbsFSAId,omitempty"`
	DistrMethod DistributionMethod `json:"distrMethod"`
	ObjDistrAnnInfo *ObjectDistMethAnmtInfo `json:"objDistrAnnInfo,omitempty"`
	SesDesInfo []string `json:"sesDesInfo"`
}

// NewMBSDistSessionAnmt instantiates a new MBSDistSessionAnmt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSDistSessionAnmt(distrMethod DistributionMethod, sesDesInfo []string) *MBSDistSessionAnmt {
	this := MBSDistSessionAnmt{}
	this.DistrMethod = distrMethod
	this.SesDesInfo = sesDesInfo
	return &this
}

// NewMBSDistSessionAnmtWithDefaults instantiates a new MBSDistSessionAnmt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSDistSessionAnmtWithDefaults() *MBSDistSessionAnmt {
	this := MBSDistSessionAnmt{}
	return &this
}

// GetMbsSessionId returns the MbsSessionId field value if set, zero value otherwise.
func (o *MBSDistSessionAnmt) GetMbsSessionId() MbsSessionId {
	if o == nil || isNil(o.MbsSessionId) {
		var ret MbsSessionId
		return ret
	}
	return *o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistSessionAnmt) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil || isNil(o.MbsSessionId) {
    return nil, false
	}
	return o.MbsSessionId, true
}

// HasMbsSessionId returns a boolean if a field has been set.
func (o *MBSDistSessionAnmt) HasMbsSessionId() bool {
	if o != nil && !isNil(o.MbsSessionId) {
		return true
	}

	return false
}

// SetMbsSessionId gets a reference to the given MbsSessionId and assigns it to the MbsSessionId field.
func (o *MBSDistSessionAnmt) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = &v
}

// GetMbsFSAId returns the MbsFSAId field value if set, zero value otherwise.
func (o *MBSDistSessionAnmt) GetMbsFSAId() string {
	if o == nil || isNil(o.MbsFSAId) {
		var ret string
		return ret
	}
	return *o.MbsFSAId
}

// GetMbsFSAIdOk returns a tuple with the MbsFSAId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistSessionAnmt) GetMbsFSAIdOk() (*string, bool) {
	if o == nil || isNil(o.MbsFSAId) {
    return nil, false
	}
	return o.MbsFSAId, true
}

// HasMbsFSAId returns a boolean if a field has been set.
func (o *MBSDistSessionAnmt) HasMbsFSAId() bool {
	if o != nil && !isNil(o.MbsFSAId) {
		return true
	}

	return false
}

// SetMbsFSAId gets a reference to the given string and assigns it to the MbsFSAId field.
func (o *MBSDistSessionAnmt) SetMbsFSAId(v string) {
	o.MbsFSAId = &v
}

// GetDistrMethod returns the DistrMethod field value
func (o *MBSDistSessionAnmt) GetDistrMethod() DistributionMethod {
	if o == nil {
		var ret DistributionMethod
		return ret
	}

	return o.DistrMethod
}

// GetDistrMethodOk returns a tuple with the DistrMethod field value
// and a boolean to check if the value has been set.
func (o *MBSDistSessionAnmt) GetDistrMethodOk() (*DistributionMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DistrMethod, true
}

// SetDistrMethod sets field value
func (o *MBSDistSessionAnmt) SetDistrMethod(v DistributionMethod) {
	o.DistrMethod = v
}

// GetObjDistrAnnInfo returns the ObjDistrAnnInfo field value if set, zero value otherwise.
func (o *MBSDistSessionAnmt) GetObjDistrAnnInfo() ObjectDistMethAnmtInfo {
	if o == nil || isNil(o.ObjDistrAnnInfo) {
		var ret ObjectDistMethAnmtInfo
		return ret
	}
	return *o.ObjDistrAnnInfo
}

// GetObjDistrAnnInfoOk returns a tuple with the ObjDistrAnnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistSessionAnmt) GetObjDistrAnnInfoOk() (*ObjectDistMethAnmtInfo, bool) {
	if o == nil || isNil(o.ObjDistrAnnInfo) {
    return nil, false
	}
	return o.ObjDistrAnnInfo, true
}

// HasObjDistrAnnInfo returns a boolean if a field has been set.
func (o *MBSDistSessionAnmt) HasObjDistrAnnInfo() bool {
	if o != nil && !isNil(o.ObjDistrAnnInfo) {
		return true
	}

	return false
}

// SetObjDistrAnnInfo gets a reference to the given ObjectDistMethAnmtInfo and assigns it to the ObjDistrAnnInfo field.
func (o *MBSDistSessionAnmt) SetObjDistrAnnInfo(v ObjectDistMethAnmtInfo) {
	o.ObjDistrAnnInfo = &v
}

// GetSesDesInfo returns the SesDesInfo field value
func (o *MBSDistSessionAnmt) GetSesDesInfo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SesDesInfo
}

// GetSesDesInfoOk returns a tuple with the SesDesInfo field value
// and a boolean to check if the value has been set.
func (o *MBSDistSessionAnmt) GetSesDesInfoOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SesDesInfo, true
}

// SetSesDesInfo sets field value
func (o *MBSDistSessionAnmt) SetSesDesInfo(v []string) {
	o.SesDesInfo = v
}

func (o MBSDistSessionAnmt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsSessionId) {
		toSerialize["mbsSessionId"] = o.MbsSessionId
	}
	if !isNil(o.MbsFSAId) {
		toSerialize["mbsFSAId"] = o.MbsFSAId
	}
	if true {
		toSerialize["distrMethod"] = o.DistrMethod
	}
	if !isNil(o.ObjDistrAnnInfo) {
		toSerialize["objDistrAnnInfo"] = o.ObjDistrAnnInfo
	}
	if true {
		toSerialize["sesDesInfo"] = o.SesDesInfo
	}
	return json.Marshal(toSerialize)
}

type NullableMBSDistSessionAnmt struct {
	value *MBSDistSessionAnmt
	isSet bool
}

func (v NullableMBSDistSessionAnmt) Get() *MBSDistSessionAnmt {
	return v.value
}

func (v *NullableMBSDistSessionAnmt) Set(val *MBSDistSessionAnmt) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSDistSessionAnmt) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSDistSessionAnmt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSDistSessionAnmt(val *MBSDistSessionAnmt) *NullableMBSDistSessionAnmt {
	return &NullableMBSDistSessionAnmt{value: val, isSet: true}
}

func (v NullableMBSDistSessionAnmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSDistSessionAnmt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


