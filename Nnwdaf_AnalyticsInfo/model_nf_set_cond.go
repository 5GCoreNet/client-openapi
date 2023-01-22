/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// NfSetCond Subscription to a set of NFs based on their Set Id
type NfSetCond struct {
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NfSetId string `json:"nfSetId"`
}

// NewNfSetCond instantiates a new NfSetCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfSetCond(nfSetId string) *NfSetCond {
	this := NfSetCond{}
	this.NfSetId = nfSetId
	return &this
}

// NewNfSetCondWithDefaults instantiates a new NfSetCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfSetCondWithDefaults() *NfSetCond {
	this := NfSetCond{}
	return &this
}

// GetNfSetId returns the NfSetId field value
func (o *NfSetCond) GetNfSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfSetId
}

// GetNfSetIdOk returns a tuple with the NfSetId field value
// and a boolean to check if the value has been set.
func (o *NfSetCond) GetNfSetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfSetId, true
}

// SetNfSetId sets field value
func (o *NfSetCond) SetNfSetId(v string) {
	o.NfSetId = v
}

func (o NfSetCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfSetId"] = o.NfSetId
	}
	return json.Marshal(toSerialize)
}

type NullableNfSetCond struct {
	value *NfSetCond
	isSet bool
}

func (v NullableNfSetCond) Get() *NfSetCond {
	return v.value
}

func (v *NullableNfSetCond) Set(val *NfSetCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfSetCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfSetCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfSetCond(val *NfSetCond) *NullableNfSetCond {
	return &NullableNfSetCond{value: val, isSet: true}
}

func (v NullableNfSetCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfSetCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


