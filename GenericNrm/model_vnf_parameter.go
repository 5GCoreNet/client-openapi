/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// VnfParameter struct for VnfParameter
type VnfParameter struct {
	VnfInstanceId *string `json:"vnfInstanceId,omitempty"`
	VnfdId *string `json:"vnfdId,omitempty"`
	FlavourId *string `json:"flavourId,omitempty"`
	AutoScalable *bool `json:"autoScalable,omitempty"`
}

// NewVnfParameter instantiates a new VnfParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnfParameter() *VnfParameter {
	this := VnfParameter{}
	return &this
}

// NewVnfParameterWithDefaults instantiates a new VnfParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnfParameterWithDefaults() *VnfParameter {
	this := VnfParameter{}
	return &this
}

// GetVnfInstanceId returns the VnfInstanceId field value if set, zero value otherwise.
func (o *VnfParameter) GetVnfInstanceId() string {
	if o == nil || isNil(o.VnfInstanceId) {
		var ret string
		return ret
	}
	return *o.VnfInstanceId
}

// GetVnfInstanceIdOk returns a tuple with the VnfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnfParameter) GetVnfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.VnfInstanceId) {
    return nil, false
	}
	return o.VnfInstanceId, true
}

// HasVnfInstanceId returns a boolean if a field has been set.
func (o *VnfParameter) HasVnfInstanceId() bool {
	if o != nil && !isNil(o.VnfInstanceId) {
		return true
	}

	return false
}

// SetVnfInstanceId gets a reference to the given string and assigns it to the VnfInstanceId field.
func (o *VnfParameter) SetVnfInstanceId(v string) {
	o.VnfInstanceId = &v
}

// GetVnfdId returns the VnfdId field value if set, zero value otherwise.
func (o *VnfParameter) GetVnfdId() string {
	if o == nil || isNil(o.VnfdId) {
		var ret string
		return ret
	}
	return *o.VnfdId
}

// GetVnfdIdOk returns a tuple with the VnfdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnfParameter) GetVnfdIdOk() (*string, bool) {
	if o == nil || isNil(o.VnfdId) {
    return nil, false
	}
	return o.VnfdId, true
}

// HasVnfdId returns a boolean if a field has been set.
func (o *VnfParameter) HasVnfdId() bool {
	if o != nil && !isNil(o.VnfdId) {
		return true
	}

	return false
}

// SetVnfdId gets a reference to the given string and assigns it to the VnfdId field.
func (o *VnfParameter) SetVnfdId(v string) {
	o.VnfdId = &v
}

// GetFlavourId returns the FlavourId field value if set, zero value otherwise.
func (o *VnfParameter) GetFlavourId() string {
	if o == nil || isNil(o.FlavourId) {
		var ret string
		return ret
	}
	return *o.FlavourId
}

// GetFlavourIdOk returns a tuple with the FlavourId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnfParameter) GetFlavourIdOk() (*string, bool) {
	if o == nil || isNil(o.FlavourId) {
    return nil, false
	}
	return o.FlavourId, true
}

// HasFlavourId returns a boolean if a field has been set.
func (o *VnfParameter) HasFlavourId() bool {
	if o != nil && !isNil(o.FlavourId) {
		return true
	}

	return false
}

// SetFlavourId gets a reference to the given string and assigns it to the FlavourId field.
func (o *VnfParameter) SetFlavourId(v string) {
	o.FlavourId = &v
}

// GetAutoScalable returns the AutoScalable field value if set, zero value otherwise.
func (o *VnfParameter) GetAutoScalable() bool {
	if o == nil || isNil(o.AutoScalable) {
		var ret bool
		return ret
	}
	return *o.AutoScalable
}

// GetAutoScalableOk returns a tuple with the AutoScalable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnfParameter) GetAutoScalableOk() (*bool, bool) {
	if o == nil || isNil(o.AutoScalable) {
    return nil, false
	}
	return o.AutoScalable, true
}

// HasAutoScalable returns a boolean if a field has been set.
func (o *VnfParameter) HasAutoScalable() bool {
	if o != nil && !isNil(o.AutoScalable) {
		return true
	}

	return false
}

// SetAutoScalable gets a reference to the given bool and assigns it to the AutoScalable field.
func (o *VnfParameter) SetAutoScalable(v bool) {
	o.AutoScalable = &v
}

func (o VnfParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VnfInstanceId) {
		toSerialize["vnfInstanceId"] = o.VnfInstanceId
	}
	if !isNil(o.VnfdId) {
		toSerialize["vnfdId"] = o.VnfdId
	}
	if !isNil(o.FlavourId) {
		toSerialize["flavourId"] = o.FlavourId
	}
	if !isNil(o.AutoScalable) {
		toSerialize["autoScalable"] = o.AutoScalable
	}
	return json.Marshal(toSerialize)
}

type NullableVnfParameter struct {
	value *VnfParameter
	isSet bool
}

func (v NullableVnfParameter) Get() *VnfParameter {
	return v.value
}

func (v *NullableVnfParameter) Set(val *VnfParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVnfParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVnfParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnfParameter(val *VnfParameter) *NullableVnfParameter {
	return &NullableVnfParameter{value: val, isSet: true}
}

func (v NullableVnfParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnfParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


