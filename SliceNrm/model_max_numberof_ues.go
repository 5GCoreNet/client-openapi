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

// MaxNumberofUEs struct for MaxNumberofUEs
type MaxNumberofUEs struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Var3GPPNoOfUEs *int32 `json:"3GPPNoOfUEs,omitempty"`
	Non3GPPNoOfUEs *int32 `json:"non3GPPNoOfUEs,omitempty"`
}

// NewMaxNumberofUEs instantiates a new MaxNumberofUEs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumberofUEs() *MaxNumberofUEs {
	this := MaxNumberofUEs{}
	return &this
}

// NewMaxNumberofUEsWithDefaults instantiates a new MaxNumberofUEs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumberofUEsWithDefaults() *MaxNumberofUEs {
	this := MaxNumberofUEs{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *MaxNumberofUEs) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEs) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *MaxNumberofUEs) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *MaxNumberofUEs) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetVar3GPPNoOfUEs returns the Var3GPPNoOfUEs field value if set, zero value otherwise.
func (o *MaxNumberofUEs) GetVar3GPPNoOfUEs() int32 {
	if o == nil || isNil(o.Var3GPPNoOfUEs) {
		var ret int32
		return ret
	}
	return *o.Var3GPPNoOfUEs
}

// GetVar3GPPNoOfUEsOk returns a tuple with the Var3GPPNoOfUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEs) GetVar3GPPNoOfUEsOk() (*int32, bool) {
	if o == nil || isNil(o.Var3GPPNoOfUEs) {
    return nil, false
	}
	return o.Var3GPPNoOfUEs, true
}

// HasVar3GPPNoOfUEs returns a boolean if a field has been set.
func (o *MaxNumberofUEs) HasVar3GPPNoOfUEs() bool {
	if o != nil && !isNil(o.Var3GPPNoOfUEs) {
		return true
	}

	return false
}

// SetVar3GPPNoOfUEs gets a reference to the given int32 and assigns it to the Var3GPPNoOfUEs field.
func (o *MaxNumberofUEs) SetVar3GPPNoOfUEs(v int32) {
	o.Var3GPPNoOfUEs = &v
}

// GetNon3GPPNoOfUEs returns the Non3GPPNoOfUEs field value if set, zero value otherwise.
func (o *MaxNumberofUEs) GetNon3GPPNoOfUEs() int32 {
	if o == nil || isNil(o.Non3GPPNoOfUEs) {
		var ret int32
		return ret
	}
	return *o.Non3GPPNoOfUEs
}

// GetNon3GPPNoOfUEsOk returns a tuple with the Non3GPPNoOfUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEs) GetNon3GPPNoOfUEsOk() (*int32, bool) {
	if o == nil || isNil(o.Non3GPPNoOfUEs) {
    return nil, false
	}
	return o.Non3GPPNoOfUEs, true
}

// HasNon3GPPNoOfUEs returns a boolean if a field has been set.
func (o *MaxNumberofUEs) HasNon3GPPNoOfUEs() bool {
	if o != nil && !isNil(o.Non3GPPNoOfUEs) {
		return true
	}

	return false
}

// SetNon3GPPNoOfUEs gets a reference to the given int32 and assigns it to the Non3GPPNoOfUEs field.
func (o *MaxNumberofUEs) SetNon3GPPNoOfUEs(v int32) {
	o.Non3GPPNoOfUEs = &v
}

func (o MaxNumberofUEs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.Var3GPPNoOfUEs) {
		toSerialize["3GPPNoOfUEs"] = o.Var3GPPNoOfUEs
	}
	if !isNil(o.Non3GPPNoOfUEs) {
		toSerialize["non3GPPNoOfUEs"] = o.Non3GPPNoOfUEs
	}
	return json.Marshal(toSerialize)
}

type NullableMaxNumberofUEs struct {
	value *MaxNumberofUEs
	isSet bool
}

func (v NullableMaxNumberofUEs) Get() *MaxNumberofUEs {
	return v.value
}

func (v *NullableMaxNumberofUEs) Set(val *MaxNumberofUEs) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumberofUEs) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumberofUEs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumberofUEs(val *MaxNumberofUEs) *NullableMaxNumberofUEs {
	return &NullableMaxNumberofUEs{value: val, isSet: true}
}

func (v NullableMaxNumberofUEs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumberofUEs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

