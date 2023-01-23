/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// DRACHOptimizationFunctionSingleAllOfAttributes struct for DRACHOptimizationFunctionSingleAllOfAttributes
type DRACHOptimizationFunctionSingleAllOfAttributes struct {
	DrachOptimizationControl *bool `json:"drachOptimizationControl,omitempty"`
	UeAccProbilityDist *UeAccProbilityDist `json:"ueAccProbilityDist,omitempty"`
	UeAccDelayProbilityDist *UeAccDelayProbilityDist `json:"ueAccDelayProbilityDist,omitempty"`
}

// NewDRACHOptimizationFunctionSingleAllOfAttributes instantiates a new DRACHOptimizationFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDRACHOptimizationFunctionSingleAllOfAttributes() *DRACHOptimizationFunctionSingleAllOfAttributes {
	this := DRACHOptimizationFunctionSingleAllOfAttributes{}
	return &this
}

// NewDRACHOptimizationFunctionSingleAllOfAttributesWithDefaults instantiates a new DRACHOptimizationFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDRACHOptimizationFunctionSingleAllOfAttributesWithDefaults() *DRACHOptimizationFunctionSingleAllOfAttributes {
	this := DRACHOptimizationFunctionSingleAllOfAttributes{}
	return &this
}

// GetDrachOptimizationControl returns the DrachOptimizationControl field value if set, zero value otherwise.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetDrachOptimizationControl() bool {
	if o == nil || isNil(o.DrachOptimizationControl) {
		var ret bool
		return ret
	}
	return *o.DrachOptimizationControl
}

// GetDrachOptimizationControlOk returns a tuple with the DrachOptimizationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetDrachOptimizationControlOk() (*bool, bool) {
	if o == nil || isNil(o.DrachOptimizationControl) {
    return nil, false
	}
	return o.DrachOptimizationControl, true
}

// HasDrachOptimizationControl returns a boolean if a field has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) HasDrachOptimizationControl() bool {
	if o != nil && !isNil(o.DrachOptimizationControl) {
		return true
	}

	return false
}

// SetDrachOptimizationControl gets a reference to the given bool and assigns it to the DrachOptimizationControl field.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) SetDrachOptimizationControl(v bool) {
	o.DrachOptimizationControl = &v
}

// GetUeAccProbilityDist returns the UeAccProbilityDist field value if set, zero value otherwise.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetUeAccProbilityDist() UeAccProbilityDist {
	if o == nil || isNil(o.UeAccProbilityDist) {
		var ret UeAccProbilityDist
		return ret
	}
	return *o.UeAccProbilityDist
}

// GetUeAccProbilityDistOk returns a tuple with the UeAccProbilityDist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetUeAccProbilityDistOk() (*UeAccProbilityDist, bool) {
	if o == nil || isNil(o.UeAccProbilityDist) {
    return nil, false
	}
	return o.UeAccProbilityDist, true
}

// HasUeAccProbilityDist returns a boolean if a field has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) HasUeAccProbilityDist() bool {
	if o != nil && !isNil(o.UeAccProbilityDist) {
		return true
	}

	return false
}

// SetUeAccProbilityDist gets a reference to the given UeAccProbilityDist and assigns it to the UeAccProbilityDist field.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) SetUeAccProbilityDist(v UeAccProbilityDist) {
	o.UeAccProbilityDist = &v
}

// GetUeAccDelayProbilityDist returns the UeAccDelayProbilityDist field value if set, zero value otherwise.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetUeAccDelayProbilityDist() UeAccDelayProbilityDist {
	if o == nil || isNil(o.UeAccDelayProbilityDist) {
		var ret UeAccDelayProbilityDist
		return ret
	}
	return *o.UeAccDelayProbilityDist
}

// GetUeAccDelayProbilityDistOk returns a tuple with the UeAccDelayProbilityDist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) GetUeAccDelayProbilityDistOk() (*UeAccDelayProbilityDist, bool) {
	if o == nil || isNil(o.UeAccDelayProbilityDist) {
    return nil, false
	}
	return o.UeAccDelayProbilityDist, true
}

// HasUeAccDelayProbilityDist returns a boolean if a field has been set.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) HasUeAccDelayProbilityDist() bool {
	if o != nil && !isNil(o.UeAccDelayProbilityDist) {
		return true
	}

	return false
}

// SetUeAccDelayProbilityDist gets a reference to the given UeAccDelayProbilityDist and assigns it to the UeAccDelayProbilityDist field.
func (o *DRACHOptimizationFunctionSingleAllOfAttributes) SetUeAccDelayProbilityDist(v UeAccDelayProbilityDist) {
	o.UeAccDelayProbilityDist = &v
}

func (o DRACHOptimizationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DrachOptimizationControl) {
		toSerialize["drachOptimizationControl"] = o.DrachOptimizationControl
	}
	if !isNil(o.UeAccProbilityDist) {
		toSerialize["ueAccProbilityDist"] = o.UeAccProbilityDist
	}
	if !isNil(o.UeAccDelayProbilityDist) {
		toSerialize["ueAccDelayProbilityDist"] = o.UeAccDelayProbilityDist
	}
	return json.Marshal(toSerialize)
}

type NullableDRACHOptimizationFunctionSingleAllOfAttributes struct {
	value *DRACHOptimizationFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableDRACHOptimizationFunctionSingleAllOfAttributes) Get() *DRACHOptimizationFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableDRACHOptimizationFunctionSingleAllOfAttributes) Set(val *DRACHOptimizationFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDRACHOptimizationFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDRACHOptimizationFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDRACHOptimizationFunctionSingleAllOfAttributes(val *DRACHOptimizationFunctionSingleAllOfAttributes) *NullableDRACHOptimizationFunctionSingleAllOfAttributes {
	return &NullableDRACHOptimizationFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableDRACHOptimizationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDRACHOptimizationFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


