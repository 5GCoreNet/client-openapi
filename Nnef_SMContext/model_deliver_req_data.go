/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_SMContext

import (
	"encoding/json"
)

// DeliverReqData The data for Deliver service request, including the Mobile Originated data to be delivered via NEF.
type DeliverReqData struct {
	Data RefToBinaryData `json:"data"`
}

// NewDeliverReqData instantiates a new DeliverReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverReqData(data RefToBinaryData) *DeliverReqData {
	this := DeliverReqData{}
	this.Data = data
	return &this
}

// NewDeliverReqDataWithDefaults instantiates a new DeliverReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverReqDataWithDefaults() *DeliverReqData {
	this := DeliverReqData{}
	return &this
}

// GetData returns the Data field value
func (o *DeliverReqData) GetData() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliverReqData) GetDataOk() (*RefToBinaryData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliverReqData) SetData(v RefToBinaryData) {
	o.Data = v
}

func (o DeliverReqData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeliverReqData struct {
	value *DeliverReqData
	isSet bool
}

func (v NullableDeliverReqData) Get() *DeliverReqData {
	return v.value
}

func (v *NullableDeliverReqData) Set(val *DeliverReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverReqData(val *DeliverReqData) *NullableDeliverReqData {
	return &NullableDeliverReqData{value: val, isSet: true}
}

func (v NullableDeliverReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


