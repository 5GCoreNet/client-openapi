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

// AssuranceClosedControlLoopSingleAllOf struct for AssuranceClosedControlLoopSingleAllOf
type AssuranceClosedControlLoopSingleAllOf struct {
	Attributes *AssuranceClosedControlLoopSingleAllOfAttributes `json:"attributes,omitempty"`
	AssuranceGoal []AssuranceGoalSingle `json:"AssuranceGoal,omitempty"`
}

// NewAssuranceClosedControlLoopSingleAllOf instantiates a new AssuranceClosedControlLoopSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceClosedControlLoopSingleAllOf() *AssuranceClosedControlLoopSingleAllOf {
	this := AssuranceClosedControlLoopSingleAllOf{}
	return &this
}

// NewAssuranceClosedControlLoopSingleAllOfWithDefaults instantiates a new AssuranceClosedControlLoopSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceClosedControlLoopSingleAllOfWithDefaults() *AssuranceClosedControlLoopSingleAllOf {
	this := AssuranceClosedControlLoopSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOf) GetAttributes() AssuranceClosedControlLoopSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret AssuranceClosedControlLoopSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOf) GetAttributesOk() (*AssuranceClosedControlLoopSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AssuranceClosedControlLoopSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AssuranceClosedControlLoopSingleAllOf) SetAttributes(v AssuranceClosedControlLoopSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetAssuranceGoal returns the AssuranceGoal field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOf) GetAssuranceGoal() []AssuranceGoalSingle {
	if o == nil || isNil(o.AssuranceGoal) {
		var ret []AssuranceGoalSingle
		return ret
	}
	return o.AssuranceGoal
}

// GetAssuranceGoalOk returns a tuple with the AssuranceGoal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOf) GetAssuranceGoalOk() ([]AssuranceGoalSingle, bool) {
	if o == nil || isNil(o.AssuranceGoal) {
    return nil, false
	}
	return o.AssuranceGoal, true
}

// HasAssuranceGoal returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOf) HasAssuranceGoal() bool {
	if o != nil && !isNil(o.AssuranceGoal) {
		return true
	}

	return false
}

// SetAssuranceGoal gets a reference to the given []AssuranceGoalSingle and assigns it to the AssuranceGoal field.
func (o *AssuranceClosedControlLoopSingleAllOf) SetAssuranceGoal(v []AssuranceGoalSingle) {
	o.AssuranceGoal = v
}

func (o AssuranceClosedControlLoopSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.AssuranceGoal) {
		toSerialize["AssuranceGoal"] = o.AssuranceGoal
	}
	return json.Marshal(toSerialize)
}

type NullableAssuranceClosedControlLoopSingleAllOf struct {
	value *AssuranceClosedControlLoopSingleAllOf
	isSet bool
}

func (v NullableAssuranceClosedControlLoopSingleAllOf) Get() *AssuranceClosedControlLoopSingleAllOf {
	return v.value
}

func (v *NullableAssuranceClosedControlLoopSingleAllOf) Set(val *AssuranceClosedControlLoopSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceClosedControlLoopSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceClosedControlLoopSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceClosedControlLoopSingleAllOf(val *AssuranceClosedControlLoopSingleAllOf) *NullableAssuranceClosedControlLoopSingleAllOf {
	return &NullableAssuranceClosedControlLoopSingleAllOf{value: val, isSet: true}
}

func (v NullableAssuranceClosedControlLoopSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceClosedControlLoopSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


