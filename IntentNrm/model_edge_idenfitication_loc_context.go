/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package IntentNrm

import (
	"encoding/json"
)

// EdgeIdenfiticationLocContext This data type is the \"ObjectContext\" data type with specialisations for EdgeIdenfiticationLocContext   
type EdgeIdenfiticationLocContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *string `json:"contextCondition,omitempty"`
	ContextValueRange []string `json:"contextValueRange,omitempty"`
}

// NewEdgeIdenfiticationLocContext instantiates a new EdgeIdenfiticationLocContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeIdenfiticationLocContext() *EdgeIdenfiticationLocContext {
	this := EdgeIdenfiticationLocContext{}
	return &this
}

// NewEdgeIdenfiticationLocContextWithDefaults instantiates a new EdgeIdenfiticationLocContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeIdenfiticationLocContextWithDefaults() *EdgeIdenfiticationLocContext {
	this := EdgeIdenfiticationLocContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *EdgeIdenfiticationLocContext) GetContextAttribute() string {
	if o == nil || isNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeIdenfiticationLocContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ContextAttribute) {
    return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *EdgeIdenfiticationLocContext) HasContextAttribute() bool {
	if o != nil && !isNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *EdgeIdenfiticationLocContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *EdgeIdenfiticationLocContext) GetContextCondition() string {
	if o == nil || isNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeIdenfiticationLocContext) GetContextConditionOk() (*string, bool) {
	if o == nil || isNil(o.ContextCondition) {
    return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *EdgeIdenfiticationLocContext) HasContextCondition() bool {
	if o != nil && !isNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *EdgeIdenfiticationLocContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *EdgeIdenfiticationLocContext) GetContextValueRange() []string {
	if o == nil || isNil(o.ContextValueRange) {
		var ret []string
		return ret
	}
	return o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeIdenfiticationLocContext) GetContextValueRangeOk() ([]string, bool) {
	if o == nil || isNil(o.ContextValueRange) {
    return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *EdgeIdenfiticationLocContext) HasContextValueRange() bool {
	if o != nil && !isNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given []string and assigns it to the ContextValueRange field.
func (o *EdgeIdenfiticationLocContext) SetContextValueRange(v []string) {
	o.ContextValueRange = v
}

func (o EdgeIdenfiticationLocContext) MarshalJSON() ([]byte, error) {
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

type NullableEdgeIdenfiticationLocContext struct {
	value *EdgeIdenfiticationLocContext
	isSet bool
}

func (v NullableEdgeIdenfiticationLocContext) Get() *EdgeIdenfiticationLocContext {
	return v.value
}

func (v *NullableEdgeIdenfiticationLocContext) Set(val *EdgeIdenfiticationLocContext) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeIdenfiticationLocContext) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeIdenfiticationLocContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeIdenfiticationLocContext(val *EdgeIdenfiticationLocContext) *NullableEdgeIdenfiticationLocContext {
	return &NullableEdgeIdenfiticationLocContext{value: val, isSet: true}
}

func (v NullableEdgeIdenfiticationLocContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeIdenfiticationLocContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


