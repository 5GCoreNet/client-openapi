/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// V2xSubscriptionData struct for V2xSubscriptionData
type V2xSubscriptionData struct {
	NrV2xServicesAuth *NrV2xAuth `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	NrUePc5Ambr *string `json:"nrUePc5Ambr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	LtePc5Ambr *string `json:"ltePc5Ambr,omitempty"`
}

// NewV2xSubscriptionData instantiates a new V2xSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2xSubscriptionData() *V2xSubscriptionData {
	this := V2xSubscriptionData{}
	return &this
}

// NewV2xSubscriptionDataWithDefaults instantiates a new V2xSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2xSubscriptionDataWithDefaults() *V2xSubscriptionData {
	this := V2xSubscriptionData{}
	return &this
}

// GetNrV2xServicesAuth returns the NrV2xServicesAuth field value if set, zero value otherwise.
func (o *V2xSubscriptionData) GetNrV2xServicesAuth() NrV2xAuth {
	if o == nil || isNil(o.NrV2xServicesAuth) {
		var ret NrV2xAuth
		return ret
	}
	return *o.NrV2xServicesAuth
}

// GetNrV2xServicesAuthOk returns a tuple with the NrV2xServicesAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xSubscriptionData) GetNrV2xServicesAuthOk() (*NrV2xAuth, bool) {
	if o == nil || isNil(o.NrV2xServicesAuth) {
    return nil, false
	}
	return o.NrV2xServicesAuth, true
}

// HasNrV2xServicesAuth returns a boolean if a field has been set.
func (o *V2xSubscriptionData) HasNrV2xServicesAuth() bool {
	if o != nil && !isNil(o.NrV2xServicesAuth) {
		return true
	}

	return false
}

// SetNrV2xServicesAuth gets a reference to the given NrV2xAuth and assigns it to the NrV2xServicesAuth field.
func (o *V2xSubscriptionData) SetNrV2xServicesAuth(v NrV2xAuth) {
	o.NrV2xServicesAuth = &v
}

// GetLteV2xServicesAuth returns the LteV2xServicesAuth field value if set, zero value otherwise.
func (o *V2xSubscriptionData) GetLteV2xServicesAuth() LteV2xAuth {
	if o == nil || isNil(o.LteV2xServicesAuth) {
		var ret LteV2xAuth
		return ret
	}
	return *o.LteV2xServicesAuth
}

// GetLteV2xServicesAuthOk returns a tuple with the LteV2xServicesAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xSubscriptionData) GetLteV2xServicesAuthOk() (*LteV2xAuth, bool) {
	if o == nil || isNil(o.LteV2xServicesAuth) {
    return nil, false
	}
	return o.LteV2xServicesAuth, true
}

// HasLteV2xServicesAuth returns a boolean if a field has been set.
func (o *V2xSubscriptionData) HasLteV2xServicesAuth() bool {
	if o != nil && !isNil(o.LteV2xServicesAuth) {
		return true
	}

	return false
}

// SetLteV2xServicesAuth gets a reference to the given LteV2xAuth and assigns it to the LteV2xServicesAuth field.
func (o *V2xSubscriptionData) SetLteV2xServicesAuth(v LteV2xAuth) {
	o.LteV2xServicesAuth = &v
}

// GetNrUePc5Ambr returns the NrUePc5Ambr field value if set, zero value otherwise.
func (o *V2xSubscriptionData) GetNrUePc5Ambr() string {
	if o == nil || isNil(o.NrUePc5Ambr) {
		var ret string
		return ret
	}
	return *o.NrUePc5Ambr
}

// GetNrUePc5AmbrOk returns a tuple with the NrUePc5Ambr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xSubscriptionData) GetNrUePc5AmbrOk() (*string, bool) {
	if o == nil || isNil(o.NrUePc5Ambr) {
    return nil, false
	}
	return o.NrUePc5Ambr, true
}

// HasNrUePc5Ambr returns a boolean if a field has been set.
func (o *V2xSubscriptionData) HasNrUePc5Ambr() bool {
	if o != nil && !isNil(o.NrUePc5Ambr) {
		return true
	}

	return false
}

// SetNrUePc5Ambr gets a reference to the given string and assigns it to the NrUePc5Ambr field.
func (o *V2xSubscriptionData) SetNrUePc5Ambr(v string) {
	o.NrUePc5Ambr = &v
}

// GetLtePc5Ambr returns the LtePc5Ambr field value if set, zero value otherwise.
func (o *V2xSubscriptionData) GetLtePc5Ambr() string {
	if o == nil || isNil(o.LtePc5Ambr) {
		var ret string
		return ret
	}
	return *o.LtePc5Ambr
}

// GetLtePc5AmbrOk returns a tuple with the LtePc5Ambr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xSubscriptionData) GetLtePc5AmbrOk() (*string, bool) {
	if o == nil || isNil(o.LtePc5Ambr) {
    return nil, false
	}
	return o.LtePc5Ambr, true
}

// HasLtePc5Ambr returns a boolean if a field has been set.
func (o *V2xSubscriptionData) HasLtePc5Ambr() bool {
	if o != nil && !isNil(o.LtePc5Ambr) {
		return true
	}

	return false
}

// SetLtePc5Ambr gets a reference to the given string and assigns it to the LtePc5Ambr field.
func (o *V2xSubscriptionData) SetLtePc5Ambr(v string) {
	o.LtePc5Ambr = &v
}

func (o V2xSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NrV2xServicesAuth) {
		toSerialize["nrV2xServicesAuth"] = o.NrV2xServicesAuth
	}
	if !isNil(o.LteV2xServicesAuth) {
		toSerialize["lteV2xServicesAuth"] = o.LteV2xServicesAuth
	}
	if !isNil(o.NrUePc5Ambr) {
		toSerialize["nrUePc5Ambr"] = o.NrUePc5Ambr
	}
	if !isNil(o.LtePc5Ambr) {
		toSerialize["ltePc5Ambr"] = o.LtePc5Ambr
	}
	return json.Marshal(toSerialize)
}

type NullableV2xSubscriptionData struct {
	value *V2xSubscriptionData
	isSet bool
}

func (v NullableV2xSubscriptionData) Get() *V2xSubscriptionData {
	return v.value
}

func (v *NullableV2xSubscriptionData) Set(val *V2xSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2xSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2xSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2xSubscriptionData(val *V2xSubscriptionData) *NullableV2xSubscriptionData {
	return &NullableV2xSubscriptionData{value: val, isSet: true}
}

func (v NullableV2xSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2xSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


