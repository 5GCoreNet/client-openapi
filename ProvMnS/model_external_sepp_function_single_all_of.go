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

// ExternalSeppFunctionSingleAllOf struct for ExternalSeppFunctionSingleAllOf
type ExternalSeppFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewExternalSeppFunctionSingleAllOf instantiates a new ExternalSeppFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSeppFunctionSingleAllOf() *ExternalSeppFunctionSingleAllOf {
	this := ExternalSeppFunctionSingleAllOf{}
	return &this
}

// NewExternalSeppFunctionSingleAllOfWithDefaults instantiates a new ExternalSeppFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSeppFunctionSingleAllOfWithDefaults() *ExternalSeppFunctionSingleAllOf {
	this := ExternalSeppFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalSeppFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSeppFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalSeppFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalSeppFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o ExternalSeppFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableExternalSeppFunctionSingleAllOf struct {
	value *ExternalSeppFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalSeppFunctionSingleAllOf) Get() *ExternalSeppFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalSeppFunctionSingleAllOf) Set(val *ExternalSeppFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSeppFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSeppFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSeppFunctionSingleAllOf(val *ExternalSeppFunctionSingleAllOf) *NullableExternalSeppFunctionSingleAllOf {
	return &NullableExternalSeppFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalSeppFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSeppFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


