/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// AmfEventArea Represents an area to be monitored by an AMF event
type AmfEventArea struct {
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
	LadnInfo *LadnInfo `json:"ladnInfo,omitempty"`
	SNssai *Snssai `json:"sNssai,omitempty"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId *string `json:"nsiId,omitempty"`
}

// NewAmfEventArea instantiates a new AmfEventArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventArea() *AmfEventArea {
	this := AmfEventArea{}
	return &this
}

// NewAmfEventAreaWithDefaults instantiates a new AmfEventArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventAreaWithDefaults() *AmfEventArea {
	this := AmfEventArea{}
	return &this
}

// GetPresenceInfo returns the PresenceInfo field value if set, zero value otherwise.
func (o *AmfEventArea) GetPresenceInfo() PresenceInfo {
	if o == nil || isNil(o.PresenceInfo) {
		var ret PresenceInfo
		return ret
	}
	return *o.PresenceInfo
}

// GetPresenceInfoOk returns a tuple with the PresenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventArea) GetPresenceInfoOk() (*PresenceInfo, bool) {
	if o == nil || isNil(o.PresenceInfo) {
    return nil, false
	}
	return o.PresenceInfo, true
}

// HasPresenceInfo returns a boolean if a field has been set.
func (o *AmfEventArea) HasPresenceInfo() bool {
	if o != nil && !isNil(o.PresenceInfo) {
		return true
	}

	return false
}

// SetPresenceInfo gets a reference to the given PresenceInfo and assigns it to the PresenceInfo field.
func (o *AmfEventArea) SetPresenceInfo(v PresenceInfo) {
	o.PresenceInfo = &v
}

// GetLadnInfo returns the LadnInfo field value if set, zero value otherwise.
func (o *AmfEventArea) GetLadnInfo() LadnInfo {
	if o == nil || isNil(o.LadnInfo) {
		var ret LadnInfo
		return ret
	}
	return *o.LadnInfo
}

// GetLadnInfoOk returns a tuple with the LadnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventArea) GetLadnInfoOk() (*LadnInfo, bool) {
	if o == nil || isNil(o.LadnInfo) {
    return nil, false
	}
	return o.LadnInfo, true
}

// HasLadnInfo returns a boolean if a field has been set.
func (o *AmfEventArea) HasLadnInfo() bool {
	if o != nil && !isNil(o.LadnInfo) {
		return true
	}

	return false
}

// SetLadnInfo gets a reference to the given LadnInfo and assigns it to the LadnInfo field.
func (o *AmfEventArea) SetLadnInfo(v LadnInfo) {
	o.LadnInfo = &v
}

// GetSNssai returns the SNssai field value if set, zero value otherwise.
func (o *AmfEventArea) GetSNssai() Snssai {
	if o == nil || isNil(o.SNssai) {
		var ret Snssai
		return ret
	}
	return *o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventArea) GetSNssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.SNssai) {
    return nil, false
	}
	return o.SNssai, true
}

// HasSNssai returns a boolean if a field has been set.
func (o *AmfEventArea) HasSNssai() bool {
	if o != nil && !isNil(o.SNssai) {
		return true
	}

	return false
}

// SetSNssai gets a reference to the given Snssai and assigns it to the SNssai field.
func (o *AmfEventArea) SetSNssai(v Snssai) {
	o.SNssai = &v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *AmfEventArea) GetNsiId() string {
	if o == nil || isNil(o.NsiId) {
		var ret string
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventArea) GetNsiIdOk() (*string, bool) {
	if o == nil || isNil(o.NsiId) {
    return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *AmfEventArea) HasNsiId() bool {
	if o != nil && !isNil(o.NsiId) {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given string and assigns it to the NsiId field.
func (o *AmfEventArea) SetNsiId(v string) {
	o.NsiId = &v
}

func (o AmfEventArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PresenceInfo) {
		toSerialize["presenceInfo"] = o.PresenceInfo
	}
	if !isNil(o.LadnInfo) {
		toSerialize["ladnInfo"] = o.LadnInfo
	}
	if !isNil(o.SNssai) {
		toSerialize["sNssai"] = o.SNssai
	}
	if !isNil(o.NsiId) {
		toSerialize["nsiId"] = o.NsiId
	}
	return json.Marshal(toSerialize)
}

type NullableAmfEventArea struct {
	value *AmfEventArea
	isSet bool
}

func (v NullableAmfEventArea) Get() *AmfEventArea {
	return v.value
}

func (v *NullableAmfEventArea) Set(val *AmfEventArea) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventArea) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventArea(val *AmfEventArea) *NullableAmfEventArea {
	return &NullableAmfEventArea{value: val, isSet: true}
}

func (v NullableAmfEventArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


