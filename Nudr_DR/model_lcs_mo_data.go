/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// LcsMoData struct for LcsMoData
type LcsMoData struct {
	AllowedServiceClasses []LcsMoServiceClass `json:"allowedServiceClasses"`
	MoAssistanceDataTypes *LcsBroadcastAssistanceTypesData `json:"moAssistanceDataTypes,omitempty"`
}

// NewLcsMoData instantiates a new LcsMoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsMoData(allowedServiceClasses []LcsMoServiceClass) *LcsMoData {
	this := LcsMoData{}
	this.AllowedServiceClasses = allowedServiceClasses
	return &this
}

// NewLcsMoDataWithDefaults instantiates a new LcsMoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsMoDataWithDefaults() *LcsMoData {
	this := LcsMoData{}
	return &this
}

// GetAllowedServiceClasses returns the AllowedServiceClasses field value
func (o *LcsMoData) GetAllowedServiceClasses() []LcsMoServiceClass {
	if o == nil {
		var ret []LcsMoServiceClass
		return ret
	}

	return o.AllowedServiceClasses
}

// GetAllowedServiceClassesOk returns a tuple with the AllowedServiceClasses field value
// and a boolean to check if the value has been set.
func (o *LcsMoData) GetAllowedServiceClassesOk() ([]LcsMoServiceClass, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedServiceClasses, true
}

// SetAllowedServiceClasses sets field value
func (o *LcsMoData) SetAllowedServiceClasses(v []LcsMoServiceClass) {
	o.AllowedServiceClasses = v
}

// GetMoAssistanceDataTypes returns the MoAssistanceDataTypes field value if set, zero value otherwise.
func (o *LcsMoData) GetMoAssistanceDataTypes() LcsBroadcastAssistanceTypesData {
	if o == nil || isNil(o.MoAssistanceDataTypes) {
		var ret LcsBroadcastAssistanceTypesData
		return ret
	}
	return *o.MoAssistanceDataTypes
}

// GetMoAssistanceDataTypesOk returns a tuple with the MoAssistanceDataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsMoData) GetMoAssistanceDataTypesOk() (*LcsBroadcastAssistanceTypesData, bool) {
	if o == nil || isNil(o.MoAssistanceDataTypes) {
    return nil, false
	}
	return o.MoAssistanceDataTypes, true
}

// HasMoAssistanceDataTypes returns a boolean if a field has been set.
func (o *LcsMoData) HasMoAssistanceDataTypes() bool {
	if o != nil && !isNil(o.MoAssistanceDataTypes) {
		return true
	}

	return false
}

// SetMoAssistanceDataTypes gets a reference to the given LcsBroadcastAssistanceTypesData and assigns it to the MoAssistanceDataTypes field.
func (o *LcsMoData) SetMoAssistanceDataTypes(v LcsBroadcastAssistanceTypesData) {
	o.MoAssistanceDataTypes = &v
}

func (o LcsMoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedServiceClasses"] = o.AllowedServiceClasses
	}
	if !isNil(o.MoAssistanceDataTypes) {
		toSerialize["moAssistanceDataTypes"] = o.MoAssistanceDataTypes
	}
	return json.Marshal(toSerialize)
}

type NullableLcsMoData struct {
	value *LcsMoData
	isSet bool
}

func (v NullableLcsMoData) Get() *LcsMoData {
	return v.value
}

func (v *NullableLcsMoData) Set(val *LcsMoData) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsMoData) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsMoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsMoData(val *LcsMoData) *NullableLcsMoData {
	return &NullableLcsMoData{value: val, isSet: true}
}

func (v NullableLcsMoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsMoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


