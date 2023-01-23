/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"time"
)

// MbsKeyInfo MBS Security Key Data Structure
type MbsKeyInfo struct {
	// string with format 'bytes' as defined in OpenAPI
	KeyDomainId string `json:"keyDomainId"`
	// string with format 'bytes' as defined in OpenAPI
	MskId string `json:"mskId"`
	// string with format 'bytes' as defined in OpenAPI
	Msk *string `json:"msk,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MskLifetime *time.Time `json:"mskLifetime,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	MtkId *string `json:"mtkId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	Mtk *string `json:"mtk,omitempty"`
}

// NewMbsKeyInfo instantiates a new MbsKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsKeyInfo(keyDomainId string, mskId string) *MbsKeyInfo {
	this := MbsKeyInfo{}
	this.KeyDomainId = keyDomainId
	this.MskId = mskId
	return &this
}

// NewMbsKeyInfoWithDefaults instantiates a new MbsKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsKeyInfoWithDefaults() *MbsKeyInfo {
	this := MbsKeyInfo{}
	return &this
}

// GetKeyDomainId returns the KeyDomainId field value
func (o *MbsKeyInfo) GetKeyDomainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyDomainId
}

// GetKeyDomainIdOk returns a tuple with the KeyDomainId field value
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetKeyDomainIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KeyDomainId, true
}

// SetKeyDomainId sets field value
func (o *MbsKeyInfo) SetKeyDomainId(v string) {
	o.KeyDomainId = v
}

// GetMskId returns the MskId field value
func (o *MbsKeyInfo) GetMskId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MskId
}

// GetMskIdOk returns a tuple with the MskId field value
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetMskIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MskId, true
}

// SetMskId sets field value
func (o *MbsKeyInfo) SetMskId(v string) {
	o.MskId = v
}

// GetMsk returns the Msk field value if set, zero value otherwise.
func (o *MbsKeyInfo) GetMsk() string {
	if o == nil || isNil(o.Msk) {
		var ret string
		return ret
	}
	return *o.Msk
}

// GetMskOk returns a tuple with the Msk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetMskOk() (*string, bool) {
	if o == nil || isNil(o.Msk) {
    return nil, false
	}
	return o.Msk, true
}

// HasMsk returns a boolean if a field has been set.
func (o *MbsKeyInfo) HasMsk() bool {
	if o != nil && !isNil(o.Msk) {
		return true
	}

	return false
}

// SetMsk gets a reference to the given string and assigns it to the Msk field.
func (o *MbsKeyInfo) SetMsk(v string) {
	o.Msk = &v
}

// GetMskLifetime returns the MskLifetime field value if set, zero value otherwise.
func (o *MbsKeyInfo) GetMskLifetime() time.Time {
	if o == nil || isNil(o.MskLifetime) {
		var ret time.Time
		return ret
	}
	return *o.MskLifetime
}

// GetMskLifetimeOk returns a tuple with the MskLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetMskLifetimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.MskLifetime) {
    return nil, false
	}
	return o.MskLifetime, true
}

// HasMskLifetime returns a boolean if a field has been set.
func (o *MbsKeyInfo) HasMskLifetime() bool {
	if o != nil && !isNil(o.MskLifetime) {
		return true
	}

	return false
}

// SetMskLifetime gets a reference to the given time.Time and assigns it to the MskLifetime field.
func (o *MbsKeyInfo) SetMskLifetime(v time.Time) {
	o.MskLifetime = &v
}

// GetMtkId returns the MtkId field value if set, zero value otherwise.
func (o *MbsKeyInfo) GetMtkId() string {
	if o == nil || isNil(o.MtkId) {
		var ret string
		return ret
	}
	return *o.MtkId
}

// GetMtkIdOk returns a tuple with the MtkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetMtkIdOk() (*string, bool) {
	if o == nil || isNil(o.MtkId) {
    return nil, false
	}
	return o.MtkId, true
}

// HasMtkId returns a boolean if a field has been set.
func (o *MbsKeyInfo) HasMtkId() bool {
	if o != nil && !isNil(o.MtkId) {
		return true
	}

	return false
}

// SetMtkId gets a reference to the given string and assigns it to the MtkId field.
func (o *MbsKeyInfo) SetMtkId(v string) {
	o.MtkId = &v
}

// GetMtk returns the Mtk field value if set, zero value otherwise.
func (o *MbsKeyInfo) GetMtk() string {
	if o == nil || isNil(o.Mtk) {
		var ret string
		return ret
	}
	return *o.Mtk
}

// GetMtkOk returns a tuple with the Mtk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsKeyInfo) GetMtkOk() (*string, bool) {
	if o == nil || isNil(o.Mtk) {
    return nil, false
	}
	return o.Mtk, true
}

// HasMtk returns a boolean if a field has been set.
func (o *MbsKeyInfo) HasMtk() bool {
	if o != nil && !isNil(o.Mtk) {
		return true
	}

	return false
}

// SetMtk gets a reference to the given string and assigns it to the Mtk field.
func (o *MbsKeyInfo) SetMtk(v string) {
	o.Mtk = &v
}

func (o MbsKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keyDomainId"] = o.KeyDomainId
	}
	if true {
		toSerialize["mskId"] = o.MskId
	}
	if !isNil(o.Msk) {
		toSerialize["msk"] = o.Msk
	}
	if !isNil(o.MskLifetime) {
		toSerialize["mskLifetime"] = o.MskLifetime
	}
	if !isNil(o.MtkId) {
		toSerialize["mtkId"] = o.MtkId
	}
	if !isNil(o.Mtk) {
		toSerialize["mtk"] = o.Mtk
	}
	return json.Marshal(toSerialize)
}

type NullableMbsKeyInfo struct {
	value *MbsKeyInfo
	isSet bool
}

func (v NullableMbsKeyInfo) Get() *MbsKeyInfo {
	return v.value
}

func (v *NullableMbsKeyInfo) Set(val *MbsKeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsKeyInfo(val *MbsKeyInfo) *NullableMbsKeyInfo {
	return &NullableMbsKeyInfo{value: val, isSet: true}
}

func (v NullableMbsKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


