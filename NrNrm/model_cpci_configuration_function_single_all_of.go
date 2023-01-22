/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// CPCIConfigurationFunctionSingleAllOf struct for CPCIConfigurationFunctionSingleAllOf
type CPCIConfigurationFunctionSingleAllOf struct {
	Attributes *CPCIConfigurationFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewCPCIConfigurationFunctionSingleAllOf instantiates a new CPCIConfigurationFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCPCIConfigurationFunctionSingleAllOf() *CPCIConfigurationFunctionSingleAllOf {
	this := CPCIConfigurationFunctionSingleAllOf{}
	return &this
}

// NewCPCIConfigurationFunctionSingleAllOfWithDefaults instantiates a new CPCIConfigurationFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCPCIConfigurationFunctionSingleAllOfWithDefaults() *CPCIConfigurationFunctionSingleAllOf {
	this := CPCIConfigurationFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CPCIConfigurationFunctionSingleAllOf) GetAttributes() CPCIConfigurationFunctionSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret CPCIConfigurationFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPCIConfigurationFunctionSingleAllOf) GetAttributesOk() (*CPCIConfigurationFunctionSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CPCIConfigurationFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CPCIConfigurationFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *CPCIConfigurationFunctionSingleAllOf) SetAttributes(v CPCIConfigurationFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o CPCIConfigurationFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCPCIConfigurationFunctionSingleAllOf struct {
	value *CPCIConfigurationFunctionSingleAllOf
	isSet bool
}

func (v NullableCPCIConfigurationFunctionSingleAllOf) Get() *CPCIConfigurationFunctionSingleAllOf {
	return v.value
}

func (v *NullableCPCIConfigurationFunctionSingleAllOf) Set(val *CPCIConfigurationFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCPCIConfigurationFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCPCIConfigurationFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCPCIConfigurationFunctionSingleAllOf(val *CPCIConfigurationFunctionSingleAllOf) *NullableCPCIConfigurationFunctionSingleAllOf {
	return &NullableCPCIConfigurationFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableCPCIConfigurationFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCPCIConfigurationFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


