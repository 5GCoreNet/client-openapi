/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// AdditionalAccessInfo Indicates the combination of additional Access Type and RAT Type for a MA PDU session.
type AdditionalAccessInfo struct {
	AccessType AccessType `json:"accessType"`
	RatType *RatType `json:"ratType,omitempty"`
}

// NewAdditionalAccessInfo instantiates a new AdditionalAccessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalAccessInfo(accessType AccessType) *AdditionalAccessInfo {
	this := AdditionalAccessInfo{}
	this.AccessType = accessType
	return &this
}

// NewAdditionalAccessInfoWithDefaults instantiates a new AdditionalAccessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalAccessInfoWithDefaults() *AdditionalAccessInfo {
	this := AdditionalAccessInfo{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *AdditionalAccessInfo) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *AdditionalAccessInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *AdditionalAccessInfo) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *AdditionalAccessInfo) GetRatType() RatType {
	if o == nil || isNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalAccessInfo) GetRatTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RatType) {
    return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *AdditionalAccessInfo) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *AdditionalAccessInfo) SetRatType(v RatType) {
	o.RatType = &v
}

func (o AdditionalAccessInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalAccessInfo struct {
	value *AdditionalAccessInfo
	isSet bool
}

func (v NullableAdditionalAccessInfo) Get() *AdditionalAccessInfo {
	return v.value
}

func (v *NullableAdditionalAccessInfo) Set(val *AdditionalAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalAccessInfo(val *AdditionalAccessInfo) *NullableAdditionalAccessInfo {
	return &NullableAdditionalAccessInfo{value: val, isSet: true}
}

func (v NullableAdditionalAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


