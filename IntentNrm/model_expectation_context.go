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

// ExpectationContext This data type is the \"ExpectationContext\" data type without specialisations       
type ExpectationContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *Condition `json:"contextCondition,omitempty"`
	ContextValueRange []float32 `json:"contextValueRange,omitempty"`
}

// NewExpectationContext instantiates a new ExpectationContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpectationContext() *ExpectationContext {
	this := ExpectationContext{}
	return &this
}

// NewExpectationContextWithDefaults instantiates a new ExpectationContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpectationContextWithDefaults() *ExpectationContext {
	this := ExpectationContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *ExpectationContext) GetContextAttribute() string {
	if o == nil || isNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ContextAttribute) {
    return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *ExpectationContext) HasContextAttribute() bool {
	if o != nil && !isNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *ExpectationContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *ExpectationContext) GetContextCondition() Condition {
	if o == nil || isNil(o.ContextCondition) {
		var ret Condition
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationContext) GetContextConditionOk() (*Condition, bool) {
	if o == nil || isNil(o.ContextCondition) {
    return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *ExpectationContext) HasContextCondition() bool {
	if o != nil && !isNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given Condition and assigns it to the ContextCondition field.
func (o *ExpectationContext) SetContextCondition(v Condition) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *ExpectationContext) GetContextValueRange() []float32 {
	if o == nil || isNil(o.ContextValueRange) {
		var ret []float32
		return ret
	}
	return o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationContext) GetContextValueRangeOk() ([]float32, bool) {
	if o == nil || isNil(o.ContextValueRange) {
    return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *ExpectationContext) HasContextValueRange() bool {
	if o != nil && !isNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given []float32 and assigns it to the ContextValueRange field.
func (o *ExpectationContext) SetContextValueRange(v []float32) {
	o.ContextValueRange = v
}

func (o ExpectationContext) MarshalJSON() ([]byte, error) {
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

type NullableExpectationContext struct {
	value *ExpectationContext
	isSet bool
}

func (v NullableExpectationContext) Get() *ExpectationContext {
	return v.value
}

func (v *NullableExpectationContext) Set(val *ExpectationContext) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectationContext) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectationContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectationContext(val *ExpectationContext) *NullableExpectationContext {
	return &NullableExpectationContext{value: val, isSet: true}
}

func (v NullableExpectationContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectationContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


