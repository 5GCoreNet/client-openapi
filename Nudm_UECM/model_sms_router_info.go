/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UECM

import (
	"encoding/json"
)

// SmsRouterInfo Addressing information of the SMS Router configured at the UDM
type SmsRouterInfo struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId *string `json:"nfInstanceId,omitempty"`
	DiameterAddress *NetworkNodeDiameterAddress `json:"diameterAddress,omitempty"`
	MapAddress *string `json:"mapAddress,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	RouterIpv4 *string `json:"routerIpv4,omitempty"`
	RouterIpv6 *Ipv6Addr `json:"routerIpv6,omitempty"`
	// Fully Qualified Domain Name
	RouterFqdn *string `json:"routerFqdn,omitempty"`
}

// NewSmsRouterInfo instantiates a new SmsRouterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsRouterInfo() *SmsRouterInfo {
	this := SmsRouterInfo{}
	return &this
}

// NewSmsRouterInfoWithDefaults instantiates a new SmsRouterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsRouterInfoWithDefaults() *SmsRouterInfo {
	this := SmsRouterInfo{}
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetNfInstanceId() string {
	if o == nil || isNil(o.NfInstanceId) {
		var ret string
		return ret
	}
	return *o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetNfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.NfInstanceId) {
    return nil, false
	}
	return o.NfInstanceId, true
}

// HasNfInstanceId returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasNfInstanceId() bool {
	if o != nil && !isNil(o.NfInstanceId) {
		return true
	}

	return false
}

// SetNfInstanceId gets a reference to the given string and assigns it to the NfInstanceId field.
func (o *SmsRouterInfo) SetNfInstanceId(v string) {
	o.NfInstanceId = &v
}

// GetDiameterAddress returns the DiameterAddress field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetDiameterAddress() NetworkNodeDiameterAddress {
	if o == nil || isNil(o.DiameterAddress) {
		var ret NetworkNodeDiameterAddress
		return ret
	}
	return *o.DiameterAddress
}

// GetDiameterAddressOk returns a tuple with the DiameterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetDiameterAddressOk() (*NetworkNodeDiameterAddress, bool) {
	if o == nil || isNil(o.DiameterAddress) {
    return nil, false
	}
	return o.DiameterAddress, true
}

// HasDiameterAddress returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasDiameterAddress() bool {
	if o != nil && !isNil(o.DiameterAddress) {
		return true
	}

	return false
}

// SetDiameterAddress gets a reference to the given NetworkNodeDiameterAddress and assigns it to the DiameterAddress field.
func (o *SmsRouterInfo) SetDiameterAddress(v NetworkNodeDiameterAddress) {
	o.DiameterAddress = &v
}

// GetMapAddress returns the MapAddress field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetMapAddress() string {
	if o == nil || isNil(o.MapAddress) {
		var ret string
		return ret
	}
	return *o.MapAddress
}

// GetMapAddressOk returns a tuple with the MapAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetMapAddressOk() (*string, bool) {
	if o == nil || isNil(o.MapAddress) {
    return nil, false
	}
	return o.MapAddress, true
}

// HasMapAddress returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasMapAddress() bool {
	if o != nil && !isNil(o.MapAddress) {
		return true
	}

	return false
}

// SetMapAddress gets a reference to the given string and assigns it to the MapAddress field.
func (o *SmsRouterInfo) SetMapAddress(v string) {
	o.MapAddress = &v
}

// GetRouterIpv4 returns the RouterIpv4 field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetRouterIpv4() string {
	if o == nil || isNil(o.RouterIpv4) {
		var ret string
		return ret
	}
	return *o.RouterIpv4
}

// GetRouterIpv4Ok returns a tuple with the RouterIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetRouterIpv4Ok() (*string, bool) {
	if o == nil || isNil(o.RouterIpv4) {
    return nil, false
	}
	return o.RouterIpv4, true
}

// HasRouterIpv4 returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasRouterIpv4() bool {
	if o != nil && !isNil(o.RouterIpv4) {
		return true
	}

	return false
}

// SetRouterIpv4 gets a reference to the given string and assigns it to the RouterIpv4 field.
func (o *SmsRouterInfo) SetRouterIpv4(v string) {
	o.RouterIpv4 = &v
}

// GetRouterIpv6 returns the RouterIpv6 field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetRouterIpv6() Ipv6Addr {
	if o == nil || isNil(o.RouterIpv6) {
		var ret Ipv6Addr
		return ret
	}
	return *o.RouterIpv6
}

// GetRouterIpv6Ok returns a tuple with the RouterIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetRouterIpv6Ok() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.RouterIpv6) {
    return nil, false
	}
	return o.RouterIpv6, true
}

// HasRouterIpv6 returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasRouterIpv6() bool {
	if o != nil && !isNil(o.RouterIpv6) {
		return true
	}

	return false
}

// SetRouterIpv6 gets a reference to the given Ipv6Addr and assigns it to the RouterIpv6 field.
func (o *SmsRouterInfo) SetRouterIpv6(v Ipv6Addr) {
	o.RouterIpv6 = &v
}

// GetRouterFqdn returns the RouterFqdn field value if set, zero value otherwise.
func (o *SmsRouterInfo) GetRouterFqdn() string {
	if o == nil || isNil(o.RouterFqdn) {
		var ret string
		return ret
	}
	return *o.RouterFqdn
}

// GetRouterFqdnOk returns a tuple with the RouterFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRouterInfo) GetRouterFqdnOk() (*string, bool) {
	if o == nil || isNil(o.RouterFqdn) {
    return nil, false
	}
	return o.RouterFqdn, true
}

// HasRouterFqdn returns a boolean if a field has been set.
func (o *SmsRouterInfo) HasRouterFqdn() bool {
	if o != nil && !isNil(o.RouterFqdn) {
		return true
	}

	return false
}

// SetRouterFqdn gets a reference to the given string and assigns it to the RouterFqdn field.
func (o *SmsRouterInfo) SetRouterFqdn(v string) {
	o.RouterFqdn = &v
}

func (o SmsRouterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NfInstanceId) {
		toSerialize["nfInstanceId"] = o.NfInstanceId
	}
	if !isNil(o.DiameterAddress) {
		toSerialize["diameterAddress"] = o.DiameterAddress
	}
	if !isNil(o.MapAddress) {
		toSerialize["mapAddress"] = o.MapAddress
	}
	if !isNil(o.RouterIpv4) {
		toSerialize["routerIpv4"] = o.RouterIpv4
	}
	if !isNil(o.RouterIpv6) {
		toSerialize["routerIpv6"] = o.RouterIpv6
	}
	if !isNil(o.RouterFqdn) {
		toSerialize["routerFqdn"] = o.RouterFqdn
	}
	return json.Marshal(toSerialize)
}

type NullableSmsRouterInfo struct {
	value *SmsRouterInfo
	isSet bool
}

func (v NullableSmsRouterInfo) Get() *SmsRouterInfo {
	return v.value
}

func (v *NullableSmsRouterInfo) Set(val *SmsRouterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsRouterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsRouterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsRouterInfo(val *SmsRouterInfo) *NullableSmsRouterInfo {
	return &NullableSmsRouterInfo{value: val, isSet: true}
}

func (v NullableSmsRouterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsRouterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


