/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
)

// BootstrappingInfoRequest Request body of the HTTP POST operation for resource /bootstrapping-info-request
type BootstrappingInfoRequest struct {
	// Bootstrapping Transaction Identifier
	BtId string `json:"btId"`
	NafId NafId `json:"nafId"`
	GbaUAware *bool `json:"gbaUAware,omitempty"`
	GsIds []int32 `json:"gsIds,omitempty"`
}

// NewBootstrappingInfoRequest instantiates a new BootstrappingInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootstrappingInfoRequest(btId string, nafId NafId) *BootstrappingInfoRequest {
	this := BootstrappingInfoRequest{}
	this.BtId = btId
	this.NafId = nafId
	var gbaUAware bool = false
	this.GbaUAware = &gbaUAware
	return &this
}

// NewBootstrappingInfoRequestWithDefaults instantiates a new BootstrappingInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootstrappingInfoRequestWithDefaults() *BootstrappingInfoRequest {
	this := BootstrappingInfoRequest{}
	var gbaUAware bool = false
	this.GbaUAware = &gbaUAware
	return &this
}

// GetBtId returns the BtId field value
func (o *BootstrappingInfoRequest) GetBtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtId
}

// GetBtIdOk returns a tuple with the BtId field value
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoRequest) GetBtIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BtId, true
}

// SetBtId sets field value
func (o *BootstrappingInfoRequest) SetBtId(v string) {
	o.BtId = v
}

// GetNafId returns the NafId field value
func (o *BootstrappingInfoRequest) GetNafId() NafId {
	if o == nil {
		var ret NafId
		return ret
	}

	return o.NafId
}

// GetNafIdOk returns a tuple with the NafId field value
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoRequest) GetNafIdOk() (*NafId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NafId, true
}

// SetNafId sets field value
func (o *BootstrappingInfoRequest) SetNafId(v NafId) {
	o.NafId = v
}

// GetGbaUAware returns the GbaUAware field value if set, zero value otherwise.
func (o *BootstrappingInfoRequest) GetGbaUAware() bool {
	if o == nil || isNil(o.GbaUAware) {
		var ret bool
		return ret
	}
	return *o.GbaUAware
}

// GetGbaUAwareOk returns a tuple with the GbaUAware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoRequest) GetGbaUAwareOk() (*bool, bool) {
	if o == nil || isNil(o.GbaUAware) {
    return nil, false
	}
	return o.GbaUAware, true
}

// HasGbaUAware returns a boolean if a field has been set.
func (o *BootstrappingInfoRequest) HasGbaUAware() bool {
	if o != nil && !isNil(o.GbaUAware) {
		return true
	}

	return false
}

// SetGbaUAware gets a reference to the given bool and assigns it to the GbaUAware field.
func (o *BootstrappingInfoRequest) SetGbaUAware(v bool) {
	o.GbaUAware = &v
}

// GetGsIds returns the GsIds field value if set, zero value otherwise.
func (o *BootstrappingInfoRequest) GetGsIds() []int32 {
	if o == nil || isNil(o.GsIds) {
		var ret []int32
		return ret
	}
	return o.GsIds
}

// GetGsIdsOk returns a tuple with the GsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoRequest) GetGsIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.GsIds) {
    return nil, false
	}
	return o.GsIds, true
}

// HasGsIds returns a boolean if a field has been set.
func (o *BootstrappingInfoRequest) HasGsIds() bool {
	if o != nil && !isNil(o.GsIds) {
		return true
	}

	return false
}

// SetGsIds gets a reference to the given []int32 and assigns it to the GsIds field.
func (o *BootstrappingInfoRequest) SetGsIds(v []int32) {
	o.GsIds = v
}

func (o BootstrappingInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btId"] = o.BtId
	}
	if true {
		toSerialize["nafId"] = o.NafId
	}
	if !isNil(o.GbaUAware) {
		toSerialize["gbaUAware"] = o.GbaUAware
	}
	if !isNil(o.GsIds) {
		toSerialize["gsIds"] = o.GsIds
	}
	return json.Marshal(toSerialize)
}

type NullableBootstrappingInfoRequest struct {
	value *BootstrappingInfoRequest
	isSet bool
}

func (v NullableBootstrappingInfoRequest) Get() *BootstrappingInfoRequest {
	return v.value
}

func (v *NullableBootstrappingInfoRequest) Set(val *BootstrappingInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBootstrappingInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBootstrappingInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootstrappingInfoRequest(val *BootstrappingInfoRequest) *NullableBootstrappingInfoRequest {
	return &NullableBootstrappingInfoRequest{value: val, isSet: true}
}

func (v NullableBootstrappingInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootstrappingInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


