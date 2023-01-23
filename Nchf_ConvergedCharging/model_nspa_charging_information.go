/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// NSPAChargingInformation struct for NSPAChargingInformation
type NSPAChargingInformation struct {
	SingleNSSAI Snssai `json:"singleNSSAI"`
}

// NewNSPAChargingInformation instantiates a new NSPAChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSPAChargingInformation(singleNSSAI Snssai) *NSPAChargingInformation {
	this := NSPAChargingInformation{}
	this.SingleNSSAI = singleNSSAI
	return &this
}

// NewNSPAChargingInformationWithDefaults instantiates a new NSPAChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSPAChargingInformationWithDefaults() *NSPAChargingInformation {
	this := NSPAChargingInformation{}
	return &this
}

// GetSingleNSSAI returns the SingleNSSAI field value
func (o *NSPAChargingInformation) GetSingleNSSAI() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SingleNSSAI
}

// GetSingleNSSAIOk returns a tuple with the SingleNSSAI field value
// and a boolean to check if the value has been set.
func (o *NSPAChargingInformation) GetSingleNSSAIOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SingleNSSAI, true
}

// SetSingleNSSAI sets field value
func (o *NSPAChargingInformation) SetSingleNSSAI(v Snssai) {
	o.SingleNSSAI = v
}

func (o NSPAChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["singleNSSAI"] = o.SingleNSSAI
	}
	return json.Marshal(toSerialize)
}

type NullableNSPAChargingInformation struct {
	value *NSPAChargingInformation
	isSet bool
}

func (v NullableNSPAChargingInformation) Get() *NSPAChargingInformation {
	return v.value
}

func (v *NullableNSPAChargingInformation) Set(val *NSPAChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNSPAChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNSPAChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSPAChargingInformation(val *NSPAChargingInformation) *NullableNSPAChargingInformation {
	return &NullableNSPAChargingInformation{value: val, isSet: true}
}

func (v NullableNSPAChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSPAChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


