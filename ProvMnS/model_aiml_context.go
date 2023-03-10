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

// AIMLContext struct for AIMLContext
type AIMLContext struct {
	InferenceEntityRef []string `json:"inferenceEntityRef,omitempty"`
	DataProviderRef []string `json:"dataProviderRef,omitempty"`
}

// NewAIMLContext instantiates a new AIMLContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIMLContext() *AIMLContext {
	this := AIMLContext{}
	return &this
}

// NewAIMLContextWithDefaults instantiates a new AIMLContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIMLContextWithDefaults() *AIMLContext {
	this := AIMLContext{}
	return &this
}

// GetInferenceEntityRef returns the InferenceEntityRef field value if set, zero value otherwise.
func (o *AIMLContext) GetInferenceEntityRef() []string {
	if o == nil || isNil(o.InferenceEntityRef) {
		var ret []string
		return ret
	}
	return o.InferenceEntityRef
}

// GetInferenceEntityRefOk returns a tuple with the InferenceEntityRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIMLContext) GetInferenceEntityRefOk() ([]string, bool) {
	if o == nil || isNil(o.InferenceEntityRef) {
    return nil, false
	}
	return o.InferenceEntityRef, true
}

// HasInferenceEntityRef returns a boolean if a field has been set.
func (o *AIMLContext) HasInferenceEntityRef() bool {
	if o != nil && !isNil(o.InferenceEntityRef) {
		return true
	}

	return false
}

// SetInferenceEntityRef gets a reference to the given []string and assigns it to the InferenceEntityRef field.
func (o *AIMLContext) SetInferenceEntityRef(v []string) {
	o.InferenceEntityRef = v
}

// GetDataProviderRef returns the DataProviderRef field value if set, zero value otherwise.
func (o *AIMLContext) GetDataProviderRef() []string {
	if o == nil || isNil(o.DataProviderRef) {
		var ret []string
		return ret
	}
	return o.DataProviderRef
}

// GetDataProviderRefOk returns a tuple with the DataProviderRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIMLContext) GetDataProviderRefOk() ([]string, bool) {
	if o == nil || isNil(o.DataProviderRef) {
    return nil, false
	}
	return o.DataProviderRef, true
}

// HasDataProviderRef returns a boolean if a field has been set.
func (o *AIMLContext) HasDataProviderRef() bool {
	if o != nil && !isNil(o.DataProviderRef) {
		return true
	}

	return false
}

// SetDataProviderRef gets a reference to the given []string and assigns it to the DataProviderRef field.
func (o *AIMLContext) SetDataProviderRef(v []string) {
	o.DataProviderRef = v
}

func (o AIMLContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InferenceEntityRef) {
		toSerialize["inferenceEntityRef"] = o.InferenceEntityRef
	}
	if !isNil(o.DataProviderRef) {
		toSerialize["dataProviderRef"] = o.DataProviderRef
	}
	return json.Marshal(toSerialize)
}

type NullableAIMLContext struct {
	value *AIMLContext
	isSet bool
}

func (v NullableAIMLContext) Get() *AIMLContext {
	return v.value
}

func (v *NullableAIMLContext) Set(val *AIMLContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAIMLContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAIMLContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIMLContext(val *AIMLContext) *NullableAIMLContext {
	return &NullableAIMLContext{value: val, isSet: true}
}

func (v NullableAIMLContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIMLContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


