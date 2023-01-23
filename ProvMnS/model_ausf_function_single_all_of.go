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

// AusfFunctionSingleAllOf struct for AusfFunctionSingleAllOf
type AusfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewAusfFunctionSingleAllOf instantiates a new AusfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfFunctionSingleAllOf() *AusfFunctionSingleAllOf {
	this := AusfFunctionSingleAllOf{}
	return &this
}

// NewAusfFunctionSingleAllOfWithDefaults instantiates a new AusfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfFunctionSingleAllOfWithDefaults() *AusfFunctionSingleAllOf {
	this := AusfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *AusfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o AusfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAusfFunctionSingleAllOf struct {
	value *AusfFunctionSingleAllOf
	isSet bool
}

func (v NullableAusfFunctionSingleAllOf) Get() *AusfFunctionSingleAllOf {
	return v.value
}

func (v *NullableAusfFunctionSingleAllOf) Set(val *AusfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfFunctionSingleAllOf(val *AusfFunctionSingleAllOf) *NullableAusfFunctionSingleAllOf {
	return &NullableAusfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableAusfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


