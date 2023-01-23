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

// PcfAddressingInfo Contains PCF address information.
type PcfAddressingInfo struct {
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
	// contains the binding indications of the PCF.
	BindingInfo *string `json:"bindingInfo,omitempty"`
}

// NewPcfAddressingInfo instantiates a new PcfAddressingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfAddressingInfo() *PcfAddressingInfo {
	this := PcfAddressingInfo{}
	return &this
}

// NewPcfAddressingInfoWithDefaults instantiates a new PcfAddressingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfAddressingInfoWithDefaults() *PcfAddressingInfo {
	this := PcfAddressingInfo{}
	return &this
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *PcfAddressingInfo) GetPcfFqdn() string {
	if o == nil || isNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfAddressingInfo) GetPcfFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PcfFqdn) {
    return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *PcfAddressingInfo) HasPcfFqdn() bool {
	if o != nil && !isNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *PcfAddressingInfo) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *PcfAddressingInfo) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || isNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfAddressingInfo) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || isNil(o.PcfIpEndPoints) {
    return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *PcfAddressingInfo) HasPcfIpEndPoints() bool {
	if o != nil && !isNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *PcfAddressingInfo) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

// GetBindingInfo returns the BindingInfo field value if set, zero value otherwise.
func (o *PcfAddressingInfo) GetBindingInfo() string {
	if o == nil || isNil(o.BindingInfo) {
		var ret string
		return ret
	}
	return *o.BindingInfo
}

// GetBindingInfoOk returns a tuple with the BindingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfAddressingInfo) GetBindingInfoOk() (*string, bool) {
	if o == nil || isNil(o.BindingInfo) {
    return nil, false
	}
	return o.BindingInfo, true
}

// HasBindingInfo returns a boolean if a field has been set.
func (o *PcfAddressingInfo) HasBindingInfo() bool {
	if o != nil && !isNil(o.BindingInfo) {
		return true
	}

	return false
}

// SetBindingInfo gets a reference to the given string and assigns it to the BindingInfo field.
func (o *PcfAddressingInfo) SetBindingInfo(v string) {
	o.BindingInfo = &v
}

func (o PcfAddressingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PcfFqdn) {
		toSerialize["pcfFqdn"] = o.PcfFqdn
	}
	if !isNil(o.PcfIpEndPoints) {
		toSerialize["pcfIpEndPoints"] = o.PcfIpEndPoints
	}
	if !isNil(o.BindingInfo) {
		toSerialize["bindingInfo"] = o.BindingInfo
	}
	return json.Marshal(toSerialize)
}

type NullablePcfAddressingInfo struct {
	value *PcfAddressingInfo
	isSet bool
}

func (v NullablePcfAddressingInfo) Get() *PcfAddressingInfo {
	return v.value
}

func (v *NullablePcfAddressingInfo) Set(val *PcfAddressingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfAddressingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfAddressingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfAddressingInfo(val *PcfAddressingInfo) *NullablePcfAddressingInfo {
	return &NullablePcfAddressingInfo{value: val, isSet: true}
}

func (v NullablePcfAddressingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfAddressingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


