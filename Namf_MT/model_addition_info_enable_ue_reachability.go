/*
Namf_MT

AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_MT

import (
	"encoding/json"
)

// AdditionInfoEnableUeReachability Additional information to be returned in EnableUeReachability error response
type AdditionInfoEnableUeReachability struct {
	// indicating a time in seconds.
	MaxWaitingTime *int32 `json:"maxWaitingTime,omitempty"`
}

// NewAdditionInfoEnableUeReachability instantiates a new AdditionInfoEnableUeReachability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionInfoEnableUeReachability() *AdditionInfoEnableUeReachability {
	this := AdditionInfoEnableUeReachability{}
	return &this
}

// NewAdditionInfoEnableUeReachabilityWithDefaults instantiates a new AdditionInfoEnableUeReachability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionInfoEnableUeReachabilityWithDefaults() *AdditionInfoEnableUeReachability {
	this := AdditionInfoEnableUeReachability{}
	return &this
}

// GetMaxWaitingTime returns the MaxWaitingTime field value if set, zero value otherwise.
func (o *AdditionInfoEnableUeReachability) GetMaxWaitingTime() int32 {
	if o == nil || isNil(o.MaxWaitingTime) {
		var ret int32
		return ret
	}
	return *o.MaxWaitingTime
}

// GetMaxWaitingTimeOk returns a tuple with the MaxWaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionInfoEnableUeReachability) GetMaxWaitingTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxWaitingTime) {
    return nil, false
	}
	return o.MaxWaitingTime, true
}

// HasMaxWaitingTime returns a boolean if a field has been set.
func (o *AdditionInfoEnableUeReachability) HasMaxWaitingTime() bool {
	if o != nil && !isNil(o.MaxWaitingTime) {
		return true
	}

	return false
}

// SetMaxWaitingTime gets a reference to the given int32 and assigns it to the MaxWaitingTime field.
func (o *AdditionInfoEnableUeReachability) SetMaxWaitingTime(v int32) {
	o.MaxWaitingTime = &v
}

func (o AdditionInfoEnableUeReachability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxWaitingTime) {
		toSerialize["maxWaitingTime"] = o.MaxWaitingTime
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionInfoEnableUeReachability struct {
	value *AdditionInfoEnableUeReachability
	isSet bool
}

func (v NullableAdditionInfoEnableUeReachability) Get() *AdditionInfoEnableUeReachability {
	return v.value
}

func (v *NullableAdditionInfoEnableUeReachability) Set(val *AdditionInfoEnableUeReachability) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionInfoEnableUeReachability) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionInfoEnableUeReachability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionInfoEnableUeReachability(val *AdditionInfoEnableUeReachability) *NullableAdditionInfoEnableUeReachability {
	return &NullableAdditionInfoEnableUeReachability{value: val, isSet: true}
}

func (v NullableAdditionInfoEnableUeReachability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionInfoEnableUeReachability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


