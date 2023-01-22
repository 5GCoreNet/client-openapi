/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
)

// ProSeCapability Indicate the supported ProSe Capability by the PCF.
type ProSeCapability struct {
	ProseDirectDiscovey *bool `json:"proseDirectDiscovey,omitempty"`
	ProseDirectCommunication *bool `json:"proseDirectCommunication,omitempty"`
	ProseL2UetoNetworkRelay *bool `json:"proseL2UetoNetworkRelay,omitempty"`
	ProseL3UetoNetworkRelay *bool `json:"proseL3UetoNetworkRelay,omitempty"`
	ProseL2RemoteUe *bool `json:"proseL2RemoteUe,omitempty"`
	ProseL3RemoteUe *bool `json:"proseL3RemoteUe,omitempty"`
}

// NewProSeCapability instantiates a new ProSeCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProSeCapability() *ProSeCapability {
	this := ProSeCapability{}
	var proseDirectDiscovey bool = false
	this.ProseDirectDiscovey = &proseDirectDiscovey
	var proseDirectCommunication bool = false
	this.ProseDirectCommunication = &proseDirectCommunication
	var proseL2UetoNetworkRelay bool = false
	this.ProseL2UetoNetworkRelay = &proseL2UetoNetworkRelay
	var proseL3UetoNetworkRelay bool = false
	this.ProseL3UetoNetworkRelay = &proseL3UetoNetworkRelay
	var proseL2RemoteUe bool = false
	this.ProseL2RemoteUe = &proseL2RemoteUe
	var proseL3RemoteUe bool = false
	this.ProseL3RemoteUe = &proseL3RemoteUe
	return &this
}

// NewProSeCapabilityWithDefaults instantiates a new ProSeCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProSeCapabilityWithDefaults() *ProSeCapability {
	this := ProSeCapability{}
	var proseDirectDiscovey bool = false
	this.ProseDirectDiscovey = &proseDirectDiscovey
	var proseDirectCommunication bool = false
	this.ProseDirectCommunication = &proseDirectCommunication
	var proseL2UetoNetworkRelay bool = false
	this.ProseL2UetoNetworkRelay = &proseL2UetoNetworkRelay
	var proseL3UetoNetworkRelay bool = false
	this.ProseL3UetoNetworkRelay = &proseL3UetoNetworkRelay
	var proseL2RemoteUe bool = false
	this.ProseL2RemoteUe = &proseL2RemoteUe
	var proseL3RemoteUe bool = false
	this.ProseL3RemoteUe = &proseL3RemoteUe
	return &this
}

// GetProseDirectDiscovey returns the ProseDirectDiscovey field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseDirectDiscovey() bool {
	if o == nil || isNil(o.ProseDirectDiscovey) {
		var ret bool
		return ret
	}
	return *o.ProseDirectDiscovey
}

// GetProseDirectDiscoveyOk returns a tuple with the ProseDirectDiscovey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseDirectDiscoveyOk() (*bool, bool) {
	if o == nil || isNil(o.ProseDirectDiscovey) {
    return nil, false
	}
	return o.ProseDirectDiscovey, true
}

// HasProseDirectDiscovey returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseDirectDiscovey() bool {
	if o != nil && !isNil(o.ProseDirectDiscovey) {
		return true
	}

	return false
}

// SetProseDirectDiscovey gets a reference to the given bool and assigns it to the ProseDirectDiscovey field.
func (o *ProSeCapability) SetProseDirectDiscovey(v bool) {
	o.ProseDirectDiscovey = &v
}

// GetProseDirectCommunication returns the ProseDirectCommunication field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseDirectCommunication() bool {
	if o == nil || isNil(o.ProseDirectCommunication) {
		var ret bool
		return ret
	}
	return *o.ProseDirectCommunication
}

// GetProseDirectCommunicationOk returns a tuple with the ProseDirectCommunication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseDirectCommunicationOk() (*bool, bool) {
	if o == nil || isNil(o.ProseDirectCommunication) {
    return nil, false
	}
	return o.ProseDirectCommunication, true
}

// HasProseDirectCommunication returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseDirectCommunication() bool {
	if o != nil && !isNil(o.ProseDirectCommunication) {
		return true
	}

	return false
}

// SetProseDirectCommunication gets a reference to the given bool and assigns it to the ProseDirectCommunication field.
func (o *ProSeCapability) SetProseDirectCommunication(v bool) {
	o.ProseDirectCommunication = &v
}

// GetProseL2UetoNetworkRelay returns the ProseL2UetoNetworkRelay field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseL2UetoNetworkRelay() bool {
	if o == nil || isNil(o.ProseL2UetoNetworkRelay) {
		var ret bool
		return ret
	}
	return *o.ProseL2UetoNetworkRelay
}

// GetProseL2UetoNetworkRelayOk returns a tuple with the ProseL2UetoNetworkRelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseL2UetoNetworkRelayOk() (*bool, bool) {
	if o == nil || isNil(o.ProseL2UetoNetworkRelay) {
    return nil, false
	}
	return o.ProseL2UetoNetworkRelay, true
}

