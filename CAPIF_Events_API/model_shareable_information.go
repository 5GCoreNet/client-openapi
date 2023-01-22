/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Events_API

import (
	"encoding/json"
)

// ShareableInformation Indicates whether the service API and/or the service API category can be shared to the list of CAPIF provider domains. 
type ShareableInformation struct {
	// Set to \"true\" indicates that the service API and/or the service API category can be shared to the list of CAPIF provider domain information. Otherwise set to \"false\". 
	IsShareable bool `json:"isShareable"`
	// List of CAPIF provider domains to which the service API information to be shared. 
	CapifProvDoms []string `json:"capifProvDoms,omitempty"`
}

// NewShareableInformation instantiates a new ShareableInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareableInformation(isShareable bool) *ShareableInformation {
	this := ShareableInformation{}
	this.IsShareable = isShareable
	return &this
}

// NewShareableInformationWithDefaults instantiates a new ShareableInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareableInformationWithDefaults() *ShareableInformation {
	this := ShareableInformation{}
	return &this
}

// GetIsShareable returns the IsShareable field value
func (o *ShareableInformation) GetIsShareable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsShareable
}

// GetIsShareableOk returns a tuple with the IsShareable field value
// and a boolean to check if the value has been set.
func (o *ShareableInformation) GetIsShareableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsShareable, true
}

// SetIsShareable sets field value
func (o *ShareableInformation) SetIsShareable(v bool) {
	o.IsShareable = v
}

// GetCapifProvDoms returns the CapifProvDoms field value if set, zero value otherwise.
func (o *ShareableInformation) GetCapifProvDoms() []string {
	if o == nil || isNil(o.CapifProvDoms) {
		var ret []string
		return ret
	}
	return o.CapifProvDoms
}

// GetCapifProvDomsOk returns a tuple with the CapifProvDoms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareableInformation) GetCapifProvDomsOk() ([]string, bool) {
	if o == nil || isNil(o.CapifProvDoms) {
    return nil, false
	}
	return o.CapifProvDoms, true
}

// HasCapifProvDoms returns a boolean if a field has been set.
func (o *ShareableInformation) HasCapifProvDoms() bool {
	if o != nil && !isNil(o.CapifProvDoms) {
		return true
	}

	return false
}

// SetCapifProvDoms gets a reference to the given []string and assigns it to the CapifProvDoms field.
func (o *ShareableInformation) SetCapifProvDoms(v []string) {
	o.CapifProvDoms = v
}

func (o ShareableInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isShareable"] = o.IsShareable
	}
	if !isNil(o.CapifProvDoms) {
		toSerialize["capifProvDoms"] = o.CapifProvDoms
	}
	return json.Marshal(toSerialize)
}

type NullableShareableInformation struct {
	value *ShareableInformation
	isSet bool
}

func (v NullableShareableInformation) Get() *ShareableInformation {
	return v.value
}

func (v *NullableShareableInformation) Set(val *ShareableInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableShareableInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableShareableInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareableInformation(val *ShareableInformation) *NullableShareableInformation {
	return &NullableShareableInformation{value: val, isSet: true}
}

func (v NullableShareableInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareableInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


