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

// AveULRANUEThptTarget This data type is the \"ExpectationTarget\" data type with specialisations for AveULRANUEThptTarget
type AveULRANUEThptTarget struct {
	TargetName *string `json:"targetName,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
	TargetFulfilmentInfo *FulfilmentInfo `json:"targetFulfilmentInfo,omitempty"`
}

// NewAveULRANUEThptTarget instantiates a new AveULRANUEThptTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAveULRANUEThptTarget() *AveULRANUEThptTarget {
	this := AveULRANUEThptTarget{}
	return &this
}

// NewAveULRANUEThptTargetWithDefaults instantiates a new AveULRANUEThptTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAveULRANUEThptTargetWithDefaults() *AveULRANUEThptTarget {
	this := AveULRANUEThptTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *AveULRANUEThptTarget) GetTargetName() string {
	if o == nil || isNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AveULRANUEThptTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetName) {
    return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *AveULRANUEThptTarget) HasTargetName() bool {
	if o != nil && !isNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *AveULRANUEThptTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *AveULRANUEThptTarget) GetTargetCondition() string {
	if o == nil || isNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AveULRANUEThptTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || isNil(o.TargetCondition) {
    return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *AveULRANUEThptTarget) HasTargetCondition() bool {
	if o != nil && !isNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *AveULRANUEThptTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *AveULRANUEThptTarget) GetTargetValueRange() int32 {
	if o == nil || isNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AveULRANUEThptTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || isNil(o.TargetValueRange) {
    return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *AveULRANUEThptTarget) HasTargetValueRange() bool {
	if o != nil && !isNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *AveULRANUEThptTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

// GetTargetFulfilmentInfo returns the TargetFulfilmentInfo field value if set, zero value otherwise.
func (o *AveULRANUEThptTarget) GetTargetFulfilmentInfo() FulfilmentInfo {
	if o == nil || isNil(o.TargetFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.TargetFulfilmentInfo
}

// GetTargetFulfilmentInfoOk returns a tuple with the TargetFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AveULRANUEThptTarget) GetTargetFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || isNil(o.TargetFulfilmentInfo) {
    return nil, false
	}
	return o.TargetFulfilmentInfo, true
}

// HasTargetFulfilmentInfo returns a boolean if a field has been set.
func (o *AveULRANUEThptTarget) HasTargetFulfilmentInfo() bool {
	if o != nil && !isNil(o.TargetFulfilmentInfo) {
		return true
	}

	return false
}

// SetTargetFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the TargetFulfilmentInfo field.
func (o *AveULRANUEThptTarget) SetTargetFulfilmentInfo(v FulfilmentInfo) {
	o.TargetFulfilmentInfo = &v
}

func (o AveULRANUEThptTarget) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.TargetFulfilmentInfo) {
		toSerialize["targetFulfilmentInfo"] = o.TargetFulfilmentInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAveULRANUEThptTarget struct {
	value *AveULRANUEThptTarget
	isSet bool
}

func (v NullableAveULRANUEThptTarget) Get() *AveULRANUEThptTarget {
	return v.value
}

func (v *NullableAveULRANUEThptTarget) Set(val *AveULRANUEThptTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAveULRANUEThptTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAveULRANUEThptTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAveULRANUEThptTarget(val *AveULRANUEThptTarget) *NullableAveULRANUEThptTarget {
	return &NullableAveULRANUEThptTarget{value: val, isSet: true}
}

func (v NullableAveULRANUEThptTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAveULRANUEThptTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


