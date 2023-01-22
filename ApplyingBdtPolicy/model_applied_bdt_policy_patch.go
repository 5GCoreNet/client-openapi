/*
3gpp-applying-bdt-policy

API for applying BDT policy   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ApplyingBdtPolicy

import (
	"encoding/json"
)

// AppliedBdtPolicyPatch Represents the parameters to request the modification of a subscription to applied BDT policy. 
type AppliedBdtPolicyPatch struct {
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
}

// NewAppliedBdtPolicyPatch instantiates a new AppliedBdtPolicyPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedBdtPolicyPatch(bdtRefId string) *AppliedBdtPolicyPatch {
	this := AppliedBdtPolicyPatch{}
	this.BdtRefId = bdtRefId
	return &this
}

// NewAppliedBdtPolicyPatchWithDefaults instantiates a new AppliedBdtPolicyPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedBdtPolicyPatchWithDefaults() *AppliedBdtPolicyPatch {
	this := AppliedBdtPolicyPatch{}
	return &this
}

// GetBdtRefId returns the BdtRefId field value
func (o *AppliedBdtPolicyPatch) GetBdtRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value
// and a boolean to check if the value has been set.
func (o *AppliedBdtPolicyPatch) GetBdtRefIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BdtRefId, true
}

// SetBdtRefId sets field value
func (o *AppliedBdtPolicyPatch) SetBdtRefId(v string) {
	o.BdtRefId = v
}

func (o AppliedBdtPolicyPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bdtRefId"] = o.BdtRefId
	}
	return json.Marshal(toSerialize)
}

type NullableAppliedBdtPolicyPatch struct {
	value *AppliedBdtPolicyPatch
	isSet bool
}

func (v NullableAppliedBdtPolicyPatch) Get() *AppliedBdtPolicyPatch {
	return v.value
}

func (v *NullableAppliedBdtPolicyPatch) Set(val *AppliedBdtPolicyPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedBdtPolicyPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedBdtPolicyPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedBdtPolicyPatch(val *AppliedBdtPolicyPatch) *NullableAppliedBdtPolicyPatch {
	return &NullableAppliedBdtPolicyPatch{value: val, isSet: true}
}

func (v NullableAppliedBdtPolicyPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedBdtPolicyPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


