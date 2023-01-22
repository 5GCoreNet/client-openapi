/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// N2InformationTransferError Data within a failure response for a non-UE related N2 Information Transfer
type N2InformationTransferError struct {
	Error ProblemDetails `json:"error"`
	PwsErrorInfo *PWSErrorData `json:"pwsErrorInfo,omitempty"`
}

// NewN2InformationTransferError instantiates a new N2InformationTransferError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2InformationTransferError(error_ ProblemDetails) *N2InformationTransferError {
	this := N2InformationTransferError{}
	this.Error = error_
	return &this
}

// NewN2InformationTransferErrorWithDefaults instantiates a new N2InformationTransferError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2InformationTransferErrorWithDefaults() *N2InformationTransferError {
	this := N2InformationTransferError{}
	return &this
}

// GetError returns the Error field value
func (o *N2InformationTransferError) GetError() ProblemDetails {
	if o == nil {
		var ret ProblemDetails
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *N2InformationTransferError) GetErrorOk() (*ProblemDetails, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *N2InformationTransferError) SetError(v ProblemDetails) {
	o.Error = v
}

// GetPwsErrorInfo returns the PwsErrorInfo field value if set, zero value otherwise.
func (o *N2InformationTransferError) GetPwsErrorInfo() PWSErrorData {
	if o == nil || isNil(o.PwsErrorInfo) {
		var ret PWSErrorData
		return ret
	}
	return *o.PwsErrorInfo
}

// GetPwsErrorInfoOk returns a tuple with the PwsErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InformationTransferError) GetPwsErrorInfoOk() (*PWSErrorData, bool) {
	if o == nil || isNil(o.PwsErrorInfo) {
    return nil, false
	}
	return o.PwsErrorInfo, true
}

// HasPwsErrorInfo returns a boolean if a field has been set.
func (o *N2InformationTransferError) HasPwsErrorInfo() bool {
	if o != nil && !isNil(o.PwsErrorInfo) {
		return true
	}

	return false
}

// SetPwsErrorInfo gets a reference to the given PWSErrorData and assigns it to the PwsErrorInfo field.
func (o *N2InformationTransferError) SetPwsErrorInfo(v PWSErrorData) {
	o.PwsErrorInfo = &v
}

func (o N2InformationTransferError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.PwsErrorInfo) {
		toSerialize["pwsErrorInfo"] = o.PwsErrorInfo
	}
	return json.Marshal(toSerialize)
}

type NullableN2InformationTransferError struct {
	value *N2InformationTransferError
	isSet bool
}

func (v NullableN2InformationTransferError) Get() *N2InformationTransferError {
	return v.value
}

func (v *NullableN2InformationTransferError) Set(val *N2InformationTransferError) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationTransferError) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationTransferError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationTransferError(val *N2InformationTransferError) *NullableN2InformationTransferError {
	return &NullableN2InformationTransferError{value: val, isSet: true}
}

func (v NullableN2InformationTransferError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationTransferError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


