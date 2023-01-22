/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// IpInterface struct for IpInterface
type IpInterface struct {
	Ipv4EndpointAddresses *string `json:"ipv4EndpointAddresses,omitempty"`
	Ipv6EndpointAddresses *Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
}

// NewIpInterface instantiates a new IpInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpInterface() *IpInterface {
	this := IpInterface{}
	return &this
}

// NewIpInterfaceWithDefaults instantiates a new IpInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpInterfaceWithDefaults() *IpInterface {
	this := IpInterface{}
	return &this
}

// GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field value if set, zero value otherwise.
func (o *IpInterface) GetIpv4EndpointAddresses() string {
	if o == nil || isNil(o.Ipv4EndpointAddresses) {
		var ret string
		return ret
	}
	return *o.Ipv4EndpointAddresses
}

// GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpInterface) GetIpv4EndpointAddressesOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4EndpointAddresses) {
    return nil, false
	}
	return o.Ipv4EndpointAddresses, true
}

// HasIpv4EndpointAddresses returns a boolean if a field has been set.
func (o *IpInterface) HasIpv4EndpointAddresses() bool {
	if o != nil && !isNil(o.Ipv4EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv4EndpointAddresses gets a reference to the given string and assigns it to the Ipv4EndpointAddresses field.
func (o *IpInterface) SetIpv4EndpointAddresses(v string) {
	o.Ipv4EndpointAddresses = &v
}

// GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field value if set, zero value otherwise.
func (o *IpInterface) GetIpv6EndpointAddresses() Ipv6Addr {
	if o == nil || isNil(o.Ipv6EndpointAddresses) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6EndpointAddresses
}

// GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpInterface) GetIpv6EndpointAddressesOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6EndpointAddresses) {
    return nil, false
	}
	return o.Ipv6EndpointAddresses, true
}

// HasIpv6EndpointAddresses returns a boolean if a field has been set.
func (o *IpInterface) HasIpv6EndpointAddresses() bool {
	if o != nil && !isNil(o.Ipv6EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv6EndpointAddresses gets a reference to the given Ipv6Addr and assigns it to the Ipv6EndpointAddresses field.
func (o *IpInterface) SetIpv6EndpointAddresses(v Ipv6Addr) {
	o.Ipv6EndpointAddresses = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *IpInterface) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpInterface) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *IpInterface) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *IpInterface) SetFqdn(v string) {
	o.Fqdn = &v
}

func (o IpInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4EndpointAddresses) {
		toSerialize["ipv4EndpointAddresses"] = o.Ipv4EndpointAddresses
	}
	if !isNil(o.Ipv6EndpointAddresses) {
		toSerialize["ipv6EndpointAddresses"] = o.Ipv6EndpointAddresses
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	return json.Marshal(toSerialize)
}

type NullableIpInterface struct {
	value *IpInterface
	isSet bool
}

func (v NullableIpInterface) Get() *IpInterface {
	return v.value
}

func (v *NullableIpInterface) Set(val *IpInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableIpInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableIpInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpInterface(val *IpInterface) *NullableIpInterface {
	return &NullableIpInterface{value: val, isSet: true}
}

func (v NullableIpInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


