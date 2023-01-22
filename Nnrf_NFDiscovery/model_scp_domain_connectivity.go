/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// ScpDomainConnectivity SCP Domain Connectivity Information
type ScpDomainConnectivity struct {
	ConnectedScpDomainList []string `json:"connectedScpDomainList"`
}

// NewScpDomainConnectivity instantiates a new ScpDomainConnectivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpDomainConnectivity(connectedScpDomainList []string) *ScpDomainConnectivity {
	this := ScpDomainConnectivity{}
	this.ConnectedScpDomainList = connectedScpDomainList
	return &this
}

// NewScpDomainConnectivityWithDefaults instantiates a new ScpDomainConnectivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpDomainConnectivityWithDefaults() *ScpDomainConnectivity {
	this := ScpDomainConnectivity{}
	return &this
}

// GetConnectedScpDomainList returns the ConnectedScpDomainList field value
func (o *ScpDomainConnectivity) GetConnectedScpDomainList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConnectedScpDomainList
}

// GetConnectedScpDomainListOk returns a tuple with the ConnectedScpDomainList field value
// and a boolean to check if the value has been set.
func (o *ScpDomainConnectivity) GetConnectedScpDomainListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ConnectedScpDomainList, true
}

// SetConnectedScpDomainList sets field value
func (o *ScpDomainConnectivity) SetConnectedScpDomainList(v []string) {
	o.ConnectedScpDomainList = v
}

func (o ScpDomainConnectivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectedScpDomainList"] = o.ConnectedScpDomainList
	}
	return json.Marshal(toSerialize)
}

type NullableScpDomainConnectivity struct {
	value *ScpDomainConnectivity
	isSet bool
}

func (v NullableScpDomainConnectivity) Get() *ScpDomainConnectivity {
	return v.value
}

func (v *NullableScpDomainConnectivity) Set(val *ScpDomainConnectivity) {
	v.value = val
	v.isSet = true
}

func (v NullableScpDomainConnectivity) IsSet() bool {
	return v.isSet
}

func (v *NullableScpDomainConnectivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpDomainConnectivity(val *ScpDomainConnectivity) *NullableScpDomainConnectivity {
	return &NullableScpDomainConnectivity{value: val, isSet: true}
}

func (v NullableScpDomainConnectivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpDomainConnectivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


