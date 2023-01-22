/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// RouteInformation At least one of the \"ipv4Addr\" attribute and the \"ipv6Addr\" attribute shall be included in the \"RouteInformation\" data type.  
type RouteInformation struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	Ipv6Addr *Ipv6Addr1 `json:"ipv6Addr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}

// NewRouteInformation instantiates a new RouteInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteInformation(portNumber int32) *RouteInformation {
	this := RouteInformation{}
	this.PortNumber = portNumber
	return &this
}

// NewRouteInformationWithDefaults instantiates a new RouteInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteInformationWithDefaults() *RouteInformation {
	this := RouteInformation{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *RouteInformation) GetIpv4Addr() string {
	if o == nil || isNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteInformation) GetIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Addr) {
    return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *RouteInformation) HasIpv4Addr() bool {
	if o != nil && !isNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *RouteInformation) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *RouteInformation) GetIpv6Addr() Ipv6Addr1 {
	if o == nil || isNil(o.Ipv6Addr) {
		var ret Ipv6Addr1
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteInformation) GetIpv6AddrOk() (*Ipv6Addr1, bool) {
	if o == nil || isNil(o.Ipv6Addr) {
    return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *RouteInformation) HasIpv6Addr() bool {
	if o != nil && !isNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given Ipv6Addr1 and assigns it to the Ipv6Addr field.
func (o *RouteInformation) SetIpv6Addr(v Ipv6Addr1) {
	o.Ipv6Addr = &v
}

// GetPortNumber returns the PortNumber field value
func (o *RouteInformation) GetPortNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value
// and a boolean to check if the value has been set.
func (o *RouteInformation) GetPortNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PortNumber, true
}

// SetPortNumber sets field value
func (o *RouteInformation) SetPortNumber(v int32) {
	o.PortNumber = v
}

func (o RouteInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !isNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if true {
		toSerialize["portNumber"] = o.PortNumber
	}
	return json.Marshal(toSerialize)
}

type NullableRouteInformation struct {
	value *RouteInformation
	isSet bool
}

func (v NullableRouteInformation) Get() *RouteInformation {
	return v.value
}

func (v *NullableRouteInformation) Set(val *RouteInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteInformation(val *RouteInformation) *NullableRouteInformation {
	return &NullableRouteInformation{value: val, isSet: true}
}

func (v NullableRouteInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


