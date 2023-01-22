/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// PWSResponseData Data related PWS included in a N2 Information Transfer response
type PWSResponseData struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NgapMessageType int32 `json:"ngapMessageType"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	SerialNumber int32 `json:"serialNumber"`
	MessageIdentifier int32 `json:"messageIdentifier"`
	UnknownTaiList []Tai `json:"unknownTaiList,omitempty"`
}

// NewPWSResponseData instantiates a new PWSResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPWSResponseData(ngapMessageType int32, serialNumber int32, messageIdentifier int32) *PWSResponseData {
	this := PWSResponseData{}
	this.NgapMessageType = ngapMessageType
	this.SerialNumber = serialNumber
	this.MessageIdentifier = messageIdentifier
	return &this
}

// NewPWSResponseDataWithDefaults instantiates a new PWSResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPWSResponseDataWithDefaults() *PWSResponseData {
	this := PWSResponseData{}
	return &this
}

// GetNgapMessageType returns the NgapMessageType field value
func (o *PWSResponseData) GetNgapMessageType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NgapMessageType
}

// GetNgapMessageTypeOk returns a tuple with the NgapMessageType field value
// and a boolean to check if the value has been set.
func (o *PWSResponseData) GetNgapMessageTypeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NgapMessageType, true
}

// SetNgapMessageType sets field value
func (o *PWSResponseData) SetNgapMessageType(v int32) {
	o.NgapMessageType = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *PWSResponseData) GetSerialNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *PWSResponseData) GetSerialNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *PWSResponseData) SetSerialNumber(v int32) {
	o.SerialNumber = v
}

// GetMessageIdentifier returns the MessageIdentifier field value
func (o *PWSResponseData) GetMessageIdentifier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageIdentifier
}

// GetMessageIdentifierOk returns a tuple with the MessageIdentifier field value
// and a boolean to check if the value has been set.
func (o *PWSResponseData) GetMessageIdentifierOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MessageIdentifier, true
}

// SetMessageIdentifier sets field value
func (o *PWSResponseData) SetMessageIdentifier(v int32) {
	o.MessageIdentifier = v
}

// GetUnknownTaiList returns the UnknownTaiList field value if set, zero value otherwise.
func (o *PWSResponseData) GetUnknownTaiList() []Tai {
	if o == nil || isNil(o.UnknownTaiList) {
		var ret []Tai
		return ret
	}
	return o.UnknownTaiList
}

// GetUnknownTaiListOk returns a tuple with the UnknownTaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PWSResponseData) GetUnknownTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.UnknownTaiList) {
    return nil, false
	}
	return o.UnknownTaiList, true
}

// HasUnknownTaiList returns a boolean if a field has been set.
func (o *PWSResponseData) HasUnknownTaiList() bool {
	if o != nil && !isNil(o.UnknownTaiList) {
		return true
	}

	return false
}

// SetUnknownTaiList gets a reference to the given []Tai and assigns it to the UnknownTaiList field.
func (o *PWSResponseData) SetUnknownTaiList(v []Tai) {
	o.UnknownTaiList = v
}

func (o PWSResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ngapMessageType"] = o.NgapMessageType
	}
	if true {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if true {
		toSerialize["messageIdentifier"] = o.MessageIdentifier
	}
	if !isNil(o.UnknownTaiList) {
		toSerialize["unknownTaiList"] = o.UnknownTaiList
	}
	return json.Marshal(toSerialize)
}

type NullablePWSResponseData struct {
	value *PWSResponseData
	isSet bool
}

func (v NullablePWSResponseData) Get() *PWSResponseData {
	return v.value
}

func (v *NullablePWSResponseData) Set(val *PWSResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePWSResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePWSResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePWSResponseData(val *PWSResponseData) *NullablePWSResponseData {
	return &NullablePWSResponseData{value: val, isSet: true}
}

func (v NullablePWSResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePWSResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


