/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// BWPSetSingleAllOf struct for BWPSetSingleAllOf
type BWPSetSingleAllOf struct {
	BWPlist []string `json:"bWPlist,omitempty"`
}

// NewBWPSetSingleAllOf instantiates a new BWPSetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBWPSetSingleAllOf() *BWPSetSingleAllOf {
	this := BWPSetSingleAllOf{}
	return &this
}

// NewBWPSetSingleAllOfWithDefaults instantiates a new BWPSetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBWPSetSingleAllOfWithDefaults() *BWPSetSingleAllOf {
	this := BWPSetSingleAllOf{}
	return &this
}

// GetBWPlist returns the BWPlist field value if set, zero value otherwise.
func (o *BWPSetSingleAllOf) GetBWPlist() []string {
	if o == nil || isNil(o.BWPlist) {
		var ret []string
		return ret
	}
	return o.BWPlist
}

// GetBWPlistOk returns a tuple with the BWPlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingleAllOf) GetBWPlistOk() ([]string, bool) {
	if o == nil || isNil(o.BWPlist) {
    return nil, false
	}
	return o.BWPlist, true
}

// HasBWPlist returns a boolean if a field has been set.
func (o *BWPSetSingleAllOf) HasBWPlist() bool {
	if o != nil && !isNil(o.BWPlist) {
		return true
	}

	return false
}

// SetBWPlist gets a reference to the given []string and assigns it to the BWPlist field.
func (o *BWPSetSingleAllOf) SetBWPlist(v []string) {
	o.BWPlist = v
}

func (o BWPSetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BWPlist) {
		toSerialize["bWPlist"] = o.BWPlist
	}
	return json.Marshal(toSerialize)
}

type NullableBWPSetSingleAllOf struct {
	value *BWPSetSingleAllOf
	isSet bool
}

func (v NullableBWPSetSingleAllOf) Get() *BWPSetSingleAllOf {
	return v.value
}

func (v *NullableBWPSetSingleAllOf) Set(val *BWPSetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBWPSetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBWPSetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBWPSetSingleAllOf(val *BWPSetSingleAllOf) *NullableBWPSetSingleAllOf {
	return &NullableBWPSetSingleAllOf{value: val, isSet: true}
}

func (v NullableBWPSetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBWPSetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


