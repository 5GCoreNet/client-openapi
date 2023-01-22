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

// PsaInformation PSA Information
type PsaInformation struct {
	PsaInd *PsaIndication `json:"psaInd,omitempty"`
	DnaiList []string `json:"dnaiList,omitempty"`
	UeIpv6Prefix *Ipv6Prefix `json:"ueIpv6Prefix,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PsaUpfId *string `json:"psaUpfId,omitempty"`
}

// NewPsaInformation instantiates a new PsaInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsaInformation() *PsaInformation {
	this := PsaInformation{}
	return &this
}

// NewPsaInformationWithDefaults instantiates a new PsaInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsaInformationWithDefaults() *PsaInformation {
	this := PsaInformation{}
	return &this
}

// GetPsaInd returns the PsaInd field value if set, zero value otherwise.
func (o *PsaInformation) GetPsaInd() PsaIndication {
	if o == nil || isNil(o.PsaInd) {
		var ret PsaIndication
		return ret
	}
	return *o.PsaInd
}

// GetPsaIndOk returns a tuple with the PsaInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsaInformation) GetPsaIndOk() (*PsaIndication, bool) {
	if o == nil || isNil(o.PsaInd) {
    return nil, false
	}
	return o.PsaInd, true
}

// HasPsaInd returns a boolean if a field has been set.
func (o *PsaInformation) HasPsaInd() bool {
	if o != nil && !isNil(o.PsaInd) {
		return true
	}

	return false
}

// SetPsaInd gets a reference to the given PsaIndication and assigns it to the PsaInd field.
func (o *PsaInformation) SetPsaInd(v PsaIndication) {
	o.PsaInd = &v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *PsaInformation) GetDnaiList() []string {
	if o == nil || isNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsaInformation) GetDnaiListOk() ([]string, bool) {
	if o == nil || isNil(o.DnaiList) {
    return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *PsaInformation) HasDnaiList() bool {
	if o != nil && !isNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *PsaInformation) SetDnaiList(v []string) {
	o.DnaiList = v
}

// GetUeIpv6Prefix returns the UeIpv6Prefix field value if set, zero value otherwise.
func (o *PsaInformation) GetUeIpv6Prefix() Ipv6Prefix {
	if o == nil || isNil(o.UeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.UeIpv6Prefix
}

// GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsaInformation) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.UeIpv6Prefix) {
    return nil, false
	}
	return o.UeIpv6Prefix, true
}

// HasUeIpv6Prefix returns a boolean if a field has been set.
func (o *PsaInformation) HasUeIpv6Prefix() bool {
	if o != nil && !isNil(o.UeIpv6Prefix) {
		return true
	}

	return false
}

// SetUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the UeIpv6Prefix field.
func (o *PsaInformation) SetUeIpv6Prefix(v Ipv6Prefix) {
	o.UeIpv6Prefix = &v
}

// GetPsaUpfId returns the PsaUpfId field value if set, zero value otherwise.
func (o *PsaInformation) GetPsaUpfId() string {
	if o == nil || isNil(o.PsaUpfId) {
		var ret string
		return ret
	}
	return *o.PsaUpfId
}

// GetPsaUpfIdOk returns a tuple with the PsaUpfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsaInformation) GetPsaUpfIdOk() (*string, bool) {
	if o == nil || isNil(o.PsaUpfId) {
    return nil, false
	}
	return o.PsaUpfId, true
}

// HasPsaUpfId returns a boolean if a field has been set.
func (o *PsaInformation) HasPsaUpfId() bool {
	if o != nil && !isNil(o.PsaUpfId) {
		return true
	}

	return false
}

// SetPsaUpfId gets a reference to the given string and assigns it to the PsaUpfId field.
func (o *PsaInformation) SetPsaUpfId(v string) {
	o.PsaUpfId = &v
}

func (o PsaInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PsaInd) {
		toSerialize["psaInd"] = o.PsaInd
	}
	if !isNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if !isNil(o.UeIpv6Prefix) {
		toSerialize["ueIpv6Prefix"] = o.UeIpv6Prefix
	}
	if !isNil(o.PsaUpfId) {
		toSerialize["psaUpfId"] = o.PsaUpfId
	}
	return json.Marshal(toSerialize)
}

type NullablePsaInformation struct {
	value *PsaInformation
	isSet bool
}

func (v NullablePsaInformation) Get() *PsaInformation {
	return v.value
}

func (v *NullablePsaInformation) Set(val *PsaInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePsaInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePsaInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsaInformation(val *PsaInformation) *NullablePsaInformation {
	return &NullablePsaInformation{value: val, isSet: true}
}

func (v NullablePsaInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsaInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

