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

// ScpDomainRoutingInformation SCP Domain Routing Information
type ScpDomainRoutingInformation struct {
	// This IE shall contain a map of SCP domain interconnection information, where the key of the map is a SCP domain. The value of each entry shall be the interconnectivity information of the the SCP domain indicated by the key. An empty map indicates that there is no SCP domain currently registered in the NRF. 
	ScpDomainList map[string]ScpDomainConnectivity `json:"scpDomainList"`
}

// NewScpDomainRoutingInformation instantiates a new ScpDomainRoutingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpDomainRoutingInformation(scpDomainList map[string]ScpDomainConnectivity) *ScpDomainRoutingInformation {
	this := ScpDomainRoutingInformation{}
	this.ScpDomainList = scpDomainList
	return &this
}

// NewScpDomainRoutingInformationWithDefaults instantiates a new ScpDomainRoutingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpDomainRoutingInformationWithDefaults() *ScpDomainRoutingInformation {
	this := ScpDomainRoutingInformation{}
	return &this
}

// GetScpDomainList returns the ScpDomainList field value
func (o *ScpDomainRoutingInformation) GetScpDomainList() map[string]ScpDomainConnectivity {
	if o == nil {
		var ret map[string]ScpDomainConnectivity
		return ret
	}

	return o.ScpDomainList
}

// GetScpDomainListOk returns a tuple with the ScpDomainList field value
// and a boolean to check if the value has been set.
func (o *ScpDomainRoutingInformation) GetScpDomainListOk() (*map[string]ScpDomainConnectivity, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScpDomainList, true
}

// SetScpDomainList sets field value
func (o *ScpDomainRoutingInformation) SetScpDomainList(v map[string]ScpDomainConnectivity) {
	o.ScpDomainList = v
}

func (o ScpDomainRoutingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scpDomainList"] = o.ScpDomainList
	}
	return json.Marshal(toSerialize)
}

type NullableScpDomainRoutingInformation struct {
	value *ScpDomainRoutingInformation
	isSet bool
}

func (v NullableScpDomainRoutingInformation) Get() *ScpDomainRoutingInformation {
	return v.value
}

func (v *NullableScpDomainRoutingInformation) Set(val *ScpDomainRoutingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableScpDomainRoutingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableScpDomainRoutingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpDomainRoutingInformation(val *ScpDomainRoutingInformation) *NullableScpDomainRoutingInformation {
	return &NullableScpDomainRoutingInformation{value: val, isSet: true}
}

func (v NullableScpDomainRoutingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpDomainRoutingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


