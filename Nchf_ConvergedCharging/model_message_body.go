/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// MessageBody struct for MessageBody
type MessageBody struct {
	ContentType string `json:"contentType"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	ContentLength int32 `json:"contentLength"`
	ContentDisposition *string `json:"contentDisposition,omitempty"`
	Originator *OriginatorPartyType `json:"originator,omitempty"`
}

// NewMessageBody instantiates a new MessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageBody(contentType string, contentLength int32) *MessageBody {
	this := MessageBody{}
	this.ContentType = contentType
	this.ContentLength = contentLength
	return &this
}

// NewMessageBodyWithDefaults instantiates a new MessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageBodyWithDefaults() *MessageBody {
	this := MessageBody{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *MessageBody) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *MessageBody) GetContentTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *MessageBody) SetContentType(v string) {
	o.ContentType = v
}

// GetContentLength returns the ContentLength field value
func (o *MessageBody) GetContentLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentLength
}

// GetContentLengthOk returns a tuple with the ContentLength field value
// and a boolean to check if the value has been set.
func (o *MessageBody) GetContentLengthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContentLength, true
}

// SetContentLength sets field value
func (o *MessageBody) SetContentLength(v int32) {
	o.ContentLength = v
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *MessageBody) GetContentDisposition() string {
	if o == nil || isNil(o.ContentDisposition) {
		var ret string
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBody) GetContentDispositionOk() (*string, bool) {
	if o == nil || isNil(o.ContentDisposition) {
    return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *MessageBody) HasContentDisposition() bool {
	if o != nil && !isNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given string and assigns it to the ContentDisposition field.
func (o *MessageBody) SetContentDisposition(v string) {
	o.ContentDisposition = &v
}

// GetOriginator returns the Originator field value if set, zero value otherwise.
func (o *MessageBody) GetOriginator() OriginatorPartyType {
	if o == nil || isNil(o.Originator) {
		var ret OriginatorPartyType
		return ret
	}
	return *o.Originator
}

// GetOriginatorOk returns a tuple with the Originator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageBody) GetOriginatorOk() (*OriginatorPartyType, bool) {
	if o == nil || isNil(o.Originator) {
    return nil, false
	}
	return o.Originator, true
}

// HasOriginator returns a boolean if a field has been set.
func (o *MessageBody) HasOriginator() bool {
	if o != nil && !isNil(o.Originator) {
		return true
	}

	return false
}

// SetOriginator gets a reference to the given OriginatorPartyType and assigns it to the Originator field.
func (o *MessageBody) SetOriginator(v OriginatorPartyType) {
	o.Originator = &v
}

func (o MessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentType"] = o.ContentType
	}
	if true {
		toSerialize["contentLength"] = o.ContentLength
	}
	if !isNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if !isNil(o.Originator) {
		toSerialize["originator"] = o.Originator
	}
	return json.Marshal(toSerialize)
}

type NullableMessageBody struct {
	value *MessageBody
	isSet bool
}

func (v NullableMessageBody) Get() *MessageBody {
	return v.value
}

func (v *NullableMessageBody) Set(val *MessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageBody(val *MessageBody) *NullableMessageBody {
	return &NullableMessageBody{value: val, isSet: true}
}

func (v NullableMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


