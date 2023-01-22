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

// AmfInfo struct for AmfInfo
type AmfInfo struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId string `json:"amfInstanceId"`
	Guami Guami `json:"guami"`
	AccessType *AccessType `json:"accessType,omitempty"`
}

// NewAmfInfo instantiates a new AmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfInfo(amfInstanceId string, guami Guami) *AmfInfo {
	this := AmfInfo{}
	this.AmfInstanceId = amfInstanceId
	this.Guami = guami
	return &this
}

// NewAmfInfoWithDefaults instantiates a new AmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfInfoWithDefaults() *AmfInfo {
	this := AmfInfo{}
	return &this
}

// GetAmfInstanceId returns the AmfInstanceId field value
func (o *AmfInfo) GetAmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfInstanceId
}

// GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmfInstanceId, true
}

// SetAmfInstanceId sets field value
func (o *AmfInfo) SetAmfInstanceId(v string) {
	o.AmfInstanceId = v
}

// GetGuami returns the Guami field value
func (o *AmfInfo) GetGuami() Guami {
	if o == nil {
		var ret Guami
		return ret
	}

	return o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetGuamiOk() (*Guami, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Guami, true
}

// SetGuami sets field value
func (o *AmfInfo) SetGuami(v Guami) {
	o.Guami = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AmfInfo) GetAccessType() AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
    return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AmfInfo) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *AmfInfo) SetAccessType(v AccessType) {
	o.AccessType = &v
}

func (o AmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amfInstanceId"] = o.AmfInstanceId
	}
	if true {
		toSerialize["guami"] = o.Guami
	}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableAmfInfo struct {
	value *AmfInfo
	isSet bool
}

func (v NullableAmfInfo) Get() *AmfInfo {
	return v.value
}

func (v *NullableAmfInfo) Set(val *AmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfInfo(val *AmfInfo) *NullableAmfInfo {
	return &NullableAmfInfo{value: val, isSet: true}
}

func (v NullableAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

