/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// UserMgmtOpen struct for UserMgmtOpen
type UserMgmtOpen struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Support *Support `json:"support,omitempty"`
}

// NewUserMgmtOpen instantiates a new UserMgmtOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMgmtOpen() *UserMgmtOpen {
	this := UserMgmtOpen{}
	return &this
}

// NewUserMgmtOpenWithDefaults instantiates a new UserMgmtOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMgmtOpenWithDefaults() *UserMgmtOpen {
	this := UserMgmtOpen{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *UserMgmtOpen) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMgmtOpen) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *UserMgmtOpen) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *UserMgmtOpen) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *UserMgmtOpen) GetSupport() Support {
	if o == nil || isNil(o.Support) {
		var ret Support
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMgmtOpen) GetSupportOk() (*Support, bool) {
	if o == nil || isNil(o.Support) {
    return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *UserMgmtOpen) HasSupport() bool {
	if o != nil && !isNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given Support and assigns it to the Support field.
func (o *UserMgmtOpen) SetSupport(v Support) {
	o.Support = &v
}

func (o UserMgmtOpen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	return json.Marshal(toSerialize)
}

type NullableUserMgmtOpen struct {
	value *UserMgmtOpen
	isSet bool
}

func (v NullableUserMgmtOpen) Get() *UserMgmtOpen {
	return v.value
}

func (v *NullableUserMgmtOpen) Set(val *UserMgmtOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMgmtOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMgmtOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMgmtOpen(val *UserMgmtOpen) *NullableUserMgmtOpen {
	return &NullableUserMgmtOpen{value: val, isSet: true}
}

func (v NullableUserMgmtOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMgmtOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


