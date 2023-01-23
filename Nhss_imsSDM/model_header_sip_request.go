/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// HeaderSipRequest Contains a header (and optionally value of the header) in the SIP request
type HeaderSipRequest struct {
	Header string `json:"header"`
	Content *string `json:"content,omitempty"`
}

// NewHeaderSipRequest instantiates a new HeaderSipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderSipRequest(header string) *HeaderSipRequest {
	this := HeaderSipRequest{}
	this.Header = header
	return &this
}

// NewHeaderSipRequestWithDefaults instantiates a new HeaderSipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderSipRequestWithDefaults() *HeaderSipRequest {
	this := HeaderSipRequest{}
	return &this
}

// GetHeader returns the Header field value
func (o *HeaderSipRequest) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *HeaderSipRequest) GetHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *HeaderSipRequest) SetHeader(v string) {
	o.Header = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *HeaderSipRequest) GetContent() string {
	if o == nil || isNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderSipRequest) GetContentOk() (*string, bool) {
	if o == nil || isNil(o.Content) {
    return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *HeaderSipRequest) HasContent() bool {
	if o != nil && !isNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *HeaderSipRequest) SetContent(v string) {
	o.Content = &v
}

func (o HeaderSipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["header"] = o.Header
	}
	if !isNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableHeaderSipRequest struct {
	value *HeaderSipRequest
	isSet bool
}

func (v NullableHeaderSipRequest) Get() *HeaderSipRequest {
	return v.value
}

func (v *NullableHeaderSipRequest) Set(val *HeaderSipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderSipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderSipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderSipRequest(val *HeaderSipRequest) *NullableHeaderSipRequest {
	return &NullableHeaderSipRequest{value: val, isSet: true}
}

func (v NullableHeaderSipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderSipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


