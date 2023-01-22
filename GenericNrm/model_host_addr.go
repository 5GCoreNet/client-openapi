/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package GenericNrm

import (
	"encoding/json"
	"fmt"
)

// HostAddr - struct for HostAddr
type HostAddr struct {
	Ipv6Addr *Ipv6Addr
	String *string
}

// Ipv6AddrAsHostAddr is a convenience function that returns Ipv6Addr wrapped in HostAddr
func Ipv6AddrAsHostAddr(v *Ipv6Addr) HostAddr {
	return HostAddr{
		Ipv6Addr: v,
	}
}

// stringAsHostAddr is a convenience function that returns string wrapped in HostAddr
func StringAsHostAddr(v *string) HostAddr {
	return HostAddr{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HostAddr) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(HostAddr)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HostAddr)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HostAddr) MarshalJSON() ([]byte, error) {
	if src.Ipv6Addr != nil {
		return json.Marshal(&src.Ipv6Addr)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HostAddr) GetActualInstance() (interface{}) {
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

type NullableHostAddr struct {
	value *HostAddr
	isSet bool
}

func (v NullableHostAddr) Get() *HostAddr {
	return v.value
}

func (v *NullableHostAddr) Set(val *HostAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableHostAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableHostAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostAddr(val *HostAddr) *NullableHostAddr {
	return &NullableHostAddr{value: val, isSet: true}
}

func (v NullableHostAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


