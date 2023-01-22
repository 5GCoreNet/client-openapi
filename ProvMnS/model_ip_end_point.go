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

// IpEndPoint struct for IpEndPoint
type IpEndPoint struct {
	Ipv4Address *string `json:"ipv4Address,omitempty"`
	Ipv6Address *Ipv6Addr `json:"ipv6Address,omitempty"`
	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty"`
	Transport *TransportProtocol `json:"transport,omitempty"`
	Port *int32 `json:"port,omitempty"`
}

// NewIpEndPoint instantiates a new IpEndPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpEndPoint() *IpEndPoint {
	this := IpEndPoint{}
	return &this
}

// NewIpEndPointWithDefaults instantiates a new IpEndPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpEndPointWithDefaults() *IpEndPoint {
	this := IpEndPoint{}
	return &this
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *IpEndPoint) GetIpv4Address() string {
	if o == nil || isNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEndPoint) GetIpv4AddressOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Address) {
    return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *IpEndPoint) HasIpv4Address() bool {
	if o != nil && !isNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *IpEndPoint) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *IpEndPoint) GetIpv6Address() Ipv6Addr {
	if o == nil || isNil(o.Ipv6Address) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEndPoint) GetIpv6AddressOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6Address) {
    return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *IpEndPoint) HasIpv6Address() bool {
	if o != nil && !isNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given Ipv6Addr and assigns it to the Ipv6Address field.
func (o *IpEndPoint) SetIpv6Address(v Ipv6Addr) {
	o.Ipv6Address = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *IpEndPoint) GetIpv6Prefix() Ipv6Prefix {
	if o == nil || isNil(o.Ipv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEndPoint) GetIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.Ipv6Prefix) {
    return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *IpEndPoint) HasIpv6Prefix() bool {
	if o != nil && !isNil(o.Ipv6Prefix) {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the Ipv6Prefix field.
func (o *IpEndPoint) SetIpv6Prefix(v Ipv6Prefix) {
	o.Ipv6Prefix = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *IpEndPoint) GetTransport() TransportProtocol {
	if o == nil || isNil(o.Transport) {
		var ret TransportProtocol
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEndPoint) GetTransportOk() (*TransportProtocol, bool) {
	if o == nil || isNil(o.Transport) {
    return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *IpEndPoint) HasTransport() bool {
	if o != nil && !isNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given TransportProtocol and assigns it to the Transport field.
func (o *IpEndPoint) SetTransport(v TransportProtocol) {
	o.Transport = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *IpEndPoint) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEndPoint) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *IpEndPoint) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *IpEndPoint) SetPort(v int32) {
	o.Port = &v
}

func (o IpEndPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4Address) {
		toSerialize["ipv4Address"] = o.Ipv4Address
	}
	if !isNil(o.Ipv6Address) {
		toSerialize["ipv6Address"] = o.Ipv6Address
	}
	if !isNil(o.Ipv6Prefix) {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix
	}
	if !isNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableIpEndPoint struct {
	value *IpEndPoint
	isSet bool
}

func (v NullableIpEndPoint) Get() *IpEndPoint {
	return v.value
}

func (v *NullableIpEndPoint) Set(val *IpEndPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIpEndPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIpEndPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpEndPoint(val *IpEndPoint) *NullableIpEndPoint {
	return &NullableIpEndPoint{value: val, isSet: true}
}

func (v NullableIpEndPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpEndPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


