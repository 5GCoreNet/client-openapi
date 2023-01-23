/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// AcceptableServiceInfo Indicates the maximum bandwidth that shall be authorized by the PCF.
type AcceptableServiceInfo struct {
	// Indicates the maximum bandwidth that shall be authorized by the PCF for each media component of the map. The key of the map is the media component number.
	AccBwMedComps *map[string]MediaComponent `json:"accBwMedComps,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwUl *string `json:"marBwUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwDl *string `json:"marBwDl,omitempty"`
}

// NewAcceptableServiceInfo instantiates a new AcceptableServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptableServiceInfo() *AcceptableServiceInfo {
	this := AcceptableServiceInfo{}
	return &this
}

// NewAcceptableServiceInfoWithDefaults instantiates a new AcceptableServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptableServiceInfoWithDefaults() *AcceptableServiceInfo {
	this := AcceptableServiceInfo{}
	return &this
}

// GetAccBwMedComps returns the AccBwMedComps field value if set, zero value otherwise.
func (o *AcceptableServiceInfo) GetAccBwMedComps() map[string]MediaComponent {
	if o == nil || isNil(o.AccBwMedComps) {
		var ret map[string]MediaComponent
		return ret
	}
	return *o.AccBwMedComps
}

// GetAccBwMedCompsOk returns a tuple with the AccBwMedComps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptableServiceInfo) GetAccBwMedCompsOk() (*map[string]MediaComponent, bool) {
	if o == nil || isNil(o.AccBwMedComps) {
    return nil, false
	}
	return o.AccBwMedComps, true
}

// HasAccBwMedComps returns a boolean if a field has been set.
func (o *AcceptableServiceInfo) HasAccBwMedComps() bool {
	if o != nil && !isNil(o.AccBwMedComps) {
		return true
	}

	return false
}

// SetAccBwMedComps gets a reference to the given map[string]MediaComponent and assigns it to the AccBwMedComps field.
func (o *AcceptableServiceInfo) SetAccBwMedComps(v map[string]MediaComponent) {
	o.AccBwMedComps = &v
}

// GetMarBwUl returns the MarBwUl field value if set, zero value otherwise.
func (o *AcceptableServiceInfo) GetMarBwUl() string {
	if o == nil || isNil(o.MarBwUl) {
		var ret string
		return ret
	}
	return *o.MarBwUl
}

// GetMarBwUlOk returns a tuple with the MarBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptableServiceInfo) GetMarBwUlOk() (*string, bool) {
	if o == nil || isNil(o.MarBwUl) {
    return nil, false
	}
	return o.MarBwUl, true
}

// HasMarBwUl returns a boolean if a field has been set.
func (o *AcceptableServiceInfo) HasMarBwUl() bool {
	if o != nil && !isNil(o.MarBwUl) {
		return true
	}

	return false
}

// SetMarBwUl gets a reference to the given string and assigns it to the MarBwUl field.
func (o *AcceptableServiceInfo) SetMarBwUl(v string) {
	o.MarBwUl = &v
}

// GetMarBwDl returns the MarBwDl field value if set, zero value otherwise.
func (o *AcceptableServiceInfo) GetMarBwDl() string {
	if o == nil || isNil(o.MarBwDl) {
		var ret string
		return ret
	}
	return *o.MarBwDl
}

// GetMarBwDlOk returns a tuple with the MarBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptableServiceInfo) GetMarBwDlOk() (*string, bool) {
	if o == nil || isNil(o.MarBwDl) {
    return nil, false
	}
	return o.MarBwDl, true
}

// HasMarBwDl returns a boolean if a field has been set.
func (o *AcceptableServiceInfo) HasMarBwDl() bool {
	if o != nil && !isNil(o.MarBwDl) {
		return true
	}

	return false
}

// SetMarBwDl gets a reference to the given string and assigns it to the MarBwDl field.
func (o *AcceptableServiceInfo) SetMarBwDl(v string) {
	o.MarBwDl = &v
}

func (o AcceptableServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccBwMedComps) {
		toSerialize["accBwMedComps"] = o.AccBwMedComps
	}
	if !isNil(o.MarBwUl) {
		toSerialize["marBwUl"] = o.MarBwUl
	}
	if !isNil(o.MarBwDl) {
		toSerialize["marBwDl"] = o.MarBwDl
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptableServiceInfo struct {
	value *AcceptableServiceInfo
	isSet bool
}

func (v NullableAcceptableServiceInfo) Get() *AcceptableServiceInfo {
	return v.value
}

func (v *NullableAcceptableServiceInfo) Set(val *AcceptableServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptableServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptableServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptableServiceInfo(val *AcceptableServiceInfo) *NullableAcceptableServiceInfo {
	return &NullableAcceptableServiceInfo{value: val, isSet: true}
}

func (v NullableAcceptableServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptableServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


