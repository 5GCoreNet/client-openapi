/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// MnsRegistrySingle struct for MnsRegistrySingle
type MnsRegistrySingle struct {
	MnsInfo []MnsInfoSingle `json:"MnsInfo,omitempty"`
}

// NewMnsRegistrySingle instantiates a new MnsRegistrySingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnsRegistrySingle() *MnsRegistrySingle {
	this := MnsRegistrySingle{}
	return &this
}

// NewMnsRegistrySingleWithDefaults instantiates a new MnsRegistrySingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnsRegistrySingleWithDefaults() *MnsRegistrySingle {
	this := MnsRegistrySingle{}
	return &this
}

// GetMnsInfo returns the MnsInfo field value if set, zero value otherwise.
func (o *MnsRegistrySingle) GetMnsInfo() []MnsInfoSingle {
	if o == nil || isNil(o.MnsInfo) {
		var ret []MnsInfoSingle
		return ret
	}
	return o.MnsInfo
}

// GetMnsInfoOk returns a tuple with the MnsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsRegistrySingle) GetMnsInfoOk() ([]MnsInfoSingle, bool) {
	if o == nil || isNil(o.MnsInfo) {
    return nil, false
	}
	return o.MnsInfo, true
}

// HasMnsInfo returns a boolean if a field has been set.
func (o *MnsRegistrySingle) HasMnsInfo() bool {
	if o != nil && !isNil(o.MnsInfo) {
		return true
	}

	return false
}

// SetMnsInfo gets a reference to the given []MnsInfoSingle and assigns it to the MnsInfo field.
func (o *MnsRegistrySingle) SetMnsInfo(v []MnsInfoSingle) {
	o.MnsInfo = v
}

func (o MnsRegistrySingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MnsInfo) {
		toSerialize["MnsInfo"] = o.MnsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableMnsRegistrySingle struct {
	value *MnsRegistrySingle
	isSet bool
}

func (v NullableMnsRegistrySingle) Get() *MnsRegistrySingle {
	return v.value
}

func (v *NullableMnsRegistrySingle) Set(val *MnsRegistrySingle) {
	v.value = val
	v.isSet = true
}

func (v NullableMnsRegistrySingle) IsSet() bool {
	return v.isSet
}

func (v *NullableMnsRegistrySingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnsRegistrySingle(val *MnsRegistrySingle) *NullableMnsRegistrySingle {
	return &NullableMnsRegistrySingle{value: val, isSet: true}
}

func (v NullableMnsRegistrySingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnsRegistrySingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


