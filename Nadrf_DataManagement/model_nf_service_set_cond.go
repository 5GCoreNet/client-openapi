/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// NfServiceSetCond Subscription to a set of NFs based on their Service Set Id
type NfServiceSetCond struct {
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	NfServiceSetId string `json:"nfServiceSetId"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NfSetId *string `json:"nfSetId,omitempty"`
}

// NewNfServiceSetCond instantiates a new NfServiceSetCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfServiceSetCond(nfServiceSetId string) *NfServiceSetCond {
	this := NfServiceSetCond{}
	this.NfServiceSetId = nfServiceSetId
	return &this
}

// NewNfServiceSetCondWithDefaults instantiates a new NfServiceSetCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfServiceSetCondWithDefaults() *NfServiceSetCond {
	this := NfServiceSetCond{}
	return &this
}

// GetNfServiceSetId returns the NfServiceSetId field value
func (o *NfServiceSetCond) GetNfServiceSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfServiceSetId
}

// GetNfServiceSetIdOk returns a tuple with the NfServiceSetId field value
// and a boolean to check if the value has been set.
func (o *NfServiceSetCond) GetNfServiceSetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfServiceSetId, true
}

// SetNfServiceSetId sets field value
func (o *NfServiceSetCond) SetNfServiceSetId(v string) {
	o.NfServiceSetId = v
}

// GetNfSetId returns the NfSetId field value if set, zero value otherwise.
func (o *NfServiceSetCond) GetNfSetId() string {
	if o == nil || isNil(o.NfSetId) {
		var ret string
		return ret
	}
	return *o.NfSetId
}

// GetNfSetIdOk returns a tuple with the NfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfServiceSetCond) GetNfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.NfSetId) {
    return nil, false
	}
	return o.NfSetId, true
}

// HasNfSetId returns a boolean if a field has been set.
func (o *NfServiceSetCond) HasNfSetId() bool {
	if o != nil && !isNil(o.NfSetId) {
		return true
	}

	return false
}

// SetNfSetId gets a reference to the given string and assigns it to the NfSetId field.
func (o *NfServiceSetCond) SetNfSetId(v string) {
	o.NfSetId = &v
}

func (o NfServiceSetCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfServiceSetId"] = o.NfServiceSetId
	}
	if !isNil(o.NfSetId) {
		toSerialize["nfSetId"] = o.NfSetId
	}
	return json.Marshal(toSerialize)
}

type NullableNfServiceSetCond struct {
	value *NfServiceSetCond
	isSet bool
}

func (v NullableNfServiceSetCond) Get() *NfServiceSetCond {
	return v.value
}

func (v *NullableNfServiceSetCond) Set(val *NfServiceSetCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfServiceSetCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfServiceSetCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfServiceSetCond(val *NfServiceSetCond) *NullableNfServiceSetCond {
	return &NullableNfServiceSetCond{value: val, isSet: true}
}

func (v NullableNfServiceSetCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfServiceSetCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


