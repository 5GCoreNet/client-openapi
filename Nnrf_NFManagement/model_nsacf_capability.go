/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// NsacfCapability NSACF service capabilities (e.g. to monitor and control the number of registered UEs or established PDU sessions per network slice) 
type NsacfCapability struct {
	// Indicates the service capability of the NSACF to monitor and control the number of registered UEs per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportUeSAC *bool `json:"supportUeSAC,omitempty"`
	// Indicates the service capability of the NSACF to monitor and control the number of established PDU sessions per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportPduSAC *bool `json:"supportPduSAC,omitempty"`
}

// NewNsacfCapability instantiates a new NsacfCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfCapability() *NsacfCapability {
	this := NsacfCapability{}
	var supportUeSAC bool = false
	this.SupportUeSAC = &supportUeSAC
	var supportPduSAC bool = false
	this.SupportPduSAC = &supportPduSAC
	return &this
}

// NewNsacfCapabilityWithDefaults instantiates a new NsacfCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfCapabilityWithDefaults() *NsacfCapability {
	this := NsacfCapability{}
	var supportUeSAC bool = false
	this.SupportUeSAC = &supportUeSAC
	var supportPduSAC bool = false
	this.SupportPduSAC = &supportPduSAC
	return &this
}

// GetSupportUeSAC returns the SupportUeSAC field value if set, zero value otherwise.
func (o *NsacfCapability) GetSupportUeSAC() bool {
	if o == nil || isNil(o.SupportUeSAC) {
		var ret bool
		return ret
	}
	return *o.SupportUeSAC
}

// GetSupportUeSACOk returns a tuple with the SupportUeSAC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfCapability) GetSupportUeSACOk() (*bool, bool) {
	if o == nil || isNil(o.SupportUeSAC) {
    return nil, false
	}
	return o.SupportUeSAC, true
}

// HasSupportUeSAC returns a boolean if a field has been set.
func (o *NsacfCapability) HasSupportUeSAC() bool {
	if o != nil && !isNil(o.SupportUeSAC) {
		return true
	}

	return false
}

// SetSupportUeSAC gets a reference to the given bool and assigns it to the SupportUeSAC field.
func (o *NsacfCapability) SetSupportUeSAC(v bool) {
	o.SupportUeSAC = &v
}

// GetSupportPduSAC returns the SupportPduSAC field value if set, zero value otherwise.
func (o *NsacfCapability) GetSupportPduSAC() bool {
	if o == nil || isNil(o.SupportPduSAC) {
		var ret bool
		return ret
	}
	return *o.SupportPduSAC
}

// GetSupportPduSACOk returns a tuple with the SupportPduSAC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfCapability) GetSupportPduSACOk() (*bool, bool) {
	if o == nil || isNil(o.SupportPduSAC) {
    return nil, false
	}
	return o.SupportPduSAC, true
}

// HasSupportPduSAC returns a boolean if a field has been set.
func (o *NsacfCapability) HasSupportPduSAC() bool {
	if o != nil && !isNil(o.SupportPduSAC) {
		return true
	}

	return false
}

// SetSupportPduSAC gets a reference to the given bool and assigns it to the SupportPduSAC field.
func (o *NsacfCapability) SetSupportPduSAC(v bool) {
	o.SupportPduSAC = &v
}

func (o NsacfCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupportUeSAC) {
		toSerialize["supportUeSAC"] = o.SupportUeSAC
	}
	if !isNil(o.SupportPduSAC) {
		toSerialize["supportPduSAC"] = o.SupportPduSAC
	}
	return json.Marshal(toSerialize)
}

type NullableNsacfCapability struct {
	value *NsacfCapability
	isSet bool
}

func (v NullableNsacfCapability) Get() *NsacfCapability {
	return v.value
}

func (v *NullableNsacfCapability) Set(val *NsacfCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfCapability(val *NsacfCapability) *NullableNsacfCapability {
	return &NullableNsacfCapability{value: val, isSet: true}
}

func (v NullableNsacfCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


