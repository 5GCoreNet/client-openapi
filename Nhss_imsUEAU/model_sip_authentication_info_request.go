/*
Nhss_imsUEAU

Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUEAU

import (
	"encoding/json"
)

// SipAuthenticationInfoRequest Contains input data to the SIP authentication request message (e.g. SIP authentication scheme, requested number of authentication items, resynchronization information) 
type SipAuthenticationInfoRequest struct {
	CscfServerName string `json:"cscfServerName"`
	SipAuthenticationScheme SipAuthenticationScheme `json:"sipAuthenticationScheme"`
	// Indicates the number of requested SIP authentication items
	SipNumberAuthItems *int32 `json:"sipNumberAuthItems,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
}

// NewSipAuthenticationInfoRequest instantiates a new SipAuthenticationInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSipAuthenticationInfoRequest(cscfServerName string, sipAuthenticationScheme SipAuthenticationScheme) *SipAuthenticationInfoRequest {
	this := SipAuthenticationInfoRequest{}
	this.CscfServerName = cscfServerName
	this.SipAuthenticationScheme = sipAuthenticationScheme
	return &this
}

// NewSipAuthenticationInfoRequestWithDefaults instantiates a new SipAuthenticationInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSipAuthenticationInfoRequestWithDefaults() *SipAuthenticationInfoRequest {
	this := SipAuthenticationInfoRequest{}
	return &this
}

// GetCscfServerName returns the CscfServerName field value
func (o *SipAuthenticationInfoRequest) GetCscfServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CscfServerName
}

// GetCscfServerNameOk returns a tuple with the CscfServerName field value
// and a boolean to check if the value has been set.
func (o *SipAuthenticationInfoRequest) GetCscfServerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CscfServerName, true
}

// SetCscfServerName sets field value
func (o *SipAuthenticationInfoRequest) SetCscfServerName(v string) {
	o.CscfServerName = v
}

// GetSipAuthenticationScheme returns the SipAuthenticationScheme field value
func (o *SipAuthenticationInfoRequest) GetSipAuthenticationScheme() SipAuthenticationScheme {
	if o == nil {
		var ret SipAuthenticationScheme
		return ret
	}

	return o.SipAuthenticationScheme
}

// GetSipAuthenticationSchemeOk returns a tuple with the SipAuthenticationScheme field value
// and a boolean to check if the value has been set.
func (o *SipAuthenticationInfoRequest) GetSipAuthenticationSchemeOk() (*SipAuthenticationScheme, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SipAuthenticationScheme, true
}

// SetSipAuthenticationScheme sets field value
func (o *SipAuthenticationInfoRequest) SetSipAuthenticationScheme(v SipAuthenticationScheme) {
	o.SipAuthenticationScheme = v
}

// GetSipNumberAuthItems returns the SipNumberAuthItems field value if set, zero value otherwise.
func (o *SipAuthenticationInfoRequest) GetSipNumberAuthItems() int32 {
	if o == nil || isNil(o.SipNumberAuthItems) {
		var ret int32
		return ret
	}
	return *o.SipNumberAuthItems
}

// GetSipNumberAuthItemsOk returns a tuple with the SipNumberAuthItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipAuthenticationInfoRequest) GetSipNumberAuthItemsOk() (*int32, bool) {
	if o == nil || isNil(o.SipNumberAuthItems) {
    return nil, false
	}
	return o.SipNumberAuthItems, true
}

// HasSipNumberAuthItems returns a boolean if a field has been set.
func (o *SipAuthenticationInfoRequest) HasSipNumberAuthItems() bool {
	if o != nil && !isNil(o.SipNumberAuthItems) {
		return true
	}

	return false
}

// SetSipNumberAuthItems gets a reference to the given int32 and assigns it to the SipNumberAuthItems field.
func (o *SipAuthenticationInfoRequest) SetSipNumberAuthItems(v int32) {
	o.SipNumberAuthItems = &v
}

// GetResynchronizationInfo returns the ResynchronizationInfo field value if set, zero value otherwise.
func (o *SipAuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo {
	if o == nil || isNil(o.ResynchronizationInfo) {
		var ret ResynchronizationInfo
		return ret
	}
	return *o.ResynchronizationInfo
}

// GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipAuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool) {
	if o == nil || isNil(o.ResynchronizationInfo) {
    return nil, false
	}
	return o.ResynchronizationInfo, true
}

// HasResynchronizationInfo returns a boolean if a field has been set.
func (o *SipAuthenticationInfoRequest) HasResynchronizationInfo() bool {
	if o != nil && !isNil(o.ResynchronizationInfo) {
		return true
	}

	return false
}

// SetResynchronizationInfo gets a reference to the given ResynchronizationInfo and assigns it to the ResynchronizationInfo field.
func (o *SipAuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo) {
	o.ResynchronizationInfo = &v
}

func (o SipAuthenticationInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cscfServerName"] = o.CscfServerName
	}
	if true {
		toSerialize["sipAuthenticationScheme"] = o.SipAuthenticationScheme
	}
	if !isNil(o.SipNumberAuthItems) {
		toSerialize["sipNumberAuthItems"] = o.SipNumberAuthItems
	}
	if !isNil(o.ResynchronizationInfo) {
		toSerialize["resynchronizationInfo"] = o.ResynchronizationInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSipAuthenticationInfoRequest struct {
	value *SipAuthenticationInfoRequest
	isSet bool
}

func (v NullableSipAuthenticationInfoRequest) Get() *SipAuthenticationInfoRequest {
	return v.value
}

func (v *NullableSipAuthenticationInfoRequest) Set(val *SipAuthenticationInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSipAuthenticationInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSipAuthenticationInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipAuthenticationInfoRequest(val *SipAuthenticationInfoRequest) *NullableSipAuthenticationInfoRequest {
	return &NullableSipAuthenticationInfoRequest{value: val, isSet: true}
}

func (v NullableSipAuthenticationInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipAuthenticationInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


