/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package JOSEProtectedMessageForwarding

import (
	"encoding/json"
	"fmt"
)

// HttpMethod Enumeration of HTTP methods
type HttpMethod struct {
	HttpMethodAnyOf *HttpMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *HttpMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into HttpMethodAnyOf
	err = json.Unmarshal(data, &dst.HttpMethodAnyOf);
	if err == nil {
		jsonHttpMethodAnyOf, _ := json.Marshal(dst.HttpMethodAnyOf)
		if string(jsonHttpMethodAnyOf) == "{}" { // empty struct
			dst.HttpMethodAnyOf = nil
		} else {
			return nil // data stored in dst.HttpMethodAnyOf, return on the first match
		}
	} else {
		dst.HttpMethodAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(HttpMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *HttpMethod) MarshalJSON() ([]byte, error) {
	if src.HttpMethodAnyOf != nil {
		return json.Marshal(&src.HttpMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableHttpMethod struct {
	value *HttpMethod
	isSet bool
}

func (v NullableHttpMethod) Get() *HttpMethod {
	return v.value
}

func (v *NullableHttpMethod) Set(val *HttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpMethod(val *HttpMethod) *NullableHttpMethod {
	return &NullableHttpMethod{value: val, isSet: true}
}

func (v NullableHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


