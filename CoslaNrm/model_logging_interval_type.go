/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// LoggingIntervalType See details in 3GPP TS 32.422 clause 5.10.8.
type LoggingIntervalType struct {
	UMTS []string `json:"UMTS,omitempty"`
	LTE []string `json:"LTE,omitempty"`
	NR []string `json:"NR,omitempty"`
}

// NewLoggingIntervalType instantiates a new LoggingIntervalType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingIntervalType() *LoggingIntervalType {
	this := LoggingIntervalType{}
	return &this
}

// NewLoggingIntervalTypeWithDefaults instantiates a new LoggingIntervalType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingIntervalTypeWithDefaults() *LoggingIntervalType {
	this := LoggingIntervalType{}
	return &this
}

// GetUMTS returns the UMTS field value if set, zero value otherwise.
func (o *LoggingIntervalType) GetUMTS() []string {
	if o == nil || isNil(o.UMTS) {
		var ret []string
		return ret
	}
	return o.UMTS
}

// GetUMTSOk returns a tuple with the UMTS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingIntervalType) GetUMTSOk() ([]string, bool) {
	if o == nil || isNil(o.UMTS) {
    return nil, false
	}
	return o.UMTS, true
}

// HasUMTS returns a boolean if a field has been set.
func (o *LoggingIntervalType) HasUMTS() bool {
	if o != nil && !isNil(o.UMTS) {
		return true
	}

	return false
}

// SetUMTS gets a reference to the given []string and assigns it to the UMTS field.
func (o *LoggingIntervalType) SetUMTS(v []string) {
	o.UMTS = v
}

// GetLTE returns the LTE field value if set, zero value otherwise.
func (o *LoggingIntervalType) GetLTE() []string {
	if o == nil || isNil(o.LTE) {
		var ret []string
		return ret
	}
	return o.LTE
}

// GetLTEOk returns a tuple with the LTE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingIntervalType) GetLTEOk() ([]string, bool) {
	if o == nil || isNil(o.LTE) {
    return nil, false
	}
	return o.LTE, true
}

// HasLTE returns a boolean if a field has been set.
func (o *LoggingIntervalType) HasLTE() bool {
	if o != nil && !isNil(o.LTE) {
		return true
	}

	return false
}

// SetLTE gets a reference to the given []string and assigns it to the LTE field.
func (o *LoggingIntervalType) SetLTE(v []string) {
	o.LTE = v
}

// GetNR returns the NR field value if set, zero value otherwise.
func (o *LoggingIntervalType) GetNR() []string {
	if o == nil || isNil(o.NR) {
		var ret []string
		return ret
	}
	return o.NR
}

// GetNROk returns a tuple with the NR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingIntervalType) GetNROk() ([]string, bool) {
	if o == nil || isNil(o.NR) {
    return nil, false
	}
	return o.NR, true
}

// HasNR returns a boolean if a field has been set.
func (o *LoggingIntervalType) HasNR() bool {
	if o != nil && !isNil(o.NR) {
		return true
	}

	return false
}

// SetNR gets a reference to the given []string and assigns it to the NR field.
func (o *LoggingIntervalType) SetNR(v []string) {
	o.NR = v
}

func (o LoggingIntervalType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UMTS) {
		toSerialize["UMTS"] = o.UMTS
	}
	if !isNil(o.LTE) {
		toSerialize["LTE"] = o.LTE
	}
	if !isNil(o.NR) {
		toSerialize["NR"] = o.NR
	}
	return json.Marshal(toSerialize)
}

type NullableLoggingIntervalType struct {
	value *LoggingIntervalType
	isSet bool
}

func (v NullableLoggingIntervalType) Get() *LoggingIntervalType {
	return v.value
}

func (v *NullableLoggingIntervalType) Set(val *LoggingIntervalType) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingIntervalType) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingIntervalType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingIntervalType(val *LoggingIntervalType) *NullableLoggingIntervalType {
	return &NullableLoggingIntervalType{value: val, isSet: true}
}

func (v NullableLoggingIntervalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingIntervalType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


