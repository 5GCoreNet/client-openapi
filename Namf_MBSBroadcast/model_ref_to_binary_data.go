/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_MBSBroadcast

import (
	"encoding/json"
)

// RefToBinaryData This parameter provides information about the referenced binary body data.
type RefToBinaryData struct {
	// This IE shall contain the value of the Content-ID header of the referenced binary body part. 
	ContentId string `json:"contentId"`
}

// NewRefToBinaryData instantiates a new RefToBinaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefToBinaryData(contentId string) *RefToBinaryData {
	this := RefToBinaryData{}
	this.ContentId = contentId
	return &this
}

// NewRefToBinaryDataWithDefaults instantiates a new RefToBinaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefToBinaryDataWithDefaults() *RefToBinaryData {
	this := RefToBinaryData{}
	return &this
}

// GetContentId returns the ContentId field value
func (o *RefToBinaryData) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *RefToBinaryData) GetContentIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *RefToBinaryData) SetContentId(v string) {
	o.ContentId = v
}

func (o RefToBinaryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentId"] = o.ContentId
	}
	return json.Marshal(toSerialize)
}

type NullableRefToBinaryData struct {
	value *RefToBinaryData
	isSet bool
}

func (v NullableRefToBinaryData) Get() *RefToBinaryData {
	return v.value
}

func (v *NullableRefToBinaryData) Set(val *RefToBinaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefToBinaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefToBinaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefToBinaryData(val *RefToBinaryData) *NullableRefToBinaryData {
	return &NullableRefToBinaryData{value: val, isSet: true}
}

func (v NullableRefToBinaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefToBinaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


