/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// SupiRange A range of SUPIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type SupiRange struct {
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}

// NewSupiRange instantiates a new SupiRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupiRange() *SupiRange {
	this := SupiRange{}
	return &this
}

// NewSupiRangeWithDefaults instantiates a new SupiRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupiRangeWithDefaults() *SupiRange {
	this := SupiRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SupiRange) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupiRange) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SupiRange) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *SupiRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SupiRange) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupiRange) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SupiRange) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *SupiRange) SetEnd(v string) {
	o.End = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SupiRange) GetPattern() string {
	if o == nil || isNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupiRange) GetPatternOk() (*string, bool) {
	if o == nil || isNil(o.Pattern) {
    return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SupiRange) HasPattern() bool {
	if o != nil && !isNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SupiRange) SetPattern(v string) {
	o.Pattern = &v
}

func (o SupiRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableSupiRange struct {
	value *SupiRange
	isSet bool
}

func (v NullableSupiRange) Get() *SupiRange {
	return v.value
}

func (v *NullableSupiRange) Set(val *SupiRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSupiRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSupiRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupiRange(val *SupiRange) *NullableSupiRange {
	return &NullableSupiRange{value: val, isSet: true}
}

func (v NullableSupiRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupiRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


