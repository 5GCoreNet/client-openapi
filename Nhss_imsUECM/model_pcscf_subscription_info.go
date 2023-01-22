/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUECM

import (
	"encoding/json"
)

// PcscfSubscriptionInfo Subscription information of the P-CSCF for the SIP Registration State event
type PcscfSubscriptionInfo struct {
	CallIdSipHeader string `json:"callIdSipHeader"`
	FromSipHeader string `json:"fromSipHeader"`
	ToSipHeader string `json:"toSipHeader"`
	Contact string `json:"contact"`
}

// NewPcscfSubscriptionInfo instantiates a new PcscfSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcscfSubscriptionInfo(callIdSipHeader string, fromSipHeader string, toSipHeader string, contact string) *PcscfSubscriptionInfo {
	this := PcscfSubscriptionInfo{}
	this.CallIdSipHeader = callIdSipHeader
	this.FromSipHeader = fromSipHeader
	this.ToSipHeader = toSipHeader
	this.Contact = contact
	return &this
}

// NewPcscfSubscriptionInfoWithDefaults instantiates a new PcscfSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcscfSubscriptionInfoWithDefaults() *PcscfSubscriptionInfo {
	this := PcscfSubscriptionInfo{}
	return &this
}

// GetCallIdSipHeader returns the CallIdSipHeader field value
func (o *PcscfSubscriptionInfo) GetCallIdSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallIdSipHeader
}

// GetCallIdSipHeaderOk returns a tuple with the CallIdSipHeader field value
// and a boolean to check if the value has been set.
func (o *PcscfSubscriptionInfo) GetCallIdSipHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallIdSipHeader, true
}

// SetCallIdSipHeader sets field value
func (o *PcscfSubscriptionInfo) SetCallIdSipHeader(v string) {
	o.CallIdSipHeader = v
}

// GetFromSipHeader returns the FromSipHeader field value
func (o *PcscfSubscriptionInfo) GetFromSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromSipHeader
}

// GetFromSipHeaderOk returns a tuple with the FromSipHeader field value
// and a boolean to check if the value has been set.
func (o *PcscfSubscriptionInfo) GetFromSipHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FromSipHeader, true
}

// SetFromSipHeader sets field value
func (o *PcscfSubscriptionInfo) SetFromSipHeader(v string) {
	o.FromSipHeader = v
}

// GetToSipHeader returns the ToSipHeader field value
func (o *PcscfSubscriptionInfo) GetToSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToSipHeader
}

// GetToSipHeaderOk returns a tuple with the ToSipHeader field value
// and a boolean to check if the value has been set.
func (o *PcscfSubscriptionInfo) GetToSipHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ToSipHeader, true
}

// SetToSipHeader sets field value
func (o *PcscfSubscriptionInfo) SetToSipHeader(v string) {
	o.ToSipHeader = v
}

// GetContact returns the Contact field value
func (o *PcscfSubscriptionInfo) GetContact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *PcscfSubscriptionInfo) GetContactOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *PcscfSubscriptionInfo) SetContact(v string) {
	o.Contact = v
}

func (o PcscfSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["callIdSipHeader"] = o.CallIdSipHeader
	}
	if true {
		toSerialize["fromSipHeader"] = o.FromSipHeader
	}
	if true {
		toSerialize["toSipHeader"] = o.ToSipHeader
	}
	if true {
		toSerialize["contact"] = o.Contact
	}
	return json.Marshal(toSerialize)
}

type NullablePcscfSubscriptionInfo struct {
	value *PcscfSubscriptionInfo
	isSet bool
}

func (v NullablePcscfSubscriptionInfo) Get() *PcscfSubscriptionInfo {
	return v.value
}

func (v *NullablePcscfSubscriptionInfo) Set(val *PcscfSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcscfSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcscfSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcscfSubscriptionInfo(val *PcscfSubscriptionInfo) *NullablePcscfSubscriptionInfo {
	return &NullablePcscfSubscriptionInfo{value: val, isSet: true}
}

func (v NullablePcscfSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcscfSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


