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

// MnSOneOf struct for MnSOneOf
type MnSOneOf struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
}

// NewMnSOneOf instantiates a new MnSOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnSOneOf() *MnSOneOf {
	this := MnSOneOf{}
	return &this
}

// NewMnSOneOfWithDefaults instantiates a new MnSOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnSOneOfWithDefaults() *MnSOneOf {
	this := MnSOneOf{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *MnSOneOf) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnSOneOf) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *MnSOneOf) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *MnSOneOf) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

func (o MnSOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableMnSOneOf struct {
	value *MnSOneOf
	isSet bool
}

func (v NullableMnSOneOf) Get() *MnSOneOf {
	return v.value
}

func (v *NullableMnSOneOf) Set(val *MnSOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMnSOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMnSOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnSOneOf(val *MnSOneOf) *NullableMnSOneOf {
	return &NullableMnSOneOf{value: val, isSet: true}
}

func (v NullableMnSOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnSOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


