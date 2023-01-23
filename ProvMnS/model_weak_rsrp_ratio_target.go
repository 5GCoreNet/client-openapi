/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// WeakRSRPRatioTarget This data type is the \"ExpectationTarget\" data type with specialisations for WeakRSRPRatioTarget
type WeakRSRPRatioTarget struct {
	TargetName *string `json:"targetName,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
	TargetContexts *WeakRSRPContext `json:"targetContexts,omitempty"`
	TargetFulfilmentInfo *FulfilmentInfo `json:"targetFulfilmentInfo,omitempty"`
}

// NewWeakRSRPRatioTarget instantiates a new WeakRSRPRatioTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeakRSRPRatioTarget() *WeakRSRPRatioTarget {
	this := WeakRSRPRatioTarget{}
	return &this
}

// NewWeakRSRPRatioTargetWithDefaults instantiates a new WeakRSRPRatioTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeakRSRPRatioTargetWithDefaults() *WeakRSRPRatioTarget {
	this := WeakRSRPRatioTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *WeakRSRPRatioTarget) GetTargetName() string {
	if o == nil || isNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeakRSRPRatioTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetName) {
    return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *WeakRSRPRatioTarget) HasTargetName() bool {
	if o != nil && !isNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *WeakRSRPRatioTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *WeakRSRPRatioTarget) GetTargetCondition() string {
	if o == nil || isNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeakRSRPRatioTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || isNil(o.TargetCondition) {
    return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *WeakRSRPRatioTarget) HasTargetCondition() bool {
	if o != nil && !isNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *WeakRSRPRatioTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *WeakRSRPRatioTarget) GetTargetValueRange() int32 {
	if o == nil || isNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeakRSRPRatioTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || isNil(o.TargetValueRange) {
    return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *WeakRSRPRatioTarget) HasTargetValueRange() bool {
	if o != nil && !isNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *WeakRSRPRatioTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

// GetTargetContexts returns the TargetContexts field value if set, zero value otherwise.
func (o *WeakRSRPRatioTarget) GetTargetContexts() WeakRSRPContext {
	if o == nil || isNil(o.TargetContexts) {
		var ret WeakRSRPContext
		return ret
	}
	return *o.TargetContexts
}

// GetTargetContextsOk returns a tuple with the TargetContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeakRSRPRatioTarget) GetTargetContextsOk() (*WeakRSRPContext, bool) {
	if o == nil || isNil(o.TargetContexts) {
    return nil, false
	}
	return o.TargetContexts, true
}

// HasTargetContexts returns a boolean if a field has been set.
func (o *WeakRSRPRatioTarget) HasTargetContexts() bool {
	if o != nil && !isNil(o.TargetContexts) {
		return true
	}

	return false
}

// SetTargetContexts gets a reference to the given WeakRSRPContext and assigns it to the TargetContexts field.
func (o *WeakRSRPRatioTarget) SetTargetContexts(v WeakRSRPContext) {
	o.TargetContexts = &v
}

// GetTargetFulfilmentInfo returns the TargetFulfilmentInfo field value if set, zero value otherwise.
func (o *WeakRSRPRatioTarget) GetTargetFulfilmentInfo() FulfilmentInfo {
	if o == nil || isNil(o.TargetFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.TargetFulfilmentInfo
}

// GetTargetFulfilmentInfoOk returns a tuple with the TargetFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeakRSRPRatioTarget) GetTargetFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || isNil(o.TargetFulfilmentInfo) {
    return nil, false
	}
	return o.TargetFulfilmentInfo, true
}

// HasTargetFulfilmentInfo returns a boolean if a field has been set.
func (o *WeakRSRPRatioTarget) HasTargetFulfilmentInfo() bool {
	if o != nil && !isNil(o.TargetFulfilmentInfo) {
		return true
	}

	return false
}

// SetTargetFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the TargetFulfilmentInfo field.
func (o *WeakRSRPRatioTarget) SetTargetFulfilmentInfo(v FulfilmentInfo) {
	o.TargetFulfilmentInfo = &v
}

func (o WeakRSRPRatioTarget) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.TargetContexts) {
		toSerialize["targetContexts"] = o.TargetContexts
	}
	if !isNil(o.TargetFulfilmentInfo) {
		toSerialize["targetFulfilmentInfo"] = o.TargetFulfilmentInfo
	}
	return json.Marshal(toSerialize)
}

type NullableWeakRSRPRatioTarget struct {
	value *WeakRSRPRatioTarget
	isSet bool
}

func (v NullableWeakRSRPRatioTarget) Get() *WeakRSRPRatioTarget {
	return v.value
}

func (v *NullableWeakRSRPRatioTarget) Set(val *WeakRSRPRatioTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableWeakRSRPRatioTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableWeakRSRPRatioTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeakRSRPRatioTarget(val *WeakRSRPRatioTarget) *NullableWeakRSRPRatioTarget {
	return &NullableWeakRSRPRatioTarget{value: val, isSet: true}
}

func (v NullableWeakRSRPRatioTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeakRSRPRatioTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


