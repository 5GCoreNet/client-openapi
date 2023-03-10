/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnsacf_NSAC

import (
	"encoding/json"
)

// PduACRequestData struct for PduACRequestData
type PduACRequestData struct {
	PduACRequestInfo []PduACRequestInfo `json:"pduACRequestInfo"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfId *string `json:"nfId,omitempty"`
	// Fully Qualified Domain Name
	PgwFqdn *string `json:"pgwFqdn,omitempty"`
}

// NewPduACRequestData instantiates a new PduACRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduACRequestData(pduACRequestInfo []PduACRequestInfo) *PduACRequestData {
	this := PduACRequestData{}
	this.PduACRequestInfo = pduACRequestInfo
	return &this
}

// NewPduACRequestDataWithDefaults instantiates a new PduACRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduACRequestDataWithDefaults() *PduACRequestData {
	this := PduACRequestData{}
	return &this
}

// GetPduACRequestInfo returns the PduACRequestInfo field value
func (o *PduACRequestData) GetPduACRequestInfo() []PduACRequestInfo {
	if o == nil {
		var ret []PduACRequestInfo
		return ret
	}

	return o.PduACRequestInfo
}

// GetPduACRequestInfoOk returns a tuple with the PduACRequestInfo field value
// and a boolean to check if the value has been set.
func (o *PduACRequestData) GetPduACRequestInfoOk() ([]PduACRequestInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.PduACRequestInfo, true
}

// SetPduACRequestInfo sets field value
func (o *PduACRequestData) SetPduACRequestInfo(v []PduACRequestInfo) {
	o.PduACRequestInfo = v
}

// GetNfId returns the NfId field value if set, zero value otherwise.
func (o *PduACRequestData) GetNfId() string {
	if o == nil || isNil(o.NfId) {
		var ret string
		return ret
	}
	return *o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduACRequestData) GetNfIdOk() (*string, bool) {
	if o == nil || isNil(o.NfId) {
    return nil, false
	}
	return o.NfId, true
}

// HasNfId returns a boolean if a field has been set.
func (o *PduACRequestData) HasNfId() bool {
	if o != nil && !isNil(o.NfId) {
		return true
	}

	return false
}

// SetNfId gets a reference to the given string and assigns it to the NfId field.
func (o *PduACRequestData) SetNfId(v string) {
	o.NfId = &v
}

// GetPgwFqdn returns the PgwFqdn field value if set, zero value otherwise.
func (o *PduACRequestData) GetPgwFqdn() string {
	if o == nil || isNil(o.PgwFqdn) {
		var ret string
		return ret
	}
	return *o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduACRequestData) GetPgwFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PgwFqdn) {
    return nil, false
	}
	return o.PgwFqdn, true
}

// HasPgwFqdn returns a boolean if a field has been set.
func (o *PduACRequestData) HasPgwFqdn() bool {
	if o != nil && !isNil(o.PgwFqdn) {
		return true
	}

	return false
}

// SetPgwFqdn gets a reference to the given string and assigns it to the PgwFqdn field.
func (o *PduACRequestData) SetPgwFqdn(v string) {
	o.PgwFqdn = &v
}

func (o PduACRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pduACRequestInfo"] = o.PduACRequestInfo
	}
	if !isNil(o.NfId) {
		toSerialize["nfId"] = o.NfId
	}
	if !isNil(o.PgwFqdn) {
		toSerialize["pgwFqdn"] = o.PgwFqdn
	}
	return json.Marshal(toSerialize)
}

type NullablePduACRequestData struct {
	value *PduACRequestData
	isSet bool
}

func (v NullablePduACRequestData) Get() *PduACRequestData {
	return v.value
}

func (v *NullablePduACRequestData) Set(val *PduACRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePduACRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePduACRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduACRequestData(val *PduACRequestData) *NullablePduACRequestData {
	return &NullablePduACRequestData{value: val, isSet: true}
}

func (v NullablePduACRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduACRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


