/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// TceMappingInfoTceIPAddress - struct for TceMappingInfoTceIPAddress
type TceMappingInfoTceIPAddress struct {
	Ipv6Addr *Ipv6Addr
	String *string
}

// Ipv6AddrAsTceMappingInfoTceIPAddress is a convenience function that returns Ipv6Addr wrapped in TceMappingInfoTceIPAddress
func Ipv6AddrAsTceMappingInfoTceIPAddress(v *Ipv6Addr) TceMappingInfoTceIPAddress {
	return TceMappingInfoTceIPAddress{
		Ipv6Addr: v,
	}
}

// stringAsTceMappingInfoTceIPAddress is a convenience function that returns string wrapped in TceMappingInfoTceIPAddress
func StringAsTceMappingInfoTceIPAddress(v *string) TceMappingInfoTceIPAddress {
	return TceMappingInfoTceIPAddress{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TceMappingInfoTceIPAddress) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(TceMappingInfoTceIPAddress)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TceMappingInfoTceIPAddress)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TceMappingInfoTceIPAddress) MarshalJSON() ([]byte, error) {
	if src.Ipv6Addr != nil {
		return json.Marshal(&src.Ipv6Addr)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TceMappingInfoTceIPAddress) GetActualInstance() (interface{}) {
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

type NullableTceMappingInfoTceIPAddress struct {
	value *TceMappingInfoTceIPAddress
	isSet bool
}

func (v NullableTceMappingInfoTceIPAddress) Get() *TceMappingInfoTceIPAddress {
	return v.value
}

func (v *NullableTceMappingInfoTceIPAddress) Set(val *TceMappingInfoTceIPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableTceMappingInfoTceIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableTceMappingInfoTceIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTceMappingInfoTceIPAddress(val *TceMappingInfoTceIPAddress) *NullableTceMappingInfoTceIPAddress {
	return &NullableTceMappingInfoTceIPAddress{value: val, isSet: true}
}

func (v NullableTceMappingInfoTceIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTceMappingInfoTceIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


