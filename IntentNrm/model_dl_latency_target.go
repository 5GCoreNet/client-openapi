/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// DLLatencyTarget This data type is the \"ExpectationTarget\" data type with specialisations for DLLatencyTarget    
type DLLatencyTarget struct {
	TargetName *string `json:"targetName,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
}

// NewDLLatencyTarget instantiates a new DLLatencyTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDLLatencyTarget() *DLLatencyTarget {
	this := DLLatencyTarget{}
	return &this
}

// NewDLLatencyTargetWithDefaults instantiates a new DLLatencyTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDLLatencyTargetWithDefaults() *DLLatencyTarget {
	this := DLLatencyTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DLLatencyTarget) GetTargetName() string {
	if o == nil || isNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLLatencyTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetName) {
    return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DLLatencyTarget) HasTargetName() bool {
	if o != nil && !isNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DLLatencyTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *DLLatencyTarget) GetTargetCondition() string {
	if o == nil || isNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLLatencyTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || isNil(o.TargetCondition) {
    return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *DLLatencyTarget) HasTargetCondition() bool {
	if o != nil && !isNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *DLLatencyTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *DLLatencyTarget) GetTargetValueRange() int32 {
	if o == nil || isNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLLatencyTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || isNil(o.TargetValueRange) {
    return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *DLLatencyTarget) HasTargetValueRange() bool {
	if o != nil && !isNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *DLLatencyTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

func (o DLLatencyTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TargetName) {
		toSerialize["targetName"] = o.TargetName
	}
	if !isNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !isNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	return json.Marshal(toSerialize)
}

type NullableDLLatencyTarget struct {
	value *DLLatencyTarget
	isSet bool
}

func (v NullableDLLatencyTarget) Get() *DLLatencyTarget {
	return v.value
}

func (v *NullableDLLatencyTarget) Set(val *DLLatencyTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableDLLatencyTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableDLLatencyTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDLLatencyTarget(val *DLLatencyTarget) *NullableDLLatencyTarget {
	return &NullableDLLatencyTarget{value: val, isSet: true}
}

func (v NullableDLLatencyTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDLLatencyTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


