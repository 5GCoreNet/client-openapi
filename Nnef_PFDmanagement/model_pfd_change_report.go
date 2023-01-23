/*
Nnef_PFDmanagement Service API

Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_PFDmanagement

import (
	"encoding/json"
)

// PfdChangeReport Represents an error report on PFD change.
type PfdChangeReport struct {
	PfdError ProblemDetails `json:"pfdError"`
	ApplicationId []string `json:"applicationId"`
}

// NewPfdChangeReport instantiates a new PfdChangeReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfdChangeReport(pfdError ProblemDetails, applicationId []string) *PfdChangeReport {
	this := PfdChangeReport{}
	this.PfdError = pfdError
	this.ApplicationId = applicationId
	return &this
}

// NewPfdChangeReportWithDefaults instantiates a new PfdChangeReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdChangeReportWithDefaults() *PfdChangeReport {
	this := PfdChangeReport{}
	return &this
}

// GetPfdError returns the PfdError field value
func (o *PfdChangeReport) GetPfdError() ProblemDetails {
	if o == nil {
		var ret ProblemDetails
		return ret
	}

	return o.PfdError
}

// GetPfdErrorOk returns a tuple with the PfdError field value
// and a boolean to check if the value has been set.
func (o *PfdChangeReport) GetPfdErrorOk() (*ProblemDetails, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PfdError, true
}

// SetPfdError sets field value
func (o *PfdChangeReport) SetPfdError(v ProblemDetails) {
	o.PfdError = v
}

// GetApplicationId returns the ApplicationId field value
func (o *PfdChangeReport) GetApplicationId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *PfdChangeReport) GetApplicationIdOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *PfdChangeReport) SetApplicationId(v []string) {
	o.ApplicationId = v
}

func (o PfdChangeReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pfdError"] = o.PfdError
	}
	if true {
		toSerialize["applicationId"] = o.ApplicationId
	}
	return json.Marshal(toSerialize)
}

type NullablePfdChangeReport struct {
	value *PfdChangeReport
	isSet bool
}

func (v NullablePfdChangeReport) Get() *PfdChangeReport {
	return v.value
}

func (v *NullablePfdChangeReport) Set(val *PfdChangeReport) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdChangeReport) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdChangeReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdChangeReport(val *PfdChangeReport) *NullablePfdChangeReport {
	return &NullablePfdChangeReport{value: val, isSet: true}
}

func (v NullablePfdChangeReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdChangeReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


