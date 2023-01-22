/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// ObservedRedundantTransExp Represents the observed redundant transmission experience related information.
type ObservedRedundantTransExp struct {
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPktDropRateUl *int32 `json:"avgPktDropRateUl,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	VarPktDropRateUl *float32 `json:"varPktDropRateUl,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPktDropRateDl *int32 `json:"avgPktDropRateDl,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	VarPktDropRateDl *float32 `json:"varPktDropRateDl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvgPktDelayUl *int32 `json:"avgPktDelayUl,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	VarPktDelayUl *float32 `json:"varPktDelayUl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvgPktDelayDl *int32 `json:"avgPktDelayDl,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	VarPktDelayDl *float32 `json:"varPktDelayDl,omitempty"`
}

// NewObservedRedundantTransExp instantiates a new ObservedRedundantTransExp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservedRedundantTransExp() *ObservedRedundantTransExp {
	this := ObservedRedundantTransExp{}
	return &this
}

// NewObservedRedundantTransExpWithDefaults instantiates a new ObservedRedundantTransExp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservedRedundantTransExpWithDefaults() *ObservedRedundantTransExp {
	this := ObservedRedundantTransExp{}
	return &this
}

// GetAvgPktDropRateUl returns the AvgPktDropRateUl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetAvgPktDropRateUl() int32 {
	if o == nil || isNil(o.AvgPktDropRateUl) {
		var ret int32
		return ret
	}
	return *o.AvgPktDropRateUl
}

// GetAvgPktDropRateUlOk returns a tuple with the AvgPktDropRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetAvgPktDropRateUlOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPktDropRateUl) {
    return nil, false
	}
	return o.AvgPktDropRateUl, true
}

// HasAvgPktDropRateUl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasAvgPktDropRateUl() bool {
	if o != nil && !isNil(o.AvgPktDropRateUl) {
		return true
	}

	return false
}

// SetAvgPktDropRateUl gets a reference to the given int32 and assigns it to the AvgPktDropRateUl field.
func (o *ObservedRedundantTransExp) SetAvgPktDropRateUl(v int32) {
	o.AvgPktDropRateUl = &v
}

// GetVarPktDropRateUl returns the VarPktDropRateUl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetVarPktDropRateUl() float32 {
	if o == nil || isNil(o.VarPktDropRateUl) {
		var ret float32
		return ret
	}
	return *o.VarPktDropRateUl
}

// GetVarPktDropRateUlOk returns a tuple with the VarPktDropRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetVarPktDropRateUlOk() (*float32, bool) {
	if o == nil || isNil(o.VarPktDropRateUl) {
    return nil, false
	}
	return o.VarPktDropRateUl, true
}

// HasVarPktDropRateUl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasVarPktDropRateUl() bool {
	if o != nil && !isNil(o.VarPktDropRateUl) {
		return true
	}

	return false
}

// SetVarPktDropRateUl gets a reference to the given float32 and assigns it to the VarPktDropRateUl field.
func (o *ObservedRedundantTransExp) SetVarPktDropRateUl(v float32) {
	o.VarPktDropRateUl = &v
}

// GetAvgPktDropRateDl returns the AvgPktDropRateDl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetAvgPktDropRateDl() int32 {
	if o == nil || isNil(o.AvgPktDropRateDl) {
		var ret int32
		return ret
	}
	return *o.AvgPktDropRateDl
}

// GetAvgPktDropRateDlOk returns a tuple with the AvgPktDropRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetAvgPktDropRateDlOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPktDropRateDl) {
    return nil, false
	}
	return o.AvgPktDropRateDl, true
}

// HasAvgPktDropRateDl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasAvgPktDropRateDl() bool {
	if o != nil && !isNil(o.AvgPktDropRateDl) {
		return true
	}

	return false
}

// SetAvgPktDropRateDl gets a reference to the given int32 and assigns it to the AvgPktDropRateDl field.
func (o *ObservedRedundantTransExp) SetAvgPktDropRateDl(v int32) {
	o.AvgPktDropRateDl = &v
}

// GetVarPktDropRateDl returns the VarPktDropRateDl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetVarPktDropRateDl() float32 {
	if o == nil || isNil(o.VarPktDropRateDl) {
		var ret float32
		return ret
	}
	return *o.VarPktDropRateDl
}

// GetVarPktDropRateDlOk returns a tuple with the VarPktDropRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetVarPktDropRateDlOk() (*float32, bool) {
	if o == nil || isNil(o.VarPktDropRateDl) {
    return nil, false
	}
	return o.VarPktDropRateDl, true
}

// HasVarPktDropRateDl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasVarPktDropRateDl() bool {
	if o != nil && !isNil(o.VarPktDropRateDl) {
		return true
	}

	return false
}

// SetVarPktDropRateDl gets a reference to the given float32 and assigns it to the VarPktDropRateDl field.
func (o *ObservedRedundantTransExp) SetVarPktDropRateDl(v float32) {
	o.VarPktDropRateDl = &v
}

// GetAvgPktDelayUl returns the AvgPktDelayUl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetAvgPktDelayUl() int32 {
	if o == nil || isNil(o.AvgPktDelayUl) {
		var ret int32
		return ret
	}
	return *o.AvgPktDelayUl
}

// GetAvgPktDelayUlOk returns a tuple with the AvgPktDelayUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetAvgPktDelayUlOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPktDelayUl) {
    return nil, false
	}
	return o.AvgPktDelayUl, true
}

// HasAvgPktDelayUl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasAvgPktDelayUl() bool {
	if o != nil && !isNil(o.AvgPktDelayUl) {
		return true
	}

	return false
}

// SetAvgPktDelayUl gets a reference to the given int32 and assigns it to the AvgPktDelayUl field.
func (o *ObservedRedundantTransExp) SetAvgPktDelayUl(v int32) {
	o.AvgPktDelayUl = &v
}

// GetVarPktDelayUl returns the VarPktDelayUl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetVarPktDelayUl() float32 {
	if o == nil || isNil(o.VarPktDelayUl) {
		var ret float32
		return ret
	}
	return *o.VarPktDelayUl
}

// GetVarPktDelayUlOk returns a tuple with the VarPktDelayUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetVarPktDelayUlOk() (*float32, bool) {
	if o == nil || isNil(o.VarPktDelayUl) {
    return nil, false
	}
	return o.VarPktDelayUl, true
}

// HasVarPktDelayUl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasVarPktDelayUl() bool {
	if o != nil && !isNil(o.VarPktDelayUl) {
		return true
	}

	return false
}

// SetVarPktDelayUl gets a reference to the given float32 and assigns it to the VarPktDelayUl field.
func (o *ObservedRedundantTransExp) SetVarPktDelayUl(v float32) {
	o.VarPktDelayUl = &v
}

// GetAvgPktDelayDl returns the AvgPktDelayDl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetAvgPktDelayDl() int32 {
	if o == nil || isNil(o.AvgPktDelayDl) {
		var ret int32
		return ret
	}
	return *o.AvgPktDelayDl
}

// GetAvgPktDelayDlOk returns a tuple with the AvgPktDelayDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetAvgPktDelayDlOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPktDelayDl) {
    return nil, false
	}
	return o.AvgPktDelayDl, true
}

// HasAvgPktDelayDl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasAvgPktDelayDl() bool {
	if o != nil && !isNil(o.AvgPktDelayDl) {
		return true
	}

	return false
}

// SetAvgPktDelayDl gets a reference to the given int32 and assigns it to the AvgPktDelayDl field.
func (o *ObservedRedundantTransExp) SetAvgPktDelayDl(v int32) {
	o.AvgPktDelayDl = &v
}

// GetVarPktDelayDl returns the VarPktDelayDl field value if set, zero value otherwise.
func (o *ObservedRedundantTransExp) GetVarPktDelayDl() float32 {
	if o == nil || isNil(o.VarPktDelayDl) {
		var ret float32
		return ret
	}
	return *o.VarPktDelayDl
}

// GetVarPktDelayDlOk returns a tuple with the VarPktDelayDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservedRedundantTransExp) GetVarPktDelayDlOk() (*float32, bool) {
	if o == nil || isNil(o.VarPktDelayDl) {
    return nil, false
	}
	return o.VarPktDelayDl, true
}

// HasVarPktDelayDl returns a boolean if a field has been set.
func (o *ObservedRedundantTransExp) HasVarPktDelayDl() bool {
	if o != nil && !isNil(o.VarPktDelayDl) {
		return true
	}

	return false
}

// SetVarPktDelayDl gets a reference to the given float32 and assigns it to the VarPktDelayDl field.
func (o *ObservedRedundantTransExp) SetVarPktDelayDl(v float32) {
	o.VarPktDelayDl = &v
}

func (o ObservedRedundantTransExp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AvgPktDropRateUl) {
		toSerialize["avgPktDropRateUl"] = o.AvgPktDropRateUl
	}
	if !isNil(o.VarPktDropRateUl) {
		toSerialize["varPktDropRateUl"] = o.VarPktDropRateUl
	}
	if !isNil(o.AvgPktDropRateDl) {
		toSerialize["avgPktDropRateDl"] = o.AvgPktDropRateDl
	}
	if !isNil(o.VarPktDropRateDl) {
		toSerialize["varPktDropRateDl"] = o.VarPktDropRateDl
	}
	if !isNil(o.AvgPktDelayUl) {
		toSerialize["avgPktDelayUl"] = o.AvgPktDelayUl
	}
	if !isNil(o.VarPktDelayUl) {
		toSerialize["varPktDelayUl"] = o.VarPktDelayUl
	}
	if !isNil(o.AvgPktDelayDl) {
		toSerialize["avgPktDelayDl"] = o.AvgPktDelayDl
	}
	if !isNil(o.VarPktDelayDl) {
		toSerialize["varPktDelayDl"] = o.VarPktDelayDl
	}
	return json.Marshal(toSerialize)
}

type NullableObservedRedundantTransExp struct {
	value *ObservedRedundantTransExp
	isSet bool
}

func (v NullableObservedRedundantTransExp) Get() *ObservedRedundantTransExp {
	return v.value
}

func (v *NullableObservedRedundantTransExp) Set(val *ObservedRedundantTransExp) {
	v.value = val
	v.isSet = true
}

func (v NullableObservedRedundantTransExp) IsSet() bool {
	return v.isSet
}

func (v *NullableObservedRedundantTransExp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservedRedundantTransExp(val *ObservedRedundantTransExp) *NullableObservedRedundantTransExp {
	return &NullableObservedRedundantTransExp{value: val, isSet: true}
}

func (v NullableObservedRedundantTransExp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservedRedundantTransExp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


