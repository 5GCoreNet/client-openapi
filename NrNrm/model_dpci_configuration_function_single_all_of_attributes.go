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

// DPCIConfigurationFunctionSingleAllOfAttributes struct for DPCIConfigurationFunctionSingleAllOfAttributes
type DPCIConfigurationFunctionSingleAllOfAttributes struct {
	DPciConfigurationControl *bool `json:"dPciConfigurationControl,omitempty"`
	NRPciList *NRPciList `json:"nRPciList,omitempty"`
}

// NewDPCIConfigurationFunctionSingleAllOfAttributes instantiates a new DPCIConfigurationFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDPCIConfigurationFunctionSingleAllOfAttributes() *DPCIConfigurationFunctionSingleAllOfAttributes {
	this := DPCIConfigurationFunctionSingleAllOfAttributes{}
	return &this
}

// NewDPCIConfigurationFunctionSingleAllOfAttributesWithDefaults instantiates a new DPCIConfigurationFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDPCIConfigurationFunctionSingleAllOfAttributesWithDefaults() *DPCIConfigurationFunctionSingleAllOfAttributes {
	this := DPCIConfigurationFunctionSingleAllOfAttributes{}
	return &this
}

// GetDPciConfigurationControl returns the DPciConfigurationControl field value if set, zero value otherwise.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) GetDPciConfigurationControl() bool {
	if o == nil || isNil(o.DPciConfigurationControl) {
		var ret bool
		return ret
	}
	return *o.DPciConfigurationControl
}

// GetDPciConfigurationControlOk returns a tuple with the DPciConfigurationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) GetDPciConfigurationControlOk() (*bool, bool) {
	if o == nil || isNil(o.DPciConfigurationControl) {
    return nil, false
	}
	return o.DPciConfigurationControl, true
}

// HasDPciConfigurationControl returns a boolean if a field has been set.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) HasDPciConfigurationControl() bool {
	if o != nil && !isNil(o.DPciConfigurationControl) {
		return true
	}

	return false
}

// SetDPciConfigurationControl gets a reference to the given bool and assigns it to the DPciConfigurationControl field.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) SetDPciConfigurationControl(v bool) {
	o.DPciConfigurationControl = &v
}

// GetNRPciList returns the NRPciList field value if set, zero value otherwise.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) GetNRPciList() NRPciList {
	if o == nil || isNil(o.NRPciList) {
		var ret NRPciList
		return ret
	}
	return *o.NRPciList
}

// GetNRPciListOk returns a tuple with the NRPciList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) GetNRPciListOk() (*NRPciList, bool) {
	if o == nil || isNil(o.NRPciList) {
    return nil, false
	}
	return o.NRPciList, true
}

// HasNRPciList returns a boolean if a field has been set.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) HasNRPciList() bool {
	if o != nil && !isNil(o.NRPciList) {
		return true
	}

	return false
}

// SetNRPciList gets a reference to the given NRPciList and assigns it to the NRPciList field.
func (o *DPCIConfigurationFunctionSingleAllOfAttributes) SetNRPciList(v NRPciList) {
	o.NRPciList = &v
}

func (o DPCIConfigurationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DPciConfigurationControl) {
		toSerialize["dPciConfigurationControl"] = o.DPciConfigurationControl
	}
	if !isNil(o.NRPciList) {
		toSerialize["nRPciList"] = o.NRPciList
	}
	return json.Marshal(toSerialize)
}

type NullableDPCIConfigurationFunctionSingleAllOfAttributes struct {
	value *DPCIConfigurationFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableDPCIConfigurationFunctionSingleAllOfAttributes) Get() *DPCIConfigurationFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableDPCIConfigurationFunctionSingleAllOfAttributes) Set(val *DPCIConfigurationFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDPCIConfigurationFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDPCIConfigurationFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDPCIConfigurationFunctionSingleAllOfAttributes(val *DPCIConfigurationFunctionSingleAllOfAttributes) *NullableDPCIConfigurationFunctionSingleAllOfAttributes {
	return &NullableDPCIConfigurationFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableDPCIConfigurationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDPCIConfigurationFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