// HasProseL2UetoNetworkRelay returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseL2UetoNetworkRelay() bool {
	if o != nil && !isNil(o.ProseL2UetoNetworkRelay) {
		return true
	}

	return false
}

// SetProseL2UetoNetworkRelay gets a reference to the given bool and assigns it to the ProseL2UetoNetworkRelay field.
func (o *ProSeCapability) SetProseL2UetoNetworkRelay(v bool) {
	o.ProseL2UetoNetworkRelay = &v
}

// GetProseL3UetoNetworkRelay returns the ProseL3UetoNetworkRelay field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseL3UetoNetworkRelay() bool {
	if o == nil || isNil(o.ProseL3UetoNetworkRelay) {
		var ret bool
		return ret
	}
	return *o.ProseL3UetoNetworkRelay
}

// GetProseL3UetoNetworkRelayOk returns a tuple with the ProseL3UetoNetworkRelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseL3UetoNetworkRelayOk() (*bool, bool) {
	if o == nil || isNil(o.ProseL3UetoNetworkRelay) {
    return nil, false
	}
	return o.ProseL3UetoNetworkRelay, true
}

// HasProseL3UetoNetworkRelay returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseL3UetoNetworkRelay() bool {
	if o != nil && !isNil(o.ProseL3UetoNetworkRelay) {
		return true
	}

	return false
}

// SetProseL3UetoNetworkRelay gets a reference to the given bool and assigns it to the ProseL3UetoNetworkRelay field.
func (o *ProSeCapability) SetProseL3UetoNetworkRelay(v bool) {
	o.ProseL3UetoNetworkRelay = &v
}

// GetProseL2RemoteUe returns the ProseL2RemoteUe field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseL2RemoteUe() bool {
	if o == nil || isNil(o.ProseL2RemoteUe) {
		var ret bool
		return ret
	}
	return *o.ProseL2RemoteUe
}

// GetProseL2RemoteUeOk returns a tuple with the ProseL2RemoteUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseL2RemoteUeOk() (*bool, bool) {
	if o == nil || isNil(o.ProseL2RemoteUe) {
    return nil, false
	}
	return o.ProseL2RemoteUe, true
}

// HasProseL2RemoteUe returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseL2RemoteUe() bool {
	if o != nil && !isNil(o.ProseL2RemoteUe) {
		return true
	}

	return false
}

// SetProseL2RemoteUe gets a reference to the given bool and assigns it to the ProseL2RemoteUe field.
func (o *ProSeCapability) SetProseL2RemoteUe(v bool) {
	o.ProseL2RemoteUe = &v
}

// GetProseL3RemoteUe returns the ProseL3RemoteUe field value if set, zero value otherwise.
func (o *ProSeCapability) GetProseL3RemoteUe() bool {
	if o == nil || isNil(o.ProseL3RemoteUe) {
		var ret bool
		return ret
	}
	return *o.ProseL3RemoteUe
}

// GetProseL3RemoteUeOk returns a tuple with the ProseL3RemoteUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeCapability) GetProseL3RemoteUeOk() (*bool, bool) {
	if o == nil || isNil(o.ProseL3RemoteUe) {
    return nil, false
	}
	return o.ProseL3RemoteUe, true
}

// HasProseL3RemoteUe returns a boolean if a field has been set.
func (o *ProSeCapability) HasProseL3RemoteUe() bool {
	if o != nil && !isNil(o.ProseL3RemoteUe) {
		return true
	}

	return false
}

// SetProseL3RemoteUe gets a reference to the given bool and assigns it to the ProseL3RemoteUe field.
func (o *ProSeCapability) SetProseL3RemoteUe(v bool) {
	o.ProseL3RemoteUe = &v
}

func (o ProSeCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProseDirectDiscovey) {
		toSerialize["proseDirectDiscovey"] = o.ProseDirectDiscovey
	}
	if !isNil(o.ProseDirectCommunication) {
		toSerialize["proseDirectCommunication"] = o.ProseDirectCommunication
	}
	if !isNil(o.ProseL2UetoNetworkRelay) {
		toSerialize["proseL2UetoNetworkRelay"] = o.ProseL2UetoNetworkRelay
	}
	if !isNil(o.ProseL3UetoNetworkRelay) {
		toSerialize["proseL3UetoNetworkRelay"] = o.ProseL3UetoNetworkRelay
	}
	if !isNil(o.ProseL2RemoteUe) {
		toSerialize["proseL2RemoteUe"] = o.ProseL2RemoteUe
	}
	if !isNil(o.ProseL3RemoteUe) {
		toSerialize["proseL3RemoteUe"] = o.ProseL3RemoteUe
	}
	return json.Marshal(toSerialize)
}

type NullableProSeCapability struct {
	value *ProSeCapability
	isSet bool
}

func (v NullableProSeCapability) Get() *ProSeCapability {
	return v.value
}

func (v *NullableProSeCapability) Set(val *ProSeCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeCapability(val *ProSeCapability) *NullableProSeCapability {
	return &NullableProSeCapability{value: val, isSet: true}
}

func (v NullableProSeCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


