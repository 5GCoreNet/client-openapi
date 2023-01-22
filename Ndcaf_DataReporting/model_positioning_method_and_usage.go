/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndcaf_DataReporting

import (
	"encoding/json"
)

// PositioningMethodAndUsage Indicates the usage of a positioning method.
type PositioningMethodAndUsage struct {
	Method PositioningMethod `json:"method"`
	Mode PositioningMode `json:"mode"`
	Usage Usage `json:"usage"`
	MethodCode *int32 `json:"methodCode,omitempty"`
}

// NewPositioningMethodAndUsage instantiates a new PositioningMethodAndUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositioningMethodAndUsage(method PositioningMethod, mode PositioningMode, usage Usage) *PositioningMethodAndUsage {
	this := PositioningMethodAndUsage{}
	this.Method = method
	this.Mode = mode
	this.Usage = usage
	return &this
}

// NewPositioningMethodAndUsageWithDefaults instantiates a new PositioningMethodAndUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositioningMethodAndUsageWithDefaults() *PositioningMethodAndUsage {
	this := PositioningMethodAndUsage{}
	return &this
}

// GetMethod returns the Method field value
func (o *PositioningMethodAndUsage) GetMethod() PositioningMethod {
	if o == nil {
		var ret PositioningMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *PositioningMethodAndUsage) GetMethodOk() (*PositioningMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *PositioningMethodAndUsage) SetMethod(v PositioningMethod) {
	o.Method = v
}

// GetMode returns the Mode field value
func (o *PositioningMethodAndUsage) GetMode() PositioningMode {
	if o == nil {
		var ret PositioningMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PositioningMethodAndUsage) GetModeOk() (*PositioningMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PositioningMethodAndUsage) SetMode(v PositioningMode) {
	o.Mode = v
}

// GetUsage returns the Usage field value
func (o *PositioningMethodAndUsage) GetUsage() Usage {
	if o == nil {
		var ret Usage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *PositioningMethodAndUsage) GetUsageOk() (*Usage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *PositioningMethodAndUsage) SetUsage(v Usage) {
	o.Usage = v
}

// GetMethodCode returns the MethodCode field value if set, zero value otherwise.
func (o *PositioningMethodAndUsage) GetMethodCode() int32 {
	if o == nil || isNil(o.MethodCode) {
		var ret int32
		return ret
	}
	return *o.MethodCode
}

// GetMethodCodeOk returns a tuple with the MethodCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositioningMethodAndUsage) GetMethodCodeOk() (*int32, bool) {
	if o == nil || isNil(o.MethodCode) {
    return nil, false
	}
	return o.MethodCode, true
}

// HasMethodCode returns a boolean if a field has been set.
func (o *PositioningMethodAndUsage) HasMethodCode() bool {
	if o != nil && !isNil(o.MethodCode) {
		return true
	}

	return false
}

// SetMethodCode gets a reference to the given int32 and assigns it to the MethodCode field.
func (o *PositioningMethodAndUsage) SetMethodCode(v int32) {
	o.MethodCode = &v
}

func (o PositioningMethodAndUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.MethodCode) {
		toSerialize["methodCode"] = o.MethodCode
	}
	return json.Marshal(toSerialize)
}

type NullablePositioningMethodAndUsage struct {
	value *PositioningMethodAndUsage
	isSet bool
}

func (v NullablePositioningMethodAndUsage) Get() *PositioningMethodAndUsage {
	return v.value
}

func (v *NullablePositioningMethodAndUsage) Set(val *PositioningMethodAndUsage) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethodAndUsage) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethodAndUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethodAndUsage(val *PositioningMethodAndUsage) *NullablePositioningMethodAndUsage {
	return &NullablePositioningMethodAndUsage{value: val, isSet: true}
}

func (v NullablePositioningMethodAndUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethodAndUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


