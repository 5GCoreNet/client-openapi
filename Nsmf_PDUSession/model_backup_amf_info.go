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

// BackupAmfInfo Provides details of the Backup AMF.
type BackupAmfInfo struct {
	// Fully Qualified Domain Name
	BackupAmf string `json:"backupAmf"`
	// If present, this IE shall contain the list of GUAMI(s) (supported by the AMF) for which the backupAmf IE applies. 
	GuamiList []Guami `json:"guamiList,omitempty"`
}

// NewBackupAmfInfo instantiates a new BackupAmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupAmfInfo(backupAmf string) *BackupAmfInfo {
	this := BackupAmfInfo{}
	this.BackupAmf = backupAmf
	return &this
}

// NewBackupAmfInfoWithDefaults instantiates a new BackupAmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupAmfInfoWithDefaults() *BackupAmfInfo {
	this := BackupAmfInfo{}
	return &this
}

// GetBackupAmf returns the BackupAmf field value
func (o *BackupAmfInfo) GetBackupAmf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackupAmf
}

// GetBackupAmfOk returns a tuple with the BackupAmf field value
// and a boolean to check if the value has been set.
func (o *BackupAmfInfo) GetBackupAmfOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackupAmf, true
}

// SetBackupAmf sets field value
func (o *BackupAmfInfo) SetBackupAmf(v string) {
	o.BackupAmf = v
}

// GetGuamiList returns the GuamiList field value if set, zero value otherwise.
func (o *BackupAmfInfo) GetGuamiList() []Guami {
	if o == nil || isNil(o.GuamiList) {
		var ret []Guami
		return ret
	}
	return o.GuamiList
}

// GetGuamiListOk returns a tuple with the GuamiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupAmfInfo) GetGuamiListOk() ([]Guami, bool) {
	if o == nil || isNil(o.GuamiList) {
    return nil, false
	}
	return o.GuamiList, true
}

// HasGuamiList returns a boolean if a field has been set.
func (o *BackupAmfInfo) HasGuamiList() bool {
	if o != nil && !isNil(o.GuamiList) {
		return true
	}

	return false
}

// SetGuamiList gets a reference to the given []Guami and assigns it to the GuamiList field.
func (o *BackupAmfInfo) SetGuamiList(v []Guami) {
	o.GuamiList = v
}

func (o BackupAmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupAmf"] = o.BackupAmf
	}
	if !isNil(o.GuamiList) {
		toSerialize["guamiList"] = o.GuamiList
	}
	return json.Marshal(toSerialize)
}

type NullableBackupAmfInfo struct {
	value *BackupAmfInfo
	isSet bool
}

func (v NullableBackupAmfInfo) Get() *BackupAmfInfo {
	return v.value
}

func (v *NullableBackupAmfInfo) Set(val *BackupAmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupAmfInfo(val *BackupAmfInfo) *NullableBackupAmfInfo {
	return &NullableBackupAmfInfo{value: val, isSet: true}
}

func (v NullableBackupAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


