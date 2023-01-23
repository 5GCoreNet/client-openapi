/*
Nhss_SDM

HSS Subscriber Data Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_SDM

import (
	"encoding/json"
)

// SubscriptionDataSets Contains data to be reported as an immediate report in the response to a subscription creation request
type SubscriptionDataSets struct {
	UeContextInPgwData *UeContextInPgwData `json:"ueContextInPgwData,omitempty"`
}

// NewSubscriptionDataSets instantiates a new SubscriptionDataSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDataSets() *SubscriptionDataSets {
	this := SubscriptionDataSets{}
	return &this
}

// NewSubscriptionDataSetsWithDefaults instantiates a new SubscriptionDataSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataSetsWithDefaults() *SubscriptionDataSets {
	this := SubscriptionDataSets{}
	return &this
}

// GetUeContextInPgwData returns the UeContextInPgwData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetUeContextInPgwData() UeContextInPgwData {
	if o == nil || isNil(o.UeContextInPgwData) {
		var ret UeContextInPgwData
		return ret
	}
	return *o.UeContextInPgwData
}

// GetUeContextInPgwDataOk returns a tuple with the UeContextInPgwData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetUeContextInPgwDataOk() (*UeContextInPgwData, bool) {
	if o == nil || isNil(o.UeContextInPgwData) {
    return nil, false
	}
	return o.UeContextInPgwData, true
}

// HasUeContextInPgwData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasUeContextInPgwData() bool {
	if o != nil && !isNil(o.UeContextInPgwData) {
		return true
	}

	return false
}

// SetUeContextInPgwData gets a reference to the given UeContextInPgwData and assigns it to the UeContextInPgwData field.
func (o *SubscriptionDataSets) SetUeContextInPgwData(v UeContextInPgwData) {
	o.UeContextInPgwData = &v
}

func (o SubscriptionDataSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UeContextInPgwData) {
		toSerialize["ueContextInPgwData"] = o.UeContextInPgwData
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionDataSets struct {
	value *SubscriptionDataSets
	isSet bool
}

func (v NullableSubscriptionDataSets) Get() *SubscriptionDataSets {
	return v.value
}

func (v *NullableSubscriptionDataSets) Set(val *SubscriptionDataSets) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDataSets) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDataSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDataSets(val *SubscriptionDataSets) *NullableSubscriptionDataSets {
	return &NullableSubscriptionDataSets{value: val, isSet: true}
}

func (v NullableSubscriptionDataSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDataSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


