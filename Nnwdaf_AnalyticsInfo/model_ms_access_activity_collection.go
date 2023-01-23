/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// MSAccessActivityCollection Contains Media Streaming access activity collected for an UE Application via AF.
type MSAccessActivityCollection struct {
	MsAccActs []MediaStreamingAccessRecord `json:"msAccActs"`
}

// NewMSAccessActivityCollection instantiates a new MSAccessActivityCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSAccessActivityCollection(msAccActs []MediaStreamingAccessRecord) *MSAccessActivityCollection {
	this := MSAccessActivityCollection{}
	this.MsAccActs = msAccActs
	return &this
}

// NewMSAccessActivityCollectionWithDefaults instantiates a new MSAccessActivityCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSAccessActivityCollectionWithDefaults() *MSAccessActivityCollection {
	this := MSAccessActivityCollection{}
	return &this
}

// GetMsAccActs returns the MsAccActs field value
func (o *MSAccessActivityCollection) GetMsAccActs() []MediaStreamingAccessRecord {
	if o == nil {
		var ret []MediaStreamingAccessRecord
		return ret
	}

	return o.MsAccActs
}

// GetMsAccActsOk returns a tuple with the MsAccActs field value
// and a boolean to check if the value has been set.
func (o *MSAccessActivityCollection) GetMsAccActsOk() ([]MediaStreamingAccessRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.MsAccActs, true
}

// SetMsAccActs sets field value
func (o *MSAccessActivityCollection) SetMsAccActs(v []MediaStreamingAccessRecord) {
	o.MsAccActs = v
}

func (o MSAccessActivityCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["msAccActs"] = o.MsAccActs
	}
	return json.Marshal(toSerialize)
}

type NullableMSAccessActivityCollection struct {
	value *MSAccessActivityCollection
	isSet bool
}

func (v NullableMSAccessActivityCollection) Get() *MSAccessActivityCollection {
	return v.value
}

func (v *NullableMSAccessActivityCollection) Set(val *MSAccessActivityCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMSAccessActivityCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMSAccessActivityCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSAccessActivityCollection(val *MSAccessActivityCollection) *NullableMSAccessActivityCollection {
	return &NullableMSAccessActivityCollection{value: val, isSet: true}
}

func (v NullableMSAccessActivityCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSAccessActivityCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


