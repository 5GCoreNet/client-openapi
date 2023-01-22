/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
)

// SmPolicySnssaiDataPatch Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiDataPatch struct {
	Snssai Snssai `json:"snssai"`
	// Modifiable Session Management Policy data per DNN for all the DNNs of the indicated S-NSSAI. The key of the map is the DNN. 
	SmPolicyDnnData *map[string]SmPolicyDnnDataPatch `json:"smPolicyDnnData,omitempty"`
}

// NewSmPolicySnssaiDataPatch instantiates a new SmPolicySnssaiDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmPolicySnssaiDataPatch(snssai Snssai) *SmPolicySnssaiDataPatch {
	this := SmPolicySnssaiDataPatch{}
	this.Snssai = snssai
	return &this
}

// NewSmPolicySnssaiDataPatchWithDefaults instantiates a new SmPolicySnssaiDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmPolicySnssaiDataPatchWithDefaults() *SmPolicySnssaiDataPatch {
	this := SmPolicySnssaiDataPatch{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *SmPolicySnssaiDataPatch) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SmPolicySnssaiDataPatch) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SmPolicySnssaiDataPatch) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetSmPolicyDnnData returns the SmPolicyDnnData field value if set, zero value otherwise.
func (o *SmPolicySnssaiDataPatch) GetSmPolicyDnnData() map[string]SmPolicyDnnDataPatch {
	if o == nil || isNil(o.SmPolicyDnnData) {
		var ret map[string]SmPolicyDnnDataPatch
		return ret
	}
	return *o.SmPolicyDnnData
}

// GetSmPolicyDnnDataOk returns a tuple with the SmPolicyDnnData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicySnssaiDataPatch) GetSmPolicyDnnDataOk() (*map[string]SmPolicyDnnDataPatch, bool) {
	if o == nil || isNil(o.SmPolicyDnnData) {
    return nil, false
	}
	return o.SmPolicyDnnData, true
}

// HasSmPolicyDnnData returns a boolean if a field has been set.
func (o *SmPolicySnssaiDataPatch) HasSmPolicyDnnData() bool {
	if o != nil && !isNil(o.SmPolicyDnnData) {
		return true
	}

	return false
}

// SetSmPolicyDnnData gets a reference to the given map[string]SmPolicyDnnDataPatch and assigns it to the SmPolicyDnnData field.
func (o *SmPolicySnssaiDataPatch) SetSmPolicyDnnData(v map[string]SmPolicyDnnDataPatch) {
	o.SmPolicyDnnData = &v
}

func (o SmPolicySnssaiDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.SmPolicyDnnData) {
		toSerialize["smPolicyDnnData"] = o.SmPolicyDnnData
	}
	return json.Marshal(toSerialize)
}

type NullableSmPolicySnssaiDataPatch struct {
	value *SmPolicySnssaiDataPatch
	isSet bool
}

func (v NullableSmPolicySnssaiDataPatch) Get() *SmPolicySnssaiDataPatch {
	return v.value
}

func (v *NullableSmPolicySnssaiDataPatch) Set(val *SmPolicySnssaiDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicySnssaiDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicySnssaiDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicySnssaiDataPatch(val *SmPolicySnssaiDataPatch) *NullableSmPolicySnssaiDataPatch {
	return &NullableSmPolicySnssaiDataPatch{value: val, isSet: true}
}

func (v NullableSmPolicySnssaiDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicySnssaiDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


