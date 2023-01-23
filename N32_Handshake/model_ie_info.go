/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N32_Handshake

import (
	"encoding/json"
)

// IeInfo Protection and modification policy for the IE
type IeInfo struct {
	IeLoc IeLocation `json:"ieLoc"`
	IeType IeType `json:"ieType"`
	ReqIe *string `json:"reqIe,omitempty"`
	RspIe *string `json:"rspIe,omitempty"`
	IsModifiable *bool `json:"isModifiable,omitempty"`
	IsModifiableByIpx *map[string]bool `json:"isModifiableByIpx,omitempty"`
}

// NewIeInfo instantiates a new IeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIeInfo(ieLoc IeLocation, ieType IeType) *IeInfo {
	this := IeInfo{}
	this.IeLoc = ieLoc
	this.IeType = ieType
	return &this
}

// NewIeInfoWithDefaults instantiates a new IeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIeInfoWithDefaults() *IeInfo {
	this := IeInfo{}
	return &this
}

// GetIeLoc returns the IeLoc field value
func (o *IeInfo) GetIeLoc() IeLocation {
	if o == nil {
		var ret IeLocation
		return ret
	}

	return o.IeLoc
}

// GetIeLocOk returns a tuple with the IeLoc field value
// and a boolean to check if the value has been set.
func (o *IeInfo) GetIeLocOk() (*IeLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IeLoc, true
}

// SetIeLoc sets field value
func (o *IeInfo) SetIeLoc(v IeLocation) {
	o.IeLoc = v
}

// GetIeType returns the IeType field value
func (o *IeInfo) GetIeType() IeType {
	if o == nil {
		var ret IeType
		return ret
	}

	return o.IeType
}

// GetIeTypeOk returns a tuple with the IeType field value
// and a boolean to check if the value has been set.
func (o *IeInfo) GetIeTypeOk() (*IeType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IeType, true
}

// SetIeType sets field value
func (o *IeInfo) SetIeType(v IeType) {
	o.IeType = v
}

// GetReqIe returns the ReqIe field value if set, zero value otherwise.
func (o *IeInfo) GetReqIe() string {
	if o == nil || isNil(o.ReqIe) {
		var ret string
		return ret
	}
	return *o.ReqIe
}

// GetReqIeOk returns a tuple with the ReqIe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IeInfo) GetReqIeOk() (*string, bool) {
	if o == nil || isNil(o.ReqIe) {
    return nil, false
	}
	return o.ReqIe, true
}

// HasReqIe returns a boolean if a field has been set.
func (o *IeInfo) HasReqIe() bool {
	if o != nil && !isNil(o.ReqIe) {
		return true
	}

	return false
}

// SetReqIe gets a reference to the given string and assigns it to the ReqIe field.
func (o *IeInfo) SetReqIe(v string) {
	o.ReqIe = &v
}

// GetRspIe returns the RspIe field value if set, zero value otherwise.
func (o *IeInfo) GetRspIe() string {
	if o == nil || isNil(o.RspIe) {
		var ret string
		return ret
	}
	return *o.RspIe
}

// GetRspIeOk returns a tuple with the RspIe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IeInfo) GetRspIeOk() (*string, bool) {
	if o == nil || isNil(o.RspIe) {
    return nil, false
	}
	return o.RspIe, true
}

// HasRspIe returns a boolean if a field has been set.
func (o *IeInfo) HasRspIe() bool {
	if o != nil && !isNil(o.RspIe) {
		return true
	}

	return false
}

// SetRspIe gets a reference to the given string and assigns it to the RspIe field.
func (o *IeInfo) SetRspIe(v string) {
	o.RspIe = &v
}

// GetIsModifiable returns the IsModifiable field value if set, zero value otherwise.
func (o *IeInfo) GetIsModifiable() bool {
	if o == nil || isNil(o.IsModifiable) {
		var ret bool
		return ret
	}
	return *o.IsModifiable
}

// GetIsModifiableOk returns a tuple with the IsModifiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IeInfo) GetIsModifiableOk() (*bool, bool) {
	if o == nil || isNil(o.IsModifiable) {
    return nil, false
	}
	return o.IsModifiable, true
}

// HasIsModifiable returns a boolean if a field has been set.
func (o *IeInfo) HasIsModifiable() bool {
	if o != nil && !isNil(o.IsModifiable) {
		return true
	}

	return false
}

// SetIsModifiable gets a reference to the given bool and assigns it to the IsModifiable field.
func (o *IeInfo) SetIsModifiable(v bool) {
	o.IsModifiable = &v
}

// GetIsModifiableByIpx returns the IsModifiableByIpx field value if set, zero value otherwise.
func (o *IeInfo) GetIsModifiableByIpx() map[string]bool {
	if o == nil || isNil(o.IsModifiableByIpx) {
		var ret map[string]bool
		return ret
	}
	return *o.IsModifiableByIpx
}

// GetIsModifiableByIpxOk returns a tuple with the IsModifiableByIpx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IeInfo) GetIsModifiableByIpxOk() (*map[string]bool, bool) {
	if o == nil || isNil(o.IsModifiableByIpx) {
    return nil, false
	}
	return o.IsModifiableByIpx, true
}

// HasIsModifiableByIpx returns a boolean if a field has been set.
func (o *IeInfo) HasIsModifiableByIpx() bool {
	if o != nil && !isNil(o.IsModifiableByIpx) {
		return true
	}

	return false
}

// SetIsModifiableByIpx gets a reference to the given map[string]bool and assigns it to the IsModifiableByIpx field.
func (o *IeInfo) SetIsModifiableByIpx(v map[string]bool) {
	o.IsModifiableByIpx = &v
}

func (o IeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ieLoc"] = o.IeLoc
	}
	if true {
		toSerialize["ieType"] = o.IeType
	}
	if !isNil(o.ReqIe) {
		toSerialize["reqIe"] = o.ReqIe
	}
	if !isNil(o.RspIe) {
		toSerialize["rspIe"] = o.RspIe
	}
	if !isNil(o.IsModifiable) {
		toSerialize["isModifiable"] = o.IsModifiable
	}
	if !isNil(o.IsModifiableByIpx) {
		toSerialize["isModifiableByIpx"] = o.IsModifiableByIpx
	}
	return json.Marshal(toSerialize)
}

type NullableIeInfo struct {
	value *IeInfo
	isSet bool
}

func (v NullableIeInfo) Get() *IeInfo {
	return v.value
}

func (v *NullableIeInfo) Set(val *IeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIeInfo(val *IeInfo) *NullableIeInfo {
	return &NullableIeInfo{value: val, isSet: true}
}

func (v NullableIeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


