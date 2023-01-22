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

// AccessControlPolicyListExt Represents the extension for access control policies.
type AccessControlPolicyListExt struct {
	// Policy of each API invoker.
	ApiInvokerPolicies []ApiInvokerPolicy `json:"apiInvokerPolicies,omitempty"`
	ApiId string `json:"apiId"`
}

// NewAccessControlPolicyListExt instantiates a new AccessControlPolicyListExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessControlPolicyListExt(apiId string) *AccessControlPolicyListExt {
	this := AccessControlPolicyListExt{}
	this.ApiId = apiId
	return &this
}

// NewAccessControlPolicyListExtWithDefaults instantiates a new AccessControlPolicyListExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessControlPolicyListExtWithDefaults() *AccessControlPolicyListExt {
	this := AccessControlPolicyListExt{}
	return &this
}

// GetApiInvokerPolicies returns the ApiInvokerPolicies field value if set, zero value otherwise.
func (o *AccessControlPolicyListExt) GetApiInvokerPolicies() []ApiInvokerPolicy {
	if o == nil || isNil(o.ApiInvokerPolicies) {
		var ret []ApiInvokerPolicy
		return ret
	}
	return o.ApiInvokerPolicies
}

// GetApiInvokerPoliciesOk returns a tuple with the ApiInvokerPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyListExt) GetApiInvokerPoliciesOk() ([]ApiInvokerPolicy, bool) {
	if o == nil || isNil(o.ApiInvokerPolicies) {
    return nil, false
	}
	return o.ApiInvokerPolicies, true
}

// HasApiInvokerPolicies returns a boolean if a field has been set.
func (o *AccessControlPolicyListExt) HasApiInvokerPolicies() bool {
	if o != nil && !isNil(o.ApiInvokerPolicies) {
		return true
	}

	return false
}

// SetApiInvokerPolicies gets a reference to the given []ApiInvokerPolicy and assigns it to the ApiInvokerPolicies field.
func (o *AccessControlPolicyListExt) SetApiInvokerPolicies(v []ApiInvokerPolicy) {
	o.ApiInvokerPolicies = v
}

// GetApiId returns the ApiId field value
func (o *AccessControlPolicyListExt) GetApiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyListExt) GetApiIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiId, true
}

// SetApiId sets field value
func (o *AccessControlPolicyListExt) SetApiId(v string) {
	o.ApiId = v
}

func (o AccessControlPolicyListExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiInvokerPolicies) {
		toSerialize["apiInvokerPolicies"] = o.ApiInvokerPolicies
	}
	if true {
		toSerialize["apiId"] = o.ApiId
	}
	return json.Marshal(toSerialize)
}

type NullableAccessControlPolicyListExt struct {
	value *AccessControlPolicyListExt
	isSet bool
}

func (v NullableAccessControlPolicyListExt) Get() *AccessControlPolicyListExt {
	return v.value
}

func (v *NullableAccessControlPolicyListExt) Set(val *AccessControlPolicyListExt) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessControlPolicyListExt) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessControlPolicyListExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessControlPolicyListExt(val *AccessControlPolicyListExt) *NullableAccessControlPolicyListExt {
	return &NullableAccessControlPolicyListExt{value: val, isSet: true}
}

func (v NullableAccessControlPolicyListExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessControlPolicyListExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


