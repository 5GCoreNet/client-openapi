/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// RedundantTransmissionExpReq Represents other redundant transmission experience analytics requirements.
type RedundantTransmissionExpReq struct {
	RedTOrderCriter *RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`
	Order *MatchingDirection `json:"order,omitempty"`
}

// NewRedundantTransmissionExpReq instantiates a new RedundantTransmissionExpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundantTransmissionExpReq() *RedundantTransmissionExpReq {
	this := RedundantTransmissionExpReq{}
	return &this
}

// NewRedundantTransmissionExpReqWithDefaults instantiates a new RedundantTransmissionExpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundantTransmissionExpReqWithDefaults() *RedundantTransmissionExpReq {
	this := RedundantTransmissionExpReq{}
	return &this
}

// GetRedTOrderCriter returns the RedTOrderCriter field value if set, zero value otherwise.
func (o *RedundantTransmissionExpReq) GetRedTOrderCriter() RedTransExpOrderingCriterion {
	if o == nil || isNil(o.RedTOrderCriter) {
		var ret RedTransExpOrderingCriterion
		return ret
	}
	return *o.RedTOrderCriter
}

// GetRedTOrderCriterOk returns a tuple with the RedTOrderCriter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpReq) GetRedTOrderCriterOk() (*RedTransExpOrderingCriterion, bool) {
	if o == nil || isNil(o.RedTOrderCriter) {
    return nil, false
	}
	return o.RedTOrderCriter, true
}

// HasRedTOrderCriter returns a boolean if a field has been set.
func (o *RedundantTransmissionExpReq) HasRedTOrderCriter() bool {
	if o != nil && !isNil(o.RedTOrderCriter) {
		return true
	}

	return false
}

// SetRedTOrderCriter gets a reference to the given RedTransExpOrderingCriterion and assigns it to the RedTOrderCriter field.
func (o *RedundantTransmissionExpReq) SetRedTOrderCriter(v RedTransExpOrderingCriterion) {
	o.RedTOrderCriter = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RedundantTransmissionExpReq) GetOrder() MatchingDirection {
	if o == nil || isNil(o.Order) {
		var ret MatchingDirection
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpReq) GetOrderOk() (*MatchingDirection, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RedundantTransmissionExpReq) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given MatchingDirection and assigns it to the Order field.
func (o *RedundantTransmissionExpReq) SetOrder(v MatchingDirection) {
	o.Order = &v
}

func (o RedundantTransmissionExpReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RedTOrderCriter) {
		toSerialize["redTOrderCriter"] = o.RedTOrderCriter
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableRedundantTransmissionExpReq struct {
	value *RedundantTransmissionExpReq
	isSet bool
}

func (v NullableRedundantTransmissionExpReq) Get() *RedundantTransmissionExpReq {
	return v.value
}

func (v *NullableRedundantTransmissionExpReq) Set(val *RedundantTransmissionExpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantTransmissionExpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantTransmissionExpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantTransmissionExpReq(val *RedundantTransmissionExpReq) *NullableRedundantTransmissionExpReq {
	return &NullableRedundantTransmissionExpReq{value: val, isSet: true}
}

func (v NullableRedundantTransmissionExpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantTransmissionExpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


