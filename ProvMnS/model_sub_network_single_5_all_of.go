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

// SubNetworkSingle5AllOf struct for SubNetworkSingle5AllOf
type SubNetworkSingle5AllOf struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
	MLTrainingFunction []MLTrainingFunctionSingle `json:"MLTrainingFunction,omitempty"`
}

// NewSubNetworkSingle5AllOf instantiates a new SubNetworkSingle5AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle5AllOf() *SubNetworkSingle5AllOf {
	this := SubNetworkSingle5AllOf{}
	return &this
}

// NewSubNetworkSingle5AllOfWithDefaults instantiates a new SubNetworkSingle5AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle5AllOfWithDefaults() *SubNetworkSingle5AllOf {
	this := SubNetworkSingle5AllOf{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle5AllOf) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5AllOf) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle5AllOf) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle5AllOf) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingle5AllOf) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5AllOf) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingle5AllOf) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingle5AllOf) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetMLTrainingFunction returns the MLTrainingFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle5AllOf) GetMLTrainingFunction() []MLTrainingFunctionSingle {
	if o == nil || isNil(o.MLTrainingFunction) {
		var ret []MLTrainingFunctionSingle
		return ret
	}
	return o.MLTrainingFunction
}

// GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5AllOf) GetMLTrainingFunctionOk() ([]MLTrainingFunctionSingle, bool) {
	if o == nil || isNil(o.MLTrainingFunction) {
    return nil, false
	}
	return o.MLTrainingFunction, true
}

// HasMLTrainingFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle5AllOf) HasMLTrainingFunction() bool {
	if o != nil && !isNil(o.MLTrainingFunction) {
		return true
	}

	return false
}

// SetMLTrainingFunction gets a reference to the given []MLTrainingFunctionSingle and assigns it to the MLTrainingFunction field.
func (o *SubNetworkSingle5AllOf) SetMLTrainingFunction(v []MLTrainingFunctionSingle) {
	o.MLTrainingFunction = v
}

func (o SubNetworkSingle5AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !isNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	if !isNil(o.MLTrainingFunction) {
		toSerialize["MLTrainingFunction"] = o.MLTrainingFunction
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle5AllOf struct {
	value *SubNetworkSingle5AllOf
	isSet bool
}

func (v NullableSubNetworkSingle5AllOf) Get() *SubNetworkSingle5AllOf {
	return v.value
}

func (v *NullableSubNetworkSingle5AllOf) Set(val *SubNetworkSingle5AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle5AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle5AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle5AllOf(val *SubNetworkSingle5AllOf) *NullableSubNetworkSingle5AllOf {
	return &NullableSubNetworkSingle5AllOf{value: val, isSet: true}
}

func (v NullableSubNetworkSingle5AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle5AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


