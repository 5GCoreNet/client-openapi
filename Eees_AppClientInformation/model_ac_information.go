/*
EES Application Client Information_API

API for EES Application Client Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_AppClientInformation

import (
	"encoding/json"
)

// ACInformation AC Information matching the filter criteria.
type ACInformation struct {
	// List of profile information of ACs.
	AcProfs []ACProfile `json:"acProfs"`
	// List of UE identifiers hosting the AC.
	UeIds []string `json:"ueIds,omitempty"`
	UeLocInfs *LocationArea5G `json:"ueLocInfs,omitempty"`
}

// NewACInformation instantiates a new ACInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACInformation(acProfs []ACProfile) *ACInformation {
	this := ACInformation{}
	this.AcProfs = acProfs
	return &this
}

// NewACInformationWithDefaults instantiates a new ACInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACInformationWithDefaults() *ACInformation {
	this := ACInformation{}
	return &this
}

// GetAcProfs returns the AcProfs field value
func (o *ACInformation) GetAcProfs() []ACProfile {
	if o == nil {
		var ret []ACProfile
		return ret
	}

	return o.AcProfs
}

// GetAcProfsOk returns a tuple with the AcProfs field value
// and a boolean to check if the value has been set.
func (o *ACInformation) GetAcProfsOk() ([]ACProfile, bool) {
	if o == nil {
    return nil, false
	}
	return o.AcProfs, true
}

// SetAcProfs sets field value
func (o *ACInformation) SetAcProfs(v []ACProfile) {
	o.AcProfs = v
}

// GetUeIds returns the UeIds field value if set, zero value otherwise.
func (o *ACInformation) GetUeIds() []string {
	if o == nil || isNil(o.UeIds) {
		var ret []string
		return ret
	}
	return o.UeIds
}

// GetUeIdsOk returns a tuple with the UeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInformation) GetUeIdsOk() ([]string, bool) {
	if o == nil || isNil(o.UeIds) {
    return nil, false
	}
	return o.UeIds, true
}

// HasUeIds returns a boolean if a field has been set.
func (o *ACInformation) HasUeIds() bool {
	if o != nil && !isNil(o.UeIds) {
		return true
	}

	return false
}

// SetUeIds gets a reference to the given []string and assigns it to the UeIds field.
func (o *ACInformation) SetUeIds(v []string) {
	o.UeIds = v
}

// GetUeLocInfs returns the UeLocInfs field value if set, zero value otherwise.
func (o *ACInformation) GetUeLocInfs() LocationArea5G {
	if o == nil || isNil(o.UeLocInfs) {
		var ret LocationArea5G
		return ret
	}
	return *o.UeLocInfs
}

// GetUeLocInfsOk returns a tuple with the UeLocInfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInformation) GetUeLocInfsOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.UeLocInfs) {
    return nil, false
	}
	return o.UeLocInfs, true
}

// HasUeLocInfs returns a boolean if a field has been set.
func (o *ACInformation) HasUeLocInfs() bool {
	if o != nil && !isNil(o.UeLocInfs) {
		return true
	}

	return false
}

// SetUeLocInfs gets a reference to the given LocationArea5G and assigns it to the UeLocInfs field.
func (o *ACInformation) SetUeLocInfs(v LocationArea5G) {
	o.UeLocInfs = &v
}

func (o ACInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["acProfs"] = o.AcProfs
	}
	if !isNil(o.UeIds) {
		toSerialize["ueIds"] = o.UeIds
	}
	if !isNil(o.UeLocInfs) {
		toSerialize["ueLocInfs"] = o.UeLocInfs
	}
	return json.Marshal(toSerialize)
}

type NullableACInformation struct {
	value *ACInformation
	isSet bool
}

func (v NullableACInformation) Get() *ACInformation {
	return v.value
}

func (v *NullableACInformation) Set(val *ACInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableACInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableACInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACInformation(val *ACInformation) *NullableACInformation {
	return &NullableACInformation{value: val, isSet: true}
}

func (v NullableACInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


