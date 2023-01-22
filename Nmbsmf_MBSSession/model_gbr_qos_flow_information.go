/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
)

// GbrQosFlowInformation GBR MBS QoS flow information
type GbrQosFlowInformation struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxFbrDl string `json:"maxFbrDl"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GuaFbrDl string `json:"guaFbrDl"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	MaxPacketLossRateDl *int32 `json:"maxPacketLossRateDl,omitempty"`
}

// NewGbrQosFlowInformation instantiates a new GbrQosFlowInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGbrQosFlowInformation(maxFbrDl string, guaFbrDl string) *GbrQosFlowInformation {
	this := GbrQosFlowInformation{}
	this.MaxFbrDl = maxFbrDl
	this.GuaFbrDl = guaFbrDl
	return &this
}

// NewGbrQosFlowInformationWithDefaults instantiates a new GbrQosFlowInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGbrQosFlowInformationWithDefaults() *GbrQosFlowInformation {
	this := GbrQosFlowInformation{}
	return &this
}

// GetMaxFbrDl returns the MaxFbrDl field value
func (o *GbrQosFlowInformation) GetMaxFbrDl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFbrDl
}

// GetMaxFbrDlOk returns a tuple with the MaxFbrDl field value
// and a boolean to check if the value has been set.
func (o *GbrQosFlowInformation) GetMaxFbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxFbrDl, true
}

// SetMaxFbrDl sets field value
func (o *GbrQosFlowInformation) SetMaxFbrDl(v string) {
	o.MaxFbrDl = v
}

// GetGuaFbrDl returns the GuaFbrDl field value
func (o *GbrQosFlowInformation) GetGuaFbrDl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuaFbrDl
}

// GetGuaFbrDlOk returns a tuple with the GuaFbrDl field value
// and a boolean to check if the value has been set.
func (o *GbrQosFlowInformation) GetGuaFbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GuaFbrDl, true
}

// SetGuaFbrDl sets field value
func (o *GbrQosFlowInformation) SetGuaFbrDl(v string) {
	o.GuaFbrDl = v
}

// GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field value if set, zero value otherwise.
func (o *GbrQosFlowInformation) GetMaxPacketLossRateDl() int32 {
	if o == nil || isNil(o.MaxPacketLossRateDl) {
		var ret int32
		return ret
	}
	return *o.MaxPacketLossRateDl
}

// GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GbrQosFlowInformation) GetMaxPacketLossRateDlOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPacketLossRateDl) {
    return nil, false
	}
	return o.MaxPacketLossRateDl, true
}

// HasMaxPacketLossRateDl returns a boolean if a field has been set.
func (o *GbrQosFlowInformation) HasMaxPacketLossRateDl() bool {
	if o != nil && !isNil(o.MaxPacketLossRateDl) {
		return true
	}

	return false
}

// SetMaxPacketLossRateDl gets a reference to the given int32 and assigns it to the MaxPacketLossRateDl field.
func (o *GbrQosFlowInformation) SetMaxPacketLossRateDl(v int32) {
	o.MaxPacketLossRateDl = &v
}

func (o GbrQosFlowInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maxFbrDl"] = o.MaxFbrDl
	}
	if true {
		toSerialize["guaFbrDl"] = o.GuaFbrDl
	}
	if !isNil(o.MaxPacketLossRateDl) {
		toSerialize["maxPacketLossRateDl"] = o.MaxPacketLossRateDl
	}
	return json.Marshal(toSerialize)
}

type NullableGbrQosFlowInformation struct {
	value *GbrQosFlowInformation
	isSet bool
}

func (v NullableGbrQosFlowInformation) Get() *GbrQosFlowInformation {
	return v.value
}

func (v *NullableGbrQosFlowInformation) Set(val *GbrQosFlowInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableGbrQosFlowInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableGbrQosFlowInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbrQosFlowInformation(val *GbrQosFlowInformation) *NullableGbrQosFlowInformation {
	return &NullableGbrQosFlowInformation{value: val, isSet: true}
}

func (v NullableGbrQosFlowInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbrQosFlowInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

