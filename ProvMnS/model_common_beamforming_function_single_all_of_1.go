/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// CommonBeamformingFunctionSingleAllOf1 struct for CommonBeamformingFunctionSingleAllOf1
type CommonBeamformingFunctionSingleAllOf1 struct {
	Beam []BeamSingle `json:"Beam,omitempty"`
}

// NewCommonBeamformingFunctionSingleAllOf1 instantiates a new CommonBeamformingFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBeamformingFunctionSingleAllOf1() *CommonBeamformingFunctionSingleAllOf1 {
	this := CommonBeamformingFunctionSingleAllOf1{}
	return &this
}

// NewCommonBeamformingFunctionSingleAllOf1WithDefaults instantiates a new CommonBeamformingFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBeamformingFunctionSingleAllOf1WithDefaults() *CommonBeamformingFunctionSingleAllOf1 {
	this := CommonBeamformingFunctionSingleAllOf1{}
	return &this
}

// GetBeam returns the Beam field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOf1) GetBeam() []BeamSingle {
	if o == nil || isNil(o.Beam) {
		var ret []BeamSingle
		return ret
	}
	return o.Beam
}

// GetBeamOk returns a tuple with the Beam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOf1) GetBeamOk() ([]BeamSingle, bool) {
	if o == nil || isNil(o.Beam) {
    return nil, false
	}
	return o.Beam, true
}

// HasBeam returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOf1) HasBeam() bool {
	if o != nil && !isNil(o.Beam) {
		return true
	}

	return false
}

// SetBeam gets a reference to the given []BeamSingle and assigns it to the Beam field.
func (o *CommonBeamformingFunctionSingleAllOf1) SetBeam(v []BeamSingle) {
	o.Beam = v
}

func (o CommonBeamformingFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Beam) {
		toSerialize["Beam"] = o.Beam
	}
	return json.Marshal(toSerialize)
}

type NullableCommonBeamformingFunctionSingleAllOf1 struct {
	value *CommonBeamformingFunctionSingleAllOf1
	isSet bool
}

func (v NullableCommonBeamformingFunctionSingleAllOf1) Get() *CommonBeamformingFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableCommonBeamformingFunctionSingleAllOf1) Set(val *CommonBeamformingFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBeamformingFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBeamformingFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBeamformingFunctionSingleAllOf1(val *CommonBeamformingFunctionSingleAllOf1) *NullableCommonBeamformingFunctionSingleAllOf1 {
	return &NullableCommonBeamformingFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableCommonBeamformingFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBeamformingFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


