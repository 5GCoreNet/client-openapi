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

// ManagementNodeSingleAllOf struct for ManagementNodeSingleAllOf
type ManagementNodeSingleAllOf struct {
	Attributes *ManagementNodeSingleAllOfAttributes `json:"attributes,omitempty"`
	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`
}

// NewManagementNodeSingleAllOf instantiates a new ManagementNodeSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementNodeSingleAllOf() *ManagementNodeSingleAllOf {
	this := ManagementNodeSingleAllOf{}
	return &this
}

// NewManagementNodeSingleAllOfWithDefaults instantiates a new ManagementNodeSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementNodeSingleAllOfWithDefaults() *ManagementNodeSingleAllOf {
	this := ManagementNodeSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOf) GetAttributes() ManagementNodeSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret ManagementNodeSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOf) GetAttributesOk() (*ManagementNodeSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagementNodeSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagementNodeSingleAllOf) SetAttributes(v ManagementNodeSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOf) GetMnsAgent() []MnsAgentSingle {
	if o == nil || isNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOf) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || isNil(o.MnsAgent) {
    return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOf) HasMnsAgent() bool {
	if o != nil && !isNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *ManagementNodeSingleAllOf) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

func (o ManagementNodeSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	return json.Marshal(toSerialize)
}

type NullableManagementNodeSingleAllOf struct {
	value *ManagementNodeSingleAllOf
	isSet bool
}

func (v NullableManagementNodeSingleAllOf) Get() *ManagementNodeSingleAllOf {
	return v.value
}

func (v *NullableManagementNodeSingleAllOf) Set(val *ManagementNodeSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementNodeSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementNodeSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementNodeSingleAllOf(val *ManagementNodeSingleAllOf) *NullableManagementNodeSingleAllOf {
	return &NullableManagementNodeSingleAllOf{value: val, isSet: true}
}

func (v NullableManagementNodeSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementNodeSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


