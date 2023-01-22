/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// ChargingNotifyResponse struct for ChargingNotifyResponse
type ChargingNotifyResponse struct {
	InvocationResult *InvocationResult `json:"invocationResult,omitempty"`
}

// NewChargingNotifyResponse instantiates a new ChargingNotifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingNotifyResponse() *ChargingNotifyResponse {
	this := ChargingNotifyResponse{}
	return &this
}

// NewChargingNotifyResponseWithDefaults instantiates a new ChargingNotifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingNotifyResponseWithDefaults() *ChargingNotifyResponse {
	this := ChargingNotifyResponse{}
	return &this
}

// GetInvocationResult returns the InvocationResult field value if set, zero value otherwise.
func (o *ChargingNotifyResponse) GetInvocationResult() InvocationResult {
	if o == nil || isNil(o.InvocationResult) {
		var ret InvocationResult
		return ret
	}
	return *o.InvocationResult
}

// GetInvocationResultOk returns a tuple with the InvocationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingNotifyResponse) GetInvocationResultOk() (*InvocationResult, bool) {
	if o == nil || isNil(o.InvocationResult) {
    return nil, false
	}
	return o.InvocationResult, true
}

// HasInvocationResult returns a boolean if a field has been set.
func (o *ChargingNotifyResponse) HasInvocationResult() bool {
	if o != nil && !isNil(o.InvocationResult) {
		return true
	}

	return false
}

// SetInvocationResult gets a reference to the given InvocationResult and assigns it to the InvocationResult field.
func (o *ChargingNotifyResponse) SetInvocationResult(v InvocationResult) {
	o.InvocationResult = &v
}

func (o ChargingNotifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InvocationResult) {
		toSerialize["invocationResult"] = o.InvocationResult
	}
	return json.Marshal(toSerialize)
}

type NullableChargingNotifyResponse struct {
	value *ChargingNotifyResponse
	isSet bool
}

func (v NullableChargingNotifyResponse) Get() *ChargingNotifyResponse {
	return v.value
}

func (v *NullableChargingNotifyResponse) Set(val *ChargingNotifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingNotifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingNotifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingNotifyResponse(val *ChargingNotifyResponse) *NullableChargingNotifyResponse {
	return &NullableChargingNotifyResponse{value: val, isSet: true}
}

func (v NullableChargingNotifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingNotifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


