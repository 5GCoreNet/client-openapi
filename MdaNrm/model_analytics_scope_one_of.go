/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MdaNrm

import (
	"encoding/json"
)

// AnalyticsScopeOneOf struct for AnalyticsScopeOneOf
type AnalyticsScopeOneOf struct {
	ManagedEntitiesScope []string `json:"managedEntitiesScope,omitempty"`
}

// NewAnalyticsScopeOneOf instantiates a new AnalyticsScopeOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsScopeOneOf() *AnalyticsScopeOneOf {
	this := AnalyticsScopeOneOf{}
	return &this
}

// NewAnalyticsScopeOneOfWithDefaults instantiates a new AnalyticsScopeOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsScopeOneOfWithDefaults() *AnalyticsScopeOneOf {
	this := AnalyticsScopeOneOf{}
	return &this
}

// GetManagedEntitiesScope returns the ManagedEntitiesScope field value if set, zero value otherwise.
func (o *AnalyticsScopeOneOf) GetManagedEntitiesScope() []string {
	if o == nil || isNil(o.ManagedEntitiesScope) {
		var ret []string
		return ret
	}
	return o.ManagedEntitiesScope
}

// GetManagedEntitiesScopeOk returns a tuple with the ManagedEntitiesScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsScopeOneOf) GetManagedEntitiesScopeOk() ([]string, bool) {
	if o == nil || isNil(o.ManagedEntitiesScope) {
    return nil, false
	}
	return o.ManagedEntitiesScope, true
}

// HasManagedEntitiesScope returns a boolean if a field has been set.
func (o *AnalyticsScopeOneOf) HasManagedEntitiesScope() bool {
	if o != nil && !isNil(o.ManagedEntitiesScope) {
		return true
	}

	return false
}

// SetManagedEntitiesScope gets a reference to the given []string and assigns it to the ManagedEntitiesScope field.
func (o *AnalyticsScopeOneOf) SetManagedEntitiesScope(v []string) {
	o.ManagedEntitiesScope = v
}

func (o AnalyticsScopeOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ManagedEntitiesScope) {
		toSerialize["managedEntitiesScope"] = o.ManagedEntitiesScope
	}
	return json.Marshal(toSerialize)
}

type NullableAnalyticsScopeOneOf struct {
	value *AnalyticsScopeOneOf
	isSet bool
}

func (v NullableAnalyticsScopeOneOf) Get() *AnalyticsScopeOneOf {
	return v.value
}

func (v *NullableAnalyticsScopeOneOf) Set(val *AnalyticsScopeOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsScopeOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsScopeOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsScopeOneOf(val *AnalyticsScopeOneOf) *NullableAnalyticsScopeOneOf {
	return &NullableAnalyticsScopeOneOf{value: val, isSet: true}
}

func (v NullableAnalyticsScopeOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsScopeOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


