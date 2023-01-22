/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_UEPolicyControl

import (
	"encoding/json"
)

// GNbId Provides the G-NB identifier.
type GNbId struct {
	// Unsigned integer representing the bit length of the gNB ID as defined in clause 9.3.1.6 of 3GPP TS 38.413 [11], within the range 22 to 32. 
	BitLength int32 `json:"bitLength"`
	// This represents the identifier of the gNB. The value of the gNB ID shall be encoded in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The padding 0 shall be added to make multiple nibbles,  the most significant character representing the padding 0 if required together with the 4 most significant bits of the gNB ID shall appear first in the string, and the character representing the 4 least significant bit of the gNB ID shall appear last in the string. 
	GNBValue string `json:"gNBValue"`
}

// NewGNbId instantiates a new GNbId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGNbId(bitLength int32, gNBValue string) *GNbId {
	this := GNbId{}
	this.BitLength = bitLength
	this.GNBValue = gNBValue
	return &this
}

// NewGNbIdWithDefaults instantiates a new GNbId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGNbIdWithDefaults() *GNbId {
	this := GNbId{}
	return &this
}

// GetBitLength returns the BitLength field value
func (o *GNbId) GetBitLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BitLength
}

// GetBitLengthOk returns a tuple with the BitLength field value
// and a boolean to check if the value has been set.
func (o *GNbId) GetBitLengthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BitLength, true
}

// SetBitLength sets field value
func (o *GNbId) SetBitLength(v int32) {
	o.BitLength = v
}

// GetGNBValue returns the GNBValue field value
func (o *GNbId) GetGNBValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GNBValue
}

// GetGNBValueOk returns a tuple with the GNBValue field value
// and a boolean to check if the value has been set.
func (o *GNbId) GetGNBValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GNBValue, true
}

// SetGNBValue sets field value
func (o *GNbId) SetGNBValue(v string) {
	o.GNBValue = v
}

func (o GNbId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bitLength"] = o.BitLength
	}
	if true {
		toSerialize["gNBValue"] = o.GNBValue
	}
	return json.Marshal(toSerialize)
}

type NullableGNbId struct {
	value *GNbId
	isSet bool
}

func (v NullableGNbId) Get() *GNbId {
	return v.value
}

func (v *NullableGNbId) Set(val *GNbId) {
	v.value = val
	v.isSet = true
}

func (v NullableGNbId) IsSet() bool {
	return v.isSet
}

func (v *NullableGNbId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGNbId(val *GNbId) *NullableGNbId {
	return &NullableGNbId{value: val, isSet: true}
}

func (v NullableGNbId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGNbId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


