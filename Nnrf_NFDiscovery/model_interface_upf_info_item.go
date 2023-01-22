/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// InterfaceUpfInfoItem Information of a given IP interface of an UPF
type InterfaceUpfInfoItem struct {
	InterfaceType UPInterfaceType `json:"interfaceType"`
	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`
	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`
	// Fully Qualified Domain Name
	EndpointFqdn *string `json:"endpointFqdn,omitempty"`
	NetworkInstance *string `json:"networkInstance,omitempty"`
}

// NewInterfaceUpfInfoItem instantiates a new InterfaceUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceUpfInfoItem(interfaceType UPInterfaceType) *InterfaceUpfInfoItem {
	this := InterfaceUpfInfoItem{}
	this.InterfaceType = interfaceType
	return &this
}

// NewInterfaceUpfInfoItemWithDefaults instantiates a new InterfaceUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceUpfInfoItemWithDefaults() *InterfaceUpfInfoItem {
	this := InterfaceUpfInfoItem{}
	return &this
}

// GetInterfaceType returns the InterfaceType field value
func (o *InterfaceUpfInfoItem) GetInterfaceType() UPInterfaceType {
	if o == nil {
		var ret UPInterfaceType
		return ret
	}

	return o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetInterfaceTypeOk() (*UPInterfaceType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InterfaceType, true
}

// SetInterfaceType sets field value
func (o *InterfaceUpfInfoItem) SetInterfaceType(v UPInterfaceType) {
	o.InterfaceType = v
}

// GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddresses() []string {
	if o == nil || isNil(o.Ipv4EndpointAddresses) {
		var ret []string
		return ret
	}
	return o.Ipv4EndpointAddresses
}

// GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Ipv4EndpointAddresses) {
    return nil, false
	}
	return o.Ipv4EndpointAddresses, true
}

// HasIpv4EndpointAddresses returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasIpv4EndpointAddresses() bool {
	if o != nil && !isNil(o.Ipv4EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv4EndpointAddresses gets a reference to the given []string and assigns it to the Ipv4EndpointAddresses field.
func (o *InterfaceUpfInfoItem) SetIpv4EndpointAddresses(v []string) {
	o.Ipv4EndpointAddresses = v
}

// GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddresses() []Ipv6Addr {
	if o == nil || isNil(o.Ipv6EndpointAddresses) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6EndpointAddresses
}

// GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6EndpointAddresses) {
    return nil, false
	}
	return o.Ipv6EndpointAddresses, true
}

// HasIpv6EndpointAddresses returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasIpv6EndpointAddresses() bool {
	if o != nil && !isNil(o.Ipv6EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv6EndpointAddresses gets a reference to the given []Ipv6Addr and assigns it to the Ipv6EndpointAddresses field.
func (o *InterfaceUpfInfoItem) SetIpv6EndpointAddresses(v []Ipv6Addr) {
	o.Ipv6EndpointAddresses = v
}

// GetEndpointFqdn returns the EndpointFqdn field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetEndpointFqdn() string {
	if o == nil || isNil(o.EndpointFqdn) {
		var ret string
		return ret
	}
	return *o.EndpointFqdn
}

// GetEndpointFqdnOk returns a tuple with the EndpointFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetEndpointFqdnOk() (*string, bool) {
	if o == nil || isNil(o.EndpointFqdn) {
    return nil, false
	}
	return o.EndpointFqdn, true
}

// HasEndpointFqdn returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasEndpointFqdn() bool {
	if o != nil && !isNil(o.EndpointFqdn) {
		return true
	}

	return false
}

// SetEndpointFqdn gets a reference to the given string and assigns it to the EndpointFqdn field.
func (o *InterfaceUpfInfoItem) SetEndpointFqdn(v string) {
	o.EndpointFqdn = &v
}

// GetNetworkInstance returns the NetworkInstance field value if set, zero value otherwise.
func (o *InterfaceUpfInfoItem) GetNetworkInstance() string {
	if o == nil || isNil(o.NetworkInstance) {
		var ret string
		return ret
	}
	return *o.NetworkInstance
}

// GetNetworkInstanceOk returns a tuple with the NetworkInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpfInfoItem) GetNetworkInstanceOk() (*string, bool) {
	if o == nil || isNil(o.NetworkInstance) {
    return nil, false
	}
	return o.NetworkInstance, true
}

// HasNetworkInstance returns a boolean if a field has been set.
func (o *InterfaceUpfInfoItem) HasNetworkInstance() bool {
	if o != nil && !isNil(o.NetworkInstance) {
		return true
	}

	return false
}

// SetNetworkInstance gets a reference to the given string and assigns it to the NetworkInstance field.
func (o *InterfaceUpfInfoItem) SetNetworkInstance(v string) {
	o.NetworkInstance = &v
}

func (o InterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if !isNil(o.Ipv4EndpointAddresses) {
		toSerialize["ipv4EndpointAddresses"] = o.Ipv4EndpointAddresses
	}
	if !isNil(o.Ipv6EndpointAddresses) {
		toSerialize["ipv6EndpointAddresses"] = o.Ipv6EndpointAddresses
	}
	if !isNil(o.EndpointFqdn) {
		toSerialize["endpointFqdn"] = o.EndpointFqdn
	}
	if !isNil(o.NetworkInstance) {
		toSerialize["networkInstance"] = o.NetworkInstance
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceUpfInfoItem struct {
	value *InterfaceUpfInfoItem
	isSet bool
}

func (v NullableInterfaceUpfInfoItem) Get() *InterfaceUpfInfoItem {
	return v.value
}

func (v *NullableInterfaceUpfInfoItem) Set(val *InterfaceUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceUpfInfoItem(val *InterfaceUpfInfoItem) *NullableInterfaceUpfInfoItem {
	return &NullableInterfaceUpfInfoItem{value: val, isSet: true}
}

func (v NullableInterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


