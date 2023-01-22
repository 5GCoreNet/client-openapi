/*
CAPIF_Access_Control_Policy_API

API for access control policy.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Access_Control_Policy_API

import (
	"encoding/json"
)

// ApiInvokerPolicy Represents the policy of an API Invoker.
type ApiInvokerPolicy struct {
	// API invoker ID assigned by the CAPIF core function
	ApiInvokerId string `json:"apiInvokerId"`
	// Total number of invocations allowed on the service API by the API invoker.
	AllowedTotalInvocations *int32 `json:"allowedTotalInvocations,omitempty"`
	// Invocations per second allowed on the service API by the API invoker.
	AllowedInvocationsPerSecond *int32 `json:"allowedInvocationsPerSecond,omitempty"`
	// The time ranges during which the invocations are allowed on the service API by the API invoker. 
	AllowedInvocationTimeRangeList []TimeRangeList `json:"allowedInvocationTimeRangeList,omitempty"`
}

// NewApiInvokerPolicy instantiates a new ApiInvokerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInvokerPolicy(apiInvokerId string) *ApiInvokerPolicy {
	this := ApiInvokerPolicy{}
	this.ApiInvokerId = apiInvokerId
	return &this
}

// NewApiInvokerPolicyWithDefaults instantiates a new ApiInvokerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInvokerPolicyWithDefaults() *ApiInvokerPolicy {
	this := ApiInvokerPolicy{}
	return &this
}

// GetApiInvokerId returns the ApiInvokerId field value
func (o *ApiInvokerPolicy) GetApiInvokerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiInvokerId
}

// GetApiInvokerIdOk returns a tuple with the ApiInvokerId field value
// and a boolean to check if the value has been set.
func (o *ApiInvokerPolicy) GetApiInvokerIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiInvokerId, true
}

// SetApiInvokerId sets field value
func (o *ApiInvokerPolicy) SetApiInvokerId(v string) {
	o.ApiInvokerId = v
}

// GetAllowedTotalInvocations returns the AllowedTotalInvocations field value if set, zero value otherwise.
func (o *ApiInvokerPolicy) GetAllowedTotalInvocations() int32 {
	if o == nil || isNil(o.AllowedTotalInvocations) {
		var ret int32
		return ret
	}
	return *o.AllowedTotalInvocations
}

// GetAllowedTotalInvocationsOk returns a tuple with the AllowedTotalInvocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInvokerPolicy) GetAllowedTotalInvocationsOk() (*int32, bool) {
	if o == nil || isNil(o.AllowedTotalInvocations) {
    return nil, false
	}
	return o.AllowedTotalInvocations, true
}

// HasAllowedTotalInvocations returns a boolean if a field has been set.
func (o *ApiInvokerPolicy) HasAllowedTotalInvocations() bool {
	if o != nil && !isNil(o.AllowedTotalInvocations) {
		return true
	}

	return false
}

// SetAllowedTotalInvocations gets a reference to the given int32 and assigns it to the AllowedTotalInvocations field.
func (o *ApiInvokerPolicy) SetAllowedTotalInvocations(v int32) {
	o.AllowedTotalInvocations = &v
}

// GetAllowedInvocationsPerSecond returns the AllowedInvocationsPerSecond field value if set, zero value otherwise.
func (o *ApiInvokerPolicy) GetAllowedInvocationsPerSecond() int32 {
	if o == nil || isNil(o.AllowedInvocationsPerSecond) {
		var ret int32
		return ret
	}
	return *o.AllowedInvocationsPerSecond
}

// GetAllowedInvocationsPerSecondOk returns a tuple with the AllowedInvocationsPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInvokerPolicy) GetAllowedInvocationsPerSecondOk() (*int32, bool) {
	if o == nil || isNil(o.AllowedInvocationsPerSecond) {
    return nil, false
	}
	return o.AllowedInvocationsPerSecond, true
}

// HasAllowedInvocationsPerSecond returns a boolean if a field has been set.
func (o *ApiInvokerPolicy) HasAllowedInvocationsPerSecond() bool {
	if o != nil && !isNil(o.AllowedInvocationsPerSecond) {
		return true
	}

	return false
}

// SetAllowedInvocationsPerSecond gets a reference to the given int32 and assigns it to the AllowedInvocationsPerSecond field.
func (o *ApiInvokerPolicy) SetAllowedInvocationsPerSecond(v int32) {
	o.AllowedInvocationsPerSecond = &v
}

// GetAllowedInvocationTimeRangeList returns the AllowedInvocationTimeRangeList field value if set, zero value otherwise.
func (o *ApiInvokerPolicy) GetAllowedInvocationTimeRangeList() []TimeRangeList {
	if o == nil || isNil(o.AllowedInvocationTimeRangeList) {
		var ret []TimeRangeList
		return ret
	}
	return o.AllowedInvocationTimeRangeList
}

// GetAllowedInvocationTimeRangeListOk returns a tuple with the AllowedInvocationTimeRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInvokerPolicy) GetAllowedInvocationTimeRangeListOk() ([]TimeRangeList, bool) {
	if o == nil || isNil(o.AllowedInvocationTimeRangeList) {
    return nil, false
	}
	return o.AllowedInvocationTimeRangeList, true
}

// HasAllowedInvocationTimeRangeList returns a boolean if a field has been set.
func (o *ApiInvokerPolicy) HasAllowedInvocationTimeRangeList() bool {
	if o != nil && !isNil(o.AllowedInvocationTimeRangeList) {
		return true
	}

	return false
}

// SetAllowedInvocationTimeRangeList gets a reference to the given []TimeRangeList and assigns it to the AllowedInvocationTimeRangeList field.
func (o *ApiInvokerPolicy) SetAllowedInvocationTimeRangeList(v []TimeRangeList) {
	o.AllowedInvocationTimeRangeList = v
}

func (o ApiInvokerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiInvokerId"] = o.ApiInvokerId
	}
	if !isNil(o.AllowedTotalInvocations) {
		toSerialize["allowedTotalInvocations"] = o.AllowedTotalInvocations
	}
	if !isNil(o.AllowedInvocationsPerSecond) {
		toSerialize["allowedInvocationsPerSecond"] = o.AllowedInvocationsPerSecond
	}
	if !isNil(o.AllowedInvocationTimeRangeList) {
		toSerialize["allowedInvocationTimeRangeList"] = o.AllowedInvocationTimeRangeList
	}
	return json.Marshal(toSerialize)
}

type NullableApiInvokerPolicy struct {
	value *ApiInvokerPolicy
	isSet bool
}

func (v NullableApiInvokerPolicy) Get() *ApiInvokerPolicy {
	return v.value
}

func (v *NullableApiInvokerPolicy) Set(val *ApiInvokerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInvokerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInvokerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInvokerPolicy(val *ApiInvokerPolicy) *NullableApiInvokerPolicy {
	return &NullableApiInvokerPolicy{value: val, isSet: true}
}

func (v NullableApiInvokerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInvokerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


