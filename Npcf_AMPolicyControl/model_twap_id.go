/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyControl

import (
	"encoding/json"
)

// TwapId Contain the TWAP Identifier as defined in clause 4.2.8.5.3 of 3GPP TS 23.501 or the WLAN location information as defined in clause 4.5.7.2.8 of 3GPP TS 23.402. 
type TwapId struct {
	// This IE shall contain the SSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.  
	SsId string `json:"ssId"`
	// When present, it shall contain the BSSID of the access point to which the UE is attached, for trusted WLAN access, see IEEE Std 802.11-2012.  
	BssId *string `json:"bssId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	CivicAddress *string `json:"civicAddress,omitempty"`
}

// NewTwapId instantiates a new TwapId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwapId(ssId string) *TwapId {
	this := TwapId{}
	this.SsId = ssId
	return &this
}

// NewTwapIdWithDefaults instantiates a new TwapId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwapIdWithDefaults() *TwapId {
	this := TwapId{}
	return &this
}

// GetSsId returns the SsId field value
func (o *TwapId) GetSsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsId
}

// GetSsIdOk returns a tuple with the SsId field value
// and a boolean to check if the value has been set.
func (o *TwapId) GetSsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SsId, true
}

// SetSsId sets field value
func (o *TwapId) SetSsId(v string) {
	o.SsId = v
}

// GetBssId returns the BssId field value if set, zero value otherwise.
func (o *TwapId) GetBssId() string {
	if o == nil || isNil(o.BssId) {
		var ret string
		return ret
	}
	return *o.BssId
}

// GetBssIdOk returns a tuple with the BssId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwapId) GetBssIdOk() (*string, bool) {
	if o == nil || isNil(o.BssId) {
    return nil, false
	}
	return o.BssId, true
}

// HasBssId returns a boolean if a field has been set.
func (o *TwapId) HasBssId() bool {
	if o != nil && !isNil(o.BssId) {
		return true
	}

	return false
}

// SetBssId gets a reference to the given string and assigns it to the BssId field.
func (o *TwapId) SetBssId(v string) {
	o.BssId = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *TwapId) GetCivicAddress() string {
	if o == nil || isNil(o.CivicAddress) {
		var ret string
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwapId) GetCivicAddressOk() (*string, bool) {
	if o == nil || isNil(o.CivicAddress) {
    return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *TwapId) HasCivicAddress() bool {
	if o != nil && !isNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given string and assigns it to the CivicAddress field.
func (o *TwapId) SetCivicAddress(v string) {
	o.CivicAddress = &v
}

func (o TwapId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssId"] = o.SsId
	}
	if !isNil(o.BssId) {
		toSerialize["bssId"] = o.BssId
	}
	if !isNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	return json.Marshal(toSerialize)
}

type NullableTwapId struct {
	value *TwapId
	isSet bool
}

func (v NullableTwapId) Get() *TwapId {
	return v.value
}

func (v *NullableTwapId) Set(val *TwapId) {
	v.value = val
	v.isSet = true
}

func (v NullableTwapId) IsSet() bool {
	return v.isSet
}

func (v *NullableTwapId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwapId(val *TwapId) *NullableTwapId {
	return &NullableTwapId{value: val, isSet: true}
}

func (v NullableTwapId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwapId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


