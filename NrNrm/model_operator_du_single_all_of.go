/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// OperatorDuSingleAllOf struct for OperatorDuSingleAllOf
type OperatorDuSingleAllOf struct {
	GnbId *string `json:"gnbId,omitempty"`
	GnbIdLength *int32 `json:"gnbIdLength,omitempty"`
}

// NewOperatorDuSingleAllOf instantiates a new OperatorDuSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorDuSingleAllOf() *OperatorDuSingleAllOf {
	this := OperatorDuSingleAllOf{}
	return &this
}

// NewOperatorDuSingleAllOfWithDefaults instantiates a new OperatorDuSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorDuSingleAllOfWithDefaults() *OperatorDuSingleAllOf {
	this := OperatorDuSingleAllOf{}
	return &this
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *OperatorDuSingleAllOf) GetGnbId() string {
	if o == nil || isNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingleAllOf) GetGnbIdOk() (*string, bool) {
	if o == nil || isNil(o.GnbId) {
    return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *OperatorDuSingleAllOf) HasGnbId() bool {
	if o != nil && !isNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *OperatorDuSingleAllOf) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *OperatorDuSingleAllOf) GetGnbIdLength() int32 {
	if o == nil || isNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingleAllOf) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || isNil(o.GnbIdLength) {
    return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *OperatorDuSingleAllOf) HasGnbIdLength() bool {
	if o != nil && !isNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *OperatorDuSingleAllOf) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

func (o OperatorDuSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !isNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorDuSingleAllOf struct {
	value *OperatorDuSingleAllOf
	isSet bool
}

func (v NullableOperatorDuSingleAllOf) Get() *OperatorDuSingleAllOf {
	return v.value
}

func (v *NullableOperatorDuSingleAllOf) Set(val *OperatorDuSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorDuSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorDuSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorDuSingleAllOf(val *OperatorDuSingleAllOf) *NullableOperatorDuSingleAllOf {
	return &NullableOperatorDuSingleAllOf{value: val, isSet: true}
}

func (v NullableOperatorDuSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorDuSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


