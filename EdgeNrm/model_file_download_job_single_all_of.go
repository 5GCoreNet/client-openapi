/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// FileDownloadJobSingleAllOf struct for FileDownloadJobSingleAllOf
type FileDownloadJobSingleAllOf struct {
	Attributes *FileDownloadJobSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFileDownloadJobSingleAllOf instantiates a new FileDownloadJobSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDownloadJobSingleAllOf() *FileDownloadJobSingleAllOf {
	this := FileDownloadJobSingleAllOf{}
	return &this
}

// NewFileDownloadJobSingleAllOfWithDefaults instantiates a new FileDownloadJobSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDownloadJobSingleAllOfWithDefaults() *FileDownloadJobSingleAllOf {
	this := FileDownloadJobSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FileDownloadJobSingleAllOf) GetAttributes() FileDownloadJobSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret FileDownloadJobSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingleAllOf) GetAttributesOk() (*FileDownloadJobSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FileDownloadJobSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FileDownloadJobSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FileDownloadJobSingleAllOf) SetAttributes(v FileDownloadJobSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FileDownloadJobSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableFileDownloadJobSingleAllOf struct {
	value *FileDownloadJobSingleAllOf
	isSet bool
}

func (v NullableFileDownloadJobSingleAllOf) Get() *FileDownloadJobSingleAllOf {
	return v.value
}

func (v *NullableFileDownloadJobSingleAllOf) Set(val *FileDownloadJobSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDownloadJobSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDownloadJobSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDownloadJobSingleAllOf(val *FileDownloadJobSingleAllOf) *NullableFileDownloadJobSingleAllOf {
	return &NullableFileDownloadJobSingleAllOf{value: val, isSet: true}
}

func (v NullableFileDownloadJobSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDownloadJobSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


