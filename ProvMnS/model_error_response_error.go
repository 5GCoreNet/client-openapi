/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// ErrorResponseError struct for ErrorResponseError
type ErrorResponseError struct {
	ErrorInfo *string `json:"errorInfo,omitempty"`
}

// NewErrorResponseError instantiates a new ErrorResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseError() *ErrorResponseError {
	this := ErrorResponseError{}
	return &this
}

// NewErrorResponseErrorWithDefaults instantiates a new ErrorResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorWithDefaults() *ErrorResponseError {
	this := ErrorResponseError{}
	return &this
}

// GetErrorInfo returns the ErrorInfo field value if set, zero value otherwise.
func (o *ErrorResponseError) GetErrorInfo() string {
	if o == nil || isNil(o.ErrorInfo) {
		var ret string
		return ret
	}
	return *o.ErrorInfo
}

// GetErrorInfoOk returns a tuple with the ErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetErrorInfoOk() (*string, bool) {
	if o == nil || isNil(o.ErrorInfo) {
    return nil, false
	}
	return o.ErrorInfo, true
}

// HasErrorInfo returns a boolean if a field has been set.
func (o *ErrorResponseError) HasErrorInfo() bool {
	if o != nil && !isNil(o.ErrorInfo) {
		return true
	}

	return false
}

// SetErrorInfo gets a reference to the given string and assigns it to the ErrorInfo field.
func (o *ErrorResponseError) SetErrorInfo(v string) {
	o.ErrorInfo = &v
}

func (o ErrorResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ErrorInfo) {
		toSerialize["errorInfo"] = o.ErrorInfo
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResponseError struct {
	value *ErrorResponseError
	isSet bool
}

func (v NullableErrorResponseError) Get() *ErrorResponseError {
	return v.value
}

func (v *NullableErrorResponseError) Set(val *ErrorResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseError(val *ErrorResponseError) *NullableErrorResponseError {
	return &NullableErrorResponseError{value: val, isSet: true}
}

func (v NullableErrorResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


