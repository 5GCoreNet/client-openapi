/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// HttpPayload Contains the encoding of JSON payload in the API request / response
type HttpPayload struct {
	IePath string `json:"iePath"`
	IeValueLocation IeLocation `json:"ieValueLocation"`
	Value map[string]interface{} `json:"value"`
}

// NewHttpPayload instantiates a new HttpPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpPayload(iePath string, ieValueLocation IeLocation, value map[string]interface{}) *HttpPayload {
	this := HttpPayload{}
	this.IePath = iePath
	this.IeValueLocation = ieValueLocation
	this.Value = value
	return &this
}

// NewHttpPayloadWithDefaults instantiates a new HttpPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpPayloadWithDefaults() *HttpPayload {
	this := HttpPayload{}
	return &this
}

// GetIePath returns the IePath field value
func (o *HttpPayload) GetIePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IePath
}

// GetIePathOk returns a tuple with the IePath field value
// and a boolean to check if the value has been set.
func (o *HttpPayload) GetIePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IePath, true
}

// SetIePath sets field value
func (o *HttpPayload) SetIePath(v string) {
	o.IePath = v
}

// GetIeValueLocation returns the IeValueLocation field value
func (o *HttpPayload) GetIeValueLocation() IeLocation {
	if o == nil {
		var ret IeLocation
		return ret
	}

	return o.IeValueLocation
}

// GetIeValueLocationOk returns a tuple with the IeValueLocation field value
// and a boolean to check if the value has been set.
func (o *HttpPayload) GetIeValueLocationOk() (*IeLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IeValueLocation, true
}

// SetIeValueLocation sets field value
func (o *HttpPayload) SetIeValueLocation(v IeLocation) {
	o.IeValueLocation = v
}

// GetValue returns the Value field value
func (o *HttpPayload) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HttpPayload) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *HttpPayload) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o HttpPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iePath"] = o.IePath
	}
	if true {
		toSerialize["ieValueLocation"] = o.IeValueLocation
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHttpPayload struct {
	value *HttpPayload
	isSet bool
}

func (v NullableHttpPayload) Get() *HttpPayload {
	return v.value
}

func (v *NullableHttpPayload) Set(val *HttpPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpPayload(val *HttpPayload) *NullableHttpPayload {
	return &NullableHttpPayload{value: val, isSet: true}
}

func (v NullableHttpPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


