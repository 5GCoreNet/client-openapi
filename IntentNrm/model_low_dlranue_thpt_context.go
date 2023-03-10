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

// LowDLRANUEThptContext This data type is the \"TargetContext\" data type with specialisations for LowDLRANUEThptContext    
type LowDLRANUEThptContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *string `json:"contextCondition,omitempty"`
	ContextValueRange *float32 `json:"contextValueRange,omitempty"`
}

// NewLowDLRANUEThptContext instantiates a new LowDLRANUEThptContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowDLRANUEThptContext() *LowDLRANUEThptContext {
	this := LowDLRANUEThptContext{}
	return &this
}

// NewLowDLRANUEThptContextWithDefaults instantiates a new LowDLRANUEThptContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowDLRANUEThptContextWithDefaults() *LowDLRANUEThptContext {
	this := LowDLRANUEThptContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *LowDLRANUEThptContext) GetContextAttribute() string {
	if o == nil || isNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowDLRANUEThptContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ContextAttribute) {
    return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *LowDLRANUEThptContext) HasContextAttribute() bool {
	if o != nil && !isNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *LowDLRANUEThptContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *LowDLRANUEThptContext) GetContextCondition() string {
	if o == nil || isNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowDLRANUEThptContext) GetContextConditionOk() (*string, bool) {
	if o == nil || isNil(o.ContextCondition) {
    return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *LowDLRANUEThptContext) HasContextCondition() bool {
	if o != nil && !isNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *LowDLRANUEThptContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *LowDLRANUEThptContext) GetContextValueRange() float32 {
	if o == nil || isNil(o.ContextValueRange) {
		var ret float32
		return ret
	}
	return *o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowDLRANUEThptContext) GetContextValueRangeOk() (*float32, bool) {
	if o == nil || isNil(o.ContextValueRange) {
    return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *LowDLRANUEThptContext) HasContextValueRange() bool {
	if o != nil && !isNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given float32 and assigns it to the ContextValueRange field.
func (o *LowDLRANUEThptContext) SetContextValueRange(v float32) {
	o.ContextValueRange = &v
}

func (o LowDLRANUEThptContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContextAttribute) {
		toSerialize["contextAttribute"] = o.ContextAttribute
	}
	if !isNil(o.ContextCondition) {
		toSerialize["contextCondition"] = o.ContextCondition
	}
	if !isNil(o.ContextValueRange) {
		toSerialize["contextValueRange"] = o.ContextValueRange
	}
	return json.Marshal(toSerialize)
}

type NullableLowDLRANUEThptContext struct {
	value *LowDLRANUEThptContext
	isSet bool
}

func (v NullableLowDLRANUEThptContext) Get() *LowDLRANUEThptContext {
	return v.value
}

func (v *NullableLowDLRANUEThptContext) Set(val *LowDLRANUEThptContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLowDLRANUEThptContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLowDLRANUEThptContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowDLRANUEThptContext(val *LowDLRANUEThptContext) *NullableLowDLRANUEThptContext {
	return &NullableLowDLRANUEThptContext{value: val, isSet: true}
}

func (v NullableLowDLRANUEThptContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowDLRANUEThptContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


