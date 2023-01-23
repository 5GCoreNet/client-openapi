/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// ServiceNameCond Subscription to a set of NFs based on their support for a given Service Name
type ServiceNameCond struct {
	ServiceName ServiceName `json:"serviceName"`
}

// NewServiceNameCond instantiates a new ServiceNameCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNameCond(serviceName ServiceName) *ServiceNameCond {
	this := ServiceNameCond{}
	this.ServiceName = serviceName
	return &this
}

// NewServiceNameCondWithDefaults instantiates a new ServiceNameCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNameCondWithDefaults() *ServiceNameCond {
	this := ServiceNameCond{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *ServiceNameCond) GetServiceName() ServiceName {
	if o == nil {
		var ret ServiceName
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ServiceNameCond) GetServiceNameOk() (*ServiceName, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *ServiceNameCond) SetServiceName(v ServiceName) {
	o.ServiceName = v
}

func (o ServiceNameCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceName"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}

type NullableServiceNameCond struct {
	value *ServiceNameCond
	isSet bool
}

func (v NullableServiceNameCond) Get() *ServiceNameCond {
	return v.value
}

func (v *NullableServiceNameCond) Set(val *ServiceNameCond) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNameCond) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNameCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNameCond(val *ServiceNameCond) *NullableServiceNameCond {
	return &NullableServiceNameCond{value: val, isSet: true}
}

func (v NullableServiceNameCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNameCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


