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

// EcRestriction struct for EcRestriction
type EcRestriction struct {
	AfInstanceId string `json:"afInstanceId"`
	ReferenceId int32 `json:"referenceId"`
	PlmnEcInfos []PlmnEcInfo `json:"plmnEcInfos,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

// NewEcRestriction instantiates a new EcRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcRestriction(afInstanceId string, referenceId int32) *EcRestriction {
	this := EcRestriction{}
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewEcRestrictionWithDefaults instantiates a new EcRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcRestrictionWithDefaults() *EcRestriction {
	this := EcRestriction{}
	return &this
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *EcRestriction) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *EcRestriction) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *EcRestriction) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *EcRestriction) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *EcRestriction) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *EcRestriction) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetPlmnEcInfos returns the PlmnEcInfos field value if set, zero value otherwise.
func (o *EcRestriction) GetPlmnEcInfos() []PlmnEcInfo {
	if o == nil || isNil(o.PlmnEcInfos) {
		var ret []PlmnEcInfo
		return ret
	}
	return o.PlmnEcInfos
}

// GetPlmnEcInfosOk returns a tuple with the PlmnEcInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcRestriction) GetPlmnEcInfosOk() ([]PlmnEcInfo, bool) {
	if o == nil || isNil(o.PlmnEcInfos) {
    return nil, false
	}
	return o.PlmnEcInfos, true
}

// HasPlmnEcInfos returns a boolean if a field has been set.
func (o *EcRestriction) HasPlmnEcInfos() bool {
	if o != nil && !isNil(o.PlmnEcInfos) {
		return true
	}

	return false
}

// SetPlmnEcInfos gets a reference to the given []PlmnEcInfo and assigns it to the PlmnEcInfos field.
func (o *EcRestriction) SetPlmnEcInfos(v []PlmnEcInfo) {
	o.PlmnEcInfos = v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *EcRestriction) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcRestriction) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
    return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *EcRestriction) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *EcRestriction) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o EcRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afInstanceId"] = o.AfInstanceId
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.PlmnEcInfos) {
		toSerialize["plmnEcInfos"] = o.PlmnEcInfos
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return json.Marshal(toSerialize)
}

type NullableEcRestriction struct {
	value *EcRestriction
	isSet bool
}

func (v NullableEcRestriction) Get() *EcRestriction {
	return v.value
}

func (v *NullableEcRestriction) Set(val *EcRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableEcRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableEcRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcRestriction(val *EcRestriction) *NullableEcRestriction {
	return &NullableEcRestriction{value: val, isSet: true}
}

func (v NullableEcRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


