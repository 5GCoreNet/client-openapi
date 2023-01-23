/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// N6Protection struct for N6Protection
type N6Protection struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	SecFuncList []SecFunc `json:"secFuncList,omitempty"`
}

// NewN6Protection instantiates a new N6Protection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN6Protection() *N6Protection {
	this := N6Protection{}
	return &this
}

// NewN6ProtectionWithDefaults instantiates a new N6Protection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN6ProtectionWithDefaults() *N6Protection {
	this := N6Protection{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *N6Protection) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N6Protection) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *N6Protection) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *N6Protection) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetSecFuncList returns the SecFuncList field value if set, zero value otherwise.
func (o *N6Protection) GetSecFuncList() []SecFunc {
	if o == nil || isNil(o.SecFuncList) {
		var ret []SecFunc
		return ret
	}
	return o.SecFuncList
}

// GetSecFuncListOk returns a tuple with the SecFuncList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N6Protection) GetSecFuncListOk() ([]SecFunc, bool) {
	if o == nil || isNil(o.SecFuncList) {
    return nil, false
	}
	return o.SecFuncList, true
}

// HasSecFuncList returns a boolean if a field has been set.
func (o *N6Protection) HasSecFuncList() bool {
	if o != nil && !isNil(o.SecFuncList) {
		return true
	}

	return false
}

// SetSecFuncList gets a reference to the given []SecFunc and assigns it to the SecFuncList field.
func (o *N6Protection) SetSecFuncList(v []SecFunc) {
	o.SecFuncList = v
}

func (o N6Protection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.SecFuncList) {
		toSerialize["secFuncList"] = o.SecFuncList
	}
	return json.Marshal(toSerialize)
}

type NullableN6Protection struct {
	value *N6Protection
	isSet bool
}

func (v NullableN6Protection) Get() *N6Protection {
	return v.value
}

func (v *NullableN6Protection) Set(val *N6Protection) {
	v.value = val
	v.isSet = true
}

func (v NullableN6Protection) IsSet() bool {
	return v.isSet
}

func (v *NullableN6Protection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN6Protection(val *N6Protection) *NullableN6Protection {
	return &NullableN6Protection{value: val, isSet: true}
}

func (v NullableN6Protection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN6Protection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


