/*
Common Type Definitions

OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ComDefs

import (
	"encoding/json"
	"fmt"
)

// IpAddr - struct for IpAddr
type IpAddr struct {
	Ipv6Addr *Ipv6Addr
	String *string
}

// Ipv6AddrAsIpAddr is a convenience function that returns Ipv6Addr wrapped in IpAddr
func Ipv6AddrAsIpAddr(v *Ipv6Addr) IpAddr {
	return IpAddr{
		Ipv6Addr: v,
	}
}

// stringAsIpAddr is a convenience function that returns string wrapped in IpAddr
func StringAsIpAddr(v *string) IpAddr {
	return IpAddr{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpAddr) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ipv6Addr
	err = newStrictDecoder(data).Decode(&dst.Ipv6Addr)
	if err == nil {
		jsonIpv6Addr, _ := json.Marshal(dst.Ipv6Addr)
		if string(jsonIpv6Addr) == "{}" { // empty struct
			dst.Ipv6Addr = nil
		} else {
			match++
		}
	} else {
		dst.Ipv6Addr = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ipv6Addr = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IpAddr)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpAddr)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpAddr) MarshalJSON() ([]byte, error) {
	if src.Ipv6Addr != nil {
		return json.Marshal(&src.Ipv6Addr)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpAddr) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Ipv6Addr != nil {
		return obj.Ipv6Addr
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIpAddr struct {
	value *IpAddr
	isSet bool
}

func (v NullableIpAddr) Get() *IpAddr {
	return v.value
}

func (v *NullableIpAddr) Set(val *IpAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddr(val *IpAddr) *NullableIpAddr {
	return &NullableIpAddr{value: val, isSet: true}
}

func (v NullableIpAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


