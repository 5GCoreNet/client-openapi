/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// TscStreamAvailability TSC stream availability information includes the stream specification and list of traffic  specifications. This response shall include stream specification matching one of the query  parameters provided in the request. 
type TscStreamAvailability struct {
	StreamSpec StreamSpecification `json:"streamSpec"`
	TrafficSpecs []TrafficSpecification `json:"trafficSpecs"`
}

// NewTscStreamAvailability instantiates a new TscStreamAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTscStreamAvailability(streamSpec StreamSpecification, trafficSpecs []TrafficSpecification) *TscStreamAvailability {
	this := TscStreamAvailability{}
	this.StreamSpec = streamSpec
	this.TrafficSpecs = trafficSpecs
	return &this
}

// NewTscStreamAvailabilityWithDefaults instantiates a new TscStreamAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTscStreamAvailabilityWithDefaults() *TscStreamAvailability {
	this := TscStreamAvailability{}
	return &this
}

// GetStreamSpec returns the StreamSpec field value
func (o *TscStreamAvailability) GetStreamSpec() StreamSpecification {
	if o == nil {
		var ret StreamSpecification
		return ret
	}

	return o.StreamSpec
}

// GetStreamSpecOk returns a tuple with the StreamSpec field value
// and a boolean to check if the value has been set.
func (o *TscStreamAvailability) GetStreamSpecOk() (*StreamSpecification, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StreamSpec, true
}

// SetStreamSpec sets field value
func (o *TscStreamAvailability) SetStreamSpec(v StreamSpecification) {
	o.StreamSpec = v
}

// GetTrafficSpecs returns the TrafficSpecs field value
func (o *TscStreamAvailability) GetTrafficSpecs() []TrafficSpecification {
	if o == nil {
		var ret []TrafficSpecification
		return ret
	}

	return o.TrafficSpecs
}

// GetTrafficSpecsOk returns a tuple with the TrafficSpecs field value
// and a boolean to check if the value has been set.
func (o *TscStreamAvailability) GetTrafficSpecsOk() ([]TrafficSpecification, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficSpecs, true
}

// SetTrafficSpecs sets field value
func (o *TscStreamAvailability) SetTrafficSpecs(v []TrafficSpecification) {
	o.TrafficSpecs = v
}

func (o TscStreamAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["streamSpec"] = o.StreamSpec
	}
	if true {
		toSerialize["trafficSpecs"] = o.TrafficSpecs
	}
	return json.Marshal(toSerialize)
}

type NullableTscStreamAvailability struct {
	value *TscStreamAvailability
	isSet bool
}

func (v NullableTscStreamAvailability) Get() *TscStreamAvailability {
	return v.value
}

func (v *NullableTscStreamAvailability) Set(val *TscStreamAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableTscStreamAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableTscStreamAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscStreamAvailability(val *TscStreamAvailability) *NullableTscStreamAvailability {
	return &NullableTscStreamAvailability{value: val, isSet: true}
}

func (v NullableTscStreamAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscStreamAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


