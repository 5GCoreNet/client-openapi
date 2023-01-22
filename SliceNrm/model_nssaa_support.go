/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
)

// NSSAASupport struct for NSSAASupport
type NSSAASupport struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Support *Support `json:"support,omitempty"`
}

// NewNSSAASupport instantiates a new NSSAASupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSSAASupport() *NSSAASupport {
	this := NSSAASupport{}
	return &this
}

// NewNSSAASupportWithDefaults instantiates a new NSSAASupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSSAASupportWithDefaults() *NSSAASupport {
	this := NSSAASupport{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *NSSAASupport) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSSAASupport) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *NSSAASupport) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *NSSAASupport) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *NSSAASupport) GetSupport() Support {
	if o == nil || isNil(o.Support) {
		var ret Support
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSSAASupport) GetSupportOk() (*Support, bool) {
	if o == nil || isNil(o.Support) {
    return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *NSSAASupport) HasSupport() bool {
	if o != nil && !isNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given Support and assigns it to the Support field.
func (o *NSSAASupport) SetSupport(v Support) {
	o.Support = &v
}

func (o NSSAASupport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	return json.Marshal(toSerialize)
}

type NullableNSSAASupport struct {
	value *NSSAASupport
	isSet bool
}

func (v NullableNSSAASupport) Get() *NSSAASupport {
	return v.value
}

func (v *NullableNSSAASupport) Set(val *NSSAASupport) {
	v.value = val
	v.isSet = true
}

func (v NullableNSSAASupport) IsSet() bool {
	return v.isSet
}

func (v *NullableNSSAASupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSSAASupport(val *NSSAASupport) *NullableNSSAASupport {
	return &NullableNSSAASupport{value: val, isSet: true}
}

func (v NullableNSSAASupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSSAASupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


