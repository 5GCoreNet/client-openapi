/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_EventExposure

import (
	"encoding/json"
)

// MediaStreamingAccessRecordAllOfResponseMessage struct for MediaStreamingAccessRecordAllOfResponseMessage
type MediaStreamingAccessRecordAllOfResponseMessage struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ResponseCode int32 `json:"responseCode"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Size int32 `json:"size"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BodySize int32 `json:"bodySize"`
	ContentType *string `json:"contentType,omitempty"`
}

// NewMediaStreamingAccessRecordAllOfResponseMessage instantiates a new MediaStreamingAccessRecordAllOfResponseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaStreamingAccessRecordAllOfResponseMessage(responseCode int32, size int32, bodySize int32) *MediaStreamingAccessRecordAllOfResponseMessage {
	this := MediaStreamingAccessRecordAllOfResponseMessage{}
	this.ResponseCode = responseCode
	this.Size = size
	this.BodySize = bodySize
	return &this
}

// NewMediaStreamingAccessRecordAllOfResponseMessageWithDefaults instantiates a new MediaStreamingAccessRecordAllOfResponseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamingAccessRecordAllOfResponseMessageWithDefaults() *MediaStreamingAccessRecordAllOfResponseMessage {
	this := MediaStreamingAccessRecordAllOfResponseMessage{}
	return &this
}

// GetResponseCode returns the ResponseCode field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetResponseCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetResponseCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetResponseCode(v int32) {
	o.ResponseCode = v
}

// GetSize returns the Size field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetSize(v int32) {
	o.Size = v
}

// GetBodySize returns the BodySize field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetBodySize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BodySize
}

// GetBodySizeOk returns a tuple with the BodySize field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetBodySizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BodySize, true
}

// SetBodySize sets field value
func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetBodySize(v int32) {
	o.BodySize = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetContentType() string {
	if o == nil || isNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetContentTypeOk() (*string, bool) {
	if o == nil || isNil(o.ContentType) {
    return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) HasContentType() bool {
	if o != nil && !isNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetContentType(v string) {
	o.ContentType = &v
}

func (o MediaStreamingAccessRecordAllOfResponseMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["bodySize"] = o.BodySize
	}
	if !isNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableMediaStreamingAccessRecordAllOfResponseMessage struct {
	value *MediaStreamingAccessRecordAllOfResponseMessage
	isSet bool
}

func (v NullableMediaStreamingAccessRecordAllOfResponseMessage) Get() *MediaStreamingAccessRecordAllOfResponseMessage {
	return v.value
}

func (v *NullableMediaStreamingAccessRecordAllOfResponseMessage) Set(val *MediaStreamingAccessRecordAllOfResponseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamingAccessRecordAllOfResponseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamingAccessRecordAllOfResponseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamingAccessRecordAllOfResponseMessage(val *MediaStreamingAccessRecordAllOfResponseMessage) *NullableMediaStreamingAccessRecordAllOfResponseMessage {
	return &NullableMediaStreamingAccessRecordAllOfResponseMessage{value: val, isSet: true}
}

func (v NullableMediaStreamingAccessRecordAllOfResponseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamingAccessRecordAllOfResponseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

