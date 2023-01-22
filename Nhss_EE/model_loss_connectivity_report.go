/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_EE

import (
	"encoding/json"
)

// LossConnectivityReport struct for LossConnectivityReport
type LossConnectivityReport struct {
	LossOfConnectReason LossOfConnectivityReason `json:"lossOfConnectReason"`
}

// NewLossConnectivityReport instantiates a new LossConnectivityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLossConnectivityReport(lossOfConnectReason LossOfConnectivityReason) *LossConnectivityReport {
	this := LossConnectivityReport{}
	this.LossOfConnectReason = lossOfConnectReason
	return &this
}

// NewLossConnectivityReportWithDefaults instantiates a new LossConnectivityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLossConnectivityReportWithDefaults() *LossConnectivityReport {
	this := LossConnectivityReport{}
	return &this
}

// GetLossOfConnectReason returns the LossOfConnectReason field value
func (o *LossConnectivityReport) GetLossOfConnectReason() LossOfConnectivityReason {
	if o == nil {
		var ret LossOfConnectivityReason
		return ret
	}

	return o.LossOfConnectReason
}

// GetLossOfConnectReasonOk returns a tuple with the LossOfConnectReason field value
// and a boolean to check if the value has been set.
func (o *LossConnectivityReport) GetLossOfConnectReasonOk() (*LossOfConnectivityReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LossOfConnectReason, true
}

// SetLossOfConnectReason sets field value
func (o *LossConnectivityReport) SetLossOfConnectReason(v LossOfConnectivityReason) {
	o.LossOfConnectReason = v
}

func (o LossConnectivityReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lossOfConnectReason"] = o.LossOfConnectReason
	}
	return json.Marshal(toSerialize)
}

type NullableLossConnectivityReport struct {
	value *LossConnectivityReport
	isSet bool
}

func (v NullableLossConnectivityReport) Get() *LossConnectivityReport {
	return v.value
}

func (v *NullableLossConnectivityReport) Set(val *LossConnectivityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLossConnectivityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLossConnectivityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossConnectivityReport(val *LossConnectivityReport) *NullableLossConnectivityReport {
	return &NullableLossConnectivityReport{value: val, isSet: true}
}

func (v NullableLossConnectivityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossConnectivityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


