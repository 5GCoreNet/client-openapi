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

// EDNConnectionInfo struct for EDNConnectionInfo
type EDNConnectionInfo struct {
	DNN *string `json:"dNN,omitempty"`
	EDNServiceArea *ServingLocation `json:"eDNServiceArea,omitempty"`
}

// NewEDNConnectionInfo instantiates a new EDNConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEDNConnectionInfo() *EDNConnectionInfo {
	this := EDNConnectionInfo{}
	return &this
}

// NewEDNConnectionInfoWithDefaults instantiates a new EDNConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEDNConnectionInfoWithDefaults() *EDNConnectionInfo {
	this := EDNConnectionInfo{}
	return &this
}

// GetDNN returns the DNN field value if set, zero value otherwise.
func (o *EDNConnectionInfo) GetDNN() string {
	if o == nil || isNil(o.DNN) {
		var ret string
		return ret
	}
	return *o.DNN
}

// GetDNNOk returns a tuple with the DNN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EDNConnectionInfo) GetDNNOk() (*string, bool) {
	if o == nil || isNil(o.DNN) {
    return nil, false
	}
	return o.DNN, true
}

// HasDNN returns a boolean if a field has been set.
func (o *EDNConnectionInfo) HasDNN() bool {
	if o != nil && !isNil(o.DNN) {
		return true
	}

	return false
}

// SetDNN gets a reference to the given string and assigns it to the DNN field.
func (o *EDNConnectionInfo) SetDNN(v string) {
	o.DNN = &v
}

// GetEDNServiceArea returns the EDNServiceArea field value if set, zero value otherwise.
func (o *EDNConnectionInfo) GetEDNServiceArea() ServingLocation {
	if o == nil || isNil(o.EDNServiceArea) {
		var ret ServingLocation
		return ret
	}
	return *o.EDNServiceArea
}

// GetEDNServiceAreaOk returns a tuple with the EDNServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EDNConnectionInfo) GetEDNServiceAreaOk() (*ServingLocation, bool) {
	if o == nil || isNil(o.EDNServiceArea) {
    return nil, false
	}
	return o.EDNServiceArea, true
}

// HasEDNServiceArea returns a boolean if a field has been set.
func (o *EDNConnectionInfo) HasEDNServiceArea() bool {
	if o != nil && !isNil(o.EDNServiceArea) {
		return true
	}

	return false
}

// SetEDNServiceArea gets a reference to the given ServingLocation and assigns it to the EDNServiceArea field.
func (o *EDNConnectionInfo) SetEDNServiceArea(v ServingLocation) {
	o.EDNServiceArea = &v
}

func (o EDNConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DNN) {
		toSerialize["dNN"] = o.DNN
	}
	if !isNil(o.EDNServiceArea) {
		toSerialize["eDNServiceArea"] = o.EDNServiceArea
	}
	return json.Marshal(toSerialize)
}

type NullableEDNConnectionInfo struct {
	value *EDNConnectionInfo
	isSet bool
}

func (v NullableEDNConnectionInfo) Get() *EDNConnectionInfo {
	return v.value
}

func (v *NullableEDNConnectionInfo) Set(val *EDNConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEDNConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEDNConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEDNConnectionInfo(val *EDNConnectionInfo) *NullableEDNConnectionInfo {
	return &NullableEDNConnectionInfo{value: val, isSet: true}
}

func (v NullableEDNConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEDNConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


