/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// RouteToLocation At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type. 
type RouteToLocation struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RouteToLocation) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RouteToLocation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RouteToLocation) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRouteToLocation struct {
	value *RouteToLocation
	isSet bool
}

func (v NullableRouteToLocation) Get() *RouteToLocation {
	return v.value
}

func (v *NullableRouteToLocation) Set(val *RouteToLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteToLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteToLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteToLocation(val *RouteToLocation) *NullableRouteToLocation {
	return &NullableRouteToLocation{value: val, isSet: true}
}

func (v NullableRouteToLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteToLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


