/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AnalyticsExposure

import (
	"encoding/json"
)

// TargetUeId Represents the target UE(s) information.
type TargetUeId struct {
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExterGroupId *string `json:"exterGroupId,omitempty"`
}

// NewTargetUeId instantiates a new TargetUeId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetUeId() *TargetUeId {
	this := TargetUeId{}
	return &this
}

// NewTargetUeIdWithDefaults instantiates a new TargetUeId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetUeIdWithDefaults() *TargetUeId {
	this := TargetUeId{}
	return &this
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *TargetUeId) GetAnyUeInd() bool {
	if o == nil || isNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeId) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || isNil(o.AnyUeInd) {
    return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *TargetUeId) HasAnyUeInd() bool {
	if o != nil && !isNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *TargetUeId) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *TargetUeId) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeId) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *TargetUeId) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *TargetUeId) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExterGroupId returns the ExterGroupId field value if set, zero value otherwise.
func (o *TargetUeId) GetExterGroupId() string {
	if o == nil || isNil(o.ExterGroupId) {
		var ret string
		return ret
	}
	return *o.ExterGroupId
}

// GetExterGroupIdOk returns a tuple with the ExterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeId) GetExterGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExterGroupId) {
    return nil, false
	}
	return o.ExterGroupId, true
}

// HasExterGroupId returns a boolean if a field has been set.
func (o *TargetUeId) HasExterGroupId() bool {
	if o != nil && !isNil(o.ExterGroupId) {
		return true
	}

	return false
}

// SetExterGroupId gets a reference to the given string and assigns it to the ExterGroupId field.
func (o *TargetUeId) SetExterGroupId(v string) {
	o.ExterGroupId = &v
}

func (o TargetUeId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.ExterGroupId) {
		toSerialize["exterGroupId"] = o.ExterGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableTargetUeId struct {
	value *TargetUeId
	isSet bool
}

func (v NullableTargetUeId) Get() *TargetUeId {
	return v.value
}

func (v *NullableTargetUeId) Set(val *TargetUeId) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetUeId) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetUeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetUeId(val *TargetUeId) *NullableTargetUeId {
	return &NullableTargetUeId{value: val, isSet: true}
}

func (v NullableTargetUeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetUeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


