/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
)

// SearchCondition A logical condition
type SearchCondition struct {
	Cond ConditionOperator `json:"cond"`
	Units []SearchExpression `json:"units"`
	// Represents the Identifier of a Meta schema.
	SchemaId *string `json:"schemaId,omitempty"`
}

// NewSearchCondition instantiates a new SearchCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCondition(cond ConditionOperator, units []SearchExpression) *SearchCondition {
	this := SearchCondition{}
	this.Cond = cond
	this.Units = units
	return &this
}

// NewSearchConditionWithDefaults instantiates a new SearchCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchConditionWithDefaults() *SearchCondition {
	this := SearchCondition{}
	return &this
}

// GetCond returns the Cond field value
func (o *SearchCondition) GetCond() ConditionOperator {
	if o == nil {
		var ret ConditionOperator
		return ret
	}

	return o.Cond
}

// GetCondOk returns a tuple with the Cond field value
// and a boolean to check if the value has been set.
func (o *SearchCondition) GetCondOk() (*ConditionOperator, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cond, true
}

// SetCond sets field value
func (o *SearchCondition) SetCond(v ConditionOperator) {
	o.Cond = v
}

// GetUnits returns the Units field value
func (o *SearchCondition) GetUnits() []SearchExpression {
	if o == nil {
		var ret []SearchExpression
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *SearchCondition) GetUnitsOk() ([]SearchExpression, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units, true
}

// SetUnits sets field value
func (o *SearchCondition) SetUnits(v []SearchExpression) {
	o.Units = v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *SearchCondition) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCondition) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *SearchCondition) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *SearchCondition) SetSchemaId(v string) {
	o.SchemaId = &v
}

func (o SearchCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cond"] = o.Cond
	}
	if true {
		toSerialize["units"] = o.Units
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableSearchCondition struct {
	value *SearchCondition
	isSet bool
}

func (v NullableSearchCondition) Get() *SearchCondition {
	return v.value
}

func (v *NullableSearchCondition) Set(val *SearchCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCondition(val *SearchCondition) *NullableSearchCondition {
	return &NullableSearchCondition{value: val, isSet: true}
}

func (v NullableSearchCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


