/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
)

// MsConsumptionCollection Contains the Media Streaming Consumption information collected for an UE Application via AF. 
type MsConsumptionCollection struct {
	MsConsumps []string `json:"msConsumps"`
}

// NewMsConsumptionCollection instantiates a new MsConsumptionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsConsumptionCollection(msConsumps []string) *MsConsumptionCollection {
	this := MsConsumptionCollection{}
	this.MsConsumps = msConsumps
	return &this
}

// NewMsConsumptionCollectionWithDefaults instantiates a new MsConsumptionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsConsumptionCollectionWithDefaults() *MsConsumptionCollection {
	this := MsConsumptionCollection{}
	return &this
}

// GetMsConsumps returns the MsConsumps field value
func (o *MsConsumptionCollection) GetMsConsumps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MsConsumps
}

// GetMsConsumpsOk returns a tuple with the MsConsumps field value
// and a boolean to check if the value has been set.
func (o *MsConsumptionCollection) GetMsConsumpsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MsConsumps, true
}

// SetMsConsumps sets field value
func (o *MsConsumptionCollection) SetMsConsumps(v []string) {
	o.MsConsumps = v
}

func (o MsConsumptionCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["msConsumps"] = o.MsConsumps
	}
	return json.Marshal(toSerialize)
}

type NullableMsConsumptionCollection struct {
	value *MsConsumptionCollection
	isSet bool
}

func (v NullableMsConsumptionCollection) Get() *MsConsumptionCollection {
	return v.value
}

func (v *NullableMsConsumptionCollection) Set(val *MsConsumptionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMsConsumptionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMsConsumptionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsConsumptionCollection(val *MsConsumptionCollection) *NullableMsConsumptionCollection {
	return &NullableMsConsumptionCollection{value: val, isSet: true}
}

func (v NullableMsConsumptionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsConsumptionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


