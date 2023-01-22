/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUECM

import (
	"encoding/json"
)

// ScscfRestorationInfoResponse S-CSCF restoration information response
type ScscfRestorationInfoResponse struct {
	ScscfRestorationInfoResponse []ScscfRestorationInfo `json:"scscfRestorationInfoResponse,omitempty"`
}

// NewScscfRestorationInfoResponse instantiates a new ScscfRestorationInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScscfRestorationInfoResponse() *ScscfRestorationInfoResponse {
	this := ScscfRestorationInfoResponse{}
	return &this
}

// NewScscfRestorationInfoResponseWithDefaults instantiates a new ScscfRestorationInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScscfRestorationInfoResponseWithDefaults() *ScscfRestorationInfoResponse {
	this := ScscfRestorationInfoResponse{}
	return &this
}

// GetScscfRestorationInfoResponse returns the ScscfRestorationInfoResponse field value if set, zero value otherwise.
func (o *ScscfRestorationInfoResponse) GetScscfRestorationInfoResponse() []ScscfRestorationInfo {
	if o == nil || isNil(o.ScscfRestorationInfoResponse) {
		var ret []ScscfRestorationInfo
		return ret
	}
	return o.ScscfRestorationInfoResponse
}

// GetScscfRestorationInfoResponseOk returns a tuple with the ScscfRestorationInfoResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfoResponse) GetScscfRestorationInfoResponseOk() ([]ScscfRestorationInfo, bool) {
	if o == nil || isNil(o.ScscfRestorationInfoResponse) {
    return nil, false
	}
	return o.ScscfRestorationInfoResponse, true
}

// HasScscfRestorationInfoResponse returns a boolean if a field has been set.
func (o *ScscfRestorationInfoResponse) HasScscfRestorationInfoResponse() bool {
	if o != nil && !isNil(o.ScscfRestorationInfoResponse) {
		return true
	}

	return false
}

// SetScscfRestorationInfoResponse gets a reference to the given []ScscfRestorationInfo and assigns it to the ScscfRestorationInfoResponse field.
func (o *ScscfRestorationInfoResponse) SetScscfRestorationInfoResponse(v []ScscfRestorationInfo) {
	o.ScscfRestorationInfoResponse = v
}

func (o ScscfRestorationInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScscfRestorationInfoResponse) {
		toSerialize["scscfRestorationInfoResponse"] = o.ScscfRestorationInfoResponse
	}
	return json.Marshal(toSerialize)
}

type NullableScscfRestorationInfoResponse struct {
	value *ScscfRestorationInfoResponse
	isSet bool
}

func (v NullableScscfRestorationInfoResponse) Get() *ScscfRestorationInfoResponse {
	return v.value
}

func (v *NullableScscfRestorationInfoResponse) Set(val *ScscfRestorationInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScscfRestorationInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScscfRestorationInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScscfRestorationInfoResponse(val *ScscfRestorationInfoResponse) *NullableScscfRestorationInfoResponse {
	return &NullableScscfRestorationInfoResponse{value: val, isSet: true}
}

func (v NullableScscfRestorationInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScscfRestorationInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

