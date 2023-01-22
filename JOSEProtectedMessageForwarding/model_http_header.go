/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// HttpHeader Contains the encoding of HTTP headers in the API request / response
type HttpHeader struct {
	Header string `json:"header"`
	Value EncodedHttpHeaderValue `json:"value"`
}

// NewHttpHeader instantiates a new HttpHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHeader(header string, value EncodedHttpHeaderValue) *HttpHeader {
	this := HttpHeader{}
	this.Header = header
	this.Value = value
	return &this
}

// NewHttpHeaderWithDefaults instantiates a new HttpHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeaderWithDefaults() *HttpHeader {
	this := HttpHeader{}
	return &this
}

// GetHeader returns the Header field value
func (o *HttpHeader) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *HttpHeader) GetHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *HttpHeader) SetHeader(v string) {
	o.Header = v
}

// GetValue returns the Value field value
func (o *HttpHeader) GetValue() EncodedHttpHeaderValue {
	if o == nil {
		var ret EncodedHttpHeaderValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HttpHeader) GetValueOk() (*EncodedHttpHeaderValue, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HttpHeader) SetValue(v EncodedHttpHeaderValue) {
	o.Value = v
}

func (o HttpHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["header"] = o.Header
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHttpHeader struct {
	value *HttpHeader
	isSet bool
}

func (v NullableHttpHeader) Get() *HttpHeader {
	return v.value
}

func (v *NullableHttpHeader) Set(val *HttpHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHeader(val *HttpHeader) *NullableHttpHeader {
	return &NullableHttpHeader{value: val, isSet: true}
}

func (v NullableHttpHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


